package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"syscall"
	"unsafe"

	"github.com/qrtc/ffmpeg-dev-go"
)

const (
	INPUT_SAMPLERATE     = 48000
	INPUT_FORMAT         = ffmpeg.AV_SAMPLE_FMT_FLTP
	INPUT_CHANNEL_LAYOUT = ffmpeg.AV_CH_LAYOUT_5POINT0
	VOLUME_VAL           = 0.90
	FRAME_SIZE           = 1024
)

func initFilterGraph() (graph *ffmpeg.AvFilterGraph, src *ffmpeg.AvFilterContext, sink *ffmpeg.AvFilterContext, ret int32) {
	var filterGraph *ffmpeg.AvFilterGraph
	var abufferCtx *ffmpeg.AvFilterContext
	var abuffer *ffmpeg.AvFilter
	var volumeCtx *ffmpeg.AvFilterContext
	var volume *ffmpeg.AvFilter
	var aformatCtx *ffmpeg.AvFilterContext
	var aformat *ffmpeg.AvFilter
	var abuffersinkCtx *ffmpeg.AvFilterContext
	var abuffersink *ffmpeg.AvFilter
	var optionsDict *ffmpeg.AvDictionary

	// Create a new filtergraph, which will contain all the filters.
	if filterGraph = ffmpeg.AvFilterGraphAlloc(); filterGraph == nil {
		fmt.Fprintf(os.Stderr, "Unable to create filter graph.\n")
		return nil, nil, nil, ffmpeg.AVERROR(int32(syscall.ENOMEM))
	}

	// Create the abuffer filter;
	// it will be used for feeding the data into the graph.
	if abuffer = ffmpeg.AvFilterGetByName("abuffer"); abuffer == nil {
		fmt.Fprintf(os.Stderr, "Could not find the abuffer filter.\n")
		return nil, nil, nil, ffmpeg.AVERROR_FILTER_NOT_FOUND
	}

	if abufferCtx = ffmpeg.AvFilterGraphAllocFilter(filterGraph, abuffer, "src"); abufferCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate the abuffer instance.\n")
		return nil, nil, nil, ffmpeg.AVERROR(int32(syscall.ENOMEM))
	}

	ffmpeg.AvOptSet(unsafe.Pointer(abufferCtx), "channel_layout",
		ffmpeg.AvGetChannelLayoutString(0, INPUT_CHANNEL_LAYOUT), ffmpeg.AV_OPT_SEARCH_CHILDREN)
	ffmpeg.AvOptSet(unsafe.Pointer(abufferCtx), "sample_fmt",
		ffmpeg.AvGetSampleFmtName(INPUT_FORMAT), ffmpeg.AV_OPT_SEARCH_CHILDREN)
	ffmpeg.AvOptSetQ(unsafe.Pointer(abufferCtx), "time_base",
		ffmpeg.AvMakeQ(1, INPUT_SAMPLERATE), ffmpeg.AV_OPT_SEARCH_CHILDREN)
	ffmpeg.AvOptSetInt(unsafe.Pointer(abufferCtx), "sample_rate",
		INPUT_SAMPLERATE, ffmpeg.AV_OPT_SEARCH_CHILDREN)

	// Now initialize the filter; we pass NULL options, since we have already set all the options above.
	if ret = ffmpeg.AvFilterInitStr(abufferCtx, ""); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not initialize the abuffer filter.\n")
		return nil, nil, nil, ret
	}

	// Create volume filter.
	if volume = ffmpeg.AvFilterGetByName("volume"); volume == nil {
		fmt.Fprintf(os.Stderr, "Could not find the volume filter.\n")
		return nil, nil, nil, ffmpeg.AVERROR_FILTER_NOT_FOUND
	}

	if volumeCtx = ffmpeg.AvFilterGraphAllocFilter(filterGraph, volume, "volume"); volumeCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate the volume instance.\n")
		return nil, nil, nil, ffmpeg.AVERROR(int32(syscall.ENOMEM))
	}

	// A different way of passing the options is as key/value pairs in a
	// dictionary.
	ffmpeg.AvDictSet(&optionsDict, "volume", fmt.Sprintf("%f", VOLUME_VAL), 0)
	ret = ffmpeg.AvFilterInitDict(volumeCtx, &optionsDict)
	ffmpeg.AvDictFree(&optionsDict)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not initialize the volume filter.\n")
		return nil, nil, nil, ret
	}

	// Create the aformat filter;
	// it ensures that the output is of the format we want.
	if aformat = ffmpeg.AvFilterGetByName("aformat"); aformat == nil {
		fmt.Fprintf(os.Stderr, "Could not find the aformat filter.\n")
		return nil, nil, nil, ffmpeg.AVERROR_FILTER_NOT_FOUND
	}

	if aformatCtx = ffmpeg.AvFilterGraphAllocFilter(filterGraph, aformat, "aformat"); aformatCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate the aformat instance.\n")
		return nil, nil, nil, ffmpeg.AVERROR(int32(syscall.ENOMEM))
	}

	// A third way of passing the options is in a string of the form
	// key1=value1:key2=value2....
	optionsStr := fmt.Sprintf("sample_fmts=%s:sample_rates=%d:channel_layouts=0x%d",
		ffmpeg.AvGetSampleFmtName(ffmpeg.AV_SAMPLE_FMT_S16), 44100, (uint64)(ffmpeg.AV_CH_LAYOUT_STEREO))
	if ret = ffmpeg.AvFilterInitStr(aformatCtx, optionsStr); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Could not initialize the aformat filter.\n")
		return nil, nil, nil, ret
	}

	// Finally create the abuffersink filter;
	// it will be used to get the filtered data out of the graph.
	if abuffersink = ffmpeg.AvFilterGetByName("abuffersink"); abuffersink == nil {
		fmt.Fprintf(os.Stderr, "Could not find the abuffersink filter.\n")
		return nil, nil, nil, ffmpeg.AVERROR_FILTER_NOT_FOUND
	}

	if abuffersinkCtx = ffmpeg.AvFilterGraphAllocFilter(filterGraph, abuffersink, "sink"); abuffersinkCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate the abuffersink instance.\n")
		return nil, nil, nil, ffmpeg.AVERROR(int32(syscall.ENOMEM))
	}

	// This filter takes no options.
	if ret = ffmpeg.AvFilterInitStr(abuffersinkCtx, ""); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Could not initialize the abuffersink instance.\n")
		return nil, nil, nil, ret
	}

	// Connect the filters;
	// in this simple case the filters just form a linear chain.
	ret = ffmpeg.AvFilterLink2(abufferCtx, 0, volumeCtx, 0)
	if ret >= 0 {
		ret = ffmpeg.AvFilterLink2(volumeCtx, 0, aformatCtx, 0)
	}
	if ret >= 0 {
		ret = ffmpeg.AvFilterLink2(aformatCtx, 0, abuffersinkCtx, 0)
	}
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Error connecting filters\n")
		return nil, nil, nil, ret
	}

	// Configure the graph.
	if ret = ffmpeg.AvFilterGraphConfig(filterGraph, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Error configuring the filter graph\n")
		return nil, nil, nil, ret
	}

	return filterGraph, abufferCtx, abuffersinkCtx, 0
}

