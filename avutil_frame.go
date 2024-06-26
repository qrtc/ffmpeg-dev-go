// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/frame.h>
*/
import "C"
import "unsafe"

// AVFrameSideDataType
type AVFrameSideDataType = C.enum_AVFrameSideDataType

const (
	AV_FRAME_DATA_PANSCAN                    = AVFrameSideDataType(C.AV_FRAME_DATA_PANSCAN)
	AV_FRAME_DATA_A53_CC                     = AVFrameSideDataType(C.AV_FRAME_DATA_A53_CC)
	AV_FRAME_DATA_STEREO3D                   = AVFrameSideDataType(C.AV_FRAME_DATA_STEREO3D)
	AV_FRAME_DATA_MATRIXENCODING             = AVFrameSideDataType(C.AV_FRAME_DATA_MATRIXENCODING)
	AV_FRAME_DATA_DOWNMIX_INFO               = AVFrameSideDataType(C.AV_FRAME_DATA_DOWNMIX_INFO)
	AV_FRAME_DATA_REPLAYGAIN                 = AVFrameSideDataType(C.AV_FRAME_DATA_REPLAYGAIN)
	AV_FRAME_DATA_DISPLAYMATRIX              = AVFrameSideDataType(C.AV_FRAME_DATA_DISPLAYMATRIX)
	AV_FRAME_DATA_AFD                        = AVFrameSideDataType(C.AV_FRAME_DATA_AFD)
	AV_FRAME_DATA_MOTION_VECTORS             = AVFrameSideDataType(C.AV_FRAME_DATA_MOTION_VECTORS)
	AV_FRAME_DATA_SKIP_SAMPLES               = AVFrameSideDataType(C.AV_FRAME_DATA_SKIP_SAMPLES)
	AV_FRAME_DATA_AUDIO_SERVICE_TYPE         = AVFrameSideDataType(C.AV_FRAME_DATA_AUDIO_SERVICE_TYPE)
	AV_FRAME_DATA_MASTERING_DISPLAY_METADATA = AVFrameSideDataType(C.AV_FRAME_DATA_MASTERING_DISPLAY_METADATA)
	AV_FRAME_DATA_GOP_TIMECODE               = AVFrameSideDataType(C.AV_FRAME_DATA_GOP_TIMECODE)
	AV_FRAME_DATA_SPHERICAL                  = AVFrameSideDataType(C.AV_FRAME_DATA_SPHERICAL)
	AV_FRAME_DATA_CONTENT_LIGHT_LEVEL        = AVFrameSideDataType(C.AV_FRAME_DATA_CONTENT_LIGHT_LEVEL)
	AV_FRAME_DATA_ICC_PROFILE                = AVFrameSideDataType(C.AV_FRAME_DATA_ICC_PROFILE)

	AV_FRAME_DATA_S12M_TIMECODE       = AVFrameSideDataType(C.AV_FRAME_DATA_S12M_TIMECODE)
	AV_FRAME_DATA_DYNAMIC_HDR_PLUS    = AVFrameSideDataType(C.AV_FRAME_DATA_DYNAMIC_HDR_PLUS)
	AV_FRAME_DATA_REGIONS_OF_INTEREST = AVFrameSideDataType(C.AV_FRAME_DATA_REGIONS_OF_INTEREST)
	AV_FRAME_DATA_VIDEO_ENC_PARAMS    = AVFrameSideDataType(C.AV_FRAME_DATA_VIDEO_ENC_PARAMS)
	AV_FRAME_DATA_SEI_UNREGISTERED    = AVFrameSideDataType(C.AV_FRAME_DATA_SEI_UNREGISTERED)
	AV_FRAME_DATA_FILM_GRAIN_PARAMS   = AVFrameSideDataType(C.AV_FRAME_DATA_FILM_GRAIN_PARAMS)
	AV_FRAME_DATA_DETECTION_BBOXES    = AVFrameSideDataType(C.AV_FRAME_DATA_DETECTION_BBOXES)
	AV_FRAME_DATA_DOVI_RPU_BUFFER     = AVFrameSideDataType(C.AV_FRAME_DATA_DOVI_RPU_BUFFER)
	AV_FRAME_DATA_DOVI_METADATA       = AVFrameSideDataType(C.AV_FRAME_DATA_DOVI_METADATA)

	AV_FRAME_DATA_DYNAMIC_HDR_VIVID           = AVFrameSideDataType(C.AV_FRAME_DATA_DYNAMIC_HDR_VIVID)
	AV_FRAME_DATA_AMBIENT_VIEWING_ENVIRONMENT = AVFrameSideDataType(C.AV_FRAME_DATA_AMBIENT_VIEWING_ENVIRONMENT)
)

// AVActiveFormatDescription
type AVActiveFormatDescription = C.enum_AVActiveFormatDescription

const (
	AV_AFD_SAME         = AVActiveFormatDescription(C.AV_AFD_SAME)
	AV_AFD_4_3          = AVActiveFormatDescription(C.AV_AFD_4_3)
	AV_AFD_16_9         = AVActiveFormatDescription(C.AV_AFD_16_9)
	AV_AFD_14_9         = AVActiveFormatDescription(C.AV_AFD_14_9)
	AV_AFD_4_3_SP_14_9  = AVActiveFormatDescription(C.AV_AFD_4_3_SP_14_9)
	AV_AFD_16_9_SP_14_9 = AVActiveFormatDescription(C.AV_AFD_16_9_SP_14_9)
	AV_AFD_SP_4_3       = AVActiveFormatDescription(C.AV_AFD_SP_4_3)
)

// Structure to hold side data for an AVFrame.
type AVFrameSideData C.struct_AVFrameSideData

// GetType gets `AVFrameSideData.type` value.
func (sd *AVFrameSideData) GetType() AVFrameSideDataType {
	return (AVFrameSideDataType)(sd._type)
}

// SetType sets `AVFrameSideData.type` value.
func (sd *AVFrameSideData) SetType(v AVFrameSideDataType) {
	sd._type = (C.enum_AVFrameSideDataType)(v)
}

