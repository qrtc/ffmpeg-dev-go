package ffmpeg

/*
#include <libavcodec/avcodec.h>

typedef int (*avcodec_excute_func)(AVCodecContext *c2, void *arg2);
typedef int (*av_lockmgr_cb)(void **mutex, enum AVLockOp op);
*/
import "C"
import "unsafe"

const (
	// Required number of additionally allocated bytes at the end of the input bitstream for decoding.
	AV_INPUT_BUFFER_PADDING_SIZE = C.AV_INPUT_BUFFER_PADDING_SIZE
	// Minimum encoding buffer size.
	AV_INPUT_BUFFER_MIN_SIZE = C.AV_INPUT_BUFFER_MIN_SIZE
)

// AvDiscard
type AvDiscard = C.enum_AVDiscard

const (
	AVDISCARD_NONE     = AvDiscard(C.AVDISCARD_NONE)
	AVDISCARD_DEFAULT  = AvDiscard(C.AVDISCARD_DEFAULT)
	AVDISCARD_NONREF   = AvDiscard(C.AVDISCARD_NONREF)
	AVDISCARD_BIDIR    = AvDiscard(C.AVDISCARD_BIDIR)
	AVDISCARD_NONINTRA = AvDiscard(C.AVDISCARD_NONINTRA)
	AVDISCARD_NONKEY   = AvDiscard(C.AVDISCARD_NONKEY)
	AVDISCARD_ALL      = AvDiscard(C.AVDISCARD_ALL)
)

// AvAudioServiceType
type AvAudioServiceType = C.enum_AVAudioServiceType

const (
	AV_AUDIO_SERVICE_TYPE_MAIN              = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_MAIN)
	AV_AUDIO_SERVICE_TYPE_EFFECTS           = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_EFFECTS)
	AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED)
	AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED  = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED)
	AV_AUDIO_SERVICE_TYPE_DIALOGUE          = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_DIALOGUE)
	AV_AUDIO_SERVICE_TYPE_COMMENTARY        = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_COMMENTARY)
	AV_AUDIO_SERVICE_TYPE_EMERGENCY         = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_EMERGENCY)
	AV_AUDIO_SERVICE_TYPE_VOICE_OVER        = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_VOICE_OVER)
	AV_AUDIO_SERVICE_TYPE_KARAOKE           = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_KARAOKE)
	AV_AUDIO_SERVICE_TYPE_NB                = AvAudioServiceType(C.AV_AUDIO_SERVICE_TYPE_NB)
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
type AvPanScan C.struct_AVPanScan

// This structure describes the bitrate properties of an encoded bitstream. It
// roughly corresponds to a subset the VBV parameters for MPEG-2 or HRD
// parameters for H.264/HEVC.
type AvCPBProperties C.struct_AVCPBProperties

// This structure supplies correlation between a packet timestamp and a wall clock
// production time. The definition follows the Producer Reference Time ('prft')
// as defined in ISO/IEC 14496-12
type AvProducerReferenceTime C.struct_AVProducerReferenceTime

const (
	AV_GET_BUFFER_FLAG_REF        = C.AV_GET_BUFFER_FLAG_REF
	AV_GET_ENCODE_BUFFER_FLAG_REF = C.AV_GET_ENCODE_BUFFER_FLAG_REF
)

// AvCodecContext is main external API structure.
type AvCodecContext C.struct_AVCodecContext

// Custom: Get AvClass gets `AVCodecContext.av_class` value.
func (avctx *AvCodecContext) GetAvClass() *AvClass {
	return (*AvClass)(avctx.av_class)
}

// Custom: SetAvClass sets `AVCodecContext.av_class` value.
func (avctx *AvCodecContext) SetAvClass(v *AvClass) {
	avctx.av_class = (*C.struct_AVClass)(v)
}

// Custom: GetAvClassAddr gets `AVCodecContext.av_class` address.
func (avctx *AvCodecContext) GetAvClassAddr() **AvClass {
	return (**AvClass)(unsafe.Pointer(&avctx.av_class))
}

// Custom: Get LogLevelOffset gets `AVCodecContext.log_level_offset` value.
func (avctx *AvCodecContext) GetLogLevelOffset() int32 {
	return (int32)(avctx.log_level_offset)
}

// Custom: SetLogLevelOffset sets `AVCodecContext.log_level_offset` value.
func (avctx *AvCodecContext) SetLogLevelOffset(v int32) {
	avctx.log_level_offset = (C.int)(v)
}

// Custom: GetLogLevelOffsetAddr gets `AVCodecContext.log_level_offset` address.
func (avctx *AvCodecContext) GetLogLevelOffsetAddr() *int32 {
	return (*int32)(&avctx.log_level_offset)
}

// Custom: Get CodecType gets `AVCodecContext.codec_type` value.
func (avctx *AvCodecContext) GetCodecType() AvMediaType {
	return (AvMediaType)(avctx.codec_type)
}

// Custom: SetCodecType sets `AVCodecContext.codec_type` value.
func (avctx *AvCodecContext) SetCodecType(v AvMediaType) {
	avctx.codec_type = (C.enum_AVMediaType)(v)
}

// Custom: GetCodecTypeAddr gets `AVCodecContext.codec_type` address.
func (avctx *AvCodecContext) GetCodecTypeAddr() *AvMediaType {
	return (*AvMediaType)(&avctx.codec_type)
}

// Custom: Get Codec gets `AVCodecContext.codec` value.
func (avctx *AvCodecContext) GetCodec() *AvCodec {
	return (*AvCodec)(avctx.codec)
}

// Custom: SetCodec sets `AVCodecContext.codec` value.
func (avctx *AvCodecContext) SetCodec(v *AvCodec) {
	avctx.codec = (*C.struct_AVCodec)(v)
}

// Custom: GetCodecAddr gets `AVCodecContext.codec` address.
func (avctx *AvCodecContext) GetCodecAddr() **AvCodec {
	return (**AvCodec)(unsafe.Pointer(&avctx.codec))
}

// Custom: Get CodecId gets `AVCodecContext.codec_id` value.
func (avctx *AvCodecContext) GetCodecId() AvCodecID {
	return (AvCodecID)(avctx.codec_id)
}

// Custom: SetCodecId sets `AVCodecContext.codec_id` value.
func (avctx *AvCodecContext) SetCodecId(v AvCodecID) {
	avctx.codec_id = (C.enum_AVCodecID)(v)
}

// Custom: GetCodecIdAddr gets `AVCodecContext.codec_id` address.
func (avctx *AvCodecContext) GetCodecIdAddr() *AvCodecID {
	return (*AvCodecID)(unsafe.Pointer(&avctx.codec_id))
}

// Custom: Get CodecTag gets `AVCodecContext.codec_tag` value.
func (avctx *AvCodecContext) GetCodecTag() uint32 {
	return (uint32)(avctx.codec_tag)
}

// Custom: SetCodecTag sets `AVCodecContext.codec_tag` value.
func (avctx *AvCodecContext) SetCodecTag(v uint32) {
	avctx.codec_tag = (C.uint)(v)
}

// Custom: GetCodecTagAddr gets `AVCodecContext.codec_tag` address.
func (avctx *AvCodecContext) GetCodecTagAddr() *uint32 {
	return (*uint32)(&avctx.codec_tag)
}

// Custom: Get PrivData gets `AVCodecContext.priv_data` value.
func (avctx *AvCodecContext) GetPrivData() unsafe.Pointer {
	return avctx.priv_data
}

// Custom: SetPrivData sets `AVCodecContext.priv_data` value.
func (avctx *AvCodecContext) SetPrivData(v unsafe.Pointer) {
	avctx.priv_data = v
}

// Custom: GetPrivDataAddr gets `AVCodecContext.priv_data` address.
func (avctx *AvCodecContext) GetPrivDataAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&avctx.priv_data)
}

// Custom: Get Opaque gets `AVCodecContext.opaque` value.
func (avctx *AvCodecContext) GetOpaque() unsafe.Pointer {
	return (unsafe.Pointer)(avctx.opaque)
}

// Custom: SetOpaque sets `AVCodecContext.opaque` value.
func (avctx *AvCodecContext) SetOpaque(v unsafe.Pointer) {
	avctx.opaque = v
}

// Custom: GetOpaqueAddr gets `AVCodecContext.opaque` address.
func (avctx *AvCodecContext) GetOpaqueAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&avctx.opaque)
}

// Custom: Get BitRate gets `AVCodecContext.bit_rate` value.
func (avctx *AvCodecContext) GetBitRate() int64 {
	return (int64)(avctx.bit_rate)
}

// Custom: SetBitRate sets `AVCodecContext.bit_rate` value.
func (avctx *AvCodecContext) SetBitRate(v int64) {
	avctx.bit_rate = (C.int64_t)(v)
}

// Custom: GetBitRateAddr gets `AVCodecContext.bit_rate` address.
func (avctx *AvCodecContext) GetBitRateAddr() *int64 {
	return (*int64)(&avctx.bit_rate)
}

// Custom: Get BitRateTolerance gets `AVCodecContext.bit_rate_tolerance` value.
func (avctx *AvCodecContext) GetBitRateTolerance() int32 {
	return (int32)(avctx.bit_rate_tolerance)
}

// Custom: SetBitRateTolerance sets `AVCodecContext.bit_rate_tolerance` value.
func (avctx *AvCodecContext) SetBitRateTolerance(v int32) {
	avctx.bit_rate_tolerance = (C.int)(v)
}

// Custom: GetBitRateToleranceAddr gets `AVCodecContext.bit_rate_tolerance` address.
func (avctx *AvCodecContext) GetBitRateToleranceAddr() *int32 {
	return (*int32)(&avctx.bit_rate_tolerance)
}

// Custom: Get GlobalQuality gets `AVCodecContext.global_quality` value.
func (avctx *AvCodecContext) GetGlobalQuality() int32 {
	return (int32)(avctx.global_quality)
}

// Custom: SetGlobalQuality sets `AVCodecContext.global_quality` value.
func (avctx *AvCodecContext) SetGlobalQuality(v int32) {
	avctx.global_quality = (C.int)(v)
}

// Custom: GetGlobalQualityAddr gets `AVCodecContext.global_quality` address.
func (avctx *AvCodecContext) GetGlobalQualityAddr() *int32 {
	return (*int32)(&avctx.global_quality)
}

// Custom: Get CompressionLevel gets `AVCodecContext.compression_level` value.
func (avctx *AvCodecContext) GetCompressionLevel() int32 {
	return (int32)(avctx.compression_level)
}

// Custom: SetCompressionLevel sets `AVCodecContext.compression_level` value.
func (avctx *AvCodecContext) SetCompressionLevel(v int32) {
	avctx.compression_level = (C.int)(v)
}

// Custom: GetCompressionLevelAddr gets `AVCodecContext.compression_level` address.
func (avctx *AvCodecContext) GetCompressionLevelAddr() *int32 {
	return (*int32)(&avctx.compression_level)
}

// Custom: Get Flags gets `AVCodecContext.flags` value.
func (avctx *AvCodecContext) GetFlags() int32 {
	return (int32)(avctx.flags)
}

// Custom: SetFlags sets `AVCodecContext.flags` value.
func (avctx *AvCodecContext) SetFlags(v int32) {
	avctx.flags = (C.int)(v)
}

// Custom: GetFlagsAddr gets `AVCodecContext.flags` address.
func (avctx *AvCodecContext) GetFlagsAddr() *int32 {
	return (*int32)(&avctx.flags)
}

// Custom: Get Flags2 gets `AVCodecContext.flags2` value.
func (avctx *AvCodecContext) GetFlags2() int32 {
	return (int32)(avctx.flags2)
}

// Custom: SetFlags2 sets `AVCodecContext.flags2` value.
func (avctx *AvCodecContext) SetFlags2(v int32) {
	avctx.flags2 = (C.int)(v)
}

// Custom: GetFlags2Addr gets `AVCodecContext.flags2` address.
func (avctx *AvCodecContext) GetFlags2Addr() *int32 {
	return (*int32)(&avctx.flags2)
}

// Custom: Get Extradata gets `AVCodecContext.extradata` value.
func (avctx *AvCodecContext) GetExtradata() *uint8 {
	return (*uint8)(avctx.extradata)
}

// Custom: SetExtradata sets `AVCodecContext.extradata` value.
func (avctx *AvCodecContext) SetExtradata(v *uint8) {
	avctx.extradata = (*C.uint8_t)(v)
}

// Custom: GetExtradataAddr gets `AVCodecContext.extradata` address.
func (avctx *AvCodecContext) GetExtradataAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&avctx.extradata))
}

// Custom: Get ExtradataSize gets `AVCodecContext.extradata_size` value.
func (avctx *AvCodecContext) GetExtradataSize() int32 {
	return (int32)(avctx.extradata_size)
}

// Custom: SetExtradataSize sets `AVCodecContext.extradata_size` value.
func (avctx *AvCodecContext) SetExtradataSize(v int32) {
	avctx.extradata_size = (C.int)(v)
}

// Custom: GetExtradataSizeAddr gets `AVCodecContext.extradata_size` address.
func (avctx *AvCodecContext) GetExtradataSizeAddr() *int32 {
	return (*int32)(&avctx.extradata_size)
}

// Custom: Get TimeBase gets `AVCodecContext.time_base` value.
func (avctx *AvCodecContext) GetTimeBase() AvRational {
	return (AvRational)(avctx.time_base)
}

// Custom: SetTimeBase sets `AVCodecContext.time_base` value.
func (avctx *AvCodecContext) SetTimeBase(v AvRational) {
	avctx.time_base = (C.AVRational)(v)
}

// Custom: GetTimeBaseAddr gets `AVCodecContext.time_base` address.
func (avctx *AvCodecContext) GetTimeBaseAddr() *AvRational {
	return (*AvRational)(&avctx.time_base)
}

// Custom: Get TicksPerFrame gets `AVCodecContext.ticks_per_frame` value.
func (avctx *AvCodecContext) GetTicksPerFrame() int32 {
	return (int32)(avctx.ticks_per_frame)
}

// Custom: SetTicksPerFrame sets `AVCodecContext.ticks_per_frame` value.
func (avctx *AvCodecContext) SetTicksPerFrame(v int32) {
	avctx.ticks_per_frame = (C.int)(v)
}

// Custom: GetTicksPerFrameAddr gets `AVCodecContext.ticks_per_frame` address.
func (avctx *AvCodecContext) GetTicksPerFrameAddr() *int32 {
	return (*int32)(&avctx.ticks_per_frame)
}

// Custom: Get Delay gets `AVCodecContext.delay` value.
func (avctx *AvCodecContext) GetDelay() int32 {
	return (int32)(avctx.delay)
}

// Custom: SetDelay sets `AVCodecContext.delay` value.
func (avctx *AvCodecContext) SetDelay(v int32) {
	avctx.delay = (C.int)(v)
}

// Custom: GetDelayAddr gets `AVCodecContext.delay` address.
func (avctx *AvCodecContext) GetDelayAddr() *int32 {
	return (*int32)(&avctx.delay)
}

// Custom: Get Width gets `AVCodecContext.width` value.
func (avctx *AvCodecContext) GetWidth() int32 {
	return (int32)(avctx.width)
}

// Custom: SetWidth sets `AVCodecContext.width` value.
func (avctx *AvCodecContext) SetWidth(v int32) {
	avctx.width = (C.int)(v)
}

// Custom: GetWidthAddr gets `AVCodecContext.width` address.
func (avctx *AvCodecContext) GetWidthAddr() *int32 {
	return (*int32)(&avctx.width)
}

// Custom: Get Height gets `AVCodecContext.height` value.
func (avctx *AvCodecContext) GetHeight() int32 {
	return (int32)(avctx.height)
}

// Custom: SetHeight sets `AVCodecContext.height` value.
func (avctx *AvCodecContext) SetHeight(v int32) {
	avctx.height = (C.int)(v)
}

// Custom: GetHeightAddr gets `AVCodecContext.height` address.
func (avctx *AvCodecContext) GetHeightAddr() *int32 {
	return (*int32)(&avctx.height)
}

// Custom: Get CodedWidth gets `AVCodecContext.coded_width` value.
func (avctx *AvCodecContext) GetCodedWidth() int32 {
	return (int32)(avctx.coded_width)
}

// Custom: SetCodedWidth sets `AVCodecContext.coded_width` value.
func (avctx *AvCodecContext) SetCodedWidth(v int32) {
	avctx.coded_width = (C.int)(v)
}

// Custom: GetCodedWidthAddr gets `AVCodecContext.coded_width` address.
func (avctx *AvCodecContext) GetCodedWidthAddr() *int32 {
	return (*int32)(&avctx.coded_width)
}

// Custom: Get CodedHeight gets `AVCodecContext.coded_height` value.
func (avctx *AvCodecContext) GetCodedHeight() int32 {
	return (int32)(avctx.coded_height)
}

// Custom: SetCodedHeight sets `AVCodecContext.coded_height` value.
func (avctx *AvCodecContext) SetCodedHeight(v int32) {
	avctx.coded_height = (C.int)(v)
}

// Custom: GetCodedHeightAddr gets `AVCodecContext.coded_height` address.
func (avctx *AvCodecContext) GetCodedHeightAddr() *int32 {
	return (*int32)(&avctx.coded_height)
}

// Custom: Get GopSize gets `AVCodecContext.gop_size` value.
func (avctx *AvCodecContext) GetGopSize() int32 {
	return (int32)(avctx.gop_size)
}

// Custom: SetGopSize sets `AVCodecContext.gop_size` value.
func (avctx *AvCodecContext) SetGopSize(v int32) {
	avctx.gop_size = (C.int)(v)
}

// Custom: GetGopSizeAddr gets `AVCodecContext.gop_size` address.
func (avctx *AvCodecContext) GetGopSizeAddr() *int32 {
	return (*int32)(&avctx.gop_size)
}

// Custom: Get PixFmt gets `AVCodecContext.pix_fmt` value.
func (avctx *AvCodecContext) GetPixFmt() AvPixelFormat {
	return (AvPixelFormat)(avctx.pix_fmt)
}

// Custom: SetPixFmt sets `AVCodecContext.pix_fmt` value.
func (avctx *AvCodecContext) SetPixFmt(v AvPixelFormat) {
	avctx.pix_fmt = (C.enum_AVPixelFormat)(v)
}

// Custom: GetPixFmtAddr gets `AVCodecContext.pix_fmt` address.
func (avctx *AvCodecContext) GetPixFmtAddr() *AvPixelFormat {
	return (*AvPixelFormat)(&avctx.pix_fmt)
}

// Custom: Get MaxBFrames gets `AVCodecContext.max_b_frames` value.
func (avctx *AvCodecContext) GetMaxBFrames() int32 {
	return (int32)(avctx.max_b_frames)
}

// Custom: SetMaxBFrames sets `AVCodecContext.max_b_frames` value.
func (avctx *AvCodecContext) SetMaxBFrames(v int32) {
	avctx.max_b_frames = (C.int)(v)
}

// Custom: GetMaxBFramesAddr gets `AVCodecContext.max_b_frames` address.
func (avctx *AvCodecContext) GetMaxBFramesAddr() *int32 {
	return (*int32)(&avctx.max_b_frames)
}

