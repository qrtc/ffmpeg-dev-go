package ffmpeg

/*
#include <libavresample/avresample.h>
*/
import "C"
import "unsafe"

const (
	AVRESAMPLE_MAX_CHANNELS = C.AVRESAMPLE_MAX_CHANNELS
)

type AvAudioResampleContext C.struct_AVAudioResampleContext

// Deprecated: Use libswresample
type AvMixCoeffType int32

const (
	AV_MIX_COEFF_TYPE_Q8  = AvMixCoeffType(C.AV_MIX_COEFF_TYPE_Q8)
	AV_MIX_COEFF_TYPE_Q15 = AvMixCoeffType(C.AV_MIX_COEFF_TYPE_Q15)
	AV_MIX_COEFF_TYPE_FLT = AvMixCoeffType(C.AV_MIX_COEFF_TYPE_FLT)
	AV_MIX_COEFF_TYPE_NB  = AvMixCoeffType(C.AV_MIX_COEFF_TYPE_NB)
)

// Deprecated: Use libswresample
type AvResampleFilterType int32

const (
	AV_RESAMPLE_FILTER_TYPE_CUBIC            = AvResampleFilterType(C.AV_RESAMPLE_FILTER_TYPE_CUBIC)
	AV_RESAMPLE_FILTER_TYPE_BLACKMAN_NUTTALL = AvResampleFilterType(C.AV_RESAMPLE_FILTER_TYPE_BLACKMAN_NUTTALL)
	AV_RESAMPLE_FILTER_TYPE_KAISER           = AvResampleFilterType(C.AV_RESAMPLE_FILTER_TYPE_KAISER)
)

type AvResampleDitherMethod int32

const (
	AV_RESAMPLE_DITHER_NONE          = AvResampleDitherMethod(C.AV_RESAMPLE_DITHER_NONE)
	AV_RESAMPLE_DITHER_RECTANGULAR   = AvResampleDitherMethod(C.AV_RESAMPLE_DITHER_RECTANGULAR)
	AV_RESAMPLE_DITHER_TRIANGULAR    = AvResampleDitherMethod(C.AV_RESAMPLE_DITHER_TRIANGULAR)
	AV_RESAMPLE_DITHER_TRIANGULAR_HP = AvResampleDitherMethod(C.AV_RESAMPLE_DITHER_TRIANGULAR_HP)
	AV_RESAMPLE_DITHER_TRIANGULAR_NS = AvResampleDitherMethod(C.AV_RESAMPLE_DITHER_TRIANGULAR_NS)
	AV_RESAMPLE_DITHER_NB            = AvResampleDitherMethod(C.AV_RESAMPLE_DITHER_NB)
)

// Deprecated: Use libswresample
func AvResampleVersion() uint32 {
	return (uint32)(C.avresample_version())
}

// Deprecated: Use libswresample
func AvResampleConfiguration() string {
	return C.GoString(C.avresample_configuration())
}

// Deprecated: Use libswresample
func AvResampleLicense() string {
	return C.GoString(C.avresample_license())
}

// Deprecated: Use libswresample
func AvResampleGetClass() *AvClass {
	return (*AvClass)(C.avresample_get_class())
}

// Deprecated: Use libswresample
func AvResampleAllocContext() *AvAudioResampleContext {
	return (*AvAudioResampleContext)(C.avresample_alloc_context())
}

// Deprecated: Use libswresample
func AvResampleOpen(avr *AvAudioResampleContext) int32 {
	return (int32)(C.avresample_open((*C.struct_AVAudioResampleContext)(avr)))
}

// Deprecated: Use libswresample
func AvResampleIsOpen(avr *AvAudioResampleContext) int32 {
	return (int32)(C.avresample_is_open((*C.struct_AVAudioResampleContext)(avr)))
}

// Deprecated: Use libswresample
func AvResampleClose(avr *AvAudioResampleContext) {
	C.avresample_close((*C.struct_AVAudioResampleContext)(avr))
}

// Deprecated: Use libswresample
func AvResampleFree(avr **AvAudioResampleContext) {
	C.avresample_free((**C.struct_AVAudioResampleContext)(unsafe.Pointer(avr)))
}

