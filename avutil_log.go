package ffmpeg

/*
#include <libavutil/log.h>
*/
import "C"

// AvClassCategory
type AvClassCategory int32

const (
	AV_CLASS_CATEGORY_NA                  = AvClassCategory(C.AV_CLASS_CATEGORY_NA)
	AV_CLASS_CATEGORY_INPUT               = AvClassCategory(C.AV_CLASS_CATEGORY_INPUT)
	AV_CLASS_CATEGORY_OUTPUT              = AvClassCategory(C.AV_CLASS_CATEGORY_OUTPUT)
	AV_CLASS_CATEGORY_MUXER               = AvClassCategory(C.AV_CLASS_CATEGORY_MUXER)
	AV_CLASS_CATEGORY_DEMUXER             = AvClassCategory(C.AV_CLASS_CATEGORY_DEMUXER)
	AV_CLASS_CATEGORY_ENCODER             = AvClassCategory(C.AV_CLASS_CATEGORY_ENCODER)
	AV_CLASS_CATEGORY_DECODER             = AvClassCategory(C.AV_CLASS_CATEGORY_DECODER)
	AV_CLASS_CATEGORY_FILTER              = AvClassCategory(C.AV_CLASS_CATEGORY_FILTER)
	AV_CLASS_CATEGORY_BITSTREAM_FILTER    = AvClassCategory(C.AV_CLASS_CATEGORY_BITSTREAM_FILTER)
	AV_CLASS_CATEGORY_SWSCALER            = AvClassCategory(C.AV_CLASS_CATEGORY_SWSCALER)
	AV_CLASS_CATEGORY_SWRESAMPLER         = AvClassCategory(C.AV_CLASS_CATEGORY_SWRESAMPLER)
	AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT = AvClassCategory(C.AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT)
	AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT  = AvClassCategory(C.AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT)
	AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT = AvClassCategory(C.AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT)
	AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT  = AvClassCategory(C.AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT)
	AV_CLASS_CATEGORY_DEVICE_OUTPUT       = AvClassCategory(C.AV_CLASS_CATEGORY_DEVICE_OUTPUT)
	AV_CLASS_CATEGORY_DEVICE_INPUT        = AvClassCategory(C.AV_CLASS_CATEGORY_DEVICE_INPUT)
	AV_CLASS_CATEGORY_NB                  = AvClassCategory(C.AV_CLASS_CATEGORY_NB)
)

func AV_IS_INPUT_DEVICE(c AvClassCategory) bool {
	return c == AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT ||
		c == AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT ||
		c == AV_CLASS_CATEGORY_DEVICE_INPUT
}

func AV_IS_OUTPUT_DEVICE(c AvClassCategory) bool {
	return c == AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT ||
		c == AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT ||
		c == AV_CLASS_CATEGORY_DEVICE_OUTPUT
}

// AvClass
type AvClass C.struct_AVClass

// AvLogLevelType
type AvLogLevelType int32

const (
	AV_LOG_QUIET   = AvLogLevelType(C.AV_LOG_QUIET)
	AV_LOG_PANIC   = AvLogLevelType(C.AV_LOG_PANIC)
	AV_LOG_FATAL   = AvLogLevelType(C.AV_LOG_FATAL)
	AV_LOG_ERROR   = AvLogLevelType(C.AV_LOG_ERROR)
	AV_LOG_WARNING = AvLogLevelType(C.AV_LOG_WARNING)
	AV_LOG_INFO    = AvLogLevelType(C.AV_LOG_INFO)
	AV_LOG_VERBOSE = AvLogLevelType(C.AV_LOG_VERBOSE)
	AV_LOG_DEBUG   = AvLogLevelType(C.AV_LOG_DEBUG)
	AV_LOG_TRACE   = AvLogLevelType(C.AV_LOG_TRACE)
)

// AvLogSetLevel sets the log level
func AvLogSetLevel(level AvLogLevelType) {
	C.av_log_set_level(C.int(level))
}

// AvLogGetLevel gets the current log level
func AvLogGetLevel() AvLogLevelType {
	return AvLogLevelType(C.av_log_get_level())
}

const (
	AV_LOG_SKIP_REPEATED = C.AV_LOG_SKIP_REPEATED
	AV_LOG_PRINT_LEVEL   = C.AV_LOG_PRINT_LEVEL
)

// AvLogSetFlags
func AvLogSetFlags(arg int32) {
	C.av_log_set_flags((C.int)(arg))
}

// AvLogGetFlags
func AvLogGetFlags() int32 {
	return (int32)(C.av_log_get_flags())
}
