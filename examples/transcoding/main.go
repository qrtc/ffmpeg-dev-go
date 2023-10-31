package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

type filteringContext struct {
	buffersinkCtx *ffmpeg.AVFilterContext
	buffersrcCtx  *ffmpeg.AVFilterContext
	filterGraph   *ffmpeg.AVFilterGraph
	encPkt        *ffmpeg.AVPacket
	filteredFrame *ffmpeg.AVFrame
}

type streamContext struct {
	decCtx   *ffmpeg.AVCodecContext
	encCtx   *ffmpeg.AVCodecContext
	decFrame *ffmpeg.AVFrame
}

// Open an input file and the required decoder.
func openInputFile(fileName string) (ifmtCtx *ffmpeg.AVFormatContext, streamCtx []streamContext, ret int32) {
	if ret = ffmpeg.AvFormatOpenInput(&ifmtCtx, fileName, nil, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot open input file\n")
		return nil, nil, ret
	}

	if ret = ffmpeg.AvFormatFindStreamInfo(ifmtCtx, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot find stream information\n")
		return nil, nil, ret
	}

	streamCtx = make([]streamContext, ifmtCtx.GetNbStreams())

	for i := 0; i < int(ifmtCtx.GetNbStreams()); i++ {
		stream := ifmtCtx.GetStreams()[i]
		dec := ffmpeg.AvCodecFindDecoder(stream.GetCodecpar().GetCodecId())
		if dec == nil {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Failed to find decoder for stream #%d\n", i)
			return nil, nil, ffmpeg.AVERROR_DECODER_NOT_FOUND
		}
		codecCtx := ffmpeg.AvCodecAllocContext3(dec)
		if codecCtx == nil {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Failed to allocate the decoder context for stream #%d\n", i)
			return nil, nil, ffmpeg.AVERROR(syscall.ENOMEM)
		}
		if ret = ffmpeg.AvCodecParametersToContext(codecCtx, stream.GetCodecpar()); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Failed to copy decoder parameters to input decoder context "+
				"for stream #%d\n", i)
			return nil, nil, ret
		}
		// Reencode video & audio and remux subtitles etc.
		if codecCtx.GetCodecType() == ffmpeg.AVMEDIA_TYPE_VIDEO ||
			codecCtx.GetCodecType() == ffmpeg.AVMEDIA_TYPE_AUDIO {
			if codecCtx.GetCodecType() == ffmpeg.AVMEDIA_TYPE_VIDEO {
				codecCtx.SetFramerate(ffmpeg.AvGuessFrameRate(ifmtCtx, stream, nil))
			}
			// Open decoder
			if ret = ffmpeg.AvCodecOpen2(codecCtx, dec, nil); ret < 0 {
				ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Failed to open decoder for stream #%d\n", i)
				return nil, nil, ret
			}
		}
		streamCtx[i].decCtx = codecCtx

		if streamCtx[i].decFrame = ffmpeg.AvFrameAlloc(); streamCtx[i].decFrame == nil {
			return nil, nil, ffmpeg.AVERROR(syscall.ENOMEM)
		}
	}

	ffmpeg.AvDumpFormat(ifmtCtx, 0, fileName, 0)
	return ifmtCtx, streamCtx, 0
}

