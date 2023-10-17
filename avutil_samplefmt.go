package ffmpeg

/*
#include <libavutil/samplefmt.h>
*/
import "C"
import "unsafe"

// AvSampleFormat
type AvSampleFormat = C.enum_AVSampleFormat

const (
	AV_SAMPLE_FMT_NONE = AvSampleFormat(C.AV_SAMPLE_FMT_NONE)
	AV_SAMPLE_FMT_U8   = AvSampleFormat(C.AV_SAMPLE_FMT_U8)
	AV_SAMPLE_FMT_S16  = AvSampleFormat(C.AV_SAMPLE_FMT_S16)
	AV_SAMPLE_FMT_S32  = AvSampleFormat(C.AV_SAMPLE_FMT_S32)
	AV_SAMPLE_FMT_FLT  = AvSampleFormat(C.AV_SAMPLE_FMT_FLT)
	AV_SAMPLE_FMT_DBL  = AvSampleFormat(C.AV_SAMPLE_FMT_DBL)
	AV_SAMPLE_FMT_U8P  = AvSampleFormat(C.AV_SAMPLE_FMT_U8P)
	AV_SAMPLE_FMT_S16P = AvSampleFormat(C.AV_SAMPLE_FMT_S16P)
	AV_SAMPLE_FMT_S32P = AvSampleFormat(C.AV_SAMPLE_FMT_S32P)
	AV_SAMPLE_FMT_FLTP = AvSampleFormat(C.AV_SAMPLE_FMT_FLTP)
	AV_SAMPLE_FMT_DBLP = AvSampleFormat(C.AV_SAMPLE_FMT_DBLP)
	AV_SAMPLE_FMT_S64  = AvSampleFormat(C.AV_SAMPLE_FMT_S64)
	AV_SAMPLE_FMT_S64P = AvSampleFormat(C.AV_SAMPLE_FMT_S64P)
	AV_SAMPLE_FMT_NB   = AvSampleFormat(C.AV_SAMPLE_FMT_NB)
)

// AvGetSampleFmtName returns the name of sample_fmt, or NULL if sample_fmt is not
// recognized.
func AvGetSampleFmtName(sampleFmt AvSampleFormat) string {
	return C.GoString(C.av_get_sample_fmt_name((C.enum_AVSampleFormat)(sampleFmt)))
}

// AvGetSampleFmt returns a sample format corresponding to name, or AV_SAMPLE_FMT_NONE
// on error.
func AvGetSampleFmt(name string) AvSampleFormat {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (AvSampleFormat)(C.av_get_sample_fmt((*C.char)(namePtr)))
}

// AvGetAltSampleFmt returns the planar<->packed alternative form of the given sample format, or
// AV_SAMPLE_FMT_NONE on error. If the passed sample_fmt is already in the
// requested planar/packed format, the format returned is the same as the
// input.
func AvGetAltSampleFmt(sampleFmt AvSampleFormat, planar int32) AvSampleFormat {
	return (AvSampleFormat)(C.av_get_alt_sample_fmt((C.enum_AVSampleFormat)(sampleFmt), C.int(planar)))
}

// AvGetPackedSampleFmt gets the packed alternative form of the given sample format.
func AvGetPackedSampleFmt(sampleFmt AvSampleFormat) AvSampleFormat {
	return (AvSampleFormat)(C.av_get_packed_sample_fmt((C.enum_AVSampleFormat)(sampleFmt)))
}

// AvGetPlanarSampleFmt gets the planar alternative form of the given sample format.
func AvGetPlanarSampleFmt(sampleFmt AvSampleFormat) AvSampleFormat {
	return (AvSampleFormat)(C.av_get_planar_sample_fmt((C.enum_AVSampleFormat)(sampleFmt)))
}

// AvGetSampleFmtString generates a string corresponding to the sample format with
// sample_fmt, or a header if sample_fmt is negative.
func AvGetSampleFmtString(buf *int8, bufSize int32, sampleFmt AvSampleFormat) string {
	return C.GoString(C.av_get_sample_fmt_string((*C.char)(buf), (C.int)(bufSize),
		(C.enum_AVSampleFormat)(sampleFmt)))
}

