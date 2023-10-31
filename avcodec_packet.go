// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavcodec/packet.h>
*/
import "C"
import "unsafe"

// AVPacketSideDataType
type AVPacketSideDataType = C.enum_AVPacketSideDataType

const (
	AV_PKT_DATA_PALETTE                    = AVPacketSideDataType(C.AV_PKT_DATA_PALETTE)
	AV_PKT_DATA_NEW_EXTRADATA              = AVPacketSideDataType(C.AV_PKT_DATA_NEW_EXTRADATA)
	AV_PKT_DATA_PARAM_CHANGE               = AVPacketSideDataType(C.AV_PKT_DATA_PARAM_CHANGE)
	AV_PKT_DATA_H263_MB_INFO               = AVPacketSideDataType(C.AV_PKT_DATA_H263_MB_INFO)
	AV_PKT_DATA_REPLAYGAIN                 = AVPacketSideDataType(C.AV_PKT_DATA_REPLAYGAIN)
	AV_PKT_DATA_DISPLAYMATRIX              = AVPacketSideDataType(C.AV_PKT_DATA_DISPLAYMATRIX)
	AV_PKT_DATA_STEREO3D                   = AVPacketSideDataType(C.AV_PKT_DATA_STEREO3D)
	AV_PKT_DATA_AUDIO_SERVICE_TYPE         = AVPacketSideDataType(C.AV_PKT_DATA_AUDIO_SERVICE_TYPE)
	AV_PKT_DATA_QUALITY_STATS              = AVPacketSideDataType(C.AV_PKT_DATA_QUALITY_STATS)
	AV_PKT_DATA_FALLBACK_TRACK             = AVPacketSideDataType(C.AV_PKT_DATA_FALLBACK_TRACK)
	AV_PKT_DATA_CPB_PROPERTIES             = AVPacketSideDataType(C.AV_PKT_DATA_CPB_PROPERTIES)
	AV_PKT_DATA_SKIP_SAMPLES               = AVPacketSideDataType(C.AV_PKT_DATA_SKIP_SAMPLES)
	AV_PKT_DATA_JP_DUALMONO                = AVPacketSideDataType(C.AV_PKT_DATA_JP_DUALMONO)
	AV_PKT_DATA_STRINGS_METADATA           = AVPacketSideDataType(C.AV_PKT_DATA_STRINGS_METADATA)
	AV_PKT_DATA_SUBTITLE_POSITION          = AVPacketSideDataType(C.AV_PKT_DATA_SUBTITLE_POSITION)
	AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL   = AVPacketSideDataType(C.AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL)
	AV_PKT_DATA_WEBVTT_IDENTIFIER          = AVPacketSideDataType(C.AV_PKT_DATA_WEBVTT_IDENTIFIER)
	AV_PKT_DATA_WEBVTT_SETTINGS            = AVPacketSideDataType(C.AV_PKT_DATA_WEBVTT_SETTINGS)
	AV_PKT_DATA_METADATA_UPDATE            = AVPacketSideDataType(C.AV_PKT_DATA_METADATA_UPDATE)
	AV_PKT_DATA_MPEGTS_STREAM_ID           = AVPacketSideDataType(C.AV_PKT_DATA_MPEGTS_STREAM_ID)
	AV_PKT_DATA_MASTERING_DISPLAY_METADATA = AVPacketSideDataType(C.AV_PKT_DATA_MASTERING_DISPLAY_METADATA)
	AV_PKT_DATA_SPHERICAL                  = AVPacketSideDataType(C.AV_PKT_DATA_SPHERICAL)
	AV_PKT_DATA_CONTENT_LIGHT_LEVEL        = AVPacketSideDataType(C.AV_PKT_DATA_CONTENT_LIGHT_LEVEL)
	AV_PKT_DATA_A53_CC                     = AVPacketSideDataType(C.AV_PKT_DATA_A53_CC)
	AV_PKT_DATA_ENCRYPTION_INIT_INFO       = AVPacketSideDataType(C.AV_PKT_DATA_ENCRYPTION_INIT_INFO)
	AV_PKT_DATA_ENCRYPTION_INFO            = AVPacketSideDataType(C.AV_PKT_DATA_ENCRYPTION_INFO)
	AV_PKT_DATA_AFD                        = AVPacketSideDataType(C.AV_PKT_DATA_AFD)
	AV_PKT_DATA_PRFT                       = AVPacketSideDataType(C.AV_PKT_DATA_PRFT)
	AV_PKT_DATA_ICC_PROFILE                = AVPacketSideDataType(C.AV_PKT_DATA_ICC_PROFILE)
	AV_PKT_DATA_DOVI_CONF                  = AVPacketSideDataType(C.AV_PKT_DATA_DOVI_CONF)
	AV_PKT_DATA_S12M_TIMECODE              = AVPacketSideDataType(C.AV_PKT_DATA_S12M_TIMECODE)
	AV_PKT_DATA_DYNAMIC_HDR10_PLUS         = AVPacketSideDataType(C.AV_PKT_DATA_DYNAMIC_HDR10_PLUS)
	AV_PKT_DATA_NB                         = AVPacketSideDataType(C.AV_PKT_DATA_NB)
)

