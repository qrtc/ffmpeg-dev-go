// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavformat/version.h>
*/
import "C"

const (
	LIBAVFORMAT_VERSION_MINOR = C.LIBAVFORMAT_VERSION_MINOR
	LIBAVFORMAT_VERSION_MICRO = C.LIBAVFORMAT_VERSION_MICRO
)
