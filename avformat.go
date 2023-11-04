// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavformat/avformat.h>

typedef int (*av_format_context_open_cb)(struct AVFormatContext *s, AVIOContext **p, const char *url, int flags,
	const AVIOInterruptCB *int_cb, AVDictionary **options);

typedef int (*av_format_context_io_open_func)(struct AVFormatContext *s, AVIOContext **pb, const char *url,
	int flags, AVDictionary **options);

typedef void (*av_format_context_io_close_func)(struct AVFormatContext *s, AVIOContext *pb);

int get_av_index_entry_flags(AVIndexEntry *ie) {
	return ie->flags;
}

void set_av_index_entry_flags(AVIndexEntry *ie, int v) {
	ie->flags = v;
}

int get_av_index_entry_size(AVIndexEntry *ie) {
	return ie->size;
}

void set_av_index_entry_size(AVIndexEntry *ie, int v) {
	ie->size = v;
}
*/
import "C"
import (
	"unsafe"
)

// AvGetPacket allocates and reads the payload of a packet and initialize its
// fields with default values.
func AvGetPacket(s *AVIOContext, pkt *AVPacket, size int32) int32 {
	return (int32)(C.av_get_packet((*C.struct_AVIOContext)(s),
		(*C.struct_AVPacket)(pkt), (C.int)(size)))
}

// Read data and append it to the current content of the AVPacket.
// If pkt->size is 0 this is identical to av_get_packet.
func AvAppendPacket(s *AVIOContext, pkt *AVPacket, size int32) int32 {
	return (int32)(C.av_append_packet((*C.struct_AVIOContext)(s),
		(*C.struct_AVPacket)(pkt), (C.int)(size)))
}

// AVCodceTag
type AVCodecTag C.struct_AVCodecTag

// AVProbeData
type AVProbeData C.struct_AVProbeData

// GetFilename gets `AVProbeData.filename` value.
func (pd *AVProbeData) GetFilename() string {
	return C.GoString(pd.filename)
}

// GetBuf gets `AVProbeData.buf` value.
func (pd *AVProbeData) GetBuf() *uint8 {
	return (*uint8)(pd.buf)
}

// SetBuf sets `AVProbeData.buf` value.
func (pd *AVProbeData) SetBuf(v *uint8) {
	pd.buf = (*C.uint8_t)(v)
}

// GetBufAddr gets `AVProbeData.buf` address.
func (pd *AVProbeData) GetBufAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&pd.buf))
}

// GetBufSize gets `AVProbeData.buf_size` value.
func (pd *AVProbeData) GetBufSize() int32 {
	return (int32)(pd.buf_size)
}

// SetBufSize sets `AVProbeData.buf_size` value.
func (pd *AVProbeData) SetBufSize(v int32) {
	pd.buf_size = (C.int)(v)
}

// GetBufSizeAddr gets `AVProbeData.buf_size` address.
func (pd *AVProbeData) GetBufSizeAddr() *int32 {
	return (*int32)(&pd.buf_size)
}

// GetMimeType gets `AVProbeData.mimetype` value.
func (pd *AVProbeData) GetMimeType() string {
	return C.GoString(pd.mime_type)
}

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

// AVOutputFormat
type AVOutputFormat C.struct_AVOutputFormat

// GetName gets `AVOutputFormat.name` value.
func (ofmt *AVOutputFormat) GetName() string {
	return C.GoString(ofmt.name)
}

// GetLongName gets `AVOutputFormat.long_name` value.
func (ofmt *AVOutputFormat) GetLongName() string {
	return C.GoString(ofmt.long_name)
}

// GetMimeType gets `AVOutputFormat.mimetype` value.
func (ofmt *AVOutputFormat) GetMimeType() string {
	return C.GoString(ofmt.mime_type)
}

// GetExtensions gets `AVOutputFormat.extensions` value.
func (ofmt *AVOutputFormat) GetExtensions() string {
	return C.GoString(ofmt.extensions)
}

// GetAudioCodec gets `AVOutputFormat.audio_codec` value.
func (ofmt *AVOutputFormat) GetAudioCodec() AVCodecID {
	return (AVCodecID)(ofmt.audio_codec)
}

// SetAudioCodec sets `AVOutputFormat.audio_codec` value.
func (ofmt *AVOutputFormat) SetAudioCodec(v AVCodecID) {
	ofmt.audio_codec = (C.enum_AVCodecID)(v)
}

// GetAudioCodecAddr gets `AVOutputFormat.audio_codec` address.
func (ofmt *AVOutputFormat) GetAudioCodecAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&ofmt.audio_codec))
}

// GetVideoCodec gets `AVOutputFormat.video_codec` value.
func (ofmt *AVOutputFormat) GetVideoCodec() AVCodecID {
	return (AVCodecID)(ofmt.video_codec)
}

// SetVideoCodec sets `AVOutputFormat.video_codec` value.
func (ofmt *AVOutputFormat) SetVideoCodec(v AVCodecID) {
	ofmt.video_codec = (C.enum_AVCodecID)(v)
}

// GetVideoCodecAddr gets `AVOutputFormat.video_codec` address.
func (ofmt *AVOutputFormat) GetVideoCodecAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&ofmt.video_codec))
}

// GetSubtitleCodec gets `AVOutputFormat.subtitle_codec` value.
func (ofmt *AVOutputFormat) GetSubtitleCodec() AVCodecID {
	return (AVCodecID)(ofmt.subtitle_codec)
}

// SetSubtitleCodec sets `AVOutputFormat.subtitle_codec` value.
func (ofmt *AVOutputFormat) SetSubtitleCodec(v AVCodecID) {
	ofmt.subtitle_codec = (C.enum_AVCodecID)(v)
}

// GetSubtitleCodecAddr gets `AVOutputFormat.subtitle_codec` address.
func (ofmt *AVOutputFormat) GetSubtitleCodecAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&ofmt.subtitle_codec))
}

// GetFlags gets `AVOutputFormat.flags` value.
func (ofmt *AVOutputFormat) GetFlags() int32 {
	return (int32)(ofmt.flags)
}

// SetFlags sets `AVOutputFormat.flags` value.
func (ofmt *AVOutputFormat) SetFlags(v int32) {
	ofmt.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVOutputFormat.flags` address.
func (ofmt *AVOutputFormat) GetFlagsAddr() *int32 {
	return (*int32)(&ofmt.flags)
}

// GetPrivClass gets `AVOutputFormat.priv_class` value.
func (ofmt *AVOutputFormat) GetPrivClass() *AVClass {
	return (*AVClass)(ofmt.priv_class)
}

// SetPrivClass sets `AVOutputFormat.priv_class` value.
func (ofmt *AVOutputFormat) SetPrivClass(v *AVClass) {
	ofmt.priv_class = (*C.struct_AVClass)(v)
}

// GetPrivClassAddr gets `AVOutputFormat.priv_class` address.
func (ofmt *AVOutputFormat) GetPrivClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&ofmt.priv_class))
}

// AVInputFormat
type AVInputFormat C.struct_AVInputFormat

// GetName gets `AVInputFormat.name` value.
func (ifmt *AVInputFormat) GetName() string {
	return C.GoString(ifmt.name)
}

// GetLongName gets `AVInputFormat.long_name` value.
func (ifmt *AVInputFormat) GetLongName() string {
	return C.GoString(ifmt.long_name)
}

// GetFlags gets `AVInputFormat.flags` value.
func (ifmt *AVInputFormat) GetFlags() int32 {
	return (int32)(ifmt.flags)
}

// SetFlags sets `AVInputFormat.flags` value.
func (ifmt *AVInputFormat) SetFlags(v int32) {
	ifmt.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVInputFormat.flags` address.
func (ifmt *AVInputFormat) GetFlagsAddr() *int32 {
	return (*int32)(&ifmt.flags)
}

// GetExtensions gets `AVInputFormat.extensions` value.
func (ifmt *AVInputFormat) GetExtensions() string {
	return C.GoString(ifmt.extensions)
}

// GetPrivClass gets `AVInputFormat.priv_class` value.
func (ifmt *AVInputFormat) GetPrivClass() *AVClass {
	return (*AVClass)(ifmt.priv_class)
}

// SetPrivClass sets `AVInputFormat.priv_class` value.
func (ifmt *AVInputFormat) SetPrivClass(v *AVClass) {
	ifmt.priv_class = (*C.struct_AVClass)(v)
}

// GetPrivClassAddr gets `AVInputFormat.priv_class` address.
func (ifmt *AVInputFormat) GetPrivClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&ifmt.priv_class))
}

// GetMimeType gets `AVInputFormat.mimetype` value.
func (ifmt *AVInputFormat) GetMimeType() string {
	return C.GoString(ifmt.mime_type)
}

// AVStreamParseType
type AVStreamParseType = C.enum_AVStreamParseType

const (
	AVSTREAM_PARSE_NONE       = AVStreamParseType(C.AVSTREAM_PARSE_NONE)
	AVSTREAM_PARSE_FULL       = AVStreamParseType(C.AVSTREAM_PARSE_FULL)
	AVSTREAM_PARSE_HEADERS    = AVStreamParseType(C.AVSTREAM_PARSE_HEADERS)
	AVSTREAM_PARSE_TIMESTAMPS = AVStreamParseType(C.AVSTREAM_PARSE_TIMESTAMPS)
	AVSTREAM_PARSE_FULL_ONCE  = AVStreamParseType(C.AVSTREAM_PARSE_FULL_ONCE)
	AVSTREAM_PARSE_FULL_RAW   = AVStreamParseType(C.AVSTREAM_PARSE_FULL_RAW)
)

// AVIndexEntry
type AVIndexEntry C.struct_AVIndexEntry

// GetPos gets `AVIndexEntry.pos` value.
func (ie *AVIndexEntry) GetPos() int64 {
	return (int64)(ie.pos)
}

// SetPos sets `AVIndexEntry.pos` value.
func (ie *AVIndexEntry) SetPos(v int64) {
	ie.pos = (C.int64_t)(v)
}

// GetPosAddr gets `AVIndexEntry.pos` address.
func (ie *AVIndexEntry) GetPosAddr() *int64 {
	return (*int64)(&ie.pos)
}

// GetTimestamp gets `AVIndexEntry.timestamp` value.
func (ie *AVIndexEntry) GetTimestamp() int64 {
	return (int64)(ie.timestamp)
}

// SetTimestamp sets `AVIndexEntry.timestamp` value.
func (ie *AVIndexEntry) SetTimestamp(v int64) {
	ie.timestamp = (C.int64_t)(v)
}

// GetTimestampAddr gets `AVIndexEntry.timestamp` address.
func (ie *AVIndexEntry) GetTimestampAddr() *int64 {
	return (*int64)(&ie.timestamp)
}

const (
	AVINDEX_KEYFRAME      = C.AVINDEX_KEYFRAME
	AVINDEX_DISCARD_FRAME = C.AVINDEX_DISCARD_FRAME
)

// GetFlags gets `AVIndexEntry.flags` value.
func (ie *AVIndexEntry) GetFlags() int32 {
	return (int32)(C.get_av_index_entry_flags((*C.struct_AVIndexEntry)(ie)))
}

// SetFlags sets `AVIndexEntry.flags` value.
func (ie *AVIndexEntry) SetFlags(v int32) {
	C.set_av_index_entry_flags((*C.struct_AVIndexEntry)(ie), (C.int)(v))
}

// GetSize gets `AVIndexEntry.size` value.
func (ie *AVIndexEntry) GetSize() int32 {
	return (int32)(C.get_av_index_entry_size((*C.struct_AVIndexEntry)(ie)))
}

// SetSize sets `AVIndexEntry.size` value.
func (ie *AVIndexEntry) SetSize(v int32) {
	C.set_av_index_entry_size((*C.struct_AVIndexEntry)(ie), (C.int)(v))
}

// GetMinDistance gets `AVIndexEntry.min_distance` value.
func (ie *AVIndexEntry) GetMinDistance() int32 {
	return (int32)(ie.min_distance)
}

// SetMinDistance sets `AVIndexEntry.min_distance` value.
func (ie *AVIndexEntry) SetMinDistance(v int32) {
	ie.min_distance = (C.int)(v)
}

// GetMinDistanceAddr gets `AVIndexEntry.min_distance` address.
func (ie *AVIndexEntry) GetMinDistanceAddr() *int32 {
	return (*int32)(&ie.min_distance)
}

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

// AVStreamInternal
type AVStreamInternal C.struct_AVStreamInternal

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

// AVStream
type AVStream C.struct_AVStream

// GetIndex gets `AVStream.index` value.
func (stm *AVStream) GetIndex() int32 {
	return (int32)(stm.index)
}

// SetIndex sets `AVStream.index` value.
func (stm *AVStream) SetIndex(v int32) {
	stm.index = (C.int)(v)
}