const (
	// Deprecated: No use.
	AV_PKT_DATA_QUALITY_FACTOR = AVPacketSideDataType(C.AV_PKT_DATA_QUALITY_FACTOR)
)

// AVPacketSideData
type AVPacketSideData C.struct_AVPacketSideData

// GetData gets `AVPacketSideData.data` value.
func (psd *AVPacketSideData) GetData() *uint8 {
	return (*uint8)(psd.data)
}

// SetData sets `AVPacketSideData.data` value.
func (psd *AVPacketSideData) SetData(v *uint8) {
	psd.data = (*C.uint8_t)(v)
}

// GetDataAddr gets `AVPacketSideData.data` address.
func (psd *AVPacketSideData) GetDataAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&psd.data))
}

// GetSize gets `AVPacketSideData.size` value.
func (psd *AVPacketSideData) GetSize() uintptr {
	return (uintptr)(psd.size)
}

// SetSize sets `AVPacketSideData.size` value.
func (psd *AVPacketSideData) SetSize(v uintptr) {
	psd.size = (C.size_t)(v)
}

// GetSizeAddr gets `AVPacketSideData.size` address.
func (psd *AVPacketSideData) GetSizeAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&psd.size))
}

// GetType gets `AVPacketSideData.type` value.
func (psd *AVPacketSideData) GetType() AVPacketSideDataType {
	return (AVPacketSideDataType)(psd._type)
}

// SetType sets `AVPacketSideData.type` value.
func (psd *AVPacketSideData) SetType(v AVPacketSideDataType) {
	psd._type = (C.enum_AVPacketSideDataType)(v)
}

// GetTypeAddr gets `AVPacketSideData.type` address.
func (psd *AVPacketSideData) GetTypeAddr() *AVPacketSideDataType {
	return (*AVPacketSideDataType)(&psd._type)
}

// AVPacket
type AVPacket C.struct_AVPacket

// GetBuf gets `AVPacket.buf` value.
func (pkt *AVPacket) GetBuf() *AVBufferRef {
	return (*AVBufferRef)(pkt.buf)
}

// SetBuf sets `AVPacket.buf` value.
func (pkt *AVPacket) SetBuf(v *AVBufferRef) {
	pkt.buf = (*C.struct_AVBufferRef)(v)
}

// GetBufAddr gets `AVPacket.buf` address.
func (pkt *AVPacket) GetBufAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&pkt.buf))
}

// GetPts gets `AVPacket.pts` value.
func (pkt *AVPacket) GetPts() int64 {
	return (int64)(pkt.pts)
}

// SetPts sets `AVPacket.pts` value.
func (pkt *AVPacket) SetPts(v int64) {
	pkt.pts = (C.int64_t)(v)
}

// GetPtsAddr gets `AVPacket.pts` address.
func (pkt *AVPacket) GetPtsAddr() *int64 {
	return (*int64)(&pkt.pts)
}

// GetDts gets `AVPacket.dts` value.
func (pkt *AVPacket) GetDts() int64 {
	return (int64)(pkt.dts)
}

// SetDts sets `AVPacket.dts` value.
func (pkt *AVPacket) SetDts(v int64) {
	pkt.dts = (C.int64_t)(v)
}

// GetDtsAddr gets `AVPacket.dts` address.
func (pkt *AVPacket) GetDtsAddr() *int64 {
	return (*int64)(&pkt.dts)
}

// GetData gets `AVPacket.data` value.
func (pkt *AVPacket) GetData() *uint8 {
	return (*uint8)(pkt.data)
}

// SetData sets `AVPacket.data` value.
func (pkt *AVPacket) SetData(v *uint8) {
	pkt.data = (*C.uint8_t)(v)
}

