package ffmpeg

/*
#include <libavutil/mastering_display_metadata.h>
*/
import "C"
import "unsafe"

// AVMasteringDisplayMetadata
type AVMasteringDisplayMetadata C.struct_AVMasteringDisplayMetadata

// GetDisplayPrimaries gets `AVMasteringDisplayMetadata.display_primaries` value.
func (mdm *AVMasteringDisplayMetadata) GetDisplayPrimaries() (v [][]AVRational) {
	for i := 0; i < 3; i++ {
		v = append(v, unsafe.Slice((*AVRational)(&mdm.display_primaries[i][0]), 2))
	}
	return v
}

// SetDisplayPrimaries sets `AVMasteringDisplayMetadata.display_primaries` value.
func (mdm *AVMasteringDisplayMetadata) SetDisplayPrimaries(v [][]AVRational) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		for j := 0; j < FFMIN(len(v[i]), 2); j++ {
			mdm.display_primaries[i][j] = (C.struct_AVRational)(v[i][j])
		}
	}
}

// GetDisplayPrimariesAddr gets `AVMasteringDisplayMetadata.display_primaries` address.
func (mdm *AVMasteringDisplayMetadata) GetDisplayPrimariesAddr() **AVRational {
	return (**AVRational)(unsafe.Pointer(&mdm.display_primaries))
}

// GetWhitePoint gets `AVMasteringDisplayMetadata.white_point` value.
func (mdm *AVMasteringDisplayMetadata) GetWhitePoint() []AVRational {
	return unsafe.Slice((*AVRational)(&mdm.white_point[0]), 2)
}

// SetWhitePoint sets `AVMasteringDisplayMetadata.white_point` value.
func (mdm *AVMasteringDisplayMetadata) SetWhitePoint(v []AVRational) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		mdm.white_point[i] = (C.struct_AVRational)(v[i])
	}
}

// GetWhitePointAddr gets `AVMasteringDisplayMetadata.white_point` address.
func (mdm *AVMasteringDisplayMetadata) GetWhitePointAddr() **AVRational {
	return (**AVRational)(unsafe.Pointer(&mdm.white_point))
}

// GetMinLuminance gets `AVMasteringDisplayMetadata.min_luminance` value.
func (mdm *AVMasteringDisplayMetadata) GetMinLuminance() AVRational {
	return (AVRational)(mdm.min_luminance)
}

// SetMinLuminance sets `AVMasteringDisplayMetadata.min_luminance` value.
func (mdm *AVMasteringDisplayMetadata) SetMinLuminance(v AVRational) {
	mdm.min_luminance = (C.struct_AVRational)(v)
}

// GetMinLuminanceAddr gets `AVMasteringDisplayMetadata.min_luminance` address.
func (mdm *AVMasteringDisplayMetadata) GetMinLuminanceAddr() *AVRational {
	return (*AVRational)(&mdm.min_luminance)
}

// GetMaxLuminance gets `AVMasteringDisplayMetadata.max_luminance` value.
func (mdm *AVMasteringDisplayMetadata) GetMaxLuminance() AVRational {
	return (AVRational)(mdm.max_luminance)
}

// SetMaxLuminance sets `AVMasteringDisplayMetadata.max_luminance` value.
func (mdm *AVMasteringDisplayMetadata) SetMaxLuminance(v AVRational) {
	mdm.max_luminance = (C.struct_AVRational)(v)
}

// GetMaxLuminanceAddr gets `AVMasteringDisplayMetadata.max_luminance` address.
func (mdm *AVMasteringDisplayMetadata) GetMaxLuminanceAddr() *AVRational {
	return (*AVRational)(&mdm.max_luminance)
}

// GetHasPrimaries gets `AVMasteringDisplayMetadata.has_primaries` value.
func (mdm *AVMasteringDisplayMetadata) GetHasPrimaries() int32 {
	return (int32)(mdm.has_primaries)
}

