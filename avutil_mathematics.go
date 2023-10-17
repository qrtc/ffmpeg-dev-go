package ffmpeg

/*
#include <libavutil/mathematics.h>
*/
import "C"

type AvRounding = C.enum_AVRounding

const (
	AV_ROUND_ZERO        = AvRounding(C.AV_ROUND_ZERO)
	AV_ROUND_INF         = AvRounding(C.AV_ROUND_INF)
	AV_ROUND_DOWN        = AvRounding(C.AV_ROUND_DOWN)
	AV_ROUND_UP          = AvRounding(C.AV_ROUND_UP)
	AV_ROUND_NEAR_INF    = AvRounding(C.AV_ROUND_NEAR_INF)
	AV_ROUND_PASS_MINMAX = AvRounding(C.AV_ROUND_PASS_MINMAX)
)

// AvGcd computes the greatest common divisor of two integer operands.
func AvGcd(a, b int64) int64 {
	return (int64)(C.av_gcd((C.int64_t)(a), (C.int64_t)(b)))
}

// AvRescale rescale a 64-bit integer with rounding to nearest.
func AvRescale(a, b, c int64) int64 {
	return (int64)(C.av_rescale((C.int64_t)(a), (C.int64_t)(b), (C.int64_t)(c)))
}

// AvRescaleRnd rescales a 64-bit integer with specified rounding.
func AvRescaleRnd(a, b, c int64, rnd AvRounding) int64 {
	return (int64)(C.av_rescale_rnd((C.int64_t)(a), (C.int64_t)(b), (C.int64_t)(c),
		(C.enum_AVRounding)(rnd)))
}

// AvRescaleQ rescales a 64-bit integer by 2 rational numbers.
func AvRescaleQ(a int64, bq, cq AvRational) int64 {
	return (int64)(C.av_rescale_q((C.int64_t)(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq)))
}

// AvRescaleQRnd rescales a 64-bit integer by 2 rational numbers with specified rounding.
func AvRescaleQRnd(a int64, bq, cq AvRational, rnd AvRounding) int64 {
	return (int64)(C.av_rescale_q_rnd((C.int64_t)(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq),
		(C.enum_AVRounding)(rnd)))
}

// AvCompareTs compares two timestamps each in its own time base.
func AvCompareTs(tsA int64, tbA AvRational, tsB int64, tbB AvRational) int32 {
	return (int32)(C.av_compare_ts((C.int64_t)(tsA), (C.struct_AVRational)(tbA),
		(C.int64_t)(tsB), (C.struct_AVRational)(tbB)))
}

// AvCompareMod compares the remainders of two integer operands divided by a common divisor.
func AvCompareMod(a, b, mod uint64) int64 {
	return (int64)(C.av_compare_mod((C.uint64_t)(a), (C.uint64_t)(b), (C.uint64_t)(mod)))
}

// AvRescaleDelta rescales a timestamp while preserving known durations.
func AvRescaleDelta(inTb AvRational, inTs int64, fsTb AvRational,
	duration int32, last *int64, outTb AvRational) int64 {
	return (int64)(C.av_rescale_delta((C.struct_AVRational)(inTb),
		(C.int64_t)(inTs),
		(C.struct_AVRational)(fsTb),
		(C.int)(duration),
		(*C.int64_t)(last),
		(C.struct_AVRational)(fsTb)))
}

// AvAddStable adds a value to a timestamp.
func AvAddStable(tsTb AvRational, ts int64, incTb AvRational, inc int64) int32 {
	return (int32)(C.av_add_stable((C.struct_AVRational)(tsTb), (C.int64_t)(ts),
		(C.struct_AVRational)(incTb), (C.int64_t)(inc)))
}
