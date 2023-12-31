// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavdevice/version.h>
*/
import "C"

const (
	LIBAVDEVICE_VERSION_MINOR = C.LIBAVDEVICE_VERSION_MINOR
	LIBAVDEVICE_VERSION_MICRO = C.LIBAVDEVICE_VERSION_MICRO
)
