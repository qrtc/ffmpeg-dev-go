// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/pixelutils.h>

*/
import "C"

// AVPixelutilsSadFn
type AVPixelutilsSadFn C.av_pixelutils_sad_fn

// AvPixelutilsGetSadFn gets a potentially optimized pointer to a Sum-of-absolute-differences
// function (see the av_pixelutils_sad_fn prototype).
func AvPixelutilsGetSadFn(wBits, hBits, aligned int32, logCtx CVoidPointer) AVPixelutilsSadFn {
	return (AVPixelutilsSadFn)(C.av_pixelutils_get_sad_fn((C.int)(wBits), (C.int)(hBits),
		(C.int)(aligned), VoidPointer(logCtx)))
}
