package main

/*
#include <stdlib.h>
#include <stdio.h>

static int hw_pix_fmt;

static inline int get_hw_pix_fmt()
{
	return hw_pix_fmt;
}

static inline void set_hw_pix_fmt(int fmt)
{
	hw_pix_fmt = fmt;
}

struct AVCodecContext;

int get_hw_format(struct AVCodecContext *ctx, const int *pix_fmts)
{
	const int *p;

	fprintf(stderr, "get_hw_format ballback has been invoked.\n");
	for (p = pix_fmts; *p != -1; p++) {
		if (*p == hw_pix_fmt)
			return *p;
	}

	fprintf(stderr, "Failed to get HW surface format.\n");
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

var (
// Note: Gobal variable as argument pass to C world may be cause panic.
// `runtime error: cgo argument has Go pointer to unpinned Go pointer`
)

func hwDecoderInit(ctx *ffmpeg.AVCodecContext, hwType ffmpeg.AVHWDeviceType) (hwDeviceCtx *ffmpeg.AVBufferRef, ret int32) {
	if ret := ffmpeg.AvHWDeviceCtxCreate(&hwDeviceCtx, hwType, ffmpeg.NIL, nil, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to create specified HW device.\n")
		return nil, ret
	}
	ctx.SetHwDeviceCtx(ffmpeg.AvBufferRef(hwDeviceCtx))
	return hwDeviceCtx, ret
}

func decodeWrite(avctx *ffmpeg.AVCodecContext, packet *ffmpeg.AVPacket, outputFile *os.File) (ret int32) {
	var frame, swFrame, tmpFrame *ffmpeg.AVFrame
	var buffer *uint8
	var size int32

	if ret = ffmpeg.AvCodecSendPacket(avctx, packet); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error during decoding\n")
		return ret
	}

	for {
		if frame = ffmpeg.AvFrameAlloc(); frame == nil {
			fmt.Fprintf(os.Stderr, "Can not alloc frame\n")
			goto fail
		}
		if swFrame = ffmpeg.AvFrameAlloc(); swFrame == nil {
			fmt.Fprintf(os.Stderr, "Can not alloc sw frame\n")
			goto fail
		}

		ret = ffmpeg.AvCodecReceiveFrame(avctx, frame)
		if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
			ffmpeg.AvFrameFree(&frame)
			ffmpeg.AvFrameFree(&swFrame)
			return 0
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while decoding\n")
			goto fail
		}

		if frame.GetFormat() == (int32)(C.get_hw_pix_fmt()) {
			// retrieve data from GPU to CPU
			if ret = ffmpeg.AvHWFrameTransferData(swFrame, frame, 0); ret < 0 {
				fmt.Fprintf(os.Stderr, "Error transferring the data to system memory\n")
				goto fail
			}
			tmpFrame = swFrame
		} else {
			tmpFrame = frame
		}

		size = ffmpeg.AvImageGetBufferSize(tmpFrame.GetFormat(), tmpFrame.GetWidth(),
			tmpFrame.GetHeight(), 1)
		buffer = (*uint8)(ffmpeg.AvMalloc(size))
		if buffer == nil {
			fmt.Fprintf(os.Stderr, "Can not alloc buffer\n")
			ret = ffmpeg.AVERROR(syscall.ENOMEM)
			goto fail
		}

		ret = ffmpeg.AvImageCopyToBuffer(buffer, size, tmpFrame.GetData(),
			tmpFrame.GetLinesize(), tmpFrame.GetFormat(),
			tmpFrame.GetWidth(), tmpFrame.GetHeight(), 1)
		if ret < 0 {
			fmt.Fprintf(os.Stderr, "Can not copy image to buffer\n")
			goto fail
		}
		if _, err := outputFile.Write(unsafe.Slice(buffer, size)); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to dump raw data.\n")
			goto fail
		}
	fail:
		ffmpeg.AvFrameFree(&frame)
		ffmpeg.AvFrameFree(&swFrame)
		ffmpeg.AvFreep(&buffer)
		if ret < 0 {
			return ret
		}
	}
}

func main() {
	var outputFile *os.File
	var hwDeviceCtx *ffmpeg.AVBufferRef
	var inputCtx *ffmpeg.AVFormatContext
	var videoStream, ret int32
	var video *ffmpeg.AVStream
	var decoderCtx *ffmpeg.AVCodecContext
	var decoder *ffmpeg.AVCodec
	var packet ffmpeg.AVPacket
	var hwType ffmpeg.AVHWDeviceType

	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <device type> <input file> <output file>\n", os.Args[0])
		os.Exit(1)
	}

	hwType = ffmpeg.AvHWDeviceFindTypeByName(os.Args[1])
	if hwType == ffmpeg.AV_HWDEVICE_TYPE_NONE {
		fmt.Fprintf(os.Stderr, "Device type %s is not supported.\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Available device types:")
		for hwType = ffmpeg.AvHWDeviceIterateTypes(hwType); hwType != ffmpeg.AV_HWDEVICE_TYPE_NONE; {
			fmt.Fprintf(os.Stderr, " %s", ffmpeg.AvHWDeviceGetTypeName(hwType))
			hwType = ffmpeg.AvHWDeviceIterateTypes(hwType)
		}
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(1)
	}

	// open the input file
	if ret = ffmpeg.AvFormatOpenInput(&inputCtx, os.Args[2], nil, nil); ret != 0 {
		fmt.Fprintf(os.Stderr, "Cannot open input file '%s'\n", os.Args[2])
		os.Exit(1)
	}

	if ret = ffmpeg.AvFormatFindStreamInfo(inputCtx, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Cannot find input stream information.\n")
		os.Exit(1)
	}

	// find the video stream information
	ret = ffmpeg.AvFindBestStream(inputCtx, ffmpeg.AVMEDIA_TYPE_VIDEO, -1, -1, &decoder, 0)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Cannot find a video stream in the input file\n")
		os.Exit(1)
	}
	videoStream = ret

	for i := 0; ; i++ {
		config := ffmpeg.AvCodecGetHwConfig(decoder, i)
		if config == nil {
			fmt.Fprintf(os.Stderr, "Decoder %s does not support device type %s.\n",
				decoder.GetName(), ffmpeg.AvHWDeviceGetTypeName(hwType))
			os.Exit(1)
		}
		if (config.GetMethods()&ffmpeg.AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX) != 0 &&
			config.GetDeviceType() == hwType {
			C.set_hw_pix_fmt((C.int)(config.GetPixFmt()))
			break
		}
	}

	if decoderCtx = ffmpeg.AvCodecAllocContext3(decoder); decoderCtx == nil {
		os.Exit(int(ffmpeg.AVERROR(syscall.ENOMEM)))
	}

	video = inputCtx.GetStreams()[videoStream]
	if ret = ffmpeg.AvCodecParametersToContext(decoderCtx, video.GetCodecpar()); ret < 0 {
		os.Exit(1)
	}

	decoderCtx.SetGetFormat((ffmpeg.AVCodecContextGetFormatFunc)(C.get_hw_format))

	if hwDeviceCtx, ret = hwDecoderInit(decoderCtx, hwType); ret < 0 {
		os.Exit(1)
	}

	if ret = ffmpeg.AvCodecOpen2(decoderCtx, decoder, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to open codec for stream #%d\n", videoStream)
		os.Exit(1)
	}

	// open the file to dump raw data
	outputFile, _ = os.OpenFile(os.Args[3], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	// actual decoding and dump the raw data
	for ret >= 0 {
		if ret = ffmpeg.AvReadFrame(inputCtx, &packet); ret < 0 {
			break
		}
		if videoStream == packet.GetStreamIndex() {
			ret = decodeWrite(decoderCtx, &packet, outputFile)
		}

		ffmpeg.AvPacketUnref(&packet)
	}

	// flush the decoder
	packet.SetData(nil)
	packet.SetSize(0)
	ret = decodeWrite(decoderCtx, &packet, outputFile)
	ffmpeg.AvPacketUnref(&packet)

	if outputFile != nil {
		outputFile.Close()
	}
	ffmpeg.AvCodecFreeContext(&decoderCtx)
	ffmpeg.AvFormatCloseInput(&inputCtx)
	ffmpeg.AvBufferUnref(&hwDeviceCtx)

	if ret != 0 {
		fmt.Fprintf(os.Stderr, "(error '%s')\n", ffmpeg.AvErr2str(ret))
	}
	os.Exit(int(ret))
}
