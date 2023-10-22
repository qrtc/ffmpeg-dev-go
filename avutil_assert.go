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
