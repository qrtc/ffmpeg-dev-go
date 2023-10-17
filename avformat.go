package ffmpeg

/*
#include <libavformat/avformat.h>
*/
import "C"
import (
	"unsafe"
)

// AvGetPacket allocates and reads the payload of a packet and initialize its
// fields with default values.
func AvGetPacket(s *AvIOContext, pkt *AvPacket, size int32) int32 {
	return (int32)(C.av_get_packet((*C.struct_AVIOContext)(s),
		(*C.struct_AVPacket)(pkt), (C.int)(size)))
}

// Read data and append it to the current content of the AVPacket.
// If pkt->size is 0 this is identical to av_get_packet.
func AvAppendPacket(s *AvIOContext, pkt *AvPacket, size int32) int32 {
	return (int32)(C.av_append_packet((*C.struct_AVIOContext)(s),
		(*C.struct_AVPacket)(pkt), (C.int)(size)))
}

// AvCodceTag
type AvCodecTag C.struct_AVCodecTag

// AvProbeData
type AvProbeData C.struct_AVProbeData

const (
	AVPROBE_SCORE_RETRY        = C.AVPROBE_SCORE_RETRY
	AVPROBE_SCORE_STREAM_RETRY = C.AVPROBE_SCORE_STREAM_RETRY
	AVPROBE_SCORE_EXTENSION    = C.AVPROBE_SCORE_EXTENSION
	AVPROBE_SCORE_MIME         = C.AVPROBE_SCORE_MIME
	AVPROBE_SCORE_MAX          = C.AVPROBE_SCORE_MAX
)

const (
	AVFMT_NOFILE        = C.AVFMT_NOFILE
	AVFMT_NEEDNUMBER    = C.AVFMT_NEEDNUMBER
	AVFMT_SHOW_IDS      = C.AVFMT_SHOW_IDS
	AVFMT_GLOBALHEADER  = C.AVFMT_GLOBALHEADER
	AVFMT_NOTIMESTAMPS  = C.AVFMT_NOTIMESTAMPS
	AVFMT_GENERIC_INDEX = C.AVFMT_GENERIC_INDEX
	AVFMT_TS_DISCONT    = C.AVFMT_TS_DISCONT
	AVFMT_VARIABLE_FPS  = C.AVFMT_VARIABLE_FPS
	AVFMT_NODIMENSIONS  = C.AVFMT_NODIMENSIONS
	AVFMT_NOSTREAMS     = C.AVFMT_NOSTREAMS
	AVFMT_NOBINSEARCH   = C.AVFMT_NOBINSEARCH
	AVFMT_NOGENSEARCH   = C.AVFMT_NOGENSEARCH
	AVFMT_NO_BYTE_SEEK  = C.AVFMT_NO_BYTE_SEEK
	AVFMT_ALLOW_FLUSH   = C.AVFMT_ALLOW_FLUSH
	AVFMT_TS_NONSTRICT  = C.AVFMT_TS_NONSTRICT
	AVFMT_TS_NEGATIVE   = C.AVFMT_TS_NEGATIVE
	AVFMT_SEEK_TO_PTS   = C.AVFMT_SEEK_TO_PTS
)

// AvOutputFormat
type AvOutputFormat C.struct_AVOutputFormat

// Custom: GetName gets `AVOutputFormat.name` value.
func (ofmt *AvOutputFormat) GetName() string {
	return C.GoString(ofmt.name)
}

// Custom: GetLongName gets `AVOutputFormat.long_name` value.
func (ofmt *AvOutputFormat) GetLongName() string {
	return C.GoString(ofmt.long_name)
}

// Custom: GetMimeType gets `AVOutputFormat.mime_type` value.
func (ofmt *AvOutputFormat) GetMimeType() string {
	return C.GoString(ofmt.mime_type)
}

// Custom: GetExtensions gets `AVOutputFormat.extensions` value.
func (ofmt *AvOutputFormat) GetExtensions() string {
	return C.GoString(ofmt.extensions)
}

// Custom: GetAudioCodec gets `AVOutputFormat.audio_codec` value.
func (ofmt *AvOutputFormat) GetAudioCodec() AvCodecID {
	return (AvCodecID)(ofmt.audio_codec)
}

// Custom: SetAudioCodec sets `AVOutputFormat.audio_codec` value.
func (ofmt *AvOutputFormat) SetAudioCodec(v AvCodecID) {
	ofmt.audio_codec = (C.enum_AVCodecID)(v)
}

// Custom: GetAudioCodecAddr gets `AVOutputFormat.audio_codec` address.
func (ofmt *AvOutputFormat) GetAudioCodecAddr() *AvCodecID {
	return (*AvCodecID)(unsafe.Pointer(&ofmt.audio_codec))
}

// Custom: GetVideoCodec gets `AVOutputFormat.video_codec` value.
func (ofmt *AvOutputFormat) GetVideoCodec() AvCodecID {
	return (AvCodecID)(ofmt.video_codec)
}

// Custom: SetVideoCodec sets `AVOutputFormat.video_codec` value.
func (ofmt *AvOutputFormat) SetVideoCodec(v AvCodecID) {
	ofmt.video_codec = (C.enum_AVCodecID)(v)
}

// Custom: GetVideoCodecAddr gets `AVOutputFormat.video_codec` address.
func (ofmt *AvOutputFormat) GetVideoCodecAddr() *AvCodecID {
	return (*AvCodecID)(unsafe.Pointer(&ofmt.video_codec))
}

// Custom: GetSubtitleCodec gets `AVOutputFormat.subtitle_codec` value.
func (ofmt *AvOutputFormat) GetSubtitleCodec() AvCodecID {
	return (AvCodecID)(ofmt.subtitle_codec)
}

// Custom: SetSubtitleCodec sets `AVOutputFormat.subtitle_codec` value.
func (ofmt *AvOutputFormat) SetSubtitleCodec(v AvCodecID) {
	ofmt.subtitle_codec = (C.enum_AVCodecID)(v)
}

// Custom: GetSubtitleCodecAddr gets `AVOutputFormat.subtitle_codec` address.
func (ofmt *AvOutputFormat) GetSubtitleCodecAddr() *AvCodecID {
	return (*AvCodecID)(unsafe.Pointer(&ofmt.subtitle_codec))
}

// Custom: GetFlags gets `AVOutputFormat.flags` value.
func (ofmt *AvOutputFormat) GetFlags() int32 {
	return (int32)(ofmt.flags)
}

// Custom: SetFlags sets `AVOutputFormat.flags` value.
func (ofmt *AvOutputFormat) SetFlags(v int32) {
	ofmt.flags = (C.int)(v)
}

// Custom: GetFlagsAddr gets `AVOutputFormat.flags` address.
func (ofmt *AvOutputFormat) GetFlagsAddr() *int32 {
	return (*int32)(&ofmt.flags)
}

// Custom: GetPrivClass gets `AVOutputFormat.priv_class` value.
func (ofmt *AvOutputFormat) GetPrivClass() *AvClass {
	return (*AvClass)(ofmt.priv_class)
}

// Custom: SetPrivClass sets `AVOutputFormat.priv_class` value.
func (ofmt *AvOutputFormat) SetPrivClass(v *AvClass) {
	ofmt.priv_class = (*C.struct_AVClass)(v)
}

// Custom: GetPrivClassAddr gets `AVOutputFormat.priv_class` address.
func (ofmt *AvOutputFormat) GetPrivClassAddr() **AvClass {
	return (**AvClass)(unsafe.Pointer(&ofmt.priv_class))
}

// AvInputFormat
type AvInputFormat C.struct_AVInputFormat

// Custom: GetName gets `AVInputFormat.name` value.
func (ifmt *AvInputFormat) GetName() string {
	return C.GoString(ifmt.name)
}

// Custom: GetLongName gets `AVInputFormat.long_name` value.
func (ifmt *AvInputFormat) GetLongName() string {
	return C.GoString(ifmt.long_name)
}

// Custom: GetFlags gets `AVInputFormat.flags` value.
func (ifmt *AvInputFormat) GetFlags() int32 {
	return (int32)(ifmt.flags)
}

// Custom: SetFlags sets `AVInputFormat.flags` value.
func (ifmt *AvInputFormat) SetFlags(v int32) {
	ifmt.flags = (C.int)(v)
}

// Custom: GetFlagsAddr gets `AVInputFormat.flags` address.
func (ifmt *AvInputFormat) GetFlagsAddr() *int32 {
	return (*int32)(&ifmt.flags)
}

// Custom: GetExtensions gets `AVInputFormat.extensions` value.
func (ifmt *AvInputFormat) GetExtensions() string {
	return C.GoString(ifmt.extensions)
}

// Custom: GetPrivClass gets `AVInputFormat.priv_class` value.
func (ifmt *AvInputFormat) GetPrivClass() *AvClass {
	return (*AvClass)(ifmt.priv_class)
}

// Custom: SetPrivClass sets `AVInputFormat.priv_class` value.
func (ifmt *AvInputFormat) SetPrivClass(v *AvClass) {
	ifmt.priv_class = (*C.struct_AVClass)(v)
}

// Custom: GetPrivClassAddr gets `AVInputFormat.priv_class` address.
func (ifmt *AvInputFormat) GetPrivClassAddr() **AvClass {
	return (**AvClass)(unsafe.Pointer(&ifmt.priv_class))
}

// Custom: GetMimeType gets `AVInputFormat.mime_type` value.
func (ifmt *AvInputFormat) GetMimeType() string {
	return C.GoString(ifmt.mime_type)
}

// AvStreamParseType
type AvStreamParseType = C.enum_AVStreamParseType

const (
	AVSTREAM_PARSE_NONE       = AvStreamParseType(C.AVSTREAM_PARSE_NONE)
	AVSTREAM_PARSE_FULL       = AvStreamParseType(C.AVSTREAM_PARSE_FULL)
	AVSTREAM_PARSE_HEADERS    = AvStreamParseType(C.AVSTREAM_PARSE_HEADERS)
	AVSTREAM_PARSE_TIMESTAMPS = AvStreamParseType(C.AVSTREAM_PARSE_TIMESTAMPS)
	AVSTREAM_PARSE_FULL_ONCE  = AvStreamParseType(C.AVSTREAM_PARSE_FULL_ONCE)
	AVSTREAM_PARSE_FULL_RAW   = AvStreamParseType(C.AVSTREAM_PARSE_FULL_RAW)
)

// AvIndexEntry
type AvIndexEntry C.struct_AVIndexEntry

const (
	AV_DISPOSITION_DEFAULT          = C.AV_DISPOSITION_DEFAULT
	AV_DISPOSITION_DUB              = C.AV_DISPOSITION_DUB
	AV_DISPOSITION_ORIGINAL         = C.AV_DISPOSITION_ORIGINAL
	AV_DISPOSITION_COMMENT          = C.AV_DISPOSITION_COMMENT
	AV_DISPOSITION_LYRICS           = C.AV_DISPOSITION_LYRICS
	AV_DISPOSITION_KARAOKE          = C.AV_DISPOSITION_KARAOKE
	AV_DISPOSITION_FORCED           = C.AV_DISPOSITION_FORCED
	AV_DISPOSITION_HEARING_IMPAIRED = C.AV_DISPOSITION_HEARING_IMPAIRED
	AV_DISPOSITION_VISUAL_IMPAIRED  = C.AV_DISPOSITION_VISUAL_IMPAIRED
	AV_DISPOSITION_CLEAN_EFFECTS    = C.AV_DISPOSITION_CLEAN_EFFECTS
	AV_DISPOSITION_ATTACHED_PIC     = C.AV_DISPOSITION_ATTACHED_PIC
	AV_DISPOSITION_TIMED_THUMBNAILS = C.AV_DISPOSITION_TIMED_THUMBNAILS
)

type AvStreamInternal C.struct_AVStreamInternal

const (
	AV_DISPOSITION_CAPTIONS     = C.AV_DISPOSITION_CAPTIONS
	AV_DISPOSITION_DESCRIPTIONS = C.AV_DISPOSITION_DESCRIPTIONS
	AV_DISPOSITION_METADATA     = C.AV_DISPOSITION_METADATA
	AV_DISPOSITION_DEPENDENT    = C.AV_DISPOSITION_DEPENDENT
	AV_DISPOSITION_STILL_IMAGE  = C.AV_DISPOSITION_STILL_IMAGE
)

const (
	AV_PTS_WRAP_IGNORE     = C.AV_PTS_WRAP_IGNORE
	AV_PTS_WRAP_ADD_OFFSET = C.AV_PTS_WRAP_ADD_OFFSET
	AV_PTS_WRAP_SUB_OFFSET = C.AV_PTS_WRAP_SUB_OFFSET
)

// AvStream
type AvStream C.struct_AVStream

// Custom: GetIndex gets `AVStream.index` value.
func (stm *AvStream) GetIndex() int32 {
	return (int32)(stm.index)
}

// Custom: SetIndex sets `AVStream.index` value.
func (stm *AvStream) SetIndex(v int32) {
	stm.index = (C.int)(v)
}

// Custom: GetIndexAddr gets `AVStream.index` address.
func (stm *AvStream) GetIndexAddr() *int32 {
	return (*int32)(&stm.index)
}

// Custom: GetId gets `AVStream.id` value.
func (stm *AvStream) GetId() int32 {
	return (int32)(stm.id)
}

// Custom: SetId sets `AVStream.id` value.
func (stm *AvStream) SetId(v int32) {
	stm.id = (C.int)(v)
}

// Custom: GetIdAddr gets `AVStream.id` address.
func (stm *AvStream) GetIdAddr() *int32 {
	return (*int32)(&stm.id)
}

