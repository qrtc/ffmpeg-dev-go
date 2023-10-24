package main

// http-multiclient README.txt http://localhost:8082

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

var (
	videoFrameCount int
	srcFilename     string
)

func decodePacket(videoDecCtx *ffmpeg.AVCodecContext,
	pkt *ffmpeg.AVPacket, frame *ffmpeg.AVFrame) (ret int32) {

	if ret = ffmpeg.AvCodecSendPacket(videoDecCtx, pkt); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error while sending a packet to the decoder: %s\n", ffmpeg.AvErr2str(ret))
		return ret
	}

	for ret >= 0 {
		ret = ffmpeg.AvCodecReceiveFrame(videoDecCtx, frame)
		if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
			break
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while receiving a frame from the decoder: %s\n", ffmpeg.AvErr2str(ret))
			return ret
		}

		if ret > 0 {
			videoFrameCount++
			if sd := ffmpeg.AvFrameGetSideData(frame, ffmpeg.AV_FRAME_DATA_MOTION_VECTORS); sd != nil {
				mvs := (*ffmpeg.AVMotionVector)(unsafe.Pointer(sd.GetData()))
				nbMvs := int(sd.GetSize()) / int(unsafe.Sizeof(*mvs))
				for i := 0; i < nbMvs; i++ {
					mv := ffmpeg.PointerOffset(mvs, i)
					fmt.Fprintf(os.Stdout, "%d,%2d,%2d,%2d,%4d,%4d,%4d,%4d,0x%d\n",
						videoFrameCount, mv.GetSource(),
						mv.GetW(), mv.GetH(), mv.GetSrcX(), mv.GetSrcY(),
						mv.GetDstX(), mv.GetDstY(), mv.GetFlags())
				}
			}
			ffmpeg.AvFrameUnref(frame)
		}
	}

	return 0
}

func openCodecContext(fmtCtx *ffmpeg.AVFormatContext, _type ffmpeg.AVMediaType) (
	videoStreamIdx int32, videoStream *ffmpeg.AVStream, videoDecCtx *ffmpeg.AVCodecContext, ret int32) {
	var (
		dec *ffmpeg.AVCodec
		opt *ffmpeg.AVDictionary
	)
	ret = ffmpeg.AvFindBestStream(fmtCtx, _type, -1, -1, &dec, 0)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not find %s stream in input file '%s'\n",
			ffmpeg.AvGetMediaTypeString(_type), srcFilename)
		return 0, nil, nil, ret
	} else {
		videoStreamIdx = ret
		videoStream = fmtCtx.GetStreams()[videoStreamIdx]

		if videoDecCtx = ffmpeg.AvCodecAllocContext3(dec); videoDecCtx == nil {
			fmt.Fprintf(os.Stderr, "Failed to allocate codec\n")
			return 0, nil, nil, ffmpeg.AVERROR(syscall.EINVAL)
		}

		if ret = ffmpeg.AvCodecParametersToContext(videoDecCtx, videoStream.GetCodecpar()); ret < 0 {
			fmt.Fprintf(os.Stderr, "Failed to copy codec parameters to codec context\n")
			return 0, nil, nil, ret
		}

		// Init the video decoder
		ffmpeg.AvDictSet(&opt, "flags2", "+export_mvs", 0)
		if ret = ffmpeg.AvCodecOpen2(videoDecCtx, dec, &opt); ret < 0 {
			fmt.Fprintf(os.Stderr, "Failed to open %s codec\n", ffmpeg.AvGetMediaTypeString(_type))
			return 0, nil, nil, ret
		}

	}
	return videoStreamIdx, videoStream, videoDecCtx, 0
}

func main() {
	var (
		fmtCtx         *ffmpeg.AVFormatContext
		videoDecCtx    *ffmpeg.AVCodecContext
		videoStream    *ffmpeg.AVStream
		videoStreamIdx int32
		frame          *ffmpeg.AVFrame

		pkt ffmpeg.AVPacket
		ret int32
	)

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <video>\n", os.Args[0])
		os.Exit(1)
	}
	srcFilename = os.Args[1]

	if ret = ffmpeg.AvFormatOpenInput(&fmtCtx, srcFilename, nil, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open source file %s\n", srcFilename)
		os.Exit(1)
	}

	if ret = ffmpeg.AvFormatFindStreamInfo(fmtCtx, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not find stream information\n")
		os.Exit(1)
	}

	videoStreamIdx, videoStream, videoDecCtx, ret = openCodecContext(fmtCtx, ffmpeg.AVMEDIA_TYPE_VIDEO)

	ffmpeg.AvDumpFormat(fmtCtx, 0, srcFilename, 0)

	if videoStream == nil {
		fmt.Fprintf(os.Stderr, "Could not find video stream in the input, aborting\n")
		ret = 1
		goto end
	}

	if frame = ffmpeg.AvFrameAlloc(); frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate frame\n")
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	fmt.Fprintf(os.Stdout, "framenum,source,blockw,blockh,srcx,srcy,dstx,dsty,flags\n")

	// read frames from the file
	for ffmpeg.AvReadFrame(fmtCtx, &pkt) >= 0 {
		if pkt.GetStreamIndex() == videoStreamIdx {
			ret = decodePacket(videoDecCtx, &pkt, frame)
		}
		ffmpeg.AvPacketUnref(&pkt)
		if ret < 0 {
			break
		}
	}

	// flush cached frames
	decodePacket(videoDecCtx, nil, frame)

end:
	ffmpeg.AvCodecFreeContext(&videoDecCtx)
	ffmpeg.AvFormatCloseInput(&fmtCtx)
	ffmpeg.AvFrameFree(&frame)
	if ret < 0 {
		os.Exit(1)
	}
}
