// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavcodec/version.h>
*/
import "C"

const (
	LIBAVCODEC_VERSION_MINOR = C.LIBAVCODEC_VERSION_MINOR
	LIBAVCODEC_VERSION_MICRO = C.LIBAVCODEC_VERSION_MICRO
)
