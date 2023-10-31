// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavformat/avio.h>

typedef int (*avio_interrupt_callback_func)(void* opaque);
typedef int (*avio_context_read_packet_func)(void *opaque, uint8_t *buf, int buf_size);
typedef int (*avio_context_write_packet_func)(void *opaque, uint8_t *buf, int buf_size);
typedef int64_t (*avio_context_seek_func)(void *opaque, int64_t offset, int whence);

int avio_printf_wrap(AVIOContext *s, const char *fmt) {
	return avio_printf(s, fmt, NULL);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	AVIO_SEEKABLE_NORMAL = C.AVIO_SEEKABLE_NORMAL
	AVIO_SEEKABLE_TIME   = C.AVIO_SEEKABLE_TIME
)

// AVIOInterruptCB
type AVIOInterruptCB C.struct_AVIOInterruptCB

// AVIOInterruptCallbackFunc
type AVIOInterruptCallbackFunc = C.avio_interrupt_callback_func

// GetCallback gets `AVIOInterruptCB.callback` value.
func (icb *AVIOInterruptCB) GetCallback() AVIOInterruptCallbackFunc {
	return (AVIOInterruptCallbackFunc)(icb.callback)
}

// SetCallback sets `AVIOInterruptCB.callback` value.
func (icb *AVIOInterruptCB) SetCallback(v AVIOInterruptCallbackFunc) {
	icb.callback = (C.avio_interrupt_callback_func)(v)
}

// GetCallbackAddr gets `AVIOInterruptCB.callback` address.
func (icb *AVIOInterruptCB) GetCallbackAddr() *AVIOInterruptCallbackFunc {
	return (*AVIOInterruptCallbackFunc)(&icb.callback)
}

// GetOpaque gets `AVIOInterruptCB.opaque` value.
func (icb *AVIOInterruptCB) GetOpaque() unsafe.Pointer {
	return icb.opaque
}

// SetOpaque sets `AVIOInterruptCB.opaque` value.
func (icb *AVIOInterruptCB) SetOpaque(v CVoidPointer) {
	icb.opaque = VoidPointer(v)
}

// GetOpaqueAddr gets `AVIOInterruptCB.opaque` address.
func (icb *AVIOInterruptCB) GetOpaqueAddr() *unsafe.Pointer {
	return &icb.opaque
}

// AVIODirEntryType
type AVIODirEntryType = int32 // C.enum_AVIODirEntryType

const (
	AVIO_ENTRY_UNKNOWN          = AVIODirEntryType(C.AVIO_ENTRY_UNKNOWN)
	AVIO_ENTRY_BLOCK_DEVICE     = AVIODirEntryType(C.AVIO_ENTRY_BLOCK_DEVICE)
	AVIO_ENTRY_CHARACTER_DEVICE = AVIODirEntryType(C.AVIO_ENTRY_CHARACTER_DEVICE)
	AVIO_ENTRY_DIRECTORY        = AVIODirEntryType(C.AVIO_ENTRY_DIRECTORY)
	AVIO_ENTRY_NAMED_PIPE       = AVIODirEntryType(C.AVIO_ENTRY_NAMED_PIPE)
	AVIO_ENTRY_SYMBOLIC_LINK    = AVIODirEntryType(C.AVIO_ENTRY_SYMBOLIC_LINK)
	AVIO_ENTRY_SOCKET           = AVIODirEntryType(C.AVIO_ENTRY_SOCKET)
	AVIO_ENTRY_FILE             = AVIODirEntryType(C.AVIO_ENTRY_FILE)
	AVIO_ENTRY_SERVER           = AVIODirEntryType(C.AVIO_ENTRY_SERVER)
	AVIO_ENTRY_SHARE            = AVIODirEntryType(C.AVIO_ENTRY_SHARE)
	AVIO_ENTRY_WORKGROUP        = AVIODirEntryType(C.AVIO_ENTRY_WORKGROUP)
)

// AVIODirEntry
type AVIODirEntry C.struct_AVIODirEntry

// GetName gets `AVIODirEntry.name` value.
func (de *AVIODirEntry) GetName() string {
	return C.GoString(de.name)
}

// GetType gets `AVIODirEntry.type` value.
func (de *AVIODirEntry) GetType() int32 {
	return (int32)(de._type)
}

// SetType sets `AVIODirEntry.type` value.
func (de *AVIODirEntry) SetType(v int32) {
	de._type = (C.int)(v)
}

// GetTypeAddr gets `AVIODirEntry.type` address.
func (de *AVIODirEntry) GetTypeAddr() *int32 {
	return (*int32)(&de._type)
}

// GetUtf8 gets `AVIODirEntry.utf8` value.
func (de *AVIODirEntry) GetUtf8() int32 {
	return (int32)(de.utf8)
}

// SetUtf8 sets `AVIODirEntry.utf8` value.
func (de *AVIODirEntry) SetUtf8(v int32) {
	de.utf8 = (C.int)(v)
}

