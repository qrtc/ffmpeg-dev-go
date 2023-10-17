package ffmpeg

/*
#include <libavutil/log.h>

void av_log_wrap(void *avcl, int level, char *fmt) {
	av_log(avcl, level, fmt, NULL);
}

void av_log_once_wrap(void* avcl, int initial_level, int subsequent_level, int *state, char *fmt) {
	av_log_once(avcl, initial_level, subsequent_level, state, fmt, NULL);
}

typedef void (*av_log_callback_func)(void*, int, const char*, va_list);

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// AvClassCategory
type AvClassCategory = int32

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

const (
	AV_LOG_QUIET   = int32(C.AV_LOG_QUIET)
	AV_LOG_PANIC   = int32(C.AV_LOG_PANIC)
	AV_LOG_FATAL   = int32(C.AV_LOG_FATAL)
	AV_LOG_ERROR   = int32(C.AV_LOG_ERROR)
	AV_LOG_WARNING = int32(C.AV_LOG_WARNING)
	AV_LOG_INFO    = int32(C.AV_LOG_INFO)
	AV_LOG_VERBOSE = int32(C.AV_LOG_VERBOSE)
	AV_LOG_DEBUG   = int32(C.AV_LOG_DEBUG)
	AV_LOG_TRACE   = int32(C.AV_LOG_TRACE)
)

const (
	AV_LOG_MAX_OFFSET = int32(C.AV_LOG_MAX_OFFSET)
)

// AV_LOG_C sets additional colors for extended debugging sessions.
func AV_LOG_C(x int32) int32 {
	return x << 8
}

// AvLog sends the specified message to the log if the level is less than or equal
// to the current av_log_level. By default, all logging messages are sent to stderr.
// This behavior can be altered by setting a different logging callback function.
func AvLog(avcl unsafe.Pointer, level int32, _fmt string, va ...any) {
	strPtr, strFunc := StringCasting(fmt.Sprintf(_fmt, va...))
	defer strFunc()
	C.av_log_wrap(avcl, (C.int)(level), (*C.char)(strPtr))
}

// AvLogOnce sends the specified message to the log once with the initial_level and then with
// the subsequent_level. By default, all logging messages are sent to stderr.
// This behavior can be altered by setting a different logging callback function.
func AvLogOnce(avcl unsafe.Pointer, initialLevel, subsequentLevel int32, state *int32, _fmt string, va ...any) {
	fmtPtr, fmtFunc := StringCasting(fmt.Sprintf(_fmt, va...))
	defer fmtFunc()
	C.av_log_once_wrap(avcl, (C.int)(initialLevel), (C.int)(subsequentLevel), (*C.int)(state), (*C.char)(fmtPtr))
}

// NONEED: av_vlog

// AvLogGetLevel gets the current log level
func AvLogGetLevel() int32 {
	return (int32)(C.av_log_get_level())
}

// AvLogSetLevel sets the log level
func AvLogSetLevel(level int32) {
	C.av_log_set_level(C.int(level))
}

// typedef void (*av_log_callback_func)(void*, int, const char*, va_list)
type AvLogCallbackFunc C.av_log_callback_func

// AvLogSetCallback sets the logging callback
func AvLogSetCallback(f AvLogCallbackFunc) {
	C.av_log_set_callback(f)
}

// NONEED: av_log_default_callback

// AvDefaultItemName returns the context name
func AvDefaultItemName(ctx unsafe.Pointer) string {
	return C.GoString(C.av_default_item_name(ctx))
}

// AvDefaultGetCategory
func AvDefaultGetCategory(ptr unsafe.Pointer) AvClassCategory {
	return (AvClassCategory)(C.av_default_get_category(ptr))
}

// NONEED: av_log_format_line

// NONEED: av_log_format_line2

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
