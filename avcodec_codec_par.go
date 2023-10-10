package ffmpeg

/*
#include <libavcodec/codec_par.h>
*/
import "C"
import "unsafe"

type AvFieldOrder int32

const (
	AV_FIELD_UNKNOWN     = AvFieldOrder(C.AV_FIELD_UNKNOWN)
	AV_FIELD_PROGRESSIVE = AvFieldOrder(C.AV_FIELD_PROGRESSIVE)
	AV_FIELD_TT          = AvFieldOrder(C.AV_FIELD_TT)
	AV_FIELD_BB          = AvFieldOrder(C.AV_FIELD_BB)
	AV_FIELD_TB          = AvFieldOrder(C.AV_FIELD_TB)
	AV_FIELD_BT          = AvFieldOrder(C.AV_FIELD_BT)
)

// AvCodecParameters
type AvCodecParameters C.struct_AVCodecParameters

// Custom: GetCodecType gets `AVCodecParameters.codec_type` value.
func (par *AvCodecParameters) GetCodecType() AvMediaType {
	return (AvMediaType)(par.codec_type)
}

// Custom: SetCodecType sets `AVCodecParameters.codec_type` value.
func (par *AvCodecParameters) SetCodecType(v AvMediaType) {
	par.codec_type = (C.enum_AVMediaType)(v)
}

// Custom: GetCodecTypeAddr gets `AVCodecParameters.codec_type` address.
func (par *AvCodecParameters) GetCodecTypeAddr() *AvMediaType {
	return (*AvMediaType)(unsafe.Pointer(&par.codec_type))
}

// Custom: GetCodecId gets `AVCodecParameters.codec_id` value.
func (par *AvCodecParameters) GetCodecId() AvCodecID {
	return (AvCodecID)(par.codec_id)
}

// Custom: SetCodecId sets `AVCodecParameters.codec_id` value.
func (par *AvCodecParameters) SetCodecId(v AvCodecID) {
	par.codec_id = (C.enum_AVCodecID)(v)
}

// Custom: GetCodecIdAddr gets `AVCodecParameters.codec_id` address.
func (par *AvCodecParameters) GetCodecIdAddr() *AvCodecID {
	return (*AvCodecID)(unsafe.Pointer(&par.codec_id))
}

// Custom: GetCodecTag gets `AVCodecParameters.codec_tag` value.
func (par *AvCodecParameters) GetCodecTag() uint32 {
	return (uint32)(par.codec_tag)
}

// Custom: SetCodecTag sets `AVCodecParameters.codec_tag` value.
func (par *AvCodecParameters) SetCodecTag(v uint32) {
	par.codec_tag = (C.uint)(v)
}

// Custom: GetCodecTagAddr gets `AVCodecParameters.codec_tag` address.
func (par *AvCodecParameters) GetCodecTagAddr() *uint32 {
	return (*uint32)(&par.codec_tag)
}

// Custom: GetExtradata gets `AVCodecParameters.extradata` value.
func (par *AvCodecParameters) GetExtradata() *uint8 {
	return (*uint8)(par.extradata)
}

// Custom: SetExtradata sets `AVCodecParameters.extradata` value.
func (par *AvCodecParameters) SetExtradata(v *uint8) {
	par.extradata = (*C.uint8_t)(v)
}

// Custom: GetExtradataAddr gets `AVCodecParameters.extradata` address.
func (par *AvCodecParameters) GetExtradataAddr() *uint8 {
	return (*uint8)(unsafe.Pointer(&par.extradata))
}

// Custom: GetExtradataSize gets `AVCodecParameters.extradata_size` value.
func (par *AvCodecParameters) GetExtradataSize() int32 {
	return (int32)(par.extradata_size)
}

// Custom: SetExtradataSize sets `AVCodecParameters.extradata_size` value.
func (par *AvCodecParameters) SetExtradataSize(v int32) {
	par.extradata_size = (C.int)(v)
}

// Custom: GetExtradataSizeAddr gets `AVCodecParameters.extradata_size` address.
func (par *AvCodecParameters) GetExtradataSizeAddr() *int32 {
	return (*int32)(&par.extradata_size)
}

