package main

import (
	"fmt"
	"math"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

// check that a given sample format is supported by the encoder.
func checkSampleFmt(codec *ffmpeg.AVCodec, sampleFmt ffmpeg.AVSampleFormat) int32 {
	for _, f := range codec.GetSampleFmts() {
		if f == sampleFmt {
			return 1
		}
	}
	return 0
}

func selectSampleRate(codec *ffmpeg.AVCodec) int32 {
	var bestSamplerate int32
	ss := codec.GetSupportedSamplerates()
	if len(ss) == 0 {
		return 44100
	}
	for _, s := range ss {
		if bestSamplerate == 0 || ffmpeg.FFABS(44100-s) < ffmpeg.FFABS(44100-bestSamplerate) {
			bestSamplerate = s
		}
	}
	return bestSamplerate
}

// select layout with the highest channel count.
func selectChannelLayout(codec *ffmpeg.AVCodec, dst *ffmpeg.AVChannelLayout) int32 {
	var (
		bestChLayout   *ffmpeg.AVChannelLayout
		bestNbChannels int32
	)

	chls := codec.GetChLayouts()
	if len(chls) == 0 {
		return ffmpeg.AvChannelLayoutCopy(dst, ffmpeg.AV_CHANNEL_LAYOUT_STEREO())
	}

	for i := range chls {
		nbChannels := chls[i].GetNbChannels()

		if nbChannels > bestNbChannels {
			bestChLayout = &chls[i]
			bestNbChannels = nbChannels
		}
	}

	return ffmpeg.AvChannelLayoutCopy(dst, bestChLayout)
}

func encode(ctx *ffmpeg.AVCodecContext, frame *ffmpeg.AVFrame, pkt *ffmpeg.AVPacket, output *os.File) {
	// send the frame for encoding
	ret := ffmpeg.AvCodecSendFrame(ctx, frame)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Error sending the frame to the encoder\n")
		os.Exit(1)
	}

	// read all the available output packets (in general there may be any number of them
	for ret >= 0 {
		ret = ffmpeg.AvCodecReceivePacket(ctx, pkt)
		if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
			return
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error encoding audio frame\n")
			os.Exit(1)
		}

		output.Write(unsafe.Slice(pkt.GetData(), pkt.GetSize()))
		ffmpeg.AvPacketUnref(pkt)
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stdout, "Usage: %s <output file>\n", os.Args[0])
		return
	}
	filename := os.Args[1]

	// find the MP2 encoder
	codec := ffmpeg.AvCodecFindEncoder(ffmpeg.AV_CODEC_ID_MP2)

	if codec == nil {
		fmt.Fprintf(os.Stderr, "Codec not found\n")
		os.Exit(1)
	}

	avctx := ffmpeg.AvCodecAllocContext3(codec)
	if avctx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate audio codec context\n")
		os.Exit(1)
	}

	// put sample parameters
	avctx.SetBitRate(64000)

	avctx.SetSampleFmt(ffmpeg.AV_SAMPLE_FMT_S16)
	if checkSampleFmt(codec, avctx.GetSampleFmt()) == 0 {
		fmt.Fprintf(os.Stderr, "Encoder does not support sample format %s",
			ffmpeg.AvGetSampleFmtName(avctx.GetSampleFmt()))
		os.Exit(1)
	}

	// select other audio parameters supported by the encoder
	avctx.SetSampleRate(selectSampleRate(codec))
	if selectChannelLayout(codec, avctx.GetChLayoutAddr()) < 0 {
		fmt.Fprintf(os.Stderr, "codec context select channel layout failed\n")
		os.Exit(1)
	}

	// open it
	if ffmpeg.AvCodecOpen2(avctx, codec, nil) < 0 {
		fmt.Fprintf(os.Stderr, "Could not open codec\n")
		os.Exit(1)
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", filename)
		os.Exit(1)
	}

	// packet for holding encoded output
	pkt := ffmpeg.AvPacketAlloc()
	if pkt == nil {
		fmt.Fprintf(os.Stderr, "could not allocate the packet\n")
		os.Exit(1)
	}

	frame := ffmpeg.AvFrameAlloc()
	if frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate audio frame\n")
		os.Exit(1)
	}

	frame.SetNbSamples(avctx.GetFrameSize())
	frame.SetFormat(avctx.GetSampleFmt())
	if ffmpeg.AvChannelLayoutCopy(frame.GetChLayoutAddr(), avctx.GetChLayoutAddr()) < 0 {
		fmt.Fprintf(os.Stderr, "frame copy channel layout failed\n")
		os.Exit(1)
	}

	// allocate the data buffers
	if ret := ffmpeg.AvFrameGetBuffer(frame, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate audio data buffers\n")
		os.Exit(1)
	}

	// encode a single tone sound
	t := float64(0)
	tincr := 2 * math.Pi * 440.0 / float64(avctx.GetSampleRate())
	for i := 0; i < 200; i++ {
		// make sure the frame is writable -- makes a copy if the encoder
		// kept a reference internally
		if ret := ffmpeg.AvFrameMakeWritable(frame); ret < 0 {
			os.Exit(1)
		}
		samples := unsafe.Slice((*uint16)(unsafe.Pointer(frame.GetData()[0])), avctx.GetFrameSize()*avctx.GetChLayoutAddr().GetNbChannels())

		for j := 0; j < int(avctx.GetFrameSize()); j++ {
			samples[2*j] = (uint16)(math.Sin(t) * 10000)

			for k := 1; k < int(avctx.GetChLayoutAddr().GetNbChannels()); k++ {
				samples[2*j+k] = samples[2*j]
			}
			t += tincr
		}
		encode(avctx, frame, pkt, f)
	}

	// flush the encoder
	encode(avctx, nil, pkt, f)

	f.Close()

	ffmpeg.AvFrameFree(&frame)
	ffmpeg.AvPacketFree(&pkt)
	ffmpeg.AvCodecFreeContext(&avctx)
}
