package main

/*
#include <stdlib.h>
#include <stdint.h>

int readPacket(void* opaque, uint8_t *buf, int bufSize);

*/
import "C"
import (
	"fmt"
	"os"
	"runtime/cgo"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

type bufferData struct {
	ptr  *uint8
	size uintptr
}

//export readPacket
func readPacket(opaque unsafe.Pointer, buf *uint8, bufSize int32) int32 {
	bd := (*(*cgo.Handle)(opaque)).Value().(*bufferData)
	bufSize = ffmpeg.FFMIN(bufSize, int32(bd.size))

	if bufSize == 0 {
		return ffmpeg.AVERROR_EOF
	}
	fmt.Fprintf(os.Stdout, "ptr:%p size:%d\n", bd.ptr, bd.size)

	// copy internal buffer data to buf
	bd.ptr = ffmpeg.PointerOffset(bd.ptr, bufSize)
	bd.size -= uintptr(bufSize)

	return bufSize
}

func main() {
	var fmtCtx *ffmpeg.AVFormatContext
	var avioCtx *ffmpeg.AVIOContext
	var buffer, avioCtxBuffer *uint8
	var bufferSize uintptr
	var avioCtxBufferSize uint = 4096
	var ret int32
	var bd bufferData
	var opaqueHandle cgo.Handle

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s input_file\n"+
			"API example program to show how to read from a custom buffer "+
			"accessed through AVIOContext.\n", os.Args[0])
		os.Exit(1)
	}
	inputFilename := os.Args[1]

	// slurp file content into buffer
	if ret = ffmpeg.AvFileMap(inputFilename, &buffer, &bufferSize, 0, nil); ret < 0 {
		goto end
	}

	// fill opaque structure used by the AVIOContext read callback
	bd.ptr = buffer
	bd.size = bufferSize

	if fmtCtx = ffmpeg.AvFormatAllocContext(); fmtCtx == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	if avioCtxBuffer = (*uint8)(ffmpeg.AvMalloc(avioCtxBufferSize)); avioCtxBuffer == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}

	opaqueHandle = cgo.NewHandle(&bd)
	if avioCtx = ffmpeg.AvIOAllocContext(avioCtxBuffer, int32(avioCtxBufferSize), 0,
		&opaqueHandle, (ffmpeg.AVIOContextReadPacketFunc)(C.readPacket), nil, nil); avioCtx == nil {
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto end
	}
	fmtCtx.SetPb(avioCtx)

	if ret = ffmpeg.AvFormatOpenInput(&fmtCtx, ffmpeg.NIL, nil, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open input\n")
		goto end
	}

	if ret = ffmpeg.AvFormatFindStreamInfo(fmtCtx, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not find stream information\n")
		goto end
	}

	ffmpeg.AvDumpFormat(fmtCtx, 0, inputFilename, 0)

end:
	ffmpeg.AvFormatCloseInput(&fmtCtx)

	if opaqueHandle != 0 {
		opaqueHandle.Delete()
	}

	// note: the internal buffer could have changed, and be != avio_ctx_buffer
	if avioCtx != nil {
		ffmpeg.AvFreep(avioCtx.GetBufferAddr())
	}
	ffmpeg.AvIOContextFree(&avioCtx)

	ffmpeg.AvFileUnmap(buffer, bufferSize)

	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Error occurred: %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(1)
	}
}
