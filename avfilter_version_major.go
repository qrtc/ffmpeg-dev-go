// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavfilter/version_major.h>
*/
import "C"

const (
	LIBAVFILTER_VERSION_MAJOR = C.LIBAVFILTER_VERSION_MAJOR
)