// Custom: GetCodec gets `AVStream.codec` value.
func (stm *AvStream) GetCodec() *AvCodecContext {
	return (*AvCodecContext)(stm.codec)
}

// Custom: SetCodec sets `AVStream.codec` value.
func (stm *AvStream) SetCodec(v *AvCodecContext) {
	stm.codec = (*C.struct_AVCodecContext)(v)
}

// Custom: GetCodecAddr gets `AVStream.codec` address.
func (stm *AvStream) GetCodecAddr() **AvCodecContext {
	return (**AvCodecContext)(unsafe.Pointer(&stm.codec))
}

// Custom: GetPrivData gets `AVStream.priv_data` value.
func (stm *AvStream) GetPrivData() unsafe.Pointer {
	return stm.priv_data
}

// Custom: SetPrivData sets `AVStream.priv_data` value.
func (stm *AvStream) SetPrivData(v unsafe.Pointer) {
	stm.priv_data = v
}

// Custom: GetPrivDataAddr gets `AVStream.priv_data` address.
func (stm *AvStream) GetPrivDataAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&stm.priv_data)
}

// Custom: GetTimeBase gets `AVStream.time_base` value.
func (stm *AvStream) GetTimeBase() AvRational {
	return (AvRational)(stm.time_base)
}

// Custom: SetTimeBase sets `AVStream.time_base` value.
func (stm *AvStream) SetTimeBase(v AvRational) {
	stm.time_base = (C.struct_AVRational)(v)
}

// Custom: GetTimeBaseAddr gets `AVStream.time_base` address.
func (stm *AvStream) GetTimeBaseAddr() *AvRational {
	return (*AvRational)(&stm.time_base)
}

// Custom: GetStartTime gets `AVStream.start_time` value.
func (stm *AvStream) GetStartTime() int64 {
	return (int64)(stm.start_time)
}

// Custom: SetStartTime sets `AVStream.start_time` value.
func (stm *AvStream) SetStartTime(v int64) {
	stm.start_time = (C.int64_t)(v)
}

// Custom: GetStartTimeAddr gets `AVStream.start_time` address.
func (stm *AvStream) GetStartTimeAddr() *int64 {
	return (*int64)(&stm.start_time)
}

// Custom: GetDuration gets `AVStream.duration` value.
func (stm *AvStream) GetDuration() int64 {
	return (int64)(stm.duration)
}

// Custom: SetDuration sets `AVStream.duration` value.
func (stm *AvStream) SetDuration(v int64) {
	stm.duration = (C.int64_t)(v)
}

// Custom: GetDurationAddr gets `AVStream.duration` address.
func (stm *AvStream) GetDurationAddr() *int64 {
	return (*int64)(&stm.duration)
}

// Custom: GetNbFrames gets `AVStream.nb_frames` value.
func (stm *AvStream) GetNbFrames() int64 {
	return (int64)(stm.nb_frames)
}

// Custom: SetNbFrames sets `AVStream.nb_frames` value.
func (stm *AvStream) SetNbFrames(v int64) {
	stm.nb_frames = (C.int64_t)(v)
}

// Custom: GetNbFramesAddr gets `AVStream.nb_frames` address.
func (stm *AvStream) GetNbFramesAddr() *int64 {
	return (*int64)(&stm.nb_frames)
}

// Custom: GetDisposition gets `AVStream.disposition` value.
func (stm *AvStream) GetDisposition() int32 {
	return (int32)(stm.disposition)
}

// Custom: SetDisposition sets `AVStream.disposition` value.
func (stm *AvStream) SetDisposition(v int32) {
	stm.disposition = (C.int)(v)
}

// Custom: GetDispositionAddr gets `AVStream.disposition` address.
func (stm *AvStream) GetDispositionAddr() *int32 {
	return (*int32)(&stm.disposition)
}

// Custom: GetDiscard gets `AVStream.discard` value.
func (stm *AvStream) GetDiscard() AvDiscard {
	return (AvDiscard)(stm.discard)
}

// Custom: SetDiscard sets `AVStream.discard` value.
func (stm *AvStream) SetDiscard(v AvDiscard) {
	stm.discard = (C.enum_AVDiscard)(v)
}

// Custom: GetDiscardAddr gets `AVStream.discard` address.
func (stm *AvStream) GetDiscardAddr() *AvDiscard {
	return (*AvDiscard)(&stm.discard)
}

// Custom: GetSampleAspectRatio gets `AVStream.sample_aspect_ratio` value.
func (stm *AvStream) GetSampleAspectRatio() AvRational {
	return (AvRational)(stm.sample_aspect_ratio)
}

// Custom: SetSampleAspectRatio sets `AVStream.sample_aspect_ratio` value.
func (stm *AvStream) SetSampleAspectRatio(v AvRational) {
	stm.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// Custom: GetSampleAspectRatioAddr gets `AVStream.sample_aspect_ratio` address.
func (stm *AvStream) GetSampleAspectRatioAddr() *AvRational {
	return (*AvRational)(&stm.sample_aspect_ratio)
}

// Custom: GetMetadata gets `AVStream.metadata` value.
func (stm *AvStream) GetMetadata() *AvDictionary {
	return (*AvDictionary)(stm.metadata)
}

// Custom: SetMetadata sets `AVStream.metadata` value.
func (stm *AvStream) SetMetadata(v *AvDictionary) {
	stm.metadata = (*C.struct_AVDictionary)(v)
}

// Custom: GetMetadataAddr gets `AVStream.metadata` address.
func (stm *AvStream) GetMetadataAddr() **AvDictionary {
	return (**AvDictionary)(unsafe.Pointer(&stm.metadata))
}

// Custom: GetAvgFrameRate gets `AVStream.avg_frame_rate` value.
func (stm *AvStream) GetAvgFrameRate() AvRational {
	return (AvRational)(stm.avg_frame_rate)
}

// Custom: SetAvgFrameRate sets `AVStream.avg_frame_rate` value.
func (stm *AvStream) SetAvgFrameRate(v AvRational) {
	stm.avg_frame_rate = (C.struct_AVRational)(v)
}

// Custom: GetAvgFrameRateAddr gets `AVStream.avg_frame_rate` address.
func (stm *AvStream) GetAvgFrameRateAddr() *AvRational {
	return (*AvRational)(&stm.avg_frame_rate)
}

// Custom: GetAttachedPic gets `AVStream.attached_pic` value.
func (stm *AvStream) GetAttachedPic() AvPacket {
	return (AvPacket)(stm.attached_pic)
}

// Custom: SetAttachedPic sets `AVStream.attached_pic` value.
func (stm *AvStream) SetAttachedPic(v AvPacket) {
	stm.attached_pic = (C.struct_AVPacket)(v)
}

// Custom: GetAttachedPicAddr gets `AVStream.attached_pic` address.
func (stm *AvStream) GetAttachedPicAddr() *AvPacket {
	return (*AvPacket)(&stm.attached_pic)
}

// Custom: GetSideData gets `AVStream.side_data` value.
func (stm *AvStream) GetSideData() *AvPacketSideData {
	return (*AvPacketSideData)(stm.side_data)
}

// Custom: SetSideData sets `AVStream.side_data` value.
func (stm *AvStream) SetSideData(v *AvPacketSideData) {
	stm.side_data = (*C.struct_AVPacketSideData)(v)
}

// Custom: GetSideDataAddr gets `AVStream.side_data` address.
func (stm *AvStream) GetSideDataAddr() **AvPacketSideData {
	return (**AvPacketSideData)(unsafe.Pointer(&stm.side_data))
}

// Custom: GetNbSideData gets `AVStream.nb_side_data` value.
func (stm *AvStream) GetNbSideData() int32 {
	return (int32)(stm.nb_side_data)
}

// Custom: SetNbSideData sets `AVStream.nb_side_data` value.
func (stm *AvStream) SetNbSideData(v int32) {
	stm.nb_side_data = (C.int)(v)
}

// Custom: GetNbSideDataAddr gets `AVStream.nb_side_data` address.
func (stm *AvStream) GetNbSideDataAddr() *int32 {
	return (*int32)(&stm.nb_side_data)
}

// Custom: GetEventFlags gets `AVStream.event_flags` value.
func (stm *AvStream) GetEventFlags() int32 {
	return (int32)(stm.event_flags)
}

// Custom: SetEventFlags sets `AVStream.event_flags` value.
func (stm *AvStream) SetEventFlags(v int32) {
	stm.event_flags = (C.int)(v)
}

// Custom: GetEventFlagsAddr gets `AVStream.event_flags` address.
func (stm *AvStream) GetEventFlagsAddr() *int32 {
	return (*int32)(&stm.event_flags)
}

const (
	AVSTREAM_EVENT_FLAG_METADATA_UPDATED = int32(C.AVSTREAM_EVENT_FLAG_METADATA_UPDATED)
	AVSTREAM_EVENT_FLAG_NEW_PACKETS      = int32(C.AVSTREAM_EVENT_FLAG_NEW_PACKETS)
)

// Custom: GetRFrameRate gets `AVStream.r_frame_rate` value.
func (stm *AvStream) GetRFrameRate() AvRational {
	return (AvRational)(stm.r_frame_rate)
}

// Custom: SetRFrameRate sets `AVStream.r_frame_rate` value.
func (stm *AvStream) SetRFrameRate(v AvRational) {
	stm.r_frame_rate = (C.struct_AVRational)(v)
}

// Custom: GetRFrameRateAddr gets `AVStream.r_frame_rate` address.
func (stm *AvStream) GetRFrameRateAddr() *AvRational {
	return (*AvRational)(&stm.r_frame_rate)
}

// Custom: GetRecommendedEncoderConfiguration gets `AVStream.recommended_encoder_configuration` value.
func (stm *AvStream) GetRecommendedEncoderConfiguration() string {
	return C.GoString(stm.recommended_encoder_configuration)
}

// Custom: GetCodecpar gets `AVStream.codecpar` value.
func (stm *AvStream) GetCodecpar() *AvCodecParameters {
	return (*AvCodecParameters)(stm.codecpar)
}

// Custom: SetCodecpar sets `AVStream.codecpar` value.
func (stm *AvStream) SetCodecpar(v *AvCodecParameters) {
	stm.codecpar = (*C.struct_AVCodecParameters)(v)
}

// Custom: GetCodecparAddr gets `AVStream.codecpar` address.
func (stm *AvStream) GetCodecparAddr() **AvCodecParameters {
	return (**AvCodecParameters)(unsafe.Pointer(&stm.codecpar))
}

// Deprecated: No use
func AvStreamGetRFrameRate(s *AvStream) AvRational {
	return (AvRational)(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

// Deprecated: No use
func AvStreamSetRFrameRate(s *AvStream, r AvRational) {
	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), (C.struct_AVRational)(r))
}

// Deprecated: No use
func AvStreamGetRecommendedEncoderConfiguration(s *AvStream) *int8 {
	return (*int8)(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(s)))
}