// GetIndexAddr gets `AVStream.index` address.
func (stm *AVStream) GetIndexAddr() *int32 {
	return (*int32)(&stm.index)
}

// GetId gets `AVStream.id` value.
func (stm *AVStream) GetId() int32 {
	return (int32)(stm.id)
}

// SetId sets `AVStream.id` value.
func (stm *AVStream) SetId(v int32) {
	stm.id = (C.int)(v)
}

// GetIdAddr gets `AVStream.id` address.
func (stm *AVStream) GetIdAddr() *int32 {
	return (*int32)(&stm.id)
}

// Deprecated: No use.
//
// GetCodec gets `AVStream.codec` value.
func (stm *AVStream) GetCodec() *AVCodecContext {
	return (*AVCodecContext)(stm.codec)
}

// Deprecated: No use.
//
// SetCodec sets `AVStream.codec` value.
func (stm *AVStream) SetCodec(v *AVCodecContext) {
	stm.codec = (*C.struct_AVCodecContext)(v)
}

// Deprecated: No use.
//
// GetCodecAddr gets `AVStream.codec` address.
func (stm *AVStream) GetCodecAddr() **AVCodecContext {
	return (**AVCodecContext)(unsafe.Pointer(&stm.codec))
}

// GetPrivData gets `AVStream.priv_data` value.
func (stm *AVStream) GetPrivData() unsafe.Pointer {
	return stm.priv_data
}

// SetPrivData sets `AVStream.priv_data` value.
func (stm *AVStream) SetPrivData(v CVoidPointer) {
	stm.priv_data = VoidPointer(v)
}

// GetPrivDataAddr gets `AVStream.priv_data` address.
func (stm *AVStream) GetPrivDataAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&stm.priv_data)
}

// GetTimeBase gets `AVStream.time_base` value.
func (stm *AVStream) GetTimeBase() AVRational {
	return (AVRational)(stm.time_base)
}

// SetTimeBase sets `AVStream.time_base` value.
func (stm *AVStream) SetTimeBase(v AVRational) {
	stm.time_base = (C.struct_AVRational)(v)
}

// GetTimeBaseAddr gets `AVStream.time_base` address.
func (stm *AVStream) GetTimeBaseAddr() *AVRational {
	return (*AVRational)(&stm.time_base)
}

// GetStartTime gets `AVStream.start_time` value.
func (stm *AVStream) GetStartTime() int64 {
	return (int64)(stm.start_time)
}

// SetStartTime sets `AVStream.start_time` value.
func (stm *AVStream) SetStartTime(v int64) {
	stm.start_time = (C.int64_t)(v)
}

// GetStartTimeAddr gets `AVStream.start_time` address.
func (stm *AVStream) GetStartTimeAddr() *int64 {
	return (*int64)(&stm.start_time)
}

// GetDuration gets `AVStream.duration` value.
func (stm *AVStream) GetDuration() int64 {
	return (int64)(stm.duration)
}

// SetDuration sets `AVStream.duration` value.
func (stm *AVStream) SetDuration(v int64) {
	stm.duration = (C.int64_t)(v)
}

// GetDurationAddr gets `AVStream.duration` address.
func (stm *AVStream) GetDurationAddr() *int64 {
	return (*int64)(&stm.duration)
}

// GetNbFrames gets `AVStream.nb_frames` value.
func (stm *AVStream) GetNbFrames() int64 {
	return (int64)(stm.nb_frames)
}

// SetNbFrames sets `AVStream.nb_frames` value.
func (stm *AVStream) SetNbFrames(v int64) {
	stm.nb_frames = (C.int64_t)(v)
}

// GetNbFramesAddr gets `AVStream.nb_frames` address.
func (stm *AVStream) GetNbFramesAddr() *int64 {
	return (*int64)(&stm.nb_frames)
}

// GetDisposition gets `AVStream.disposition` value.
func (stm *AVStream) GetDisposition() int32 {
	return (int32)(stm.disposition)
}

// SetDisposition sets `AVStream.disposition` value.
func (stm *AVStream) SetDisposition(v int32) {
	stm.disposition = (C.int)(v)
}

// GetDispositionAddr gets `AVStream.disposition` address.
func (stm *AVStream) GetDispositionAddr() *int32 {
	return (*int32)(&stm.disposition)
}

// GetDiscard gets `AVStream.discard` value.
func (stm *AVStream) GetDiscard() AVDiscard {
	return (AVDiscard)(stm.discard)
}

// SetDiscard sets `AVStream.discard` value.
func (stm *AVStream) SetDiscard(v AVDiscard) {
	stm.discard = (C.enum_AVDiscard)(v)
}

// GetDiscardAddr gets `AVStream.discard` address.
func (stm *AVStream) GetDiscardAddr() *AVDiscard {
	return (*AVDiscard)(&stm.discard)
}

// GetSampleAspectRatio gets `AVStream.sample_aspect_ratio` value.
func (stm *AVStream) GetSampleAspectRatio() AVRational {
	return (AVRational)(stm.sample_aspect_ratio)
}

// SetSampleAspectRatio sets `AVStream.sample_aspect_ratio` value.
func (stm *AVStream) SetSampleAspectRatio(v AVRational) {
	stm.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// GetSampleAspectRatioAddr gets `AVStream.sample_aspect_ratio` address.
func (stm *AVStream) GetSampleAspectRatioAddr() *AVRational {
	return (*AVRational)(&stm.sample_aspect_ratio)
}

// GetMetadata gets `AVStream.metadata` value.
func (stm *AVStream) GetMetadata() *AVDictionary {
	return (*AVDictionary)(stm.metadata)
}

// SetMetadata sets `AVStream.metadata` value.
func (stm *AVStream) SetMetadata(v *AVDictionary) {
	stm.metadata = (*C.struct_AVDictionary)(v)
}

// GetMetadataAddr gets `AVStream.metadata` address.
func (stm *AVStream) GetMetadataAddr() **AVDictionary {
	return (**AVDictionary)(unsafe.Pointer(&stm.metadata))
}

// GetAvgFrameRate gets `AVStream.avg_frame_rate` value.
func (stm *AVStream) GetAvgFrameRate() AVRational {
	return (AVRational)(stm.avg_frame_rate)
}

// SetAvgFrameRate sets `AVStream.avg_frame_rate` value.
func (stm *AVStream) SetAvgFrameRate(v AVRational) {
	stm.avg_frame_rate = (C.struct_AVRational)(v)
}

// GetAvgFrameRateAddr gets `AVStream.avg_frame_rate` address.
func (stm *AVStream) GetAvgFrameRateAddr() *AVRational {
	return (*AVRational)(&stm.avg_frame_rate)
}

// GetAttachedPic gets `AVStream.attached_pic` value.
func (stm *AVStream) GetAttachedPic() AVPacket {
	return (AVPacket)(stm.attached_pic)
}

// SetAttachedPic sets `AVStream.attached_pic` value.
func (stm *AVStream) SetAttachedPic(v AVPacket) {
	stm.attached_pic = (C.struct_AVPacket)(v)
}

// GetAttachedPicAddr gets `AVStream.attached_pic` address.
func (stm *AVStream) GetAttachedPicAddr() *AVPacket {
	return (*AVPacket)(&stm.attached_pic)
}

// GetSideData gets `AVStream.side_data` value.
func (stm *AVStream) GetSideData() *AVPacketSideData {
	return (*AVPacketSideData)(stm.side_data)
}

// SetSideData sets `AVStream.side_data` value.
func (stm *AVStream) SetSideData(v *AVPacketSideData) {
	stm.side_data = (*C.struct_AVPacketSideData)(v)
}

// GetSideDataAddr gets `AVStream.side_data` address.
func (stm *AVStream) GetSideDataAddr() **AVPacketSideData {
	return (**AVPacketSideData)(unsafe.Pointer(&stm.side_data))
}

// GetNbSideData gets `AVStream.nb_side_data` value.
func (stm *AVStream) GetNbSideData() int32 {
	return (int32)(stm.nb_side_data)
}

// SetNbSideData sets `AVStream.nb_side_data` value.
func (stm *AVStream) SetNbSideData(v int32) {
	stm.nb_side_data = (C.int)(v)
}

// GetNbSideDataAddr gets `AVStream.nb_side_data` address.
func (stm *AVStream) GetNbSideDataAddr() *int32 {
	return (*int32)(&stm.nb_side_data)
}

// GetEventFlags gets `AVStream.event_flags` value.
func (stm *AVStream) GetEventFlags() int32 {
	return (int32)(stm.event_flags)
}

// SetEventFlags sets `AVStream.event_flags` value.
func (stm *AVStream) SetEventFlags(v int32) {
	stm.event_flags = (C.int)(v)
}

// GetEventFlagsAddr gets `AVStream.event_flags` address.
func (stm *AVStream) GetEventFlagsAddr() *int32 {
	return (*int32)(&stm.event_flags)
}

const (
	AVSTREAM_EVENT_FLAG_METADATA_UPDATED = int32(C.AVSTREAM_EVENT_FLAG_METADATA_UPDATED)
)

// GetRFrameRate gets `AVStream.r_frame_rate` value.
func (stm *AVStream) GetRFrameRate() AVRational {
	return (AVRational)(stm.r_frame_rate)
}

// SetRFrameRate sets `AVStream.r_frame_rate` value.
func (stm *AVStream) SetRFrameRate(v AVRational) {
	stm.r_frame_rate = (C.struct_AVRational)(v)
}

// GetRFrameRateAddr gets `AVStream.r_frame_rate` address.
func (stm *AVStream) GetRFrameRateAddr() *AVRational {
	return (*AVRational)(&stm.r_frame_rate)
}

// Deprecated: Unused.
//
// GetRecommendedEncoderConfiguration gets `AVStream.recommended_encoder_configuration` value.
func (stm *AVStream) GetRecommendedEncoderConfiguration() string {
	return C.GoString(stm.recommended_encoder_configuration)
}

// GetCodecpar gets `AVStream.codecpar` value.
func (stm *AVStream) GetCodecpar() *AVCodecParameters {
	return (*AVCodecParameters)(stm.codecpar)
}

// SetCodecpar sets `AVStream.codecpar` value.
func (stm *AVStream) SetCodecpar(v *AVCodecParameters) {
	stm.codecpar = (*C.struct_AVCodecParameters)(v)
}

// GetCodecparAddr gets `AVStream.codecpar` address.
func (stm *AVStream) GetCodecparAddr() **AVCodecParameters {
	return (**AVCodecParameters)(unsafe.Pointer(&stm.codecpar))
}

// Deprecated: No use.
//
// AvStreamGetRFrameRate
func AvStreamGetRFrameRate(s *AVStream) AVRational {
	return (AVRational)(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

// Deprecated: No use.
//
// AvStreamSetRFrameRate
func AvStreamSetRFrameRate(s *AVStream, r AVRational) {
	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), (C.struct_AVRational)(r))
}

// Deprecated: No use.
//
// AvStreamGetRecommendedEncoderConfiguration
func AvStreamGetRecommendedEncoderConfiguration(s *AVStream) string {
	return C.GoString(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(s)))
}

// Deprecated: No use.
//
// AvStreamSetRecommendedEncoderConfiguration
func AvStreamSetRecommendedEncoderConfiguration(s *AVStream, configuration string) {
	configurationPtr, configurationFunc := StringCasting(configuration)
	defer configurationFunc()
	C.av_stream_set_recommended_encoder_configuration((*C.struct_AVStream)(s), (*C.char)(configurationPtr))
}

