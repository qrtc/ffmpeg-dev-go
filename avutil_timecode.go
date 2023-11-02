// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/timecode.h>
*/
import "C"

const (
	AV_TIMECODE_STR_SIZE = C.AV_TIMECODE_STR_SIZE
)

type AVTimecodeFlag = C.enum_AVTimecodeFlag

const (
	AV_TIMECODE_FLAG_DROPFRAME     = AVTimecodeFlag(C.AV_TIMECODE_FLAG_DROPFRAME)
	AV_TIMECODE_FLAG_24HOURSMAX    = AVTimecodeFlag(C.AV_TIMECODE_FLAG_24HOURSMAX)
	AV_TIMECODE_FLAG_ALLOWNEGATIVE = AVTimecodeFlag(C.AV_TIMECODE_FLAG_ALLOWNEGATIVE)
)

type AVTimecode C.AVTimecode

// GetStart gets `AVTimecode.start` value.
func (tc *AVTimecode) GetStart() int32 {
	return (int32)(tc.start)
}

// SetStart sets `AVTimecode.start` value.
func (tc *AVTimecode) SetStart(v int32) {
	tc.start = (C.int)(v)
}

// GetStartAddr gets `AVTimecode.start` address.
func (tc *AVTimecode) GetStartAddr() *int32 {
	return (*int32)(&tc.start)
}

// GetFlags gets `AVTimecode.flags` value.
func (tc *AVTimecode) GetFlags() uint32 {
	return (uint32)(tc.flags)
}

// SetFlags sets `AVTimecode.flags` value.
func (tc *AVTimecode) SetFlags(v uint32) {
	tc.flags = (C.uint32_t)(v)
}

// GetFlagsAddr gets `AVTimecode.flags` address.
func (tc *AVTimecode) GetFlagsAddr() *uint32 {
	return (*uint32)(&tc.flags)
}

// GetRate gets `AVTimecode.rate` value.
func (tc *AVTimecode) GetRate() AVRational {
	return (AVRational)(tc.rate)
}

// SetRate sets `AVTimecode.rate` value.
func (tc *AVTimecode) SetRate(v AVRational) {
	tc.rate = (C.struct_AVRational)(v)
}

// GetRateAddr gets `AVTimecode.rate` address.
func (tc *AVTimecode) GetRateAddr() *AVRational {
	return (*AVRational)(&tc.rate)
}

// GetFps gets `AVTimecode.fps` value.
func (tc *AVTimecode) GetFps() uint32 {
	return (uint32)(tc.fps)
}

// SetFps sets `AVTimecode.fps` value.
func (tc *AVTimecode) SetFps(v uint32) {
	tc.fps = (C.uint)(v)
}

// GetFpsAddr gets `AVTimecode.fps` address.
func (tc *AVTimecode) GetFpsAddr() *uint32 {
	return (*uint32)(&tc.fps)
}

// AvTimecodeAdjustNtscFramenum2 adjusts frame number for NTSC drop frame time code.
func AvTimecodeAdjustNtscFramenum2(framenum, fps int32) int32 {
	return (int32)(C.av_timecode_adjust_ntsc_framenum2((C.int)(framenum), (C.int)(fps)))
}

// AvTimecodeGetSmpteFromFramenum converts frame number to SMPTE 12M binary representation.
func AvTimecodeGetSmpteFromFramenum(tc *AVTimecode, framenum int32) uint32 {
	return (uint32)(C.av_timecode_get_smpte_from_framenum((*C.AVTimecode)(tc), (C.int)(framenum)))
}

// AvTimecodeMakeString loads timecode string in buf.
func AvTimecodeMakeString(tc *AVTimecode, framenum int32) (buf, bufPar string) {
	b := make([]C.char, AV_TIMECODE_STR_SIZE+1)
	ret := C.av_timecode_make_string((*C.AVTimecode)(tc), (*C.char)(&b[0]), (C.int)(framenum))
	return C.GoString(&b[0]), C.GoString(ret)
}

// AvTimecodeMakeSmpteTcString gets the timecode string from the SMPTE timecode format.
func AvTimecodeMakeSmpteTcString(tcsmpte uint32, preventDf int32) (buf, bufPar string) {
	b := make([]C.char, AV_TIMECODE_STR_SIZE+1)
	ret := C.av_timecode_make_smpte_tc_string((*C.char)(&b[0]), (C.uint32_t)(tcsmpte), (C.int)(preventDf))
	return C.GoString(&b[0]), C.GoString(ret)
}

// AvTimecodeMakeMpegTcString gets the timecode string from the 25-bit timecode format (MPEG GOP format).
func AvTimecodeMakeMpegTcString(tc25bit uint32) (buf, bufPar string) {
	b := make([]C.char, AV_TIMECODE_STR_SIZE+1)
	ret := C.av_timecode_make_mpeg_tc_string((*C.char)(&b[0]), (C.uint32_t)(tc25bit))
	return C.GoString(&b[0]), C.GoString(ret)
}

// AvTimecodeInit initializes a timecode struct with the passed parameters.
func AvTimecodeInit(tc *AVTimecode, rate AVRational, flags, frameStart int32, logCtx CVoidPointer) int32 {
	return (int32)(C.av_timecode_init((*C.AVTimecode)(tc), (C.struct_AVRational)(rate),
		(C.int)(flags), (C.int)(frameStart), VoidPointer(logCtx)))
}

// AvTimecodeInitFromString parses timecode representation (hh:mm:ss[:;.]ff).
func AvTimecodeInitFromString(tc *AVTimecode, rate AVRational, str string, logCtx CVoidPointer) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.av_timecode_init_from_string((*C.AVTimecode)(tc), (C.struct_AVRational)(rate),
		(*C.char)(strPtr), VoidPointer(logCtx)))
}

// AvTimecodeCheckFrameRate checks if the timecode feature is available for the given frame rate
func AvTimecodeCheckFrameRate(rate AVRational) int32 {
	return (int32)(C.av_timecode_check_frame_rate((C.struct_AVRational)(rate)))
}
