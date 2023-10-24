package ffmpeg

/*
#include <libavcodec/codec_par.h>
*/
import "C"
import "unsafe"

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

// AVCodecParameters
type AVCodecParameters C.struct_AVCodecParameters

// Custom: GetCodecType gets `AVCodecParameters.codectype` value.
func (par *AVCodecParameters) GetCodecType() AVMediaType {
	return (AVMediaType)(par.codec_type)
}

// Custom: SetCodecType sets `AVCodecParameters.codectype` value.
func (par *AVCodecParameters) SetCodecType(v AVMediaType) {
	par.codec_type = (C.enum_AVMediaType)(v)
}

// Custom: GetCodecTypeAddr gets `AVCodecParameters.codectype` address.
func (par *AVCodecParameters) GetCodecTypeAddr() *AVMediaType {
	return (*AVMediaType)(unsafe.Pointer(&par.codec_type))
}

// Custom: GetCodecId gets `AVCodecParameters.codec_id` value.
func (par *AVCodecParameters) GetCodecId() AVCodecID {
	return (AVCodecID)(par.codec_id)
}

// Custom: SetCodecId sets `AVCodecParameters.codec_id` value.
func (par *AVCodecParameters) SetCodecId(v AVCodecID) {
	par.codec_id = (C.enum_AVCodecID)(v)
}

// Custom: GetCodecIdAddr gets `AVCodecParameters.codec_id` address.
func (par *AVCodecParameters) GetCodecIdAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&par.codec_id))
}

// Custom: GetCodecTag gets `AVCodecParameters.codec_tag` value.
func (par *AVCodecParameters) GetCodecTag() uint32 {
	return (uint32)(par.codec_tag)
}

// Custom: SetCodecTag sets `AVCodecParameters.codec_tag` value.
func (par *AVCodecParameters) SetCodecTag(v uint32) {
	par.codec_tag = (C.uint)(v)
}

// Custom: GetCodecTagAddr gets `AVCodecParameters.codec_tag` address.
func (par *AVCodecParameters) GetCodecTagAddr() *uint32 {
	return (*uint32)(&par.codec_tag)
}

// Custom: GetExtradata gets `AVCodecParameters.extradata` value.
func (par *AVCodecParameters) GetExtradata() *uint8 {
	return (*uint8)(par.extradata)
}

// Custom: SetExtradata sets `AVCodecParameters.extradata` value.
func (par *AVCodecParameters) SetExtradata(v *uint8) {
	par.extradata = (*C.uint8_t)(v)
}

// Custom: GetExtradataAddr gets `AVCodecParameters.extradata` address.
func (par *AVCodecParameters) GetExtradataAddr() *uint8 {
	return (*uint8)(unsafe.Pointer(&par.extradata))
}

// Custom: GetExtradataSize gets `AVCodecParameters.extradata_size` value.
func (par *AVCodecParameters) GetExtradataSize() int32 {
	return (int32)(par.extradata_size)
}

// Custom: SetExtradataSize sets `AVCodecParameters.extradata_size` value.
func (par *AVCodecParameters) SetExtradataSize(v int32) {
	par.extradata_size = (C.int)(v)
}

// Custom: GetExtradataSizeAddr gets `AVCodecParameters.extradata_size` address.
func (par *AVCodecParameters) GetExtradataSizeAddr() *int32 {
	return (*int32)(&par.extradata_size)
}

// Custom: GetFormat gets `AVCodecParameters.format` value.
func (par *AVCodecParameters) GetFormat() int32 {
	return (int32)(par.format)
}

// Custom: SetFormat sets `AVCodecParameters.format` value.
func (par *AVCodecParameters) SetFormat(v int32) {
	par.format = (C.int)(v)
}

// Custom: GetFormatAddr gets `AVCodecParameters.format` address.
func (par *AVCodecParameters) GetFormatAddr() *int32 {
	return (*int32)(&par.format)
}

