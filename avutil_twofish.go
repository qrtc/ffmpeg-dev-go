// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/twofish.h>
*/
import "C"

// AVTWOFISH
type AVTWOFISH C.struct_AVTWOFISH

// AvTwofishAlloc allocates an AVTWOFISH context.
func AvTwofishAlloc() *AVTWOFISH {
	return (*AVTWOFISH)(C.av_twofish_alloc())
}

// AvTwofishInit initializes an AVTWOFISH context.
func AvTwofishInit(d *AVTWOFISH, key *uint8, keyBits int32) int32 {
	return (int32)(C.av_twofish_init((*C.struct_AVTWOFISH)(d),
		(*C.uint8_t)(key), (C.int)(keyBits)))
}

// AvTwofishCrypt encrypts or decrypts a buffer using a previously initialized context
func AvTwofishCrypt(d *AVTWOFISH, dst, src *uint8, count int32, iv *uint8, decrypt int32) {
	C.av_twofish_crypt((*C.struct_AVTWOFISH)(d),
		(*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}