// GetUtf8Addr gets `AVIODirEntry.utf8` address.
func (de *AVIODirEntry) GetUtf8Addr() *int32 {
	return (*int32)(&de.utf8)
}

// GetSize gets `AVIODirEntry.size` value.
func (de *AVIODirEntry) GetSize() int64 {
	return (int64)(de.size)
}

// SetSize sets `AVIODirEntry.size` value.
func (de *AVIODirEntry) SetSize(v int64) {
	de.size = (C.int64_t)(v)
}

// GetSizeAddr gets `AVIODirEntry.size` address.
func (de *AVIODirEntry) GetSizeAddr() *int64 {
	return (*int64)(&de.size)
}

// GetModificationTimestamp gets `AVIODirEntry.modification_timestamp` value.
func (de *AVIODirEntry) GetModificationTimestamp() int64 {
	return (int64)(de.modification_timestamp)
}

// SetModificationTimestamp sets `AVIODirEntry.modification_timestamp` value.
func (de *AVIODirEntry) SetModificationTimestamp(v int64) {
	de.modification_timestamp = (C.int64_t)(v)
}

// GetModificationTimestampAddr gets `AVIODirEntry.modification_timestamp` address.
func (de *AVIODirEntry) GetModificationTimestampAddr() *int64 {
	return (*int64)(&de.modification_timestamp)
}

// GetAccessTimestamp gets `AVIODirEntry.access_timestamp` value.
func (de *AVIODirEntry) GetAccessTimestamp() int64 {
	return (int64)(de.access_timestamp)
}

// SetAccessTimestamp sets `AVIODirEntry.access_timestamp` value.
func (de *AVIODirEntry) SetAccessTimestamp(v int64) {
	de.access_timestamp = (C.int64_t)(v)
}

// GetAccessTimestampAddr gets `AVIODirEntry.access_timestamp` address.
func (de *AVIODirEntry) GetAccessTimestampAddr() *int64 {
	return (*int64)(&de.access_timestamp)
}

// GetStatusChangeTimestamp gets `AVIODirEntry.status_change_timestamp` value.
func (de *AVIODirEntry) GetStatusChangeTimestamp() int64 {
	return (int64)(de.status_change_timestamp)
}

// SetStatusChangeTimestamp sets `AVIODirEntry.status_change_timestamp` value.
func (de *AVIODirEntry) SetStatusChangeTimestamp(v int64) {
	de.status_change_timestamp = (C.int64_t)(v)
}

// GetStatusChangeTimestampAddr gets `AVIODirEntry.status_change_timestamp` address.
func (de *AVIODirEntry) GetStatusChangeTimestampAddr() *int64 {
	return (*int64)(&de.status_change_timestamp)
}

// GetUserId gets `AVIODirEntry.user_id` value.
func (de *AVIODirEntry) GetUserId() int64 {
	return (int64)(de.user_id)
}

// SetUserId sets `AVIODirEntry.user_id` value.
func (de *AVIODirEntry) SetUserId(v int64) {
	de.user_id = (C.int64_t)(v)
}

// GetUserIdAddr gets `AVIODirEntry.user_id` address.
func (de *AVIODirEntry) GetUserIdAddr() *int64 {
	return (*int64)(&de.user_id)
}

// GetGroupId gets `AVIODirEntry.group_id` value.
func (de *AVIODirEntry) GetGroupId() int64 {
	return (int64)(de.group_id)
}

// SetGroupId sets `AVIODirEntry.group_id` value.
func (de *AVIODirEntry) SetGroupId(v int64) {
	de.group_id = (C.int64_t)(v)
}

// GetGroupIdAddr gets `AVIODirEntry.group_id` address.
func (de *AVIODirEntry) GetGroupIdAddr() *int64 {
	return (*int64)(&de.group_id)
}

// GetFilemode gets `AVIODirEntry.filemode` value.
func (de *AVIODirEntry) GetFilemode() int64 {
	return (int64)(de.filemode)
}

// SetFilemode sets `AVIODirEntry.filemode` value.
func (de *AVIODirEntry) SetFilemode(v int64) {
	de.filemode = (C.int64_t)(v)
}

// GetFilemodeAddr gets `AVIODirEntry.filemode` address.
func (de *AVIODirEntry) GetFilemodeAddr() *int64 {
	return (*int64)(&de.filemode)
}

// AVIODirContext
type AVIODirContext C.struct_AVIODirContext

// URLContext
type URLContext C.struct_URLContext

// Deprecated: No use.
//
// GetURLContext gets `AVIODirContext.url_context` value.
func (dc *AVIODirContext) GetURLContext() *URLContext {
	return (*URLContext)(dc.url_context)
}

// Deprecated: No use.
//
// SetURLContext sets `AVIODirContext.url_context` value.
func (dc *AVIODirContext) SetURLContext(v *URLContext) {
	dc.url_context = (*C.struct_URLContext)(v)
}

