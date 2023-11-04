// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/buffer.h>

typedef void (*av_buffer_free_func)(void *opaque, uint8_t *data);
typedef AVBufferRef* (*av_buffer_pool_alloc_func)(int size);
typedef AVBufferRef* (*av_buffer_pool_alloc2_func)(void* opaque, int size);
typedef void (*av_buffer_pool_free_func)(void* opaque);
*/
import "C"
import "unsafe"

// AVBufferRef
type AVBuffer C.struct_AVBuffer

// AVBufferRef
type AVBufferRef C.struct_AVBufferRef

// GetBuffer gets `AVBufferRef.buffer` value.
func (bf *AVBufferRef) GetBuffer() *AVBuffer {
	return (*AVBuffer)(bf.buffer)
}

// GetData gets `AVBufferRef.data` value.
func (bf *AVBufferRef) GetData() *uint8 {
	return (*uint8)(bf.data)
}

// GetSize gets `AVBufferRef.size` value.
func (bf *AVBufferRef) GetSize() int32 {
	return (int32)(bf.size)
}

// AvBufferAlloc allocates an AVBuffer of the given size using AvMalloc().
func AvBufferAlloc(size int32) *AVBufferRef {
	return (*AVBufferRef)(C.av_buffer_alloc((C.int)(size)))
}

// AvBufferAllocz same as AVBufferAlloc(), except the returned buffer will be initialized
// to zero.
func AvBufferAllocz(size int32) *AVBufferRef {
	return (*AVBufferRef)(C.av_buffer_allocz((C.int)(size)))
}

const (
	AV_BUFFER_FLAG_READONLY = C.AV_BUFFER_FLAG_READONLY
)

// typedef void (*av_buffer_free_func)(void *opaque, uint8_t *data);
type AVBufferFreeFunc = C.av_buffer_free_func

// AvBufferCreate Create an AVBuffer from an existing array.
func AvBufferCreate(data *uint8, size int32, free AVBufferFreeFunc, opaque CVoidPointer, flags int32) *AVBufferRef {
	return (*AVBufferRef)(C.av_buffer_create((*C.uint8_t)(data), (C.int)(size),
		(C.av_buffer_free_func)(free), VoidPointer(opaque), (C.int)(flags)))
}

// AvBufferDefaultFree frees buffer data.
func AvBufferDefaultFree(opaque CVoidPointer, data *uint8) {
	C.av_buffer_default_free(VoidPointer(opaque), (*C.uint8_t)(data))
}

// AvBufferRef creates a new reference to an AVBuffer.
func AvBufferRef(buf *AVBufferRef) *AVBufferRef {
	return (*AVBufferRef)(C.av_buffer_ref((*C.struct_AVBufferRef)(buf)))
}

// AvBufferUnref frees a given reference and automatically free the buffer if there are no more
// references to it.
func AvBufferUnref(buf **AVBufferRef) {
	C.av_buffer_unref((**C.struct_AVBufferRef)(unsafe.Pointer(buf)))
}

// AvBufferIsWritable returns 1 if the caller may write to the data referred to by buf (which is
// true if and only if buf is the only reference to the underlying AVBuffer).
// Return 0 otherwise.
func AvBufferIsWritable(buf *AVBufferRef) int32 {
	return (int32)(C.av_buffer_is_writable((*C.struct_AVBufferRef)(buf)))
}

// AvBufferGetOpaque returns the opaque parameter set by AVBufferCreate.
func AvBufferGetOpaque(buf *AVBufferRef) unsafe.Pointer {
	return C.av_buffer_get_opaque((*C.struct_AVBufferRef)(buf))
}

// AvBufferGetRefCount
func AvBufferGetRefCount(buf *AVBufferRef) int32 {
	return (int32)(C.av_buffer_get_ref_count((*C.struct_AVBufferRef)(buf)))
}

// AvBufferMakeWritable creates a writable reference from a given buffer reference,
// avoiding data copy if possible.
func AvBufferMakeWritable(buf **AVBufferRef) int32 {
	return (int32)(C.av_buffer_make_writable((**C.struct_AVBufferRef)(unsafe.Pointer(buf))))
}

// AvBufferRealloc reallocates a given buffer.
func AvBufferRealloc(buf **AVBufferRef, size int32) int32 {
	return (int32)(C.av_buffer_realloc((**C.struct_AVBufferRef)(unsafe.Pointer(buf)), (C.int)(size)))
}

// AVBufferPool
type AVBufferPool C.struct_AVBufferPool

// typedef AVBufferRef* (*av_buffer_pool_alloc_func)(int size);
type AVBufferPoolAllocFunc = C.av_buffer_pool_alloc_func

// typedef AVBufferRef* (*av_buffer_pool_alloc2_func)(void* opaque, int size);
type AVBufferPoolAlloc2Func = C.av_buffer_pool_alloc2_func

// typedef void (*av_buffer_pool_free_func)(void* opaque);
type AVBufferPoolFreeFunc = C.av_buffer_pool_free_func

// AvBufferPoolInit allocates and initializes a buffer pool.
func AvBufferPoolInit(size int32, alloc AVBufferPoolAllocFunc) *AVBufferPool {
	return (*AVBufferPool)(C.av_buffer_pool_init((C.int)(size), (C.av_buffer_pool_alloc_func)(alloc)))
}

// AvBufferPoolInit2 allocates and initialize a buffer pool with a more complex allocator.
func AvBufferPoolInit2(size int32, opaque CVoidPointer,
	alloc AVBufferPoolAllocFunc, free AVBufferPoolFreeFunc) *AVBufferPool {
	return (*AVBufferPool)(C.av_buffer_pool_init2((C.int)(size), VoidPointer(opaque),
		(C.av_buffer_pool_alloc_func)(alloc),
		(C.av_buffer_pool_free_func)(free)))
}

// AvBufferPoolUninit marks the pool as being available for freeing.
func AvBufferPoolUninit(pool **AVBufferPool) {
	C.av_buffer_pool_uninit((**C.struct_AVBufferPool)(unsafe.Pointer(pool)))
}

// AvBufferPoolGet allocates a new AVBuffer, reusing an old buffer from the pool when available.
func AvBufferPoolGet(pool *AVBufferPool) *AVBufferRef {
	return (*AVBufferRef)(C.av_buffer_pool_get((*C.struct_AVBufferPool)(pool)))
}

// AvBufferPoolBufferGetOpaque queries the original opaque parameter of an allocated buffer in the pool.
func AvBufferPoolBufferGetOpaque(buf *AVBufferRef) unsafe.Pointer {
	return C.av_buffer_pool_buffer_get_opaque((*C.struct_AVBufferRef)(buf))
}
