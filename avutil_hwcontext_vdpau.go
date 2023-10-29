// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_vdpau

package ffmpeg

/*
#include <libavutil/hwcontext_vdpau.h>
*/
import "C"

// AVVDPAUDeviceContext
type AVVDPAUDeviceContext C.struct_AVVDPAUDeviceContext