// AvStreamGetParser
func AvStreamGetParser(s *AvStream) *AvCodecParserContext {
	return (*AvCodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

// AvStreamGetEndPts returns the pts of the last muxed packet + its duration.
// the retuned value is undefined when used with a demuxer.
func AvStreamGetEndPts(s *AvStream) int64 {
	return (int64)(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}

const (
	AV_PROGRAM_RUNNING = C.AV_PROGRAM_RUNNING
)

// AvProgram
type AvProgram C.struct_AVProgram

// Custom: GetId gets `AVProgram.id` value.
func (pgm *AvProgram) GetId() int32 {
	return (int32)(pgm.id)
}

// Custom: SetId sets `AVProgram.id` value.
func (pgm *AvProgram) SetId(v int32) {
	pgm.id = (C.int)(v)
}

// Custom: GetIdAddr gets `AVProgram.id` address.
func (pgm *AvProgram) GetIdAddr() *int32 {
	return (*int32)(&pgm.id)
}

// Custom: GetFlags gets `AVProgram.flags` value.
func (pgm *AvProgram) GetFlags() int32 {
	return (int32)(pgm.flags)
}

// Custom: SetFlags sets `AVProgram.flags` value.
func (pgm *AvProgram) SetFlags(v int32) {
	pgm.flags = (C.int)(v)
}

// Custom: GetFlagsAddr gets `AVProgram.flags` address.
func (pgm *AvProgram) GetFlagsAddr() *int32 {
	return (*int32)(&pgm.flags)
}

// Custom: GetDiscard gets `AVProgram.discard` value.
func (pgm *AvProgram) GetDiscard() AvDiscard {
	return (AvDiscard)(pgm.discard)
}

// Custom: SetDiscard sets `AVProgram.discard` value.
func (pgm *AvProgram) SetDiscard(v AvDiscard) {
	pgm.discard = (C.enum_AVDiscard)(v)
}

// Custom: GetDiscardAddr gets `AVProgram.discard` address.
func (pgm *AvProgram) GetDiscardAddr() *AvDiscard {
	return (*AvDiscard)(unsafe.Pointer(&pgm.discard))
}

// Custom: GetStreamIndex gets `AVProgram.stream_index` value.
func (pgm *AvProgram) GetStreamIndex() (v []uint32) {
	if pgm.stream_index == nil {
		return v
	}
	ptr := (*uint32)(unsafe.Pointer(pgm.stream_index))
	for i := 0; i < int(pgm.nb_stream_indexes); i++ {
		v = append(v, *ptr)
		ptr = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: SetStreamIndex sets `AVProgram.stream_index` value.
func (pgm *AvProgram) SetStreamIndex(v *uint32) {
	pgm.stream_index = (*C.uint)(v)
}

// Custom: GetStreamIndexAddr gets `AVProgram.stream_index` address.
func (pgm *AvProgram) GetStreamIndexAddr() **uint32 {
	return (**uint32)(unsafe.Pointer(&pgm.stream_index))
}

// Custom: GetNbStreamIndexes gets `AVProgram.nb_stream_indexes` value.
func (pgm *AvProgram) GetNbStreamIndexes() uint32 {
	return (uint32)(pgm.nb_stream_indexes)
}

// Custom: SetNbStreamIndexes sets `AVProgram.nb_stream_indexes` value.
func (pgm *AvProgram) SetNbStreamIndexes(v uint32) {
	pgm.nb_stream_indexes = (C.uint)(v)
}

// Custom: GetNbStreamIndexesAddr gets `AVProgram.nb_stream_indexes` address.
func (pgm *AvProgram) GetNbStreamIndexesAddr() *uint32 {
	return (*uint32)(&pgm.nb_stream_indexes)
}

// Custom: GetMetadata gets `AVProgram.metadata` value.
func (pgm *AvProgram) GetMetadata() *AvDictionary {
	return (*AvDictionary)(pgm.metadata)
}

// Custom: SetMetadata sets `AVProgram.metadata` value.
func (pgm *AvProgram) SetMetadata(v *AvDictionary) {
	pgm.metadata = (*C.struct_AVDictionary)(v)
}

// Custom: GetMetadataAddr gets `AVProgram.metadata` address.
func (pgm *AvProgram) GetMetadataAddr() **AvDictionary {
	return (**AvDictionary)(unsafe.Pointer(&pgm.metadata))
}

// Custom: GetProgramNum gets `AVProgram.program_num` value.
func (pgm *AvProgram) GetProgramNum() int32 {
	return (int32)(pgm.program_num)
}

// Custom: SetProgramNum sets `AVProgram.program_num` value.
func (pgm *AvProgram) SetProgramNum(v int32) {
	pgm.program_num = (C.int)(v)
}

// Custom: GetProgramNumAddr gets `AVProgram.program_num` address.
func (pgm *AvProgram) GetProgramNumAddr() *int32 {
	return (*int32)(&pgm.program_num)
}

// Custom: GetPmtPid gets `AVProgram.pmt_pid` value.
func (pgm *AvProgram) GetPmtPid() int32 {
	return (int32)(pgm.pmt_pid)
}

// Custom: SetPmtPid sets `AVProgram.pmt_pid` value.
func (pgm *AvProgram) SetPmtPid(v int32) {
	pgm.pmt_pid = (C.int)(v)
}

// Custom: GetPmtPidAddr gets `AVProgram.pmt_pid` address.
func (pgm *AvProgram) GetPmtPidAddr() *int32 {
	return (*int32)(&pgm.pmt_pid)
}

// Custom: GetPcrPid gets `AVProgram.pcr_pid` value.
func (pgm *AvProgram) GetPcrPid() int32 {
	return (int32)(pgm.pcr_pid)
}

// Custom: SetPcrPid sets `AVProgram.pcr_pid` value.
func (pgm *AvProgram) SetPcrPid(v int32) {
	pgm.pcr_pid = (C.int)(v)
}

// Custom: GetPcrPidAddr gets `AVProgram.pcr_pid` address.
func (pgm *AvProgram) GetPcrPidAddr() *int32 {
	return (*int32)(&pgm.pcr_pid)
}

// Custom: GetPmtVersion gets `AVProgram.pmt_version` value.
func (pgm *AvProgram) GetPmtVersion() int32 {
	return (int32)(pgm.pmt_version)
}

// Custom: SetPmtVersion sets `AVProgram.pmt_version` value.
func (pgm *AvProgram) SetPmtVersion(v int32) {
	pgm.pmt_version = (C.int)(v)
}

// Custom: GetPmtVersionAddr gets `AVProgram.pmt_version` address.
func (pgm *AvProgram) GetPmtVersionAddr() *int32 {
	return (*int32)(&pgm.pmt_version)
}

const (
	AVFMTCTX_NOHEADER   = C.AVFMTCTX_NOHEADER
	AVFMTCTX_UNSEEKABLE = C.AVFMTCTX_UNSEEKABLE
)

// AvChapter
type AvChapter = C.struct_AVChapter

// Custom: GetId gets `AVChapter.id` value.
func (cpt *AvChapter) GetId() int32 {
	return (int32)(cpt.id)
}

// Custom: SetId sets `AVChapter.id` value.
func (cpt *AvChapter) SetId(v int32) {
	cpt.id = (C.int)(v)
}

// Custom: GetIdAddr gets `AVChapter.id` address.
func (cpt *AvChapter) GetIdAddr() *int32 {
	return (*int32)(&cpt.id)
}

// Custom: GetTimeBase gets `AVChapter.time_base` value.
func (cpt *AvChapter) GetTimeBase() AvRational {
	return (AvRational)(cpt.time_base)
}

// Custom: SetTimeBase sets `AVChapter.time_base` value.
func (cpt *AvChapter) SetTimeBase(v AvRational) {
	cpt.time_base = (C.struct_AVRational)(v)
}

// Custom: GetTimeBaseAddr gets `AVChapter.time_base` address.
func (cpt *AvChapter) GetTimeBaseAddr() *AvRational {
	return (*AvRational)(&cpt.time_base)
}

// Custom: GetStart gets `AVChapter.start` value.
func (cpt *AvChapter) GetStart() int64 {
	return (int64)(cpt.start)
}

// Custom: SetStart sets `AVChapter.start` value.
func (cpt *AvChapter) SetStart(v int64) {
	cpt.start = (C.int64_t)(v)
}

// Custom: GetStartAddr gets `AVChapter.start` address.
func (cpt *AvChapter) GetStartAddr() *int64 {
	return (*int64)(&cpt.start)
}

// Custom: GetEnd gets `AVChapter.end` value.
func (cpt *AvChapter) GetEnd() int64 {
	return (int64)(cpt.end)
}

// Custom: SetEnd sets `AVChapter.end` value.
func (cpt *AvChapter) SetEnd(v int64) {
	cpt.end = (C.int64_t)(v)
}

// Custom: GetEndAddr gets `AVChapter.end` address.
func (cpt *AvChapter) GetEndAddr() *int64 {
	return (*int64)(&cpt.end)
}

// Custom: GetMetadata gets `AVChapter.metadata` value.
func (cpt *AvChapter) GetMetadata() *AvDictionary {
	return (*AvDictionary)(cpt.metadata)
}

// Custom: SetMetadata sets `AVChapter.metadata` value.
func (cpt *AvChapter) SetMetadata(v *AvDictionary) {
	cpt.metadata = (*C.struct_AVDictionary)(v)
}

// Custom: GetMetadataAddr gets `AVChapter.metadata` address.
func (cpt *AvChapter) GetMetadataAddr() **AvDictionary {
	return (**AvDictionary)(unsafe.Pointer(&cpt.metadata))
}

type AvFormatControlMessageFunc C.av_format_control_message

type AvOpenCallbackFunc C.AVOpenCallback

// AvDurationEstimationMethod
type AvDurationEstimationMethod = C.enum_AVDurationEstimationMethod

const (
	AVFMT_DURATION_FROM_PTS     = AvDurationEstimationMethod(C.AVFMT_DURATION_FROM_PTS)
	AVFMT_DURATION_FROM_STREAM  = AvDurationEstimationMethod(C.AVFMT_DURATION_FROM_STREAM)
	AVFMT_DURATION_FROM_BITRATE = AvDurationEstimationMethod(C.AVFMT_DURATION_FROM_BITRATE)
)

// AvFormatInternal
type AvFormatInternal C.struct_AVFormatInternal

// Format I/O context.
type AvFormatContext C.struct_AVFormatContext

// Custom: GetAvClass gets `AVFormatContext.av_class` value.
func (s *AvFormatContext) GetAvClass() *AvClass {
	return (*AvClass)(s.av_class)
}

// Custom: SetAvClass sets `AVFormatContext.av_class` value.
func (s *AvFormatContext) SetAvClass(v *AvClass) {
	s.av_class = (*C.struct_AVClass)(v)
}

// Custom: GetAvClassAddr gets `AVFormatContext.av_class` address.
func (s *AvFormatContext) GetAvClassAddr() **AvClass {
	return (**AvClass)(unsafe.Pointer(&s.av_class))
}

// Custom: GetIformat gets `AVFormatContext.iformat` value.
func (s *AvFormatContext) GetIformat() *AvInputFormat {
	return (*AvInputFormat)(s.iformat)
}

// Custom: SetIformat sets `AVFormatContext.iformat` value.
func (s *AvFormatContext) SetIformat(v *AvInputFormat) {
	s.iformat = (*C.struct_AVInputFormat)(v)
}

// Custom: GetIformatAddr gets `AVFormatContext.iformat` address.
func (s *AvFormatContext) GetIformatAddr() **AvInputFormat {
	return (**AvInputFormat)(unsafe.Pointer(&s.iformat))
}

// Custom: GetOformat gets `AVFormatContext.oformat` value.
func (s *AvFormatContext) GetOformat() *AvOutputFormat {
	return (*AvOutputFormat)(s.oformat)
}

// Custom: SetOformat sets `AVFormatContext.oformat` value.
func (s *AvFormatContext) SetOformat(v *AvOutputFormat) {
	s.oformat = (*C.struct_AVOutputFormat)(v)
}

// Custom: GetOformatAddr gets `AVFormatContext.oformat` address.
func (s *AvFormatContext) GetOformatAddr() **AvOutputFormat {
	return (**AvOutputFormat)(unsafe.Pointer(&s.oformat))
}

// Custom: GetPrivData gets `AVFormatContext.priv_data` value.
func (s *AvFormatContext) GetPrivData() unsafe.Pointer {
	return s.priv_data
}

// Custom: SetPrivData sets `AVFormatContext.priv_data` value.
func (s *AvFormatContext) SetPrivData(v unsafe.Pointer) {
	s.priv_data = v
}

// Custom: GetPrivDataAddr gets `AVFormatContext.priv_data` address.
func (s *AvFormatContext) GetPrivDataAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&s.priv_data)
}

// Custom: GetPb gets `AVFormatContext.pb` value.
func (s *AvFormatContext) GetPb() *AvIOContext {
	return (*AvIOContext)(s.pb)
}

// Custom: SetPb sets `AVFormatContext.pb` value.
func (s *AvFormatContext) SetPb(v *AvIOContext) {
	s.pb = (*C.struct_AVIOContext)(v)
}

// Custom: GetPbAddr gets `AVFormatContext.pb` address.
func (s *AvFormatContext) GetPbAddr() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&s.pb))
}

// Custom: GetCtxFlags gets `AVFormatContext.ctx_flags` value.
func (s *AvFormatContext) GetCtxFlags() int32 {
	return (int32)(s.ctx_flags)
}

// Custom: SetCtxFlags sets `AVFormatContext.ctx_flags` value.
func (s *AvFormatContext) SetCtxFlags(v int32) {
	s.ctx_flags = (C.int)(v)
}

// Custom: GetCtxFlagsAddr gets `AVFormatContext.ctx_flags` address.
func (s *AvFormatContext) GetCtxFlagsAddr() *int32 {
	return (*int32)(&s.ctx_flags)
}

// Custom: GetNbStreams gets `AVFormatContext.nb_streams` value.
func (s *AvFormatContext) GetNbStreams() uint32 {
	return (uint32)(s.nb_streams)
}

// Custom: SetNbStreams sets `AVFormatContext.nb_streams` value.
func (s *AvFormatContext) SetNbStreams(v uint32) {
	s.nb_streams = (C.uint)(v)
}

// Custom: GetNbStreamsAddr gets `AVFormatContext.nb_streams` address.
func (s *AvFormatContext) GetNbStreamsAddr() *uint32 {
	return (*uint32)(&s.nb_streams)
}

