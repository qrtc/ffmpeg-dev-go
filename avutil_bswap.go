package ffmpeg

/*
#include <libavutil/bswap.h>
*/
import "C"

// AvBswap16
func AvBswap16(x uint16) uint16 {
	return (uint16)(C.av_bswap16((C.uint16_t)(x)))
}

// AvBswap32
func AvBswap32(x uint32) uint32 {
	return (uint32)(C.av_bswap32((C.uint32_t)(x)))
}

// AvBswap64
func AvBswap64(x uint64) uint64 {
	return (uint64)(C.av_bswap64((C.uint64_t)(x)))
}