// AvGetBytesPerSample returns number of bytes per sample.
func AvGetBytesPerSample(sampleFmt AvSampleFormat) int32 {
	return (int32)(C.av_get_bytes_per_sample((C.enum_AVSampleFormat)(sampleFmt)))
}

// AvSampleFmtIsPlanar checks if the sample format is planar.
func AvSampleFmtIsPlanar(sampleFmt AvSampleFormat) int32 {
	return (int32)(C.av_sample_fmt_is_planar((C.enum_AVSampleFormat)(sampleFmt)))
}

// AvSamplesGetBufferSize gets the required buffer size for the given audio parameters.
func AvSamplesGetBufferSize(linesize *int32, nbChannels, nbSamples int32,
	sampleFmt AvSampleFormat, align int32) int32 {
	return (int32)(C.av_samples_get_buffer_size((*C.int)(linesize), (C.int)(nbChannels),
		(C.int)(nbSamples), (C.enum_AVSampleFormat)(sampleFmt), (C.int)(align)))
}

// AvSamplesFillArrays fills plane data pointers and linesize for samples with sample
// format sample_fmt.
func AvSamplesFillArrays(audioData **uint8, linesize *int32,
	buf *uint8, nbChannels, nbSamples int32,
	sampleFmt AvSampleFormat, align int32) int32 {
	return (int32)(C.av_samples_fill_arrays((**C.uint8_t)(unsafe.Pointer(audioData)),
		(*C.int)(linesize), (*C.uint8_t)(buf),
		(C.int)(nbChannels), (C.int)(nbSamples),
		(C.enum_AVSampleFormat)(sampleFmt), (C.int)(align)))
}

// AvSamplesAlloc allocates a samples buffer for nb_samples samples, and fill data pointers and
// linesize accordingly.
func AvSamplesAlloc(audioData **uint8, linesize *int32,
	nbChannels, nbSamples int32,
	sampleFmt AvSampleFormat, align int32) int32 {
	return (int32)(C.av_samples_alloc((**C.uint8_t)(unsafe.Pointer(audioData)),
		(*C.int)(linesize),
		(C.int)(nbChannels), (C.int)(nbSamples),
		(C.enum_AVSampleFormat)(sampleFmt), (C.int)(align)))
}

// AvSamplesAllocArrayAndSamples allocates a data pointers array, samples buffer for nb_samples
// samples, and fill data pointers and linesize accordingly.
func AvSamplesAllocArrayAndSamples(audioData ***uint8, linesize *int32, nbChannels, nbSamples int32,
	sampleFmt AvSampleFormat, align int32) int32 {
	return (int32)(C.av_samples_alloc_array_and_samples((***C.uint8_t)(unsafe.Pointer(audioData)),
		(*C.int)(linesize), (C.int)(nbChannels), (C.int)(nbSamples),
		(C.enum_AVSampleFormat)(sampleFmt), (C.int)(align)))
}

// AvSamplesCopy copies samples from src to dst.
func AvSamplesCopy(dst, src **uint8, dstOffset, srcOffset int32,
	nbSamples, nbChannels int32, sampleFmt AvSampleFormat) int32 {
	return (int32)(C.av_samples_copy((**C.uint8_t)(unsafe.Pointer(dst)),
		(**C.uint8_t)(unsafe.Pointer(src)),
		(C.int)(dstOffset), (C.int)(srcOffset),
		(C.int)(nbSamples), (C.int)(nbChannels),
		(C.enum_AVSampleFormat)(sampleFmt)))
}

// AvSamplesSetSilence fills an audio buffer with silence.
func AvSamplesSetSilence(audioData **uint8, offset int32,
	nbSamples, nbChannels int32, sampleFmt AvSampleFormat) int32 {
	return (int32)(C.av_samples_set_silence((**C.uint8_t)(unsafe.Pointer(audioData)), (C.int)(offset),
		(C.int)(nbSamples), (C.int)(nbChannels), (C.enum_AVSampleFormat)(sampleFmt)))
}
