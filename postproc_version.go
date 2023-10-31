// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libpostproc/version.h>
*/
import "C"

const (
	LIBPOSTPROC_VERSION_MINOR = C.LIBPOSTPROC_VERSION_MINOR
	LIBPOSTPROC_VERSION_MICRO = C.LIBPOSTPROC_VERSION_MICRO
)
