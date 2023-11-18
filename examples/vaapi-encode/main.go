package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

func setHwFrameCtx(ctx *ffmpeg.AVCodecContext, hwDeviceCtx *ffmpeg.AVBufferRef, width, height int32) (ret int32) {
	var (
		hwFramesRef *ffmpeg.AVBufferRef
		framesCtx   *ffmpeg.AVHWFramesContext
	)

	if hwFramesRef = ffmpeg.AvHWFrameCtxAlloc(hwDeviceCtx); hwFramesRef == nil {
		fmt.Fprintf(os.Stderr, "Failed to create VAAPI frame context.\n")
		return -1
	}
	framesCtx = (*ffmpeg.AVHWFramesContext)(unsafe.Pointer(hwFramesRef.GetData()))
	framesCtx.SetFormat(ffmpeg.AV_PIX_FMT_VAAPI)
	framesCtx.SetSwFormat(ffmpeg.AV_PIX_FMT_NV12)
	framesCtx.SetWidth(width)
	framesCtx.SetHeight(height)
	framesCtx.SetInitialPoolSize(20)
	if ret = ffmpeg.AvHWFrameCtxInit(hwFramesRef); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to initialize VAAPI frame context."+
			"Error code: %s\n", ffmpeg.AvErr2str(ret))
		ffmpeg.AvBufferUnref(&hwFramesRef)
		return ret
	}
	ctx.SetHwFramesCtx(ffmpeg.AvBufferRef(hwFramesRef))
	if ctx.GetHwFramesCtx() == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
	}

	ffmpeg.AvBufferUnref(&hwFramesRef)
	return ret
}

func encodeWrite(avctx *ffmpeg.AVCodecContext, frame *ffmpeg.AVFrame, fout *os.File) (ret int32) {
	var (
		encPkt *ffmpeg.AVPacket
	)

	if encPkt = ffmpeg.AvPacketAlloc(); encPkt == nil {
		return ffmpeg.AVERROR(syscall.ENOMEM)
	}

	if ret = ffmpeg.AvCodecSendFrame(avctx, frame); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error code: %s\n", ffmpeg.AvErr2str(ret))
		goto end
	}
	for {
		if ret = ffmpeg.AvCodecReceivePacket(avctx, encPkt); ret != 0 {
			break
		}

		encPkt.SetStreamIndex(0)
		fout.Write(unsafe.Slice(encPkt.GetData(), encPkt.GetSize()))
	}
end:
	ffmpeg.AvPacketFree(&encPkt)
	ret = ffmpeg.CondExpr(ret == ffmpeg.AVERROR(syscall.EAGAIN), 0, int32(-1))
	return ret
}

func main() {
	var (
		hwDeviceCtx         *ffmpeg.AVBufferRef
		fin, fout           *os.File
		swFrame, hwFrame    *ffmpeg.AVFrame
		avctx               *ffmpeg.AVCodecContext
		codec               *ffmpeg.AVCodec
		width, height, size int32
		encName             = "h264_vaapi"
		ret                 int32
		err                 error
	)

	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Usage: %s <width> <height> <input file> <output file>\n", os.Args[0])
		os.Exit(1)
	}

	if n, err := strconv.ParseInt(os.Args[1], 10, 32); err != nil {
		width = int32(n)
	}
	if n, err := strconv.ParseInt(os.Args[2], 10, 32); err != nil {
		height = int32(n)
	}
	size = width * height

	fin, err = os.OpenFile(os.Args[3], os.O_RDONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", os.Args[3])
		os.Exit(1)
	}
	fout, err = os.OpenFile(os.Args[4], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", os.Args[4])
		goto close
	}

	if ret = ffmpeg.AvHWDeviceCtxCreate(&hwDeviceCtx, ffmpeg.AV_HWDEVICE_TYPE_VAAPI,
		ffmpeg.NIL, nil, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to create a VAAPI device. Error code: %s\n", ffmpeg.AvErr2str(ret))
		goto close
	}

	if codec = ffmpeg.AvCodecFindEncoderByName(encName); codec == nil {
		fmt.Fprintf(os.Stderr, "Could not find encoder.\n")
		ret = -1
		goto close
	}

	if avctx = ffmpeg.AvCodecAllocContext3(codec); avctx == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto close
	}
	avctx.SetWidth(width)
	avctx.SetHeight(height)
	avctx.SetTimeBase(ffmpeg.AvMakeQ(1, 25))
	avctx.SetFramerate(ffmpeg.AvMakeQ(25, 1))
	avctx.SetPixFmt(ffmpeg.AV_PIX_FMT_VAAPI)

	// set hw_frames_ctx for encoder's AVCodecContext
	if ret = setHwFrameCtx(avctx, hwDeviceCtx, width, height); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to set hwframe context.\n")
		goto close
	}

	if ret = ffmpeg.AvCodecOpen2(avctx, codec, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Cannot open video encoder codec. Error code: %s\n", ffmpeg.AvErr2str(ret))
		goto close
	}

	for {
		if swFrame = ffmpeg.AvFrameAlloc(); swFrame == nil {
			ret = ffmpeg.AVERROR(syscall.ENOMEM)
			goto close
		}
		// read data into software frame, and transfer them into hw frame
		swFrame.SetWidth(width)
		swFrame.SetHeight(height)
		swFrame.SetFormat(ffmpeg.AV_PIX_FMT_NV12)
		if ret = ffmpeg.AvFrameGetBuffer(swFrame, 0); ret < 0 {
			goto close
		}
		frameData := ffmpeg.SliceSlice(&swFrame.GetData()[0], 2, size)
		if _, err = fin.Read(frameData[0]); err != nil {
			break
		}
		if _, err = fin.Read(frameData[1][0 : size/2]); err != nil {
			break
		}

		if hwFrame = ffmpeg.AvFrameAlloc(); hwFrame == nil {
			ret = ffmpeg.AVERROR(syscall.ENOMEM)
			goto close
		}
		if ret = ffmpeg.AvHWFrameGetBuffer(avctx.GetHwDeviceCtx(), hwFrame, 0); ret < 0 {
			fmt.Fprintf(os.Stderr, "Error code: %s\n", ffmpeg.AvErr2str(ret))
			goto close
		}
		if hwFrame.GetHwFramesCtx() == nil {
			ret = ffmpeg.AVERROR(syscall.ENOMEM)
			goto close
		}
		if ret = ffmpeg.AvHWFrameTransferData(hwFrame, swFrame, 0); ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while transferring frame data to surface."+
				"Error code: %s\n", ffmpeg.AvErr2str(ret))
			goto close
		}

		if ret = encodeWrite(avctx, hwFrame, fout); ret < 0 {
			fmt.Fprintf(os.Stderr, "Failed to encode.\n")
			goto close
		}
		ffmpeg.AvFrameFree(&hwFrame)
		ffmpeg.AvFrameFree(&swFrame)
	}

	// flush encoder
	if ret = encodeWrite(avctx, nil, fout); ret == ffmpeg.AVERROR_EOF {
		ret = 0
	}

close:
	if fin != nil {
		fin.Close()
	}
	if fout != nil {
		fout.Close()
	}
	ffmpeg.AvFrameFree(&swFrame)
	ffmpeg.AvFrameFree(&hwFrame)
	ffmpeg.AvCodecFreeContext(&avctx)
	ffmpeg.AvBufferUnref(&hwDeviceCtx)

	os.Exit(int(ret))
}
