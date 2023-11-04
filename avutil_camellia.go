// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/camellia.h>
*/
import "C"

// AVCAMELLIA
type AVCAMELLIA C.struct_AVCAMELLIA

// AvCamelliaAlloc allocates an AVCAMELLIA context.
func AvCamelliaAlloc() *AVCAMELLIA {
	return (*AVCAMELLIA)(C.av_camellia_alloc())
}

// AvCamelliaInit initializes an AVCAMELLIA context.
func AvCamelliaInit(ctx *AVCAMELLIA, key *uint8, keyBits int32) int32 {
	return (int32)(C.av_camellia_init((*C.struct_AVCAMELLIA)(ctx),
		(*C.uint8_t)(key), (C.int)(keyBits)))
}

// AvCamelliaCrypt encrypts or decrypts a buffer using a previously initialized context.
func AvCamelliaCrypt(ctx *AVCAMELLIA, dst, src *uint8, count int32, iv *uint8, decrypt int32) {
	C.av_camellia_crypt((*C.struct_AVCAMELLIA)(ctx),
		(*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}
