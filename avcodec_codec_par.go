// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

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

// GetCodecType gets `AVCodecParameters.codectype` value.
func (par *AVCodecParameters) GetCodecType() AVMediaType {
	return (AVMediaType)(par.codec_type)
}

// SetCodecType sets `AVCodecParameters.codectype` value.
func (par *AVCodecParameters) SetCodecType(v AVMediaType) {
	par.codec_type = (C.enum_AVMediaType)(v)
}

// GetCodecTypeAddr gets `AVCodecParameters.codectype` address.
func (par *AVCodecParameters) GetCodecTypeAddr() *AVMediaType {
	return (*AVMediaType)(unsafe.Pointer(&par.codec_type))
}

// GetCodecId gets `AVCodecParameters.codec_id` value.
func (par *AVCodecParameters) GetCodecId() AVCodecID {
	return (AVCodecID)(par.codec_id)
}

// SetCodecId sets `AVCodecParameters.codec_id` value.
func (par *AVCodecParameters) SetCodecId(v AVCodecID) {
	par.codec_id = (C.enum_AVCodecID)(v)
}

// GetCodecIdAddr gets `AVCodecParameters.codec_id` address.
func (par *AVCodecParameters) GetCodecIdAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&par.codec_id))
}

// GetCodecTag gets `AVCodecParameters.codec_tag` value.
func (par *AVCodecParameters) GetCodecTag() uint32 {
	return (uint32)(par.codec_tag)
}

// SetCodecTag sets `AVCodecParameters.codec_tag` value.
func (par *AVCodecParameters) SetCodecTag(v uint32) {
	par.codec_tag = (C.uint)(v)
}

// GetCodecTagAddr gets `AVCodecParameters.codec_tag` address.
func (par *AVCodecParameters) GetCodecTagAddr() *uint32 {
	return (*uint32)(&par.codec_tag)
}

// GetExtradata gets `AVCodecParameters.extradata` value.
func (par *AVCodecParameters) GetExtradata() *uint8 {
	return (*uint8)(par.extradata)
}

// SetExtradata sets `AVCodecParameters.extradata` value.
func (par *AVCodecParameters) SetExtradata(v *uint8) {
	par.extradata = (*C.uint8_t)(v)
}

// GetExtradataAddr gets `AVCodecParameters.extradata` address.
func (par *AVCodecParameters) GetExtradataAddr() *uint8 {
	return (*uint8)(unsafe.Pointer(&par.extradata))
}

// GetExtradataSize gets `AVCodecParameters.extradata_size` value.
func (par *AVCodecParameters) GetExtradataSize() int32 {
	return (int32)(par.extradata_size)
}

// SetExtradataSize sets `AVCodecParameters.extradata_size` value.
func (par *AVCodecParameters) SetExtradataSize(v int32) {
	par.extradata_size = (C.int)(v)
}

// GetExtradataSizeAddr gets `AVCodecParameters.extradata_size` address.
func (par *AVCodecParameters) GetExtradataSizeAddr() *int32 {
	return (*int32)(&par.extradata_size)
}

// GetFormat gets `AVCodecParameters.format` value.
func (par *AVCodecParameters) GetFormat() int32 {
	return (int32)(par.format)
}

// SetFormat sets `AVCodecParameters.format` value.
func (par *AVCodecParameters) SetFormat(v int32) {
	par.format = (C.int)(v)
}

// GetFormatAddr gets `AVCodecParameters.format` address.
func (par *AVCodecParameters) GetFormatAddr() *int32 {
	return (*int32)(&par.format)
}

// GetBitRate gets `AVCodecParameters.bit_rate` value.
func (par *AVCodecParameters) GetBitRate() int64 {
	return (int64)(par.bit_rate)
}

// SetBitRate sets `AVCodecParameters.bit_rate` value.
func (par *AVCodecParameters) SetBitRate(v int64) {
	par.bit_rate = (C.int64_t)(v)
}

// GetBitRateAddr gets `AVCodecParameters.bit_rate` address.
func (par *AVCodecParameters) GetBitRateAddr() *int64 {
	return (*int64)(&par.bit_rate)
}

// GetBitsPerCodedSample gets `AVCodecParameters.bits_per_coded_sample` value.
func (par *AVCodecParameters) GetBitsPerCodedSample() int32 {
	return (int32)(par.bits_per_coded_sample)
}

// SetBitsPerCodedSample sets `AVCodecParameters.bits_per_coded_sample` value.
func (par *AVCodecParameters) SetBitsPerCodedSample(v int32) {
	par.bits_per_coded_sample = (C.int)(v)
}

