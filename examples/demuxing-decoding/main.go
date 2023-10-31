package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

var (
	audioStream      *ffmpeg.AVStream
	videoStream      *ffmpeg.AVStream
	videoFrameCount  int
	audioFrameCount  int
	width, height    int32
	pixFmt           ffmpeg.AVPixelFormat
	audioDstFile     *os.File
	videoDstFile     *os.File
	audioStreamIdx   int32
	videoStreamIdx   int32
	videoDstData     [4]*uint8
	videoDstLinesize [4]int32
	videoDstBufSize  int32
	srcFilename      string
)

func outputVideoFrame(videoDecCtx *ffmpeg.AVCodecContext, frame *ffmpeg.AVFrame) int32 {

	if frame.GetWidth() != width || frame.GetHeight() != height || frame.GetFormat() != pixFmt {
		// To handle this change, one could call av_image_alloc again and
		// decode the following frames into another rawvideo file.
		fmt.Fprintf(os.Stderr, "Error: Width, height and pixel format have to be "+
			"constant in a rawvideo file, but the width, height or "+
			"pixel format of the input video changed:\n"+
			"old: width = %d, height = %d, format = %s\n"+
			"new: width = %d, height = %d, format = %s\n",
			width, height, ffmpeg.AvGetPixFmtName(pixFmt),
			frame.GetWidth(), frame.GetHeight(),
			ffmpeg.AvGetPixFmtName(frame.GetFormat()))
		return -1
	}

	videoFrameCount++
	fmt.Fprintf(os.Stdout, "video_frame n:%d coded_n:%d\n", videoFrameCount, frame.GetCodedPictureNumber())

	// copy decoded frame to destination buffer:
	// this is required since rawvideo expects non aligned data
	ffmpeg.AvImageCopy(videoDstData[:], videoDstLinesize[:], frame.GetData(), frame.GetLinesize(), pixFmt, width, height)

	// write to rawvideo file
	videoDstFile.Write(unsafe.Slice(videoDstData[0], videoDstBufSize))
	return 0
}

func outputAudioFrame(audioDecCtx *ffmpeg.AVCodecContext, frame *ffmpeg.AVFrame) (ret int32) {
	unpaddedLinesize := frame.GetNbSamples() * ffmpeg.AvGetBytesPerSample(frame.GetFormat())
	audioFrameCount++
	fmt.Fprintf(os.Stdout, "audio_frame n:%d nb_samples:%d pts:%s\n",
		audioFrameCount, frame.GetNbSamples(),
		ffmpeg.AvTs2timestr(frame.GetPts(), audioDecCtx.GetPktTimebaseAddr()))

	// Write the raw audio data samples of the first plane. This works
	// fine for packed formats (e.g. AV_SAMPLE_FMT_S16). However,
	// most audio decoders output planar audio, which uses a separate
	// plane of audio samples for each channel (e.g. AV_SAMPLE_FMT_S16P).
	// In other words, this code will write only the first audio channel
	// in these cases.
	// You should use libswresample or libavfilter to convert the frame
	// to packed data.
	audioDstFile.Write(ffmpeg.SliceSlice(frame.GetExtendedData(), 1, unpaddedLinesize)[0])

	return 0
}

func deocdePacket(dec *ffmpeg.AVCodecContext, pkt *ffmpeg.AVPacket, frame *ffmpeg.AVFrame) (ret int32) {

	// submit the packet to the decoder
	if ret = ffmpeg.AvCodecSendPacket(dec, pkt); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error submitting a packet for decoding (%s)\n", ffmpeg.AvErr2str(ret))
		return ret
	}

	// get all the available frames from the decoder
	for ret >= 0 {
		if ret = ffmpeg.AvCodecReceiveFrame(dec, frame); ret < 0 {
			// those two return values are special and mean there is no output
			// frame available, but there were no errors during decoding
			if ret == ffmpeg.AVERROR_EOF || ret == ffmpeg.AVERROR(syscall.EAGAIN) {
				return 0
			}

			fmt.Fprintf(os.Stderr, "Error during decoding (%s)\n", ffmpeg.AvErr2str(ret))
			return ret
		}

		// write the frame data to output file
		if dec.GetCodec().GetType() == ffmpeg.AVMEDIA_TYPE_VIDEO {
			ret = outputVideoFrame(dec, frame)
		} else {
			ret = outputAudioFrame(dec, frame)
		}

		ffmpeg.AvFrameUnref(frame)
		if ret < 0 {
			return ret
		}
	}

	return 0
}

