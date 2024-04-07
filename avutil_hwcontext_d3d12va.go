// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_d3d12va

package ffmpeg

/*
#include <libavutil/hwcontext_d3d12va.h>
*/
import "C"

// AVD3D12VADeviceContext
type AVD3D12VADeviceContext C.struct_AVD3D12VADeviceContext

// AVD3D12VASyncContext
type AVD3D12VASyncContext C.struct_AVD3D12VASyncContext

// AVD3D12VAFrame
type AVD3D12VAFrame C.struct_AVD3D12VAFrame

// AVD3D12VAFramesContext
type AVD3D12VAFramesContext C.struct_AVD3D12VAFramesContext
