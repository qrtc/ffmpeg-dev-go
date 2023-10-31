// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/fifo.h>

typedef void (*av_fifo_peek_func)(void*, void*, int);
typedef void (*av_fifo_read_func)(void*, void*, int);
typedef int (*av_fifo_write_func)(void*, void*, int);

*/
import "C"
import "unsafe"

// AVFifo
type AVFifo C.struct_AVFifo

// typedef int AVFifoCB(void *opaque, void *buf, size_t *nb_elems);
type AVFifoCB = C.AVFifoCB

const (
	AV_FIFO_FLAG_AUTO_GROW = C.AV_FIFO_FLAG_AUTO_GROW
)

// AvFifoAlloc2 allocates and initializes an AVFifo with a given element size.
func AvFifoAlloc2[U, V Integer](elems U, elemSize V, flags uint32) *AVFifo {
	return (*AVFifo)(C.av_fifo_alloc2((C.size_t)(elems), (C.size_t)(elemSize), (C.uint)(flags)))
}

// AvFifoElemSize returns Element size for FIFO operations.
func AvFifoElemSize(f *AVFifo) uintptr {
	return (uintptr)(C.av_fifo_elem_size((*C.struct_AVFifo)(f)))
}

// AvFifoAutoGrowLimit sets the maximum size (in elements) to which the FIFO can be resized
// automatically.
func AvFifoAutoGrowLimit[T Integer](f *AVFifo, maxElems T) {
	C.av_fifo_auto_grow_limit((*C.struct_AVFifo)(f), (C.size_t)(maxElems))
}

// AvFifoCanRead returns number of elements available for reading from the given FIFO.
func AvFifoCanRead(f *AVFifo) uintptr {
	return (uintptr)(C.av_fifo_can_read((*C.struct_AVFifo)(f)))
}

// AvFifoCanWrite returns number of elements that can be written into the given FIFO.
func AvFifoCanWrite(f *AVFifo) uintptr {
	return (uintptr)(C.av_fifo_can_write((*C.struct_AVFifo)(f)))
}

// AvFifoGrow2 enlarges an AVFifo.
func AvFifoGrow2[T Integer](f *AVFifo, inc T) int32 {
	return (int32)(C.av_fifo_grow2((*C.struct_AVFifo)(f), (C.size_t)(inc)))
}

// AvFifoWrite writes data into a FIFO.
func AvFifoWrite[T Integer](f *AVFifo, buf CVoidPointer, nbElems T) int32 {
	return (int32)(C.av_fifo_write((*C.struct_AVFifo)(f), VoidPointer(buf), (C.size_t)(nbElems)))
}

// AvFifoWriteFromCb writes data from a user-provided callback into a FIFO.
func AvFifoWriteFromCb(f *AVFifo, readCb *AVFifoCB, opaque CVoidPointer, nbElems *uintptr) int32 {
	return (int32)(C.av_fifo_write_from_cb((*C.struct_AVFifo)(f),
		(*C.AVFifoCB)(readCb), VoidPointer(opaque), (*C.size_t)(unsafe.Pointer(nbElems))))
}

// AvFifoRead reads data from a FIFO.
func AvFifoRead[T Integer](f *AVFifo, buf CVoidPointer, nbElems T) int32 {
	return (int32)(C.av_fifo_read((*C.struct_AVFifo)(f), VoidPointer(buf), (C.size_t)(nbElems)))
}

// AvFifoReadToCb feeds data from a FIFO into a user-provided callback.
func AvFifoReadToCb(f *AVFifo, writeCb *AVFifoCB, opaque CVoidPointer, nbElems *uintptr) int32 {
	return (int32)(C.av_fifo_read_to_cb((*C.struct_AVFifo)(f),
		(*C.AVFifoCB)(writeCb), VoidPointer(opaque), (*C.size_t)(unsafe.Pointer(nbElems))))
}

// AvFifoPeek reads data from a FIFO without modifying FIFO state.
func AvFifoPeek[U, V Integer](f *AVFifo, buf CVoidPointer, nbElems U, offset V) int32 {
	return (int32)(C.av_fifo_peek((*C.struct_AVFifo)(f), VoidPointer(buf),
		(C.size_t)(nbElems), (C.size_t)(offset)))
}

