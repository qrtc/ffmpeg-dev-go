// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavfilter/buffersrc.h>
*/
import "C"
import "unsafe"

const (
	AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT = int32(C.AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT)
	AV_BUFFERSRC_FLAG_PUSH            = int32(C.AV_BUFFERSRC_FLAG_PUSH)
	AV_BUFFERSRC_FLAG_KEEP_REF        = int32(C.AV_BUFFERSRC_FLAG_KEEP_REF)
)

// AvBuffersrcGetNbFailedRequests gets the number of failed requests.
func AvBuffersrcGetNbFailedRequests(bufferSrc *AVFilterContext) uint32 {
	return (uint32)(C.av_buffersrc_get_nb_failed_requests((*C.struct_AVFilterContext)(bufferSrc)))
}

// AVBufferSrcParameters
type AVBufferSrcParameters C.struct_AVBufferSrcParameters

// GetFormat gets `AVBufferSrcParameters.format` value.
func (bsrc *AVBufferSrcParameters) GetFormat() int32 {
	return (int32)(bsrc.format)
}

// SetFormat sets `AVBufferSrcParameters.format` value.
func (bsrc *AVBufferSrcParameters) SetFormat(v int32) {
	bsrc.format = (C.int)(v)
}

// GetFormatAddr gets `AVBufferSrcParameters.format` address.
func (bsrc *AVBufferSrcParameters) GetFormatAddr() *int32 {
	return (*int32)(&bsrc.format)
}

// GetTimeBase gets `AVBufferSrcParameters.time_base` value.
func (bsrc *AVBufferSrcParameters) GetTimeBase() AVRational {
	return (AVRational)(bsrc.time_base)
}

// SetTimeBase sets `AVBufferSrcParameters.time_base` value.
func (bsrc *AVBufferSrcParameters) SetTimeBase(v AVRational) {
	bsrc.time_base = (C.struct_AVRational)(v)
}

// GetTimeBaseAddr gets `AVBufferSrcParameters.time_base` address.
func (bsrc *AVBufferSrcParameters) GetTimeBaseAddr() *AVRational {
	return (*AVRational)(&bsrc.time_base)
}

// GetWidth gets `AVBufferSrcParameters.width` value.
func (bsrc *AVBufferSrcParameters) GetWidth() int32 {
	return (int32)(bsrc.width)
}

// SetWidth sets `AVBufferSrcParameters.width` value.
func (bsrc *AVBufferSrcParameters) SetWidth(v int32) {
	bsrc.width = (C.int)(v)
}

// GetWidthAddr gets `AVBufferSrcParameters.width` address.
func (bsrc *AVBufferSrcParameters) GetWidthAddr() *int32 {
	return (*int32)(&bsrc.width)
}

// GetHeight gets `AVBufferSrcParameters.height` value.
func (bsrc *AVBufferSrcParameters) GetHeight() int32 {
	return (int32)(bsrc.height)
}

// SetHeight sets `AVBufferSrcParameters.height` value.
func (bsrc *AVBufferSrcParameters) SetHeight(v int32) {
	bsrc.height = (C.int)(v)
}

// GetHeightAddr gets `AVBufferSrcParameters.height` address.
func (bsrc *AVBufferSrcParameters) GetHeightAddr() *int32 {
	return (*int32)(&bsrc.height)
}

// GetSampleAspectRatio gets `AVBufferSrcParameters.sample_aspect_ratio` value.
func (bsrc *AVBufferSrcParameters) GetSampleAspectRatio() AVRational {
	return (AVRational)(bsrc.sample_aspect_ratio)
}

// SetSampleAspectRatio sets `AVBufferSrcParameters.sample_aspect_ratio` value.
func (bsrc *AVBufferSrcParameters) SetSampleAspectRatio(v AVRational) {
	bsrc.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// GetSampleAspectRatioAddr gets `AVBufferSrcParameters.sample_aspect_ratio` address.
func (bsrc *AVBufferSrcParameters) GetSampleAspectRatioAddr() *AVRational {
	return (*AVRational)(&bsrc.sample_aspect_ratio)
}

// GetFrameRate gets `AVBufferSrcParameters.frame_rate` value.
func (bsrc *AVBufferSrcParameters) GetFrameRate() AVRational {
	return (AVRational)(bsrc.frame_rate)
}

// SetFrameRate sets `AVBufferSrcParameters.frame_rate` value.
func (bsrc *AVBufferSrcParameters) SetFrameRate(v AVRational) {
	bsrc.frame_rate = (C.struct_AVRational)(v)
}

// GetFrameRateAddr gets `AVBufferSrcParameters.frame_rate` address.
func (bsrc *AVBufferSrcParameters) GetFrameRateAddr() *AVRational {
	return (*AVRational)(&bsrc.frame_rate)
}

// GetHwFramesCtx gets `AVBufferSrcParameters.hw_frames_ctx` value.
func (bsrc *AVBufferSrcParameters) GetHwFramesCtx() *AVBufferRef {
	return (*AVBufferRef)(bsrc.hw_frames_ctx)
}

// SetHwFramesCtx sets `AVBufferSrcParameters.hw_frames_ctx` value.
func (bsrc *AVBufferSrcParameters) SetHwFramesCtx(v *AVBufferRef) {
	bsrc.hw_frames_ctx = (*C.struct_AVBufferRef)(v)
}

// GetHwFramesCtxAddr gets `AVBufferSrcParameters.hw_frames_ctx` address.
func (bsrc *AVBufferSrcParameters) GetHwFramesCtxAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&bsrc.hw_frames_ctx))
}