// Custom: GetBitRate gets `AVCodecParameters.bit_rate` value.
func (par *AVCodecParameters) GetBitRate() int64 {
	return (int64)(par.bit_rate)
}

// Custom: SetBitRate sets `AVCodecParameters.bit_rate` value.
func (par *AVCodecParameters) SetBitRate(v int64) {
	par.bit_rate = (C.int64_t)(v)
}

// Custom: GetBitRateAddr gets `AVCodecParameters.bit_rate` address.
func (par *AVCodecParameters) GetBitRateAddr() *int64 {
	return (*int64)(&par.bit_rate)
}

// Custom: GetBitsPerCodedSample gets `AVCodecParameters.bits_per_coded_sample` value.
func (par *AVCodecParameters) GetBitsPerCodedSample() int32 {
	return (int32)(par.bits_per_coded_sample)
}

// Custom: SetBitsPerCodedSample sets `AVCodecParameters.bits_per_coded_sample` value.
func (par *AVCodecParameters) SetBitsPerCodedSample(v int32) {
	par.bits_per_coded_sample = (C.int)(v)
}

// Custom: GetBitsPerCodedSampleAddr gets `AVCodecParameters.bits_per_coded_sample` address.
func (par *AVCodecParameters) GetBitsPerCodedSampleAddr() *int32 {
	return (*int32)(&par.bits_per_coded_sample)
}

// Custom: GetBitsPerRawSample gets `AVCodecParameters.bits_per_raw_sample` value.
func (par *AVCodecParameters) GetBitsPerRawSample() int32 {
	return (int32)(par.bits_per_raw_sample)
}

// Custom: SetBitsPerRawSample sets `AVCodecParameters.bits_per_raw_sample` value.
func (par *AVCodecParameters) SetBitsPerRawSample(v int32) {
	par.bits_per_raw_sample = (C.int)(v)
}

// Custom: GetBitsPerRawSampleAddr gets `AVCodecParameters.bits_per_raw_sample` address.
func (par *AVCodecParameters) GetBitsPerRawSampleAddr() *int32 {
	return (*int32)(&par.bits_per_raw_sample)
}

// Custom: GetProfile gets `AVCodecParameters.profile` value.
func (par *AVCodecParameters) GetProfile() int32 {
	return (int32)(par.profile)
}

// Custom: SetProfile sets `AVCodecParameters.profile` value.
func (par *AVCodecParameters) SetProfile(v int32) {
	par.profile = (C.int)(v)
}

// Custom: GetProfileAddr gets `AVCodecParameters.profile` address.
func (par *AVCodecParameters) GetProfileAddr() *int32 {
	return (*int32)(&par.profile)
}

// Custom: GetLevel gets `AVCodecParameters.level` value.
func (par *AVCodecParameters) GetLevel() int32 {
	return (int32)(par.level)
}

// Custom: SetLevel sets `AVCodecParameters.level` value.
func (par *AVCodecParameters) SetLevel(v int32) {
	par.level = (C.int)(v)
}

// Custom: GetLevelAddr gets `AVCodecParameters.level` address.
func (par *AVCodecParameters) GetLevelAddr() *int32 {
	return (*int32)(&par.level)
}

// Custom: GetWidth gets `AVCodecParameters.width` value.
func (par *AVCodecParameters) GetWidth() int32 {
	return (int32)(par.width)
}

// Custom: SetWidth sets `AVCodecParameters.width` value.
func (par *AVCodecParameters) SetWidth(v int32) {
	par.width = (C.int)(v)
}

// Custom: GetWidthAddr gets `AVCodecParameters.width` address.
func (par *AVCodecParameters) GetWidthAddr() *int32 {
	return (*int32)(&par.width)
}

// Custom: GetHeight gets `AVCodecParameters.height` value.
func (par *AVCodecParameters) GetHeight() int32 {
	return (int32)(par.height)
}

