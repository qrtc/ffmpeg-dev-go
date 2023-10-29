// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_dxva2

package ffmpeg

/*
#include <libavutil/hwcontext_dxva2.h>
*/
import "C"

// AVDXVA2DeviceContext
type AVDXVA2DeviceContext C.struct_AVDXVA2DeviceContext

// AVDXVA2FramesContext
type AVDXVA2FramesContext C.struct_AVDXVA2FramesContext
