package ffmpeg

/*
#include <libavutil/frame.h>
*/
import "C"
import "unsafe"

// AvFrameSideDataType
type AvFrameSideDataType int32

const (
	AV_FRAME_DATA_PANSCAN                    = AvFrameSideDataType(C.AV_FRAME_DATA_PANSCAN)
	AV_FRAME_DATA_A53_CC                     = AvFrameSideDataType(C.AV_FRAME_DATA_A53_CC)
	AV_FRAME_DATA_STEREO3D                   = AvFrameSideDataType(C.AV_FRAME_DATA_STEREO3D)
	AV_FRAME_DATA_MATRIXENCODING             = AvFrameSideDataType(C.AV_FRAME_DATA_MATRIXENCODING)
	AV_FRAME_DATA_DOWNMIX_INFO               = AvFrameSideDataType(C.AV_FRAME_DATA_DOWNMIX_INFO)
	AV_FRAME_DATA_REPLAYGAIN                 = AvFrameSideDataType(C.AV_FRAME_DATA_REPLAYGAIN)
	AV_FRAME_DATA_DISPLAYMATRIX              = AvFrameSideDataType(C.AV_FRAME_DATA_DISPLAYMATRIX)
	AV_FRAME_DATA_AFD                        = AvFrameSideDataType(C.AV_FRAME_DATA_AFD)
	AV_FRAME_DATA_MOTION_VECTORS             = AvFrameSideDataType(C.AV_FRAME_DATA_MOTION_VECTORS)
	AV_FRAME_DATA_SKIP_SAMPLES               = AvFrameSideDataType(C.AV_FRAME_DATA_SKIP_SAMPLES)
	AV_FRAME_DATA_AUDIO_SERVICE_TYPE         = AvFrameSideDataType(C.AV_FRAME_DATA_AUDIO_SERVICE_TYPE)
	AV_FRAME_DATA_MASTERING_DISPLAY_METADATA = AvFrameSideDataType(C.AV_FRAME_DATA_MASTERING_DISPLAY_METADATA)
	AV_FRAME_DATA_GOP_TIMECODE               = AvFrameSideDataType(C.AV_FRAME_DATA_GOP_TIMECODE)
	AV_FRAME_DATA_SPHERICAL                  = AvFrameSideDataType(C.AV_FRAME_DATA_SPHERICAL)
	AV_FRAME_DATA_CONTENT_LIGHT_LEVEL        = AvFrameSideDataType(C.AV_FRAME_DATA_CONTENT_LIGHT_LEVEL)
	AV_FRAME_DATA_ICC_PROFILE                = AvFrameSideDataType(C.AV_FRAME_DATA_ICC_PROFILE)
	AV_FRAME_DATA_QP_TABLE_PROPERTIES        = AvFrameSideDataType(C.AV_FRAME_DATA_QP_TABLE_PROPERTIES)
	AV_FRAME_DATA_QP_TABLE_DATA              = AvFrameSideDataType(C.AV_FRAME_DATA_QP_TABLE_DATA)
	AV_FRAME_DATA_S12M_TIMECODE              = AvFrameSideDataType(C.AV_FRAME_DATA_S12M_TIMECODE)
	AV_FRAME_DATA_DYNAMIC_HDR_PLUS           = AvFrameSideDataType(C.AV_FRAME_DATA_DYNAMIC_HDR_PLUS)
	AV_FRAME_DATA_REGIONS_OF_INTEREST        = AvFrameSideDataType(C.AV_FRAME_DATA_REGIONS_OF_INTEREST)
	AV_FRAME_DATA_VIDEO_ENC_PARAMS           = AvFrameSideDataType(C.AV_FRAME_DATA_VIDEO_ENC_PARAMS)
	AV_FRAME_DATA_SEI_UNREGISTERED           = AvFrameSideDataType(C.AV_FRAME_DATA_SEI_UNREGISTERED)
	AV_FRAME_DATA_FILM_GRAIN_PARAMS          = AvFrameSideDataType(C.AV_FRAME_DATA_FILM_GRAIN_PARAMS)
)

// AvActiveFormatDescription
type AvActiveFormatDescription int32

const (
	AV_AFD_SAME         = AvActiveFormatDescription(C.AV_AFD_SAME)
	AV_AFD_4_3          = AvActiveFormatDescription(C.AV_AFD_4_3)
	AV_AFD_16_9         = AvActiveFormatDescription(C.AV_AFD_16_9)
	AV_AFD_14_9         = AvActiveFormatDescription(C.AV_AFD_14_9)
	AV_AFD_4_3_SP_14_9  = AvActiveFormatDescription(C.AV_AFD_4_3_SP_14_9)
	AV_AFD_16_9_SP_14_9 = AvActiveFormatDescription(C.AV_AFD_16_9_SP_14_9)
	AV_AFD_SP_4_3       = AvActiveFormatDescription(C.AV_AFD_SP_4_3)
)

// Structure to hold side data for an AVFrame.
type AvFrameSideData C.struct_AVFrameSideData

// Structure describing a single Region Of Interest.
type AvRegionOfInterest C.struct_AVRegionOfInterest

// AvFrame
type AvFrame C.struct_AVFrame

const (
	AV_NUM_DATA_POINTERS = C.AV_NUM_DATA_POINTERS
)

// Custom: GetDataIdx gets `AVFrame.data` index value.
func (frame *AvFrame) GetDataIdx(idx int) *uint8 {
	return (*uint8)(frame.data[idx])
}

// Custom: SetDataIdx sets `AVFrame.data` index value.
func (frame *AvFrame) SetDataIdx(idx int, v *uint8) {
	frame.data[idx] = (*C.uint8_t)(v)
}

// Custom: GetDataIdxAddr gets `AVFrame.data` index address.
func (frame *AvFrame) GetDataIdxAddr(idx int) **uint8 {
	return (**uint8)(unsafe.Pointer(&frame.data[idx]))
}

// Custom: GetData gets `AVFrame.data` value.
func (frame *AvFrame) GetData() **uint8 {
	return (**uint8)(unsafe.Pointer(&frame.data[0]))
}

// Custom: GetDataAddr gets `AVFrame.data` address.
func (frame *AvFrame) GetDataAddr() ***uint8 {
	return (***uint8)(unsafe.Pointer(&frame.data))
}

