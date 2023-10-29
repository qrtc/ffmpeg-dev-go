// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/version.h>
*/
import "C"

const (
	LIBAVUTIL_VERSION_MAJOR = C.LIBAVUTIL_VERSION_MAJOR
	LIBAVUTIL_VERSION_MINOR = C.LIBAVUTIL_VERSION_MINOR
	LIBAVUTIL_VERSION_MICRO = C.LIBAVUTIL_VERSION_MICRO
)
