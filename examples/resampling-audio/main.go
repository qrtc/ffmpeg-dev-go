package main

import (
	"fmt"
	"math"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

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

// Fill dst buffer with nb_samples, generated starting from t.
func fileSamples(dst []float64, nbSamples, nbChannels, sampleRate int32, t *float64) {
	var (
		tincr = 1.0 / float64(sampleRate)
		c     = 2 * math.Pi * 440.0
	)

	// generate sin tone with 440Hz frequency and duplicated channels
	for i := int32(0); i < nbSamples; i++ {
		dst[i*nbChannels] = math.Sin(c * (*t))
		for j := int32(1); j < nbChannels; j++ {
			dst[i*nbChannels+j] = dst[i*nbChannels]
		}
		*t += tincr
	}
}

func main() {
	var (
		srcRate, dstRate             int32 = 48000, 44100
		srcData, dstData             **uint8
		srcNbChannels, dstNbChannels int32
		srcLinesize, dstLinesize     int32
		srcNbSamples, dstNbSamples   int32 = 1024, 0
		maxDstNbSamples              int32
		srcChLayout, dstChLayout     = ffmpeg.AV_CH_LAYOUT_STEREO, ffmpeg.AV_CH_LAYOUT_SURROUND
		srcSampleFmt, dstSampleFmt   = ffmpeg.AV_SAMPLE_FMT_DBL, ffmpeg.AV_SAMPLE_FMT_S16
		swrCtx                       *ffmpeg.SwrContext
		ret                          int32
		t                            float64
		dstBufsize                   int32
		_fmt                         string
	)

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Usage: %s output_file\n"+
			"API example program to show how to resample an audio stream with libswresample.\n"+
			"This program generates a series of audio frames, resamples them to a specified "+
			"output format and rate and saves them to an output file named output_file.\n",
			os.Args[0])
		os.Exit(1)
	}
	dstFilename := os.Args[1]

	dstFile, err := os.OpenFile(dstFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open destination file %s\n", dstFilename)
		os.Exit(1)
	}

	// create resampler context
	if swrCtx = ffmpeg.SwrAlloc(); swrCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate resampler context\n")
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	// set options
	ffmpeg.AvOptSetInt(swrCtx, "in_channel_layout", srcChLayout, 0)
	ffmpeg.AvOptSetInt(swrCtx, "in_sample_rate", srcRate, 0)
	ffmpeg.AvOptSetSampleFmt(swrCtx, "in_sample_fmt", srcSampleFmt, 0)

	ffmpeg.AvOptSetInt(swrCtx, "out_channel_layout", dstChLayout, 0)
	ffmpeg.AvOptSetInt(swrCtx, "out_sample_rate", dstRate, 0)
	ffmpeg.AvOptSetSampleFmt(swrCtx, "out_sample_fmt", dstSampleFmt, 0)

	// initialize the resampling context
	if ret = ffmpeg.SwrInit(swrCtx); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to initialize the resampling context\n")
		goto end
	}

	// allocate source and destination samples buffers
	srcNbChannels = ffmpeg.AvGetChannelLayoutNbChannels(srcChLayout)
	if ret = ffmpeg.AvSamplesAllocArrayAndSamples(&srcData, &srcLinesize, srcNbChannels,
		srcNbSamples, srcSampleFmt, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate source samples\n")
		goto end
	}

	// compute the number of converted samples: buffering is avoided
	// ensuring that the output buffer will contain at least all the
	// converted input samples
	dstNbSamples = ffmpeg.AvRescaleRnd(srcNbSamples, dstRate, srcRate, ffmpeg.AV_ROUND_UP)
	maxDstNbSamples = dstNbSamples

	// buffer is going to be directly written to a rawaudio file, no alignment
	dstNbChannels = ffmpeg.AvGetChannelLayoutNbChannels(dstChLayout)
	if ret = ffmpeg.AvSamplesAllocArrayAndSamples(&dstData, &dstLinesize, dstNbChannels,
		dstNbSamples, dstSampleFmt, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate destination samples\n")
		goto end
	}

	for ok := true; ok; ok = (t < 10) {
		// generate synthetic audio
		fileSamples(unsafe.Slice((*float64)(unsafe.Pointer(*srcData)), srcNbSamples*srcNbChannels),
			srcNbSamples, srcNbChannels, srcRate, &t)

		//compute destination number of samples
		dstNbSamples = ffmpeg.AvRescaleRnd(ffmpeg.SwrGetDelay(swrCtx, srcRate)+srcNbSamples,
			dstRate, srcRate, ffmpeg.AV_ROUND_UP)
		if dstNbSamples > maxDstNbSamples {
			ffmpeg.AvFreep(dstData)
			if ret = ffmpeg.AvSamplesAlloc(dstData, &dstLinesize, dstNbChannels,
				dstNbSamples, dstSampleFmt, 1); ret < 0 {
				break
			}
			maxDstNbSamples = dstNbSamples
		}

		// convert to destination format
		if ret = ffmpeg.SwrConvert(swrCtx, dstData, dstNbSamples, srcData, srcNbSamples); ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while converting\n")
			goto end
		}
		if dstBufsize = ffmpeg.AvSamplesGetBufferSize(&dstLinesize, dstNbChannels,
			ret, dstSampleFmt, 1); dstBufsize < 0 {
			fmt.Fprintf(os.Stderr, "Could not get sample buffer size\n")
			goto end
		}
		fmt.Fprintf(os.Stdout, "t:%f in:%d out:%d\n", t, srcNbSamples, ret)
		dstFile.Write(ffmpeg.SliceSlice(dstData, 1, dstBufsize)[0])
	}

	if _fmt, ret = getFormatFromSampleFmt(dstSampleFmt); ret < 0 {
		goto end
	}
	fmt.Fprintf(os.Stderr, "Resampling succeeded. Play the output file with the command:\n"+
		"ffplay -f %s -channel_layout %d -channels %d -ar %d %s\n",
		_fmt, dstChLayout, dstNbChannels, dstRate, dstFilename)

end:
	dstFile.Close()

	if srcData != nil {
		ffmpeg.AvFreep(srcData)
	}
	ffmpeg.AvFreep(&srcData)

	if dstData != nil {
		ffmpeg.AvFreep(dstData)
	}
	ffmpeg.AvFreep(&dstData)

	ffmpeg.SwrFree(&swrCtx)
	if ret < 0 {
		os.Exit(1)
	}
}
