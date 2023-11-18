// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/video_hint.h>
*/
import "C"
import "unsafe"

// AVVideoRect
type AVVideoRect C.struct_AVVideoRect

// GetX gets `AVVideoRect.x` value.
func (rect *AVVideoRect) GetX() uint32 {
	return (uint32)(rect.x)
}

// SetX sets `AVVideoRect.x` value.
func (rect *AVVideoRect) SetX(v uint32) {
	rect.x = (C.uint32_t)(v)
}

// GetXAddr gets `AVVideoRect.x` address.
func (rect *AVVideoRect) GetXAddr() *uint32 {
	return (*uint32)(&rect.x)
}

// GetY gets `AVVideoRect.y` value.
func (rect *AVVideoRect) GetY() uint32 {
	return (uint32)(rect.y)
}

// SetY sets `AVVideoRect.y` value.
func (rect *AVVideoRect) SetY(v uint32) {
	rect.y = (C.uint32_t)(v)
}

// GetYAddr gets `AVVideoRect.y` address.
func (rect *AVVideoRect) GetYAddr() *uint32 {
	return (*uint32)(&rect.y)
}

// GetWidth gets `AVVideoRect.width` value.
func (rect *AVVideoRect) GetWidth() uint32 {
	return (uint32)(rect.width)
}

// SetWidth sets `AVVideoRect.width` value.
func (rect *AVVideoRect) SetWidth(v uint32) {
	rect.width = (C.uint32_t)(v)
}

// GetWidthAddr gets `AVVideoRect.width` address.
func (rect *AVVideoRect) GetWidthAddr() *uint32 {
	return (*uint32)(&rect.width)
}

// GetHeight gets `AVVideoRect.height` value.
func (rect *AVVideoRect) GetHeight() uint32 {
	return (uint32)(rect.height)
}

// SetHeight sets `AVVideoRect.height` value.
func (rect *AVVideoRect) SetHeight(v uint32) {
	rect.height = (C.uint32_t)(v)
}

// GetHeightAddr gets `AVVideoRect.height` address.
func (rect *AVVideoRect) GetHeightAddr() *uint32 {
	return (*uint32)(&rect.height)
}

// AVVideoHintType
type AVVideoHintType = C.AVVideoHintType // C.enum_AVVideoHintType

const (
	AV_VIDEO_HINT_TYPE_CONSTANT = AVVideoHintType(C.AV_VIDEO_HINT_TYPE_CONSTANT)
	AV_VIDEO_HINT_TYPE_CHANGED  = AVVideoHintType(C.AV_VIDEO_HINT_TYPE_CHANGED)
)

// AVVideoHint
type AVVideoHint C.struct_AVVideoHint

// GetNbRects gets `AVVideoHint.nb_rects` value.
func (hint *AVVideoHint) GetNbRects() uintptr {
	return (uintptr)(hint.nb_rects)
}

// SetNbRects sets `AVVideoHint.nb_rects` value.
func (hint *AVVideoHint) SetNbRects(v uintptr) {
	hint.nb_rects = (C.size_t)(v)
}

// GetNbRectsAddr gets `AVVideoHint.nb_rects` address.
func (hint *AVVideoHint) GetNbRectsAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&hint.nb_rects))
}

// GetRectOffset gets `AVVideoHint.rect_offset` value.
func (hint *AVVideoHint) GetRectOffset() uintptr {
	return (uintptr)(hint.rect_offset)
}

// SetRectOffset sets `AVVideoHint.rect_offset` value.
func (hint *AVVideoHint) SetRectOffset(v uintptr) {
	hint.rect_offset = (C.size_t)(v)
}

// GetRectOffsetAddr gets `AVVideoHint.rect_offset` address.
func (hint *AVVideoHint) GetRectOffsetAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&hint.rect_offset))
}

// GetRectSize gets `AVVideoHint.rect_size` value.
func (hint *AVVideoHint) GetRectSize() uintptr {
	return (uintptr)(hint.rect_size)
}

// SetRectSize sets `AVVideoHint.rect_size` value.
func (hint *AVVideoHint) SetRectSize(v uintptr) {
	hint.rect_size = (C.size_t)(v)
}

// GetRectSizeAddr gets `AVVideoHint.rect_size` address.
func (hint *AVVideoHint) GetRectSizeAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&hint.rect_size))
}

// GetType gets `AVVideoHint.type` value.
func (hint *AVVideoHint) GetType() AVVideoHintType {
	return (AVVideoHintType)(hint._type)
}

// SetType sets `AVVideoHint.type` value.
func (hint *AVVideoHint) SetType(v AVVideoHintType) {
	hint._type = (C.AVVideoHintType)(v)
}

// GetTypeAddr gets `AVVideoHint.type` address.
func (hint *AVVideoHint) GetTypeAddr() *AVVideoHintType {
	return (*AVVideoHintType)(&hint._type)
}

// AvVideoHintRects
func AvVideoHintRects(hints *AVVideoHint) *AVVideoRect {
	return (*AVVideoRect)(C.av_video_hint_rects((*C.struct_AVVideoHint)(hints)))
}

// AvVideoHintGetRect
func AvVideoHintGetRect(hints *AVVideoHint, idx uintptr) *AVVideoRect {
	return (*AVVideoRect)(C.av_video_hint_get_rect((*C.struct_AVVideoHint)(hints), (C.size_t)(idx)))
}

// AvVideoHintAlloc allocates memory for the AVVideoHint struct along with an nb_rects-sized
// arrays of AVVideoRect.
func AvVideoHintAlloc(nbRects uintptr, outSize *uintptr) *AVVideoHint {
	return (*AVVideoHint)(C.av_video_hint_alloc((C.size_t)(nbRects), (*C.size_t)(unsafe.Pointer(outSize))))
}

// AvVideoHintCreateSideData same as AvVideoHintAlloc(), except newly-allocated AVVideoHint is attached
// as side data of type AV_FRAME_DATA_VIDEO_HINT_INFO to frame.
func AvVideoHintCreateSideData(frame *AVFrame, nbRects uintptr) *AVVideoHint {
	return (*AVVideoHint)(C.av_video_hint_create_side_data((*C.struct_AVFrame)(frame), (C.size_t)(nbRects)))
}
