package ffmpeg

/*
#include <libavutil/dovi_meta.h>
*/
import "C"
import "unsafe"

// AVDOVIDecoderConfigurationRecord
type AVDOVIDecoderConfigurationRecord C.struct_AVDOVIDecoderConfigurationRecord

// Custom: GetDvVersionMajor gets `AVDOVIDecoderConfigurationRecord.dv_version_major` value.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvVersionMajor() uint8 {
	return (uint8)(dcr.dv_version_major)
}

// Custom: SetDvVersionMajor sets `AVDOVIDecoderConfigurationRecord.dv_version_major` value.
func (dcr *AVDOVIDecoderConfigurationRecord) SetDvVersionMajor(v uint8) {
	dcr.dv_version_major = (C.uint8_t)(v)
}

// Custom: GetDvVersionMajorAddr gets `AVDOVIDecoderConfigurationRecord.dv_version_major` address.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvVersionMajorAddr() *uint8 {
	return (*uint8)(&dcr.dv_version_major)
}

// Custom: GetDvVersionMinor gets `AVDOVIDecoderConfigurationRecord.dv_version_minor` value.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvVersionMinor() uint8 {
	return (uint8)(dcr.dv_version_minor)
}

// Custom: SetDvVersionMinor sets `AVDOVIDecoderConfigurationRecord.dv_version_minor` value.
func (dcr *AVDOVIDecoderConfigurationRecord) SetDvVersionMinor(v uint8) {
	dcr.dv_version_minor = (C.uint8_t)(v)
}

// Custom: GetDvVersionMinorAddr gets `AVDOVIDecoderConfigurationRecord.dv_version_minor` address.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvVersionMinorAddr() *uint8 {
	return (*uint8)(&dcr.dv_version_minor)
}

// Custom: GetDvProfile gets `AVDOVIDecoderConfigurationRecord.dv_profile` value.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvProfile() uint8 {
	return (uint8)(dcr.dv_profile)
}

// Custom: SetDvProfile sets `AVDOVIDecoderConfigurationRecord.dv_profile` value.
func (dcr *AVDOVIDecoderConfigurationRecord) SetDvProfile(v uint8) {
	dcr.dv_profile = (C.uint8_t)(v)
}

// Custom: GetDvProfileAddr gets `AVDOVIDecoderConfigurationRecord.dv_profile` address.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvProfileAddr() *uint8 {
	return (*uint8)(&dcr.dv_profile)
}

// Custom: GetDvLevel gets `AVDOVIDecoderConfigurationRecord.dv_level` value.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvLevel() uint8 {
	return (uint8)(dcr.dv_level)
}

// Custom: SetDvLevel sets `AVDOVIDecoderConfigurationRecord.dv_level` value.
func (dcr *AVDOVIDecoderConfigurationRecord) SetDvLevel(v uint8) {
	dcr.dv_level = (C.uint8_t)(v)
}

// Custom: GetDvLevelAddr gets `AVDOVIDecoderConfigurationRecord.dv_level` address.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvLevelAddr() *uint8 {
	return (*uint8)(&dcr.dv_level)
}

// Custom: GetRpuPresentFlag gets `AVDOVIDecoderConfigurationRecord.rpu_present_flag` value.
func (dcr *AVDOVIDecoderConfigurationRecord) GetRpuPresentFlag() uint8 {
	return (uint8)(dcr.rpu_present_flag)
}

// Custom: SetRpuPresentFlag sets `AVDOVIDecoderConfigurationRecord.rpu_present_flag` value.
func (dcr *AVDOVIDecoderConfigurationRecord) SetRpuPresentFlag(v uint8) {
	dcr.rpu_present_flag = (C.uint8_t)(v)
}

// Custom: GetRpuPresentFlagAddr gets `AVDOVIDecoderConfigurationRecord.rpu_present_flag` address.
func (dcr *AVDOVIDecoderConfigurationRecord) GetRpuPresentFlagAddr() *uint8 {
	return (*uint8)(&dcr.rpu_present_flag)
}

// Custom: GetElPresentFlag gets `AVDOVIDecoderConfigurationRecord.el_present_flag` value.
func (dcr *AVDOVIDecoderConfigurationRecord) GetElPresentFlag() uint8 {
	return (uint8)(dcr.el_present_flag)
}

// Custom: SetElPresentFlag sets `AVDOVIDecoderConfigurationRecord.el_present_flag` value.
func (dcr *AVDOVIDecoderConfigurationRecord) SetElPresentFlag(v uint8) {
	dcr.el_present_flag = (C.uint8_t)(v)
}

// Custom: GetElPresentFlagAddr gets `AVDOVIDecoderConfigurationRecord.el_present_flag` address.
func (dcr *AVDOVIDecoderConfigurationRecord) GetElPresentFlagAddr() *uint8 {
	return (*uint8)(&dcr.el_present_flag)
}

// Custom: GetBlPresentFlag gets `AVDOVIDecoderConfigurationRecord.bl_present_flag` value.
func (dcr *AVDOVIDecoderConfigurationRecord) GetBlPresentFlag() uint8 {
	return (uint8)(dcr.bl_present_flag)
}

// Custom: SetBlPresentFlag sets `AVDOVIDecoderConfigurationRecord.bl_present_flag` value.
func (dcr *AVDOVIDecoderConfigurationRecord) SetBlPresentFlag(v uint8) {
	dcr.bl_present_flag = (C.uint8_t)(v)
}

// Custom: GetBlPresentFlagAddr gets `AVDOVIDecoderConfigurationRecord.bl_present_flag` address.
func (dcr *AVDOVIDecoderConfigurationRecord) GetBlPresentFlagAddr() *uint8 {
	return (*uint8)(&dcr.bl_present_flag)
}

// Custom: GetDvBlSignalCompatibilityId
// gets `AVDOVIDecoderConfigurationRecord.dv_bl_signal_compatibility_id` value.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvBlSignalCompatibilityId() uint8 {
	return (uint8)(dcr.dv_bl_signal_compatibility_id)
}

// Custom: SetDvBlSignalCompatibilityId
// sets `AVDOVIDecoderConfigurationRecord.dv_bl_signal_compatibility_id` value.
func (dcr *AVDOVIDecoderConfigurationRecord) SetDvBlSignalCompatibilityId(v uint8) {
	dcr.dv_bl_signal_compatibility_id = (C.uint8_t)(v)
}

// Custom: GetDvBlSignalCompatibilityIdAddr
// gets `AVDOVIDecoderConfigurationRecord.dv_bl_signal_compatibility_id` address.
func (dcr *AVDOVIDecoderConfigurationRecord) GetDvBlSignalCompatibilityIdAddr() *uint8 {
	return (*uint8)(&dcr.dv_bl_signal_compatibility_id)
}

// AvDoviAlloc allocates a AVDOVIDecoderConfigurationRecord structure and initialize its
// fields to default values.
func AvDoviAlloc(size *uintptr) *AVDOVIDecoderConfigurationRecord {
	return (*AVDOVIDecoderConfigurationRecord)(C.av_dovi_alloc((*C.size_t)(unsafe.Pointer(size))))
}
