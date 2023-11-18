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

const (
	AV_EF_CRCCHECK   = int32(C.AV_EF_CRCCHECK)
	AV_EF_BITSTREAM  = int32(C.AV_EF_BITSTREAM)
	AV_EF_BUFFER     = int32(C.AV_EF_BUFFER)
	AV_EF_EXPLODE    = int32(C.AV_EF_EXPLODE)
	AV_EF_IGNORE_ERR = int32(C.AV_EF_IGNORE_ERR)
	AV_EF_CAREFUL    = int32(C.AV_EF_CAREFUL)
	AV_EF_COMPLIANT  = int32(C.AV_EF_COMPLIANT)
	AV_EF_AGGRESSIVE = int32(C.AV_EF_AGGRESSIVE)
)

const (
	FF_COMPLIANCE_VERY_STRICT  = int32(C.FF_COMPLIANCE_VERY_STRICT)
	FF_COMPLIANCE_STRICT       = int32(C.FF_COMPLIANCE_STRICT)
	FF_COMPLIANCE_NORMAL       = int32(C.FF_COMPLIANCE_NORMAL)
	FF_COMPLIANCE_UNOFFICIAL   = int32(C.FF_COMPLIANCE_UNOFFICIAL)
	FF_COMPLIANCE_EXPERIMENTAL = int32(C.FF_COMPLIANCE_EXPERIMENTAL)
)

