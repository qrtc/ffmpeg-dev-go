// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_vdpau

package ffmpeg

/*
#include <libavcodec/vdpau.h>
*/
import "C"

// typedef int (*AVVDPAU_Render2)(struct AVCodecContext *, struct AVFrame *,
// const VdpPictureInfo *, uint32_t,
// const VdpBitstreamBuffer *);
type AVVDPAURenderFunc2 = C.AVVDPAU_Render2

// AVVDPAUContext
type AVVDPAUContext C.struct_AVVDPAUContext

// Deprecated: No use.
//
// AvAllocVdpaucontext allocates AVVDPAUContext.
func AvAllocVdpaucontext() *AVVDPAUContext {
	return (*AVVDPAUContext)(C.av_alloc_vdpaucontext())
}

// Deprecated: No use.
//
// AvVdpauHwaccelGetRender2
func AvVdpauHwaccelGetRender2(ctx *AVVDPAUContext) AVVDPAURenderFunc2 {
	return (AVVDPAURenderFunc2)(C.av_vdpau_hwaccel_get_render2((*C.struct_AVVDPAUContext)(ctx)))
}

// Deprecated: No use.
//
// AvVdpauHwaccelSetRender2
func AvVdpauHwaccelSetRender2(ctx *AVVDPAUContext, f AVVDPAURenderFunc2) {
	C.av_vdpau_hwaccel_set_render2((*C.struct_AVVDPAUContext)(ctx), (C.AVVDPAU_Render2)(f))
}

// NONEED: av_vdpau_get_surface_parameters

// NONEED: av_vdpau_get_surface_parameters

// NONEED: av_vdpau_bind_context

// Deprecated: No use.
//
// AvVdpauAllocContext
func AvVdpauAllocContext() *AVVDPAUContext {
	return (*AVVDPAUContext)(C.av_vdpau_alloc_context())
}

// NONEED: av_vdpau_get_profile
