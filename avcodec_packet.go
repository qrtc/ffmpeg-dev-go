package ffmpeg

/*
#include <libavcodec/packet.h>
*/
import "C"
import "unsafe"

// AvPacketSideDataType
type AvPacketSideDataType int32

const (
	AV_PKT_DATA_PALETTE                    = AvPacketSideDataType(C.AV_PKT_DATA_PALETTE)
	AV_PKT_DATA_NEW_EXTRADATA              = AvPacketSideDataType(C.AV_PKT_DATA_NEW_EXTRADATA)
	AV_PKT_DATA_PARAM_CHANGE               = AvPacketSideDataType(C.AV_PKT_DATA_PARAM_CHANGE)
	AV_PKT_DATA_H263_MB_INFO               = AvPacketSideDataType(C.AV_PKT_DATA_H263_MB_INFO)
	AV_PKT_DATA_REPLAYGAIN                 = AvPacketSideDataType(C.AV_PKT_DATA_REPLAYGAIN)
	AV_PKT_DATA_DISPLAYMATRIX              = AvPacketSideDataType(C.AV_PKT_DATA_DISPLAYMATRIX)
	AV_PKT_DATA_STEREO3D                   = AvPacketSideDataType(C.AV_PKT_DATA_STEREO3D)
	AV_PKT_DATA_AUDIO_SERVICE_TYPE         = AvPacketSideDataType(C.AV_PKT_DATA_AUDIO_SERVICE_TYPE)
	AV_PKT_DATA_QUALITY_STATS              = AvPacketSideDataType(C.AV_PKT_DATA_QUALITY_STATS)
	AV_PKT_DATA_FALLBACK_TRACK             = AvPacketSideDataType(C.AV_PKT_DATA_FALLBACK_TRACK)
	AV_PKT_DATA_CPB_PROPERTIES             = AvPacketSideDataType(C.AV_PKT_DATA_CPB_PROPERTIES)
	AV_PKT_DATA_SKIP_SAMPLES               = AvPacketSideDataType(C.AV_PKT_DATA_SKIP_SAMPLES)
	AV_PKT_DATA_JP_DUALMONO                = AvPacketSideDataType(C.AV_PKT_DATA_JP_DUALMONO)
	AV_PKT_DATA_STRINGS_METADATA           = AvPacketSideDataType(C.AV_PKT_DATA_STRINGS_METADATA)
	AV_PKT_DATA_SUBTITLE_POSITION          = AvPacketSideDataType(C.AV_PKT_DATA_SUBTITLE_POSITION)
	AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL   = AvPacketSideDataType(C.AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL)
	AV_PKT_DATA_WEBVTT_IDENTIFIER          = AvPacketSideDataType(C.AV_PKT_DATA_WEBVTT_IDENTIFIER)
	AV_PKT_DATA_WEBVTT_SETTINGS            = AvPacketSideDataType(C.AV_PKT_DATA_WEBVTT_SETTINGS)
	AV_PKT_DATA_METADATA_UPDATE            = AvPacketSideDataType(C.AV_PKT_DATA_METADATA_UPDATE)
	AV_PKT_DATA_MPEGTS_STREAM_ID           = AvPacketSideDataType(C.AV_PKT_DATA_MPEGTS_STREAM_ID)
	AV_PKT_DATA_MASTERING_DISPLAY_METADATA = AvPacketSideDataType(C.AV_PKT_DATA_MASTERING_DISPLAY_METADATA)
	AV_PKT_DATA_SPHERICAL                  = AvPacketSideDataType(C.AV_PKT_DATA_SPHERICAL)
	AV_PKT_DATA_CONTENT_LIGHT_LEVEL        = AvPacketSideDataType(C.AV_PKT_DATA_CONTENT_LIGHT_LEVEL)
	AV_PKT_DATA_A53_CC                     = AvPacketSideDataType(C.AV_PKT_DATA_A53_CC)
	AV_PKT_DATA_ENCRYPTION_INIT_INFO       = AvPacketSideDataType(C.AV_PKT_DATA_ENCRYPTION_INIT_INFO)
	AV_PKT_DATA_ENCRYPTION_INFO            = AvPacketSideDataType(C.AV_PKT_DATA_ENCRYPTION_INFO)
	AV_PKT_DATA_AFD                        = AvPacketSideDataType(C.AV_PKT_DATA_AFD)
	AV_PKT_DATA_PRFT                       = AvPacketSideDataType(C.AV_PKT_DATA_PRFT)
	AV_PKT_DATA_ICC_PROFILE                = AvPacketSideDataType(C.AV_PKT_DATA_ICC_PROFILE)
	AV_PKT_DATA_DOVI_CONF                  = AvPacketSideDataType(C.AV_PKT_DATA_DOVI_CONF)
	AV_PKT_DATA_S12M_TIMECODE              = AvPacketSideDataType(C.AV_PKT_DATA_S12M_TIMECODE)
	AV_PKT_DATA_NB                         = AvPacketSideDataType(C.AV_PKT_DATA_NB)
)