func openOutputFile(ifmtCtx *ffmpeg.AVFormatContext, streamCtx []streamContext,
	filename string) (ofmtCtx *ffmpeg.AVFormatContext, ret int32) {
	if ffmpeg.AvFormatAllocOutputContext2(&ofmtCtx, nil, ffmpeg.NIL, filename); ofmtCtx == nil {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Could not create output context\n")
		return nil, ret
	}

	for i := 0; i < len(streamCtx); i++ {
		outStream := ffmpeg.AvFormatNewStream(ofmtCtx, nil)
		if outStream == nil {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Failed allocating output stream\n")
			return nil, ret
		}

		inStream := ifmtCtx.GetStreams()[i]
		decCtx := streamCtx[i].decCtx

		if decCtx.GetCodecType() == ffmpeg.AVMEDIA_TYPE_VIDEO ||
			decCtx.GetCodecType() == ffmpeg.AVMEDIA_TYPE_AUDIO {
			encoder := ffmpeg.AvCodecFindEncoder(decCtx.GetCodecId())
			if encoder == nil {
				ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Necessary encoder not found\n")
				return nil, ffmpeg.AVERROR_INVALIDDATA
			}
			encCtx := ffmpeg.AvCodecAllocContext3(encoder)
			if encCtx == nil {
				ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Failed to allocate the encoder context\n")
				return nil, ffmpeg.AVERROR(syscall.ENOMEM)
			}

			// In this example, we transcode to same properties (picture size,
			// sample rate etc.). These properties can be changed for output
			// streams easily using filters
			if decCtx.GetCodecType() == ffmpeg.AVMEDIA_TYPE_VIDEO {
				encCtx.SetHeight(decCtx.GetHeight())
				encCtx.SetWidth(decCtx.GetWidth())
				encCtx.SetSampleAspectRatio(decCtx.GetSampleAspectRatio())
				// take first format from list of supported formats
				if len(encoder.GetPixFmts()) != 0 {
					encCtx.SetPixFmt(encoder.GetPixFmts()[0])
				} else {
					encCtx.SetPixFmt(decCtx.GetPixFmt())
				}
				// video time_base can be set to whatever is handy and supported by encoder
				encCtx.SetTimeBase(ffmpeg.AvInvQ(decCtx.GetFramerate()))
			} else {
				encCtx.SetSampleRate(decCtx.GetSampleRate())
				if ret = ffmpeg.AvChannelLayoutCopy(encCtx.GetChLayoutAddr(),
					decCtx.GetChLayoutAddr()); ret < 0 {
					return nil, ret
				}
				// take first format from list of supported formats
				encCtx.SetSampleFmt(encoder.GetSampleFmts()[0])
				encCtx.SetTimeBase(ffmpeg.AvMakeQ(1, encCtx.GetSampleRate()))
			}

			if (ofmtCtx.GetOformat().GetFlags() & ffmpeg.AVFMT_GLOBALHEADER) != 0 {
				encCtx.SetFlags(encCtx.GetFlags() | ffmpeg.AV_CODEC_FLAG_GLOBAL_HEADER)
			}

			// Third parameter can be used to pass settings to encoder
			if ret = ffmpeg.AvCodecOpen2(encCtx, encoder, nil); ret < 0 {
				ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot open video encoder for stream #%d\n", i)
				return nil, ret
			}
			if ret = ffmpeg.AvCodecParametersFromContext(outStream.GetCodecpar(), encCtx); ret < 0 {
				ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Failed to copy encoder parameters to output stream #%d\n", i)
				return nil, ret
			}

			outStream.SetTimeBase(encCtx.GetTimeBase())
			streamCtx[i].encCtx = encCtx
		} else if decCtx.GetCodecType() == ffmpeg.AVMEDIA_TYPE_UNKNOWN {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Elementary stream #%d is of unknown type, cannot proceed\n", i)
			return nil, ffmpeg.AVERROR_INVALIDDATA
		} else {
			// if this stream must be remuxed
			if ret = ffmpeg.AvCodecParametersCopy(outStream.GetCodecpar(), inStream.GetCodecpar()); ret < 0 {
				ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Copying parameters for stream #%d failed\n", i)
				return nil, ret
			}
			outStream.SetTimeBase(inStream.GetTimeBase())
		}
	}
	ffmpeg.AvDumpFormat(ofmtCtx, 0, filename, 1)

	if (ofmtCtx.GetOformat().GetFlags() & ffmpeg.AVFMT_NOFILE) == 0 {
		if ret = ffmpeg.AvIOOpen(ofmtCtx.GetPbAddr(), filename, ffmpeg.AVIO_FLAG_WRITE); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Could not open output file '%s'", filename)
			return nil, ret
		}
	}

	// init muxer, write output file header
	if ret = ffmpeg.AvFormatWriteHeader(ofmtCtx, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Error occurred when opening output file\n")
		return nil, ret
	}

	return ofmtCtx, 0
}

