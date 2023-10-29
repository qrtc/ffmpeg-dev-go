// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libswresample/version.h>
*/
import "C"

const (
	LIBSWRESAMPLE_VERSION_MAJOR = C.LIBSWRESAMPLE_VERSION_MAJOR
	LIBSWRESAMPLE_VERSION_MINOR = C.LIBSWRESAMPLE_VERSION_MINOR
	LIBSWRESAMPLE_VERSION_MICRO = C.LIBSWRESAMPLE_VERSION_MICRO
)
