// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavfilter/buffersink.h>
*/
import "C"

// AvBuffersinkGetFrameFlags gets a frame with filtered data from sink and put it in frame.
func AvBuffersinkGetFrameFlags(ctx *AVFilterContext, frame *AVFrame, flags int32) int32 {
	return (int32)(C.av_buffersink_get_frame_flags((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame), (C.int)(flags)))
}

const (
	AV_BUFFERSINK_FLAG_PEEK       = C.AV_BUFFERSINK_FLAG_PEEK
	AV_BUFFERSINK_FLAG_NO_REQUEST = C.AV_BUFFERSINK_FLAG_NO_REQUEST
)

// AVBufferSinkParams
type AVBufferSinkParams C.struct_AVBufferSinkParams

// AVABufferSinkParams
type AVABufferSinkParams C.struct_AVABufferSinkParams

// AvBuffersinkSetFrameSize sets the frame size for an audio buffer sink.
func AvBuffersinkSetFrameSize(ctx *AVFilterContext, frameSize uint32) {
	C.av_buffersink_set_frame_size((*C.struct_AVFilterContext)(ctx), (C.uint)(frameSize))
}

// AvBuffersinkGetType
func AvBuffersinkGetType(ctx *AVFilterContext) AVMediaType {
	return (AVMediaType)(C.av_buffersink_get_type((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetTimeBase
func AvBuffersinkGetTimeBase(ctx *AVFilterContext) AVRational {
	return (AVRational)(C.av_buffersink_get_time_base((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetFormat
func AvBuffersinkGetFormat(ctx *AVFilterContext) int32 {
	return (int32)(C.av_buffersink_get_format((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetFrameRate
func AvBuffersinkGetFrameRate(ctx *AVFilterContext) AVRational {
	return (AVRational)(C.av_buffersink_get_frame_rate((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetW
func AvBuffersinkGetW(ctx *AVFilterContext) int32 {
	return (int32)(C.av_buffersink_get_w((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetH
func AvBuffersinkGetH(ctx *AVFilterContext) int32 {
	return (int32)(C.av_buffersink_get_h((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetSampleAspectRatio
func AvBuffersinkGetSampleAspectRatio(ctx *AVFilterContext) AVRational {
	return (AVRational)(C.av_buffersink_get_sample_aspect_ratio((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetColorspace
func AvBuffersinkGetColorspace(ctx *AVFilterContext) AVColorSpace {
	return (AVColorSpace)(C.av_buffersink_get_colorspace((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetColorRange
func AvBuffersinkGetColorRange(ctx *AVFilterContext) AVColorRange {
	return (AVColorRange)(C.av_buffersink_get_color_range((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetChannels
func AvBuffersinkGetChannels(ctx *AVFilterContext) int32 {
	return (int32)(C.av_buffersink_get_channels((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetChLayout
func AvBuffersinkGetChLayout(ctx *AVFilterContext, chLayout *AVChannelLayout) int32 {
	return (int32)(C.av_buffersink_get_ch_layout((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVChannelLayout)(chLayout)))
}

// AvBuffersinkGetSampleRate
func AvBuffersinkGetSampleRate(ctx *AVFilterContext) int32 {
	return (int32)(C.av_buffersink_get_sample_rate((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetHwFramesCtx
func AvBuffersinkGetHwFramesCtx(ctx *AVFilterContext) *AVBufferRef {
	return (*AVBufferRef)(C.av_buffersink_get_hw_frames_ctx((*C.struct_AVFilterContext)(ctx)))
}

// AvBuffersinkGetFrame gets a frame with filtered data from sink and put it in frame.
func AvBuffersinkGetFrame(ctx *AVFilterContext, frame *AVFrame) int32 {
	return (int32)(C.av_buffersink_get_frame((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame)))
}

// AvBuffersinkGetSamples same as AVBuffersinkGetFrame(), but with the ability to specify the number
// of samples read. This function is less efficient than AVBuffersinkGetFrame(), because it copies the data around.
func AvBuffersinkGetSamples(ctx *AVFilterContext, frame *AVFrame, nbSamples int32) int32 {
	return (int32)(C.av_buffersink_get_samples((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame), (C.int)(nbSamples)))
}