// Custom: GetFormat gets `AVCodecParameters.format` value.
func (par *AvCodecParameters) GetFormat() int32 {
	return (int32)(par.format)
}

// Custom: SetFormat sets `AVCodecParameters.format` value.
func (par *AvCodecParameters) SetFormat(v int32) {
	par.format = (C.int)(v)
}

// Custom: GetFormatAddr gets `AVCodecParameters.format` address.
func (par *AvCodecParameters) GetFormatAddr() *int32 {
	return (*int32)(&par.format)
}

// Custom: GetBitRate gets `AVCodecParameters.bit_rate` value.
func (par *AvCodecParameters) GetBitRate() int64 {
	return (int64)(par.bit_rate)
}

// Custom: SetBitRate sets `AVCodecParameters.bit_rate` value.
func (par *AvCodecParameters) SetBitRate(v int64) {
	par.bit_rate = (C.int64_t)(v)
}

// Custom: GetBitRateAddr gets `AVCodecParameters.bit_rate` address.
func (par *AvCodecParameters) GetBitRateAddr() *int64 {
	return (*int64)(&par.bit_rate)
}

// Custom: GetBitsPerCodedSample gets `AVCodecParameters.bits_per_coded_sample` value.
func (par *AvCodecParameters) GetBitsPerCodedSample() int32 {
	return (int32)(par.bits_per_coded_sample)
}

// Custom: SetBitsPerCodedSample sets `AVCodecParameters.bits_per_coded_sample` value.
func (par *AvCodecParameters) SetBitsPerCodedSample(v int32) {
	par.bits_per_coded_sample = (C.int)(v)
}

// Custom: GetBitsPerCodedSampleAddr gets `AVCodecParameters.bits_per_coded_sample` address.
func (par *AvCodecParameters) GetBitsPerCodedSampleAddr() *int32 {
	return (*int32)(&par.bits_per_coded_sample)
}

// Custom: GetBitsPerRawSample gets `AVCodecParameters.bits_per_raw_sample` value.
func (par *AvCodecParameters) GetBitsPerRawSample() int32 {
	return (int32)(par.bits_per_raw_sample)
}

// Custom: SetBitsPerRawSample sets `AVCodecParameters.bits_per_raw_sample` value.
func (par *AvCodecParameters) SetBitsPerRawSample(v int32) {
	par.bits_per_raw_sample = (C.int)(v)
}

// Custom: GetBitsPerRawSampleAddr gets `AVCodecParameters.bits_per_raw_sample` address.
func (par *AvCodecParameters) GetBitsPerRawSampleAddr() *int32 {
	return (*int32)(&par.bits_per_raw_sample)
}

// Custom: GetProfile gets `AVCodecParameters.profile` value.
func (par *AvCodecParameters) GetProfile() int32 {
	return (int32)(par.profile)
}

// Custom: SetProfile sets `AVCodecParameters.profile` value.
func (par *AvCodecParameters) SetProfile(v int32) {
	par.profile = (C.int)(v)
}

// Custom: GetProfileAddr gets `AVCodecParameters.profile` address.
func (par *AvCodecParameters) GetProfileAddr() *int32 {
	return (*int32)(&par.profile)
}

// Custom: GetLevel gets `AVCodecParameters.level` value.
func (par *AvCodecParameters) GetLevel() int32 {
	return (int32)(par.level)
}

// Custom: SetLevel sets `AVCodecParameters.level` value.
func (par *AvCodecParameters) SetLevel(v int32) {
	par.level = (C.int)(v)
}

// Custom: GetLevelAddr gets `AVCodecParameters.level` address.
func (par *AvCodecParameters) GetLevelAddr() *int32 {
	return (*int32)(&par.level)
}

// Custom: GetWidth gets `AVCodecParameters.width` value.
func (par *AvCodecParameters) GetWidth() int32 {
	return (int32)(par.width)
}

