package main

import (
	"fmt"
	"math"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

const (
	STREAM_DURATION   = 10.0
	STREAM_FRAME_RATE = 25                        // 25 images/s
	STREAM_PIX_FMT    = ffmpeg.AV_PIX_FMT_YUV420P //default pix_fmt
	SCALE_FLAGS       = ffmpeg.SWS_BICUBIC
)

type outputStream struct {
	st  *ffmpeg.AVStream
	enc *ffmpeg.AVCodecContext

	// pts of the next frame that will be generated
	nextPts      int64
	samplesCount int32
	frame        *ffmpeg.AVFrame
	tmpFrame     *ffmpeg.AVFrame

	t, tincr, tincr2 float32

	swsCtx *ffmpeg.SwsContext
	swrCtx *ffmpeg.SwrContext
}

func logPacket(fmtCtx *ffmpeg.AVFormatContext, pkt *ffmpeg.AVPacket) {
	timeBase := fmtCtx.GetStreams()[pkt.GetStreamIndex()].GetTimeBaseAddr()
	fmt.Fprintf(os.Stdout, "pts:%s pts_time:%s dts:%s dts_time:%s duration:%s duration_time:%s stream_index:%d\n",
		ffmpeg.AvTs2str(pkt.GetPts()), ffmpeg.AvTs2timestr(pkt.GetPts(), timeBase),
		ffmpeg.AvTs2str(pkt.GetDts()), ffmpeg.AvTs2timestr(pkt.GetDts(), timeBase),
		ffmpeg.AvTs2str(pkt.GetDuration()), ffmpeg.AvTs2timestr(pkt.GetDuration(), timeBase),
		pkt.GetStreamIndex())
}

func writeFrame(fmtCtx *ffmpeg.AVFormatContext, c *ffmpeg.AVCodecContext, st *ffmpeg.AVStream, frame *ffmpeg.AVFrame) bool {
	var (
		ret int32
	)
	// send the frame to the encoder
	if ret = ffmpeg.AvCodecSendFrame(c, frame); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error sending a frame to the encoder: %s\n",
			ffmpeg.AvErr2str(ret))
		os.Exit(1)
	}

	for ret >= 0 {
		var pkt ffmpeg.AVPacket

		ret = ffmpeg.AvCodecReceivePacket(c, &pkt)
		if ret == ffmpeg.AVERROR(syscall.EAGAIN) || ret == ffmpeg.AVERROR_EOF {
			break
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error encoding a frame: %s\n", ffmpeg.AvErr2str(ret))
			os.Exit(1)
		}

		// rescale output packet timestamp values from codec to stream timebase
		ffmpeg.AvPacketRescaleTs(&pkt, c.GetTimeBase(), st.GetTimeBase())
		pkt.SetStreamIndex(st.GetIndex())

		// Write the compressed frame to the media file.
		logPacket(fmtCtx, &pkt)
		ret = ffmpeg.AvInterleavedWriteFrame(fmtCtx, &pkt)
		ffmpeg.AvPacketUnref(&pkt)
		if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while writing output packet: %s\n", ffmpeg.AvErr2str(ret))
			os.Exit(1)
		}
	}

	return ret == ffmpeg.AVERROR_EOF
}

