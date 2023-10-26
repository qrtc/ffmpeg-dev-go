package ffmpeg

/*
#include <libavcodec/dirac.h>
*/
import "C"
import "unsafe"

const (
	MAX_DWT_LEVELS = C.MAX_DWT_LEVELS
)

type DiracParseCodes = C.enum_DiracParseCodes

const (
	DIRAC_PCODE_SEQ_HEADER      = DiracParseCodes(C.DIRAC_PCODE_SEQ_HEADER)
	DIRAC_PCODE_END_SEQ         = DiracParseCodes(C.DIRAC_PCODE_END_SEQ)
	DIRAC_PCODE_AUX             = DiracParseCodes(C.DIRAC_PCODE_AUX)
	DIRAC_PCODE_PAD             = DiracParseCodes(C.DIRAC_PCODE_PAD)
	DIRAC_PCODE_PICTURE_CODED   = DiracParseCodes(C.DIRAC_PCODE_PICTURE_CODED)
	DIRAC_PCODE_PICTURE_RAW     = DiracParseCodes(C.DIRAC_PCODE_PICTURE_RAW)
	DIRAC_PCODE_PICTURE_LOW_DEL = DiracParseCodes(C.DIRAC_PCODE_PICTURE_LOW_DEL)
	DIRAC_PCODE_PICTURE_HQ      = DiracParseCodes(C.DIRAC_PCODE_PICTURE_HQ)
	DIRAC_PCODE_INTER_NOREF_CO1 = DiracParseCodes(C.DIRAC_PCODE_INTER_NOREF_CO1)
	DIRAC_PCODE_INTER_NOREF_CO2 = DiracParseCodes(C.DIRAC_PCODE_INTER_NOREF_CO2)
	DIRAC_PCODE_INTER_REF_CO1   = DiracParseCodes(C.DIRAC_PCODE_INTER_REF_CO1)
	DIRAC_PCODE_INTER_REF_CO2   = DiracParseCodes(C.DIRAC_PCODE_INTER_REF_CO2)
	DIRAC_PCODE_INTRA_REF_CO    = DiracParseCodes(C.DIRAC_PCODE_INTRA_REF_CO)
	DIRAC_PCODE_INTRA_REF_RAW   = DiracParseCodes(C.DIRAC_PCODE_INTRA_REF_RAW)
	DIRAC_PCODE_INTRA_REF_PICT  = DiracParseCodes(C.DIRAC_PCODE_INTRA_REF_PICT)
	DIRAC_PCODE_MAGIC           = DiracParseCodes(C.DIRAC_PCODE_MAGIC)
)

// DiracVersionInfo
type DiracVersionInfo C.struct_DiracVersionInfo

// GetMajor gets `DiracVersionInfo.major` value.
func (dvi *DiracVersionInfo) GetMajor() int32 {
	return (int32)(dvi.major)
}

// SetMajor sets `DiracVersionInfo.major` value.
func (dvi *DiracVersionInfo) SetMajor(v int32) {
	dvi.major = (C.int)(v)
}

// GetMajorAddr gets `DiracVersionInfo.major` address.
func (dvi *DiracVersionInfo) GetMajorAddr() *int32 {
	return (*int32)(&dvi.major)
}

// GetMinor gets `DiracVersionInfo.minor` value.
func (dvi *DiracVersionInfo) GetMinor() int32 {
	return (int32)(dvi.minor)
}

// SetMinor sets `DiracVersionInfo.minor` value.
func (dvi *DiracVersionInfo) SetMinor(v int32) {
	dvi.minor = (C.int)(v)
}

// GetMinorAddr gets `DiracVersionInfo.minor` address.
func (dvi *DiracVersionInfo) GetMinorAddr() *int32 {
	return (*int32)(&dvi.minor)
}

// AVDiracSeqHeader
type AVDiracSeqHeader C.struct_AVDiracSeqHeader

// GetWidth gets `AVDiracSeqHeader.width` value.
func (dsh *AVDiracSeqHeader) GetWidth() uint32 {
	return (uint32)(dsh.width)
}

// SetWidth sets `AVDiracSeqHeader.width` value.
func (dsh *AVDiracSeqHeader) SetWidth(v uint32) {
	dsh.width = (C.uint)(v)
}

