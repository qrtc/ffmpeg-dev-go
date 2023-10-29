// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/avassert.h>
*/
import "C"

// AvAssert0
func AvAssert0(cond bool) {
	if !cond {
		AvLog(nil, AV_LOG_PANIC, "Assertion failed\n")
		panic("AvAssert0 assert failed")
	}
}

// AvAssert0Fpu asserts that floating point operations can be executed.
func AvAssert0Fpu() {
	C.av_assert0_fpu()
}