// Add an output stream.
func addStream(ost *outputStream, oc *ffmpeg.AVFormatContext, codecId ffmpeg.AVCodecID) (codec *ffmpeg.AVCodec) {
	var (
		c *ffmpeg.AVCodecContext
	)
	if codec = ffmpeg.AvCodecFindEncoder(codecId); codec == nil {
		fmt.Fprintf(os.Stderr, "Could not find encoder for '%s'\n", ffmpeg.AvCodecGetName(codecId))
		os.Exit(1)
	}

	if ost.st = ffmpeg.AvFormatNewStream(oc, nil); ost.st == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate stream\n")
		os.Exit(1)
	}
	ost.st.SetId(int32(oc.GetNbStreams() - 1))
	if c = ffmpeg.AvCodecAllocContext3(codec); c == nil {
		fmt.Fprintf(os.Stderr, "Could not alloc an encoding context\n")
		os.Exit(1)
	}
	ost.enc = c

	switch codec.GetType() {
	case ffmpeg.AVMEDIA_TYPE_AUDIO:
		c.SetSampleFmt(ffmpeg.CondExpr(len(codec.GetSampleFmts()) > 0,
			codec.GetSampleFmts()[0], ffmpeg.AV_SAMPLE_FMT_FLTP))
		c.SetBitRate(64_000)
		c.SetSampleRate(44_100)
		if len(codec.GetSupportedSamplerates()) > 0 {
			c.SetSampleRate(codec.GetSupportedSamplerates()[0])
			for _, sr := range codec.GetSupportedSamplerates() {
				if sr == 44_100 {
					c.SetSampleRate(44_100)
				}
			}
		}
		c.SetChannels(ffmpeg.AvGetChannelLayoutNbChannels(c.GetChannelLayout()))
		c.SetChannelLayout(ffmpeg.AV_CH_LAYOUT_STEREO)
		if len(codec.GetChannelLayouts()) > 0 {
			c.SetChannelLayout(codec.GetChannelLayouts()[0])
			for _, cl := range codec.GetChannelLayouts() {
				if cl == ffmpeg.AV_CH_LAYOUT_STEREO {
					c.SetChannelLayout(ffmpeg.AV_CH_LAYOUT_STEREO)
				}
			}
		}
		c.SetChannels(ffmpeg.AvGetChannelLayoutNbChannels(c.GetChannelLayout()))
		ost.st.SetTimeBase(ffmpeg.AvMakeQ(1, c.GetSampleRate()))

	case ffmpeg.AVMEDIA_TYPE_VIDEO:
		c.SetCodecId(codecId)

		c.SetBitRate(4_000_000)
		// Resolution must be a multiple of two.
		c.SetWidth(352)
		c.SetHeight(288)
		// timebase: This is the fundamental unit of time (in seconds) in terms
		// of which frame timestamps are represented. For fixed-fps content,
		// timebase should be 1/framerate and timestamp increments should be
		// identical to 1.
		ost.st.SetTimeBase(ffmpeg.AvMakeQ(1, STREAM_FRAME_RATE))
		c.SetTimeBase(ost.st.GetTimeBase())

		c.SetGopSize(12) // emit one intra frame every twelve frames at most
		c.SetPixFmt(STREAM_PIX_FMT)
		if c.GetCodecId() == ffmpeg.AV_CODEC_ID_MPEG2VIDEO {
			// just for testing, we also add B-frames
			c.SetMaxBFrames(2)
		}
		if c.GetCodecId() == ffmpeg.AV_CODEC_ID_MPEG1VIDEO {
			// Needed to avoid using macroblocks in which some coeffs overflow.
			// This does not happen with normal video, it just happens here as
			// the motion of the chroma plane does not match the luma plane.
			c.SetMbDecision(2)
		}

	default:
		break
	}

	if (oc.GetOformat().GetFlags() & ffmpeg.AVFMT_GLOBALHEADER) != 0 {
		c.SetFlags(c.GetFlags() | ffmpeg.AV_CODEC_FLAG_GLOBAL_HEADER)
	}

	return codec
}

// **************************************************************
// audio output

func allocAudioFrame(sampleFmt ffmpeg.AVSampleFormat, channelLayout uint64, sampleRate int32, nbSamples int32) (
	frame *ffmpeg.AVFrame) {
	if frame = ffmpeg.AvFrameAlloc(); frame == nil {
		panic("Error allocating an audio frame")
	}

	frame.SetFormat(sampleFmt)
	frame.SetChannelLayout(channelLayout)
	frame.SetSampleRate(sampleRate)
	frame.SetNbSamples(nbSamples)

	if nbSamples != 0 {
		if ret := ffmpeg.AvFrameGetBuffer(frame, 0); ret < 0 {
			panic("Error allocating an audio buffer")
		}
	}
	return frame
}

