package ffmpeg

/*
#include <libavutil/common.h>
#include <libavutil/error.h>
*/
import "C"

// AVERROR returns a negative error code from a POSIX error code, to return from library functions.
func AVERROR[T Integer](ie T) (oe int32) {
	oe = (int32)(ie)
	if C.EDOM > 0 {
		return (-oe)
	}
	return oe
}

// AVUNERROR returns a POSIX error code from a library function error return value.
func AVUNERROR(e int32) int32 {
	if C.EDOM > 0 {
		return (-e)
	}
	return e
}

// Error handling
const (
	AVERROR_BSF_NOT_FOUND      = int32(C.AVERROR_BSF_NOT_FOUND)
	AVERROR_BUG                = int32(C.AVERROR_BUG)
	AVERROR_BUFFER_TOO_SMALL   = int32(C.AVERROR_BUFFER_TOO_SMALL)
	AVERROR_DECODER_NOT_FOUND  = int32(C.AVERROR_DECODER_NOT_FOUND)
	AVERROR_DEMUXER_NOT_FOUND  = int32(C.AVERROR_DEMUXER_NOT_FOUND)
	AVERROR_ENCODER_NOT_FOUND  = int32(C.AVERROR_ENCODER_NOT_FOUND)
	AVERROR_EOF                = int32(C.AVERROR_EOF)
	AVERROR_EXIT               = int32(C.AVERROR_EXIT)
	AVERROR_EXTERNAL           = int32(C.AVERROR_EXTERNAL)
	AVERROR_FILTER_NOT_FOUND   = int32(C.AVERROR_FILTER_NOT_FOUND)
	AVERROR_INVALIDDATA        = int32(C.AVERROR_INVALIDDATA)
	AVERROR_MUXER_NOT_FOUND    = int32(C.AVERROR_MUXER_NOT_FOUND)
	AVERROR_OPTION_NOT_FOUND   = int32(C.AVERROR_OPTION_NOT_FOUND)
	AVERROR_PATCHWELCOME       = int32(C.AVERROR_PATCHWELCOME)
	AVERROR_PROTOCOL_NOT_FOUND = int32(C.AVERROR_PROTOCOL_NOT_FOUND)

	AVERROR_STREAM_NOT_FOUND = int32(C.AVERROR_STREAM_NOT_FOUND)

	AVERROR_BUG2              = int32(C.AVERROR_BUG2)
	AVERROR_UNKNOWN           = int32(C.AVERROR_UNKNOWN)
	AVERROR_EXPERIMENTAL      = int32(C.AVERROR_EXPERIMENTAL)
	AVERROR_INPUT_CHANGED     = int32(C.AVERROR_INPUT_CHANGED)
	AVERROR_OUTPUT_CHANGED    = int32(C.AVERROR_OUTPUT_CHANGED)
	AVERROR_HTTP_BAD_REQUEST  = int32(C.AVERROR_HTTP_BAD_REQUEST)
	AVERROR_HTTP_UNAUTHORIZED = int32(C.AVERROR_HTTP_UNAUTHORIZED)
	AVERROR_HTTP_FORBIDDEN    = int32(C.AVERROR_HTTP_FORBIDDEN)
	AVERROR_HTTP_NOT_FOUND    = int32(C.AVERROR_HTTP_NOT_FOUND)
	AVERROR_HTTP_OTHER_4XX    = int32(C.AVERROR_HTTP_OTHER_4XX)
	AVERROR_HTTP_SERVER_ERROR = int32(C.AVERROR_HTTP_SERVER_ERROR)
)

const (
	AV_ERROR_MAX_STRING_SIZE = int32(C.AV_ERROR_MAX_STRING_SIZE)
)

// AvStrerror puts a description of the AVERROR code errnum in errbuf.
// In case of failure the global variable errno is set to indicate the
// error. Even in case of failure AvStrerror() will print a generic
// error message indicating the errnum provided to errbuf.
func AvStrerror(errnum int32, errbuf *int8, errbufSize uintptr) int32 {
	return (int32)(C.av_strerror((C.int)(errnum), (*C.char)(errbuf), (C.size_t)(errbufSize)))
}

// AvMakeErrorString fills the provided buffer with a string containing an error string
// corresponding to the AVERROR code errnum.
func AvMakeErrorString(errbuf *int8, errbufSize uintptr, errnum int32) *int8 {
	return (*int8)(C.av_make_error_string((*C.char)(errbuf), (C.size_t)(errbufSize), (C.int)(errnum)))
}

// AvErr2str
func AvErr2str(errnum int32) string {
	buf := make([]int8, AV_ERROR_MAX_STRING_SIZE)
	return C.GoString(C.av_make_error_string((*C.char)(&buf[0]),
		(C.size_t)(AV_ERROR_MAX_STRING_SIZE), (C.int)(errnum)))
}
