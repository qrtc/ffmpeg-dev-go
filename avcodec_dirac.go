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

// Custom: GetMajor gets `DiracVersionInfo.major` value.
func (dvi *DiracVersionInfo) GetMajor() int32 {
	return (int32)(dvi.major)
}

// Custom: SetMajor sets `DiracVersionInfo.major` value.
func (dvi *DiracVersionInfo) SetMajor(v int32) {
	dvi.major = (C.int)(v)
}

// Custom: GetMajorAddr gets `DiracVersionInfo.major` address.
func (dvi *DiracVersionInfo) GetMajorAddr() *int32 {
	return (*int32)(&dvi.major)
}

// Custom: GetMinor gets `DiracVersionInfo.minor` value.
func (dvi *DiracVersionInfo) GetMinor() int32 {
	return (int32)(dvi.minor)
}

// Custom: SetMinor sets `DiracVersionInfo.minor` value.
func (dvi *DiracVersionInfo) SetMinor(v int32) {
	dvi.minor = (C.int)(v)
}

// Custom: GetMinorAddr gets `DiracVersionInfo.minor` address.
func (dvi *DiracVersionInfo) GetMinorAddr() *int32 {
	return (*int32)(&dvi.minor)
}

// AVDiracSeqHeader
type AVDiracSeqHeader C.struct_AVDiracSeqHeader

// Custom: GetWidth gets `AVDiracSeqHeader.width` value.
func (dsh *AVDiracSeqHeader) GetWidth() uint32 {
	return (uint32)(dsh.width)
}

// Custom: SetWidth sets `AVDiracSeqHeader.width` value.
func (dsh *AVDiracSeqHeader) SetWidth(v uint32) {
	dsh.width = (C.uint)(v)
}

// Custom: GetWidthAddr gets `AVDiracSeqHeader.width` address.
func (dsh *AVDiracSeqHeader) GetWidthAddr() *uint32 {
	return (*uint32)(&dsh.width)
}

// Custom: GetHeight gets `AVDiracSeqHeader.height` value.
func (dsh *AVDiracSeqHeader) GetHeight() uint32 {
	return (uint32)(dsh.height)
}

// Custom: SetHeight sets `AVDiracSeqHeader.height` value.
func (dsh *AVDiracSeqHeader) SetHeight(v uint32) {
	dsh.height = (C.uint)(v)
}

// Custom: GetHeightAddr gets `AVDiracSeqHeader.height` address.
func (dsh *AVDiracSeqHeader) GetHeightAddr() *uint32 {
	return (*uint32)(&dsh.height)
}

// Custom: GetChromaFormat gets `AVDiracSeqHeader.chroma_format` value.
func (dsh *AVDiracSeqHeader) GetChromaFormat() uint8 {
	return (uint8)(dsh.chroma_format)
}

// Custom: SetChromaFormat sets `AVDiracSeqHeader.chroma_format` value.
func (dsh *AVDiracSeqHeader) SetChromaFormat(v uint8) {
	dsh.chroma_format = (C.uint8_t)(v)
}

// Custom: GetChromaFormatAddr gets `AVDiracSeqHeader.chroma_format` address.
func (dsh *AVDiracSeqHeader) GetChromaFormatAddr() *uint8 {
	return (*uint8)(&dsh.chroma_format)
}

// Custom: GetInterlaced gets `AVDiracSeqHeader.interlaced` value.
func (dsh *AVDiracSeqHeader) GetInterlaced() uint8 {
	return (uint8)(dsh.interlaced)
}

// Custom: SetInterlaced sets `AVDiracSeqHeader.interlaced` value.
func (dsh *AVDiracSeqHeader) SetInterlaced(v uint8) {
	dsh.interlaced = (C.uint8_t)(v)
}

// Custom: GetInterlacedAddr gets `AVDiracSeqHeader.interlaced` address.
func (dsh *AVDiracSeqHeader) GetInterlacedAddr() *uint8 {
	return (*uint8)(&dsh.interlaced)
}

// Custom: GetTopFieldFirst gets `AVDiracSeqHeader.top_field_first` value.
func (dsh *AVDiracSeqHeader) GetTopFieldFirst() uint8 {
	return (uint8)(dsh.top_field_first)
}

// Custom: SetTopFieldFirst sets `AVDiracSeqHeader.top_field_first` value.
func (dsh *AVDiracSeqHeader) SetTopFieldFirst(v uint8) {
	dsh.top_field_first = (C.uint8_t)(v)
}

// Custom: GetTopFieldFirstAddr gets `AVDiracSeqHeader.top_field_first` address.
func (dsh *AVDiracSeqHeader) GetTopFieldFirstAddr() *uint8 {
	return (*uint8)(&dsh.top_field_first)
}

