package ffmpeg

/*
#include <libavcodec/codec.h>
*/
import "C"
import "unsafe"

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

// AvProfile
type AvProfile C.struct_AVProfile

// Custom: GetProfile gets `AVProfile.profile` value.
func (p *AvProfile) GetProfile() int32 {
	return (int32)(p.profile)
}

// Custom: GetName gets `AVProfile.name` value.
func (p *AvProfile) GetName() string {
	return C.GoString(p.name)
}

// AvCodec
type AvCodec C.struct_AVCodec

// Custom: GetName gets `AVCodec.name` value.
func (codec *AvCodec) GetName() string {
	return C.GoString(codec.name)
}

// Custom: GetLongName gets `AVCodec.long_name` value.
func (codec *AvCodec) GetLongName() string {
	return C.GoString(codec.long_name)
}

// Custom: GetType gets `AVCodec.type` value.
func (codec *AvCodec) GetType() AvMediaType {
	return (AvMediaType)(codec._type)
}

// Custom: GetType gets `AVCodec.id` value.
func (codec *AvCodec) GetID() AvCodecID {
	return (AvCodecID)(codec.id)
}

// Custom: GetCapabilities gets `AVCodec.capabilities` value.
func (codec *AvCodec) GetCapabilities() int32 {
	return (int32)(codec.capabilities)
}

// Custom: GetSupportedFramerates gets `AVCodec.supportedFramerates` value.
func (codec *AvCodec) GetSupportedFramerates() (v []AvRational) {
	if codec.supported_framerates == nil {
		return v
	}
	zeroQ := AvMakeQ(0, 0)
	ptr := (*AvRational)(codec.supported_framerates)
	for AvCmpQ(zeroQ, *ptr) != 0 {
		v = append(v, *ptr)
		ptr = (*AvRational)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: GetPixFmts gets `AVCodec.pix_fmts` value.
func (codec *AvCodec) GetPixFmts() (v []AvPixelFormat) {
	if codec.pix_fmts == nil {
		return v
	}
	ptr := (*AvPixelFormat)(codec.pix_fmts)
	for *ptr != AV_PIX_FMT_NONE {
		v = append(v, *ptr)
		ptr = (*AvPixelFormat)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: GetSupportedSamplerates gets `AVCodec.supported_samplerates` value.
func (codec *AvCodec) GetSupportedSamplerates() (v []int32) {
	if codec.supported_samplerates == nil {
		return v
	}
	ptr := (*int32)(codec.supported_samplerates)
	for *ptr != 0 {
		v = append(v, *ptr)
		ptr = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: GetSampleFmts gets `AVCodec.sample_fmts` value.
func (codec *AvCodec) GetSampleFmts() (v []AvSampleFormat) {
	if codec.sample_fmts == nil {
		return v
	}
	ptr := (*AvSampleFormat)(codec.sample_fmts)
	for *ptr != AV_SAMPLE_FMT_NONE {
		v = append(v, *ptr)
		ptr = (*AvSampleFormat)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: GetChannelLayouts gets `AVCodec.channel_layouts` value.
func (codec *AvCodec) GetChannelLayouts() (v []uint64) {
	if codec.channel_layouts == nil {
		return v
	}
	ptr := (*uint64)(codec.channel_layouts)
	for *ptr != 0 {
		v = append(v, *ptr)
		ptr = (*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: GetMaxLowres gets `AVCodec.max_lowres` value.
func (codec *AvCodec) GetMaxLowres() uint8 {
	return (uint8)(codec.max_lowres)
}

// Custom: GetProfiles gets `AVCodec.profiles` value.
func (codec *AvCodec) GetProfiles() (v []AvProfile) {
	if codec.profiles == nil {
		return v
	}
	ptr := (*AvProfile)(codec.profiles)
	for ptr.GetProfile() != FF_PROFILE_UNKNOWN {
		v = append(v, *ptr)
		ptr = (*AvProfile)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: GetWrapperName gets `AVCodec.wrapper_name` value.
func (codec *AvCodec) GetWrapperName() string {
	return C.GoString(codec.wrapper_name)
}

// AvCodecIterate iterates over all registered codecs.
func AvCodecIterate(opaque *unsafe.Pointer) *AvCodec {
	return (*AvCodec)(C.av_codec_iterate(opaque))
}

// AvCodecFindDecoder finds a registered decoder with a matching codec ID.
func AvCodecFindDecoder(id AvCodecID) *AvCodec {
	return (*AvCodec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// AvCodecFindDecoderByName finds a registered decoder with the specified name.
func AvCodecFindDecoderByName(name string) *AvCodec {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AvCodec)(C.avcodec_find_decoder_by_name((*C.char)(namePtr)))
}

// AvCodecFindEncoder finds a registered encoder with a matching codec ID.
func AvCodecFindEncoder(id AvCodecID) *AvCodec {
	return (*AvCodec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

// AvCodecFindEncoderByName finds a registered encoder with the specified name.
func AvCodecFindEncoderByName(name string) *AvCodec {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AvCodec)(C.avcodec_find_encoder_by_name((*C.char)(namePtr)))
}

// AvCodecIsEncoder returns a non-zero number if codec is an encoder, zero otherwise
func AvCodecIsEncoder(codec *AvCodec) int32 {
	return (int32)(C.av_codec_is_encoder((*C.struct_AVCodec)(codec)))
}

// AvCodecIsDecoder returns a non-zero number if codec is an decoder, zero otherwise
func AvCodecIsDecoder(codec *AvCodec) int32 {
	return (int32)(C.av_codec_is_decoder((*C.struct_AVCodec)(codec)))
}

const (
	AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX = int32(C.AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX)
	AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX = int32(C.AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX)
	AV_CODEC_HW_CONFIG_METHOD_INTERNAL      = int32(C.AV_CODEC_HW_CONFIG_METHOD_INTERNAL)
	AV_CODEC_HW_CONFIG_METHOD_AD_HOC        = int32(C.AV_CODEC_HW_CONFIG_METHOD_AD_HOC)
)

// AvCodecHWConfig
type AvCodecHWConfig C.struct_AVCodecHWConfig

// AvCodecGetHwConfig retrieves supported hardware configurations for a codec.
func AvCodecGetHwConfig(codec *AvCodec, index int32) *AvCodecHWConfig {
	return (*AvCodecHWConfig)(C.avcodec_get_hw_config((*C.struct_AVCodec)(codec), (C.int)(index)))
}