// GetWidthAddr gets `AVDiracSeqHeader.width` address.
func (dsh *AVDiracSeqHeader) GetWidthAddr() *uint32 {
	return (*uint32)(&dsh.width)
}

// GetHeight gets `AVDiracSeqHeader.height` value.
func (dsh *AVDiracSeqHeader) GetHeight() uint32 {
	return (uint32)(dsh.height)
}

// SetHeight sets `AVDiracSeqHeader.height` value.
func (dsh *AVDiracSeqHeader) SetHeight(v uint32) {
	dsh.height = (C.uint)(v)
}

// GetHeightAddr gets `AVDiracSeqHeader.height` address.
func (dsh *AVDiracSeqHeader) GetHeightAddr() *uint32 {
	return (*uint32)(&dsh.height)
}

// GetChromaFormat gets `AVDiracSeqHeader.chroma_format` value.
func (dsh *AVDiracSeqHeader) GetChromaFormat() uint8 {
	return (uint8)(dsh.chroma_format)
}

// SetChromaFormat sets `AVDiracSeqHeader.chroma_format` value.
func (dsh *AVDiracSeqHeader) SetChromaFormat(v uint8) {
	dsh.chroma_format = (C.uint8_t)(v)
}

// GetChromaFormatAddr gets `AVDiracSeqHeader.chroma_format` address.
func (dsh *AVDiracSeqHeader) GetChromaFormatAddr() *uint8 {
	return (*uint8)(&dsh.chroma_format)
}

// GetInterlaced gets `AVDiracSeqHeader.interlaced` value.
func (dsh *AVDiracSeqHeader) GetInterlaced() uint8 {
	return (uint8)(dsh.interlaced)
}

// SetInterlaced sets `AVDiracSeqHeader.interlaced` value.
func (dsh *AVDiracSeqHeader) SetInterlaced(v uint8) {
	dsh.interlaced = (C.uint8_t)(v)
}

// GetInterlacedAddr gets `AVDiracSeqHeader.interlaced` address.
func (dsh *AVDiracSeqHeader) GetInterlacedAddr() *uint8 {
	return (*uint8)(&dsh.interlaced)
}

// GetTopFieldFirst gets `AVDiracSeqHeader.top_field_first` value.
func (dsh *AVDiracSeqHeader) GetTopFieldFirst() uint8 {
	return (uint8)(dsh.top_field_first)
}

// SetTopFieldFirst sets `AVDiracSeqHeader.top_field_first` value.
func (dsh *AVDiracSeqHeader) SetTopFieldFirst(v uint8) {
	dsh.top_field_first = (C.uint8_t)(v)
}

// GetTopFieldFirstAddr gets `AVDiracSeqHeader.top_field_first` address.
func (dsh *AVDiracSeqHeader) GetTopFieldFirstAddr() *uint8 {
	return (*uint8)(&dsh.top_field_first)
}

// GetFrameRateIndex gets `AVDiracSeqHeader.frame_rate_index` value.
func (dsh *AVDiracSeqHeader) GetFrameRateIndex() uint8 {
	return (uint8)(dsh.frame_rate_index)
}

// SetFrameRateIndex sets `AVDiracSeqHeader.frame_rate_index` value.
func (dsh *AVDiracSeqHeader) SetFrameRateIndex(v uint8) {
	dsh.frame_rate_index = (C.uint8_t)(v)
}

// GetFrameRateIndexAddr gets `AVDiracSeqHeader.frame_rate_index` address.
func (dsh *AVDiracSeqHeader) GetFrameRateIndexAddr() *uint8 {
	return (*uint8)(&dsh.frame_rate_index)
}

// GetAspectRatioIndex gets `AVDiracSeqHeader.aspect_ratio_index` value.
func (dsh *AVDiracSeqHeader) GetAspectRatioIndex() uint8 {
	return (uint8)(dsh.aspect_ratio_index)
}

// SetAspectRatioIndex sets `AVDiracSeqHeader.aspect_ratio_index` value.
func (dsh *AVDiracSeqHeader) SetAspectRatioIndex(v uint8) {
	dsh.aspect_ratio_index = (C.uint8_t)(v)
}

