package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/qrtc/ffmpeg-dev-go"
)

var (
	filterDescr            = "aresample=8000,aformat=sample_fmts=s16:channel_layouts=mono"
	player                 = "ffplay -f s16le -ar 8000 -ac 1 -"
	audioStreamIndex int32 = -1
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

	// select the audio stream
	if ret = ffmpeg.AvFindBestStream(fmtCtx, ffmpeg.AVMEDIA_TYPE_AUDIO, -1, -1, &dec, 0); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot find an audio stream in the input file\n")
		return nil, nil, ret
	}
	audioStreamIndex = ret

	// create decoding context
	if decCtx = ffmpeg.AvCodecAllocContext3(dec); decCtx == nil {
		return nil, nil, ffmpeg.AVERROR(syscall.ENOMEM)
	}
	ffmpeg.AvCodecParametersToContext(decCtx, fmtCtx.GetStreams()[audioStreamIndex].GetCodecpar())

	// init the audio decoder
	if ret = ffmpeg.AvCodecOpen2(decCtx, dec, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot open audio decoder\n")
		return nil, nil, ret
	}

	return decCtx, fmtCtx, 0
}

func initFilters(decCtx *ffmpeg.AVCodecContext, fmtCtx *ffmpeg.AVFormatContext, filtersDescr string) (
	filterGraph *ffmpeg.AVFilterGraph, buffersinkCtx, buffersrcCtx *ffmpeg.AVFilterContext, ret int32) {
	var (
		abuffersrc        = ffmpeg.AvFilterGetByName("abuffer")
		abuffersink       = ffmpeg.AvFilterGetByName("abuffersink")
		outputs           = ffmpeg.AvFilterInoutAlloc()
		inputs            = ffmpeg.AvFilterInoutAlloc()
		outSampleFmts     = []ffmpeg.AVSampleFormat{ffmpeg.AV_SAMPLE_FMT_S16, -1}
		outChannelLayouts = []int64{int64(ffmpeg.AV_CH_LAYOUT_MONO), -1}
		outSampleRates    = []int32{8000, -1}
		timeBase          = fmtCtx.GetStreams()[audioStreamIndex].GetTimeBase()
		args              string
		outlink           *ffmpeg.AVFilterLink
	)

	filterGraph = ffmpeg.AvFilterGraphAlloc()
	if outputs == nil || inputs == nil || filterGraph == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	// buffer audio source: the decoded frames from the decoder will be inserted here.
	if decCtx.GetChannelLayout() == 0 {
		decCtx.SetChannelLayout(uint64(ffmpeg.AvGetDefaultChannelLayout(decCtx.GetChannels())))
	}
	args = fmt.Sprintf("time_base=%d/%d:sample_rate=%d:sample_fmt=%s:channel_layout=0x%d",
		timeBase.GetNum(), timeBase.GetDen(), decCtx.GetSampleRate(),
		ffmpeg.AvGetSampleFmtName(decCtx.GetSampleFmt()), decCtx.GetChannelLayout())

	ffmpeg.AvLog(nil, ffmpeg.AV_LOG_INFO, "audio source args: %s\n", args)

	if ret = ffmpeg.AvFilterGraphCreateFilter(&buffersrcCtx, abuffersrc, "in",
		args, nil, filterGraph); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot create audio buffer source\n")
		goto end
	}

	// buffer audio sink: to terminate the filter chain.
	if ret = ffmpeg.AvFilterGraphCreateFilter(&buffersinkCtx, abuffersink, "out",
		ffmpeg.NIL, nil, filterGraph); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot create audio buffer sink\n")
		goto end
	}

	if ret = ffmpeg.AvOptSetIntList(buffersinkCtx, "sample_fmts", &outSampleFmts[0], -1,
		ffmpeg.AV_OPT_SEARCH_CHILDREN); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot set output sample format\n")
		goto end
	}

	if ret = ffmpeg.AvOptSetIntList(buffersinkCtx, "channel_layouts", &outChannelLayouts[0], -1,
		ffmpeg.AV_OPT_SEARCH_CHILDREN); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot set output channel layout\n")
		goto end
	}

	if ret = ffmpeg.AvOptSetIntList(buffersinkCtx, "sample_rates", &outSampleRates[0], -1,
		ffmpeg.AV_OPT_SEARCH_CHILDREN); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot set output sample rate\n")
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

	// Print summary of the sink buffer
	// Note: args buffer is reused to store channel layout string
	outlink = buffersinkCtx.GetInputs()[0]
	args = ffmpeg.AvGetChannelLayoutString(-1, outlink.GetChannelLayout())
	ffmpeg.AvLog(nil, ffmpeg.AV_LOG_INFO, "Output: srate:%dHz fmt:%s chlayout:%s\n",
		outlink.GetSampleRate(),
		ffmpeg.AvGetSampleFmtName(outlink.GetFormat()),
		args)

end:
	ffmpeg.AvFilterInoutFree(&inputs)
	ffmpeg.AvFilterInoutFree(&outputs)
	return filterGraph, buffersinkCtx, buffersrcCtx, ret
}

func printFrame(frame *ffmpeg.AVFrame) {
	n := frame.GetNbSamples() * ffmpeg.AvGetChannelLayoutNbChannels(frame.GetChannelLayout())
	p := (*uint16)(unsafe.Pointer(frame.GetData()[0]))
	data := unsafe.Slice(p, n)
	for i := int32(0); i < n; i++ {
		os.Stdout.Write([]byte{byte(data[i] & 0xFF), byte(data[i] >> 8 & 0xFF)})
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
		fmt.Fprintf(os.Stderr, "Usage: %s file | %s\n", os.Args[0], player)
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

		if packet.GetStreamIndex() == audioStreamIndex {
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

				if ret >= 0 {
					// push the audio data from decoded frame into the filtergraph
					if ffmpeg.AvBuffersrcAddFrameFlags(buffersrcCtx, frame, ffmpeg.AV_BUFFERSRC_FLAG_KEEP_REF) < 0 {
						ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Error while feeding the audio filtergraph\n")
						break
					}

					// pull filtered audio from the filtergraph
					for {
						ret = ffmpeg.AvBuffersinkGetFrame(buffersinkCtx, filtFrame)
						if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
							break
						}
						if ret < 0 {
							goto end
						}
						printFrame(filtFrame)
						ffmpeg.AvFrameUnref(filtFrame)
					}
					ffmpeg.AvFrameUnref(frame)
				}
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
