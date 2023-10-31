// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/motion_vector.h>
*/
import "C"

// AVMotionVector
type AVMotionVector C.struct_AVMotionVector

// GetSource gets `AVMotionVector.source` value.
func (mv *AVMotionVector) GetSource() int32 {
	return (int32)(mv.source)
}

// SetSource sets `AVMotionVector.source` value.
func (mv *AVMotionVector) SetSource(v int32) {
	mv.source = (C.int32_t)(v)
}

// GetSourceAddr gets `AVMotionVector.source` address.
func (mv *AVMotionVector) GetSourceAddr() *int32 {
	return (*int32)(&mv.source)
}

// GetW gets `AVMotionVector.w` value.
func (mv *AVMotionVector) GetW() uint8 {
	return (uint8)(mv.w)
}

// SetW sets `AVMotionVector.w` value.
func (mv *AVMotionVector) SetW(v uint8) {
	mv.w = (C.uint8_t)(v)
}

// GetWAddr gets `AVMotionVector.w` address.
func (mv *AVMotionVector) GetWAddr() *uint8 {
	return (*uint8)(&mv.w)
}

// GetH gets `AVMotionVector.h` value.
func (mv *AVMotionVector) GetH() uint8 {
	return (uint8)(mv.h)
}

// SetH sets `AVMotionVector.h` value.
func (mv *AVMotionVector) SetH(v uint8) {
	mv.h = (C.uint8_t)(v)
}

// GetHAddr gets `AVMotionVector.h` address.
func (mv *AVMotionVector) GetHAddr() *uint8 {
	return (*uint8)(&mv.h)
}

// GetSrcX gets `AVMotionVector.src_x` value.
func (mv *AVMotionVector) GetSrcX() int16 {
	return (int16)(mv.src_x)
}

// SetSrcX sets `AVMotionVector.src_x` value.
func (mv *AVMotionVector) SetSrcX(v int16) {
	mv.src_x = (C.int16_t)(v)
}

// GetSrcXAddr gets `AVMotionVector.src_x` address.
func (mv *AVMotionVector) GetSrcXAddr() *int16 {
	return (*int16)(&mv.src_x)
}

// GetSrcY gets `AVMotionVector.src_y` value.
func (mv *AVMotionVector) GetSrcY() int16 {
	return (int16)(mv.src_y)
}

// SetSrcY sets `AVMotionVector.src_y` value.
func (mv *AVMotionVector) SetSrcY(v int16) {
	mv.src_y = (C.int16_t)(v)
}

// GetSrcYAddr gets `AVMotionVector.src_y` address.
func (mv *AVMotionVector) GetSrcYAddr() *int16 {
	return (*int16)(&mv.src_y)
}

// GetDstX gets `AVMotionVector.dst_x` value.
func (mv *AVMotionVector) GetDstX() int16 {
	return (int16)(mv.dst_x)
}

// SetDstX sets `AVMotionVector.dst_x` value.
func (mv *AVMotionVector) SetDstX(v int16) {
	mv.dst_x = (C.int16_t)(v)
}

// GetDstXAddr gets `AVMotionVector.dst_x` address.
func (mv *AVMotionVector) GetDstXAddr() *int16 {
	return (*int16)(&mv.dst_x)
}

// GetDstY gets `AVMotionVector.dst_y` value.
func (mv *AVMotionVector) GetDstY() int16 {
	return (int16)(mv.dst_y)
}

// SetDstY sets `AVMotionVector.dst_y` value.
func (mv *AVMotionVector) SetDstY(v int16) {
	mv.dst_y = (C.int16_t)(v)
}

// GetDstYAddr gets `AVMotionVector.dst_y` address.
func (mv *AVMotionVector) GetDstYAddr() *int16 {
	return (*int16)(&mv.dst_y)
}

// GetFlags gets `AVMotionVector.flags` value.
func (mv *AVMotionVector) GetFlags() uint64 {
	return (uint64)(mv.flags)
}

// SetFlags sets `AVMotionVector.flags` value.
func (mv *AVMotionVector) SetFlags(v uint64) {
	mv.flags = (C.uint64_t)(v)
}

// GetFlagsAddr gets `AVMotionVector.flags` address.
func (mv *AVMotionVector) GetFlagsAddr() *uint64 {
	return (*uint64)(&mv.flags)
}

// GetMotionX gets `AVMotionVector.motion_x` value.
func (mv *AVMotionVector) GetMotionX() int32 {
	return (int32)(mv.motion_x)
}

// SetMotionX sets `AVMotionVector.motion_x` value.
func (mv *AVMotionVector) SetMotionX(v int32) {
	mv.motion_x = (C.int32_t)(v)
}

// GetMotionXAddr gets `AVMotionVector.motion_x` address.
func (mv *AVMotionVector) GetMotionXAddr() *int32 {
	return (*int32)(&mv.motion_x)
}

// GetMotionY gets `AVMotionVector.motion_y` value.
func (mv *AVMotionVector) GetMotionY() int32 {
	return (int32)(mv.motion_y)
}

// SetMotionY sets `AVMotionVector.motion_y` value.
func (mv *AVMotionVector) SetMotionY(v int32) {
	mv.motion_y = (C.int32_t)(v)
}

// GetMotionYAddr gets `AVMotionVector.motion_y` address.
func (mv *AVMotionVector) GetMotionYAddr() *int32 {
	return (*int32)(&mv.motion_y)
}

// GetMotionScale gets `AVMotionVector.motion_scale` value.
func (mv *AVMotionVector) GetMotionScale() uint16 {
	return (uint16)(mv.motion_scale)
}

// SetMotionScale sets `AVMotionVector.motion_scale` value.
func (mv *AVMotionVector) SetMotionScale(v uint16) {
	mv.motion_scale = (C.uint16_t)(v)
}

// GetMotionScaleAddr gets `AVMotionVector.motion_scale` address.
func (mv *AVMotionVector) GetMotionScaleAddr() *uint16 {
	return (*uint16)(&mv.motion_scale)
}
