// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/log.h>

typedef const char* (*av_class_item_name_func)(void* ctx);

typedef void* (*av_class_child_next_func)(void *obj, void *prev);

typedef const struct AVClass* (*av_class_child_class_next_func)(const struct AVClass *prev);

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
)

// AVClassCategory
type AVClassCategory = int32

const (
	AV_CLASS_CATEGORY_NA                  = AVClassCategory(C.AV_CLASS_CATEGORY_NA)
	AV_CLASS_CATEGORY_INPUT               = AVClassCategory(C.AV_CLASS_CATEGORY_INPUT)
	AV_CLASS_CATEGORY_OUTPUT              = AVClassCategory(C.AV_CLASS_CATEGORY_OUTPUT)
	AV_CLASS_CATEGORY_MUXER               = AVClassCategory(C.AV_CLASS_CATEGORY_MUXER)
	AV_CLASS_CATEGORY_DEMUXER             = AVClassCategory(C.AV_CLASS_CATEGORY_DEMUXER)
	AV_CLASS_CATEGORY_ENCODER             = AVClassCategory(C.AV_CLASS_CATEGORY_ENCODER)
	AV_CLASS_CATEGORY_DECODER             = AVClassCategory(C.AV_CLASS_CATEGORY_DECODER)
	AV_CLASS_CATEGORY_FILTER              = AVClassCategory(C.AV_CLASS_CATEGORY_FILTER)
	AV_CLASS_CATEGORY_BITSTREAM_FILTER    = AVClassCategory(C.AV_CLASS_CATEGORY_BITSTREAM_FILTER)
	AV_CLASS_CATEGORY_SWSCALER            = AVClassCategory(C.AV_CLASS_CATEGORY_SWSCALER)
	AV_CLASS_CATEGORY_SWRESAMPLER         = AVClassCategory(C.AV_CLASS_CATEGORY_SWRESAMPLER)
	AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT = AVClassCategory(C.AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT)
	AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT  = AVClassCategory(C.AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT)
	AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT = AVClassCategory(C.AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT)
	AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT  = AVClassCategory(C.AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT)
	AV_CLASS_CATEGORY_DEVICE_OUTPUT       = AVClassCategory(C.AV_CLASS_CATEGORY_DEVICE_OUTPUT)
	AV_CLASS_CATEGORY_DEVICE_INPUT        = AVClassCategory(C.AV_CLASS_CATEGORY_DEVICE_INPUT)
	AV_CLASS_CATEGORY_NB                  = AVClassCategory(C.AV_CLASS_CATEGORY_NB)
)

func Av_IS_INPUT_DEVICE(c AVClassCategory) bool {
	return c == AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT ||
		c == AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT ||
		c == AV_CLASS_CATEGORY_DEVICE_INPUT
}

func Av_IS_OUTPUT_DEVICE(c AVClassCategory) bool {
	return c == AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT ||
		c == AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT ||
		c == AV_CLASS_CATEGORY_DEVICE_OUTPUT
}

// AVClass
type AVClass C.struct_AVClass

// GetClassName gets `AVClass.class_name` value.
func (cls *AVClass) GetClassName() string {
	return C.GoString(cls.class_name)
}

// typedef const char* (*av_class_item_name_func)(void* ctx);
type AvClassItemNameFunc = C.av_class_item_name_func

// GetItemName gets `AVClass.item_name` value.
func (cls *AVClass) GetItemName() AvClassItemNameFunc {
	return (AvClassItemNameFunc)(cls.item_name)
}

// SetItemName sets `AVClass.item_name` value.
func (cls *AVClass) SetItemName(v AvClassItemNameFunc) {
	cls.item_name = (C.av_class_item_name_func)(v)
}

// GetItemNameAddr gets `AVClass.item_name` address.
func (cls *AVClass) GetItemNameAddr() *AvClassItemNameFunc {
	return (*AvClassItemNameFunc)(&cls.item_name)
}

// GetOption gets `AVClass.option` value.
func (cls *AVClass) GetOption() *AVOption {
	return (*AVOption)(cls.option)
}

// GetVersion gets `AVClass.version` value.
func (cls *AVClass) GetVersion() int32 {
	return (int32)(cls.version)
}

// GetLogLevelOffsetOffset gets `AVClass.log_level_offset_offset` value.
func (cls *AVClass) GetLogLevelOffsetOffset() int32 {
	return (int32)(cls.log_level_offset_offset)
}

