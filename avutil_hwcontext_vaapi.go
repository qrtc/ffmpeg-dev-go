// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_vaapi

package ffmpeg

/*
#include <libavutil/hwcontext_vaapi.h>
*/
import "C"

const (
	AV_VAAPI_DRIVER_QUIRK_USER_SET             = C.AV_VAAPI_DRIVER_QUIRK_USER_SET
	AV_VAAPI_DRIVER_QUIRK_RENDER_PARAM_BUFFERS = C.AV_VAAPI_DRIVER_QUIRK_RENDER_PARAM_BUFFERS
	AV_VAAPI_DRIVER_QUIRK_ATTRIB_MEMTYPE       = C.AV_VAAPI_DRIVER_QUIRK_ATTRIB_MEMTYPE
	AV_VAAPI_DRIVER_QUIRK_SURFACE_ATTRIBUTES   = C.AV_VAAPI_DRIVER_QUIRK_SURFACE_ATTRIBUTES
)

// AVVAAPIDeviceContext
type AVVAAPIDeviceContext C.struct_AVVAAPIDeviceContext

// AVVAAPIFramesContext
type AVVAAPIFramesContext C.struct_AVVAAPIFramesContext

// AVVAAPIHWConfig
type AVVAAPIHWConfig C.struct_AVVAAPIHWConfig
