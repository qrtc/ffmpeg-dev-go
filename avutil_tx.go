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

// AVComplexInt32
type AVComplexInt32 C.struct_AVComplexInt32

// AVTXType
type AVTXType = C.enum_AVTXType

const (
	AV_TX_FLOAT_FFT  = AVTXType(C.AV_TX_FLOAT_FFT)
	AV_TX_FLOAT_MDCT = AVTXType(C.AV_TX_FLOAT_MDCT)
)

// typedef void (*av_tx_fn)(AVTXContext *s, void *out, void *in, ptrdiff_t stride);
type AvTxFn = C.av_tx_fn

// AvTxInit initializes a transform context with the given configuration
// (i)MDCTs with an odd length are currently not supported.
func AvTxInit(ctx **AVTXContext, tx *AvTxFn, _type AVTXType,
	inv, len int32, scale CVoidPointer, flags uint64) int32 {
	return (int32)(C.av_tx_init((**C.struct_AVTXContext)(unsafe.Pointer(ctx)),
		(*C.av_tx_fn)(tx), (C.enum_AVTXType)(_type),
		(C.int)(inv), (C.int)(len), VoidPointer(scale), (C.uint64_t)(flags)))
}

// AvTxUninit frees a context and sets ctx to NULL, does nothing when ctx == NULL
func AvTxUninit(ctx **AVTXContext) {
	C.av_tx_uninit((**C.struct_AVTXContext)(unsafe.Pointer(ctx)))
}
