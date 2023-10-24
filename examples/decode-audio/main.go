package main

import (
	"fmt"
	"os"
	"syscall"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

const (
	AUDIO_INBUF_SIZE    = 20480
	AUDIO_REFILL_THRESH = 4096
)

func getFormatFromSampleFmt(sampleFmt ffmpeg.AVSampleFormat) (string, int32) {
	sampleFmtEntry := []struct {
		sampleFmt ffmpeg.AVSampleFormat
		fmtBe     string
		fmtLe     string
	}{
		{ffmpeg.AV_SAMPLE_FMT_U8, "u8", "u8"},
		{ffmpeg.AV_SAMPLE_FMT_S16, "s16be", "s16le"},
		{ffmpeg.AV_SAMPLE_FMT_S32, "s32be", "s32le"},
		{ffmpeg.AV_SAMPLE_FMT_FLT, "f32be", "f32le"},
		{ffmpeg.AV_SAMPLE_FMT_DBL, "f64be", "f64le"},
	}

	for _, entry := range sampleFmtEntry {
		if sampleFmt == entry.sampleFmt {
			return ffmpeg.AV_NE(entry.fmtBe, entry.fmtLe), 0
		}
	}

	fmt.Fprintf(os.Stderr, "sample format %s is not supported as output format\n",
		ffmpeg.AvGetSampleFmtName(sampleFmt))
	return ffmpeg.NIL, -1
}

func decode(decCtx *ffmpeg.AVCodecContext, pkt *ffmpeg.AVPacket, frame *ffmpeg.AVFrame, outfile *os.File) {
	// send the packet with the compressed data to the decoder
	ret := ffmpeg.AvCodecSendPacket(decCtx, pkt)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Error submitting the packet to the decoder\n")
		os.Exit(1)
	}

	// read all the output frames (in general there may be any number of them
	for ret >= 0 {
		ret = ffmpeg.AvCodecReceiveFrame(decCtx, frame)
		if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
			return
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error during decoding\n")
			os.Exit(1)
		}
		dataSize := ffmpeg.AvGetBytesPerSample(decCtx.GetSampleFmt())
		if dataSize < 0 {
			// This should not occur, checking just for paranoia
			fmt.Fprintf(os.Stderr, "Failed to calculate data size\n")
			os.Exit(1)
		}
		data := ffmpeg.SliceSlice(&frame.GetData()[0], decCtx.GetChannels(), frame.GetNbSamples()*dataSize)
		for i := int32(0); i < frame.GetNbSamples(); i++ {
			for ch := 0; ch < int(decCtx.GetChannels()); ch++ {
				outfile.Write(data[ch][dataSize*i : dataSize*(i+1)])
			}
		}
	}
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Fprintf(os.Stdout, "Usage: %s <input file> <output file>\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]
	outfilename := os.Args[2]

	pkt := ffmpeg.AvPacketAlloc()

	codec := ffmpeg.AvCodecFindDecoder(ffmpeg.AV_CODEC_ID_MP2)
	if codec == nil {
		fmt.Fprintf(os.Stderr, "Codec not found\n")
		os.Exit(1)
	}

	parser := ffmpeg.AvParserInit(codec.GetID())
	if parser == nil {
		fmt.Fprintf(os.Stderr, "Parser not found\n")
		os.Exit(1)
	}

	avctx := ffmpeg.AvCodecAllocContext3(codec)
	if avctx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate audio codec context\n")
		os.Exit(1)
	}

	// open it
	if ffmpeg.AvCodecOpen2(avctx, codec, nil) < 0 {
		fmt.Fprintf(os.Stderr, "Could not open codec\n")
		os.Exit(1)
	}

	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", filename)
		os.Exit(1)
	}
	outfile, err := os.OpenFile(outfilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", outfilename)
		os.Exit(1)
	}

	// decode until eof
	var decodedFrame *ffmpeg.AVFrame
	inbuf := make([]byte, AUDIO_INBUF_SIZE+ffmpeg.AV_INPUT_BUFFER_PADDING_SIZE)
	dataOffset := 0
	dataSize, err := f.Read(inbuf[:AUDIO_INBUF_SIZE])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	for dataSize > 0 {
		if decodedFrame == nil {
			if decodedFrame = ffmpeg.AvFrameAlloc(); decodedFrame == nil {
				fmt.Fprintf(os.Stderr, "Could not allocate audio frame\n")
				os.Exit(1)
			}
		}

		ret := ffmpeg.AvParserParse2(parser, avctx, pkt.GetDataAddr(), pkt.GetSizeAddr(),
			&inbuf[dataOffset], int32(dataSize),
			ffmpeg.AV_NOPTS_VALUE, ffmpeg.AV_NOPTS_VALUE, 0)
		if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while parsing\n")
			os.Exit(1)
		}

		dataOffset += int(ret)
		dataSize -= int(ret)

		if pkt.GetSize() > 0 {
			decode(avctx, pkt, decodedFrame, outfile)
		}

		if dataSize < AUDIO_REFILL_THRESH {
			copy(inbuf, inbuf[dataOffset:dataOffset+dataSize])
			dataOffset = 0
			length, _ := f.Read(inbuf[dataOffset+dataSize : AUDIO_INBUF_SIZE-dataSize])
			if length > 0 {
				dataSize += int(length)
			}
		}
	}

	// flush the decoder
	pkt.SetData(nil)
	pkt.SetSize(0)
	decode(avctx, pkt, decodedFrame, outfile)

	// print output pcm infomations, because there have no metadata of pcm
	sfmt := avctx.GetSampleFmt()

	if ffmpeg.AvSampleFmtIsPlanar(sfmt) > 0 {
		packed := ffmpeg.AvGetSampleFmtName(sfmt)
		if len(packed) == 0 {
			packed = "?"
		}
		fmt.Fprintf(os.Stdout, "Warning: the sample format the decoder produced is planar (%s)."+
			" This example will output the first channel only.\n", packed)
		sfmt = ffmpeg.AvGetPackedSampleFmt(sfmt)
	}

	nChannels := avctx.GetChannels()
	fmtStr, ret := getFormatFromSampleFmt(sfmt)
	if ret < 0 {
		goto end
	}

	fmt.Fprintf(os.Stdout, "Play the output audio file with the command:\n"+
		"ffplay -f %s -ac %d -ar %d %s\n",
		fmtStr, nChannels, avctx.GetSampleRate(), outfilename)

end:
	f.Close()
	outfile.Close()

	ffmpeg.AvCodecFreeContext(&avctx)
	ffmpeg.AvParserClose(parser)
	ffmpeg.AvFrameFree(&decodedFrame)
	ffmpeg.AvPacketFree(&pkt)
}
