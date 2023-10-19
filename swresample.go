package ffmpeg

/*
#include <libswresample/swresample.h>
*/
import "C"
import "unsafe"

const (
	SWR_FLAG_RESAMPLE = C.SWR_FLAG_RESAMPLE
)

// Dithering algorithms
type SwrDitherType = C.enum_SwrDitherType

const (
	SWR_DITHER_NONE                   = SwrDitherType(C.SWR_DITHER_NONE)
	SWR_DITHER_RECTANGULAR            = SwrDitherType(C.SWR_DITHER_RECTANGULAR)
	SWR_DITHER_TRIANGULAR             = SwrDitherType(C.SWR_DITHER_TRIANGULAR)
	SWR_DITHER_TRIANGULAR_HIGHPASS    = SwrDitherType(C.SWR_DITHER_TRIANGULAR_HIGHPASS)
	SWR_DITHER_NS                     = SwrDitherType(C.SWR_DITHER_NS)
	SWR_DITHER_NS_LIPSHITZ            = SwrDitherType(C.SWR_DITHER_NS_LIPSHITZ)
	SWR_DITHER_NS_F_WEIGHTED          = SwrDitherType(C.SWR_DITHER_NS_F_WEIGHTED)
	SWR_DITHER_NS_MODIFIED_E_WEIGHTED = SwrDitherType(C.SWR_DITHER_NS_MODIFIED_E_WEIGHTED)
	SWR_DITHER_NS_IMPROVED_E_WEIGHTED = SwrDitherType(C.SWR_DITHER_NS_IMPROVED_E_WEIGHTED)
	SWR_DITHER_NS_SHIBATA             = SwrDitherType(C.SWR_DITHER_NS_SHIBATA)
	SWR_DITHER_NS_LOW_SHIBATA         = SwrDitherType(C.SWR_DITHER_NS_LOW_SHIBATA)
	SWR_DITHER_NS_HIGH_SHIBATA        = SwrDitherType(C.SWR_DITHER_NS_HIGH_SHIBATA)
	SWR_DITHER_NB                     = SwrDitherType(C.SWR_DITHER_NB)
)

// Resampling Engines
type SwrEngine = C.enum_SwrEngine

const (
	SWR_ENGINE_SWR  = SwrEngine(C.SWR_ENGINE_SWR)
	SWR_ENGINE_SOXR = SwrEngine(C.SWR_ENGINE_SOXR)
	SWR_ENGINE_NB   = SwrEngine(C.SWR_ENGINE_NB)
)

// Resampling Filter Types
type SwrFilterType = C.enum_SwrFilterType

const (
	SWR_FILTER_TYPE_CUBIC            = SwrFilterType(C.SWR_FILTER_TYPE_CUBIC)
	SWR_FILTER_TYPE_BLACKMAN_NUTTALL = SwrFilterType(C.SWR_FILTER_TYPE_BLACKMAN_NUTTALL)
	SWR_FILTER_TYPE_KAISER           = SwrFilterType(C.SWR_FILTER_TYPE_KAISER)
)

// SwrContext
type SwrContext C.struct_SwrContext

// SwrGetClass gets the AVClass for SwrContext. It can be used in combination with
// Av_OPT_SEARCH_FAKE_OBJ for examining options.
func SwrGetClass() *AVClass {
	return (*AVClass)(C.swr_get_class())
}

// SwrAlloc allocates SwrContext.
func SwrAlloc() *SwrContext {
	return (*SwrContext)(C.swr_alloc())
}

// SwrInit initializes context after user parameters have been set.
func SwrInit(s *SwrContext) int32 {
	return (int32)(C.swr_init((*C.struct_SwrContext)(s)))
}

// SwrIsInitialized checks whether an swr context has been initialized or not.
func SwrIsInitialized(s *SwrContext) int32 {
	return (int32)(C.swr_is_initialized((*C.struct_SwrContext)(s)))
}

// SwrAllocSetOpts allocates SwrContext if needed and set/reset common parameters.
func SwrAllocSetOpts(s *SwrContext,
	outChLayout int64, outSampleFmt AVSampleFormat, outSampleRate int32,
	inChLayout int64, inSampleFmt AVSampleFormat, inSampleRate int32,
	logOffset int32, logCtx CVoidPointer) *SwrContext {
	return (*SwrContext)(C.swr_alloc_set_opts((*C.struct_SwrContext)(s),
		(C.int64_t)(outChLayout), (C.enum_AVSampleFormat)(outSampleFmt), (C.int)(outSampleRate),
		(C.int64_t)(inChLayout), (C.enum_AVSampleFormat)(inSampleFmt), (C.int)(inSampleRate),
		(C.int)(logOffset), VoidPointer(logCtx)))
}

// SwrFree frees the given SwrContext and set the pointer to NULL.
func SwrFree(s **SwrContext) {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(s)))
}