func initFilter(fctx *filteringContext, decCtx, encCtx *ffmpeg.AVCodecContext, filterSpec string) (ret int32) {
	var (
		buffersrc     *ffmpeg.AVFilter
		buffersink    *ffmpeg.AVFilter
		buffersrcCtx  *ffmpeg.AVFilterContext
		buffersinkCtx *ffmpeg.AVFilterContext
		outputs       = ffmpeg.AvFilterInoutAlloc()
		inputs        = ffmpeg.AvFilterInoutAlloc()
		filterGraph   = ffmpeg.AvFilterGraphAlloc()
		args          string
	)

	if outputs == nil || inputs == nil || filterGraph == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	if decCtx.GetCodecType() == ffmpeg.AVMEDIA_TYPE_VIDEO {
		buffersrc = ffmpeg.AvFilterGetByName("buffer")
		buffersink = ffmpeg.AvFilterGetByName("buffersink")
		if buffersrc == nil || buffersink == nil {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "filtering source or sink element not found\n")
			ret = ffmpeg.AVERROR_UNKNOWN
			goto end
		}

		args = fmt.Sprintf("video_size=%dx%d:pix_fmt=%d:time_base=%d/%d:pixel_aspect=%d/%d",
			decCtx.GetWidth(), decCtx.GetHeight(), decCtx.GetPixFmt(),
			decCtx.GetTimeBaseAddr().GetNum(), decCtx.GetTimeBaseAddr().GetDen(),
			decCtx.GetSampleAspectRatioAddr().GetNum(), decCtx.GetSampleAspectRatioAddr().GetDen())

		if ret = ffmpeg.AvFilterGraphCreateFilter(&buffersrcCtx, buffersrc, "in",
			args, nil, filterGraph); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot create buffer source\n")
			goto end
		}

		if ret = ffmpeg.AvFilterGraphCreateFilter(&buffersinkCtx, buffersink, "out",
			ffmpeg.NIL, nil, filterGraph); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot create buffer sink\n")
			goto end
		}

		if ret = ffmpeg.AvOptSetBin(buffersinkCtx, "pix_fmts",
			encCtx.GetPixFmtAddr(), unsafe.Sizeof(encCtx.GetPixFmt()), ffmpeg.AV_OPT_SEARCH_CHILDREN); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot set output pixel format\n")
			goto end
		}
	} else if decCtx.GetCodecType() == ffmpeg.AVMEDIA_TYPE_AUDIO {
		buffersrc = ffmpeg.AvFilterGetByName("abuffer")
		buffersink = ffmpeg.AvFilterGetByName("abuffersink")
		if buffersrc == nil || buffersink == nil {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "filtering source or sink element not found\n")
			ret = ffmpeg.AVERROR_UNKNOWN
			goto end
		}

		if decCtx.GetChLayoutAddr().GetOrder() == ffmpeg.AV_CHANNEL_ORDER_UNSPEC {
			ffmpeg.AvChannelLayoutDefault(decCtx.GetChLayoutAddr(), decCtx.GetChLayoutAddr().GetNbChannels())
		}

		args = fmt.Sprintf("time_base=%d/%d:sample_rate=%d:sample_fmt=%s:channel_layout=%s",
			decCtx.GetTimeBaseAddr().GetNum(), decCtx.GetTimeBaseAddr().GetDen(), decCtx.GetSampleRate(),
			ffmpeg.AvGetSampleFmtName(decCtx.GetSampleFmt()),
			ffmpeg.AvChannelLayoutDescribe(decCtx.GetChLayoutAddr()))

		if ret = ffmpeg.AvFilterGraphCreateFilter(&buffersrcCtx, buffersrc, "in",
			args, nil, filterGraph); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot create audio buffer source\n")
			goto end
		}

		if ret = ffmpeg.AvFilterGraphCreateFilter(&buffersinkCtx, buffersink, "out",
			ffmpeg.NIL, nil, filterGraph); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot create audio buffer sink\n")
			goto end
		}

		if ret = ffmpeg.AvOptSetBin(buffersinkCtx, "sample_fmts",
			encCtx.GetSampleFmtAddr(), unsafe.Sizeof(encCtx.GetSampleFmt()), ffmpeg.AV_OPT_SEARCH_CHILDREN); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot set output sample format\n")
			goto end
		}

		if ret = ffmpeg.AvOptSet(buffersinkCtx, "ch_layouts",
			ffmpeg.AvChannelLayoutDescribe(encCtx.GetChLayoutAddr()), ffmpeg.AV_OPT_SEARCH_CHILDREN); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot set output channel layout\n")
			goto end
		}

		if ret = ffmpeg.AvOptSetBin(buffersinkCtx, "sample_rates",
			encCtx.GetSampleRateAddr(), unsafe.Sizeof(encCtx.GetSampleRate()), ffmpeg.AV_OPT_SEARCH_CHILDREN); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot set output sample rate\n")
			goto end
		}
	} else {
		ret = ffmpeg.AVERROR_UNKNOWN
		goto end
	}

	// Endpoints for the filter graph.
	outputs.SetName("in")
	outputs.SetFilterCtx(buffersrcCtx)
	outputs.SetPadIdx(0)
	outputs.SetNext(nil)

	inputs.SetName("out")
	inputs.SetFilterCtx(buffersinkCtx)
	inputs.SetPadIdx(0)
	inputs.SetNext(nil)

	if len(outputs.GetName()) == 0 || len(inputs.GetName()) == 0 {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	if ret = ffmpeg.AvFilterGraphParsePtr(filterGraph, filterSpec, &inputs, &outputs, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot parse filter graph\n")
		goto end
	}

	if ret = ffmpeg.AvFilterGraphConfig(filterGraph, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot config filter graph\n")
		goto end
	}

	// Fill FilteringContext
	fctx.buffersinkCtx = buffersinkCtx
	fctx.buffersrcCtx = buffersrcCtx
	fctx.filterGraph = filterGraph

end:
	ffmpeg.AvFilterInoutFree(&inputs)
	ffmpeg.AvFilterInoutFree(&outputs)

	return ret
}

