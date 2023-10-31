// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavcodec/defs.h>
*/
import "C"
import "unsafe"

const (
	// Required number of additionally allocated bytes at the end of the input bitstream for decoding.
	AV_INPUT_BUFFER_PADDING_SIZE = C.AV_INPUT_BUFFER_PADDING_SIZE
)

// AVDiscard
type AVDiscard = C.enum_AVDiscard

const (
	AVDISCARD_NONE     = AVDiscard(C.AVDISCARD_NONE)
	AVDISCARD_DEFAULT  = AVDiscard(C.AVDISCARD_DEFAULT)
	AVDISCARD_NONREF   = AVDiscard(C.AVDISCARD_NONREF)
	AVDISCARD_BIDIR    = AVDiscard(C.AVDISCARD_BIDIR)
	AVDISCARD_NONINTRA = AVDiscard(C.AVDISCARD_NONINTRA)
	AVDISCARD_NONKEY   = AVDiscard(C.AVDISCARD_NONKEY)
	AVDISCARD_ALL      = AVDiscard(C.AVDISCARD_ALL)
)

// AVAudioServiceType
type AVAudioServiceType = C.enum_AVAudioServiceType

const (
	AV_AUDIO_SERVICE_TYPE_MAIN              = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_MAIN)
	AV_AUDIO_SERVICE_TYPE_EFFECTS           = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_EFFECTS)
	AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED)
	AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED  = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED)
	AV_AUDIO_SERVICE_TYPE_DIALOGUE          = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_DIALOGUE)
	AV_AUDIO_SERVICE_TYPE_COMMENTARY        = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_COMMENTARY)
	AV_AUDIO_SERVICE_TYPE_EMERGENCY         = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_EMERGENCY)
	AV_AUDIO_SERVICE_TYPE_VOICE_OVER        = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_VOICE_OVER)
	AV_AUDIO_SERVICE_TYPE_KARAOKE           = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_KARAOKE)
	AV_AUDIO_SERVICE_TYPE_NB                = AVAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_NB)
)

// Pan Scan area.
// This specifies the area which should be displayed.
// Note there may be multiple such areas for one frame.
type AVPanScan C.struct_AVPanScan

// GetId gets `AVPanScan.id` value.
func (psn *AVPanScan) GetId() int32 {
	return (int32)(psn.id)
}

// SetId sets `AVPanScan.id` value.
func (psn *AVPanScan) SetId(v int32) {
	psn.id = (C.int)(v)
}

// GetIdAddr gets `AVPanScan.id` address.
func (psn *AVPanScan) GetIdAddr() *int32 {
	return (*int32)(&psn.id)
}

// GetWidth gets `AVPanScan.width` value.
func (psn *AVPanScan) GetWidth() int32 {
	return (int32)(psn.width)
}

// SetWidth sets `AVPanScan.width` value.
func (psn *AVPanScan) SetWidth(v int32) {
	psn.width = (C.int)(v)
}

// GetWidthAddr gets `AVPanScan.width` address.
func (psn *AVPanScan) GetWidthAddr() *int32 {
	return (*int32)(&psn.width)
}

// GetHeight gets `AVPanScan.height` value.
func (psn *AVPanScan) GetHeight() int32 {
	return (int32)(psn.height)
}

// SetHeight sets `AVPanScan.height` value.
func (psn *AVPanScan) SetHeight(v int32) {
	psn.height = (C.int)(v)
}

// GetHeightAddr gets `AVPanScan.height` address.
func (psn *AVPanScan) GetHeightAddr() *int32 {
	return (*int32)(&psn.height)
}

// GetPosition gets `AVPanScan.position` value.
func (psn *AVPanScan) GetPosition() (v [][]int16) {
	for i := 0; i < 3; i++ {
		v = append(v, unsafe.Slice((*int16)(&psn.position[i][0]), 2))
	}
	return v
}

// SetPosition sets `AVPanScan.position` value.
func (psn *AVPanScan) SetPosition(v [][]int16) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		for j := 0; j < FFMIN(len(v[i]), 2); j++ {
			psn.position[i][j] = (C.int16_t)(v[i][j])
		}
	}
}

// GetPositionAddr gets `AVPanScan.position` address.
func (psn *AVPanScan) GetPositionAddr() **int16 {
	return (**int16)(unsafe.Pointer(&psn.position))
}

// This structure describes the bitrate properties of an encoded bitstream. It
// roughly corresponds to a subset the VBV parameters for MPEG-2 or HRD
// parameters for H.264/HEVC.
type AVCPBProperties C.struct_AVCPBProperties

// GetMaxBitrate gets `AVCPBProperties.max_bitrate` value.
func (cpbp *AVCPBProperties) GetMaxBitrate() int64 {
	return (int64)(cpbp.max_bitrate)
}

// SetMaxBitrate sets `AVCPBProperties.max_bitrate` value.
func (cpbp *AVCPBProperties) SetMaxBitrate(v int64) {
	cpbp.max_bitrate = (C.int64_t)(v)
}

