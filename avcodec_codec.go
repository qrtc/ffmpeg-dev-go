package ffmpeg

/*
#include <libavcodec/codec.h>
*/
import "C"

const (
	AV_CODEC_CAP_DRAW_HORIZ_BAND          = C.AV_CODEC_CAP_DRAW_HORIZ_BAND
	AV_CODEC_CAP_DR1                      = C.AV_CODEC_CAP_DR1
	AV_CODEC_CAP_TRUNCATED                = C.AV_CODEC_CAP_TRUNCATED
	AV_CODEC_CAP_DELAY                    = C.AV_CODEC_CAP_DELAY
	AV_CODEC_CAP_SMALL_LAST_FRAME         = C.AV_CODEC_CAP_SMALL_LAST_FRAME
	AV_CODEC_CAP_SUBFRAMES                = C.AV_CODEC_CAP_SUBFRAMES
	AV_CODEC_CAP_EXPERIMENTAL             = C.AV_CODEC_CAP_EXPERIMENTAL
	AV_CODEC_CAP_CHANNEL_CONF             = C.AV_CODEC_CAP_CHANNEL_CONF
	AV_CODEC_CAP_FRAME_THREADS            = C.AV_CODEC_CAP_FRAME_THREADS
	AV_CODEC_CAP_SLICE_THREADS            = C.AV_CODEC_CAP_SLICE_THREADS
	AV_CODEC_CAP_PARAM_CHANGE             = C.AV_CODEC_CAP_PARAM_CHANGE
	AV_CODEC_CAP_OTHER_THREADS            = C.AV_CODEC_CAP_OTHER_THREADS
	AV_CODEC_CAP_AUTO_THREADS             = C.AV_CODEC_CAP_AUTO_THREADS
	AV_CODEC_CAP_VARIABLE_FRAME_SIZE      = C.AV_CODEC_CAP_VARIABLE_FRAME_SIZE
	AV_CODEC_CAP_AVOID_PROBING            = C.AV_CODEC_CAP_AVOID_PROBING
	AV_CODEC_CAP_INTRA_ONLY               = C.AV_CODEC_CAP_INTRA_ONLY
	AV_CODEC_CAP_LOSSLESS                 = C.AV_CODEC_CAP_LOSSLESS
	AV_CODEC_CAP_HARDWARE                 = C.AV_CODEC_CAP_HARDWARE
	AV_CODEC_CAP_HYBRID                   = C.AV_CODEC_CAP_HYBRID
	AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE = C.AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE
	AV_CODEC_CAP_ENCODER_FLUSH            = C.AV_CODEC_CAP_ENCODER_FLUSH
)

// AVProfile
type AVProfile C.struct_AVProfile

// Custom: GetProfile gets `AVProfile.profile` value.
func (p *AVProfile) GetProfile() int32 {
	return (int32)(p.profile)
}

// Custom: GetName gets `AVProfile.name` value.
func (p *AVProfile) GetName() string {
	return C.GoString(p.name)
}

// AVCodec
type AVCodec C.struct_AVCodec

// Custom: GetName gets `AVCodec.name` value.
func (codec *AVCodec) GetName() string {
	return C.GoString(codec.name)
}

// Custom: GetLongName gets `AVCodec.long_name` value.
func (codec *AVCodec) GetLongName() string {
	return C.GoString(codec.long_name)
}

// Custom: GetType gets `AVCodec.type` value.
func (codec *AVCodec) GetType() AVMediaType {
	return (AVMediaType)(codec._type)
}

// Custom: GetType gets `AVCodec.id` value.
func (codec *AVCodec) GetID() AVCodecID {
	return (AVCodecID)(codec.id)
}

// Custom: GetCapabilities gets `AVCodec.capabilities` value.
func (codec *AVCodec) GetCapabilities() int32 {
	return (int32)(codec.capabilities)
}

// Custom: GetSupportedFramerates gets `AVCodec.supportedFramerates` value.
func (codec *AVCodec) GetSupportedFramerates() []AVRational {
	return TruncSlice((*AVRational)(codec.supported_framerates), func(ar AVRational) bool {
		return ar.GetNum() == 0 && ar.GetDen() == 0
	})
}

// Custom: GetPixFmts gets `AVCodec.pix_fmts` value.
func (codec *AVCodec) GetPixFmts() []AVPixelFormat {
	return TruncSlice((*AVPixelFormat)(codec.pix_fmts), func(pf AVPixelFormat) bool {
		return pf == AV_PIX_FMT_NONE
	})
}

// Custom: GetSupportedSamplerates gets `AVCodec.supported_samplerates` value.
func (codec *AVCodec) GetSupportedSamplerates() []int32 {
	return TruncSlice((*int32)(codec.supported_samplerates), func(i int32) bool {
		return i == 0
	})
}

// Custom: GetSampleFmts gets `AVCodec.sample_fmts` value.
func (codec *AVCodec) GetSampleFmts() []AVSampleFormat {
	return TruncSlice((*AVSampleFormat)(codec.sample_fmts), func(sf AVSampleFormat) bool {
		return sf == AV_SAMPLE_FMT_NONE
	})
}

// Custom: GetChannelLayouts gets `AVCodec.channel_layouts` value.
func (codec *AVCodec) GetChannelLayouts() []uint64 {
	return TruncSlice((*uint64)(codec.channel_layouts), func(u uint64) bool {
		return u == 0
	})
}