// GetAspectRatioIndexAddr gets `AVDiracSeqHeader.aspect_ratio_index` address.
func (dsh *AVDiracSeqHeader) GetAspectRatioIndexAddr() *uint8 {
	return (*uint8)(&dsh.aspect_ratio_index)
}

// GetCleanWidth gets `AVDiracSeqHeader.clean_width` value.
func (dsh *AVDiracSeqHeader) GetCleanWidth() uint16 {
	return (uint16)(dsh.clean_width)
}

// SetCleanWidth sets `AVDiracSeqHeader.clean_width` value.
func (dsh *AVDiracSeqHeader) SetCleanWidth(v uint16) {
	dsh.clean_width = (C.uint16_t)(v)
}

// GetCleanWidthAddr gets `AVDiracSeqHeader.clean_width` address.
func (dsh *AVDiracSeqHeader) GetCleanWidthAddr() *uint16 {
	return (*uint16)(&dsh.clean_width)
}

// GetCleanHeight gets `AVDiracSeqHeader.clean_height` value.
func (dsh *AVDiracSeqHeader) GetCleanHeight() uint16 {
	return (uint16)(dsh.clean_height)
}

// SetCleanHeight sets `AVDiracSeqHeader.clean_height` value.
func (dsh *AVDiracSeqHeader) SetCleanHeight(v uint16) {
	dsh.clean_height = (C.uint16_t)(v)
}

// GetCleanHeightAddr gets `AVDiracSeqHeader.clean_height` address.
func (dsh *AVDiracSeqHeader) GetCleanHeightAddr() *uint16 {
	return (*uint16)(&dsh.clean_height)
}

// GetCleanLeftOffset gets `AVDiracSeqHeader.clean_left_offset` value.
func (dsh *AVDiracSeqHeader) GetCleanLeftOffset() uint16 {
	return (uint16)(dsh.clean_left_offset)
}

// SetCleanLeftOffset sets `AVDiracSeqHeader.clean_left_offset` value.
func (dsh *AVDiracSeqHeader) SetCleanLeftOffset(v uint16) {
	dsh.clean_left_offset = (C.uint16_t)(v)
}

// GetCleanLeftOffsetAddr gets `AVDiracSeqHeader.clean_left_offset` address.
func (dsh *AVDiracSeqHeader) GetCleanLeftOffsetAddr() *uint16 {
	return (*uint16)(&dsh.clean_left_offset)
}

// GetCleanRightOffset gets `AVDiracSeqHeader.clean_right_offset` value.
func (dsh *AVDiracSeqHeader) GetCleanRightOffset() uint16 {
	return (uint16)(dsh.clean_right_offset)
}

// SetCleanRightOffset sets `AVDiracSeqHeader.clean_right_offset` value.
func (dsh *AVDiracSeqHeader) SetCleanRightOffset(v uint16) {
	dsh.clean_right_offset = (C.uint16_t)(v)
}

// GetCleanRightOffsetAddr gets `AVDiracSeqHeader.clean_right_offset` address.
func (dsh *AVDiracSeqHeader) GetCleanRightOffsetAddr() *uint16 {
	return (*uint16)(&dsh.clean_right_offset)
}

// GetPixelRangeIndex gets `AVDiracSeqHeader.pixel_range_index` value.
func (dsh *AVDiracSeqHeader) GetPixelRangeIndex() uint8 {
	return (uint8)(dsh.pixel_range_index)
}

// SetPixelRangeIndex sets `AVDiracSeqHeader.pixel_range_index` value.
func (dsh *AVDiracSeqHeader) SetPixelRangeIndex(v uint8) {
	dsh.pixel_range_index = (C.uint8_t)(v)
}

// GetPixelRangeIndexAddr gets `AVDiracSeqHeader.pixel_range_index` address.
func (dsh *AVDiracSeqHeader) GetPixelRangeIndexAddr() *uint8 {
	return (*uint8)(&dsh.pixel_range_index)
}

// GetColorSpecIndex gets `AVDiracSeqHeader.color_spec_index` value.
func (dsh *AVDiracSeqHeader) GetColorSpecIndex() uint8 {
	return (uint8)(dsh.color_spec_index)
}