// Custom: GetStreams gets `AVFormatContext.streams` value.
func (s *AvFormatContext) GetStreams() (v []*AvStream) {
	if s.streams == nil {
		return v
	}
	ptr := (**AvStream)(unsafe.Pointer(s.streams))
	for i := 0; i < int(s.nb_streams); i++ {
		v = append(v, *ptr)
		ptr = (**AvStream)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: SetStreams sets `AVFormatContext.streams` value.
func (s *AvFormatContext) SetStreams(v **AvStream) {
	s.streams = (**C.struct_AVStream)(unsafe.Pointer(v))
}

// Custom: GetStreamsAddr gets `AVFormatContext.streams` address.
func (s *AvFormatContext) GetStreamsAddr() ***AvStream {
	return (***AvStream)(unsafe.Pointer(&s.streams))
}

// Custom: GetFilename gets `AVFormatContext.filename` value.
func (s *AvFormatContext) GetFilename() string {
	return C.GoString((*C.char)(&s.filename[0]))
}

// Custom: GetUrl gets `AVFormatContext.url` value.
func (s *AvFormatContext) GetUrl() string {
	return C.GoString(s.url)
}

// Custom: SetUrl sets `AVFormatContext.url` value.
func (s *AvFormatContext) SetUrl(v string) {
	vPtr, _ := StringCasting(v)
	if s.url != nil {
		C.free(unsafe.Pointer(s.url))
	}
	s.url = (*C.char)(vPtr)
}

// Custom: GetStartTime gets `AVFormatContext.start_time` value.
func (s *AvFormatContext) GetStartTime() int64 {
	return (int64)(s.start_time)
}

// Custom: SetStartTime sets `AVFormatContext.start_time` value.
func (s *AvFormatContext) SetStartTime(v int64) {
	s.start_time = (C.int64_t)(v)
}

// Custom: GetStartTimeAddr gets `AVFormatContext.start_time` address.
func (s *AvFormatContext) GetStartTimeAddr() *int64 {
	return (*int64)(&s.start_time)
}

// Custom: GetDuration gets `AVFormatContext.duration` value.
func (s *AvFormatContext) GetDuration() int64 {
	return (int64)(s.duration)
}

// Custom: SetDuration sets `AVFormatContext.duration` value.
func (s *AvFormatContext) SetDuration(v int64) {
	s.duration = (C.int64_t)(v)
}

// Custom: GetDurationAddr gets `AVFormatContext.duration` address.
func (s *AvFormatContext) GetDurationAddr() *int64 {
	return (*int64)(&s.duration)
}

// Custom: GetBitRate gets `AVFormatContext.bit_rate` value.
func (s *AvFormatContext) GetBitRate() int64 {
	return (int64)(s.bit_rate)
}

// Custom: SetBitRate sets `AVFormatContext.bit_rate` value.
func (s *AvFormatContext) SetBitRate(v int64) {
	s.bit_rate = (C.int64_t)(v)
}

// Custom: GetBitRateAddr gets `AVFormatContext.bit_rate` address.
func (s *AvFormatContext) GetBitRateAddr() *int64 {
	return (*int64)(&s.bit_rate)
}

// Custom: GetPacketSize gets `AVFormatContext.packet_size` value.
func (s *AvFormatContext) GetPacketSize() uint32 {
	return (uint32)(s.packet_size)
}

// Custom: SetPacketSize sets `AVFormatContext.packet_size` value.
func (s *AvFormatContext) SetPacketSize(v uint32) {
	s.packet_size = (C.uint)(v)
}

// Custom: GetPacketSizeAddr gets `AVFormatContext.packet_size` address.
func (s *AvFormatContext) GetPacketSizeAddr() *uint32 {
	return (*uint32)(&s.packet_size)
}

// Custom: GetMaxDelay gets `AVFormatContext.max_delay` value.
func (s *AvFormatContext) GetMaxDelay() int32 {
	return (int32)(s.max_delay)
}

// Custom: SetMaxDelay sets `AVFormatContext.max_delay` value.
func (s *AvFormatContext) SetMaxDelay(v int32) {
	s.max_delay = (C.int)(v)
}

// Custom: GetMaxDelayAddr gets `AVFormatContext.max_delay` address.
func (s *AvFormatContext) GetMaxDelayAddr() *int32 {
	return (*int32)(&s.max_delay)
}

// Custom: GetFlags gets `AVFormatContext.flags` value.
func (s *AvFormatContext) GetFlags() int32 {
	return (int32)(s.flags)
}

// Custom: SetFlags sets `AVFormatContext.flags` value.
func (s *AvFormatContext) SetFlags(v int32) {
	s.flags = (C.int)(v)
}

// Custom: GetFlagsAddr gets `AVFormatContext.flags` address.
func (s *AvFormatContext) GetFlagsAddr() *int32 {
	return (*int32)(&s.flags)
}

// Custom: GetProbesize gets `AVFormatContext.probesize` value.
func (s *AvFormatContext) GetProbesize() int64 {
	return (int64)(s.probesize)
}

// Custom: SetProbesize sets `AVFormatContext.probesize` value.
func (s *AvFormatContext) SetProbesize(v int64) {
	s.probesize = (C.int64_t)(v)
}

// Custom: GetProbesizeAddr gets `AVFormatContext.probesize` address.
func (s *AvFormatContext) GetProbesizeAddr() *int64 {
	return (*int64)(&s.probesize)
}

// Custom: GetMaxAnalyzeDuration gets `AVFormatContext.max_analyze_duration` value.
func (s *AvFormatContext) GetMaxAnalyzeDuration() int64 {
	return (int64)(s.max_analyze_duration)
}

// Custom: SetMaxAnalyzeDuration sets `AVFormatContext.max_analyze_duration` value.
func (s *AvFormatContext) SetMaxAnalyzeDuration(v int64) {
	s.max_analyze_duration = (C.int64_t)(v)
}

// Custom: GetMaxAnalyzeDurationAddr gets `AVFormatContext.max_analyze_duration` address.
func (s *AvFormatContext) GetMaxAnalyzeDurationAddr() *int64 {
	return (*int64)(&s.max_analyze_duration)
}

// Custom: GetKey gets `AVFormatContext.key` value.
func (s *AvFormatContext) GetKey() *uint8 {
	return (*uint8)(s.key)
}

// Custom: SetKey sets `AVFormatContext.key` value.
func (s *AvFormatContext) SetKey(v *uint8) {
	s.key = (*C.uint8_t)(v)
}

// Custom: GetKeyAddr gets `AVFormatContext.key` address.
func (s *AvFormatContext) GetKeyAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&s.key))
}

// Custom: GetKeylen gets `AVFormatContext.keylen` value.
func (s *AvFormatContext) GetKeylen() int32 {
	return (int32)(s.keylen)
}

// Custom: SetKeylen sets `AVFormatContext.keylen` value.
func (s *AvFormatContext) SetKeylen(v int32) {
	s.keylen = (C.int)(v)
}

// Custom: GetKeylenAddr gets `AVFormatContext.keylen` address.
func (s *AvFormatContext) GetKeylenAddr() *int32 {
	return (*int32)(&s.keylen)
}

// Custom: GetNbPrograms gets `AVFormatContext.nb_programs` value.
func (s *AvFormatContext) GetNbPrograms() uint32 {
	return (uint32)(s.nb_programs)
}

// Custom: SetNbPrograms sets `AVFormatContext.nb_programs` value.
func (s *AvFormatContext) SetNbPrograms(v uint32) {
	s.nb_programs = (C.uint)(v)
}

// Custom: GetNbProgramsAddr gets `AVFormatContext.nb_programs` address.
func (s *AvFormatContext) GetNbProgramsAddr() *uint32 {
	return (*uint32)(&s.nb_programs)
}