// Custom: GetLinesizeIdx gets `AVFrame.linesize` index value.
func (frame *AvFrame) GetLinesizeIdx(idx int) int32 {
	return (int32)(frame.linesize[idx])
}

// Custom: SetLinesizeIdx sets `AVFrame.linesize` index value.
func (frame *AvFrame) SetLinesizeIdx(idx int, v int32) {
	frame.linesize[idx] = (C.int)(v)
}

// Custom: GetLinesizeIdxAddr gets `AVFrame.linesize` index address.
func (frame *AvFrame) GetLinesizeIdxAddr(idx int) *int32 {
	return (*int32)(&frame.linesize[idx])
}

// Custom: GetLinesize gets `AVFrame.linesize` value.
func (frame *AvFrame) GetLinesize() *int32 {
	return (*int32)(&frame.linesize[0])
}

// Custom: GetLinesizeAddr gets `AVFrame.linesize` address.
func (frame *AvFrame) GetLinesizeAddr() **int32 {
	return (**int32)(unsafe.Pointer(&frame.linesize))
}

// Custom: GetExtendedData gets `AVFrame.extended_data` value.
func (frame *AvFrame) GetExtendedData() **uint8 {
	return (**uint8)(unsafe.Pointer(frame.extended_data))
}

// Custom: SetExtendedData sets `AVFrame.extended_data` value.
func (frame *AvFrame) SetExtendedData(v **uint8) {
	frame.extended_data = (**C.uint8_t)(unsafe.Pointer(v))
}

// Custom: GetExtendedDataAddr gets `AVFrame.extended_data` address.
func (frame *AvFrame) GetExtendedDataAddr() ***uint8 {
	return (***uint8)(unsafe.Pointer(&frame.extended_data))
}

// Custom: GetWidth gets `AVFrame.width` value.
func (frame *AvFrame) GetWidth() int32 {
	return (int32)(frame.width)
}

// Custom: SetWidth sets `AVFrame.width` value.
func (frame *AvFrame) SetWidth(v int32) {
	frame.width = (C.int)(v)
}

// Custom: GetWidthAddr gets `AVFrame.width` address.
func (frame *AvFrame) GetWidthAddr() *int32 {
	return (*int32)(&frame.width)
}

// Custom: GetHeight gets `AVFrame.height` value.
func (frame *AvFrame) GetHeight() int32 {
	return (int32)(frame.height)
}

// Custom: SetHeight sets `AVFrame.height` value.
func (frame *AvFrame) SetHeight(v int32) {
	frame.height = (C.int)(v)
}

// Custom: GetHeightAddr gets `AVFrame.height` address.
func (frame *AvFrame) GetHeightAddr() *int32 {
	return (*int32)(&frame.height)
}

// Custom: GetNbSamples gets `AVFrame.nb_samples` value.
func (frame *AvFrame) GetNbSamples() int32 {
	return (int32)(frame.nb_samples)
}

// Custom: SetNbSamples sets `AVFrame.nb_samples` value.
func (frame *AvFrame) SetNbSamples(v int32) {
	frame.nb_samples = (C.int)(v)
}

// Custom: GetNbSamplesAddr gets `AVFrame.nb_samples` address.
func (frame *AvFrame) GetNbSamplesAddr() *int32 {
	return (*int32)(&frame.nb_samples)
}

// Custom: GetFormat gets `AVFrame.format` value.
func (frame *AvFrame) GetFormat() int32 {
	return (int32)(frame.format)
}

// Custom: SetFormat sets `AVFrame.format` value.
func (frame *AvFrame) SetFormat(v int32) {
	frame.format = (C.int)(v)
}

// Custom: GetFormatAddr gets `AVFrame.format` address.
func (frame *AvFrame) GetFormatAddr() *int32 {
	return (*int32)(&frame.format)
}

// Custom: GetKeyFrame gets `AVFrame.key_frame` value.
func (frame *AvFrame) GetKeyFrame() int32 {
	return (int32)(frame.key_frame)
}

// Custom: SetKeyFrame sets `AVFrame.key_frame` value.
func (frame *AvFrame) SetKeyFrame(v int32) {
	frame.key_frame = (C.int)(v)
}

// Custom: GetKeyFrameAddr gets `AVFrame.key_frame` address.
func (frame *AvFrame) GetKeyFrameAddr() *int32 {
	return (*int32)(&frame.key_frame)
}

// Custom: GetPictType gets `AVFrame.pict_type` value.
func (frame *AvFrame) GetPictType() AvPictureType {
	return (AvPictureType)(frame.pict_type)
}

// Custom: SetPictType sets `AVFrame.pict_type` value.
func (frame *AvFrame) SetPictType(v AvPictureType) {
	frame.pict_type = (C.enum_AVPictureType)(v)
}

// Custom: GetPictTypeAddr gets `AVFrame.pict_type` address.
func (frame *AvFrame) GetPictTypeAddr() *AvPictureType {
	return (*AvPictureType)(unsafe.Pointer(&frame.pict_type))
}

// Custom: GetSampleAspectRatio gets `AVFrame.sample_aspect_ratio` value.
func (frame *AvFrame) GetSampleAspectRatio() AvRational {
	return (AvRational)(frame.sample_aspect_ratio)
}