// SetColorSpecIndex sets `AVDiracSeqHeader.color_spec_index` value.
func (dsh *AVDiracSeqHeader) SetColorSpecIndex(v uint8) {
	dsh.color_spec_index = (C.uint8_t)(v)
}

// GetColorSpecIndexAddr gets `AVDiracSeqHeader.color_spec_index` address.
func (dsh *AVDiracSeqHeader) GetColorSpecIndexAddr() *uint8 {
	return (*uint8)(&dsh.color_spec_index)
}

// GetProfile gets `AVDiracSeqHeader.profile` value.
func (dsh *AVDiracSeqHeader) GetProfile() int32 {
	return (int32)(dsh.profile)
}

// SetProfile sets `AVDiracSeqHeader.profile` value.
func (dsh *AVDiracSeqHeader) SetProfile(v int32) {
	dsh.profile = (C.int)(v)
}

// GetProfileAddr gets `AVDiracSeqHeader.profile` address.
func (dsh *AVDiracSeqHeader) GetProfileAddr() *int32 {
	return (*int32)(&dsh.profile)
}

// GetLevel gets `AVDiracSeqHeader.level` value.
func (dsh *AVDiracSeqHeader) GetLevel() int32 {
	return (int32)(dsh.level)
}

// SetLevel sets `AVDiracSeqHeader.level` value.
func (dsh *AVDiracSeqHeader) SetLevel(v int32) {
	dsh.level = (C.int)(v)
}

// GetLevelAddr gets `AVDiracSeqHeader.level` address.
func (dsh *AVDiracSeqHeader) GetLevelAddr() *int32 {
	return (*int32)(&dsh.level)
}

// GetFramerate gets `AVDiracSeqHeader.framerate` value.
func (dsh *AVDiracSeqHeader) GetFramerate() AVRational {
	return (AVRational)(dsh.framerate)
}

// SetFramerate sets `AVDiracSeqHeader.framerate` value.
func (dsh *AVDiracSeqHeader) SetFramerate(v AVRational) {
	dsh.framerate = (C.struct_AVRational)(v)
}

// GetFramerateAddr gets `AVDiracSeqHeader.framerate` address.
func (dsh *AVDiracSeqHeader) GetFramerateAddr() *AVRational {
	return (*AVRational)(&dsh.framerate)
}

// GetSampleAspectRatio gets `AVDiracSeqHeader.sample_aspect_ratio` value.
func (dsh *AVDiracSeqHeader) GetSampleAspectRatio() AVRational {
	return (AVRational)(dsh.sample_aspect_ratio)
}

// SetSampleAspectRatio sets `AVDiracSeqHeader.sample_aspect_ratio` value.
func (dsh *AVDiracSeqHeader) SetSampleAspectRatio(v AVRational) {
	dsh.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// GetSampleAspectRatioAddr gets `AVDiracSeqHeader.sample_aspect_ratio` address.
func (dsh *AVDiracSeqHeader) GetSampleAspectRatioAddr() *AVRational {
	return (*AVRational)(&dsh.sample_aspect_ratio)
}

// GetPixFmt gets `AVDiracSeqHeader.pix_fmt` value.
func (dsh *AVDiracSeqHeader) GetPixFmt() AVPixelFormat {
	return (AVPixelFormat)(dsh.pix_fmt)
}

// SetPixFmt sets `AVDiracSeqHeader.pix_fmt` value.
func (dsh *AVDiracSeqHeader) SetPixFmt(v AVPixelFormat) {
	dsh.pix_fmt = (C.enum_AVPixelFormat)(v)
}

// GetPixFmtAddr gets `AVDiracSeqHeader.pix_fmt` address.
func (dsh *AVDiracSeqHeader) GetPixFmtAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&dsh.pix_fmt)
}

// GetColorRange gets `AVDiracSeqHeader.colorrange` value.
func (dsh *AVDiracSeqHeader) GetColorRange() AVColorRange {
	return (AVColorRange)(dsh.color_range)
}

// SetColorRange sets `AVDiracSeqHeader.colorrange` value.
func (dsh *AVDiracSeqHeader) SetColorRange(v AVColorRange) {
	dsh.color_range = (C.enum_AVColorRange)(v)
}

