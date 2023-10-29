// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_qsv

package ffmpeg

/*
#include <libavutil/hwcontext_qsv.h>
*/
import "C"

// AVQSVDeviceContext
type AVQSVDeviceContext C.struct_AVQSVDeviceContext

// AVQSVFramesContext
type AVQSVFramesContext C.struct_AVQSVFramesContext