// Custom: SetSampleAspectRatio sets `AVFrame.sample_aspect_ratio` value.
func (frame *AvFrame) SetSampleAspectRatio(v AvRational) {
	frame.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// Custom: GetSampleAspectRatioAddr gets `AVFrame.sample_aspect_ratio` address.
func (frame *AvFrame) GetSampleAspectRatioAddr() *AvRational {
	return (*AvRational)(&frame.sample_aspect_ratio)
}

// Custom: GetPts gets `AVFrame.pts` value.
func (frame *AvFrame) GetPts() int64 {
	return (int64)(frame.pts)
}

// Custom: SetPts sets `AVFrame.pts` value.
func (frame *AvFrame) SetPts(v int64) {
	frame.pts = (C.int64_t)(v)
}

// Custom: GetPtsAddr gets `AVFrame.pts` address.
func (frame *AvFrame) GetPtsAddr() *int64 {
	return (*int64)(&frame.pts)
}

// Custom: GetPktPts gets `AVFrame.pkt_pts` value.
func (frame *AvFrame) GetPktPts() int64 {
	return (int64)(frame.pkt_pts)
}

// Custom: SetPktPts sets `AVFrame.pkt_pts` value.
func (frame *AvFrame) SetPktPts(v int64) {
	frame.pkt_pts = (C.int64_t)(v)
}

// Custom: GetPktPtsAddr gets `AVFrame.pkt_pts` address.
func (frame *AvFrame) GetPktPtsAddr() *int64 {
	return (*int64)(&frame.pkt_pts)
}

// Custom: GetPktDts gets `AVFrame.pkt_dts` value.
func (frame *AvFrame) GetPktDts() int64 {
	return (int64)(frame.pkt_dts)
}

// Custom: SetPktDts sets `AVFrame.pkt_dts` value.
func (frame *AvFrame) SetPktDts(v int64) {
	frame.pkt_dts = (C.int64_t)(v)
}

// Custom: GetPktDtsAddr gets `AVFrame.pkt_dts` address.
func (frame *AvFrame) GetPktDtsAddr() *int64 {
	return (*int64)(&frame.pkt_dts)
}

// Custom: GetCodedPictureNumber gets `AVFrame.coded_picture_number` value.
func (frame *AvFrame) GetCodedPictureNumber() int32 {
	return (int32)(frame.coded_picture_number)
}

// Custom: SetCodedPictureNumber sets `AVFrame.coded_picture_number` value.
func (frame *AvFrame) SetCodedPictureNumber(v int32) {
	frame.coded_picture_number = (C.int)(v)
}

// Custom: GetCodedPictureNumberAddr gets `AVFrame.coded_picture_number` address.
func (frame *AvFrame) GetCodedPictureNumberAddr() *int32 {
	return (*int32)(&frame.coded_picture_number)
}

// Custom: GetDisplayPictureNumber gets `AVFrame.display_picture_number` value.
func (frame *AvFrame) GetDisplayPictureNumber() int32 {
	return (int32)(frame.display_picture_number)
}

// Custom: SetDisplayPictureNumber sets `AVFrame.display_picture_number` value.
func (frame *AvFrame) SetDisplayPictureNumber(v int32) {
	frame.display_picture_number = (C.int)(v)
}

// Custom: GetDisplayPictureNumberAddr gets `AVFrame.display_picture_number` address.
func (frame *AvFrame) GetDisplayPictureNumberAddr() *int32 {
	return (*int32)(&frame.display_picture_number)
}

// Custom: GetQuality gets `AVFrame.quality` value.
func (frame *AvFrame) GetQuality() int32 {
	return (int32)(frame.quality)
}

// Custom: SetQuality sets `AVFrame.quality` value.
func (frame *AvFrame) SetQuality(v int32) {
	frame.quality = (C.int)(v)
}

// Custom: GetQualityAddr gets `AVFrame.quality` address.
func (frame *AvFrame) GetQualityAddr() *int32 {
	return (*int32)(&frame.quality)
}

// Custom: GetIdxError gets `AVFrame.error` index value.
func (frame *AvFrame) GetErrorIdx(idx int) uint64 {
	return (uint64)(frame.error[idx])
}

// Custom: SetIdxError sets `AVFrame.error` index value.
func (frame *AvFrame) SetErrorIdx(idx int, v uint64) {
	frame.error[idx] = (C.uint64_t)(v)
}

// Custom: GetErrorIdxAddr gets `AVFrame.error` index address.
func (frame *AvFrame) GetErrorIdxAddr(idx int) *uint64 {
	return (*uint64)(&frame.error[idx])
}

// Custom: GetError gets `AVFrame.error` value.
func (frame *AvFrame) GetError() *uint64 {
	return (*uint64)(&frame.error[0])
}

// Custom: GetErrorAddr gets `AVFrame.error` address.
func (frame *AvFrame) GetErrorAddr() **uint64 {
	return (**uint64)(unsafe.Pointer(&frame.error))
}

// Custom: GetRepeatPict gets `AVFrame.repeat_pict` value.
func (frame *AvFrame) GetRepeatPict() int32 {
	return (int32)(frame.repeat_pict)
}

// Custom: SetRepeatPict sets `AVFrame.repeat_pict` value.
func (frame *AvFrame) SetRepeatPict(v int32) {
	frame.repeat_pict = (C.int)(v)
}

// Custom: GetRepeatPictAddr gets `AVFrame.repeat_pict` address.
func (frame *AvFrame) GetRepeatPictAddr() *int32 {
	return (*int32)(&frame.repeat_pict)
}

// Custom: GetInterlacedFrame gets `AVFrame.interlaced_frame` value.
func (frame *AvFrame) GetInterlacedFrame() int32 {
	return (int32)(frame.interlaced_frame)
}

// Custom: SetInterlacedFrame sets `AVFrame.interlaced_frame` value.
func (frame *AvFrame) SetInterlacedFrame(v int32) {
	frame.interlaced_frame = (C.int)(v)
}

// Custom: GetInterlacedFrameAddr gets `AVFrame.interlaced_frame` address.
func (frame *AvFrame) GetInterlacedFrameAddr() *int32 {
	return (*int32)(&frame.interlaced_frame)
}

// Custom: GetTopFieldFirst gets `AVFrame.top_field_first` value.
func (frame *AvFrame) GetTopFieldFirst() int32 {
	return (int32)(frame.top_field_first)
}

// Custom: SetTopFieldFirst sets `AVFrame.top_field_first` value.
func (frame *AvFrame) SetTopFieldFirst(v int32) {
	frame.top_field_first = (C.int)(v)
}

// Custom: GetTopFieldFirstAddr gets `AVFrame.top_field_first` address.
func (frame *AvFrame) GetTopFieldFirstAddr() *int32 {
	return (*int32)(&frame.top_field_first)
}

// Custom: GetPaletteHasChanged gets `AVFrame.palette_has_changed` value.
func (frame *AvFrame) GetPaletteHasChanged() int32 {
	return (int32)(frame.palette_has_changed)
}

// Custom: SetPaletteHasChanged sets `AVFrame.palette_has_changed` value.
func (frame *AvFrame) SetPaletteHasChanged(v int32) {
	frame.palette_has_changed = (C.int)(v)
}

// Custom: GetPaletteHasChangedAddr gets `AVFrame.palette_has_changed` address.
func (frame *AvFrame) GetPaletteHasChangedAddr() *int32 {
	return (*int32)(&frame.palette_has_changed)
}

// Custom: GetReorderedOpaque gets `AVFrame.reordered_opaque` value.
func (frame *AvFrame) GetReorderedOpaque() int64 {
	return (int64)(frame.reordered_opaque)
}

// Custom: SetReorderedOpaque sets `AVFrame.reordered_opaque` value.
func (frame *AvFrame) SetReorderedOpaque(v int64) {
	frame.reordered_opaque = (C.int64_t)(v)
}

// Custom: GetReorderedOpaqueAddr gets `AVFrame.reordered_opaque` address.
func (frame *AvFrame) GetReorderedOpaqueAddr() *int64 {
	return (*int64)(&frame.reordered_opaque)
}

// Custom: GetSampleRate gets `AVFrame.sample_rate` value.
func (frame *AvFrame) GetSampleRate() int32 {
	return (int32)(frame.sample_rate)
}

// Custom: SetSampleRate sets `AVFrame.sample_rate` value.
func (frame *AvFrame) SetSampleRate(v int32) {
	frame.sample_rate = (C.int)(v)
}

// Custom: GetSampleRateAddr gets `AVFrame.sample_rate` address.
func (frame *AvFrame) GetSampleRateAddr() *int32 {
	return (*int32)(&frame.sample_rate)
}

// Custom: GetChannelLayout gets `AVFrame.channel_layout` value.
func (frame *AvFrame) GetChannelLayout() uint64 {
	return (uint64)(frame.channel_layout)
}

// Custom: SetChannelLayout sets `AVFrame.channel_layout` value.
func (frame *AvFrame) SetChannelLayout(v uint64) {
	frame.channel_layout = (C.uint64_t)(v)
}

// Custom: GetChannelLayoutAddr gets `AVFrame.channel_layout` address.
func (frame *AvFrame) GetChannelLayoutAddr() *uint64 {
	return (*uint64)(&frame.channel_layout)
}

// Custom: GetBufIdx gets `AVFrame.buf` value.
func (frame *AvFrame) GetBufIdx(idx int) *AvBufferRef {
	return (*AvBufferRef)(frame.buf[idx])
}

// Custom: SetBufIdx sets `AVFrame.buf` value.
func (frame *AvFrame) SetBufIdx(idx int, v *AvBufferRef) {
	frame.buf[idx] = (*C.struct_AVBufferRef)(v)
}

// Custom: GetBufIdxAddr gets `AVFrame.buf` address.
func (frame *AvFrame) GetBufIdxAddr(idx int) **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer(&frame.buf[idx]))
}

