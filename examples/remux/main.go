package main

import (
	"fmt"
	"os"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

func logPacket(fmtCtx *ffmpeg.AVFormatContext, pkt *ffmpeg.AVPacket, tag string) {
	timeBase := fmtCtx.GetStreams()[pkt.GetStreamIndex()].GetTimeBaseAddr()
	fmt.Fprintf(os.Stdout, "%s: pts:%s pts_time:%s dts:%s dts_time:%s duration:%s duration_time:%s stream_index:%d\n",
		tag,
		ffmpeg.AvTs2str(pkt.GetPts()), ffmpeg.AvTs2timestr(pkt.GetPts(), timeBase),
		ffmpeg.AvTs2str(pkt.GetDts()), ffmpeg.AvTs2timestr(pkt.GetDts(), timeBase),
		ffmpeg.AvTs2str(pkt.GetDuration()), ffmpeg.AvTs2timestr(pkt.GetDuration(), timeBase),
		pkt.GetStreamIndex())
}

func main() {
	var (
		ofmt             *ffmpeg.AVOutputFormat
		ifmtCtx, ofmtCtx *ffmpeg.AVFormatContext
		pkt              ffmpeg.AVPacket
		ret              int32
		streamMapping    []int32
		streamIndex      int32
	)

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s input output\n"+
			"API example program to remux a media file with libavformat and libavcodec.\n"+
			"The output format is guessed according to the file extension.\n"+
			"\n", os.Args[0])
		os.Exit(1)
	}

	inFilename := os.Args[1]
	outFilename := os.Args[2]

	if ret = ffmpeg.AvFormatOpenInput(&ifmtCtx, inFilename, nil, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open input file '%s'", inFilename)
		goto end
	}
	if ret = ffmpeg.AvFormatFindStreamInfo(ifmtCtx, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to retrieve input stream information")
		goto end
	}

	ffmpeg.AvDumpFormat(ifmtCtx, 0, inFilename, 0)

	if ffmpeg.AvFormatAllocOutputContext2(&ofmtCtx, nil, ffmpeg.NIL, outFilename); ofmtCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not create output context\n")
		ret = ffmpeg.AVERROR_UNKNOWN
		goto end
	}

	streamMapping = make([]int32, ifmtCtx.GetNbStreams())

	ofmt = ofmtCtx.GetOformat()

	for i := 0; i < len(streamMapping); i++ {
		inStream := ifmtCtx.GetStreams()[i]
		inCodecPar := inStream.GetCodecpar()

		if inCodecPar.GetCodecType() != ffmpeg.AVMEDIA_TYPE_AUDIO &&
			inCodecPar.GetCodecType() != ffmpeg.AVMEDIA_TYPE_VIDEO &&
			inCodecPar.GetCodecType() != ffmpeg.AVMEDIA_TYPE_SUBTITLE {
			streamMapping[i] = -1
			continue
		}

		streamMapping[i] = streamIndex
		streamIndex++

		outStream := ffmpeg.AvFormatNewStream(ofmtCtx, nil)
		if outStream == nil {
			fmt.Fprintf(os.Stderr, "Failed allocating output stream\n")
			ret = ffmpeg.AVERROR_UNKNOWN
			goto end
		}

		if ret = ffmpeg.AvCodecParametersCopy(outStream.GetCodecpar(), inCodecPar); ret < 0 {
			fmt.Fprintf(os.Stderr, "Failed to copy codec parameters\n")
			goto end
		}
		outStream.GetCodecpar().SetCodecTag(0)
	}
	ffmpeg.AvDumpFormat(ofmtCtx, 0, outFilename, 1)

	if (ofmt.GetFlags() & ffmpeg.AVFMT_NOFILE) == 0 {
		if ret = ffmpeg.AvIOOpen(ofmtCtx.GetPbAddr(), outFilename, ffmpeg.AVIO_FLAG_WRITE); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Could not open output file '%s'", outFilename)
			goto end
		}
	}

	if ret = ffmpeg.AvFormatWriteHeader(ofmtCtx, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error occurred when opening output file\n")
		goto end
	}

	for {
		var inStream, outStream *ffmpeg.AVStream

		if ret = ffmpeg.AvReadFrame(ifmtCtx, &pkt); ret < 0 {
			break
		}

		inStream = ifmtCtx.GetStreams()[pkt.GetStreamIndex()]
		if int(pkt.GetStreamIndex()) >= len(streamMapping) ||
			streamMapping[pkt.GetStreamIndex()] < 0 {
			ffmpeg.AvPacketUnref(&pkt)
			continue
		}

		pkt.SetStreamIndex(streamMapping[pkt.GetStreamIndex()])
		outStream = ofmtCtx.GetStreams()[pkt.GetStreamIndex()]
		logPacket(ifmtCtx, &pkt, "in")

		// copy packet
		pkt.SetPts(ffmpeg.AvRescaleQRnd(pkt.GetPts(), inStream.GetTimeBase(), outStream.GetTimeBase(),
			ffmpeg.AV_ROUND_NEAR_INF|ffmpeg.AV_ROUND_PASS_MINMAX))
		pkt.SetDts(ffmpeg.AvRescaleQRnd(pkt.GetDts(), inStream.GetTimeBase(), outStream.GetTimeBase(),
			ffmpeg.AV_ROUND_NEAR_INF|ffmpeg.AV_ROUND_PASS_MINMAX))
		pkt.SetDuration(ffmpeg.AvRescaleQ(pkt.GetDuration(), inStream.GetTimeBase(), outStream.GetTimeBase()))
		pkt.SetPos(-1)
		logPacket(ofmtCtx, &pkt, "out")

		if ret = ffmpeg.AvInterleavedWriteFrame(ofmtCtx, &pkt); ret < 0 {
			fmt.Fprintf(os.Stderr, "Error muxing packet\n")
			break
		}
		ffmpeg.AvPacketUnref(&pkt)
	}
	ffmpeg.AvWriteTrailer(ofmtCtx)
end:
	// close output
	ffmpeg.AvFormatCloseInput(&ifmtCtx)
	if ofmtCtx != nil && (ofmt.GetFlags()&ffmpeg.AVFMT_NOFILE) == 0 {
		ffmpeg.AvIOClosep(ofmtCtx.GetPbAddr())
	}
	ffmpeg.AvFormatFreeContext(ofmtCtx)

	if ret < 0 && ret != ffmpeg.AVERROR_EOF {
		fmt.Fprintf(os.Stderr, "Error occurred: %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(1)
	}
	os.Exit(0)
}