// Custom: GetFrameRateIndex gets `AVDiracSeqHeader.frame_rate_index` value.
func (dsh *AVDiracSeqHeader) GetFrameRateIndex() uint8 {
	return (uint8)(dsh.frame_rate_index)
}

// Custom: SetFrameRateIndex sets `AVDiracSeqHeader.frame_rate_index` value.
func (dsh *AVDiracSeqHeader) SetFrameRateIndex(v uint8) {
	dsh.frame_rate_index = (C.uint8_t)(v)
}

// Custom: GetFrameRateIndexAddr gets `AVDiracSeqHeader.frame_rate_index` address.
func (dsh *AVDiracSeqHeader) GetFrameRateIndexAddr() *uint8 {
	return (*uint8)(&dsh.frame_rate_index)
}

// Custom: GetAspectRatioIndex gets `AVDiracSeqHeader.aspect_ratio_index` value.
func (dsh *AVDiracSeqHeader) GetAspectRatioIndex() uint8 {
	return (uint8)(dsh.aspect_ratio_index)
}

// Custom: SetAspectRatioIndex sets `AVDiracSeqHeader.aspect_ratio_index` value.
func (dsh *AVDiracSeqHeader) SetAspectRatioIndex(v uint8) {
	dsh.aspect_ratio_index = (C.uint8_t)(v)
}

// Custom: GetAspectRatioIndexAddr gets `AVDiracSeqHeader.aspect_ratio_index` address.
func (dsh *AVDiracSeqHeader) GetAspectRatioIndexAddr() *uint8 {
	return (*uint8)(&dsh.aspect_ratio_index)
}

// Custom: GetCleanWidth gets `AVDiracSeqHeader.clean_width` value.
func (dsh *AVDiracSeqHeader) GetCleanWidth() uint16 {
	return (uint16)(dsh.clean_width)
}

// Custom: SetCleanWidth sets `AVDiracSeqHeader.clean_width` value.
func (dsh *AVDiracSeqHeader) SetCleanWidth(v uint16) {
	dsh.clean_width = (C.uint16_t)(v)
}

// Custom: GetCleanWidthAddr gets `AVDiracSeqHeader.clean_width` address.
func (dsh *AVDiracSeqHeader) GetCleanWidthAddr() *uint16 {
	return (*uint16)(&dsh.clean_width)
}

// Custom: GetCleanHeight gets `AVDiracSeqHeader.clean_height` value.
func (dsh *AVDiracSeqHeader) GetCleanHeight() uint16 {
	return (uint16)(dsh.clean_height)
}

// Custom: SetCleanHeight sets `AVDiracSeqHeader.clean_height` value.
func (dsh *AVDiracSeqHeader) SetCleanHeight(v uint16) {
	dsh.clean_height = (C.uint16_t)(v)
}

// Custom: GetCleanHeightAddr gets `AVDiracSeqHeader.clean_height` address.
func (dsh *AVDiracSeqHeader) GetCleanHeightAddr() *uint16 {
	return (*uint16)(&dsh.clean_height)
}

// Custom: GetCleanLeftOffset gets `AVDiracSeqHeader.clean_left_offset` value.
func (dsh *AVDiracSeqHeader) GetCleanLeftOffset() uint16 {
	return (uint16)(dsh.clean_left_offset)
}

// Custom: SetCleanLeftOffset sets `AVDiracSeqHeader.clean_left_offset` value.
func (dsh *AVDiracSeqHeader) SetCleanLeftOffset(v uint16) {
	dsh.clean_left_offset = (C.uint16_t)(v)
}

// Custom: GetCleanLeftOffsetAddr gets `AVDiracSeqHeader.clean_left_offset` address.
func (dsh *AVDiracSeqHeader) GetCleanLeftOffsetAddr() *uint16 {
	return (*uint16)(&dsh.clean_left_offset)
}

// Custom: GetCleanRightOffset gets `AVDiracSeqHeader.clean_right_offset` value.
func (dsh *AVDiracSeqHeader) GetCleanRightOffset() uint16 {
	return (uint16)(dsh.clean_right_offset)
}

// Custom: SetCleanRightOffset sets `AVDiracSeqHeader.clean_right_offset` value.
func (dsh *AVDiracSeqHeader) SetCleanRightOffset(v uint16) {
	dsh.clean_right_offset = (C.uint16_t)(v)
}

// Custom: GetCleanRightOffsetAddr gets `AVDiracSeqHeader.clean_right_offset` address.
func (dsh *AVDiracSeqHeader) GetCleanRightOffsetAddr() *uint16 {
	return (*uint16)(&dsh.clean_right_offset)
}

// Custom: GetPixelRangeIndex gets `AVDiracSeqHeader.pixel_range_index` value.
func (dsh *AVDiracSeqHeader) GetPixelRangeIndex() uint8 {
	return (uint8)(dsh.pixel_range_index)
}

