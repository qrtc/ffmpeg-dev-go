package ffmpeg

/*
#include <libavfilter/buffersrc.h>
*/
import "C"

const (
	AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT = int32(C.AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT)
	AV_BUFFERSRC_FLAG_PUSH            = int32(C.AV_BUFFERSRC_FLAG_PUSH)
	AV_BUFFERSRC_FLAG_KEEP_REF        = int32(C.AV_BUFFERSRC_FLAG_KEEP_REF)
)

// AvBuffersrcGetNbFailedRequests gets the number of failed requests.
func AvBuffersrcGetNbFailedRequests(bufferSrc *AvFilterContext) uint32 {
	return (uint32)(C.av_buffersrc_get_nb_failed_requests((*C.struct_AVFilterContext)(bufferSrc)))
}

type AvBufferSrcParameters C.struct_AVBufferSrcParameters

// AvBuffersrcParametersAlloc allocates a new AVBufferSrcParameters instance. It should be freed by the
// caller with AvFree().
func AvBuffersrcParametersAlloc() *AvBufferSrcParameters {
	return (*AvBufferSrcParameters)(C.av_buffersrc_parameters_alloc())
}

// AvBuffersrcParametersSet initializes the buffersrc or abuffersrc filter with the provided parameters.
func AvBuffersrcParametersSet(ctx *AvFilterContext, param *AvBufferSrcParameters) int32 {
	return (int32)(C.av_buffersrc_parameters_set((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVBufferSrcParameters)(param)))
}

// AvBuffersrcWriteFrame adds a frame to the buffer source.
func AvBuffersrcWriteFrame(ctx *AvFilterContext, frame *AvFrame) int32 {
	return (int32)(C.av_buffersrc_write_frame((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame)))
}

// AvBuffersrcAddFrame adds a frame to the buffer source.
func AvBuffersrcAddFrame(ctx *AvFilterContext, frame *AvFrame) int32 {
	return (int32)(C.av_buffersrc_add_frame((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame)))
}

// AvBuffersrcAddFrameFlags adds a frame to the buffer source.
func AvBuffersrcAddFrameFlags(ctx *AvFilterContext, frame *AvFrame, flags int32) int32 {
	return (int32)(C.av_buffersrc_add_frame_flags((*C.struct_AVFilterContext)(ctx),
		(*C.struct_AVFrame)(frame), (C.int)(flags)))
}

// AvBuffersrcClose closes the buffer source after EOF.
func AvBuffersrcClose(ctx *AvFilterContext, pts int64, flags uint32) int32 {
	return (int32)(C.av_buffersrc_close((*C.struct_AVFilterContext)(ctx),
		(C.int64_t)(pts), (C.uint)(flags)))
}