// GetBitsPerCodedSampleAddr gets `AVCodecParameters.bits_per_coded_sample` address.
func (par *AVCodecParameters) GetBitsPerCodedSampleAddr() *int32 {
	return (*int32)(&par.bits_per_coded_sample)
}

// GetBitsPerRawSample gets `AVCodecParameters.bits_per_raw_sample` value.
func (par *AVCodecParameters) GetBitsPerRawSample() int32 {
	return (int32)(par.bits_per_raw_sample)
}

// SetBitsPerRawSample sets `AVCodecParameters.bits_per_raw_sample` value.
func (par *AVCodecParameters) SetBitsPerRawSample(v int32) {
	par.bits_per_raw_sample = (C.int)(v)
}

// GetBitsPerRawSampleAddr gets `AVCodecParameters.bits_per_raw_sample` address.
func (par *AVCodecParameters) GetBitsPerRawSampleAddr() *int32 {
	return (*int32)(&par.bits_per_raw_sample)
}

// GetProfile gets `AVCodecParameters.profile` value.
func (par *AVCodecParameters) GetProfile() int32 {
	return (int32)(par.profile)
}

// SetProfile sets `AVCodecParameters.profile` value.
func (par *AVCodecParameters) SetProfile(v int32) {
	par.profile = (C.int)(v)
}

// GetProfileAddr gets `AVCodecParameters.profile` address.
func (par *AVCodecParameters) GetProfileAddr() *int32 {
	return (*int32)(&par.profile)
}

// GetLevel gets `AVCodecParameters.level` value.
func (par *AVCodecParameters) GetLevel() int32 {
	return (int32)(par.level)
}

// SetLevel sets `AVCodecParameters.level` value.
func (par *AVCodecParameters) SetLevel(v int32) {
	par.level = (C.int)(v)
}

// GetLevelAddr gets `AVCodecParameters.level` address.
func (par *AVCodecParameters) GetLevelAddr() *int32 {
	return (*int32)(&par.level)
}

// GetWidth gets `AVCodecParameters.width` value.
func (par *AVCodecParameters) GetWidth() int32 {
	return (int32)(par.width)
}

// SetWidth sets `AVCodecParameters.width` value.
func (par *AVCodecParameters) SetWidth(v int32) {
	par.width = (C.int)(v)
}

// GetWidthAddr gets `AVCodecParameters.width` address.
func (par *AVCodecParameters) GetWidthAddr() *int32 {
	return (*int32)(&par.width)
}

// GetHeight gets `AVCodecParameters.height` value.
func (par *AVCodecParameters) GetHeight() int32 {
	return (int32)(par.height)
}

// SetHeight sets `AVCodecParameters.height` value.
func (par *AVCodecParameters) SetHeight(v int32) {
	par.height = (C.int)(v)
}

// GetHeightAddr gets `AVCodecParameters.height` address.
func (par *AVCodecParameters) GetHeightAddr() *int32 {
	return (*int32)(&par.height)
}

// GetSampleAspectRatio gets `AVCodecParameters.sample_aspect_ratio` value.
func (par *AVCodecParameters) GetSampleAspectRatio() AVRational {
	return (AVRational)(par.sample_aspect_ratio)
}