// Custom: Get BQuantFactor gets `AVCodecContext.b_quant_factor` value.
func (avctx *AvCodecContext) GetBQuantFactor() float32 {
	return (float32)(avctx.b_quant_factor)
}

// Custom: SetBQuantFactor sets `AVCodecContext.b_quant_factor` value.
func (avctx *AvCodecContext) SetBQuantFactor(v float32) {
	avctx.b_quant_factor = (C.float)(v)
}

// Custom: GetBQuantFactorAddr gets `AVCodecContext.b_quant_factor` address.
func (avctx *AvCodecContext) GetBQuantFactorAddr() *float32 {
	return (*float32)(&avctx.b_quant_factor)
}

// Custom: Get BFrameStrategy gets `AVCodecContext.b_frame_strategy` value.
func (avctx *AvCodecContext) GetBFrameStrategy() int32 {
	return (int32)(avctx.b_frame_strategy)
}

// Custom: SetBFrameStrategy sets `AVCodecContext.b_frame_strategy` value.
func (avctx *AvCodecContext) SetBFrameStrategy(v int32) {
	avctx.b_frame_strategy = (C.int)(v)
}

// Custom: GetBFrameStrategyAddr gets `AVCodecContext.b_frame_strategy` address.
func (avctx *AvCodecContext) GetBFrameStrategyAddr() *int32 {
	return (*int32)(&avctx.b_frame_strategy)
}

// Custom: Get BQuantOffset gets `AVCodecContext.b_quant_offset` value.
func (avctx *AvCodecContext) GetBQuantOffset() float32 {
	return (float32)(avctx.b_quant_offset)
}

// Custom: SetBQuantOffset sets `AVCodecContext.b_quant_offset` value.
func (avctx *AvCodecContext) SetBQuantOffset(v float32) {
	avctx.b_quant_offset = (C.float)(v)
}

// Custom: GetBQuantOffsetAddr gets `AVCodecContext.b_quant_offset` address.
func (avctx *AvCodecContext) GetBQuantOffsetAddr() *float32 {
	return (*float32)(&avctx.b_quant_offset)
}

// Custom: Get HasBFrames gets `AVCodecContext.has_b_frames` value.
func (avctx *AvCodecContext) GetHasBFrames() int32 {
	return (int32)(avctx.has_b_frames)
}

// Custom: SetHasBFrames sets `AVCodecContext.has_b_frames` value.
func (avctx *AvCodecContext) SetHasBFrames(v int32) {
	avctx.has_b_frames = (C.int)(v)
}

// Custom: GetHasBFramesAddr gets `AVCodecContext.has_b_frames` address.
func (avctx *AvCodecContext) GetHasBFramesAddr() *int32 {
	return (*int32)(&avctx.has_b_frames)
}

// Custom: Get MpegQuant gets `AVCodecContext.mpeg_quant` value.
func (avctx *AvCodecContext) GetMpegQuant() int32 {
	return (int32)(avctx.mpeg_quant)
}

// Custom: SetMpegQuant sets `AVCodecContext.mpeg_quant` value.
func (avctx *AvCodecContext) SetMpegQuant(v int32) {
	avctx.mpeg_quant = (C.int)(v)
}

// Custom: GetMpegQuantAddr gets `AVCodecContext.mpeg_quant` address.
func (avctx *AvCodecContext) GetMpegQuantAddr() *int32 {
	return (*int32)(&avctx.mpeg_quant)
}

// Custom: Get IQuantFactor gets `AVCodecContext.i_quant_factor` value.
func (avctx *AvCodecContext) GetIQuantFactor() float32 {
	return (float32)(avctx.i_quant_factor)
}

// Custom: SetIQuantFactor sets `AVCodecContext.i_quant_factor` value.
func (avctx *AvCodecContext) SetIQuantFactor(v float32) {
	avctx.i_quant_factor = (C.float)(v)
}

// Custom: GetIQuantFactorAddr gets `AVCodecContext.i_quant_factor` address.
func (avctx *AvCodecContext) GetIQuantFactorAddr() *float32 {
	return (*float32)(&avctx.i_quant_factor)
}

// Custom: Get IQuantOffset gets `AVCodecContext.i_quant_offset` value.
func (avctx *AvCodecContext) GetIQuantOffset() float32 {
	return (float32)(avctx.i_quant_offset)
}

// Custom: SetIQuantOffset sets `AVCodecContext.i_quant_offset` value.
func (avctx *AvCodecContext) SetIQuantOffset(v float32) {
	avctx.i_quant_offset = (C.float)(v)
}

// Custom: GetIQuantOffsetAddr gets `AVCodecContext.i_quant_offset` address.
func (avctx *AvCodecContext) GetIQuantOffsetAddr() *float32 {
	return (*float32)(&avctx.i_quant_offset)
}

// Custom: Get LumiMasking gets `AVCodecContext.lumi_masking` value.
func (avctx *AvCodecContext) GetLumiMasking() float32 {
	return (float32)(avctx.lumi_masking)
}

// Custom: SetLumiMasking sets `AVCodecContext.lumi_masking` value.
func (avctx *AvCodecContext) SetLumiMasking(v float32) {
	avctx.lumi_masking = (C.float)(v)
}

// Custom: GetLumiMaskingAddr gets `AVCodecContext.lumi_masking` address.
func (avctx *AvCodecContext) GetLumiMaskingAddr() *float32 {
	return (*float32)(&avctx.lumi_masking)
}

// Custom: Get TemporalCplxMasking gets `AVCodecContext.temporal_cplx_masking` value.
func (avctx *AvCodecContext) GetTemporalCplxMasking() float32 {
	return (float32)(avctx.temporal_cplx_masking)
}

// Custom: SetTemporalCplxMasking sets `AVCodecContext.temporal_cplx_masking` value.
func (avctx *AvCodecContext) SetTemporalCplxMasking(v float32) {
	avctx.temporal_cplx_masking = (C.float)(v)
}

// Custom: GetTemporalCplxMaskingAddr gets `AVCodecContext.temporal_cplx_masking` address.
func (avctx *AvCodecContext) GetTemporalCplxMaskingAddr() *float32 {
	return (*float32)(&avctx.temporal_cplx_masking)
}

// Custom: Get SpatialCplxMasking gets `AVCodecContext.spatial_cplx_masking` value.
func (avctx *AvCodecContext) GetSpatialCplxMasking() float32 {
	return (float32)(avctx.spatial_cplx_masking)
}

// Custom: SetSpatialCplxMasking sets `AVCodecContext.spatial_cplx_masking` value.
func (avctx *AvCodecContext) SetSpatialCplxMasking(v float32) {
	avctx.spatial_cplx_masking = (C.float)(v)
}

// Custom: GetSpatialCplxMaskingAddr gets `AVCodecContext.spatial_cplx_masking` address.
func (avctx *AvCodecContext) GetSpatialCplxMaskingAddr() *float32 {
	return (*float32)(&avctx.spatial_cplx_masking)
}

// Custom: Get PMasking gets `AVCodecContext.p_masking` value.
func (avctx *AvCodecContext) GetPMasking() float32 {
	return (float32)(avctx.p_masking)
}

// Custom: SetPMasking sets `AVCodecContext.p_masking` value.
func (avctx *AvCodecContext) SetPMasking(v float32) {
	avctx.p_masking = (C.float)(v)
}

// Custom: GetPMaskingAddr gets `AVCodecContext.p_masking` address.
func (avctx *AvCodecContext) GetPMaskingAddr() *float32 {
	return (*float32)(&avctx.p_masking)
}

// Custom: Get DarkMasking gets `AVCodecContext.dark_masking` value.
func (avctx *AvCodecContext) GetDarkMasking() float32 {
	return (float32)(avctx.dark_masking)
}

// Custom: SetDarkMasking sets `AVCodecContext.dark_masking` value.
func (avctx *AvCodecContext) SetDarkMasking(v float32) {
	avctx.dark_masking = (C.float)(v)
}

// Custom: GetDarkMaskingAddr gets `AVCodecContext.dark_masking` address.
func (avctx *AvCodecContext) GetDarkMaskingAddr() *float32 {
	return (*float32)(&avctx.dark_masking)
}

// Custom: Get SliceCount gets `AVCodecContext.slice_count` value.
func (avctx *AvCodecContext) GetSliceCount() int32 {
	return (int32)(avctx.slice_count)
}

// Custom: SetSliceCount sets `AVCodecContext.slice_count` value.
func (avctx *AvCodecContext) SetSliceCount(v int32) {
	avctx.slice_count = (C.int)(v)
}

// Custom: GetSliceCountAddr gets `AVCodecContext.slice_count` address.
func (avctx *AvCodecContext) GetSliceCountAddr() *int32 {
	return (*int32)(&avctx.slice_count)
}

// Custom: Get PredictionMethod gets `AVCodecContext.prediction_method` value.
func (avctx *AvCodecContext) GetPredictionMethod() int32 {
	return (int32)(avctx.prediction_method)
}

// Custom: SetPredictionMethod sets `AVCodecContext.prediction_method` value.
func (avctx *AvCodecContext) SetPredictionMethod(v int32) {
	avctx.prediction_method = (C.int)(v)
}

// Custom: GetPredictionMethodAddr gets `AVCodecContext.prediction_method` address.
func (avctx *AvCodecContext) GetPredictionMethodAddr() *int32 {
	return (*int32)(&avctx.prediction_method)
}

const (
	FF_PRED_LEFT   = int32(C.FF_PRED_LEFT)
	FF_PRED_PLANE  = int32(C.FF_PRED_PLANE)
	FF_PRED_MEDIAN = int32(C.FF_PRED_MEDIAN)
)

// Custom: Get SliceOffset gets `AVCodecContext.slice_offset` value.
func (avctx *AvCodecContext) GetSliceOffset() *int32 {
	return (*int32)(avctx.slice_offset)
}

// Custom: SetSliceOffset sets `AVCodecContext.slice_offset` value.
func (avctx *AvCodecContext) SetSliceOffset(v *int32) {
	avctx.slice_offset = (*C.int)(v)
}

// Custom: GetSliceOffsetAddr gets `AVCodecContext.slice_offset` address.
func (avctx *AvCodecContext) GetSliceOffsetAddr() **int32 {
	return (**int32)(unsafe.Pointer(&avctx.slice_offset))
}

// Custom: Get SampleAspectRatio gets `AVCodecContext.sample_aspect_ratio` value.
func (avctx *AvCodecContext) GetSampleAspectRatio() AvRational {
	return (AvRational)(avctx.sample_aspect_ratio)
}

// Custom: SetSampleAspectRatio sets `AVCodecContext.sample_aspect_ratio` value.
func (avctx *AvCodecContext) SetSampleAspectRatio(v AvRational) {
	avctx.sample_aspect_ratio = (C.AVRational)(v)
}

// Custom: GetSampleAspectRatioAddr gets `AVCodecContext.sample_aspect_ratio` address.
func (avctx *AvCodecContext) GetSampleAspectRatioAddr() *AvRational {
	return (*AvRational)(&avctx.sample_aspect_ratio)
}

// Custom: Get MeCmp gets `AVCodecContext.me_cmp` value.
func (avctx *AvCodecContext) GetMeCmp() int32 {
	return (int32)(avctx.me_cmp)
}

// Custom: SetMeCmp sets `AVCodecContext.me_cmp` value.
func (avctx *AvCodecContext) SetMeCmp(v int32) {
	avctx.me_cmp = (C.int)(v)
}

// Custom: GetMeCmpAddr gets `AVCodecContext.me_cmp` address.
func (avctx *AvCodecContext) GetMeCmpAddr() *int32 {
	return (*int32)(&avctx.me_cmp)
}

// Custom: Get MeSubCmp gets `AVCodecContext.me_sub_cmp` value.
func (avctx *AvCodecContext) GetMeSubCmp() int32 {
	return (int32)(avctx.me_sub_cmp)
}

// Custom: SetMeSubCmp sets `AVCodecContext.me_sub_cmp` value.
func (avctx *AvCodecContext) SetMeSubCmp(v int32) {
	avctx.me_sub_cmp = (C.int)(v)
}

// Custom: GetMeSubCmpAddr gets `AVCodecContext.me_sub_cmp` address.
func (avctx *AvCodecContext) GetMeSubCmpAddr() *int32 {
	return (*int32)(&avctx.me_sub_cmp)
}

// Custom: Get MbCmp gets `AVCodecContext.mb_cmp` value.
func (avctx *AvCodecContext) GetMbCmp() int32 {
	return (int32)(avctx.mb_cmp)
}

// Custom: SetMbCmp sets `AVCodecContext.mb_cmp` value.
func (avctx *AvCodecContext) SetMbCmp(v int32) {
	avctx.mb_cmp = (C.int)(v)
}

// Custom: GetMbCmpAddr gets `AVCodecContext.mb_cmp` address.
func (avctx *AvCodecContext) GetMbCmpAddr() *int32 {
	return (*int32)(&avctx.mb_cmp)
}

// Custom: Get IldctCmp gets `AVCodecContext.ildct_cmp` value.
func (avctx *AvCodecContext) GetIldctCmp() int32 {
	return (int32)(avctx.ildct_cmp)
}

// Custom: SetIldctCmp sets `AVCodecContext.ildct_cmp` value.
func (avctx *AvCodecContext) SetIldctCmp(v int32) {
	avctx.ildct_cmp = (C.int)(v)
}

