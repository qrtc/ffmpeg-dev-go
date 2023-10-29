// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/aes.h>
*/
import "C"

// AVAES
type AVAES C.struct_AVAES

// AvAesAlloc allocates an AVAES context.
func AvAesAlloc() *AVAES {
	return (*AVAES)(C.av_aes_alloc())
}

// AvAesInit initializes an AVAES context.
func AvAesInit(a *AVAES, key *uint8, keyBits, decrypt int32) int32 {
	return (int32)(C.av_aes_init((*C.struct_AVAES)(a), (*C.uint8_t)(key),
		(C.int)(keyBits), (C.int)(decrypt)))
}

// AvAesCrypt encrypts or decrypts a buffer using a previously initialized context.
func AvAesCrypt(a *AVAES, dst, src *uint8, count int32, iv *uint8, decrypt int32) {
	C.av_aes_crypt((*C.struct_AVAES)(a), (*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}
