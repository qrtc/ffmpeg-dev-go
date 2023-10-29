// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_xvmc

package ffmpeg

/*
#include <libavcodec/videotoolbox.h>
*/
import "C"

const (
	AV_XVMC_ID = C.AV_XVMC_ID
)

// Deprecated: No use.
//
// XvmcPixFmt
type XvmcPixFmt C.struct_xvmc_pix_fmt
