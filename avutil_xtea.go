// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/xtea.h>
*/
import "C"
import "unsafe"

type AVXTEA C.struct_AVXTEA

// GetKey gets `AVXTEA.key` value.
func (dct *AVXTEA) GetKey() []uint32 {
	return unsafe.Slice((*uint32)(&dct.key[0]), 16)
}

// SetKey sets `AVXTEA.key` value.
func (dct *AVXTEA) SetKey(v []uint32) {
	for i := 0; i < FFMIN(len(v), 16); i++ {
		dct.key[i] = (C.uint32_t)(v[i])
	}
}

// GetKeyAddr gets `AVXTEA.key` address.
func (dct *AVXTEA) GetKeyAddr() **uint32 {
	return (**uint32)(unsafe.Pointer(&dct.key))
}

// AvXteaAlloc allocates an AVXTEA context.
func AvXteaAlloc() *AVXTEA {
	return (*AVXTEA)(C.av_xtea_alloc())
}

// AvXteaInit initializes an AVXTEA context.
func AvXteaInit(ctx *AVXTEA, key []uint8) {
	if len(key) < 16 {
		panic("key len < 16")
	}
	C.av_xtea_init((*C.struct_AVXTEA)(ctx), (*C.uint8_t)(&key[0]))
}

// AvXteaLeInit initializes an AVXTEA context.
func AvXteaLeInit(ctx *AVXTEA, key []uint8) {
	if len(key) < 16 {
		panic("key len < 16")
	}
	C.av_xtea_le_init((*C.struct_AVXTEA)(ctx), (*C.uint8_t)(&key[0]))
}

// AvXteaCrypt encrypts or decrypts a buffer using a previously initialized context,
// in big endian format.
func AvXteaCrypt(ctx *AVXTEA, dst, src *uint8, count int32, iv *uint8, decrypt int32) {
	C.av_xtea_crypt((*C.struct_AVXTEA)(ctx), (*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}

// AvXteaLeCrypt encrypts or decrypts a buffer using a previously initialized context,
// in little endian format.
func AvXteaLeCrypt(ctx *AVXTEA, dst, src *uint8, count int32, iv *uint8, decrypt int32) {
	C.av_xtea_le_crypt((*C.struct_AVXTEA)(ctx), (*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}