// Custom: GetIldctCmpAddr gets `AVCodecContext.ildct_cmp` address.
func (avctx *AvCodecContext) GetIldctCmpAddr() *int32 {
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

// Custom: Get DiaSize gets `AVCodecContext.dia_size` value.
func (avctx *AvCodecContext) GetDiaSize() int32 {
	return (int32)(avctx.dia_size)
}

// Custom: SetDiaSize sets `AVCodecContext.dia_size` value.
func (avctx *AvCodecContext) SetDiaSize(v int32) {
	avctx.dia_size = (C.int)(v)
}

// Custom: GetDiaSizeAddr gets `AVCodecContext.dia_size` address.
func (avctx *AvCodecContext) GetDiaSizeAddr() *int32 {
	return (*int32)(&avctx.dia_size)
}

// Custom: Get LastPredictorCount gets `AVCodecContext.last_predictor_count` value.
func (avctx *AvCodecContext) GetLastPredictorCount() int32 {
	return (int32)(avctx.last_predictor_count)
}

// Custom: SetLastPredictorCount sets `AVCodecContext.last_predictor_count` value.
func (avctx *AvCodecContext) SetLastPredictorCount(v int32) {
	avctx.last_predictor_count = (C.int)(v)
}

// Custom: GetLastPredictorCountAddr gets `AVCodecContext.last_predictor_count` address.
func (avctx *AvCodecContext) GetLastPredictorCountAddr() *int32 {
	return (*int32)(&avctx.last_predictor_count)
}

// Custom: Get PreMe gets `AVCodecContext.pre_me` value.
func (avctx *AvCodecContext) GetPreMe() int32 {
	return (int32)(avctx.pre_me)
}

// Custom: SetPreMe sets `AVCodecContext.pre_me` value.
func (avctx *AvCodecContext) SetPreMe(v int32) {
	avctx.pre_me = (C.int)(v)
}

// Custom: GetPreMeAddr gets `AVCodecContext.pre_me` address.
func (avctx *AvCodecContext) GetPreMeAddr() *int32 {
	return (*int32)(&avctx.pre_me)
}

// Custom: Get MePreCmp gets `AVCodecContext.me_pre_cmp` value.
func (avctx *AvCodecContext) GetMePreCmp() int32 {
	return (int32)(avctx.me_pre_cmp)
}

// Custom: SetMePreCmp sets `AVCodecContext.me_pre_cmp` value.
func (avctx *AvCodecContext) SetMePreCmp(v int32) {
	avctx.me_pre_cmp = (C.int)(v)
}

// Custom: GetMePreCmpAddr gets `AVCodecContext.me_pre_cmp` address.
func (avctx *AvCodecContext) GetMePreCmpAddr() *int32 {
	return (*int32)(&avctx.me_pre_cmp)
}

// Custom: Get PreDiaSize gets `AVCodecContext.pre_dia_size` value.
func (avctx *AvCodecContext) GetPreDiaSize() int32 {
	return (int32)(avctx.pre_dia_size)
}

// Custom: SetPreDiaSize sets `AVCodecContext.pre_dia_size` value.
func (avctx *AvCodecContext) SetPreDiaSize(v int32) {
	avctx.pre_dia_size = (C.int)(v)
}

// Custom: GetPreDiaSizeAddr gets `AVCodecContext.pre_dia_size` address.
func (avctx *AvCodecContext) GetPreDiaSizeAddr() *int32 {
	return (*int32)(&avctx.pre_dia_size)
}

// Custom: Get MeSubpelQuality gets `AVCodecContext.me_subpel_quality` value.
func (avctx *AvCodecContext) GetMeSubpelQuality() int32 {
	return (int32)(avctx.me_subpel_quality)
}

// Custom: SetMeSubpelQuality sets `AVCodecContext.me_subpel_quality` value.
func (avctx *AvCodecContext) SetMeSubpelQuality(v int32) {
	avctx.me_subpel_quality = (C.int)(v)
}

// Custom: GetMeSubpelQualityAddr gets `AVCodecContext.me_subpel_quality` address.
func (avctx *AvCodecContext) GetMeSubpelQualityAddr() *int32 {
	return (*int32)(&avctx.me_subpel_quality)
}

// Custom: Get SliceFlags gets `AVCodecContext.slice_flags` value.
func (avctx *AvCodecContext) GetSliceFlags() int32 {
	return (int32)(avctx.slice_flags)
}

// Custom: SetSliceFlags sets `AVCodecContext.slice_flags` value.
func (avctx *AvCodecContext) SetSliceFlags(v int32) {
	avctx.slice_flags = (C.int)(v)
}

// Custom: GetSliceFlagsAddr gets `AVCodecContext.slice_flags` address.
func (avctx *AvCodecContext) GetSliceFlagsAddr() *int32 {
	return (*int32)(&avctx.slice_flags)
}

const (
	SLICE_FLAG_CODED_ORDER = int32(C.SLICE_FLAG_CODED_ORDER)
	SLICE_FLAG_ALLOW_FIELD = int32(C.SLICE_FLAG_ALLOW_FIELD)
	SLICE_FLAG_ALLOW_PLANE = int32(C.SLICE_FLAG_ALLOW_PLANE)
)

// Custom: Get MbDecision gets `AVCodecContext.mb_decision` value.
func (avctx *AvCodecContext) GetMbDecision() int32 {
	return (int32)(avctx.mb_decision)
}

// Custom: SetMbDecision sets `AVCodecContext.mb_decision` value.
func (avctx *AvCodecContext) SetMbDecision(v int32) {
	avctx.mb_decision = (C.int)(v)
}

// Custom: GetMbDecisionAddr gets `AVCodecContext.mb_decision` address.
func (avctx *AvCodecContext) GetMbDecisionAddr() *int32 {
	return (*int32)(&avctx.mb_decision)
}

const (
	FF_MB_DECISION_SIMPLE = int32(C.FF_MB_DECISION_SIMPLE)
	FF_MB_DECISION_BITS   = int32(C.FF_MB_DECISION_BITS)
	FF_MB_DECISION_RD     = int32(C.FF_MB_DECISION_RD)
)

// Custom: Get IntraMatrix gets `AVCodecContext.intra_matrix` value.
func (avctx *AvCodecContext) GetIntraMatrix() *uint16 {
	return (*uint16)(avctx.intra_matrix)
}

// Custom: SetIntraMatrix sets `AVCodecContext.intra_matrix` value.
func (avctx *AvCodecContext) SetIntraMatrix(v *uint16) {
	avctx.intra_matrix = (*C.uint16_t)(v)
}

// Custom: GetIntraMatrixAddr gets `AVCodecContext.intra_matrix` address.
func (avctx *AvCodecContext) GetIntraMatrixAddr() **uint16 {
	return (**uint16)(unsafe.Pointer(&avctx.intra_matrix))
}

// Custom: Get InterMatrix gets `AVCodecContext.inter_matrix` value.
func (avctx *AvCodecContext) GetInterMatrix() *uint16 {
	return (*uint16)(avctx.inter_matrix)
}

// Custom: SetInterMatrix sets `AVCodecContext.inter_matrix` value.
func (avctx *AvCodecContext) SetInterMatrix(v *uint16) {
	avctx.inter_matrix = (*C.uint16_t)(v)
}

// Custom: GetInterMatrixAddr gets `AVCodecContext.inter_matrix` address.
func (avctx *AvCodecContext) GetInterMatrixAddr() **uint16 {
	return (**uint16)(unsafe.Pointer(&avctx.inter_matrix))
}

// Custom: Get ScenechangeThreshold gets `AVCodecContext.scenechange_threshold` value.
func (avctx *AvCodecContext) GetScenechangeThreshold() int32 {
	return (int32)(avctx.scenechange_threshold)
}

// Custom: SetScenechangeThreshold sets `AVCodecContext.scenechange_threshold` value.
func (avctx *AvCodecContext) SetScenechangeThreshold(v int32) {
	avctx.scenechange_threshold = (C.int)(v)
}

// Custom: GetScenechangeThresholdAddr gets `AVCodecContext.scenechange_threshold` address.
func (avctx *AvCodecContext) GetScenechangeThresholdAddr() *int32 {
	return (*int32)(&avctx.scenechange_threshold)
}

// Custom: Get NoiseReduction gets `AVCodecContext.noise_reduction` value.
func (avctx *AvCodecContext) GetNoiseReduction() int32 {
	return (int32)(avctx.noise_reduction)
}

// Custom: SetNoiseReduction sets `AVCodecContext.noise_reduction` value.
func (avctx *AvCodecContext) SetNoiseReduction(v int32) {
	avctx.noise_reduction = (C.int)(v)
}

// Custom: GetNoiseReductionAddr gets `AVCodecContext.noise_reduction` address.
func (avctx *AvCodecContext) GetNoiseReductionAddr() *int32 {
	return (*int32)(&avctx.noise_reduction)
}

// Custom: Get IntraDcPrecision gets `AVCodecContext.intra_dc_precision` value.
func (avctx *AvCodecContext) GetIntraDcPrecision() int32 {
	return (int32)(avctx.intra_dc_precision)
}

// Custom: SetIntraDcPrecision sets `AVCodecContext.intra_dc_precision` value.
func (avctx *AvCodecContext) SetIntraDcPrecision(v int32) {
	avctx.intra_dc_precision = (C.int)(v)
}

// Custom: GetIntraDcPrecisionAddr gets `AVCodecContext.intra_dc_precision` address.
func (avctx *AvCodecContext) GetIntraDcPrecisionAddr() *int32 {
	return (*int32)(&avctx.intra_dc_precision)
}

// Custom: Get SkipTop gets `AVCodecContext.skip_top` value.
func (avctx *AvCodecContext) GetSkipTop() int32 {
	return (int32)(avctx.skip_top)
}

// Custom: SetSkipTop sets `AVCodecContext.skip_top` value.
func (avctx *AvCodecContext) SetSkipTop(v int32) {
	avctx.skip_top = (C.int)(v)
}

// Custom: GetSkipTopAddr gets `AVCodecContext.skip_top` address.
func (avctx *AvCodecContext) GetSkipTopAddr() *int32 {
	return (*int32)(&avctx.skip_top)
}

// Custom: Get SkipBottom gets `AVCodecContext.skip_bottom` value.
func (avctx *AvCodecContext) GetSkipBottom() int32 {
	return (int32)(avctx.skip_bottom)
}

// Custom: SetSkipBottom sets `AVCodecContext.skip_bottom` value.
func (avctx *AvCodecContext) SetSkipBottom(v int32) {
	avctx.skip_bottom = (C.int)(v)
}

// Custom: GetSkipBottomAddr gets `AVCodecContext.skip_bottom` address.
func (avctx *AvCodecContext) GetSkipBottomAddr() *int32 {
	return (*int32)(&avctx.skip_bottom)
}

// Custom: Get MbLmin gets `AVCodecContext.mb_lmin` value.
func (avctx *AvCodecContext) GetMbLmin() int32 {
	return (int32)(avctx.mb_lmin)
}

// Custom: SetMbLmin sets `AVCodecContext.mb_lmin` value.
func (avctx *AvCodecContext) SetMbLmin(v int32) {
	avctx.mb_lmin = (C.int)(v)
}

// Custom: GetMbLminAddr gets `AVCodecContext.mb_lmin` address.
func (avctx *AvCodecContext) GetMbLminAddr() *int32 {
	return (*int32)(&avctx.mb_lmin)
}

// Custom: Get MbLmax gets `AVCodecContext.mb_lmax` value.
func (avctx *AvCodecContext) GetMbLmax() int32 {
	return (int32)(avctx.mb_lmax)
}

// Custom: SetMbLmax sets `AVCodecContext.mb_lmax` value.
func (avctx *AvCodecContext) SetMbLmax(v int32) {
	avctx.mb_lmax = (C.int)(v)
}

// Custom: GetMbLmaxAddr gets `AVCodecContext.mb_lmax` address.
func (avctx *AvCodecContext) GetMbLmaxAddr() *int32 {
	return (*int32)(&avctx.mb_lmax)
}

// Custom: Get MePenaltyCompensation gets `AVCodecContext.me_penalty_compensation` value.
func (avctx *AvCodecContext) GetMePenaltyCompensation() int32 {
	return (int32)(avctx.me_penalty_compensation)
}

// Custom: SetMePenaltyCompensation sets `AVCodecContext.me_penalty_compensation` value.
func (avctx *AvCodecContext) SetMePenaltyCompensation(v int32) {
	avctx.me_penalty_compensation = (C.int)(v)
}

// Custom: GetMePenaltyCompensationAddr gets `AVCodecContext.me_penalty_compensation` address.
func (avctx *AvCodecContext) GetMePenaltyCompensationAddr() *int32 {
	return (*int32)(&avctx.me_penalty_compensation)
}

// Custom: Get BidirRefine gets `AVCodecContext.bidir_refine` value.
func (avctx *AvCodecContext) GetBidirRefine() int32 {
	return (int32)(avctx.bidir_refine)
}

// Custom: SetBidirRefine sets `AVCodecContext.bidir_refine` value.
func (avctx *AvCodecContext) SetBidirRefine(v int32) {
	avctx.bidir_refine = (C.int)(v)
}

// Custom: GetBidirRefineAddr gets `AVCodecContext.bidir_refine` address.
func (avctx *AvCodecContext) GetBidirRefineAddr() *int32 {
	return (*int32)(&avctx.bidir_refine)
}

// Custom: Get BrdScale gets `AVCodecContext.brd_scale` value.
func (avctx *AvCodecContext) GetBrdScale() int32 {
	return (int32)(avctx.brd_scale)
}

// Custom: SetBrdScale sets `AVCodecContext.brd_scale` value.
func (avctx *AvCodecContext) SetBrdScale(v int32) {
	avctx.brd_scale = (C.int)(v)
}

// Custom: GetBrdScaleAddr gets `AVCodecContext.brd_scale` address.
func (avctx *AvCodecContext) GetBrdScaleAddr() *int32 {
	return (*int32)(&avctx.brd_scale)
}

// Custom: Get KeyintMin gets `AVCodecContext.keyint_min` value.
func (avctx *AvCodecContext) GetKeyintMin() int32 {
	return (int32)(avctx.keyint_min)
}

// Custom: SetKeyintMin sets `AVCodecContext.keyint_min` value.
func (avctx *AvCodecContext) SetKeyintMin(v int32) {
	avctx.keyint_min = (C.int)(v)
}

// Custom: GetKeyintMinAddr gets `AVCodecContext.keyint_min` address.
func (avctx *AvCodecContext) GetKeyintMinAddr() *int32 {
	return (*int32)(&avctx.keyint_min)
}

// Custom: Get Refs gets `AVCodecContext.refs` value.
func (avctx *AvCodecContext) GetRefs() int32 {
	return (int32)(avctx.refs)
}

// Custom: SetRefs sets `AVCodecContext.refs` value.
func (avctx *AvCodecContext) SetRefs(v int32) {
	avctx.refs = (C.int)(v)
}

// Custom: GetRefsAddr gets `AVCodecContext.refs` address.
func (avctx *AvCodecContext) GetRefsAddr() *int32 {
	return (*int32)(&avctx.refs)
}

// Custom: Get Mv0Threshold gets `AVCodecContext.mv0_threshold` value.
func (avctx *AvCodecContext) GetMv0Threshold() int32 {
	return (int32)(avctx.mv0_threshold)
}

// Custom: SetMv0Threshold sets `AVCodecContext.mv0_threshold` value.
func (avctx *AvCodecContext) SetMv0Threshold(v int32) {
	avctx.mv0_threshold = (C.int)(v)
}

// Custom: GetMv0ThresholdAddr gets `AVCodecContext.mv0_threshold` address.
func (avctx *AvCodecContext) GetMv0ThresholdAddr() *int32 {
	return (*int32)(&avctx.mv0_threshold)
}

// Custom: Get BSensitivity gets `AVCodecContext.b_sensitivity` value.
func (avctx *AvCodecContext) GetBSensitivity() int32 {
	return (int32)(avctx.b_sensitivity)
}

// Custom: SetBSensitivity sets `AVCodecContext.b_sensitivity` value.
func (avctx *AvCodecContext) SetBSensitivity(v int32) {
	avctx.b_sensitivity = (C.int)(v)
}

// Custom: GetBSensitivityAddr gets `AVCodecContext.b_sensitivity` address.
func (avctx *AvCodecContext) GetBSensitivityAddr() *int32 {
	return (*int32)(&avctx.b_sensitivity)
}

// Custom: Get ColorPrimaries gets `AVCodecContext.color_primaries` value.
func (avctx *AvCodecContext) GetColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(avctx.color_primaries)
}

// Custom: SetColorPrimaries sets `AVCodecContext.color_primaries` value.
func (avctx *AvCodecContext) SetColorPrimaries(v AvColorPrimaries) {
	avctx.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// Custom: GetColorPrimariesAddr gets `AVCodecContext.color_primaries` address.
func (avctx *AvCodecContext) GetColorPrimariesAddr() *AvColorPrimaries {
	return (*AvColorPrimaries)(unsafe.Pointer(&avctx.color_primaries))
}

// Custom: Get ColorTrc gets `AVCodecContext.color_trc` value.
func (avctx *AvCodecContext) GetColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(avctx.color_trc)
}

// Custom: SetColorTrc sets `AVCodecContext.color_trc` value.
func (avctx *AvCodecContext) SetColorTrc(v AvColorTransferCharacteristic) {
	avctx.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// Custom: GetColorTrcAddr gets `AVCodecContext.color_trc` address.
func (avctx *AvCodecContext) GetColorTrcAddr() *AvColorTransferCharacteristic {
	return (*AvColorTransferCharacteristic)(unsafe.Pointer(&avctx.color_trc))
}

// Custom: Get Colorspace gets `AVCodecContext.colorspace` value.
func (avctx *AvCodecContext) GetColorspace() AvColorSpace {
	return (AvColorSpace)(avctx.colorspace)
}

// Custom: SetColorspace sets `AVCodecContext.colorspace` value.
func (avctx *AvCodecContext) SetColorspace(v AvColorSpace) {
	avctx.colorspace = (C.enum_AVColorSpace)(v)
}

// Custom: GetColorspaceAddr gets `AVCodecContext.colorspace` address.
func (avctx *AvCodecContext) GetColorspaceAddr() *AvColorSpace {
	return (*AvColorSpace)(unsafe.Pointer(&avctx.colorspace))
}

// Custom: Get ColorRange gets `AVCodecContext.color_range` value.
func (avctx *AvCodecContext) GetColorRange() AvColorRange {
	return (AvColorRange)(avctx.color_range)
}

// Custom: SetColorRange sets `AVCodecContext.color_range` value.
func (avctx *AvCodecContext) SetColorRange(v AvColorRange) {
	avctx.color_range = (C.enum_AVColorRange)(v)
}

// Custom: GetColorRangeAddr gets `AVCodecContext.color_range` address.
func (avctx *AvCodecContext) GetColorRangeAddr() *AvColorRange {
	return (*AvColorRange)(unsafe.Pointer(&avctx.color_range))
}

// Custom: Get ChromaSampleLocation gets `AVCodecContext.chroma_sample_location` value.
func (avctx *AvCodecContext) GetChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(avctx.chroma_sample_location)
}

// Custom: SetChromaSampleLocation sets `AVCodecContext.chroma_sample_location` value.
func (avctx *AvCodecContext) SetChromaSampleLocation(v AvChromaLocation) {
	avctx.chroma_sample_location = (C.enum_AVChromaLocation)(v)
}

// Custom: GetChromaSampleLocationAddr gets `AVCodecContext.chroma_sample_location` address.
func (avctx *AvCodecContext) GetChromaSampleLocationAddr() *AvChromaLocation {
	return (*AvChromaLocation)(unsafe.Pointer(&avctx.chroma_sample_location))
}

// Custom: Get Slices gets `AVCodecContext.slices` value.
func (avctx *AvCodecContext) GetSlices() int32 {
	return (int32)(avctx.slices)
}

// Custom: SetSlices sets `AVCodecContext.slices` value.
func (avctx *AvCodecContext) SetSlices(v int32) {
	avctx.slices = (C.int)(v)
}

// Custom: GetSlicesAddr gets `AVCodecContext.slices` address.
func (avctx *AvCodecContext) GetSlicesAddr() *int32 {
	return (*int32)(&avctx.slices)
}

// Custom: Get FieldOrder gets `AVCodecContext.field_order` value.
func (avctx *AvCodecContext) GetFieldOrder() AvFieldOrder {
	return (AvFieldOrder)(avctx.field_order)
}

// Custom: SetFieldOrder sets `AVCodecContext.field_order` value.
func (avctx *AvCodecContext) SetFieldOrder(v AvFieldOrder) {
	avctx.field_order = (C.enum_AVFieldOrder)(v)
}

// Custom: GetFieldOrderAddr gets `AVCodecContext.field_order` address.
func (avctx *AvCodecContext) GetFieldOrderAddr() *AvFieldOrder {
	return (*AvFieldOrder)(unsafe.Pointer(&avctx.field_order))
}

// Custom: Get SampleRate gets `AVCodecContext.sample_rate` value.
func (avctx *AvCodecContext) GetSampleRate() int32 {
	return (int32)(avctx.sample_rate)
}

// Custom: SetSampleRate sets `AVCodecContext.sample_rate` value.
func (avctx *AvCodecContext) SetSampleRate(v int32) {
	avctx.sample_rate = (C.int)(v)
}

// Custom: GetSampleRateAddr gets `AVCodecContext.sample_rate` address.
func (avctx *AvCodecContext) GetSampleRateAddr() *int32 {
	return (*int32)(&avctx.sample_rate)
}

// Custom: Get Channels gets `AVCodecContext.channels` value.
func (avctx *AvCodecContext) GetChannels() int32 {
	return (int32)(avctx.channels)
}

// Custom: SetChannels sets `AVCodecContext.channels` value.
func (avctx *AvCodecContext) SetChannels(v int32) {
	avctx.channels = (C.int)(v)
}

// Custom: GetChannelsAddr gets `AVCodecContext.channels` address.
func (avctx *AvCodecContext) GetChannelsAddr() *int32 {
	return (*int32)(&avctx.channels)
}

// Custom: Get SampleFmt gets `AVCodecContext.sample_fmt` value.
func (avctx *AvCodecContext) GetSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(avctx.sample_fmt)
}

// Custom: SetSampleFmt sets `AVCodecContext.sample_fmt` value.
func (avctx *AvCodecContext) SetSampleFmt(v AvSampleFormat) {
	avctx.sample_fmt = (C.enum_AVSampleFormat)(v)
}

// Custom: GetSampleFmtAddr gets `AVCodecContext.sample_fmt` address.
func (avctx *AvCodecContext) GetSampleFmtAddr() *AvSampleFormat {
	return (*AvSampleFormat)(&avctx.sample_fmt)
}

// Custom: Get FrameSize gets `AVCodecContext.frame_size` value.
func (avctx *AvCodecContext) GetFrameSize() int32 {
	return (int32)(avctx.frame_size)
}

// Custom: SetFrameSize sets `AVCodecContext.frame_size` value.
func (avctx *AvCodecContext) SetFrameSize(v int32) {
	avctx.frame_size = (C.int)(v)
}

// Custom: GetFrameSizeAddr gets `AVCodecContext.frame_size` address.
func (avctx *AvCodecContext) GetFrameSizeAddr() *int32 {
	return (*int32)(&avctx.frame_size)
}

// Custom: Get FrameNumber gets `AVCodecContext.frame_number` value.
func (avctx *AvCodecContext) GetFrameNumber() int32 {
	return (int32)(avctx.frame_number)
}

// Custom: SetFrameNumber sets `AVCodecContext.frame_number` value.
func (avctx *AvCodecContext) SetFrameNumber(v int32) {
	avctx.frame_number = (C.int)(v)
}

// Custom: GetFrameNumberAddr gets `AVCodecContext.frame_number` address.
func (avctx *AvCodecContext) GetFrameNumberAddr() *int32 {
	return (*int32)(&avctx.frame_number)
}

// Custom: Get BlockAlign gets `AVCodecContext.block_align` value.
func (avctx *AvCodecContext) GetBlockAlign() int32 {
	return (int32)(avctx.block_align)
}

// Custom: SetBlockAlign sets `AVCodecContext.block_align` value.
func (avctx *AvCodecContext) SetBlockAlign(v int32) {
	avctx.block_align = (C.int)(v)
}

// Custom: GetBlockAlignAddr gets `AVCodecContext.block_align` address.
func (avctx *AvCodecContext) GetBlockAlignAddr() *int32 {
	return (*int32)(&avctx.block_align)
}

// Custom: Get Cutoff gets `AVCodecContext.cutoff` value.
func (avctx *AvCodecContext) GetCutoff() int32 {
	return (int32)(avctx.cutoff)
}

// Custom: SetCutoff sets `AVCodecContext.cutoff` value.
func (avctx *AvCodecContext) SetCutoff(v int32) {
	avctx.cutoff = (C.int)(v)
}