// AvStreamGetParser
func AvStreamGetParser(s *AVStream) *AVCodecParserContext {
	return (*AVCodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

// AvStreamGetEndPts returns the pts of the last muxed packet + its duration.
// the retuned value is undefined when used with a demuxer.
func AvStreamGetEndPts(s *AVStream) int64 {
	return (int64)(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}

const (
	AV_PROGRAM_RUNNING = C.AV_PROGRAM_RUNNING
)

// AVProgram
type AVProgram C.struct_AVProgram

// GetId gets `AVProgram.id` value.
func (pgm *AVProgram) GetId() int32 {
	return (int32)(pgm.id)
}

// SetId sets `AVProgram.id` value.
func (pgm *AVProgram) SetId(v int32) {
	pgm.id = (C.int)(v)
}

// GetIdAddr gets `AVProgram.id` address.
func (pgm *AVProgram) GetIdAddr() *int32 {
	return (*int32)(&pgm.id)
}

// GetFlags gets `AVProgram.flags` value.
func (pgm *AVProgram) GetFlags() int32 {
	return (int32)(pgm.flags)
}

// SetFlags sets `AVProgram.flags` value.
func (pgm *AVProgram) SetFlags(v int32) {
	pgm.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVProgram.flags` address.
func (pgm *AVProgram) GetFlagsAddr() *int32 {
	return (*int32)(&pgm.flags)
}

// GetDiscard gets `AVProgram.discard` value.
func (pgm *AVProgram) GetDiscard() AVDiscard {
	return (AVDiscard)(pgm.discard)
}

// SetDiscard sets `AVProgram.discard` value.
func (pgm *AVProgram) SetDiscard(v AVDiscard) {
	pgm.discard = (C.enum_AVDiscard)(v)
}

// GetDiscardAddr gets `AVProgram.discard` address.
func (pgm *AVProgram) GetDiscardAddr() *AVDiscard {
	return (*AVDiscard)(unsafe.Pointer(&pgm.discard))
}

// GetStreamIndex gets `AVProgram.stream_index` value.
func (pgm *AVProgram) GetStreamIndex() (v []uint32) {
	if pgm.stream_index == nil {
		return v
	}
	return unsafe.Slice((*uint32)(unsafe.Pointer(pgm.stream_index)), pgm.nb_stream_indexes)
}

// SetStreamIndex sets `AVProgram.stream_index` value.
func (pgm *AVProgram) SetStreamIndex(v *uint32) {
	pgm.stream_index = (*C.uint)(v)
}

// GetStreamIndexAddr gets `AVProgram.stream_index` address.
func (pgm *AVProgram) GetStreamIndexAddr() **uint32 {
	return (**uint32)(unsafe.Pointer(&pgm.stream_index))
}

// GetNbStreamIndexes gets `AVProgram.nb_stream_indexes` value.
func (pgm *AVProgram) GetNbStreamIndexes() uint32 {
	return (uint32)(pgm.nb_stream_indexes)
}

// SetNbStreamIndexes sets `AVProgram.nb_stream_indexes` value.
func (pgm *AVProgram) SetNbStreamIndexes(v uint32) {
	pgm.nb_stream_indexes = (C.uint)(v)
}

// GetNbStreamIndexesAddr gets `AVProgram.nb_stream_indexes` address.
func (pgm *AVProgram) GetNbStreamIndexesAddr() *uint32 {
	return (*uint32)(&pgm.nb_stream_indexes)
}

// GetMetadata gets `AVProgram.metadata` value.
func (pgm *AVProgram) GetMetadata() *AVDictionary {
	return (*AVDictionary)(pgm.metadata)
}

// SetMetadata sets `AVProgram.metadata` value.
func (pgm *AVProgram) SetMetadata(v *AVDictionary) {
	pgm.metadata = (*C.struct_AVDictionary)(v)
}

// GetMetadataAddr gets `AVProgram.metadata` address.
func (pgm *AVProgram) GetMetadataAddr() **AVDictionary {
	return (**AVDictionary)(unsafe.Pointer(&pgm.metadata))
}

// GetProgramNum gets `AVProgram.program_num` value.
func (pgm *AVProgram) GetProgramNum() int32 {
	return (int32)(pgm.program_num)
}

// SetProgramNum sets `AVProgram.program_num` value.
func (pgm *AVProgram) SetProgramNum(v int32) {
	pgm.program_num = (C.int)(v)
}

// GetProgramNumAddr gets `AVProgram.program_num` address.
func (pgm *AVProgram) GetProgramNumAddr() *int32 {
	return (*int32)(&pgm.program_num)
}

// GetPmtPid gets `AVProgram.pmt_pid` value.
func (pgm *AVProgram) GetPmtPid() int32 {
	return (int32)(pgm.pmt_pid)
}

// SetPmtPid sets `AVProgram.pmt_pid` value.
func (pgm *AVProgram) SetPmtPid(v int32) {
	pgm.pmt_pid = (C.int)(v)
}

// GetPmtPidAddr gets `AVProgram.pmt_pid` address.
func (pgm *AVProgram) GetPmtPidAddr() *int32 {
	return (*int32)(&pgm.pmt_pid)
}

// GetPcrPid gets `AVProgram.pcr_pid` value.
func (pgm *AVProgram) GetPcrPid() int32 {
	return (int32)(pgm.pcr_pid)
}

// SetPcrPid sets `AVProgram.pcr_pid` value.
func (pgm *AVProgram) SetPcrPid(v int32) {
	pgm.pcr_pid = (C.int)(v)
}

// GetPcrPidAddr gets `AVProgram.pcr_pid` address.
func (pgm *AVProgram) GetPcrPidAddr() *int32 {
	return (*int32)(&pgm.pcr_pid)
}

// GetPmtVersion gets `AVProgram.pmt_version` value.
func (pgm *AVProgram) GetPmtVersion() int32 {
	return (int32)(pgm.pmt_version)
}

// SetPmtVersion sets `AVProgram.pmt_version` value.
func (pgm *AVProgram) SetPmtVersion(v int32) {
	pgm.pmt_version = (C.int)(v)
}

// GetPmtVersionAddr gets `AVProgram.pmt_version` address.
func (pgm *AVProgram) GetPmtVersionAddr() *int32 {
	return (*int32)(&pgm.pmt_version)
}

const (
	AVFMTCTX_NOHEADER   = C.AVFMTCTX_NOHEADER
	AVFMTCTX_UNSEEKABLE = C.AVFMTCTX_UNSEEKABLE
)

// AVChapter
type AVChapter = C.struct_AVChapter

// GetId gets `AVChapter.id` value.
func (cpt *AVChapter) GetId() int32 {
	return (int32)(cpt.id)
}

// SetId sets `AVChapter.id` value.
func (cpt *AVChapter) SetId(v int32) {
	cpt.id = (C.int)(v)
}

// GetIdAddr gets `AVChapter.id` address.
func (cpt *AVChapter) GetIdAddr() *int32 {
	return (*int32)(&cpt.id)
}

// GetTimeBase gets `AVChapter.time_base` value.
func (cpt *AVChapter) GetTimeBase() AVRational {
	return (AVRational)(cpt.time_base)
}

// SetTimeBase sets `AVChapter.time_base` value.
func (cpt *AVChapter) SetTimeBase(v AVRational) {
	cpt.time_base = (C.struct_AVRational)(v)
}

// GetTimeBaseAddr gets `AVChapter.time_base` address.
func (cpt *AVChapter) GetTimeBaseAddr() *AVRational {
	return (*AVRational)(&cpt.time_base)
}

// GetStart gets `AVChapter.start` value.
func (cpt *AVChapter) GetStart() int64 {
	return (int64)(cpt.start)
}

// SetStart sets `AVChapter.start` value.
func (cpt *AVChapter) SetStart(v int64) {
	cpt.start = (C.int64_t)(v)
}

// GetStartAddr gets `AVChapter.start` address.
func (cpt *AVChapter) GetStartAddr() *int64 {
	return (*int64)(&cpt.start)
}

// GetEnd gets `AVChapter.end` value.
func (cpt *AVChapter) GetEnd() int64 {
	return (int64)(cpt.end)
}

// SetEnd sets `AVChapter.end` value.
func (cpt *AVChapter) SetEnd(v int64) {
	cpt.end = (C.int64_t)(v)
}

// GetEndAddr gets `AVChapter.end` address.
func (cpt *AVChapter) GetEndAddr() *int64 {
	return (*int64)(&cpt.end)
}

// GetMetadata gets `AVChapter.metadata` value.
func (cpt *AVChapter) GetMetadata() *AVDictionary {
	return (*AVDictionary)(cpt.metadata)
}

// SetMetadata sets `AVChapter.metadata` value.
func (cpt *AVChapter) SetMetadata(v *AVDictionary) {
	cpt.metadata = (*C.struct_AVDictionary)(v)
}

// GetMetadataAddr gets `AVChapter.metadata` address.
func (cpt *AVChapter) GetMetadataAddr() **AVDictionary {
	return (**AVDictionary)(unsafe.Pointer(&cpt.metadata))
}

// typedef int (*av_format_control_message)(struct AVFormatContext *s, int type,
// void *data, size_t data_size);
type AVFormatControlMessageFunc C.av_format_control_message

// typedef int (*AVOpenCallback)(struct AVFormatContext *s,
// AVIOContext **pb, const char *url, int flags,
// const AVIOInterruptCB *int_cb, AVDictionary **options);
type AVOpenCallbackFunc C.AVOpenCallback

// AVDurationEstimationMethod
type AVDurationEstimationMethod = C.enum_AVDurationEstimationMethod

const (
	AVFMT_DURATION_FROM_PTS     = AVDurationEstimationMethod(C.AVFMT_DURATION_FROM_PTS)
	AVFMT_DURATION_FROM_STREAM  = AVDurationEstimationMethod(C.AVFMT_DURATION_FROM_STREAM)
	AVFMT_DURATION_FROM_BITRATE = AVDurationEstimationMethod(C.AVFMT_DURATION_FROM_BITRATE)
)

// AVFormatInternal
type AVFormatInternal C.struct_AVFormatInternal

// Format I/O context.
type AVFormatContext C.struct_AVFormatContext

// GetAvClass gets `AVFormatContext.av_class` value.
func (s *AVFormatContext) GetAvClass() *AVClass {
	return (*AVClass)(s.av_class)
}

// SetAvClass sets `AVFormatContext.av_class` value.
func (s *AVFormatContext) SetAvClass(v *AVClass) {
	s.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVFormatContext.av_class` address.
func (s *AVFormatContext) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&s.av_class))
}

// GetIformat gets `AVFormatContext.iformat` value.
func (s *AVFormatContext) GetIformat() *AVInputFormat {
	return (*AVInputFormat)(s.iformat)
}

// SetIformat sets `AVFormatContext.iformat` value.
func (s *AVFormatContext) SetIformat(v *AVInputFormat) {
	s.iformat = (*C.struct_AVInputFormat)(v)
}

// GetIformatAddr gets `AVFormatContext.iformat` address.
func (s *AVFormatContext) GetIformatAddr() **AVInputFormat {
	return (**AVInputFormat)(unsafe.Pointer(&s.iformat))
}

// GetOformat gets `AVFormatContext.oformat` value.
func (s *AVFormatContext) GetOformat() *AVOutputFormat {
	return (*AVOutputFormat)(s.oformat)
}

// SetOformat sets `AVFormatContext.oformat` value.
func (s *AVFormatContext) SetOformat(v *AVOutputFormat) {
	s.oformat = (*C.struct_AVOutputFormat)(v)
}

// GetOformatAddr gets `AVFormatContext.oformat` address.
func (s *AVFormatContext) GetOformatAddr() **AVOutputFormat {
	return (**AVOutputFormat)(unsafe.Pointer(&s.oformat))
}

// GetPrivData gets `AVFormatContext.priv_data` value.
func (s *AVFormatContext) GetPrivData() unsafe.Pointer {
	return s.priv_data
}

// SetPrivData sets `AVFormatContext.priv_data` value.
func (s *AVFormatContext) SetPrivData(v CVoidPointer) {
	s.priv_data = VoidPointer(v)
}

// GetPrivDataAddr gets `AVFormatContext.priv_data` address.
func (s *AVFormatContext) GetPrivDataAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&s.priv_data)
}

// GetPb gets `AVFormatContext.pb` value.
func (s *AVFormatContext) GetPb() *AVIOContext {
	return (*AVIOContext)(s.pb)
}

// SetPb sets `AVFormatContext.pb` value.
func (s *AVFormatContext) SetPb(v *AVIOContext) {
	s.pb = (*C.struct_AVIOContext)(v)
}

// GetPbAddr gets `AVFormatContext.pb` address.
func (s *AVFormatContext) GetPbAddr() **AVIOContext {
	return (**AVIOContext)(unsafe.Pointer(&s.pb))
}

// GetCtxFlags gets `AVFormatContext.ctx_flags` value.
func (s *AVFormatContext) GetCtxFlags() int32 {
	return (int32)(s.ctx_flags)
}

// SetCtxFlags sets `AVFormatContext.ctx_flags` value.
func (s *AVFormatContext) SetCtxFlags(v int32) {
	s.ctx_flags = (C.int)(v)
}

// GetCtxFlagsAddr gets `AVFormatContext.ctx_flags` address.
func (s *AVFormatContext) GetCtxFlagsAddr() *int32 {
	return (*int32)(&s.ctx_flags)
}

// GetNbStreams gets `AVFormatContext.nb_streams` value.
func (s *AVFormatContext) GetNbStreams() uint32 {
	return (uint32)(s.nb_streams)
}

// SetNbStreams sets `AVFormatContext.nb_streams` value.
func (s *AVFormatContext) SetNbStreams(v uint32) {
	s.nb_streams = (C.uint)(v)
}

// GetNbStreamsAddr gets `AVFormatContext.nb_streams` address.
func (s *AVFormatContext) GetNbStreamsAddr() *uint32 {
	return (*uint32)(&s.nb_streams)
}

// GetStreams gets `AVFormatContext.streams` value.
func (s *AVFormatContext) GetStreams() (v []*AVStream) {
	if s.streams == nil {
		return v
	}
	return unsafe.Slice((**AVStream)(unsafe.Pointer(s.streams)), s.nb_streams)
}

