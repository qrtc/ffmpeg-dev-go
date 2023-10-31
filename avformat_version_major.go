// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavformat/version_major.h>
*/
import "C"

const (
	LIBAVFORMAT_VERSION_MAJOR = C.LIBAVFORMAT_VERSION_MAJOR
)
