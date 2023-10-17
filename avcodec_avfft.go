package ffmpeg

/*
#include <libavcodec/avfft.h>
*/
import "C"

type FftSample C.FFTSample

type FftComplex C.struct_FFTComplex

type FftContext C.struct_FFTContext

// AvFftInit sets up a complex FFT.
func AvFftInit(nbits, inverse int32) *FftContext {
	return (*FftContext)(C.av_fft_init((C.int)(nbits), (C.int)(inverse)))
}

// AvFftPermute does the permutation needed BEFORE calling FfFftCalc().
func AvFftPermute(s *FftContext, z *FftComplex) {
	C.av_fft_permute((*C.struct_FFTContext)(s), (*C.struct_FFTComplex)(z))
}

// FfFftCalc does a complex FFT with the parameters defined in AvFftInit().
func FfFftCalc(s *FftContext, z *FftComplex) {
	C.av_fft_calc((*C.struct_FFTContext)(s), (*C.struct_FFTComplex)(z))
}

// AvFftEnd
func AvFftEnd(s *FftContext) {
	C.av_fft_end((*C.struct_FFTContext)(s))
}

// AvMdctInit
func AvMdctInit(nbits, inverse int32, scale float64) *FftContext {
	return (*FftContext)(C.av_mdct_init((C.int)(nbits), (C.int)(inverse), (C.double)(scale)))
}

// AvImdctCalc
func AvImdctCalc(s *FftContext, output, input *FftSample) {
	C.av_imdct_calc((*C.struct_FFTContext)(s),
		(*C.FFTSample)(output), (*C.FFTSample)(input))
}

// AvImdctHalf
func AvImdctHalf(s *FftContext, output, input *FftSample) {
	C.av_imdct_half((*C.struct_FFTContext)(s),
		(*C.FFTSample)(output), (*C.FFTSample)(input))
}

// AvMdctCalc
func AvMdctCalc(s *FftContext, output, input *FftSample) {
	C.av_mdct_calc((*C.struct_FFTContext)(s),
		(*C.FFTSample)(output), (*C.FFTSample)(input))
}

// AvMdctEnd
func AvMdctEnd(s *FftContext) {
	C.av_mdct_end((*C.struct_FFTContext)(s))
}

type RDFTransformType = C.enum_RDFTransformType

const (
	DFT_R2C  = RDFTransformType(C.DFT_R2C)
	IDFT_C2R = RDFTransformType(C.IDFT_C2R)
	IDFT_R2C = RDFTransformType(C.IDFT_R2C)
	DFT_C2R  = RDFTransformType(C.DFT_C2R)
)

type RDFTContext C.struct_RDFTContext

// AvRdftInit
func AvRdftInit(nbits int32, trans RDFTransformType) *RDFTContext {
	return (*RDFTContext)(C.av_rdft_init((C.int)(nbits),
		(C.enum_RDFTransformType)(trans)))
}

// AvRdftCalc
func AvRdftCalc(r *RDFTContext, data *FftSample) {
	C.av_rdft_calc((*C.struct_RDFTContext)(r), (*C.FFTSample)(data))
}

// AvRdftEnd
func AvRdftEnd(r *RDFTContext) {
	C.av_rdft_end((*C.struct_RDFTContext)(r))
}

type DCTContext C.struct_DCTContext

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
func AvDctCalc(d *DCTContext, data *FftSample) {
	C.av_dct_calc((*C.struct_DCTContext)(d), (*C.FFTSample)(data))
}

// AvDctEnd
func AvDctEnd(d *DCTContext) {
	C.av_dct_end((*C.struct_DCTContext)(d))
}
