// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_videotoolbox

package ffmpeg

/*
#include <libavcodec/videotoolbox.h>
*/
import "C"

type AVVideotoolboxContext C.struct_AVVideotoolboxContext

// AvVideotoolboxAllocContext allocates and initializes a Videotoolbox context.
func AvVideotoolboxAllocContext() *AVVideotoolboxContext {
	return (*AVVideotoolboxContext)(C.av_videotoolbox_alloc_context())
}

// AvVideotoolboxDefaultInit creates and sets up the Videotoolbox context using
// an internal implementation.
func AvVideotoolboxDefaultInit(avctx *AVCodecContext) int32 {
	return (int32)(C.av_videotoolbox_default_init((*C.struct_AVCodecContext)(avctx)))
}

// AvVideotoolboxDefaultInit2 creates and sets up the Videotoolbox context using
// an internal implementation.
func AvVideotoolboxDefaultInit2(avctx *AVCodecContext, vtctx *AVVideotoolboxContext) int32 {
	return (int32)(C.av_videotoolbox_default_init2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVVideotoolboxContext)(vtctx)))
}

// AvVideotoolboxDefaultFree frees the Videotoolbox context.
func AvVideotoolboxDefaultFree(avctx *AVCodecContext) {
	C.av_videotoolbox_default_free((*C.struct_AVCodecContext)(avctx))
}