func openAudio(oc *ffmpeg.AVFormatContext, codec *ffmpeg.AVCodec, ost *outputStream, optArg *ffmpeg.AVDictionary) {
	var (
		c         *ffmpeg.AVCodecContext
		opt       *ffmpeg.AVDictionary
		nbSamples int32
	)
	c = ost.enc

	// open it
	ffmpeg.AvDictCopy(&opt, optArg, 0)
	ret := ffmpeg.AvCodecOpen2(c, codec, &opt)
	ffmpeg.AvDictFree(&opt)
	if ret < 0 {
		panic(fmt.Sprintf("Could not open audio codec: %s", ffmpeg.AvErr2str(ret)))
	}

	// init signal generator
	ost.t = 0
	ost.tincr = float32(2 * math.Pi * 110.0 / float64(c.GetSampleRate()))
	// increment frequency by 110 Hz per second
	ost.tincr2 = float32(2 * math.Pi * 110.0 / float64(c.GetSampleRate()) / float64(c.GetSampleRate()))

	if (c.GetCodec().GetCapabilities() & ffmpeg.AV_CODEC_CAP_VARIABLE_FRAME_SIZE) != 0 {
		nbSamples = 10_000
	} else {
		nbSamples = c.GetFrameSize()
	}
	ost.frame = allocAudioFrame(c.GetSampleFmt(), c.GetChannelLayout(),
		c.GetSampleRate(), nbSamples)
	ost.tmpFrame = allocAudioFrame(ffmpeg.AV_SAMPLE_FMT_S16, c.GetChannelLayout(),
		c.GetSampleRate(), nbSamples)

	// copy the stream parameters to the muxer
	if ret = ffmpeg.AvCodecParametersFromContext(ost.st.GetCodecpar(), c); ret < 0 {
		panic("Could not copy the stream parameters")
	}

	// create resampler context
	if ost.swrCtx = ffmpeg.SwrAlloc(); ost.swrCtx == nil {
		panic("Could not allocate resampler context")
	}

	// set options
	ffmpeg.AvOptSetInt(ost.swrCtx, "in_channel_count", c.GetChannels(), 0)
	ffmpeg.AvOptSetInt(ost.swrCtx, "in_sample_rate", c.GetSampleRate(), 0)
	ffmpeg.AvOptSetSampleFmt(ost.swrCtx, "in_sample_fmt", ffmpeg.AV_SAMPLE_FMT_S16, 0)
	ffmpeg.AvOptSetInt(ost.swrCtx, "out_channel_count", c.GetChannels(), 0)
	ffmpeg.AvOptSetInt(ost.swrCtx, "out_sample_rate", c.GetSampleRate(), 0)
	ffmpeg.AvOptSetSampleFmt(ost.swrCtx, "out_sample_fmt", c.GetSampleFmt(), 0)

	// initialize the resampling context
	if ret = ffmpeg.SwrInit(ost.swrCtx); ret < 0 {
		panic("Failed to initialize the resampling context")
	}
}

// Prepare a 16 bit dummy audio frame of 'frame_size' samples and
// 'nb_channels' channels.
func getAudioFrame(ost *outputStream) (frame *ffmpeg.AVFrame) {
	frame = ost.tmpFrame
	data := unsafe.Slice((*int16)(unsafe.Pointer(frame.GetData()[0])),
		frame.GetNbSamples()*ost.enc.GetChannels())

	if ffmpeg.AvCompareTs(ost.nextPts, ost.enc.GetTimeBase(), STREAM_DURATION, ffmpeg.AvMakeQ(1, 1)) > 0 {
		return nil
	}
	idx := 0
	for j := 0; j < int(frame.GetNbSamples()); j++ {
		v := (int16)(math.Sin(float64(ost.t)) * 10_000)
		for i := 0; i < int(ost.enc.GetChannels()); i++ {
			data[idx] = v
			idx++
			ost.t += ost.tincr
			ost.tincr += ost.tincr2
		}
	}
	frame.SetPts(ost.nextPts)
	ost.nextPts += int64(frame.GetNbSamples())

	return frame
}