// Custom: SetWidth sets `AVCodecParameters.width` value.
func (par *AvCodecParameters) SetWidth(v int32) {
	par.width = (C.int)(v)
}

// Custom: GetWidthAddr gets `AVCodecParameters.width` address.
func (par *AvCodecParameters) GetWidthAddr() *int32 {
	return (*int32)(&par.width)
}

// Custom: GetHeight gets `AVCodecParameters.height` value.
func (par *AvCodecParameters) GetHeight() int32 {
	return (int32)(par.height)
}

// Custom: SetHeight sets `AVCodecParameters.height` value.
func (par *AvCodecParameters) SetHeight(v int32) {
	par.height = (C.int)(v)
}

// Custom: GetHeightAddr gets `AVCodecParameters.height` address.
func (par *AvCodecParameters) GetHeightAddr() *int32 {
	return (*int32)(&par.height)
}

// Custom: GetSampleAspectRatio gets `AVCodecParameters.sample_aspect_ratio` value.
func (par *AvCodecParameters) GetSampleAspectRatio() AvRational {
	return (AvRational)(par.sample_aspect_ratio)
}

// Custom: SetSampleAspectRatio sets `AVCodecParameters.sample_aspect_ratio` value.
func (par *AvCodecParameters) SetSampleAspectRatio(v AvRational) {
	par.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// Custom: GetSampleAspectRatioAddr gets `AVCodecParameters.sample_aspect_ratio` address.
func (par *AvCodecParameters) GetSampleAspectRatioAddr() *AvRational {
	return (*AvRational)(&par.sample_aspect_ratio)
}

// Custom: GetFieldOrder gets `AVCodecParameters.field_order` value.
func (par *AvCodecParameters) GetFieldOrder() AvFieldOrder {
	return (AvFieldOrder)(par.field_order)
}

// Custom: SetFieldOrder sets `AVCodecParameters.field_order` value.
func (par *AvCodecParameters) SetFieldOrder(v AvFieldOrder) {
	par.field_order = (C.enum_AVFieldOrder)(v)
}

// Custom: GetFieldOrderAddr gets `AVCodecParameters.field_order` address.
func (par *AvCodecParameters) GetFieldOrderAddr() *AvFieldOrder {
	return (*AvFieldOrder)(unsafe.Pointer(&par.field_order))
}

// Custom: GetColorRange gets `AVCodecParameters.color_range` value.
func (par *AvCodecParameters) GetColorRange() AvColorRange {
	return (AvColorRange)(par.color_range)
}

// Custom: SetColorRange sets `AVCodecParameters.color_range` value.
func (par *AvCodecParameters) SetColorRange(v AvColorRange) {
	par.color_range = (C.enum_AVColorRange)(v)
}

// Custom: GetColorRangeAddr gets `AVCodecParameters.color_range` address.
func (par *AvCodecParameters) GetColorRangeAddr() *AvColorRange {
	return (*AvColorRange)(unsafe.Pointer(&par.color_range))
}

// Custom: GetColorPrimaries gets `AVCodecParameters.color_primaries` value.
func (par *AvCodecParameters) GetColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(par.color_primaries)
}

// Custom: SetColorPrimaries sets `AVCodecParameters.color_primaries` value.
func (par *AvCodecParameters) SetColorPrimaries(v AvColorPrimaries) {
	par.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// Custom: GetColorPrimariesAddr gets `AVCodecParameters.color_primaries` address.
func (par *AvCodecParameters) GetColorPrimariesAddr() *AvColorPrimaries {
	return (*AvColorPrimaries)(unsafe.Pointer(&par.color_primaries))
}

// Custom: GetColorTrc gets `AVCodecParameters.color_trc` value.
func (par *AvCodecParameters) GetColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(par.color_trc)
}