// Custom: GetPrograms gets `AVFormatContext.programs` value.
func (s *AvFormatContext) GetPrograms() (v []*AvProgram) {
	if s.programs == nil {
		return v
	}
	ptr := (**AvProgram)(unsafe.Pointer(s.programs))
	for i := 0; i < int(s.nb_programs); i++ {
		v = append(v, *ptr)
		ptr = (**AvProgram)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: SetPrograms sets `AVFormatContext.AvProgram` value.
func (s *AvFormatContext) SetPrograms(v **AvProgram) {
	s.programs = (**C.struct_AVProgram)(unsafe.Pointer(v))
}

// Custom: GetProgramsAddr gets `AVFormatContext.AvProgram` address.
func (s *AvFormatContext) GetProgramsAddr() ***AvProgram {
	return (***AvProgram)(unsafe.Pointer(&s.programs))
}

// Custom: GetVideoCodecId gets `AVFormatContext.video_codec_id` value.
func (s *AvFormatContext) GetVideoCodecId() AvCodecID {
	return (AvCodecID)(s.video_codec_id)
}

// Custom: SetVideoCodecId sets `AVFormatContext.video_codec_id` value.
func (s *AvFormatContext) SetVideoCodecId(v AvCodecID) {
	s.video_codec_id = (C.enum_AVCodecID)(v)
}

// Custom: GetVideoCodecIdAddr gets `AVFormatContext.video_codec_id` address.
func (s *AvFormatContext) GetVideoCodecIdAddr() *AvCodecID {
	return (*AvCodecID)(unsafe.Pointer(&s.video_codec_id))
}

// Custom: GetAudioCodecId gets `AVFormatContext.audio_codec_id` value.
func (s *AvFormatContext) GetAudioCodecId() AvCodecID {
	return (AvCodecID)(s.audio_codec_id)
}

// Custom: SetAudioCodecId sets `AVFormatContext.audio_codec_id` value.
func (s *AvFormatContext) SetAudioCodecId(v AvCodecID) {
	s.audio_codec_id = (C.enum_AVCodecID)(v)
}

// Custom: GetAudioCodecIdAddr gets `AVFormatContext.audio_codec_id` address.
func (s *AvFormatContext) GetAudioCodecIdAddr() *AvCodecID {
	return (*AvCodecID)(unsafe.Pointer(&s.audio_codec_id))
}

// Custom: GetSubtitleCodecId gets `AVFormatContext.subtitle_codec_id` value.
func (s *AvFormatContext) GetSubtitleCodecId() AvCodecID {
	return (AvCodecID)(s.subtitle_codec_id)
}

// Custom: SetSubtitleCodecId sets `AVFormatContext.subtitle_codec_id` value.
func (s *AvFormatContext) SetSubtitleCodecId(v AvCodecID) {
	s.subtitle_codec_id = (C.enum_AVCodecID)(v)
}

// Custom: GetSubtitleCodecIdAddr gets `AVFormatContext.subtitle_codec_id` address.
func (s *AvFormatContext) GetSubtitleCodecIdAddr() *AvCodecID {
	return (*AvCodecID)(unsafe.Pointer(&s.subtitle_codec_id))
}

// Custom: GetMaxIndexSize gets `AVFormatContext.max_index_size` value.
func (s *AvFormatContext) GetMaxIndexSize() uint32 {
	return (uint32)(s.max_index_size)
}

// Custom: SetMaxIndexSize sets `AVFormatContext.max_index_size` value.
func (s *AvFormatContext) SetMaxIndexSize(v uint32) {
	s.max_index_size = (C.uint)(v)
}

// Custom: GetMaxIndexSizeAddr gets `AVFormatContext.max_index_size` address.
func (s *AvFormatContext) GetMaxIndexSizeAddr() *uint32 {
	return (*uint32)(&s.max_index_size)
}

// Custom: GetMaxPictureBuffer gets `AVFormatContext.max_picture_buffer` value.
func (s *AvFormatContext) GetMaxPictureBuffer() uint32 {
	return (uint32)(s.max_picture_buffer)
}

// Custom: SetMaxPictureBuffer sets `AVFormatContext.max_picture_buffer` value.
func (s *AvFormatContext) SetMaxPictureBuffer(v uint32) {
	s.max_picture_buffer = (C.uint)(v)
}

// Custom: GetMaxPictureBufferAddr gets `AVFormatContext.max_picture_buffer` address.
func (s *AvFormatContext) GetMaxPictureBufferAddr() *uint32 {
	return (*uint32)(&s.max_picture_buffer)
}

// Custom: GetNbChapters gets `AVFormatContext.nb_chapters` value.
func (s *AvFormatContext) GetNbChapters() uint32 {
	return (uint32)(s.nb_chapters)
}

// Custom: SetNbChapters sets `AVFormatContext.nb_chapters` value.
func (s *AvFormatContext) SetNbChapters(v uint32) {
	s.nb_chapters = (C.uint)(v)
}

// Custom: GetNbChaptersAddr gets `AVFormatContext.nb_chapters` address.
func (s *AvFormatContext) GetNbChaptersAddr() *uint32 {
	return (*uint32)(&s.nb_chapters)
}

// Custom: GetChapters gets `AVFormatContext.chapters` value.
func (s *AvFormatContext) GetChapters() (v []*AvChapter) {
	if s.chapters == nil {
		return v
	}
	ptr := (**AvChapter)(unsafe.Pointer(s.chapters))
	for i := 0; i < int(s.nb_chapters); i++ {
		v = append(v, *ptr)
		ptr = (**AvChapter)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: SetChapters sets `AVFormatContext.chapters` value.
func (s *AvFormatContext) SetChapters(v **AvChapter) {
	s.chapters = (**C.struct_AVChapter)(v)
}

// Custom: GetChaptersAddr gets `AVFormatContext.chapters` address.
func (s *AvFormatContext) GetChaptersAddr() ***AvChapter {
	return (***AvChapter)(&s.chapters)
}

// Custom: GetMetadata gets `AVFormatContext.metadata` value.
func (s *AvFormatContext) GetMetadata() *AvDictionary {
	return (*AvDictionary)(s.metadata)
}

// Custom: SetMetadata sets `AVFormatContext.metadata` value.
func (s *AvFormatContext) SetMetadata(v *AvDictionary) {
	s.metadata = (*C.struct_AVDictionary)(v)
}

// Custom: GetMetadataAddr gets `AVFormatContext.metadata` address.
func (s *AvFormatContext) GetMetadataAddr() **AvDictionary {
	return (**AvDictionary)(unsafe.Pointer(&s.metadata))
}

// Custom: GetStartTimeRealtime gets `AVFormatContext.start_time_realtime` value.
func (s *AvFormatContext) GetStartTimeRealtime() int64 {
	return (int64)(s.start_time_realtime)
}

// Custom: SetStartTimeRealtime sets `AVFormatContext.start_time_realtime` value.
func (s *AvFormatContext) SetStartTimeRealtime(v int64) {
	s.start_time_realtime = (C.int64_t)(v)
}

// Custom: GetStartTimeRealtimeAddr gets `AVFormatContext.start_time_realtime` address.
func (s *AvFormatContext) GetStartTimeRealtimeAddr() *int64 {
	return (*int64)(&s.start_time_realtime)
}

// Custom: GetFpsProbeSize gets `AVFormatContext.fps_probe_size` value.
func (s *AvFormatContext) GetFpsProbeSize() int32 {
	return (int32)(s.fps_probe_size)
}

// Custom: SetFpsProbeSize sets `AVFormatContext.fps_probe_size` value.
func (s *AvFormatContext) SetFpsProbeSize(v int32) {
	s.fps_probe_size = (C.int)(v)
}

// Custom: GetFpsProbeSizeAddr gets `AVFormatContext.fps_probe_size` address.
func (s *AvFormatContext) GetFpsProbeSizeAddr() *int32 {
	return (*int32)(&s.fps_probe_size)
}

// Custom: GetErrorRecognition gets `AVFormatContext.error_recognition` value.
func (s *AvFormatContext) GetErrorRecognition() int32 {
	return (int32)(s.error_recognition)
}

// Custom: SetErrorRecognition sets `AVFormatContext.error_recognition` value.
func (s *AvFormatContext) SetErrorRecognition(v int32) {
	s.error_recognition = (C.int)(v)
}

// Custom: GetErrorRecognitionAddr gets `AVFormatContext.error_recognition` address.
func (s *AvFormatContext) GetErrorRecognitionAddr() *int32 {
	return (*int32)(&s.error_recognition)
}

// Custom: GetInterruptCallback gets `AVFormatContext.interrupt_callback` value.
func (s *AvFormatContext) GetInterruptCallback() AvIOInterruptCB {
	return (AvIOInterruptCB)(s.interrupt_callback)
}

// Custom: SetInterruptCallback sets `AVFormatContext.interrupt_callback` value.
func (s *AvFormatContext) SetInterruptCallback(v AvIOInterruptCB) {
	s.interrupt_callback = (C.AVIOInterruptCB)(v)
}

// Custom: GetInterruptCallbackAddr gets `AVFormatContext.interrupt_callback` address.
func (s *AvFormatContext) GetInterruptCallbackAddr() *AvIOInterruptCB {
	return (*AvIOInterruptCB)(&s.interrupt_callback)
}

// Custom: GetDebug gets `AVFormatContext.debug` value.
func (s *AvFormatContext) GetDebug() int32 {
	return (int32)(s.debug)
}

// Custom: SetDebug sets `AVFormatContext.debug` value.
func (s *AvFormatContext) SetDebug(v int32) {
	s.debug = (C.int)(v)
}

// Custom: GetDebugAddr gets `AVFormatContext.debug` address.
func (s *AvFormatContext) GetDebugAddr() *int32 {
	return (*int32)(&s.debug)
}

// Custom: GetMaxInterleaveDelta gets `AVFormatContext.max_interleave_delta` value.
func (s *AvFormatContext) GetMaxInterleaveDelta() int64 {
	return (int64)(s.max_interleave_delta)
}

// Custom: SetMaxInterleaveDelta sets `AVFormatContext.max_interleave_delta` value.
func (s *AvFormatContext) SetMaxInterleaveDelta(v int64) {
	s.max_interleave_delta = (C.int64_t)(v)
}

// Custom: GetMaxInterleaveDeltaAddr gets `AVFormatContext.max_interleave_delta` address.
func (s *AvFormatContext) GetMaxInterleaveDeltaAddr() *int64 {
	return (*int64)(&s.max_interleave_delta)
}

// Custom: GetStrictStdCompliance gets `AVFormatContext.strict_std_compliance` value.
func (s *AvFormatContext) GetStrictStdCompliance() int32 {
	return (int32)(s.strict_std_compliance)
}

// Custom: SetStrictStdCompliance sets `AVFormatContext.strict_std_compliance` value.
func (s *AvFormatContext) SetStrictStdCompliance(v int32) {
	s.strict_std_compliance = (C.int)(v)
}

// Custom: GetStrictStdComplianceAddr gets `AVFormatContext.strict_std_compliance` address.
func (s *AvFormatContext) GetStrictStdComplianceAddr() *int32 {
	return (*int32)(&s.strict_std_compliance)
}

// Custom: GetEventFlags gets `AVFormatContext.event_flags` value.
func (s *AvFormatContext) GetEventFlags() int32 {
	return (int32)(s.event_flags)
}

// Custom: SetEventFlags sets `AVFormatContext.event_flags` value.
func (s *AvFormatContext) SetEventFlags(v int32) {
	s.event_flags = (C.int)(v)
}

// Custom: GetEventFlagsAddr gets `AVFormatContext.event_flags` address.
func (s *AvFormatContext) GetEventFlagsAddr() *int32 {
	return (*int32)(&s.event_flags)
}

// Custom: GetMaxTsProbe gets `AVFormatContext.max_ts_probe` value.
func (s *AvFormatContext) GetMaxTsProbe() int32 {
	return (int32)(s.max_ts_probe)
}

// Custom: SetMaxTsProbe sets `AVFormatContext.max_ts_probe` value.
func (s *AvFormatContext) SetMaxTsProbe(v int32) {
	s.max_ts_probe = (C.int)(v)
}

// Custom: GetMaxTsProbeAddr gets `AVFormatContext.max_ts_probe` address.
func (s *AvFormatContext) GetMaxTsProbeAddr() *int32 {
	return (*int32)(&s.max_ts_probe)
}

// Custom: GetAvoidNegativeTs gets `AVFormatContext.avoid_negative_ts` value.
func (s *AvFormatContext) GetAvoidNegativeTs() int32 {
	return (int32)(s.avoid_negative_ts)
}

// Custom: SetAvoidNegativeTs sets `AVFormatContext.avoid_negative_ts` value.
func (s *AvFormatContext) SetAvoidNegativeTs(v int32) {
	s.avoid_negative_ts = (C.int)(v)
}

// Custom: GetAvoidNegativeTsAddr gets `AVFormatContext.avoid_negative_ts` address.
func (s *AvFormatContext) GetAvoidNegativeTsAddr() *int32 {
	return (*int32)(&s.avoid_negative_ts)
}

// Custom: GetTsId gets `AVFormatContext.ts_id` value.
func (s *AvFormatContext) GetTsId() int32 {
	return (int32)(s.ts_id)
}

// Custom: SetTsId sets `AVFormatContext.ts_id` value.
func (s *AvFormatContext) SetTsId(v int32) {
	s.ts_id = (C.int)(v)
}

// Custom: GetTsIdAddr gets `AVFormatContext.ts_id` address.
func (s *AvFormatContext) GetTsIdAddr() *int32 {
	return (*int32)(&s.ts_id)
}

// Custom: GetAudioPreload gets `AVFormatContext.audio_preload` value.
func (s *AvFormatContext) GetAudioPreload() int32 {
	return (int32)(s.audio_preload)
}

// Custom: SetAudioPreload sets `AVFormatContext.audio_preload` value.
func (s *AvFormatContext) SetAudioPreload(v int32) {
	s.audio_preload = (C.int)(v)
}

// Custom: GetAudioPreloadAddr gets `AVFormatContext.audio_preload` address.
func (s *AvFormatContext) GetAudioPreloadAddr() *int32 {
	return (*int32)(&s.audio_preload)
}

// Custom: GetMaxChunkDuration gets `AVFormatContext.max_chunk_duration` value.
func (s *AvFormatContext) GetMaxChunkDuration() int32 {
	return (int32)(s.max_chunk_duration)
}

// Custom: SetMaxChunkDuration sets `AVFormatContext.max_chunk_duration` value.
func (s *AvFormatContext) SetMaxChunkDuration(v int32) {
	s.max_chunk_duration = (C.int)(v)
}

// Custom: GetMaxChunkDurationAddr gets `AVFormatContext.max_chunk_duration` address.
func (s *AvFormatContext) GetMaxChunkDurationAddr() *int32 {
	return (*int32)(&s.max_chunk_duration)
}

// Custom: GetMaxChunkSize gets `AVFormatContext.max_chunk_size` value.
func (s *AvFormatContext) GetMaxChunkSize() int32 {
	return (int32)(s.max_chunk_size)
}

// Custom: SetMaxChunkSize sets `AVFormatContext.max_chunk_size` value.
func (s *AvFormatContext) SetMaxChunkSize(v int32) {
	s.max_chunk_size = (C.int)(v)
}

// Custom: GetMaxChunkSizeAddr gets `AVFormatContext.max_chunk_size` address.
func (s *AvFormatContext) GetMaxChunkSizeAddr() *int32 {
	return (*int32)(&s.max_chunk_size)
}

// Custom: GetUseWallclockAsTimestamps gets `AVFormatContext.use_wallclock_as_timestamps` value.
func (s *AvFormatContext) GetUseWallclockAsTimestamps() int32 {
	return (int32)(s.use_wallclock_as_timestamps)
}

// Custom: SetUseWallclockAsTimestamps sets `AVFormatContext.use_wallclock_as_timestamps` value.
func (s *AvFormatContext) SetUseWallclockAsTimestamps(v int32) {
	s.use_wallclock_as_timestamps = (C.int)(v)
}

// Custom: GetUseWallclockAsTimestampsAddr gets `AVFormatContext.use_wallclock_as_timestamps` address.
func (s *AvFormatContext) GetUseWallclockAsTimestampsAddr() *int32 {
	return (*int32)(&s.use_wallclock_as_timestamps)
}

// Custom: GetAvioFlags gets `AVFormatContext.avio_flags` value.
func (s *AvFormatContext) GetAvioFlags() int32 {
	return (int32)(s.avio_flags)
}

// Custom: SetAvioFlags sets `AVFormatContext.avio_flags` value.
func (s *AvFormatContext) SetAvioFlags(v int32) {
	s.avio_flags = (C.int)(v)
}

// Custom: GetAvioFlagsAddr gets `AVFormatContext.avio_flags` address.
func (s *AvFormatContext) GetAvioFlagsAddr() *int32 {
	return (*int32)(&s.avio_flags)
}

// Custom: GetDurationEstimationMethod gets `AVFormatContext.duration_estimation_method` value.
func (s *AvFormatContext) GetDurationEstimationMethod() AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(s.duration_estimation_method)
}

// Custom: SetDurationEstimationMethod sets `AVFormatContext.duration_estimation_method` value.
func (s *AvFormatContext) SetDurationEstimationMethod(v AvDurationEstimationMethod) {
	s.duration_estimation_method = (C.enum_AVDurationEstimationMethod)(v)
}

// Custom: GetDurationEstimationMethodAddr gets `AVFormatContext.duration_estimation_method` address.
func (s *AvFormatContext) GetDurationEstimationMethodAddr() *AvDurationEstimationMethod {
	return (*AvDurationEstimationMethod)(unsafe.Pointer(&s.duration_estimation_method))
}

// Custom: GetSkipInitialBytes gets `AVFormatContext.skip_initial_bytes` value.
func (s *AvFormatContext) GetSkipInitialBytes() int64 {
	return (int64)(s.skip_initial_bytes)
}

// Custom: SetSkipInitialBytes sets `AVFormatContext.skip_initial_bytes` value.
func (s *AvFormatContext) SetSkipInitialBytes(v int64) {
	s.skip_initial_bytes = (C.int64_t)(v)
}

// Custom: GetSkipInitialBytesAddr gets `AVFormatContext.skip_initial_bytes` address.
func (s *AvFormatContext) GetSkipInitialBytesAddr() *int64 {
	return (*int64)(&s.skip_initial_bytes)
}

// Custom: GetCorrectTsOverflow gets `AVFormatContext.correct_ts_overflow` value.
func (s *AvFormatContext) GetCorrectTsOverflow() uint32 {
	return (uint32)(s.correct_ts_overflow)
}

// Custom: SetCorrectTsOverflow sets `AVFormatContext.correct_ts_overflow` value.
func (s *AvFormatContext) SetCorrectTsOverflow(v uint32) {
	s.correct_ts_overflow = (C.uint)(v)
}

// Custom: GetCorrectTsOverflowAddr gets `AVFormatContext.correct_ts_overflow` address.
func (s *AvFormatContext) GetCorrectTsOverflowAddr() *uint32 {
	return (*uint32)(&s.correct_ts_overflow)
}

// Custom: GetSeek2any gets `AVFormatContext.seek2any` value.
func (s *AvFormatContext) GetSeek2any() int32 {
	return (int32)(s.seek2any)
}

// Custom: SetSeek2any sets `AVFormatContext.seek2any` value.
func (s *AvFormatContext) SetSeek2any(v int32) {
	s.seek2any = (C.int)(v)
}

// Custom: GetSeek2anyAddr gets `AVFormatContext.seek2any` address.
func (s *AvFormatContext) GetSeek2anyAddr() *int32 {
	return (*int32)(&s.seek2any)
}

// Custom: GetFlushPackets gets `AVFormatContext.flush_packets` value.
func (s *AvFormatContext) GetFlushPackets() int32 {
	return (int32)(s.flush_packets)
}

// Custom: SetFlushPackets sets `AVFormatContext.flush_packets` value.
func (s *AvFormatContext) SetFlushPackets(v int32) {
	s.flush_packets = (C.int)(v)
}

// Custom: GetFlushPacketsAddr gets `AVFormatContext.flush_packets` address.
func (s *AvFormatContext) GetFlushPacketsAddr() *int32 {
	return (*int32)(&s.flush_packets)
}

// Custom: GetProbeScore gets `AVFormatContext.probe_score` value.
func (s *AvFormatContext) GetProbeScore() int32 {
	return (int32)(s.probe_score)
}

// Custom: SetProbeScore sets `AVFormatContext.probe_score` value.
func (s *AvFormatContext) SetProbeScore(v int32) {
	s.probe_score = (C.int)(v)
}

// Custom: GetProbeScoreAddr gets `AVFormatContext.probe_score` address.
func (s *AvFormatContext) GetProbeScoreAddr() *int32 {
	return (*int32)(&s.probe_score)
}

// Custom: GetFormatProbesize gets `AVFormatContext.format_probesize` value.
func (s *AvFormatContext) GetFormatProbesize() int32 {
	return (int32)(s.format_probesize)
}

// Custom: SetFormatProbesize sets `AVFormatContext.format_probesize` value.
func (s *AvFormatContext) SetFormatProbesize(v int32) {
	s.format_probesize = (C.int)(v)
}

// Custom: GetFormatProbesizeAddr gets `AVFormatContext.format_probesize` address.
func (s *AvFormatContext) GetFormatProbesizeAddr() *int32 {
	return (*int32)(&s.format_probesize)
}

// Custom: GetCodecWhitelist gets `AVFormatContext.codec_whitelist` value.
func (s *AvFormatContext) GetCodecWhitelist() string {
	return C.GoString(s.codec_whitelist)
}

// Custom: GetFormatWhitelist gets `AVFormatContext.format_whitelist` value.
func (s *AvFormatContext) GetFormatWhitelist() string {
	return C.GoString(s.format_whitelist)
}

// Custom: GetInternal gets `AVFormatContext.internal` value.
func (s *AvFormatContext) GetInternal() *AvFormatInternal {
	return (*AvFormatInternal)(s.internal)
}

// Custom: SetInternal sets `AVFormatContext.internal` value.
func (s *AvFormatContext) SetInternal(v *AvFormatInternal) {
	s.internal = (*C.struct_AVFormatInternal)(v)
}

// Custom: GetInternalAddr gets `AVFormatContext.internal` address.
func (s *AvFormatContext) GetInternalAddr() **AvFormatInternal {
	return (**AvFormatInternal)(unsafe.Pointer(&s.internal))
}

// Custom: GetIoRepositioned gets `AVFormatContext.io_repositioned` value.
func (s *AvFormatContext) GetIoRepositioned() int32 {
	return (int32)(s.io_repositioned)
}

// Custom: SetIoRepositioned sets `AVFormatContext.io_repositioned` value.
func (s *AvFormatContext) SetIoRepositioned(v int32) {
	s.io_repositioned = (C.int)(v)
}

// Custom: GetIoRepositionedAddr gets `AVFormatContext.io_repositioned` address.
func (s *AvFormatContext) GetIoRepositionedAddr() *int32 {
	return (*int32)(&s.io_repositioned)
}

// Custom: GetVideoCodec gets `AVFormatContext.video_codec` value.
func (s *AvFormatContext) GetVideoCodec() *AvCodec {
	return (*AvCodec)(s.video_codec)
}

// Custom: SetVideoCodec sets `AVFormatContext.video_codec` value.
func (s *AvFormatContext) SetVideoCodec(v *AvCodec) {
	s.video_codec = (*C.struct_AVCodec)(v)
}

// Custom: GetVideoCodecAddr gets `AVFormatContext.video_codec` address.
func (s *AvFormatContext) GetVideoCodecAddr() **AvCodec {
	return (**AvCodec)(unsafe.Pointer(&s.video_codec))
}

// Custom: GetAudioCodec gets `AVFormatContext.audio_codec` value.
func (s *AvFormatContext) GetAudioCodec() *AvCodec {
	return (*AvCodec)(s.audio_codec)
}

// Custom: SetAudioCodec sets `AVFormatContext.audio_codec` value.
func (s *AvFormatContext) SetAudioCodec(v *AvCodec) {
	s.audio_codec = (*C.struct_AVCodec)(v)
}

// Custom: GetAudioCodecAddr gets `AVFormatContext.audio_codec` address.
func (s *AvFormatContext) GetAudioCodecAddr() **AvCodec {
	return (**AvCodec)(unsafe.Pointer(&s.audio_codec))
}

// Custom: GetSubtitleCodec gets `AVFormatContext.subtitle_codec` value.
func (s *AvFormatContext) GetSubtitleCodec() *AvCodec {
	return (*AvCodec)(s.subtitle_codec)
}

// Custom: SetSubtitleCodec sets `AVFormatContext.subtitle_codec` value.
func (s *AvFormatContext) SetSubtitleCodec(v *AvCodec) {
	s.subtitle_codec = (*C.struct_AVCodec)(v)
}

// Custom: GetSubtitleCodecAddr gets `AVFormatContext.subtitle_codec` address.
func (s *AvFormatContext) GetSubtitleCodecAddr() **AvCodec {
	return (**AvCodec)(unsafe.Pointer(&s.subtitle_codec))
}

// Custom: GetDataCodec gets `AVFormatContext.data_codec` value.
func (s *AvFormatContext) GetDataCodec() *AvCodec {
	return (*AvCodec)(s.data_codec)
}

// Custom: SetDataCodec sets `AVFormatContext.data_codec` value.
func (s *AvFormatContext) SetDataCodec(v *AvCodec) {
	s.data_codec = (*C.struct_AVCodec)(v)
}

// Custom: GetDataCodecAddr gets `AVFormatContext.data_codec` address.
func (s *AvFormatContext) GetDataCodecAddr() **AvCodec {
	return (**AvCodec)(unsafe.Pointer(&s.data_codec))
}

// Custom: GetMetadataHeaderPadding gets `AVFormatContext.metadata_header_padding` value.
func (s *AvFormatContext) GetMetadataHeaderPadding() int32 {
	return (int32)(s.metadata_header_padding)
}

// Custom: SetMetadataHeaderPadding sets `AVFormatContext.metadata_header_padding` value.
func (s *AvFormatContext) SetMetadataHeaderPadding(v int32) {
	s.metadata_header_padding = (C.int)(v)
}

// Custom: GetMetadataHeaderPaddingAddr gets `AVFormatContext.metadata_header_padding` address.
func (s *AvFormatContext) GetMetadataHeaderPaddingAddr() *int32 {
	return (*int32)(&s.metadata_header_padding)
}

// Custom: GetOpaque gets `AVFormatContext.opaque` value.
func (s *AvFormatContext) GetOpaque() unsafe.Pointer {
	return s.opaque
}

// Custom: SetOpaque sets `AVFormatContext.opaque` value.
func (s *AvFormatContext) SetOpaque(v unsafe.Pointer) {
	s.opaque = v
}

// Custom: GetOpaqueAddr gets `AVFormatContext.opaque` address.
func (s *AvFormatContext) GetOpaqueAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&s.opaque)
}

// Custom: GetOutputTsOffset gets `AVFormatContext.output_ts_offset` value.
func (s *AvFormatContext) GetOutputTsOffset() int64 {
	return (int64)(s.output_ts_offset)
}

// Custom: SetOutputTsOffset sets `AVFormatContext.output_ts_offset` value.
func (s *AvFormatContext) SetOutputTsOffset(v int64) {
	s.output_ts_offset = (C.int64_t)(v)
}

// Custom: GetOutputTsOffsetAddr gets `AVFormatContext.output_ts_offset` address.
func (s *AvFormatContext) GetOutputTsOffsetAddr() *int64 {
	return (*int64)(&s.output_ts_offset)
}

// Custom: GetDumpSeparator gets `AVFormatContext.dump_separator` value.
func (s *AvFormatContext) GetDumpSeparator() *uint8 {
	return (*uint8)(s.dump_separator)
}

// Custom: SetDumpSeparator sets `AVFormatContext.dump_separator` value.
func (s *AvFormatContext) SetDumpSeparator(v *uint8) {
	s.dump_separator = (*C.uint8_t)(v)
}

// Custom: GetDumpSeparatorAddr gets `AVFormatContext.dump_separator` address.
func (s *AvFormatContext) GetDumpSeparatorAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&s.dump_separator))
}

