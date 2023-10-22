package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

func encode(encCtx *ffmpeg.AVCodecContext, frame *ffmpeg.AVFrame, pkt *ffmpeg.AVPacket, outfile *os.File) {
	if frame != nil {
		fmt.Fprintf(os.Stdout, "Send frame %3d\n", frame.GetPts())
	}

	ret := ffmpeg.AvCodecSendFrame(encCtx, frame)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Error sending a frame for encoding\n")
		os.Exit(1)
	}

	for ret >= 0 {
		ret = ffmpeg.AvCodecReceivePacket(encCtx, pkt)
		if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
			return
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error during encoding\n")
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Write packet %3d (size=%5d)\n", pkt.GetPts(), pkt.GetSize())
		outfile.Write(unsafe.Slice(pkt.GetData(), pkt.GetSize()))
		ffmpeg.AvPacketUnref(pkt)
	}
}

func main() {
	endcode := []uint8{0x00, 0x00, 0x01, 0xb7}
	if len(os.Args) <= 2 {
		fmt.Fprintf(os.Stdout, "Usage: %s <output file> <codec name>\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]
	codecName := os.Args[2]

	// find the mpeg1video encoder
	codec := ffmpeg.AvCodecFindEncoderByName(codecName)
	if codec == nil {
		fmt.Fprintf(os.Stderr, "Codec '%s' not found\n", codecName)
		os.Exit(1)
	}

	avctx := ffmpeg.AvCodecAllocContext3(codec)
	if avctx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate video codec context\n")
		os.Exit(1)
	}

	pkt := ffmpeg.AvPacketAlloc()
	if pkt == nil {
		os.Exit(1)
	}

	// put sample parameters
	avctx.SetBitRate(400000)
	// resolution must be a multiple of two
	avctx.SetWidth(352)
	avctx.SetHeight(288)
	// frames per second
	avctx.SetTimeBase(ffmpeg.AvMakeQ(1, 25))
	avctx.SetFramerate(ffmpeg.AvMakeQ(25, 1))

	// emit one intra frame every ten frames check frame pict_type before passing frame
	// to encoder, if frame->pict_type is AV_PICTURE_TYPE_I
	// then gop_size is ignored and the output of encoder
	// will always be I frame irrespective to gop_size
	avctx.SetGopSize(10)
	avctx.SetMaxBFrames(1)
	avctx.SetPixFmt(ffmpeg.AV_PIX_FMT_YUV420P)

	if codec.GetID() == ffmpeg.AV_CODEC_ID_H264 {
		ffmpeg.AvOptSet(avctx.GetPrivData(), "preset", "slow", 0)
	}

	// open it
	ret := ffmpeg.AvCodecOpen2(avctx, codec, nil)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open codec %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(1)
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", filename)
		os.Exit(1)
	}

	frame := ffmpeg.AvFrameAlloc()
	if frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate video frame\n")
		os.Exit(1)
	}
	frame.SetFormat(avctx.GetPixFmt())
	frame.SetWidth(avctx.GetWidth())
	frame.SetHeight(avctx.GetHeight())

	ret = ffmpeg.AvFrameGetBuffer(frame, 0)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate the video frame data\n")
		os.Exit(1)
	}

	// encode 1 second of video
	for i := 0; i < 25; i++ {
		// make sure the frame data is writable
		ret = ffmpeg.AvFrameMakeWritable(frame)
		if ret < 0 {
			os.Exit(1)
		}

		// prepare a dummy image
		data0 := unsafe.Slice(frame.GetData()[0], avctx.GetHeight()*frame.GetLinesize()[0])
		data1 := unsafe.Slice(frame.GetData()[1], (avctx.GetHeight()/2)*frame.GetLinesize()[1])
		data2 := unsafe.Slice(frame.GetData()[2], (avctx.GetHeight()/2)*frame.GetLinesize()[2])
		// Y
		for y := 0; y < int(avctx.GetHeight()); y++ {
			for x := 0; x < int(avctx.GetWidth()); x++ {
				data0[y*int(frame.GetLinesize()[0])+x] = uint8(x + y + i*3)
			}
		}
		// Cb and Cr
		for y := 0; y < int(avctx.GetHeight()/2); y++ {
			for x := 0; x < int(avctx.GetWidth()/2); x++ {
				data1[y*int(frame.GetLinesize()[1])+x] = uint8(128 + y + i*2)
				data2[y*int(frame.GetLinesize()[2])+x] = uint8(64 + x + i*5)
			}
		}

		frame.SetPts(int64(i))

		// encode the image
		encode(avctx, frame, pkt, f)
	}

	// flush the encoder
	encode(avctx, nil, pkt, f)

	// add sequence end code to have a real MPEG file
	if codec.GetID() == ffmpeg.AV_CODEC_ID_MPEG1VIDEO || codec.GetID() == ffmpeg.AV_CODEC_ID_MPEG2VIDEO {
		f.Write(endcode)
	}
	f.Close()

	ffmpeg.AvCodecFreeContext(&avctx)
	ffmpeg.AvFrameFree(&frame)
	ffmpeg.AvPacketFree(&pkt)
}