const (
	// Deprecated: No use
	AV_PKT_DATA_QUALITY_FACTOR = AvPacketSideDataType(C.AV_PKT_DATA_QUALITY_FACTOR)
)

// AvPacketSideData
type AvPacketSideData C.struct_AVPacketSideData

// AvPacket
type AvPacket C.struct_AVPacket

// Custom: GetBuf gets `AVPacket.buf` value.
func (pkt *AvPacket) GetBuf() *AvBufferRef {
	return (*AvBufferRef)(pkt.buf)
}

// Custom: SetBuf sets `AVPacket.buf` value.
func (pkt *AvPacket) SetBuf(v *AvBufferRef) {
	pkt.buf = (*C.struct_AVBufferRef)(v)
}

// Custom: GetBufAddr gets `AVPacket.buf` address.
func (pkt *AvPacket) GetBufAddr() **AvBufferRef {
	return (**AvBufferRef)(unsafe.Pointer(&pkt.buf))
}

// Custom: GetPts gets `AVPacket.pts` value.
func (pkt *AvPacket) GetPts() int64 {
	return (int64)(pkt.pts)
}

// Custom: SetPts sets `AVPacket.pts` value.
func (pkt *AvPacket) SetPts(v int64) {
	pkt.pts = (C.int64_t)(v)
}

// Custom: GetPtsAddr gets `AVPacket.pts` address.
func (pkt *AvPacket) GetPtsAddr() *int64 {
	return (*int64)(&pkt.pts)
}

// Custom: GetDts gets `AVPacket.dts` value.
func (pkt *AvPacket) GetDts() int64 {
	return (int64)(pkt.dts)
}

// Custom: SetDts sets `AVPacket.dts` value.
func (pkt *AvPacket) SetDts(v int64) {
	pkt.dts = (C.int64_t)(v)
}

// Custom: GetDtsAddr gets `AVPacket.dts` address.
func (pkt *AvPacket) GetDtsAddr() *int64 {
	return (*int64)(&pkt.dts)
}

// Custom: GetData gets `AVPacket.data` value.
func (pkt *AvPacket) GetData() *uint8 {
	return (*uint8)(pkt.data)
}

// Custom: SetData sets `AVPacket.data` value.
func (pkt *AvPacket) SetData(v *uint8) {
	pkt.data = (*C.uint8_t)(v)
}

// Custom: GetDataAddr gets `AVPacket.data` address.
func (pkt *AvPacket) GetDataAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&pkt.data))
}

// Custom: GetSize gets `AVPacket.size` value.
func (pkt *AvPacket) GetSize() int32 {
	return (int32)(pkt.size)
}

// Custom: SetSize sets `AVPacket.size` value.
func (pkt *AvPacket) SetSize(v int32) {
	pkt.size = (C.int)(v)
}

// Custom: GetSizeAddr gets `AVPacket.size` address.
func (pkt *AvPacket) GetSizeAddr() *int32 {
	return (*int32)(&pkt.size)
}

// Custom: GetStreamIndex gets `AVPacket.stream_index` value.
func (pkt *AvPacket) GetStreamIndex() int32 {
	return (int32)(pkt.stream_index)
}

