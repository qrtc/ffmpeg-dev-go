package ffmpeg

/*
#include <libavutil/common.h>
*/
import "C"

func AV_NE[T any](be, le T) T {
	if C.AV_HAVE_BIGENDIAN > 0 {
		return be
	}
	return le
}

func FFABS[T HelperSingedInteger](a T) T {
	if a >= 0 {
		return a
	}
	return 0 - a
}

func FFSIGNT[T HelperSingedInteger](a T) T {
	if a > 0 {
		return 1
	}
	return -1
}

func FFMAX[T HelperInteger](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func FFMAX3[T HelperInteger](a, b, c T) T {
	return FFMAX(FFMAX(a, b), c)
}

func FFMIN[T HelperInteger](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func FFMIN3[T HelperInteger](a, b, c T) T {
	return FFMIN(FFMIN(a, b), c)
}
