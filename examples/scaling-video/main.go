package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

func fillYuvImage(data [4]*uint8, linesize [4]int32, width, height, frameIndex int32) {
	// Y
	data0 := unsafe.Slice(data[0], height*linesize[0])
	for y := int32(0); y < height; y++ {
		for x := int32(0); x < width; x++ {
			data0[y*linesize[0]+x] = (uint8)(x + y + frameIndex*3)
		}
	}

	// Cb and Cr
	data1 := unsafe.Slice(data[1], height*width/4)
	data2 := unsafe.Slice(data[2], height*width/4)
	for y := int32(0); y < height/2; y++ {
		for x := int32(0); x < width/2; x++ {
			data1[y*linesize[1]+x] = (uint8)(128 + y + frameIndex*2)
			data2[y*linesize[2]+x] = (uint8)(64 + x + frameIndex*5)
		}
	}
}

func main() {
	var ret int32
	var dstBufsize int32
	var srcW int32 = 320
	var srcH int32 = 240
	var dstW, dstH int32
	var srcData, dstData [4]*uint8
	var srcLinesize, dstLinesize [4]int32
	var srcPixFmt = ffmpeg.AV_PIX_FMT_YUV420P
	var dstPixFmt = ffmpeg.AV_PIX_FMT_RGB24

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s output_file output_size\n"+
			"API example program to show how to scale an image with libswscale.\n"+
			"This program generates a series of pictures, rescales them to the given "+
			"output_size and saves them to an output file named output_file\n."+
			"\n", os.Args[0])
		os.Exit(1)
	}
	dstFilename := os.Args[1]
	dstSize := os.Args[2]

	if ret = ffmpeg.AvParseVideoSize(&dstW, &dstH, dstSize); ret < 0 {
		fmt.Fprintf(os.Stderr,
			"Invalid size '%s', must be in the form WxH or a valid size abbreviation\n", dstSize)
		os.Exit(1)
	}

	dstFile, err := os.OpenFile(dstFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", dstFilename)
		os.Exit(1)
	}

	// create scaling context
	swsCtx := ffmpeg.SwsGetContext(srcW, srcH, srcPixFmt,
		dstW, dstH, dstPixFmt,
		ffmpeg.SWS_BILINEAR, nil, nil, nil)
	if swsCtx == nil {
		fmt.Fprintf(os.Stderr, "Impossible to create scale context for the conversion "+
			"fmt:%s s:%dx%d -> fmt:%s s:%dx%d\n",
			ffmpeg.AvGetPixFmtName(srcPixFmt), srcW, srcH,
			ffmpeg.AvGetPixFmtName(dstPixFmt), dstW, dstH)
		ret = ffmpeg.AVERROR(syscall.EINVAL)
		goto end
	}

	// allocate source and destination image buffers
	if ret = ffmpeg.AvImageAlloc(srcData[:], srcLinesize[:], srcW, srcH, srcPixFmt, 16); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate source image\n")
		goto end
	}

	// buffer is going to be written to rawvideo file, no alignment
	if ret = ffmpeg.AvImageAlloc(dstData[:], dstLinesize[:], dstW, dstH, dstPixFmt, 1); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate destination image\n")
		goto end
	}
	dstBufsize = ret

	for i := int32(0); i < 100; i++ {
		// generate synthetic video
		fillYuvImage(srcData, srcLinesize, srcW, srcH, i)

		// convert to destination format
		ffmpeg.SwsScale(swsCtx, srcData[:], srcLinesize[:], 0, srcH, dstData[:], dstLinesize[:])

		// write scaled image to file
		dstFile.Write(unsafe.Slice(dstData[0], dstBufsize))
	}

	fmt.Fprintf(os.Stderr, "Scaling succeeded. Play the output file with the command:\n"+
		"ffplay -f rawvideo -pix_fmt %s -video_size %dx%d %s\n",
		ffmpeg.AvGetPixFmtName(dstPixFmt), dstW, dstH, dstFilename)

end:
	dstFile.Close()
	ffmpeg.AvFreep(&srcData[0])
	ffmpeg.AvFreep(&dstData[0])
	ffmpeg.SwsFreeContext(swsCtx)
	if ret < 0 {
		os.Exit(1)
	}
}
