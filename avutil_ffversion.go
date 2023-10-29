// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/ffversion.h>
*/
import "C"

const (
	FFMPEG_VERSION = C.FFMPEG_VERSION
)
