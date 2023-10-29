// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_vaapi

package ffmpeg

/*
#include <libavcodec/vappi.h>
*/
import "C"

// Deprecated: No use.
//
// VaapiContext
type VaapiContext C.struct_vaapi_context