// Custom: GetBuf gets `AVFrame.buf` value.
func (frame *AvFrame) GetBuf() **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer((&frame.buf[0])))
}

// Custom: GetBufAddr gets `AVFrame.buf` address.
func (frame *AvFrame) GetBufAddr() ***AvBufferRef {
	return (***AvBufferRef)(unsafe.Pointer(&frame.buf))
}

// Custom: GetExtendedBuf gets `AVFrame.extended_buf` value.
func (frame *AvFrame) GetExtendedBuf() **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer(frame.extended_buf))
}

// Custom: SetExtendedBuf sets `AVFrame.extended_buf` value.
func (frame *AvFrame) SetExtendedBuf(v **AvBufferRef) {
	frame.extended_buf = (**C.struct_AVBufferRef)(unsafe.Pointer(v))
}

// Custom: GetExtendedBufAddr gets `AVFrame.extended_buf` address.
func (frame *AvFrame) GetExtendedBufAddr() ***AvBufferRef {
	return (***AvBufferRef)(unsafe.Pointer(&frame.extended_buf))
}

// Custom: GetNbExtendedBuf gets `AVFrame.nb_extended_buf` value.
func (frame *AvFrame) GetNbExtendedBuf() int32 {
	return (int32)(frame.nb_extended_buf)
}

// Custom: SetNbExtendedBuf sets `AVFrame.nb_extended_buf` value.
func (frame *AvFrame) SetNbExtendedBuf(v int32) {
	frame.nb_extended_buf = (C.int)(v)
}

// Custom: GetNbExtendedBufAddr gets `AVFrame.nb_extended_buf` address.
func (frame *AvFrame) GetNbExtendedBufAddr() *int32 {
	return (*int32)(&frame.nb_extended_buf)
}

// Custom: GetSideData gets `AVFrame.side_data` value.
func (frame *AvFrame) GetSideData() **AvFrameSideData {
	return (**AvFrameSideData)(unsafe.Pointer(frame.side_data))
}

// Custom: SetSideData sets `AVFrame.side_data` value.
func (frame *AvFrame) SetSideData(v **AvFrameSideData) {
	frame.side_data = (**C.struct_AVFrameSideData)(unsafe.Pointer(v))
}

// Custom: GetSideDataAddr gets `AVFrame.side_data` address.
func (frame *AvFrame) GetSideDataAddr() ***AvFrameSideData {
	return (***AvFrameSideData)(unsafe.Pointer(&frame.side_data))
}

// Custom: GetNbSideData gets `AVFrame.nb_side_data` value.
func (frame *AvFrame) GetNbSideData() int32 {
	return (int32)(frame.nb_side_data)
}

// Custom: SetNbSideData sets `AVFrame.nb_side_data` value.
func (frame *AvFrame) SetNbSideData(v int32) {
	frame.nb_side_data = (C.int)(v)
}

// Custom: GetNbSideDataAddr gets `AVFrame.nb_side_data` address.
func (frame *AvFrame) GetNbSideDataAddr() *int32 {
	return (*int32)(&frame.nb_side_data)
}

const (
	AV_FRAME_FLAG_CORRUPT = int32(C.AV_FRAME_FLAG_CORRUPT)
	AV_FRAME_FLAG_DISCARD = int32(C.AV_FRAME_FLAG_DISCARD)
)

