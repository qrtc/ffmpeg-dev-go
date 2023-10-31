package main

/*
#include <stdlib.h>
#include <stdio.h>

struct AVCodecContext;

static int av_pix_fmt_none;
static int av_pix_fmt_vaapi;

static inline void initial_pix_fmt(int x, int y)
{
	av_pix_fmt_none = x;
	av_pix_fmt_vaapi = y;
}

int get_vaapi_format(struct AVCodecContext *ctx, int *pix_fmts)
{
    const int *p;

    for (p = pix_fmts; *p != av_pix_fmt_none; p++) {
        if (*p == av_pix_fmt_vaapi)
            return *p;
    }

    fprintf(stderr, "Unable to decode this file using VA-API.\n");
    return av_pix_fmt_none;
}

*/
import "C"

import (
	"fmt"
	"os"
	"syscall"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

type streamContext struct {
	ifmtCtx *ffmpeg.AVFormatContext
	ofmtCtx *ffmpeg.AVFormatContext

	hwDeviceCtx *ffmpeg.AVBufferRef

	decoderCtx *ffmpeg.AVCodecContext
	encoderCtx *ffmpeg.AVCodecContext

	videoStream int32
	ost         *ffmpeg.AVStream
	initialized bool
}

func openInputFile(filename string, sc *streamContext) (ret int32) {
	var (
		decoder *ffmpeg.AVCodec
		video   *ffmpeg.AVStream
	)

	if ret = ffmpeg.AvFormatOpenInput(&sc.ifmtCtx, filename, nil, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Cannot open input file '%s', Error code: %s\n",
			filename, ffmpeg.AvErr2str(ret))
		return ret
	}

	if ret = ffmpeg.AvFormatFindStreamInfo(sc.ifmtCtx, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Cannot find input stream information. Error code: %s\n",
			ffmpeg.AvErr2str(ret))
		return ret
	}

	if ret = ffmpeg.AvFindBestStream(sc.ifmtCtx, ffmpeg.AVMEDIA_TYPE_VIDEO, -1, -1, &decoder, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Cannot find a video stream in the input file."+
			" Error code: %s\n", ffmpeg.AvErr2str(ret))
		return ret
	}
	sc.videoStream = ret

	if sc.decoderCtx = ffmpeg.AvCodecAllocContext3(decoder); sc.decoderCtx == nil {
		return ffmpeg.AVERROR(syscall.ENOMEM)
	}

	video = sc.ifmtCtx.GetStreams()[sc.videoStream]
	if ret = ffmpeg.AvCodecParametersToContext(sc.decoderCtx, video.GetCodecpar()); ret < 0 {
		fmt.Fprintf(os.Stderr, "avcodec_parameters_to_context error. Error code: %s\n",
			ffmpeg.AvErr2str(ret))
		return ret
	}

	sc.decoderCtx.SetHwDeviceCtx(ffmpeg.AvBufferRef(sc.hwDeviceCtx))
	if sc.decoderCtx.GetHwDeviceCtx() == nil {
		fmt.Fprintf(os.Stderr, "A hardware device reference create failed.\n")
		return ffmpeg.AVERROR(syscall.ENOMEM)
	}
	C.initial_pix_fmt((C.int)(ffmpeg.AV_PIX_FMT_NONE), (C.int)(ffmpeg.AV_PIX_FMT_VAAPI))
	sc.decoderCtx.SetGetFormat((ffmpeg.AVCodecContextGetFormatFunc)(C.get_vaapi_format))

	if ret = ffmpeg.AvCodecOpen2(sc.decoderCtx, decoder, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to open codec for decoding. Error code: %s\n",
			ffmpeg.AvErr2str(ret))
	}

	return 0
}

func encodeWrite(encPkt *ffmpeg.AVPacket, frame *ffmpeg.AVFrame, sc *streamContext) (ret int32) {

	ffmpeg.AvPacketUnref(encPkt)

	if ret = ffmpeg.AvCodecSendFrame(sc.encoderCtx, frame); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error during encoding. Error code: %s\n", ffmpeg.AvErr2str(ret))
		goto end
	}
	for {
		if ret = ffmpeg.AvCodecReceivePacket(sc.encoderCtx, encPkt); ret != 0 {
			break
		}

		encPkt.SetStreamIndex(0)
		ffmpeg.AvPacketRescaleTs(encPkt, sc.ifmtCtx.GetStreams()[sc.videoStream].GetTimeBase(),
			sc.ofmtCtx.GetStreams()[0].GetTimeBase())
		if ret = ffmpeg.AvInterleavedWriteFrame(sc.ofmtCtx, encPkt); ret < 0 {
			fmt.Fprintf(os.Stderr, "Error during writing data to output file. "+
				"Error code: %s\n", ffmpeg.AvErr2str(ret))
			return -1
		}
	}
end:
	if ret == ffmpeg.AVERROR_EOF {
		return 0
	}
	return ffmpeg.CondExpr(ret == ffmpeg.AVERROR(syscall.EAGAIN), 0, int32(-1))
}