// GetParentLogContextOffset gets `AVClass.parent_log_context_offset` value.
func (cls *AVClass) GetParentLogContextOffset() int32 {
	return (int32)(cls.parent_log_context_offset)
}

// typedef void* (*av_class_child_next_func)(void *obj, void *prev);
type AvClassChildNextFunc = C.av_class_child_next_func

// typedef const struct AVClass* (*av_class_child_class_next_func)(const struct AVClass *prev);
type AvClassChildClassNextFunc = C.av_class_child_class_next_func

// GetChildNext gets `AVClass.child_next` value.
func (cls *AVClass) GetChildNext() AvClassChildNextFunc {
	return (AvClassChildNextFunc)(cls.child_next)
}

// SetChildNext sets `AVClass.child_next` value.
func (cls *AVClass) SetChildNext(v AvClassChildNextFunc) {
	cls.child_next = (C.av_class_child_next_func)(v)
}

// GetChildNextAddr gets `AVClass.child_next` address.
func (cls *AVClass) GetChildNextAddr() *AvClassChildNextFunc {
	return (*AvClassChildNextFunc)(&cls.child_next)
}

// GetChildClassNext gets `AVClass.child_class_next` value.
func (cls *AVClass) GetChildClassNext() AvClassChildClassNextFunc {
	return (AvClassChildClassNextFunc)(cls.child_class_next)
}

// SetChildClassNext sets `AVClass.child_class_next` value.
func (cls *AVClass) SetChildClassNext(v AvClassChildClassNextFunc) {
	cls.child_class_next = (C.av_class_child_class_next_func)(v)
}

// GetChildClassNextAddr gets `AVClass.child_class_next` address.
func (cls *AVClass) GetChildClassNextAddr() *AvClassChildClassNextFunc {
	return (*AvClassChildClassNextFunc)(&cls.child_class_next)
}

// GetCategory gets `AVClass.category` value.
func (cls *AVClass) GetCategory() AVClassCategory {
	return (AVClassCategory)(cls.category)
}

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

// Av_LOG_C sets additional colors for extended debugging sessions.
func Av_LOG_C(x int32) int32 {
	return x << 8
}

// AvLog sends the specified message to the log if the level is less than or equal
// to the current av_log_level. By default, all logging messages are sent to stderr.
// This behavior can be altered by setting a different logging callback function.
func AvLog(avcl CVoidPointer, level int32, _fmt string, va ...any) {
	strPtr, strFunc := StringCasting(fmt.Sprintf(_fmt, va...))
	defer strFunc()
	C.av_log_wrap(VoidPointer(avcl), (C.int)(level), (*C.char)(strPtr))
}

// AvLogOnce sends the specified message to the log once with the initial_level and then with
// the subsequent_level. By default, all logging messages are sent to stderr.
// This behavior can be altered by setting a different logging callback function.
func AvLogOnce(avcl CVoidPointer, initialLevel, subsequentLevel int32, state *int32, _fmt string, va ...any) {
	fmtPtr, fmtFunc := StringCasting(fmt.Sprintf(_fmt, va...))
	defer fmtFunc()
	C.av_log_once_wrap(VoidPointer(avcl), (C.int)(initialLevel), (C.int)(subsequentLevel), (*C.int)(state), (*C.char)(fmtPtr))
}

// NONEED: av_vlog

// AvLogGetLevel gets the current log level.
func AvLogGetLevel() int32 {
	return (int32)(C.av_log_get_level())
}

// AvLogSetLevel sets the log level.
func AvLogSetLevel(level int32) {
	C.av_log_set_level((C.int)(level))
}

// typedef void (*av_log_callback_func)(void*, int, const char*, va_list);
type AVLogCallbackFunc = C.av_log_callback_func

// AvLogSetCallback sets the logging callback.
func AvLogSetCallback(f AVLogCallbackFunc) {
	C.av_log_set_callback(f)
}

// NONEED: av_log_default_callback

// AvDefaultItemName returns the context name.
func AvDefaultItemName(ctx CVoidPointer) string {
	return C.GoString(C.av_default_item_name(VoidPointer(ctx)))
}

// AvDefaultGetCategory
func AvDefaultGetCategory(ptr CVoidPointer) AVClassCategory {
	return (AVClassCategory)(C.av_default_get_category(VoidPointer(ptr)))
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
