// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/macros.h>
*/
import "C"

// AV_NE
func AV_NE[T any](be, le T) T {
	if C.AV_HAVE_BIGENDIAN > 0 {
		return be
	}
	return le
}

// FFDIFFSIGN
func FFDIFFSIGN[T Integer](x, y T) int {
	if x > y {
		return 1
	} else if x < y {
		return -1
	} else {
		return 0
	}
}

// FFMAX
func FFMAX[T Integer](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// FFMAX3
func FFMAX3[T Integer](a, b, c T) T {
	return FFMAX(FFMAX(a, b), c)
}

// FFMIN
func FFMIN[T Integer](a, b T) T {
	if a > b {
		return b
	}
	return a
}

// FFMIN3
func FFMIN3[T Integer](a, b, c T) T {
	return FFMIN(FFMIN(a, b), c)
}

// FFSWAP
func FFSWAP[T any](a, b *T) {
	swapTmp := *b
	*b = *a
	*a = swapTmp
}

// NONEED: FF_ARRAY_ELEMS

// MKTAG
func MKTAG(a, b, c, d uint32) uint32 {
	return (a) | ((b) << 8) | ((c) << 16) | ((uint32)(d) << 24)
}

// MKBETAG
func MKBETAG(a, b, c, d uint32) uint32 {
	return (d) | ((c) << 8) | ((b) << 16) | ((uint32)(a) << 24)
}

// FFALIGN
func FFALIGN[T Integer](x, a T) T {
	return ((x) + (a) - 1) & ^((a) - 1)
}