// SwrClose closes the context so that SwrIsInitialized() returns 0.
func SwrClose(s *SwrContext) {
	C.swr_close((*C.struct_SwrContext)(s))
}

// SwrConvert converts audio.
func SwrConvert(s *SwrContext, out **uint8, outCount int32, in **uint8, inCount int32) int32 {
	return (int32)(C.swr_convert((*C.struct_SwrContext)(s),
		(**C.uint8_t)(unsafe.Pointer(out)), (C.int)(outCount),
		(**C.uint8_t)(unsafe.Pointer(in)), (C.int)(inCount)))
}

// SwrNextPts converts the next timestamp from input to output
// timestamps are in 1/(in_sample_rate * out_sample_rate) units.
func SwrNextPts(s *SwrContext, pts int64) int64 {
	return (int64)(C.swr_next_pts((*C.struct_SwrContext)(s), (C.int64_t)(pts)))
}

// SwrSetCompensation  activates resampling compensation ("soft" compensation).
// This function is internally called when needed in SwrNextPts().
func SwrSetCompensation(s *SwrContext, sampleDelta, compensationDistance int32) int32 {
	return (int32)(C.swr_set_compensation((*C.struct_SwrContext)(s),
		(C.int)(sampleDelta), (C.int)(compensationDistance)))
}

// SwrSetChannelMapping sets a customized input channel mapping.
func SwrSetChannelMapping(s *SwrContext, channelMap *int32) int32 {
	return (int32)(C.swr_set_channel_mapping((*C.struct_SwrContext)(s), (*C.int)(channelMap)))
}

// SwrBuildMatrix generates a channel mixing matrix.
func SwrBuildMatrix(inLayout, outLayout uint64,
	centerMixLevel, surroundMixLevel, lfeMixLevel float64,
	rematrixMaxval, rematrixVolume float64,
	matrix *float64, stride int32, matrixEncoding AVMatrixEncoding, logCtx CVoidPointer) int32 {
	return (int32)(C.swr_build_matrix((C.uint64_t)(inLayout), (C.uint64_t)(outLayout),
		(C.double)(centerMixLevel), (C.double)(surroundMixLevel), (C.double)(lfeMixLevel),
		(C.double)(rematrixMaxval), (C.double)(rematrixVolume),
		(*C.double)(matrix), (C.int)(stride), (C.enum_AVMatrixEncoding)(matrixEncoding), VoidPointer(logCtx)))
}

// SwrSetMatrix sets a customized remix matrix.
func SwrSetMatrix(s *SwrContext, matrix *float64, stride int32) int32 {
	return (int32)(C.swr_set_matrix((*C.struct_SwrContext)(s), (*C.double)(matrix), (C.int)(stride)))
}

// SwrDropOutput drops the specified number of output samples.
func SwrDropOutput(s *SwrContext, count int32) int32 {
	return (int32)(C.swr_drop_output((*C.struct_SwrContext)(s), (C.int)(count)))
}

// SwrInjectSilence injects the specified number of silence samples.
func SwrInjectSilence(s *SwrContext, count int32) int32 {
	return (int32)(C.swr_inject_silence((*C.struct_SwrContext)(s), (C.int)(count)))
}

// SwrGetDelay gets the delay the next input sample will experience relative to the next output sample.
func SwrGetDelay(s *SwrContext, base int64) int64 {
	return (int64)(C.swr_get_delay((*C.struct_SwrContext)(s), (C.int64_t)(base)))
}

// SwrGetOutSamples Find an upper bound on the number of samples that the next swr_convert
// call will output, if called with in_samples of input samples.
func SwrGetOutSamples(s *SwrContext, inSamples int32) int32 {
	return (int32)(C.swr_get_out_samples((*C.struct_SwrContext)(s), (C.int)(inSamples)))
}

// SwResampleVersion returns the LIBSWRESAMPLE_VERSION_INT constant.
func SwResampleVersion() uint32 {
	return (uint32)(C.swresample_version())
}

// SwResampleConfiguration returns the swr build-time configuration.
func SwResampleConfiguration() string {
	return C.GoString(C.swresample_configuration())
}

// SwResampleLicense returns the swr license.
func SwResampleLicense() string {
	return C.GoString(C.swresample_license())
}

// SwrConvertFrame converts the samples in the input AVFrame and write them to the output AVFrame.
func SwrConvertFrame(s *SwrContext, output, input *AVFrame) int32 {
	return (int32)(C.swr_convert_frame((*C.struct_SwrContext)(s),
		(*C.struct_AVFrame)(output), (*C.struct_AVFrame)(input)))
}

// SwrConfigFrame configures or reconfigure the SwrContext using the information provided by the AVFrames.
func SwrConfigFrame(s *SwrContext, out, in *AVFrame) int32 {
	return (int32)(C.swr_config_frame((*C.struct_SwrContext)(s),
		(*C.struct_AVFrame)(out), (*C.struct_AVFrame)(in)))
}
