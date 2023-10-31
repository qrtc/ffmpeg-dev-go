// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavdevice/version_major.h>
*/
import "C"

const (
	LIBAVDEVICE_VERSION_MAJOR = C.LIBAVDEVICE_VERSION_MAJOR
)
