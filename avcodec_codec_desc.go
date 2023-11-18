// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavcodec/codec_desc.h>
*/
import "C"

// AVCodecDescriptor
type AVCodecDescriptor C.struct_AVCodecDescriptor

// GetId gets `AVCodecDescriptor.id` value.
func (hwc *AVCodecDescriptor) GetId() AVCodecID {
	return (AVCodecID)(hwc.id)
}

// GetType gets `AVCodecDescriptor.type` value.
func (hwc *AVCodecDescriptor) GetType() AVMediaType {
	return (AVMediaType)(hwc._type)
}

// GetName gets `AVCodecDescriptor.name` value.
func (hwc *AVCodecDescriptor) GetName() string {
	return C.GoString(hwc.name)
}

// GetLongName gets `AVCodecDescriptor.long_name` value.
func (hwc *AVCodecDescriptor) GetLongName() string {
	return C.GoString(hwc.long_name)
}

// GetProps gets `AVCodecDescriptor.props` value.
func (hwc *AVCodecDescriptor) GetProps() int32 {
	return (int32)(hwc.props)
}

// GetMimeTypes gets `AVCodecDescriptor.mime_types` value.
func (hwc *AVCodecDescriptor) GetMimeTypes() (v []string) {
	return SliceTruncString(hwc.mime_types)
}

// GetProfiles gets `AVCodecDescriptor.profiles` value.
func (hwc *AVCodecDescriptor) GetProfiles() []AVProfile {
	return SliceTrunc((*AVProfile)(hwc.profiles), func(ap AVProfile) bool {
		return ap.GetProfile() == FF_PROFILE_UNKNOWN
	})
}

const (
	AV_CODEC_PROP_INTRA_ONLY = C.AV_CODEC_PROP_INTRA_ONLY
	AV_CODEC_PROP_LOSSY      = C.AV_CODEC_PROP_LOSSY
	AV_CODEC_PROP_LOSSLESS   = C.AV_CODEC_PROP_LOSSLESS
	AV_CODEC_PROP_REORDER    = C.AV_CODEC_PROP_REORDER
	AV_CODEC_PROP_FIELDS     = C.AV_CODEC_PROP_FIELDS
	AV_CODEC_PROP_BITMAP_SUB = C.AV_CODEC_PROP_BITMAP_SUB
	AV_CODEC_PROP_TEXT_SUB   = C.AV_CODEC_PROP_TEXT_SUB
)

// AvCodecDescriptorGet returns descriptor for given codec ID or NULL if no descriptor exists.
func AvCodecDescriptorGet(id AVCodecID) *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

// AvCodecDescriptorNext iterates over all codec descriptors known to libavcodec.
func AvCodecDescriptorNext(prev *AVCodecDescriptor) *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(prev)))
}

// AvCodecDescriptorGetByName returns codec descriptor with the given name or NULL
// if no such descriptor exists.
func AvCodecDescriptorGetByName(name string) *AVCodecDescriptor {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AVCodecDescriptor)(C.avcodec_descriptor_get_by_name((*C.char)(namePtr)))
}