// GetTypeAddr gets `AVFrameSideData.type` address.
func (sd *AVFrameSideData) GetTypeAddr() *AVFrameSideDataType {
	return (*AVFrameSideDataType)(&sd._type)
}

// GetData gets `AVFrameSideData.data` value.
func (sd *AVFrameSideData) GetData() *uint8 {
	return (*uint8)(sd.data)
}

// SetData sets `AVFrameSideData.data` value.
func (sd *AVFrameSideData) SetData(v *uint8) {
	sd.data = (*C.uint8_t)(v)
}

// GetDataAddr gets `AVFrameSideData.data` address.
func (sd *AVFrameSideData) GetDataAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&sd.data))
}

// GetSize gets `AVFrameSideData.size` value.
func (sd *AVFrameSideData) GetSize() uintptr {
	return (uintptr)(sd.size)
}

// SetSize sets `AVFrameSideData.size` value.
func (sd *AVFrameSideData) SetSize(v uintptr) {
	sd.size = (C.size_t)(v)
}

// GetSizeAddr gets `AVFrameSideData.size` address.
func (sd *AVFrameSideData) GetSizeAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&sd.size))
}

// GetMetadata gets `AVFrameSideData.metadata` value.
func (sd *AVFrameSideData) GetMetadata() *AVDictionary {
	return (*AVDictionary)(sd.metadata)
}

// SetMetadata sets `AVFrameSideData.metadata` value.
func (sd *AVFrameSideData) SetMetadata(v *AVDictionary) {
	sd.metadata = (*C.struct_AVDictionary)(v)
}

// GetMetadataAddr gets `AVFrameSideData.metadata` address.
func (sd *AVFrameSideData) GetMetadataAddr() **AVDictionary {
	return (**AVDictionary)(unsafe.Pointer(&sd.metadata))
}

// GetBuf gets `AVFrameSideData.buf` value.
func (sd *AVFrameSideData) GetBuf() *AVBufferRef {
	return (*AVBufferRef)(sd.buf)
}

// SetBuf sets `AVFrameSideData.buf` value.
func (sd *AVFrameSideData) SetBuf(v *AVBufferRef) {
	sd.buf = (*C.struct_AVBufferRef)(v)
}

// GetBufAddr gets `AVFrameSideData.buf` address.
func (sd *AVFrameSideData) GetBufAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&sd.buf))
}

// Structure describing a single Region Of Interest.
type AVRegionOfInterest C.struct_AVRegionOfInterest

// GetSelfSize gets `AVRegionOfInterest.self_size` value.
func (roi *AVRegionOfInterest) GetSelfSize() uint32 {
	return (uint32)(roi.self_size)
}

// SetSelfSize sets `AVRegionOfInterest.self_size` value.
func (roi *AVRegionOfInterest) SetSelfSize(v uint32) {
	roi.self_size = (C.uint32_t)(v)
}

// GetSelfSizeAddr gets `AVRegionOfInterest.self_size` address.
func (roi *AVRegionOfInterest) GetSelfSizeAddr() *uint32 {
	return (*uint32)(&roi.self_size)
}

// GetTop gets `AVRegionOfInterest.top` value.
func (roi *AVRegionOfInterest) GetTop() int32 {
	return (int32)(roi.top)
}

// SetTop sets `AVRegionOfInterest.top` value.
func (roi *AVRegionOfInterest) SetTop(v int32) {
	roi.top = (C.int)(v)
}

// GetTopAddr gets `AVRegionOfInterest.top` address.
func (roi *AVRegionOfInterest) GetTopAddr() *int32 {
	return (*int32)(&roi.top)
}

// GetBottom gets `AVRegionOfInterest.bottom` value.
func (roi *AVRegionOfInterest) GetBottom() int32 {
	return (int32)(roi.bottom)
}

// SetBottom sets `AVRegionOfInterest.bottom` value.
func (roi *AVRegionOfInterest) SetBottom(v int32) {
	roi.bottom = (C.int)(v)
}

// GetBottomAddr gets `AVRegionOfInterest.bottom` address.
func (roi *AVRegionOfInterest) GetBottomAddr() *int32 {
	return (*int32)(&roi.bottom)
}

// GetLeft gets `AVRegionOfInterest.left` value.
func (roi *AVRegionOfInterest) GetLeft() int32 {
	return (int32)(roi.left)
}

// SetLeft sets `AVRegionOfInterest.left` value.
func (roi *AVRegionOfInterest) SetLeft(v int32) {
	roi.left = (C.int)(v)
}

// GetLeftAddr gets `AVRegionOfInterest.left` address.
func (roi *AVRegionOfInterest) GetLeftAddr() *int32 {
	return (*int32)(&roi.left)
}

// GetRight gets `AVRegionOfInterest.right` value.
func (roi *AVRegionOfInterest) GetRight() int32 {
	return (int32)(roi.right)
}

// SetRight sets `AVRegionOfInterest.right` value.
func (roi *AVRegionOfInterest) SetRight(v int32) {
	roi.right = (C.int)(v)
}

// GetRightAddr gets `AVRegionOfInterest.right` address.
func (roi *AVRegionOfInterest) GetRightAddr() *int32 {
	return (*int32)(&roi.right)
}

// GetQoffset gets `AVRegionOfInterest.qoffset` value.
func (roi *AVRegionOfInterest) GetQoffset() AVRational {
	return (AVRational)(roi.qoffset)
}

// SetQoffset sets `AVRegionOfInterest.qoffset` value.
func (roi *AVRegionOfInterest) SetQoffset(v AVRational) {
	roi.qoffset = (C.struct_AVRational)(v)
}

// GetQoffsetAddr gets `AVRegionOfInterest.qoffset` address.
func (roi *AVRegionOfInterest) GetQoffsetAddr() *AVRational {
	return (*AVRational)(&roi.qoffset)
}

// AVFrame
type AVFrame C.struct_AVFrame

