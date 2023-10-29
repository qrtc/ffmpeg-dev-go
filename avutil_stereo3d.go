// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/stereo3d.h>
*/
import "C"

// AVStereo3DType
type AVStereo3DType = C.enum_AVStereo3DType

const (
	AV_STEREO3D_2D                  = AVStereo3DType(C.AV_STEREO3D_2D)
	AV_STEREO3D_SIDEBYSIDE          = AVStereo3DType(C.AV_STEREO3D_SIDEBYSIDE)
	AV_STEREO3D_TOPBOTTOM           = AVStereo3DType(C.AV_STEREO3D_TOPBOTTOM)
	AV_STEREO3D_FRAMESEQUENCE       = AVStereo3DType(C.AV_STEREO3D_FRAMESEQUENCE)
	AV_STEREO3D_CHECKERBOARD        = AVStereo3DType(C.AV_STEREO3D_CHECKERBOARD)
	AV_STEREO3D_SIDEBYSIDE_QUINCUNX = AVStereo3DType(C.AV_STEREO3D_SIDEBYSIDE_QUINCUNX)
	AV_STEREO3D_LINES               = AVStereo3DType(C.AV_STEREO3D_LINES)
	AV_STEREO3D_COLUMNS             = AVStereo3DType(C.AV_STEREO3D_COLUMNS)
)

// AVStereo3DView
type AVStereo3DView = C.enum_AVStereo3DView

const (
	AV_STEREO3D_VIEW_PACKED = AVStereo3DView(C.AV_STEREO3D_VIEW_PACKED)
	AV_STEREO3D_VIEW_LEFT   = AVStereo3DView(C.AV_STEREO3D_VIEW_LEFT)
	AV_STEREO3D_VIEW_RIGHT  = AVStereo3DView(C.AV_STEREO3D_VIEW_RIGHT)
)

const (
	AV_STEREO3D_FLAG_INVERT = C.AV_STEREO3D_FLAG_INVERT
)

type AVStereo3D C.struct_AVStereo3D

// GetType gets `AVStereo3D.type` value.
func (s3d *AVStereo3D) GetType() AVStereo3DType {
	return (AVStereo3DType)(s3d._type)
}

// SetType sets `AVStereo3D.type` value.
func (s3d *AVStereo3D) SetType(v AVStereo3DType) {
	s3d._type = (C.enum_AVStereo3DType)(v)
}

// GetTypeAddr gets `AVStereo3D.type` address.
func (s3d *AVStereo3D) GetTypeAddr() *AVStereo3DType {
	return (*AVStereo3DType)(&s3d._type)
}

// GetFlags gets `AVStereo3D.flags` value.
func (s3d *AVStereo3D) GetFlags() int32 {
	return (int32)(s3d.flags)
}

// SetFlags sets `AVStereo3D.flags` value.
func (s3d *AVStereo3D) SetFlags(v int32) {
	s3d.flags = (C.int32_t)(v)
}

// GetFlagsAddr gets `AVStereo3D.flags` address.
func (s3d *AVStereo3D) GetFlagsAddr() *int32 {
	return (*int32)(&s3d.flags)
}

// GetView gets `AVStereo3D.view` value.
func (s3d *AVStereo3D) GetView() AVStereo3DView {
	return (AVStereo3DView)(s3d.view)
}

// SetView sets `AVStereo3D.view` value.
func (s3d *AVStereo3D) SetView(v AVStereo3DView) {
	s3d.view = (C.enum_AVStereo3DView)(v)
}

// GetViewAddr gets `AVStereo3D.view` address.
func (s3d *AVStereo3D) GetViewAddr() *AVStereo3DView {
	return (*AVStereo3DView)(&s3d.view)
}

// AvStereo3dAlloc allocates an AVStereo3D structure and set its fields to default values.
func AvStereo3dAlloc() *AVStereo3D {
	return (*AVStereo3D)(C.av_stereo3d_alloc())
}

// AvStereo3dCreateSideData allocates a complete AVFrameSideData and add it to the frame.
func AvStereo3dCreateSideData(frame *AVFrame) *AVStereo3D {
	return (*AVStereo3D)(C.av_stereo3d_create_side_data((*C.struct_AVFrame)(frame)))
}

// AvStereo3dTypeName provides a human-readable name of a given stereo3d type.
func AvStereo3dTypeName(_type uint32) string {
	return C.GoString(C.av_stereo3d_type_name((C.uint)(_type)))
}

// AvStereo3dFromName gets the AVStereo3DType form a human-readable name.
func AvStereo3dFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_stereo3d_from_name((*C.char)(namePtr)))
}
