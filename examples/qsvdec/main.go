//go:build ffmpeg_hw_qsv

package main

/*
#include <stdlib.h>

struct AVCodecContext;

int get_format(struct AVCodecContext *avctx, int *pix_fmts) {
	// TODO.
	return -1;
}

*/
import "C"

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

type decodeContext struct {
	hwDeviceRef *ffmpeg.AVBufferRef
}

func decodePacket(decode *decodeContext, decoderCtx *ffmpeg.AVCodecContext,
	frame *ffmpeg.AVFrame, swFrame *ffmpeg.AVFrame,
	pkt *ffmpeg.AVPacket, outputCtx *ffmpeg.AVIOContext) (ret int32) {
	var (
		data [][]uint8
	)
	if ret = ffmpeg.AvCodecSendPacket(decoderCtx, pkt); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error during decoding\n")
		return ret
	}

	for ret >= 0 {
		ret = ffmpeg.AvCodecReceiveFrame(decoderCtx, frame)
		if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
			break
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while decoding.\n")
			return ret
		}

		// A real program would do something useful with the decoded frame here.
		// We just retrieve the raw data and write it to a file, which is rather
		// useless but pedagogic.
		if ret = ffmpeg.AvHWFrameTransferData(swFrame, frame, 0); ret < 0 {
			fmt.Fprintf(os.Stderr, "Error transferring the data to system memory\n")
			goto fail
		}

		data = ffmpeg.SliceSlice(&swFrame.GetData()[0], len(swFrame.GetData()),
			swFrame.GetWidth()*swFrame.GetHeight())
		for i := 0; i < len(swFrame.GetData()) && data[i] != nil; i++ {
			for j := 0; j < int(swFrame.GetHeight()>>ffmpeg.CondExpr(i > 0, 1, 0)); j++ {
				ffmpeg.AvIOWrite(outputCtx, &data[i][j*int(swFrame.GetLinesize()[i])], swFrame.GetWidth())
			}
		}
	fail:
		ffmpeg.AvFrameUnref(swFrame)
		ffmpeg.AvFrameUnref(frame)

		if ret < 0 {
			return ret
		}
	}

	return 0
}

func main() {
	var (
		intputCtx      *ffmpeg.AVFormatContext
		videoSt        *ffmpeg.AVStream
		decoderCtx     *ffmpeg.AVCodecContext
		decoder        *ffmpeg.AVCodec
		pkt            ffmpeg.AVPacket
		frame, swFrame *ffmpeg.AVFrame
		decode         = &decodeContext{}
		outputCtx      *ffmpeg.AVIOContext
		ret            int32
	)

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input file> <output file>\n", os.Args[0])
		os.Exit(1)
	}

	// open the input file
	if ret = ffmpeg.AvFormatOpenInput(&intputCtx, os.Args[1], nil, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Cannot open input file '%s': ", os.Args[1])
		goto finish
	}

	// find the first H.264 video stream
	for i := 0; i < int(intputCtx.GetNbStreams()); i++ {
		st := intputCtx.GetStreams()[i]
		if st.GetCodecpar().GetCodecId() == ffmpeg.AV_CODEC_ID_H264 && videoSt == nil {
			videoSt = st
		} else {
			st.SetDiscard(ffmpeg.AVDISCARD_ALL)
		}
	}
	if videoSt == nil {
		fmt.Fprintf(os.Stderr, "No H.264 video stream in the input file\n")
		goto finish
	}

	// open the hardware device
	if ret = ffmpeg.AvHWDeviceCtxCreate(&decode.hwDeviceRef, ffmpeg.AV_HWDEVICE_TYPE_QSV,
		"auto", nil, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Cannot open the hardware device\n")
		goto finish
	}

	// initialize the decoder
	if decoder = ffmpeg.AvCodecFindDecoderByName("h264_qsv"); decoder == nil {
		fmt.Fprintf(os.Stderr, "The QSV decoder is not present in libavcodec\n")
		goto finish
	}

	decoderCtx.SetCodecId(ffmpeg.AV_CODEC_ID_H264)
	if videoSt.GetCodecpar().GetExtradataSize() != 0 {
		decoderCtx.SetExtradata((*uint8)(ffmpeg.AvMallocz(videoSt.GetCodecpar().GetExtradataSize() +
			ffmpeg.AV_INPUT_BUFFER_PADDING_SIZE)))
		if decoderCtx.GetExtradata() == nil {
			ret = ffmpeg.AVERROR(syscall.ENOMEM)
			goto finish
		}
		copy(unsafe.Slice(decoderCtx.GetExtradata(), videoSt.GetCodecpar().GetExtradataSize()),
			unsafe.Slice(videoSt.GetCodecpar().GetExtradata(), videoSt.GetCodecpar().GetExtradataSize()))
	}

	decoderCtx.SetOpaque(decode.hwDeviceRef)
	decoderCtx.SetGetFormat((ffmpeg.AVCodecContextGetFormatFunc)(C.get_format))

	if ret = ffmpeg.AvCodecOpen2(decoderCtx, nil, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error opening the decoder: ")
		goto finish
	}

	// open the output stream
	if ret = ffmpeg.AvIOOpen(&outputCtx, os.Args[2], ffmpeg.AVIO_FLAG_WRITE); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error opening the output context: ")
		goto finish
	}

	frame = ffmpeg.AvFrameAlloc()
	swFrame = ffmpeg.AvFrameAlloc()
	if frame == nil || swFrame == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto finish
	}

	// actual decoding
	for ret >= 0 {
		if ret = ffmpeg.AvReadFrame(intputCtx, &pkt); ret < 0 {
			break
		}

		if pkt.GetStreamIndex() == videoSt.GetIndex() {
			ret = decodePacket(decode, decoderCtx, frame, swFrame, &pkt, outputCtx)
		}

		ffmpeg.AvPacketUnref(&pkt)
	}

finish:
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "%s\n", ffmpeg.AvErr2str(ret))
	}

	ffmpeg.AvFormatCloseInput(&intputCtx)

	ffmpeg.AvFrameFree(&frame)
	ffmpeg.AvFrameFree(&swFrame)

	ffmpeg.AvCodecFreeContext(&decoderCtx)

	ffmpeg.AvBufferUnref(&decode.hwDeviceRef)

	ffmpeg.AvIOClose(outputCtx)

	os.Exit(int(ret))
}
