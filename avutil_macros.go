// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/macros.h>
*/
import "C"

func FFALIGN[T Integer](x, a T) T {
	return ((x) + (a) - 1) & ^((a) - 1)
}