// GetDataAddr gets `AVPacket.data` address.
func (pkt *AVPacket) GetDataAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&pkt.data))
}

// GetSize gets `AVPacket.size` value.
func (pkt *AVPacket) GetSize() int32 {
	return (int32)(pkt.size)
}

// SetSize sets `AVPacket.size` value.
func (pkt *AVPacket) SetSize(v int32) {
	pkt.size = (C.int)(v)
}

// GetSizeAddr gets `AVPacket.size` address.
func (pkt *AVPacket) GetSizeAddr() *int32 {
	return (*int32)(&pkt.size)
}

// GetStreamIndex gets `AVPacket.stream_index` value.
func (pkt *AVPacket) GetStreamIndex() int32 {
	return (int32)(pkt.stream_index)
}

// SetStreamIndex sets `AVPacket.stream_index` value.
func (pkt *AVPacket) SetStreamIndex(v int32) {
	pkt.stream_index = (C.int)(v)
}

// GetStreamIndexAddr gets `AVPacket.stream_index` address.
func (pkt *AVPacket) GetStreamIndexAddr() *int32 {
	return (*int32)(&pkt.stream_index)
}

// GetFlags gets `AVPacket.flags` value.
func (pkt *AVPacket) GetFlags() int32 {
	return (int32)(pkt.flags)
}

