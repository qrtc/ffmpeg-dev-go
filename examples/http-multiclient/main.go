package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

func processClient(client *ffmpeg.AVIOContext, inUri string) {
	var (
		input       *ffmpeg.AVIOContext
		resource    *C.char
		buf         [1024]uint8
		resourceStr string
		ret         int32
		replyCode   int32
	)

	for {
		if ret = ffmpeg.AvIOHandshake(client); ret > 0 {
			ffmpeg.AvOptGet(client, "resource",
				ffmpeg.AV_OPT_SEARCH_CHILDREN, (**uint8)(unsafe.Pointer(&resource)))
			// check for strlen(resource) is necessary, because av_opt_get()
			// may return empty string.
			resourceStr = C.GoString(resource)
			if resource != nil && len(resourceStr) > 0 {
				break
			}
			ffmpeg.AvFreep(&resource)
		} else {
			break
		}
	}
	if ret < 0 {
		goto end
	}
	ffmpeg.AvLog(client, ffmpeg.AV_LOG_TRACE, "resource=%s\n", resourceStr)
	if resource != nil && resourceStr[0] == '/' && resourceStr[1:] == inUri {
		replyCode = 200
	} else {
		ffmpeg.AvLog(client, ffmpeg.AV_LOG_TRACE, "resource: %s", resourceStr)
		replyCode = ffmpeg.AVERROR_HTTP_NOT_FOUND
	}
	if ret = ffmpeg.AvOptSetInt(client, "reply_code", int64(replyCode), ffmpeg.AV_OPT_SEARCH_CHILDREN); ret < 0 {
		ffmpeg.AvLog(client, ffmpeg.AV_LOG_ERROR, "Failed to set reply_code: %s.\n", ffmpeg.AvErr2str(ret))
		goto end
	}
	ffmpeg.AvLog(client, ffmpeg.AV_LOG_TRACE, "Set reply code to %d\n", replyCode)

	for {
		if ret = ffmpeg.AvIOHandshake(client); ret <= 0 {
			break
		}
	}

	if ret < 0 {
		goto end
	}

	fmt.Fprintf(os.Stderr, "Handshake performed.\n")

	if replyCode != 200 {
		goto end
	}
	fmt.Fprintf(os.Stderr, "Opening input file.\n")
	if ret = ffmpeg.AvIOOpen2(&input, inUri, ffmpeg.AVIO_FLAG_READ, nil, nil); ret < 0 {
		ffmpeg.AvLog(client, ffmpeg.AV_LOG_ERROR, "Failed to open input: %s: %s.\n", inUri, ffmpeg.AvErr2str(ret))
		goto end
	}
	for {
		n := ffmpeg.AvIORead(input, &buf[0], 1024)
		if n < 0 {
			if n == ffmpeg.AVERROR_EOF {
				break
			}
			ffmpeg.AvLog(client, ffmpeg.AV_LOG_ERROR, "Error reading from input: %s.\n", ffmpeg.AvErr2str(ret))
			break
		}
		ffmpeg.AvIOWrite(client, &buf[0], n)
		ffmpeg.AvIOFlush(client)
	}
end:
	fmt.Fprintf(os.Stderr, "Flushing client\n")
	ffmpeg.AvIOFlush(client)
	fmt.Fprintf(os.Stderr, "Closing client\n")
	ffmpeg.AvIOClose(client)
	fmt.Fprintf(os.Stderr, "Closing input\n")
	ffmpeg.AvIOClose(input)
	ffmpeg.AvFreep(&resource)
}

func main() {
	var (
		options        *ffmpeg.AVDictionary
		client, server *ffmpeg.AVIOContext
		ret            int32
	)
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s input http://hostname[:port]\n"+
			"API example program to serve http to multiple clients.\n"+
			"\n", os.Args[0])
		os.Exit(1)
	}

	inUri := os.Args[1]
	outUri := os.Args[2]

	ffmpeg.AvFormatNetworkInit()
	ffmpeg.AvLogSetLevel(ffmpeg.AV_LOG_TRACE)

	if ret = ffmpeg.AvDictSet(&options, "listen", "2", 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to set listen mode for server: %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(int(ret))
	}
	if ret = ffmpeg.AvIOOpen2(&server, outUri, ffmpeg.AVIO_FLAG_WRITE, nil, &options); ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to open server: %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(int(ret))
	}
	fmt.Fprintf(os.Stderr, "Entering main loop.\n")
	for {
		if ret = ffmpeg.AvIOAccept(server, &client); ret < 0 {
			goto end
		}
		go processClient(client, inUri)
	}
end:
	ffmpeg.AvIOClose(server)
	if ret < 0 && ret != ffmpeg.AVERROR_EOF {
		fmt.Fprintf(os.Stderr, "Some errors occurred: %s\n", ffmpeg.AvErr2str(ret))
		os.Exit(1)
	}
}
