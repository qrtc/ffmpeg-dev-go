package ffmpeg

/*
#include <libavutil/mem.h>
*/
import "C"
import "unsafe"

// AvMalloc allocates a memory block with alignment suitable for all memory accesses
// (including vectors if available on the CPU).
func AvMalloc(size uint) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_malloc((C.size_t)(size)))
}

// AvMallocz allocates a memory block with alignment suitable for all memory accesses
// (including vectors if available on the CPU) and zero all the bytes of the
// block.
func AvMallocz(size uint) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_mallocz((C.size_t)(size)))
}

// AvMallocArray allocates a memory block for an array with AvMalloc().
func AvMallocArray(nmemb, size uint) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_malloc_array((C.size_t)(nmemb), (C.size_t)(size)))
}

// AvMalloczArray allocates a memory block for an array with AvMallocz().
func AvMalloczArray(nmemb, size uint) unsafe.Pointer {
	return C.av_mallocz_array((C.size_t)(nmemb), (C.size_t)(size))
}

// AvCalloc is non-inlined equivalent of AvMalloczArray().
func AvCalloc(nmemb, size uint) unsafe.Pointer {
	return C.av_calloc((C.size_t)(nmemb), (C.size_t)(size))
}

// AvRealloc allocates, reallocates, or frees a block of memory.
func AvRealloc(ptr unsafe.Pointer, size uint) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_realloc(ptr, (C.size_t)(size)))
}

// AvReallocp allocates, reallocates, or frees a block of memory through a pointer to a pointer.
func AvReallocp(ptr unsafe.Pointer, size uint) int32 {
	return (int32)(C.av_reallocp(ptr, (C.size_t)(size)))
}

// AvReallocF allocates, reallocates, or frees a block of memory.
func AvReallocF(ptr unsafe.Pointer, nelem, elsize uint) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_realloc_f(ptr, (C.size_t)(nelem), (C.size_t)(elsize)))
}

// AvReallocpArray allocates, reallocates, or frees an array through a pointer to a pointer.
func AvReallocpArray(ptr unsafe.Pointer, nmemb, size uint) int32 {
	return (int32)(C.av_reallocp_array(ptr, (C.size_t)(nmemb), (C.size_t)(size)))
}

// AvFastRealloc reallocates the given buffer if it is not large enough, otherwise do nothing.
func AvFastRealloc(ptr unsafe.Pointer, size *uint32, minSize uint) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_fast_realloc(ptr, (*C.uint)(size), (C.size_t)(minSize)))
}

// AvFastMalloc allocates a buffer, reusing the given one if large enough.
func AvFastMalloc(ptr unsafe.Pointer, size *uint32, minSize uint) {
	C.av_fast_malloc(ptr, (*C.uint)(size), (C.size_t)(minSize))
}

// AvFastMallocz allocates and clear a buffer, reusing the given one if large enough.
func AvFastMallocz(ptr unsafe.Pointer, size *uint32, minSize uint) {
	C.av_fast_mallocz(ptr, (*C.uint)(size), (C.size_t)(minSize))
}

// AvFree free a memory block which has been allocated with a function of AvMalloc()
// or AvRealloc() family.
func AvFree(ptr unsafe.Pointer) {
	C.av_free(ptr)
}

// AvFreep frees a memory block which has been allocated with a function of AvMalloc()
// or AvRealloc() family, and set the pointer pointing to it to `NULL`.
func AvFreep(ptr unsafe.Pointer) {
	C.av_freep(ptr)
}

// AvStrdup
func AvStrdup(s *int8) *int8 {
	return (*int8)(C.av_strdup((*C.char)(s)))
}

// AvStrndup
func AvStrndup(s *int8, len uint) *int8 {
	return (*int8)(C.av_strndup((*C.char)(s), (C.size_t)(len)))
}

// AvMemdup duplicates a buffer with av_malloc().
func AvMemdup(p unsafe.Pointer, size uint) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_memdup(p, (C.size_t)(size)))
}

// Overlapping memcpy() implementation.
func av_memcpy_backptr(dst *uint8, back, cnt int32) {
	C.av_memcpy_backptr((*C.uint8_t)(dst), (C.int)(back), (C.int)(cnt))
}

// AvDynarrayAdd adds the pointer to an element to a dynamic array.
func AvDynarrayAdd(tabPtr unsafe.Pointer, nbPtr *int32, elem unsafe.Pointer) {
	C.av_dynarray_add(tabPtr, (*C.int)(nbPtr), elem)
}

// AvDynarrayAddNofree adds an element to a dynamic array.
func AvDynarrayAddNofree(tabPtr unsafe.Pointer, nbPtr *int32, elem unsafe.Pointer) int32 {
	return (int32)(C.av_dynarray_add_nofree(tabPtr, (*C.int)(nbPtr), elem))
}

// AvDynarray2Add adds an element of size `elem_size` to a dynamic array.
func AvDynarray2Add(tabPtr *unsafe.Pointer, nbPtr *int32,
	elemSize uint, elemData *uint8) unsafe.Pointer {
	return (unsafe.Pointer)(C.av_dynarray2_add(tabPtr,
		(*C.int)(nbPtr), (C.size_t)(elemSize), (*C.uint8_t)(elemData)))
}

// AvSizeMult multiplies two `size_t` values checking for overflow.
func AvSizeMult(a, b uint, r *uint) int32 {
	return (int32)(C.av_size_mult((C.size_t)(a), (C.size_t)(b), (*C.size_t)(unsafe.Pointer(r))))
}

// AvMaxAlloc sets the maximum size that may be allocated in one block.
func AvMaxAlloc(max uint) {
	C.av_max_alloc((C.size_t)(max))
}
