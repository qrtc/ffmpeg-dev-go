// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/ripemd.h>
*/
import "C"

// AVRIPEMD
type AVRIPEMD C.struct_AVRIPEMD

// AvRipemdAlloc allocates an AVRIPEMD context.
func AvRipemdAlloc() *AVRIPEMD {
	return (*AVRIPEMD)(C.av_ripemd_alloc())
}

// AvRipemdInit initializes RIPEMD hashing.
func AvRipemdInit(context *AVRIPEMD, bits int32) int32 {
	return (int32)(C.av_ripemd_init((*C.struct_AVRIPEMD)(context), (C.int)(bits)))
}

// AvRipemdUpdate updates hash value.
func AvRipemdUpdate(context *AVRIPEMD, data *uint8, len uint32) {
	C.av_ripemd_update((*C.struct_AVRIPEMD)(context), (*C.uint8_t)(data), (C.uint)(len))
}

// AvRipemdFinal finishes hashing and output digest value.
func AvRipemdFinal(context *AVRIPEMD, digest *uint8) {
	C.av_ripemd_final((*C.struct_AVRIPEMD)(context), (*C.uint8_t)(digest))
}
