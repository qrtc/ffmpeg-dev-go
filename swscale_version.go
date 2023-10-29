// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libswscale/version.h>
*/
import "C"

const (
	LIBSWSCALE_VERSION_MAJOR = C.LIBSWSCALE_VERSION_MAJOR
	LIBSWSCALE_VERSION_MINOR = C.LIBSWSCALE_VERSION_MINOR
	LIBSWSCALE_VERSION_MICRO = C.LIBSWSCALE_VERSION_MICRO
)