// Custom: GetCutoffAddr gets `AVCodecContext.cutoff` address.
func (avctx *AvCodecContext) GetCutoffAddr() *int32 {
	return (*int32)(&avctx.cutoff)
}

// Custom: Get ChannelLayout gets `AVCodecContext.channel_layout` value.
func (avctx *AvCodecContext) GetChannelLayout() uint64 {
	return (uint64)(avctx.channel_layout)
}

// Custom: SetChannelLayout sets `AVCodecContext.channel_layout` value.
func (avctx *AvCodecContext) SetChannelLayout(v uint64) {
	avctx.channel_layout = (C.uint64_t)(v)
}

// Custom: GetChannelLayoutAddr gets `AVCodecContext.channel_layout` address.
func (avctx *AvCodecContext) GetChannelLayoutAddr() *uint64 {
	return (*uint64)(&avctx.channel_layout)
}

// Custom: Get RequestChannelLayout gets `AVCodecContext.request_channel_layout` value.
func (avctx *AvCodecContext) GetRequestChannelLayout() uint64 {
	return (uint64)(avctx.request_channel_layout)
}

// Custom: SetRequestChannelLayout sets `AVCodecContext.request_channel_layout` value.
func (avctx *AvCodecContext) SetRequestChannelLayout(v uint64) {
	avctx.request_channel_layout = (C.uint64_t)(v)
}

// Custom: GetRequestChannelLayoutAddr gets `AVCodecContext.request_channel_layout` address.
func (avctx *AvCodecContext) GetRequestChannelLayoutAddr() *uint64 {
	return (*uint64)(&avctx.request_channel_layout)
}

// Custom: Get AudioServiceType gets `AVCodecContext.audio_service_type` value.
func (avctx *AvCodecContext) GetAudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(avctx.audio_service_type)
}

// Custom: SetAudioServiceType sets `AVCodecContext.audio_service_type` value.
func (avctx *AvCodecContext) SetAudioServiceType(v AvAudioServiceType) {
	avctx.audio_service_type = (C.enum_AVAudioServiceType)(v)
}

// Custom: GetAudioServiceTypeAddr gets `AVCodecContext.audio_service_type` address.
func (avctx *AvCodecContext) GetAudioServiceTypeAddr() *AvAudioServiceType {
	return (*AvAudioServiceType)(unsafe.Pointer(&avctx.audio_service_type))
}

// Custom: Get RequestSampleFmt gets `AVCodecContext.request_sample_fmt` value.
func (avctx *AvCodecContext) GetRequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(avctx.request_sample_fmt)
}

// Custom: SetRequestSampleFmt sets `AVCodecContext.request_sample_fmt` value.
func (avctx *AvCodecContext) SetRequestSampleFmt(v AvSampleFormat) {
	avctx.request_sample_fmt = (C.enum_AVSampleFormat)(v)
}

// Custom: GetRequestSampleFmtAddr gets `AVCodecContext.request_sample_fmt` address.
func (avctx *AvCodecContext) GetRequestSampleFmtAddr() *AvSampleFormat {
	return (*AvSampleFormat)(&avctx.request_sample_fmt)
}

// Custom: Get RefcountedFrames gets `AVCodecContext.refcounted_frames` value.
func (avctx *AvCodecContext) GetRefcountedFrames() int32 {
	return (int32)(avctx.refcounted_frames)
}

// Custom: SetRefcountedFrames sets `AVCodecContext.refcounted_frames` value.
func (avctx *AvCodecContext) SetRefcountedFrames(v int32) {
	avctx.refcounted_frames = (C.int)(v)
}

// Custom: GetRefcountedFramesAddr gets `AVCodecContext.refcounted_frames` address.
func (avctx *AvCodecContext) GetRefcountedFramesAddr() *int32 {
	return (*int32)(&avctx.refcounted_frames)
}

// Custom: Get Qcompress gets `AVCodecContext.qcompress` value.
func (avctx *AvCodecContext) GetQcompress() float32 {
	return (float32)(avctx.qcompress)
}

// Custom: SetQcompress sets `AVCodecContext.qcompress` value.
func (avctx *AvCodecContext) SetQcompress(v float32) {
	avctx.qcompress = (C.float)(v)
}

// Custom: GetQcompressAddr gets `AVCodecContext.qcompress` address.
func (avctx *AvCodecContext) GetQcompressAddr() *float32 {
	return (*float32)(&avctx.qcompress)
}

// Custom: Get Qblur gets `AVCodecContext.qblur` value.
func (avctx *AvCodecContext) GetQblur() float32 {
	return (float32)(avctx.qblur)
}

// Custom: SetQblur sets `AVCodecContext.qblur` value.
func (avctx *AvCodecContext) SetQblur(v float32) {
	avctx.qblur = (C.float)(v)
}

// Custom: GetQblurAddr gets `AVCodecContext.qblur` address.
func (avctx *AvCodecContext) GetQblurAddr() *float32 {
	return (*float32)(&avctx.qblur)
}

// Custom: Get Qmin gets `AVCodecContext.qmin` value.
func (avctx *AvCodecContext) GetQmin() int32 {
	return (int32)(avctx.qmin)
}

// Custom: SetQmin sets `AVCodecContext.qmin` value.
func (avctx *AvCodecContext) SetQmin(v int32) {
	avctx.qmin = (C.int)(v)
}

// Custom: GetQminAddr gets `AVCodecContext.qmin` address.
func (avctx *AvCodecContext) GetQminAddr() *int32 {
	return (*int32)(&avctx.qmin)
}

// Custom: Get Qmax gets `AVCodecContext.qmax` value.
func (avctx *AvCodecContext) GetQmax() int32 {
	return (int32)(avctx.qmax)
}

// Custom: SetQmax sets `AVCodecContext.qmax` value.
func (avctx *AvCodecContext) SetQmax(v int32) {
	avctx.qmax = (C.int)(v)
}

// Custom: GetQmaxAddr gets `AVCodecContext.qmax` address.
func (avctx *AvCodecContext) GetQmaxAddr() *int32 {
	return (*int32)(&avctx.qmax)
}

// Custom: Get MaxQdiff gets `AVCodecContext.max_qdiff` value.
func (avctx *AvCodecContext) GetMaxQdiff() int32 {
	return (int32)(avctx.max_qdiff)
}

// Custom: SetMaxQdiff sets `AVCodecContext.max_qdiff` value.
func (avctx *AvCodecContext) SetMaxQdiff(v int32) {
	avctx.max_qdiff = (C.int)(v)
}

// Custom: GetMaxQdiffAddr gets `AVCodecContext.max_qdiff` address.
func (avctx *AvCodecContext) GetMaxQdiffAddr() *int32 {
	return (*int32)(&avctx.max_qdiff)
}

// Custom: Get RcBufferSize gets `AVCodecContext.rc_buffer_size` value.
func (avctx *AvCodecContext) GetRcBufferSize() int32 {
	return (int32)(avctx.rc_buffer_size)
}

// Custom: SetRcBufferSize sets `AVCodecContext.rc_buffer_size` value.
func (avctx *AvCodecContext) SetRcBufferSize(v int32) {
	avctx.rc_buffer_size = (C.int)(v)
}

// Custom: GetRcBufferSizeAddr gets `AVCodecContext.rc_buffer_size` address.
func (avctx *AvCodecContext) GetRcBufferSizeAddr() *int32 {
	return (*int32)(&avctx.rc_buffer_size)
}

// Custom: Get RcOverrideCount gets `AVCodecContext.rc_override_count` value.
func (avctx *AvCodecContext) GetRcOverrideCount() int32 {
	return (int32)(avctx.rc_override_count)
}

// Custom: SetRcOverrideCount sets `AVCodecContext.rc_override_count` value.
func (avctx *AvCodecContext) SetRcOverrideCount(v int32) {
	avctx.rc_override_count = (C.int)(v)
}

// Custom: GetRcOverrideCountAddr gets `AVCodecContext.rc_override_count` address.
func (avctx *AvCodecContext) GetRcOverrideCountAddr() *int32 {
	return (*int32)(&avctx.rc_override_count)
}

// Custom: Get RcOverride gets `AVCodecContext.rc_override` value.
func (avctx *AvCodecContext) GetRcOverride() *RcOverride {
	return (*RcOverride)(avctx.rc_override)
}

// Custom: SetRcOverride sets `AVCodecContext.rc_override` value.
func (avctx *AvCodecContext) SetRcOverride(v *RcOverride) {
	avctx.rc_override = (*C.RcOverride)(v)
}

// Custom: GetRcOverrideAddr gets `AVCodecContext.rc_override` address.
func (avctx *AvCodecContext) GetRcOverrideAddr() **RcOverride {
	return (**RcOverride)(unsafe.Pointer(&avctx.rc_override))
}

// Custom: Get RcMaxRate gets `AVCodecContext.rc_max_rate` value.
func (avctx *AvCodecContext) GetRcMaxRate() int64 {
	return (int64)(avctx.rc_max_rate)
}

// Custom: SetRcMaxRate sets `AVCodecContext.rc_max_rate` value.
func (avctx *AvCodecContext) SetRcMaxRate(v int64) {
	avctx.rc_max_rate = (C.int64_t)(v)
}

// Custom: GetRcMaxRateAddr gets `AVCodecContext.rc_max_rate` address.
func (avctx *AvCodecContext) GetRcMaxRateAddr() *int64 {
	return (*int64)(&avctx.rc_max_rate)
}

// Custom: Get RcMinRate gets `AVCodecContext.rc_min_rate` value.
func (avctx *AvCodecContext) GetRcMinRate() int64 {
	return (int64)(avctx.rc_min_rate)
}

// Custom: SetRcMinRate sets `AVCodecContext.rc_min_rate` value.
func (avctx *AvCodecContext) SetRcMinRate(v int64) {
	avctx.rc_min_rate = (C.int64_t)(v)
}

// Custom: GetRcMinRateAddr gets `AVCodecContext.rc_min_rate` address.
func (avctx *AvCodecContext) GetRcMinRateAddr() *int64 {
	return (*int64)(&avctx.rc_min_rate)
}

// Custom: Get RcMaxAvailableVbvUse gets `AVCodecContext.rc_max_available_vbv_use` value.
func (avctx *AvCodecContext) GetRcMaxAvailableVbvUse() float32 {
	return (float32)(avctx.rc_max_available_vbv_use)
}

// Custom: SetRcMaxAvailableVbvUse sets `AVCodecContext.rc_max_available_vbv_use` value.
func (avctx *AvCodecContext) SetRcMaxAvailableVbvUse(v float32) {
	avctx.rc_max_available_vbv_use = (C.float)(v)
}

// Custom: GetRcMaxAvailableVbvUseAddr gets `AVCodecContext.rc_max_available_vbv_use` address.
func (avctx *AvCodecContext) GetRcMaxAvailableVbvUseAddr() *float32 {
	return (*float32)(&avctx.rc_max_available_vbv_use)
}

// Custom: Get RcMinVbvOverflowUse gets `AVCodecContext.rc_min_vbv_overflow_use` value.
func (avctx *AvCodecContext) GetRcMinVbvOverflowUse() float32 {
	return (float32)(avctx.rc_min_vbv_overflow_use)
}

// Custom: SetRcMinVbvOverflowUse sets `AVCodecContext.rc_min_vbv_overflow_use` value.
func (avctx *AvCodecContext) SetRcMinVbvOverflowUse(v float32) {
	avctx.rc_min_vbv_overflow_use = (C.float)(v)
}

// Custom: GetRcMinVbvOverflowUseAddr gets `AVCodecContext.rc_min_vbv_overflow_use` address.
func (avctx *AvCodecContext) GetRcMinVbvOverflowUseAddr() *float32 {
	return (*float32)(&avctx.rc_min_vbv_overflow_use)
}

// Custom: Get RcInitialBufferOccupancy gets `AVCodecContext.rc_initial_buffer_occupancy` value.
func (avctx *AvCodecContext) GetRcInitialBufferOccupancy() int32 {
	return (int32)(avctx.rc_initial_buffer_occupancy)
}

// Custom: SetRcInitialBufferOccupancy sets `AVCodecContext.rc_initial_buffer_occupancy` value.
func (avctx *AvCodecContext) SetRcInitialBufferOccupancy(v int32) {
	avctx.rc_initial_buffer_occupancy = (C.int)(v)
}

// Custom: GetRcInitialBufferOccupancyAddr gets `AVCodecContext.rc_initial_buffer_occupancy` address.
func (avctx *AvCodecContext) GetRcInitialBufferOccupancyAddr() *int32 {
	return (*int32)(&avctx.rc_initial_buffer_occupancy)
}

const (
	FF_CODER_TYPE_VLC = int32(C.FF_CODER_TYPE_VLC)
	FF_CODER_TYPE_AC  = int32(C.FF_CODER_TYPE_AC)
	FF_CODER_TYPE_RAW = int32(C.FF_CODER_TYPE_RAW)
	FF_CODER_TYPE_RLE = int32(C.FF_CODER_TYPE_RLE)
)

// Custom: Get CoderType gets `AVCodecContext.coder_type` value.
func (avctx *AvCodecContext) GetCoderType() int32 {
	return (int32)(avctx.coder_type)
}

// Custom: SetCoderType sets `AVCodecContext.coder_type` value.
func (avctx *AvCodecContext) SetCoderType(v int32) {
	avctx.coder_type = (C.int)(v)
}

// Custom: GetCoderTypeAddr gets `AVCodecContext.coder_type` address.
func (avctx *AvCodecContext) GetCoderTypeAddr() *int32 {
	return (*int32)(&avctx.coder_type)
}

// Custom: Get ContextModel gets `AVCodecContext.context_model` value.
func (avctx *AvCodecContext) GetContextModel() int32 {
	return (int32)(avctx.context_model)
}

// Custom: SetContextModel sets `AVCodecContext.context_model` value.
func (avctx *AvCodecContext) SetContextModel(v int32) {
	avctx.context_model = (C.int)(v)
}

// Custom: GetContextModelAddr gets `AVCodecContext.context_model` address.
func (avctx *AvCodecContext) GetContextModelAddr() *int32 {
	return (*int32)(&avctx.context_model)
}

// Custom: Get FrameSkipThreshold gets `AVCodecContext.frame_skip_threshold` value.
func (avctx *AvCodecContext) GetFrameSkipThreshold() int32 {
	return (int32)(avctx.frame_skip_threshold)
}

// Custom: SetFrameSkipThreshold sets `AVCodecContext.frame_skip_threshold` value.
func (avctx *AvCodecContext) SetFrameSkipThreshold(v int32) {
	avctx.frame_skip_threshold = (C.int)(v)
}

// Custom: GetFrameSkipThresholdAddr gets `AVCodecContext.frame_skip_threshold` address.
func (avctx *AvCodecContext) GetFrameSkipThresholdAddr() *int32 {
	return (*int32)(&avctx.frame_skip_threshold)
}

// Custom: Get FrameSkipFactor gets `AVCodecContext.frame_skip_factor` value.
func (avctx *AvCodecContext) GetFrameSkipFactor() int32 {
	return (int32)(avctx.frame_skip_factor)
}

// Custom: SetFrameSkipFactor sets `AVCodecContext.frame_skip_factor` value.
func (avctx *AvCodecContext) SetFrameSkipFactor(v int32) {
	avctx.frame_skip_factor = (C.int)(v)
}

// Custom: GetFrameSkipFactorAddr gets `AVCodecContext.frame_skip_factor` address.
func (avctx *AvCodecContext) GetFrameSkipFactorAddr() *int32 {
	return (*int32)(&avctx.frame_skip_factor)
}

// Custom: Get FrameSkipExp gets `AVCodecContext.frame_skip_exp` value.
func (avctx *AvCodecContext) GetFrameSkipExp() int32 {
	return (int32)(avctx.frame_skip_exp)
}

// Custom: SetFrameSkipExp sets `AVCodecContext.frame_skip_exp` value.
func (avctx *AvCodecContext) SetFrameSkipExp(v int32) {
	avctx.frame_skip_exp = (C.int)(v)
}

// Custom: GetFrameSkipExpAddr gets `AVCodecContext.frame_skip_exp` address.
func (avctx *AvCodecContext) GetFrameSkipExpAddr() *int32 {
	return (*int32)(&avctx.frame_skip_exp)
}

// Custom: Get FrameSkipCmp gets `AVCodecContext.frame_skip_cmp` value.
func (avctx *AvCodecContext) GetFrameSkipCmp() int32 {
	return (int32)(avctx.frame_skip_cmp)
}

// Custom: SetFrameSkipCmp sets `AVCodecContext.frame_skip_cmp` value.
func (avctx *AvCodecContext) SetFrameSkipCmp(v int32) {
	avctx.frame_skip_cmp = (C.int)(v)
}

// Custom: GetFrameSkipCmpAddr gets `AVCodecContext.frame_skip_cmp` address.
func (avctx *AvCodecContext) GetFrameSkipCmpAddr() *int32 {
	return (*int32)(&avctx.frame_skip_cmp)
}

// Custom: Get Trellis gets `AVCodecContext.trellis` value.
func (avctx *AvCodecContext) GetTrellis() int32 {
	return (int32)(avctx.trellis)
}

// Custom: SetTrellis sets `AVCodecContext.trellis` value.
func (avctx *AvCodecContext) SetTrellis(v int32) {
	avctx.trellis = (C.int)(v)
}

// Custom: GetTrellisAddr gets `AVCodecContext.trellis` address.
func (avctx *AvCodecContext) GetTrellisAddr() *int32 {
	return (*int32)(&avctx.trellis)
}

// Custom: Get MinPredictionOrder gets `AVCodecContext.min_prediction_order` value.
func (avctx *AvCodecContext) GetMinPredictionOrder() int32 {
	return (int32)(avctx.min_prediction_order)
}

// Custom: SetMinPredictionOrder sets `AVCodecContext.min_prediction_order` value.
func (avctx *AvCodecContext) SetMinPredictionOrder(v int32) {
	avctx.min_prediction_order = (C.int)(v)
}

// Custom: GetMinPredictionOrderAddr gets `AVCodecContext.min_prediction_order` address.
func (avctx *AvCodecContext) GetMinPredictionOrderAddr() *int32 {
	return (*int32)(&avctx.min_prediction_order)
}

// Custom: Get MaxPredictionOrder gets `AVCodecContext.max_prediction_order` value.
func (avctx *AvCodecContext) GetMaxPredictionOrder() int32 {
	return (int32)(avctx.max_prediction_order)
}

// Custom: SetMaxPredictionOrder sets `AVCodecContext.max_prediction_order` value.
func (avctx *AvCodecContext) SetMaxPredictionOrder(v int32) {
	avctx.max_prediction_order = (C.int)(v)
}

// Custom: GetMaxPredictionOrderAddr gets `AVCodecContext.max_prediction_order` address.
func (avctx *AvCodecContext) GetMaxPredictionOrderAddr() *int32 {
	return (*int32)(&avctx.max_prediction_order)
}