// Custom: SetStreamIndex sets `AVPacket.stream_index` value.
func (pkt *AvPacket) SetStreamIndex(v int32) {
	pkt.stream_index = (C.int)(v)
}

// Custom: GetStreamIndexAddr gets `AVPacket.stream_index` address.
func (pkt *AvPacket) GetStreamIndexAddr() *int32 {
	return (*int32)(&pkt.stream_index)
}

// Custom: GetFlags gets `AVPacket.flags` value.
func (pkt *AvPacket) GetFlags() int32 {
	return (int32)(pkt.flags)
}

// Custom: SetFlags sets `AVPacket.flags` value.
func (pkt *AvPacket) SetFlags(v int32) {
	pkt.flags = (C.int)(v)
}

// Custom: GetFlagsAddr gets `AVPacket.flags` address.
func (pkt *AvPacket) GetFlagsAddr() *int32 {
	return (*int32)(&pkt.flags)
}

// Custom: GetSideData gets `AVPacket.side_data` value.
func (pkt *AvPacket) GetSideData() *AvPacketSideData {
	return (*AvPacketSideData)(pkt.side_data)
}

// Custom: SetSideData sets `AVPacket.side_data` value.
func (pkt *AvPacket) SetSideData(v *AvPacketSideData) {
	pkt.side_data = (*C.struct_AVPacketSideData)(v)
}

// Custom: GetSideDataAddr gets `AVPacket.side_data` address.
func (pkt *AvPacket) GetSideDataAddr() **AvPacketSideData {
	return (**AvPacketSideData)(unsafe.Pointer(&pkt.side_data))
}

// Custom: GetSideDataElems gets `AVPacket.side_data_elems` value.
func (pkt *AvPacket) GetSideDataElems() int32 {
	return (int32)(pkt.side_data_elems)
}

// Custom: SetSideDataElems sets `AVPacket.side_data_elems` value.
func (pkt *AvPacket) SetSideDataElems(v int32) {
	pkt.side_data_elems = (C.int)(v)
}

// Custom: GetSideDataElemsAddr gets `AVPacket.side_data_elems` address.
func (pkt *AvPacket) GetSideDataElemsAddr() *int32 {
	return (*int32)(&pkt.side_data_elems)
}

// Custom: GetDuration gets `AVPacket.duration` value.
func (pkt *AvPacket) GetDuration() int64 {
	return (int64)(pkt.duration)
}

// Custom: SetDuration sets `AVPacket.duration` value.
func (pkt *AvPacket) SetDuration(v int64) {
	pkt.duration = (C.int64_t)(v)
}

// Custom: GetDurationAddr gets `AVPacket.duration` address.
func (pkt *AvPacket) GetDurationAddr() *int64 {
	return (*int64)(&pkt.duration)
}

// Custom: GetPos gets `AVPacket.pos` value.
func (pkt *AvPacket) GetPos() int64 {
	return (int64)(pkt.pos)
}

// Custom: SetPos sets `AVPacket.pos` value.
func (pkt *AvPacket) SetPos(v int64) {
	pkt.pos = (C.int64_t)(v)
}

// Custom: GetPosAddr gets `AVPacket.pos` address.
func (pkt *AvPacket) GetPosAddr() *int64 {
	return (*int64)(&pkt.pos)
}

// Custom: GetConvergenceDuration gets `AVPacket.convergence_duration` value.
func (pkt *AvPacket) GetConvergenceDuration() int64 {
	return (int64)(pkt.convergence_duration)
}

// Custom: SetConvergenceDuration sets `AVPacket.convergence_duration` value.
func (pkt *AvPacket) SetConvergenceDuration(v int64) {
	pkt.convergence_duration = (C.int64_t)(v)
}

// Custom: GetConvergenceDurationAddr gets `AVPacket.convergence_duration` address.
func (pkt *AvPacket) GetConvergenceDurationAddr() *int64 {
	return (*int64)(&pkt.convergence_duration)
}

// AvPacketList
type AvPacketList C.struct_AVPacketList