// SetStreams sets `AVFormatContext.streams` value.
func (s *AVFormatContext) SetStreams(v **AVStream) {
	s.streams = (**C.struct_AVStream)(unsafe.Pointer(v))
}

// GetStreamsAddr gets `AVFormatContext.streams` address.
func (s *AVFormatContext) GetStreamsAddr() ***AVStream {
	return (***AVStream)(unsafe.Pointer(&s.streams))
}

// Deprecated: Use url instead.
//
// GetFilename gets `AVFormatContext.filename` value.
func (s *AVFormatContext) GetFilename() string {
	return C.GoString((*C.char)(&s.filename[0]))
}

// GetUrl gets `AVFormatContext.url` value.
func (s *AVFormatContext) GetUrl() string {
	return C.GoString(s.url)
}

// SetUrl sets `AVFormatContext.url` value.
func (s *AVFormatContext) SetUrl(v string) {
	vPtr, _ := StringCasting(v)
	if s.url != nil {
		FreePointer(s.url)
	}
	s.url = (*C.char)(vPtr)
}

// GetStartTime gets `AVFormatContext.start_time` value.
func (s *AVFormatContext) GetStartTime() int64 {
	return (int64)(s.start_time)
}

// SetStartTime sets `AVFormatContext.start_time` value.
func (s *AVFormatContext) SetStartTime(v int64) {
	s.start_time = (C.int64_t)(v)
}

// GetStartTimeAddr gets `AVFormatContext.start_time` address.
func (s *AVFormatContext) GetStartTimeAddr() *int64 {
	return (*int64)(&s.start_time)
}

// GetDuration gets `AVFormatContext.duration` value.
func (s *AVFormatContext) GetDuration() int64 {
	return (int64)(s.duration)
}

// SetDuration sets `AVFormatContext.duration` value.
func (s *AVFormatContext) SetDuration(v int64) {
	s.duration = (C.int64_t)(v)
}

// GetDurationAddr gets `AVFormatContext.duration` address.
func (s *AVFormatContext) GetDurationAddr() *int64 {
	return (*int64)(&s.duration)
}

// GetBitRate gets `AVFormatContext.bit_rate` value.
func (s *AVFormatContext) GetBitRate() int64 {
	return (int64)(s.bit_rate)
}

// SetBitRate sets `AVFormatContext.bit_rate` value.
func (s *AVFormatContext) SetBitRate(v int64) {
	s.bit_rate = (C.int64_t)(v)
}

// GetBitRateAddr gets `AVFormatContext.bit_rate` address.
func (s *AVFormatContext) GetBitRateAddr() *int64 {
	return (*int64)(&s.bit_rate)
}

// GetPacketSize gets `AVFormatContext.packet_size` value.
func (s *AVFormatContext) GetPacketSize() uint32 {
	return (uint32)(s.packet_size)
}

// SetPacketSize sets `AVFormatContext.packet_size` value.
func (s *AVFormatContext) SetPacketSize(v uint32) {
	s.packet_size = (C.uint)(v)
}

// GetPacketSizeAddr gets `AVFormatContext.packet_size` address.
func (s *AVFormatContext) GetPacketSizeAddr() *uint32 {
	return (*uint32)(&s.packet_size)
}

// GetMaxDelay gets `AVFormatContext.max_delay` value.
func (s *AVFormatContext) GetMaxDelay() int32 {
	return (int32)(s.max_delay)
}

// SetMaxDelay sets `AVFormatContext.max_delay` value.
func (s *AVFormatContext) SetMaxDelay(v int32) {
	s.max_delay = (C.int)(v)
}

// GetMaxDelayAddr gets `AVFormatContext.max_delay` address.
func (s *AVFormatContext) GetMaxDelayAddr() *int32 {
	return (*int32)(&s.max_delay)
}

// GetFlags gets `AVFormatContext.flags` value.
func (s *AVFormatContext) GetFlags() int32 {
	return (int32)(s.flags)
}