// Custom: Get TimecodeFrameStart gets `AVCodecContext.timecode_frame_start` value.
func (avctx *AvCodecContext) GetTimecodeFrameStart() int64 {
	return (int64)(avctx.timecode_frame_start)
}

// Custom: SetTimecodeFrameStart sets `AVCodecContext.timecode_frame_start` value.
func (avctx *AvCodecContext) SetTimecodeFrameStart(v int64) {
	avctx.timecode_frame_start = (C.int64_t)(v)
}

// Custom: GetTimecodeFrameStartAddr gets `AVCodecContext.timecode_frame_start` address.
func (avctx *AvCodecContext) GetTimecodeFrameStartAddr() *int64 {
	return (*int64)(&avctx.timecode_frame_start)
}

// Custom: Get RtpPayloadSize gets `AVCodecContext.rtp_payload_size` value.
func (avctx *AvCodecContext) GetRtpPayloadSize() int32 {
	return (int32)(avctx.rtp_payload_size)
}

// Custom: SetRtpPayloadSize sets `AVCodecContext.rtp_payload_size` value.
func (avctx *AvCodecContext) SetRtpPayloadSize(v int32) {
	avctx.rtp_payload_size = (C.int)(v)
}

// Custom: GetRtpPayloadSizeAddr gets `AVCodecContext.rtp_payload_size` address.
func (avctx *AvCodecContext) GetRtpPayloadSizeAddr() *int32 {
	return (*int32)(&avctx.rtp_payload_size)
}

// Custom: Get MvBits gets `AVCodecContext.mv_bits` value.
func (avctx *AvCodecContext) GetMvBits() int32 {
	return (int32)(avctx.mv_bits)
}

// Custom: SetMvBits sets `AVCodecContext.mv_bits` value.
func (avctx *AvCodecContext) SetMvBits(v int32) {
	avctx.mv_bits = (C.int)(v)
}

// Custom: GetMvBitsAddr gets `AVCodecContext.mv_bits` address.
func (avctx *AvCodecContext) GetMvBitsAddr() *int32 {
	return (*int32)(&avctx.mv_bits)
}

// Custom: Get HeaderBits gets `AVCodecContext.header_bits` value.
func (avctx *AvCodecContext) GetHeaderBits() int32 {
	return (int32)(avctx.header_bits)
}

// Custom: SetHeaderBits sets `AVCodecContext.header_bits` value.
func (avctx *AvCodecContext) SetHeaderBits(v int32) {
	avctx.header_bits = (C.int)(v)
}

// Custom: GetHeaderBitsAddr gets `AVCodecContext.header_bits` address.
func (avctx *AvCodecContext) GetHeaderBitsAddr() *int32 {
	return (*int32)(&avctx.header_bits)
}

// Custom: Get ITexBits gets `AVCodecContext.i_tex_bits` value.
func (avctx *AvCodecContext) GetITexBits() int32 {
	return (int32)(avctx.i_tex_bits)
}

// Custom: SetITexBits sets `AVCodecContext.i_tex_bits` value.
func (avctx *AvCodecContext) SetITexBits(v int32) {
	avctx.i_tex_bits = (C.int)(v)
}

// Custom: GetITexBitsAddr gets `AVCodecContext.i_tex_bits` address.
func (avctx *AvCodecContext) GetITexBitsAddr() *int32 {
	return (*int32)(&avctx.i_tex_bits)
}

// Custom: Get PTexBits gets `AVCodecContext.p_tex_bits` value.
func (avctx *AvCodecContext) GetPTexBits() int32 {
	return (int32)(avctx.p_tex_bits)
}

// Custom: SetPTexBits sets `AVCodecContext.p_tex_bits` value.
func (avctx *AvCodecContext) SetPTexBits(v int32) {
	avctx.p_tex_bits = (C.int)(v)
}

// Custom: GetPTexBitsAddr gets `AVCodecContext.p_tex_bits` address.
func (avctx *AvCodecContext) GetPTexBitsAddr() *int32 {
	return (*int32)(&avctx.p_tex_bits)
}

// Custom: Get ICount gets `AVCodecContext.i_count` value.
func (avctx *AvCodecContext) GetICount() int32 {
	return (int32)(avctx.i_count)
}

// Custom: SetICount sets `AVCodecContext.i_count` value.
func (avctx *AvCodecContext) SetICount(v int32) {
	avctx.i_count = (C.int)(v)
}

// Custom: GetICountAddr gets `AVCodecContext.i_count` address.
func (avctx *AvCodecContext) GetICountAddr() *int32 {
	return (*int32)(&avctx.i_count)
}

// Custom: Get PCount gets `AVCodecContext.p_count` value.
func (avctx *AvCodecContext) GetPCount() int32 {
	return (int32)(avctx.p_count)
}

// Custom: SetPCount sets `AVCodecContext.p_count` value.
func (avctx *AvCodecContext) SetPCount(v int32) {
	avctx.p_count = (C.int)(v)
}

// Custom: GetPCountAddr gets `AVCodecContext.p_count` address.
func (avctx *AvCodecContext) GetPCountAddr() *int32 {
	return (*int32)(&avctx.p_count)
}

// Custom: Get SkipCount gets `AVCodecContext.skip_count` value.
func (avctx *AvCodecContext) GetSkipCount() int32 {
	return (int32)(avctx.skip_count)
}

// Custom: SetSkipCount sets `AVCodecContext.skip_count` value.
func (avctx *AvCodecContext) SetSkipCount(v int32) {
	avctx.skip_count = (C.int)(v)
}

// Custom: GetSkipCountAddr gets `AVCodecContext.skip_count` address.
func (avctx *AvCodecContext) GetSkipCountAddr() *int32 {
	return (*int32)(&avctx.skip_count)
}

// Custom: Get MiscBits gets `AVCodecContext.misc_bits` value.
func (avctx *AvCodecContext) GetMiscBits() int32 {
	return (int32)(avctx.misc_bits)
}

// Custom: SetMiscBits sets `AVCodecContext.misc_bits` value.
func (avctx *AvCodecContext) SetMiscBits(v int32) {
	avctx.misc_bits = (C.int)(v)
}

// Custom: GetMiscBitsAddr gets `AVCodecContext.misc_bits` address.
func (avctx *AvCodecContext) GetMiscBitsAddr() *int32 {
	return (*int32)(&avctx.misc_bits)
}

// Custom: Get FrameBits gets `AVCodecContext.frame_bits` value.
func (avctx *AvCodecContext) GetFrameBits() int32 {
	return (int32)(avctx.frame_bits)
}

// Custom: SetFrameBits sets `AVCodecContext.frame_bits` value.
func (avctx *AvCodecContext) SetFrameBits(v int32) {
	avctx.frame_bits = (C.int)(v)
}

// Custom: GetFrameBitsAddr gets `AVCodecContext.frame_bits` address.
func (avctx *AvCodecContext) GetFrameBitsAddr() *int32 {
	return (*int32)(&avctx.frame_bits)
}

// Custom: Get StatsOut gets `AVCodecContext.stats_out` value.
func (avctx *AvCodecContext) GetStatsOut() string {
	return C.GoString(avctx.stats_out)
}

// Custom: Get StatsIn gets `AVCodecContext.stats_in` value.
func (avctx *AvCodecContext) GetStatsIn() string {
	return C.GoString(avctx.stats_in)
}

// Custom: Get WorkaroundBugs gets `AVCodecContext.workaround_bugs` value.
func (avctx *AvCodecContext) GetWorkaroundBugs() int32 {
	return (int32)(avctx.workaround_bugs)
}

// Custom: SetWorkaroundBugs sets `AVCodecContext.workaround_bugs` value.
func (avctx *AvCodecContext) SetWorkaroundBugs(v int32) {
	avctx.workaround_bugs = (C.int)(v)
}

