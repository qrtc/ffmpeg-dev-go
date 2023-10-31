// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavcodec/avcodec.h>

typedef void (*avcodec_context_draw_horiz_band_func)(struct AVCodecContext *s,
	const AVFrame *src, int offset[AV_NUM_DATA_POINTERS],
	int y, int type, int height);

typedef enum AVPixelFormat (*avcodec_context_get_format_func)(struct AVCodecContext *s,
	const enum AVPixelFormat * fmt);

typedef int (*avcodec_context_get_buffer2_func)(struct AVCodecContext *s, AVFrame *frame, int flags);

typedef void (*avcodec_context_rtp_callback_func)(struct AVCodecContext *avctx, void *data, int size, int mb_nb);

typedef int (*avcodec_context_get_encode_buffer_func)(struct AVCodecContext *s,
	 AVPacket *pkt, int flags);

typedef int (*avcodec_context_execute_func)(AVCodecContext *c2, void *arg2);

typedef int (*avcodec_context_internal_execute_func)(struct AVCodecContext *c,
	avcodec_context_execute_func func, void *arg2, int *ret, int count, int size);

typedef int (*avcodec_context_execute2_func)(AVCodecContext *c2, void *arg2, int, int);

typedef int (*avcodec_context_internal_execute2_func)(struct AVCodecContext *c,
	avcodec_context_execute2_func func, void *arg2, int *ret, int count);

typedef int (*av_lockmgr_cb)(void **mutex, enum AVLockOp op);

typedef int (*avcodec_parser_init_func)(AVCodecParserContext *s);

typedef int (*avcodec_parser_parse_func)(AVCodecParserContext *s,
		AVCodecContext *avctx,
		const uint8_t **poutbuf, int *poutbuf_size,
		const uint8_t *buf, int buf_size);

typedef void (*avcodec_parser_close_func)(AVCodecParserContext *s);

typedef int (*avcodec_parser_split_func)(AVCodecContext *avctx, const uint8_t *buf, int buf_size);

*/
import "C"
import (
	"unsafe"
)

const (
	// Required number of additionally allocated bytes at the end of the input bitstream for decoding.
	AV_INPUT_BUFFER_PADDING_SIZE = C.AV_INPUT_BUFFER_PADDING_SIZE
	// Minimum encoding buffer size.
	AV_INPUT_BUFFER_MIN_SIZE = C.AV_INPUT_BUFFER_MIN_SIZE
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

// RcOverride
type RcOverride C.struct_RcOverride

// These flags can be passed in AVCodecContext.flags before initialization.
// Note: Not everything is supported yet.
const (
	AV_CODEC_FLAG_UNALIGNED            = C.AV_CODEC_FLAG_UNALIGNED
	AV_CODEC_FLAG_QSCALE               = C.AV_CODEC_FLAG_QSCALE
	AV_CODEC_FLAG_4MV                  = C.AV_CODEC_FLAG_4MV
	AV_CODEC_FLAG_OUTPUT_CORRUPT       = C.AV_CODEC_FLAG_OUTPUT_CORRUPT
	AV_CODEC_FLAG_QPEL                 = C.AV_CODEC_FLAG_QPEL
	AV_CODEC_FLAG_DROPCHANGED          = C.AV_CODEC_FLAG_DROPCHANGED
	AV_CODEC_FLAG_PASS1                = C.AV_CODEC_FLAG_PASS1
	AV_CODEC_FLAG_PASS2                = C.AV_CODEC_FLAG_PASS2
	AV_CODEC_FLAG_LOOP_FILTER          = C.AV_CODEC_FLAG_LOOP_FILTER
	AV_CODEC_FLAG_GRAY                 = C.AV_CODEC_FLAG_GRAY
	AV_CODEC_FLAG_PSNR                 = C.AV_CODEC_FLAG_PSNR
	AV_CODEC_FLAG_TRUNCATED            = C.AV_CODEC_FLAG_TRUNCATED
	AV_CODEC_FLAG_INTERLACED_DCT       = C.AV_CODEC_FLAG_INTERLACED_DCT
	AV_CODEC_FLAG_LOW_DELAY            = C.AV_CODEC_FLAG_LOW_DELAY
	AV_CODEC_FLAG_GLOBAL_HEADER        = C.AV_CODEC_FLAG_GLOBAL_HEADER
	AV_CODEC_FLAG_BITEXACT             = C.AV_CODEC_FLAG_BITEXACT
	AV_CODEC_FLAG_AC_PRED              = C.AV_CODEC_FLAG_AC_PRED
	AV_CODEC_FLAG_INTERLACED_ME        = C.AV_CODEC_FLAG_INTERLACED_ME
	AV_CODEC_FLAG_CLOSED_GOP           = C.AV_CODEC_FLAG_CLOSED_GOP
	AV_CODEC_FLAG2_FAST                = C.AV_CODEC_FLAG2_FAST
	AV_CODEC_FLAG2_NO_OUTPUT           = C.AV_CODEC_FLAG2_NO_OUTPUT
	AV_CODEC_FLAG2_LOCAL_HEADER        = C.AV_CODEC_FLAG2_LOCAL_HEADER
	AV_CODEC_FLAG2_DROP_FRAME_TIMECODE = C.AV_CODEC_FLAG2_DROP_FRAME_TIMECODE
	AV_CODEC_FLAG2_CHUNKS              = C.AV_CODEC_FLAG2_CHUNKS
	AV_CODEC_FLAG2_IGNORE_CROP         = C.AV_CODEC_FLAG2_IGNORE_CROP
	AV_CODEC_FLAG2_SHOW_ALL            = C.AV_CODEC_FLAG2_SHOW_ALL
	AV_CODEC_FLAG2_EXPORT_MVS          = C.AV_CODEC_FLAG2_EXPORT_MVS
	AV_CODEC_FLAG2_SKIP_MANUAL         = C.AV_CODEC_FLAG2_SKIP_MANUAL
	AV_CODEC_FLAG2_RO_FLUSH_NOOP       = C.AV_CODEC_FLAG2_RO_FLUSH_NOOP
)

// Exported side data.
// These flags can be passed in AVCodecContext.export_side_data before initialization.
const (
	AV_CODEC_EXPORT_DATA_MVS              = C.AV_CODEC_EXPORT_DATA_MVS
	AV_CODEC_EXPORT_DATA_PRFT             = C.AV_CODEC_EXPORT_DATA_PRFT
	AV_CODEC_EXPORT_DATA_VIDEO_ENC_PARAMS = C.AV_CODEC_EXPORT_DATA_VIDEO_ENC_PARAMS
	AV_CODEC_EXPORT_DATA_FILM_GRAIN       = C.AV_CODEC_EXPORT_DATA_FILM_GRAIN
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
func (cpbp *AVCPBProperties) GetMaxBitrate() int32 {
	return (int32)(cpbp.max_bitrate)
}

// SetMaxBitrate sets `AVCPBProperties.max_bitrate` value.
func (cpbp *AVCPBProperties) SetMaxBitrate(v int32) {
	cpbp.max_bitrate = (C.int)(v)
}

// GetMaxBitrateAddr gets `AVCPBProperties.max_bitrate` address.
func (cpbp *AVCPBProperties) GetMaxBitrateAddr() *int32 {
	return (*int32)(&cpbp.max_bitrate)
}

// GetMinBitrate gets `AVCPBProperties.min_bitrate` value.
func (cpbp *AVCPBProperties) GetMinBitrate() int32 {
	return (int32)(cpbp.min_bitrate)
}

// SetMinBitrate sets `AVCPBProperties.min_bitrate` value.
func (cpbp *AVCPBProperties) SetMinBitrate(v int32) {
	cpbp.min_bitrate = (C.int)(v)
}

// GetMinBitrateAddr gets `AVCPBProperties.min_bitrate` address.
func (cpbp *AVCPBProperties) GetMinBitrateAddr() *int32 {
	return (*int32)(&cpbp.min_bitrate)
}

// GetAvgBitrate gets `AVCPBProperties.avg_bitrate` value.
func (cpbp *AVCPBProperties) GetAvgBitrate() int32 {
	return (int32)(cpbp.avg_bitrate)
}

// SetAvgBitrate sets `AVCPBProperties.avg_bitrate` value.
func (cpbp *AVCPBProperties) SetAvgBitrate(v int32) {
	cpbp.avg_bitrate = (C.int)(v)
}

// GetAvgBitrateAddr gets `AVCPBProperties.avg_bitrate` address.
func (cpbp *AVCPBProperties) GetAvgBitrateAddr() *int32 {
	return (*int32)(&cpbp.avg_bitrate)
}

// GetBufferSize gets `AVCPBProperties.buffer_size` value.
func (cpbp *AVCPBProperties) GetBufferSize() int32 {
	return (int32)(cpbp.buffer_size)
}

// SetBufferSize sets `AVCPBProperties.buffer_size` value.
func (cpbp *AVCPBProperties) SetBufferSize(v int32) {
	cpbp.buffer_size = (C.int)(v)
}

// GetBufferSizeAddr gets `AVCPBProperties.buffer_size` address.
func (cpbp *AVCPBProperties) GetBufferSizeAddr() *int32 {
	return (*int32)(&cpbp.buffer_size)
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

const (
	AV_GET_BUFFER_FLAG_REF        = C.AV_GET_BUFFER_FLAG_REF
	AV_GET_ENCODE_BUFFER_FLAG_REF = C.AV_GET_ENCODE_BUFFER_FLAG_REF
)

// AvCodecContext is main external API structure.
type AVCodecContext C.struct_AVCodecContext

// GetAVClass gets `AVCodecContext.av_class` value.
func (avctx *AVCodecContext) GetAvClass() *AVClass {
	return (*AVClass)(avctx.av_class)
}

// SetAvClass sets `AVCodecContext.av_class` value.
func (avctx *AVCodecContext) SetAvClass(v *AVClass) {
	avctx.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVCodecContext.av_class` address.
func (avctx *AVCodecContext) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&avctx.av_class))
}

// GetLogLevelOffset gets `AVCodecContext.log_level_offset` value.
func (avctx *AVCodecContext) GetLogLevelOffset() int32 {
	return (int32)(avctx.log_level_offset)
}

// SetLogLevelOffset sets `AVCodecContext.log_level_offset` value.
func (avctx *AVCodecContext) SetLogLevelOffset(v int32) {
	avctx.log_level_offset = (C.int)(v)
}

// GetLogLevelOffsetAddr gets `AVCodecContext.log_level_offset` address.
func (avctx *AVCodecContext) GetLogLevelOffsetAddr() *int32 {
	return (*int32)(&avctx.log_level_offset)
}

// GetCodecType gets `AVCodecContext.codectype` value.
func (avctx *AVCodecContext) GetCodecType() AVMediaType {
	return (AVMediaType)(avctx.codec_type)
}

// SetCodecType sets `AVCodecContext.codectype` value.
func (avctx *AVCodecContext) SetCodecType(v AVMediaType) {
	avctx.codec_type = (C.enum_AVMediaType)(v)
}

// GetCodecTypeAddr gets `AVCodecContext.codectype` address.
func (avctx *AVCodecContext) GetCodecTypeAddr() *AVMediaType {
	return (*AVMediaType)(&avctx.codec_type)
}

// GetCodec gets `AVCodecContext.codec` value.
func (avctx *AVCodecContext) GetCodec() *AVCodec {
	return (*AVCodec)(avctx.codec)
}

// SetCodec sets `AVCodecContext.codec` value.
func (avctx *AVCodecContext) SetCodec(v *AVCodec) {
	avctx.codec = (*C.struct_AVCodec)(v)
}

// GetCodecAddr gets `AVCodecContext.codec` address.
func (avctx *AVCodecContext) GetCodecAddr() **AVCodec {
	return (**AVCodec)(unsafe.Pointer(&avctx.codec))
}

// GetCodecId gets `AVCodecContext.codec_id` value.
func (avctx *AVCodecContext) GetCodecId() AVCodecID {
	return (AVCodecID)(avctx.codec_id)
}

// SetCodecId sets `AVCodecContext.codec_id` value.
func (avctx *AVCodecContext) SetCodecId(v AVCodecID) {
	avctx.codec_id = (C.enum_AVCodecID)(v)
}

// GetCodecIdAddr gets `AVCodecContext.codec_id` address.
func (avctx *AVCodecContext) GetCodecIdAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&avctx.codec_id))
}

// GetCodecTag gets `AVCodecContext.codec_tag` value.
func (avctx *AVCodecContext) GetCodecTag() uint32 {
	return (uint32)(avctx.codec_tag)
}

// SetCodecTag sets `AVCodecContext.codec_tag` value.
func (avctx *AVCodecContext) SetCodecTag(v uint32) {
	avctx.codec_tag = (C.uint)(v)
}

// GetCodecTagAddr gets `AVCodecContext.codec_tag` address.
func (avctx *AVCodecContext) GetCodecTagAddr() *uint32 {
	return (*uint32)(&avctx.codec_tag)
}

// GetPrivData gets `AVCodecContext.priv_data` value.
func (avctx *AVCodecContext) GetPrivData() unsafe.Pointer {
	return avctx.priv_data
}

// SetPrivData sets `AVCodecContext.priv_data` value.
func (avctx *AVCodecContext) SetPrivData(v CVoidPointer) {
	avctx.priv_data = VoidPointer(v)
}

// GetPrivDataAddr gets `AVCodecContext.priv_data` address.
func (avctx *AVCodecContext) GetPrivDataAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&avctx.priv_data)
}

// GetOpaque gets `AVCodecContext.opaque` value.
func (avctx *AVCodecContext) GetOpaque() unsafe.Pointer {
	return avctx.opaque
}

// SetOpaque sets `AVCodecContext.opaque` value.
func (avctx *AVCodecContext) SetOpaque(v CVoidPointer) {
	avctx.opaque = VoidPointer(v)
}

// GetOpaqueAddr gets `AVCodecContext.opaque` address.
func (avctx *AVCodecContext) GetOpaqueAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&avctx.opaque)
}

// GetBitRate gets `AVCodecContext.bit_rate` value.
func (avctx *AVCodecContext) GetBitRate() int64 {
	return (int64)(avctx.bit_rate)
}

// SetBitRate sets `AVCodecContext.bit_rate` value.
func (avctx *AVCodecContext) SetBitRate(v int64) {
	avctx.bit_rate = (C.int64_t)(v)
}

// GetBitRateAddr gets `AVCodecContext.bit_rate` address.
func (avctx *AVCodecContext) GetBitRateAddr() *int64 {
	return (*int64)(&avctx.bit_rate)
}

// GetBitRateTolerance gets `AVCodecContext.bit_rate_tolerance` value.
func (avctx *AVCodecContext) GetBitRateTolerance() int32 {
	return (int32)(avctx.bit_rate_tolerance)
}

// SetBitRateTolerance sets `AVCodecContext.bit_rate_tolerance` value.
func (avctx *AVCodecContext) SetBitRateTolerance(v int32) {
	avctx.bit_rate_tolerance = (C.int)(v)
}

// GetBitRateToleranceAddr gets `AVCodecContext.bit_rate_tolerance` address.
func (avctx *AVCodecContext) GetBitRateToleranceAddr() *int32 {
	return (*int32)(&avctx.bit_rate_tolerance)
}

// GetGlobalQuality gets `AVCodecContext.global_quality` value.
func (avctx *AVCodecContext) GetGlobalQuality() int32 {
	return (int32)(avctx.global_quality)
}

// SetGlobalQuality sets `AVCodecContext.global_quality` value.
func (avctx *AVCodecContext) SetGlobalQuality(v int32) {
	avctx.global_quality = (C.int)(v)
}

// GetGlobalQualityAddr gets `AVCodecContext.global_quality` address.
func (avctx *AVCodecContext) GetGlobalQualityAddr() *int32 {
	return (*int32)(&avctx.global_quality)
}

// GetCompressionLevel gets `AVCodecContext.compression_level` value.
func (avctx *AVCodecContext) GetCompressionLevel() int32 {
	return (int32)(avctx.compression_level)
}

// SetCompressionLevel sets `AVCodecContext.compression_level` value.
func (avctx *AVCodecContext) SetCompressionLevel(v int32) {
	avctx.compression_level = (C.int)(v)
}

// GetCompressionLevelAddr gets `AVCodecContext.compression_level` address.
func (avctx *AVCodecContext) GetCompressionLevelAddr() *int32 {
	return (*int32)(&avctx.compression_level)
}

// GetFlags gets `AVCodecContext.flags` value.
func (avctx *AVCodecContext) GetFlags() int32 {
	return (int32)(avctx.flags)
}

// SetFlags sets `AVCodecContext.flags` value.
func (avctx *AVCodecContext) SetFlags(v int32) {
	avctx.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVCodecContext.flags` address.
func (avctx *AVCodecContext) GetFlagsAddr() *int32 {
	return (*int32)(&avctx.flags)
}

// GetFlags2 gets `AVCodecContext.flags2` value.
func (avctx *AVCodecContext) GetFlags2() int32 {
	return (int32)(avctx.flags2)
}

// SetFlags2 sets `AVCodecContext.flags2` value.
func (avctx *AVCodecContext) SetFlags2(v int32) {
	avctx.flags2 = (C.int)(v)
}

// GetFlags2Addr gets `AVCodecContext.flags2` address.
func (avctx *AVCodecContext) GetFlags2Addr() *int32 {
	return (*int32)(&avctx.flags2)
}

// GetExtradata gets `AVCodecContext.extradata` value.
func (avctx *AVCodecContext) GetExtradata() *uint8 {
	return (*uint8)(avctx.extradata)
}

// SetExtradata sets `AVCodecContext.extradata` value.
func (avctx *AVCodecContext) SetExtradata(v *uint8) {
	avctx.extradata = (*C.uint8_t)(v)
}

// GetExtradataAddr gets `AVCodecContext.extradata` address.
func (avctx *AVCodecContext) GetExtradataAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&avctx.extradata))
}

// GetExtradataSize gets `AVCodecContext.extradata_size` value.
func (avctx *AVCodecContext) GetExtradataSize() int32 {
	return (int32)(avctx.extradata_size)
}

// SetExtradataSize sets `AVCodecContext.extradata_size` value.
func (avctx *AVCodecContext) SetExtradataSize(v int32) {
	avctx.extradata_size = (C.int)(v)
}

// GetExtradataSizeAddr gets `AVCodecContext.extradata_size` address.
func (avctx *AVCodecContext) GetExtradataSizeAddr() *int32 {
	return (*int32)(&avctx.extradata_size)
}

// GetTimeBase gets `AVCodecContext.time_base` value.
func (avctx *AVCodecContext) GetTimeBase() AVRational {
	return (AVRational)(avctx.time_base)
}

// SetTimeBase sets `AVCodecContext.time_base` value.
func (avctx *AVCodecContext) SetTimeBase(v AVRational) {
	avctx.time_base = (C.struct_AVRational)(v)
}

// GetTimeBaseAddr gets `AVCodecContext.time_base` address.
func (avctx *AVCodecContext) GetTimeBaseAddr() *AVRational {
	return (*AVRational)(&avctx.time_base)
}

// GetTicksPerFrame gets `AVCodecContext.ticks_per_frame` value.
func (avctx *AVCodecContext) GetTicksPerFrame() int32 {
	return (int32)(avctx.ticks_per_frame)
}

// SetTicksPerFrame sets `AVCodecContext.ticks_per_frame` value.
func (avctx *AVCodecContext) SetTicksPerFrame(v int32) {
	avctx.ticks_per_frame = (C.int)(v)
}

// GetTicksPerFrameAddr gets `AVCodecContext.ticks_per_frame` address.
func (avctx *AVCodecContext) GetTicksPerFrameAddr() *int32 {
	return (*int32)(&avctx.ticks_per_frame)
}

// GetDelay gets `AVCodecContext.delay` value.
func (avctx *AVCodecContext) GetDelay() int32 {
	return (int32)(avctx.delay)
}

// SetDelay sets `AVCodecContext.delay` value.
func (avctx *AVCodecContext) SetDelay(v int32) {
	avctx.delay = (C.int)(v)
}

// GetDelayAddr gets `AVCodecContext.delay` address.
func (avctx *AVCodecContext) GetDelayAddr() *int32 {
	return (*int32)(&avctx.delay)
}

// GetWidth gets `AVCodecContext.width` value.
func (avctx *AVCodecContext) GetWidth() int32 {
	return (int32)(avctx.width)
}

// SetWidth sets `AVCodecContext.width` value.
func (avctx *AVCodecContext) SetWidth(v int32) {
	avctx.width = (C.int)(v)
}

// GetWidthAddr gets `AVCodecContext.width` address.
func (avctx *AVCodecContext) GetWidthAddr() *int32 {
	return (*int32)(&avctx.width)
}

// GetHeight gets `AVCodecContext.height` value.
func (avctx *AVCodecContext) GetHeight() int32 {
	return (int32)(avctx.height)
}

// SetHeight sets `AVCodecContext.height` value.
func (avctx *AVCodecContext) SetHeight(v int32) {
	avctx.height = (C.int)(v)
}

// GetHeightAddr gets `AVCodecContext.height` address.
func (avctx *AVCodecContext) GetHeightAddr() *int32 {
	return (*int32)(&avctx.height)
}

// GetCodedWidth gets `AVCodecContext.coded_width` value.
func (avctx *AVCodecContext) GetCodedWidth() int32 {
	return (int32)(avctx.coded_width)
}

// SetCodedWidth sets `AVCodecContext.coded_width` value.
func (avctx *AVCodecContext) SetCodedWidth(v int32) {
	avctx.coded_width = (C.int)(v)
}

// GetCodedWidthAddr gets `AVCodecContext.coded_width` address.
func (avctx *AVCodecContext) GetCodedWidthAddr() *int32 {
	return (*int32)(&avctx.coded_width)
}

// GetCodedHeight gets `AVCodecContext.coded_height` value.
func (avctx *AVCodecContext) GetCodedHeight() int32 {
	return (int32)(avctx.coded_height)
}

// SetCodedHeight sets `AVCodecContext.coded_height` value.
func (avctx *AVCodecContext) SetCodedHeight(v int32) {
	avctx.coded_height = (C.int)(v)
}

// GetCodedHeightAddr gets `AVCodecContext.coded_height` address.
func (avctx *AVCodecContext) GetCodedHeightAddr() *int32 {
	return (*int32)(&avctx.coded_height)
}

// GetGopSize gets `AVCodecContext.gop_size` value.
func (avctx *AVCodecContext) GetGopSize() int32 {
	return (int32)(avctx.gop_size)
}

// SetGopSize sets `AVCodecContext.gop_size` value.
func (avctx *AVCodecContext) SetGopSize(v int32) {
	avctx.gop_size = (C.int)(v)
}

// GetGopSizeAddr gets `AVCodecContext.gop_size` address.
func (avctx *AVCodecContext) GetGopSizeAddr() *int32 {
	return (*int32)(&avctx.gop_size)
}

// GetPixFmt gets `AVCodecContext.pix_fmt` value.
func (avctx *AVCodecContext) GetPixFmt() AVPixelFormat {
	return (AVPixelFormat)(avctx.pix_fmt)
}

// SetPixFmt sets `AVCodecContext.pix_fmt` value.
func (avctx *AVCodecContext) SetPixFmt(v AVPixelFormat) {
	avctx.pix_fmt = (C.enum_AVPixelFormat)(v)
}

// GetPixFmtAddr gets `AVCodecContext.pix_fmt` address.
func (avctx *AVCodecContext) GetPixFmtAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&avctx.pix_fmt)
}

// typedef void (*avcodec_context_draw_horiz_band_func)(struct AVCodecContext *s,
// const AVFrame *src, int offset[AV_NUM_DATA_POINTERS],
// int y, int type, int height);
type AVCodecContextDrawHorizBandFunc = C.avcodec_context_draw_horiz_band_func

// GetDrawHorizBand gets `AVCodecContext.draw_horiz_band` value.
func (avctx *AVCodecContext) GetDrawHorizBand() AVCodecContextDrawHorizBandFunc {
	return (AVCodecContextDrawHorizBandFunc)(avctx.draw_horiz_band)
}

// SetDrawHorizBand sets `AVCodecContext.draw_horiz_band` value.
func (avctx *AVCodecContext) SetDrawHorizBand(v AVCodecContextDrawHorizBandFunc) {
	avctx.draw_horiz_band = (C.avcodec_context_draw_horiz_band_func)(v)
}

// GetDrawHorizBandAddr gets `AVCodecContext.draw_horiz_band` value.
func (avctx *AVCodecContext) GetDrawHorizBandAddr() *AVCodecContextDrawHorizBandFunc {
	return (*AVCodecContextDrawHorizBandFunc)(&avctx.draw_horiz_band)
}

// GetMaxBFrames gets `AVCodecContext.max_b_frames` value.
func (avctx *AVCodecContext) GetMaxBFrames() int32 {
	return (int32)(avctx.max_b_frames)
}

// SetMaxBFrames sets `AVCodecContext.max_b_frames` value.
func (avctx *AVCodecContext) SetMaxBFrames(v int32) {
	avctx.max_b_frames = (C.int)(v)
}

// GetMaxBFramesAddr gets `AVCodecContext.max_b_frames` address.
func (avctx *AVCodecContext) GetMaxBFramesAddr() *int32 {
	return (*int32)(&avctx.max_b_frames)
}

// typedef enum AVPixelFormat (*avcodec_context_get_format)(struct AVCodecContext *s,
// const enum AVPixelFormat * fmt);
type AVCodecContextGetFormatFunc = C.avcodec_context_get_format_func

// GetGetFormat gets `AVCodecContext.get_format` value.
func (avctx *AVCodecContext) GetGetFormat() AVCodecContextGetFormatFunc {
	return (AVCodecContextGetFormatFunc)(avctx.get_format)
}

// SetGetFormat sets `AVCodecContext.get_format` value.
func (avctx *AVCodecContext) SetGetFormat(v AVCodecContextGetFormatFunc) {
	avctx.get_format = (C.avcodec_context_get_format_func)(v)
}

// GetGetFormatAddr gets `AVCodecContext.get_format` address.
func (avctx *AVCodecContext) GetGetFormatAddr() *AVCodecContextGetFormatFunc {
	return (*AVCodecContextGetFormatFunc)(&avctx.get_format)
}

// GetBQuantFactor gets `AVCodecContext.b_quant_factor` value.
func (avctx *AVCodecContext) GetBQuantFactor() float32 {
	return (float32)(avctx.b_quant_factor)
}

// SetBQuantFactor sets `AVCodecContext.b_quant_factor` value.
func (avctx *AVCodecContext) SetBQuantFactor(v float32) {
	avctx.b_quant_factor = (C.float)(v)
}

// GetBQuantFactorAddr gets `AVCodecContext.b_quant_factor` address.
func (avctx *AVCodecContext) GetBQuantFactorAddr() *float32 {
	return (*float32)(&avctx.b_quant_factor)
}

// Deprecated: Use encoder private options instead.
//
// GetBFrameStrategy gets `AVCodecContext.b_frame_strategy` value.
func (avctx *AVCodecContext) GetBFrameStrategy() int32 {
	return (int32)(avctx.b_frame_strategy)
}