// Deprecated: No use.
//
// GetURLContextAddr gets `AVIODirContext.url_context` address.
func (dc *AVIODirContext) GetURLContextAddr() **URLContext {
	return (**URLContext)(unsafe.Pointer(&dc.url_context))
}

// AVIODataMarkerType
type AVIODataMarkerType = C.enum_AVIODataMarkerType

const (
	AVIO_DATA_MARKER_HEADER         = AVIODataMarkerType(C.AVIO_DATA_MARKER_HEADER)
	AVIO_DATA_MARKER_SYNC_POINT     = AVIODataMarkerType(C.AVIO_DATA_MARKER_SYNC_POINT)
	AVIO_DATA_MARKER_BOUNDARY_POINT = AVIODataMarkerType(C.AVIO_DATA_MARKER_BOUNDARY_POINT)
	AVIO_DATA_MARKER_UNKNOWN        = AVIODataMarkerType(C.AVIO_DATA_MARKER_UNKNOWN)
	AVIO_DATA_MARKER_TRAILER        = AVIODataMarkerType(C.AVIO_DATA_MARKER_TRAILER)
	AVIO_DATA_MARKER_FLUSH_POINT    = AVIODataMarkerType(C.AVIO_DATA_MARKER_FLUSH_POINT)
)

// AVIOContext
type AVIOContext C.struct_AVIOContext

// GetAvClass gets `AVIOContext.av_class` value.
func (ctx *AVIOContext) GetAvClass() *AVClass {
	return (*AVClass)(ctx.av_class)
}

