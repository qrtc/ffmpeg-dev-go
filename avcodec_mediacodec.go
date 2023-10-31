// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_mediacodec

package ffmpeg

/*
#include <libavcodec/mediacodec.h>
*/
import "C"

// AVMediaCodecContext
type AVMediaCodecContext C.struct_AVMediaCodecContext

// AvMediacodecAllocContext allocates and initializes a MediaCodec context.
func AvMediacodecAllocContext() *AVMediaCodecContext {
	return (*AVMediaCodecContext)(C.av_mediacodec_alloc_context())
}

// AvMediacodecDefaultInit sets up the MediaCodec context.
func AvMediacodecDefaultInit(avctx *AVCodecContext, ctx *AVMediaCodecContext, surface CVoidPointer) int32 {
	return (int32)(C.av_mediacodec_default_init((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVMediaCodecContext)(ctx), VoidPointer(surface)))
}

// AvMediacodecDefaultFree frees the MediaCodec context.
func AvMediacodecDefaultFree(avctx *AVCodecContext) {
	C.av_mediacodec_default_free((*C.struct_AVCodecContext)(avctx))
}

// MediaCodecBuffer
type MediaCodecBuffer C.struct_MediaCodecBuffer

// AVMediaCodecBuffer
type AVMediaCodecBuffer = MediaCodecBuffer

// AvMediacodecReleaseBuffer releases a MediaCodec buffer and render it to the surface that is associated
// with the decoder.
func AvMediacodecReleaseBuffer(buffer *AVMediaCodecBuffer, render int32) int32 {
	return (int32)(C.av_mediacodec_release_buffer((*C.struct_MediaCodecBuffer)(buffer), (C.int)(render)))
}

// AvMediacodecRenderBufferAtTime release a MediaCodec buffer and render it at the given time to the surface
// that is associated with the decoder.
func AvMediacodecRenderBufferAtTime(buffer *AVMediaCodecBuffer, time int64) int32 {
	return (int32)(C.av_mediacodec_render_buffer_at_time((*C.struct_MediaCodecBuffer)(buffer), (C.int64_t)(time)))
}