// Custom: GetDataCodecId gets `AVFormatContext.data_codec_id` value.
func (s *AvFormatContext) GetDataCodecId() AvCodecID {
	return (AvCodecID)(s.data_codec_id)
}

// Custom: SetDataCodecId sets `AVFormatContext.data_codec_id` value.
func (s *AvFormatContext) SetDataCodecId(v AvCodecID) {
	s.data_codec_id = (C.enum_AVCodecID)(v)
}

// Custom: GetDataCodecIdAddr gets `AVFormatContext.data_codec_id` address.
func (s *AvFormatContext) GetDataCodecIdAddr() *AvCodecID {
	return (*AvCodecID)(unsafe.Pointer(&s.data_codec_id))
}

// Custom: GetProtocolWhitelist gets `AVFormatContext.protocol_whitelist` value.
func (s *AvFormatContext) GetProtocolWhitelist() string {
	return C.GoString(s.protocol_whitelist)
}

// Custom: GetProtocolBlacklist gets `AVFormatContext.protocol_blacklist` value.
func (s *AvFormatContext) GetProtocolBlacklist() string {
	return C.GoString(s.protocol_blacklist)
}

// Custom: GetMaxStreams gets `AVFormatContext.max_streams` value.
func (s *AvFormatContext) GetMaxStreams() int32 {
	return (int32)(s.max_streams)
}

// Custom: SetMaxStreams sets `AVFormatContext.max_streams` value.
func (s *AvFormatContext) SetMaxStreams(v int32) {
	s.max_streams = (C.int)(v)
}

// Custom: GetMaxStreamsAddr gets `AVFormatContext.max_streams` address.
func (s *AvFormatContext) GetMaxStreamsAddr() *int32 {
	return (*int32)(&s.max_streams)
}

// Custom: GetSkipEstimateDurationFromPts gets `AVFormatContext.skip_estimate_duration_from_pts` value.
func (s *AvFormatContext) GetSkipEstimateDurationFromPts() int32 {
	return (int32)(s.skip_estimate_duration_from_pts)
}

// Custom: SetSkipEstimateDurationFromPts sets `AVFormatContext.skip_estimate_duration_from_pts` value.
func (s *AvFormatContext) SetSkipEstimateDurationFromPts(v int32) {
	s.skip_estimate_duration_from_pts = (C.int)(v)
}

// Custom: GetSkipEstimateDurationFromPtsAddr gets `AVFormatContext.skip_estimate_duration_from_pts` address.
func (s *AvFormatContext) GetSkipEstimateDurationFromPtsAddr() *int32 {
	return (*int32)(&s.skip_estimate_duration_from_pts)
}

// Custom: GetMaxProbePackets gets `AVFormatContext.max_probe_packets` value.
func (s *AvFormatContext) GetMaxProbePackets() int32 {
	return (int32)(s.max_probe_packets)
}

// Custom: SetMaxProbePackets sets `AVFormatContext.max_probe_packets` value.
func (s *AvFormatContext) SetMaxProbePackets(v int32) {
	s.max_probe_packets = (C.int)(v)
}

// Custom: GetMaxProbePacketsAddr gets `AVFormatContext.max_probe_packets` address.
func (s *AvFormatContext) GetMaxProbePacketsAddr() *int32 {
	return (*int32)(&s.max_probe_packets)
}

const (
	AVFMT_FLAG_GENPTS          = C.AVFMT_FLAG_GENPTS
	AVFMT_FLAG_IGNIDX          = C.AVFMT_FLAG_IGNIDX
	AVFMT_FLAG_NONBLOCK        = C.AVFMT_FLAG_NONBLOCK
	AVFMT_FLAG_IGNDTS          = C.AVFMT_FLAG_IGNDTS
	AVFMT_FLAG_NOFILLIN        = C.AVFMT_FLAG_NOFILLIN
	AVFMT_FLAG_NOPARSE         = C.AVFMT_FLAG_NOPARSE
	AVFMT_FLAG_NOBUFFER        = C.AVFMT_FLAG_NOBUFFER
	AVFMT_FLAG_CUSTOM_IO       = C.AVFMT_FLAG_CUSTOM_IO
	AVFMT_FLAG_DISCARD_CORRUPT = C.AVFMT_FLAG_DISCARD_CORRUPT
	AVFMT_FLAG_FLUSH_PACKETS   = C.AVFMT_FLAG_FLUSH_PACKETS
	AVFMT_FLAG_BITEXACT        = C.AVFMT_FLAG_BITEXACT
	AVFMT_FLAG_MP4A_LATM       = C.AVFMT_FLAG_MP4A_LATM
	AVFMT_FLAG_SORT_DTS        = C.AVFMT_FLAG_SORT_DTS
	AVFMT_FLAG_PRIV_OPT        = C.AVFMT_FLAG_PRIV_OPT
	AVFMT_FLAG_KEEP_SIDE_DATA  = C.AVFMT_FLAG_KEEP_SIDE_DATA
	AVFMT_FLAG_FAST_SEEK       = C.AVFMT_FLAG_FAST_SEEK
	AVFMT_FLAG_SHORTEST        = C.AVFMT_FLAG_SHORTEST
	AVFMT_FLAG_AUTO_BSF        = C.AVFMT_FLAG_AUTO_BSF
)

const (
	AVFMT_EVENT_FLAG_METADATA_UPDATED = C.AVFMT_EVENT_FLAG_METADATA_UPDATED
)