func initFilters(ifmtCtx *ffmpeg.AVFormatContext, streamCtx []streamContext) (filterCtx []filteringContext, ret int32) {
	var (
		filterSpec string
	)

	filterCtx = make([]filteringContext, ifmtCtx.GetNbStreams())

	for i := 0; i < int(ifmtCtx.GetNbStreams()); i++ {
		if !(ifmtCtx.GetStreams()[i].GetCodecpar().GetCodecType() == ffmpeg.AVMEDIA_TYPE_AUDIO ||
			ifmtCtx.GetStreams()[i].GetCodecpar().GetCodecType() == ffmpeg.AVMEDIA_TYPE_VIDEO) {
			continue
		}

		if ifmtCtx.GetStreams()[i].GetCodecpar().GetCodecType() == ffmpeg.AVMEDIA_TYPE_VIDEO {
			// passthrough (dummy) filter for video
			filterSpec = "null"
		} else {
			// passthrough (dummy) filter for audio
			filterSpec = "anull"
		}
		if ret = initFilter(&filterCtx[i], streamCtx[i].decCtx, streamCtx[i].encCtx, filterSpec); ret > 0 {
			return nil, ret
		}

		if filterCtx[i].encPkt = ffmpeg.AvPacketAlloc(); filterCtx[i].encPkt == nil {
			return nil, ffmpeg.AVERROR(syscall.ENOMEM)
		}

		if filterCtx[i].filteredFrame = ffmpeg.AvFrameAlloc(); filterCtx[i].filteredFrame == nil {
			return nil, ffmpeg.AVERROR(syscall.ENOMEM)
		}
	}
	return filterCtx, 0
}