// Do something useful with the filtered data: this simple
// example just prints the MD5 checksum of each plane to stdout.
func processOutput(md5 *ffmpeg.AvMD5, frame *ffmpeg.AvFrame) int32 {
	planar := ffmpeg.AvSampleFmtIsPlanar(frame.GetFormat())
	channels := ffmpeg.AvGetChannelLayoutNbChannels(frame.GetChannelLayout())
	planes := channels
	if planar == 0 {
		planes = 1
	}
	bps := ffmpeg.AvGetBytesPerSample(frame.GetFormat())
	planeSize := bps * frame.GetNbSamples()
	if planar == 0 {
		planeSize *= channels
	}

	for i := 0; i < int(planes); i++ {
		var checksum [16]uint8

		ffmpeg.AvMd5Init(md5)
		ffmpeg.AvMd5Sum(&checksum[0], frame.GetExtendedDataIdx(i), planeSize)

		fmt.Fprintf(os.Stdout, "plane %d: 0x", i)
		for j := 0; j < len(checksum); j++ {
			fmt.Fprintf(os.Stdout, "%02X", checksum[j])
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
	fmt.Fprintf(os.Stdout, "\n")

	return 0
}

// Construct a frame of audio data to be filtered;
// this simple example just synthesizes a sine wave.
func getInput(frame *ffmpeg.AvFrame, frameNum int32) int32 {
	// Set up the frame properties and allocate the buffer for the data.
	frame.SetSampleRate(INPUT_SAMPLERATE)
	frame.SetFormat(INPUT_FORMAT)
	frame.SetChannelLayout(INPUT_CHANNEL_LAYOUT)
	frame.SetNbSamples(FRAME_SIZE)
	frame.SetPts(int64(frameNum) * FRAME_SIZE)

	if ret := ffmpeg.AvFrameGetBuffer(frame, 0); ret < 0 {
		return ret
	}

	// Fill the data for each channel.
	for i := 0; i < 5; i++ {
		dataLen := int(frame.GetNbSamples())
		data := (*float32)(unsafe.Pointer(frame.GetExtendedDataIdx(i)))
		dataSlice := unsafe.Slice(data, dataLen)

		for j := 0; j < dataLen; j++ {
			dataSlice[j] = (float32)(math.Sin(2 * math.Pi * (float64)(int(frameNum)+j) * (float64)(i+1) / FRAME_SIZE))
		}
	}

	return 0
}

func main() {
	var md5 *ffmpeg.AvMD5
	var graph *ffmpeg.AvFilterGraph
	var src, sink *ffmpeg.AvFilterContext
	var frame *ffmpeg.AvFrame
	var ret int32

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <duration>\n", os.Args[0])
		os.Exit(1)
	}

	duration, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid duration: %s\n", os.Args[1])
		os.Exit(1)
	}
	nbFrames := (int32)(duration * INPUT_SAMPLERATE / FRAME_SIZE)
	if nbFrames <= 0 {
		fmt.Fprintf(os.Stderr, "Invalid duration: %s\n", os.Args[1])
		os.Exit(1)
	}

	// Allocate the frame we will be using to store the data.
	if frame = ffmpeg.AvFrameAlloc(); frame == nil {
		fmt.Fprintf(os.Stderr, "Error allocating the frame\n")
		os.Exit(1)
	}

	if md5 = ffmpeg.AvMd5Alloc(); md5 == nil {
		fmt.Fprintf(os.Stderr, "Error allocating the MD5 context\n")
		os.Exit(1)
	}

	// Set up the filtergraph.
	if graph, src, sink, ret = initFilterGraph(); ret < 0 {
		fmt.Fprintf(os.Stderr, "Unable to init filter graph:")
		goto fail
	}

	// the main filtering loop
	for i := int32(0); i < nbFrames; i++ {
		// get an input frame to be filtered
		if ret = getInput(frame, i); ret < 0 {
			fmt.Fprintf(os.Stderr, "Error generating input frame:")
			goto fail
		}

		// Send the frame to the input of the filtergraph.
		if ret = ffmpeg.AvBuffersrcAddFrame(src, frame); ret < 0 {
			ffmpeg.AvFrameUnref(frame)
			fmt.Fprintf(os.Stderr, "Error submitting the frame to the filtergraph:")
			goto fail
		}

		// Get all the filtered output that is available.
		for ffmpeg.AvBuffersinkGetFrame(sink, frame) >= 0 {
			// now do something with our filtered frame
			if ret = processOutput(md5, frame); ret < 0 {
				fmt.Fprintf(os.Stderr, "Error processing the filtered frame:")
				goto fail
			}
			ffmpeg.AvFrameUnref(frame)
		}

		if ret == ffmpeg.AVERROR(int32(syscall.EAGAIN)) {
			// Need to feed more frames in.
			continue
		} else if ret == ffmpeg.AVERROR_EOF {
			// Nothing more to do, finish.
			break
		} else if ret < 0 {
			// An error occurred.
			fmt.Fprintf(os.Stderr, "Error filtering the data:")
			goto fail
		}
	}

	ffmpeg.AvFilterGraphFree(&graph)
	ffmpeg.AvFrameFree(&frame)
	ffmpeg.AvFreep(unsafe.Pointer(&md5))
	return

fail:
	fmt.Fprintf(os.Stderr, "%s\n", ffmpeg.AvErr2str(ret))
	os.Exit(1)
}
