// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavfilter/version.h>
*/
import "C"

const (
	LIBAVFILTER_VERSION_MINOR = C.LIBAVFILTER_VERSION_MINOR
	LIBAVFILTER_VERSION_MICRO = C.LIBAVFILTER_VERSION_MICRO
)
