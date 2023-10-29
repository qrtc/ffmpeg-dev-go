// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/display.h>
*/
import "C"

// AvDisplayRotationGet extracts the rotation component of the transformation matrix.
func AvDisplayRotationGet(matrix []int32) float64 {
	if len(matrix) < 9 {
		panic("matrix len < 9")
	}
	return (float64)(C.av_display_rotation_get((*C.int32_t)(&matrix[0])))
}

// AvDisplayRotationSet initializes a transformation matrix describing a pure counterclockwise
// rotation by the specified angle (in degrees).
func AvDisplayRotationSet(matrix []int32, angle float64) {
	C.av_display_rotation_set((*C.int32_t)(&matrix[0]), (C.double)(angle))
}

// AvDisplayMatrixFlip flips the input matrix horizontally and/or vertically.
func AvDisplayMatrixFlip(matrix []int32, hflip, vflip int32) {
	C.av_display_matrix_flip((*C.int32_t)(&matrix[0]), (C.int)(hflip), (C.int)(vflip))
}