// Deprecated: No use
func AvFormatGetProbeScore(s *AvFormatContext) int32 {
	return (int32)(C.av_format_get_probe_score((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use
func AvFormatGetVideoCodec(s *AvFormatContext) *AvCodec {
	return (*AvCodec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use
func AvFormatSetVideoCodec(s *AvFormatContext, c *AvCodec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// Deprecated: No use
func AvFormatGetAudioCodec(s *AvFormatContext) *AvCodec {
	return (*AvCodec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use
func AvFormatSetAudioCodec(s *AvFormatContext, c *AvCodec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// Deprecated: No use
func AvFormatGetSubtitleCodec(s *AvFormatContext) *AvCodec {
	return (*AvCodec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use
func AvFormatSetSubtitleCodec(s *AvFormatContext, c *AvCodec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// Deprecated: No use
func AvFormatGetDataCodec(s *AvFormatContext) *AvCodec {
	return (*AvCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use
func AvFormatSetDataCodec(s *AvFormatContext, c *AvCodec) {
	C.av_format_set_data_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// Deprecated: No use
func AvFormatGetOpaque(s *AvFormatContext) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_format_get_opaque((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use
func AvFormatSetOpaque(s *AvFormatContext, opaque unsafe.Pointer) {
	C.av_format_set_opaque((*C.struct_AVFormatContext)(s), opaque)
}

// Deprecated: No use
func AvFormatGetControlMessageCb(s *AvFormatContext) AvFormatControlMessageFunc {
	return (AvFormatControlMessageFunc)(C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use
func AvFormatSetControlMessageCb(s *AvFormatContext, callback AvFormatControlMessageFunc) {
	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(s), (C.av_format_control_message)(callback))
}

// Deprecated: No use
func AvFormatGetOpenCb(s *AvFormatContext) AvOpenCallbackFunc {
	return (AvOpenCallbackFunc)(C.av_format_get_open_cb((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use
func AvFormatSetOpenCb(s *AvFormatContext, callback AvOpenCallbackFunc) {
	C.av_format_set_open_cb((*C.struct_AVFormatContext)(s), (C.AVOpenCallback)(callback))
}

// AvFormatInjectGlobalSideData will cause global side data to be injected in the next packet
// of each stream as well as after any subsequent seek.
func AvFormatInjectGlobalSideData(s *AvFormatContext) {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(s))
}

// AvFmtCtxGetDurationEstimationMethod returns the method used to set ctx->duration.
func AvFmtCtxGetDurationEstimationMethod(s *AvFormatContext) AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(s)))
}

// AvFormatVersion returns the LIBAVFORMAT_VERSION_INT constant.
func AvFormatVersion() uint32 {
	return (uint32)(C.avformat_version())
}

// AvFormatConfiguration returns the libavformat build-time configuration.
func AvFormatConfiguration() string {
	return C.GoString(C.avformat_configuration())
}

// AvFormatLicense returns the libavformat license.
func AvFormatLicense() string {
	return C.GoString(C.avformat_license())
}

// Deprecated: No use
func AvRegisterAll() {
	C.av_register_all()
}

// Deprecated: No use
func AvRegisterInputFormat(format *AvInputFormat) {
	C.av_register_input_format((*C.struct_AVInputFormat)(format))
}

// Deprecated: No use
func AvRegisterOutputFormat(format *AvOutputFormat) {
	C.av_register_output_format((*C.struct_AVOutputFormat)(format))
}

// AvFormatNetworkInit does global initialization of network libraries. This is optional,
// and not recommended anymore.
func AvFormatNetworkInit() int32 {
	return (int32)(C.avformat_network_init())
}

// AvFormatNetworkDeinit undoes the initialization done by avformat_network_init. Call it only
// once for each time you called avformat_network_init.
func AvFormatNetworkDeinit() int32 {
	return (int32)(C.avformat_network_deinit())
}

// Deprecated: No use
func AvIformatNext(f *AvInputFormat) *AvInputFormat {
	return (*AvInputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

// Deprecated: No use
func AvOformatNext(f *AvOutputFormat) *AvOutputFormat {
	return (*AvOutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

// AvMuxerIterate iterates over all registered muxers.
func AvMuxerIterate(opaque *unsafe.Pointer) *AvOutputFormat {
	return (*AvOutputFormat)(C.av_muxer_iterate(opaque))
}

// AvDemuxerIterate iterates over all registered demuxers.
func AvDemuxerIterate(opaque *unsafe.Pointer) *AvInputFormat {
	return (*AvInputFormat)(C.av_demuxer_iterate(opaque))
}

// AvFormatAllocContext allocates an AVFormatContext.
func AvFormatAllocContext() *AvFormatContext {
	return (*AvFormatContext)(C.avformat_alloc_context())
}

// AvFormatFreeContext frees an AVFormatContext and all its streams.
func AvFormatFreeContext(s *AvFormatContext) {
	C.avformat_free_context((*C.struct_AVFormatContext)(s))
}

// AvFormatGetClass gets the AVClass for AVFormatContext. It can be used in combination with
// AV_OPT_SEARCH_FAKE_OBJ for examining options.
func AvFormatGetClass() *AvClass {
	return (*AvClass)(C.avformat_get_class())
}

// AvFormatNewStream adds a new stream to a media file.
func AvFormatNewStream(s *AvFormatContext, c *AvCodec) *AvStream {
	return (*AvStream)(C.avformat_new_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c)))
}

// AvStreamAddSideData wraps an existing array as stream side data.
func AvStreamAddSideData(st *AvStream, _type AvPacketSideDataType, data *uint8, size uint) int32 {
	return (int32)(C.av_stream_add_side_data((*C.struct_AVStream)(st),
		(C.enum_AVPacketSideDataType)(_type), (*C.uint8_t)(data), (C.size_t)(size)))
}

// AvStreamNewSideData allocates new information from stream.
func AvStreamNewSideData(st *AvStream, _type AvPacketSideDataType, size int32) *uint8 {
	return (*uint8)(C.av_stream_new_side_data((*C.struct_AVStream)(st),
		(C.enum_AVPacketSideDataType)(_type), (C.int)(size)))
}

// AvStreamGetSideData gets side information from stream.
func AvStreamGetSideData(st *AvStream, _type AvPacketSideDataType, size *int32) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(st),
		(C.enum_AVPacketSideDataType)(_type), (*C.int)(size)))
}

// AvNewProgram
func AvNewProgram(s *AvFormatContext, id int32) *AvProgram {
	return (*AvProgram)(C.av_new_program((*C.struct_AVFormatContext)(s), (C.int)(id)))
}

// AvFormatAllocOutputContext2 allocates an AVFormatContext for an output format.
func AvFormatAllocOutputContext2(ctx **AvFormatContext,
	oformat *AvOutputFormat, formatName, filename string) int32 {
	formatNamePtr, formatNameFunc := StringCasting(formatName)
	defer formatNameFunc()
	filenamePtr, filenameFunc := StringCasting(filename)
	defer filenameFunc()
	return (int32)(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(ctx)),
		(*C.struct_AVOutputFormat)(oformat), (*C.char)(formatNamePtr), (*C.char)(filenamePtr)))
}

// AvFindInputFormat finds AVInputFormat based on the short name of the input format.
func AvFindInputFormat(shortName string) *AvInputFormat {
	shortNamePtr, shortNameFunc := StringCasting(shortName)
	defer shortNameFunc()
	return (*AvInputFormat)(C.av_find_input_format((*C.char)(shortNamePtr)))
}

// AvProbeInputFormat guesses the file format.
func AvProbeInputFormat(pd *AvProbeData, isOpened int32) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), (C.int)(isOpened)))
}

// AvProbeInputFormat2 guesses the file format.
func AvProbeInputFormat2(pd *AvProbeData, isOpened int32, scoreMax *int32) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd),
		(C.int)(isOpened), (*C.int)(scoreMax)))
}

// AvProbeInputFormat3 guesses the file format.
func AvProbeInputFormat3(pd *AvProbeData, isOpened int32, scoreRet *int32) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd),
		(C.int)(isOpened), (*C.int)(scoreRet)))
}

// AvProbeInputBuffer2 probes a bytestream to determine the input format. Each time a probe returns
// with a score that is too low, the probe buffer size is increased and another
// attempt is made. When the maximum probe size is reached, the input format
// with the highest score is returned.
func AvProbeInputBuffer2(pb *AvIOContext, fmt **AvInputFormat,
	url string, logctx unsafe.Pointer, offset, maxProbeSize uint32) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb), (**C.AVInputFormat)(unsafe.Pointer(fmt)),
		(*C.char)(urlPtr), logctx, (C.uint)(offset), (C.uint)(maxProbeSize)))
}

// AvProbeInputBuffer likes AvProbeInputBuffer2() but returns 0 on success
func AvProbeInputBuffer(pb *AvIOContext, fmt **AvInputFormat,
	url string, logctx unsafe.Pointer, offset, maxProbeSize uint32) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.av_probe_input_buffer((*C.struct_AVIOContext)(pb), (**C.AVInputFormat)(unsafe.Pointer(fmt)),
		(*C.char)(urlPtr), logctx, (C.uint)(offset), (C.uint)(maxProbeSize)))
}

// AvFormatOpenInput open an input stream and read the header. The codecs are not opened.
// The stream must be closed with AvFormatCloseInput().
func AvFormatOpenInput(ps **AvFormatContext, url string, fmt *AvInputFormat, options **AvDictionary) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)),
		(*C.char)(urlPtr), (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// Deprecated: Use an AVDictionary to pass options to a demuxer.
func AvDemuxerOpen(ic *AvFormatContext) int32 {
	return (int32)(C.av_demuxer_open((*C.struct_AVFormatContext)(ic)))
}

// AvFormatFindStreamInfo reads packets of a media file to get stream information.
func AvFormatFindStreamInfo(ic *AvFormatContext, options **AvDictionary) int32 {
	return (int32)(C.avformat_find_stream_info((*C.struct_AVFormatContext)(ic),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvFindProgramFromStream finds the programs which belong to a given stream.
func AvFindProgramFromStream(ic *AvFormatContext, last *AvProgram, s int32) *AvProgram {
	return (*AvProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(ic),
		(*C.struct_AVProgram)(last), (C.int)(s)))
}

// AvProgramAddStreamIndex
func AvProgramAddStreamIndex(ic *AvFormatContext, progid int32, idx uint32) {
	C.av_program_add_stream_index((*C.struct_AVFormatContext)(ic), (C.int)(progid), (C.uint)(idx))
}

// AvFindBestStream finds the "best" stream in the file.
func AvFindBestStream(ic *AvFormatContext, _type AvMediaType,
	wantedStreamNb, relatedStream int32,
	decoderRet **AvCodec, flags int32) int32 {
	return (int32)(C.av_find_best_stream((*C.struct_AVFormatContext)(ic),
		(C.enum_AVMediaType)(_type), (C.int)(wantedStreamNb), (C.int)(relatedStream),
		(**C.struct_AVCodec)(unsafe.Pointer(decoderRet)), (C.int)(flags)))
}

// AvReadFrame returns the next frame of a stream.
func AvReadFrame(ic *AvFormatContext, pkt *AvPacket) int32 {
	return (int32)(C.av_read_frame((*C.struct_AVFormatContext)(ic), (*C.struct_AVPacket)(pkt)))
}

// AvSeekFrame seeks to the keyframe at timestamp.
func AvSeekFrame(ic *AvFormatContext, streamIndex int32, timestamp int64, flags int32) int32 {
	return (int32)(C.av_seek_frame((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex), (C.int64_t)(timestamp), (C.int)(flags)))
}

// AvFormatSeekFile seeks to timestamp ts.
func AvFormatSeekFile(ic *AvFormatContext, streamIndex int32, minTs, ts, maxTs int64, flags int32) int32 {
	return (int32)(C.avformat_seek_file((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex), (C.int64_t)(minTs), (C.int64_t)(ts), (C.int64_t)(maxTs), (C.int)(flags)))
}

// AvFormatFlush siscards all internally buffered data. This can be useful when dealing with
// discontinuities in the byte stream. Generally works only with formats that
// can resync. This includes headerless formats like MPEG-TS/TS but should also
// work with NUT, Ogg and in a limited way AVI for example.
func AvFormatFlush(ic *AvFormatContext) int32 {
	return (int32)(C.avformat_flush((*C.struct_AVFormatContext)(ic)))
}

// AvReadPlay starts playing a network-based stream (e.g. RTSP stream) at the
// current position.
func AvReadPlay(ic *AvFormatContext) int32 {
	return (int32)(C.av_read_play((*C.struct_AVFormatContext)(ic)))
}

// AvReadPause pauses a network-based stream (e.g. RTSP stream).
func AvReadPause(ic *AvFormatContext) int32 {
	return (int32)(C.av_read_pause((*C.struct_AVFormatContext)(ic)))
}

// AvFormatCloseInput closes an opened input AVFormatContext. Free it and all its contents
// and set *s to NULL.
func AvFormatCloseInput(ic **AvFormatContext) {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(ic)))
}

const (
	AVSEEK_FLAG_BACKWARD = C.AVSEEK_FLAG_BACKWARD
	AVSEEK_FLAG_BYTE     = C.AVSEEK_FLAG_BYTE
	AVSEEK_FLAG_ANY      = C.AVSEEK_FLAG_ANY
	AVSEEK_FLAG_FRAME    = C.AVSEEK_FLAG_FRAME
)

const (
	AVSTREAM_INIT_IN_WRITE_HEADER = C.AVSTREAM_INIT_IN_WRITE_HEADER
	AVSTREAM_INIT_IN_INIT_OUTPUT  = C.AVSTREAM_INIT_IN_INIT_OUTPUT
)

// AvFormatWriteHeader allocates the stream private data and write the stream header to
// an output media file.
func AvFormatWriteHeader(ic *AvFormatContext, options **AvDictionary) int32 {
	return (int32)(C.avformat_write_header((*C.struct_AVFormatContext)(ic),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvFormatInitOutput allocates the stream private data and initialize the codec,
// but do not write the header.
func AvFormatInitOutput(ic *AvFormatContext, options **AvDictionary) int32 {
	return (int32)(C.avformat_init_output((*C.struct_AVFormatContext)(ic),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvWriteFrame writes a packet to an output media file.
func AvWriteFrame(ic *AvFormatContext, pkt *AvPacket) int32 {
	return (int32)(C.av_write_frame((*C.struct_AVFormatContext)(ic), (*C.struct_AVPacket)(pkt)))
}

// AvInterleavedWriteFrame writes a packet to an output media file ensuring correct interleaving.
func AvInterleavedWriteFrame(ic *AvFormatContext, pkt *AvPacket) int32 {
	return (int32)(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(ic), (*C.struct_AVPacket)(pkt)))
}

// AvWriteUncodedFrame writes an uncoded frame to an output media file.
func AvWriteUncodedFrame(ic *AvFormatContext, streamIndex int32, frame *AvFrame) int32 {
	return (int32)(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex), (*C.struct_AVFrame)(frame)))
}

// AvInterleavedWriteUncodedFrame Write an uncoded frame to an output media file.
func AvInterleavedWriteUncodedFrame(ic *AvFormatContext, streamIndex int32, frame *AvFrame) int32 {
	return (int32)(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex), (*C.struct_AVFrame)(frame)))
}

// AvWriteUncodedFrameQuery tests whether a muxer supports uncoded frame.
func AvWriteUncodedFrameQuery(ic *AvFormatContext, streamIndex int32) int32 {
	return (int32)(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex)))
}

