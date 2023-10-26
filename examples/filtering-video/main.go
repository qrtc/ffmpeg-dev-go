package main

/*
#include <stdlib.h>
#include <stdio.h>

static void putcs(uint16_t *p, int n)
{
    const uint16_t *p_end = p + n;
    while (p < p_end) {
        fputc(*p    & 0xff, stdout);
        fputc(*p>>8 & 0xff, stdout);
        p++;
    }
    fflush(stdout);
}
*/
import "C"

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/qrtc/ffmpeg-dev-go"
)

var (
	filterDescr = "scale=78:24,transpose=cclock"
	// other way:
	// scale=78:24 [scl]; [scl] transpose=cclock // assumes "[in]" and "[out]" to be input output pads respectively
	videoStreamIndex int32 = -1
	lastPts          int64 = ffmpeg.AV_NOPTS_VALUE
	displayStr             = " .-+#"
)

func openInputFile(filename string) (decCtx *ffmpeg.AVCodecContext, fmtCtx *ffmpeg.AVFormatContext, ret int32) {
	var (
		dec *ffmpeg.AVCodec
	)

	if ret = ffmpeg.AvFormatOpenInput(&fmtCtx, filename, nil, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot open input file\n")
		return nil, nil, ret
	}

	if ret = ffmpeg.AvFormatFindStreamInfo(fmtCtx, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot find stream information\n")
		return nil, nil, ret
	}

	// select the video stream
	if ret = ffmpeg.AvFindBestStream(fmtCtx, ffmpeg.AVMEDIA_TYPE_VIDEO, -1, -1, &dec, 0); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot find a video stream in the input file\n")
		return nil, nil, ret
	}
	videoStreamIndex = ret

	// create decoding context
	if decCtx = ffmpeg.AvCodecAllocContext3(dec); decCtx == nil {
		return nil, nil, ffmpeg.AVERROR(syscall.ENOMEM)
	}
	ffmpeg.AvCodecParametersToContext(decCtx, fmtCtx.GetStreams()[videoStreamIndex].GetCodecpar())

	// init the audio decoder
	if ret = ffmpeg.AvCodecOpen2(decCtx, dec, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot open video decoder\n")
		return nil, nil, ret
	}

	return decCtx, fmtCtx, 0
}

func initFilters(decCtx *ffmpeg.AVCodecContext, fmtCtx *ffmpeg.AVFormatContext, filtersDescr string) (
	filterGraph *ffmpeg.AVFilterGraph, buffersinkCtx, buffersrcCtx *ffmpeg.AVFilterContext, ret int32) {
	var (
		buffersrc  = ffmpeg.AvFilterGetByName("buffer")
		buffersink = ffmpeg.AvFilterGetByName("buffersink")
		outputs    = ffmpeg.AvFilterInoutAlloc()
		inputs     = ffmpeg.AvFilterInoutAlloc()
		timeBase   = fmtCtx.GetStreams()[videoStreamIndex].GetTimeBase()
		pixFmts    = []ffmpeg.AVPixelFormat{ffmpeg.AV_PIX_FMT_GRAY8, ffmpeg.AV_PIX_FMT_NONE}
		args       string
	)

	filterGraph = ffmpeg.AvFilterGraphAlloc()
	if outputs == nil || inputs == nil || filterGraph == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	// buffer video source: the decoded frames from the decoder will be inserted here.
	args = fmt.Sprintf("video_size=%dx%d:pix_fmt=%d:time_base=%d/%d:pixel_aspect=%d/%d",
		decCtx.GetWidth(), decCtx.GetHeight(), decCtx.GetPixFmt(),
		timeBase.GetNum(), timeBase.GetDen(),
		decCtx.GetSampleAspectRatioAddr().GetNum(), decCtx.GetSampleAspectRatioAddr().GetDen())

	ffmpeg.AvLog(nil, ffmpeg.AV_LOG_INFO, "video source args: %s\n", args)

	if ret = ffmpeg.AvFilterGraphCreateFilter(&buffersrcCtx, buffersrc, "in",
		args, nil, filterGraph); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot create buffer source\n")
		goto end
	}

	// buffer audio sink: to terminate the filter chain.
	if ret = ffmpeg.AvFilterGraphCreateFilter(&buffersinkCtx, buffersink, "out",
		ffmpeg.NIL, nil, filterGraph); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot create buffer sink\n")
		goto end
	}

	if ret = ffmpeg.AvOptSetIntList(buffersinkCtx, "pix_fmts", &pixFmts[0], ffmpeg.AV_PIX_FMT_NONE,
		ffmpeg.AV_OPT_SEARCH_CHILDREN); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot set output pixel format\n")
		goto end
	}

	// Set the endpoints for the filter graph. The filter_graph will
	// be linked to the graph described by filters_descr.

	// The buffer source output must be connected to the input pad of
	// the first filter described by filters_descr; since the first
	// filter input label is not specified, it is set to "in" by
	// default.
	outputs.SetName("in")
	outputs.SetFilterCtx(buffersrcCtx)
	outputs.SetPadIdx(0)
	outputs.SetNext(nil)

	// The buffer sink input must be connected to the output pad of
	// the last filter described by filters_descr; since the last
	// filter output label is not specified, it is set to "out" by
	// default.
	inputs.SetName("out")
	inputs.SetFilterCtx(buffersinkCtx)
	inputs.SetPadIdx(0)
	inputs.SetNext(nil)

	if ret = ffmpeg.AvFilterGraphParsePtr(filterGraph, filtersDescr, &inputs, &outputs, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot parse filter graph\n")
		goto end
	}

	if ret = ffmpeg.AvFilterGraphConfig(filterGraph, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot config filter graph\n")
		goto end
	}