const (
	AV_PKT_FLAG_KEY        = C.AV_PKT_FLAG_KEY
	AV_PKT_FLAG_CORRUPT    = C.AV_PKT_FLAG_CORRUPT
	AV_PKT_FLAG_DISCARD    = C.AV_PKT_FLAG_DISCARD
	AV_PKT_FLAG_TRUSTED    = C.AV_PKT_FLAG_TRUSTED
	AV_PKT_FLAG_DISPOSABLE = C.AV_PKT_FLAG_DISPOSABLE
)

// AvSideDataParamChangeFlags
type AvSideDataParamChangeFlags int32

const (
	AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT  = AvSideDataParamChangeFlags(C.AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT)
	AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT = AvSideDataParamChangeFlags(C.AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT)
	AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE    = AvSideDataParamChangeFlags(C.AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE)
	AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS     = AvSideDataParamChangeFlags(C.AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS)
)

// AvPacketAlloc allocates an AVPacket and set its fields to default values.  The resulting
// struct must be freed using AvPacketFree().
func AvPacketAlloc() *AvPacket {
	return (*AvPacket)(C.av_packet_alloc())
}

// AvPacketClone creates a new packet that references the same data as src.
func AvPacketClone(pkt *AvPacket) *AvPacket {
	return (*AvPacket)(C.av_packet_clone((*C.struct_AVPacket)(pkt)))
}

// AvPacketFree frees the packet, if the packet is reference counted, it will be
// unreferenced first.
func AvPacketFree(pkt **AvPacket) {
	C.av_packet_free((**C.struct_AVPacket)(unsafe.Pointer(pkt)))
}

// AvNewPacket allocates the payload of a packet and initialize its fields with
// default values.
func AvNewPacket(pkt *AvPacket, size int32) int32 {
	return (int32)(C.av_new_packet((*C.struct_AVPacket)(pkt), (C.int)(size)))
}

// AvShrinkPacket reduces packet size, correctly zeroing padding
func AvShrinkPacket(pkt *AvPacket, size int32) {
	C.av_shrink_packet((*C.struct_AVPacket)(pkt), (C.int)(size))
}

// AvGrowPacket increases packet size, correctly zeroing padding
func AvGrowPacket(pkt *AvPacket, growBy int32) int32 {
	return (int32)(C.av_grow_packet((*C.struct_AVPacket)(pkt), (C.int)(growBy)))
}

// AvPacketFromData initializes a reference-counted packet from AvMalloc()ed data.
func AvPacketFromData(pkt *AvPacket, data *uint8, size int32) int32 {
	return (int32)(C.av_packet_from_data((*C.struct_AVPacket)(pkt),
		(*C.uint8_t)(data), (C.int)(size)))
}

// Deprecated: Use AvPacketRef() or AvPacketMakeRefcounted() instead.
func AvDupPacket(pkt *AvPacket) {
	C.av_dup_packet((*C.struct_AVPacket)(pkt))
}

// Deprecated: Use AvPacketRef instead.
// AvCopyPacket copies packet, including contents
func AvCopyPacket(dst, src *AvPacket) int32 {
	return (int32)(C.av_copy_packet((*C.struct_AVPacket)(dst), (*C.struct_AVPacket)(src)))
}

// Deprecated: Use AvPacketCopyProps instead.
// AvCopyPacketSideData copies packet side data
func AvCopyPacketSideData(dst, src *AvPacket) int32 {
	return (int32)(C.av_copy_packet_side_data((*C.struct_AVPacket)(dst), (*C.struct_AVPacket)(src)))
}

// Deprecated: Use AvPacketUnref() instead.
// AvFreePacket frees a packet.
func AvFreePacket(pkt *AvPacket) {
	C.av_free_packet((*C.struct_AVPacket)(pkt))
}

// AvPacketNewSideData allocates new information of a packet.
func AvPacketNewSideData(pkt *AvPacket, _type AvPacketSideDataType, size int32) *uint8 {
	return (*uint8)(C.av_packet_new_side_data((*C.struct_AVPacket)(pkt),
		(C.enum_AVPacketSideDataType)(_type), (C.int)(size)))
}