// AvFifoPeekToCb feeds data from a FIFO into a user-provided callback.
func AvFifoPeekToCb[T Integer](f *AVFifo, writeCb *AVFifoCB, opaque CVoidPointer,
	nbElems *uintptr, offset T) int32 {
	return (int32)(C.av_fifo_peek_to_cb((*C.struct_AVFifo)(f),
		(*C.AVFifoCB)(writeCb), VoidPointer(opaque),
		(*C.size_t)(unsafe.Pointer(nbElems)), (C.size_t)(offset)))
}

// AvFifoDrain2 discards the specified amount of data from an AVFifo.
func AvFifoDrain2[T Integer](f *AVFifo, size T) {
	C.av_fifo_drain2((*C.struct_AVFifo)(f), (C.size_t)(size))
}

// AvFifoReset2 empties the AVFifo.
func AvFifoReset2(f *AVFifo) {
	C.av_fifo_reset2((*C.struct_AVFifo)(f))
}

// AvFifoFreep2 frees an AVFifo and reset pointer to NULL.
func AvFifoFreep2(f **AVFifo) {
	C.av_fifo_freep2((**C.struct_AVFifo)(unsafe.Pointer(f)))
}

// AVFifoBuffer
type AVFifoBuffer C.struct_AVFifoBuffer

// Deprecated: Use AvFifoAlloc2() instread.
//
// AvFifoAlloc initializes an AVFifoBuffer.
func AvFifoAlloc[T Integer](size T) *AVFifoBuffer {
	return (*AVFifoBuffer)(C.av_fifo_alloc((C.uint)(size)))
}

// Deprecated: Use AvFifoAlloc2() instread.
//
// AvFifoAllocArray initializes an AVFifoBuffer.
func AvFifoAllocArray[U, V Integer](nmemb U, size V) *AVFifoBuffer {
	return (*AVFifoBuffer)(C.av_fifo_alloc_array((C.size_t)(nmemb), (C.size_t)(size)))
}

// Deprecated: Use AvFifoFreep2() instread.
//
// AvFifoFree frees an AVFifoBuffer.
func AvFifoFree(f *AVFifoBuffer) {
	C.av_fifo_free((*C.struct_AVFifoBuffer)(f))
}

// Deprecated: Use AvFifoFreep2() instread.
//
// AvFifoFreep frees an AVFifoBuffer and reset pointer to NULL.
func AvFifoFreep(f **AVFifoBuffer) {
	C.av_fifo_freep((**C.struct_AVFifoBuffer)(unsafe.Pointer(f)))
}

// Deprecated: Use AvFifoReset2() instread.
//
// AvFifoReset resets the AVFifoBuffer to the state right after AvFifoAlloc, in particular it is emptied.
func AvFifoReset(f *AVFifoBuffer) {
	C.av_fifo_reset((*C.struct_AVFifoBuffer)(f))
}

// Deprecated: Use AvFifoCanRead() instread.
//
// AvFifoSize returns the amount of data in bytes in the AVFifoBuffer, that is the
// amount of data you can read from it.
func AvFifoSize(f *AVFifoBuffer) int32 {
	return (int32)(C.av_fifo_size((*C.struct_AVFifoBuffer)(f)))
}

// Deprecated: Use AvFifoCanWrite() instread.
//
// AvFifoSpace returns the amount of space in bytes in the AVFifoBuffer, that is the
// amount of data you can write into it.
func AvFifoSpace(f *AVFifoBuffer) int32 {
	return (int32)(C.av_fifo_space((*C.struct_AVFifoBuffer)(f)))
}

// Deprecated: No use.
//
// typedef void (*av_fifo_peek_func)(void*, void*, int);
type AvFifoPeekFunc = C.av_fifo_peek_func