// encode one audio frame and send it to the muxer
// return 1 when encoding is finished, 0 otherwise
func writeAudioFrame(oc *ffmpeg.AVFormatContext, ost *outputStream) bool {
	var (
		c            *ffmpeg.AVCodecContext
		frame        *ffmpeg.AVFrame
		ret          int32
		dstNbSamples int32
	)
	c = ost.enc

	frame = getAudioFrame(ost)

	if frame != nil {
		// convert samples from native format to destination codec format, using the resampler */
		// compute destination number of samples
		dstNbSamples = ffmpeg.AvRescaleRnd(ffmpeg.SwrGetDelay(ost.swrCtx, c.GetSampleRate())+frame.GetNbSamples(),
			c.GetSampleRate(), c.GetSampleRate(), ffmpeg.AV_ROUND_UP)
		ffmpeg.AvAssert0(dstNbSamples == frame.GetNbSamples())

		// when we pass a frame to the encoder, it may keep a reference to it
		// internally;
		// make sure we do not overwrite it here
		if ret = ffmpeg.AvFrameMakeWritable(ost.frame); ret < 0 {
			panic("Make frame writeable failed")
		}

		if ret = ffmpeg.SwrConvert(ost.swrCtx,
			&ost.frame.GetData()[0], dstNbSamples,
			&frame.GetData()[0], frame.GetNbSamples()); ret < 0 {
			panic("Error while converting")
		}
		frame = ost.frame

		frame.SetPts(ffmpeg.AvRescaleQ(int64(ost.samplesCount), ffmpeg.AvMakeQ(1, c.GetSampleRate()), c.GetTimeBase()))
		ost.samplesCount += dstNbSamples
	}

	return writeFrame(oc, c, ost.st, frame)
}

// **************************************************************
// video output

func allocPicture(pixFmt ffmpeg.AVPixelFormat, width, height int32) (picture *ffmpeg.AVFrame) {
	if picture = ffmpeg.AvFrameAlloc(); picture == nil {
		return nil
	}

	picture.SetFormat(pixFmt)
	picture.SetWidth(width)
	picture.SetHeight(height)

	// allocate the buffers for the frame data
	if ret := ffmpeg.AvFrameGetBuffer(picture, 0); ret < 0 {
		panic("Could not allocate frame data.")
	}
	return picture
}

func openVideo(oc *ffmpeg.AVFormatContext, codec *ffmpeg.AVCodec,
	ost *outputStream, optArg *ffmpeg.AVDictionary) {
	var (
		ret int32
		c   = ost.enc
		opt *ffmpeg.AVDictionary
	)

	ffmpeg.AvDictCopy(&opt, optArg, 0)

	// open the codec
	ret = ffmpeg.AvCodecOpen2(c, codec, &opt)
	ffmpeg.AvDictFree(&opt)
	if ret < 0 {
		panic(fmt.Sprintf("Could not video audio codec: %s", ffmpeg.AvErr2str(ret)))
	}

	// allocate and init a re-usable frame
	if ost.frame = allocPicture(c.GetPixFmt(), c.GetWidth(), c.GetHeight()); ost.frame == nil {
		panic("Could not allocate video frame")
	}

	// If the output format is not YUV420P, then a temporary YUV420P
	// picture is needed too. It is then converted to the required
	// output format.
	ost.tmpFrame = nil
	if c.GetPixFmt() != ffmpeg.AV_PIX_FMT_YUV420P {
		ost.tmpFrame = allocPicture(ffmpeg.AV_PIX_FMT_YUV420P, c.GetWidth(), c.GetHeight())
		if ost.tmpFrame == nil {
			panic("Could not allocate temporary picture")
		}
	}

	// copy the stream parameters to the muxer
	if ret = ffmpeg.AvCodecParametersFromContext(ost.st.GetCodecpar(), c); ret < 0 {
		panic("Could not copy the stream parameters")
	}
}