// Custom: SetColorTrc sets `AVCodecParameters.color_trc` value.
func (par *AvCodecParameters) SetColorTrc(v AvColorTransferCharacteristic) {
	par.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// Custom: GetColorTrcAddr gets `AVCodecParameters.color_trc` address.
func (par *AvCodecParameters) GetColorTrcAddr() *AvColorTransferCharacteristic {
	return (*AvColorTransferCharacteristic)(unsafe.Pointer(&par.color_trc))
}

// Custom: GetColorSpace gets `AVCodecParameters.color_space` value.
func (par *AvCodecParameters) GetColorSpace() AvColorSpace {
	return (AvColorSpace)(par.color_space)
}

// Custom: SetColorSpace sets `AVCodecParameters.color_space` value.
func (par *AvCodecParameters) SetColorSpace(v AvColorSpace) {
	par.color_space = (C.enum_AVColorSpace)(v)
}

// Custom: GetColorSpaceAddr gets `AVCodecParameters.color_space` address.
func (par *AvCodecParameters) GetColorSpaceAddr() *AvColorSpace {
	return (*AvColorSpace)(unsafe.Pointer(&par.color_space))
}

// Custom: GetChromaLocation gets `AVCodecParameters.chroma_location` value.
func (par *AvCodecParameters) GetChromaLocation() AvChromaLocation {
	return (AvChromaLocation)(par.chroma_location)
}

// Custom: SetChromaLocation sets `AVCodecParameters.chroma_location` value.
func (par *AvCodecParameters) SetChromaLocation(v AvChromaLocation) {
	par.chroma_location = (C.enum_AVChromaLocation)(v)
}

// Custom: GetChromaLocationAddr gets `AVCodecParameters.chroma_location` address.
func (par *AvCodecParameters) GetChromaLocationAddr() *AvChromaLocation {
	return (*AvChromaLocation)(unsafe.Pointer(&par.chroma_location))
}

// Custom: GetVideoDelay gets `AVCodecParameters.video_delay` value.
func (par *AvCodecParameters) GetVideoDelay() int32 {
	return (int32)(par.video_delay)
}

// Custom: SetVideoDelay sets `AVCodecParameters.video_delay` value.
func (par *AvCodecParameters) SetVideoDelay(v int32) {
	par.video_delay = (C.int)(v)
}

// Custom: GetVideoDelayAddr gets `AVCodecParameters.video_delay` address.
func (par *AvCodecParameters) GetVideoDelayAddr() *int32 {
	return (*int32)(&par.video_delay)
}

// Custom: GetChannelLayout gets `AVCodecParameters.channel_layout` value.
func (par *AvCodecParameters) GetChannelLayout() uint64 {
	return (uint64)(par.channel_layout)
}

// Custom: SetChannelLayout sets `AVCodecParameters.channel_layout` value.
func (par *AvCodecParameters) SetChannelLayout(v uint64) {
	par.channel_layout = (C.uint64_t)(v)
}

// Custom: GetChannelLayoutAddr gets `AVCodecParameters.channel_layout` address.
func (par *AvCodecParameters) GetChannelLayoutAddr() *uint64 {
	return (*uint64)(&par.channel_layout)
}

// Custom: GetChannels gets `AVCodecParameters.channels` value.
func (par *AvCodecParameters) GetChannels() int32 {
	return (int32)(par.channels)
}

// Custom: SetChannels sets `AVCodecParameters.channels` value.
func (par *AvCodecParameters) SetChannels(v int32) {
	par.channels = (C.int)(v)
}

// Custom: GetChannelsAddr gets `AVCodecParameters.channels` address.
func (par *AvCodecParameters) GetChannelsAddr() *int32 {
	return (*int32)(&par.channels)
}

// Custom: GetSampleRate gets `AVCodecParameters.sample_rate` value.
func (par *AvCodecParameters) GetSampleRate() int32 {
	return (int32)(par.sample_rate)
}

// Custom: SetSampleRate sets `AVCodecParameters.sample_rate` value.
func (par *AvCodecParameters) SetSampleRate(v int32) {
	par.sample_rate = (C.int)(v)
}

// Custom: GetSampleRateAddr gets `AVCodecParameters.sample_rate` address.
func (par *AvCodecParameters) GetSampleRateAddr() *int32 {
	return (*int32)(&par.sample_rate)
}

// Custom: GetBlockAlign gets `AVCodecParameters.block_align` value.
func (par *AvCodecParameters) GetBlockAlign() int32 {
	return (int32)(par.block_align)
}

// Custom: SetBlockAlign sets `AVCodecParameters.block_align` value.
func (par *AvCodecParameters) SetBlockAlign(v int32) {
	par.block_align = (C.int)(v)
}

// Custom: GetBlockAlignAddr gets `AVCodecParameters.block_align` address.
func (par *AvCodecParameters) GetBlockAlignAddr() *int32 {
	return (*int32)(&par.block_align)
}

// Custom: GetFrameSize gets `AVCodecParameters.frame_size` value.
func (par *AvCodecParameters) GetFrameSize() int32 {
	return (int32)(par.frame_size)
}

// Custom: SetFrameSize sets `AVCodecParameters.frame_size` value.
func (par *AvCodecParameters) SetFrameSize(v int32) {
	par.frame_size = (C.int)(v)
}

// Custom: GetFrameSizeAddr gets `AVCodecParameters.frame_size` address.
func (par *AvCodecParameters) GetFrameSizeAddr() *int32 {
	return (*int32)(&par.frame_size)
}

// Custom: GetInitialPadding gets `AVCodecParameters.initial_padding` value.
func (par *AvCodecParameters) GetInitialPadding() int32 {
	return (int32)(par.initial_padding)
}

// Custom: SetInitialPadding sets `AVCodecParameters.initial_padding` value.
func (par *AvCodecParameters) SetInitialPadding(v int32) {
	par.initial_padding = (C.int)(v)
}

// Custom: GetInitialPaddingAddr gets `AVCodecParameters.initial_padding` address.
func (par *AvCodecParameters) GetInitialPaddingAddr() *int32 {
	return (*int32)(&par.initial_padding)
}

// Custom: GetTrailingPadding gets `AVCodecParameters.trailing_padding` value.
func (par *AvCodecParameters) GetTrailingPadding() int32 {
	return (int32)(par.trailing_padding)
}

// Custom: SetTrailingPadding sets `AVCodecParameters.trailing_padding` value.
func (par *AvCodecParameters) SetTrailingPadding(v int32) {
	par.trailing_padding = (C.int)(v)
}

// Custom: GetTrailingPaddingAddr gets `AVCodecParameters.trailing_padding` address.
func (par *AvCodecParameters) GetTrailingPaddingAddr() *int32 {
	return (*int32)(&par.trailing_padding)
}

// Custom: GetSeekPreroll gets `AVCodecParameters.seek_preroll` value.
func (par *AvCodecParameters) GetSeekPreroll() int32 {
	return (int32)(par.seek_preroll)
}

// Custom: SetSeekPreroll sets `AVCodecParameters.seek_preroll` value.
func (par *AvCodecParameters) SetSeekPreroll(v int32) {
	par.seek_preroll = (C.int)(v)
}

// Custom: GetSeekPrerollAddr gets `AVCodecParameters.seek_preroll` address.
func (par *AvCodecParameters) GetSeekPrerollAddr() *int32 {
	return (*int32)(&par.seek_preroll)
}

// AvCodecParametersAlloc allocates a new AVCodecParameters and set its fields to default values
// (unknown/invalid/0). The returned struct must be freed with AvCodecParametersFree().
func AvCodecParametersAlloc() *AvCodecParameters {
	return (*AvCodecParameters)(C.avcodec_parameters_alloc())
}

// AvCodecParametersFree frees an AVCodecParameters instance and everything associated with it and
// write NULL to the supplied pointer.
func AvCodecParametersFree(par **AvCodecParameters) {
	C.avcodec_parameters_free((**C.struct_AVCodecParameters)(unsafe.Pointer(par)))
}

// AvCodecParametersCopy copies the contents of src to dst.
func AvCodecParametersCopy(dst, src *AvCodecParameters) int32 {
	return (int32)(C.avcodec_parameters_copy((*C.struct_AVCodecParameters)(dst),
		(*C.struct_AVCodecParameters)(src)))
}