// Custom: SetPixelRangeIndex sets `AVDiracSeqHeader.pixel_range_index` value.
func (dsh *AVDiracSeqHeader) SetPixelRangeIndex(v uint8) {
	dsh.pixel_range_index = (C.uint8_t)(v)
}

// Custom: GetPixelRangeIndexAddr gets `AVDiracSeqHeader.pixel_range_index` address.
func (dsh *AVDiracSeqHeader) GetPixelRangeIndexAddr() *uint8 {
	return (*uint8)(&dsh.pixel_range_index)
}

// Custom: GetColorSpecIndex gets `AVDiracSeqHeader.color_spec_index` value.
func (dsh *AVDiracSeqHeader) GetColorSpecIndex() uint8 {
	return (uint8)(dsh.color_spec_index)
}

// Custom: SetColorSpecIndex sets `AVDiracSeqHeader.color_spec_index` value.
func (dsh *AVDiracSeqHeader) SetColorSpecIndex(v uint8) {
	dsh.color_spec_index = (C.uint8_t)(v)
}

// Custom: GetColorSpecIndexAddr gets `AVDiracSeqHeader.color_spec_index` address.
func (dsh *AVDiracSeqHeader) GetColorSpecIndexAddr() *uint8 {
	return (*uint8)(&dsh.color_spec_index)
}

// Custom: GetProfile gets `AVDiracSeqHeader.profile` value.
func (dsh *AVDiracSeqHeader) GetProfile() int32 {
	return (int32)(dsh.profile)
}

// Custom: SetProfile sets `AVDiracSeqHeader.profile` value.
func (dsh *AVDiracSeqHeader) SetProfile(v int32) {
	dsh.profile = (C.int)(v)
}

// Custom: GetProfileAddr gets `AVDiracSeqHeader.profile` address.
func (dsh *AVDiracSeqHeader) GetProfileAddr() *int32 {
	return (*int32)(&dsh.profile)
}

// Custom: GetLevel gets `AVDiracSeqHeader.level` value.
func (dsh *AVDiracSeqHeader) GetLevel() int32 {
	return (int32)(dsh.level)
}

// Custom: SetLevel sets `AVDiracSeqHeader.level` value.
func (dsh *AVDiracSeqHeader) SetLevel(v int32) {
	dsh.level = (C.int)(v)
}

// Custom: GetLevelAddr gets `AVDiracSeqHeader.level` address.
func (dsh *AVDiracSeqHeader) GetLevelAddr() *int32 {
	return (*int32)(&dsh.level)
}

// Custom: GetFramerate gets `AVDiracSeqHeader.framerate` value.
func (dsh *AVDiracSeqHeader) GetFramerate() AVRational {
	return (AVRational)(dsh.framerate)
}

// Custom: SetFramerate sets `AVDiracSeqHeader.framerate` value.
func (dsh *AVDiracSeqHeader) SetFramerate(v AVRational) {
	dsh.framerate = (C.struct_AVRational)(v)
}

// Custom: GetFramerateAddr gets `AVDiracSeqHeader.framerate` address.
func (dsh *AVDiracSeqHeader) GetFramerateAddr() *AVRational {
	return (*AVRational)(&dsh.framerate)
}

// Custom: GetSampleAspectRatio gets `AVDiracSeqHeader.sample_aspect_ratio` value.
func (dsh *AVDiracSeqHeader) GetSampleAspectRatio() AVRational {
	return (AVRational)(dsh.sample_aspect_ratio)
}

// Custom: SetSampleAspectRatio sets `AVDiracSeqHeader.sample_aspect_ratio` value.
func (dsh *AVDiracSeqHeader) SetSampleAspectRatio(v AVRational) {
	dsh.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// Custom: GetSampleAspectRatioAddr gets `AVDiracSeqHeader.sample_aspect_ratio` address.
func (dsh *AVDiracSeqHeader) GetSampleAspectRatioAddr() *AVRational {
	return (*AVRational)(&dsh.sample_aspect_ratio)
}

// Custom: GetPixFmt gets `AVDiracSeqHeader.pix_fmt` value.
func (dsh *AVDiracSeqHeader) GetPixFmt() AVPixelFormat {
	return (AVPixelFormat)(dsh.pix_fmt)
}

// Custom: SetPixFmt sets `AVDiracSeqHeader.pix_fmt` value.
func (dsh *AVDiracSeqHeader) SetPixFmt(v AVPixelFormat) {
	dsh.pix_fmt = (C.enum_AVPixelFormat)(v)
}

// Custom: GetPixFmtAddr gets `AVDiracSeqHeader.pix_fmt` address.
func (dsh *AVDiracSeqHeader) GetPixFmtAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&dsh.pix_fmt)
}

