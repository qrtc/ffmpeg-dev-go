package main

import (
	"fmt"
	"io"
	"os"
	"syscall"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

const (
	INBUF_SIZE = 4096
)

func pgmSave(buf *uint8, wrap, xsize, ysize int32, filename string) {
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	fmt.Fprintf(f, "P5\n%d %d\n%d\n", xsize, ysize, 255)
	for i := int32(0); i < ysize; i++ {
		f.Write(ffmpeg.SliceWithOffset(buf, i+wrap, xsize))
	}
	f.Close()
}

func decode(decCtx *ffmpeg.AvCodecContext, frame *ffmpeg.AvFrame, pkt *ffmpeg.AvPacket, filename string) {
	ret := ffmpeg.AvCodecSendPacket(decCtx, pkt)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Error sending a packet for decoding\n")
		os.Exit(1)
	}

	for ret >= 0 {
		ret = ffmpeg.AvCodecReceiveFrame(decCtx, frame)
		if ret == ffmpeg.AVERROR(int32(syscall.EAGAIN)) || ret == ffmpeg.AVERROR_EOF {
			return
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error during decoding\n")
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "saving frame %3d\n", decCtx.GetFrameNumber())

		// the picture is allocated by the decoder. no need to free it
		fname := fmt.Sprintf("%s-%d", filename, decCtx.GetFrameNumber())
		pgmSave(frame.GetDataIdx(0), frame.GetLinesizeIdx(0), frame.GetWidth(), frame.GetHeight(), fname)
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
	if pkt == nil {
		os.Exit(1)
	}

	codec := ffmpeg.AvCodecFindDecoder(ffmpeg.AV_CODEC_ID_MPEG1VIDEO)
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
		fmt.Fprintf(os.Stderr, "Could not allocate video codec context\n")
		os.Exit(1)
	}

	// For some codecs, such as msmpeg4 and mpeg4, width and height MUST be initialized
	// there because this information is not available in the bitstream.

	// open it
	if ffmpeg.AvCodecOpen2(avctx, codec, nil) < 0 {
		fmt.Fprintf(os.Stderr, "Could not open codec\n")
		os.Exit(1)
	}

	inbuf := make([]byte, INBUF_SIZE+ffmpeg.AV_INPUT_BUFFER_PADDING_SIZE)

	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", filename)
		os.Exit(1)
	}

	frame := ffmpeg.AvFrameAlloc()
	if frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate video frame\n")
		os.Exit(1)
	}

	for {
		// read raw data from the input file
		dataSize, err := f.Read(inbuf[:INBUF_SIZE])
		if err == io.EOF || dataSize == 0 {
			break
		}

		data := inbuf
		// use the parser to split the data into frames
		for dataSize > 0 {
			ret := ffmpeg.AvParserParse2(parser, avctx, pkt.GetDataAddr(), pkt.GetSizeAddr(),
				&data[0], int32(dataSize), ffmpeg.AV_NOPTS_VALUE, ffmpeg.AV_NOPTS_VALUE, 0)
			if ret < 0 {
				fmt.Fprintf(os.Stderr, "Error while parsing\n")
				os.Exit(1)
			}
			data = data[ret:]
			dataSize -= int(ret)

			if pkt.GetSize() > 0 {
				decode(avctx, frame, pkt, outfilename)
			}
		}
	}

	// flush the decoder
	decode(avctx, frame, nil, outfilename)

	f.Close()

	ffmpeg.AvCodecFreeContext(&avctx)
	ffmpeg.AvParserClose(parser)
	ffmpeg.AvFrameFree(&frame)
	ffmpeg.AvPacketFree(&pkt)
}
