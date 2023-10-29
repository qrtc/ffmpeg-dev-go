// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_drm

package ffmpeg

/*
#include <libavutil/hwcontext_drm.h>
*/
import "C"

const (
	AV_DRM_MAX_PLANES = C.AV_DRM_MAX_PLANES
)

// AVDRMObjectDescriptor
type AVDRMObjectDescriptor C.struct_AVDRMObjectDescriptor

// AVDRMPlaneDescriptor
type AVDRMPlaneDescriptor C.struct_AVDRMPlaneDescriptor

// AVDRMLayerDescriptor
type AVDRMLayerDescriptor C.struct_AVDRMLayerDescriptor

// AVDRMFrameDescriptor
type AVDRMFrameDescriptor C.struct_AVDRMFrameDescriptor

// AVDRMDeviceContext
type AVDRMDeviceContext C.struct_AVDRMDeviceContext
