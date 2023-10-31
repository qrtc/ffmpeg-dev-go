// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavcodec/avfft.h>
*/
import "C"

// FFTSample
type FFTSample = C.FFTSample

// FFTComplex
type FFTComplex C.struct_FFTComplex

// GetRe gets `FFTComplex.re` value.
func (fc *FFTComplex) GetRe() FFTSample {
	return (FFTSample)(fc.re)
}

// GetIm gets `FFTComplex.im` value.
func (fc *FFTComplex) GetIm() FFTSample {
	return (FFTSample)(fc.im)
}

// FFTContext
type FFTContext C.struct_FFTContext

// AvFftInit sets up a complex FFT.
func AvFftInit(nbits, inverse int32) *FFTContext {
	return (*FFTContext)(C.av_fft_init((C.int)(nbits), (C.int)(inverse)))
}

// AvFftPermute does the permutation needed BEFORE calling FfFftCalc().
func AvFftPermute(s *FFTContext, z *FFTComplex) {
	C.av_fft_permute((*C.struct_FFTContext)(s), (*C.struct_FFTComplex)(z))
}

// FfFftCalc does a complex FFT with the parameters defined in AvFftInit().
func FfFftCalc(s *FFTContext, z *FFTComplex) {
	C.av_fft_calc((*C.struct_FFTContext)(s), (*C.struct_FFTComplex)(z))
}

// AvFftEnd
func AvFftEnd(s *FFTContext) {
	C.av_fft_end((*C.struct_FFTContext)(s))
}

// RDFTransformType
type RDFTransformType = C.enum_RDFTransformType

const (
	DFT_R2C  = RDFTransformType(C.DFT_R2C)
	IDFT_C2R = RDFTransformType(C.IDFT_C2R)
	IDFT_R2C = RDFTransformType(C.IDFT_R2C)
	DFT_C2R  = RDFTransformType(C.DFT_C2R)
)

// RDFTContext
type RDFTContext C.struct_RDFTContext

// AvRdftInit
func AvRdftInit(nbits int32, trans RDFTransformType) *RDFTContext {
	return (*RDFTContext)(C.av_rdft_init((C.int)(nbits),
		(C.enum_RDFTransformType)(trans)))
}

// AvRdftCalc
func AvRdftCalc(r *RDFTContext, data *FFTSample) {
	C.av_rdft_calc((*C.struct_RDFTContext)(r), (*C.FFTSample)(data))
}

// AvRdftEnd
func AvRdftEnd(r *RDFTContext) {
	C.av_rdft_end((*C.struct_RDFTContext)(r))
}

// DCTContext
type DCTContext C.struct_DCTContext

// DCTTransformType
type DCTTransformType = C.enum_DCTTransformType

const (
	DCT_II  = DCTTransformType(C.DCT_II)
	DCT_III = DCTTransformType(C.DCT_III)
	DCT_I   = DCTTransformType(C.DCT_I)
	DST_I   = DCTTransformType(C.DST_I)
)

// AvDctInit
func AvDctInit(nbits int32, _type RDFTransformType) *DCTContext {
	return (*DCTContext)(C.av_dct_init((C.int)(nbits),
		(C.enum_RDFTransformType)(_type)))
}

// AvDctCalc
func AvDctCalc(d *DCTContext, data *FFTSample) {
	C.av_dct_calc((*C.struct_DCTContext)(d), (*C.FFTSample)(data))
}

// AvDctEnd
func AvDctEnd(d *DCTContext) {
	C.av_dct_end((*C.struct_DCTContext)(d))
}