func openCodecContext(fmtCtx *ffmpeg.AVFormatContext, _type ffmpeg.AVMediaType) (
	streamIdx int32, decCtx *ffmpeg.AVCodecContext, ret int32) {
	var (
		opts *ffmpeg.AVDictionary
	)
	if ret = ffmpeg.AvFindBestStream(fmtCtx, _type, -1, -1, nil, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not find %s stream in input file '%s'\n",
			ffmpeg.AvGetMediaTypeString(_type), srcFilename)
		return 0, nil, ret
	} else {
		streamIdx = ret
		st := fmtCtx.GetStreams()[streamIdx]

		// find decoder for the stream
		dec := ffmpeg.AvCodecFindDecoder(st.GetCodecpar().GetCodecId())
		if dec == nil {
			fmt.Fprintf(os.Stderr, "Failed to find %s codec\n",
				ffmpeg.AvGetMediaTypeString(_type))
			return 0, nil, ffmpeg.AVERROR(syscall.EINVAL)
		}

		// Allocate a codec context for the decoder
		if decCtx = ffmpeg.AvCodecAllocContext3(dec); decCtx == nil {
			fmt.Fprintf(os.Stderr, "Failed to allocate the %s codec context\n",
				ffmpeg.AvGetMediaTypeString(_type))
			return 0, nil, ffmpeg.AVERROR(syscall.ENOMEM)
		}

		// Copy codec parameters from input stream to output codec context
		if ret = ffmpeg.AvCodecParametersToContext(decCtx, st.GetCodecpar()); ret < 0 {
			fmt.Fprintf(os.Stderr, "Failed to copy %s codec parameters to decoder context\n",
				ffmpeg.AvGetMediaTypeString(_type))
			return 0, nil, ret
		}

		// Init the decoders
		if ret = ffmpeg.AvCodecOpen2(decCtx, dec, &opts); ret < 0 {
			fmt.Fprintf(os.Stderr, "Failed to open %s codec\n",
				ffmpeg.AvGetMediaTypeString(_type))
			return 0, nil, ret
		}
	}

	return streamIdx, decCtx, 0
}

func getFormatFromSampleFmt(sampleFmt ffmpeg.AVSampleFormat) (string, int32) {
	sampleFmtEntry := []struct {
		sampleFmt ffmpeg.AVSampleFormat
		fmtBe     string
		fmtLe     string
	}{
		{ffmpeg.AV_SAMPLE_FMT_U8, "u8", "u8"},
		{ffmpeg.AV_SAMPLE_FMT_S16, "s16be", "s16le"},
		{ffmpeg.AV_SAMPLE_FMT_S32, "s32be", "s32le"},
		{ffmpeg.AV_SAMPLE_FMT_FLT, "f32be", "f32le"},
		{ffmpeg.AV_SAMPLE_FMT_DBL, "f64be", "f64le"},
	}

	for _, entry := range sampleFmtEntry {
		if sampleFmt == entry.sampleFmt {
			return ffmpeg.AV_NE(entry.fmtBe, entry.fmtLe), 0
		}
	}

	fmt.Fprintf(os.Stderr, "sample format %s is not supported as output format\n",
		ffmpeg.AvGetSampleFmtName(sampleFmt))
	return ffmpeg.NIL, -1
}

