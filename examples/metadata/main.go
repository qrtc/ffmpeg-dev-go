package main

import (
	"fmt"
	"os"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

func main() {
	var fmtCtx *ffmpeg.AVFormatContext
	var tag *ffmpeg.AVDictionaryEntry
	var ret int32

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <input_file>\n"+
			"example program to demonstrate the use of the libavformat metadata API.\n"+
			"\n", os.Args[0])
		os.Exit(1)
	}

	if ret = ffmpeg.AvFormatOpenInput(&fmtCtx, os.Args[1], nil, nil); ret != 0 {
		os.Exit(int(ret))
	}

	if ret = ffmpeg.AvFormatFindStreamInfo(fmtCtx, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot find stream information\n")
		os.Exit(int(ret))
	}

	for {
		if tag = ffmpeg.AvDictGet(fmtCtx.GetMetadata(), "", tag, ffmpeg.AV_DICT_IGNORE_SUFFIX); tag == nil {
			break
		}
		fmt.Fprintf(os.Stdout, "%s=%s\n", tag.GetKey(), tag.GetValue())
	}

	ffmpeg.AvFormatCloseInput(&fmtCtx)
	os.Exit(0)
}