end:
	ffmpeg.AvFilterInoutFree(&inputs)
	ffmpeg.AvFilterInoutFree(&outputs)
	return filterGraph, buffersinkCtx, buffersrcCtx, ret
}

func displayFrame(frame *ffmpeg.AVFrame, timeBase ffmpeg.AVRational) {
	if frame.GetPts() != ffmpeg.AV_NOPTS_VALUE {
		if lastPts != ffmpeg.AV_NOPTS_VALUE {
			// sleep roughly the right amount of time;
			// usleep is in microseconds, just like AV_TIME_BASE.
			delay := ffmpeg.AvRescaleQ(frame.GetPts()-lastPts, timeBase, ffmpeg.AV_TIME_BASE_Q)
			if delay > 0 && delay < 1000000 {
				time.Sleep(time.Duration(delay) * time.Microsecond)
			}
		}
		lastPts = frame.GetPts()
	}

	// Trivial ASCII grayscale display.
	data := unsafe.Slice(frame.GetData()[0], frame.GetHeight()*frame.GetHeight())

	os.Stdout.Write([]byte{byte(033), byte('c'), byte('\n')})
	fmt.Fprintf(os.Stdout, "\n")
	for y := int32(0); y < frame.GetHeight(); y++ {
		idx := frame.GetLinesize()[0] * y
		for x := int32(0); x < frame.GetWidth(); x++ {
			fmt.Fprintf(os.Stdout, "%c", displayStr[data[idx+x]/52])
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
}

func main() {
	var (
		packet        ffmpeg.AVPacket
		ret           int32
		decCtx        *ffmpeg.AVCodecContext
		fmtCtx        *ffmpeg.AVFormatContext
		filterGraph   *ffmpeg.AVFilterGraph
		buffersinkCtx *ffmpeg.AVFilterContext
		buffersrcCtx  *ffmpeg.AVFilterContext
	)
	frame := ffmpeg.AvFrameAlloc()
	filtFrame := ffmpeg.AvFrameAlloc()
	if frame == nil || filtFrame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate frame")
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s file\n", os.Args[0])
		os.Exit(1)
	}

	if decCtx, fmtCtx, ret = openInputFile(os.Args[1]); ret < 0 {
		goto end
	}
	if filterGraph, buffersinkCtx, buffersrcCtx, ret = initFilters(decCtx, fmtCtx, filterDescr); ret < 0 {
		goto end
	}

	// read all packets
	for {
		if ret = ffmpeg.AvReadFrame(fmtCtx, &packet); ret < 0 {
			break
		}

		if packet.GetStreamIndex() == videoStreamIndex {
			if ret = ffmpeg.AvCodecSendPacket(decCtx, &packet); ret < 0 {
				ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Error while sending a packet to the decoder\n")
				break
			}

			for ret >= 0 {
				ret = ffmpeg.AvCodecReceiveFrame(decCtx, frame)
				if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
					break
				} else if ret < 0 {
					ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Error while receiving a frame from the decoder\n")
					goto end
				}

				frame.SetPts(frame.GetBestEffortTimestamp())

				// push the decoded frame into the filtergraph
				if ffmpeg.AvBuffersrcAddFrameFlags(buffersrcCtx, frame, ffmpeg.AV_BUFFERSRC_FLAG_KEEP_REF) < 0 {
					ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Error while feeding the filtergraph\n")
					break
				}

				// pull filtered frames from the filtergraph
				for {
					ret = ffmpeg.AvBuffersinkGetFrame(buffersinkCtx, filtFrame)
					if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
						break
					}
					if ret < 0 {
						goto end
					}
					displayFrame(filtFrame, buffersinkCtx.GetInputs()[0].GetTimeBase())
					ffmpeg.AvFrameUnref(filtFrame)
				}
				ffmpeg.AvFrameUnref(frame)
			}
		}
		ffmpeg.AvPacketUnref(&packet)
	}

end:
	ffmpeg.AvFilterGraphFree(&filterGraph)
	ffmpeg.AvCodecFreeContext(&decCtx)
	ffmpeg.AvFormatCloseInput(&fmtCtx)
	ffmpeg.AvFrameFree(&frame)
	ffmpeg.AvFrameFree(&filtFrame)

	if ret < 0 && ret != ffmpeg.AVERROR_EOF {
		fmt.Fprintf(os.Stderr, "Error occurred: %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(1)
	}
	os.Exit(0)
}