// AvPacketAddSideData wraps an existing array as a packet side data.
func AvPacketAddSideData(pkt *AvPacket, _type AvPacketSideDataType, data *uint8, size uint) int32 {
	return (int32)(C.av_packet_add_side_data((*C.struct_AVPacket)(pkt),
		(C.enum_AVPacketSideDataType)(_type), (*C.uint8_t)(data), (C.size_t)(size)))
}

// AvPacketShrinkSideData shrinks the already allocated side data buffer.
func AvPacketShrinkSideData(pkt *AvPacket, _type AvPacketSideDataType, size int32) int32 {
	return (int32)(C.av_packet_shrink_side_data((*C.struct_AVPacket)(pkt),
		(C.enum_AVPacketSideDataType)(_type), (C.int)(size)))
}

// AvPacketGetSideData gets side information from packet.
func AvPacketGetSideData(pkt *AvPacket, _type AvPacketSideDataType, size *int32) *uint8 {
	return (*uint8)(C.av_packet_get_side_data((*C.struct_AVPacket)(pkt),
		(C.enum_AVPacketSideDataType)(_type), (*C.int)(size)))
}

// Deprecated: No use
func AvPacketMergeSideData(pkt *AvPacket) int32 {
	return (int32)(C.av_packet_merge_side_data((*C.struct_AVPacket)(pkt)))
}

// Deprecated: No use
func AvPacketSplitSideData(pkt *AvPacket) int32 {
	return (int32)(C.av_packet_split_side_data((*C.struct_AVPacket)(pkt)))
}

// AvPacketPackDictionary packs a dictionary for use in side_data.
func AvPacketPackDictionary(dict *AvDictionary, size *int32) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(dict),
		(*C.int)(size)))
}

// AvPacketUnpackDictionary unpacks a dictionary from side_data.
func AvPacketUnpackDictionary(data *uint8, size int32, dict **AvDictionary) int32 {
	return (int32)(C.av_packet_unpack_dictionary((*C.uint8_t)(data), (C.int)(size),
		(**C.struct_AVDictionary)(unsafe.Pointer(dict))))
}

// AvPacketFreeSideData is a convenience function to free all the side data stored.
func AvPacketFreeSideData(pkt *AvPacket) {
	C.av_packet_free_side_data((*C.struct_AVPacket)(pkt))
}

// AvPacketRef setups a new reference to the data described by a given packet
func AvPacketRef(dst, src *AvPacket) int32 {
	return (int32)(C.av_packet_ref((*C.struct_AVPacket)(dst), (*C.struct_AVPacket)(src)))
}

// AvPacketUnref unreferences the buffer referenced by the packet and reset the
// remaining packet fields to their default values.
func AvPacketUnref(pkt *AvPacket) {
	C.av_packet_unref((*C.struct_AVPacket)(pkt))
}

// AvPacketMoveRef moves every field in src to dst and reset src.
func AvPacketMoveRef(dst, src *AvPacket) {
	C.av_packet_move_ref((*C.struct_AVPacket)(dst), (*C.struct_AVPacket)(src))
}

// AvPacketCopyProps copies only "properties" fields from src to dst.
func AvPacketCopyProps(dst, src *AvPacket) int32 {
	return (int32)(C.av_packet_copy_props((*C.struct_AVPacket)(dst), (*C.struct_AVPacket)(src)))
}

// AvPacketMakeRefcounted ensures the data described by a given packet is reference counted.
func AvPacketMakeRefcounted(pkt *AvPacket) {
	C.av_packet_make_refcounted((*C.struct_AVPacket)(pkt))
}

// AvPacketMakeWritable  creates a writable reference for the data described by a given packet,
// avoiding data copy if possible.
func AvPacketMakeWritable(pkt *AvPacket) {
	C.av_packet_make_writable((*C.struct_AVPacket)(pkt))
}

// AvPacketRescaleTs converts valid timing fields (timestamps / durations) in a packet from one
// timebase to another. Timestamps with unknown values (AV_NOPTS_VALUE) will be ignored.
func AvPacketRescaleTs(pkt *AvPacket, tbSrc, tbDst AvRational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(pkt),
		(C.struct_AVRational)(tbSrc), (C.struct_AVRational)(tbDst))
}