// Custom: GetFlags gets `AVFrame.flags` value.
func (frame *AvFrame) GetFlags() int32 {
	return (int32)(frame.flags)
}

// Custom: SetFlags sets `AVFrame.flags` value.
func (frame *AvFrame) SetFlags(v int32) {
	frame.flags = (C.int)(v)
}

// Custom: GetFlagsAddr gets `AVFrame.flags` address.
func (frame *AvFrame) GetFlagsAddr() *int32 {
	return (*int32)(&frame.flags)
}

// Custom: GetColorRange gets `AVFrame.color_range` value.
func (frame *AvFrame) GetColorRange() AvColorRange {
	return (AvColorRange)(frame.color_range)
}

// Custom: SetColorRange sets `AVFrame.color_range` value.
func (frame *AvFrame) SetColorRange(v AvColorRange) {
	frame.color_range = (C.enum_AVColorRange)(v)
}

// Custom: GetColorRangeAddr gets `AVFrame.color_range` address.
func (frame *AvFrame) GetColorRangeAddr() *AvColorRange {
	return (*AvColorRange)(unsafe.Pointer(&frame.color_range))
}

// Custom: GetColorPrimaries gets `AVFrame.color_primaries` value.
func (frame *AvFrame) GetColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(frame.color_primaries)
}

// Custom: SetColorPrimaries sets `AVFrame.color_primaries` value.
func (frame *AvFrame) SetColorPrimaries(v AvColorPrimaries) {
	frame.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// Custom: GetColorPrimariesAddr gets `AVFrame.color_primaries` address.
func (frame *AvFrame) GetColorPrimariesAddr() *AvColorPrimaries {
	return (*AvColorPrimaries)(unsafe.Pointer(&frame.color_primaries))
}

// Custom: GetColorTrc gets `AVFrame.color_trc` value.
func (frame *AvFrame) GetColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(frame.color_trc)
}