// Custom: GetWorkaroundBugsAddr gets `AVCodecContext.workaround_bugs` address.
func (avctx *AvCodecContext) GetWorkaroundBugsAddr() *int32 {
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

// Custom: Get StrictStdCompliance gets `AVCodecContext.strict_std_compliance` value.
func (avctx *AvCodecContext) GetStrictStdCompliance() int32 {
	return (int32)(avctx.strict_std_compliance)
}

// Custom: SetStrictStdCompliance sets `AVCodecContext.strict_std_compliance` value.
func (avctx *AvCodecContext) SetStrictStdCompliance(v int32) {
	avctx.strict_std_compliance = (C.int)(v)
}

// Custom: GetStrictStdComplianceAddr gets `AVCodecContext.strict_std_compliance` address.
func (avctx *AvCodecContext) GetStrictStdComplianceAddr() *int32 {
	return (*int32)(&avctx.strict_std_compliance)
}

const (
	FF_COMPLIANCE_VERY_STRICT  = int32(C.FF_COMPLIANCE_VERY_STRICT)
	FF_COMPLIANCE_STRICT       = int32(C.FF_COMPLIANCE_STRICT)
	FF_COMPLIANCE_NORMAL       = int32(C.FF_COMPLIANCE_NORMAL)
	FF_COMPLIANCE_UNOFFICIAL   = int32(C.FF_COMPLIANCE_UNOFFICIAL)
	FF_COMPLIANCE_EXPERIMENTAL = int32(C.FF_COMPLIANCE_EXPERIMENTAL)
)

// Custom: Get ErrorConcealment gets `AVCodecContext.error_concealment` value.
func (avctx *AvCodecContext) GetErrorConcealment() int32 {
	return (int32)(avctx.error_concealment)
}

// Custom: SetErrorConcealment sets `AVCodecContext.error_concealment` value.
func (avctx *AvCodecContext) SetErrorConcealment(v int32) {
	avctx.error_concealment = (C.int)(v)
}

// Custom: GetErrorConcealmentAddr gets `AVCodecContext.error_concealment` address.
func (avctx *AvCodecContext) GetErrorConcealmentAddr() *int32 {
	return (*int32)(&avctx.error_concealment)
}

const (
	FF_EC_GUESS_MVS   = int32(C.FF_EC_GUESS_MVS)
	FF_EC_DEBLOCK     = int32(C.FF_EC_DEBLOCK)
	FF_EC_FAVOR_INTER = int32(C.FF_EC_FAVOR_INTER)
)

// Custom: Get Debug gets `AVCodecContext.debug` value.
func (avctx *AvCodecContext) GetDebug() int32 {
	return (int32)(avctx.debug)
}

// Custom: SetDebug sets `AVCodecContext.debug` value.
func (avctx *AvCodecContext) SetDebug(v int32) {
	avctx.debug = (C.int)(v)
}

// Custom: GetDebugAddr gets `AVCodecContext.debug` address.
func (avctx *AvCodecContext) GetDebugAddr() *int32 {
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

// Custom: Get ErrRecognition gets `AVCodecContext.err_recognition` value.
func (avctx *AvCodecContext) GetErrRecognition() int32 {
	return (int32)(avctx.err_recognition)
}

// Custom: SetErrRecognition sets `AVCodecContext.err_recognition` value.
func (avctx *AvCodecContext) SetErrRecognition(v int32) {
	avctx.err_recognition = (C.int)(v)
}

// Custom: GetErrRecognitionAddr gets `AVCodecContext.err_recognition` address.
func (avctx *AvCodecContext) GetErrRecognitionAddr() *int32 {
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

// Custom: Get ReorderedOpaque gets `AVCodecContext.reordered_opaque` value.
func (avctx *AvCodecContext) GetReorderedOpaque() int64 {
	return (int64)(avctx.reordered_opaque)
}

// Custom: SetReorderedOpaque sets `AVCodecContext.reordered_opaque` value.
func (avctx *AvCodecContext) SetReorderedOpaque(v int64) {
	avctx.reordered_opaque = (C.int64_t)(v)
}

// Custom: GetReorderedOpaqueAddr gets `AVCodecContext.reordered_opaque` address.
func (avctx *AvCodecContext) GetReorderedOpaqueAddr() *int64 {
	return (*int64)(&avctx.reordered_opaque)
}

// Custom: Get Hwaccel gets `AVCodecContext.hwaccel` value.
func (avctx *AvCodecContext) GetHwaccel() *AvHWAccel {
	return (*AvHWAccel)(avctx.hwaccel)
}

// Custom: SetHwaccel sets `AVCodecContext.hwaccel` value.
func (avctx *AvCodecContext) SetHwaccel(v *AvHWAccel) {
	avctx.hwaccel = (*C.AVHWAccel)(v)
}

// Custom: GetHwaccelAddr gets `AVCodecContext.hwaccel` address.
func (avctx *AvCodecContext) GetHwaccelAddr() **AvHWAccel {
	return (**AvHWAccel)(unsafe.Pointer(&avctx.hwaccel))
}

// Custom: Get HwaccelContext gets `AVCodecContext.hwaccel_context` value.
func (avctx *AvCodecContext) GetHwaccelContext() unsafe.Pointer {
	return (unsafe.Pointer)(avctx.hwaccel_context)
}

// Custom: SetHwaccelContext sets `AVCodecContext.hwaccel_context` value.
func (avctx *AvCodecContext) SetHwaccelContext(v unsafe.Pointer) {
	avctx.hwaccel_context = v
}

// Custom: GetHwaccelContextAddr gets `AVCodecContext.hwaccel_context` address.
func (avctx *AvCodecContext) GetHwaccelContextAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&avctx.hwaccel_context)
}

// Custom: Get DctAlgo gets `AVCodecContext.dct_algo` value.
func (avctx *AvCodecContext) GetDctAlgo() int32 {
	return (int32)(avctx.dct_algo)
}

// Custom: SetDctAlgo sets `AVCodecContext.dct_algo` value.
func (avctx *AvCodecContext) SetDctAlgo(v int32) {
	avctx.dct_algo = (C.int)(v)
}

// Custom: GetDctAlgoAddr gets `AVCodecContext.dct_algo` address.
func (avctx *AvCodecContext) GetDctAlgoAddr() *int32 {
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

// Custom: Get IdctAlgo gets `AVCodecContext.idct_algo` value.
func (avctx *AvCodecContext) GetIdctAlgo() int32 {
	return (int32)(avctx.idct_algo)
}

// Custom: SetIdctAlgo sets `AVCodecContext.idct_algo` value.
func (avctx *AvCodecContext) SetIdctAlgo(v int32) {
	avctx.idct_algo = (C.int)(v)
}

// Custom: GetIdctAlgoAddr gets `AVCodecContext.idct_algo` address.
func (avctx *AvCodecContext) GetIdctAlgoAddr() *int32 {
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

// Custom: Get BitsPerCodedSample gets `AVCodecContext.bits_per_coded_sample` value.
func (avctx *AvCodecContext) GetBitsPerCodedSample() int32 {
	return (int32)(avctx.bits_per_coded_sample)
}

// Custom: SetBitsPerCodedSample sets `AVCodecContext.bits_per_coded_sample` value.
func (avctx *AvCodecContext) SetBitsPerCodedSample(v int32) {
	avctx.bits_per_coded_sample = (C.int)(v)
}

// Custom: GetBitsPerCodedSampleAddr gets `AVCodecContext.bits_per_coded_sample` address.
func (avctx *AvCodecContext) GetBitsPerCodedSampleAddr() *int32 {
	return (*int32)(&avctx.bits_per_coded_sample)
}

// Custom: Get BitsPerRawSample gets `AVCodecContext.bits_per_raw_sample` value.
func (avctx *AvCodecContext) GetBitsPerRawSample() int32 {
	return (int32)(avctx.bits_per_raw_sample)
}

// Custom: SetBitsPerRawSample sets `AVCodecContext.bits_per_raw_sample` value.
func (avctx *AvCodecContext) SetBitsPerRawSample(v int32) {
	avctx.bits_per_raw_sample = (C.int)(v)
}

// Custom: GetBitsPerRawSampleAddr gets `AVCodecContext.bits_per_raw_sample` address.
func (avctx *AvCodecContext) GetBitsPerRawSampleAddr() *int32 {
	return (*int32)(&avctx.bits_per_raw_sample)
}

// Custom: Get Lowres gets `AVCodecContext.lowres` value.
func (avctx *AvCodecContext) GetLowres() int32 {
	return (int32)(avctx.lowres)
}

// Custom: SetLowres sets `AVCodecContext.lowres` value.
func (avctx *AvCodecContext) SetLowres(v int32) {
	avctx.lowres = (C.int)(v)
}

// Custom: GetLowresAddr gets `AVCodecContext.lowres` address.
func (avctx *AvCodecContext) GetLowresAddr() *int32 {
	return (*int32)(&avctx.lowres)
}

// Custom: Get CodedFrame gets `AVCodecContext.coded_frame` value.
func (avctx *AvCodecContext) GetCodedFrame() *AvFrame {
	return (*AvFrame)(avctx.coded_frame)
}

// Custom: SetCodedFrame sets `AVCodecContext.coded_frame` value.
func (avctx *AvCodecContext) SetCodedFrame(v *AvFrame) {
	avctx.coded_frame = (*C.struct_AVFrame)(v)
}

// Custom: GetCodedFrameAddr gets `AVCodecContext.coded_frame` address.
func (avctx *AvCodecContext) GetCodedFrameAddr() **AvFrame {
	return (**AvFrame)(unsafe.Pointer(&avctx.coded_frame))
}

// Custom: Get ThreadCount gets `AVCodecContext.thread_count` value.
func (avctx *AvCodecContext) GetThreadCount() int32 {
	return (int32)(avctx.thread_count)
}

// Custom: SetThreadCount sets `AVCodecContext.thread_count` value.
func (avctx *AvCodecContext) SetThreadCount(v int32) {
	avctx.thread_count = (C.int)(v)
}

// Custom: GetThreadCountAddr gets `AVCodecContext.thread_count` address.
func (avctx *AvCodecContext) GetThreadCountAddr() *int32 {
	return (*int32)(&avctx.thread_count)
}

// Custom: Get ThreadType gets `AVCodecContext.thread_type` value.
func (avctx *AvCodecContext) GetThreadType() int32 {
	return (int32)(avctx.thread_type)
}

// Custom: SetThreadType sets `AVCodecContext.thread_type` value.
func (avctx *AvCodecContext) SetThreadType(v int32) {
	avctx.thread_type = (C.int)(v)
}

// Custom: GetThreadTypeAddr gets `AVCodecContext.thread_type` address.
func (avctx *AvCodecContext) GetThreadTypeAddr() *int32 {
	return (*int32)(&avctx.thread_type)
}

const (
	FF_THREAD_FRAME = int32(C.FF_THREAD_FRAME)
	FF_THREAD_SLICE = int32(C.FF_THREAD_SLICE)
)

// Custom: Get ActiveThreadType gets `AVCodecContext.active_thread_type` value.
func (avctx *AvCodecContext) GetActiveThreadType() int32 {
	return (int32)(avctx.active_thread_type)
}

// Custom: SetActiveThreadType sets `AVCodecContext.active_thread_type` value.
func (avctx *AvCodecContext) SetActiveThreadType(v int32) {
	avctx.active_thread_type = (C.int)(v)
}

// Custom: GetActiveThreadTypeAddr gets `AVCodecContext.active_thread_type` address.
func (avctx *AvCodecContext) GetActiveThreadTypeAddr() *int32 {
	return (*int32)(&avctx.active_thread_type)
}

// Custom: Get ThreadSafeCallbacks gets `AVCodecContext.thread_safe_callbacks` value.
func (avctx *AvCodecContext) GetThreadSafeCallbacks() int32 {
	return (int32)(avctx.thread_safe_callbacks)
}

// Custom: SetThreadSafeCallbacks sets `AVCodecContext.thread_safe_callbacks` value.
func (avctx *AvCodecContext) SetThreadSafeCallbacks(v int32) {
	avctx.thread_safe_callbacks = (C.int)(v)
}

// Custom: GetThreadSafeCallbacksAddr gets `AVCodecContext.thread_safe_callbacks` address.
func (avctx *AvCodecContext) GetThreadSafeCallbacksAddr() *int32 {
	return (*int32)(&avctx.thread_safe_callbacks)
}

// Custom: Get NsseWeight gets `AVCodecContext.nsse_weight` value.
func (avctx *AvCodecContext) GetNsseWeight() int32 {
	return (int32)(avctx.nsse_weight)
}

// Custom: SetNsseWeight sets `AVCodecContext.nsse_weight` value.
func (avctx *AvCodecContext) SetNsseWeight(v int32) {
	avctx.nsse_weight = (C.int)(v)
}

// Custom: GetNsseWeightAddr gets `AVCodecContext.nsse_weight` address.
func (avctx *AvCodecContext) GetNsseWeightAddr() *int32 {
	return (*int32)(&avctx.nsse_weight)
}

// Custom: Get Profile gets `AVCodecContext.profile` value.
func (avctx *AvCodecContext) GetProfile() int32 {
	return (int32)(avctx.profile)
}

// Custom: SetProfile sets `AVCodecContext.profile` value.
func (avctx *AvCodecContext) SetProfile(v int32) {
	avctx.profile = (C.int)(v)
}

// Custom: GetProfileAddr gets `AVCodecContext.profile` address.
func (avctx *AvCodecContext) GetProfileAddr() *int32 {
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

// Custom: Get Level gets `AVCodecContext.level` value.
func (avctx *AvCodecContext) GetLevel() int32 {
	return (int32)(avctx.level)
}

// Custom: SetLevel sets `AVCodecContext.level` value.
func (avctx *AvCodecContext) SetLevel(v int32) {
	avctx.level = (C.int)(v)
}

// Custom: GetLevelAddr gets `AVCodecContext.level` address.
func (avctx *AvCodecContext) GetLevelAddr() *int32 {
	return (*int32)(&avctx.level)
}

// Custom: Get SkipLoopFilter gets `AVCodecContext.skip_loop_filter` value.
func (avctx *AvCodecContext) GetSkipLoopFilter() AvDiscard {
	return (AvDiscard)(avctx.skip_loop_filter)
}

// Custom: SetSkipLoopFilter sets `AVCodecContext.skip_loop_filter` value.
func (avctx *AvCodecContext) SetSkipLoopFilter(v AvDiscard) {
	avctx.skip_loop_filter = (C.enum_AVDiscard)(v)
}

// Custom: GetSkipLoopFilterAddr gets `AVCodecContext.skip_loop_filter` address.
func (avctx *AvCodecContext) GetSkipLoopFilterAddr() *AvDiscard {
	return (*AvDiscard)(&avctx.skip_loop_filter)
}

// Custom: Get SkipIdct gets `AVCodecContext.skip_idct` value.
func (avctx *AvCodecContext) GetSkipIdct() AvDiscard {
	return (AvDiscard)(avctx.skip_idct)
}

// Custom: SetSkipIdct sets `AVCodecContext.skip_idct` value.
func (avctx *AvCodecContext) SetSkipIdct(v AvDiscard) {
	avctx.skip_idct = (C.enum_AVDiscard)(v)
}

// Custom: GetSkipIdctAddr gets `AVCodecContext.skip_idct` address.
func (avctx *AvCodecContext) GetSkipIdctAddr() *AvDiscard {
	return (*AvDiscard)(&avctx.skip_idct)
}

// Custom: Get SkipFrame gets `AVCodecContext.skip_frame` value.
func (avctx *AvCodecContext) GetSkipFrame() AvDiscard {
	return (AvDiscard)(avctx.skip_frame)
}

// Custom: SetSkipFrame sets `AVCodecContext.skip_frame` value.
func (avctx *AvCodecContext) SetSkipFrame(v AvDiscard) {
	avctx.skip_frame = (C.enum_AVDiscard)(v)
}

// Custom: GetSkipFrameAddr gets `AVCodecContext.skip_frame` address.
func (avctx *AvCodecContext) GetSkipFrameAddr() *AvDiscard {
	return (*AvDiscard)(&avctx.skip_frame)
}

// Custom: Get SubtitleHeader gets `AVCodecContext.subtitle_header` value.
func (avctx *AvCodecContext) GetSubtitleHeader() *uint8 {
	return (*uint8)(avctx.subtitle_header)
}

// Custom: SetSubtitleHeader sets `AVCodecContext.subtitle_header` value.
func (avctx *AvCodecContext) SetSubtitleHeader(v *uint8) {
	avctx.subtitle_header = (*C.uint8_t)(v)
}

// Custom: GetSubtitleHeaderAddr gets `AVCodecContext.subtitle_header` address.
func (avctx *AvCodecContext) GetSubtitleHeaderAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&avctx.subtitle_header))
}

// Custom: Get SubtitleHeaderSize gets `AVCodecContext.subtitle_header_size` value.
func (avctx *AvCodecContext) GetSubtitleHeaderSize() int32 {
	return (int32)(avctx.subtitle_header_size)
}

// Custom: SetSubtitleHeaderSize sets `AVCodecContext.subtitle_header_size` value.
func (avctx *AvCodecContext) SetSubtitleHeaderSize(v int32) {
	avctx.subtitle_header_size = (C.int)(v)
}

// Custom: GetSubtitleHeaderSizeAddr gets `AVCodecContext.subtitle_header_size` address.
func (avctx *AvCodecContext) GetSubtitleHeaderSizeAddr() *int32 {
	return (*int32)(&avctx.subtitle_header_size)
}

// Custom: Get VbvDelay gets `AVCodecContext.vbv_delay` value.
func (avctx *AvCodecContext) GetVbvDelay() uint64 {
	return (uint64)(avctx.vbv_delay)
}

// Custom: SetVbvDelay sets `AVCodecContext.vbv_delay` value.
func (avctx *AvCodecContext) SetVbvDelay(v uint64) {
	avctx.vbv_delay = (C.uint64_t)(v)
}

// Custom: GetVbvDelayAddr gets `AVCodecContext.vbv_delay` address.
func (avctx *AvCodecContext) GetVbvDelayAddr() *uint64 {
	return (*uint64)(&avctx.vbv_delay)
}

// Custom: Get SideDataOnlyPackets gets `AVCodecContext.side_data_only_packets` value.
func (avctx *AvCodecContext) GetSideDataOnlyPackets() int32 {
	return (int32)(avctx.side_data_only_packets)
}

// Custom: SetSideDataOnlyPackets sets `AVCodecContext.side_data_only_packets` value.
func (avctx *AvCodecContext) SetSideDataOnlyPackets(v int32) {
	avctx.side_data_only_packets = (C.int)(v)
}

// Custom: GetSideDataOnlyPacketsAddr gets `AVCodecContext.side_data_only_packets` address.
func (avctx *AvCodecContext) GetSideDataOnlyPacketsAddr() *int32 {
	return (*int32)(&avctx.side_data_only_packets)
}

// Custom: Get InitialPadding gets `AVCodecContext.initial_padding` value.
func (avctx *AvCodecContext) GetInitialPadding() int32 {
	return (int32)(avctx.initial_padding)
}

// Custom: SetInitialPadding sets `AVCodecContext.initial_padding` value.
func (avctx *AvCodecContext) SetInitialPadding(v int32) {
	avctx.initial_padding = (C.int)(v)
}

// Custom: GetInitialPaddingAddr gets `AVCodecContext.initial_padding` address.
func (avctx *AvCodecContext) GetInitialPaddingAddr() *int32 {
	return (*int32)(&avctx.initial_padding)
}

// Custom: Get Framerate gets `AVCodecContext.framerate` value.
func (avctx *AvCodecContext) GetFramerate() AvRational {
	return (AvRational)(avctx.framerate)
}

// Custom: SetFramerate sets `AVCodecContext.framerate` value.
func (avctx *AvCodecContext) SetFramerate(v AvRational) {
	avctx.framerate = (C.AVRational)(v)
}

// Custom: GetFramerateAddr gets `AVCodecContext.framerate` address.
func (avctx *AvCodecContext) GetFramerateAddr() *AvRational {
	return (*AvRational)(&avctx.framerate)
}

// Custom: Get SwPixFmt gets `AVCodecContext.sw_pix_fmt` value.
func (avctx *AvCodecContext) GetSwPixFmt() AvPixelFormat {
	return (AvPixelFormat)(avctx.sw_pix_fmt)
}

// Custom: SetSwPixFmt sets `AVCodecContext.sw_pix_fmt` value.
func (avctx *AvCodecContext) SetSwPixFmt(v AvPixelFormat) {
	avctx.sw_pix_fmt = (C.enum_AVPixelFormat)(v)
}

// Custom: GetSwPixFmtAddr gets `AVCodecContext.sw_pix_fmt` address.
func (avctx *AvCodecContext) GetSwPixFmtAddr() *AvPixelFormat {
	return (*AvPixelFormat)(&avctx.sw_pix_fmt)
}

// Custom: Get PktTimebase gets `AVCodecContext.pkt_timebase` value.
func (avctx *AvCodecContext) GetPktTimebase() AvRational {
	return (AvRational)(avctx.pkt_timebase)
}

// Custom: SetPktTimebase sets `AVCodecContext.pkt_timebase` value.
func (avctx *AvCodecContext) SetPktTimebase(v AvRational) {
	avctx.pkt_timebase = (C.AVRational)(v)
}

// Custom: GetPktTimebaseAddr gets `AVCodecContext.pkt_timebase` address.
func (avctx *AvCodecContext) GetPktTimebaseAddr() *AvRational {
	return (*AvRational)(&avctx.pkt_timebase)
}

// Custom: Get CodecDescriptor gets `AVCodecContext.codec_descriptor` value.
func (avctx *AvCodecContext) GetCodecDescriptor() *AvCodecDescriptor {
	return (*AvCodecDescriptor)(avctx.codec_descriptor)
}

// Custom: SetCodecDescriptor sets `AVCodecContext.codec_descriptor` value.
func (avctx *AvCodecContext) SetCodecDescriptor(v *AvCodecDescriptor) {
	avctx.codec_descriptor = (*C.struct_AVCodecDescriptor)(v)
}

// Custom: GetCodecDescriptorAddr gets `AVCodecContext.codec_descriptor` address.
func (avctx *AvCodecContext) GetCodecDescriptorAddr() **AvCodecDescriptor {
	return (**AvCodecDescriptor)(unsafe.Pointer(&avctx.codec_descriptor))
}

// Custom: Get PtsCorrectionNumFaultyPts gets `AVCodecContext.pts_correction_num_faulty_pts` value.
func (avctx *AvCodecContext) GetPtsCorrectionNumFaultyPts() int64 {
	return (int64)(avctx.pts_correction_num_faulty_pts)
}

// Custom: SetPtsCorrectionNumFaultyPts sets `AVCodecContext.pts_correction_num_faulty_pts` value.
func (avctx *AvCodecContext) SetPtsCorrectionNumFaultyPts(v int64) {
	avctx.pts_correction_num_faulty_pts = (C.int64_t)(v)
}

// Custom: GetPtsCorrectionNumFaultyPtsAddr gets `AVCodecContext.pts_correction_num_faulty_pts` address.
func (avctx *AvCodecContext) GetPtsCorrectionNumFaultyPtsAddr() *int64 {
	return (*int64)(&avctx.pts_correction_num_faulty_pts)
}

// Custom: Get PtsCorrectionNumFaultyDts gets `AVCodecContext.pts_correction_num_faulty_dts` value.
func (avctx *AvCodecContext) GetPtsCorrectionNumFaultyDts() int64 {
	return (int64)(avctx.pts_correction_num_faulty_dts)
}

// Custom: SetPtsCorrectionNumFaultyDts sets `AVCodecContext.pts_correction_num_faulty_dts` value.
func (avctx *AvCodecContext) SetPtsCorrectionNumFaultyDts(v int64) {
	avctx.pts_correction_num_faulty_dts = (C.int64_t)(v)
}

// Custom: GetPtsCorrectionNumFaultyDtsAddr gets `AVCodecContext.pts_correction_num_faulty_dts` address.
func (avctx *AvCodecContext) GetPtsCorrectionNumFaultyDtsAddr() *int64 {
	return (*int64)(&avctx.pts_correction_num_faulty_dts)
}

// Custom: Get PtsCorrectionLastPts gets `AVCodecContext.pts_correction_last_pts` value.
func (avctx *AvCodecContext) GetPtsCorrectionLastPts() int64 {
	return (int64)(avctx.pts_correction_last_pts)
}

// Custom: SetPtsCorrectionLastPts sets `AVCodecContext.pts_correction_last_pts` value.
func (avctx *AvCodecContext) SetPtsCorrectionLastPts(v int64) {
	avctx.pts_correction_last_pts = (C.int64_t)(v)
}

// Custom: GetPtsCorrectionLastPtsAddr gets `AVCodecContext.pts_correction_last_pts` address.
func (avctx *AvCodecContext) GetPtsCorrectionLastPtsAddr() *int64 {
	return (*int64)(&avctx.pts_correction_last_pts)
}

// Custom: Get PtsCorrectionLastDts gets `AVCodecContext.pts_correction_last_dts` value.
func (avctx *AvCodecContext) GetPtsCorrectionLastDts() int64 {
	return (int64)(avctx.pts_correction_last_dts)
}

// Custom: SetPtsCorrectionLastDts sets `AVCodecContext.pts_correction_last_dts` value.
func (avctx *AvCodecContext) SetPtsCorrectionLastDts(v int64) {
	avctx.pts_correction_last_dts = (C.int64_t)(v)
}

// Custom: GetPtsCorrectionLastDtsAddr gets `AVCodecContext.pts_correction_last_dts` address.
func (avctx *AvCodecContext) GetPtsCorrectionLastDtsAddr() *int64 {
	return (*int64)(&avctx.pts_correction_last_dts)
}

// Custom: Get SubCharenc gets `AVCodecContext.sub_charenc` value.
func (avctx *AvCodecContext) GetSubCharenc() string {
	return C.GoString(avctx.sub_charenc)
}

// Custom: Get SubCharencMode gets `AVCodecContext.sub_charenc_mode` value.
func (avctx *AvCodecContext) GetSubCharencMode() int32 {
	return (int32)(avctx.sub_charenc_mode)
}

// Custom: SetSubCharencMode sets `AVCodecContext.sub_charenc_mode` value.
func (avctx *AvCodecContext) SetSubCharencMode(v int32) {
	avctx.sub_charenc_mode = (C.int)(v)
}

// Custom: GetSubCharencModeAddr gets `AVCodecContext.sub_charenc_mode` address.
func (avctx *AvCodecContext) GetSubCharencModeAddr() *int32 {
	return (*int32)(&avctx.sub_charenc_mode)
}

const (
	FF_SUB_CHARENC_MODE_DO_NOTHING  = int32(C.FF_SUB_CHARENC_MODE_DO_NOTHING)
	FF_SUB_CHARENC_MODE_AUTOMATIC   = int32(C.FF_SUB_CHARENC_MODE_AUTOMATIC)
	FF_SUB_CHARENC_MODE_PRE_DECODER = int32(C.FF_SUB_CHARENC_MODE_PRE_DECODER)
	FF_SUB_CHARENC_MODE_IGNORE      = int32(C.FF_SUB_CHARENC_MODE_IGNORE)
)

// Custom: Get SkipAlpha gets `AVCodecContext.skip_alpha` value.
func (avctx *AvCodecContext) GetSkipAlpha() int32 {
	return (int32)(avctx.skip_alpha)
}

// Custom: SetSkipAlpha sets `AVCodecContext.skip_alpha` value.
func (avctx *AvCodecContext) SetSkipAlpha(v int32) {
	avctx.skip_alpha = (C.int)(v)
}

// Custom: GetSkipAlphaAddr gets `AVCodecContext.skip_alpha` address.
func (avctx *AvCodecContext) GetSkipAlphaAddr() *int32 {
	return (*int32)(&avctx.skip_alpha)
}

// Custom: Get SeekPreroll gets `AVCodecContext.seek_preroll` value.
func (avctx *AvCodecContext) GetSeekPreroll() int32 {
	return (int32)(avctx.seek_preroll)
}

// Custom: SetSeekPreroll sets `AVCodecContext.seek_preroll` value.
func (avctx *AvCodecContext) SetSeekPreroll(v int32) {
	avctx.seek_preroll = (C.int)(v)
}

// Custom: GetSeekPrerollAddr gets `AVCodecContext.seek_preroll` address.
func (avctx *AvCodecContext) GetSeekPrerollAddr() *int32 {
	return (*int32)(&avctx.seek_preroll)
}

// Custom: Get DebugMv gets `AVCodecContext.debug_mv` value.
func (avctx *AvCodecContext) GetDebugMv() int32 {
	return (int32)(avctx.debug_mv)
}

// Custom: SetDebugMv sets `AVCodecContext.debug_mv` value.
func (avctx *AvCodecContext) SetDebugMv(v int32) {
	avctx.debug_mv = (C.int)(v)
}

// Custom: GetDebugMvAddr gets `AVCodecContext.debug_mv` address.
func (avctx *AvCodecContext) GetDebugMvAddr() *int32 {
	return (*int32)(&avctx.debug_mv)
}

const (
	FF_DEBUG_VIS_MV_P_FOR  = int32(C.FF_DEBUG_VIS_MV_P_FOR)
	FF_DEBUG_VIS_MV_B_FOR  = int32(C.FF_DEBUG_VIS_MV_B_FOR)
	FF_DEBUG_VIS_MV_B_BACK = int32(C.FF_DEBUG_VIS_MV_B_BACK)
)

// Custom: Get ChromaIntraMatrix gets `AVCodecContext.chroma_intra_matrix` value.
func (avctx *AvCodecContext) GetChromaIntraMatrix() *uint16 {
	return (*uint16)(avctx.chroma_intra_matrix)
}

// Custom: SetChromaIntraMatrix sets `AVCodecContext.chroma_intra_matrix` value.
func (avctx *AvCodecContext) SetChromaIntraMatrix(v *uint16) {
	avctx.chroma_intra_matrix = (*C.uint16_t)(v)
}

// Custom: GetChromaIntraMatrixAddr gets `AVCodecContext.chroma_intra_matrix` address.
func (avctx *AvCodecContext) GetChromaIntraMatrixAddr() **uint16 {
	return (**uint16)(unsafe.Pointer(&avctx.chroma_intra_matrix))
}

// Custom: Get DumpSeparator gets `AVCodecContext.dump_separator` value.
func (avctx *AvCodecContext) GetDumpSeparator() *uint8 {
	return (*uint8)(avctx.dump_separator)
}

// Custom: SetDumpSeparator sets `AVCodecContext.dump_separator` value.
func (avctx *AvCodecContext) SetDumpSeparator(v *uint8) {
	avctx.dump_separator = (*C.uint8_t)(v)
}

// Custom: GetDumpSeparatorAddr gets `AVCodecContext.dump_separator` address.
func (avctx *AvCodecContext) GetDumpSeparatorAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&avctx.dump_separator))
}