const (
	AV_NUM_DATA_POINTERS = C.AV_NUM_DATA_POINTERS
)

// GetData gets `AVFrame.data` value.
func (frame *AVFrame) GetData() []*uint8 {
	return unsafe.Slice((**uint8)(unsafe.Pointer(&frame.data[0])), AV_NUM_DATA_POINTERS)
}

// SetData sets `AVFrame.data` value.
func (frame *AVFrame) SetData(v []*uint8) {
	for i := 0; i < FFMIN(len(v), AV_NUM_DATA_POINTERS); i++ {
		frame.data[i] = (*C.uint8_t)(v[i])
	}
}

// GetDataAddr gets `AVFrame.data` address.
func (frame *AVFrame) GetDataAddr() ***uint8 {
	return (***uint8)(unsafe.Pointer(&frame.data))
}

// GetLinesize gets `AVFrame.linesize` value.
func (frame *AVFrame) GetLinesize() []int32 {
	return unsafe.Slice((*int32)(&frame.linesize[0]), AV_NUM_DATA_POINTERS)
}

// SetLinesize sets `AVFrame.linesize` value.
func (frame *AVFrame) SetLinesize(v []int32) {
	for i := 0; i < FFMIN(len(v), AV_NUM_DATA_POINTERS); i++ {
		frame.linesize[i] = (C.int)(v[i])
	}
}

// GetLinesizeAddr gets `AVFrame.linesize` address.
func (frame *AVFrame) GetLinesizeAddr() **int32 {
	return (**int32)(unsafe.Pointer(&frame.linesize))
}

// GetExtendedData gets `AVFrame.extended_data` value.
func (frame *AVFrame) GetExtendedData() **uint8 {
	return (**uint8)(unsafe.Pointer(frame.extended_data))
}

// SetExtendedData sets `AVFrame.extended_data` value.
func (frame *AVFrame) SetExtendedData(v **uint8) {
	frame.extended_data = (**C.uint8_t)(unsafe.Pointer(v))
}

// GetExtendedDataAddr gets `AVFrame.extended_data` address.
func (frame *AVFrame) GetExtendedDataAddr() ***uint8 {
	return (***uint8)(unsafe.Pointer(&frame.extended_data))
}

// GetWidth gets `AVFrame.width` value.
func (frame *AVFrame) GetWidth() int32 {
	return (int32)(frame.width)
}

// SetWidth sets `AVFrame.width` value.
func (frame *AVFrame) SetWidth(v int32) {
	frame.width = (C.int)(v)
}

// GetWidthAddr gets `AVFrame.width` address.
func (frame *AVFrame) GetWidthAddr() *int32 {
	return (*int32)(&frame.width)
}

// GetHeight gets `AVFrame.height` value.
func (frame *AVFrame) GetHeight() int32 {
	return (int32)(frame.height)
}

// SetHeight sets `AVFrame.height` value.
func (frame *AVFrame) SetHeight(v int32) {
	frame.height = (C.int)(v)
}

// GetHeightAddr gets `AVFrame.height` address.
func (frame *AVFrame) GetHeightAddr() *int32 {
	return (*int32)(&frame.height)
}

// GetNbSamples gets `AVFrame.nb_samples` value.
func (frame *AVFrame) GetNbSamples() int32 {
	return (int32)(frame.nb_samples)
}

// SetNbSamples sets `AVFrame.nb_samples` value.
func (frame *AVFrame) SetNbSamples(v int32) {
	frame.nb_samples = (C.int)(v)
}

// GetNbSamplesAddr gets `AVFrame.nb_samples` address.
func (frame *AVFrame) GetNbSamplesAddr() *int32 {
	return (*int32)(&frame.nb_samples)
}

// GetFormat gets `AVFrame.format` value.
func (frame *AVFrame) GetFormat() int32 {
	return (int32)(frame.format)
}

// SetFormat sets `AVFrame.format` value.
func (frame *AVFrame) SetFormat(v int32) {
	frame.format = (C.int)(v)
}

// GetFormatAddr gets `AVFrame.format` address.
func (frame *AVFrame) GetFormatAddr() *int32 {
	return (*int32)(&frame.format)
}

// Deprecated: Use AV_FRAME_FLAG_KEY instead.
//
// GetKeyFrame gets `AVFrame.key_frame` value.
func (frame *AVFrame) GetKeyFrame() int32 {
	return (int32)(frame.key_frame)
}

// Deprecated: Use AV_FRAME_FLAG_KEY instead.
//
// SetKeyFrame sets `AVFrame.key_frame` value.
func (frame *AVFrame) SetKeyFrame(v int32) {
	frame.key_frame = (C.int)(v)
}

// Deprecated: Use AV_FRAME_FLAG_KEY instead.
//
// GetKeyFrameAddr gets `AVFrame.key_frame` address.
func (frame *AVFrame) GetKeyFrameAddr() *int32 {
	return (*int32)(&frame.key_frame)
}

// GetPictType gets `AVFrame.picttype` value.
func (frame *AVFrame) GetPictType() AVPictureType {
	return (AVPictureType)(frame.pict_type)
}

// SetPictType sets `AVFrame.picttype` value.
func (frame *AVFrame) SetPictType(v AVPictureType) {
	frame.pict_type = (C.enum_AVPictureType)(v)
}

// GetPictTypeAddr gets `AVFrame.picttype` address.
func (frame *AVFrame) GetPictTypeAddr() *AVPictureType {
	return (*AVPictureType)(unsafe.Pointer(&frame.pict_type))
}

// GetSampleAspectRatio gets `AVFrame.sample_aspect_ratio` value.
func (frame *AVFrame) GetSampleAspectRatio() AVRational {
	return (AVRational)(frame.sample_aspect_ratio)
}

