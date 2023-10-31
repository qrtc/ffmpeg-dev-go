// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/common.h>
*/
import "C"

// AV_NE
func AV_NE[T any](be, le T) T {
	if C.AV_HAVE_BIGENDIAN > 0 {
		return be
	}
	return le
}

// RSHIFT
func RSHIFT[U, V Integer](a U, b V) U {
	if a > 0 {
		return ((a) + ((1 << (b)) >> 1)) >> b
	}
	return ((a) + ((1 << (b)) >> 1) - 1) >> b
}

// ROUNDED_DIV
func ROUNDED_DIV[T Integer](a, b T) T {
	if a >= 0 {
		return ((a) + ((b) >> 1)) / b
	}
	return ((a) - ((b) >> 1)) / b
}

// AV_CEIL_RSHIFT
func AV_CEIL_RSHIFT[U, V Integer](a U, b V) U {
	return -((-(a)) >> (b))
}

// FF_CEIL_RSHIFT
func FF_CEIL_RSHIFT[U, V Integer](a U, b V) U {
	return AV_CEIL_RSHIFT(a, b)
}

// FFUDIV
func FFUDIV[T Integer](a, b T) T {
	if a > 0 {
		return a / b
	}
	return ((a) - (b) + 1) / b
}

// FFUMOD
func FFUMOD[T Integer](a, b T) T {
	return a - b*FFUDIV(a, b)
}

// FFABS
func FFABS[T SingedInteger](a T) T {
	if a >= 0 {
		return a
	}
	return -a
}

// FFSIGNT
func FFSIGNT[T SingedInteger](a T) T {
	if a > 0 {
		return 1
	}
	return -1
}

// FFNABS
func FFNABS[T SingedInteger](a T) T {
	if a <= 0 {
		return a
	}
	return -a
}

// FFABSU
func FFABSU[T SingedInteger](a T) uint32 {
	if a <= 0 {
		return -(uint32)(a)
	}
	return (uint32)(a)
}

// FFABS64U
func FFABS64U[T SingedInteger](a T) uint64 {
	if a <= 0 {
		return -(uint64)(a)
	}
	return (uint64)(a)
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

// AvLog2
func AvLog2(v uint32) int32 {
	return (int32)(C.av_log2((C.uint)(v)))
}

// AvLog2With16bit
func AvLog2With16bit(v uint32) int32 {
	return (int32)(C.av_log2_16bit((C.uint)(v)))
}

// AvClipC clips a signed integer value into the amin-amax range.
func AvClipC(a, amin, amax int32) int32 {
	return (int32)(C.av_clip_c((C.int)(a), (C.int)(amin), (C.int)(amax)))
}

// AvClip64C clips a signed 64bit integer value into the amin-amax range.
func AvClip64C(a, amin, amax int64) int64 {
	return (int64)(C.av_clip64_c((C.int64_t)(a), (C.int64_t)(amin), (C.int64_t)(amax)))
}

// AvClipUint8C clips a signed integer value into the 0-255 range.
func AvClipUint8C(a int32) uint8 {
	return (uint8)(C.av_clip_uint8_c((C.int)(a)))
}

// AvClipInt8C clips a signed integer value into the -128,127 range.
func AvClipInt8C(a int32) int8 {
	return (int8)(C.av_clip_int8_c((C.int)(a)))
}

// AvClipUint16C clips a signed integer value into the 0-65535 range.
func AvClipUint16C(a int32) uint16 {
	return (uint16)(C.av_clip_uint16_c((C.int)(a)))
}

// AvClipInt16C clips a signed integer value into the -32768,32767 range.
func AvClipInt16C(a int32) int16 {
	return (int16)(C.av_clip_int16_c((C.int)(a)))
}

// AvCliplInt32C clips a signed 64-bit integer value into the -2147483648,2147483647 range.
func AvCliplInt32C(a int64) int32 {
	return (int32)(C.av_clipl_int32_c((C.int64_t)(a)))
}

// AvClipIntp2C clips a signed integer into the -(2^p),(2^p-1) range.
func AvClipIntp2C(a, p int32) int32 {
	return (int32)(C.av_clip_intp2_c((C.int)(a), (C.int)(p)))
}

// AvClipUintp2C clips a signed integer to an unsigned power of two range.
func AvClipUintp2C(a, p int32) uint32 {
	return (uint32)(C.av_clip_uintp2_c((C.int)(a), (C.int)(p)))
}

// AvModUintp2C clears high bits from an unsigned integer starting with specific bit position.
func AvModUintp2C(a, p uint32) uint32 {
	return (uint32)(C.av_mod_uintp2_c((C.uint)(a), (C.uint)(p)))
}

// AvSatAdd32C adds two signed 32-bit values with saturation.
func AvSatAdd32C(a, b int32) int32 {
	return (int32)(C.av_sat_add32_c((C.int)(a), (C.int)(b)))
}

// AvSatDadd32C adds a doubled value to another value with saturation at both stages.
func AvSatDadd32C(a, b int32) int32 {
	return (int32)(C.av_sat_dadd32_c((C.int)(a), (C.int)(b)))
}

// AvSatSub32C subtracts two signed 32-bit values with saturation.
func AvSatSub32C(a, b int32) int32 {
	return (int32)(C.av_sat_sub32_c((C.int)(a), (C.int)(b)))
}

// AvSatDsub32C subtracts a doubled value from another value with saturation at both stages.
func AvSatDsub32C(a, b int32) int32 {
	return (int32)(C.av_sat_dsub32_c((C.int)(a), (C.int)(b)))
}

// AvClipfC clip a float value into the amin-amax range.
func AvClipfC(a, amin, amax float32) float32 {
	return (float32)(C.av_clipf_c((C.float)(a), (C.float)(amin), (C.float)(amax)))
}

// AvClipdC clips a double value into the amin-amax range.
func AvClipdC(a, amin, amax float64) float64 {
	return (float64)(C.av_clipd_c((C.double)(a), (C.double)(amin), (C.double)(amax)))
}

// AvCeilLog2C computes ceil(log2(x)).
func AvCeilLog2C(x int32) int32 {
	return (int32)(C.av_ceil_log2_c((C.int)(x)))
}

// AvPopcountC counts number of bits set to one in x.
func AvPopcountC(x uint32) int32 {
	return (int32)(C.av_popcount_c((C.uint)(x)))
}

// AvPopcount64C counts number of bits set to one in x.
func AvPopcount64C(x uint64) int32 {
	return (int32)(C.av_popcount64_c((C.uint64_t)(x)))
}

// AvParityC
func AvParityC(x uint32) int32 {
	return (int32)(C.av_parity_c((C.uint32_t)(x)))
}

// MKTAG
func MKTAG(a, b, c, d uint32) uint32 {
	return (a) | ((b) << 8) | ((c) << 16) | ((uint32)(d) << 24)
}

// MKBETAG
func MKBETAG(a, b, c, d uint32) uint32 {
	return (d) | ((c) << 8) | ((b) << 16) | ((uint32)(a) << 24)
}

// See https://pkg.go.dev/unicode/utf8

// NONEED: GET_UTF8

// NONEED: GET_UTF16

// NONEED: PUT_UTF8

// NONEED: PUT_UTF16