func main() {
	var (
		ret         int32
		err         error
		fmtCtx      *ffmpeg.AVFormatContext
		audioDecCtx *ffmpeg.AVCodecContext
		videoDecCtx *ffmpeg.AVCodecContext
		frame       *ffmpeg.AVFrame
		pkt         *ffmpeg.AVPacket
	)
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "usage: %s  input_file video_output_file audio_output_file\n"+
			"API example program to show how to read frames from an input file.\n"+
			"This program reads frames from a file, decodes them, and writes decoded\n"+
			"video frames to a rawvideo file named video_output_file, and decoded\n"+
			"audio frames to a rawaudio file named audio_output_file.\n", os.Args[0])
		os.Exit(1)
	}

	srcFilename = os.Args[1]
	videoDstFilename := os.Args[2]
	audioDstFilename := os.Args[3]

	// open input file, and allocate format context
	if ret = ffmpeg.AvFormatOpenInput(&fmtCtx, srcFilename, nil, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open source file %s\n", srcFilename)
		os.Exit(1)
	}

	// retrieve stream information
	if videoStreamIdx, videoDecCtx, ret = openCodecContext(fmtCtx, ffmpeg.AVMEDIA_TYPE_VIDEO); ret >= 0 {
		videoStream = fmtCtx.GetStreams()[videoStreamIdx]

		videoDstFile, err = os.OpenFile(videoDstFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open destination file %s\n", videoDstFilename)
			ret = 1
			goto end
		}

		// allocate image where the decoded image will be put
		width = videoDecCtx.GetWidth()
		height = videoDecCtx.GetHeight()
		pixFmt = videoDecCtx.GetPixFmt()
		if pixFmt == ffmpeg.AV_PIX_FMT_NONE {
			pixFmt = ffmpeg.AV_PIX_FMT_YUV420P
		}
		if ret = ffmpeg.AvImageAlloc(videoDstData[:], videoDstLinesize[:],
			width, height, pixFmt, 1); ret < 0 {
			fmt.Fprintf(os.Stderr, "Could not allocate raw video buffer (%s)\n", ffmpeg.AvErr2str(ret))
			goto end
		}
		videoDstBufSize = ret
	}

	if audioStreamIdx, audioDecCtx, ret = openCodecContext(fmtCtx, ffmpeg.AVMEDIA_TYPE_AUDIO); ret >= 0 {
		audioStream = fmtCtx.GetStreams()[audioStreamIdx]

		audioDstFile, err = os.OpenFile(audioDstFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open destination file %s\n", audioDstFilename)
			ret = 1
			goto end
		}
	}

	// dump input information to stderr
	ffmpeg.AvDumpFormat(fmtCtx, 0, srcFilename, 0)

	if audioStream == nil && videoStream == nil {
		fmt.Fprintf(os.Stderr, "Could not find audio or video stream in the input, aborting\n")
		ret = 1
		goto end
	}

	if frame = ffmpeg.AvFrameAlloc(); frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate frame\n")
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	if pkt = ffmpeg.AvPacketAlloc(); pkt == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate packet\n")
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	if videoStream != nil {
		fmt.Fprintf(os.Stdout, "Demuxing video from file '%s' into '%s'\n", srcFilename, videoDstFilename)
	}
	if audioStream != nil {
		fmt.Fprintf(os.Stdout, "Demuxing audio from file '%s' into '%s'\n", srcFilename, audioDstFilename)
	}

	// read frames from the file
	for ffmpeg.AvReadFrame(fmtCtx, pkt) >= 0 {
		// check if the packet belongs to a stream we are interested in, otherwise
		// skip it
		if pkt.GetStreamIndex() == videoStreamIdx {
			ret = deocdePacket(videoDecCtx, pkt, frame)
		} else if pkt.GetStreamIndex() == audioStreamIdx {
			ret = deocdePacket(audioDecCtx, pkt, frame)
		}
		ffmpeg.AvPacketUnref(pkt)
		if ret < 0 {
			break
		}
	}

	// flush the decoders
	if videoDecCtx != nil {
		deocdePacket(videoDecCtx, nil, frame)
	}
	if audioDecCtx != nil {
		deocdePacket(audioDecCtx, nil, frame)
	}

	fmt.Fprintf(os.Stdout, "Demuxing succeeded.\n")

	if videoStream != nil {
		fmt.Fprintf(os.Stdout, "Play the output video file with the command:\n"+
			"ffplay -f rawvideo -pixel_format %s -video_size %dx%d %s\n",
			ffmpeg.AvGetPixFmtName(videoDecCtx.GetPixFmt()),
			videoDecCtx.GetWidth(), videoDecCtx.GetHeight(), videoDstFilename)
	}

	if audioStream != nil {
		sfmt := audioDecCtx.GetSampleFmt()
		nChannels := audioDecCtx.GetChannels()
		var _fmt string

		if ffmpeg.AvSampleFmtIsPlanar(sfmt) != 0 {
			packed := ffmpeg.AvGetSampleFmtName(sfmt)
			if len(packed) == 0 {
				packed = "?"
			}
			fmt.Fprintf(os.Stdout, "Warning: the sample format the decoder produced is planar "+
				"(%s). This example will output the first channel only.\n", packed)
			sfmt = ffmpeg.AvGetPackedSampleFmt(sfmt)
			nChannels = 1
		}

		if _fmt, ret = getFormatFromSampleFmt(sfmt); ret < 0 {
			goto end
		}

		fmt.Fprintf(os.Stdout, "Play the output audio file with the command:\n"+
			"ffplay -f %s -ac %d -ar %d %s\n",
			_fmt, nChannels, audioDecCtx.GetSampleRate(), audioDstFilename)
	}

end:
	ffmpeg.AvCodecFreeContext(&videoDecCtx)
	ffmpeg.AvCodecFreeContext(&audioDecCtx)
	ffmpeg.AvFormatCloseInput(&fmtCtx)
	if videoDstFile != nil {
		videoDstFile.Close()
	}
	if audioDstFile != nil {
		audioDstFile.Close()
	}
	ffmpeg.AvPacketFree(&pkt)
	ffmpeg.AvFrameFree(&frame)
	ffmpeg.AvFree(videoDstData[0])
	if ret < 0 {
		os.Exit(int(ret))
	}
}
