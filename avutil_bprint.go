package ffmpeg

/*
#include <libavutil/bprint.h>

void av_bprintf_wrap(AVBPrint *buf, const char *fmt) {
	av_bprintf(buf, fmt, NULL);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type AVBPrint C.struct_AVBPrint

const (
	AV_BPRINT_SIZE_UNLIMITED  = C.AV_BPRINT_SIZE_UNLIMITED
	AV_BPRINT_SIZE_AUTOMATIC  = C.AV_BPRINT_SIZE_AUTOMATIC
	AV_BPRINT_SIZE_COUNT_ONLY = C.AV_BPRINT_SIZE_COUNT_ONLY
)

// AvBprintInit initializes a print buffer.
func AvBprintInit(buf *AVBPrint, sizeInit, sizeMax uint32) {
	C.av_bprint_init((*C.struct_AVBPrint)(buf), (C.uint)(sizeInit), (C.uint)(sizeMax))
}

// AvBprintInitForBuffer initializes a print buffer using a pre-existing buffer.
func AvBprintInitForBuffer(buf *AVBPrint, buffer *int8, size uint32) {
	C.av_bprint_init_for_buffer((*C.struct_AVBPrint)(buf), (*C.char)(buffer), (C.uint)(size))
}

// AvBprintf appends a formatted string to a print buffer.
func AvBprintf(buf *AVBPrint, _fmt string, va ...any) {
	strPtr, strFunc := StringCasting(fmt.Sprintf(_fmt, va...))
	defer strFunc()
	C.av_bprintf_wrap((*C.struct_AVBPrint)(buf), (*C.char)(strPtr))
}

// NONEED: av_vbprintf

// AvBprintChars appends char c n times to a print buffer.
func AvBprintChars(buf *AVBPrint, c int8, n uint32) {
	C.av_bprint_chars((*C.struct_AVBPrint)(buf), (C.char)(c), (C.uint)(n))
}

// AvBprintAppendData appends data to a print buffer.
func AvBprintAppendData(buf *AVBPrint, data *int8, size uint32) {
	C.av_bprint_append_data((*C.struct_AVBPrint)(buf), (*C.char)(data), (C.uint)(size))
}

// NONEED: av_bprint_strftime

// AvBprintGetBuffer allocates bytes in the buffer for external use.
func AvBprintGetBuffer(buf *AVBPrint, size uint32, mem **uint8, actualSize *uint32) {
	C.av_bprint_get_buffer((*C.struct_AVBPrint)(buf), (C.uint)(size),
		(**C.uint8_t)(unsafe.Pointer(mem)), (*C.uint)(actualSize))
}

// AvBprintClear resets the string to "" but keep internal allocated data.
func AvBprintClear(buf *AVBPrint) {
	C.av_bprint_clear((*C.struct_AVBPrint)(buf))
}

// AvBprintIsComplete tests if the print buffer is complete (not truncated).
func AvBprintIsComplete(buf *AVBPrint) int32 {
	return (int32)(C.av_bprint_is_complete((*C.struct_AVBPrint)(buf)))
}

// AvBprintFinalize finalizes a print buffer.
func AvBprintFinalize(buf *AVBPrint, retStr **int8) int32 {
	return (int32)(C.av_bprint_finalize((*C.struct_AVBPrint)(buf),
		(**C.char)(unsafe.Pointer(retStr))))
}

// AvBprintEscape escapes the content in src and append it to dstbuf.
func AvBprintEscape(dstbuf *AVBPrint, src, specialChars *int8, mode AVEscapeMode, flags int32) {
	C.av_bprint_escape((*C.struct_AVBPrint)(dstbuf),
		(*C.char)(src), (*C.char)(specialChars),
		(C.enum_AVEscapeMode)(mode), (C.int)(flags))
}