// Prepare a dummy image.
func fillYuvImage(pict *ffmpeg.AVFrame, frameIndex int32, width, height int32) {
	var (
		data     = ffmpeg.SliceSlice(&pict.GetData()[0], 3, width*height)
		linesize = pict.GetLinesize()
		i        = frameIndex
	)

	// Y
	for y := int32(0); y < height; y++ {
		for x := int32(0); x < width; x++ {
			data[0][y*linesize[0]+x] = uint8(x + y + i*3)
		}
	}
	// Cb and Cr
	for y := int32(0); y < height/2; y++ {
		for x := int32(0); x < width/2; x++ {
			data[1][y*linesize[1]+x] = uint8(128 + y + i*2)
			data[2][y*linesize[2]+x] = uint8(64 + x + i*5)
		}
	}
}

func getVideoFrame(ost *outputStream) *ffmpeg.AVFrame {
	c := ost.enc

	// check if we want to generate more frames
	if ffmpeg.AvCompareTs(ost.nextPts, c.GetTimeBase(), STREAM_DURATION, ffmpeg.AvMakeQ(1, 1)) > 0 {
		return nil
	}

	//  when we pass a frame to the encoder, it may keep a reference to it
	// internally; make sure we do not overwrite it here
	if ffmpeg.AvFrameMakeWritable(ost.frame) < 0 {
		panic("Make video frame writable failed")
	}

	if c.GetPixFmt() != ffmpeg.AV_PIX_FMT_YUV420P {
		// as we only generate a YUV420P picture, we must convert it
		// to the codec pixel format if needed
		if ost.swsCtx == nil {
			ost.swsCtx = ffmpeg.SwsGetContext(c.GetWidth(), c.GetHeight(),
				ffmpeg.AV_PIX_FMT_YUV420P,
				c.GetWidth(), c.GetHeight(),
				c.GetPixFmt(),
				SCALE_FLAGS, nil, nil, nil)
			if ost.swsCtx == nil {
				panic("Could not initialize the conversion context")
			}
		}
		fillYuvImage(ost.tmpFrame, int32(ost.nextPts), c.GetWidth(), c.GetHeight())
		ffmpeg.SwsScale(ost.swsCtx, ost.tmpFrame.GetData(),
			ost.tmpFrame.GetLinesize(), 0, c.GetHeight(), ost.frame.GetData(),
			ost.frame.GetLinesize())
	} else {
		fillYuvImage(ost.frame, int32(ost.nextPts), c.GetWidth(), c.GetHeight())
	}

	ost.frame.SetPts(ffmpeg.PlusPlus(&ost.nextPts))

	return ost.frame
}

// encode one video frame and send it to the muxer
// return 1 when encoding is finished, 0 otherwise
func writeVideoFrame(oc *ffmpeg.AVFormatContext, ost *outputStream) bool {
	return writeFrame(oc, ost.enc, ost.st, getVideoFrame(ost))
}

func closeStream(oc *ffmpeg.AVFormatContext, ost *outputStream) {
	ffmpeg.AvCodecFreeContext(&ost.enc)
	ffmpeg.AvFrameFree(&ost.frame)
	ffmpeg.AvFrameFree(&ost.tmpFrame)
	ffmpeg.SwsFreeContext(ost.swsCtx)
	ffmpeg.SwrFree(&ost.swrCtx)
}

// **************************************************************
// media file output