// AvWriteTrailer writes the stream trailer to an output media file and free the file private data.
func AvWriteTrailer(ic *AvFormatContext) int32 {
	return (int32)(C.av_write_trailer((*C.struct_AVFormatContext)(ic)))
}

// AvGuessFormat returns the output format in the list of registered output formats
// which best matches the provided parameters, or return NULL if there is no match.
func AvGuessFormat(shortName, filename, mimeType string) *AvOutputFormat {
	shortNamePtr, shortNameFunc := StringCasting(shortName)
	defer shortNameFunc()
	filenamePtr, filenameFunc := StringCasting(filename)
	defer filenameFunc()
	mimeTypePtr, mimeTypeFunc := StringCasting(mimeType)
	defer mimeTypeFunc()
	return (*AvOutputFormat)(C.av_guess_format((*C.char)(shortNamePtr),
		(*C.char)(filenamePtr),
		(*C.char)(mimeTypePtr)))
}

// AvGuessCodec guesses the codec ID based upon muxer and filename.
func AvGuessCodec(fmt *AvOutputFormat, shortName,
	filename, mimeType string, _type AvMediaType) AvCodecID {
	shortNamePtr, shortNameFunc := StringCasting(shortName)
	defer shortNameFunc()
	filenamePtr, filenameFunc := StringCasting(filename)
	defer filenameFunc()
	mimeTypePtr, mimeTypeFunc := StringCasting(mimeType)
	defer mimeTypeFunc()
	return (AvCodecID)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt),
		(*C.char)(shortNamePtr),
		(*C.char)(filenamePtr),
		(*C.char)(mimeTypePtr),
		(C.enum_AVMediaType)(_type)))
}

// AvGetOutputTimestamp gets timing information for the data currently output.
func AvGetOutputTimestamp(ic *AvFormatContext, stream int32, dts, wall *int64) int32 {
	return (int32)(C.av_get_output_timestamp((*C.struct_AVFormatContext)(ic),
		(C.int)(stream), (*C.int64_t)(dts), (*C.int64_t)(wall)))
}

// AvHexDump sends a nice hexadecimal dump of a buffer to the specified file stream.
func AvHexDump(f *FILE, buf *uint8, size int32) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(buf), (C.int)(size))
}

// AvHexDumpLog sends a nice hexadecimal dump of a buffer to the log.
func AvHexDumpLog(avcl unsafe.Pointer, level int32, buf *uint8, size int32) {
	C.av_hex_dump_log(avcl, (C.int)(level), (*C.uint8_t)(buf), (C.int)(size))
}

// AvPktDump2 sends a nice dump of a packet to the specified file stream.
func AvPktDump2(f *FILE, pkt *AvPacket, dumpPayload int32, st *AvStream) {
	C.av_pkt_dump2((*C.FILE)(f), (*C.struct_AVPacket)(pkt),
		(C.int)(dumpPayload), (*C.struct_AVStream)(st))
}

// AvPktDumpLog2 sends a nice dump of a packet to the log.
func av_pkt_dump_log2(avcl unsafe.Pointer, level int32, pkt *AvPacket, dumpPayload int32, st *AvStream) {
	C.av_pkt_dump_log2(avcl, (C.int)(level), (*C.struct_AVPacket)(pkt),
		(C.int)(dumpPayload), (*C.struct_AVStream)(st))
}

// AvCodecGetId gets the AVCodecID for the given codec tag tag.
func AvCodecGetId(tags **AvCodecTag, tag uint32) AvCodecID {
	return (AvCodecID)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), (C.uint)(tag)))
}

// AvCodecGetTag gets the codec tag for the given codec id.
func AvCodecGetTag(tags **AvCodecTag, id AvCodecID) uint32 {
	return (uint32)(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(tags)),
		(C.enum_AVCodecID)(id)))
}

// AvCodecGetTag2 gets the codec tag for the given codec id.
func AvCodecGetTag2(tags **AvCodecTag, id AvCodecID, tag *uint32) int32 {
	return (int32)(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(tags)),
		(C.enum_AVCodecID)(id), (*C.uint)(tag)))
}

// AvFindDefaultStreamIndex
func AvFindDefaultStreamIndex(ic *AvFormatContext) int32 {
	return (int32)(C.av_find_default_stream_index((*C.struct_AVFormatContext)(ic)))
}

// AvIndexSearchTimestamp gets the index for a specific timestamp.
func AvIndexSearchTimestamp(st *AvStream, timestamp int64, flags int32) int32 {
	return (int32)(C.av_index_search_timestamp((*C.struct_AVStream)(st),
		(C.int64_t)(timestamp), (C.int)(flags)))
}

// AvAddIndexEntry add an index entry into a sorted list. Update the entry if the list
// already contains it.
func AvAddIndexEntry(st *AvStream, pos, timestamp int64, size, distance, flags int32) int32 {
	return (int32)(C.av_add_index_entry((*C.struct_AVStream)(st),
		(C.int64_t)(pos), (C.int64_t)(timestamp),
		(C.int)(size), (C.int)(distance), (C.int)(flags)))
}

// AvUrlSplit splits a URL string into components.
func AvUrlSplit(url string) (proto, authorization, hostname string, port int32, path string) {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	urlLen := len(url)
	protoBuf := make([]C.char, urlLen)
	authorizationBuf := make([]C.char, urlLen)
	hostnameBuf := make([]C.char, urlLen)
	pathBuf := make([]C.char, urlLen)

	C.av_url_split((*C.char)(&protoBuf[0]), (C.int)(urlLen),
		&authorizationBuf[0], (C.int)(urlLen),
		&hostnameBuf[0], (C.int)(urlLen),
		(*C.int)(&port),
		&pathBuf[0], (C.int)(urlLen),
		(*C.char)(urlPtr))

	return C.GoString(&protoBuf[0]),
		C.GoString(&authorizationBuf[0]),
		C.GoString(&hostnameBuf[0]),
		port,
		C.GoString(&pathBuf[0])
}

// AvDumpFormat prints detailed information about the input or output format, such as
// duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func AvDumpFormat(ic *AvFormatContext, index int32, url string, isOutput int32) {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	C.av_dump_format((*C.struct_AVFormatContext)(ic),
		(C.int)(index), (*C.char)(urlPtr), (C.int)(isOutput))
}

// AvGetFrameFilename2 returns in 'buf' the path with '%d' replaced by a number.
func AvGetFrameFilename2(bufSize int32, path string, number, flags int32) (output string, ret int32) {
	pathPtr, pathFunc := StringCasting(path)
	defer pathFunc()
	outputBuf := make([]C.char, bufSize)
	ret = (int32)(C.av_get_frame_filename2(&outputBuf[0], (C.int)(bufSize),
		(*C.char)(pathPtr), (C.int)(number), (C.int)(flags)))
	return C.GoString(&outputBuf[0]), ret
}

// AvGetFrameFilename
func AvGetFrameFilename(bufSize int32, path string, number int32) (output string, ret int32) {
	pathPtr, pathFunc := StringCasting(path)
	defer pathFunc()
	outputBuf := make([]C.char, bufSize)
	ret = (int32)(C.av_get_frame_filename(&outputBuf[0], (C.int)(bufSize),
		(*C.char)(pathPtr), (C.int)(number)))
	return C.GoString(&outputBuf[0]), ret
}

// AvFilenameNumberTest
func AvFilenameNumberTest(filename string) int32 {
	filenamePtr, filenameFunc := StringCasting(filename)
	defer filenameFunc()
	return (int32)(C.av_filename_number_test((*C.char)(filenamePtr)))
}

// AvSdpCreate generates an SDP for an RTP session.
func AvSdpCreate(ac []*AvFormatContext, nFiles, sdpSize int32) (sdp string, ret int32) {
	sdpBuf := make([]C.char, sdpSize)
	ret = (int32)(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(&ac[0])),
		(C.int)(nFiles),
		&sdpBuf[0], (C.int)(sdpSize)))
	return C.GoString(&sdpBuf[0]), ret
}

// AvMatchExt returns a positive value if the given filename has one of the given
// extensions, 0 otherwise.
func AvMatchExt(filename, extensions string) int32 {
	filenamePtr, filenameFunc := StringCasting(filename)
	defer filenameFunc()
	extensionsPtr, extensionsFunc := StringCasting(extensions)
	defer extensionsFunc()
	return (int32)(C.av_match_ext((*C.char)(filenamePtr), (*C.char)(extensionsPtr)))
}

// AvFormatQueryCodec tests if the given container can store a codec.
func AvFormatQueryCodec(ofmt *AvOutputFormat, codeID AvCodecID, stdCompliance int32) int32 {
	return (int32)(C.avformat_query_codec((*C.struct_AVOutputFormat)(ofmt),
		(C.enum_AVCodecID)(codeID), (C.int)(stdCompliance)))
}

// AvFormatGetRiffVideoTags returns the table mapping RIFF FourCCs for video to libavcodec AVCodecID.
func AvFormatGetRiffVideoTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_video_tags())
}

// AvFormatGetRiffAudioTags returns the table mapping RIFF FourCCs for audio to AVCodecID.
func AvFormatGetRiffAudioTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_audio_tags())
}

// AvFormatGetMovVideoTags returns the table mapping MOV FourCCs for video to libavcodec AVCodecID.
func AvFormatGetMovVideoTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_video_tags())
}

// AvFormatGetMovAudioTags returns the table mapping MOV FourCCs for audio to AVCodecID.
func AvFormatGetMovAudioTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_audio_tags())
}

// AvGuessSampleAspectRatio guesses the sample aspect ratio of a frame, based on both the stream and the
// frame aspect ratio.
func AvGuessSampleAspectRatio(format *AvFormatContext, stream *AvStream, frame *AvFrame) AvRational {
	return (AvRational)(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(format),
		(*C.struct_AVStream)(stream), (*C.struct_AVFrame)(frame)))
}

// AvGuessFrameRate guesses the frame rate, based on both the container and codec information.
func AvGuessFrameRate(ctx *AvFormatContext, stream *AvStream, frame *AvFrame) AvRational {
	return (AvRational)(C.av_guess_frame_rate((*C.struct_AVFormatContext)(ctx),
		(*C.struct_AVStream)(stream), (*C.struct_AVFrame)(frame)))
}

// AvFormatMatchStreamSpecifier checks if the stream st contained in s is
// matched by the stream specifier spec.
func AvFormatMatchStreamSpecifier(ic *AvFormatContext, st *AvStream, spec string) int32 {
	specPtr, specFunc := StringCasting(spec)
	defer specFunc()
	return (int32)(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(ic),
		(*C.struct_AVStream)(st), (*C.char)(specPtr)))
}

// AvFormatQueueAttachedPictures
func AvFormatQueueAttachedPictures(ic *AvFormatContext) int32 {
	return (int32)(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(ic)))
}

// AvApplyBitstreamFilters applies a list of bitstream filters to a packet.
func AvApplyBitstreamFilters(codec *AvCodecContext, pkt *AvPacket, bsfc *AvBitStreamFilterContext) int32 {
	return (int32)(C.av_apply_bitstream_filters((*C.struct_AVCodecContext)(codec),
		(*C.struct_AVPacket)(pkt), (*C.struct_AVBitStreamFilterContext)(bsfc)))
}

// AvTimebaseSource
type AvTimebaseSource = C.enum_AVTimebaseSource

const (
	AVFMT_TBCF_AUTO        = AvTimebaseSource(C.AVFMT_TBCF_AUTO)
	AVFMT_TBCF_DECODER     = AvTimebaseSource(C.AVFMT_TBCF_DECODER)
	AVFMT_TBCF_DEMUXER     = AvTimebaseSource(C.AVFMT_TBCF_DEMUXER)
	AVFMT_TBCF_R_FRAMERATE = AvTimebaseSource(C.AVFMT_TBCF_R_FRAMERATE)
)

// AvFormatTransferInternalStreamTimingInfo transfers internal timing information from one stream to another.
func AvFormatTransferInternalStreamTimingInfo(ofmt *AvOutputFormat,
	ost, ist *AvStream,
	copyTb AvTimebaseSource) int32 {
	return (int32)(C.avformat_transfer_internal_stream_timing_info((*C.struct_AVOutputFormat)(ofmt),
		(*C.struct_AVStream)(ost), (*C.struct_AVStream)(ist),
		(C.enum_AVTimebaseSource)(copyTb)))
}

// AvStreamGetCodecTimebase gets the internal codec timebase from a stream.
func AvStreamGetCodecTimebase(st *AvStream) AvRational {
	return (AvRational)(C.av_stream_get_codec_timebase((*C.struct_AVStream)(st)))
}