// Custom: Get CodecWhitelist gets `AVCodecContext.codec_whitelist` value.
func (avctx *AvCodecContext) GetCodecWhitelist() string {
	return C.GoString(avctx.codec_whitelist)
}

// Custom: Get Properties gets `AVCodecContext.properties` value.
func (avctx *AvCodecContext) GetProperties() uint32 {
	return (uint32)(avctx.properties)
}

// Custom: SetProperties sets `AVCodecContext.properties` value.
func (avctx *AvCodecContext) SetProperties(v uint32) {
	avctx.properties = (C.uint)(v)
}

// Custom: GetPropertiesAddr gets `AVCodecContext.properties` address.
func (avctx *AvCodecContext) GetPropertiesAddr() *uint32 {
	return (*uint32)(&avctx.properties)
}

// Custom: Get CodedSideData gets `AVCodecContext.coded_side_data` value.
func (avctx *AvCodecContext) GetCodedSideData() *AvPacketSideData {
	return (*AvPacketSideData)(avctx.coded_side_data)
}

// Custom: SetCodedSideData sets `AVCodecContext.coded_side_data` value.
func (avctx *AvCodecContext) SetCodedSideData(v *AvPacketSideData) {
	avctx.coded_side_data = (*C.AVPacketSideData)(v)
}

// Custom: GetCodedSideDataAddr gets `AVCodecContext.coded_side_data` address.
func (avctx *AvCodecContext) GetCodedSideDataAddr() **AvPacketSideData {
	return (**AvPacketSideData)(unsafe.Pointer(&avctx.coded_side_data))
}

// Custom: Get NbCodedSideData gets `AVCodecContext.nb_coded_side_data` value.
func (avctx *AvCodecContext) GetNbCodedSideData() int32 {
	return (int32)(avctx.nb_coded_side_data)
}

// Custom: SetNbCodedSideData sets `AVCodecContext.nb_coded_side_data` value.
func (avctx *AvCodecContext) SetNbCodedSideData(v int32) {
	avctx.nb_coded_side_data = (C.int)(v)
}

// Custom: GetNbCodedSideDataAddr gets `AVCodecContext.nb_coded_side_data` address.
func (avctx *AvCodecContext) GetNbCodedSideDataAddr() *int32 {
	return (*int32)(&avctx.nb_coded_side_data)
}

// Custom: Get HwFramesCtx gets `AVCodecContext.hw_frames_ctx` value.
func (avctx *AvCodecContext) GetHwFramesCtx() *AvBufferRef {
	return (*AvBufferRef)(avctx.hw_frames_ctx)
}

// Custom: SetHwFramesCtx sets `AVCodecContext.hw_frames_ctx` value.
func (avctx *AvCodecContext) SetHwFramesCtx(v *AvBufferRef) {
	avctx.hw_frames_ctx = (*C.AVBufferRef)(v)
}

// Custom: GetHwFramesCtxAddr gets `AVCodecContext.hw_frames_ctx` address.
func (avctx *AvCodecContext) GetHwFramesCtxAddr() **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer(&avctx.hw_frames_ctx))
}

// Custom: Get SubTextFormat gets `AVCodecContext.sub_text_format` value.
func (avctx *AvCodecContext) GetSubTextFormat() int32 {
	return (int32)(avctx.sub_text_format)
}

// Custom: SetSubTextFormat sets `AVCodecContext.sub_text_format` value.
func (avctx *AvCodecContext) SetSubTextFormat(v int32) {
	avctx.sub_text_format = (C.int)(v)
}

// Custom: GetSubTextFormatAddr gets `AVCodecContext.sub_text_format` address.
func (avctx *AvCodecContext) GetSubTextFormatAddr() *int32 {
	return (*int32)(&avctx.sub_text_format)
}

const (
	FF_SUB_TEXT_FMT_ASS              = int32(C.FF_SUB_TEXT_FMT_ASS)
	FF_SUB_TEXT_FMT_ASS_WITH_TIMINGS = int32(C.FF_SUB_TEXT_FMT_ASS_WITH_TIMINGS)
)

// Custom: Get TrailingPadding gets `AVCodecContext.trailing_padding` value.
func (avctx *AvCodecContext) GetTrailingPadding() int32 {
	return (int32)(avctx.trailing_padding)
}

// Custom: SetTrailingPadding sets `AVCodecContext.trailing_padding` value.
func (avctx *AvCodecContext) SetTrailingPadding(v int32) {
	avctx.trailing_padding = (C.int)(v)
}

// Custom: GetTrailingPaddingAddr gets `AVCodecContext.trailing_padding` address.
func (avctx *AvCodecContext) GetTrailingPaddingAddr() *int32 {
	return (*int32)(&avctx.trailing_padding)
}

// Custom: Get MaxPixels gets `AVCodecContext.max_pixels` value.
func (avctx *AvCodecContext) GetMaxPixels() int64 {
	return (int64)(avctx.max_pixels)
}

// Custom: SetMaxPixels sets `AVCodecContext.max_pixels` value.
func (avctx *AvCodecContext) SetMaxPixels(v int64) {
	avctx.max_pixels = (C.int64_t)(v)
}

// Custom: GetMaxPixelsAddr gets `AVCodecContext.max_pixels` address.
func (avctx *AvCodecContext) GetMaxPixelsAddr() *int64 {
	return (*int64)(&avctx.max_pixels)
}

// Custom: Get HwDeviceCtx gets `AVCodecContext.hw_device_ctx` value.
func (avctx *AvCodecContext) GetHwDeviceCtx() *AvBufferRef {
	return (*AvBufferRef)(avctx.hw_device_ctx)
}

// Custom: SetHwDeviceCtx sets `AVCodecContext.hw_device_ctx` value.
func (avctx *AvCodecContext) SetHwDeviceCtx(v *AvBufferRef) {
	avctx.hw_device_ctx = (*C.AVBufferRef)(v)
}

// Custom: GetHwDeviceCtxAddr gets `AVCodecContext.hw_device_ctx` address.
func (avctx *AvCodecContext) GetHwDeviceCtxAddr() **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer(&avctx.hw_device_ctx))
}

// Custom: Get HwaccelFlags gets `AVCodecContext.hwaccel_flags` value.
func (avctx *AvCodecContext) GetHwaccelFlags() int32 {
	return (int32)(avctx.hwaccel_flags)
}

// Custom: SetHwaccelFlags sets `AVCodecContext.hwaccel_flags` value.
func (avctx *AvCodecContext) SetHwaccelFlags(v int32) {
	avctx.hwaccel_flags = (C.int)(v)
}

// Custom: GetHwaccelFlagsAddr gets `AVCodecContext.hwaccel_flags` address.
func (avctx *AvCodecContext) GetHwaccelFlagsAddr() *int32 {
	return (*int32)(&avctx.hwaccel_flags)
}

// Custom: Get ApplyCropping gets `AVCodecContext.apply_cropping` value.
func (avctx *AvCodecContext) GetApplyCropping() int32 {
	return (int32)(avctx.apply_cropping)
}

// Custom: SetApplyCropping sets `AVCodecContext.apply_cropping` value.
func (avctx *AvCodecContext) SetApplyCropping(v int32) {
	avctx.apply_cropping = (C.int)(v)
}

// Custom: GetApplyCroppingAddr gets `AVCodecContext.apply_cropping` address.
func (avctx *AvCodecContext) GetApplyCroppingAddr() *int32 {
	return (*int32)(&avctx.apply_cropping)
}

// Custom: Get ExtraHwFrames gets `AVCodecContext.extra_hw_frames` value.
func (avctx *AvCodecContext) GetExtraHwFrames() int32 {
	return (int32)(avctx.extra_hw_frames)
}

// Custom: SetExtraHwFrames sets `AVCodecContext.extra_hw_frames` value.
func (avctx *AvCodecContext) SetExtraHwFrames(v int32) {
	avctx.extra_hw_frames = (C.int)(v)
}

// Custom: GetExtraHwFramesAddr gets `AVCodecContext.extra_hw_frames` address.
func (avctx *AvCodecContext) GetExtraHwFramesAddr() *int32 {
	return (*int32)(&avctx.extra_hw_frames)
}

// Custom: Get DiscardDamagedPercentage gets `AVCodecContext.discard_damaged_percentage` value.
func (avctx *AvCodecContext) GetDiscardDamagedPercentage() int32 {
	return (int32)(avctx.discard_damaged_percentage)
}

// Custom: SetDiscardDamagedPercentage sets `AVCodecContext.discard_damaged_percentage` value.
func (avctx *AvCodecContext) SetDiscardDamagedPercentage(v int32) {
	avctx.discard_damaged_percentage = (C.int)(v)
}

// Custom: GetDiscardDamagedPercentageAddr gets `AVCodecContext.discard_damaged_percentage` address.
func (avctx *AvCodecContext) GetDiscardDamagedPercentageAddr() *int32 {
	return (*int32)(&avctx.discard_damaged_percentage)
}

// Custom: Get MaxSamples gets `AVCodecContext.max_samples` value.
func (avctx *AvCodecContext) GetMaxSamples() int64 {
	return (int64)(avctx.max_samples)
}

// Custom: SetMaxSamples sets `AVCodecContext.max_samples` value.
func (avctx *AvCodecContext) SetMaxSamples(v int64) {
	avctx.max_samples = (C.int64_t)(v)
}

// Custom: GetMaxSamplesAddr gets `AVCodecContext.max_samples` address.
func (avctx *AvCodecContext) GetMaxSamplesAddr() *int64 {
	return (*int64)(&avctx.max_samples)
}

// Custom: Get ExportSideData gets `AVCodecContext.export_side_data` value.
func (avctx *AvCodecContext) GetExportSideData() int32 {
	return (int32)(avctx.export_side_data)
}

// Custom: SetExportSideData sets `AVCodecContext.export_side_data` value.
func (avctx *AvCodecContext) SetExportSideData(v int32) {
	avctx.export_side_data = (C.int)(v)
}

// Custom: GetExportSideDataAddr gets `AVCodecContext.export_side_data` address.
func (avctx *AvCodecContext) GetExportSideDataAddr() *int32 {
	return (*int32)(&avctx.export_side_data)
}

// Deprecated: No use
func AvCodecGetPktTimebase(avctx *AvCodecContext) AvRational {
	return AvRational(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use
func AvCodecSetPktTimebase(avctx *AvCodecContext, r AvRational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(avctx), (C.struct_AVRational)(r))
}

// Deprecated: No use
func AvCodecGetCodecDescriptor(avctx *AvCodecContext) *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use
func AvCodecSetCodecDescriptor(avctx *AvCodecContext, d *AvCodecDescriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(avctx), (*C.struct_AVCodecDescriptor)(d))
}

// Deprecated: No use
func AvCodecGetLowres(avctx *AvCodecContext) int32 {
	return (int32)(C.av_codec_get_lowres((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use
func AvCodecSetLowres(avctx *AvCodecContext, i int32) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(avctx), C.int(i))
}

// Deprecated: No use
func AvCodecGetSeekPreroll(avctx *AvCodecContext) int32 {
	return (int32)(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use
func AvCodecSetSeekPreroll(avctx *AvCodecContext, i int32) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(avctx), C.int(i))
}

// Deprecated: No use
func AvCodecGetChromaIntraMatrix(avctx *AvCodecContext) *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(avctx)))
}

// Deprecated: No use
func AvCodecSetChromaIntraMatrix(avctx *AvCodecContext, t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(avctx), (*C.uint16_t)(t))
}

// Deprecated: No use
func AvCodecGetMaxLowres(c *AvCodec) int32 {
	return (int32)(C.av_codec_get_max_lowres((*C.struct_AVCodec)(c)))
}

// MpegEncContext
type MpegEncContext C.struct_MpegEncContext

// AvHWAccel
type AvHWAccel C.struct_AVHWAccel

const (
	AV_HWACCEL_CODEC_CAP_EXPERIMENTAL      = int(C.AV_HWACCEL_CODEC_CAP_EXPERIMENTAL)
	AV_HWACCEL_FLAG_IGNORE_LEVEL           = int(C.AV_HWACCEL_FLAG_IGNORE_LEVEL)
	AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH       = int(C.AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH)
	AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH = int(C.AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH)
)

// AvPicture
type AvPicture C.struct_AVPicture

// AvSubtitleType
type AvSubtitleType = C.enum_AVSubtitleType

const (
	SUBTITLE_NONE   = AvSubtitleType(C.SUBTITLE_NONE)
	SUBTITLE_BITMAP = AvSubtitleType(C.SUBTITLE_BITMAP)
	SUBTITLE_TEXT   = AvSubtitleType(C.SUBTITLE_TEXT)
	SUBTITLE_ASS    = AvSubtitleType(C.SUBTITLE_ASS)
)

const AV_SUBTITLE_FLAG_FORCED = C.AV_SUBTITLE_FLAG_FORCED

// AvSubtitleRect
type AvSubtitleRect C.struct_AVSubtitleRect

// AvSubtitle
type AvSubtitle C.struct_AVSubtitle

// If c is NULL, returns the first registered codec,
// if c is non-NULL, returns the next registered codec after c,
// or NULL if c is the last one.
func AvCodecNext(c *AvCodec) *AvCodec {
	return (*AvCodec)(C.av_codec_next((*C.struct_AVCodec)(c)))
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
func AvCodecRegister(c *AvCodec) {
	C.avcodec_register((*C.struct_AVCodec)(c))
}

// Deprecated: Calling this function is unnecessary.
func AvCodecRegisterAll() {
	C.avcodec_register_all()
}

// AvCodecAllocContext3 allocates an AVCodecContext and set its fields to default values.
// The resulting struct should be freed with AvCodecFreeContext().
func AvCodecAllocContext3(c *AvCodec) *AvCodecContext {
	return (*AvCodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
}

// AvCodecFreeContext frees the codec context and everything associated with it
// and write NULL to the provided pointer.
func AvCodecFreeContext(avctx **AvCodecContext) {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(avctx)))
}

// Deprecated: No use
func AvCodecGetContextDefaults3(avctx *AvCodecContext, c *AvCodec) int32 {
	return (int32)(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVCodec)(c)))
}

// AvCodecGetClass gets the AvClass for AvCodecContext.
func AvCodecGetClass() *AvClass {
	return (*AvClass)(C.avcodec_get_class())
}

// Deprecated: No use
func AvCodecGetFrameClass() *AvClass {
	return (*AvClass)(C.avcodec_get_frame_class())
}

// AvCodecGetSubtitleRectClass gets the AvClass for AvSubtitleRect.
func AvCodecGetSubtitleRectClass() *AvClass {
	return (*AvClass)(C.avcodec_get_subtitle_rect_class())
}

// Deprecated: Use an intermediate AVCodecParameters instance and the
// AvCodecParametersFromContext() / AvCodecParametersToContext() functions.
func AvCodecCopyContext(dest, src *AvCodecContext) int32 {
	return (int32)(C.avcodec_copy_context((*C.struct_AVCodecContext)(dest),
		(*C.struct_AVCodecContext)(src)))
}

// AvCodecParametersFromContext fills the parameters struct based on the values from the supplied codec
// context. Any allocated fields in par are freed and replaced with duplicates
// of the corresponding fields in codec.
func AvCodecParametersFromContext(par *AvCodecParameters, avctx *AvCodecContext) int32 {
	return (int32)(C.avcodec_parameters_from_context((*C.struct_AVCodecParameters)(par),
		(*C.struct_AVCodecContext)(avctx)))
}

// AvCodecParametersToContext fills the codec context based on the values from the supplied codec
// parameters. Any allocated fields in codec that have a corresponding field in
// par are freed and replaced with duplicates of the corresponding field in par.
// Fields in codec that do not have a counterpart in par are not touched.
func AvCodecParametersToContext(avctx *AvCodecContext, par *AvCodecParameters) int32 {
	return (int32)(C.avcodec_parameters_to_context((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVCodecParameters)(par)))
}

// AvCodecContext initializes the context to use the given codec.
func AvCodecOpen2(avctx *AvCodecContext, c *AvCodec, d **AvDictionary) int32 {
	return (int32)(C.avcodec_open2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

// AvCodecClose closes a given context and free all the data associated with it (but not the context itself).
func AvCodecClose(avctx *AvCodecContext) int32 {
	return (int32)(C.avcodec_close((*C.struct_AVCodecContext)(avctx)))
}

// AvSubtitleFree frees all allocated data in the given subtitle struct.
func AvSubtitleFree(s *AvSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

// The default callback for AVCodecContext.get_buffer2().
func AvCodecDefaultGetBuffer2(avctx *AvCodecContext, frame *AvFrame, flags int32) int32 {
	return (int32)(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(frame), C.int(flags)))
}

// The default callback for AVCodecContext.get_encode_buffer().
func AvCodecDefaultGetEncodeBuffer(avctx *AvCodecContext, pkt *AvPacket, flags int32) int32 {
	return (int32)(C.avcodec_default_get_encode_buffer((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(pkt), C.int(flags)))
}

// AvCodecAlignDimensions modifies width and height values so that they will result in a memory
// buffer that is acceptable for the codec if you do not use any horizontal padding.
func AvCodecAlignDimensions(avctx *AvCodecContext, width, height *int32) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(avctx), (*C.int)(width), (*C.int)(height))
}

// AvCodecAlignDimensions2 modifies width and height values so that they will result in a memory
// buffer that is acceptable for the codec if you also ensure that all
// line sizes are a multiple of the respective linesize_align[i].
func AvCodecAlignDimensions2(avctx *AvCodecContext, width, height *int32,
	linesizeAlign [AV_NUM_DATA_POINTERS]int32) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(avctx),
		(*C.int)(width), (*C.int)(height), (*C.int)(unsafe.Pointer(&linesizeAlign[0])))
}

// AvCodecEnumToChromaPos converts AvChromaLocation to swscale x/y chroma position.
func AvCodecEnumToChromaPos(xpos, ypos *int32, pos AvChromaLocation) int32 {
	return (int32)(C.avcodec_enum_to_chroma_pos((*C.int)(xpos), (*C.int)(ypos), (C.enum_AVChromaLocation)(pos)))
}

// AvCodecChromaPosToEnum converts swscale x/y chroma position to AvChromaLocation.
func AvCodecChromaPosToEnum(xpos, ypos int32) AvChromaLocation {
	return (AvChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(xpos), C.int(ypos)))
}

// Deprecated: Use AvCodecSendPacket() and AvCodecReceiveFrame().
func AvCodecDecodeAudio4(avctx *AvCodecContext, frame *AvFrame, gotFramePtr *int32, avpkt *AvPacket) int32 {
	return (int32)(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(frame), (*C.int)(gotFramePtr), (*C.struct_AVPacket)(avpkt)))
}