// Deprecated: Use the new AVFifo-API with AvFifoPeek() or AvFifoPeekToCb() instread.
//
// AvFifoGenericPeekAt feeds data at specific position from an AVFifoBuffer to a user-supplied callback.
func AvFifoGenericPeekAt(f *AVFifoBuffer, dest CVoidPointer,
	offset, bufSize int32, fn AvFifoPeekFunc) int32 {
	return (int32)(C.av_fifo_generic_peek_at((*C.struct_AVFifoBuffer)(f),
		VoidPointer(dest), (C.int)(offset), (C.int)(bufSize), (C.av_fifo_peek_func)(fn)))
}

// Deprecated: Use the new AVFifo-API with AvFifoPeek() or AvFifoPeekToCb() instread.
//
// AvFifoGenericPeek feeds data from an AVFifoBuffer to a user-supplied callback.
func AvFifoGenericPeek(f *AVFifoBuffer, dest CVoidPointer,
	bufSize int32, fn AvFifoPeekFunc) int32 {
	return (int32)(C.av_fifo_generic_peek((*C.struct_AVFifoBuffer)(f),
		VoidPointer(dest), (C.int)(bufSize), (C.av_fifo_peek_func)(fn)))
}

// typedef void (*av_fifo_read_func)(void*, void*, int);
type AvFifoReadFunc = C.av_fifo_read_func

// Deprecated: Use the new AVFifo-API with AvFifoRead() or AvFifoReadToCb() instread.
//
// AvFifoGenericRead feeds data from an AVFifoBuffer to a user-supplied callback.
func AvFifoGenericRead(f *AVFifoBuffer, dest CVoidPointer,
	bufSize int32, fn AvFifoReadFunc) int32 {
	return (int32)(C.av_fifo_generic_read((*C.struct_AVFifoBuffer)(f),
		VoidPointer(dest), (C.int)(bufSize), (C.av_fifo_read_func)(fn)))
}

// typedef int (*av_fifo_write_func)(void*, void*, int);
type AvFifoWriteFunc = C.av_fifo_write_func

// Deprecated: Use the new AVFifo-API with AvFifoWrite() or AvFifoWriteFromCb() instread.
//
// AvFifoGenericWrite feeds data from a user-supplied callback to an AVFifoBuffer.
func AvFifoGenericWrite(f *AVFifoBuffer, src CVoidPointer,
	size int32, fn AvFifoWriteFunc) int32 {
	return (int32)(C.av_fifo_generic_write((*C.struct_AVFifoBuffer)(f),
		VoidPointer(src), (C.int)(size), (C.av_fifo_write_func)(fn)))
}

// Deprecated: Use the new AVFifo-API with AvFifoGrow2() instread.
//
// AvFifoRealloc2 resizes an AVFifoBuffer.
func AvFifoRealloc2[T Integer](f *AVFifoBuffer, size T) int32 {
	return (int32)(C.av_fifo_realloc2((*C.struct_AVFifoBuffer)(f), (C.uint)(size)))
}

// Deprecated: Use the new AVFifo-API with av_fifo_grow2() instread.
//
// AvFifoGrow enlarges an AVFifoBuffer.
func AvFifoGrow[T Integer](f *AVFifoBuffer, additionalSpace T) int32 {
	return (int32)(C.av_fifo_grow((*C.struct_AVFifoBuffer)(f), (C.uint)(additionalSpace)))
}

// Deprecated: Use the new AVFifo-API with AvFifoDrain2() instread.
//
// AvFifoDrain reads and discards the specified amount of data from an AVFifoBuffer.
func AvFifoDrain[T Integer](f *AVFifoBuffer, size T) {
	C.av_fifo_drain((*C.struct_AVFifoBuffer)(f), (C.int)(size))
}

// Deprecated: Use the new AVFifo-API with AvFifoPeek() or AvFifoPeekToCb() instread.
//
// AvFifoPeek2 returns a pointer to the data stored in a FIFO buffer at a certain offset.
// The FIFO buffer is not modified.
func AvFifoPeek2[T Integer](f *AVFifoBuffer, offs T) *uint8 {
	return (*uint8)(C.av_fifo_peek2((*C.struct_AVFifoBuffer)(f), (C.int)(offs)))
}