const (
	AV_PROFILE_UNKNOWN                               = int32(C.AV_PROFILE_UNKNOWN)
	AV_PROFILE_RESERVED                              = int32(C.AV_PROFILE_RESERVED)
	AV_PROFILE_AAC_MAIN                              = int32(C.AV_PROFILE_AAC_MAIN)
	AV_PROFILE_AAC_LOW                               = int32(C.AV_PROFILE_AAC_LOW)
	AV_PROFILE_AAC_SSR                               = int32(C.AV_PROFILE_AAC_SSR)
	AV_PROFILE_AAC_LTP                               = int32(C.AV_PROFILE_AAC_LTP)
	AV_PROFILE_AAC_HE                                = int32(C.AV_PROFILE_AAC_HE)
	AV_PROFILE_AAC_HE_V2                             = int32(C.AV_PROFILE_AAC_HE_V2)
	AV_PROFILE_AAC_LD                                = int32(C.AV_PROFILE_AAC_LD)
	AV_PROFILE_AAC_ELD                               = int32(C.AV_PROFILE_AAC_ELD)
	AV_PROFILE_MPEG2_AAC_LOW                         = int32(C.AV_PROFILE_MPEG2_AAC_LOW)
	AV_PROFILE_MPEG2_AAC_HE                          = int32(C.AV_PROFILE_MPEG2_AAC_HE)
	AV_PROFILE_DNXHD                                 = int32(C.AV_PROFILE_DNXHD)
	AV_PROFILE_DNXHR_LB                              = int32(C.AV_PROFILE_DNXHR_LB)
	AV_PROFILE_DNXHR_SQ                              = int32(C.AV_PROFILE_DNXHR_SQ)
	AV_PROFILE_DNXHR_HQ                              = int32(C.AV_PROFILE_DNXHR_HQ)
	AV_PROFILE_DNXHR_HQX                             = int32(C.AV_PROFILE_DNXHR_HQX)
	AV_PROFILE_DNXHR_444                             = int32(C.AV_PROFILE_DNXHR_444)
	AV_PROFILE_DTS                                   = int32(C.AV_PROFILE_DTS)
	AV_PROFILE_DTS_ES                                = int32(C.AV_PROFILE_DTS_ES)
	AV_PROFILE_DTS_96_24                             = int32(C.AV_PROFILE_DTS_96_24)
	AV_PROFILE_DTS_HD_HRA                            = int32(C.AV_PROFILE_DTS_HD_HRA)
	AV_PROFILE_DTS_HD_MA                             = int32(C.AV_PROFILE_DTS_HD_MA)
	AV_PROFILE_DTS_EXPRESS                           = int32(C.AV_PROFILE_DTS_EXPRESS)
	AV_PROFILE_DTS_HD_MA_X                           = int32(C.AV_PROFILE_DTS_HD_MA_X)
	AV_PROFILE_DTS_HD_MA_X_IMAX                      = int32(C.AV_PROFILE_DTS_HD_MA_X_IMAX)
	AV_PROFILE_EAC3_DDP_ATMOS                        = int32(C.AV_PROFILE_EAC3_DDP_ATMOS)
	AV_PROFILE_TRUEHD_ATMOS                          = int32(C.AV_PROFILE_TRUEHD_ATMOS)
	AV_PROFILE_MPEG2_422                             = int32(C.AV_PROFILE_MPEG2_422)
	AV_PROFILE_MPEG2_HIGH                            = int32(C.AV_PROFILE_MPEG2_HIGH)
	AV_PROFILE_MPEG2_SS                              = int32(C.AV_PROFILE_MPEG2_SS)
	AV_PROFILE_MPEG2_SNR_SCALABLE                    = int32(C.AV_PROFILE_MPEG2_SNR_SCALABLE)
	AV_PROFILE_MPEG2_MAIN                            = int32(C.AV_PROFILE_MPEG2_MAIN)
	AV_PROFILE_MPEG2_SIMPLE                          = int32(C.AV_PROFILE_MPEG2_SIMPLE)
	AV_PROFILE_H264_CONSTRAINED                      = int32(C.AV_PROFILE_H264_CONSTRAINED)
	AV_PROFILE_H264_INTRA                            = int32(C.AV_PROFILE_H264_INTRA)
	AV_PROFILE_H264_BASELINE                         = int32(C.AV_PROFILE_H264_BASELINE)
	AV_PROFILE_H264_CONSTRAINED_BASELINE             = int32(C.AV_PROFILE_H264_CONSTRAINED_BASELINE)
	AV_PROFILE_H264_MAIN                             = int32(C.AV_PROFILE_H264_MAIN)
	AV_PROFILE_H264_EXTENDED                         = int32(C.AV_PROFILE_H264_EXTENDED)
	AV_PROFILE_H264_HIGH                             = int32(C.AV_PROFILE_H264_HIGH)
	AV_PROFILE_H264_HIGH_10                          = int32(C.AV_PROFILE_H264_HIGH_10)
	AV_PROFILE_H264_HIGH_10_INTRA                    = int32(C.AV_PROFILE_H264_HIGH_10_INTRA)
	AV_PROFILE_H264_MULTIVIEW_HIGH                   = int32(C.AV_PROFILE_H264_MULTIVIEW_HIGH)
	AV_PROFILE_H264_HIGH_422                         = int32(C.AV_PROFILE_H264_HIGH_422)
	AV_PROFILE_H264_HIGH_422_INTRA                   = int32(C.AV_PROFILE_H264_HIGH_422_INTRA)
	AV_PROFILE_H264_STEREO_HIGH                      = int32(C.AV_PROFILE_H264_STEREO_HIGH)
	AV_PROFILE_H264_HIGH_444                         = int32(C.AV_PROFILE_H264_HIGH_444)
	AV_PROFILE_H264_HIGH_444_PREDICTIVE              = int32(C.AV_PROFILE_H264_HIGH_444_PREDICTIVE)
	AV_PROFILE_H264_HIGH_444_INTRA                   = int32(C.AV_PROFILE_H264_HIGH_444_INTRA)
	AV_PROFILE_H264_CAVLC_444                        = int32(C.AV_PROFILE_H264_CAVLC_444)
	AV_PROFILE_VC1_SIMPLE                            = int32(C.AV_PROFILE_VC1_SIMPLE)
	AV_PROFILE_VC1_MAIN                              = int32(C.AV_PROFILE_VC1_MAIN)
	AV_PROFILE_VC1_COMPLEX                           = int32(C.AV_PROFILE_VC1_COMPLEX)
	AV_PROFILE_VC1_ADVANCED                          = int32(C.AV_PROFILE_VC1_ADVANCED)
	AV_PROFILE_MPEG4_SIMPLE                          = int32(C.AV_PROFILE_MPEG4_SIMPLE)
	AV_PROFILE_MPEG4_SIMPLE_SCALABLE                 = int32(C.AV_PROFILE_MPEG4_SIMPLE_SCALABLE)
	AV_PROFILE_MPEG4_CORE                            = int32(C.AV_PROFILE_MPEG4_CORE)
	AV_PROFILE_MPEG4_MAIN                            = int32(C.AV_PROFILE_MPEG4_MAIN)
	AV_PROFILE_MPEG4_N_BIT                           = int32(C.AV_PROFILE_MPEG4_N_BIT)
	AV_PROFILE_MPEG4_SCALABLE_TEXTURE                = int32(C.AV_PROFILE_MPEG4_SCALABLE_TEXTURE)
	AV_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION           = int32(C.AV_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION)
	AV_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE          = int32(C.AV_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE)
	AV_PROFILE_MPEG4_HYBRID                          = int32(C.AV_PROFILE_MPEG4_HYBRID)
	AV_PROFILE_MPEG4_ADVANCED_REAL_TIME              = int32(C.AV_PROFILE_MPEG4_ADVANCED_REAL_TIME)
	AV_PROFILE_MPEG4_CORE_SCALABLE                   = int32(C.AV_PROFILE_MPEG4_CORE_SCALABLE)
	AV_PROFILE_MPEG4_ADVANCED_CODING                 = int32(C.AV_PROFILE_MPEG4_ADVANCED_CODING)
	AV_PROFILE_MPEG4_ADVANCED_CORE                   = int32(C.AV_PROFILE_MPEG4_ADVANCED_CORE)
	AV_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE       = int32(C.AV_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE)
	AV_PROFILE_MPEG4_SIMPLE_STUDIO                   = int32(C.AV_PROFILE_MPEG4_SIMPLE_STUDIO)
	AV_PROFILE_MPEG4_ADVANCED_SIMPLE                 = int32(C.AV_PROFILE_MPEG4_ADVANCED_SIMPLE)
	AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0        = int32(C.AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0)
	AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1        = int32(C.AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1)
	AV_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION       = int32(C.AV_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION)
	AV_PROFILE_JPEG2000_DCINEMA_2K                   = int32(C.AV_PROFILE_JPEG2000_DCINEMA_2K)
	AV_PROFILE_JPEG2000_DCINEMA_4K                   = int32(C.AV_PROFILE_JPEG2000_DCINEMA_4K)
	AV_PROFILE_VP9_0                                 = int32(C.AV_PROFILE_VP9_0)
	AV_PROFILE_VP9_1                                 = int32(C.AV_PROFILE_VP9_1)
	AV_PROFILE_VP9_2                                 = int32(C.AV_PROFILE_VP9_2)
	AV_PROFILE_VP9_3                                 = int32(C.AV_PROFILE_VP9_3)
	AV_PROFILE_HEVC_MAIN                             = int32(C.AV_PROFILE_HEVC_MAIN)
	AV_PROFILE_HEVC_MAIN_10                          = int32(C.AV_PROFILE_HEVC_MAIN_10)
	AV_PROFILE_HEVC_MAIN_STILL_PICTURE               = int32(C.AV_PROFILE_HEVC_MAIN_STILL_PICTURE)
	AV_PROFILE_HEVC_REXT                             = int32(C.AV_PROFILE_HEVC_REXT)
	AV_PROFILE_HEVC_SCC                              = int32(C.AV_PROFILE_HEVC_SCC)
	AV_PROFILE_VVC_MAIN_10                           = int32(C.AV_PROFILE_VVC_MAIN_10)
	AV_PROFILE_VVC_MAIN_10_444                       = int32(C.AV_PROFILE_VVC_MAIN_10_444)
	AV_PROFILE_AV1_MAIN                              = int32(C.AV_PROFILE_AV1_MAIN)
	AV_PROFILE_AV1_HIGH                              = int32(C.AV_PROFILE_AV1_HIGH)
	AV_PROFILE_AV1_PROFESSIONAL                      = int32(C.AV_PROFILE_AV1_PROFESSIONAL)
	AV_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT            = int32(C.AV_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT)
	AV_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT = int32(C.AV_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT)
	AV_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT         = int32(C.AV_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT)
	AV_PROFILE_MJPEG_HUFFMAN_LOSSLESS                = int32(C.AV_PROFILE_MJPEG_HUFFMAN_LOSSLESS)
	AV_PROFILE_MJPEG_JPEG_LS                         = int32(C.AV_PROFILE_MJPEG_JPEG_LS)
	AV_PROFILE_SBC_MSBC                              = int32(C.AV_PROFILE_SBC_MSBC)
	AV_PROFILE_PRORES_PROXY                          = int32(C.AV_PROFILE_PRORES_PROXY)
	AV_PROFILE_PRORES_LT                             = int32(C.AV_PROFILE_PRORES_LT)
	AV_PROFILE_PRORES_STANDARD                       = int32(C.AV_PROFILE_PRORES_STANDARD)
	AV_PROFILE_PRORES_HQ                             = int32(C.AV_PROFILE_PRORES_HQ)
	AV_PROFILE_PRORES_4444                           = int32(C.AV_PROFILE_PRORES_4444)
	AV_PROFILE_PRORES_XQ                             = int32(C.AV_PROFILE_PRORES_XQ)
	AV_PROFILE_ARIB_PROFILE_A                        = int32(C.AV_PROFILE_ARIB_PROFILE_A)
	AV_PROFILE_ARIB_PROFILE_C                        = int32(C.AV_PROFILE_ARIB_PROFILE_C)
	AV_PROFILE_KLVA_SYNC                             = int32(C.AV_PROFILE_KLVA_SYNC)
	AV_PROFILE_KLVA_ASYNC                            = int32(C.AV_PROFILE_KLVA_ASYNC)
	AV_PROFILE_EVC_BASELINE                          = int32(C.AV_PROFILE_EVC_BASELINE)
	AV_PROFILE_EVC_MAIN                              = int32(C.AV_PROFILE_EVC_MAIN)
)

const (
	AV_LEVEL_UNKNOWN = int32(C.AV_LEVEL_UNKNOWN)
)

// AVFieldOrder
type AVFieldOrder = C.enum_AVFieldOrder

const (
	AV_FIELD_UNKNOWN     = AVFieldOrder(C.AV_FIELD_UNKNOWN)
	AV_FIELD_PROGRESSIVE = AVFieldOrder(C.AV_FIELD_PROGRESSIVE)
	AV_FIELD_TT          = AVFieldOrder(C.AV_FIELD_TT)
	AV_FIELD_BB          = AVFieldOrder(C.AV_FIELD_BB)
	AV_FIELD_TB          = AVFieldOrder(C.AV_FIELD_TB)
	AV_FIELD_BT          = AVFieldOrder(C.AV_FIELD_BT)
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