// GetMaxBitrateAddr gets `AVCPBProperties.max_bitrate` address.
func (cpbp *AVCPBProperties) GetMaxBitrateAddr() *int64 {
	return (*int64)(&cpbp.max_bitrate)
}

// GetMinBitrate gets `AVCPBProperties.min_bitrate` value.
func (cpbp *AVCPBProperties) GetMinBitrate() int64 {
	return (int64)(cpbp.min_bitrate)
}

// SetMinBitrate sets `AVCPBProperties.min_bitrate` value.
func (cpbp *AVCPBProperties) SetMinBitrate(v int64) {
	cpbp.min_bitrate = (C.int64_t)(v)
}

// GetMinBitrateAddr gets `AVCPBProperties.min_bitrate` address.
func (cpbp *AVCPBProperties) GetMinBitrateAddr() *int64 {
	return (*int64)(&cpbp.min_bitrate)
}

// GetAvgBitrate gets `AVCPBProperties.avg_bitrate` value.
func (cpbp *AVCPBProperties) GetAvgBitrate() int64 {
	return (int64)(cpbp.avg_bitrate)
}

// SetAvgBitrate sets `AVCPBProperties.avg_bitrate` value.
func (cpbp *AVCPBProperties) SetAvgBitrate(v int64) {
	cpbp.avg_bitrate = (C.int64_t)(v)
}

// GetAvgBitrateAddr gets `AVCPBProperties.avg_bitrate` address.
func (cpbp *AVCPBProperties) GetAvgBitrateAddr() *int64 {
	return (*int64)(&cpbp.avg_bitrate)
}

// GetBufferSize gets `AVCPBProperties.buffer_size` value.
func (cpbp *AVCPBProperties) GetBufferSize() int64 {
	return (int64)(cpbp.buffer_size)
}

// SetBufferSize sets `AVCPBProperties.buffer_size` value.
func (cpbp *AVCPBProperties) SetBufferSize(v int64) {
	cpbp.buffer_size = (C.int64_t)(v)
}

// GetBufferSizeAddr gets `AVCPBProperties.buffer_size` address.
func (cpbp *AVCPBProperties) GetBufferSizeAddr() *int64 {
	return (*int64)(&cpbp.buffer_size)
}

// GetVbvDelay gets `AVCPBProperties.vbv_delay` value.
func (cpbp *AVCPBProperties) GetVbvDelay() uint64 {
	return (uint64)(cpbp.vbv_delay)
}

// SetVbvDelay sets `AVCPBProperties.vbv_delay` value.
func (cpbp *AVCPBProperties) SetVbvDelay(v uint64) {
	cpbp.vbv_delay = (C.uint64_t)(v)
}

// GetVbvDelayAddr gets `AVCPBProperties.vbv_delay` address.
func (cpbp *AVCPBProperties) GetVbvDelayAddr() *uint64 {
	return (*uint64)(&cpbp.vbv_delay)
}

// AvCpbPropertiesAlloc allocates a CPB properties structure and initialize its fields to default
// values.
func AvCpbPropertiesAlloc(size *uintptr) *AVCPBProperties {
	return (*AVCPBProperties)(C.av_cpb_properties_alloc((*C.size_t)(unsafe.Pointer(size))))
}

// This structure supplies correlation between a packet timestamp and a wall clock
// production time. The definition follows the Producer Reference Time ('prft')
// as defined in ISO/IEC 14496-12
type AVProducerReferenceTime C.struct_AVProducerReferenceTime

// GetWallclock gets `AVProducerReferenceTime.wallclock` value.
func (prt *AVProducerReferenceTime) GetWallclock() int64 {
	return (int64)(prt.wallclock)
}

// SetWallclock sets `AVProducerReferenceTime.wallclock` value.
func (prt *AVProducerReferenceTime) SetWallclock(v int64) {
	prt.wallclock = (C.int64_t)(v)
}

// GetWallclockAddr gets `AVProducerReferenceTime.wallclock` address.
func (prt *AVProducerReferenceTime) GetWallclockAddr() *int64 {
	return (*int64)(&prt.wallclock)
}

// GetFlags gets `AVProducerReferenceTime.flags` value.
func (prt *AVProducerReferenceTime) GetFlags() int32 {
	return (int32)(prt.flags)
}

// SetFlags sets `AVProducerReferenceTime.flags` value.
func (prt *AVProducerReferenceTime) SetFlags(v int32) {
	prt.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVProducerReferenceTime.flags` address.
func (prt *AVProducerReferenceTime) GetFlagsAddr() *int32 {
	return (*int32)(&prt.flags)
}

// AvXiphlacing encodes extradata length to a buffer. Used by xiph codecs.
func AvXiphlacing(s *uint8, v int32) int32 {
	return (int32)(C.av_xiphlacing((*C.uchar)(s), (C.uint)(v)))
}