func main() {
	var (
		videoSt, audioSt         outputStream
		videoCodec, audioCodec   *ffmpeg.AVCodec
		haveVideo, haveAudio     bool
		encodeVideo, encodeAudio bool
		opt                      *ffmpeg.AVDictionary
		ret                      int32
		_fmt                     *ffmpeg.AVOutputFormat
		oc                       *ffmpeg.AVFormatContext
	)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stdout, "usage: %s output_file\n"+
			"API example program to output a media file with libavformat.\n"+
			"This program generates a synthetic audio and video stream, encodes and\n"+
			"muxes them into a file named output_file.\n"+
			"The output format is automatically guessed according to the file extension.\n"+
			"Raw images can also be output by using '%%d' in the filename.\n"+
			"\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	for i := 2; i+1 < len(os.Args); i += 2 {
		if os.Args[i] == "-flags" || os.Args[i] == "-fflags" {
			ffmpeg.AvDictSet(&opt, os.Args[i][1:], os.Args[i+1], 0)
		}
	}

	// allocate the output media context
	ffmpeg.AvFormatAllocOutputContext2(&oc, nil, ffmpeg.NIL, filename)
	if oc == nil {
		fmt.Fprintf(os.Stderr, "Could not deduce output format from file extension: using MPEG.\n")
		ffmpeg.AvFormatAllocOutputContext2(&oc, nil, "mpeg", filename)
	}
	if oc == nil {
		panic("Allocate the output media context failed")
	}

	_fmt = oc.GetOformat()

	// Add the audio and video streams using the default format codecs
	// and initialize the codecs.
	if _fmt.GetVideoCodec() != ffmpeg.AV_CODEC_ID_NONE {
		videoCodec = addStream(&videoSt, oc, _fmt.GetVideoCodec())
		haveVideo = true
		encodeVideo = true
	}
	if _fmt.GetAudioCodec() != ffmpeg.AV_CODEC_ID_NONE {
		audioCodec = addStream(&audioSt, oc, _fmt.GetAudioCodec())
		haveAudio = true
		encodeAudio = true
	}

	// Now that all the parameters are set, we can open the audio and
	// video codecs and allocate the necessary encode buffers.
	if haveVideo {
		openVideo(oc, videoCodec, &videoSt, opt)
	}
	if haveAudio {
		openAudio(oc, audioCodec, &audioSt, opt)
	}

	ffmpeg.AvDumpFormat(oc, 0, filename, 1)

	// open the output file, if needed
	if (_fmt.GetFlags() & ffmpeg.AVFMT_NOFILE) == 0 {
		if ret = ffmpeg.AvIOOpen(oc.GetPbAddr(), filename, ffmpeg.AVIO_FLAG_WRITE); ret < 0 {
			fmt.Fprintf(os.Stderr, "Could not open '%s': %s\n", filename, ffmpeg.AvErr2str(ret))
			os.Exit(1)
		}
	}

	// Write the stream header, if any.
	if ret = ffmpeg.AvFormatWriteHeader(oc, &opt); ret < 0 {
		fmt.Fprintf(os.Stderr, "Error occurred when opening output file: %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(1)
	}

	for encodeVideo || encodeAudio {
		// select the stream to encode
		if encodeVideo &&
			(!encodeAudio || ffmpeg.AvCompareTs(videoSt.nextPts, videoSt.enc.GetTimeBase(),
				audioSt.nextPts, audioSt.enc.GetTimeBase()) <= 0) {
			encodeVideo = !writeVideoFrame(oc, &videoSt)
		} else {
			encodeAudio = !writeAudioFrame(oc, &audioSt)
		}
	}

	// Write the trailer, if any. The trailer must be written before you
	// close the CodecContexts open when you wrote the header; otherwise
	// AvWriteTrailer() may try to use memory that was freed on
	// AvCodecClose().
	ffmpeg.AvWriteTrailer(oc)

	// Close each codec.
	if haveVideo {
		closeStream(oc, &videoSt)
	}
	if haveAudio {
		closeStream(oc, &audioSt)
	}

	if (_fmt.GetFlags() & ffmpeg.AVFMT_NOFILE) == 0 {
		// Close the output file.
		ffmpeg.AvIOClosep(oc.GetPbAddr())
	}

	// free the stream
	ffmpeg.AvFormatFreeContext(oc)

	os.Exit(0)
}