// Custom: SetColorTrc sets `AVFrame.color_trc` value.
func (frame *AvFrame) SetColorTrc(v AvColorTransferCharacteristic) {
	frame.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// Custom: GetColorTrcAddr gets `AVFrame.color_trc` address.
func (frame *AvFrame) GetColorTrcAddr() *AvColorTransferCharacteristic {
	return (*AvColorTransferCharacteristic)(unsafe.Pointer(&frame.color_trc))
}

// Custom: GetColorspace gets `AVFrame.colorspace` value.
func (frame *AvFrame) GetColorspace() AvColorSpace {
	return (AvColorSpace)(frame.colorspace)
}

// Custom: SetColorspace sets `AVFrame.colorspace` value.
func (frame *AvFrame) SetColorspace(v AvColorSpace) {
	frame.colorspace = (C.enum_AVColorSpace)(v)
}

// Custom: GetColorspaceAddr gets `AVFrame.colorspace` address.
func (frame *AvFrame) GetColorspaceAddr() *AvColorSpace {
	return (*AvColorSpace)(unsafe.Pointer(&frame.colorspace))
}

// Custom: GetChromaLocation gets `AVFrame.chroma_location` value.
func (frame *AvFrame) GetChromaLocation() AvChromaLocation {
	return (AvChromaLocation)(frame.chroma_location)
}

// Custom: SetChromaLocation sets `AVFrame.chroma_location` value.
func (frame *AvFrame) SetChromaLocation(v AvChromaLocation) {
	frame.chroma_location = (C.enum_AVChromaLocation)(v)
}

// Custom: GetChromaLocationAddr gets `AVFrame.chroma_location` address.
func (frame *AvFrame) GetChromaLocationAddr() *AvChromaLocation {
	return (*AvChromaLocation)(unsafe.Pointer(&frame.chroma_location))
}

// Custom: GetBestEffortTimestamp gets `AVFrame.best_effort_timestamp` value.
func (frame *AvFrame) GetBestEffortTimestamp() int64 {
	return (int64)(frame.best_effort_timestamp)
}

// Custom: SetBestEffortTimestamp sets `AVFrame.best_effort_timestamp` value.
func (frame *AvFrame) SetBestEffortTimestamp(v int64) {
	frame.best_effort_timestamp = (C.int64_t)(v)
}

// Custom: GetBestEffortTimestampAddr gets `AVFrame.best_effort_timestamp` address.
func (frame *AvFrame) GetBestEffortTimestampAddr() *int64 {
	return (*int64)(&frame.best_effort_timestamp)
}

// Custom: GetPktPos gets `AVFrame.pkt_pos` value.
func (frame *AvFrame) GetPktPos() int64 {
	return (int64)(frame.pkt_pos)
}

// Custom: SetPktPos sets `AVFrame.pkt_pos` value.
func (frame *AvFrame) SetPktPos(v int64) {
	frame.pkt_pos = (C.int64_t)(v)
}

// Custom: GetPktPosAddr gets `AVFrame.pkt_pos` address.
func (frame *AvFrame) GetPktPosAddr() *int64 {
	return (*int64)(&frame.pkt_pos)
}

// Custom: GetPktDuration gets `AVFrame.pkt_duration` value.
func (frame *AvFrame) GetPktDuration() int64 {
	return (int64)(frame.pkt_duration)
}

// Custom: SetPktDuration sets `AVFrame.pkt_duration` value.
func (frame *AvFrame) SetPktDuration(v int64) {
	frame.pkt_duration = (C.int64_t)(v)
}

// Custom: GetPktDurationAddr gets `AVFrame.pkt_duration` address.
func (frame *AvFrame) GetPktDurationAddr() *int64 {
	return (*int64)(&frame.pkt_duration)
}

// Custom: GetMetadata gets `AVFrame.metadata` value.
func (frame *AvFrame) GetMetadata() *AvDictionary {
	return (*AvDictionary)(frame.metadata)
}

// Custom: SetMetadata sets `AVFrame.metadata` value.
func (frame *AvFrame) SetMetadata(v *AvDictionary) {
	frame.metadata = (*C.struct_AVDictionary)(v)
}

// Custom: GetMetadataAddr gets `AVFrame.metadata` address.
func (frame *AvFrame) GetMetadataAddr() **AvDictionary {
	return (**AvDictionary)(unsafe.Pointer(&frame.metadata))
}

// Custom: GetDecodeErrorFlags gets `AVFrame.decode_error_flags` value.
func (frame *AvFrame) GetDecodeErrorFlags() int32 {
	return (int32)(frame.decode_error_flags)
}

// Custom: SetDecodeErrorFlags sets `AVFrame.decode_error_flags` value.
func (frame *AvFrame) SetDecodeErrorFlags(v int32) {
	frame.decode_error_flags = (C.int)(v)
}

// Custom: GetDecodeErrorFlagsAddr gets `AVFrame.decode_error_flags` address.
func (frame *AvFrame) GetDecodeErrorFlagsAddr() *int32 {
	return (*int32)(&frame.decode_error_flags)
}

const (
	FF_DECODE_ERROR_INVALID_BITSTREAM  = int32(C.FF_DECODE_ERROR_INVALID_BITSTREAM)
	FF_DECODE_ERROR_MISSING_REFERENCE  = int32(C.FF_DECODE_ERROR_MISSING_REFERENCE)
	FF_DECODE_ERROR_CONCEALMENT_ACTIVE = int32(C.FF_DECODE_ERROR_CONCEALMENT_ACTIVE)
	FF_DECODE_ERROR_DECODE_SLICES      = int32(C.FF_DECODE_ERROR_DECODE_SLICES)
)

// Custom: GetChannels gets `AVFrame.channels` value.
func (frame *AvFrame) GetChannels() int32 {
	return (int32)(frame.channels)
}

// Custom: SetChannels sets `AVFrame.channels` value.
func (frame *AvFrame) SetChannels(v int32) {
	frame.channels = (C.int)(v)
}

// Custom: GetChannelsAddr gets `AVFrame.channels` address.
func (frame *AvFrame) GetChannelsAddr() *int32 {
	return (*int32)(&frame.channels)
}

// Custom: GetPktSize gets `AVFrame.pkt_size` value.
func (frame *AvFrame) GetPktSize() int32 {
	return (int32)(frame.pkt_size)
}

// Custom: SetPktSize sets `AVFrame.pkt_size` value.
func (frame *AvFrame) SetPktSize(v int32) {
	frame.pkt_size = (C.int)(v)
}

// Custom: GetPktSizeAddr gets `AVFrame.pkt_size` address.
func (frame *AvFrame) GetPktSizeAddr() *int32 {
	return (*int32)(&frame.pkt_size)
}

// Custom: GetQscaleTable gets `AVFrame.qscale_table` value.
func (frame *AvFrame) GetQscaleTable() *int8 {
	return (*int8)(frame.qscale_table)
}

// Custom: SetQscaleTable sets `AVFrame.qscale_table` value.
func (frame *AvFrame) SetQscaleTable(v *int8) {
	frame.qscale_table = (*C.int8_t)(v)
}

// Custom: GetQscaleTableAddr gets `AVFrame.qscale_table` address.
func (frame *AvFrame) GetQscaleTableAddr() **int8 {
	return (**int8)(unsafe.Pointer(&frame.qscale_table))
}

// Custom: GetQstride gets `AVFrame.qstride` value.
func (frame *AvFrame) GetQstride() int32 {
	return (int32)(frame.qstride)
}

// Custom: SetQstride sets `AVFrame.qstride` value.
func (frame *AvFrame) SetQstride(v int32) {
	frame.qstride = (C.int)(v)
}

// Custom: GetQstrideAddr gets `AVFrame.qstride` address.
func (frame *AvFrame) GetQstrideAddr() *int32 {
	return (*int32)(&frame.qstride)
}

// Custom: GetQscaleType gets `AVFrame.qscale_type` value.
func (frame *AvFrame) GetQscaleType() int32 {
	return (int32)(frame.qscale_type)
}

// Custom: SetQscaleType sets `AVFrame.qscale_type` value.
func (frame *AvFrame) SetQscaleType(v int32) {
	frame.qscale_type = (C.int)(v)
}

// Custom: GetQscaleTypeAddr gets `AVFrame.qscale_type` address.
func (frame *AvFrame) GetQscaleTypeAddr() *int32 {
	return (*int32)(&frame.qscale_type)
}

// Custom: GetQpTableBuf gets `AVFrame.qp_table_buf` value.
func (frame *AvFrame) GetQpTableBuf() *AvBufferRef {
	return (*AvBufferRef)(frame.qp_table_buf)
}

// Custom: SetQpTableBuf sets `AVFrame.qp_table_buf` value.
func (frame *AvFrame) SetQpTableBuf(v *AvBufferRef) {
	frame.qp_table_buf = (*C.struct_AVBufferRef)(v)
}

// Custom: GetQpTableBufAddr gets `AVFrame.qp_table_buf` address.
func (frame *AvFrame) GetQpTableBufAddr() **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer(&frame.qp_table_buf))
}

// Custom: GetHwFramesCtx gets `AVFrame.hw_frames_ctx` value.
func (frame *AvFrame) GetHwFramesCtx() *AvBufferRef {
	return (*AvBufferRef)(frame.hw_frames_ctx)
}

// Custom: SetHwFramesCtx sets `AVFrame.hw_frames_ctx` value.
func (frame *AvFrame) SetHwFramesCtx(v *AvBufferRef) {
	frame.hw_frames_ctx = (*C.struct_AVBufferRef)(v)
}

// Custom: GetHwFramesCtxAddr gets `AVFrame.hw_frames_ctx` address.
func (frame *AvFrame) GetHwFramesCtxAddr() **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer(&frame.hw_frames_ctx))
}

// Custom: GetOpaqueRef gets `AVFrame.opaque_ref` value.
func (frame *AvFrame) GetOpaqueRef() *AvBufferRef {
	return (*AvBufferRef)(frame.opaque_ref)
}

// Custom: SetOpaqueRef sets `AVFrame.opaque_ref` value.
func (frame *AvFrame) SetOpaqueRef(v *AvBufferRef) {
	frame.opaque_ref = (*C.AVBufferRef)(v)
}

// Custom: GetOpaqueRefAddr gets `AVFrame.opaque_ref` address.
func (frame *AvFrame) GetOpaqueRefAddr() **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer(&frame.opaque_ref))
}

