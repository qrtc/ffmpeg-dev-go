package ffmpeg

/*
#include <libavfilter/buffersink.h>
*/
import "C"

// AvBuffersinkGetFrameFlags gets a frame with filtered data from sink and put it in frame.
func AvBuffersinkGetFrameFlags(ctx *AvFilterContext, frame *AvFrame, flags int32) int32 {
	return (int32)(C.av_buffersink_get_frame_flags((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame), (C.int)(flags)))
}

const (
	AV_BUFFERSINK_FLAG_PEEK       = C.AV_BUFFERSINK_FLAG_PEEK
	AV_BUFFERSINK_FLAG_NO_REQUEST = C.AV_BUFFERSINK_FLAG_NO_REQUEST
)

type AvBufferSinkParams C.struct_AVBufferSinkParams

// Deprecated: No use
func AvBuffersinkParamsAlloc() *AvBufferSinkParams {
	return (*AvBufferSinkParams)(C.av_buffersink_params_alloc())
}

type AvABufferSinkParams C.struct_AVABufferSinkParams

// Deprecated: No use
func AvAbuffersinkParamsAlloc() *AvABufferSinkParams {
	return (*AvABufferSinkParams)(C.av_abuffersink_params_alloc())
}

// AvBuffersinkSetFrameSize sets the frame size for an audio buffer sink.
func AvBuffersinkSetFrameSize(ctx *AvFilterContext, frameSize uint32) {
	C.av_buffersink_set_frame_size((*C.struct_AVFilterContext)(ctx), (C.uint)(frameSize))
}

// AvBuffersinkGetType
func AvBuffersinkGetType(ctx *AvFilterContext) AvMediaType {
	return (AvMediaType)(C.av_buffersink_get_type((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetTimeBase
func AvBuffersinkGetTimeBase(ctx *AvFilterContext) AvRational {
	return (AvRational)(C.av_buffersink_get_time_base((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetFormat
func AvBuffersinkGetFormat(ctx *AvFilterContext) int32 {
	return (int32)(C.av_buffersink_get_format((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetFrameRate
func AvBuffersinkGetFrameRate(ctx *AvFilterContext) AvRational {
	return (AvRational)(C.av_buffersink_get_frame_rate((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetW
func AvBuffersinkGetW(ctx *AvFilterContext) int32 {
	return (int32)(C.av_buffersink_get_w((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetH
func AvBuffersinkGetH(ctx *AvFilterContext) int32 {
	return (int32)(C.av_buffersink_get_h((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetSampleAspectRatio
func AvBuffersinkGetSampleAspectRatio(ctx *AvFilterContext) AvRational {
	return (AvRational)(C.av_buffersink_get_sample_aspect_ratio((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetChannels
func AvBuffersinkGetChannels(ctx *AvFilterContext) int32 {
	return (int32)(C.av_buffersink_get_channels((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetChannelLayout
func AvBuffersinkGetChannelLayout(ctx *AvFilterContext) uint64 {
	return (uint64)(C.av_buffersink_get_channel_layout((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetSampleRate
func AvBuffersinkGetSampleRate(ctx *AvFilterContext) int32 {
	return (int32)(C.av_buffersink_get_sample_rate((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetHwFramesCtx
func AvBuffersinkGetHwFramesCtx(ctx *AvFilterContext) *AvBufferRef {
	return (*AvBufferRef)(C.av_buffersink_get_hw_frames_ctx((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetFrame gets a frame with filtered data from sink and put it in frame.
func AvBuffersinkGetFrame(ctx *AvFilterContext, frame *AvFrame) int32 {
	return (int32)(C.av_buffersink_get_frame((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame)))
}

// AvBuffersinkGetSamples same as AvBuffersinkGetFrame(), but with the ability to specify the number
// of samples read. This function is less efficient than AvBuffersinkGetFrame(), because it copies the data around.
func AvBuffersinkGetSamples(ctx *AvFilterContext, frame *AvFrame, nbSamples int32) int32 {
	return (int32)(C.av_buffersink_get_samples((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame), (C.int)(nbSamples)))
}