// Custom: GetMaxLowres gets `AVCodec.max_lowres` value.
func (codec *AVCodec) GetMaxLowres() uint8 {
	return (uint8)(codec.max_lowres)
}

// Custom: GetProfiles gets `AVCodec.profiles` value.
func (codec *AVCodec) GetProfiles() []AVProfile {
	return TruncSlice((*AVProfile)(codec.profiles), func(ap AVProfile) bool {
		return ap.GetProfile() == FF_PROFILE_UNKNOWN
	})
}

// Custom: GetWrapperName gets `AVCodec.wrapper_name` value.
func (codec *AVCodec) GetWrapperName() string {
	return C.GoString(codec.wrapper_name)
}

// AvCodecIterate iterates over all registered codecs.
func AvCodecIterate(opaque CVoidPointerPointer) *AVCodec {
	return (*AVCodec)(C.av_codec_iterate(VoidPointerPointer(opaque)))
}

// AvCodecFindDecoder finds a registered decoder with a matching codec ID.
func AvCodecFindDecoder(id AVCodecID) *AVCodec {
	return (*AVCodec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// AvCodecFindDecoderByName finds a registered decoder with the specified name.
func AvCodecFindDecoderByName(name string) *AVCodec {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AVCodec)(C.avcodec_find_decoder_by_name((*C.char)(namePtr)))
}

// AvCodecFindEncoder finds a registered encoder with a matching codec ID.
func AvCodecFindEncoder(id AVCodecID) *AVCodec {
	return (*AVCodec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

// AvCodecFindEncoderByName finds a registered encoder with the specified name.
func AvCodecFindEncoderByName(name string) *AVCodec {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AVCodec)(C.avcodec_find_encoder_by_name((*C.char)(namePtr)))
}

// AvCodecIsEncoder returns a non-zero number if codec is an encoder, zero otherwise
func AvCodecIsEncoder(codec *AVCodec) int32 {
	return (int32)(C.av_codec_is_encoder((*C.struct_AVCodec)(codec)))
}

// AvCodecIsDecoder returns a non-zero number if codec is an decoder, zero otherwise
func AvCodecIsDecoder(codec *AVCodec) int32 {
	return (int32)(C.av_codec_is_decoder((*C.struct_AVCodec)(codec)))
}

const (
	AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX = int32(C.AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX)
	AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX = int32(C.AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX)
	AV_CODEC_HW_CONFIG_METHOD_INTERNAL      = int32(C.AV_CODEC_HW_CONFIG_METHOD_INTERNAL)
	AV_CODEC_HW_CONFIG_METHOD_AD_HOC        = int32(C.AV_CODEC_HW_CONFIG_METHOD_AD_HOC)
)

// AVCodecHWConfig
type AVCodecHWConfig C.struct_AVCodecHWConfig

// Custom: GetPixFmt gets `AVCodecHWConfig.pix_fmt` value.
func (hwc *AVCodecHWConfig) GetPixFmt() AVPixelFormat {
	return (AVPixelFormat)(hwc.pix_fmt)
}

// Custom: SetPixFmt sets `AVCodecHWConfig.pix_fmt` value.
func (hwc *AVCodecHWConfig) SetPixFmt(v AVPixelFormat) {
	hwc.pix_fmt = (C.enum_AVPixelFormat)(v)
}

// Custom: GetPixFmtAddr gets `AVCodecHWConfig.pix_fmt` address.
func (hwc *AVCodecHWConfig) GetPixFmtAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&hwc.pix_fmt)
}

// Custom: GetMethods gets `AVCodecHWConfig.methods` value.
func (hwc *AVCodecHWConfig) GetMethods() int32 {
	return (int32)(hwc.methods)
}

// Custom: SetMethods sets `AVCodecHWConfig.methods` value.
func (hwc *AVCodecHWConfig) SetMethods(v int32) {
	hwc.methods = (C.int)(v)
}

// Custom: GetMethodsAddr gets `AVCodecHWConfig.methods` address.
func (hwc *AVCodecHWConfig) GetMethodsAddr() *int32 {
	return (*int32)(&hwc.methods)
}

// Custom: GetDeviceType gets `AVCodecHWConfig.device_type` value.
func (hwc *AVCodecHWConfig) GetDeviceType() AVHWDeviceType {
	return (AVHWDeviceType)(hwc.device_type)
}

// Custom: SetDeviceType sets `AVCodecHWConfig.device_type` value.
func (hwc *AVCodecHWConfig) SetDeviceType(v AVHWDeviceType) {
	hwc.device_type = (C.enum_AVHWDeviceType)(v)
}

// Custom: GetDeviceTypeAddr gets `AVCodecHWConfig.device_type` address.
func (hwc *AVCodecHWConfig) GetDeviceTypeAddr() *AVHWDeviceType {
	return (*AVHWDeviceType)(&hwc.device_type)
}

// AvCodecGetHwConfig retrieves supported hardware configurations for a codec.
func AvCodecGetHwConfig[T HelperInteger](codec *AVCodec, index T) *AVCodecHWConfig {
	return (*AVCodecHWConfig)(C.avcodec_get_hw_config((*C.struct_AVCodec)(codec), (C.int)(index)))
}
