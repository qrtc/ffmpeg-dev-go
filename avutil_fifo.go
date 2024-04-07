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