// Deprecated: Use libswresample
func AvResampleBuildMatrix(inLayout, outLayout uint64,
	centerMixLevel, surroundMixLevel, lfeMixLevel float64,
	normalize int32, matrix *float64, stride int32, matrixEncoding AvMatrixEncoding) int32 {
	return (int32)(C.avresample_build_matrix((C.uint64_t)(inLayout), (C.uint64_t)(outLayout),
		(C.double)(centerMixLevel), (C.double)(surroundMixLevel), (C.double)(lfeMixLevel),
		(C.int)(normalize), (*C.double)(matrix), (C.int)(stride),
		(C.enum_AVMatrixEncoding)(matrixEncoding)))
}

// Deprecated: Use libswresample
func AvResampleGetMatrix(avr *AvAudioResampleContext, matrix *float64, stride int32) int32 {
	return (int32)(C.avresample_get_matrix((*C.struct_AVAudioResampleContext)(avr),
		(*C.double)(matrix), (C.int)(stride)))
}

// Deprecated: Use libswresample
func AvResampleSetMatrix(avr *AvAudioResampleContext, matrix *float64, stride int32) int32 {
	return (int32)(C.avresample_set_matrix((*C.struct_AVAudioResampleContext)(avr),
		(*C.double)(matrix), (C.int)(stride)))
}

// Deprecated: Use libswresample
func AvResampleSetChannelMapping(avr *AvAudioResampleContext, channelMap *int32) int32 {
	return (int32)(C.avresample_set_channel_mapping((*C.struct_AVAudioResampleContext)(avr),
		(*C.int)(channelMap)))
}

// Deprecated: Use libswresample
func AvResampleSetCompensation(avr *AvAudioResampleContext, sampleDelta, compensationDistance int32) int32 {
	return (int32)(C.avresample_set_compensation((*C.struct_AVAudioResampleContext)(avr),
		(C.int)(sampleDelta), (C.int)(compensationDistance)))
}

// Deprecated: Use libswresample
func AvResampleGetOutSamples(avr *AvAudioResampleContext, inNbSamples int32) int32 {
	return (int32)(C.avresample_get_out_samples((*C.struct_AVAudioResampleContext)(avr),
		(C.int)(inNbSamples)))
}

// Deprecated: Use libswresample
func AvResampleConvert(avr *AvAudioResampleContext, output **uint8, outPlaneSize, outSamples int32,
	input **uint8, inPlaneSize, inSamples int32) int32 {
	return (int32)(C.avresample_convert((*C.struct_AVAudioResampleContext)(avr),
		(**C.uint8_t)(unsafe.Pointer(output)), (C.int)(outPlaneSize), (C.int)(outSamples),
		(**C.uint8_t)(unsafe.Pointer(input)), (C.int)(inPlaneSize), (C.int)(inSamples)))
}

// Deprecated: Use libswresample
func AvResampleGetDelay(avr *AvAudioResampleContext) int32 {
	return (int32)(C.avresample_get_delay((*C.struct_AVAudioResampleContext)(avr)))
}

// Deprecated: Use libswresample
func AvResampleAvailable(avr *AvAudioResampleContext) int32 {
	return (int32)(C.avresample_available((*C.struct_AVAudioResampleContext)(avr)))
}

// Deprecated: Use libswresample
func AvResampleRead(avr *AvAudioResampleContext, output **uint8, nbSamples int32) int32 {
	return (int32)(C.avresample_read((*C.struct_AVAudioResampleContext)(avr),
		(**C.uint8_t)(unsafe.Pointer(output)), (C.int)(nbSamples)))
}

// Deprecated: Use libswresample
func AvResampleConvertFrame(avr *AvAudioResampleContext, output, input *AvFrame) int32 {
	return (int32)(C.avresample_convert_frame((*C.struct_AVAudioResampleContext)(avr),
		(*C.struct_AVFrame)(output), (*C.struct_AVFrame)(input)))
}

// Deprecated: Use libswresample
func AvResampleConfig(avr *AvAudioResampleContext, out, in *AvFrame) int32 {
	return (int32)(C.avresample_config((*C.struct_AVAudioResampleContext)(avr),
		(*C.struct_AVFrame)(out), (*C.struct_AVFrame)(in)))
}
