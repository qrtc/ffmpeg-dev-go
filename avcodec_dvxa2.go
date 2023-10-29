// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_dvxa2

package ffmpeg

/*
#include <libavcodec/dvxa2.h>
*/
import "C"

const {
	FF_DXVA2_WORKAROUND_SCALING_LIST_ZIGZAG = C.FF_DXVA2_WORKAROUND_SCALING_LIST_ZIGZAG
	FF_DXVA2_WORKAROUND_INTEL_CLEARVIDEO = C.FF_DXVA2_WORKAROUND_INTEL_CLEARVIDEO
}

// DxvaContext
type DxvaContext C.struct_dxva_context