// Deprecated: Use AvCodecSendPacket() and AvCodecReceiveFrame().
func AvCodecDecodeVideo2(avctx *AvCodecContext, picture *AvFrame, gotPicturePtr *int32, avpkt *AvPacket) int32 {
	return (int32)(C.avcodec_decode_video2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(picture), (*C.int)(gotPicturePtr), (*C.struct_AVPacket)(avpkt)))
}

// AvCodecDecodeSubtitle2 decodes a subtitle message.
func AvCodecDecodeSubtitle2(avctx *AvCodecContext, sub *AvSubtitle, gotSubPtr *int32, avpkt *AvPacket) int32 {
	return (int32)(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVSubtitle)(sub), (*C.int)(gotSubPtr), (*C.struct_AVPacket)(avpkt)))
}

// AvCodecSendPacket supplies raw packet data as input to a decoder.
func AvCodecSendPacket(avctx *AvCodecContext, avpkt *AvPacket) int32 {
	return (int32)(C.avcodec_send_packet((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(avpkt)))
}

// AvCodecReceiveFrame returns decoded output data from a decoder.
func AvCodecReceiveFrame(avctx *AvCodecContext, frame *AvFrame) int32 {
	return (int32)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(frame)))
}

// AvCodecSendFrame supplies a raw video or audio frame to the encoder. Use AvCodecReceivePacket()
// to retrieve buffered output packets.
func AvCodecSendFrame(avctx *AvCodecContext, frame *AvFrame) int32 {
	return (int32)(C.avcodec_send_frame((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVFrame)(frame)))
}

// AvCodecReceivePacket reads encoded data from the encoder.
func AvCodecReceivePacket(avctx *AvCodecContext, avpkt *AvPacket) int32 {
	return (int32)(C.avcodec_receive_packet((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(avpkt)))
}

// AvCodecGetHwFramesParameters create and return a AVHWFramesContext with values adequate for hardware
// decoding.
func AvCodecGetHwFramesParameters(avctx *AvCodecContext, deviceRef *AvBufferRef,
	hwPixFmt AvPixelFormat,
	outFramesRef **AvBufferRef) int32 {
	return (int32)(C.avcodec_get_hw_frames_parameters((*C.struct_AVCodecContext)(avctx),
		(*C.AVBufferRef)(deviceRef),
		(C.enum_AVPixelFormat)(hwPixFmt),
		(**C.AVBufferRef)(unsafe.Pointer(outFramesRef))))
}

// AvPictureStructure
type AvPictureStructure = C.enum_AVPictureStructure

const (
	AV_PICTURE_STRUCTURE_UNKNOWN      = AvPictureStructure(C.AV_PICTURE_STRUCTURE_UNKNOWN)
	AV_PICTURE_STRUCTURE_TOP_FIELD    = AvPictureStructure(C.AV_PICTURE_STRUCTURE_TOP_FIELD)
	AV_PICTURE_STRUCTURE_BOTTOM_FIELD = AvPictureStructure(C.AV_PICTURE_STRUCTURE_BOTTOM_FIELD)
	AV_PICTURE_STRUCTURE_FRAME        = AvPictureStructure(C.AV_PICTURE_STRUCTURE_FRAME)
)

// AvCodecParserContext
type AvCodecParserContext C.struct_AVCodecParserContext

// AvCodecParser
type AvCodecParser C.struct_AVCodecParser

// AvParserIterate iterates over all registered codec parsers.
func AvParserIterate(p *unsafe.Pointer) *AvCodecParser {
	return (*AvCodecParser)(C.av_parser_iterate(p))
}

// Deprecated: No use
func AvParserNext(c *AvCodecParser) *AvCodecParser {
	return (*AvCodecParser)(C.av_parser_next((*C.struct_AVCodecParser)(c)))
}

// Deprecated: No use
func AvRegisterCodecParser(parser *AvCodecParser) {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(parser))
}

// AvParserInit
func AvParserInit(codecID AvCodecID) *AvCodecParserContext {
	return (*AvCodecParserContext)(C.av_parser_init((C.int)(codecID)))
}

// Parse a packet.
func AvParserParse2(s *AvCodecParserContext, avctx *AvCodecContext,
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
func AvParserChange(s *AvCodecParserContext, avctx *AvCodecContext,
	outbuf **uint8, poutbufSize *int32,
	buf *uint8, bufSize int32, keyframe int32) int32 {
	return (int32)(C.av_parser_change((*C.AVCodecParserContext)(s),
		(*C.struct_AVCodecContext)(avctx),
		(**C.uint8_t)(unsafe.Pointer(outbuf)), (*C.int)(poutbufSize),
		(*C.uint8_t)(buf), (C.int)(bufSize), (C.int)(keyframe)))
}

// AvParserClose
func AvParserClose(s *AvCodecParserContext) {
	C.av_parser_close((*C.AVCodecParserContext)(s))
}

// Deprecated: Use AvCodecSendFrame()/AvCodecReceivePacket() instead.
func AvCodecEncodeAudio2(avctx *AvCodecContext,
	avpkt *AvPacket, frame *AvFrame, gotPacketPtr *int32) int32 {
	return (int32)(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(avpkt), (*C.struct_AVFrame)(frame), (*C.int)(gotPacketPtr)))
}

// Deprecated: Use AvCodecSendFrame()/AvCodecReceivePacket() instead.
func AvCodecEncodeVideo2(avctx *AvCodecContext,
	avpkt *AvPacket, frame *AvFrame, gotPacketPtr *int32) int32 {
	return (int32)(C.avcodec_encode_video2((*C.struct_AVCodecContext)(avctx),
		(*C.struct_AVPacket)(avpkt), (*C.struct_AVFrame)(frame), (*C.int)(gotPacketPtr)))
}

// AvCodecEncodeSubtitle
func AvCodecEncodeSubtitle(avctx *AvCodecContext,
	buf *uint8, bufSize int32, sub *AvSubtitle) int32 {
	return (int32)(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(avctx),
		(*C.uint8_t)(buf), (C.int)(bufSize), (*C.struct_AVSubtitle)(sub)))
}

// Deprecated: No use
func AvPictureAlloc(picture *AvPicture, pixFmt AvPixelFormat, width, height int32) int32 {
	return (int32)(C.avpicture_alloc((*C.struct_AVPicture)(picture),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height)))
}

// Deprecated: No use
func AvPictureFree(picture *AvPicture) {
	C.avpicture_free((*C.struct_AVPicture)(picture))
}

// Deprecated: Use AvImageFillArrays() instead.
func AvPictureFill(picture *AvPicture, ptr *uint8, pixFmt AvPixelFormat, width, height int32) int32 {
	return (int32)(C.avpicture_fill((*C.struct_AVPicture)(picture),
		(*C.uint8_t)(ptr), (C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height)))
}

// Deprecated: Use AvImageCopyToBuffer() instead.
func AvPictureLayout(src *AvPicture, pixFmt AvPixelFormat, width, height int32, dest *uint8, destSize int32) int32 {
	return (int32)(C.avpicture_layout((*C.struct_AVPicture)(src),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height),
		(*C.uchar)(dest), (C.int)(destSize)))
}

// Deprecated: Use AvImageGetBufferSize() instead.
func AvPictureGetSize(pixFmt AvPixelFormat, width, height int32) int32 {
	return (int32)(C.avpicture_get_size((C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height)))
}

// Deprecated: Use AvImageCopy() instead.
func AvPictureCopy(dst, src *AvPicture, pixFmt AvPixelFormat, width, height int32) {
	C.av_picture_copy((*C.struct_AVPicture)(dst), (*C.struct_AVPicture)(src),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height))
}

// Deprecated: No use
func AvPictureCrop(dst, src *AvPicture, pixFmt AvPixelFormat, topBand, leftBand int32) int32 {
	return (int32)(C.av_picture_crop((*C.struct_AVPicture)(dst), (*C.struct_AVPicture)(src),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(topBand), (C.int)(leftBand)))
}

// Deprecated: No use
func AvPicturePad(dst, src *AvPicture, width, height int32, pixFmt AvPixelFormat,
	padTop, padBottom, padLeft, padRight int32, color *int32) int32 {
	return (int32)(C.av_picture_pad((*C.struct_AVPicture)(dst), (*C.struct_AVPicture)(src),
		(C.int)(width), (C.int)(height), (C.enum_AVPixelFormat)(pixFmt),
		(C.int)(padTop), (C.int)(padBottom), (C.int)(padLeft), (C.int)(padRight),
		(*C.int)(color)))
}

// Deprecated: Use AvPixFmtGetChromaSubSample() instead.
func AvCodecGetChromaSubSample(pixFmt AvPixelFormat, hShift, vShift *int32) {
	C.avcodec_get_chroma_sub_sample((C.enum_AVPixelFormat)(pixFmt),
		(*C.int)(hShift), (*C.int)(vShift))
}

// AvCodecPixFmtToCodecTag returns a value representing the fourCC code associated to the
// pixel format pix_fmt, or 0 if no associated fourCC code can be found.
func AvCodecPixFmtToCodecTag(pixFmt AvPixelFormat) uint {
	return (uint)(C.avcodec_pix_fmt_to_codec_tag((C.enum_AVPixelFormat)(pixFmt)))
}

// AvCodecFindBestPixFmtOfList finds the best pixel format
// to convert to given a certain source pixel format.
func AvCodecFindBestPixFmtOfList(pixFmtList *AvPixelFormat,
	srcPixFmt AvPixelFormat, hasAlpha int32, lossPtr *int32) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_find_best_pix_fmt_of_list(
		(*C.enum_AVPixelFormat)(pixFmtList),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha),
		(*C.int)(lossPtr)))
}

// Deprecated: Use AvGetPixFmtLoss() instead.
func AvCodecGetPixFmtLoss(dstPixFmt, srcPixFmt AvPixelFormat, hasAlpha int32) int32 {
	return (int32)(C.avcodec_get_pix_fmt_loss((C.enum_AVPixelFormat)(dstPixFmt),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha)))
}

// Deprecated: Use AvFindBestPixFmtOf2() instead.
func AvCodecFindBestPixFmtOf2(dstPixFmt1, dstPixFmt2, srcPixFmt AvPixelFormat,
	hasAlpha int32, lossPtr *int32) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_find_best_pix_fmt_of_2(
		(C.enum_AVPixelFormat)(dstPixFmt1), (C.enum_AVPixelFormat)(dstPixFmt2),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha), (*C.int)(lossPtr)))
}

// Deprecated: No use
func AvCodecFindBestPixFmt2(dstPixFmt1, dstPixFmt2, srcPixFmt AvPixelFormat,
	hasAlpha int32, lossPtr *int32) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_find_best_pix_fmt2(
		(C.enum_AVPixelFormat)(dstPixFmt1), (C.enum_AVPixelFormat)(dstPixFmt2),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha), (*C.int)(lossPtr)))
}

// AvCodecDefaultGetFormat
func AvCodecDefaultGetFormat(avctx *AvCodecContext, fmt *AvPixelFormat) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(avctx),
		(*C.enum_AVPixelFormat)(fmt)))
}

// Deprecated: Use  AvFourccMakeString() or AvFourcc2str() instead.
func AvGetCodecTagString(buf *int8, bufSize uint, codecTag uint32) int32 {
	return (int32)(C.av_get_codec_tag_string((*C.char)(buf),
		(C.size_t)(bufSize), (C.uint)(codecTag)))
}

// AvCodecString
func AvCodecString(buf *int8, bufSize int32, enc *AvCodecContext, encode int32) {
	C.avcodec_string((*C.char)(buf), (C.int)(bufSize),
		(*C.struct_AVCodecContext)(enc), (C.int)(encode))
}

// AvGetProfileName returns a name for the specified profile, if available.
func AvGetProfileName(c *AvCodec, profile int32) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), (C.int)(profile)))
}

// AvCodecProfileName returns a name for the specified profile, if available.
func AvCodecProfileName(codecID AvCodecID, profile int32) string {
	return C.GoString(C.avcodec_profile_name((C.enum_AVCodecID)(codecID), (C.int)(profile)))
}

// typedef int (*avcodec_excute_func)(AvCodecContext *c2, void *arg2)
type AvcodecExcuteFunc C.avcodec_excute_func

// AvCodecDefaultExecute
func AvCodecDefaultExecute(avctx *AvCodecContext, f AvcodecExcuteFunc, arg unsafe.Pointer,
	ret *int32, count, size int32) int32 {
	return (int32)(C.avcodec_default_execute((*C.struct_AVCodecContext)(avctx),
		(C.avcodec_excute_func)(f), arg, (*C.int)(ret), (C.int)(count), (C.int)(size)))
}

// AvcodecDefaultExecute2
func AvcodecDefaultExecute2(avctx *AvCodecContext, f AvcodecExcuteFunc, arg unsafe.Pointer,
	ret *int32, count int32) int32 {
	return (int32)(C.avcodec_default_execute2((*C.struct_AVCodecContext)(avctx),
		(C.avcodec_excute_func)(f), arg, (*C.int)(ret), (C.int)(count)))
}

// AvCodecFillAudioFrame fills AVFrame audio data and linesize pointers.
func AvCodecFillAudioFrame(frame *AvFrame, nbChannels int32,
	sampleFmt AvSampleFormat, buf *uint8,
	bufSize int32, align int32) int32 {
	return (int32)(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(frame),
		(C.int)(nbChannels), (C.enum_AVSampleFormat)(sampleFmt),
		(*C.uint8_t)(buf), (C.int)(bufSize), (C.int)(align)))
}

// AvCodecFlushBuffers resets the internal codec state / flush internal buffers. Should be called
// e.g. when seeking or when switching to a different stream.
func AvCodecFlushBuffers(avctx *AvCodecContext) {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(avctx))
}

// AvGetBitsPerSample returns codec bits per sample.
func AvGetBitsPerSample(codecID AvCodecID) int32 {
	return (int32)(C.av_get_bits_per_sample((C.enum_AVCodecID)(codecID)))
}

// AvGetPcmCodec returns the PCM codec associated with a sample format.
func AvGetPcmCodec(fmt AvSampleFormat, be int32) AvCodecID {
	return (AvCodecID)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(fmt), (C.int)(be)))
}

// AvGetExactBitsPerSample returns codec bits per sample.
func AvGetExactBitsPerSample(codecID AvCodecID) int32 {
	return (int32)(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(codecID)))
}

// AvGetAudioFrameDuration returns audio frame duration.
func AvGetAudioFrameDuration(avctx *AvCodecContext, frameBytes int32) int32 {
	return (int32)(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(avctx), (C.int)(frameBytes)))
}

// AvGetAudioFrameDuration2 returns audio frame duration.
func AvGetAudioFrameDuration2(par *AvCodecParameters, frameBytes int32) int32 {
	return (int32)(C.av_get_audio_frame_duration2((*C.struct_AVCodecParameters)(par), (C.int)(frameBytes)))
}

type AvBitStreamFilterContext C.struct_AVBitStreamFilterContext

// Deprecated: Use AvBSFContext instead.
func AvRegisterBitstreamFilter(bsf *AvBitStreamFilter) {
	C.av_register_bitstream_filter((*C.struct_AVBitStreamFilter)(bsf))
}

// Deprecated: Use AvBSFContext instead.
func AvBitstreamFilterInit(name string) *AvBitStreamFilterContext {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AvBitStreamFilterContext)(C.av_bitstream_filter_init((*C.char)(namePtr)))
}

// Deprecated: Use AvBSFContext instead.
func AvBitstreamFilterFilter(bsfc *AvBitStreamFilterContext, avctx *AvCodecContext, args string,
	outbuf **uint8, poutbufSize *int32,
	buf *uint8, bufSize int32, keyframe int32) int32 {
	argsPtr, nameFunc := StringCasting(args)
	defer nameFunc()
	return (int32)(C.av_bitstream_filter_filter((*C.struct_AVBitStreamFilterContext)(bsfc),
		(*C.struct_AVCodecContext)(avctx), (*C.char)(argsPtr),
		(**C.uint8_t)(unsafe.Pointer(outbuf)), (*C.int)(poutbufSize),
		(*C.uint8_t)(buf), (C.int)(bufSize), (C.int)(keyframe)))
}

// Deprecated: Use AvBSFContext instead.
func AvBitstreamFilterClose(bsfc *AvBitStreamFilterContext) {
	C.av_bitstream_filter_close((*C.struct_AVBitStreamFilterContext)(bsfc))
}

// Deprecated: Use AvBSFContext instead.
func AvBitstreamFilterNext(f *AvBitStreamFilter) *AvBitStreamFilter {
	return (*AvBitStreamFilter)(C.av_bitstream_filter_next((*C.struct_AVBitStreamFilter)(f)))
}

// Deprecated: No use
func AvBsfNext(opaque *unsafe.Pointer) *AvBitStreamFilter {
	return (*AvBitStreamFilter)(C.av_bsf_next(opaque))
}

// AvFastPaddedMalloc
func AvFastPaddedMalloc(ptr unsafe.Pointer, size *uint32, minSize uint) {
	C.av_fast_padded_malloc(ptr, (*C.uint)(size), (C.size_t)(minSize))
}

// AvFastPaddedMallocz
func AvFastPaddedMallocz(ptr unsafe.Pointer, size *uint32, minSize uint) {
	C.av_fast_padded_mallocz(ptr, (*C.uint)(size), (C.size_t)(minSize))
}

// AvXiphlacing encodes extradata length to a buffer. Used by xiph codecs.
func AvXiphlacing(s *uint8, v int32) int32 {
	return (int32)(C.av_xiphlacing((*C.uchar)(s), (C.uint)(v)))
}

// Deprecated: No use
func AvRegisterHwaccel(hwaccel *AvHWAccel) {
	C.av_register_hwaccel((*C.struct_AVHWAccel)(hwaccel))
}

// Deprecated: No use
func AvHwaccelNext(hwaccel *AvHWAccel) *AvHWAccel {
	return (*AvHWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(hwaccel)))
}

// AvLockOp
type AvLockOp = C.enum_AVLockOp

const (
	AV_LOCK_CREATE  = AvLockOp(C.AV_LOCK_CREATE)
	AV_LOCK_OBTAIN  = AvLockOp(C.AV_LOCK_OBTAIN)
	AV_LOCK_RELEASE = AvLockOp(C.AV_LOCK_RELEASE)
	AV_LOCK_DESTROY = AvLockOp(C.AV_LOCK_DESTROY)
)

// typedef int (*av_lockmgr_cb)(void **mutex, enum AvLockOp op)
type AvLockmgrCb C.av_lockmgr_cb

// AvLockmgrRegister
func AvLockmgrRegister(cb AvLockmgrCb) int32 {
	return (int32)(C.av_lockmgr_register((C.av_lockmgr_cb)(cb)))
}

// A positive value if s is open,
// 0 otherwise.
func AvCodecIsOpen(avctx *AvCodecContext) int32 {
	return (int32)(C.avcodec_is_open((*C.struct_AVCodecContext)(avctx)))
}

// AvCpbPropertiesAlloc allocates a CPB properties structure and initialize its fields to default
// values.
func AvCpbPropertiesAlloc(size *uint) *AvCPBProperties {
	return (*AvCPBProperties)(C.av_cpb_properties_alloc((*C.size_t)(unsafe.Pointer(size))))
}