// Deprecated: Use encoder private options instead.
//
// SetBFrameStrategy sets `AVCodecContext.b_frame_strategy` value.
func (avctx *AVCodecContext) SetBFrameStrategy(v int32) {
	avctx.b_frame_strategy = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetBFrameStrategyAddr gets `AVCodecContext.b_frame_strategy` address.
func (avctx *AVCodecContext) GetBFrameStrategyAddr() *int32 {
	return (*int32)(&avctx.b_frame_strategy)
}

// GetBQuantOffset gets `AVCodecContext.b_quant_offset` value.
func (avctx *AVCodecContext) GetBQuantOffset() float32 {
	return (float32)(avctx.b_quant_offset)
}

// SetBQuantOffset sets `AVCodecContext.b_quant_offset` value.
func (avctx *AVCodecContext) SetBQuantOffset(v float32) {
	avctx.b_quant_offset = (C.float)(v)
}

// GetBQuantOffsetAddr gets `AVCodecContext.b_quant_offset` address.
func (avctx *AVCodecContext) GetBQuantOffsetAddr() *float32 {
	return (*float32)(&avctx.b_quant_offset)
}

// GetHasBFrames gets `AVCodecContext.has_b_frames` value.
func (avctx *AVCodecContext) GetHasBFrames() int32 {
	return (int32)(avctx.has_b_frames)
}

// SetHasBFrames sets `AVCodecContext.has_b_frames` value.
func (avctx *AVCodecContext) SetHasBFrames(v int32) {
	avctx.has_b_frames = (C.int)(v)
}

// GetHasBFramesAddr gets `AVCodecContext.has_b_frames` address.
func (avctx *AVCodecContext) GetHasBFramesAddr() *int32 {
	return (*int32)(&avctx.has_b_frames)
}

// Deprecated: Use encoder private options instead.
//
// GetMpegQuant gets `AVCodecContext.mpeg_quant` value.
func (avctx *AVCodecContext) GetMpegQuant() int32 {
	return (int32)(avctx.mpeg_quant)
}

// Deprecated: Use encoder private options instead.
//
// SetMpegQuant sets `AVCodecContext.mpeg_quant` value.
func (avctx *AVCodecContext) SetMpegQuant(v int32) {
	avctx.mpeg_quant = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetMpegQuantAddr gets `AVCodecContext.mpeg_quant` address.
func (avctx *AVCodecContext) GetMpegQuantAddr() *int32 {
	return (*int32)(&avctx.mpeg_quant)
}

// GetIQuantFactor gets `AVCodecContext.i_quant_factor` value.
func (avctx *AVCodecContext) GetIQuantFactor() float32 {
	return (float32)(avctx.i_quant_factor)
}

// SetIQuantFactor sets `AVCodecContext.i_quant_factor` value.
func (avctx *AVCodecContext) SetIQuantFactor(v float32) {
	avctx.i_quant_factor = (C.float)(v)
}

// GetIQuantFactorAddr gets `AVCodecContext.i_quant_factor` address.
func (avctx *AVCodecContext) GetIQuantFactorAddr() *float32 {
	return (*float32)(&avctx.i_quant_factor)
}

// GetIQuantOffset gets `AVCodecContext.i_quant_offset` value.
func (avctx *AVCodecContext) GetIQuantOffset() float32 {
	return (float32)(avctx.i_quant_offset)
}

// SetIQuantOffset sets `AVCodecContext.i_quant_offset` value.
func (avctx *AVCodecContext) SetIQuantOffset(v float32) {
	avctx.i_quant_offset = (C.float)(v)
}

// GetIQuantOffsetAddr gets `AVCodecContext.i_quant_offset` address.
func (avctx *AVCodecContext) GetIQuantOffsetAddr() *float32 {
	return (*float32)(&avctx.i_quant_offset)
}

// GetLumiMasking gets `AVCodecContext.lumi_masking` value.
func (avctx *AVCodecContext) GetLumiMasking() float32 {
	return (float32)(avctx.lumi_masking)
}

// SetLumiMasking sets `AVCodecContext.lumi_masking` value.
func (avctx *AVCodecContext) SetLumiMasking(v float32) {
	avctx.lumi_masking = (C.float)(v)
}

// GetLumiMaskingAddr gets `AVCodecContext.lumi_masking` address.
func (avctx *AVCodecContext) GetLumiMaskingAddr() *float32 {
	return (*float32)(&avctx.lumi_masking)
}

// GetTemporalCplxMasking gets `AVCodecContext.temporal_cplx_masking` value.
func (avctx *AVCodecContext) GetTemporalCplxMasking() float32 {
	return (float32)(avctx.temporal_cplx_masking)
}

// SetTemporalCplxMasking sets `AVCodecContext.temporal_cplx_masking` value.
func (avctx *AVCodecContext) SetTemporalCplxMasking(v float32) {
	avctx.temporal_cplx_masking = (C.float)(v)
}

// GetTemporalCplxMaskingAddr gets `AVCodecContext.temporal_cplx_masking` address.
func (avctx *AVCodecContext) GetTemporalCplxMaskingAddr() *float32 {
	return (*float32)(&avctx.temporal_cplx_masking)
}

// GetSpatialCplxMasking gets `AVCodecContext.spatial_cplx_masking` value.
func (avctx *AVCodecContext) GetSpatialCplxMasking() float32 {
	return (float32)(avctx.spatial_cplx_masking)
}

// SetSpatialCplxMasking sets `AVCodecContext.spatial_cplx_masking` value.
func (avctx *AVCodecContext) SetSpatialCplxMasking(v float32) {
	avctx.spatial_cplx_masking = (C.float)(v)
}

// GetSpatialCplxMaskingAddr gets `AVCodecContext.spatial_cplx_masking` address.
func (avctx *AVCodecContext) GetSpatialCplxMaskingAddr() *float32 {
	return (*float32)(&avctx.spatial_cplx_masking)
}

// GetPMasking gets `AVCodecContext.p_masking` value.
func (avctx *AVCodecContext) GetPMasking() float32 {
	return (float32)(avctx.p_masking)
}

// SetPMasking sets `AVCodecContext.p_masking` value.
func (avctx *AVCodecContext) SetPMasking(v float32) {
	avctx.p_masking = (C.float)(v)
}

// GetPMaskingAddr gets `AVCodecContext.p_masking` address.
func (avctx *AVCodecContext) GetPMaskingAddr() *float32 {
	return (*float32)(&avctx.p_masking)
}

// GetDarkMasking gets `AVCodecContext.dark_masking` value.
func (avctx *AVCodecContext) GetDarkMasking() float32 {
	return (float32)(avctx.dark_masking)
}

// SetDarkMasking sets `AVCodecContext.dark_masking` value.
func (avctx *AVCodecContext) SetDarkMasking(v float32) {
	avctx.dark_masking = (C.float)(v)
}

// GetDarkMaskingAddr gets `AVCodecContext.dark_masking` address.
func (avctx *AVCodecContext) GetDarkMaskingAddr() *float32 {
	return (*float32)(&avctx.dark_masking)
}

// GetSliceCount gets `AVCodecContext.slice_count` value.
func (avctx *AVCodecContext) GetSliceCount() int32 {
	return (int32)(avctx.slice_count)
}

// SetSliceCount sets `AVCodecContext.slice_count` value.
func (avctx *AVCodecContext) SetSliceCount(v int32) {
	avctx.slice_count = (C.int)(v)
}

// GetSliceCountAddr gets `AVCodecContext.slice_count` address.
func (avctx *AVCodecContext) GetSliceCountAddr() *int32 {
	return (*int32)(&avctx.slice_count)
}

// Deprecated: Use encoder private options instead.
//
// GetPredictionMethod gets `AVCodecContext.prediction_method` value.
func (avctx *AVCodecContext) GetPredictionMethod() int32 {
	return (int32)(avctx.prediction_method)
}

// Deprecated: Use encoder private options instead.
//
// SetPredictionMethod sets `AVCodecContext.prediction_method` value.
func (avctx *AVCodecContext) SetPredictionMethod(v int32) {
	avctx.prediction_method = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetPredictionMethodAddr gets `AVCodecContext.prediction_method` address.
func (avctx *AVCodecContext) GetPredictionMethodAddr() *int32 {
	return (*int32)(&avctx.prediction_method)
}

const (
	FF_PRED_LEFT   = int32(C.FF_PRED_LEFT)
	FF_PRED_PLANE  = int32(C.FF_PRED_PLANE)
	FF_PRED_MEDIAN = int32(C.FF_PRED_MEDIAN)
)

// GetSliceOffset gets `AVCodecContext.slice_offset` value.
func (avctx *AVCodecContext) GetSliceOffset() *int32 {
	return (*int32)(avctx.slice_offset)
}

// SetSliceOffset sets `AVCodecContext.slice_offset` value.
func (avctx *AVCodecContext) SetSliceOffset(v *int32) {
	avctx.slice_offset = (*C.int)(v)
}

// GetSliceOffsetAddr gets `AVCodecContext.slice_offset` address.
func (avctx *AVCodecContext) GetSliceOffsetAddr() **int32 {
	return (**int32)(unsafe.Pointer(&avctx.slice_offset))
}

// GetSampleAspectRatio gets `AVCodecContext.sample_aspect_ratio` value.
func (avctx *AVCodecContext) GetSampleAspectRatio() AVRational {
	return (AVRational)(avctx.sample_aspect_ratio)
}

// SetSampleAspectRatio sets `AVCodecContext.sample_aspect_ratio` value.
func (avctx *AVCodecContext) SetSampleAspectRatio(v AVRational) {
	avctx.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// GetSampleAspectRatioAddr gets `AVCodecContext.sample_aspect_ratio` address.
func (avctx *AVCodecContext) GetSampleAspectRatioAddr() *AVRational {
	return (*AVRational)(&avctx.sample_aspect_ratio)
}

// GetMeCmp gets `AVCodecContext.me_cmp` value.
func (avctx *AVCodecContext) GetMeCmp() int32 {
	return (int32)(avctx.me_cmp)
}

// SetMeCmp sets `AVCodecContext.me_cmp` value.
func (avctx *AVCodecContext) SetMeCmp(v int32) {
	avctx.me_cmp = (C.int)(v)
}

// GetMeCmpAddr gets `AVCodecContext.me_cmp` address.
func (avctx *AVCodecContext) GetMeCmpAddr() *int32 {
	return (*int32)(&avctx.me_cmp)
}

// GetMeSubCmp gets `AVCodecContext.me_sub_cmp` value.
func (avctx *AVCodecContext) GetMeSubCmp() int32 {
	return (int32)(avctx.me_sub_cmp)
}

// SetMeSubCmp sets `AVCodecContext.me_sub_cmp` value.
func (avctx *AVCodecContext) SetMeSubCmp(v int32) {
	avctx.me_sub_cmp = (C.int)(v)
}

// GetMeSubCmpAddr gets `AVCodecContext.me_sub_cmp` address.
func (avctx *AVCodecContext) GetMeSubCmpAddr() *int32 {
	return (*int32)(&avctx.me_sub_cmp)
}

// GetMbCmp gets `AVCodecContext.mb_cmp` value.
func (avctx *AVCodecContext) GetMbCmp() int32 {
	return (int32)(avctx.mb_cmp)
}

// SetMbCmp sets `AVCodecContext.mb_cmp` value.
func (avctx *AVCodecContext) SetMbCmp(v int32) {
	avctx.mb_cmp = (C.int)(v)
}

// GetMbCmpAddr gets `AVCodecContext.mb_cmp` address.
func (avctx *AVCodecContext) GetMbCmpAddr() *int32 {
	return (*int32)(&avctx.mb_cmp)
}

// GetIldctCmp gets `AVCodecContext.ildct_cmp` value.
func (avctx *AVCodecContext) GetIldctCmp() int32 {
	return (int32)(avctx.ildct_cmp)
}

// SetIldctCmp sets `AVCodecContext.ildct_cmp` value.
func (avctx *AVCodecContext) SetIldctCmp(v int32) {
	avctx.ildct_cmp = (C.int)(v)
}

// GetIldctCmpAddr gets `AVCodecContext.ildct_cmp` address.
func (avctx *AVCodecContext) GetIldctCmpAddr() *int32 {
	return (*int32)(&avctx.ildct_cmp)
}

const (
	FF_CMP_SAD        = int32(C.FF_CMP_SAD)
	FF_CMP_SSE        = int32(C.FF_CMP_SSE)
	FF_CMP_SATD       = int32(C.FF_CMP_SATD)
	FF_CMP_DCT        = int32(C.FF_CMP_DCT)
	FF_CMP_PSNR       = int32(C.FF_CMP_PSNR)
	FF_CMP_BIT        = int32(C.FF_CMP_BIT)
	FF_CMP_RD         = int32(C.FF_CMP_RD)
	FF_CMP_ZERO       = int32(C.FF_CMP_ZERO)
	FF_CMP_VSAD       = int32(C.FF_CMP_VSAD)
	FF_CMP_VSSE       = int32(C.FF_CMP_VSSE)
	FF_CMP_NSSE       = int32(C.FF_CMP_NSSE)
	FF_CMP_W53        = int32(C.FF_CMP_W53)
	FF_CMP_W97        = int32(C.FF_CMP_W97)
	FF_CMP_DCTMAX     = int32(C.FF_CMP_DCTMAX)
	FF_CMP_DCT264     = int32(C.FF_CMP_DCT264)
	FF_CMP_MEDIAN_SAD = int32(C.FF_CMP_MEDIAN_SAD)
	FF_CMP_CHROMA     = int32(C.FF_CMP_CHROMA)
)

// GetDiaSize gets `AVCodecContext.dia_size` value.
func (avctx *AVCodecContext) GetDiaSize() int32 {
	return (int32)(avctx.dia_size)
}

// SetDiaSize sets `AVCodecContext.dia_size` value.
func (avctx *AVCodecContext) SetDiaSize(v int32) {
	avctx.dia_size = (C.int)(v)
}

// GetDiaSizeAddr gets `AVCodecContext.dia_size` address.
func (avctx *AVCodecContext) GetDiaSizeAddr() *int32 {
	return (*int32)(&avctx.dia_size)
}

// GetLastPredictorCount gets `AVCodecContext.last_predictor_count` value.
func (avctx *AVCodecContext) GetLastPredictorCount() int32 {
	return (int32)(avctx.last_predictor_count)
}

// SetLastPredictorCount sets `AVCodecContext.last_predictor_count` value.
func (avctx *AVCodecContext) SetLastPredictorCount(v int32) {
	avctx.last_predictor_count = (C.int)(v)
}

// GetLastPredictorCountAddr gets `AVCodecContext.last_predictor_count` address.
func (avctx *AVCodecContext) GetLastPredictorCountAddr() *int32 {
	return (*int32)(&avctx.last_predictor_count)
}

// Deprecated: Use encoder private options instead.
//
// GetPreMe gets `AVCodecContext.pre_me` value.
func (avctx *AVCodecContext) GetPreMe() int32 {
	return (int32)(avctx.pre_me)
}

// Deprecated: Use encoder private options instead.
//
// SetPreMe sets `AVCodecContext.pre_me` value.
func (avctx *AVCodecContext) SetPreMe(v int32) {
	avctx.pre_me = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetPreMeAddr gets `AVCodecContext.pre_me` address.
func (avctx *AVCodecContext) GetPreMeAddr() *int32 {
	return (*int32)(&avctx.pre_me)
}

// GetMePreCmp gets `AVCodecContext.me_pre_cmp` value.
func (avctx *AVCodecContext) GetMePreCmp() int32 {
	return (int32)(avctx.me_pre_cmp)
}

// SetMePreCmp sets `AVCodecContext.me_pre_cmp` value.
func (avctx *AVCodecContext) SetMePreCmp(v int32) {
	avctx.me_pre_cmp = (C.int)(v)
}

// GetMePreCmpAddr gets `AVCodecContext.me_pre_cmp` address.
func (avctx *AVCodecContext) GetMePreCmpAddr() *int32 {
	return (*int32)(&avctx.me_pre_cmp)
}

// GetPreDiaSize gets `AVCodecContext.pre_dia_size` value.
func (avctx *AVCodecContext) GetPreDiaSize() int32 {
	return (int32)(avctx.pre_dia_size)
}

// SetPreDiaSize sets `AVCodecContext.pre_dia_size` value.
func (avctx *AVCodecContext) SetPreDiaSize(v int32) {
	avctx.pre_dia_size = (C.int)(v)
}

// GetPreDiaSizeAddr gets `AVCodecContext.pre_dia_size` address.
func (avctx *AVCodecContext) GetPreDiaSizeAddr() *int32 {
	return (*int32)(&avctx.pre_dia_size)
}

// GetMeSubpelQuality gets `AVCodecContext.me_subpel_quality` value.
func (avctx *AVCodecContext) GetMeSubpelQuality() int32 {
	return (int32)(avctx.me_subpel_quality)
}

// SetMeSubpelQuality sets `AVCodecContext.me_subpel_quality` value.
func (avctx *AVCodecContext) SetMeSubpelQuality(v int32) {
	avctx.me_subpel_quality = (C.int)(v)
}

// GetMeSubpelQualityAddr gets `AVCodecContext.me_subpel_quality` address.
func (avctx *AVCodecContext) GetMeSubpelQualityAddr() *int32 {
	return (*int32)(&avctx.me_subpel_quality)
}

// GetSliceFlags gets `AVCodecContext.slice_flags` value.
func (avctx *AVCodecContext) GetSliceFlags() int32 {
	return (int32)(avctx.slice_flags)
}

// SetSliceFlags sets `AVCodecContext.slice_flags` value.
func (avctx *AVCodecContext) SetSliceFlags(v int32) {
	avctx.slice_flags = (C.int)(v)
}

// GetSliceFlagsAddr gets `AVCodecContext.slice_flags` address.
func (avctx *AVCodecContext) GetSliceFlagsAddr() *int32 {
	return (*int32)(&avctx.slice_flags)
}

const (
	SLICE_FLAG_CODED_ORDER = int32(C.SLICE_FLAG_CODED_ORDER)
	SLICE_FLAG_ALLOW_FIELD = int32(C.SLICE_FLAG_ALLOW_FIELD)
	SLICE_FLAG_ALLOW_PLANE = int32(C.SLICE_FLAG_ALLOW_PLANE)
)

// GetMbDecision gets `AVCodecContext.mb_decision` value.
func (avctx *AVCodecContext) GetMbDecision() int32 {
	return (int32)(avctx.mb_decision)
}

// SetMbDecision sets `AVCodecContext.mb_decision` value.
func (avctx *AVCodecContext) SetMbDecision(v int32) {
	avctx.mb_decision = (C.int)(v)
}

// GetMbDecisionAddr gets `AVCodecContext.mb_decision` address.
func (avctx *AVCodecContext) GetMbDecisionAddr() *int32 {
	return (*int32)(&avctx.mb_decision)
}

const (
	FF_MB_DECISION_SIMPLE = int32(C.FF_MB_DECISION_SIMPLE)
	FF_MB_DECISION_BITS   = int32(C.FF_MB_DECISION_BITS)
	FF_MB_DECISION_RD     = int32(C.FF_MB_DECISION_RD)
)

// GetIntraMatrix gets `AVCodecContext.intra_matrix` value.
func (avctx *AVCodecContext) GetIntraMatrix() *uint16 {
	return (*uint16)(avctx.intra_matrix)
}

// SetIntraMatrix sets `AVCodecContext.intra_matrix` value.
func (avctx *AVCodecContext) SetIntraMatrix(v *uint16) {
	avctx.intra_matrix = (*C.uint16_t)(v)
}

// GetIntraMatrixAddr gets `AVCodecContext.intra_matrix` address.
func (avctx *AVCodecContext) GetIntraMatrixAddr() **uint16 {
	return (**uint16)(unsafe.Pointer(&avctx.intra_matrix))
}

// GetInterMatrix gets `AVCodecContext.inter_matrix` value.
func (avctx *AVCodecContext) GetInterMatrix() *uint16 {
	return (*uint16)(avctx.inter_matrix)
}

// SetInterMatrix sets `AVCodecContext.inter_matrix` value.
func (avctx *AVCodecContext) SetInterMatrix(v *uint16) {
	avctx.inter_matrix = (*C.uint16_t)(v)
}

// GetInterMatrixAddr gets `AVCodecContext.inter_matrix` address.
func (avctx *AVCodecContext) GetInterMatrixAddr() **uint16 {
	return (**uint16)(unsafe.Pointer(&avctx.inter_matrix))
}

// Deprecated: Use encoder private options instead.
//
// GetScenechangeThreshold gets `AVCodecContext.scenechange_threshold` value.
func (avctx *AVCodecContext) GetScenechangeThreshold() int32 {
	return (int32)(avctx.scenechange_threshold)
}

// Deprecated: Use encoder private options instead.
//
// SetScenechangeThreshold sets `AVCodecContext.scenechange_threshold` value.
func (avctx *AVCodecContext) SetScenechangeThreshold(v int32) {
	avctx.scenechange_threshold = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetScenechangeThresholdAddr gets `AVCodecContext.scenechange_threshold` address.
func (avctx *AVCodecContext) GetScenechangeThresholdAddr() *int32 {
	return (*int32)(&avctx.scenechange_threshold)
}

// Deprecated: Use encoder private options instead.
//
// GetNoiseReduction gets `AVCodecContext.noise_reduction` value.
func (avctx *AVCodecContext) GetNoiseReduction() int32 {
	return (int32)(avctx.noise_reduction)
}

// Deprecated: Use encoder private options instead.
//
// SetNoiseReduction sets `AVCodecContext.noise_reduction` value.
func (avctx *AVCodecContext) SetNoiseReduction(v int32) {
	avctx.noise_reduction = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetNoiseReductionAddr gets `AVCodecContext.noise_reduction` address.
func (avctx *AVCodecContext) GetNoiseReductionAddr() *int32 {
	return (*int32)(&avctx.noise_reduction)
}

// GetIntraDcPrecision gets `AVCodecContext.intra_dc_precision` value.
func (avctx *AVCodecContext) GetIntraDcPrecision() int32 {
	return (int32)(avctx.intra_dc_precision)
}

// SetIntraDcPrecision sets `AVCodecContext.intra_dc_precision` value.
func (avctx *AVCodecContext) SetIntraDcPrecision(v int32) {
	avctx.intra_dc_precision = (C.int)(v)
}

// GetIntraDcPrecisionAddr gets `AVCodecContext.intra_dc_precision` address.
func (avctx *AVCodecContext) GetIntraDcPrecisionAddr() *int32 {
	return (*int32)(&avctx.intra_dc_precision)
}

// GetSkipTop gets `AVCodecContext.skip_top` value.
func (avctx *AVCodecContext) GetSkipTop() int32 {
	return (int32)(avctx.skip_top)
}

// SetSkipTop sets `AVCodecContext.skip_top` value.
func (avctx *AVCodecContext) SetSkipTop(v int32) {
	avctx.skip_top = (C.int)(v)
}

// GetSkipTopAddr gets `AVCodecContext.skip_top` address.
func (avctx *AVCodecContext) GetSkipTopAddr() *int32 {
	return (*int32)(&avctx.skip_top)
}

// GetSkipBottom gets `AVCodecContext.skip_bottom` value.
func (avctx *AVCodecContext) GetSkipBottom() int32 {
	return (int32)(avctx.skip_bottom)
}

// SetSkipBottom sets `AVCodecContext.skip_bottom` value.
func (avctx *AVCodecContext) SetSkipBottom(v int32) {
	avctx.skip_bottom = (C.int)(v)
}

// GetSkipBottomAddr gets `AVCodecContext.skip_bottom` address.
func (avctx *AVCodecContext) GetSkipBottomAddr() *int32 {
	return (*int32)(&avctx.skip_bottom)
}

// GetMbLmin gets `AVCodecContext.mb_lmin` value.
func (avctx *AVCodecContext) GetMbLmin() int32 {
	return (int32)(avctx.mb_lmin)
}

// SetMbLmin sets `AVCodecContext.mb_lmin` value.
func (avctx *AVCodecContext) SetMbLmin(v int32) {
	avctx.mb_lmin = (C.int)(v)
}

// GetMbLminAddr gets `AVCodecContext.mb_lmin` address.
func (avctx *AVCodecContext) GetMbLminAddr() *int32 {
	return (*int32)(&avctx.mb_lmin)
}

// GetMbLmax gets `AVCodecContext.mb_lmax` value.
func (avctx *AVCodecContext) GetMbLmax() int32 {
	return (int32)(avctx.mb_lmax)
}

// SetMbLmax sets `AVCodecContext.mb_lmax` value.
func (avctx *AVCodecContext) SetMbLmax(v int32) {
	avctx.mb_lmax = (C.int)(v)
}

// GetMbLmaxAddr gets `AVCodecContext.mb_lmax` address.
func (avctx *AVCodecContext) GetMbLmaxAddr() *int32 {
	return (*int32)(&avctx.mb_lmax)
}

// Deprecated: Use encoder private options instead.
//
// GetMePenaltyCompensation gets `AVCodecContext.me_penalty_compensation` value.
func (avctx *AVCodecContext) GetMePenaltyCompensation() int32 {
	return (int32)(avctx.me_penalty_compensation)
}

// Deprecated: Use encoder private options instead.
//
// SetMePenaltyCompensation sets `AVCodecContext.me_penalty_compensation` value.
func (avctx *AVCodecContext) SetMePenaltyCompensation(v int32) {
	avctx.me_penalty_compensation = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetMePenaltyCompensationAddr gets `AVCodecContext.me_penalty_compensation` address.
func (avctx *AVCodecContext) GetMePenaltyCompensationAddr() *int32 {
	return (*int32)(&avctx.me_penalty_compensation)
}

// GetBidirRefine gets `AVCodecContext.bidir_refine` value.
func (avctx *AVCodecContext) GetBidirRefine() int32 {
	return (int32)(avctx.bidir_refine)
}

// SetBidirRefine sets `AVCodecContext.bidir_refine` value.
func (avctx *AVCodecContext) SetBidirRefine(v int32) {
	avctx.bidir_refine = (C.int)(v)
}

// GetBidirRefineAddr gets `AVCodecContext.bidir_refine` address.
func (avctx *AVCodecContext) GetBidirRefineAddr() *int32 {
	return (*int32)(&avctx.bidir_refine)
}

// Deprecated: Use encoder private options instead.
//
// GetBrdScale gets `AVCodecContext.brd_scale` value.
func (avctx *AVCodecContext) GetBrdScale() int32 {
	return (int32)(avctx.brd_scale)
}

// Deprecated: Use encoder private options instead.
//
// SetBrdScale sets `AVCodecContext.brd_scale` value.
func (avctx *AVCodecContext) SetBrdScale(v int32) {
	avctx.brd_scale = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetBrdScaleAddr gets `AVCodecContext.brd_scale` address.
func (avctx *AVCodecContext) GetBrdScaleAddr() *int32 {
	return (*int32)(&avctx.brd_scale)
}

// GetKeyintMin gets `AVCodecContext.keyint_min` value.
func (avctx *AVCodecContext) GetKeyintMin() int32 {
	return (int32)(avctx.keyint_min)
}

// SetKeyintMin sets `AVCodecContext.keyint_min` value.
func (avctx *AVCodecContext) SetKeyintMin(v int32) {
	avctx.keyint_min = (C.int)(v)
}

// GetKeyintMinAddr gets `AVCodecContext.keyint_min` address.
func (avctx *AVCodecContext) GetKeyintMinAddr() *int32 {
	return (*int32)(&avctx.keyint_min)
}

// GetRefs gets `AVCodecContext.refs` value.
func (avctx *AVCodecContext) GetRefs() int32 {
	return (int32)(avctx.refs)
}

// SetRefs sets `AVCodecContext.refs` value.
func (avctx *AVCodecContext) SetRefs(v int32) {
	avctx.refs = (C.int)(v)
}

// GetRefsAddr gets `AVCodecContext.refs` address.
func (avctx *AVCodecContext) GetRefsAddr() *int32 {
	return (*int32)(&avctx.refs)
}

// Deprecated: Use encoder private options instead.
//
// GetChromaoffset gets `AVCodecContext.chromaoffset` value.
func (avctx *AVCodecContext) GetChromaoffset() int32 {
	return (int32)(avctx.chromaoffset)
}

// Deprecated: Use encoder private options instead.
//
// SetChromaoffset sets `AVCodecContext.chromaoffset` value.
func (avctx *AVCodecContext) SetChromaoffset(v int32) {
	avctx.chromaoffset = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetChromaoffsetAddr gets `AVCodecContext.chromaoffset` address.
func (avctx *AVCodecContext) GetChromaoffsetAddr() *int32 {
	return (*int32)(&avctx.chromaoffset)
}

// GetMv0Threshold gets `AVCodecContext.mv0_threshold` value.
func (avctx *AVCodecContext) GetMv0Threshold() int32 {
	return (int32)(avctx.mv0_threshold)
}

// SetMv0Threshold sets `AVCodecContext.mv0_threshold` value.
func (avctx *AVCodecContext) SetMv0Threshold(v int32) {
	avctx.mv0_threshold = (C.int)(v)
}

// GetMv0ThresholdAddr gets `AVCodecContext.mv0_threshold` address.
func (avctx *AVCodecContext) GetMv0ThresholdAddr() *int32 {
	return (*int32)(&avctx.mv0_threshold)
}

// Deprecated: Use encoder private options instead.
//
// GetBSensitivity gets `AVCodecContext.b_sensitivity` value.
func (avctx *AVCodecContext) GetBSensitivity() int32 {
	return (int32)(avctx.b_sensitivity)
}

// Deprecated: Use encoder private options instead.
//
// SetBSensitivity sets `AVCodecContext.b_sensitivity` value.
func (avctx *AVCodecContext) SetBSensitivity(v int32) {
	avctx.b_sensitivity = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetBSensitivityAddr gets `AVCodecContext.b_sensitivity` address.
func (avctx *AVCodecContext) GetBSensitivityAddr() *int32 {
	return (*int32)(&avctx.b_sensitivity)
}

// GetColorPrimaries gets `AVCodecContext.color_primaries` value.
func (avctx *AVCodecContext) GetColorPrimaries() AVColorPrimaries {
	return (AVColorPrimaries)(avctx.color_primaries)
}

// SetColorPrimaries sets `AVCodecContext.color_primaries` value.
func (avctx *AVCodecContext) SetColorPrimaries(v AVColorPrimaries) {
	avctx.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// GetColorPrimariesAddr gets `AVCodecContext.color_primaries` address.
func (avctx *AVCodecContext) GetColorPrimariesAddr() *AVColorPrimaries {
	return (*AVColorPrimaries)(unsafe.Pointer(&avctx.color_primaries))
}

// GetColorTrc gets `AVCodecContext.color_trc` value.
func (avctx *AVCodecContext) GetColorTrc() AVColorTransferCharacteristic {
	return (AVColorTransferCharacteristic)(avctx.color_trc)
}

// SetColorTrc sets `AVCodecContext.color_trc` value.
func (avctx *AVCodecContext) SetColorTrc(v AVColorTransferCharacteristic) {
	avctx.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// GetColorTrcAddr gets `AVCodecContext.color_trc` address.
func (avctx *AVCodecContext) GetColorTrcAddr() *AVColorTransferCharacteristic {
	return (*AVColorTransferCharacteristic)(unsafe.Pointer(&avctx.color_trc))
}

// GetColorspace gets `AVCodecContext.colorspace` value.
func (avctx *AVCodecContext) GetColorspace() AVColorSpace {
	return (AVColorSpace)(avctx.colorspace)
}

// SetColorspace sets `AVCodecContext.colorspace` value.
func (avctx *AVCodecContext) SetColorspace(v AVColorSpace) {
	avctx.colorspace = (C.enum_AVColorSpace)(v)
}

// GetColorspaceAddr gets `AVCodecContext.colorspace` address.
func (avctx *AVCodecContext) GetColorspaceAddr() *AVColorSpace {
	return (*AVColorSpace)(unsafe.Pointer(&avctx.colorspace))
}

// GetColorRange gets `AVCodecContext.colorrange` value.
func (avctx *AVCodecContext) GetColorRange() AVColorRange {
	return (AVColorRange)(avctx.color_range)
}

// SetColorRange sets `AVCodecContext.colorrange` value.
func (avctx *AVCodecContext) SetColorRange(v AVColorRange) {
	avctx.color_range = (C.enum_AVColorRange)(v)
}

// GetColorRangeAddr gets `AVCodecContext.colorrange` address.
func (avctx *AVCodecContext) GetColorRangeAddr() *AVColorRange {
	return (*AVColorRange)(unsafe.Pointer(&avctx.color_range))
}

// GetChromaSampleLocation gets `AVCodecContext.chroma_sample_location` value.
func (avctx *AVCodecContext) GetChromaSampleLocation() AVChromaLocation {
	return (AVChromaLocation)(avctx.chroma_sample_location)
}

// SetChromaSampleLocation sets `AVCodecContext.chroma_sample_location` value.
func (avctx *AVCodecContext) SetChromaSampleLocation(v AVChromaLocation) {
	avctx.chroma_sample_location = (C.enum_AVChromaLocation)(v)
}

// GetChromaSampleLocationAddr gets `AVCodecContext.chroma_sample_location` address.
func (avctx *AVCodecContext) GetChromaSampleLocationAddr() *AVChromaLocation {
	return (*AVChromaLocation)(unsafe.Pointer(&avctx.chroma_sample_location))
}

// GetSlices gets `AVCodecContext.slices` value.
func (avctx *AVCodecContext) GetSlices() int32 {
	return (int32)(avctx.slices)
}

// SetSlices sets `AVCodecContext.slices` value.
func (avctx *AVCodecContext) SetSlices(v int32) {
	avctx.slices = (C.int)(v)
}

// GetSlicesAddr gets `AVCodecContext.slices` address.
func (avctx *AVCodecContext) GetSlicesAddr() *int32 {
	return (*int32)(&avctx.slices)
}

// GetFieldOrder gets `AVCodecContext.field_order` value.
func (avctx *AVCodecContext) GetFieldOrder() AVFieldOrder {
	return (AVFieldOrder)(avctx.field_order)
}

// SetFieldOrder sets `AVCodecContext.field_order` value.
func (avctx *AVCodecContext) SetFieldOrder(v AVFieldOrder) {
	avctx.field_order = (C.enum_AVFieldOrder)(v)
}

// GetFieldOrderAddr gets `AVCodecContext.field_order` address.
func (avctx *AVCodecContext) GetFieldOrderAddr() *AVFieldOrder {
	return (*AVFieldOrder)(unsafe.Pointer(&avctx.field_order))
}

// GetSampleRate gets `AVCodecContext.sample_rate` value.
func (avctx *AVCodecContext) GetSampleRate() int32 {
	return (int32)(avctx.sample_rate)
}

// SetSampleRate sets `AVCodecContext.sample_rate` value.
func (avctx *AVCodecContext) SetSampleRate(v int32) {
	avctx.sample_rate = (C.int)(v)
}

// GetSampleRateAddr gets `AVCodecContext.sample_rate` address.
func (avctx *AVCodecContext) GetSampleRateAddr() *int32 {
	return (*int32)(&avctx.sample_rate)
}

// GetChannels gets `AVCodecContext.channels` value.
func (avctx *AVCodecContext) GetChannels() int32 {
	return (int32)(avctx.channels)
}

// SetChannels sets `AVCodecContext.channels` value.
func (avctx *AVCodecContext) SetChannels(v int32) {
	avctx.channels = (C.int)(v)
}

// GetChannelsAddr gets `AVCodecContext.channels` address.
func (avctx *AVCodecContext) GetChannelsAddr() *int32 {
	return (*int32)(&avctx.channels)
}

// GetSampleFmt gets `AVCodecContext.sample_fmt` value.
func (avctx *AVCodecContext) GetSampleFmt() AVSampleFormat {
	return (AVSampleFormat)(avctx.sample_fmt)
}

// SetSampleFmt sets `AVCodecContext.sample_fmt` value.
func (avctx *AVCodecContext) SetSampleFmt(v AVSampleFormat) {
	avctx.sample_fmt = (C.enum_AVSampleFormat)(v)
}

// GetSampleFmtAddr gets `AVCodecContext.sample_fmt` address.
func (avctx *AVCodecContext) GetSampleFmtAddr() *AVSampleFormat {
	return (*AVSampleFormat)(&avctx.sample_fmt)
}

// GetFrameSize gets `AVCodecContext.frame_size` value.
func (avctx *AVCodecContext) GetFrameSize() int32 {
	return (int32)(avctx.frame_size)
}

// SetFrameSize sets `AVCodecContext.frame_size` value.
func (avctx *AVCodecContext) SetFrameSize(v int32) {
	avctx.frame_size = (C.int)(v)
}

// GetFrameSizeAddr gets `AVCodecContext.frame_size` address.
func (avctx *AVCodecContext) GetFrameSizeAddr() *int32 {
	return (*int32)(&avctx.frame_size)
}

// GetFrameNumber gets `AVCodecContext.frame_number` value.
func (avctx *AVCodecContext) GetFrameNumber() int32 {
	return (int32)(avctx.frame_number)
}

// SetFrameNumber sets `AVCodecContext.frame_number` value.
func (avctx *AVCodecContext) SetFrameNumber(v int32) {
	avctx.frame_number = (C.int)(v)
}

// GetFrameNumberAddr gets `AVCodecContext.frame_number` address.
func (avctx *AVCodecContext) GetFrameNumberAddr() *int32 {
	return (*int32)(&avctx.frame_number)
}

// GetBlockAlign gets `AVCodecContext.block_align` value.
func (avctx *AVCodecContext) GetBlockAlign() int32 {
	return (int32)(avctx.block_align)
}

// SetBlockAlign sets `AVCodecContext.block_align` value.
func (avctx *AVCodecContext) SetBlockAlign(v int32) {
	avctx.block_align = (C.int)(v)
}

// GetBlockAlignAddr gets `AVCodecContext.block_align` address.
func (avctx *AVCodecContext) GetBlockAlignAddr() *int32 {
	return (*int32)(&avctx.block_align)
}

// GetCutoff gets `AVCodecContext.cutoff` value.
func (avctx *AVCodecContext) GetCutoff() int32 {
	return (int32)(avctx.cutoff)
}

// SetCutoff sets `AVCodecContext.cutoff` value.
func (avctx *AVCodecContext) SetCutoff(v int32) {
	avctx.cutoff = (C.int)(v)
}

// GetCutoffAddr gets `AVCodecContext.cutoff` address.
func (avctx *AVCodecContext) GetCutoffAddr() *int32 {
	return (*int32)(&avctx.cutoff)
}

// GetChannelLayout gets `AVCodecContext.channel_layout` value.
func (avctx *AVCodecContext) GetChannelLayout() uint64 {
	return (uint64)(avctx.channel_layout)
}

// SetChannelLayout sets `AVCodecContext.channel_layout` value.
func (avctx *AVCodecContext) SetChannelLayout(v uint64) {
	avctx.channel_layout = (C.uint64_t)(v)
}

// GetChannelLayoutAddr gets `AVCodecContext.channel_layout` address.
func (avctx *AVCodecContext) GetChannelLayoutAddr() *uint64 {
	return (*uint64)(&avctx.channel_layout)
}

// GetRequestChannelLayout gets `AVCodecContext.request_channel_layout` value.
func (avctx *AVCodecContext) GetRequestChannelLayout() uint64 {
	return (uint64)(avctx.request_channel_layout)
}

// SetRequestChannelLayout sets `AVCodecContext.request_channel_layout` value.
func (avctx *AVCodecContext) SetRequestChannelLayout(v uint64) {
	avctx.request_channel_layout = (C.uint64_t)(v)
}

// GetRequestChannelLayoutAddr gets `AVCodecContext.request_channel_layout` address.
func (avctx *AVCodecContext) GetRequestChannelLayoutAddr() *uint64 {
	return (*uint64)(&avctx.request_channel_layout)
}

// GetAudioServiceType gets `AVCodecContext.audio_servicetype` value.
func (avctx *AVCodecContext) GetAudioServiceType() AVAudioServiceType {
	return (AVAudioServiceType)(avctx.audio_service_type)
}

// SetAudioServiceType sets `AVCodecContext.audio_servicetype` value.
func (avctx *AVCodecContext) SetAudioServiceType(v AVAudioServiceType) {
	avctx.audio_service_type = (C.enum_AVAudioServiceType)(v)
}

// GetAudioServiceTypeAddr gets `AVCodecContext.audio_servicetype` address.
func (avctx *AVCodecContext) GetAudioServiceTypeAddr() *AVAudioServiceType {
	return (*AVAudioServiceType)(unsafe.Pointer(&avctx.audio_service_type))
}

// GetRequestSampleFmt gets `AVCodecContext.request_sample_fmt` value.
func (avctx *AVCodecContext) GetRequestSampleFmt() AVSampleFormat {
	return (AVSampleFormat)(avctx.request_sample_fmt)
}

// SetRequestSampleFmt sets `AVCodecContext.request_sample_fmt` value.
func (avctx *AVCodecContext) SetRequestSampleFmt(v AVSampleFormat) {
	avctx.request_sample_fmt = (C.enum_AVSampleFormat)(v)
}

// GetRequestSampleFmtAddr gets `AVCodecContext.request_sample_fmt` address.
func (avctx *AVCodecContext) GetRequestSampleFmtAddr() *AVSampleFormat {
	return (*AVSampleFormat)(&avctx.request_sample_fmt)
}

// typedef int (*avcodec_context_get_buffer2_func)(struct AVCodecContext *s, AVFrame *frame, int flags);
type AVCodecContextGetBuffer2Func = C.avcodec_context_get_buffer2_func

// GetGetBuffer2 gets `AVCodecContext.get_buffer2` value.
func (avctx *AVCodecContext) GetGetBuffer2() AVCodecContextGetBuffer2Func {
	return (AVCodecContextGetBuffer2Func)(avctx.get_buffer2)
}

// SetGetBuffer2 sets `AVCodecContext.get_buffer2` value.
func (avctx *AVCodecContext) SetGetBuffer2(v AVCodecContextGetBuffer2Func) {
	avctx.get_buffer2 = (C.avcodec_context_get_buffer2_func)(v)
}

// GetGetBuffer2Addr gets `AVCodecContext.get_buffer2` address.
func (avctx *AVCodecContext) GetGetBuffer2Addr() *AVCodecContextGetBuffer2Func {
	return (*AVCodecContextGetBuffer2Func)(&avctx.get_buffer2)
}

// Deprecated: Use encoder private options instead.
//
// GetRefcountedFrames gets `AVCodecContext.refcounted_frames` value.
func (avctx *AVCodecContext) GetRefcountedFrames() int32 {
	return (int32)(avctx.refcounted_frames)
}

// Deprecated: Use encoder private options instead.
//
// SetRefcountedFrames sets `AVCodecContext.refcounted_frames` value.
func (avctx *AVCodecContext) SetRefcountedFrames(v int32) {
	avctx.refcounted_frames = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetRefcountedFramesAddr gets `AVCodecContext.refcounted_frames` address.
func (avctx *AVCodecContext) GetRefcountedFramesAddr() *int32 {
	return (*int32)(&avctx.refcounted_frames)
}

// GetQcompress gets `AVCodecContext.qcompress` value.
func (avctx *AVCodecContext) GetQcompress() float32 {
	return (float32)(avctx.qcompress)
}

// SetQcompress sets `AVCodecContext.qcompress` value.
func (avctx *AVCodecContext) SetQcompress(v float32) {
	avctx.qcompress = (C.float)(v)
}

// GetQcompressAddr gets `AVCodecContext.qcompress` address.
func (avctx *AVCodecContext) GetQcompressAddr() *float32 {
	return (*float32)(&avctx.qcompress)
}

// GetQblur gets `AVCodecContext.qblur` value.
func (avctx *AVCodecContext) GetQblur() float32 {
	return (float32)(avctx.qblur)
}

// SetQblur sets `AVCodecContext.qblur` value.
func (avctx *AVCodecContext) SetQblur(v float32) {
	avctx.qblur = (C.float)(v)
}

// GetQblurAddr gets `AVCodecContext.qblur` address.
func (avctx *AVCodecContext) GetQblurAddr() *float32 {
	return (*float32)(&avctx.qblur)
}

// GetQmin gets `AVCodecContext.qmin` value.
func (avctx *AVCodecContext) GetQmin() int32 {
	return (int32)(avctx.qmin)
}

// SetQmin sets `AVCodecContext.qmin` value.
func (avctx *AVCodecContext) SetQmin(v int32) {
	avctx.qmin = (C.int)(v)
}

// GetQminAddr gets `AVCodecContext.qmin` address.
func (avctx *AVCodecContext) GetQminAddr() *int32 {
	return (*int32)(&avctx.qmin)
}

// GetQmax gets `AVCodecContext.qmax` value.
func (avctx *AVCodecContext) GetQmax() int32 {
	return (int32)(avctx.qmax)
}

// SetQmax sets `AVCodecContext.qmax` value.
func (avctx *AVCodecContext) SetQmax(v int32) {
	avctx.qmax = (C.int)(v)
}

// GetQmaxAddr gets `AVCodecContext.qmax` address.
func (avctx *AVCodecContext) GetQmaxAddr() *int32 {
	return (*int32)(&avctx.qmax)
}

// GetMaxQdiff gets `AVCodecContext.max_qdiff` value.
func (avctx *AVCodecContext) GetMaxQdiff() int32 {
	return (int32)(avctx.max_qdiff)
}

// SetMaxQdiff sets `AVCodecContext.max_qdiff` value.
func (avctx *AVCodecContext) SetMaxQdiff(v int32) {
	avctx.max_qdiff = (C.int)(v)
}

// GetMaxQdiffAddr gets `AVCodecContext.max_qdiff` address.
func (avctx *AVCodecContext) GetMaxQdiffAddr() *int32 {
	return (*int32)(&avctx.max_qdiff)
}

// GetRcBufferSize gets `AVCodecContext.rc_buffer_size` value.
func (avctx *AVCodecContext) GetRcBufferSize() int32 {
	return (int32)(avctx.rc_buffer_size)
}

// SetRcBufferSize sets `AVCodecContext.rc_buffer_size` value.
func (avctx *AVCodecContext) SetRcBufferSize(v int32) {
	avctx.rc_buffer_size = (C.int)(v)
}

// GetRcBufferSizeAddr gets `AVCodecContext.rc_buffer_size` address.
func (avctx *AVCodecContext) GetRcBufferSizeAddr() *int32 {
	return (*int32)(&avctx.rc_buffer_size)
}

// GetRcOverrideCount gets `AVCodecContext.rc_override_count` value.
func (avctx *AVCodecContext) GetRcOverrideCount() int32 {
	return (int32)(avctx.rc_override_count)
}

// SetRcOverrideCount sets `AVCodecContext.rc_override_count` value.
func (avctx *AVCodecContext) SetRcOverrideCount(v int32) {
	avctx.rc_override_count = (C.int)(v)
}

// GetRcOverrideCountAddr gets `AVCodecContext.rc_override_count` address.
func (avctx *AVCodecContext) GetRcOverrideCountAddr() *int32 {
	return (*int32)(&avctx.rc_override_count)
}

// GetRcOverride gets `AVCodecContext.rc_override` value.
func (avctx *AVCodecContext) GetRcOverride() *RcOverride {
	return (*RcOverride)(avctx.rc_override)
}

// SetRcOverride sets `AVCodecContext.rc_override` value.
func (avctx *AVCodecContext) SetRcOverride(v *RcOverride) {
	avctx.rc_override = (*C.RcOverride)(v)
}

// GetRcOverrideAddr gets `AVCodecContext.rc_override` address.
func (avctx *AVCodecContext) GetRcOverrideAddr() **RcOverride {
	return (**RcOverride)(unsafe.Pointer(&avctx.rc_override))
}

// GetRcMaxRate gets `AVCodecContext.rc_max_rate` value.
func (avctx *AVCodecContext) GetRcMaxRate() int64 {
	return (int64)(avctx.rc_max_rate)
}

// SetRcMaxRate sets `AVCodecContext.rc_max_rate` value.
func (avctx *AVCodecContext) SetRcMaxRate(v int64) {
	avctx.rc_max_rate = (C.int64_t)(v)
}

// GetRcMaxRateAddr gets `AVCodecContext.rc_max_rate` address.
func (avctx *AVCodecContext) GetRcMaxRateAddr() *int64 {
	return (*int64)(&avctx.rc_max_rate)
}

// GetRcMinRate gets `AVCodecContext.rc_min_rate` value.
func (avctx *AVCodecContext) GetRcMinRate() int64 {
	return (int64)(avctx.rc_min_rate)
}

// SetRcMinRate sets `AVCodecContext.rc_min_rate` value.
func (avctx *AVCodecContext) SetRcMinRate(v int64) {
	avctx.rc_min_rate = (C.int64_t)(v)
}

// GetRcMinRateAddr gets `AVCodecContext.rc_min_rate` address.
func (avctx *AVCodecContext) GetRcMinRateAddr() *int64 {
	return (*int64)(&avctx.rc_min_rate)
}

// GetRcMaxAvailableVbvUse gets `AVCodecContext.rc_max_available_vbv_use` value.
func (avctx *AVCodecContext) GetRcMaxAvailableVbvUse() float32 {
	return (float32)(avctx.rc_max_available_vbv_use)
}

// SetRcMaxAvailableVbvUse sets `AVCodecContext.rc_max_available_vbv_use` value.
func (avctx *AVCodecContext) SetRcMaxAvailableVbvUse(v float32) {
	avctx.rc_max_available_vbv_use = (C.float)(v)
}

// GetRcMaxAvailableVbvUseAddr gets `AVCodecContext.rc_max_available_vbv_use` address.
func (avctx *AVCodecContext) GetRcMaxAvailableVbvUseAddr() *float32 {
	return (*float32)(&avctx.rc_max_available_vbv_use)
}

// GetRcMinVbvOverflowUse gets `AVCodecContext.rc_min_vbv_overflow_use` value.
func (avctx *AVCodecContext) GetRcMinVbvOverflowUse() float32 {
	return (float32)(avctx.rc_min_vbv_overflow_use)
}

// SetRcMinVbvOverflowUse sets `AVCodecContext.rc_min_vbv_overflow_use` value.
func (avctx *AVCodecContext) SetRcMinVbvOverflowUse(v float32) {
	avctx.rc_min_vbv_overflow_use = (C.float)(v)
}

// GetRcMinVbvOverflowUseAddr gets `AVCodecContext.rc_min_vbv_overflow_use` address.
func (avctx *AVCodecContext) GetRcMinVbvOverflowUseAddr() *float32 {
	return (*float32)(&avctx.rc_min_vbv_overflow_use)
}

// GetRcInitialBufferOccupancy gets `AVCodecContext.rc_initial_buffer_occupancy` value.
func (avctx *AVCodecContext) GetRcInitialBufferOccupancy() int32 {
	return (int32)(avctx.rc_initial_buffer_occupancy)
}

// SetRcInitialBufferOccupancy sets `AVCodecContext.rc_initial_buffer_occupancy` value.
func (avctx *AVCodecContext) SetRcInitialBufferOccupancy(v int32) {
	avctx.rc_initial_buffer_occupancy = (C.int)(v)
}

// GetRcInitialBufferOccupancyAddr gets `AVCodecContext.rc_initial_buffer_occupancy` address.
func (avctx *AVCodecContext) GetRcInitialBufferOccupancyAddr() *int32 {
	return (*int32)(&avctx.rc_initial_buffer_occupancy)
}

const (
	FF_CODER_TYPE_VLC = int32(C.FF_CODER_TYPE_VLC)
	FF_CODER_TYPE_AC  = int32(C.FF_CODER_TYPE_AC)
	FF_CODER_TYPE_RAW = int32(C.FF_CODER_TYPE_RAW)
	FF_CODER_TYPE_RLE = int32(C.FF_CODER_TYPE_RLE)
)

// Deprecated: Use encoder private options instead.
//
// GetCoderType gets `AVCodecContext.codertype` value.
func (avctx *AVCodecContext) GetCoderType() int32 {
	return (int32)(avctx.coder_type)
}

// Deprecated: Use encoder private options instead.
//
// SetCoderType sets `AVCodecContext.codertype` value.
func (avctx *AVCodecContext) SetCoderType(v int32) {
	avctx.coder_type = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetCoderTypeAddr gets `AVCodecContext.codertype` address.
func (avctx *AVCodecContext) GetCoderTypeAddr() *int32 {
	return (*int32)(&avctx.coder_type)
}

// Deprecated: Use encoder private options instead.
//
// GetContextModel gets `AVCodecContext.context_model` value.
func (avctx *AVCodecContext) GetContextModel() int32 {
	return (int32)(avctx.context_model)
}

// Deprecated: Use encoder private options instead.
//
// SetContextModel sets `AVCodecContext.context_model` value.
func (avctx *AVCodecContext) SetContextModel(v int32) {
	avctx.context_model = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetContextModelAddr gets `AVCodecContext.context_model` address.
func (avctx *AVCodecContext) GetContextModelAddr() *int32 {
	return (*int32)(&avctx.context_model)
}

// Deprecated: Use encoder private options instead.
//
// GetFrameSkipThreshold gets `AVCodecContext.frame_skip_threshold` value.
func (avctx *AVCodecContext) GetFrameSkipThreshold() int32 {
	return (int32)(avctx.frame_skip_threshold)
}

// Deprecated: Use encoder private options instead.
//
// SetFrameSkipThreshold sets `AVCodecContext.frame_skip_threshold` value.
func (avctx *AVCodecContext) SetFrameSkipThreshold(v int32) {
	avctx.frame_skip_threshold = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetFrameSkipThresholdAddr gets `AVCodecContext.frame_skip_threshold` address.
func (avctx *AVCodecContext) GetFrameSkipThresholdAddr() *int32 {
	return (*int32)(&avctx.frame_skip_threshold)
}

// Deprecated: Use encoder private options instead.
//
// GetFrameSkipFactor gets `AVCodecContext.frame_skip_factor` value.
func (avctx *AVCodecContext) GetFrameSkipFactor() int32 {
	return (int32)(avctx.frame_skip_factor)
}

// Deprecated: Use encoder private options instead.
//
// SetFrameSkipFactor sets `AVCodecContext.frame_skip_factor` value.
func (avctx *AVCodecContext) SetFrameSkipFactor(v int32) {
	avctx.frame_skip_factor = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetFrameSkipFactorAddr gets `AVCodecContext.frame_skip_factor` address.
func (avctx *AVCodecContext) GetFrameSkipFactorAddr() *int32 {
	return (*int32)(&avctx.frame_skip_factor)
}

// Deprecated: Use encoder private options instead.
//
// GetFrameSkipExp gets `AVCodecContext.frame_skip_exp` value.
func (avctx *AVCodecContext) GetFrameSkipExp() int32 {
	return (int32)(avctx.frame_skip_exp)
}

// Deprecated: Use encoder private options instead.
//
// SetFrameSkipExp sets `AVCodecContext.frame_skip_exp` value.
func (avctx *AVCodecContext) SetFrameSkipExp(v int32) {
	avctx.frame_skip_exp = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetFrameSkipExpAddr gets `AVCodecContext.frame_skip_exp` address.
func (avctx *AVCodecContext) GetFrameSkipExpAddr() *int32 {
	return (*int32)(&avctx.frame_skip_exp)
}

// Deprecated: Use encoder private options instead.
//
// GetFrameSkipCmp gets `AVCodecContext.frame_skip_cmp` value.
func (avctx *AVCodecContext) GetFrameSkipCmp() int32 {
	return (int32)(avctx.frame_skip_cmp)
}

// Deprecated: Use encoder private options instead.
//
// SetFrameSkipCmp sets `AVCodecContext.frame_skip_cmp` value.
func (avctx *AVCodecContext) SetFrameSkipCmp(v int32) {
	avctx.frame_skip_cmp = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetFrameSkipCmpAddr gets `AVCodecContext.frame_skip_cmp` address.
func (avctx *AVCodecContext) GetFrameSkipCmpAddr() *int32 {
	return (*int32)(&avctx.frame_skip_cmp)
}

// GetTrellis gets `AVCodecContext.trellis` value.
func (avctx *AVCodecContext) GetTrellis() int32 {
	return (int32)(avctx.trellis)
}

// SetTrellis sets `AVCodecContext.trellis` value.
func (avctx *AVCodecContext) SetTrellis(v int32) {
	avctx.trellis = (C.int)(v)
}

// GetTrellisAddr gets `AVCodecContext.trellis` address.
func (avctx *AVCodecContext) GetTrellisAddr() *int32 {
	return (*int32)(&avctx.trellis)
}

// Deprecated: Use encoder private options instead.
//
// GetMinPredictionOrder gets `AVCodecContext.min_prediction_order` value.
func (avctx *AVCodecContext) GetMinPredictionOrder() int32 {
	return (int32)(avctx.min_prediction_order)
}

// Deprecated: Use encoder private options instead.
//
// SetMinPredictionOrder sets `AVCodecContext.min_prediction_order` value.
func (avctx *AVCodecContext) SetMinPredictionOrder(v int32) {
	avctx.min_prediction_order = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetMinPredictionOrderAddr gets `AVCodecContext.min_prediction_order` address.
func (avctx *AVCodecContext) GetMinPredictionOrderAddr() *int32 {
	return (*int32)(&avctx.min_prediction_order)
}

// Deprecated: Use encoder private options instead.
//
// GetMaxPredictionOrder gets `AVCodecContext.max_prediction_order` value.
func (avctx *AVCodecContext) GetMaxPredictionOrder() int32 {
	return (int32)(avctx.max_prediction_order)
}

// Deprecated: Use encoder private options instead.
//
// SetMaxPredictionOrder sets `AVCodecContext.max_prediction_order` value.
func (avctx *AVCodecContext) SetMaxPredictionOrder(v int32) {
	avctx.max_prediction_order = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetMaxPredictionOrderAddr gets `AVCodecContext.max_prediction_order` address.
func (avctx *AVCodecContext) GetMaxPredictionOrderAddr() *int32 {
	return (*int32)(&avctx.max_prediction_order)
}

// Deprecated: Use encoder private options instead.
//
// GetTimecodeFrameStart gets `AVCodecContext.timecode_frame_start` value.
func (avctx *AVCodecContext) GetTimecodeFrameStart() int64 {
	return (int64)(avctx.timecode_frame_start)
}

// Deprecated: Use encoder private options instead.
//
// SetTimecodeFrameStart sets `AVCodecContext.timecode_frame_start` value.
func (avctx *AVCodecContext) SetTimecodeFrameStart(v int64) {
	avctx.timecode_frame_start = (C.int64_t)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetTimecodeFrameStartAddr gets `AVCodecContext.timecode_frame_start` address.
func (avctx *AVCodecContext) GetTimecodeFrameStartAddr() *int64 {
	return (*int64)(&avctx.timecode_frame_start)
}

// Deprecated: Unused.
//
// typedef void (*avcodec_context_rtp_callback_func)(struct AVCodecContext *avctx, void *data, int size, int mb_nb);
type AvCodecContextRtpCallbackFunc = C.avcodec_context_rtp_callback_func

// Deprecated: Unused.
//
// GetRtpCallback gets `AVCodecContext.rtp_callback` value.
func (avctx *AVCodecContext) GetRtpCallback() AvCodecContextRtpCallbackFunc {
	return (AvCodecContextRtpCallbackFunc)(avctx.rtp_callback)
}

// Deprecated: Unused.
//
// SetRtpCallback sets `AVCodecContext.rtp_callback` value.
func (avctx *AVCodecContext) SetRtpCallback(v AvCodecContextRtpCallbackFunc) {
	avctx.rtp_callback = (C.avcodec_context_rtp_callback_func)(v)
}

// Deprecated: Unused.
//
// GetRtpCallbackAddr gets `AVCodecContext.rtp_callback` address.
func (avctx *AVCodecContext) GetRtpCallbackAddr() *AvCodecContextRtpCallbackFunc {
	return (*AvCodecContextRtpCallbackFunc)(&avctx.rtp_callback)
}

// Deprecated: Use encoder private options instead.
//
// GetRtpPayloadSize gets `AVCodecContext.rtp_payload_size` value.
func (avctx *AVCodecContext) GetRtpPayloadSize() int32 {
	return (int32)(avctx.rtp_payload_size)
}

// Deprecated: Use encoder private options instead.
//
// SetRtpPayloadSize sets `AVCodecContext.rtp_payload_size` value.
func (avctx *AVCodecContext) SetRtpPayloadSize(v int32) {
	avctx.rtp_payload_size = (C.int)(v)
}

// Deprecated: Use encoder private options instead.
//
// GetRtpPayloadSizeAddr gets `AVCodecContext.rtp_payload_size` address.
func (avctx *AVCodecContext) GetRtpPayloadSizeAddr() *int32 {
	return (*int32)(&avctx.rtp_payload_size)
}

// Deprecated: No use.
//
// GetMvBits gets `AVCodecContext.mv_bits` value.
func (avctx *AVCodecContext) GetMvBits() int32 {
	return (int32)(avctx.mv_bits)
}

// Deprecated: No use.
//
// SetMvBits sets `AVCodecContext.mv_bits` value.
func (avctx *AVCodecContext) SetMvBits(v int32) {
	avctx.mv_bits = (C.int)(v)
}

// Deprecated: No use.
//
// GetMvBitsAddr gets `AVCodecContext.mv_bits` address.
func (avctx *AVCodecContext) GetMvBitsAddr() *int32 {
	return (*int32)(&avctx.mv_bits)
}

// Deprecated: No use.
//
// GetHeaderBits gets `AVCodecContext.header_bits` value.
func (avctx *AVCodecContext) GetHeaderBits() int32 {
	return (int32)(avctx.header_bits)
}

// Deprecated: No use.
//
// SetHeaderBits sets `AVCodecContext.header_bits` value.
func (avctx *AVCodecContext) SetHeaderBits(v int32) {
	avctx.header_bits = (C.int)(v)
}

// Deprecated: No use.
//
// GetHeaderBitsAddr gets `AVCodecContext.header_bits` address.
func (avctx *AVCodecContext) GetHeaderBitsAddr() *int32 {
	return (*int32)(&avctx.header_bits)
}

// Deprecated: No use.
//
// GetITexBits gets `AVCodecContext.i_tex_bits` value.
func (avctx *AVCodecContext) GetITexBits() int32 {
	return (int32)(avctx.i_tex_bits)
}

// Deprecated: No use.
//
// SetITexBits sets `AVCodecContext.i_tex_bits` value.
func (avctx *AVCodecContext) SetITexBits(v int32) {
	avctx.i_tex_bits = (C.int)(v)
}

// Deprecated: No use.
//
// GetITexBitsAddr gets `AVCodecContext.i_tex_bits` address.
func (avctx *AVCodecContext) GetITexBitsAddr() *int32 {
	return (*int32)(&avctx.i_tex_bits)
}

// Deprecated: No use.
//
// GetPTexBits gets `AVCodecContext.p_tex_bits` value.
func (avctx *AVCodecContext) GetPTexBits() int32 {
	return (int32)(avctx.p_tex_bits)
}

// Deprecated: No use.
//
// SetPTexBits sets `AVCodecContext.p_tex_bits` value.
func (avctx *AVCodecContext) SetPTexBits(v int32) {
	avctx.p_tex_bits = (C.int)(v)
}

// Deprecated: No use.
//
// GetPTexBitsAddr gets `AVCodecContext.p_tex_bits` address.
func (avctx *AVCodecContext) GetPTexBitsAddr() *int32 {
	return (*int32)(&avctx.p_tex_bits)
}

// Deprecated: No use.
//
// GetICount gets `AVCodecContext.i_count` value.
func (avctx *AVCodecContext) GetICount() int32 {
	return (int32)(avctx.i_count)
}

// Deprecated: No use.
//
// SetICount sets `AVCodecContext.i_count` value.
func (avctx *AVCodecContext) SetICount(v int32) {
	avctx.i_count = (C.int)(v)
}

// Deprecated: No use.
//
// GetICountAddr gets `AVCodecContext.i_count` address.
func (avctx *AVCodecContext) GetICountAddr() *int32 {
	return (*int32)(&avctx.i_count)
}

// Deprecated: No use.
//
// GetPCount gets `AVCodecContext.p_count` value.
func (avctx *AVCodecContext) GetPCount() int32 {
	return (int32)(avctx.p_count)
}

// Deprecated: No use.
//
// SetPCount sets `AVCodecContext.p_count` value.
func (avctx *AVCodecContext) SetPCount(v int32) {
	avctx.p_count = (C.int)(v)
}

// Deprecated: No use.
//
// GetPCountAddr gets `AVCodecContext.p_count` address.
func (avctx *AVCodecContext) GetPCountAddr() *int32 {
	return (*int32)(&avctx.p_count)
}

// Deprecated: No use.
//
// GetSkipCount gets `AVCodecContext.skip_count` value.
func (avctx *AVCodecContext) GetSkipCount() int32 {
	return (int32)(avctx.skip_count)
}

// Deprecated: No use.
//
// SetSkipCount sets `AVCodecContext.skip_count` value.
func (avctx *AVCodecContext) SetSkipCount(v int32) {
	avctx.skip_count = (C.int)(v)
}

// Deprecated: No use.
//
// GetSkipCountAddr gets `AVCodecContext.skip_count` address.
func (avctx *AVCodecContext) GetSkipCountAddr() *int32 {
	return (*int32)(&avctx.skip_count)
}

// Deprecated: No use.
//
// GetMiscBits gets `AVCodecContext.misc_bits` value.
func (avctx *AVCodecContext) GetMiscBits() int32 {
	return (int32)(avctx.misc_bits)
}

// Deprecated: No use.
//
// SetMiscBits sets `AVCodecContext.misc_bits` value.
func (avctx *AVCodecContext) SetMiscBits(v int32) {
	avctx.misc_bits = (C.int)(v)
}

// Deprecated: No use.
//
// GetMiscBitsAddr gets `AVCodecContext.misc_bits` address.
func (avctx *AVCodecContext) GetMiscBitsAddr() *int32 {
	return (*int32)(&avctx.misc_bits)
}

// Deprecated: Unused.
//
// GetFrameBits gets `AVCodecContext.frame_bits` value.
func (avctx *AVCodecContext) GetFrameBits() int32 {
	return (int32)(avctx.frame_bits)
}

// Deprecated: Unused.
//
// SetFrameBits sets `AVCodecContext.frame_bits` value.
func (avctx *AVCodecContext) SetFrameBits(v int32) {
	avctx.frame_bits = (C.int)(v)
}

// Deprecated: Unused.
//
// GetFrameBitsAddr gets `AVCodecContext.frame_bits` address.
func (avctx *AVCodecContext) GetFrameBitsAddr() *int32 {
	return (*int32)(&avctx.frame_bits)
}

// GetStatsOut gets `AVCodecContext.stats_out` value.
func (avctx *AVCodecContext) GetStatsOut() string {
	return C.GoString(avctx.stats_out)
}

// GetStatsIn gets `AVCodecContext.stats_in` value.
func (avctx *AVCodecContext) GetStatsIn() string {
	return C.GoString(avctx.stats_in)
}

// GetWorkaroundBugs gets `AVCodecContext.workaround_bugs` value.
func (avctx *AVCodecContext) GetWorkaroundBugs() int32 {
	return (int32)(avctx.workaround_bugs)
}

// SetWorkaroundBugs sets `AVCodecContext.workaround_bugs` value.
func (avctx *AVCodecContext) SetWorkaroundBugs(v int32) {
	avctx.workaround_bugs = (C.int)(v)
}

// GetWorkaroundBugsAddr gets `AVCodecContext.workaround_bugs` address.
func (avctx *AVCodecContext) GetWorkaroundBugsAddr() *int32 {
	return (*int32)(&avctx.workaround_bugs)
}

const (
	FF_BUG_AUTODETECT       = int32(C.FF_BUG_AUTODETECT)
	FF_BUG_XVID_ILACE       = int32(C.FF_BUG_XVID_ILACE)
	FF_BUG_UMP4             = int32(C.FF_BUG_UMP4)
	FF_BUG_NO_PADDING       = int32(C.FF_BUG_NO_PADDING)
	FF_BUG_AMV              = int32(C.FF_BUG_AMV)
	FF_BUG_QPEL_CHROMA      = int32(C.FF_BUG_QPEL_CHROMA)
	FF_BUG_STD_QPEL         = int32(C.FF_BUG_STD_QPEL)
	FF_BUG_QPEL_CHROMA2     = int32(C.FF_BUG_QPEL_CHROMA2)
	FF_BUG_DIRECT_BLOCKSIZE = int32(C.FF_BUG_DIRECT_BLOCKSIZE)
	FF_BUG_EDGE             = int32(C.FF_BUG_EDGE)
	FF_BUG_HPEL_CHROMA      = int32(C.FF_BUG_HPEL_CHROMA)
	FF_BUG_DC_CLIP          = int32(C.FF_BUG_DC_CLIP)
	FF_BUG_MS               = int32(C.FF_BUG_MS)
	FF_BUG_TRUNCATED        = int32(C.FF_BUG_TRUNCATED)
	FF_BUG_IEDGE            = int32(C.FF_BUG_IEDGE)
)

// GetStrictStdCompliance gets `AVCodecContext.strict_std_compliance` value.
func (avctx *AVCodecContext) GetStrictStdCompliance() int32 {
	return (int32)(avctx.strict_std_compliance)
}

// SetStrictStdCompliance sets `AVCodecContext.strict_std_compliance` value.
func (avctx *AVCodecContext) SetStrictStdCompliance(v int32) {
	avctx.strict_std_compliance = (C.int)(v)
}

// GetStrictStdComplianceAddr gets `AVCodecContext.strict_std_compliance` address.
func (avctx *AVCodecContext) GetStrictStdComplianceAddr() *int32 {
	return (*int32)(&avctx.strict_std_compliance)
}

const (
	FF_COMPLIANCE_VERY_STRICT  = int32(C.FF_COMPLIANCE_VERY_STRICT)
	FF_COMPLIANCE_STRICT       = int32(C.FF_COMPLIANCE_STRICT)
	FF_COMPLIANCE_NORMAL       = int32(C.FF_COMPLIANCE_NORMAL)
	FF_COMPLIANCE_UNOFFICIAL   = int32(C.FF_COMPLIANCE_UNOFFICIAL)
	FF_COMPLIANCE_EXPERIMENTAL = int32(C.FF_COMPLIANCE_EXPERIMENTAL)
)

// GetErrorConcealment gets `AVCodecContext.error_concealment` value.
func (avctx *AVCodecContext) GetErrorConcealment() int32 {
	return (int32)(avctx.error_concealment)
}

// SetErrorConcealment sets `AVCodecContext.error_concealment` value.
func (avctx *AVCodecContext) SetErrorConcealment(v int32) {
	avctx.error_concealment = (C.int)(v)
}

// GetErrorConcealmentAddr gets `AVCodecContext.error_concealment` address.
func (avctx *AVCodecContext) GetErrorConcealmentAddr() *int32 {
	return (*int32)(&avctx.error_concealment)
}

const (
	FF_EC_GUESS_MVS   = int32(C.FF_EC_GUESS_MVS)
	FF_EC_DEBLOCK     = int32(C.FF_EC_DEBLOCK)
	FF_EC_FAVOR_INTER = int32(C.FF_EC_FAVOR_INTER)
)

// GetDebug gets `AVCodecContext.debug` value.
func (avctx *AVCodecContext) GetDebug() int32 {
	return (int32)(avctx.debug)
}

// SetDebug sets `AVCodecContext.debug` value.
func (avctx *AVCodecContext) SetDebug(v int32) {
	avctx.debug = (C.int)(v)
}

// GetDebugAddr gets `AVCodecContext.debug` address.
func (avctx *AVCodecContext) GetDebugAddr() *int32 {
	return (*int32)(&avctx.debug)
}

const (
	FF_DEBUG_PICT_INFO = int32(C.FF_DEBUG_PICT_INFO)
	FF_DEBUG_RC        = int32(C.FF_DEBUG_RC)
	FF_DEBUG_BITSTREAM = int32(C.FF_DEBUG_BITSTREAM)
	FF_DEBUG_MB_TYPE   = int32(C.FF_DEBUG_MB_TYPE)
	FF_DEBUG_QP        = int32(C.FF_DEBUG_QP)
	FF_DEBUG_DCT_COEFF = int32(C.FF_DEBUG_DCT_COEFF)
	FF_DEBUG_SKIP      = int32(C.FF_DEBUG_SKIP)
	FF_DEBUG_STARTCODE = int32(C.FF_DEBUG_STARTCODE)
	FF_DEBUG_ER        = int32(C.FF_DEBUG_ER)
	FF_DEBUG_MMCO      = int32(C.FF_DEBUG_MMCO)
	FF_DEBUG_BUGS      = int32(C.FF_DEBUG_BUGS)
	FF_DEBUG_BUFFERS   = int32(C.FF_DEBUG_BUFFERS)
	FF_DEBUG_THREADS   = int32(C.FF_DEBUG_THREADS)
	FF_DEBUG_GREEN_MD  = int32(C.FF_DEBUG_GREEN_MD)
	FF_DEBUG_NOMC      = int32(C.FF_DEBUG_NOMC)
)

// GetErrRecognition gets `AVCodecContext.err_recognition` value.
func (avctx *AVCodecContext) GetErrRecognition() int32 {
	return (int32)(avctx.err_recognition)
}

// SetErrRecognition sets `AVCodecContext.err_recognition` value.
func (avctx *AVCodecContext) SetErrRecognition(v int32) {
	avctx.err_recognition = (C.int)(v)
}

// GetErrRecognitionAddr gets `AVCodecContext.err_recognition` address.
func (avctx *AVCodecContext) GetErrRecognitionAddr() *int32 {
	return (*int32)(&avctx.err_recognition)
}

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

// GetReorderedOpaque gets `AVCodecContext.reordered_opaque` value.
func (avctx *AVCodecContext) GetReorderedOpaque() int64 {
	return (int64)(avctx.reordered_opaque)
}

// SetReorderedOpaque sets `AVCodecContext.reordered_opaque` value.
func (avctx *AVCodecContext) SetReorderedOpaque(v int64) {
	avctx.reordered_opaque = (C.int64_t)(v)
}

// GetReorderedOpaqueAddr gets `AVCodecContext.reordered_opaque` address.
func (avctx *AVCodecContext) GetReorderedOpaqueAddr() *int64 {
	return (*int64)(&avctx.reordered_opaque)
}

// GetHwaccel gets `AVCodecContext.hwaccel` value.
func (avctx *AVCodecContext) GetHwaccel() *AVHWAccel {
	return (*AVHWAccel)(avctx.hwaccel)
}

// SetHwaccel sets `AVCodecContext.hwaccel` value.
func (avctx *AVCodecContext) SetHwaccel(v *AVHWAccel) {
	avctx.hwaccel = (*C.AVHWAccel)(v)
}

// GetHwaccelAddr gets `AVCodecContext.hwaccel` address.
func (avctx *AVCodecContext) GetHwaccelAddr() **AVHWAccel {
	return (**AVHWAccel)(unsafe.Pointer(&avctx.hwaccel))
}

// GetHwaccelContext gets `AVCodecContext.hwaccel_context` value.
func (avctx *AVCodecContext) GetHwaccelContext() unsafe.Pointer {
	return (unsafe.Pointer)(avctx.hwaccel_context)
}

// SetHwaccelContext sets `AVCodecContext.hwaccel_context` value.
func (avctx *AVCodecContext) SetHwaccelContext(v CVoidPointer) {
	avctx.hwaccel_context = VoidPointer(v)
}

// GetHwaccelContextAddr gets `AVCodecContext.hwaccel_context` address.
func (avctx *AVCodecContext) GetHwaccelContextAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&avctx.hwaccel_context)
}

// GetError gets `AVCodecContext.error` value.
func (avctx *AVCodecContext) GetError() []uint64 {
	return unsafe.Slice((*uint64)(&avctx.error[0]), AV_NUM_DATA_POINTERS)
}

// SetError sets `AVCodecContext.error` value.
func (avctx *AVCodecContext) SetError(v []uint64) {
	for i := 0; i < FFMIN(len(v), AV_NUM_DATA_POINTERS); i++ {
		avctx.error[i] = (C.uint64_t)(v[i])
	}
}

// GetErrorAddr gets `AVCodecContext.error` address.
func (avctx *AVCodecContext) GetErrorAddr() **uint64 {
	return (**uint64)(unsafe.Pointer(&avctx.error))
}

// GetDctAlgo gets `AVCodecContext.dct_algo` value.
func (avctx *AVCodecContext) GetDctAlgo() int32 {
	return (int32)(avctx.dct_algo)
}

// SetDctAlgo sets `AVCodecContext.dct_algo` value.
func (avctx *AVCodecContext) SetDctAlgo(v int32) {
	avctx.dct_algo = (C.int)(v)
}

// GetDctAlgoAddr gets `AVCodecContext.dct_algo` address.
func (avctx *AVCodecContext) GetDctAlgoAddr() *int32 {
	return (*int32)(&avctx.dct_algo)
}

const (
	FF_DCT_AUTO    = int32(C.FF_DCT_AUTO)
	FF_DCT_FASTINT = int32(C.FF_DCT_FASTINT)
	FF_DCT_INT     = int32(C.FF_DCT_INT)
	FF_DCT_MMX     = int32(C.FF_DCT_MMX)
	FF_DCT_ALTIVEC = int32(C.FF_DCT_ALTIVEC)
	FF_DCT_FAAN    = int32(C.FF_DCT_FAAN)
)

// GetIdctAlgo gets `AVCodecContext.idct_algo` value.
func (avctx *AVCodecContext) GetIdctAlgo() int32 {
	return (int32)(avctx.idct_algo)
}

// SetIdctAlgo sets `AVCodecContext.idct_algo` value.
func (avctx *AVCodecContext) SetIdctAlgo(v int32) {
	avctx.idct_algo = (C.int)(v)
}

// GetIdctAlgoAddr gets `AVCodecContext.idct_algo` address.
func (avctx *AVCodecContext) GetIdctAlgoAddr() *int32 {
	return (*int32)(&avctx.idct_algo)
}

const (
	FF_IDCT_AUTO          = int32(C.FF_IDCT_AUTO)
	FF_IDCT_INT           = int32(C.FF_IDCT_INT)
	FF_IDCT_SIMPLE        = int32(C.FF_IDCT_SIMPLE)
	FF_IDCT_SIMPLEMMX     = int32(C.FF_IDCT_SIMPLEMMX)
	FF_IDCT_ARM           = int32(C.FF_IDCT_ARM)
	FF_IDCT_ALTIVEC       = int32(C.FF_IDCT_ALTIVEC)
	FF_IDCT_SIMPLEARM     = int32(C.FF_IDCT_SIMPLEARM)
	FF_IDCT_XVID          = int32(C.FF_IDCT_XVID)
	FF_IDCT_SIMPLEARMV5TE = int32(C.FF_IDCT_SIMPLEARMV5TE)
	FF_IDCT_SIMPLEARMV6   = int32(C.FF_IDCT_SIMPLEARMV6)
	FF_IDCT_FAAN          = int32(C.FF_IDCT_FAAN)
	FF_IDCT_SIMPLENEON    = int32(C.FF_IDCT_SIMPLENEON)
	FF_IDCT_NONE          = int32(C.FF_IDCT_NONE)
	FF_IDCT_SIMPLEAUTO    = int32(C.FF_IDCT_SIMPLEAUTO)
)

// GetBitsPerCodedSample gets `AVCodecContext.bits_per_coded_sample` value.
func (avctx *AVCodecContext) GetBitsPerCodedSample() int32 {
	return (int32)(avctx.bits_per_coded_sample)
}

// SetBitsPerCodedSample sets `AVCodecContext.bits_per_coded_sample` value.
func (avctx *AVCodecContext) SetBitsPerCodedSample(v int32) {
	avctx.bits_per_coded_sample = (C.int)(v)
}

// GetBitsPerCodedSampleAddr gets `AVCodecContext.bits_per_coded_sample` address.
func (avctx *AVCodecContext) GetBitsPerCodedSampleAddr() *int32 {
	return (*int32)(&avctx.bits_per_coded_sample)
}

// GetBitsPerRawSample gets `AVCodecContext.bits_per_raw_sample` value.
func (avctx *AVCodecContext) GetBitsPerRawSample() int32 {
	return (int32)(avctx.bits_per_raw_sample)
}

// SetBitsPerRawSample sets `AVCodecContext.bits_per_raw_sample` value.
func (avctx *AVCodecContext) SetBitsPerRawSample(v int32) {
	avctx.bits_per_raw_sample = (C.int)(v)
}

// GetBitsPerRawSampleAddr gets `AVCodecContext.bits_per_raw_sample` address.
func (avctx *AVCodecContext) GetBitsPerRawSampleAddr() *int32 {
	return (*int32)(&avctx.bits_per_raw_sample)
}

// GetLowres gets `AVCodecContext.lowres` value.
func (avctx *AVCodecContext) GetLowres() int32 {
	return (int32)(avctx.lowres)
}

// SetLowres sets `AVCodecContext.lowres` value.
func (avctx *AVCodecContext) SetLowres(v int32) {
	avctx.lowres = (C.int)(v)
}

// GetLowresAddr gets `AVCodecContext.lowres` address.
func (avctx *AVCodecContext) GetLowresAddr() *int32 {
	return (*int32)(&avctx.lowres)
}

// Deprecated: Use the quality factor packet side data instead.
//
// GetCodedFrame gets `AVCodecContext.coded_frame` value.
func (avctx *AVCodecContext) GetCodedFrame() *AVFrame {
	return (*AVFrame)(avctx.coded_frame)
}

// Deprecated: Use the quality factor packet side data instead.
//
// SetCodedFrame sets `AVCodecContext.coded_frame` value.
func (avctx *AVCodecContext) SetCodedFrame(v *AVFrame) {
	avctx.coded_frame = (*C.struct_AVFrame)(v)
}

// Deprecated: Use the quality factor packet side data instead.
//
// GetCodedFrameAddr gets `AVCodecContext.coded_frame` address.
func (avctx *AVCodecContext) GetCodedFrameAddr() **AVFrame {
	return (**AVFrame)(unsafe.Pointer(&avctx.coded_frame))
}

// GetThreadCount gets `AVCodecContext.thread_count` value.
func (avctx *AVCodecContext) GetThreadCount() int32 {
	return (int32)(avctx.thread_count)
}

// SetThreadCount sets `AVCodecContext.thread_count` value.
func (avctx *AVCodecContext) SetThreadCount(v int32) {
	avctx.thread_count = (C.int)(v)
}

// GetThreadCountAddr gets `AVCodecContext.thread_count` address.
func (avctx *AVCodecContext) GetThreadCountAddr() *int32 {
	return (*int32)(&avctx.thread_count)
}

// GetThreadType gets `AVCodecContext.threadtype` value.
func (avctx *AVCodecContext) GetThreadType() int32 {
	return (int32)(avctx.thread_type)
}

// SetThreadType sets `AVCodecContext.threadtype` value.
func (avctx *AVCodecContext) SetThreadType(v int32) {
	avctx.thread_type = (C.int)(v)
}

// GetThreadTypeAddr gets `AVCodecContext.threadtype` address.
func (avctx *AVCodecContext) GetThreadTypeAddr() *int32 {
	return (*int32)(&avctx.thread_type)
}

const (
	FF_THREAD_FRAME = int32(C.FF_THREAD_FRAME)
	FF_THREAD_SLICE = int32(C.FF_THREAD_SLICE)
)

// GetActiveThreadType gets `AVCodecContext.active_threadtype` value.
func (avctx *AVCodecContext) GetActiveThreadType() int32 {
	return (int32)(avctx.active_thread_type)
}

// SetActiveThreadType sets `AVCodecContext.active_threadtype` value.
func (avctx *AVCodecContext) SetActiveThreadType(v int32) {
	avctx.active_thread_type = (C.int)(v)
}

// GetActiveThreadTypeAddr gets `AVCodecContext.active_threadtype` address.
func (avctx *AVCodecContext) GetActiveThreadTypeAddr() *int32 {
	return (*int32)(&avctx.active_thread_type)
}

// Deprecated: Unused.
//
// GetThreadSafeCallbacks gets `AVCodecContext.thread_safe_callbacks` value.
func (avctx *AVCodecContext) GetThreadSafeCallbacks() int32 {
	return (int32)(avctx.thread_safe_callbacks)
}

// Deprecated: Unused.
//
// SetThreadSafeCallbacks sets `AVCodecContext.thread_safe_callbacks` value.
func (avctx *AVCodecContext) SetThreadSafeCallbacks(v int32) {
	avctx.thread_safe_callbacks = (C.int)(v)
}

// Deprecated: Unused.
//
// GetThreadSafeCallbacksAddr gets `AVCodecContext.thread_safe_callbacks` address.
func (avctx *AVCodecContext) GetThreadSafeCallbacksAddr() *int32 {
	return (*int32)(&avctx.thread_safe_callbacks)
}

// typedef int (*avcodec_context_internal_execute_func)(struct AVCodecContext *c,
// avcodec_context_execute_func func, void *arg2, int *ret, int count, int size);
type AvCodecContextInternalExecuteFunc = C.avcodec_context_internal_execute_func

// GetExecute gets `AVCodecContext.execute` value.
func (avctx *AVCodecContext) GetExecute() AvCodecContextInternalExecuteFunc {
	return (AvCodecContextInternalExecuteFunc)(avctx.execute)
}

// SetExecute sets `AVCodecContext.execute` value.
func (avctx *AVCodecContext) SetExecute(v AvCodecContextInternalExecuteFunc) {
	avctx.execute = (C.avcodec_context_internal_execute_func)(v)
}

// GetExecuteAddr gets `AVCodecContext.execute` address.
func (avctx *AVCodecContext) GetExecuteAddr() *AvCodecContextInternalExecuteFunc {
	return (*AvCodecContextInternalExecuteFunc)(&avctx.execute)
}

// typedef int (*avcodec_context_internal_execute2_func)(struct AVCodecContext *c,
// avcodec_context_execute2_func func, void *arg2, int *ret, int count);
type AvCodecContextInternalExecute2Func = C.avcodec_context_internal_execute2_func

// GetExecute2 gets `AVCodecContext.execute2` value.
func (avctx *AVCodecContext) GetExecute2() AvCodecContextInternalExecute2Func {
	return (AvCodecContextInternalExecute2Func)(avctx.execute2)
}

// SetExecute2 sets `AVCodecContext.execute2` value.
func (avctx *AVCodecContext) SetExecute2(v AvCodecContextInternalExecute2Func) {
	avctx.execute2 = (C.avcodec_context_internal_execute2_func)(v)
}

// GetExecute2Addr gets `AVCodecContext.execute2` address.
func (avctx *AVCodecContext) GetExecute2Addr() *AvCodecContextInternalExecute2Func {
	return (*AvCodecContextInternalExecute2Func)(&avctx.execute2)
}

// GetNsseWeight gets `AVCodecContext.nsse_weight` value.
func (avctx *AVCodecContext) GetNsseWeight() int32 {
	return (int32)(avctx.nsse_weight)
}

// SetNsseWeight sets `AVCodecContext.nsse_weight` value.
func (avctx *AVCodecContext) SetNsseWeight(v int32) {
	avctx.nsse_weight = (C.int)(v)
}

// GetNsseWeightAddr gets `AVCodecContext.nsse_weight` address.
func (avctx *AVCodecContext) GetNsseWeightAddr() *int32 {
	return (*int32)(&avctx.nsse_weight)
}

// GetProfile gets `AVCodecContext.profile` value.
func (avctx *AVCodecContext) GetProfile() int32 {
	return (int32)(avctx.profile)
}

// SetProfile sets `AVCodecContext.profile` value.
func (avctx *AVCodecContext) SetProfile(v int32) {
	avctx.profile = (C.int)(v)
}

// GetProfileAddr gets `AVCodecContext.profile` address.
func (avctx *AVCodecContext) GetProfileAddr() *int32 {
	return (*int32)(&avctx.profile)
}

const (
	FF_PROFILE_UNKNOWN  = int32(C.FF_PROFILE_UNKNOWN)
	FF_PROFILE_RESERVED = int32(C.FF_PROFILE_RESERVED)

	FF_PROFILE_AAC_MAIN      = int32(C.FF_PROFILE_AAC_MAIN)
	FF_PROFILE_AAC_LOW       = int32(C.FF_PROFILE_AAC_LOW)
	FF_PROFILE_AAC_SSR       = int32(C.FF_PROFILE_AAC_SSR)
	FF_PROFILE_AAC_LTP       = int32(C.FF_PROFILE_AAC_LTP)
	FF_PROFILE_AAC_HE        = int32(C.FF_PROFILE_AAC_HE)
	FF_PROFILE_AAC_HE_V2     = int32(C.FF_PROFILE_AAC_HE_V2)
	FF_PROFILE_AAC_LD        = int32(C.FF_PROFILE_AAC_LD)
	FF_PROFILE_AAC_ELD       = int32(C.FF_PROFILE_AAC_ELD)
	FF_PROFILE_MPEG2_AAC_LOW = int32(C.FF_PROFILE_MPEG2_AAC_LOW)
	FF_PROFILE_MPEG2_AAC_HE  = int32(C.FF_PROFILE_MPEG2_AAC_HE)

	FF_PROFILE_DNXHD     = int32(C.FF_PROFILE_DNXHD)
	FF_PROFILE_DNXHR_LB  = int32(C.FF_PROFILE_DNXHR_LB)
	FF_PROFILE_DNXHR_SQ  = int32(C.FF_PROFILE_DNXHR_SQ)
	FF_PROFILE_DNXHR_HQ  = int32(C.FF_PROFILE_DNXHR_HQ)
	FF_PROFILE_DNXHR_HQX = int32(C.FF_PROFILE_DNXHR_HQX)
	FF_PROFILE_DNXHR_444 = int32(C.FF_PROFILE_DNXHR_444)

	FF_PROFILE_DTS         = int32(C.FF_PROFILE_DTS)
	FF_PROFILE_DTS_ES      = int32(C.FF_PROFILE_DTS_ES)
	FF_PROFILE_DTS_96_24   = int32(C.FF_PROFILE_DTS_96_24)
	FF_PROFILE_DTS_HD_HRA  = int32(C.FF_PROFILE_DTS_HD_HRA)
	FF_PROFILE_DTS_HD_MA   = int32(C.FF_PROFILE_DTS_HD_MA)
	FF_PROFILE_DTS_EXPRESS = int32(C.FF_PROFILE_DTS_EXPRESS)

	FF_PROFILE_MPEG2_422          = int32(C.FF_PROFILE_MPEG2_422)
	FF_PROFILE_MPEG2_HIGH         = int32(C.FF_PROFILE_MPEG2_HIGH)
	FF_PROFILE_MPEG2_SS           = int32(C.FF_PROFILE_MPEG2_SS)
	FF_PROFILE_MPEG2_SNR_SCALABLE = int32(C.FF_PROFILE_MPEG2_SNR_SCALABLE)
	FF_PROFILE_MPEG2_MAIN         = int32(C.FF_PROFILE_MPEG2_MAIN)
	FF_PROFILE_MPEG2_SIMPLE       = int32(C.FF_PROFILE_MPEG2_SIMPLE)

	FF_PROFILE_H264_CONSTRAINED = int32(C.FF_PROFILE_H264_CONSTRAINED)
	FF_PROFILE_H264_INTRA       = int32(C.FF_PROFILE_H264_INTRA)

	FF_PROFILE_H264_BASELINE             = int32(C.FF_PROFILE_H264_BASELINE)
	FF_PROFILE_H264_CONSTRAINED_BASELINE = int32(C.FF_PROFILE_H264_CONSTRAINED_BASELINE)
	FF_PROFILE_H264_MAIN                 = int32(C.FF_PROFILE_H264_MAIN)
	FF_PROFILE_H264_EXTENDED             = int32(C.FF_PROFILE_H264_EXTENDED)
	FF_PROFILE_H264_HIGH                 = int32(C.FF_PROFILE_H264_HIGH)
	FF_PROFILE_H264_HIGH_10              = int32(C.FF_PROFILE_H264_HIGH_10)
	FF_PROFILE_H264_HIGH_10_INTRA        = int32(C.FF_PROFILE_H264_HIGH_10_INTRA)
	FF_PROFILE_H264_MULTIVIEW_HIGH       = int32(C.FF_PROFILE_H264_MULTIVIEW_HIGH)
	FF_PROFILE_H264_HIGH_422             = int32(C.FF_PROFILE_H264_HIGH_422)
	FF_PROFILE_H264_HIGH_422_INTRA       = int32(C.FF_PROFILE_H264_HIGH_422_INTRA)
	FF_PROFILE_H264_STEREO_HIGH          = int32(C.FF_PROFILE_H264_STEREO_HIGH)
	FF_PROFILE_H264_HIGH_444             = int32(C.FF_PROFILE_H264_HIGH_444)
	FF_PROFILE_H264_HIGH_444_PREDICTIVE  = int32(C.FF_PROFILE_H264_HIGH_444_PREDICTIVE)
	FF_PROFILE_H264_HIGH_444_INTRA       = int32(C.FF_PROFILE_H264_HIGH_444_INTRA)
	FF_PROFILE_H264_CAVLC_444            = int32(C.FF_PROFILE_H264_CAVLC_444)

	FF_PROFILE_VC1_SIMPLE   = int32(C.FF_PROFILE_VC1_SIMPLE)
	FF_PROFILE_VC1_MAIN     = int32(C.FF_PROFILE_VC1_MAIN)
	FF_PROFILE_VC1_COMPLEX  = int32(C.FF_PROFILE_VC1_COMPLEX)
	FF_PROFILE_VC1_ADVANCED = int32(C.FF_PROFILE_VC1_ADVANCED)

	FF_PROFILE_MPEG4_SIMPLE                    = int32(C.FF_PROFILE_MPEG4_SIMPLE)
	FF_PROFILE_MPEG4_SIMPLE_SCALABLE           = int32(C.FF_PROFILE_MPEG4_SIMPLE_SCALABLE)
	FF_PROFILE_MPEG4_CORE                      = int32(C.FF_PROFILE_MPEG4_CORE)
	FF_PROFILE_MPEG4_MAIN                      = int32(C.FF_PROFILE_MPEG4_MAIN)
	FF_PROFILE_MPEG4_N_BIT                     = int32(C.FF_PROFILE_MPEG4_N_BIT)
	FF_PROFILE_MPEG4_SCALABLE_TEXTURE          = int32(C.FF_PROFILE_MPEG4_SCALABLE_TEXTURE)
	FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION     = int32(C.FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION)
	FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE    = int32(C.FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE)
	FF_PROFILE_MPEG4_HYBRID                    = int32(C.FF_PROFILE_MPEG4_HYBRID)
	FF_PROFILE_MPEG4_ADVANCED_REAL_TIME        = int32(C.FF_PROFILE_MPEG4_ADVANCED_REAL_TIME)
	FF_PROFILE_MPEG4_CORE_SCALABLE             = int32(C.FF_PROFILE_MPEG4_CORE_SCALABLE)
	FF_PROFILE_MPEG4_ADVANCED_CODING           = int32(C.FF_PROFILE_MPEG4_ADVANCED_CODING)
	FF_PROFILE_MPEG4_ADVANCED_CORE             = int32(C.FF_PROFILE_MPEG4_ADVANCED_CORE)
	FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE = int32(C.FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE)
	FF_PROFILE_MPEG4_SIMPLE_STUDIO             = int32(C.FF_PROFILE_MPEG4_SIMPLE_STUDIO)
	FF_PROFILE_MPEG4_ADVANCED_SIMPLE           = int32(C.FF_PROFILE_MPEG4_ADVANCED_SIMPLE)

	FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0  = int32(C.FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0)
	FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1  = int32(C.FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1)
	FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION = int32(C.FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION)
	FF_PROFILE_JPEG2000_DCINEMA_2K             = int32(C.FF_PROFILE_JPEG2000_DCINEMA_2K)
	FF_PROFILE_JPEG2000_DCINEMA_4K             = int32(C.FF_PROFILE_JPEG2000_DCINEMA_4K)

	FF_PROFILE_VP9_0 = int32(C.FF_PROFILE_VP9_0)
	FF_PROFILE_VP9_1 = int32(C.FF_PROFILE_VP9_1)
	FF_PROFILE_VP9_2 = int32(C.FF_PROFILE_VP9_2)
	FF_PROFILE_VP9_3 = int32(C.FF_PROFILE_VP9_3)

	FF_PROFILE_HEVC_MAIN               = int32(C.FF_PROFILE_HEVC_MAIN)
	FF_PROFILE_HEVC_MAIN_10            = int32(C.FF_PROFILE_HEVC_MAIN_10)
	FF_PROFILE_HEVC_MAIN_STILL_PICTURE = int32(C.FF_PROFILE_HEVC_MAIN_STILL_PICTURE)
	FF_PROFILE_HEVC_REXT               = int32(C.FF_PROFILE_HEVC_REXT)

	FF_PROFILE_VVC_MAIN_10     = int32(C.FF_PROFILE_VVC_MAIN_10)
	FF_PROFILE_VVC_MAIN_10_444 = int32(C.FF_PROFILE_VVC_MAIN_10_444)

	FF_PROFILE_AV1_MAIN         = int32(C.FF_PROFILE_AV1_MAIN)
	FF_PROFILE_AV1_HIGH         = int32(C.FF_PROFILE_AV1_HIGH)
	FF_PROFILE_AV1_PROFESSIONAL = int32(C.FF_PROFILE_AV1_PROFESSIONAL)

	FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT            = int32(C.FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT)
	FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT = int32(C.FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT)
	FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT         = int32(C.FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT)
	FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS                = int32(C.FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS)
	FF_PROFILE_MJPEG_JPEG_LS                         = int32(C.FF_PROFILE_MJPEG_JPEG_LS)

	FF_PROFILE_SBC_MSBC = int32(C.FF_PROFILE_SBC_MSBC)

	FF_PROFILE_PRORES_PROXY    = int32(C.FF_PROFILE_PRORES_PROXY)
	FF_PROFILE_PRORES_LT       = int32(C.FF_PROFILE_PRORES_LT)
	FF_PROFILE_PRORES_STANDARD = int32(C.FF_PROFILE_PRORES_STANDARD)
	FF_PROFILE_PRORES_HQ       = int32(C.FF_PROFILE_PRORES_HQ)
	FF_PROFILE_PRORES_4444     = int32(C.FF_PROFILE_PRORES_4444)
	FF_PROFILE_PRORES_XQ       = int32(C.FF_PROFILE_PRORES_XQ)

	FF_PROFILE_ARIB_PROFILE_A = int32(C.FF_PROFILE_ARIB_PROFILE_A)
	FF_PROFILE_ARIB_PROFILE_C = int32(C.FF_PROFILE_ARIB_PROFILE_C)

	FF_PROFILE_KLVA_SYNC  = int32(C.FF_PROFILE_KLVA_SYNC)
	FF_PROFILE_KLVA_ASYNC = int32(C.FF_PROFILE_KLVA_ASYNC)
)

// GetLevel gets `AVCodecContext.level` value.
func (avctx *AVCodecContext) GetLevel() int32 {
	return (int32)(avctx.level)
}

// SetLevel sets `AVCodecContext.level` value.
func (avctx *AVCodecContext) SetLevel(v int32) {
	avctx.level = (C.int)(v)
}

// GetLevelAddr gets `AVCodecContext.level` address.
func (avctx *AVCodecContext) GetLevelAddr() *int32 {
	return (*int32)(&avctx.level)
}

// GetSkipLoopFilter gets `AVCodecContext.skip_loop_filter` value.
func (avctx *AVCodecContext) GetSkipLoopFilter() AVDiscard {
	return (AVDiscard)(avctx.skip_loop_filter)
}

// SetSkipLoopFilter sets `AVCodecContext.skip_loop_filter` value.
func (avctx *AVCodecContext) SetSkipLoopFilter(v AVDiscard) {
	avctx.skip_loop_filter = (C.enum_AVDiscard)(v)
}

// GetSkipLoopFilterAddr gets `AVCodecContext.skip_loop_filter` address.
func (avctx *AVCodecContext) GetSkipLoopFilterAddr() *AVDiscard {
	return (*AVDiscard)(&avctx.skip_loop_filter)
}

// GetSkipIdct gets `AVCodecContext.skip_idct` value.
func (avctx *AVCodecContext) GetSkipIdct() AVDiscard {
	return (AVDiscard)(avctx.skip_idct)
}

// SetSkipIdct sets `AVCodecContext.skip_idct` value.
func (avctx *AVCodecContext) SetSkipIdct(v AVDiscard) {
	avctx.skip_idct = (C.enum_AVDiscard)(v)
}

// GetSkipIdctAddr gets `AVCodecContext.skip_idct` address.
func (avctx *AVCodecContext) GetSkipIdctAddr() *AVDiscard {
	return (*AVDiscard)(&avctx.skip_idct)
}

// GetSkipFrame gets `AVCodecContext.skip_frame` value.
func (avctx *AVCodecContext) GetSkipFrame() AVDiscard {
	return (AVDiscard)(avctx.skip_frame)
}

// SetSkipFrame sets `AVCodecContext.skip_frame` value.
func (avctx *AVCodecContext) SetSkipFrame(v AVDiscard) {
	avctx.skip_frame = (C.enum_AVDiscard)(v)
}

// GetSkipFrameAddr gets `AVCodecContext.skip_frame` address.
func (avctx *AVCodecContext) GetSkipFrameAddr() *AVDiscard {
	return (*AVDiscard)(&avctx.skip_frame)
}

// GetSubtitleHeader gets `AVCodecContext.subtitle_header` value.
func (avctx *AVCodecContext) GetSubtitleHeader() *uint8 {
	return (*uint8)(avctx.subtitle_header)
}

// SetSubtitleHeader sets `AVCodecContext.subtitle_header` value.
func (avctx *AVCodecContext) SetSubtitleHeader(v *uint8) {
	avctx.subtitle_header = (*C.uint8_t)(v)
}

// GetSubtitleHeaderAddr gets `AVCodecContext.subtitle_header` address.
func (avctx *AVCodecContext) GetSubtitleHeaderAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&avctx.subtitle_header))
}

// GetSubtitleHeaderSize gets `AVCodecContext.subtitle_header_size` value.
func (avctx *AVCodecContext) GetSubtitleHeaderSize() int32 {
	return (int32)(avctx.subtitle_header_size)
}

// SetSubtitleHeaderSize sets `AVCodecContext.subtitle_header_size` value.
func (avctx *AVCodecContext) SetSubtitleHeaderSize(v int32) {
	avctx.subtitle_header_size = (C.int)(v)
}

// GetSubtitleHeaderSizeAddr gets `AVCodecContext.subtitle_header_size` address.
func (avctx *AVCodecContext) GetSubtitleHeaderSizeAddr() *int32 {
	return (*int32)(&avctx.subtitle_header_size)
}

// Deprecated: No use.
//
// GetVbvDelay gets `AVCodecContext.vbv_delay` value.
func (avctx *AVCodecContext) GetVbvDelay() uint64 {
	return (uint64)(avctx.vbv_delay)
}

// Deprecated: No use.
//
// SetVbvDelay sets `AVCodecContext.vbv_delay` value.
func (avctx *AVCodecContext) SetVbvDelay(v uint64) {
	avctx.vbv_delay = (C.uint64_t)(v)
}

// Deprecated: No use.
//
// GetVbvDelayAddr gets `AVCodecContext.vbv_delay` address.
func (avctx *AVCodecContext) GetVbvDelayAddr() *uint64 {
	return (*uint64)(&avctx.vbv_delay)
}

// Deprecated: No use.
//
// GetSideDataOnlyPackets gets `AVCodecContext.side_data_only_packets` value.
func (avctx *AVCodecContext) GetSideDataOnlyPackets() int32 {
	return (int32)(avctx.side_data_only_packets)
}

// Deprecated: No use.
//
// SetSideDataOnlyPackets sets `AVCodecContext.side_data_only_packets` value.
func (avctx *AVCodecContext) SetSideDataOnlyPackets(v int32) {
	avctx.side_data_only_packets = (C.int)(v)
}

// Deprecated: No use.
//
// GetSideDataOnlyPacketsAddr gets `AVCodecContext.side_data_only_packets` address.
func (avctx *AVCodecContext) GetSideDataOnlyPacketsAddr() *int32 {
	return (*int32)(&avctx.side_data_only_packets)
}

// GetInitialPadding gets `AVCodecContext.initial_padding` value.
func (avctx *AVCodecContext) GetInitialPadding() int32 {
	return (int32)(avctx.initial_padding)
}

// SetInitialPadding sets `AVCodecContext.initial_padding` value.
func (avctx *AVCodecContext) SetInitialPadding(v int32) {
	avctx.initial_padding = (C.int)(v)
}

// GetInitialPaddingAddr gets `AVCodecContext.initial_padding` address.
func (avctx *AVCodecContext) GetInitialPaddingAddr() *int32 {
	return (*int32)(&avctx.initial_padding)
}

// GetFramerate gets `AVCodecContext.framerate` value.
func (avctx *AVCodecContext) GetFramerate() AVRational {
	return (AVRational)(avctx.framerate)
}

// SetFramerate sets `AVCodecContext.framerate` value.
func (avctx *AVCodecContext) SetFramerate(v AVRational) {
	avctx.framerate = (C.struct_AVRational)(v)
}

// GetFramerateAddr gets `AVCodecContext.framerate` address.
func (avctx *AVCodecContext) GetFramerateAddr() *AVRational {
	return (*AVRational)(&avctx.framerate)
}

// GetSwPixFmt gets `AVCodecContext.sw_pix_fmt` value.
func (avctx *AVCodecContext) GetSwPixFmt() AVPixelFormat {
	return (AVPixelFormat)(avctx.sw_pix_fmt)
}

// SetSwPixFmt sets `AVCodecContext.sw_pix_fmt` value.
func (avctx *AVCodecContext) SetSwPixFmt(v AVPixelFormat) {
	avctx.sw_pix_fmt = (C.enum_AVPixelFormat)(v)
}

// GetSwPixFmtAddr gets `AVCodecContext.sw_pix_fmt` address.
func (avctx *AVCodecContext) GetSwPixFmtAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&avctx.sw_pix_fmt)
}

// GetPktTimebase gets `AVCodecContext.pkt_timebase` value.
func (avctx *AVCodecContext) GetPktTimebase() AVRational {
	return (AVRational)(avctx.pkt_timebase)
}

// SetPktTimebase sets `AVCodecContext.pkt_timebase` value.
func (avctx *AVCodecContext) SetPktTimebase(v AVRational) {
	avctx.pkt_timebase = (C.struct_AVRational)(v)
}

// GetPktTimebaseAddr gets `AVCodecContext.pkt_timebase` address.
func (avctx *AVCodecContext) GetPktTimebaseAddr() *AVRational {
	return (*AVRational)(&avctx.pkt_timebase)
}

// GetCodecDescriptor gets `AVCodecContext.codec_descriptor` value.
func (avctx *AVCodecContext) GetCodecDescriptor() *AVCodecDescriptor {
	return (*AVCodecDescriptor)(avctx.codec_descriptor)
}

// SetCodecDescriptor sets `AVCodecContext.codec_descriptor` value.
func (avctx *AVCodecContext) SetCodecDescriptor(v *AVCodecDescriptor) {
	avctx.codec_descriptor = (*C.struct_AVCodecDescriptor)(v)
}

// GetCodecDescriptorAddr gets `AVCodecContext.codec_descriptor` address.
func (avctx *AVCodecContext) GetCodecDescriptorAddr() **AVCodecDescriptor {
	return (**AVCodecDescriptor)(unsafe.Pointer(&avctx.codec_descriptor))
}

// GetPtsCorrectionNumFaultyPts gets `AVCodecContext.pts_correction_num_faulty_pts` value.
func (avctx *AVCodecContext) GetPtsCorrectionNumFaultyPts() int64 {
	return (int64)(avctx.pts_correction_num_faulty_pts)
}

// SetPtsCorrectionNumFaultyPts sets `AVCodecContext.pts_correction_num_faulty_pts` value.
func (avctx *AVCodecContext) SetPtsCorrectionNumFaultyPts(v int64) {
	avctx.pts_correction_num_faulty_pts = (C.int64_t)(v)
}

// GetPtsCorrectionNumFaultyPtsAddr gets `AVCodecContext.pts_correction_num_faulty_pts` address.
func (avctx *AVCodecContext) GetPtsCorrectionNumFaultyPtsAddr() *int64 {
	return (*int64)(&avctx.pts_correction_num_faulty_pts)
}

// GetPtsCorrectionNumFaultyDts gets `AVCodecContext.pts_correction_num_faulty_dts` value.
func (avctx *AVCodecContext) GetPtsCorrectionNumFaultyDts() int64 {
	return (int64)(avctx.pts_correction_num_faulty_dts)
}

// SetPtsCorrectionNumFaultyDts sets `AVCodecContext.pts_correction_num_faulty_dts` value.
func (avctx *AVCodecContext) SetPtsCorrectionNumFaultyDts(v int64) {
	avctx.pts_correction_num_faulty_dts = (C.int64_t)(v)
}

// GetPtsCorrectionNumFaultyDtsAddr gets `AVCodecContext.pts_correction_num_faulty_dts` address.
func (avctx *AVCodecContext) GetPtsCorrectionNumFaultyDtsAddr() *int64 {
	return (*int64)(&avctx.pts_correction_num_faulty_dts)
}

// GetPtsCorrectionLastPts gets `AVCodecContext.pts_correction_last_pts` value.
func (avctx *AVCodecContext) GetPtsCorrectionLastPts() int64 {
	return (int64)(avctx.pts_correction_last_pts)
}

// SetPtsCorrectionLastPts sets `AVCodecContext.pts_correction_last_pts` value.
func (avctx *AVCodecContext) SetPtsCorrectionLastPts(v int64) {
	avctx.pts_correction_last_pts = (C.int64_t)(v)
}

// GetPtsCorrectionLastPtsAddr gets `AVCodecContext.pts_correction_last_pts` address.
func (avctx *AVCodecContext) GetPtsCorrectionLastPtsAddr() *int64 {
	return (*int64)(&avctx.pts_correction_last_pts)
}

// GetPtsCorrectionLastDts gets `AVCodecContext.pts_correction_last_dts` value.
func (avctx *AVCodecContext) GetPtsCorrectionLastDts() int64 {
	return (int64)(avctx.pts_correction_last_dts)
}

// SetPtsCorrectionLastDts sets `AVCodecContext.pts_correction_last_dts` value.
func (avctx *AVCodecContext) SetPtsCorrectionLastDts(v int64) {
	avctx.pts_correction_last_dts = (C.int64_t)(v)
}

// GetPtsCorrectionLastDtsAddr gets `AVCodecContext.pts_correction_last_dts` address.
func (avctx *AVCodecContext) GetPtsCorrectionLastDtsAddr() *int64 {
	return (*int64)(&avctx.pts_correction_last_dts)
}

// GetSubCharenc gets `AVCodecContext.sub_charenc` value.
func (avctx *AVCodecContext) GetSubCharenc() string {
	return C.GoString(avctx.sub_charenc)
}

// GetSubCharencMode gets `AVCodecContext.sub_charenc_mode` value.
func (avctx *AVCodecContext) GetSubCharencMode() int32 {
	return (int32)(avctx.sub_charenc_mode)
}

// SetSubCharencMode sets `AVCodecContext.sub_charenc_mode` value.
func (avctx *AVCodecContext) SetSubCharencMode(v int32) {
	avctx.sub_charenc_mode = (C.int)(v)
}

// GetSubCharencModeAddr gets `AVCodecContext.sub_charenc_mode` address.
func (avctx *AVCodecContext) GetSubCharencModeAddr() *int32 {
	return (*int32)(&avctx.sub_charenc_mode)
}

const (
	FF_SUB_CHARENC_MODE_DO_NOTHING  = int32(C.FF_SUB_CHARENC_MODE_DO_NOTHING)
	FF_SUB_CHARENC_MODE_AUTOMATIC   = int32(C.FF_SUB_CHARENC_MODE_AUTOMATIC)
	FF_SUB_CHARENC_MODE_PRE_DECODER = int32(C.FF_SUB_CHARENC_MODE_PRE_DECODER)
	FF_SUB_CHARENC_MODE_IGNORE      = int32(C.FF_SUB_CHARENC_MODE_IGNORE)
)

// GetSkipAlpha gets `AVCodecContext.skip_alpha` value.
func (avctx *AVCodecContext) GetSkipAlpha() int32 {
	return (int32)(avctx.skip_alpha)
}

// SetSkipAlpha sets `AVCodecContext.skip_alpha` value.
func (avctx *AVCodecContext) SetSkipAlpha(v int32) {
	avctx.skip_alpha = (C.int)(v)
}

// GetSkipAlphaAddr gets `AVCodecContext.skip_alpha` address.
func (avctx *AVCodecContext) GetSkipAlphaAddr() *int32 {
	return (*int32)(&avctx.skip_alpha)
}

// GetSeekPreroll gets `AVCodecContext.seek_preroll` value.
func (avctx *AVCodecContext) GetSeekPreroll() int32 {
	return (int32)(avctx.seek_preroll)
}

// SetSeekPreroll sets `AVCodecContext.seek_preroll` value.
func (avctx *AVCodecContext) SetSeekPreroll(v int32) {
	avctx.seek_preroll = (C.int)(v)
}

// GetSeekPrerollAddr gets `AVCodecContext.seek_preroll` address.
func (avctx *AVCodecContext) GetSeekPrerollAddr() *int32 {
	return (*int32)(&avctx.seek_preroll)
}

// Deprecated: Unused.
//
// GetDebugMv gets `AVCodecContext.debug_mv` value.
func (avctx *AVCodecContext) GetDebugMv() int32 {
	return (int32)(avctx.debug_mv)
}

// Deprecated: Unused.
//
// SetDebugMv sets `AVCodecContext.debug_mv` value.
func (avctx *AVCodecContext) SetDebugMv(v int32) {
	avctx.debug_mv = (C.int)(v)
}

// Deprecated: Unused.
//
// GetDebugMvAddr gets `AVCodecContext.debug_mv` address.
func (avctx *AVCodecContext) GetDebugMvAddr() *int32 {
	return (*int32)(&avctx.debug_mv)
}

const (
	FF_DEBUG_VIS_MV_P_FOR  = int32(C.FF_DEBUG_VIS_MV_P_FOR)
	FF_DEBUG_VIS_MV_B_FOR  = int32(C.FF_DEBUG_VIS_MV_B_FOR)
	FF_DEBUG_VIS_MV_B_BACK = int32(C.FF_DEBUG_VIS_MV_B_BACK)
)

// GetChromaIntraMatrix gets `AVCodecContext.chroma_intra_matrix` value.
func (avctx *AVCodecContext) GetChromaIntraMatrix() *uint16 {
	return (*uint16)(avctx.chroma_intra_matrix)
}

// SetChromaIntraMatrix sets `AVCodecContext.chroma_intra_matrix` value.
func (avctx *AVCodecContext) SetChromaIntraMatrix(v *uint16) {
	avctx.chroma_intra_matrix = (*C.uint16_t)(v)
}

// GetChromaIntraMatrixAddr gets `AVCodecContext.chroma_intra_matrix` address.
func (avctx *AVCodecContext) GetChromaIntraMatrixAddr() **uint16 {
	return (**uint16)(unsafe.Pointer(&avctx.chroma_intra_matrix))
}

// GetDumpSeparator gets `AVCodecContext.dump_separator` value.
func (avctx *AVCodecContext) GetDumpSeparator() *uint8 {
	return (*uint8)(avctx.dump_separator)
}

// SetDumpSeparator sets `AVCodecContext.dump_separator` value.
func (avctx *AVCodecContext) SetDumpSeparator(v *uint8) {
	avctx.dump_separator = (*C.uint8_t)(v)
}

// GetDumpSeparatorAddr gets `AVCodecContext.dump_separator` address.
func (avctx *AVCodecContext) GetDumpSeparatorAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&avctx.dump_separator))
}

// GetCodecWhitelist gets `AVCodecContext.codec_whitelist` value.
func (avctx *AVCodecContext) GetCodecWhitelist() string {
	return C.GoString(avctx.codec_whitelist)
}

// GetProperties gets `AVCodecContext.properties` value.
func (avctx *AVCodecContext) GetProperties() uint32 {
	return (uint32)(avctx.properties)
}

// SetProperties sets `AVCodecContext.properties` value.
func (avctx *AVCodecContext) SetProperties(v uint32) {
	avctx.properties = (C.uint)(v)
}

// GetPropertiesAddr gets `AVCodecContext.properties` address.
func (avctx *AVCodecContext) GetPropertiesAddr() *uint32 {
	return (*uint32)(&avctx.properties)
}

const (
	FF_CODEC_PROPERTY_LOSSLESS        = uint32(C.FF_CODEC_PROPERTY_LOSSLESS)
	FF_CODEC_PROPERTY_CLOSED_CAPTIONS = uint32(C.FF_CODEC_PROPERTY_CLOSED_CAPTIONS)
)

// GetCodedSideData gets `AVCodecContext.coded_side_data` value.
func (avctx *AVCodecContext) GetCodedSideData() *AVPacketSideData {
	return (*AVPacketSideData)(avctx.coded_side_data)
}

// SetCodedSideData sets `AVCodecContext.coded_side_data` value.
func (avctx *AVCodecContext) SetCodedSideData(v *AVPacketSideData) {
	avctx.coded_side_data = (*C.AVPacketSideData)(v)
}

// GetCodedSideDataAddr gets `AVCodecContext.coded_side_data` address.
func (avctx *AVCodecContext) GetCodedSideDataAddr() **AVPacketSideData {
	return (**AVPacketSideData)(unsafe.Pointer(&avctx.coded_side_data))
}

// GetNbCodedSideData gets `AVCodecContext.nb_coded_side_data` value.
func (avctx *AVCodecContext) GetNbCodedSideData() int32 {
	return (int32)(avctx.nb_coded_side_data)
}

// SetNbCodedSideData sets `AVCodecContext.nb_coded_side_data` value.
func (avctx *AVCodecContext) SetNbCodedSideData(v int32) {
	avctx.nb_coded_side_data = (C.int)(v)
}

// GetNbCodedSideDataAddr gets `AVCodecContext.nb_coded_side_data` address.
func (avctx *AVCodecContext) GetNbCodedSideDataAddr() *int32 {
	return (*int32)(&avctx.nb_coded_side_data)
}

// GetHwFramesCtx gets `AVCodecContext.hw_frames_ctx` value.
func (avctx *AVCodecContext) GetHwFramesCtx() *AVBufferRef {
	return (*AVBufferRef)(avctx.hw_frames_ctx)
}

// SetHwFramesCtx sets `AVCodecContext.hw_frames_ctx` value.
func (avctx *AVCodecContext) SetHwFramesCtx(v *AVBufferRef) {
	avctx.hw_frames_ctx = (*C.AVBufferRef)(v)
}

// GetHwFramesCtxAddr gets `AVCodecContext.hw_frames_ctx` address.
func (avctx *AVCodecContext) GetHwFramesCtxAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&avctx.hw_frames_ctx))
}

// GetSubTextFormat gets `AVCodecContext.sub_text_format` value.
func (avctx *AVCodecContext) GetSubTextFormat() int32 {
	return (int32)(avctx.sub_text_format)
}

// SetSubTextFormat sets `AVCodecContext.sub_text_format` value.
func (avctx *AVCodecContext) SetSubTextFormat(v int32) {
	avctx.sub_text_format = (C.int)(v)
}

// GetSubTextFormatAddr gets `AVCodecContext.sub_text_format` address.
func (avctx *AVCodecContext) GetSubTextFormatAddr() *int32 {
	return (*int32)(&avctx.sub_text_format)
}

const (
	FF_SUB_TEXT_FMT_ASS              = int32(C.FF_SUB_TEXT_FMT_ASS)
	FF_SUB_TEXT_FMT_ASS_WITH_TIMINGS = int32(C.FF_SUB_TEXT_FMT_ASS_WITH_TIMINGS)
)

// GetTrailingPadding gets `AVCodecContext.trailing_padding` value.
func (avctx *AVCodecContext) GetTrailingPadding() int32 {
	return (int32)(avctx.trailing_padding)
}

// SetTrailingPadding sets `AVCodecContext.trailing_padding` value.
func (avctx *AVCodecContext) SetTrailingPadding(v int32) {
	avctx.trailing_padding = (C.int)(v)
}

// GetTrailingPaddingAddr gets `AVCodecContext.trailing_padding` address.
func (avctx *AVCodecContext) GetTrailingPaddingAddr() *int32 {
	return (*int32)(&avctx.trailing_padding)
}

// GetMaxPixels gets `AVCodecContext.max_pixels` value.
func (avctx *AVCodecContext) GetMaxPixels() int64 {
	return (int64)(avctx.max_pixels)
}

// SetMaxPixels sets `AVCodecContext.max_pixels` value.
func (avctx *AVCodecContext) SetMaxPixels(v int64) {
	avctx.max_pixels = (C.int64_t)(v)
}

// GetMaxPixelsAddr gets `AVCodecContext.max_pixels` address.
func (avctx *AVCodecContext) GetMaxPixelsAddr() *int64 {
	return (*int64)(&avctx.max_pixels)
}

// GetHwDeviceCtx gets `AVCodecContext.hw_device_ctx` value.
func (avctx *AVCodecContext) GetHwDeviceCtx() *AVBufferRef {
	return (*AVBufferRef)(avctx.hw_device_ctx)
}

// SetHwDeviceCtx sets `AVCodecContext.hw_device_ctx` value.
func (avctx *AVCodecContext) SetHwDeviceCtx(v *AVBufferRef) {
	avctx.hw_device_ctx = (*C.AVBufferRef)(v)
}

// GetHwDeviceCtxAddr gets `AVCodecContext.hw_device_ctx` address.
func (avctx *AVCodecContext) GetHwDeviceCtxAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&avctx.hw_device_ctx))
}

// GetHwaccelFlags gets `AVCodecContext.hwaccel_flags` value.
func (avctx *AVCodecContext) GetHwaccelFlags() int32 {
	return (int32)(avctx.hwaccel_flags)
}

// SetHwaccelFlags sets `AVCodecContext.hwaccel_flags` value.
func (avctx *AVCodecContext) SetHwaccelFlags(v int32) {
	avctx.hwaccel_flags = (C.int)(v)
}

// GetHwaccelFlagsAddr gets `AVCodecContext.hwaccel_flags` address.
func (avctx *AVCodecContext) GetHwaccelFlagsAddr() *int32 {
	return (*int32)(&avctx.hwaccel_flags)
}

// GetApplyCropping gets `AVCodecContext.apply_cropping` value.
func (avctx *AVCodecContext) GetApplyCropping() int32 {
	return (int32)(avctx.apply_cropping)
}

// SetApplyCropping sets `AVCodecContext.apply_cropping` value.
func (avctx *AVCodecContext) SetApplyCropping(v int32) {
	avctx.apply_cropping = (C.int)(v)
}

// GetApplyCroppingAddr gets `AVCodecContext.apply_cropping` address.
func (avctx *AVCodecContext) GetApplyCroppingAddr() *int32 {
	return (*int32)(&avctx.apply_cropping)
}

// GetExtraHwFrames gets `AVCodecContext.extra_hw_frames` value.
func (avctx *AVCodecContext) GetExtraHwFrames() int32 {
	return (int32)(avctx.extra_hw_frames)
}

// SetExtraHwFrames sets `AVCodecContext.extra_hw_frames` value.
func (avctx *AVCodecContext) SetExtraHwFrames(v int32) {
	avctx.extra_hw_frames = (C.int)(v)
}

// GetExtraHwFramesAddr gets `AVCodecContext.extra_hw_frames` address.
func (avctx *AVCodecContext) GetExtraHwFramesAddr() *int32 {
	return (*int32)(&avctx.extra_hw_frames)
}

// GetDiscardDamagedPercentage gets `AVCodecContext.discard_damaged_percentage` value.
func (avctx *AVCodecContext) GetDiscardDamagedPercentage() int32 {
	return (int32)(avctx.discard_damaged_percentage)
}

// SetDiscardDamagedPercentage sets `AVCodecContext.discard_damaged_percentage` value.
func (avctx *AVCodecContext) SetDiscardDamagedPercentage(v int32) {
	avctx.discard_damaged_percentage = (C.int)(v)
}

// GetDiscardDamagedPercentageAddr gets `AVCodecContext.discard_damaged_percentage` address.
func (avctx *AVCodecContext) GetDiscardDamagedPercentageAddr() *int32 {
	return (*int32)(&avctx.discard_damaged_percentage)
}

// GetMaxSamples gets `AVCodecContext.max_samples` value.
func (avctx *AVCodecContext) GetMaxSamples() int64 {
	return (int64)(avctx.max_samples)
}

// SetMaxSamples sets `AVCodecContext.max_samples` value.
func (avctx *AVCodecContext) SetMaxSamples(v int64) {
	avctx.max_samples = (C.int64_t)(v)
}

// GetMaxSamplesAddr gets `AVCodecContext.max_samples` address.
func (avctx *AVCodecContext) GetMaxSamplesAddr() *int64 {
	return (*int64)(&avctx.max_samples)
}

// GetExportSideData gets `AVCodecContext.export_side_data` value.
func (avctx *AVCodecContext) GetExportSideData() int32 {
	return (int32)(avctx.export_side_data)
}

// SetExportSideData sets `AVCodecContext.export_side_data` value.
func (avctx *AVCodecContext) SetExportSideData(v int32) {
	avctx.export_side_data = (C.int)(v)
}

// GetExportSideDataAddr gets `AVCodecContext.export_side_data` address.
func (avctx *AVCodecContext) GetExportSideDataAddr() *int32 {
	return (*int32)(&avctx.export_side_data)
}

// typedef int (*avcodec_context_get_encode_buffer_func)(struct AVCodecContext *s,
// AvPacket *pkt, int flags);
type AVCodecContextGetEncodeBufferFunc = C.avcodec_context_get_encode_buffer_func

// GetGetEncodeBuffer gets `AVCodecContext.get_encode_buffer` value.
func (avctx *AVCodecContext) GetGetEncodeBuffer() AVCodecContextGetEncodeBufferFunc {
	return (AVCodecContextGetEncodeBufferFunc)(avctx.get_encode_buffer)
}

// SetGetEncodeBuffer sets `AVCodecContext.get_encode_buffer` value.
func (avctx *AVCodecContext) SetGetEncodeBuffer(v AVCodecContextGetEncodeBufferFunc) {
	avctx.get_encode_buffer = (C.avcodec_context_get_encode_buffer_func)(v)
}

// GetGetEncodeBufferAddr gets `AVCodecContext.get_encode_buffer` address.
func (avctx *AVCodecContext) GetGetEncodeBufferAddr() *AVCodecContextGetEncodeBufferFunc {
	return (*AVCodecContextGetEncodeBufferFunc)(&avctx.get_encode_buffer)
}

// Deprecated: No use.
//
// AvCodecGetPktTimebase
func AvCodecGetPktTimebase(avctx *AVCodecContext) AVRational {
	return AVRational(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use.
//
// AvCodecSetPktTimebase
func AvCodecSetPktTimebase(avctx *AVCodecContext, r AVRational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(avctx), (C.struct_AVRational)(r))
}

// Deprecated: No use.
//
// AvCodecGetCodecDescriptor
func AvCodecGetCodecDescriptor(avctx *AVCodecContext) *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use.
//
// AvCodecSetCodecDescriptor
func AvCodecSetCodecDescriptor(avctx *AVCodecContext, d *AVCodecDescriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(avctx), (*C.struct_AVCodecDescriptor)(d))
}

// Deprecated: No use.
//
// AvCodecGetLowres
func AvCodecGetLowres(avctx *AVCodecContext) int32 {
	return (int32)(C.av_codec_get_lowres((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use.
//
// AvCodecSetLowres
func AvCodecSetLowres(avctx *AVCodecContext, i int32) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(avctx), (C.int)(i))
}

// Deprecated: No use.
//
// AvCodecGetSeekPreroll
func AvCodecGetSeekPreroll(avctx *AVCodecContext) int32 {
	return (int32)(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use.
//
// AvCodecSetSeekPreroll
func AvCodecSetSeekPreroll(avctx *AVCodecContext, i int32) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(avctx), (C.int)(i))
}

// Deprecated: No use.
//
// AvCodecGetChromaIntraMatrix
func AvCodecGetChromaIntraMatrix(avctx *AVCodecContext) *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use.
//
// AvCodecSetChromaIntraMatrix
func AvCodecSetChromaIntraMatrix(avctx *AVCodecContext, t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(avctx), (*C.uint16_t)(t))
}

// Deprecated: No use.
//
// AvCodecGetMaxLowres
func AvCodecGetMaxLowres(c *AVCodec) int32 {
	return (int32)(C.av_codec_get_max_lowres((*C.struct_AVCodec)(c)))
}

// MpegEncContext
type MpegEncContext C.struct_MpegEncContext

// AVHWAccel
type AVHWAccel C.struct_AVHWAccel

// GetName gets `AVHWAccel.name` value.
func (hwa *AVHWAccel) GetName() string {
	return C.GoString(hwa.name)
}

// GetType gets `AVHWAccel.type` value.
func (hwa *AVHWAccel) GetType() AVMediaType {
	return (AVMediaType)(hwa._type)
}

// SetType sets `AVHWAccel.type` value.
func (hwa *AVHWAccel) SetType(v AVMediaType) {
	hwa._type = (C.enum_AVMediaType)(v)
}

// GetTypeAddr gets `AVHWAccel.type` address.
func (hwa *AVHWAccel) GetTypeAddr() *AVMediaType {
	return (*AVMediaType)(&hwa._type)
}

// GetId gets `AVHWAccel.id` value.
func (hwa *AVHWAccel) GetId() AVCodecID {
	return (AVCodecID)(hwa.id)
}

// SetId sets `AVHWAccel.id` value.
func (hwa *AVHWAccel) SetId(v AVCodecID) {
	hwa.id = (C.enum_AVCodecID)(v)
}

// GetIdAddr gets `AVHWAccel.id` address.
func (hwa *AVHWAccel) GetIdAddr() *AVCodecID {
	return (*AVCodecID)(&hwa.id)
}

// GetPixFmt gets `AVHWAccel.pix_fmt` value.
func (hwa *AVHWAccel) GetPixFmt() AVPixelFormat {
	return (AVPixelFormat)(hwa.pix_fmt)
}

// SetPixFmt sets `AVHWAccel.pix_fmt` value.
func (hwa *AVHWAccel) SetPixFmt(v AVPixelFormat) {
	hwa.pix_fmt = (C.enum_AVPixelFormat)(v)
}

// GetPixFmtAddr gets `AVHWAccel.pix_fmt` address.
func (hwa *AVHWAccel) GetPixFmtAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&hwa.pix_fmt)
}

// GetCapabilities gets `AVHWAccel.capabilities` value.
func (hwa *AVHWAccel) GetCapabilities() int32 {
	return (int32)(hwa.capabilities)
}

// SetCapabilities sets `AVHWAccel.capabilities` value.
func (hwa *AVHWAccel) SetCapabilities(v int32) {
	hwa.capabilities = (C.int)(v)
}

// GetCapabilitiesAddr gets `AVHWAccel.capabilities` address.
func (hwa *AVHWAccel) GetCapabilitiesAddr() *int32 {
	return (*int32)(&hwa.capabilities)
}

const (
	AV_HWACCEL_CODEC_CAP_EXPERIMENTAL      = int(C.AV_HWACCEL_CODEC_CAP_EXPERIMENTAL)
	AV_HWACCEL_FLAG_IGNORE_LEVEL           = int(C.AV_HWACCEL_FLAG_IGNORE_LEVEL)
	AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH       = int(C.AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH)
	AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH = int(C.AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH)
)

// AVPicture
type AVPicture C.struct_AVPicture

// Deprecated: No use.
//
// GetData gets `AVPicture.data` value.
func (pct *AVPicture) GetData() []*uint8 {
	return unsafe.Slice((**uint8)(unsafe.Pointer(&pct.data[0])), AV_NUM_DATA_POINTERS)
}

// Deprecated: No use.
//
// SetData sets `AVPicture.data` value.
func (pct *AVPicture) SetData(v []*uint8) {
	for i := 0; i < FFMIN(len(v), AV_NUM_DATA_POINTERS); i++ {
		pct.data[i] = (*C.uint8_t)(v[i])
	}
}

// Deprecated: No use.
//
// GetDataAddr gets `AVPicture.data` address.
func (pct *AVPicture) GetDataAddr() ***uint8 {
	return (***uint8)(unsafe.Pointer(&pct.data))
}

// Deprecated: No use.
//
// GetLinesize gets `AVPicture.linesize` value.
func (pct *AVPicture) GetLinesize() []int32 {
	return unsafe.Slice((*int32)(&pct.linesize[0]), AV_NUM_DATA_POINTERS)
}

// Deprecated: No use.
//
// SetLinesize sets `AVPicture.linesize` value.
func (pct *AVPicture) SetLinesize(v []int32) {
	for i := 0; i < FFMIN(len(v), AV_NUM_DATA_POINTERS); i++ {
		pct.linesize[i] = (C.int)(v[i])
	}
}

// Deprecated: No use.
//
// GetLinesizeAddr gets `AVPicture.linesize` address.
func (pct *AVPicture) GetLinesizeAddr() **int32 {
	return (**int32)(unsafe.Pointer(&pct.linesize))
}

// AVSubtitleType
type AVSubtitleType = C.enum_AVSubtitleType

const (
	SUBTITLE_NONE   = AVSubtitleType(C.SUBTITLE_NONE)
	SUBTITLE_BITMAP = AVSubtitleType(C.SUBTITLE_BITMAP)
	SUBTITLE_TEXT   = AVSubtitleType(C.SUBTITLE_TEXT)
	SUBTITLE_ASS    = AVSubtitleType(C.SUBTITLE_ASS)
)

const AV_SUBTITLE_FLAG_FORCED = C.AV_SUBTITLE_FLAG_FORCED

// AVSubtitleRect
type AVSubtitleRect C.struct_AVSubtitleRect

// GetX gets `AVSubtitleRect.x` value.
func (sbtr *AVSubtitleRect) GetX() int32 {
	return (int32)(sbtr.x)
}

// SetX sets `AVSubtitleRect.x` value.
func (sbtr *AVSubtitleRect) SetX(v int32) {
	sbtr.x = (C.int)(v)
}

// GetXAddr gets `AVSubtitleRect.x` address.
func (sbtr *AVSubtitleRect) GetXAddr() *int32 {
	return (*int32)(&sbtr.x)
}

// GetY gets `AVSubtitleRect.y` value.
func (sbtr *AVSubtitleRect) GetY() int32 {
	return (int32)(sbtr.y)
}

// SetY sets `AVSubtitleRect.y` value.
func (sbtr *AVSubtitleRect) SetY(v int32) {
	sbtr.y = (C.int)(v)
}

// GetYAddr gets `AVSubtitleRect.y` address.
func (sbtr *AVSubtitleRect) GetYAddr() *int32 {
	return (*int32)(&sbtr.y)
}

// GetW gets `AVSubtitleRect.w` value.
func (sbtr *AVSubtitleRect) GetW() int32 {
	return (int32)(sbtr.w)
}

// SetW sets `AVSubtitleRect.w` value.
func (sbtr *AVSubtitleRect) SetW(v int32) {
	sbtr.w = (C.int)(v)
}

// GetWAddr gets `AVSubtitleRect.w` address.
func (sbtr *AVSubtitleRect) GetWAddr() *int32 {
	return (*int32)(&sbtr.w)
}

// GetH gets `AVSubtitleRect.h` value.
func (sbtr *AVSubtitleRect) GetH() int32 {
	return (int32)(sbtr.h)
}

// SetH sets `AVSubtitleRect.h` value.
func (sbtr *AVSubtitleRect) SetH(v int32) {
	sbtr.h = (C.int)(v)
}

// GetHAddr gets `AVSubtitleRect.h` address.
func (sbtr *AVSubtitleRect) GetHAddr() *int32 {
	return (*int32)(&sbtr.h)
}

// GetNbColors gets `AVSubtitleRect.nb_colors` value.
func (sbtr *AVSubtitleRect) GetNbColors() int32 {
	return (int32)(sbtr.nb_colors)
}

// SetNbColors sets `AVSubtitleRect.nb_colors` value.
func (sbtr *AVSubtitleRect) SetNbColors(v int32) {
	sbtr.nb_colors = (C.int)(v)
}

// GetNbColorsAddr gets `AVSubtitleRect.nb_colors` address.
func (sbtr *AVSubtitleRect) GetNbColorsAddr() *int32 {
	return (*int32)(&sbtr.nb_colors)
}

// Deprecated: Unused.
//
// GetPict gets `AVSubtitleRect.pict` value.
func (sbtr *AVSubtitleRect) GetPict() AVPicture {
	return (AVPicture)(sbtr.pict)
}

// Deprecated: Unused.
//
// SetPict sets `AVSubtitleRect.pict` value.
func (sbtr *AVSubtitleRect) SetPict(v AVPicture) {
	sbtr.pict = (C.struct_AVPicture)(v)
}

// Deprecated: Unused.
//
// GetPictAddr gets `AVSubtitleRect.pict` address.
func (sbtr *AVSubtitleRect) GetPictAddr() *AVPicture {
	return (*AVPicture)(&sbtr.pict)
}

// GetData gets `AVSubtitleRect.data` value.
func (sbtr *AVSubtitleRect) GetData() []*uint8 {
	return unsafe.Slice((**uint8)(unsafe.Pointer(&sbtr.data[0])), 4)
}

// SetData sets `AVSubtitleRect.data` value.
func (sbtr *AVSubtitleRect) SetData(v []*uint8) {
	for i := 0; i < FFMIN(len(v), 4); i++ {
		sbtr.data[i] = (*C.uint8_t)(v[i])
	}
}

// GetDataAddr gets `AVSubtitleRect.data` address.
func (sbtr *AVSubtitleRect) GetDataAddr() ***uint8 {
	return (***uint8)(unsafe.Pointer(&sbtr.data))
}

// GetLinesize gets `AVSubtitleRect.linesize` value.
func (sbtr *AVSubtitleRect) GetLinesize() []int32 {
	return unsafe.Slice((*int32)(&sbtr.linesize[0]), 4)
}

// SetLinesize sets `AVSubtitleRect.linesize` value.
func (sbtr *AVSubtitleRect) SetLinesize(v []int32) {
	for i := 0; i < FFMIN(len(v), 4); i++ {
		sbtr.linesize[i] = (C.int)(v[i])
	}
}

// GetLinesize gets `AVSubtitleRect.linesize` address.
func (sbtr *AVSubtitleRect) GetLinesizeAddr() **int32 {
	return (**int32)(unsafe.Pointer(&sbtr.linesize))
}

// GetType gets `AVSubtitleRect.type` value.
func (sbtr *AVSubtitleRect) GetType() AVSubtitleType {
	return (AVSubtitleType)(sbtr._type)
}

// SetType sets `AVSubtitleRect.type` value.
func (sbtr *AVSubtitleRect) SetType(v AVSubtitleType) {
	sbtr._type = (C.enum_AVSubtitleType)(v)
}

// GetTypeAddr gets `AVSubtitleRect.type` address.
func (sbtr *AVSubtitleRect) GetTypeAddr() *AVSubtitleType {
	return (*AVSubtitleType)(&sbtr._type)
}

// GetText gets `AVSubtitleRect.text` value.
func (sbtr *AVSubtitleRect) GetText() string {
	return C.GoString(sbtr.text)
}

// GetAss gets `AVSubtitleRect.ass` value.
func (sbtr *AVSubtitleRect) GetAss() string {
	return C.GoString(sbtr.ass)
}

// GetFlags gets `AVSubtitleRect.flags` value.
func (sbtr *AVSubtitleRect) GetFlags() int32 {
	return (int32)(sbtr.flags)
}

// SetFlags sets `AVSubtitleRect.flags` value.
func (sbtr *AVSubtitleRect) SetFlags(v int32) {
	sbtr.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVSubtitleRect.flags` address.
func (sbtr *AVSubtitleRect) GetFlagsAddr() *int32 {
	return (*int32)(&sbtr.flags)
}

// AVSubtitle
type AVSubtitle C.struct_AVSubtitle

// GetFormat gets `AVSubtitle.format` value.
func (sbt *AVSubtitle) GetFormat() uint16 {
	return (uint16)(sbt.format)
}

// SetFormat sets `AVSubtitle.format` value.
func (sbt *AVSubtitle) SetFormat(v uint16) {
	sbt.format = (C.uint16_t)(v)
}

// GetFormatAddr gets `AVSubtitle.format` address.
func (sbt *AVSubtitle) GetFormatAddr() *uint16 {
	return (*uint16)(&sbt.format)
}

// GetStartDisplayTime gets `AVSubtitle.start_display_time` value.
func (sbt *AVSubtitle) GetStartDisplayTime() uint32 {
	return (uint32)(sbt.start_display_time)
}

// SetStartDisplayTime sets `AVSubtitle.start_display_time` value.
func (sbt *AVSubtitle) SetStartDisplayTime(v uint32) {
	sbt.start_display_time = (C.uint32_t)(v)
}

// GetStartDisplayTimeAddr gets `AVSubtitle.start_display_time` address.
func (sbt *AVSubtitle) GetStartDisplayTimeAddr() *uint32 {
	return (*uint32)(&sbt.start_display_time)
}

// GetEndDisplayTime gets `AVSubtitle.end_display_time` value.
func (sbt *AVSubtitle) GetEndDisplayTime() uint32 {
	return (uint32)(sbt.end_display_time)
}

// SetEndDisplayTime sets `AVSubtitle.end_display_time` value.
func (sbt *AVSubtitle) SetEndDisplayTime(v uint32) {
	sbt.end_display_time = (C.uint32_t)(v)
}

// GetEndDisplayTimeAddr gets `AVSubtitle.end_display_time` address.
func (sbt *AVSubtitle) GetEndDisplayTimeAddr() *uint32 {
	return (*uint32)(&sbt.end_display_time)
}

// GetNumRects gets `AVSubtitle.num_rects` value.
func (sbt *AVSubtitle) GetNumRects() uint32 {
	return (uint32)(sbt.num_rects)
}

// SetNumRects sets `AVSubtitle.num_rects` value.
func (sbt *AVSubtitle) SetNumRects(v uint32) {
	sbt.num_rects = (C.uint)(v)
}

// GetNumRectsAddr gets `AVSubtitle.num_rects` address.
func (sbt *AVSubtitle) GetNumRectsAddr() *uint32 {
	return (*uint32)(&sbt.num_rects)
}

// GetRects gets `AVSubtitle.rects` value.
func (sbt *AVSubtitle) GetRects() []*AVSubtitleRect {
	if sbt.rects == nil {
		return nil
	}
	return unsafe.Slice((**AVSubtitleRect)(unsafe.Pointer(sbt.rects)), sbt.num_rects)
}

// SetRects sets `AVSubtitle.rects` value.
func (sbt *AVSubtitle) SetRects(v **AVSubtitleRect) {
	sbt.rects = (**C.struct_AVSubtitleRect)(unsafe.Pointer(v))
}

// GetRectsAddr gets `AVSubtitle.rects` address.
func (sbt *AVSubtitle) GetRectsAddr() ***AVSubtitleRect {
	return (***AVSubtitleRect)(unsafe.Pointer(&sbt.rects))
}

// GetPts gets `AVSubtitle.pts` value.
func (sbt *AVSubtitle) GetPts() int64 {
	return (int64)(sbt.pts)
}

// SetPts sets `AVSubtitle.pts` value.
func (sbt *AVSubtitle) SetPts(v int64) {
	sbt.pts = (C.int64_t)(v)
}

// GetPtsAddr gets `AVSubtitle.pts` address.
func (sbt *AVSubtitle) GetPtsAddr() *int64 {
	return (*int64)(&sbt.pts)
}

// Deprecated: No use.
//
// AvCodecNext returns the first registered codec if c is NULL,
// returns the next registered codec after c if c is non-NULL,
// or NULL if c is the last one.
func AvCodecNext(c *AVCodec) *AVCodec {
	return (*AVCodec)(C.av_codec_next((*C.struct_AVCodec)(c)))
}

// AvCodecVersion returns the LIBAVCODEC_VERSION_INT constant.
func AvCodecVersion() uint32 {
	return (uint32)(C.avcodec_version())
}

// AvCodecConfiguration returns the libavcodec build-time configuration.
func AvCodecConfiguration() string {
	return C.GoString(C.avcodec_configuration())
}

// Deprecated: Calling this function is unnecessary.
//
// AvCodecRegister
func AvCodecRegister(c *AVCodec) {
	C.avcodec_register((*C.struct_AVCodec)(c))
}

// Deprecated: Calling this function is unnecessary.
//
// AvCodecRegisterAll
func AvCodecRegisterAll() {
	C.avcodec_register_all()
}

// AvCodecAllocContext3 allocates an AVCodecContext and set its fields to default values.
// The resulting struct should be freed with AVCodecFreeContext().
func AvCodecAllocContext3(c *AVCodec) *AVCodecContext {
	return (*AVCodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
}

// AvCodecFreeContext frees the codec context and everything associated with it
// and write NULL to the provided pointer.
func AvCodecFreeContext(avctx **AVCodecContext) {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(avctx)))
}

// Deprecated: No use.
func AvCodecGetContextDefaults3(avctx *AVCodecContext, c *AVCodec) int32 {
	return (int32)(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVCodec)(c)))
}

// AvCodecGetClass gets the AVClass for AVCodecContext.
func AvCodecGetClass() *AVClass {
	return (*AVClass)(C.avcodec_get_class())
}

// Deprecated: This function should not be used.
//
// AvCodecGetFrameClass
func AvCodecGetFrameClass() *AVClass {
	return (*AVClass)(C.avcodec_get_frame_class())
}

// AvCodecGetSubtitleRectClass gets the AVClass for AVSubtitleRect.
func AvCodecGetSubtitleRectClass() *AVClass {
	return (*AVClass)(C.avcodec_get_subtitle_rect_class())
}

// Deprecated: Use an intermediate AVCodecParameters instance and the
// AvCodecParametersFromContext() / AVCodecParametersToContext() functions.
//
// AvCodecCopyContext copies the settings of the source AVCodecContext into the destination
// AVCodecContext.
func AvCodecCopyContext(dest, src *AVCodecContext) int32 {
	return (int32)(C.avcodec_copy_context((*C.struct_AVCodecContext)(dest),
		(*C.struct_AVCodecContext)(src)))
}

// AvCodecParametersFromContext fills the parameters struct based on the values from the supplied codec
// context. Any allocated fields in par are freed and replaced with duplicates
// of the corresponding fields in codec.
func AvCodecParametersFromContext(par *AVCodecParameters, avctx *AVCodecContext) int32 {
	return (int32)(C.avcodec_parameters_from_context((*C.struct_AVCodecParameters)(par),
		(*C.struct_AVCodecContext)(avctx)))
}

// AvCodecParametersToContext fills the codec context based on the values from the supplied codec
// parameters. Any allocated fields in codec that have a corresponding field in
// par are freed and replaced with duplicates of the corresponding field in par.
// Fields in codec that do not have a counterpart in par are not touched.
func AvCodecParametersToContext(avctx *AVCodecContext, par *AVCodecParameters) int32 {
	return (int32)(C.avcodec_parameters_to_context((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVCodecParameters)(par)))
}

// AvCodecContext initializes the context to use the given codec.
func AvCodecOpen2(avctx *AVCodecContext, c *AVCodec, d **AVDictionary) int32 {
	return (int32)(C.avcodec_open2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

// AvCodecClose closes a given context and free all the data associated with it (but not the context itself).
func AvCodecClose(avctx *AVCodecContext) int32 {
	return (int32)(C.avcodec_close((*C.struct_AVCodecContext)(avctx)))
}

// AvSubtitleFree frees all allocated data in the given subtitle struct.
func AvSubtitleFree(s *AVSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

// The default callback for AVCodecContext.get_buffer2().
func AvCodecDefaultGetBuffer2(avctx *AVCodecContext, frame *AVFrame, flags int32) int32 {
	return (int32)(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(frame), (C.int)(flags)))
}

// The default callback for AVCodecContext.get_encode_buffer().
func AvCodecDefaultGetEncodeBuffer(avctx *AVCodecContext, pkt *AVPacket, flags int32) int32 {
	return (int32)(C.avcodec_default_get_encode_buffer((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(pkt), (C.int)(flags)))
}

// AvCodecAlignDimensions modifies width and height values so that they will result in a memory
// buffer that is acceptable for the codec if you do not use any horizontal padding.
func AvCodecAlignDimensions(avctx *AVCodecContext, width, height *int32) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(avctx), (*C.int)(width), (*C.int)(height))
}

// AvCodecAlignDimensions2 modifies width and height values so that they will result in a memory
// buffer that is acceptable for the codec if you also ensure that all
// line sizes are a multiple of the respective linesize_align[i].
func AvCodecAlignDimensions2(avctx *AVCodecContext, width, height *int32,
	linesizeAlign [AV_NUM_DATA_POINTERS]int32) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(avctx),
		(*C.int)(width), (*C.int)(height), (*C.int)(unsafe.Pointer(&linesizeAlign[0])))
}

// AvCodecEnumToChromaPos converts AVChromaLocation to swscale x/y chroma position.
func AvCodecEnumToChromaPos(xpos, ypos *int32, pos AVChromaLocation) int32 {
	return (int32)(C.avcodec_enum_to_chroma_pos((*C.int)(xpos), (*C.int)(ypos), (C.enum_AVChromaLocation)(pos)))
}

// AvCodecChromaPosToEnum converts swscale x/y chroma position to AVChromaLocation.
func AvCodecChromaPosToEnum(xpos, ypos int32) AVChromaLocation {
	return (AVChromaLocation)(C.avcodec_chroma_pos_to_enum((C.int)(xpos), (C.int)(ypos)))
}

// Deprecated: Use AVCodecSendPacket() and AVCodecReceiveFrame().
//
// AvCodecDecodeAudio4 decodes the audio frame of size avpkt->size from avpkt->data into frame.
func AvCodecDecodeAudio4(avctx *AVCodecContext, frame *AVFrame, gotFramePtr *int32, avpkt *AVPacket) int32 {
	return (int32)(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(frame), (*C.int)(gotFramePtr), (*C.struct_AVPacket)(avpkt)))
}

// Deprecated: Use AVCodecSendPacket() and AVCodecReceiveFrame().
//
// AvCodecDecodeVideo2 decodes the video frame of size avpkt->size from avpkt->data into picture.
func AvCodecDecodeVideo2(avctx *AVCodecContext, picture *AVFrame, gotPicturePtr *int32, avpkt *AVPacket) int32 {
	return (int32)(C.avcodec_decode_video2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(picture), (*C.int)(gotPicturePtr), (*C.struct_AVPacket)(avpkt)))
}

// AvCodecDecodeSubtitle2 decodes a subtitle message.
func AvCodecDecodeSubtitle2(avctx *AVCodecContext, sub *AVSubtitle, gotSubPtr *int32, avpkt *AVPacket) int32 {
	return (int32)(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVSubtitle)(sub), (*C.int)(gotSubPtr), (*C.struct_AVPacket)(avpkt)))
}

// AvCodecSendPacket supplies raw packet data as input to a decoder.
func AvCodecSendPacket(avctx *AVCodecContext, avpkt *AVPacket) int32 {
	return (int32)(C.avcodec_send_packet((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(avpkt)))
}

// AvCodecReceiveFrame returns decoded output data from a decoder.
func AvCodecReceiveFrame(avctx *AVCodecContext, frame *AVFrame) int32 {
	return (int32)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(frame)))
}

// AvCodecSendFrame supplies a raw video or audio frame to the encoder. Use AVCodecReceivePacket()
// to retrieve buffered output packets.
func AvCodecSendFrame(avctx *AVCodecContext, frame *AVFrame) int32 {
	return (int32)(C.avcodec_send_frame((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(frame)))
}

// AvCodecReceivePacket reads encoded data from the encoder.
func AvCodecReceivePacket(avctx *AVCodecContext, avpkt *AVPacket) int32 {
	return (int32)(C.avcodec_receive_packet((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(avpkt)))
}

// AvCodecGetHwFramesParameters create and return a AVHWFramesContext with values adequate for hardware
// decoding.
func AvCodecGetHwFramesParameters(avctx *AVCodecContext, deviceRef *AVBufferRef,
	hwPixFmt AVPixelFormat,
	outFramesRef **AVBufferRef) int32 {
	return (int32)(C.avcodec_get_hw_frames_parameters((*C.struct_AVCodecContext)(avctx),
		(*C.AVBufferRef)(deviceRef),
		(C.enum_AVPixelFormat)(hwPixFmt),
		(**C.AVBufferRef)(unsafe.Pointer(outFramesRef))))
}

// AVPictureStructure
type AVPictureStructure = C.enum_AVPictureStructure

const (
	AV_PICTURE_STRUCTURE_UNKNOWN      = AVPictureStructure(C.AV_PICTURE_STRUCTURE_UNKNOWN)
	AV_PICTURE_STRUCTURE_TOP_FIELD    = AVPictureStructure(C.AV_PICTURE_STRUCTURE_TOP_FIELD)
	AV_PICTURE_STRUCTURE_BOTTOM_FIELD = AVPictureStructure(C.AV_PICTURE_STRUCTURE_BOTTOM_FIELD)
	AV_PICTURE_STRUCTURE_FRAME        = AVPictureStructure(C.AV_PICTURE_STRUCTURE_FRAME)
)

// AVCodecParserContext
type AVCodecParserContext C.struct_AVCodecParserContext

// GetPrivData gets `AVCodecParserContext.priv_data` value.
func (cpc *AVCodecParserContext) GetPrivData() unsafe.Pointer {
	return cpc.priv_data
}

// SetPrivData sets `AVCodecParserContext.priv_data` value.
func (cpc *AVCodecParserContext) SetPrivData(v CVoidPointer) {
	cpc.priv_data = VoidPointer(v)
}

// GetPrivDataAddr gets `AVCodecParserContext.priv_data` address.
func (cpc *AVCodecParserContext) GetPrivDataAddr() *unsafe.Pointer {
	return (*unsafe.Pointer)(&cpc.priv_data)
}

// GetParser gets `AVCodecParserContext.parser` value.
func (cpc *AVCodecParserContext) GetParser() *AVCodecParser {
	return (*AVCodecParser)(cpc.parser)
}

// SetParser sets `AVCodecParserContext.parser` value.
func (cpc *AVCodecParserContext) SetParser(v *AVCodecParser) {
	cpc.parser = (*C.struct_AVCodecParser)(v)
}

// GetParserAddr gets `AVCodecParserContext.parser` address.
func (cpc *AVCodecParserContext) GetParserAddr() **AVCodecParser {
	return (**AVCodecParser)(unsafe.Pointer(&cpc.parser))
}

// GetFrameOffset gets `AVCodecParserContext.frame_offset` value.
func (cpc *AVCodecParserContext) GetFrameOffset() int64 {
	return (int64)(cpc.frame_offset)
}

// SetFrameOffset sets `AVCodecParserContext.frame_offset` value.
func (cpc *AVCodecParserContext) SetFrameOffset(v int64) {
	cpc.frame_offset = (C.int64_t)(v)
}

// GetFrameOffsetAddr gets `AVCodecParserContext.frame_offset` address.
func (cpc *AVCodecParserContext) GetFrameOffsetAddr() *int64 {
	return (*int64)(&cpc.frame_offset)
}

// GetCurOffset gets `AVCodecParserContext.cur_offset` value.
func (cpc *AVCodecParserContext) GetCurOffset() int64 {
	return (int64)(cpc.cur_offset)
}

// SetCurOffset sets `AVCodecParserContext.cur_offset` value.
func (cpc *AVCodecParserContext) SetCurOffset(v int64) {
	cpc.cur_offset = (C.int64_t)(v)
}

// GetCurOffsetAddr gets `AVCodecParserContext.cur_offset` address.
func (cpc *AVCodecParserContext) GetCurOffsetAddr() *int64 {
	return (*int64)(&cpc.cur_offset)
}

// GetNextFrameOffset gets `AVCodecParserContext.next_frame_offset` value.
func (cpc *AVCodecParserContext) GetNextFrameOffset() int64 {
	return (int64)(cpc.next_frame_offset)
}

// SetNextFrameOffset sets `AVCodecParserContext.next_frame_offset` value.
func (cpc *AVCodecParserContext) SetNextFrameOffset(v int64) {
	cpc.next_frame_offset = (C.int64_t)(v)
}

// GetNextFrameOffsetAddr gets `AVCodecParserContext.next_frame_offset` address.
func (cpc *AVCodecParserContext) GetNextFrameOffsetAddr() *int64 {
	return (*int64)(&cpc.next_frame_offset)
}

// GetPictType gets `AVCodecParserContext.picttype` value.
func (cpc *AVCodecParserContext) GetPictType() int32 {
	return (int32)(cpc.pict_type)
}

// SetPictType sets `AVCodecParserContext.picttype` value.
func (cpc *AVCodecParserContext) SetPictType(v int32) {
	cpc.pict_type = (C.int)(v)
}

// GetPictTypeAddr gets `AVCodecParserContext.picttype` address.
func (cpc *AVCodecParserContext) GetPictTypeAddr() *int32 {
	return (*int32)(&cpc.pict_type)
}

// GetPts gets `AVCodecParserContext.pts` value.
func (cpc *AVCodecParserContext) GetPts() int64 {
	return (int64)(cpc.pts)
}

// SetPts sets `AVCodecParserContext.pts` value.
func (cpc *AVCodecParserContext) SetPts(v int64) {
	cpc.pts = (C.int64_t)(v)
}

// GetPtsAddr gets `AVCodecParserContext.pts` address.
func (cpc *AVCodecParserContext) GetPtsAddr() *int64 {
	return (*int64)(&cpc.pts)
}

// GetDts gets `AVCodecParserContext.dts` value.
func (cpc *AVCodecParserContext) GetDts() int64 {
	return (int64)(cpc.dts)
}

// SetDts sets `AVCodecParserContext.dts` value.
func (cpc *AVCodecParserContext) SetDts(v int64) {
	cpc.dts = (C.int64_t)(v)
}

// GetDtsAddr gets `AVCodecParserContext.dts` address.
func (cpc *AVCodecParserContext) GetDtsAddr() *int64 {
	return (*int64)(&cpc.dts)
}

// GetLastPts gets `AVCodecParserContext.last_pts` value.
func (cpc *AVCodecParserContext) GetLastPts() int64 {
	return (int64)(cpc.last_pts)
}

// SetLastPts sets `AVCodecParserContext.last_pts` value.
func (cpc *AVCodecParserContext) SetLastPts(v int64) {
	cpc.last_pts = (C.int64_t)(v)
}

// GetLastPtsAddr gets `AVCodecParserContext.last_pts` address.
func (cpc *AVCodecParserContext) GetLastPtsAddr() *int64 {
	return (*int64)(&cpc.last_pts)
}

// GetLastDts gets `AVCodecParserContext.last_dts` value.
func (cpc *AVCodecParserContext) GetLastDts() int64 {
	return (int64)(cpc.last_dts)
}

// SetLastDts sets `AVCodecParserContext.last_dts` value.
func (cpc *AVCodecParserContext) SetLastDts(v int64) {
	cpc.last_dts = (C.int64_t)(v)
}

// GetLastDtsAddr gets `AVCodecParserContext.last_dts` address.
func (cpc *AVCodecParserContext) GetLastDtsAddr() *int64 {
	return (*int64)(&cpc.last_dts)
}

// GetFetchTimestamp gets `AVCodecParserContext.fetch_timestamp` value.
func (cpc *AVCodecParserContext) GetFetchTimestamp() int32 {
	return (int32)(cpc.fetch_timestamp)
}

// SetFetchTimestamp sets `AVCodecParserContext.fetch_timestamp` value.
func (cpc *AVCodecParserContext) SetFetchTimestamp(v int32) {
	cpc.fetch_timestamp = (C.int)(v)
}

// GetFetchTimestampAddr gets `AVCodecParserContext.fetch_timestamp` address.
func (cpc *AVCodecParserContext) GetFetchTimestampAddr() *int32 {
	return (*int32)(&cpc.fetch_timestamp)
}

const (
	AV_PARSER_PTS_NB = C.AV_PARSER_PTS_NB
)

// GetCurFrameStartIndex gets `AVCodecParserContext.cur_frame_start_index` value.
func (cpc *AVCodecParserContext) GetCurFrameStartIndex() int32 {
	return (int32)(cpc.cur_frame_start_index)
}

// SetCurFrameStartIndex sets `AVCodecParserContext.cur_frame_start_index` value.
func (cpc *AVCodecParserContext) SetCurFrameStartIndex(v int32) {
	cpc.cur_frame_start_index = (C.int)(v)
}

// GetCurFrameStartIndexAddr gets `AVCodecParserContext.cur_frame_start_index` address.
func (cpc *AVCodecParserContext) GetCurFrameStartIndexAddr() *int32 {
	return (*int32)(&cpc.cur_frame_start_index)
}

// GetCurFrameOffset gets `AVCodecParserContext.cur_frame_offset` value.
func (cpc *AVCodecParserContext) GetCurFrameOffset() []int64 {
	return unsafe.Slice((*int64)(&cpc.cur_frame_offset[0]), AV_PARSER_PTS_NB)
}

// SetCurFrameOffset sets `AVCodecParserContext.cur_frame_offset` value.
func (cpc *AVCodecParserContext) SetCurFrameOffset(v []int64) {
	for i := 0; i < FFMIN(len(v), AV_PARSER_PTS_NB); i++ {
		cpc.cur_frame_offset[i] = (C.int64_t)(v[i])
	}
}

// GetCurFrameOffsetAddr gets `AVCodecParserContext.cur_frame_offset` address.
func (cpc *AVCodecParserContext) GetCurFrameOffsetAddr() **int64 {
	return (**int64)(unsafe.Pointer(&cpc.cur_frame_offset))
}

// GetCurFramePts gets `AVCodecParserContext.cur_frame_pts` value.
func (cpc *AVCodecParserContext) GetCurFramePts() []int64 {
	return unsafe.Slice((*int64)(&cpc.cur_frame_pts[0]), AV_PARSER_PTS_NB)
}

// SetCurFramePts sets `AVCodecParserContext.cur_frame_pts` value.
func (cpc *AVCodecParserContext) SetCurFramePts(v []int64) {
	for i := 0; i < FFMIN(len(v), AV_PARSER_PTS_NB); i++ {
		cpc.cur_frame_pts[i] = (C.int64_t)(v[i])
	}
}

// GetCurFramePtsAddr gets `AVCodecParserContext.cur_frame_pts` address.
func (cpc *AVCodecParserContext) GetCurFramePtsAddr() **int64 {
	return (**int64)(unsafe.Pointer(&cpc.cur_frame_pts))
}

// GetCurFrameDts gets `AVCodecParserContext.cur_frame_dts` value.
func (cpc *AVCodecParserContext) GetCurFrameDts() []int64 {
	return unsafe.Slice((*int64)(&cpc.cur_frame_dts[0]), AV_PARSER_PTS_NB)
}

// SetCurFrameDts sets `AVCodecParserContext.cur_frame_dts` value.
func (cpc *AVCodecParserContext) SetCurFrameDts(v []int64) {
	for i := 0; i < FFMIN(len(v), AV_PARSER_PTS_NB); i++ {
		cpc.cur_frame_dts[i] = (C.int64_t)(v[i])
	}
}

// GetCurFrameDtsAddr gets `AVCodecParserContext.cur_frame_dts` address.
func (cpc *AVCodecParserContext) GetCurFrameDtsAddr() **int64 {
	return (**int64)(unsafe.Pointer(&cpc.cur_frame_dts))
}

// GetOffset gets `AVCodecParserContext.offset` value.
func (cpc *AVCodecParserContext) GetOffset() int64 {
	return (int64)(cpc.offset)
}

// SetOffset sets `AVCodecParserContext.offset` value.
func (cpc *AVCodecParserContext) SetOffset(v int64) {
	cpc.offset = (C.int64_t)(v)
}

// GetOffsetAddr gets `AVCodecParserContext.offset` address.
func (cpc *AVCodecParserContext) GetOffsetAddr() *int64 {
	return (*int64)(&cpc.offset)
}

// GetCurFrameEnd gets `AVCodecParserContext.cur_frame_end` value.
func (cpc *AVCodecParserContext) GetCurFrameEnd() []int64 {
	return unsafe.Slice((*int64)(&cpc.cur_frame_end[0]), AV_PARSER_PTS_NB)
}

// SetCurFrameEnd sets `AVCodecParserContext.cur_frame_end` value.
func (cpc *AVCodecParserContext) SetCurFrameEnd(v []int64) {
	for i := 0; i < FFMIN(len(v), AV_PARSER_PTS_NB); i++ {
		cpc.cur_frame_end[i] = (C.int64_t)(v[i])
	}
}

// GetCurFrameEndAddr gets `AVCodecParserContext.cur_frame_end` address.
func (cpc *AVCodecParserContext) GetCurFrameEndAddr() **int64 {
	return (**int64)(unsafe.Pointer(&cpc.cur_frame_end))
}

// GetKeyFrame gets `AVCodecParserContext.key_frame` value.
func (cpc *AVCodecParserContext) GetKeyFrame() int32 {
	return (int32)(cpc.key_frame)
}

// SetKeyFrame sets `AVCodecParserContext.key_frame` value.
func (cpc *AVCodecParserContext) SetKeyFrame(v int32) {
	cpc.key_frame = (C.int)(v)
}

// GetKeyFrameAddr gets `AVCodecParserContext.key_frame` address.
func (cpc *AVCodecParserContext) GetKeyFrameAddr() *int32 {
	return (*int32)(&cpc.key_frame)
}

// Deprecated: Unused.
//
// GetConvergenceDuration gets `AVCodecParserContext.convergence_duration` value.
func (cpc *AVCodecParserContext) GetConvergenceDuration() int64 {
	return (int64)(cpc.convergence_duration)
}

// Deprecated: Unused.
//
// SetConvergenceDuration sets `AVCodecParserContext.convergence_duration` value.
func (cpc *AVCodecParserContext) SetConvergenceDuration(v int64) {
	cpc.convergence_duration = (C.int64_t)(v)
}

// Deprecated: Unused.
//
// GetConvergenceDurationAddr gets `AVCodecParserContext.convergence_duration` address.
func (cpc *AVCodecParserContext) GetConvergenceDurationAddr() *int64 {
	return (*int64)(&cpc.convergence_duration)
}

// GetDtsSyncPoint gets `AVCodecParserContext.dts_sync_point` value.
func (cpc *AVCodecParserContext) GetDtsSyncPoint() int32 {
	return (int32)(cpc.dts_sync_point)
}

// SetDtsSyncPoint sets `AVCodecParserContext.dts_sync_point` value.
func (cpc *AVCodecParserContext) SetDtsSyncPoint(v int32) {
	cpc.dts_sync_point = (C.int)(v)
}

// GetDtsSyncPointAddr gets `AVCodecParserContext.dts_sync_point` address.
func (cpc *AVCodecParserContext) GetDtsSyncPointAddr() *int32 {
	return (*int32)(&cpc.dts_sync_point)
}

// GetDtsRefDtsDelta gets `AVCodecParserContext.dts_ref_dts_delta` value.
func (cpc *AVCodecParserContext) GetDtsRefDtsDelta() int32 {
	return (int32)(cpc.dts_ref_dts_delta)
}

// SetDtsRefDtsDelta sets `AVCodecParserContext.dts_ref_dts_delta` value.
func (cpc *AVCodecParserContext) SetDtsRefDtsDelta(v int32) {
	cpc.dts_ref_dts_delta = (C.int)(v)
}

// GetDtsRefDtsDeltaAddr gets `AVCodecParserContext.dts_ref_dts_delta` address.
func (cpc *AVCodecParserContext) GetDtsRefDtsDeltaAddr() *int32 {
	return (*int32)(&cpc.dts_ref_dts_delta)
}

// GetPtsDtsDelta gets `AVCodecParserContext.pts_dts_delta` value.
func (cpc *AVCodecParserContext) GetPtsDtsDelta() int32 {
	return (int32)(cpc.pts_dts_delta)
}

// SetPtsDtsDelta sets `AVCodecParserContext.pts_dts_delta` value.
func (cpc *AVCodecParserContext) SetPtsDtsDelta(v int32) {
	cpc.pts_dts_delta = (C.int)(v)
}

// GetPtsDtsDeltaAddr gets `AVCodecParserContext.pts_dts_delta` address.
func (cpc *AVCodecParserContext) GetPtsDtsDeltaAddr() *int32 {
	return (*int32)(&cpc.pts_dts_delta)
}

// GetCurFramePos gets `AVCodecParserContext.cur_frame_pos` value.
func (cpc *AVCodecParserContext) GetCurFramePos() []int64 {
	return unsafe.Slice((*int64)(&cpc.cur_frame_pos[0]), AV_PARSER_PTS_NB)
}

// SetCurFramePos sets `AVCodecParserContext.cur_frame_pos` value.
func (cpc *AVCodecParserContext) SetCurFramePos(v []int64) {
	for i := 0; i < FFMIN(len(v), AV_PARSER_PTS_NB); i++ {
		cpc.cur_frame_pos[i] = (C.int64_t)(v[i])
	}
}

// GetCurFramePosAddr gets `AVCodecParserContext.cur_frame_pos` address.
func (cpc *AVCodecParserContext) GetCurFramePosAddr() **int64 {
	return (**int64)(unsafe.Pointer(&cpc.cur_frame_pos))
}

// GetPos gets `AVCodecParserContext.pos` value.
func (cpc *AVCodecParserContext) GetPos() int64 {
	return (int64)(cpc.pos)
}

// SetPos sets `AVCodecParserContext.pos` value.
func (cpc *AVCodecParserContext) SetPos(v int64) {
	cpc.pos = (C.int64_t)(v)
}

// GetPosAddr gets `AVCodecParserContext.pos` address.
func (cpc *AVCodecParserContext) GetPosAddr() *int64 {
	return (*int64)(&cpc.pos)
}

// GetLastPos gets `AVCodecParserContext.last_pos` value.
func (cpc *AVCodecParserContext) GetLastPos() int64 {
	return (int64)(cpc.last_pos)
}

// SetLastPos sets `AVCodecParserContext.last_pos` value.
func (cpc *AVCodecParserContext) SetLastPos(v int64) {
	cpc.last_pos = (C.int64_t)(v)
}

// GetLastPosAddr gets `AVCodecParserContext.last_pos` address.
func (cpc *AVCodecParserContext) GetLastPosAddr() *int64 {
	return (*int64)(&cpc.last_pos)
}

// GetDuration gets `AVCodecParserContext.duration` value.
func (cpc *AVCodecParserContext) GetDuration() int32 {
	return (int32)(cpc.duration)
}

// SetDuration sets `AVCodecParserContext.duration` value.
func (cpc *AVCodecParserContext) SetDuration(v int32) {
	cpc.duration = (C.int)(v)
}

// GetDurationAddr gets `AVCodecParserContext.duration` address.
func (cpc *AVCodecParserContext) GetDurationAddr() *int32 {
	return (*int32)(&cpc.duration)
}

// GetFieldOrder gets `AVCodecParserContext.field_order` value.
func (cpc *AVCodecParserContext) GetFieldOrder() AVFieldOrder {
	return (AVFieldOrder)(cpc.field_order)
}

// SetFieldOrder sets `AVCodecParserContext.field_order` value.
func (cpc *AVCodecParserContext) SetFieldOrder(v AVFieldOrder) {
	cpc.field_order = (C.enum_AVFieldOrder)(v)
}

// GetFieldOrderAddr gets `AVCodecParserContext.field_order` address.
func (cpc *AVCodecParserContext) GetFieldOrderAddr() *AVFieldOrder {
	return (*AVFieldOrder)(&cpc.field_order)
}

// GetPictureStructure gets `AVCodecParserContext.picture_structure` value.
func (cpc *AVCodecParserContext) GetPictureStructure() AVPictureStructure {
	return (AVPictureStructure)(cpc.picture_structure)
}

// SetPictureStructure sets `AVCodecParserContext.picture_structure` value.
func (cpc *AVCodecParserContext) SetPictureStructure(v AVPictureStructure) {
	cpc.picture_structure = (C.enum_AVPictureStructure)(v)
}

// GetPictureStructureAddr gets `AVCodecParserContext.picture_structure` address.
func (cpc *AVCodecParserContext) GetPictureStructureAddr() *AVPictureStructure {
	return (*AVPictureStructure)(&cpc.picture_structure)
}

// GetOutputPictureNumber gets `AVCodecParserContext.output_picture_number` value.
func (cpc *AVCodecParserContext) GetOutputPictureNumber() int32 {
	return (int32)(cpc.output_picture_number)
}

// SetOutputPictureNumber sets `AVCodecParserContext.output_picture_number` value.
func (cpc *AVCodecParserContext) SetOutputPictureNumber(v int32) {
	cpc.output_picture_number = (C.int)(v)
}

// GetOutputPictureNumberAddr gets `AVCodecParserContext.output_picture_number` address.
func (cpc *AVCodecParserContext) GetOutputPictureNumberAddr() *int32 {
	return (*int32)(&cpc.output_picture_number)
}

// GetWidth gets `AVCodecParserContext.width` value.
func (cpc *AVCodecParserContext) GetWidth() int32 {
	return (int32)(cpc.width)
}

// SetWidth sets `AVCodecParserContext.width` value.
func (cpc *AVCodecParserContext) SetWidth(v int32) {
	cpc.width = (C.int)(v)
}

// GetWidthAddr gets `AVCodecParserContext.width` address.
func (cpc *AVCodecParserContext) GetWidthAddr() *int32 {
	return (*int32)(&cpc.width)
}

// GetHeight gets `AVCodecParserContext.height` value.
func (cpc *AVCodecParserContext) GetHeight() int32 {
	return (int32)(cpc.height)
}

// SetHeight sets `AVCodecParserContext.height` value.
func (cpc *AVCodecParserContext) SetHeight(v int32) {
	cpc.height = (C.int)(v)
}

// GetHeightAddr gets `AVCodecParserContext.height` address.
func (cpc *AVCodecParserContext) GetHeightAddr() *int32 {
	return (*int32)(&cpc.height)
}

// GetCodedWidth gets `AVCodecParserContext.coded_width` value.
func (cpc *AVCodecParserContext) GetCodedWidth() int32 {
	return (int32)(cpc.coded_width)
}

// SetCodedWidth sets `AVCodecParserContext.coded_width` value.
func (cpc *AVCodecParserContext) SetCodedWidth(v int32) {
	cpc.coded_width = (C.int)(v)
}

// GetCodedWidthAddr gets `AVCodecParserContext.coded_width` address.
func (cpc *AVCodecParserContext) GetCodedWidthAddr() *int32 {
	return (*int32)(&cpc.coded_width)
}

// GetCodedHeight gets `AVCodecParserContext.coded_height` value.
func (cpc *AVCodecParserContext) GetCodedHeight() int32 {
	return (int32)(cpc.coded_height)
}

// SetCodedHeight sets `AVCodecParserContext.coded_height` value.
func (cpc *AVCodecParserContext) SetCodedHeight(v int32) {
	cpc.coded_height = (C.int)(v)
}

// GetCodedHeightAddr gets `AVCodecParserContext.coded_height` address.
func (cpc *AVCodecParserContext) GetCodedHeightAddr() *int32 {
	return (*int32)(&cpc.coded_height)
}

// GetFormat gets `AVCodecParserContext.format` value.
func (cpc *AVCodecParserContext) GetFormat() int32 {
	return (int32)(cpc.format)
}

// SetFormat sets `AVCodecParserContext.format` value.
func (cpc *AVCodecParserContext) SetFormat(v int32) {
	cpc.format = (C.int)(v)
}

// GetFormatAddr gets `AVCodecParserContext.format` address.
func (cpc *AVCodecParserContext) GetFormatAddr() *int32 {
	return (*int32)(&cpc.format)
}

// AVCodecParser
type AVCodecParser C.struct_AVCodecParser

// GetCodecIds gets `AVCodecParser.codec_ids` value.
func (cp *AVCodecParser) GetCodecIds() []int32 {
	return unsafe.Slice((*int32)(&cp.codec_ids[0]), 5)
}

// SetCodecIds sets `AVCodecParser.codec_ids` value.
func (cp *AVCodecParser) SetCodecIds(v []int32) {
	for i := 0; i < FFMIN(len(v), 5); i++ {
		cp.codec_ids[i] = (C.int)(v[i])
	}
}

// GetCodecIdsAddr gets `AVCodecParser.codec_ids` address.
func (cp *AVCodecParser) GetCodecIdsAddr() *int32 {
	return (*int32)(unsafe.Pointer(&cp.codec_ids))
}

// GetPrivDataSize gets `AVCodecParser.priv_data_size` value.
func (cp *AVCodecParser) GetPrivDataSize() int32 {
	return (int32)(cp.priv_data_size)
}

// SetPrivDataSize sets `AVCodecParser.priv_data_size` value.
func (cp *AVCodecParser) SetPrivDataSize(v int32) {
	cp.priv_data_size = (C.int)(v)
}

// GetPrivDataSizeAddr gets `AVCodecParser.priv_data_size` address.
func (cp *AVCodecParser) GetPrivDataSizeAddr() *int32 {
	return (*int32)(&cp.priv_data_size)
}

// typedef int (*avcodec_parser_init_func)(AVCodecParserContext *s);
type AvcodecParserInitFunc = C.avcodec_parser_init_func

// typedef int (*avcodec_parser_parse_func)(AVCodecParserContext *s,
// AVCodecContext *avctx,
// const uint8_t **poutbuf, int *poutbuf_size,
// const uint8_t *buf, int buf_size);
type AvcodecParserParseFunc = C.avcodec_parser_parse_func

// typedef void (*avcodec_parser_close_func)(AVCodecParserContext *s);
type AvcodecParserCloseFunc = C.avcodec_parser_close_func

// typedef int (*avcodec_parser_split_func)(AVCodecContext *avctx, const uint8_t *buf, int buf_size);
type AvcodecParserSplitFunc = C.avcodec_parser_split_func

// GetParserInit gets `AVCodecParser.parser_init` value.
func (cp *AVCodecParser) GetParserInit() AvcodecParserInitFunc {
	return (AvcodecParserInitFunc)(cp.parser_init)
}

// SetParserInit sets `AVCodecParser.parser_init` value.
func (cp *AVCodecParser) SetParserInit(v AvcodecParserInitFunc) {
	cp.parser_init = (C.avcodec_parser_init_func)(v)
}

// GetParserInitAddr gets `AVCodecParser.parser_init` address.
func (cp *AVCodecParser) GetParserInitAddr() *AvcodecParserInitFunc {
	return (*AvcodecParserInitFunc)(&cp.parser_init)
}

// GetParserParse gets `AVCodecParser.parser_parse` value.
func (cp *AVCodecParser) GetParserParse() AvcodecParserParseFunc {
	return (AvcodecParserParseFunc)(cp.parser_parse)
}

// SetParserParse sets `AVCodecParser.parser_parse` value.
func (cp *AVCodecParser) SetParserParse(v AvcodecParserParseFunc) {
	cp.parser_parse = (C.avcodec_parser_parse_func)(v)
}

// GetParserParseAddr gets `AVCodecParser.parser_parse` address.
func (cp *AVCodecParser) GetParserParseAddr() *AvcodecParserParseFunc {
	return (*AvcodecParserParseFunc)(&cp.parser_parse)
}

// GetParserClose gets `AVCodecParser.parser_close` value.
func (cp *AVCodecParser) GetParserClose() AvcodecParserCloseFunc {
	return (AvcodecParserCloseFunc)(cp.parser_close)
}

// SetParserClose sets `AVCodecParser.parser_close` value.
func (cp *AVCodecParser) SetParserClose(v AvcodecParserCloseFunc) {
	cp.parser_close = (C.avcodec_parser_close_func)(v)
}

// GetParserCloseAddr gets `AVCodecParser.parser_close` address.
func (cp *AVCodecParser) GetParserCloseAddr() *AvcodecParserCloseFunc {
	return (*AvcodecParserCloseFunc)(&cp.parser_close)
}

// GetSplit gets `AVCodecParser.split` value.
func (cp *AVCodecParser) GetSplit() AvcodecParserSplitFunc {
	return (AvcodecParserSplitFunc)(cp.split)
}

// SetSplit sets `AVCodecParser.split` value.
func (cp *AVCodecParser) SetSplit(v AvcodecParserSplitFunc) {
	cp.split = (C.avcodec_parser_split_func)(v)
}

// GetSplitAddr gets `AVCodecParser.split` address.
func (cp *AVCodecParser) GetSplitAddr() *AvcodecParserSplitFunc {
	return (*AvcodecParserSplitFunc)(&cp.split)
}

// Deprecated: No use.
//
// GetNext gets `AVCodecParser.next` value.
func (cp *AVCodecParser) GetNext() *AVCodecParser {
	return (*AVCodecParser)(cp.next)
}

// Deprecated: No use.
//
// SetNext sets `AVCodecParser.next` value.
func (cp *AVCodecParser) SetNext(v *AVCodecParser) {
	cp.next = (*C.struct_AVCodecParser)(v)
}

// Deprecated: No use.
//
// GetNextAddr gets `AVCodecParser.next` address.
func (cp *AVCodecParser) GetNextAddr() **AVCodecParser {
	return (**AVCodecParser)(unsafe.Pointer(&cp.next))
}

// AvParserIterate iterates over all registered codec parsers.
func AvParserIterate(p CVoidPointerPointer) *AVCodecParser {
	return (*AVCodecParser)(C.av_parser_iterate(VoidPointerPointer(p)))
}

// Deprecated: No use.
//
// AvParserNext
func AvParserNext(c *AVCodecParser) *AVCodecParser {
	return (*AVCodecParser)(C.av_parser_next((*C.struct_AVCodecParser)(c)))
}

// Deprecated: No use.
//
// AvRegisterCodecParser
func AvRegisterCodecParser(parser *AVCodecParser) {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(parser))
}

// AvParserInit
func AvParserInit(codecID AVCodecID) *AVCodecParserContext {
	return (*AVCodecParserContext)(C.av_parser_init((C.int)(codecID)))
}

// Parse a packet.
func AvParserParse2(s *AVCodecParserContext, avctx *AVCodecContext,
	outbuf **uint8, poutbufSize *int32,
	buf *uint8, bufSize int32,
	pts, dts, pos int64) int32 {
	return (int32)(C.av_parser_parse2((*C.AVCodecParserContext)(s),
		(*C.struct_AVCodecContext)(avctx),
		(**C.uint8_t)(unsafe.Pointer(outbuf)), (*C.int)(poutbufSize),
		(*C.uint8_t)(buf), (C.int)(bufSize),
		(C.int64_t)(pts), (C.int64_t)(dts), (C.int64_t)(pos)))
}

// Deprecated: Use DumpExtradata, RemoveExtra or ExtractExtradata bitstream filters instead.
//
// AvParserChange
func AvParserChange(s *AVCodecParserContext, avctx *AVCodecContext,
	outbuf **uint8, poutbufSize *int32,
	buf *uint8, bufSize int32, keyframe int32) int32 {
	return (int32)(C.av_parser_change((*C.AVCodecParserContext)(s),
		(*C.struct_AVCodecContext)(avctx),
		(**C.uint8_t)(unsafe.Pointer(outbuf)), (*C.int)(poutbufSize),
		(*C.uint8_t)(buf), (C.int)(bufSize), (C.int)(keyframe)))
}

// AvParserClose
func AvParserClose(s *AVCodecParserContext) {
	C.av_parser_close((*C.AVCodecParserContext)(s))
}

// Deprecated: Use AVCodecSendFrame()/AVCodecReceivePacket() instead.
//
// AvCodecEncodeAudio2 encodes a frame of audio.
func AvCodecEncodeAudio2(avctx *AVCodecContext,
	avpkt *AVPacket, frame *AVFrame, gotPacketPtr *int32) int32 {
	return (int32)(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(avpkt), (*C.struct_AVFrame)(frame), (*C.int)(gotPacketPtr)))
}

// Deprecated: Use AVCodecSendFrame()/AVCodecReceivePacket() instead.
//
// AvCodecEncodeVideo2 encodes a frame of video.
func AvCodecEncodeVideo2(avctx *AVCodecContext,
	avpkt *AVPacket, frame *AVFrame, gotPacketPtr *int32) int32 {
	return (int32)(C.avcodec_encode_video2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(avpkt), (*C.struct_AVFrame)(frame), (*C.int)(gotPacketPtr)))
}

// AvCodecEncodeSubtitle
func AvCodecEncodeSubtitle(avctx *AVCodecContext,
	buf *uint8, bufSize int32, sub *AVSubtitle) int32 {
	return (int32)(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(avctx),
		(*C.uint8_t)(buf), (C.int)(bufSize), (*C.struct_AVSubtitle)(sub)))
}

// Deprecated: Unused.
//
// AvPictureAlloc
func AvPictureAlloc(picture *AVPicture, pixFmt AVPixelFormat, width, height int32) int32 {
	return (int32)(C.avpicture_alloc((*C.struct_AVPicture)(picture),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height)))
}

// Deprecated: Unused.
//
// AvPictureFree
func AvPictureFree(picture *AVPicture) {
	C.avpicture_free((*C.struct_AVPicture)(picture))
}

// Deprecated: Use AvImageFillArrays() instead.
//
// AvPictureFill
func AvPictureFill(picture *AVPicture, ptr *uint8, pixFmt AVPixelFormat, width, height int32) int32 {
	return (int32)(C.avpicture_fill((*C.struct_AVPicture)(picture),
		(*C.uint8_t)(ptr), (C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height)))
}

// Deprecated: Use AvImageCopyToBuffer() instead.
//
// AvPictureLayout
func AvPictureLayout(src *AVPicture, pixFmt AVPixelFormat, width, height int32, dest *uint8, destSize int32) int32 {
	return (int32)(C.avpicture_layout((*C.struct_AVPicture)(src),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height),
		(*C.uchar)(dest), (C.int)(destSize)))
}

// Deprecated: Use AvImageGetBufferSize() instead.
//
// AvPictureGetSize
func AvPictureGetSize(pixFmt AVPixelFormat, width, height int32) int32 {
	return (int32)(C.avpicture_get_size((C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height)))
}

// Deprecated: Use AvImageCopy() instead.
//
// AvPictureCopy
func AvPictureCopy(dst, src *AVPicture, pixFmt AVPixelFormat, width, height int32) {
	C.av_picture_copy((*C.struct_AVPicture)(dst), (*C.struct_AVPicture)(src),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height))
}

// Deprecated: Unused.
//
// AvPictureCrop
func AvPictureCrop(dst, src *AVPicture, pixFmt AVPixelFormat, topBand, leftBand int32) int32 {
	return (int32)(C.av_picture_crop((*C.struct_AVPicture)(dst), (*C.struct_AVPicture)(src),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(topBand), (C.int)(leftBand)))
}

// Deprecated: Unused.
//
// AvPicturePad
func AvPicturePad(dst, src *AVPicture, width, height int32, pixFmt AVPixelFormat,
	padTop, padBottom, padLeft, padRight int32, color *int32) int32 {
	return (int32)(C.av_picture_pad((*C.struct_AVPicture)(dst), (*C.struct_AVPicture)(src),
		(C.int)(width), (C.int)(height), (C.enum_AVPixelFormat)(pixFmt),
		(C.int)(padTop), (C.int)(padBottom), (C.int)(padLeft), (C.int)(padRight),
		(*C.int)(color)))
}

// Deprecated: Use AvPixFmtGetChromaSubSample() instead.
//
// AvCodecGetChromaSubSample
func AvCodecGetChromaSubSample(pixFmt AVPixelFormat, hShift, vShift *int32) {
	C.avcodec_get_chroma_sub_sample((C.enum_AVPixelFormat)(pixFmt),
		(*C.int)(hShift), (*C.int)(vShift))
}

// AvCodecPixFmtToCodecTag returns a value representing the fourCC code associated to the
// pixel format pix_fmt, or 0 if no associated fourCC code can be found.
func AvCodecPixFmtToCodecTag(pixFmt AVPixelFormat) uint {
	return (uint)(C.avcodec_pix_fmt_to_codec_tag((C.enum_AVPixelFormat)(pixFmt)))
}

// AvCodecFindBestPixFmtOfList finds the best pixel format
// to convert to given a certain source pixel format.
func AvCodecFindBestPixFmtOfList(pixFmtList *AVPixelFormat,
	srcPixFmt AVPixelFormat, hasAlpha int32, lossPtr *int32) AVPixelFormat {
	return (AVPixelFormat)(C.avcodec_find_best_pix_fmt_of_list(
		(*C.enum_AVPixelFormat)(pixFmtList),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha),
		(*C.int)(lossPtr)))
}

// Deprecated: Use AvGetPixFmtLoss() instead.
//
// AvCodecGetPixFmtLoss
func AvCodecGetPixFmtLoss(dstPixFmt, srcPixFmt AVPixelFormat, hasAlpha int32) int32 {
	return (int32)(C.avcodec_get_pix_fmt_loss((C.enum_AVPixelFormat)(dstPixFmt),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha)))
}

// Deprecated: Use AvFindBestPixFmtOf2() instead.
//
// AvCodecFindBestPixFmtOf2
func AvCodecFindBestPixFmtOf2(dstPixFmt1, dstPixFmt2, srcPixFmt AVPixelFormat,
	hasAlpha int32, lossPtr *int32) AVPixelFormat {
	return (AVPixelFormat)(C.avcodec_find_best_pix_fmt_of_2(
		(C.enum_AVPixelFormat)(dstPixFmt1), (C.enum_AVPixelFormat)(dstPixFmt2),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha), (*C.int)(lossPtr)))
}

// Deprecated: No use.
//
// AvCodecFindBestPixFmt2
func AvCodecFindBestPixFmt2(dstPixFmt1, dstPixFmt2, srcPixFmt AVPixelFormat,
	hasAlpha int32, lossPtr *int32) AVPixelFormat {
	return (AVPixelFormat)(C.avcodec_find_best_pix_fmt2(
		(C.enum_AVPixelFormat)(dstPixFmt1), (C.enum_AVPixelFormat)(dstPixFmt2),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha), (*C.int)(lossPtr)))
}

// AvCodecDefaultGetFormat
func AvCodecDefaultGetFormat(avctx *AVCodecContext, fmt *AVPixelFormat) AVPixelFormat {
	return (AVPixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(avctx),
		(*C.enum_AVPixelFormat)(fmt)))
}

// Deprecated: Use  AvFourccMakeString() or AvFourcc2str() instead.
//
// AvGetCodecTagString
func AvGetCodecTagString(buf *int8, bufSize uintptr, codecTag uint32) int32 {
	return (int32)(C.av_get_codec_tag_string((*C.char)(buf),
		(C.size_t)(bufSize), (C.uint)(codecTag)))
}

// AvCodecString
func AvCodecString(buf *int8, bufSize int32, enc *AVCodecContext, encode int32) {
	C.avcodec_string((*C.char)(buf), (C.int)(bufSize),
		(*C.struct_AVCodecContext)(enc), (C.int)(encode))
}

// AvGetProfileName returns a name for the specified profile, if available.
func AvGetProfileName(c *AVCodec, profile int32) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), (C.int)(profile)))
}

// AvCodecProfileName returns a name for the specified profile, if available.
func AvCodecProfileName(codecID AVCodecID, profile int32) string {
	return C.GoString(C.avcodec_profile_name((C.enum_AVCodecID)(codecID), (C.int)(profile)))
}

// typedef int (*avcodec_context_execute_func)(AVCodecContext *c2, void *arg2);
type AVCodecContextExecuteFunc = C.avcodec_context_execute_func

// typedef int (*avcodec_context_execute2_func)(AVCodecContext *c2, void *arg2, int, int);
type AVCodecContextExecute2Func = C.avcodec_context_execute2_func

// AvCodecDefaultExecute
func AvCodecDefaultExecute(avctx *AVCodecContext, f AVCodecContextExecuteFunc, arg CVoidPointer,
	ret *int32, count, size int32) int32 {
	return (int32)(C.avcodec_default_execute((*C.struct_AVCodecContext)(avctx),
		(C.avcodec_context_execute_func)(f), VoidPointer(arg), (*C.int)(ret), (C.int)(count), (C.int)(size)))
}

// AvCodecDefaultExecute2
func AvCodecDefaultExecute2(avctx *AVCodecContext, f AVCodecContextExecute2Func, arg CVoidPointer,
	ret *int32, count int32) int32 {
	return (int32)(C.avcodec_default_execute2((*C.struct_AVCodecContext)(avctx),
		(C.avcodec_context_execute2_func)(f), VoidPointer(arg), (*C.int)(ret), (C.int)(count)))
}

// AvCodecFillAudioFrame fills AVFrame audio data and linesize pointers.
func AvCodecFillAudioFrame(frame *AVFrame, nbChannels int32,
	sampleFmt AVSampleFormat, buf *uint8,
	bufSize int32, align int32) int32 {
	return (int32)(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(frame),
		(C.int)(nbChannels), (C.enum_AVSampleFormat)(sampleFmt),
		(*C.uint8_t)(buf), (C.int)(bufSize), (C.int)(align)))
}

// AvCodecFlushBuffers resets the internal codec state / flush internal buffers. Should be called
// e.g. when seeking or when switching to a different stream.
func AvCodecFlushBuffers(avctx *AVCodecContext) {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(avctx))
}

// AvGetBitsPerSample returns codec bits per sample.
func AvGetBitsPerSample(codecID AVCodecID) int32 {
	return (int32)(C.av_get_bits_per_sample((C.enum_AVCodecID)(codecID)))
}

// AvGetPcmCodec returns the PCM codec associated with a sample format.
func AvGetPcmCodec(fmt AVSampleFormat, be int32) AVCodecID {
	return (AVCodecID)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(fmt), (C.int)(be)))
}

// AvGetExactBitsPerSample returns codec bits per sample.
func AvGetExactBitsPerSample(codecID AVCodecID) int32 {
	return (int32)(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(codecID)))
}

// AvGetAudioFrameDuration returns audio frame duration.
func AvGetAudioFrameDuration(avctx *AVCodecContext, frameBytes int32) int32 {
	return (int32)(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(avctx), (C.int)(frameBytes)))
}

// AvGetAudioFrameDuration2 returns audio frame duration.
func AvGetAudioFrameDuration2(par *AVCodecParameters, frameBytes int32) int32 {
	return (int32)(C.av_get_audio_frame_duration2((*C.struct_AVCodecParameters)(par), (C.int)(frameBytes)))
}

// AVBitStreamFilterContext
type AVBitStreamFilterContext C.struct_AVBitStreamFilterContext

// GetPrivData gets `AVBitStreamFilterContext.priv_data` value.
func (obsfc *AVBitStreamFilterContext) GetPrivData() unsafe.Pointer {
	return obsfc.priv_data
}

// GetFilter gets `AVBitStreamFilterContext.filter` value.
func (obsfc *AVBitStreamFilterContext) GetFilter() *AVBitStreamFilter {
	return (*AVBitStreamFilter)(obsfc.filter)
}

// GetParser gets `AVBitStreamFilterContext.parser` value.
func (obsfc *AVBitStreamFilterContext) GetParser() *AVCodecParserContext {
	return (*AVCodecParserContext)(obsfc.parser)
}

// GetNext gets `AVBitStreamFilterContext.next` value.
func (obsfc *AVBitStreamFilterContext) GetNext() *AVBitStreamFilterContext {
	return (*AVBitStreamFilterContext)(obsfc.next)
}

// GetArgs gets `AVBitStreamFilterContext.args` value.
func (obsfc *AVBitStreamFilterContext) GetArgs() string {
	return C.GoString(obsfc.args)
}

// Deprecated: Use AVBSFContext instead.
//
// AvRegisterBitstreamFilter
func AvRegisterBitstreamFilter(bsf *AVBitStreamFilter) {
	C.av_register_bitstream_filter((*C.struct_AVBitStreamFilter)(bsf))
}

// Deprecated: Use AVBSFContext instead.
//
// AvBitstreamFilterInit
func AvBitstreamFilterInit(name string) *AVBitStreamFilterContext {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AVBitStreamFilterContext)(C.av_bitstream_filter_init((*C.char)(namePtr)))
}

// Deprecated: Use AVBSFContext instead.
//
// AvBitstreamFilterFilter
func AvBitstreamFilterFilter(bsfc *AVBitStreamFilterContext, avctx *AVCodecContext, args string,
	outbuf **uint8, poutbufSize *int32,
	buf *uint8, bufSize int32, keyframe int32) int32 {
	argsPtr, nameFunc := StringCasting(args)
	defer nameFunc()
	return (int32)(C.av_bitstream_filter_filter((*C.struct_AVBitStreamFilterContext)(bsfc),
		(*C.struct_AVCodecContext)(avctx), (*C.char)(argsPtr),
		(**C.uint8_t)(unsafe.Pointer(outbuf)), (*C.int)(poutbufSize),
		(*C.uint8_t)(buf), (C.int)(bufSize), (C.int)(keyframe)))
}

// Deprecated: Use AVBSFContext instead.
//
// AvBitstreamFilterClose
func AvBitstreamFilterClose(bsfc *AVBitStreamFilterContext) {
	C.av_bitstream_filter_close((*C.struct_AVBitStreamFilterContext)(bsfc))
}

// Deprecated: Use AVBSFContext instead.
//
// AvBitstreamFilterNext
func AvBitstreamFilterNext(f *AVBitStreamFilter) *AVBitStreamFilter {
	return (*AVBitStreamFilter)(C.av_bitstream_filter_next((*C.struct_AVBitStreamFilter)(f)))
}

// Deprecated: No use.
//
// AvBsfNext
func AvBsfNext(opaque CVoidPointerPointer) *AVBitStreamFilter {
	return (*AVBitStreamFilter)(C.av_bsf_next(VoidPointerPointer(opaque)))
}

// AvFastPaddedMalloc
func AvFastPaddedMalloc(ptr CVoidPointer, size *uint32, minSize uintptr) {
	C.av_fast_padded_malloc(VoidPointer(ptr), (*C.uint)(size), (C.size_t)(minSize))
}

// AvFastPaddedMallocz
func AvFastPaddedMallocz(ptr CVoidPointer, size *uint32, minSize uintptr) {
	C.av_fast_padded_mallocz(VoidPointer(ptr), (*C.uint)(size), (C.size_t)(minSize))
}

// AvXiphlacing encodes extradata length to a buffer. Used by xiph codecs.
func AvXiphlacing(s *uint8, v int32) int32 {
	return (int32)(C.av_xiphlacing((*C.uchar)(s), (C.uint)(v)))
}

// Deprecated: No use.
//
// AvRegisterHwaccel
func AvRegisterHwaccel(hwaccel *AVHWAccel) {
	C.av_register_hwaccel((*C.struct_AVHWAccel)(hwaccel))
}

// Deprecated: No use.
//
// AvHwaccelNext returns the first registered hardware accelerator if hwaccel is NULL,
// returns the next registered hardware accelerator after hwaccelif hwaccel is non-NULL,
// or NULL if hwaccel is the last one.
func AvHwaccelNext(hwaccel *AVHWAccel) *AVHWAccel {
	return (*AVHWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(hwaccel)))
}

// AVLockOp
type AVLockOp = C.enum_AVLockOp

const (
	AV_LOCK_CREATE  = AVLockOp(C.AV_LOCK_CREATE)
	AV_LOCK_OBTAIN  = AVLockOp(C.AV_LOCK_OBTAIN)
	AV_LOCK_RELEASE = AVLockOp(C.AV_LOCK_RELEASE)
	AV_LOCK_DESTROY = AVLockOp(C.AV_LOCK_DESTROY)
)

// typedef int (*av_lockmgr_cb)(void **mutex, enum AVLockOp op);
type AVLockmgrCb C.av_lockmgr_cb

// Deprecated: No use.
//
// AvLockmgrRegister
func AvLockmgrRegister(cb AVLockmgrCb) int32 {
	return (int32)(C.av_lockmgr_register((C.av_lockmgr_cb)(cb)))
}

// A positive value if s is open,
// 0 otherwise.
func AvCodecIsOpen(avctx *AVCodecContext) int32 {
	return (int32)(C.avcodec_is_open((*C.struct_AVCodecContext)(avctx)))
}

// AvCpbPropertiesAlloc allocates a CPB properties structure and initialize its fields to default
// values.
func AvCpbPropertiesAlloc(size *uintptr) *AVCPBProperties {
	return (*AVCPBProperties)(C.av_cpb_properties_alloc((*C.size_t)(unsafe.Pointer(size))))
}