// SetSampleAspectRatio sets `AVCodecParameters.sample_aspect_ratio` value.
func (par *AVCodecParameters) SetSampleAspectRatio(v AVRational) {
	par.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// GetSampleAspectRatioAddr gets `AVCodecParameters.sample_aspect_ratio` address.
func (par *AVCodecParameters) GetSampleAspectRatioAddr() *AVRational {
	return (*AVRational)(&par.sample_aspect_ratio)
}

// GetFieldOrder gets `AVCodecParameters.field_order` value.
func (par *AVCodecParameters) GetFieldOrder() AVFieldOrder {
	return (AVFieldOrder)(par.field_order)
}

// SetFieldOrder sets `AVCodecParameters.field_order` value.
func (par *AVCodecParameters) SetFieldOrder(v AVFieldOrder) {
	par.field_order = (C.enum_AVFieldOrder)(v)
}

// GetFieldOrderAddr gets `AVCodecParameters.field_order` address.
func (par *AVCodecParameters) GetFieldOrderAddr() *AVFieldOrder {
	return (*AVFieldOrder)(unsafe.Pointer(&par.field_order))
}

// GetColorRange gets `AVCodecParameters.colorrange` value.
func (par *AVCodecParameters) GetColorRange() AVColorRange {
	return (AVColorRange)(par.color_range)
}

// SetColorRange sets `AVCodecParameters.colorrange` value.
func (par *AVCodecParameters) SetColorRange(v AVColorRange) {
	par.color_range = (C.enum_AVColorRange)(v)
}

// GetColorRangeAddr gets `AVCodecParameters.colorrange` address.
func (par *AVCodecParameters) GetColorRangeAddr() *AVColorRange {
	return (*AVColorRange)(unsafe.Pointer(&par.color_range))
}

// GetColorPrimaries gets `AVCodecParameters.color_primaries` value.
func (par *AVCodecParameters) GetColorPrimaries() AVColorPrimaries {
	return (AVColorPrimaries)(par.color_primaries)
}

// SetColorPrimaries sets `AVCodecParameters.color_primaries` value.
func (par *AVCodecParameters) SetColorPrimaries(v AVColorPrimaries) {
	par.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// GetColorPrimariesAddr gets `AVCodecParameters.color_primaries` address.
func (par *AVCodecParameters) GetColorPrimariesAddr() *AVColorPrimaries {
	return (*AVColorPrimaries)(unsafe.Pointer(&par.color_primaries))
}

// GetColorTrc gets `AVCodecParameters.color_trc` value.
func (par *AVCodecParameters) GetColorTrc() AVColorTransferCharacteristic {
	return (AVColorTransferCharacteristic)(par.color_trc)
}

// SetColorTrc sets `AVCodecParameters.color_trc` value.
func (par *AVCodecParameters) SetColorTrc(v AVColorTransferCharacteristic) {
	par.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// GetColorTrcAddr gets `AVCodecParameters.color_trc` address.
func (par *AVCodecParameters) GetColorTrcAddr() *AVColorTransferCharacteristic {
	return (*AVColorTransferCharacteristic)(unsafe.Pointer(&par.color_trc))
}

// GetColorSpace gets `AVCodecParameters.color_space` value.
func (par *AVCodecParameters) GetColorSpace() AVColorSpace {
	return (AVColorSpace)(par.color_space)
}

// SetColorSpace sets `AVCodecParameters.color_space` value.
func (par *AVCodecParameters) SetColorSpace(v AVColorSpace) {
	par.color_space = (C.enum_AVColorSpace)(v)
}

// GetColorSpaceAddr gets `AVCodecParameters.color_space` address.
func (par *AVCodecParameters) GetColorSpaceAddr() *AVColorSpace {
	return (*AVColorSpace)(unsafe.Pointer(&par.color_space))
}

// GetChromaLocation gets `AVCodecParameters.chroma_location` value.
func (par *AVCodecParameters) GetChromaLocation() AVChromaLocation {
	return (AVChromaLocation)(par.chroma_location)
}

// SetChromaLocation sets `AVCodecParameters.chroma_location` value.
func (par *AVCodecParameters) SetChromaLocation(v AVChromaLocation) {
	par.chroma_location = (C.enum_AVChromaLocation)(v)
}

// GetChromaLocationAddr gets `AVCodecParameters.chroma_location` address.
func (par *AVCodecParameters) GetChromaLocationAddr() *AVChromaLocation {
	return (*AVChromaLocation)(unsafe.Pointer(&par.chroma_location))
}

// GetVideoDelay gets `AVCodecParameters.video_delay` value.
func (par *AVCodecParameters) GetVideoDelay() int32 {
	return (int32)(par.video_delay)
}

// SetVideoDelay sets `AVCodecParameters.video_delay` value.
func (par *AVCodecParameters) SetVideoDelay(v int32) {
	par.video_delay = (C.int)(v)
}

// GetVideoDelayAddr gets `AVCodecParameters.video_delay` address.
func (par *AVCodecParameters) GetVideoDelayAddr() *int32 {
	return (*int32)(&par.video_delay)
}

// GetChannelLayout gets `AVCodecParameters.channel_layout` value.
func (par *AVCodecParameters) GetChannelLayout() uint64 {
	return (uint64)(par.channel_layout)
}

// SetChannelLayout sets `AVCodecParameters.channel_layout` value.
func (par *AVCodecParameters) SetChannelLayout(v uint64) {
	par.channel_layout = (C.uint64_t)(v)
}

// GetChannelLayoutAddr gets `AVCodecParameters.channel_layout` address.
func (par *AVCodecParameters) GetChannelLayoutAddr() *uint64 {
	return (*uint64)(&par.channel_layout)
}

// GetChannels gets `AVCodecParameters.channels` value.
func (par *AVCodecParameters) GetChannels() int32 {
	return (int32)(par.channels)
}

// SetChannels sets `AVCodecParameters.channels` value.
func (par *AVCodecParameters) SetChannels(v int32) {
	par.channels = (C.int)(v)
}

// GetChannelsAddr gets `AVCodecParameters.channels` address.
func (par *AVCodecParameters) GetChannelsAddr() *int32 {
	return (*int32)(&par.channels)
}

// GetSampleRate gets `AVCodecParameters.sample_rate` value.
func (par *AVCodecParameters) GetSampleRate() int32 {
	return (int32)(par.sample_rate)
}

// SetSampleRate sets `AVCodecParameters.sample_rate` value.
func (par *AVCodecParameters) SetSampleRate(v int32) {
	par.sample_rate = (C.int)(v)
}

// GetSampleRateAddr gets `AVCodecParameters.sample_rate` address.
func (par *AVCodecParameters) GetSampleRateAddr() *int32 {
	return (*int32)(&par.sample_rate)
}

// GetBlockAlign gets `AVCodecParameters.block_align` value.
func (par *AVCodecParameters) GetBlockAlign() int32 {
	return (int32)(par.block_align)
}

// SetBlockAlign sets `AVCodecParameters.block_align` value.
func (par *AVCodecParameters) SetBlockAlign(v int32) {
	par.block_align = (C.int)(v)
}

// GetBlockAlignAddr gets `AVCodecParameters.block_align` address.
func (par *AVCodecParameters) GetBlockAlignAddr() *int32 {
	return (*int32)(&par.block_align)
}

// GetFrameSize gets `AVCodecParameters.frame_size` value.
func (par *AVCodecParameters) GetFrameSize() int32 {
	return (int32)(par.frame_size)
}

// SetFrameSize sets `AVCodecParameters.frame_size` value.
func (par *AVCodecParameters) SetFrameSize(v int32) {
	par.frame_size = (C.int)(v)
}

// GetFrameSizeAddr gets `AVCodecParameters.frame_size` address.
func (par *AVCodecParameters) GetFrameSizeAddr() *int32 {
	return (*int32)(&par.frame_size)
}

// GetInitialPadding gets `AVCodecParameters.initial_padding` value.
func (par *AVCodecParameters) GetInitialPadding() int32 {
	return (int32)(par.initial_padding)
}

// SetInitialPadding sets `AVCodecParameters.initial_padding` value.
func (par *AVCodecParameters) SetInitialPadding(v int32) {
	par.initial_padding = (C.int)(v)
}

// GetInitialPaddingAddr gets `AVCodecParameters.initial_padding` address.
func (par *AVCodecParameters) GetInitialPaddingAddr() *int32 {
	return (*int32)(&par.initial_padding)
}

// GetTrailingPadding gets `AVCodecParameters.trailing_padding` value.
func (par *AVCodecParameters) GetTrailingPadding() int32 {
	return (int32)(par.trailing_padding)
}

// SetTrailingPadding sets `AVCodecParameters.trailing_padding` value.
func (par *AVCodecParameters) SetTrailingPadding(v int32) {
	par.trailing_padding = (C.int)(v)
}

// GetTrailingPaddingAddr gets `AVCodecParameters.trailing_padding` address.
func (par *AVCodecParameters) GetTrailingPaddingAddr() *int32 {
	return (*int32)(&par.trailing_padding)
}

// GetSeekPreroll gets `AVCodecParameters.seek_preroll` value.
func (par *AVCodecParameters) GetSeekPreroll() int32 {
	return (int32)(par.seek_preroll)
}

// SetSeekPreroll sets `AVCodecParameters.seek_preroll` value.
func (par *AVCodecParameters) SetSeekPreroll(v int32) {
	par.seek_preroll = (C.int)(v)
}

// GetSeekPrerollAddr gets `AVCodecParameters.seek_preroll` address.
func (par *AVCodecParameters) GetSeekPrerollAddr() *int32 {
	return (*int32)(&par.seek_preroll)
}

// GetChLayout gets `AVCodecParameters.ch_layout` value.
func (par *AVCodecParameters) GetChLayout() AVChannelLayout {
	return (AVChannelLayout)(par.ch_layout)
}

// SetChLayout sets `AVCodecParameters.ch_layout` value.
func (par *AVCodecParameters) SetChLayout(v AVChannelLayout) {
	par.ch_layout = (C.struct_AVChannelLayout)(v)
}

// GetChLayoutAddr gets `AVCodecParameters.ch_layout` address.
func (par *AVCodecParameters) GetChLayoutAddr() *AVChannelLayout {
	return (*AVChannelLayout)(&par.ch_layout)
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

// AvGetAudioFrameDuration2 returns audio frame duration.
func AvGetAudioFrameDuration2(par *AVCodecParameters, frameBytes int32) int32 {
	return (int32)(C.av_get_audio_frame_duration2((*C.struct_AVCodecParameters)(par), (C.int)(frameBytes)))
}