// SetSampleAspectRatio sets `AVFrame.sample_aspect_ratio` value.
func (frame *AVFrame) SetSampleAspectRatio(v AVRational) {
	frame.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// GetSampleAspectRatioAddr gets `AVFrame.sample_aspect_ratio` address.
func (frame *AVFrame) GetSampleAspectRatioAddr() *AVRational {
	return (*AVRational)(&frame.sample_aspect_ratio)
}

// GetPts gets `AVFrame.pts` value.
func (frame *AVFrame) GetPts() int64 {
	return (int64)(frame.pts)
}

// SetPts sets `AVFrame.pts` value.
func (frame *AVFrame) SetPts(v int64) {
	frame.pts = (C.int64_t)(v)
}

// GetPtsAddr gets `AVFrame.pts` address.
func (frame *AVFrame) GetPtsAddr() *int64 {
	return (*int64)(&frame.pts)
}

// GetPktDts gets `AVFrame.pkt_dts` value.
func (frame *AVFrame) GetPktDts() int64 {
	return (int64)(frame.pkt_dts)
}

// SetPktDts sets `AVFrame.pkt_dts` value.
func (frame *AVFrame) SetPktDts(v int64) {
	frame.pkt_dts = (C.int64_t)(v)
}

// GetPktDtsAddr gets `AVFrame.pkt_dts` address.
func (frame *AVFrame) GetPktDtsAddr() *int64 {
	return (*int64)(&frame.pkt_dts)
}

// GetTimeBase gets `AVFrame.time_base` value.
func (frame *AVFrame) GetTimeBase() AVRational {
	return (AVRational)(frame.time_base)
}

// SetTimeBase sets `AVFrame.time_base` value.
func (frame *AVFrame) SetTimeBase(v AVRational) {
	frame.time_base = (C.struct_AVRational)(v)
}

// GetTimeBaseAddr gets `AVFrame.time_base` address.
func (frame *AVFrame) GetTimeBaseAddr() *AVRational {
	return (*AVRational)(&frame.time_base)
}

// GetQuality gets `AVFrame.quality` value.
func (frame *AVFrame) GetQuality() int32 {
	return (int32)(frame.quality)
}

// SetQuality sets `AVFrame.quality` value.
func (frame *AVFrame) SetQuality(v int32) {
	frame.quality = (C.int)(v)
}

// GetQualityAddr gets `AVFrame.quality` address.
func (frame *AVFrame) GetQualityAddr() *int32 {
	return (*int32)(&frame.quality)
}

// GetRepeatPict gets `AVFrame.repeat_pict` value.
func (frame *AVFrame) GetRepeatPict() int32 {
	return (int32)(frame.repeat_pict)
}

// SetRepeatPict sets `AVFrame.repeat_pict` value.
func (frame *AVFrame) SetRepeatPict(v int32) {
	frame.repeat_pict = (C.int)(v)
}

// GetRepeatPictAddr gets `AVFrame.repeat_pict` address.
func (frame *AVFrame) GetRepeatPictAddr() *int32 {
	return (*int32)(&frame.repeat_pict)
}

// Deprecated: Use AV_FRAME_FLAG_INTERLACED instead.
//
// GetInterlacedFrame gets `AVFrame.interlaced_frame` value.
func (frame *AVFrame) GetInterlacedFrame() int32 {
	return (int32)(frame.interlaced_frame)
}

// Deprecated: Use AV_FRAME_FLAG_INTERLACED instead.
//
// SetInterlacedFrame sets `AVFrame.interlaced_frame` value.
func (frame *AVFrame) SetInterlacedFrame(v int32) {
	frame.interlaced_frame = (C.int)(v)
}

// Deprecated: Use AV_FRAME_FLAG_INTERLACED instead.
//
// GetInterlacedFrameAddr gets `AVFrame.interlaced_frame` address.
func (frame *AVFrame) GetInterlacedFrameAddr() *int32 {
	return (*int32)(&frame.interlaced_frame)
}

// Deprecated: Use AV_FRAME_FLAG_TOP_FIELD_FIRST instead.
//
// GetTopFieldFirst gets `AVFrame.top_field_first` value.
func (frame *AVFrame) GetTopFieldFirst() int32 {
	return (int32)(frame.top_field_first)
}

// Deprecated: Use AV_FRAME_FLAG_TOP_FIELD_FIRST instead.
//
// SetTopFieldFirst sets `AVFrame.top_field_first` value.
func (frame *AVFrame) SetTopFieldFirst(v int32) {
	frame.top_field_first = (C.int)(v)
}

// Deprecated: Use AV_FRAME_FLAG_TOP_FIELD_FIRST instead.
//
// GetTopFieldFirstAddr gets `AVFrame.top_field_first` address.
func (frame *AVFrame) GetTopFieldFirstAddr() *int32 {
	return (*int32)(&frame.top_field_first)
}

// Deprecated: no use.
//
// GetPaletteHasChanged gets `AVFrame.palette_has_changed` value.
func (frame *AVFrame) GetPaletteHasChanged() int32 {
	return (int32)(frame.palette_has_changed)
}

// Deprecated: no use.
//
// SetPaletteHasChanged sets `AVFrame.palette_has_changed` value.
func (frame *AVFrame) SetPaletteHasChanged(v int32) {
	frame.palette_has_changed = (C.int)(v)
}

// Deprecated: no use.
//
// GetPaletteHasChangedAddr gets `AVFrame.palette_has_changed` address.
func (frame *AVFrame) GetPaletteHasChangedAddr() *int32 {
	return (*int32)(&frame.palette_has_changed)
}

// GetSampleRate gets `AVFrame.sample_rate` value.
func (frame *AVFrame) GetSampleRate() int32 {
	return (int32)(frame.sample_rate)
}

// SetSampleRate sets `AVFrame.sample_rate` value.
func (frame *AVFrame) SetSampleRate(v int32) {
	frame.sample_rate = (C.int)(v)
}

// GetSampleRateAddr gets `AVFrame.sample_rate` address.
func (frame *AVFrame) GetSampleRateAddr() *int32 {
	return (*int32)(&frame.sample_rate)
}

// GetBuf gets `AVFrame.buf` value.
func (frame *AVFrame) GetBuf() []*AVBufferRef {
	return unsafe.Slice((**AVBufferRef)(unsafe.Pointer(&frame.buf[0])), AV_NUM_DATA_POINTERS)
}

// SetBuf sets `AVFrame.buf` value.
func (frame *AVFrame) SetBuf(v []*AVBufferRef) {
	for i := 0; i < FFMIN(len(v), AV_NUM_DATA_POINTERS); i++ {
		frame.buf[i] = (*C.struct_AVBufferRef)(v[i])
	}
}

// GetBufAddr gets `AVFrame.buf` address.
func (frame *AVFrame) GetBufAddr() ***AVBufferRef {
	return (***AVBufferRef)(unsafe.Pointer(&frame.buf))
}

// GetExtendedBuf gets `AVFrame.extended_buf` value.
func (frame *AVFrame) GetExtendedBuf() []*AVBufferRef {
	return unsafe.Slice((**AVBufferRef)(unsafe.Pointer(frame.extended_buf)),
		frame.nb_extended_buf)
}

// SetExtendedBuf sets `AVFrame.extended_buf` value.
func (frame *AVFrame) SetExtendedBuf(v **AVBufferRef) {
	frame.extended_buf = (**C.struct_AVBufferRef)(unsafe.Pointer(v))
}

// GetExtendedBufAddr gets `AVFrame.extended_buf` address.
func (frame *AVFrame) GetExtendedBufAddr() ***AVBufferRef {
	return (***AVBufferRef)(unsafe.Pointer(&frame.extended_buf))
}

// GetNbExtendedBuf gets `AVFrame.nb_extended_buf` value.
func (frame *AVFrame) GetNbExtendedBuf() int32 {
	return (int32)(frame.nb_extended_buf)
}

// SetNbExtendedBuf sets `AVFrame.nb_extended_buf` value.
func (frame *AVFrame) SetNbExtendedBuf(v int32) {
	frame.nb_extended_buf = (C.int)(v)
}

// GetNbExtendedBufAddr gets `AVFrame.nb_extended_buf` address.
func (frame *AVFrame) GetNbExtendedBufAddr() *int32 {
	return (*int32)(&frame.nb_extended_buf)
}

// GetSideData gets `AVFrame.side_data` value.
func (frame *AVFrame) GetSideData() []*AVFrameSideData {
	return unsafe.Slice((**AVFrameSideData)(unsafe.Pointer(frame.side_data)), frame.nb_side_data)
}

// SetSideData sets `AVFrame.side_data` value.
func (frame *AVFrame) SetSideData(v **AVFrameSideData) {
	frame.side_data = (**C.struct_AVFrameSideData)(unsafe.Pointer(v))
}

// GetSideDataAddr gets `AVFrame.side_data` address.
func (frame *AVFrame) GetSideDataAddr() ***AVFrameSideData {
	return (***AVFrameSideData)(unsafe.Pointer(&frame.side_data))
}

// GetNbSideData gets `AVFrame.nb_side_data` value.
func (frame *AVFrame) GetNbSideData() int32 {
	return (int32)(frame.nb_side_data)
}

// SetNbSideData sets `AVFrame.nb_side_data` value.
func (frame *AVFrame) SetNbSideData(v int32) {
	frame.nb_side_data = (C.int)(v)
}

// GetNbSideDataAddr gets `AVFrame.nb_side_data` address.
func (frame *AVFrame) GetNbSideDataAddr() *int32 {
	return (*int32)(&frame.nb_side_data)
}

const (
	AV_FRAME_FLAG_CORRUPT         = int32(C.AV_FRAME_FLAG_CORRUPT)
	AV_FRAME_FLAG_KEY             = int32(C.AV_FRAME_FLAG_KEY)
	AV_FRAME_FLAG_DISCARD         = int32(C.AV_FRAME_FLAG_DISCARD)
	AV_FRAME_FLAG_INTERLACED      = int32(C.AV_FRAME_FLAG_INTERLACED)
	AV_FRAME_FLAG_TOP_FIELD_FIRST = int32(C.AV_FRAME_FLAG_TOP_FIELD_FIRST)
)

// GetFlags gets `AVFrame.flags` value.
func (frame *AVFrame) GetFlags() int32 {
	return (int32)(frame.flags)
}

// SetFlags sets `AVFrame.flags` value.
func (frame *AVFrame) SetFlags(v int32) {
	frame.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVFrame.flags` address.
func (frame *AVFrame) GetFlagsAddr() *int32 {
	return (*int32)(&frame.flags)
}

// GetColorRange gets `AVFrame.colorrange` value.
func (frame *AVFrame) GetColorRange() AVColorRange {
	return (AVColorRange)(frame.color_range)
}

// SetColorRange sets `AVFrame.colorrange` value.
func (frame *AVFrame) SetColorRange(v AVColorRange) {
	frame.color_range = (C.enum_AVColorRange)(v)
}

// GetColorRangeAddr gets `AVFrame.colorrange` address.
func (frame *AVFrame) GetColorRangeAddr() *AVColorRange {
	return (*AVColorRange)(unsafe.Pointer(&frame.color_range))
}

// GetColorPrimaries gets `AVFrame.color_primaries` value.
func (frame *AVFrame) GetColorPrimaries() AVColorPrimaries {
	return (AVColorPrimaries)(frame.color_primaries)
}

// SetColorPrimaries sets `AVFrame.color_primaries` value.
func (frame *AVFrame) SetColorPrimaries(v AVColorPrimaries) {
	frame.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// GetColorPrimariesAddr gets `AVFrame.color_primaries` address.
func (frame *AVFrame) GetColorPrimariesAddr() *AVColorPrimaries {
	return (*AVColorPrimaries)(unsafe.Pointer(&frame.color_primaries))
}

// GetColorTrc gets `AVFrame.color_trc` value.
func (frame *AVFrame) GetColorTrc() AVColorTransferCharacteristic {
	return (AVColorTransferCharacteristic)(frame.color_trc)
}

// SetColorTrc sets `AVFrame.color_trc` value.
func (frame *AVFrame) SetColorTrc(v AVColorTransferCharacteristic) {
	frame.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// GetColorTrcAddr gets `AVFrame.color_trc` address.
func (frame *AVFrame) GetColorTrcAddr() *AVColorTransferCharacteristic {
	return (*AVColorTransferCharacteristic)(unsafe.Pointer(&frame.color_trc))
}

// GetColorspace gets `AVFrame.colorspace` value.
func (frame *AVFrame) GetColorspace() AVColorSpace {
	return (AVColorSpace)(frame.colorspace)
}

// SetColorspace sets `AVFrame.colorspace` value.
func (frame *AVFrame) SetColorspace(v AVColorSpace) {
	frame.colorspace = (C.enum_AVColorSpace)(v)
}

// GetColorspaceAddr gets `AVFrame.colorspace` address.
func (frame *AVFrame) GetColorspaceAddr() *AVColorSpace {
	return (*AVColorSpace)(unsafe.Pointer(&frame.colorspace))
}

// GetChromaLocation gets `AVFrame.chroma_location` value.
func (frame *AVFrame) GetChromaLocation() AVChromaLocation {
	return (AVChromaLocation)(frame.chroma_location)
}

// SetChromaLocation sets `AVFrame.chroma_location` value.
func (frame *AVFrame) SetChromaLocation(v AVChromaLocation) {
	frame.chroma_location = (C.enum_AVChromaLocation)(v)
}

// GetChromaLocationAddr gets `AVFrame.chroma_location` address.
func (frame *AVFrame) GetChromaLocationAddr() *AVChromaLocation {
	return (*AVChromaLocation)(unsafe.Pointer(&frame.chroma_location))
}

// GetBestEffortTimestamp gets `AVFrame.best_effort_timestamp` value.
func (frame *AVFrame) GetBestEffortTimestamp() int64 {
	return (int64)(frame.best_effort_timestamp)
}

// SetBestEffortTimestamp sets `AVFrame.best_effort_timestamp` value.
func (frame *AVFrame) SetBestEffortTimestamp(v int64) {
	frame.best_effort_timestamp = (C.int64_t)(v)
}

// GetBestEffortTimestampAddr gets `AVFrame.best_effort_timestamp` address.
func (frame *AVFrame) GetBestEffortTimestampAddr() *int64 {
	return (*int64)(&frame.best_effort_timestamp)
}

// GetPktPos gets `AVFrame.pkt_pos` value.
func (frame *AVFrame) GetPktPos() int64 {
	return (int64)(frame.pkt_pos)
}

// SetPktPos sets `AVFrame.pkt_pos` value.
func (frame *AVFrame) SetPktPos(v int64) {
	frame.pkt_pos = (C.int64_t)(v)
}

// GetPktPosAddr gets `AVFrame.pkt_pos` address.
func (frame *AVFrame) GetPktPosAddr() *int64 {
	return (*int64)(&frame.pkt_pos)
}

// GetMetadata gets `AVFrame.metadata` value.
func (frame *AVFrame) GetMetadata() *AVDictionary {
	return (*AVDictionary)(frame.metadata)
}

// SetMetadata sets `AVFrame.metadata` value.
func (frame *AVFrame) SetMetadata(v *AVDictionary) {
	frame.metadata = (*C.struct_AVDictionary)(v)
}

// GetMetadataAddr gets `AVFrame.metadata` address.
func (frame *AVFrame) GetMetadataAddr() **AVDictionary {
	return (**AVDictionary)(unsafe.Pointer(&frame.metadata))
}

// GetDecodeErrorFlags gets `AVFrame.decode_error_flags` value.
func (frame *AVFrame) GetDecodeErrorFlags() int32 {
	return (int32)(frame.decode_error_flags)
}

// SetDecodeErrorFlags sets `AVFrame.decode_error_flags` value.
func (frame *AVFrame) SetDecodeErrorFlags(v int32) {
	frame.decode_error_flags = (C.int)(v)
}

// GetDecodeErrorFlagsAddr gets `AVFrame.decode_error_flags` address.
func (frame *AVFrame) GetDecodeErrorFlagsAddr() *int32 {
	return (*int32)(&frame.decode_error_flags)
}

const (
	FF_DECODE_ERROR_INVALID_BITSTREAM  = int32(C.FF_DECODE_ERROR_INVALID_BITSTREAM)
	FF_DECODE_ERROR_MISSING_REFERENCE  = int32(C.FF_DECODE_ERROR_MISSING_REFERENCE)
	FF_DECODE_ERROR_CONCEALMENT_ACTIVE = int32(C.FF_DECODE_ERROR_CONCEALMENT_ACTIVE)
	FF_DECODE_ERROR_DECODE_SLICES      = int32(C.FF_DECODE_ERROR_DECODE_SLICES)
)

// GetPktSize gets `AVFrame.pkt_size` value.
func (frame *AVFrame) GetPktSize() int32 {
	return (int32)(frame.pkt_size)
}

// SetPktSize sets `AVFrame.pkt_size` value.
func (frame *AVFrame) SetPktSize(v int32) {
	frame.pkt_size = (C.int)(v)
}

// GetPktSizeAddr gets `AVFrame.pkt_size` address.
func (frame *AVFrame) GetPktSizeAddr() *int32 {
	return (*int32)(&frame.pkt_size)
}

// GetHwFramesCtx gets `AVFrame.hw_frames_ctx` value.
func (frame *AVFrame) GetHwFramesCtx() *AVBufferRef {
	return (*AVBufferRef)(frame.hw_frames_ctx)
}

// SetHwFramesCtx sets `AVFrame.hw_frames_ctx` value.
func (frame *AVFrame) SetHwFramesCtx(v *AVBufferRef) {
	frame.hw_frames_ctx = (*C.struct_AVBufferRef)(v)
}

// GetHwFramesCtxAddr gets `AVFrame.hw_frames_ctx` address.
func (frame *AVFrame) GetHwFramesCtxAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&frame.hw_frames_ctx))
}

// GetOpaqueRef gets `AVFrame.opaque_ref` value.
func (frame *AVFrame) GetOpaqueRef() *AVBufferRef {
	return (*AVBufferRef)(frame.opaque_ref)
}

// SetOpaqueRef sets `AVFrame.opaque_ref` value.
func (frame *AVFrame) SetOpaqueRef(v *AVBufferRef) {
	frame.opaque_ref = (*C.AVBufferRef)(v)
}

// GetOpaqueRefAddr gets `AVFrame.opaque_ref` address.
func (frame *AVFrame) GetOpaqueRefAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&frame.opaque_ref))
}

// GetCropTop gets `AVFrame.crop_top` value.
func (frame *AVFrame) GetCropTop() uintptr {
	return (uintptr)(frame.crop_top)
}

// SetCropTop sets `AVFrame.crop_top` value.
func (frame *AVFrame) SetCropTop(v uintptr) {
	frame.crop_top = (C.size_t)(v)
}

// GetCropTopAddr gets `AVFrame.crop_top` address.
func (frame *AVFrame) GetCropTopAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&frame.crop_top))
}

// GetCropBottom gets `AVFrame.crop_bottom` value.
func (frame *AVFrame) GetCropBottom() uintptr {
	return (uintptr)(frame.crop_bottom)
}

// SetCropBottom sets `AVFrame.crop_bottom` value.
func (frame *AVFrame) SetCropBottom(v uintptr) {
	frame.crop_bottom = (C.size_t)(v)
}

// GetCropBottomAddr gets `AVFrame.crop_bottom` address.
func (frame *AVFrame) GetCropBottomAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&frame.crop_bottom))
}

// GetCropLeft gets `AVFrame.crop_left` value.
func (frame *AVFrame) GetCropLeft() uintptr {
	return (uintptr)(frame.crop_left)
}

// SetCropLeft sets `AVFrame.crop_left` value.
func (frame *AVFrame) SetCropLeft(v uintptr) {
	frame.crop_left = (C.size_t)(v)
}

// GetCropLeftAddr gets `AVFrame.crop_left` address.
func (frame *AVFrame) GetCropLeftAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&frame.crop_left))
}

// GetCropRight gets `AVFrame.crop_right` value.
func (frame *AVFrame) GetCropRight() uintptr {
	return (uintptr)(frame.crop_right)
}

// SetCropRight sets `AVFrame.crop_right` value.
func (frame *AVFrame) SetCropRight(v uintptr) {
	frame.crop_right = (C.size_t)(v)
}

// GetCropRightAddr gets `AVFrame.crop_right` address.
func (frame *AVFrame) GetCropRightAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&frame.crop_right))
}

// GetPrivateRef gets `AVFrame.private_ref` value.
func (frame *AVFrame) GetPrivateRef() *AVBufferRef {
	return (*AVBufferRef)(frame.private_ref)
}

// SetPrivateRef sets `AVFrame.private_ref` value.
func (frame *AVFrame) SetPrivateRef(v *AVBufferRef) {
	frame.private_ref = (*C.struct_AVBufferRef)(v)
}

// GetPrivateRefAddr gets `AVFrame.private_ref` address.
func (frame *AVFrame) GetPrivateRefAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&frame.private_ref))
}

// GetChLayout gets `AVFrame.ch_layout` value.
func (frame *AVFrame) GetChLayout() AVChannelLayout {
	return (AVChannelLayout)(frame.ch_layout)
}

// SetChLayout sets `AVFrame.ch_layout` value.
func (frame *AVFrame) SetChLayout(v AVChannelLayout) {
	frame.ch_layout = (C.struct_AVChannelLayout)(v)
}

// GetChLayoutAddr gets `AVFrame.ch_layout` address.
func (frame *AVFrame) GetChLayoutAddr() *AVChannelLayout {
	return (*AVChannelLayout)(&frame.ch_layout)
}

// GetDuration gets `AVFrame.duration` value.
func (frame *AVFrame) GetDuration() int64 {
	return (int64)(frame.duration)
}

// SetDuration sets `AVFrame.duration` value.
func (frame *AVFrame) SetDuration(v int64) {
	frame.duration = (C.int64_t)(v)
}

// GetDurationAddr gets `AVFrame.duration` address.
func (frame *AVFrame) GetDurationAddr() *int64 {
	return (*int64)(&frame.duration)
}

// AvFrameAlloc allocates an AVFrame and set its fields to default values.
// The resulting struct must be freed using AVFrameFree().
func AvFrameAlloc() *AVFrame {
	return (*AVFrame)(C.av_frame_alloc())
}

// AvFrameFree frees the frame and any dynamically allocated objects in it,
// e.g. extended_data. If the frame is reference counted, it will be
// unreferenced first.
func AvFrameFree(frame **AVFrame) {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(frame)))
}

// AvFrameRef sets up a new reference to the data described by the source frame.
func AvFrameRef(dst, src *AVFrame) int32 {
	return (int32)(C.av_frame_ref((*C.struct_AVFrame)(dst), (*C.struct_AVFrame)(src)))
}

// AvFrameReplace ensures the destination frame refers to the same data described by the source frame.
func AvFrameReplace(dst, src *AVFrame) int32 {
	return (int32)(C.av_frame_replace((*C.struct_AVFrame)(dst), (*C.struct_AVFrame)(src)))
}

// AvFrameClone creates a new frame that references the same data as src.
func AvFrameClone(frame *AVFrame) *AVFrame {
	return (*AVFrame)(C.av_frame_clone((*C.struct_AVFrame)(frame)))
}

// AvFrameUnref unreferences all the buffers referenced by frame and reset the frame fields.
func AvFrameUnref(frame *AVFrame) {
	C.av_frame_unref((*C.struct_AVFrame)(frame))
}

// AvFrameMoveRef moves everything contained in src to dst and reset src.
func AvFrameMoveRef(dst, src *AVFrame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(dst), (*C.struct_AVFrame)(src))
}

// AvFrameGetBuffer allocates new buffer(s) for audio or video data.
func AvFrameGetBuffer(frame *AVFrame, align int32) int32 {
	return (int32)(C.av_frame_get_buffer((*C.struct_AVFrame)(frame), (C.int)(align)))
}

// AvFrameIsWritable checks if the frame data is writable.
func AvFrameIsWritable(frame *AVFrame) int32 {
	return (int32)(C.av_frame_is_writable((*C.struct_AVFrame)(frame)))
}

// AvFrameMakeWritable ensures that the frame data is writable, avoiding data copy if possible.
func AvFrameMakeWritable(frame *AVFrame) int32 {
	return (int32)(C.av_frame_make_writable((*C.struct_AVFrame)(frame)))
}

// AvFrameCopy copies the frame data from src to dst.
func AvFrameCopy(dst, src *AVFrame) int32 {
	return (int32)(C.av_frame_copy((*C.struct_AVFrame)(dst), (*C.struct_AVFrame)(src)))
}

// AvFrameCopyProps copies only "metadata" fields from src to dst.
func AvFrameCopyProps(dst, src *AVFrame) int32 {
	return (int32)(C.av_frame_copy_props((*C.struct_AVFrame)(dst), (*C.struct_AVFrame)(src)))
}

// AvFrameGetPlaneBuffer gets the buffer reference a given data plane is stored in.
func AvFrameGetPlaneBuffer(frame *AVFrame, plane int32) *AVBufferRef {
	return (*AVBufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(frame), (C.int)(plane)))
}

// AvFrameNewSideData adds a new side data to a frame.
func AvFrameNewSideData(frame *AVFrame, _type AVFrameSideDataType, size uintptr) *AVFrameSideData {
	return (*AVFrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(frame),
		(C.enum_AVFrameSideDataType)(_type), (C.size_t)(size)))
}

// AvFrameNewSideDataFromBuf adds a new side data to a frame from an existing AVBufferRef.
func AvFrameNewSideDataFromBuf(frame *AVFrame, _type AVFrameSideDataType,
	buf *AVBufferRef) *AVFrameSideData {
	return (*AVFrameSideData)(C.av_frame_new_side_data_from_buf((*C.struct_AVFrame)(frame),
		(C.enum_AVFrameSideDataType)(_type), (*C.struct_AVBufferRef)(buf)))
}

// AvFrameGetSideData gets a pointer to the side data of a given type on success, NULL if there
// is no side data with such type in this frame.
func AvFrameGetSideData(frame *AVFrame, _type AVFrameSideDataType) *AVFrameSideData {
	return (*AVFrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(frame),
		(C.enum_AVFrameSideDataType)(_type)))
}

// AvFrameRemoveSideData remove and free all side data instances of the given type.
func AvFrameRemoveSideData(frame *AVFrame, _type AVFrameSideDataType) {
	C.av_frame_remove_side_data((*C.struct_AVFrame)(frame), (C.enum_AVFrameSideDataType)(_type))
}

// Flags for frame cropping.
const (
	AV_FRAME_CROP_UNALIGNED = int32(C.AV_FRAME_CROP_UNALIGNED)
)

// AvFrameApplyCropping crops the given video AVFrame according to its crop_left/crop_top/crop_right/
// crop_bottom fields.
func AvFrameApplyCropping(frame *AVFrame, flags int32) int32 {
	return (int32)(C.av_frame_apply_cropping((*C.struct_AVFrame)(frame), (C.int)(flags)))
}

// AvFrameSideDataName returns a string identifying the side data type.
func AvFrameSideDataName(_type AVFrameSideDataType) string {
	return C.GoString(C.av_frame_side_data_name((C.enum_AVFrameSideDataType)(_type)))
}

// AvFrameSideDataFree frees all side data entries and their contents.
func AvFrameSideDataFree(sd ***AVFrameSideData, nbSd *int32) {
	C.av_frame_side_data_free((***C.struct_AVFrameSideData)(unsafe.Pointer(sd)), (*C.int)(nbSd))
}

const (
	AV_FRAME_SIDE_DATA_FLAG_UNIQUE = uint32(C.AV_FRAME_SIDE_DATA_FLAG_UNIQUE)
)

// AvFrameSideDataNew adds new side data entry to an array.
func AvFrameSideDataNew(sd ***AVFrameSideData, nbSd *int32, _type AVFrameSideDataType, size uintptr, flags uint32) *AVFrameSideData {
	return (*AVFrameSideData)(C.av_frame_side_data_new(
		(***C.struct_AVFrameSideData)(unsafe.Pointer(sd)), (*C.int)(nbSd),
		(C.enum_AVFrameSideDataType)(_type), (C.size_t)(size), (C.uint)(flags)))
}

// AvFrameSideDataClone
func AvFrameSideDataClone(sd ***AVFrameSideData, nbSd *int32, src *AVFrameSideData, flags uint32) int32 {
	return (int32)(C.av_frame_side_data_clone((***C.struct_AVFrameSideData)(unsafe.Pointer(sd)), (*C.int)(nbSd),
		(*C.struct_AVFrameSideData)(src), (C.uint)(flags)))
}

// AvFrameSideDataGetC
func AvFrameSideDataGetC(sd **AVFrameSideData, nbSd int32, _type AVFrameSideDataType) *AVFrameSideData {
	return (*AVFrameSideData)(C.av_frame_side_data_get_c((**C.struct_AVFrameSideData)(unsafe.Pointer(sd)), (C.int)(nbSd),
		(C.enum_AVFrameSideDataType)(_type)))
}

// AvFrameSideDataGet
func AvFrameSideDataGet(sd **AVFrameSideData, nbSd int32, _type AVFrameSideDataType) *AVFrameSideData {
	return (*AVFrameSideData)(C.av_frame_side_data_get((**C.struct_AVFrameSideData)(unsafe.Pointer(sd)), (C.int)(nbSd),
		(C.enum_AVFrameSideDataType)(_type)))
}