// GetSampleRate gets `AVBufferSrcParameters.sample_rate` value.
func (bsrc *AVBufferSrcParameters) GetSampleRate() int32 {
	return (int32)(bsrc.sample_rate)
}

// SetSampleRate sets `AVBufferSrcParameters.sample_rate` value.
func (bsrc *AVBufferSrcParameters) SetSampleRate(v int32) {
	bsrc.sample_rate = (C.int)(v)
}

// GetSampleRateAddr gets `AVBufferSrcParameters.sample_rate` address.
func (bsrc *AVBufferSrcParameters) GetSampleRateAddr() *int32 {
	return (*int32)(&bsrc.sample_rate)
}

// Deprecated: use ChLayout instead.
//
// GetChannelLayout gets `AVBufferSrcParameters.channel_layout` value.
func (bsrc *AVBufferSrcParameters) GetChannelLayout() uint64 {
	return (uint64)(bsrc.channel_layout)
}

// Deprecated: use ChLayout instead.
//
// SetChannelLayout sets `AVBufferSrcParameters.channel_layout` value.
func (bsrc *AVBufferSrcParameters) SetChannelLayout(v uint64) {
	bsrc.channel_layout = (C.uint64_t)(v)
}

// Deprecated: use ChLayout instead.
//
// GetChannelLayoutAddr gets `AVBufferSrcParameters.channel_layout` address.
func (bsrc *AVBufferSrcParameters) GetChannelLayoutAddr() *uint64 {
	return (*uint64)(&bsrc.channel_layout)
}

// GetChLayout gets `AVBufferSrcParameters.ch_layout` value.
func (bsrc *AVBufferSrcParameters) GetChLayout() AVChannelLayout {
	return (AVChannelLayout)(bsrc.ch_layout)
}

// SetChLayout sets `AVBufferSrcParameters.ch_layout` value.
func (bsrc *AVBufferSrcParameters) SetChLayout(v AVChannelLayout) {
	bsrc.ch_layout = (C.struct_AVChannelLayout)(v)
}

// GetChLayoutAddr gets `AVBufferSrcParameters.ch_layout` address.
func (bsrc *AVBufferSrcParameters) GetChLayoutAddr() *AVChannelLayout {
	return (*AVChannelLayout)(&bsrc.ch_layout)
}

// AvBuffersrcParametersAlloc allocates a new AVBufferSrcParameters instance. It should be freed by the
// caller with AvFree().
func AvBuffersrcParametersAlloc() *AVBufferSrcParameters {
	return (*AVBufferSrcParameters)(C.av_buffersrc_parameters_alloc())
}

// AvBuffersrcParametersSet initializes the buffersrc or abuffersrc filter with the provided parameters.
func AvBuffersrcParametersSet(ctx *AVFilterContext, param *AVBufferSrcParameters) int32 {
	return (int32)(C.av_buffersrc_parameters_set((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVBufferSrcParameters)(param)))
}

// AvBuffersrcWriteFrame adds a frame to the buffer source.
func AvBuffersrcWriteFrame(ctx *AVFilterContext, frame *AVFrame) int32 {
	return (int32)(C.av_buffersrc_write_frame((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame)))
}

// AvBuffersrcAddFrame adds a frame to the buffer source.
func AvBuffersrcAddFrame(ctx *AVFilterContext, frame *AVFrame) int32 {
	return (int32)(C.av_buffersrc_add_frame((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame)))
}

// AvBuffersrcAddFrameFlags adds a frame to the buffer source.
func AvBuffersrcAddFrameFlags(ctx *AVFilterContext, frame *AVFrame, flags int32) int32 {
	return (int32)(C.av_buffersrc_add_frame_flags((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame), (C.int)(flags)))
}

// AvBuffersrcClose closes the buffer source after EOF.
func AvBuffersrcClose(ctx *AVFilterContext, pts int64, flags uint32) int32 {
	return (int32)(C.av_buffersrc_close((*C.struct_AVFilterContext)(ctx),
		(C.int64_t)(pts), (C.uint)(flags)))
}
