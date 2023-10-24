package ffmpeg

/*
#include <libavutil/lzo.h>
*/
import "C"

const (
	AV_LZO_INPUT_DEPLETED  = C.AV_LZO_INPUT_DEPLETED
	AV_LZO_OUTPUT_FULL     = C.AV_LZO_OUTPUT_FULL
	AV_LZO_INVALID_BACKPTR = C.AV_LZO_INVALID_BACKPTR
	AV_LZO_ERROR           = C.AV_LZO_ERROR
	AV_LZO_INPUT_PADDING   = C.AV_LZO_INPUT_PADDING
	AV_LZO_OUTPUT_PADDING  = C.AV_LZO_OUTPUT_PADDING
)

// AvLzo1xDecode decodes LZO 1x compressed data.
func AvLzo1xDecode(out CVoidPointer, outlen *int32, in CVoidPointer, inlen *int32) int32 {
	return (int32)(C.av_lzo1x_decode(VoidPointer(out),
		(*C.int)(outlen), VoidPointer(in), (*C.int)(inlen)))
}
