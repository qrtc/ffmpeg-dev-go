// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/rational.h>
*/
import "C"

// AVRational
type AVRational C.struct_AVRational

// GetNum gets `AVRational.num` value.
func (q *AVRational) GetNum() int32 {
	return (int32)(q.num)
}

// GetDen gets `AVRational.den` value.
func (q *AVRational) GetDen() int32 {
	return (int32)(q.den)
}

// AvMakeQ creates an AVRational with numerator and denominator.
func AvMakeQ(num, den int32) AVRational {
	return (AVRational)(C.av_make_q((C.int)(num), (C.int)(den)))
}

const (
	INT_MIN = int32(C.INT_MIN)
)

// AvCmpQ compares two rationals.
// returns One of the following values:
// 0 if `a == b`
// 1 if `a > b`
// -1 if `a < b`
// `INT_MIN` if one of the values is of the form `0 / 0`
func AvCmpQ(a, b AVRational) int32 {
	return (int32)(C.av_cmp_q((C.struct_AVRational)(a), (C.struct_AVRational)(b)))
}

// AvQ2d converts an AVRational to a `float64`.
func AvQ2d(a AVRational) float64 {
	return (float64)(C.av_q2d((C.struct_AVRational)(a)))
}

// AvReduce reduces a fraction.
func AvReduce(dstNum, dstDen *int32, num, den, max int64) int32 {
	return (int32)(C.av_reduce((*C.int)(dstNum), (*C.int)(dstDen),
		(C.int64_t)(num), (C.int64_t)(den), (C.int64_t)(max)))
}

// AvMulQ multiplies two rationals.
func AvMulQ(a, b AVRational) AVRational {
	return (AVRational)(C.av_mul_q((C.struct_AVRational)(a), (C.struct_AVRational)(b)))
}

// AvDivQ divides one rational by another.
func AvDivQ(a, b AVRational) AVRational {
	return (AVRational)(C.av_div_q((C.struct_AVRational)(a), (C.struct_AVRational)(b)))
}

// AvAddQ adds two rationals.
func AvAddQ(a, b AVRational) AVRational {
	return (AVRational)(C.av_add_q((C.struct_AVRational)(a), (C.struct_AVRational)(b)))
}

// AvSubQ subtracts one rational from another.
func AvSubQ(a, b AVRational) AVRational {
	return (AVRational)(C.av_sub_q((C.struct_AVRational)(a), (C.struct_AVRational)(b)))
}

// AvInvQ invert a rational.
// return 1 / q
func AvInvQ(q AVRational) AVRational {
	return (AVRational)(C.av_inv_q((C.struct_AVRational)(q)))
}

// AvD2Q converts a double precision floating point number to a rational.
func AvD2Q(d float64, max int32) AVRational {
	return (AVRational)(C.av_d2q((C.double)(d), (C.int)(max)))
}

// AvNearerQ finds which of the two rationals is closer to another rational.
// return One of the following values:
// 1 if `q1` is nearer to `q` than `q2`
// -1 if `q2` is nearer to `q` than `q1`
// 0 if they have the same distance.
func AvNearerQ(q, q1, q2 AVRational) int32 {
	return (int32)(C.av_nearer_q((C.struct_AVRational)(q),
		(C.struct_AVRational)(q1), (C.struct_AVRational)(q2)))
}

// AvFindNearestQIdx finds the value in a list of rationals nearest a given reference rational.
func AvFindNearestQIdx(q AVRational, qList *AVRational) int32 {
	return (int32)(C.av_find_nearest_q_idx((C.struct_AVRational)(q), (*C.struct_AVRational)(qList)))
}

// AvQ2intfloat Convert an AVRational to a IEEE 32-bit `float` expressed in fixed-point format.
func AvQ2intfloat(q AVRational) uint32 {
	return (uint32)(C.av_q2intfloat((C.struct_AVRational)(q)))
}

// AvGcdQ returns the best rational so that a and b are multiple of it.
// If the resulting denominator is larger than max_den, return def.
func AvGcdQ(a, b AVRational, maxDen int32, def AVRational) AVRational {
	return (AVRational)(C.av_gcd_q((C.struct_AVRational)(a), (C.struct_AVRational)(b),
		(C.int)(maxDen), (C.struct_AVRational)(def)))
}
