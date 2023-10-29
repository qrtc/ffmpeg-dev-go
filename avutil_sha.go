// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/sha.h>
*/
import "C"

// AVSHA
type AVSHA C.struct_AVSHA

// AvShaAlloc allocates an AVSHA context.
func AvShaAlloc() *AVSHA {
	return (*AVSHA)(C.av_sha_alloc())
}

// AvShaInit initializes SHA-1 or SHA-2 hashing.
func AvShaInit(context *AVSHA, bits int32) int32 {
	return (int32)(C.av_sha_init((*C.struct_AVSHA)(context), (C.int)(bits)))
}

// AvShaUpdate updates hash value.
func AvShaUpdate(ctx *AVSHA, data *uint8, len uint32) {
	C.av_sha_update((*C.struct_AVSHA)(ctx), (*C.uint8_t)(data), (C.uint)(len))
}

// AvShaFinal finishes hashing and output digest value.
func AvShaFinal(ctx *AVSHA, digest *uint8) {
	C.av_sha_final((*C.struct_AVSHA)(ctx), (*C.uint8_t)(digest))
}