// SetFlags sets `AVPacket.flags` value.
func (pkt *AVPacket) SetFlags(v int32) {
	pkt.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVPacket.flags` address.
func (pkt *AVPacket) GetFlagsAddr() *int32 {
	return (*int32)(&pkt.flags)
}

// GetSideData gets `AVPacket.side_data` value.
func (pkt *AVPacket) GetSideData() *AVPacketSideData {
	return (*AVPacketSideData)(pkt.side_data)
}

// SetSideData sets `AVPacket.side_data` value.
func (pkt *AVPacket) SetSideData(v *AVPacketSideData) {
	pkt.side_data = (*C.struct_AVPacketSideData)(v)
}

// GetSideDataAddr gets `AVPacket.side_data` address.
func (pkt *AVPacket) GetSideDataAddr() **AVPacketSideData {
	return (**AVPacketSideData)(unsafe.Pointer(&pkt.side_data))
}

// GetSideDataElems gets `AVPacket.side_data_elems` value.
func (pkt *AVPacket) GetSideDataElems() int32 {
	return (int32)(pkt.side_data_elems)
}

// SetSideDataElems sets `AVPacket.side_data_elems` value.
func (pkt *AVPacket) SetSideDataElems(v int32) {
	pkt.side_data_elems = (C.int)(v)
}

// GetSideDataElemsAddr gets `AVPacket.side_data_elems` address.
func (pkt *AVPacket) GetSideDataElemsAddr() *int32 {
	return (*int32)(&pkt.side_data_elems)
}

// GetDuration gets `AVPacket.duration` value.
func (pkt *AVPacket) GetDuration() int64 {
	return (int64)(pkt.duration)
}

// SetDuration sets `AVPacket.duration` value.
func (pkt *AVPacket) SetDuration(v int64) {
	pkt.duration = (C.int64_t)(v)
}

// GetDurationAddr gets `AVPacket.duration` address.
func (pkt *AVPacket) GetDurationAddr() *int64 {
	return (*int64)(&pkt.duration)
}

// GetPos gets `AVPacket.pos` value.
func (pkt *AVPacket) GetPos() int64 {
	return (int64)(pkt.pos)
}

// SetPos sets `AVPacket.pos` value.
func (pkt *AVPacket) SetPos(v int64) {
	pkt.pos = (C.int64_t)(v)
}

// GetPosAddr gets `AVPacket.pos` address.
func (pkt *AVPacket) GetPosAddr() *int64 {
	return (*int64)(&pkt.pos)
}

// GetOpaque gets `AVPacket.opaque` value.
func (pkt *AVPacket) GetOpaque() unsafe.Pointer {
	return pkt.opaque
}

// SetOpaque sets `AVPacket.opaque` value.
func (pkt *AVPacket) SetOpaque(v CVoidPointer) {
	pkt.opaque = VoidPointer(v)
}

// GetOpaqueAddr gets `AVPacket.opaque` address.
func (pkt *AVPacket) GetOpaqueAddr() *unsafe.Pointer {
	return &pkt.opaque
}

// GetOpaqueRef gets `AVPacket.opaque_ref` value.
func (pkt *AVPacket) GetOpaqueRef() *AVBufferRef {
	return (*AVBufferRef)(pkt.opaque_ref)
}

// SetOpaqueRef sets `AVPacket.opaque_ref` value.
func (pkt *AVPacket) SetOpaqueRef(v *AVBufferRef) {
	pkt.opaque_ref = (*C.struct_AVBufferRef)(v)
}

// GetOpaqueRefAddr gets `AVPacket.opaque_ref` address.
func (pkt *AVPacket) GetOpaqueRefAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&pkt.opaque_ref))
}

// GetTimeBase gets `AVPacket.time_base` value.
func (pkt *AVPacket) GetTimeBase() AVRational {
	return (AVRational)(pkt.time_base)
}

// SetTimeBase sets `AVPacket.time_base` value.
func (pkt *AVPacket) SetTimeBase(v AVRational) {
	pkt.time_base = (C.struct_AVRational)(v)
}

// GetTimeBaseAddr gets `AVPacket.time_base` address.
func (pkt *AVPacket) GetTimeBaseAddr() *AVRational {
	return (*AVRational)(&pkt.time_base)
}

// Deprecated: No use.
//
// AVPacketList
type AVPacketList C.struct_AVPacketList

// Getgets `AVPacketList.next` value.
func (pl *AVPacketList) Get() *AVPacketList {
	return (*AVPacketList)(pl.next)
}

// Set sets `AVPacketList.next` value.
func (pl *AVPacketList) Set(v *AVPacketList) {
	pl.next = (*C.struct_AVPacketList)(v)
}

// GetAddr gets `AVPacketList.next` address.
func (pl *AVPacketList) GetAddr() **AVPacketList {
	return (**AVPacketList)(unsafe.Pointer(&pl.next))
}

const (
	AV_PKT_FLAG_KEY        = C.AV_PKT_FLAG_KEY
	AV_PKT_FLAG_CORRUPT    = C.AV_PKT_FLAG_CORRUPT
	AV_PKT_FLAG_DISCARD    = C.AV_PKT_FLAG_DISCARD
	AV_PKT_FLAG_TRUSTED    = C.AV_PKT_FLAG_TRUSTED
	AV_PKT_FLAG_DISPOSABLE = C.AV_PKT_FLAG_DISPOSABLE
)

// AVSideDataParamChangeFlags
type AVSideDataParamChangeFlags = C.enum_AVSideDataParamChangeFlags

const (
	AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT  = AVSideDataParamChangeFlags(C.AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT)
	AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT = AVSideDataParamChangeFlags(C.AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT)
	AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE    = AVSideDataParamChangeFlags(C.AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE)
	AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS     = AVSideDataParamChangeFlags(C.AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS)
)

// AvPacketAlloc allocates an AVPacket and set its fields to default values. The resulting
// struct must be freed using AVPacketFree().
func AvPacketAlloc() *AVPacket {
	return (*AVPacket)(C.av_packet_alloc())
}

// AvPacketClone creates a new packet that references the same data as src.
func AvPacketClone(pkt *AVPacket) *AVPacket {
	return (*AVPacket)(C.av_packet_clone((*C.struct_AVPacket)(pkt)))
}

// AvPacketFree frees the packet, if the packet is reference counted, it will be
// unreferenced first.
func AvPacketFree(pkt **AVPacket) {
	C.av_packet_free((**C.struct_AVPacket)(unsafe.Pointer(pkt)))
}

// Deprecated: No use.
//
// AvInitPacket initializes optional fields of a packet with default values.
func AvInitPacket(pkt *AVPacket) {
	C.av_init_packet((*C.struct_AVPacket)(pkt))
}

// AvNewPacket allocates the payload of a packet and initialize its fields with
// default values.
func AvNewPacket(pkt *AVPacket, size int32) int32 {
	return (int32)(C.av_new_packet((*C.struct_AVPacket)(pkt), (C.int)(size)))
}

// AvShrinkPacket reduces packet size, correctly zeroing padding.
func AvShrinkPacket(pkt *AVPacket, size int32) {
	C.av_shrink_packet((*C.struct_AVPacket)(pkt), (C.int)(size))
}

// AvGrowPacket increases packet size, correctly zeroing padding.
func AvGrowPacket(pkt *AVPacket, growBy int32) int32 {
	return (int32)(C.av_grow_packet((*C.struct_AVPacket)(pkt), (C.int)(growBy)))
}

// AvPacketFromData initializes a reference-counted packet from AvMalloc()ed data.
func AvPacketFromData(pkt *AVPacket, data *uint8, size int32) int32 {
	return (int32)(C.av_packet_from_data((*C.struct_AVPacket)(pkt),
		(*C.uint8_t)(data), (C.int)(size)))
}

// AvPacketNewSideData allocates new information of a packet.
func AvPacketNewSideData(pkt *AVPacket, _type AVPacketSideDataType, size uintptr) *uint8 {
	return (*uint8)(C.av_packet_new_side_data((*C.struct_AVPacket)(pkt),
		(C.enum_AVPacketSideDataType)(_type), (C.size_t)(size)))
}

// AvPacketAddSideData wraps an existing array as a packet side data.
func AvPacketAddSideData(pkt *AVPacket, _type AVPacketSideDataType, data *uint8, size uintptr) int32 {
	return (int32)(C.av_packet_add_side_data((*C.struct_AVPacket)(pkt),
		(C.enum_AVPacketSideDataType)(_type), (*C.uint8_t)(data), (C.size_t)(size)))
}

// AvPacketShrinkSideData shrinks the already allocated side data buffer.
func AvPacketShrinkSideData(pkt *AVPacket, _type AVPacketSideDataType, size uintptr) int32 {
	return (int32)(C.av_packet_shrink_side_data((*C.struct_AVPacket)(pkt),
		(C.enum_AVPacketSideDataType)(_type), (C.size_t)(size)))
}

// AvPacketGetSideData gets side information from packet.
func AvPacketGetSideData(pkt *AVPacket, _type AVPacketSideDataType, size *uintptr) *uint8 {
	return (*uint8)(C.av_packet_get_side_data((*C.struct_AVPacket)(pkt),
		(C.enum_AVPacketSideDataType)(_type), (*C.size_t)(unsafe.Pointer(size))))
}

// AvPacketPackDictionary packs a dictionary for use in side_data.
func AvPacketPackDictionary(dict *AVDictionary, size *uintptr) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(dict),
		(*C.size_t)(unsafe.Pointer(size))))
}

// AvPacketUnpackDictionary unpacks a dictionary from side_data.
func AvPacketUnpackDictionary(data *uint8, size uintptr, dict **AVDictionary) int32 {
	return (int32)(C.av_packet_unpack_dictionary((*C.uint8_t)(data), (C.size_t)(size),
		(**C.struct_AVDictionary)(unsafe.Pointer(dict))))
}

// AvPacketFreeSideData is a convenience function to free all the side data stored.
func AvPacketFreeSideData(pkt *AVPacket) {
	C.av_packet_free_side_data((*C.struct_AVPacket)(pkt))
}

// AvPacketRef setups a new reference to the data described by a given packet.
func AvPacketRef(dst, src *AVPacket) int32 {
	return (int32)(C.av_packet_ref((*C.struct_AVPacket)(dst), (*C.struct_AVPacket)(src)))
}

// AvPacketUnref unreferences the buffer referenced by the packet and reset the
// remaining packet fields to their default values.
func AvPacketUnref(pkt *AVPacket) {
	C.av_packet_unref((*C.struct_AVPacket)(pkt))
}

// AvPacketMoveRef moves every field in src to dst and reset src.
func AvPacketMoveRef(dst, src *AVPacket) {
	C.av_packet_move_ref((*C.struct_AVPacket)(dst), (*C.struct_AVPacket)(src))
}

// AvPacketCopyProps copies only "properties" fields from src to dst.
func AvPacketCopyProps(dst, src *AVPacket) int32 {
	return (int32)(C.av_packet_copy_props((*C.struct_AVPacket)(dst), (*C.struct_AVPacket)(src)))
}

// AvPacketMakeRefcounted ensures the data described by a given packet is reference counted.
func AvPacketMakeRefcounted(pkt *AVPacket) {
	C.av_packet_make_refcounted((*C.struct_AVPacket)(pkt))
}

// AvPacketMakeWritable creates a writable reference for the data described by a given packet,
// avoiding data copy if possible.
func AvPacketMakeWritable(pkt *AVPacket) {
	C.av_packet_make_writable((*C.struct_AVPacket)(pkt))
}

// AvPacketRescaleTs converts valid timing fields (timestamps / durations) in a packet from one
// timebase to another. Timestamps with unknown values (AV_NOPTS_VALUE) will be ignored.
func AvPacketRescaleTs(pkt *AVPacket, tbSrc, tbDst AVRational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(pkt),
		(C.struct_AVRational)(tbSrc), (C.struct_AVRational)(tbDst))
}