func initialEncCodecCtx(encCodec *ffmpeg.AVCodec, sc *streamContext) (ret int32) {

	// we need to ref hw_frames_ctx of decoder to initialize encoder's codec.
	// Only after we get a decoded frame, can we obtain its hw_frames_ctx
	sc.encoderCtx.SetHwFramesCtx(ffmpeg.AvBufferRef(sc.decoderCtx.GetHwFramesCtx()))
	if sc.encoderCtx.GetHwFramesCtx() == nil {
		return ffmpeg.AVERROR(syscall.ENOMEM)
	}
	// set AVCodecContext Parameters for encoder, here we keep them stay
	// the same as decoder.
	// xxx: now the sample can't handle resolution change case.
	sc.encoderCtx.SetTimeBase(ffmpeg.AvInvQ(sc.decoderCtx.GetFramerate()))
	sc.encoderCtx.SetPixFmt(ffmpeg.AV_PIX_FMT_VAAPI)
	sc.encoderCtx.SetWidth(sc.decoderCtx.GetWidth())
	sc.encoderCtx.SetHeight(sc.decoderCtx.GetHeight())

	if ret = ffmpeg.AvCodecOpen2(sc.encoderCtx, encCodec, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to open encode codec. Error code: %s\n",
			ffmpeg.AvErr2str(ret))
		return ret
	}
	sc.ost.SetTimeBase(sc.encoderCtx.GetTimeBase())
	if ret = ffmpeg.AvCodecParametersFromContext(sc.ost.GetCodecpar(), sc.encoderCtx); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to copy the stream parameters. "+
			"Error code: %s\n", ffmpeg.AvErr2str(ret))
		return ret
	}

	// write the stream header
	if ret = ffmpeg.AvFormatWriteHeader(sc.ofmtCtx, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error while writing stream header. "+
			"Error code: %s\n", ffmpeg.AvErr2str(ret))
		return ret
	}

	return 0
}

func decEnc(pkt *ffmpeg.AVPacket, encCodec *ffmpeg.AVCodec, sc *streamContext) (ret int32) {
	var (
		frame *ffmpeg.AVFrame
	)

	if ret = ffmpeg.AvCodecSendPacket(sc.decoderCtx, pkt); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error during decoding. Error code: %s\n", ffmpeg.AvErr2str(ret))
		return ret
	}

	for ret >= 0 {
		if frame = ffmpeg.AvFrameAlloc(); frame == nil {
			return ffmpeg.AVERROR(syscall.ENOMEM)
		}

		ret = ffmpeg.AvCodecReceiveFrame(sc.decoderCtx, frame)
		if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
			ffmpeg.AvFrameFree(&frame)
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while decoding. Error code: %s\n", ffmpeg.AvErr2str(ret))
			goto fail
		}

		if !sc.initialized {
			if ret = initialEncCodecCtx(encCodec, sc); ret != 0 {
				fmt.Fprintf(os.Stderr, "Initial EncodeCtx failed.\n")
				goto fail
			}
			sc.initialized = true
		}

		if ret = encodeWrite(pkt, frame, sc); ret < 0 {
			fmt.Fprintf(os.Stderr, "Error during encoding and writing.\n")
		}
	fail:
		ffmpeg.AvFrameFree(&frame)
		if ret < 0 {
			return ret
		}
	}
	return 0
}

func main() {
	var (
		ret      int32
		decPkt   *ffmpeg.AVPacket
		encCodec *ffmpeg.AVCodec
		sc       = &streamContext{}
	)

	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input file> <encode codec> <output file>\n"+
			"The output format is guessed according to the file extension.\n"+
			"\n", os.Args[0])
		os.Exit(1)
	}

	if ret = ffmpeg.AvHWDeviceCtxCreate(&sc.hwDeviceCtx, ffmpeg.AV_HWDEVICE_TYPE_VAAPI,
		ffmpeg.NIL, nil, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to create a VAAPI device. Error code: %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(1)
	}

	if decPkt = ffmpeg.AvPacketAlloc(); decPkt == nil {
		fmt.Fprintf(os.Stderr, "Failed to allocate decode packet\n")
		goto end
	}

	if ret = openInputFile(os.Args[1], sc); ret < 0 {
		goto end
	}

	if encCodec = ffmpeg.AvCodecFindEncoderByName(os.Args[2]); encCodec == nil {
		fmt.Fprintf(os.Stderr, "Could not find encoder '%s'\n", os.Args[2])
		ret = -1
		goto end
	}

	if ret = ffmpeg.AvFormatAllocOutputContext2(&sc.ofmtCtx, nil, ffmpeg.NIL, os.Args[3]); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to deduce output format from file extension. Error code: "+
			"%s\n", ffmpeg.AvErr2str(ret))
		goto end
	}

	if sc.encoderCtx = ffmpeg.AvCodecAllocContext3(encCodec); sc.encoderCtx == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	if ret = ffmpeg.AvIOOpen(sc.ofmtCtx.GetPbAddr(), os.Args[3], ffmpeg.AVIO_FLAG_WRITE); ret < 0 {
		fmt.Fprintf(os.Stderr, "Cannot open output file. "+
			"Error code: %s\n", ffmpeg.AvErr2str(ret))
		goto end
	}

	// read all packets and only transcoding video
	for ret >= 0 {
		if ret = ffmpeg.AvReadFrame(sc.ifmtCtx, decPkt); ret < 0 {
			break
		}

		if sc.videoStream == decPkt.GetStreamIndex() {
			ret = decEnc(decPkt, encCodec, sc)
		}

		ffmpeg.AvPacketUnref(decPkt)
	}

	// flush decoder
	ffmpeg.AvPacketUnref(decPkt)
	decEnc(decPkt, encCodec, sc)

	// flush encoder
	encodeWrite(decPkt, nil, sc)

	// write the trailer for output stream
	ffmpeg.AvWriteTrailer(sc.ofmtCtx)

end:
	ffmpeg.AvFormatCloseInput(&sc.ifmtCtx)
	ffmpeg.AvFormatCloseInput(&sc.ofmtCtx)
	ffmpeg.AvCodecFreeContext(&sc.decoderCtx)
	ffmpeg.AvCodecFreeContext(&sc.encoderCtx)
	ffmpeg.AvBufferUnref(&sc.hwDeviceCtx)
	ffmpeg.AvPacketFree(&decPkt)
	os.Exit(int(ret))
}
