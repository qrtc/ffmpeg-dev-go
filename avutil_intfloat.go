// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/intfloat.h>
*/
import "C"

// AvInt2float reinterprets a 32-bit integer as a float.
func AvInt2float(i uint32) float32 {
	return (float32)(C.av_int2float((C.uint32_t)(i)))
}

// AvFloat2int reinterprets a float as a 32-bit integer.
func AvFloat2int(f float32) uint32 {
	return (uint32)(C.av_float2int((C.float)(f)))
}

// AvInt2double reinterprets a 64-bit integer as a double.
func AvInt2double(i uint64) float64 {
	return (float64)(C.av_int2double((C.uint64_t)(i)))
}

// AvDouble2int reinterprets a double as a 64-bit integer.
func AvDouble2int(f float64) uint64 {
	return (uint64)(C.av_double2int((C.double)(f)))
}
