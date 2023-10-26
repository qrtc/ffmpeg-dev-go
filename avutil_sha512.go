package ffmpeg

/*
#include <libavutil/sha512.h>
*/
import "C"

// AVSHA512
type AVSHA512 C.struct_AVSHA512

// AvSha512Alloc allocates an AVSHA512 context.
func AvSha512Alloc() *AVSHA512 {
	return (*AVSHA512)(C.av_sha512_alloc())
}

// AvSha512Init initializes SHA-2 512 hashing.
func AvSha512Init(context *AVSHA512, bits int32) int32 {
	return (int32)(C.av_sha512_init((*C.struct_AVSHA512)(context), (C.int)(bits)))
}

// AvSha512Update updates hash value.
func AvSha512Update(ctx *AVSHA512, data *uint8, len uint32) {
	C.av_sha512_update((*C.struct_AVSHA512)(ctx), (*C.uint8_t)(data), (C.uint)(len))
}

// AvSha512Final finishes hashing and output digest value.
func AvSha512Final(ctx *AVSHA512, digest *uint8) {
	C.av_sha512_final((*C.struct_AVSHA512)(ctx), (*C.uint8_t)(digest))
}