// GetColorRangeAddr gets `AVDiracSeqHeader.colorrange` address.
func (dsh *AVDiracSeqHeader) GetColorRangeAddr() *AVColorRange {
	return (*AVColorRange)(&dsh.color_range)
}

// GetColorPrimaries gets `AVDiracSeqHeader.color_primaries` value.
func (dsh *AVDiracSeqHeader) GetColorPrimaries() AVColorPrimaries {
	return (AVColorPrimaries)(dsh.color_primaries)
}

// SetColorPrimaries sets `AVDiracSeqHeader.color_primaries` value.
func (dsh *AVDiracSeqHeader) SetColorPrimaries(v AVColorPrimaries) {
	dsh.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// GetColorPrimariesAddr gets `AVDiracSeqHeader.color_primaries` address.
func (dsh *AVDiracSeqHeader) GetColorPrimariesAddr() *AVColorPrimaries {
	return (*AVColorPrimaries)(&dsh.color_primaries)
}

// GetColorTrc gets `AVDiracSeqHeader.color_trc` value.
func (dsh *AVDiracSeqHeader) GetColorTrc() AVColorTransferCharacteristic {
	return (AVColorTransferCharacteristic)(dsh.color_trc)
}

// SetColorTrc sets `AVDiracSeqHeader.color_trc` value.
func (dsh *AVDiracSeqHeader) SetColorTrc(v AVColorTransferCharacteristic) {
	dsh.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// GetColorTrcAddr gets `AVDiracSeqHeader.color_trc` address.
func (dsh *AVDiracSeqHeader) GetColorTrcAddr() *AVColorTransferCharacteristic {
	return (*AVColorTransferCharacteristic)(&dsh.color_trc)
}

// GetColorspace gets `AVDiracSeqHeader.colorspace` value.
func (dsh *AVDiracSeqHeader) GetColorspace() AVColorSpace {
	return (AVColorSpace)(dsh.colorspace)
}

// SetColorspace sets `AVDiracSeqHeader.colorspace` value.
func (dsh *AVDiracSeqHeader) SetColorspace(v AVColorSpace) {
	dsh.colorspace = (C.enum_AVColorSpace)(v)
}

// GetColorspaceAddr gets `AVDiracSeqHeader.colorspace` address.
func (dsh *AVDiracSeqHeader) GetColorspaceAddr() *AVColorSpace {
	return (*AVColorSpace)(&dsh.colorspace)
}

// GetVersion gets `AVDiracSeqHeader.version` value.
func (dsh *AVDiracSeqHeader) GetVersion() DiracVersionInfo {
	return (DiracVersionInfo)(dsh.version)
}

// SetVersion sets `AVDiracSeqHeader.version` value.
func (dsh *AVDiracSeqHeader) SetVersion(v DiracVersionInfo) {
	dsh.version = (C.struct_DiracVersionInfo)(v)
}

// GetVersionAddr gets `AVDiracSeqHeader.version` address.
func (dsh *AVDiracSeqHeader) GetVersionAddr() *DiracVersionInfo {
	return (*DiracVersionInfo)(&dsh.version)
}

// GetBitDepth gets `AVDiracSeqHeader.bit_depth` value.
func (dsh *AVDiracSeqHeader) GetBitDepth() int32 {
	return (int32)(dsh.bit_depth)
}

// SetBitDepth sets `AVDiracSeqHeader.bit_depth` value.
func (dsh *AVDiracSeqHeader) SetBitDepth(v int32) {
	dsh.bit_depth = (C.int)(v)
}

// GetBitDepthAddr gets `AVDiracSeqHeader.bit_depth` address.
func (dsh *AVDiracSeqHeader) GetBitDepthAddr() *int32 {
	return (*int32)(&dsh.bit_depth)
}

// AvDiracParseSequenceHeader parses a Dirac sequence header.
func AvDiracParseSequenceHeader(dsh **AVDiracSeqHeader, buf *uint8, bufSize uintptr, logCtx CVoidPointer) int32 {
	return (int32)(C.av_dirac_parse_sequence_header((**C.struct_AVDiracSeqHeader)(unsafe.Pointer(dsh)),
		(*C.uint8_t)(buf), (C.size_t)(bufSize), VoidPointer(logCtx)))
}