// Custom: GetCropTop gets `AVFrame.crop_top` value.
func (frame *AvFrame) GetCropTop() uint {
	return (uint)(frame.crop_top)
}

// Custom: SetCropTop sets `AVFrame.crop_top` value.
func (frame *AvFrame) SetCropTop(v uint) {
	frame.crop_top = (C.size_t)(v)
}

// Custom: GetCropTopAddr gets `AVFrame.crop_top` address.
func (frame *AvFrame) GetCropTopAddr() *uint {
	return (*uint)(unsafe.Pointer(&frame.crop_top))
}

// Custom: GetCropBottom gets `AVFrame.crop_bottom` value.
func (frame *AvFrame) GetCropBottom() uint {
	return (uint)(frame.crop_bottom)
}

// Custom: SetCropBottom sets `AVFrame.crop_bottom` value.
func (frame *AvFrame) SetCropBottom(v uint) {
	frame.crop_bottom = (C.size_t)(v)
}

// Custom: GetCropBottomAddr gets `AVFrame.crop_bottom` address.
func (frame *AvFrame) GetCropBottomAddr() *uint {
	return (*uint)(unsafe.Pointer(&frame.crop_bottom))
}

// Custom: GetCropLeft gets `AVFrame.crop_left` value.
func (frame *AvFrame) GetCropLeft() uint {
	return (uint)(frame.crop_left)
}

// Custom: SetCropLeft sets `AVFrame.crop_left` value.
func (frame *AvFrame) SetCropLeft(v uint) {
	frame.crop_left = (C.size_t)(v)
}

// Custom: GetCropLeftAddr gets `AVFrame.crop_left` address.
func (frame *AvFrame) GetCropLeftAddr() *uint {
	return (*uint)(unsafe.Pointer(&frame.crop_left))
}

// Custom: GetCropRight gets `AVFrame.crop_right` value.
func (frame *AvFrame) GetCropRight() uint {
	return (uint)(frame.crop_right)
}

// Custom: SetCropRight sets `AVFrame.crop_right` value.
func (frame *AvFrame) SetCropRight(v uint) {
	frame.crop_right = (C.size_t)(v)
}

// Custom: GetCropRightAddr gets `AVFrame.crop_right` address.
func (frame *AvFrame) GetCropRightAddr() *uint {
	return (*uint)(unsafe.Pointer(&frame.crop_right))
}

// Custom: GetPrivateRef gets `AVFrame.private_ref` value.
func (frame *AvFrame) GetPrivateRef() *AvBufferRef {
	return (*AvBufferRef)(frame.private_ref)
}

// Custom: SetPrivateRef sets `AVFrame.private_ref` value.
func (frame *AvFrame) SetPrivateRef(v *AvBufferRef) {
	frame.private_ref = (*C.struct_AVBufferRef)(v)
}

// Custom: GetPrivateRefAddr gets `AVFrame.private_ref` address.
func (frame *AvFrame) GetPrivateRefAddr() **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer(&frame.private_ref))
}