func encodeWriteFrame(ofmtCtx *ffmpeg.AVFormatContext, streamCtx []streamContext, filterCtx []filteringContext,
	streamIndex int32, flush int32) (ret int32) {
	var (
		stream    = &streamCtx[streamIndex]
		filter    = &filterCtx[streamIndex]
		filtFrame *ffmpeg.AVFrame
		encPkt    = filter.encPkt
	)
	if flush == 0 {
		filtFrame = filter.filteredFrame
	}

	ffmpeg.AvLog(nil, ffmpeg.AV_LOG_INFO, "Encoding frame\n")
	// encode filtered frame
	ffmpeg.AvPacketUnref(encPkt)

	if ret = ffmpeg.AvCodecSendFrame(stream.encCtx, filtFrame); ret < 0 {
		return ret
	}

	for ret >= 0 {
		ret = ffmpeg.AvCodecReceivePacket(stream.encCtx, encPkt)

		if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
			return 0
		}

		// prepare packet for muxing
		encPkt.SetStreamIndex(streamIndex)
		ffmpeg.AvPacketRescaleTs(encPkt, stream.encCtx.GetTimeBase(),
			ofmtCtx.GetStreams()[streamIndex].GetTimeBase())

		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_DEBUG, "Muxing frame\n")
		// mux encoded frame
		ret = ffmpeg.AvInterleavedWriteFrame(ofmtCtx, encPkt)
	}
	return ret
}

func filterEncodeWriteFrame(ofmtCtx *ffmpeg.AVFormatContext, streamCtx []streamContext, filterCtx []filteringContext,
	frame *ffmpeg.AVFrame, streamIndex int32) (ret int32) {
	var (
		filter = &filterCtx[streamIndex]
	)

	ffmpeg.AvLog(nil, ffmpeg.AV_LOG_INFO, "Pushing decoded frame to filters\n")
	// push the decoded frame into the filtergraph
	if ret = ffmpeg.AvBuffersrcAddFrameFlags(filter.buffersrcCtx, frame, 0); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Error while feeding the filtergraph\n")
		return ret
	}

	// pull filtered frames from the filtergraph
	for {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_INFO, "Pulling filtered frame from filters\n")
		if ret = ffmpeg.AvBuffersinkGetFrame(filter.buffersinkCtx, filter.filteredFrame); ret < 0 {
			// if no more frames for output - returns AVERROR(EAGAIN)
			// if flushed and no more frames for output - returns AVERROR_EOF
			// rewrite retcode to 0 to show it as normal procedure completion
			if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
				ret = 0
			}
			break
		}

		filter.filteredFrame.SetPictType(ffmpeg.AV_PICTURE_TYPE_NONE)
		ret = encodeWriteFrame(ofmtCtx, streamCtx, filterCtx, streamIndex, 0)
		ffmpeg.AvFrameUnref(filter.filteredFrame)
		if ret < 0 {
			break
		}
	}

	return ret
}

func flushEncoder(ofmtCtx *ffmpeg.AVFormatContext, streamCtx []streamContext, filterCtx []filteringContext,
	streamIndex int32) (ret int32) {
	if (streamCtx[streamIndex].encCtx.GetCodec().GetCapabilities() & ffmpeg.AV_CODEC_CAP_DELAY) == 0 {
		return 0
	}
	ffmpeg.AvLog(nil, ffmpeg.AV_LOG_INFO, "Flushing stream #%d encoder\n", streamIndex)
	return encodeWriteFrame(ofmtCtx, streamCtx, filterCtx, streamIndex, 1)
}

