package ffmpeg

/*
#include <libavformat/avio.h>

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

// AvIOInterruptCB
type AvIOInterruptCB C.struct_AVIOInterruptCB

// AvIODirEntry
type AvIODirEntry C.struct_AVIODirEntry

// AvIODirEntryType
type AvIODirEntryType = C.enum_AVIODirEntryType

const (
	AVIO_ENTRY_UNKNOWN          = AvIODirEntryType(C.AVIO_ENTRY_UNKNOWN)
	AVIO_ENTRY_BLOCK_DEVICE     = AvIODirEntryType(C.AVIO_ENTRY_BLOCK_DEVICE)
	AVIO_ENTRY_CHARACTER_DEVICE = AvIODirEntryType(C.AVIO_ENTRY_CHARACTER_DEVICE)
	AVIO_ENTRY_DIRECTORY        = AvIODirEntryType(C.AVIO_ENTRY_DIRECTORY)
	AVIO_ENTRY_NAMED_PIPE       = AvIODirEntryType(C.AVIO_ENTRY_NAMED_PIPE)
	AVIO_ENTRY_SYMBOLIC_LINK    = AvIODirEntryType(C.AVIO_ENTRY_SYMBOLIC_LINK)
	AVIO_ENTRY_SOCKET           = AvIODirEntryType(C.AVIO_ENTRY_SOCKET)
	AVIO_ENTRY_FILE             = AvIODirEntryType(C.AVIO_ENTRY_FILE)
	AVIO_ENTRY_SERVER           = AvIODirEntryType(C.AVIO_ENTRY_SERVER)
	AVIO_ENTRY_SHARE            = AvIODirEntryType(C.AVIO_ENTRY_SHARE)
	AVIO_ENTRY_WORKGROUP        = AvIODirEntryType(C.AVIO_ENTRY_WORKGROUP)
)

// AvIODirContext
type AvIODirContext C.struct_AVIODirContext

// AvIODataMarkerType
type AvIODataMarkerType = C.enum_AVIODataMarkerType

const (
	AVIO_DATA_MARKER_HEADER         = AvIODataMarkerType(C.AVIO_DATA_MARKER_HEADER)
	AVIO_DATA_MARKER_SYNC_POINT     = AvIODataMarkerType(C.AVIO_DATA_MARKER_SYNC_POINT)
	AVIO_DATA_MARKER_BOUNDARY_POINT = AvIODataMarkerType(C.AVIO_DATA_MARKER_BOUNDARY_POINT)
	AVIO_DATA_MARKER_UNKNOWN        = AvIODataMarkerType(C.AVIO_DATA_MARKER_UNKNOWN)
	AVIO_DATA_MARKER_TRAILER        = AvIODataMarkerType(C.AVIO_DATA_MARKER_TRAILER)
	AVIO_DATA_MARKER_FLUSH_POINT    = AvIODataMarkerType(C.AVIO_DATA_MARKER_FLUSH_POINT)
)

// AvIOContext
type AvIOContext C.struct_AVIOContext

// AvIOFindProtocolName returns the name of the protocol that will handle the passed URL.
func AvIOFindProtocolName(url string) string {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return C.GoString(C.avio_find_protocol_name((*C.char)(urlPtr)))
}

// AvIOCheck returns AVIO_FLAG_* access flags corresponding to the access permissions
// of the resource in url, or a negative value corresponding to an
// AVERROR code in case of failure. The returned access flags are
// masked by the value in flags.
// avio_check
func AvIOCheck(url string, flags int32) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avio_check((*C.char)(urlPtr), (C.int)(flags)))
}

// AvPrivIoMove moves or renames a resource.
func AvPrivIoMove(urlSrc, urlDst string) int32 {
	urlSrcPtr, urlSrcFunc := StringCasting(urlSrc)
	defer urlSrcFunc()
	urlDstPtr, urlDstFunc := StringCasting(urlDst)
	defer urlDstFunc()
	return (int32)(C.avpriv_io_move((*C.char)(urlSrcPtr), (*C.char)(urlDstPtr)))
}

// AvPrivIoDelete deletes a resource.
func AvPrivIoDelete(url string) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avpriv_io_delete((*C.char)(urlPtr)))
}

// AvIOOpenDir opens directory for reading.
func AvIOOpenDir(s **AvIODirContext, url string, options **AvDictionary) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avio_open_dir((**C.struct_AVIODirContext)(unsafe.Pointer(s)),
		(*C.char)(urlPtr), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvIOReadDir gets next directory entry.
func AvIOReadDir(s *AvIODirContext, next **AvIODirEntry) int32 {
	return (int32)(C.avio_read_dir((*C.struct_AVIODirContext)(s),
		(**C.struct_AVIODirEntry)(unsafe.Pointer(next))))
}

// AvIOCloseDir closes directory.
func AvIOCloseDir(s **AvIODirContext) int32 {
	return (int32)(C.avio_close_dir((**C.struct_AVIODirContext)(unsafe.Pointer(s))))
}

// AvIOFreeDirectoryEntry frees entry allocated by AvIOReadDir().
func AvIOFreeDirectoryEntry(entry **AvIODirEntry) {
	C.avio_free_directory_entry((**C.struct_AVIODirEntry)(unsafe.Pointer(entry)))
}

// typedef int (*avio_context_read_packet_func)(void *opaque, uint8_t *buf, int buf_size)
type AvIOContextReadPacketFunc C.avio_context_read_packet_func

// typedef int (*avio_context_write_packet_func)(void *opaque, uint8_t *buf, int buf_size)
type AvIOContextWritePacketFunc C.avio_context_write_packet_func

// typedef int64_t (*avio_context_seek_func)(void *opaque, int64_t offset, int whence)
type AvIOContextSeekFunc C.avio_context_seek_func

// avio_alloc_context
func avio_alloc_context(buffer *uint8, bufferSize, writeFlag int32,
	opaque unsafe.Pointer,
	readPacket AvIOContextReadPacketFunc,
	writePacket AvIOContextWritePacketFunc,
	seek AvIOContextSeekFunc) *AvIOContext {
	return (*AvIOContext)(C.avio_alloc_context((*C.uint8_t)(buffer), (C.int)(bufferSize),
		(C.int)(writeFlag), opaque,
		(C.avio_context_read_packet_func)(readPacket),
		(C.avio_context_write_packet_func)(writePacket),
		(C.avio_context_seek_func)(seek)))
}

// AvIOContextFree frees the supplied IO context and everything associated with it.
func AvIOContextFree(s **AvIOContext) {
	C.avio_context_free((**C.struct_AVIOContext)(unsafe.Pointer(s)))
}

// AvIOW8
func AvIOW8(s *AvIOContext, b int32) {
	C.avio_w8((*C.struct_AVIOContext)(s), (C.int)(b))
}

// AvIOWrite
func AvIOWrite(s *AvIOContext, buf *uint8, size int32) {
	C.avio_write((*C.struct_AVIOContext)(s), (*C.uchar)(buf), (C.int)(size))
}

// AvIOWl64
func AvIOWl64(s *AvIOContext, val uint64) {
	C.avio_wl64((*C.struct_AVIOContext)(s), (C.uint64_t)(val))
}

// AvIOWb64
func AvIOWb64(s *AvIOContext, val uint64) {
	C.avio_wb64((*C.struct_AVIOContext)(s), (C.uint64_t)(val))
}

// AvIOWl32
func AvIOWl32(s *AvIOContext, val uint32) {
	C.avio_wl32((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWb32
func AvIOWb32(s *AvIOContext, val uint32) {
	C.avio_wb32((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWl24
func AvIOWl24(s *AvIOContext, val uint32) {
	C.avio_wl24((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWb24
func AvIOWb24(s *AvIOContext, val uint32) {
	C.avio_wb24((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWl16
func AvIOWl16(s *AvIOContext, val uint32) {
	C.avio_wl16((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOWb16
func AvIOWb16(s *AvIOContext, val uint32) {
	C.avio_wb16((*C.struct_AVIOContext)(s), (C.uint32_t)(val))
}

// AvIOPutStr Write a string.
func AvIOPutStr(s *AvIOContext, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.avio_put_str((*C.struct_AVIOContext)(s), (*C.char)(strPtr)))
}

// AvIOPutStr16le converts an UTF-8 string to UTF-16LE and write it.
func AvIOPutStr16le(s *AvIOContext, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.avio_put_str16le((*C.struct_AVIOContext)(s), (*C.char)(strPtr)))
}

// AvIOPutStr16be converts an UTF-8 string to UTF-16BE and write it.
func AvIOPutStr16be(s *AvIOContext, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.avio_put_str16be((*C.struct_AVIOContext)(s), (*C.char)(strPtr)))
}

// AvIOWriteMarker
func AvIOWriteMarker(s *AvIOContext, time int64, _type AvIODataMarkerType) {
	C.avio_write_marker((*C.struct_AVIOContext)(s), (C.int64_t)(time), (C.enum_AVIODataMarkerType)(_type))
}

const (
	AVSEEK_SIZE  = C.AVSEEK_SIZE
	AVSEEK_FORCE = C.AVSEEK_FORCE
)

// AvIOSeek equivalents fseek().
func AvIOSeek(s *AvIOContext, offset int64, whence int32) int64 {
	return (int64)(C.avio_seek((*C.struct_AVIOContext)(s), (C.int64_t)(offset), (C.int)(whence)))
}

// AvIOSkip skips given number of bytes forward.
func AvIOSkip(s *AvIOContext, offset int64) int64 {
	return (int64)(C.avio_skip((*C.struct_AVIOContext)(s), (C.int64_t)(offset)))
}

// AvIOTell equivalents ftell().
func AvIOTell(s *AvIOContext) int64 {
	return (int64)(C.avio_tell((*C.struct_AVIOContext)(s)))
}

// AvIOSize gets the filesize.
func AvIOSize(s *AvIOContext) int64 {
	return (int64)(C.avio_size((*C.struct_AVIOContext)(s)))
}

// AvIOFeof similar to feof() but also returns nonzero on read errors.
func AvIOFeof(s *AvIOContext) int32 {
	return (int32)(C.avio_feof((*C.struct_AVIOContext)(s)))
}

// AvIOPrintf Writes a formatted string to the context.
func AvIOPrintf(s *AvIOContext, _fmt string, va ...any) int32 {
	fmtPtr, fmtFunc := StringCasting(fmt.Sprintf(_fmt, va...))
	defer fmtFunc()
	return (int32)(C.avio_printf_wrap((*C.struct_AVIOContext)(s), (*C.char)(fmtPtr)))
}

// NONEED: avio_print_string_array

// AvIOPrint
func AvIOPrint(s *AvIOContext, va ...any) {
	fmtPtr, fmtFunc := StringCasting(fmt.Sprint(va...))
	defer fmtFunc()
	fmtArray := []*C.char{(*C.char)(fmtPtr), nil}
	C.avio_print_string_array((*C.struct_AVIOContext)(s), &fmtArray[0])
}

// AvIOFlush forces flushing of buffered data.
func AvIOFlush(s *AvIOContext) {
	C.avio_flush((*C.struct_AVIOContext)(s))
}

// AvIORead reads size bytes from AVIOContext into buf.
func AvIORead(s *AvIOContext, buf *uint8, size int32) int32 {
	return (int32)(C.avio_read((*C.struct_AVIOContext)(s), (*C.uchar)(buf), (C.int)(size)))
}

// AvIOReadPartial sead size bytes from AVIOContext into buf. Unlike avio_read(), this is allowed
// to read fewer bytes than requested. The missing bytes can be read in the next
// call. This always tries to read at least 1 byte.
// Useful to reduce latency in certain cases.
func AvIOReadPartial(s *AvIOContext, buf *uint8, size int32) int32 {
	return (int32)(C.avio_read_partial((*C.struct_AVIOContext)(s), (*C.uchar)(buf), (C.int)(size)))
}

// AvIOR8
func AvIOR8(s *AvIOContext) int32 {
	return (int32)(C.avio_r8((*C.struct_AVIOContext)(s)))
}

// AvIORl16
func AvIORl16(s *AvIOContext) uint32 {
	return (uint32)(C.avio_rl16((*C.struct_AVIOContext)(s)))
}

// AvIORl24
func AvIORl24(s *AvIOContext) uint32 {
	return (uint32)(C.avio_rl24((*C.struct_AVIOContext)(s)))
}

// AvIORl32
func AvIORl32(s *AvIOContext) uint32 {
	return (uint32)(C.avio_rl32((*C.struct_AVIOContext)(s)))
}

// AvIORl64
func AvIORl64(s *AvIOContext) uint64 {
	return (uint64)(C.avio_rl64((*C.struct_AVIOContext)(s)))
}

// AvIORb16
func AvIORb16(s *AvIOContext) uint32 {
	return (uint32)(C.avio_rb16((*C.struct_AVIOContext)(s)))
}

// AvIORb24
func AvIORb24(s *AvIOContext) uint32 {
	return (uint32)(C.avio_rb24((*C.struct_AVIOContext)(s)))
}

// AvIORb32
func AvIORb32(s *AvIOContext) uint32 {
	return (uint32)(C.avio_rb32((*C.struct_AVIOContext)(s)))
}

// AvIORb64
func AvIORb64(s *AvIOContext) uint64 {
	return (uint64)(C.avio_rb64((*C.struct_AVIOContext)(s)))
}

// AvIOGetStr reads a string from pb into buf. The reading will terminate when either
// a NULL character was encountered, maxlen bytes have been read, or nothing
// more can be read from pb. The result is guaranteed to be NULL-terminated, it
// will be truncated if buf is too small.
// Note that the string is not interpreted or validated in any way, it
// might get truncated in the middle of a sequence for multi-byte encodings.
func AvIOGetStr(s *AvIOContext, maxLen int32, buf *int8, buflen int32) int32 {
	return (int32)(C.avio_get_str((*C.struct_AVIOContext)(s), (C.int)(maxLen), (*C.char)(buf), (C.int)(buflen)))
}

// AvIOGetStr16le reads a UTF-16LE string from pb and convert it to UTF-8.
func AvIOGetStr16le(s *AvIOContext, maxLen int32, buf *int8, buflen int32) int32 {
	return (int32)(C.avio_get_str16le((*C.struct_AVIOContext)(s), (C.int)(maxLen), (*C.char)(buf), (C.int)(buflen)))
}

// AvIOGetStr16be reads a UTF-16BE string from pb and convert it to UTF-8.
func AvIOGetStr16be(s *AvIOContext, maxLen int32, buf *int8, buflen int32) int32 {
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
func AvIOOpen(s **AvIOContext, url string, flags int32) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avio_open((**C.struct_AVIOContext)(unsafe.Pointer(s)),
		(*C.char)(urlPtr), (C.int)(flags)))
}

// avio_open2 creates and initializes a AVIOContext for accessing the resource indicated by url.
func avio_open2(s **AvIOContext, url string, flags int32,
	intCb *AvIOInterruptCB, options **AvDictionary) int32 {
	urlPtr, urlFunc := StringCasting(url)
	defer urlFunc()
	return (int32)(C.avio_open2((**C.struct_AVIOContext)(unsafe.Pointer(s)),
		(*C.char)(urlPtr), (C.int)(flags),
		(*C.struct_AVIOInterruptCB)(intCb), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvIOClose
func AvIOClose(s *AvIOContext) int32 {
	return (int32)(C.avio_close((*C.struct_AVIOContext)(s)))
}

// AvIOClosep closes the resource accessed by the AVIOContext *s, free it
// and set the pointer pointing to it to NULL.
func AvIOClosep(s **AvIOContext) int32 {
	return (int32)(C.avio_closep((**C.struct_AVIOContext)(unsafe.Pointer(s))))
}

// AvIOOpenDynBuf opens a write only memory stream.
func AvIOOpenDynBuf(s **AvIOContext) int32 {
	return (int32)(C.avio_open_dyn_buf((**C.struct_AVIOContext)(unsafe.Pointer(s))))
}

// AvIOGetDynBuf returns the written size and a pointer to the buffer.
func AvIOGetDynBuf(s *AvIOContext, pbuffer **uint8) int32 {
	return (int32)(C.avio_get_dyn_buf((*C.struct_AVIOContext)(s),
		(**C.uint8_t)(unsafe.Pointer(pbuffer))))
}

// AvIOCloseDynBuf returns the written size and a pointer to the buffer. The buffer
// must be freed with AvFree().
func AvIOCloseDynBuf(s *AvIOContext, pbuffer **uint8) int32 {
	return (int32)(C.avio_close_dyn_buf((*C.struct_AVIOContext)(s),
		(**C.uint8_t)(unsafe.Pointer(pbuffer))))
}

// AvIOEnumProtocols iterates through names of available protocols.
func AvIOEnumProtocols(opaque *unsafe.Pointer, output int32) string {
	return C.GoString(C.avio_enum_protocols(opaque, (C.int)(output)))
}

// AvIOProtocolGetClass gets AVClass by names of available protocols.
func AvIOProtocolGetClass(name string) *AvClass {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AvClass)(C.avio_protocol_get_class((*C.char)(namePtr)))
}

// AvIOPause pauses and resumes playing - only meaningful if using a network streaming
// protocol (e.g. MMS).
func AvIOPause(s *AvIOContext, pause int32) int32 {
	return (int32)(C.avio_pause((*C.struct_AVIOContext)(s), (C.int)(pause)))
}

// AvIOSeekTime seeks to a given timestamp relative to some component stream.
// Only meaningful if using a network streaming protocol (e.g. MMS.).
func AvIOSeekTime(s *AvIOContext, streamIndex int32, timestamp int64, flags int32) int64 {
	return (int64)(C.avio_seek_time((*C.struct_AVIOContext)(s),
		(C.int)(streamIndex), (C.int64_t)(timestamp), (C.int)(flags)))
}

type AvBPrint C.struct_AVBPrint

// avio_read_to_bprint
func avio_read_to_bprint(s *AvIOContext, pb *AvBPrint, maxSize uint) int32 {
	return (int32)(C.avio_read_to_bprint((*C.struct_AVIOContext)(s),
		(*C.struct_AVBPrint)(pb), (C.size_t)(maxSize)))
}

// AvIOAccept accepts and allocates a client context on a server context.
func AvIOAccept(s *AvIOContext, c **AvIOContext) int32 {
	return (int32)(C.avio_accept((*C.struct_AVIOContext)(s), (**C.struct_AVIOContext)(unsafe.Pointer(c))))
}

// AvIOHandshake performs one step of the protocol handshake to accept a new client.
func AvIOHandshake(s *AvIOContext) int32 {
	return (int32)(C.avio_handshake((*C.struct_AVIOContext)(s)))
}