// Custom: GetColorRange gets `AVDiracSeqHeader.colorrange` value.
func (dsh *AVDiracSeqHeader) GetColorRange() AVColorRange {
	return (AVColorRange)(dsh.color_range)
}

// Custom: SetColorRange sets `AVDiracSeqHeader.colorrange` value.
func (dsh *AVDiracSeqHeader) SetColorRange(v AVColorRange) {
	dsh.color_range = (C.enum_AVColorRange)(v)
}

// Custom: GetColorRangeAddr gets `AVDiracSeqHeader.colorrange` address.
func (dsh *AVDiracSeqHeader) GetColorRangeAddr() *AVColorRange {
	return (*AVColorRange)(&dsh.color_range)
}

// Custom: GetColorPrimaries gets `AVDiracSeqHeader.color_primaries` value.
func (dsh *AVDiracSeqHeader) GetColorPrimaries() AVColorPrimaries {
	return (AVColorPrimaries)(dsh.color_primaries)
}

// Custom: SetColorPrimaries sets `AVDiracSeqHeader.color_primaries` value.
func (dsh *AVDiracSeqHeader) SetColorPrimaries(v AVColorPrimaries) {
	dsh.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// Custom: GetColorPrimariesAddr gets `AVDiracSeqHeader.color_primaries` address.
func (dsh *AVDiracSeqHeader) GetColorPrimariesAddr() *AVColorPrimaries {
	return (*AVColorPrimaries)(&dsh.color_primaries)
}

// Custom: GetColorTrc gets `AVDiracSeqHeader.color_trc` value.
func (dsh *AVDiracSeqHeader) GetColorTrc() AVColorTransferCharacteristic {
	return (AVColorTransferCharacteristic)(dsh.color_trc)
}

// Custom: SetColorTrc sets `AVDiracSeqHeader.color_trc` value.
func (dsh *AVDiracSeqHeader) SetColorTrc(v AVColorTransferCharacteristic) {
	dsh.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// Custom: GetColorTrcAddr gets `AVDiracSeqHeader.color_trc` address.
func (dsh *AVDiracSeqHeader) GetColorTrcAddr() *AVColorTransferCharacteristic {
	return (*AVColorTransferCharacteristic)(&dsh.color_trc)
}

// Custom: GetColorspace gets `AVDiracSeqHeader.colorspace` value.
func (dsh *AVDiracSeqHeader) GetColorspace() AVColorSpace {
	return (AVColorSpace)(dsh.colorspace)
}

// Custom: SetColorspace sets `AVDiracSeqHeader.colorspace` value.
func (dsh *AVDiracSeqHeader) SetColorspace(v AVColorSpace) {
	dsh.colorspace = (C.enum_AVColorSpace)(v)
}

// Custom: GetColorspaceAddr gets `AVDiracSeqHeader.colorspace` address.
func (dsh *AVDiracSeqHeader) GetColorspaceAddr() *AVColorSpace {
	return (*AVColorSpace)(&dsh.colorspace)
}

// Custom: GetVersion gets `AVDiracSeqHeader.version` value.
func (dsh *AVDiracSeqHeader) GetVersion() DiracVersionInfo {
	return (DiracVersionInfo)(dsh.version)
}

// Custom: SetVersion sets `AVDiracSeqHeader.version` value.
func (dsh *AVDiracSeqHeader) SetVersion(v DiracVersionInfo) {
	dsh.version = (C.struct_DiracVersionInfo)(v)
}

// Custom: GetVersionAddr gets `AVDiracSeqHeader.version` address.
func (dsh *AVDiracSeqHeader) GetVersionAddr() *DiracVersionInfo {
	return (*DiracVersionInfo)(&dsh.version)
}

// Custom: GetBitDepth gets `AVDiracSeqHeader.bit_depth` value.
func (dsh *AVDiracSeqHeader) GetBitDepth() int32 {
	return (int32)(dsh.bit_depth)
}

// Custom: SetBitDepth sets `AVDiracSeqHeader.bit_depth` value.
func (dsh *AVDiracSeqHeader) SetBitDepth(v int32) {
	dsh.bit_depth = (C.int)(v)
}

// Custom: GetBitDepthAddr gets `AVDiracSeqHeader.bit_depth` address.
func (dsh *AVDiracSeqHeader) GetBitDepthAddr() *int32 {
	return (*int32)(&dsh.bit_depth)
}

// AvDiracParseSequenceHeader parses a Dirac sequence header.
func AvDiracParseSequenceHeader(dsh **AVDiracSeqHeader, buf *uint8, bufSize uintptr, logCtx CVoidPointer) int32 {
	return (int32)(C.av_dirac_parse_sequence_header((**C.struct_AVDiracSeqHeader)(unsafe.Pointer(dsh)),
		(*C.uint8_t)(buf), (C.size_t)(bufSize), VoidPointer(logCtx)))
}
