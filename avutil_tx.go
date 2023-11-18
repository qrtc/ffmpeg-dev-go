// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/tx.h>
*/
import "C"
import "unsafe"

// AVTXContext
type AVTXContext C.struct_AVTXContext

// AVComplexFloat
type AVComplexFloat C.struct_AVComplexFloat

// Make new AVComplexFloat.
func AvMakeComxFloat(re, im float32) AVComplexFloat {
	return (AVComplexFloat)(C.struct_AVComplexFloat{
		re: (C.float)(re),
		im: (C.float)(im)})
}

// GetRe gets `AVComplexFloat.re` value.
func (cf *AVComplexFloat) GetRe() float32 {
	return (float32)(cf.re)
}

// SetRe sets `AVComplexFloat.re` value.
func (cf *AVComplexFloat) SetRe(v float32) {
	cf.re = (C.float)(v)
}

// GetReAddr gets `AVComplexFloat.re` address.
func (cf *AVComplexFloat) GetReAddr() *float32 {
	return (*float32)(&cf.re)
}

// GetIm gets `AVComplexFloat.im` value.
func (cf *AVComplexFloat) GetIm() float32 {
	return (float32)(cf.im)
}

// SetIm sets `AVComplexFloat.im` value.
func (cf *AVComplexFloat) SetIm(v float32) {
	cf.im = (C.float)(v)
}

// GetImAddr gets `AVComplexFloat.im` address.
func (cf *AVComplexFloat) GetImAddr() *float32 {
	return (*float32)(&cf.im)
}

// AVComplexDouble
type AVComplexDouble C.struct_AVComplexDouble

// Make new AVComplexDouble.
func AvMakeComxDouble(re, im float64) AVComplexDouble {
	return (AVComplexDouble)(C.struct_AVComplexDouble{
		re: (C.double)(re),
		im: (C.double)(im)})
}

// GetRe gets `AVComplexDouble.re` value.
func (cd *AVComplexDouble) GetRe() float64 {
	return (float64)(cd.re)
}

// SetRe sets `AVComplexDouble.re` value.
func (cd *AVComplexDouble) SetRe(v float64) {
	cd.re = (C.double)(v)
}

// GetReAddr gets `AVComplexDouble.re` address.
func (cd *AVComplexDouble) GetReAddr() *float64 {
	return (*float64)(&cd.re)
}

// GetIm gets `AVComplexDouble.im` value.
func (cd *AVComplexDouble) GetIm() float64 {
	return (float64)(cd.im)
}

// SetIm sets `AVComplexDouble.im` value.
func (cd *AVComplexDouble) SetIm(v float64) {
	cd.im = (C.double)(v)
}

// GetImAddr gets `AVComplexDouble.im` address.
func (cd *AVComplexDouble) GetImAddr() *float64 {
	return (*float64)(&cd.im)
}

// AVComplexInt32
type AVComplexInt32 C.struct_AVComplexInt32

// Make new AVComplexFloat.
func AvMakeComxInt32(re, im int32) AVComplexInt32 {
	return (AVComplexInt32)(C.struct_AVComplexInt32{
		re: (C.int32_t)(re),
		im: (C.int32_t)(im)})
}

// GetRe gets `AVComplexInt32.re` value.
func (ci *AVComplexInt32) GetRe() int32 {
	return (int32)(ci.re)
}

// SetRe sets `AVComplexInt32.re` value.
func (ci *AVComplexInt32) SetRe(v int32) {
	ci.re = (C.int32_t)(v)
}

// GetReAddr gets `AVComplexInt32.re` address.
func (ci *AVComplexInt32) GetReAddr() *int32 {
	return (*int32)(&ci.re)
}

// GetIm gets `AVComplexInt32.im` value.
func (ci *AVComplexInt32) GetIm() int32 {
	return (int32)(ci.im)
}

// SetIm sets `AVComplexInt32.im` value.
func (ci *AVComplexInt32) SetIm(v int32) {
	ci.im = (C.int32_t)(v)
}

// GetImAddr gets `AVComplexInt32.im` address.
func (ci *AVComplexInt32) GetImAddr() *int32 {
	return (*int32)(&ci.im)
}

// AVTXType
type AVTXType = C.enum_AVTXType

const (
	AV_TX_FLOAT_FFT   = AVTXType(C.AV_TX_FLOAT_FFT)
	AV_TX_FLOAT_MDCT  = AVTXType(C.AV_TX_FLOAT_MDCT)
	AV_TX_DOUBLE_FFT  = AVTXType(C.AV_TX_DOUBLE_FFT)
	AV_TX_DOUBLE_MDCT = AVTXType(C.AV_TX_DOUBLE_MDCT)
	AV_TX_INT32_FFT   = AVTXType(C.AV_TX_INT32_FFT)
	AV_TX_INT32_MDCT  = AVTXType(C.AV_TX_INT32_MDCT)

	AV_TX_FLOAT_RDFT  = AVTXType(C.AV_TX_FLOAT_RDFT)
	AV_TX_DOUBLE_RDFT = AVTXType(C.AV_TX_DOUBLE_RDFT)
	AV_TX_INT32_RDFT  = AVTXType(C.AV_TX_INT32_RDFT)

	AV_TX_FLOAT_DCT  = AVTXType(C.AV_TX_FLOAT_DCT)
	AV_TX_DOUBLE_DCT = AVTXType(C.AV_TX_DOUBLE_DCT)
	AV_TX_INT32_DCT  = AVTXType(C.AV_TX_INT32_DCT)

	AV_TX_NB = AVTXType(C.AV_TX_NB)
)

// typedef void (*av_tx_fn)(AVTXContext *s, void *out, void *in, ptrdiff_t stride);
type AvTxFn = C.av_tx_fn

// AVTXFlags
type AVTXFlags = C.enum_AVTXFlags

const (
	AV_TX_INPLACE           = AVTXFlags(C.AV_TX_INPLACE)
	AV_TX_UNALIGNED         = AVTXFlags(C.AV_TX_UNALIGNED)
	AV_TX_FULL_IMDCT        = AVTXFlags(C.AV_TX_FULL_IMDCT)
	AV_TX_REAL_TO_REAL      = AVTXFlags(C.AV_TX_REAL_TO_REAL)
	AV_TX_REAL_TO_IMAGINARY = AVTXFlags(C.AV_TX_REAL_TO_IMAGINARY)
)

// AvTxInit initializes a transform context with the given configuration
// (i)MDCTs with an odd length are currently not supported.
func AvTxInit(ctx **AVTXContext, tx *AvTxFn, _type AVTXType,
	inv, len int32, scale CVoidPointer, flags uint64) int32 {
	return (int32)(C.av_tx_init((**C.struct_AVTXContext)(unsafe.Pointer(ctx)),
		(*C.av_tx_fn)(tx), (C.enum_AVTXType)(_type),
		(C.int)(inv), (C.int)(len), VoidPointer(scale), (C.uint64_t)(flags)))
}

// AvTxUninit frees a context and sets ctx to NULL, does nothing when ctx == NULL.
func AvTxUninit(ctx **AVTXContext) {
	C.av_tx_uninit((**C.struct_AVTXContext)(unsafe.Pointer(ctx)))
}