// SetHasPrimaries sets `AVMasteringDisplayMetadata.has_primaries` value.
func (mdm *AVMasteringDisplayMetadata) SetHasPrimaries(v int32) {
	mdm.has_primaries = (C.int)(v)
}

// GetHasPrimariesAddr gets `AVMasteringDisplayMetadata.has_primaries` address.
func (mdm *AVMasteringDisplayMetadata) GetHasPrimariesAddr() *int32 {
	return (*int32)(&mdm.has_primaries)
}

// GetHasLuminance gets `AVMasteringDisplayMetadata.has_luminance` value.
func (mdm *AVMasteringDisplayMetadata) GetHasLuminance() int32 {
	return (int32)(mdm.has_luminance)
}

// SetHasLuminance sets `AVMasteringDisplayMetadata.has_luminance` value.
func (mdm *AVMasteringDisplayMetadata) SetHasLuminance(v int32) {
	mdm.has_luminance = (C.int)(v)
}

// GetHasLuminanceAddr gets `AVMasteringDisplayMetadata.has_luminance` address.
func (mdm *AVMasteringDisplayMetadata) GetHasLuminanceAddr() *int32 {
	return (*int32)(&mdm.has_luminance)
}

// AvMasteringDisplayMetadataAlloc allocates an AVMasteringDisplayMetadata structure
// and set its fields to default values.
func AvMasteringDisplayMetadataAlloc() *AVMasteringDisplayMetadata {
	return (*AVMasteringDisplayMetadata)(C.av_mastering_display_metadata_alloc())
}

// AvMasteringDisplayMetadataCreateSideData allocates a complete AVMasteringDisplayMetadata
// and add it to the frame.
func AvMasteringDisplayMetadataCreateSideData(frame *AVFrame) *AVMasteringDisplayMetadata {
	return (*AVMasteringDisplayMetadata)(C.av_mastering_display_metadata_create_side_data(
		(*C.struct_AVFrame)(frame)))
}

// AVContentLightMetadata
type AVContentLightMetadata C.struct_AVContentLightMetadata

// GetMaxCLL gets `AVContentLightMetadata.MaxCLL` value.
func (clm *AVContentLightMetadata) GetMaxCLL() uint32 {
	return (uint32)(clm.MaxCLL)
}

// SetMaxCLL sets `AVContentLightMetadata.MaxCLL` value.
func (clm *AVContentLightMetadata) SetMaxCLL(v uint32) {
	clm.MaxCLL = (C.uint)(v)
}

// GetMaxCLLAddr gets `AVContentLightMetadata.MaxCLL` address.
func (clm *AVContentLightMetadata) GetMaxCLLAddr() *uint32 {
	return (*uint32)(&clm.MaxCLL)
}

// GetMaxFALL gets `AVContentLightMetadata.MaxFALL` value.
func (clm *AVContentLightMetadata) GetMaxFALL() uint32 {
	return (uint32)(clm.MaxFALL)
}

// SetMaxFALL sets `AVContentLightMetadata.MaxFALL` value.
func (clm *AVContentLightMetadata) SetMaxFALL(v uint32) {
	clm.MaxFALL = (C.uint)(v)
}

// GetMaxFALLAddr gets `AVContentLightMetadata.MaxFALL` address.
func (clm *AVContentLightMetadata) GetMaxFALLAddr() *uint32 {
	return (*uint32)(&clm.MaxFALL)
}

// AvContentLightMetadataAlloc allocates an AVContentLightMetadata structure and set its fields to
// default values.
func AvContentLightMetadataAlloc(size *uintptr) *AVContentLightMetadata {
	return (*AVContentLightMetadata)(C.av_content_light_metadata_alloc(
		(*C.size_t)(unsafe.Pointer(size))))
}

// AvContentLightMetadataCreateSideData allocates a complete AVContentLightMetadata and add it to the frame.
func AvContentLightMetadataCreateSideData(frame *AVFrame) *AVContentLightMetadata {
	return (*AVContentLightMetadata)(C.av_content_light_metadata_create_side_data(
		(*C.struct_AVFrame)(frame)))
}