func main() {
	var (
		ret         int32
		packet      *ffmpeg.AVPacket
		streamIndex int32
		ifmtCtx     *ffmpeg.AVFormatContext
		ofmtCtx     *ffmpeg.AVFormatContext
		streamCtx   []streamContext
		filterCtx   []filteringContext
	)

	if len(os.Args) != 3 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Usage: %s <input file> <output file>\n", os.Args[0])
		os.Exit(1)
	}

	if ifmtCtx, streamCtx, ret = openInputFile(os.Args[1]); ret < 0 {
		goto end
	}
	if ofmtCtx, ret = openOutputFile(ifmtCtx, streamCtx, os.Args[2]); ret < 0 {
		goto end
	}
	if filterCtx, ret = initFilters(ifmtCtx, streamCtx); ret < 0 {
		goto end
	}
	if packet = ffmpeg.AvPacketAlloc(); packet == nil {
		goto end
	}

	// read all packets
	for {
		if ret = ffmpeg.AvReadFrame(ifmtCtx, packet); ret < 0 {
			break
		}
		streamIndex = packet.GetStreamIndex()
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_DEBUG, "Demuxer gave frame of stream_index %d\n", streamIndex)

		if filterCtx[streamIndex].filterGraph != nil {
			stream := &streamCtx[streamIndex]

			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_DEBUG, "Going to reencode&filter the frame\n")

			ffmpeg.AvPacketRescaleTs(packet,
				ifmtCtx.GetStreams()[streamIndex].GetTimeBase(),
				stream.decCtx.GetTimeBase())
			if ret = ffmpeg.AvCodecSendPacket(stream.decCtx, packet); ret < 0 {
				ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Decoding failed\n")
				break
			}

			for ret >= 0 {
				ret = ffmpeg.AvCodecReceiveFrame(stream.decCtx, stream.decFrame)
				if ret == ffmpeg.AVERROR_EOF || ret == ffmpeg.AVERROR(syscall.EAGAIN) {
					break
				} else if ret < 0 {
					goto end
				}

				stream.decFrame.SetPts(stream.decFrame.GetBestEffortTimestamp())
				if ret = filterEncodeWriteFrame(ofmtCtx, streamCtx, filterCtx,
					stream.decFrame, streamIndex); ret < 0 {
					goto end
				}
			}
		} else {
			// remux this frame without reencoding
			ffmpeg.AvPacketRescaleTs(packet,
				ifmtCtx.GetStreams()[streamIndex].GetTimeBase(),
				ofmtCtx.GetStreams()[streamIndex].GetTimeBase())
			if ret = ffmpeg.AvInterleavedWriteFrame(ofmtCtx, packet); ret < 0 {
				goto end
			}
		}
		ffmpeg.AvPacketUnref(packet)
	}

	// flush filters and encoders
	for i := int32(0); i < int32(ifmtCtx.GetNbStreams()); i++ {
		// flush filter
		if filterCtx[i].filterGraph == nil {
			continue
		}
		if ret = filterEncodeWriteFrame(ofmtCtx, streamCtx, filterCtx, nil, i); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Flushing filter failed\n")
			goto end
		}

		// flush encoder
		if ret = flushEncoder(ofmtCtx, streamCtx, filterCtx, i); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Flushing encoder failed\n")
			goto end
		}
	}

	ffmpeg.AvWriteTrailer(ofmtCtx)
end:
	ffmpeg.AvPacketFree(&packet)
	for i := 0; i < len(streamCtx); i++ {
		ffmpeg.AvCodecFreeContext(&streamCtx[i].decCtx)
		if ofmtCtx != nil && ofmtCtx.GetNbStreams() > uint32(i) &&
			ofmtCtx.GetStreams()[i] != nil && streamCtx[i].encCtx != nil {
			ffmpeg.AvCodecFreeContext(&streamCtx[i].encCtx)
		}
		if filterCtx != nil && filterCtx[i].filterGraph != nil {
			ffmpeg.AvFilterGraphFree(&filterCtx[i].filterGraph)
			ffmpeg.AvPacketFree(&filterCtx[i].encPkt)
			ffmpeg.AvFrameFree(&filterCtx[i].filteredFrame)
		}

		ffmpeg.AvFrameFree(&streamCtx[i].decFrame)
	}
	ffmpeg.AvFormatCloseInput(&ifmtCtx)
	if ofmtCtx != nil && (ofmtCtx.GetOformat().GetFlags()&ffmpeg.AVFMT_NOFILE) == 0 {
		ffmpeg.AvIOClosep(ofmtCtx.GetPbAddr())
	}
	ffmpeg.AvFormatFreeContext(ofmtCtx)

	if ret < 0 && ret != ffmpeg.AVERROR_EOF {
		fmt.Fprintf(os.Stderr, "Error occurred: %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(1)
	}
	os.Exit(0)
}