// Custom: SetHeight sets `AVCodecParameters.height` value.
func (par *AVCodecParameters) SetHeight(v int32) {
	par.height = (C.int)(v)
}

// Custom: GetHeightAddr gets `AVCodecParameters.height` address.
func (par *AVCodecParameters) GetHeightAddr() *int32 {
	return (*int32)(&par.height)
}

// Custom: GetSampleAspectRatio gets `AVCodecParameters.sample_aspect_ratio` value.
func (par *AVCodecParameters) GetSampleAspectRatio() AVRational {
	return (AVRational)(par.sample_aspect_ratio)
}

// Custom: SetSampleAspectRatio sets `AVCodecParameters.sample_aspect_ratio` value.
func (par *AVCodecParameters) SetSampleAspectRatio(v AVRational) {
	par.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// Custom: GetSampleAspectRatioAddr gets `AVCodecParameters.sample_aspect_ratio` address.
func (par *AVCodecParameters) GetSampleAspectRatioAddr() *AVRational {
	return (*AVRational)(&par.sample_aspect_ratio)
}

// Custom: GetFieldOrder gets `AVCodecParameters.field_order` value.
func (par *AVCodecParameters) GetFieldOrder() AVFieldOrder {
	return (AVFieldOrder)(par.field_order)
}

// Custom: SetFieldOrder sets `AVCodecParameters.field_order` value.
func (par *AVCodecParameters) SetFieldOrder(v AVFieldOrder) {
	par.field_order = (C.enum_AVFieldOrder)(v)
}

// Custom: GetFieldOrderAddr gets `AVCodecParameters.field_order` address.
func (par *AVCodecParameters) GetFieldOrderAddr() *AVFieldOrder {
	return (*AVFieldOrder)(unsafe.Pointer(&par.field_order))
}

// Custom: GetColorRange gets `AVCodecParameters.colorrange` value.
func (par *AVCodecParameters) GetColorRange() AVColorRange {
	return (AVColorRange)(par.color_range)
}

// Custom: SetColorRange sets `AVCodecParameters.colorrange` value.
func (par *AVCodecParameters) SetColorRange(v AVColorRange) {
	par.color_range = (C.enum_AVColorRange)(v)
}

// Custom: GetColorRangeAddr gets `AVCodecParameters.colorrange` address.
func (par *AVCodecParameters) GetColorRangeAddr() *AVColorRange {
	return (*AVColorRange)(unsafe.Pointer(&par.color_range))
}

// Custom: GetColorPrimaries gets `AVCodecParameters.color_primaries` value.
func (par *AVCodecParameters) GetColorPrimaries() AVColorPrimaries {
	return (AVColorPrimaries)(par.color_primaries)
}

// Custom: SetColorPrimaries sets `AVCodecParameters.color_primaries` value.
func (par *AVCodecParameters) SetColorPrimaries(v AVColorPrimaries) {
	par.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// Custom: GetColorPrimariesAddr gets `AVCodecParameters.color_primaries` address.
func (par *AVCodecParameters) GetColorPrimariesAddr() *AVColorPrimaries {
	return (*AVColorPrimaries)(unsafe.Pointer(&par.color_primaries))
}

// Custom: GetColorTrc gets `AVCodecParameters.color_trc` value.
func (par *AVCodecParameters) GetColorTrc() AVColorTransferCharacteristic {
	return (AVColorTransferCharacteristic)(par.color_trc)
}

// Custom: SetColorTrc sets `AVCodecParameters.color_trc` value.
func (par *AVCodecParameters) SetColorTrc(v AVColorTransferCharacteristic) {
	par.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// Custom: GetColorTrcAddr gets `AVCodecParameters.color_trc` address.
func (par *AVCodecParameters) GetColorTrcAddr() *AVColorTransferCharacteristic {
	return (*AVColorTransferCharacteristic)(unsafe.Pointer(&par.color_trc))
}

// Custom: GetColorSpace gets `AVCodecParameters.color_space` value.
func (par *AVCodecParameters) GetColorSpace() AVColorSpace {
	return (AVColorSpace)(par.color_space)
}

// Custom: SetColorSpace sets `AVCodecParameters.color_space` value.
func (par *AVCodecParameters) SetColorSpace(v AVColorSpace) {
	par.color_space = (C.enum_AVColorSpace)(v)
}

// Custom: GetColorSpaceAddr gets `AVCodecParameters.color_space` address.
func (par *AVCodecParameters) GetColorSpaceAddr() *AVColorSpace {
	return (*AVColorSpace)(unsafe.Pointer(&par.color_space))
}

// Custom: GetChromaLocation gets `AVCodecParameters.chroma_location` value.
func (par *AVCodecParameters) GetChromaLocation() AVChromaLocation {
	return (AVChromaLocation)(par.chroma_location)
}

// Custom: SetChromaLocation sets `AVCodecParameters.chroma_location` value.
func (par *AVCodecParameters) SetChromaLocation(v AVChromaLocation) {
	par.chroma_location = (C.enum_AVChromaLocation)(v)
}

// Custom: GetChromaLocationAddr gets `AVCodecParameters.chroma_location` address.
func (par *AVCodecParameters) GetChromaLocationAddr() *AVChromaLocation {
	return (*AVChromaLocation)(unsafe.Pointer(&par.chroma_location))
}

// Custom: GetVideoDelay gets `AVCodecParameters.video_delay` value.
func (par *AVCodecParameters) GetVideoDelay() int32 {
	return (int32)(par.video_delay)
}

// Custom: SetVideoDelay sets `AVCodecParameters.video_delay` value.
func (par *AVCodecParameters) SetVideoDelay(v int32) {
	par.video_delay = (C.int)(v)
}

// Custom: GetVideoDelayAddr gets `AVCodecParameters.video_delay` address.
func (par *AVCodecParameters) GetVideoDelayAddr() *int32 {
	return (*int32)(&par.video_delay)
}

// Custom: GetChannelLayout gets `AVCodecParameters.channel_layout` value.
func (par *AVCodecParameters) GetChannelLayout() uint64 {
	return (uint64)(par.channel_layout)
}

// Custom: SetChannelLayout sets `AVCodecParameters.channel_layout` value.
func (par *AVCodecParameters) SetChannelLayout(v uint64) {
	par.channel_layout = (C.uint64_t)(v)
}

// Custom: GetChannelLayoutAddr gets `AVCodecParameters.channel_layout` address.
func (par *AVCodecParameters) GetChannelLayoutAddr() *uint64 {
	return (*uint64)(&par.channel_layout)
}

// Custom: GetChannels gets `AVCodecParameters.channels` value.
func (par *AVCodecParameters) GetChannels() int32 {
	return (int32)(par.channels)
}

// Custom: SetChannels sets `AVCodecParameters.channels` value.
func (par *AVCodecParameters) SetChannels(v int32) {
	par.channels = (C.int)(v)
}

// Custom: GetChannelsAddr gets `AVCodecParameters.channels` address.
func (par *AVCodecParameters) GetChannelsAddr() *int32 {
	return (*int32)(&par.channels)
}

// Custom: GetSampleRate gets `AVCodecParameters.sample_rate` value.
func (par *AVCodecParameters) GetSampleRate() int32 {
	return (int32)(par.sample_rate)
}

// Custom: SetSampleRate sets `AVCodecParameters.sample_rate` value.
func (par *AVCodecParameters) SetSampleRate(v int32) {
	par.sample_rate = (C.int)(v)
}

// Custom: GetSampleRateAddr gets `AVCodecParameters.sample_rate` address.
func (par *AVCodecParameters) GetSampleRateAddr() *int32 {
	return (*int32)(&par.sample_rate)
}

// Custom: GetBlockAlign gets `AVCodecParameters.block_align` value.
func (par *AVCodecParameters) GetBlockAlign() int32 {
	return (int32)(par.block_align)
}

// Custom: SetBlockAlign sets `AVCodecParameters.block_align` value.
func (par *AVCodecParameters) SetBlockAlign(v int32) {
	par.block_align = (C.int)(v)
}

// Custom: GetBlockAlignAddr gets `AVCodecParameters.block_align` address.
func (par *AVCodecParameters) GetBlockAlignAddr() *int32 {
	return (*int32)(&par.block_align)
}

// Custom: GetFrameSize gets `AVCodecParameters.frame_size` value.
func (par *AVCodecParameters) GetFrameSize() int32 {
	return (int32)(par.frame_size)
}

// Custom: SetFrameSize sets `AVCodecParameters.frame_size` value.
func (par *AVCodecParameters) SetFrameSize(v int32) {
	par.frame_size = (C.int)(v)
}

// Custom: GetFrameSizeAddr gets `AVCodecParameters.frame_size` address.
func (par *AVCodecParameters) GetFrameSizeAddr() *int32 {
	return (*int32)(&par.frame_size)
}

// Custom: GetInitialPadding gets `AVCodecParameters.initial_padding` value.
func (par *AVCodecParameters) GetInitialPadding() int32 {
	return (int32)(par.initial_padding)
}

// Custom: SetInitialPadding sets `AVCodecParameters.initial_padding` value.
func (par *AVCodecParameters) SetInitialPadding(v int32) {
	par.initial_padding = (C.int)(v)
}

// Custom: GetInitialPaddingAddr gets `AVCodecParameters.initial_padding` address.
func (par *AVCodecParameters) GetInitialPaddingAddr() *int32 {
	return (*int32)(&par.initial_padding)
}

// Custom: GetTrailingPadding gets `AVCodecParameters.trailing_padding` value.
func (par *AVCodecParameters) GetTrailingPadding() int32 {
	return (int32)(par.trailing_padding)
}

// Custom: SetTrailingPadding sets `AVCodecParameters.trailing_padding` value.
func (par *AVCodecParameters) SetTrailingPadding(v int32) {
	par.trailing_padding = (C.int)(v)
}

// Custom: GetTrailingPaddingAddr gets `AVCodecParameters.trailing_padding` address.
func (par *AVCodecParameters) GetTrailingPaddingAddr() *int32 {
	return (*int32)(&par.trailing_padding)
}

// Custom: GetSeekPreroll gets `AVCodecParameters.seek_preroll` value.
func (par *AVCodecParameters) GetSeekPreroll() int32 {
	return (int32)(par.seek_preroll)
}

// Custom: SetSeekPreroll sets `AVCodecParameters.seek_preroll` value.
func (par *AVCodecParameters) SetSeekPreroll(v int32) {
	par.seek_preroll = (C.int)(v)
}

// Custom: GetSeekPrerollAddr gets `AVCodecParameters.seek_preroll` address.
func (par *AVCodecParameters) GetSeekPrerollAddr() *int32 {
	return (*int32)(&par.seek_preroll)
}

// AvCodecParametersAlloc allocates a new AVCodecParameters and set its fields to default values
// (unknown/invalid/0). The returned struct must be freed with AVCodecParametersFree().
func AvCodecParametersAlloc() *AVCodecParameters {
	return (*AVCodecParameters)(C.avcodec_parameters_alloc())
}

// AvCodecParametersFree frees an AVCodecParameters instance and everything associated with it and
// write NULL to the supplied pointer.
func AvCodecParametersFree(par **AVCodecParameters) {
	C.avcodec_parameters_free((**C.struct_AVCodecParameters)(unsafe.Pointer(par)))
}

// AvCodecParametersCopy copies the contents of src to dst.
func AvCodecParametersCopy(dst, src *AVCodecParameters) int32 {
	return (int32)(C.avcodec_parameters_copy((*C.struct_AVCodecParameters)(dst),
		(*C.struct_AVCodecParameters)(src)))
}
