// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/cast5.h>
*/
import "C"

type AVCAST5 C.struct_AVCAST5

// AvCast5Alloc allocates an AVCAST5 context.
func AvCast5Alloc() *AVCAST5 {
	return (*AVCAST5)(C.av_cast5_alloc())
}

// AvCast5Init initializes an AVCAST5 context.
func AvCast5Init(ctx *AVCAST5, key *uint8, keyBits int32) int32 {
	return (int32)(C.av_cast5_init((*C.struct_AVCAST5)(ctx),
		(*C.uint8_t)(key), (C.int)(keyBits)))
}

// AvCast5Crypt encrypts or decrypts a buffer using a previously initialized context.
func AvCast5Crypt(ctx *AVCAST5, dst, src *uint8, count, decrypt int32) {
	C.av_cast5_crypt((*C.struct_AVCAST5)(ctx),
		(*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (C.int)(decrypt))
}

// AvCast5Crypt2 encrypts or decrypts a buffer using a previously initialized context.
func AvCast5Crypt2(ctx *AVCAST5, dst, src *uint8, count int32, iv *uint8, decrypt int32) {
	C.av_cast5_crypt2((*C.struct_AVCAST5)(ctx),
		(*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}
