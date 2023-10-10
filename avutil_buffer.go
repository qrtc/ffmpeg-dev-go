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

// AvBufferRef
type AvBuffer C.struct_AVBuffer

// AvBufferRef
type AvBufferRef C.struct_AVBufferRef

// AvBufferAlloc allocates an AVBuffer of the given size using AvMalloc().
func AvBufferAlloc(size int32) *AvBufferRef {
	return (*AvBufferRef)(C.av_buffer_alloc((C.int)(size)))
}

// AvBufferAllocz same as AvBufferAlloc(), except the returned buffer will be initialized
// to zero.
func AvBufferAllocz(size int32) *AvBufferRef {
	return (*AvBufferRef)(C.av_buffer_allocz((C.int)(size)))
}

const AV_BUFFER_FLAG_READONLY = C.AV_BUFFER_FLAG_READONLY

// typedef void (*av_buffer_free_func)(void *opaque, uint8_t *data)
type AvBufferFreeFunc C.av_buffer_free_func

// AvBufferCreate Create an AvBuffer from an existing array.
func AvBufferCreate(data *uint8, size int32, free AvBufferFreeFunc, opaque unsafe.Pointer, flags int32) *AvBufferRef {
	return (*AvBufferRef)(C.av_buffer_create((*C.uint8_t)(data), (C.int)(size),
		(C.av_buffer_free_func)(free), opaque, (C.int)(flags)))
}

// AvBufferDefaultFree frees buffer data.
func AvBufferDefaultFree(opaque unsafe.Pointer, data *uint8) {
	C.av_buffer_default_free(opaque, (*C.uint8_t)(data))
}

// AvBufferRef creates a new reference to an AVBuffer.
func AvBufferReference(buf *AvBufferRef) *AvBufferRef {
	return (*AvBufferRef)(C.av_buffer_ref((*C.struct_AVBufferRef)(buf)))
}

// AvBufferUnreference frees a given reference and automatically free the buffer if there are no more
// references to it.
func AvBufferUnreference(buf **AvBufferRef) {
	C.av_buffer_unref((**C.struct_AVBufferRef)(unsafe.Pointer(buf)))
}

// AvBufferIsWritable returns 1 if the caller may write to the data referred to by buf (which is
// true if and only if buf is the only reference to the underlying AVBuffer).
// Return 0 otherwise.
func AvBufferIsWritable(buf *AvBufferRef) int32 {
	return (int32)(C.av_buffer_is_writable((*C.struct_AVBufferRef)(buf)))
}

// AvBufferGetOpaque returns the opaque parameter set by AvBufferCreate.
func AvBufferGetOpaque(buf *AvBufferRef) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_buffer_get_opaque((*C.struct_AVBufferRef)(buf)))
}

// AvBufferGetRefCount
func AvBufferGetRefCount(buf *AvBufferRef) int32 {
	return (int32)(C.av_buffer_get_ref_count((*C.struct_AVBufferRef)(buf)))
}

// AvBufferMakeWritable creates a writable reference from a given buffer reference,
// avoiding data copy if possible.
func AvBufferMakeWritable(buf **AvBufferRef) int32 {
	return (int32)(C.av_buffer_make_writable((**C.struct_AVBufferRef)(unsafe.Pointer(buf))))
}

// AvBufferRealloc reallocates a given buffer.
func AvBufferRealloc(buf **AvBufferRef, size int32) int32 {
	return (int32)(C.av_buffer_realloc((**C.struct_AVBufferRef)(unsafe.Pointer(buf)), (C.int)(size)))
}

// AvBufferReplace ensures dst refers to the same data as src.
func AvBufferReplace(dst **AvBufferRef, src *AvBufferRef) int32 {
	return (int32)(C.av_buffer_replace((**C.struct_AVBufferRef)(unsafe.Pointer(dst)),
		(*C.struct_AVBufferRef)(src)))
}

type AvBufferPool C.struct_AVBufferPool

// typedef AVBufferRef* (*av_buffer_pool_alloc_func)(int size)
type AvBufferPoolAllocFunc C.av_buffer_pool_alloc_func

// typedef AVBufferRef* (*av_buffer_pool_alloc2_func)(void* opaque, int size)
type AvBufferPoolAlloc2Func C.av_buffer_pool_alloc2_func

// typedef void (*av_buffer_pool_free_func)(void* opaque)
type AvBufferPoolFreeFunc C.av_buffer_pool_free_func

// AvBufferPoolInit allocates and initializes a buffer pool.
func av_buffer_pool_init(size int32, alloc AvBufferPoolAllocFunc) *AvBufferPool {
	return (*AvBufferPool)(C.av_buffer_pool_init((C.int)(size), (C.av_buffer_pool_alloc_func)(alloc)))
}

// AvBufferPoolInit2 allocates and initialize a buffer pool with a more complex allocator.
func AvBufferPoolInit2(size int32, opaque unsafe.Pointer,
	alloc AvBufferPoolAllocFunc, free AvBufferPoolFreeFunc) *AvBufferPool {
	return (*AvBufferPool)(C.av_buffer_pool_init2((C.int)(size), opaque,
		(C.av_buffer_pool_alloc_func)(alloc),
		(C.av_buffer_pool_free_func)(free)))
}

// AvBufferPoolUninit marks the pool as being available for freeing.
func AvBufferPoolUninit(pool **AvBufferPool) {
	C.av_buffer_pool_uninit((**C.struct_AVBufferPool)(unsafe.Pointer(pool)))
}

// AvBufferPoolGet allocates a new AVBuffer, reusing an old buffer from the pool when available.
func AvBufferPoolGet(pool *AvBufferPool) *AvBufferRef {
	return (*AvBufferRef)(C.av_buffer_pool_get((*C.struct_AVBufferPool)(pool)))
}

// AvBufferPoolBufferGetOpaque queries the original opaque parameter of an allocated buffer in the pool.
func AvBufferPoolBufferGetOpaque(buf *AvBufferRef) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_buffer_pool_buffer_get_opaque((*C.struct_AVBufferRef)(buf)))
}