// SetFlags sets `AVFormatContext.flags` value.
func (s *AVFormatContext) SetFlags(v int32) {
	s.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVFormatContext.flags` address.
func (s *AVFormatContext) GetFlagsAddr() *int32 {
	return (*int32)(&s.flags)
}

// GetProbesize gets `AVFormatContext.probesize` value.
func (s *AVFormatContext) GetProbesize() int64 {
	return (int64)(s.probesize)
}

// SetProbesize sets `AVFormatContext.probesize` value.
func (s *AVFormatContext) SetProbesize(v int64) {
	s.probesize = (C.int64_t)(v)
}

// GetProbesizeAddr gets `AVFormatContext.probesize` address.
func (s *AVFormatContext) GetProbesizeAddr() *int64 {
	return (*int64)(&s.probesize)
}

// GetMaxAnalyzeDuration gets `AVFormatContext.max_analyze_duration` value.
func (s *AVFormatContext) GetMaxAnalyzeDuration() int64 {
	return (int64)(s.max_analyze_duration)
}

// SetMaxAnalyzeDuration sets `AVFormatContext.max_analyze_duration` value.
func (s *AVFormatContext) SetMaxAnalyzeDuration(v int64) {
	s.max_analyze_duration = (C.int64_t)(v)
}

// GetMaxAnalyzeDurationAddr gets `AVFormatContext.max_analyze_duration` address.
func (s *AVFormatContext) GetMaxAnalyzeDurationAddr() *int64 {
	return (*int64)(&s.max_analyze_duration)
}

// GetKey gets `AVFormatContext.key` value.
func (s *AVFormatContext) GetKey() *uint8 {
	return (*uint8)(s.key)
}

// SetKey sets `AVFormatContext.key` value.
func (s *AVFormatContext) SetKey(v *uint8) {
	s.key = (*C.uint8_t)(v)
}

// GetKeyAddr gets `AVFormatContext.key` address.
func (s *AVFormatContext) GetKeyAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&s.key))
}

// GetKeylen gets `AVFormatContext.keylen` value.
func (s *AVFormatContext) GetKeylen() int32 {
	return (int32)(s.keylen)
}

// SetKeylen sets `AVFormatContext.keylen` value.
func (s *AVFormatContext) SetKeylen(v int32) {
	s.keylen = (C.int)(v)
}

// GetKeylenAddr gets `AVFormatContext.keylen` address.
func (s *AVFormatContext) GetKeylenAddr() *int32 {
	return (*int32)(&s.keylen)
}

// GetNbPrograms gets `AVFormatContext.nb_programs` value.
func (s *AVFormatContext) GetNbPrograms() uint32 {
	return (uint32)(s.nb_programs)
}

// SetNbPrograms sets `AVFormatContext.nb_programs` value.
func (s *AVFormatContext) SetNbPrograms(v uint32) {
	s.nb_programs = (C.uint)(v)
}

// GetNbProgramsAddr gets `AVFormatContext.nb_programs` address.
func (s *AVFormatContext) GetNbProgramsAddr() *uint32 {
	return (*uint32)(&s.nb_programs)
}

// GetPrograms gets `AVFormatContext.programs` value.
func (s *AVFormatContext) GetPrograms() (v []*AVProgram) {
	if s.programs == nil {
		return v
	}
	return unsafe.Slice((**AVProgram)(unsafe.Pointer(s.programs)), s.nb_programs)
}

// SetPrograms sets `AVFormatContext.programs` value.
func (s *AVFormatContext) SetPrograms(v **AVProgram) {
	s.programs = (**C.struct_AVProgram)(unsafe.Pointer(v))
}

// GetProgramsAddr gets `AVFormatContext.programs` address.
func (s *AVFormatContext) GetProgramsAddr() ***AVProgram {
	return (***AVProgram)(unsafe.Pointer(&s.programs))
}

// GetVideoCodecId gets `AVFormatContext.video_codec_id` value.
func (s *AVFormatContext) GetVideoCodecId() AVCodecID {
	return (AVCodecID)(s.video_codec_id)
}

// SetVideoCodecId sets `AVFormatContext.video_codec_id` value.
func (s *AVFormatContext) SetVideoCodecId(v AVCodecID) {
	s.video_codec_id = (C.enum_AVCodecID)(v)
}

// GetVideoCodecIdAddr gets `AVFormatContext.video_codec_id` address.
func (s *AVFormatContext) GetVideoCodecIdAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&s.video_codec_id))
}

// GetAudioCodecId gets `AVFormatContext.audio_codec_id` value.
func (s *AVFormatContext) GetAudioCodecId() AVCodecID {
	return (AVCodecID)(s.audio_codec_id)
}

// SetAudioCodecId sets `AVFormatContext.audio_codec_id` value.
func (s *AVFormatContext) SetAudioCodecId(v AVCodecID) {
	s.audio_codec_id = (C.enum_AVCodecID)(v)
}

// GetAudioCodecIdAddr gets `AVFormatContext.audio_codec_id` address.
func (s *AVFormatContext) GetAudioCodecIdAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&s.audio_codec_id))
}

// GetSubtitleCodecId gets `AVFormatContext.subtitle_codec_id` value.
func (s *AVFormatContext) GetSubtitleCodecId() AVCodecID {
	return (AVCodecID)(s.subtitle_codec_id)
}

// SetSubtitleCodecId sets `AVFormatContext.subtitle_codec_id` value.
func (s *AVFormatContext) SetSubtitleCodecId(v AVCodecID) {
	s.subtitle_codec_id = (C.enum_AVCodecID)(v)
}

// GetSubtitleCodecIdAddr gets `AVFormatContext.subtitle_codec_id` address.
func (s *AVFormatContext) GetSubtitleCodecIdAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&s.subtitle_codec_id))
}

// GetMaxIndexSize gets `AVFormatContext.max_index_size` value.
func (s *AVFormatContext) GetMaxIndexSize() uint32 {
	return (uint32)(s.max_index_size)
}

// SetMaxIndexSize sets `AVFormatContext.max_index_size` value.
func (s *AVFormatContext) SetMaxIndexSize(v uint32) {
	s.max_index_size = (C.uint)(v)
}

// GetMaxIndexSizeAddr gets `AVFormatContext.max_index_size` address.
func (s *AVFormatContext) GetMaxIndexSizeAddr() *uint32 {
	return (*uint32)(&s.max_index_size)
}

// GetMaxPictureBuffer gets `AVFormatContext.max_picture_buffer` value.
func (s *AVFormatContext) GetMaxPictureBuffer() uint32 {
	return (uint32)(s.max_picture_buffer)
}

// SetMaxPictureBuffer sets `AVFormatContext.max_picture_buffer` value.
func (s *AVFormatContext) SetMaxPictureBuffer(v uint32) {
	s.max_picture_buffer = (C.uint)(v)
}

// GetMaxPictureBufferAddr gets `AVFormatContext.max_picture_buffer` address.
func (s *AVFormatContext) GetMaxPictureBufferAddr() *uint32 {
	return (*uint32)(&s.max_picture_buffer)
}

// GetNbChapters gets `AVFormatContext.nb_chapters` value.
func (s *AVFormatContext) GetNbChapters() uint32 {
	return (uint32)(s.nb_chapters)
}

// SetNbChapters sets `AVFormatContext.nb_chapters` value.
func (s *AVFormatContext) SetNbChapters(v uint32) {
	s.nb_chapters = (C.uint)(v)
}

// GetNbChaptersAddr gets `AVFormatContext.nb_chapters` address.
func (s *AVFormatContext) GetNbChaptersAddr() *uint32 {
	return (*uint32)(&s.nb_chapters)
}

// GetChapters gets `AVFormatContext.chapters` value.
func (s *AVFormatContext) GetChapters() (v []*AVChapter) {
	if s.chapters == nil {
		return v
	}
	return unsafe.Slice((**AVChapter)(unsafe.Pointer(s.chapters)), s.nb_chapters)
}

// SetChapters sets `AVFormatContext.chapters` value.
func (s *AVFormatContext) SetChapters(v **AVChapter) {
	s.chapters = (**C.struct_AVChapter)(v)
}

// GetChaptersAddr gets `AVFormatContext.chapters` address.
func (s *AVFormatContext) GetChaptersAddr() ***AVChapter {
	return (***AVChapter)(&s.chapters)
}

// GetMetadata gets `AVFormatContext.metadata` value.
func (s *AVFormatContext) GetMetadata() *AVDictionary {
	return (*AVDictionary)(s.metadata)
}

// SetMetadata sets `AVFormatContext.metadata` value.
func (s *AVFormatContext) SetMetadata(v *AVDictionary) {
	s.metadata = (*C.struct_AVDictionary)(v)
}

// GetMetadataAddr gets `AVFormatContext.metadata` address.
func (s *AVFormatContext) GetMetadataAddr() **AVDictionary {
	return (**AVDictionary)(unsafe.Pointer(&s.metadata))
}

// GetStartTimeRealtime gets `AVFormatContext.start_time_realtime` value.
func (s *AVFormatContext) GetStartTimeRealtime() int64 {
	return (int64)(s.start_time_realtime)
}

// SetStartTimeRealtime sets `AVFormatContext.start_time_realtime` value.
func (s *AVFormatContext) SetStartTimeRealtime(v int64) {
	s.start_time_realtime = (C.int64_t)(v)
}

// GetStartTimeRealtimeAddr gets `AVFormatContext.start_time_realtime` address.
func (s *AVFormatContext) GetStartTimeRealtimeAddr() *int64 {
	return (*int64)(&s.start_time_realtime)
}

// GetFpsProbeSize gets `AVFormatContext.fps_probe_size` value.
func (s *AVFormatContext) GetFpsProbeSize() int32 {
	return (int32)(s.fps_probe_size)
}

// SetFpsProbeSize sets `AVFormatContext.fps_probe_size` value.
func (s *AVFormatContext) SetFpsProbeSize(v int32) {
	s.fps_probe_size = (C.int)(v)
}

// GetFpsProbeSizeAddr gets `AVFormatContext.fps_probe_size` address.
func (s *AVFormatContext) GetFpsProbeSizeAddr() *int32 {
	return (*int32)(&s.fps_probe_size)
}

// GetErrorRecognition gets `AVFormatContext.error_recognition` value.
func (s *AVFormatContext) GetErrorRecognition() int32 {
	return (int32)(s.error_recognition)
}

// SetErrorRecognition sets `AVFormatContext.error_recognition` value.
func (s *AVFormatContext) SetErrorRecognition(v int32) {
	s.error_recognition = (C.int)(v)
}

// GetErrorRecognitionAddr gets `AVFormatContext.error_recognition` address.
func (s *AVFormatContext) GetErrorRecognitionAddr() *int32 {
	return (*int32)(&s.error_recognition)
}

// GetInterruptCallback gets `AVFormatContext.interrupt_callback` value.
func (s *AVFormatContext) GetInterruptCallback() AVIOInterruptCB {
	return (AVIOInterruptCB)(s.interrupt_callback)
}

// SetInterruptCallback sets `AVFormatContext.interrupt_callback` value.
func (s *AVFormatContext) SetInterruptCallback(v AVIOInterruptCB) {
	s.interrupt_callback = (C.AVIOInterruptCB)(v)
}

// GetInterruptCallbackAddr gets `AVFormatContext.interrupt_callback` address.
func (s *AVFormatContext) GetInterruptCallbackAddr() *AVIOInterruptCB {
	return (*AVIOInterruptCB)(&s.interrupt_callback)
}

// GetDebug gets `AVFormatContext.debug` value.
func (s *AVFormatContext) GetDebug() int32 {
	return (int32)(s.debug)
}

// SetDebug sets `AVFormatContext.debug` value.
func (s *AVFormatContext) SetDebug(v int32) {
	s.debug = (C.int)(v)
}

// GetDebugAddr gets `AVFormatContext.debug` address.
func (s *AVFormatContext) GetDebugAddr() *int32 {
	return (*int32)(&s.debug)
}

// GetMaxInterleaveDelta gets `AVFormatContext.max_interleave_delta` value.
func (s *AVFormatContext) GetMaxInterleaveDelta() int64 {
	return (int64)(s.max_interleave_delta)
}

// SetMaxInterleaveDelta sets `AVFormatContext.max_interleave_delta` value.
func (s *AVFormatContext) SetMaxInterleaveDelta(v int64) {
	s.max_interleave_delta = (C.int64_t)(v)
}

// GetMaxInterleaveDeltaAddr gets `AVFormatContext.max_interleave_delta` address.
func (s *AVFormatContext) GetMaxInterleaveDeltaAddr() *int64 {
	return (*int64)(&s.max_interleave_delta)
}

// GetStrictStdCompliance gets `AVFormatContext.strict_std_compliance` value.
func (s *AVFormatContext) GetStrictStdCompliance() int32 {
	return (int32)(s.strict_std_compliance)
}

// SetStrictStdCompliance sets `AVFormatContext.strict_std_compliance` value.
func (s *AVFormatContext) SetStrictStdCompliance(v int32) {
	s.strict_std_compliance = (C.int)(v)
}

// GetStrictStdComplianceAddr gets `AVFormatContext.strict_std_compliance` address.
func (s *AVFormatContext) GetStrictStdComplianceAddr() *int32 {
	return (*int32)(&s.strict_std_compliance)
}

// GetEventFlags gets `AVFormatContext.event_flags` value.
func (s *AVFormatContext) GetEventFlags() int32 {
	return (int32)(s.event_flags)
}

// SetEventFlags sets `AVFormatContext.event_flags` value.
func (s *AVFormatContext) SetEventFlags(v int32) {
	s.event_flags = (C.int)(v)
}

// GetEventFlagsAddr gets `AVFormatContext.event_flags` address.
func (s *AVFormatContext) GetEventFlagsAddr() *int32 {
	return (*int32)(&s.event_flags)
}

// GetMaxTsProbe gets `AVFormatContext.max_ts_probe` value.
func (s *AVFormatContext) GetMaxTsProbe() int32 {
	return (int32)(s.max_ts_probe)
}

// SetMaxTsProbe sets `AVFormatContext.max_ts_probe` value.
func (s *AVFormatContext) SetMaxTsProbe(v int32) {
	s.max_ts_probe = (C.int)(v)
}

// GetMaxTsProbeAddr gets `AVFormatContext.max_ts_probe` address.
func (s *AVFormatContext) GetMaxTsProbeAddr() *int32 {
	return (*int32)(&s.max_ts_probe)
}

// GetAvoidNegativeTs gets `AVFormatContext.avoid_negative_ts` value.
func (s *AVFormatContext) GetAvoidNegativeTs() int32 {
	return (int32)(s.avoid_negative_ts)
}

// SetAvoidNegativeTs sets `AVFormatContext.avoid_negative_ts` value.
func (s *AVFormatContext) SetAvoidNegativeTs(v int32) {
	s.avoid_negative_ts = (C.int)(v)
}

// GetAvoidNegativeTsAddr gets `AVFormatContext.avoid_negative_ts` address.
func (s *AVFormatContext) GetAvoidNegativeTsAddr() *int32 {
	return (*int32)(&s.avoid_negative_ts)
}

// GetTsId gets `AVFormatContext.ts_id` value.
func (s *AVFormatContext) GetTsId() int32 {
	return (int32)(s.ts_id)
}

// SetTsId sets `AVFormatContext.ts_id` value.
func (s *AVFormatContext) SetTsId(v int32) {
	s.ts_id = (C.int)(v)
}

// GetTsIdAddr gets `AVFormatContext.ts_id` address.
func (s *AVFormatContext) GetTsIdAddr() *int32 {
	return (*int32)(&s.ts_id)
}

// GetAudioPreload gets `AVFormatContext.audio_preload` value.
func (s *AVFormatContext) GetAudioPreload() int32 {
	return (int32)(s.audio_preload)
}

// SetAudioPreload sets `AVFormatContext.audio_preload` value.
func (s *AVFormatContext) SetAudioPreload(v int32) {
	s.audio_preload = (C.int)(v)
}

// GetAudioPreloadAddr gets `AVFormatContext.audio_preload` address.
func (s *AVFormatContext) GetAudioPreloadAddr() *int32 {
	return (*int32)(&s.audio_preload)
}

// GetMaxChunkDuration gets `AVFormatContext.max_chunk_duration` value.
func (s *AVFormatContext) GetMaxChunkDuration() int32 {
	return (int32)(s.max_chunk_duration)
}

// SetMaxChunkDuration sets `AVFormatContext.max_chunk_duration` value.
func (s *AVFormatContext) SetMaxChunkDuration(v int32) {
	s.max_chunk_duration = (C.int)(v)
}

// GetMaxChunkDurationAddr gets `AVFormatContext.max_chunk_duration` address.
func (s *AVFormatContext) GetMaxChunkDurationAddr() *int32 {
	return (*int32)(&s.max_chunk_duration)
}

// GetMaxChunkSize gets `AVFormatContext.max_chunk_size` value.
func (s *AVFormatContext) GetMaxChunkSize() int32 {
	return (int32)(s.max_chunk_size)
}

// SetMaxChunkSize sets `AVFormatContext.max_chunk_size` value.
func (s *AVFormatContext) SetMaxChunkSize(v int32) {
	s.max_chunk_size = (C.int)(v)
}

// GetMaxChunkSizeAddr gets `AVFormatContext.max_chunk_size` address.
func (s *AVFormatContext) GetMaxChunkSizeAddr() *int32 {
	return (*int32)(&s.max_chunk_size)
}

// GetUseWallclockAsTimestamps gets `AVFormatContext.use_wallclock_as_timestamps` value.
func (s *AVFormatContext) GetUseWallclockAsTimestamps() int32 {
	return (int32)(s.use_wallclock_as_timestamps)
}

// SetUseWallclockAsTimestamps sets `AVFormatContext.use_wallclock_as_timestamps` value.
func (s *AVFormatContext) SetUseWallclockAsTimestamps(v int32) {
	s.use_wallclock_as_timestamps = (C.int)(v)
}

// GetUseWallclockAsTimestampsAddr gets `AVFormatContext.use_wallclock_as_timestamps` address.
func (s *AVFormatContext) GetUseWallclockAsTimestampsAddr() *int32 {
	return (*int32)(&s.use_wallclock_as_timestamps)
}

// GetAvioFlags gets `AVFormatContext.avio_flags` value.
func (s *AVFormatContext) GetAvioFlags() int32 {
	return (int32)(s.avio_flags)
}

// SetAvioFlags sets `AVFormatContext.avio_flags` value.
func (s *AVFormatContext) SetAvioFlags(v int32) {
	s.avio_flags = (C.int)(v)
}

// GetAvioFlagsAddr gets `AVFormatContext.avio_flags` address.
func (s *AVFormatContext) GetAvioFlagsAddr() *int32 {
	return (*int32)(&s.avio_flags)
}

// GetDurationEstimationMethod gets `AVFormatContext.duration_estimation_method` value.
func (s *AVFormatContext) GetDurationEstimationMethod() AVDurationEstimationMethod {
	return (AVDurationEstimationMethod)(s.duration_estimation_method)
}

// SetDurationEstimationMethod sets `AVFormatContext.duration_estimation_method` value.
func (s *AVFormatContext) SetDurationEstimationMethod(v AVDurationEstimationMethod) {
	s.duration_estimation_method = (C.enum_AVDurationEstimationMethod)(v)
}

// GetDurationEstimationMethodAddr gets `AVFormatContext.duration_estimation_method` address.
func (s *AVFormatContext) GetDurationEstimationMethodAddr() *AVDurationEstimationMethod {
	return (*AVDurationEstimationMethod)(unsafe.Pointer(&s.duration_estimation_method))
}

// GetSkipInitialBytes gets `AVFormatContext.skip_initial_bytes` value.
func (s *AVFormatContext) GetSkipInitialBytes() int64 {
	return (int64)(s.skip_initial_bytes)
}

// SetSkipInitialBytes sets `AVFormatContext.skip_initial_bytes` value.
func (s *AVFormatContext) SetSkipInitialBytes(v int64) {
	s.skip_initial_bytes = (C.int64_t)(v)
}

// GetSkipInitialBytesAddr gets `AVFormatContext.skip_initial_bytes` address.
func (s *AVFormatContext) GetSkipInitialBytesAddr() *int64 {
	return (*int64)(&s.skip_initial_bytes)
}

// GetCorrectTsOverflow gets `AVFormatContext.correct_ts_overflow` value.
func (s *AVFormatContext) GetCorrectTsOverflow() uint32 {
	return (uint32)(s.correct_ts_overflow)
}

// SetCorrectTsOverflow sets `AVFormatContext.correct_ts_overflow` value.
func (s *AVFormatContext) SetCorrectTsOverflow(v uint32) {
	s.correct_ts_overflow = (C.uint)(v)
}

// GetCorrectTsOverflowAddr gets `AVFormatContext.correct_ts_overflow` address.
func (s *AVFormatContext) GetCorrectTsOverflowAddr() *uint32 {
	return (*uint32)(&s.correct_ts_overflow)
}

// GetSeek2any gets `AVFormatContext.seek2any` value.
func (s *AVFormatContext) GetSeek2any() int32 {
	return (int32)(s.seek2any)
}

// SetSeek2any sets `AVFormatContext.seek2any` value.
func (s *AVFormatContext) SetSeek2any(v int32) {
	s.seek2any = (C.int)(v)
}

// GetSeek2anyAddr gets `AVFormatContext.seek2any` address.
func (s *AVFormatContext) GetSeek2anyAddr() *int32 {
	return (*int32)(&s.seek2any)
}

// GetFlushPackets gets `AVFormatContext.flush_packets` value.
func (s *AVFormatContext) GetFlushPackets() int32 {
	return (int32)(s.flush_packets)
}

// SetFlushPackets sets `AVFormatContext.flush_packets` value.
func (s *AVFormatContext) SetFlushPackets(v int32) {
	s.flush_packets = (C.int)(v)
}

// GetFlushPacketsAddr gets `AVFormatContext.flush_packets` address.
func (s *AVFormatContext) GetFlushPacketsAddr() *int32 {
	return (*int32)(&s.flush_packets)
}

// GetProbeScore gets `AVFormatContext.probe_score` value.
func (s *AVFormatContext) GetProbeScore() int32 {
	return (int32)(s.probe_score)
}

// SetProbeScore sets `AVFormatContext.probe_score` value.
func (s *AVFormatContext) SetProbeScore(v int32) {
	s.probe_score = (C.int)(v)
}

// GetProbeScoreAddr gets `AVFormatContext.probe_score` address.
func (s *AVFormatContext) GetProbeScoreAddr() *int32 {
	return (*int32)(&s.probe_score)
}

// GetFormatProbesize gets `AVFormatContext.format_probesize` value.
func (s *AVFormatContext) GetFormatProbesize() int32 {
	return (int32)(s.format_probesize)
}

// SetFormatProbesize sets `AVFormatContext.format_probesize` value.
func (s *AVFormatContext) SetFormatProbesize(v int32) {
	s.format_probesize = (C.int)(v)
}

// GetFormatProbesizeAddr gets `AVFormatContext.format_probesize` address.
func (s *AVFormatContext) GetFormatProbesizeAddr() *int32 {
	return (*int32)(&s.format_probesize)
}

// GetCodecWhitelist gets `AVFormatContext.codec_whitelist` value.
func (s *AVFormatContext) GetCodecWhitelist() string {
	return C.GoString(s.codec_whitelist)
}

// GetFormatWhitelist gets `AVFormatContext.format_whitelist` value.
func (s *AVFormatContext) GetFormatWhitelist() string {
	return C.GoString(s.format_whitelist)
}

// GetInternal gets `AVFormatContext.internal` value.
func (s *AVFormatContext) GetInternal() *AVFormatInternal {
	return (*AVFormatInternal)(s.internal)
}

// SetInternal sets `AVFormatContext.internal` value.
func (s *AVFormatContext) SetInternal(v *AVFormatInternal) {
	s.internal = (*C.struct_AVFormatInternal)(v)
}

// GetInternalAddr gets `AVFormatContext.internal` address.
func (s *AVFormatContext) GetInternalAddr() **AVFormatInternal {
	return (**AVFormatInternal)(unsafe.Pointer(&s.internal))
}

// GetIoRepositioned gets `AVFormatContext.io_repositioned` value.
func (s *AVFormatContext) GetIoRepositioned() int32 {
	return (int32)(s.io_repositioned)
}

// SetIoRepositioned sets `AVFormatContext.io_repositioned` value.
func (s *AVFormatContext) SetIoRepositioned(v int32) {
	s.io_repositioned = (C.int)(v)
}

// GetIoRepositionedAddr gets `AVFormatContext.io_repositioned` address.
func (s *AVFormatContext) GetIoRepositionedAddr() *int32 {
	return (*int32)(&s.io_repositioned)
}

// GetVideoCodec gets `AVFormatContext.video_codec` value.
func (s *AVFormatContext) GetVideoCodec() *AVCodec {
	return (*AVCodec)(s.video_codec)
}

// SetVideoCodec sets `AVFormatContext.video_codec` value.
func (s *AVFormatContext) SetVideoCodec(v *AVCodec) {
	s.video_codec = (*C.struct_AVCodec)(v)
}

// GetVideoCodecAddr gets `AVFormatContext.video_codec` address.
func (s *AVFormatContext) GetVideoCodecAddr() **AVCodec {
	return (**AVCodec)(unsafe.Pointer(&s.video_codec))
}

// GetAudioCodec gets `AVFormatContext.audio_codec` value.
func (s *AVFormatContext) GetAudioCodec() *AVCodec {
	return (*AVCodec)(s.audio_codec)
}

// SetAudioCodec sets `AVFormatContext.audio_codec` value.
func (s *AVFormatContext) SetAudioCodec(v *AVCodec) {
	s.audio_codec = (*C.struct_AVCodec)(v)
}

// GetAudioCodecAddr gets `AVFormatContext.audio_codec` address.
func (s *AVFormatContext) GetAudioCodecAddr() **AVCodec {
	return (**AVCodec)(unsafe.Pointer(&s.audio_codec))
}

// GetSubtitleCodec gets `AVFormatContext.subtitle_codec` value.
func (s *AVFormatContext) GetSubtitleCodec() *AVCodec {
	return (*AVCodec)(s.subtitle_codec)
}

// SetSubtitleCodec sets `AVFormatContext.subtitle_codec` value.
func (s *AVFormatContext) SetSubtitleCodec(v *AVCodec) {
	s.subtitle_codec = (*C.struct_AVCodec)(v)
}

// GetSubtitleCodecAddr gets `AVFormatContext.subtitle_codec` address.
func (s *AVFormatContext) GetSubtitleCodecAddr() **AVCodec {
	return (**AVCodec)(unsafe.Pointer(&s.subtitle_codec))
}

// GetDataCodec gets `AVFormatContext.data_codec` value.
func (s *AVFormatContext) GetDataCodec() *AVCodec {
	return (*AVCodec)(s.data_codec)
}

// SetDataCodec sets `AVFormatContext.data_codec` value.
func (s *AVFormatContext) SetDataCodec(v *AVCodec) {
	s.data_codec = (*C.struct_AVCodec)(v)
}

// GetDataCodecAddr gets `AVFormatContext.data_codec` address.
func (s *AVFormatContext) GetDataCodecAddr() **AVCodec {
	return (**AVCodec)(unsafe.Pointer(&s.data_codec))
}

// GetMetadataHeaderPadding gets `AVFormatContext.metadata_header_padding` value.
func (s *AVFormatContext) GetMetadataHeaderPadding() int32 {
	return (int32)(s.metadata_header_padding)
}

// SetMetadataHeaderPadding sets `AVFormatContext.metadata_header_padding` value.
func (s *AVFormatContext) SetMetadataHeaderPadding(v int32) {
	s.metadata_header_padding = (C.int)(v)
}

// GetMetadataHeaderPaddingAddr gets `AVFormatContext.metadata_header_padding` address.
func (s *AVFormatContext) GetMetadataHeaderPaddingAddr() *int32 {
	return (*int32)(&s.metadata_header_padding)
}

// GetOpaque gets `AVFormatContext.opaque` value.
func (s *AVFormatContext) GetOpaque() unsafe.Pointer {
	return s.opaque
}

// SetOpaque sets `AVFormatContext.opaque` value.
func (s *AVFormatContext) SetOpaque(v CVoidPointer) {
	s.opaque = VoidPointer(v)
}

// GetOpaqueAddr gets `AVFormatContext.opaque` address.
func (s *AVFormatContext) GetOpaqueAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&s.opaque)
}

// GetOutputTsOffset gets `AVFormatContext.output_ts_offset` value.
func (s *AVFormatContext) GetOutputTsOffset() int64 {
	return (int64)(s.output_ts_offset)
}

// SetOutputTsOffset sets `AVFormatContext.output_ts_offset` value.
func (s *AVFormatContext) SetOutputTsOffset(v int64) {
	s.output_ts_offset = (C.int64_t)(v)
}

// GetOutputTsOffsetAddr gets `AVFormatContext.output_ts_offset` address.
func (s *AVFormatContext) GetOutputTsOffsetAddr() *int64 {
	return (*int64)(&s.output_ts_offset)
}

// GetDumpSeparator gets `AVFormatContext.dump_separator` value.
func (s *AVFormatContext) GetDumpSeparator() *uint8 {
	return (*uint8)(s.dump_separator)
}

// SetDumpSeparator sets `AVFormatContext.dump_separator` value.
func (s *AVFormatContext) SetDumpSeparator(v *uint8) {
	s.dump_separator = (*C.uint8_t)(v)
}

// GetDumpSeparatorAddr gets `AVFormatContext.dump_separator` address.
func (s *AVFormatContext) GetDumpSeparatorAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&s.dump_separator))
}

// GetDataCodecId gets `AVFormatContext.data_codec_id` value.
func (s *AVFormatContext) GetDataCodecId() AVCodecID {
	return (AVCodecID)(s.data_codec_id)
}

// SetDataCodecId sets `AVFormatContext.data_codec_id` value.
func (s *AVFormatContext) SetDataCodecId(v AVCodecID) {
	s.data_codec_id = (C.enum_AVCodecID)(v)
}

// GetDataCodecIdAddr gets `AVFormatContext.data_codec_id` address.
func (s *AVFormatContext) GetDataCodecIdAddr() *AVCodecID {
	return (*AVCodecID)(unsafe.Pointer(&s.data_codec_id))
}

// typedef int (*av_format_context_open_cb)(struct AVFormatContext *s, AVIOContext **p, const char *url, int flags,
// const AVIOInterruptCB *int_cb, AVDictionary **options);
type AvFormatContextOpenCb = C.av_format_context_open_cb

// Deprecated: Use io_open and io_close.
//
// GetOpenCb gets `AVFormatContext.open_cb` value.
func (s *AVFormatContext) GetOpenCb() AvFormatContextOpenCb {
	return (AvFormatContextOpenCb)(s.open_cb)
}

// Deprecated: Use io_open and io_close.
//
// SetOpenCb sets `AVFormatContext.open_cb` value.
func (s *AVFormatContext) SetOpenCb(v AvFormatContextOpenCb) {
	s.open_cb = (C.av_format_context_open_cb)(v)
}

// Deprecated: Use io_open and io_close.
//
// GetOpenCbAddr gets `AVFormatContext.open_cb` address.
func (s *AVFormatContext) GetOpenCbAddr() *AvFormatContextOpenCb {
	return (*AvFormatContextOpenCb)(&s.open_cb)
}

// GetProtocolWhitelist gets `AVFormatContext.protocol_whitelist` value.
func (s *AVFormatContext) GetProtocolWhitelist() string {
	return C.GoString(s.protocol_whitelist)
}

// typedef int (*av_format_context_io_open_func)(struct AVFormatContext *s, AVIOContext **pb, const char *url,
// int flags, AVDictionary **options);
type AvFormatContextIoOpenFunc = C.av_format_context_io_open_func

// typedef void (*av_format_context_io_close_func)(struct AVFormatContext *s, AVIOContext *pb);
type AvFormatContextIoCloseFunc = C.av_format_context_io_close_func

// GetIoOpen gets `AVFormatContext.io_open` value.
func (s *AVFormatContext) GetIoOpen() AvFormatContextIoOpenFunc {
	return (AvFormatContextIoOpenFunc)(s.io_open)
}

// SetIoOpen sets `AVFormatContext.io_open` value.
func (s *AVFormatContext) SetIoOpen(v AvFormatContextIoOpenFunc) {
	s.io_open = (C.av_format_context_io_open_func)(v)
}

// GetIoOpenAddr gets `AVFormatContext.io_open` address.
func (s *AVFormatContext) GetIoOpenAddr() *AvFormatContextIoOpenFunc {
	return (*AvFormatContextIoOpenFunc)(&s.io_open)
}

// GetIoClose gets `AVFormatContext.io_close` value.
func (s *AVFormatContext) GetIoClose() AvFormatContextIoCloseFunc {
	return (AvFormatContextIoCloseFunc)(s.io_close)
}

// SetIoClose sets `AVFormatContext.io_close` value.
func (s *AVFormatContext) SetIoClose(v AvFormatContextIoCloseFunc) {
	s.io_close = (C.av_format_context_io_close_func)(v)
}

// GetIoCloseAddr gets `AVFormatContext.io_close` address.
func (s *AVFormatContext) GetIoCloseAddr() *AvFormatContextIoCloseFunc {
	return (*AvFormatContextIoCloseFunc)(&s.io_close)
}

// GetProtocolBlacklist gets `AVFormatContext.protocol_blacklist` value.
func (s *AVFormatContext) GetProtocolBlacklist() string {
	return C.GoString(s.protocol_blacklist)
}

// GetMaxStreams gets `AVFormatContext.max_streams` value.
func (s *AVFormatContext) GetMaxStreams() int32 {
	return (int32)(s.max_streams)
}

// SetMaxStreams sets `AVFormatContext.max_streams` value.
func (s *AVFormatContext) SetMaxStreams(v int32) {
	s.max_streams = (C.int)(v)
}

// GetMaxStreamsAddr gets `AVFormatContext.max_streams` address.
func (s *AVFormatContext) GetMaxStreamsAddr() *int32 {
	return (*int32)(&s.max_streams)
}

// GetSkipEstimateDurationFromPts gets `AVFormatContext.skip_estimate_duration_from_pts` value.
func (s *AVFormatContext) GetSkipEstimateDurationFromPts() int32 {
	return (int32)(s.skip_estimate_duration_from_pts)
}

// SetSkipEstimateDurationFromPts sets `AVFormatContext.skip_estimate_duration_from_pts` value.
func (s *AVFormatContext) SetSkipEstimateDurationFromPts(v int32) {
	s.skip_estimate_duration_from_pts = (C.int)(v)
}

// GetSkipEstimateDurationFromPtsAddr gets `AVFormatContext.skip_estimate_duration_from_pts` address.
func (s *AVFormatContext) GetSkipEstimateDurationFromPtsAddr() *int32 {
	return (*int32)(&s.skip_estimate_duration_from_pts)
}

// GetMaxProbePackets gets `AVFormatContext.max_probe_packets` value.
func (s *AVFormatContext) GetMaxProbePackets() int32 {
	return (int32)(s.max_probe_packets)
}

// SetMaxProbePackets sets `AVFormatContext.max_probe_packets` value.
func (s *AVFormatContext) SetMaxProbePackets(v int32) {
	s.max_probe_packets = (C.int)(v)
}

// GetMaxProbePacketsAddr gets `AVFormatContext.max_probe_packets` address.
func (s *AVFormatContext) GetMaxProbePacketsAddr() *int32 {
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

// Deprecated: No use.
//
// AvFormatGetProbeScore
func AvFormatGetProbeScore(s *AVFormatContext) int32 {
	return (int32)(C.av_format_get_probe_score((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use.
//
// AvFormatGetVideoCodec
func AvFormatGetVideoCodec(s *AVFormatContext) *AVCodec {
	return (*AVCodec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use.
//
// AvFormatSetVideoCodec
func AvFormatSetVideoCodec(s *AVFormatContext, c *AVCodec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// Deprecated: No use.
//
// AvFormatGetAudioCodec
func AvFormatGetAudioCodec(s *AVFormatContext) *AVCodec {
	return (*AVCodec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use.
//
// AvFormatSetAudioCodec
func AvFormatSetAudioCodec(s *AVFormatContext, c *AVCodec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// Deprecated: No use.
//
// AvFormatGetSubtitleCodec
func AvFormatGetSubtitleCodec(s *AVFormatContext) *AVCodec {
	return (*AVCodec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use.
//
// AvFormatSetSubtitleCodec
func AvFormatSetSubtitleCodec(s *AVFormatContext, c *AVCodec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// Deprecated: No use.
//
// AvFormatGetDataCodec
func AvFormatGetDataCodec(s *AVFormatContext) *AVCodec {
	return (*AVCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use.
//
// AvFormatSetDataCodec
func AvFormatSetDataCodec(s *AVFormatContext, c *AVCodec) {
	C.av_format_set_data_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// Deprecated: No use.
//
// AvFormatGetOpaque
func AvFormatGetOpaque(s *AVFormatContext) unsafe.Pointer {
	return C.av_format_get_opaque((*C.struct_AVFormatContext)(s))
}

// Deprecated: No use.
//
// AvFormatSetOpaque
func AvFormatSetOpaque(s *AVFormatContext, opaque CVoidPointer) {
	C.av_format_set_opaque((*C.struct_AVFormatContext)(s), VoidPointer(opaque))
}

// Deprecated: No use.
//
// AvFormatGetControlMessageCb
func AvFormatGetControlMessageCb(s *AVFormatContext) AVFormatControlMessageFunc {
	return (AVFormatControlMessageFunc)(C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use.
//
// AvFormatSetControlMessageCb
func AvFormatSetControlMessageCb(s *AVFormatContext, callback AVFormatControlMessageFunc) {
	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(s), (C.av_format_control_message)(callback))
}

// Deprecated: No use.
//
// AvFormatGetOpenCb
func AvFormatGetOpenCb(s *AVFormatContext) AVOpenCallbackFunc {
	return (AVOpenCallbackFunc)(C.av_format_get_open_cb((*C.struct_AVFormatContext)(s)))
}

// Deprecated: No use.
//
// AvFormatSetOpenCb
func AvFormatSetOpenCb(s *AVFormatContext, callback AVOpenCallbackFunc) {
	C.av_format_set_open_cb((*C.struct_AVFormatContext)(s), (C.AVOpenCallback)(callback))
}

// AvFormatInjectGlobalSideData will cause global side data to be injected in the next packet
// of each stream as well as after any subsequent seek.
func AvFormatInjectGlobalSideData(s *AVFormatContext) {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(s))
}

// AvFmtCtxGetDurationEstimationMethod returns the method used to set ctx->duration.
func AvFmtCtxGetDurationEstimationMethod(s *AVFormatContext) AVDurationEstimationMethod {
	return (AVDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(s)))
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

// Deprecated: No use.
//
// AvRegisterAll
func AvRegisterAll() {
	C.av_register_all()
}

// Deprecated: No use.
//
// AvRegisterInputFormat
func AvRegisterInputFormat(format *AVInputFormat) {
	C.av_register_input_format((*C.struct_AVInputFormat)(format))
}

// Deprecated: No use.
//
// AvRegisterOutputFormat
func AvRegisterOutputFormat(format *AVOutputFormat) {
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

// Deprecated: No use.
//
// AvIformatNext returns the first registered input format if f is NULL,
// returns the next registered input format after f if f is non-NULL,
// or NULL if f is the last one.
func AvIformatNext(f *AVInputFormat) *AVInputFormat {
	return (*AVInputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

// Deprecated: No use.
//
// AvOformatNext returns the first registered output format if f is NULL,
// returns the next registered output format after f if f is non-NULL,
// or NULL if f is the last one.
func AvOformatNext(f *AVOutputFormat) *AVOutputFormat {
	return (*AVOutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

// AvMuxerIterate iterates over all registered muxers.
func AvMuxerIterate(opaque CVoidPointerPointer) *AVOutputFormat {
	return (*AVOutputFormat)(C.av_muxer_iterate(VoidPointerPointer(opaque)))
}

// AvDemuxerIterate iterates over all registered demuxers.
func AvDemuxerIterate(opaque CVoidPointerPointer) *AVInputFormat {
	return (*AVInputFormat)(C.av_demuxer_iterate(VoidPointerPointer(opaque)))
}

// AvFormatAllocContext allocates an AVFormatContext.
func AvFormatAllocContext() *AVFormatContext {
	return (*AVFormatContext)(C.avformat_alloc_context())
}

// AvFormatFreeContext frees an AVFormatContext and all its streams.
func AvFormatFreeContext(s *AVFormatContext) {
	C.avformat_free_context((*C.struct_AVFormatContext)(s))
}

// AvFormatGetClass gets the AVClass for AVFormatContext. It can be used in combination with
// Av_OPT_SEARCH_FAKE_OBJ for examining options.
func AvFormatGetClass() *AVClass {
	return (*AVClass)(C.avformat_get_class())
}

// AvFormatNewStream adds a new stream to a media file.
func AvFormatNewStream(s *AVFormatContext, c *AVCodec) *AVStream {
	return (*AVStream)(C.avformat_new_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c)))
}

// AvStreamAddSideData wraps an existing array as stream side data.
func AvStreamAddSideData(st *AVStream, _type AVPacketSideDataType, data *uint8, size uintptr) int32 {
	return (int32)(C.av_stream_add_side_data((*C.struct_AVStream)(st),
		(C.enum_AVPacketSideDataType)(_type), (*C.uint8_t)(data), (C.size_t)(size)))
}

// AvStreamNewSideData allocates new information from stream.
func AvStreamNewSideData(st *AVStream, _type AVPacketSideDataType, size int32) *uint8 {
	return (*uint8)(C.av_stream_new_side_data((*C.struct_AVStream)(st),
		(C.enum_AVPacketSideDataType)(_type), (C.int)(size)))
}

// AvStreamGetSideData gets side information from stream.
func AvStreamGetSideData(st *AVStream, _type AVPacketSideDataType, size *int32) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(st),
		(C.enum_AVPacketSideDataType)(_type), (*C.int)(size)))
}

// AvNewProgram
func AvNewProgram(s *AVFormatContext, id int32) *AVProgram {
	return (*AVProgram)(C.av_new_program((*C.struct_AVFormatContext)(s), (C.int)(id)))
}

// AvFormatAllocOutputContext2 allocates an AVFormatContext for an output format.
func AvFormatAllocOutputContext2(ctx **AVFormatContext,
	oformat *AVOutputFormat, formatName, filename string) int32 {
	formatNamePtr, formatNameFunc := StringCasting(formatName)
	defer formatNameFunc()
	filenamePtr, filenameFunc := StringCasting(filename)
	defer filenameFunc()
	return (int32)(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(ctx)),
		(*C.struct_AVOutputFormat)(oformat), (*C.char)(formatNamePtr), (*C.char)(filenamePtr)))
}

// AvFindInputFormat finds AVInputFormat based on the short name of the input format.
func AvFindInputFormat(shortName string) *AVInputFormat {
	shortNamePtr, shortNameFunc := StringCasting(shortName)
	defer shortNameFunc()
	return (*AVInputFormat)(C.av_find_input_format((*C.char)(shortNamePtr)))
}

// AvProbeInputFormat guesses the file format.
func AvProbeInputFormat(pd *AVProbeData, isOpened int32) *AVInputFormat {
	return (*AVInputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), (C.int)(isOpened)))
}

// AvProbeInputFormat2 guesses the file format.
func AvProbeInputFormat2(pd *AVProbeData, isOpened int32, scoreMax *int32) *AVInputFormat {
	return (*AVInputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd),
		(C.int)(isOpened), (*C.int)(scoreMax)))
}

// AvProbeInputFormat3 guesses the file format.
func AvProbeInputFormat3(pd *AVProbeData, isOpened int32, scoreRet *int32) *AVInputFormat {
	return (*AVInputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd),
		(C.int)(isOpened), (*C.int)(scoreRet)))
}

// AvProbeInputBuffer2 probes a bytestream to determine the input format. Each time a probe returns
// with a score that is too low, the probe buffer size is increased and another
// attempt is made. When the maximum probe size is reached, the input format
// with the highest score is returned.
func AvProbeInputBuffer2(pb *AVIOContext, fmt **AVInputFormat,
	url string, logctx CVoidPointer, offset, maxProbeSize uint32) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb), (**C.AVInputFormat)(unsafe.Pointer(fmt)),
		(*C.char)(urlPtr), VoidPointer(logctx), (C.uint)(offset), (C.uint)(maxProbeSize)))
}

// AvProbeInputBuffer likes AvProbeInputBuffer2() but returns 0 on success.
func AvProbeInputBuffer(pb *AVIOContext, fmt **AVInputFormat,
	url string, logctx CVoidPointer, offset, maxProbeSize uint32) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.av_probe_input_buffer((*C.struct_AVIOContext)(pb), (**C.AVInputFormat)(unsafe.Pointer(fmt)),
		(*C.char)(urlPtr), VoidPointer(logctx), (C.uint)(offset), (C.uint)(maxProbeSize)))
}

// AvFormatOpenInput open an input stream and read the header. The codecs are not opened.
// The stream must be closed with AvFormatCloseInput().
func AvFormatOpenInput(ps **AVFormatContext, url string, fmt *AVInputFormat, options **AVDictionary) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)),
		(*C.char)(urlPtr), (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvDemuxerOpen
func AvDemuxerOpen(ic *AVFormatContext) int32 {
	return (int32)(C.av_demuxer_open((*C.struct_AVFormatContext)(ic)))
}

// AvFormatFindStreamInfo reads packets of a media file to get stream information.
func AvFormatFindStreamInfo(ic *AVFormatContext, options **AVDictionary) int32 {
	return (int32)(C.avformat_find_stream_info((*C.struct_AVFormatContext)(ic),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvFindProgramFromStream finds the programs which belong to a given stream.
func AvFindProgramFromStream(ic *AVFormatContext, last *AVProgram, s int32) *AVProgram {
	return (*AVProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(ic),
		(*C.struct_AVProgram)(last), (C.int)(s)))
}

// AvProgramAddStreamIndex
func AvProgramAddStreamIndex(ic *AVFormatContext, progid int32, idx uint32) {
	C.av_program_add_stream_index((*C.struct_AVFormatContext)(ic), (C.int)(progid), (C.uint)(idx))
}

// AvFindBestStream finds the "best" stream in the file.
func AvFindBestStream(ic *AVFormatContext, _type AVMediaType,
	wantedStreamNb, relatedStream int32,
	decoderRet **AVCodec, flags int32) int32 {
	return (int32)(C.av_find_best_stream((*C.struct_AVFormatContext)(ic),
		(C.enum_AVMediaType)(_type), (C.int)(wantedStreamNb), (C.int)(relatedStream),
		(**C.struct_AVCodec)(unsafe.Pointer(decoderRet)), (C.int)(flags)))
}

// AvReadFrame returns the next frame of a stream.
func AvReadFrame(ic *AVFormatContext, pkt *AVPacket) int32 {
	return (int32)(C.av_read_frame((*C.struct_AVFormatContext)(ic), (*C.struct_AVPacket)(pkt)))
}

// AvSeekFrame seeks to the keyframe at timestamp.
func AvSeekFrame(ic *AVFormatContext, streamIndex int32, timestamp int64, flags int32) int32 {
	return (int32)(C.av_seek_frame((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex), (C.int64_t)(timestamp), (C.int)(flags)))
}

// AvFormatSeekFile seeks to timestamp ts.
func AvFormatSeekFile(ic *AVFormatContext, streamIndex int32, minTs, ts, maxTs int64, flags int32) int32 {
	return (int32)(C.avformat_seek_file((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex), (C.int64_t)(minTs), (C.int64_t)(ts), (C.int64_t)(maxTs), (C.int)(flags)))
}

// AvFormatFlush siscards all internally buffered data. This can be useful when dealing with
// discontinuities in the byte stream. Generally works only with formats that
// can resync. This includes headerless formats like MPEG-TS/TS but should also
// work with NUT, Ogg and in a limited way AVI for example.
func AvFormatFlush(ic *AVFormatContext) int32 {
	return (int32)(C.avformat_flush((*C.struct_AVFormatContext)(ic)))
}

// AvReadPlay starts playing a network-based stream (e.g. RTSP stream) at the
// current position.
func AvReadPlay(ic *AVFormatContext) int32 {
	return (int32)(C.av_read_play((*C.struct_AVFormatContext)(ic)))
}

// AvReadPause pauses a network-based stream (e.g. RTSP stream).
func AvReadPause(ic *AVFormatContext) int32 {
	return (int32)(C.av_read_pause((*C.struct_AVFormatContext)(ic)))
}

// AvFormatCloseInput closes an opened input AVFormatContext. Free it and all its contents
// and set *s to NULL.
func AvFormatCloseInput(ic **AVFormatContext) {
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
func AvFormatWriteHeader(ic *AVFormatContext, options **AVDictionary) int32 {
	return (int32)(C.avformat_write_header((*C.struct_AVFormatContext)(ic),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvFormatInitOutput allocates the stream private data and initialize the codec,
// but do not write the header.
func AvFormatInitOutput(ic *AVFormatContext, options **AVDictionary) int32 {
	return (int32)(C.avformat_init_output((*C.struct_AVFormatContext)(ic),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvWriteFrame writes a packet to an output media file.
func AvWriteFrame(ic *AVFormatContext, pkt *AVPacket) int32 {
	return (int32)(C.av_write_frame((*C.struct_AVFormatContext)(ic), (*C.struct_AVPacket)(pkt)))
}

// AvInterleavedWriteFrame writes a packet to an output media file ensuring correct interleaving.
func AvInterleavedWriteFrame(ic *AVFormatContext, pkt *AVPacket) int32 {
	return (int32)(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(ic), (*C.struct_AVPacket)(pkt)))
}

// AvWriteUncodedFrame writes an uncoded frame to an output media file.
func AvWriteUncodedFrame(ic *AVFormatContext, streamIndex int32, frame *AVFrame) int32 {
	return (int32)(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex), (*C.struct_AVFrame)(frame)))
}

// AvInterleavedWriteUncodedFrame Write an uncoded frame to an output media file.
func AvInterleavedWriteUncodedFrame(ic *AVFormatContext, streamIndex int32, frame *AVFrame) int32 {
	return (int32)(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex), (*C.struct_AVFrame)(frame)))
}

// AvWriteUncodedFrameQuery tests whether a muxer supports uncoded frame.
func AvWriteUncodedFrameQuery(ic *AVFormatContext, streamIndex int32) int32 {
	return (int32)(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(ic),
		(C.int)(streamIndex)))
}

// AvWriteTrailer writes the stream trailer to an output media file and free the file private data.
func AvWriteTrailer(ic *AVFormatContext) int32 {
	return (int32)(C.av_write_trailer((*C.struct_AVFormatContext)(ic)))
}

// AvGuessFormat returns the output format in the list of registered output formats
// which best matches the provided parameters, or return NULL if there is no match.
func AvGuessFormat(shortName, filename, mimeType string) *AVOutputFormat {
	shortNamePtr, shortNameFunc := StringCasting(shortName)
	defer shortNameFunc()
	filenamePtr, filenameFunc := StringCasting(filename)
	defer filenameFunc()
	mimeTypePtr, mimeTypeFunc := StringCasting(mimeType)
	defer mimeTypeFunc()
	return (*AVOutputFormat)(C.av_guess_format((*C.char)(shortNamePtr),
		(*C.char)(filenamePtr),
		(*C.char)(mimeTypePtr)))
}

// AvGuessCodec guesses the codec ID based upon muxer and filename.
func AvGuessCodec(fmt *AVOutputFormat, shortName,
	filename, mimeType string, _type AVMediaType) AVCodecID {
	shortNamePtr, shortNameFunc := StringCasting(shortName)
	defer shortNameFunc()
	filenamePtr, filenameFunc := StringCasting(filename)
	defer filenameFunc()
	mimeTypePtr, mimeTypeFunc := StringCasting(mimeType)
	defer mimeTypeFunc()
	return (AVCodecID)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt),
		(*C.char)(shortNamePtr),
		(*C.char)(filenamePtr),
		(*C.char)(mimeTypePtr),
		(C.enum_AVMediaType)(_type)))
}

// AvGetOutputTimestamp gets timing information for the data currently output.
func AvGetOutputTimestamp(ic *AVFormatContext, stream int32, dts, wall *int64) int32 {
	return (int32)(C.av_get_output_timestamp((*C.struct_AVFormatContext)(ic),
		(C.int)(stream), (*C.int64_t)(dts), (*C.int64_t)(wall)))
}

// AvHexDump sends a nice hexadecimal dump of a buffer to the specified file stream.
func AvHexDump(f *FILE, buf *uint8, size int32) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(buf), (C.int)(size))
}

// AvHexDumpLog sends a nice hexadecimal dump of a buffer to the log.
func AvHexDumpLog(avcl CVoidPointer, level int32, buf *uint8, size int32) {
	C.av_hex_dump_log(VoidPointer(avcl), (C.int)(level), (*C.uint8_t)(buf), (C.int)(size))
}

// AvPktDump2 sends a nice dump of a packet to the specified file stream.
func AvPktDump2(f *FILE, pkt *AVPacket, dumpPayload int32, st *AVStream) {
	C.av_pkt_dump2((*C.FILE)(f), (*C.struct_AVPacket)(pkt),
		(C.int)(dumpPayload), (*C.struct_AVStream)(st))
}

// AvPktDumpLog2 sends a nice dump of a packet to the log.
func AvPktDumpLog2(avcl CVoidPointer, level int32, pkt *AVPacket, dumpPayload int32, st *AVStream) {
	C.av_pkt_dump_log2(VoidPointer(avcl), (C.int)(level), (*C.struct_AVPacket)(pkt),
		(C.int)(dumpPayload), (*C.struct_AVStream)(st))
}

// AvCodecGetId gets the AVCodecID for the given codec tag tag.
func AvCodecGetId(tags **AVCodecTag, tag uint32) AVCodecID {
	return (AVCodecID)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), (C.uint)(tag)))
}

// AvCodecGetTag gets the codec tag for the given codec id.
func AvCodecGetTag(tags **AVCodecTag, id AVCodecID) uint32 {
	return (uint32)(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(tags)),
		(C.enum_AVCodecID)(id)))
}

// AvCodecGetTag2 gets the codec tag for the given codec id.
func AvCodecGetTag2(tags **AVCodecTag, id AVCodecID, tag *uint32) int32 {
	return (int32)(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(tags)),
		(C.enum_AVCodecID)(id), (*C.uint)(tag)))
}

// AvFindDefaultStreamIndex
func AvFindDefaultStreamIndex(ic *AVFormatContext) int32 {
	return (int32)(C.av_find_default_stream_index((*C.struct_AVFormatContext)(ic)))
}

// AvIndexSearchTimestamp gets the index for a specific timestamp.
func AvIndexSearchTimestamp(st *AVStream, timestamp int64, flags int32) int32 {
	return (int32)(C.av_index_search_timestamp((*C.struct_AVStream)(st),
		(C.int64_t)(timestamp), (C.int)(flags)))
}

// AvAddIndexEntry add an index entry into a sorted list. Update the entry if the list
// already contains it.
func AvAddIndexEntry(st *AVStream, pos, timestamp int64, size, distance, flags int32) int32 {
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
func AvDumpFormat(ic *AVFormatContext, index int32, url string, isOutput int32) {
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
func AvSdpCreate(ac []*AVFormatContext, nFiles, sdpSize int32) (sdp string, ret int32) {
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
func AvFormatQueryCodec(ofmt *AVOutputFormat, codeID AVCodecID, stdCompliance int32) int32 {
	return (int32)(C.avformat_query_codec((*C.struct_AVOutputFormat)(ofmt),
		(C.enum_AVCodecID)(codeID), (C.int)(stdCompliance)))
}

// AvFormatGetRiffVideoTags returns the table mapping RIFF FourCCs for video to libavcodec AVCodecID.
func AvFormatGetRiffVideoTags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_riff_video_tags())
}

// AvFormatGetRiffAudioTags returns the table mapping RIFF FourCCs for audio to AVCodecID.
func AvFormatGetRiffAudioTags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_riff_audio_tags())
}

// AvFormatGetMovVideoTags returns the table mapping MOV FourCCs for video to libavcodec AVCodecID.
func AvFormatGetMovVideoTags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_mov_video_tags())
}

// AvFormatGetMovAudioTags returns the table mapping MOV FourCCs for audio to AVCodecID.
func AvFormatGetMovAudioTags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_mov_audio_tags())
}

// AvGuessSampleAspectRatio guesses the sample aspect ratio of a frame, based on both the stream and the
// frame aspect ratio.
func AvGuessSampleAspectRatio(format *AVFormatContext, stream *AVStream, frame *AVFrame) AVRational {
	return (AVRational)(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(format),
		(*C.struct_AVStream)(stream), (*C.struct_AVFrame)(frame)))
}

// AvGuessFrameRate guesses the frame rate, based on both the container and codec information.
func AvGuessFrameRate(ctx *AVFormatContext, stream *AVStream, frame *AVFrame) AVRational {
	return (AVRational)(C.av_guess_frame_rate((*C.struct_AVFormatContext)(ctx),
		(*C.struct_AVStream)(stream), (*C.struct_AVFrame)(frame)))
}

// AvFormatMatchStreamSpecifier checks if the stream st contained in s is
// matched by the stream specifier spec.
func AvFormatMatchStreamSpecifier(ic *AVFormatContext, st *AVStream, spec string) int32 {
	specPtr, specFunc := StringCasting(spec)
	defer specFunc()
	return (int32)(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(ic),
		(*C.struct_AVStream)(st), (*C.char)(specPtr)))
}

// AvFormatQueueAttachedPictures
func AvFormatQueueAttachedPictures(ic *AVFormatContext) int32 {
	return (int32)(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(ic)))
}

// Deprecated: No use.
//
// AvApplyBitstreamFilters applies a list of bitstream filters to a packet.
func AvApplyBitstreamFilters(codec *AVCodecContext, pkt *AVPacket, bsfc *AVBitStreamFilterContext) int32 {
	return (int32)(C.av_apply_bitstream_filters((*C.struct_AVCodecContext)(codec),
		(*C.struct_AVPacket)(pkt), (*C.struct_AVBitStreamFilterContext)(bsfc)))
}

// AVTimebaseSource
type AVTimebaseSource = C.enum_AVTimebaseSource

const (
	AVFMT_TBCF_AUTO        = AVTimebaseSource(C.AVFMT_TBCF_AUTO)
	AVFMT_TBCF_DECODER     = AVTimebaseSource(C.AVFMT_TBCF_DECODER)
	AVFMT_TBCF_DEMUXER     = AVTimebaseSource(C.AVFMT_TBCF_DEMUXER)
	AVFMT_TBCF_R_FRAMERATE = AVTimebaseSource(C.AVFMT_TBCF_R_FRAMERATE)
)

// AvFormatTransferInternalStreamTimingInfo transfers internal timing information from one stream to another.
func AvFormatTransferInternalStreamTimingInfo(ofmt *AVOutputFormat,
	ost, ist *AVStream,
	copyTb AVTimebaseSource) int32 {
	return (int32)(C.avformat_transfer_internal_stream_timing_info((*C.struct_AVOutputFormat)(ofmt),
		(*C.struct_AVStream)(ost), (*C.struct_AVStream)(ist),
		(C.enum_AVTimebaseSource)(copyTb)))
}

// AvStreamGetCodecTimebase gets the internal codec timebase from a stream.
func AvStreamGetCodecTimebase(st *AVStream) AVRational {
	return (AVRational)(C.av_stream_get_codec_timebase((*C.struct_AVStream)(st)))
}
