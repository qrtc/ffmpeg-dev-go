package ffmpeg

/*
#include <libavutil/fifo.h>

typedef void (*av_fifo_peek_func)(void*, void*, int);
typedef void (*av_fifo_read_func)(void*, void*, int);
typedef int (*av_fifo_write_func)(void*, void*, int);

*/
import "C"
import "unsafe"

// AVFifoBuffer
type AVFifoBuffer C.struct_AVFifoBuffer

// AvFifoAlloc initializes an AVFifoBuffer.
func AvFifoAlloc[T HelperInteger](size T) *AVFifoBuffer {
	return (*AVFifoBuffer)(C.av_fifo_alloc((C.uint)(size)))
}

// AvFifoAllocArray initializes an AVFifoBuffer.
func AvFifoAllocArray[U, V HelperInteger](nmemb U, size V) *AVFifoBuffer {
	return (*AVFifoBuffer)(C.av_fifo_alloc_array((C.size_t)(nmemb), (C.size_t)(size)))
}

// AvFifoFree frees an AVFifoBuffer.
func AvFifoFree(f *AVFifoBuffer) {
	C.av_fifo_free((*C.struct_AVFifoBuffer)(f))
}

// AvFifoFreep frees an AVFifoBuffer and reset pointer to NULL.
func AvFifoFreep(f **AVFifoBuffer) {
	C.av_fifo_freep((**C.struct_AVFifoBuffer)(unsafe.Pointer(f)))
}

// AvFifoReset resets the AVFifoBuffer to the state right after AvFifoAlloc, in particular it is emptied.
func AvFifoReset(f *AVFifoBuffer) {
	C.av_fifo_reset((*C.struct_AVFifoBuffer)(f))
}

// AvFifoSize returns the amount of data in bytes in the AVFifoBuffer, that is the
// amount of data you can read from it.
func AvFifoSize(f *AVFifoBuffer) int32 {
	return (int32)(C.av_fifo_size((*C.struct_AVFifoBuffer)(f)))
}

// AvFifoSpace returns the amount of space in bytes in the AVFifoBuffer, that is the
// amount of data you can write into it.
func AvFifoSpace(f *AVFifoBuffer) int32 {
	return (int32)(C.av_fifo_space((*C.struct_AVFifoBuffer)(f)))
}

// typedef void (*av_fifo_peek_func)(void*, void*, int);
type AvFifoPeekFunc = C.av_fifo_peek_func

// AvFifoGenericPeekAt feeds data at specific position from an AVFifoBuffer to a user-supplied callback.
func AvFifoGenericPeekAt(f *AVFifoBuffer, dest CVoidPointer,
	offset, bufSize int32, fn AvFifoPeekFunc) int32 {
	return (int32)(C.av_fifo_generic_peek_at((*C.struct_AVFifoBuffer)(f),
		VoidPointer(dest), (C.int)(offset), (C.int)(bufSize), (C.av_fifo_peek_func)(fn)))
}

// AvFifoGenericPeek feeds data from an AVFifoBuffer to a user-supplied callback.
func AvFifoGenericPeek(f *AVFifoBuffer, dest CVoidPointer,
	bufSize int32, fn AvFifoPeekFunc) int32 {
	return (int32)(C.av_fifo_generic_peek((*C.struct_AVFifoBuffer)(f),
		VoidPointer(dest), (C.int)(bufSize), (C.av_fifo_peek_func)(fn)))
}

// typedef void (*av_fifo_read_func)(void*, void*, int);
type AvFifoReadFunc = C.av_fifo_read_func

// AvFifoGenericRead feeds data from an AVFifoBuffer to a user-supplied callback.
func AvFifoGenericRead(f *AVFifoBuffer, dest CVoidPointer,
	bufSize int32, fn AvFifoReadFunc) int32 {
	return (int32)(C.av_fifo_generic_read((*C.struct_AVFifoBuffer)(f),
		VoidPointer(dest), (C.int)(bufSize), (C.av_fifo_read_func)(fn)))
}

// typedef int (*av_fifo_write_func)(void*, void*, int);
type AvFifoWriteFunc = C.av_fifo_write_func

// AvFifoGenericWrite feeds data from a user-supplied callback to an AVFifoBuffer.
func AvFifoGenericWrite(f *AVFifoBuffer, src CVoidPointer,
	size int32, fn AvFifoWriteFunc) int32 {
	return (int32)(C.av_fifo_generic_write((*C.struct_AVFifoBuffer)(f),
		VoidPointer(src), (C.int)(size), (C.av_fifo_write_func)(fn)))
}

// AvFifoRealloc2 resizes an AVFifoBuffer.
func AvFifoRealloc2[T HelperInteger](f *AVFifoBuffer, size T) int32 {
	return (int32)(C.av_fifo_realloc2((*C.struct_AVFifoBuffer)(f), (C.uint)(size)))
}

// AvFifoGrow enlarges an AVFifoBuffer.
func AvFifoGrow[T HelperInteger](f *AVFifoBuffer, additionalSpace T) int32 {
	return (int32)(C.av_fifo_grow((*C.struct_AVFifoBuffer)(f), (C.uint)(additionalSpace)))
}

// AvFifoDrain reads and discards the specified amount of data from an AVFifoBuffer.
func AvFifoDrain[T HelperInteger](f *AVFifoBuffer, size T) {
	C.av_fifo_drain((*C.struct_AVFifoBuffer)(f), (C.int)(size))
}

// AvFifoPeek2 returns a pointer to the data stored in a FIFO buffer at a certain offset.
// The FIFO buffer is not modified.
func AvFifoPeek2[T HelperInteger](f *AVFifoBuffer, offs T) *uint8 {
	return (*uint8)(C.av_fifo_peek2((*C.struct_AVFifoBuffer)(f), (C.int)(offs)))
}
