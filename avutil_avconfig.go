// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/avconfig.h>
*/
import "C"

const (
	AV_HAVE_BIGENDIAN      = C.AV_HAVE_BIGENDIAN
	AV_HAVE_FAST_UNALIGNED = C.AV_HAVE_FAST_UNALIGNED
)