// Deprecated: No use
func AvFrameGetBestEffortTimestamp(frame *AvFrame) int64 {
	return (int64)(C.av_frame_get_best_effort_timestamp((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetBestEffortTimestamp(frame *AvFrame, val int64) {
	C.av_frame_set_best_effort_timestamp((*C.struct_AVFrame)(frame), (C.int64_t)(val))
}

// Deprecated: No use
func AvFrameGetPktDuration(frame *AvFrame) int64 {
	return (int64)(C.av_frame_get_pkt_duration((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetPktDuration(frame *AvFrame, val int64) {
	C.av_frame_set_pkt_duration((*C.struct_AVFrame)(frame), (C.int64_t)(val))
}

// Deprecated: No use
func AvFrameGetPktPos(frame *AvFrame) int64 {
	return (int64)(C.av_frame_get_pkt_pos((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetPktPos(frame *AvFrame, val int64) {
	C.av_frame_set_pkt_pos((*C.struct_AVFrame)(frame), (C.int64_t)(val))
}

// Deprecated: No use
func AvFrameGetChannelLayout(frame *AvFrame) int64 {
	return (int64)(C.av_frame_get_channel_layout((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetChannelLayout(frame *AvFrame, val int64) {
	C.av_frame_set_channel_layout((*C.struct_AVFrame)(frame), (C.int64_t)(val))
}

// Deprecated: No use
func AvFrameGetChannels(frame *AvFrame) int32 {
	return (int32)(C.av_frame_get_channels((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetChannels(frame *AvFrame, val int32) {
	C.av_frame_set_channels((*C.struct_AVFrame)(frame), (C.int)(val))
}

// Deprecated: No use
func AvFrameGetSampleRate(frame *AvFrame) int32 {
	return (int32)(C.av_frame_get_sample_rate((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetSampleRate(frame *AvFrame, val int32) {
	C.av_frame_set_sample_rate((*C.struct_AVFrame)(frame), (C.int)(val))
}

// Deprecated: No use
func AvFrameGetMetadata(frame *AvFrame) *AvDictionary {
	return (*AvDictionary)(C.av_frame_get_metadata((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetMetadata(frame *AvFrame, val *AvDictionary) {
	C.av_frame_set_metadata((*C.struct_AVFrame)(frame), (*C.struct_AVDictionary)(val))
}

// Deprecated: No use
func AvFrameGetDecodeErrorFlags(frame *AvFrame) int32 {
	return (int32)(C.av_frame_get_decode_error_flags((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetDecodeErrorFlags(frame *AvFrame, val int32) {
	C.av_frame_set_decode_error_flags((*C.struct_AVFrame)(frame), (C.int)(val))
}

// Deprecated: No use
func AvFrameGetPktSize(frame *AvFrame) int32 {
	return (int32)(C.av_frame_get_pkt_size((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetPktSize(frame *AvFrame, val int32) {
	C.av_frame_set_pkt_size((*C.struct_AVFrame)(frame), (C.int)(val))
}

// Deprecated: No use
func AvFrameGetQpTable(frame *AvFrame, stride, _type *int32) *int8 {
	return (*int8)(C.av_frame_get_qp_table((*C.struct_AVFrame)(frame),
		(*C.int)(stride), (*C.int)(_type)))
}

// Deprecated: No use
func AvFrameSetQpTable(frame *AvFrame, buf *AvBufferRef, stride, _type int32) int32 {
	return (int32)(C.av_frame_set_qp_table((*C.struct_AVFrame)(frame),
		(*C.struct_AVBufferRef)(buf), (C.int)(stride), (C.int)(_type)))
}

// Deprecated: No use
func AvFrameGetColorspace(frame *AvFrame) AvColorSpace {
	return (AvColorSpace)(C.av_frame_get_colorspace((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetColorspace(frame *AvFrame, val AvColorSpace) {
	C.av_frame_set_colorspace((*C.struct_AVFrame)(frame), (C.enum_AVColorSpace)(val))
}

// Deprecated: No use
func AvFrameGetColorRange(frame *AvFrame) AvColorRange {
	return (AvColorRange)(C.av_frame_get_color_range((*C.struct_AVFrame)(frame)))
}

// Deprecated: No use
func AvFrameSetColorRange(frame *AvFrame, val AvColorRange) {
	C.av_frame_set_color_range((*C.struct_AVFrame)(frame), (C.enum_AVColorRange)(val))
}

// AvGetColorspaceName gets the name of a colorspace.
func AvGetColorspaceName(val AvColorSpace) string {
	return C.GoString(C.av_get_colorspace_name((C.enum_AVColorSpace)(val)))
}

// AvFrameAlloc allocates an AVFrame and set its fields to default values.
// The resulting struct must be freed using AvFrameFree().
func AvFrameAlloc() *AvFrame {
	return (*AvFrame)(C.av_frame_alloc())
}

// AvFrameFree frees the frame and any dynamically allocated objects in it,
// e.g. extended_data. If the frame is reference counted, it will be
// unreferenced first.
func AvFrameFree(frame **AvFrame) {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(frame)))
}

// AvFrameRef sets up a new reference to the data described by the source frame.
func AvFrameRef(dst, src *AvFrame) int32 {
	return (int32)(C.av_frame_ref((*C.struct_AVFrame)(dst), (*C.struct_AVFrame)(src)))
}

// AvFrameClone creates a new frame that references the same data as src.
func AvFrameClone(frame *AvFrame) *AvFrame {
	return (*AvFrame)(C.av_frame_clone((*C.struct_AVFrame)(frame)))
}

// AvFrameUnref unreferences all the buffers referenced by frame and reset the frame fields.
func AvFrameUnref(frame *AvFrame) {
	C.av_frame_unref((*C.struct_AVFrame)(frame))
}

// AvFrameMoveRef moves everything contained in src to dst and reset src.
func AvFrameMoveRef(dst, src *AvFrame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(dst), (*C.struct_AVFrame)(src))
}

// AvFrameGetBuffer allocates new buffer(s) for audio or video data.
func AvFrameGetBuffer(frame *AvFrame, align int32) int32 {
	return (int32)(C.av_frame_get_buffer((*C.struct_AVFrame)(frame), (C.int)(align)))
}

// AvFrameIsWritable checks if the frame data is writable.
func AvFrameIsWritable(frame *AvFrame) int32 {
	return (int32)(C.av_frame_is_writable((*C.struct_AVFrame)(frame)))
}

// AvFrameMakeWritable ensures that the frame data is writable, avoiding data copy if possible.
func AvFrameMakeWritable(frame *AvFrame) int32 {
	return (int32)(C.av_frame_make_writable((*C.struct_AVFrame)(frame)))
}

// AvFrameCopy copies the frame data from src to dst.
func AvFrameCopy(dst, src *AvFrame) int32 {
	return (int32)(C.av_frame_copy((*C.struct_AVFrame)(dst), (*C.struct_AVFrame)(src)))
}

// AvFrameCopyProps copies only "metadata" fields from src to dst.
func AvFrameCopyProps(dst, src *AvFrame) int32 {
	return (int32)(C.av_frame_copy_props((*C.struct_AVFrame)(dst), (*C.struct_AVFrame)(src)))
}

// AvFrameGetPlaneBuffer gets the buffer reference a given data plane is stored in.
func AvFrameGetPlaneBuffer(frame *AvFrame, plane int32) *AvBufferRef {
	return (*AvBufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(frame), (C.int)(plane)))
}

// AvFrameNewSideData adds a new side data to a frame.
func AvFrameNewSideData(frame *AvFrame, _type AvFrameSideDataType, size int32) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(frame),
		(C.enum_AVFrameSideDataType)(_type), (C.int)(size)))
}

// AvFrameNewSideDataFromBuf adds a new side data to a frame from an existing AVBufferRef.
func AvFrameNewSideDataFromBuf(frame *AvFrame, _type AvFrameSideDataType,
	buf *AvBufferRef) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_new_side_data_from_buf((*C.struct_AVFrame)(frame),
		(C.enum_AVFrameSideDataType)(_type), (*C.struct_AVBufferRef)(buf)))
}

// AvFrameGetSideData gets a pointer to the side data of a given type on success, NULL if there
// is no side data with such type in this frame.
func AvFrameGetSideData(frame *AvFrame, _type AvFrameSideDataType) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(frame),
		(C.enum_AVFrameSideDataType)(_type)))
}

// AvFrameRemoveSideData remove and free all side data instances of the given type.
func AvFrameRemoveSideData(frame *AvFrame, _type AvFrameSideDataType) {
	C.av_frame_remove_side_data((*C.struct_AVFrame)(frame), (C.enum_AVFrameSideDataType)(_type))
}

// Flags for frame cropping.
const (
	AV_FRAME_CROP_UNALIGNED = int32(C.AV_FRAME_CROP_UNALIGNED)
)

// AvFrameApplyCropping crops the given video AVFrame according to its crop_left/crop_top/crop_right/
// crop_bottom fields.
func AvFrameApplyCropping(frame *AvFrame, flags int32) int32 {
	return (int32)(C.av_frame_apply_cropping((*C.struct_AVFrame)(frame), (C.int)(flags)))
}

// AvFrameSideDataName returns a string identifying the side data type
func AvFrameSideDataName(_type AvFrameSideDataType) string {
	return C.GoString(C.av_frame_side_data_name((C.enum_AVFrameSideDataType)(_type)))
}