// SetAvClass sets `AVIOContext.av_class` value.
func (ctx *AVIOContext) SetAvClass(v *AVClass) {
	ctx.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIOContext.av_class` address.
func (ctx *AVIOContext) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&ctx.av_class))
}

// GetBuffer gets `AVIOContext.buffer` value.
func (ctx *AVIOContext) GetBuffer() *uint8 {
	return (*uint8)(ctx.buffer)
}

// SetBuffer sets `AVIOContext.buffer` value.
func (ctx *AVIOContext) SetBuffer(v *uint8) {
	ctx.buffer = (*C.uint8_t)(v)
}

// GetBufferAddr gets `AVIOContext.buffer` address.
func (ctx *AVIOContext) GetBufferAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&ctx.buffer))
}

// GetBufferSize gets `AVIOContext.buffer_size` value.
func (ctx *AVIOContext) GetBufferSize() int32 {
	return (int32)(ctx.buffer_size)
}

// SetBufferSize sets `AVIOContext.buffer_size` value.
func (ctx *AVIOContext) SetBufferSize(v int32) {
	ctx.buffer_size = (C.int)(v)
}

// GetBufferSizeAddr gets `AVIOContext.buffer_size` address.
func (ctx *AVIOContext) GetBufferSizeAddr() *int32 {
	return (*int32)(&ctx.buffer_size)
}

// GetBufPtr gets `AVIOContext.buf_ptr` value.
func (ctx *AVIOContext) GetBufPtr() *uint8 {
	return (*uint8)(ctx.buf_ptr)
}

// SetBufPtr sets `AVIOContext.buf_ptr` value.
func (ctx *AVIOContext) SetBufPtr(v *uint8) {
	ctx.buf_ptr = (*C.uint8_t)(v)
}

// GetBufPtrAddr gets `AVIOContext.buf_ptr` address.
func (ctx *AVIOContext) GetBufPtrAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&ctx.buf_ptr))
}

// GetBufEnd gets `AVIOContext.buf_end` value.
func (ctx *AVIOContext) GetBufEnd() *uint8 {
	return (*uint8)(ctx.buf_end)
}

// SetBufEnd sets `AVIOContext.buf_end` value.
func (ctx *AVIOContext) SetBufEnd(v *uint8) {
	ctx.buf_end = (*C.uint8_t)(v)
}

// GetBufEndAddr gets `AVIOContext.buf_end` address.
func (ctx *AVIOContext) GetBufEndAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&ctx.buf_end))
}

// GetOpaque gets `AVIOContext.opaque` value.
func (ctx *AVIOContext) GetOpaque() unsafe.Pointer {
	return ctx.opaque
}

// SetOpaque sets `AVIOContext.opaque` value.
func (ctx *AVIOContext) SetOpaque(v CVoidPointer) {
	ctx.opaque = VoidPointer(v)
}

// GetOpaqueAddr gets `AVIOContext.opaque` address.
func (ctx *AVIOContext) GetOpaqueAddr() *unsafe.Pointer {
	return (*unsafe.Pointer)(&ctx.opaque)
}

// GetPos gets `AVIOContext.pos` value.
func (ctx *AVIOContext) GetPos() int64 {
	return (int64)(ctx.pos)
}

// SetPos sets `AVIOContext.pos` value.
func (ctx *AVIOContext) SetPos(v int64) {
	ctx.pos = (C.int64_t)(v)
}

// GetPosAddr gets `AVIOContext.pos` address.
func (ctx *AVIOContext) GetPosAddr() *int64 {
	return (*int64)(&ctx.pos)
}

// GetEofReached gets `AVIOContext.eof_reached` value.
func (ctx *AVIOContext) GetEofReached() int32 {
	return (int32)(ctx.eof_reached)
}

// SetEofReached sets `AVIOContext.eof_reached` value.
func (ctx *AVIOContext) SetEofReached(v int32) {
	ctx.eof_reached = (C.int)(v)
}

// GetEofReachedAddr gets `AVIOContext.eof_reached` address.
func (ctx *AVIOContext) GetEofReachedAddr() *int32 {
	return (*int32)(&ctx.eof_reached)
}

// GetWriteFlag gets `AVIOContext.write_flag` value.
func (ctx *AVIOContext) GetWriteFlag() int32 {
	return (int32)(ctx.write_flag)
}

// SetWriteFlag sets `AVIOContext.write_flag` value.
func (ctx *AVIOContext) SetWriteFlag(v int32) {
	ctx.write_flag = (C.int)(v)
}

// GetWriteFlagAddr gets `AVIOContext.write_flag` address.
func (ctx *AVIOContext) GetWriteFlagAddr() *int32 {
	return (*int32)(&ctx.write_flag)
}

// GetMaxPacketSize gets `AVIOContext.max_packet_size` value.
func (ctx *AVIOContext) GetMaxPacketSize() int32 {
	return (int32)(ctx.max_packet_size)
}

// SetMaxPacketSize sets `AVIOContext.max_packet_size` value.
func (ctx *AVIOContext) SetMaxPacketSize(v int32) {
	ctx.max_packet_size = (C.int)(v)
}

// GetMaxPacketSizeAddr gets `AVIOContext.max_packet_size` address.
func (ctx *AVIOContext) GetMaxPacketSizeAddr() *int32 {
	return (*int32)(&ctx.max_packet_size)
}

// GetChecksum gets `AVIOContext.checksum` value.
func (ctx *AVIOContext) GetChecksum() uint32 {
	return (uint32)(ctx.checksum)
}

// SetChecksum sets `AVIOContext.checksum` value.
func (ctx *AVIOContext) SetChecksum(v uint32) {
	ctx.checksum = (C.ulong)(v)
}

// GetChecksumAddr gets `AVIOContext.checksum` address.
func (ctx *AVIOContext) GetChecksumAddr() *uint32 {
	return (*uint32)(unsafe.Pointer(&ctx.checksum))
}

// GetChecksumPtr gets `AVIOContext.checksum_ptr` value.
func (ctx *AVIOContext) GetChecksumPtr() *uint8 {
	return (*uint8)(ctx.checksum_ptr)
}

// SetChecksumPtr sets `AVIOContext.checksum_ptr` value.
func (ctx *AVIOContext) SetChecksumPtr(v *uint8) {
	ctx.checksum_ptr = (*C.uint8_t)(v)
}

// GetChecksumPtrAddr gets `AVIOContext.checksum_ptr` address.
func (ctx *AVIOContext) GetChecksumPtrAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&ctx.checksum_ptr))
}

// GetError gets `AVIOContext.error` value.
func (ctx *AVIOContext) GetError() int32 {
	return (int32)(ctx.error)
}

// SetError sets `AVIOContext.error` value.
func (ctx *AVIOContext) SetError(v int32) {
	ctx.error = (C.int)(v)
}

// GetErrorAddr gets `AVIOContext.error` address.
func (ctx *AVIOContext) GetErrorAddr() *int32 {
	return (*int32)(&ctx.error)
}

// GetSeekable gets `AVIOContext.seekable` value.
func (ctx *AVIOContext) GetSeekable() int32 {
	return (int32)(ctx.seekable)
}

// SetSeekable sets `AVIOContext.seekable` value.
func (ctx *AVIOContext) SetSeekable(v int32) {
	ctx.seekable = (C.int)(v)
}

// GetSeekableAddr gets `AVIOContext.seekable` address.
func (ctx *AVIOContext) GetSeekableAddr() *int32 {
	return (*int32)(&ctx.seekable)
}

// GetDirect gets `AVIOContext.direct` value.
func (ctx *AVIOContext) GetDirect() int32 {
	return (int32)(ctx.direct)
}

// SetDirect sets `AVIOContext.direct` value.
func (ctx *AVIOContext) SetDirect(v int32) {
	ctx.direct = (C.int)(v)
}

// GetDirectAddr gets `AVIOContext.direct` address.
func (ctx *AVIOContext) GetDirectAddr() *int32 {
	return (*int32)(&ctx.direct)
}

// GetBytesRead gets `AVIOContext.bytes_read` value.
func (ctx *AVIOContext) GetBytesRead() int64 {
	return (int64)(ctx.bytes_read)
}

// SetBytesRead sets `AVIOContext.bytes_read` value.
func (ctx *AVIOContext) SetBytesRead(v int64) {
	ctx.bytes_read = (C.int64_t)(v)
}

// GetBytesReadAddr gets `AVIOContext.bytes_read` address.
func (ctx *AVIOContext) GetBytesReadAddr() *int64 {
	return (*int64)(&ctx.bytes_read)
}

// GetBytesWritten gets `AVIOContext.bytes_written` value.
func (ctx *AVIOContext) GetBytesWritten() int64 {
	return (int64)(ctx.bytes_written)
}

// SetBytesWritten sets `AVIOContext.bytes_written` value.
func (ctx *AVIOContext) SetBytesWritten(v int64) {
	ctx.bytes_written = (C.int64_t)(v)
}

// GetBytesWrittenAddr gets `AVIOContext.bytes_written` address.
func (ctx *AVIOContext) GetBytesWrittenAddr() *int64 {
	return (*int64)(&ctx.bytes_written)
}

// GetProtocolWhitelist gets `AVIOContext.protocol_whitelist` value.
func (ctx *AVIOContext) GetProtocolWhitelist() string {
	return C.GoString(ctx.protocol_whitelist)
}

// GetProtocolBlacklist gets `AVIOContext.protocol_blacklist` value.
func (ctx *AVIOContext) GetProtocolBlacklist() string {
	return C.GoString(ctx.protocol_blacklist)
}

// GetIgnoreBoundaryPoint gets `AVIOContext.ignore_boundary_point` value.
func (ctx *AVIOContext) GetIgnoreBoundaryPoint() int32 {
	return (int32)(ctx.ignore_boundary_point)
}

// SetIgnoreBoundaryPoint sets `AVIOContext.ignore_boundary_point` value.
func (ctx *AVIOContext) SetIgnoreBoundaryPoint(v int32) {
	ctx.ignore_boundary_point = (C.int)(v)
}

// GetIgnoreBoundaryPointAddr gets `AVIOContext.ignore_boundary_point` address.
func (ctx *AVIOContext) GetIgnoreBoundaryPointAddr() *int32 {
	return (*int32)(&ctx.ignore_boundary_point)
}

// GetBufPtrMax gets `AVIOContext.buf_ptr_max` value.
func (ctx *AVIOContext) GetBufPtrMax() *uint8 {
	return (*uint8)(ctx.buf_ptr_max)
}

// SetBufPtrMax sets `AVIOContext.buf_ptr_max` value.
func (ctx *AVIOContext) SetBufPtrMax(v *uint8) {
	ctx.buf_ptr_max = (*C.uint8_t)(v)
}

// GetBufPtrMaxAddr gets `AVIOContext.buf_ptr_max` address.
func (ctx *AVIOContext) GetBufPtrMaxAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&ctx.buf_ptr_max))
}

// GetMinPacketSize gets `AVIOContext.min_packet_size` value.
func (ctx *AVIOContext) GetMinPacketSize() int32 {
	return (int32)(ctx.min_packet_size)
}

// SetMinPacketSize sets `AVIOContext.min_packet_size` value.
func (ctx *AVIOContext) SetMinPacketSize(v int32) {
	ctx.min_packet_size = (C.int)(v)
}

// GetMinPacketSizeAddr gets `AVIOContext.min_packet_size` address.
func (ctx *AVIOContext) GetMinPacketSizeAddr() *int32 {
	return (*int32)(&ctx.min_packet_size)
}

// AvIOFindProtocolName returns the name of the protocol that will handle the passed URL.
func AvIOFindProtocolName(url string) string {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return C.GoString(C.avio_find_protocol_name((*C.char)(urlPtr)))
}

// AvIOCheck returns AVIO_FLAG_* access flags corresponding to the access permissions
// of the resource in url, or a negative value corresponding to an
// AvERROR code in case of failure. The returned access flags are
// masked by the value in flags.
func AvIOCheck(url string, flags int32) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avio_check((*C.char)(urlPtr), (C.int)(flags)))
}

// AvIOOpenDir opens directory for reading.
func AvIOOpenDir(s **AVIODirContext, url string, options **AVDictionary) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avio_open_dir((**C.struct_AVIODirContext)(unsafe.Pointer(s)),
		(*C.char)(urlPtr), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvIOReadDir gets next directory entry.
func AvIOReadDir(s *AVIODirContext, next **AVIODirEntry) int32 {
	return (int32)(C.avio_read_dir((*C.struct_AVIODirContext)(s),
		(**C.struct_AVIODirEntry)(unsafe.Pointer(next))))
}

// AvIOCloseDir closes directory.
func AvIOCloseDir(s **AVIODirContext) int32 {
	return (int32)(C.avio_close_dir((**C.struct_AVIODirContext)(unsafe.Pointer(s))))
}

// AvIOFreeDirectoryEntry frees entry allocated by AvIOReadDir().
func AvIOFreeDirectoryEntry(entry **AVIODirEntry) {
	C.avio_free_directory_entry((**C.struct_AVIODirEntry)(unsafe.Pointer(entry)))
}

// typedef int (*avio_context_read_packet_func)(void *opaque, uint8_t *buf, int buf_size);
type AVIOContextReadPacketFunc = C.avio_context_read_packet_func

// typedef int (*avio_context_write_packet_func)(void *opaque, uint8_t *buf, int buf_size);
type AVIOContextWritePacketFunc = C.avio_context_write_packet_func

// typedef int64_t (*avio_context_seek_func)(void *opaque, int64_t offset, int whence);
type AVIOContextSeekFunc = C.avio_context_seek_func

// AvIOAllocContext sllocates and initialize an AVIOContext for buffered I/O. It must be later
// freed with AVIOContextFree().
func AvIOAllocContext(buffer *uint8, bufferSize, writeFlag int32,
	opaque CVoidPointer,
	readPacket AVIOContextReadPacketFunc,
	writePacket AVIOContextWritePacketFunc,
	seek AVIOContextSeekFunc) *AVIOContext {
	return (*AVIOContext)(C.avio_alloc_context((*C.uint8_t)(buffer), (C.int)(bufferSize),
		(C.int)(writeFlag), VoidPointer(opaque),
		(C.avio_context_read_packet_func)(readPacket),
		(C.avio_context_write_packet_func)(writePacket),
		(C.avio_context_seek_func)(seek)))
}

// AvIOContextFree frees the supplied IO context and everything associated with it.
func AvIOContextFree(s **AVIOContext) {
	C.avio_context_free((**C.struct_AVIOContext)(unsafe.Pointer(s)))
}

// AvIOW8
func AvIOW8(s *AVIOContext, b int32) {
	C.avio_w8((*C.struct_AVIOContext)(s), (C.int)(b))
}

// AvIOWrite
func AvIOWrite(s *AVIOContext, buf *uint8, size int32) {
	C.avio_write((*C.struct_AVIOContext)(s), (*C.uchar)(buf), (C.int)(size))
}

// AvIOWl64
func AvIOWl64(s *AVIOContext, val uint64) {
	C.avio_wl64((*C.struct_AVIOContext)(s), (C.uint64_t)(val))
}

// AvIOWb64
func AvIOWb64(s *AVIOContext, val uint64) {
	C.avio_wb64((*C.struct_AVIOContext)(s), (C.uint64_t)(val))
}

// AvIOWl32
func AvIOWl32(s *AVIOContext, val uint32) {
	C.avio_wl32((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWb32
func AvIOWb32(s *AVIOContext, val uint32) {
	C.avio_wb32((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWl24
func AvIOWl24(s *AVIOContext, val uint32) {
	C.avio_wl24((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWb24
func AvIOWb24(s *AVIOContext, val uint32) {
	C.avio_wb24((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWl16
func AvIOWl16(s *AVIOContext, val uint32) {
	C.avio_wl16((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWb16
func AvIOWb16(s *AVIOContext, val uint32) {
	C.avio_wb16((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOPutStr Write a string.
func AvIOPutStr(s *AVIOContext, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.avio_put_str((*C.struct_AVIOContext)(s), (*C.char)(strPtr)))
}

// AvIOPutStr16le converts an UTF-8 string to UTF-16LE and write it.
func AvIOPutStr16le(s *AVIOContext, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.avio_put_str16le((*C.struct_AVIOContext)(s), (*C.char)(strPtr)))
}

// AvIOPutStr16be converts an UTF-8 string to UTF-16BE and write it.
func AvIOPutStr16be(s *AVIOContext, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.avio_put_str16be((*C.struct_AVIOContext)(s), (*C.char)(strPtr)))
}

// AvIOWriteMarker
func AvIOWriteMarker(s *AVIOContext, time int64, _type AVIODataMarkerType) {
	C.avio_write_marker((*C.struct_AVIOContext)(s), (C.int64_t)(time), (C.enum_AVIODataMarkerType)(_type))
}

const (
	AVSEEK_SIZE  = C.AVSEEK_SIZE
	AVSEEK_FORCE = C.AVSEEK_FORCE
)

// AvIOSeek equivalents fseek().
func AvIOSeek(s *AVIOContext, offset int64, whence int32) int64 {
	return (int64)(C.avio_seek((*C.struct_AVIOContext)(s), (C.int64_t)(offset), (C.int)(whence)))
}

// AvIOSkip skips given number of bytes forward.
func AvIOSkip(s *AVIOContext, offset int64) int64 {
	return (int64)(C.avio_skip((*C.struct_AVIOContext)(s), (C.int64_t)(offset)))
}

// AvIOTell equivalents ftell().
func AvIOTell(s *AVIOContext) int64 {
	return (int64)(C.avio_tell((*C.struct_AVIOContext)(s)))
}

// AvIOSize gets the filesize.
func AvIOSize(s *AVIOContext) int64 {
	return (int64)(C.avio_size((*C.struct_AVIOContext)(s)))
}

// AvIOFeof similar to feof() but also returns nonzero on read errors.
func AvIOFeof(s *AVIOContext) int32 {
	return (int32)(C.avio_feof((*C.struct_AVIOContext)(s)))
}

// NONEED: avio_vprintf

// AvIOPrintf Writes a formatted string to the context.
func AvIOPrintf(s *AVIOContext, _fmt string, va ...any) int32 {
	fmtPtr, fmtFunc := StringCasting(fmt.Sprintf(_fmt, va...))
	defer fmtFunc()
	return (int32)(C.avio_printf_wrap((*C.struct_AVIOContext)(s), (*C.char)(fmtPtr)))
}

// NONEED: avio_print_string_array

// AvIOPrint
func AvIOPrint(s *AVIOContext, va ...any) {
	fmtPtr, fmtFunc := StringCasting(fmt.Sprint(va...))
	defer fmtFunc()
	fmtArray := []*C.char{(*C.char)(fmtPtr), nil}
	C.avio_print_string_array((*C.struct_AVIOContext)(s), &fmtArray[0])
}

// AvIOFlush forces flushing of buffered data.
func AvIOFlush(s *AVIOContext) {
	C.avio_flush((*C.struct_AVIOContext)(s))
}

// AvIORead reads size bytes from AVIOContext into buf.
func AvIORead(s *AVIOContext, buf *uint8, size int32) int32 {
	return (int32)(C.avio_read((*C.struct_AVIOContext)(s), (*C.uchar)(buf), (C.int)(size)))
}

// AvIOReadPartial sead size bytes from AVIOContext into buf. Unlike avio_read(), this is allowed
// to read fewer bytes than requested. The missing bytes can be read in the next
// call. This always tries to read at least 1 byte.
// Useful to reduce latency in certain cases.
func AvIOReadPartial(s *AVIOContext, buf *uint8, size int32) int32 {
	return (int32)(C.avio_read_partial((*C.struct_AVIOContext)(s), (*C.uchar)(buf), (C.int)(size)))
}

// AvIOR8
func AvIOR8(s *AVIOContext) int32 {
	return (int32)(C.avio_r8((*C.struct_AVIOContext)(s)))
}

// AvIORl16
func AvIORl16(s *AVIOContext) uint32 {
	return (uint32)(C.avio_rl16((*C.struct_AVIOContext)(s)))
}

// AvIORl24
func AvIORl24(s *AVIOContext) uint32 {
	return (uint32)(C.avio_rl24((*C.struct_AVIOContext)(s)))
}

// AvIORl32
func AvIORl32(s *AVIOContext) uint32 {
	return (uint32)(C.avio_rl32((*C.struct_AVIOContext)(s)))
}

// AvIORl64
func AvIORl64(s *AVIOContext) uint64 {
	return (uint64)(C.avio_rl64((*C.struct_AVIOContext)(s)))
}

// AvIORb16
func AvIORb16(s *AVIOContext) uint32 {
	return (uint32)(C.avio_rb16((*C.struct_AVIOContext)(s)))
}

// AvIORb24
func AvIORb24(s *AVIOContext) uint32 {
	return (uint32)(C.avio_rb24((*C.struct_AVIOContext)(s)))
}

// AvIORb32
func AvIORb32(s *AVIOContext) uint32 {
	return (uint32)(C.avio_rb32((*C.struct_AVIOContext)(s)))
}

// AvIORb64
func AvIORb64(s *AVIOContext) uint64 {
	return (uint64)(C.avio_rb64((*C.struct_AVIOContext)(s)))
}

// AvIOGetStr reads a string from pb into buf. The reading will terminate when either
// a NULL character was encountered, maxlen bytes have been read, or nothing
// more can be read from pb. The result is guaranteed to be NULL-terminated, it
// will be truncated if buf is too small.
// Note that the string is not interpreted or validated in any way, it
// might get truncated in the middle of a sequence for multi-byte encodings.
func AvIOGetStr(s *AVIOContext, maxLen int32, buf *int8, buflen int32) int32 {
	return (int32)(C.avio_get_str((*C.struct_AVIOContext)(s), (C.int)(maxLen), (*C.char)(buf), (C.int)(buflen)))
}

// AvIOGetStr16le reads a UTF-16LE string from pb and convert it to UTF-8.
func AvIOGetStr16le(s *AVIOContext, maxLen int32, buf *int8, buflen int32) int32 {
	return (int32)(C.avio_get_str16le((*C.struct_AVIOContext)(s), (C.int)(maxLen), (*C.char)(buf), (C.int)(buflen)))
}

// AvIOGetStr16be reads a UTF-16BE string from pb and convert it to UTF-8.
func AvIOGetStr16be(s *AVIOContext, maxLen int32, buf *int8, buflen int32) int32 {
	return (int32)(C.avio_get_str16be((*C.struct_AVIOContext)(s), (C.int)(maxLen), (*C.char)(buf), (C.int)(buflen)))
}

const (
	AVIO_FLAG_READ       = C.AVIO_FLAG_READ
	AVIO_FLAG_WRITE      = C.AVIO_FLAG_WRITE
	AVIO_FLAG_READ_WRITE = C.AVIO_FLAG_READ_WRITE
	AVIO_FLAG_NONBLOCK   = C.AVIO_FLAG_NONBLOCK
	AVIO_FLAG_DIRECT     = C.AVIO_FLAG_DIRECT
)

// AvIOOpen creates and initializes a AVIOContext for accessing the resource indicated by url.
func AvIOOpen(s **AVIOContext, url string, flags int32) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avio_open((**C.struct_AVIOContext)(unsafe.Pointer(s)),
		(*C.char)(urlPtr), (C.int)(flags)))
}

// AvIOOpen2 creates and initializes a AVIOContext for accessing the resource indicated by url.
func AvIOOpen2(s **AVIOContext, url string, flags int32,
	intCb *AVIOInterruptCB, options **AVDictionary) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avio_open2((**C.struct_AVIOContext)(unsafe.Pointer(s)),
		(*C.char)(urlPtr), (C.int)(flags),
		(*C.struct_AVIOInterruptCB)(intCb), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvIOClose
func AvIOClose(s *AVIOContext) int32 {
	return (int32)(C.avio_close((*C.struct_AVIOContext)(s)))
}

// AvIOClosep closes the resource accessed by the AVIOContext *s, free it
// and set the pointer pointing to it to NULL.
func AvIOClosep(s **AVIOContext) int32 {
	return (int32)(C.avio_closep((**C.struct_AVIOContext)(unsafe.Pointer(s))))
}

// AvIOOpenDynBuf opens a write only memory stream.
func AvIOOpenDynBuf(s **AVIOContext) int32 {
	return (int32)(C.avio_open_dyn_buf((**C.struct_AVIOContext)(unsafe.Pointer(s))))
}

// AvIOGetDynBuf returns the written size and a pointer to the buffer.
func AvIOGetDynBuf(s *AVIOContext, pbuffer **uint8) int32 {
	return (int32)(C.avio_get_dyn_buf((*C.struct_AVIOContext)(s),
		(**C.uint8_t)(unsafe.Pointer(pbuffer))))
}

// AvIOCloseDynBuf returns the written size and a pointer to the buffer. The buffer
// must be freed with AvFree().
func AvIOCloseDynBuf(s *AVIOContext, pbuffer **uint8) int32 {
	return (int32)(C.avio_close_dyn_buf((*C.struct_AVIOContext)(s),
		(**C.uint8_t)(unsafe.Pointer(pbuffer))))
}

// AvIOEnumProtocols iterates through names of available protocols.
func AvIOEnumProtocols(opaque CVoidPointerPointer, output int32) string {
	return C.GoString(C.avio_enum_protocols(VoidPointerPointer(opaque), (C.int)(output)))
}

// AvIOProtocolGetClass gets AVClass by names of available protocols.
func AvIOProtocolGetClass(name string) *AVClass {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AVClass)(C.avio_protocol_get_class((*C.char)(namePtr)))
}

// AvIOPause pauses and resumes playing - only meaningful if using a network streaming
// protocol (e.g. MMS).
func AvIOPause(s *AVIOContext, pause int32) int32 {
	return (int32)(C.avio_pause((*C.struct_AVIOContext)(s), (C.int)(pause)))
}

// AvIOSeekTime seeks to a given timestamp relative to some component stream.
// Only meaningful if using a network streaming protocol (e.g. MMS.).
func AvIOSeekTime(s *AVIOContext, streamIndex int32, timestamp int64, flags int32) int64 {
	return (int64)(C.avio_seek_time((*C.struct_AVIOContext)(s),
		(C.int)(streamIndex), (C.int64_t)(timestamp), (C.int)(flags)))
}

// AvIOReadToBprint reads contents of h into print buffer, up to max_size bytes, or up to EOF.
func AvIOReadToBprint(s *AVIOContext, pb *AVBPrint, maxSize uintptr) int32 {
	return (int32)(C.avio_read_to_bprint((*C.struct_AVIOContext)(s),
		(*C.struct_AVBPrint)(pb), (C.size_t)(maxSize)))
}

// AvIOAccept accepts and allocates a client context on a server context.
func AvIOAccept(s *AVIOContext, c **AVIOContext) int32 {
	return (int32)(C.avio_accept((*C.struct_AVIOContext)(s), (**C.struct_AVIOContext)(unsafe.Pointer(c))))
}

// AvIOHandshake performs one step of the protocol handshake to accept a new client.
func AvIOHandshake(s *AVIOContext) int32 {
	return (int32)(C.avio_handshake((*C.struct_AVIOContext)(s)))
}
