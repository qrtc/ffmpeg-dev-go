package ffmpeg

/*
#include <libavutil/tea.h>
*/
import "C"

// AVTEA
type AVTEA C.struct_AVTEA

// AvTeaAlloc allocates an AVTEA context.
func AvTeaAlloc() *AVTEA {
	return (*AVTEA)(C.av_tea_alloc())
}

// AvTeaInit initializes an AVTEA context.
func AvTeaInit(ctx *AVTEA, key []uint8, rounds int32) {
	if len(key) < 16 {
		panic("key len < 16")
	}
	C.av_tea_init((*C.struct_AVTEA)(ctx), (*C.uint8_t)(&key[0]), (C.int)(rounds))
}

// AvTeaCrypt encrypts or decrypts a buffer using a previously initialized context.
func AvTeaCrypt(ctx *AVTEA, dst, src *uint8, count int32, iv *uint8, decrypt int32) {
	C.av_tea_crypt((*C.struct_AVTEA)(ctx), (*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}
