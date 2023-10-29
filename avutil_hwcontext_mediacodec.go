// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_mediacodec

package ffmpeg

/*
#include <libavutil/hwcontext_mediacodec.h>
*/
import "C"

// AVMediaCodecDeviceContext
type AVMediaCodecDeviceContext C.struct_AVMediaCodecDeviceContext
