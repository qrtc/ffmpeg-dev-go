// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/parseutils.h>
*/
import "C"
import "unsafe"

// AvParseRatio parses str and store the parsed ratio in q.
func AvParseRatio(q *AVRational, str string, max, logOffset int32, logCtx CVoidPointer) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.av_parse_ratio((*C.struct_AVRational)(q),
		(*C.char)(strPtr), (C.int)(max), (C.int)(logOffset), VoidPointer(logCtx)))
}

// AvParseRatioQuiet
func AvParseRatioQuiet(q *AVRational, str string, max int32) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.av_parse_ratio((*C.struct_AVRational)(q),
		(*C.char)(strPtr), (C.int)(max), (C.int)(AV_LOG_MAX_OFFSET), nil))
}

// AvParseVideoSize parses str and put in width_ptr and height_ptr the detected values.
func AvParseVideoSize(widthPtr, heightPtr *int32, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.av_parse_video_size((*C.int)(widthPtr), (*C.int)(heightPtr), (*C.char)(strPtr)))
}

// AvParseVideoRate parses str and store the detected values in *rate.
func AvParseVideoRate(rate *AVRational, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.av_parse_video_rate((*C.struct_AVRational)(rate), (*C.char)(strPtr)))
}

// AvParseColor puts the RGBA values that correspond to color_string in rgba_color.
func AvParseColor(rgbaColor *uint8, colorString string, slen int32, logCtx CVoidPointer) int32 {
	strPtr, strFunc := StringCasting(colorString)
	defer strFunc()
	return (int32)(C.av_parse_color((*C.uint8_t)(rgbaColor), (*C.char)(strPtr), (C.int)(slen), VoidPointer(logCtx)))
}

// AvGetKnownColorName gets the name of a color from the internal table of hard-coded named colors.
func AvGetKnownColorName(colorIdx int32, rgb **uint8) string {
	return C.GoString(C.av_get_known_color_name((C.int)(colorIdx), (**C.uint8_t)(unsafe.Pointer(rgb))))
}

// AvParseTime parse timestr and return in *time a corresponding number of microseconds.
func AvParseTime(timeval *int64, timestr string, duration int32) int32 {
	strPtr, strFunc := StringCasting(timestr)
	defer strFunc()
	return (int32)(C.av_parse_time((*C.int64_t)(timeval), (*C.char)(strPtr), (C.int)(duration)))
}

// AvFindInfoTag attempts to find a specific tag in a URL.
func AvFindInfoTag(tag1, info string) (val string, ret int32) {
	tag1Ptr, tag1Func := StringCasting(tag1)
	defer tag1Func()
	infoPtr, infoFunc := StringCasting(info)
	defer infoFunc()
	infoBuf := make([]C.char, len(info))
	ret = (int32)(C.av_find_info_tag(&infoBuf[0], (C.int)(len(info)),
		(*C.char)(tag1Ptr), (*C.char)(infoPtr)))
	return C.GoString(&infoBuf[0]), ret
}

// NONEED: av_small_strptime

// NONEED: av_timegm
