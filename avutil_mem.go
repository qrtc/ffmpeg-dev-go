// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/mem.h>
*/
import "C"
import "unsafe"

// AvMalloc allocates a memory block with alignment suitable for all memory accesses
// (including vectors if available on the CPU).
func AvMalloc[T Integer](size T) unsafe.Pointer {
	return C.av_malloc((C.size_t)(size))
}

// AvMallocz allocates a memory block with alignment suitable for all memory accesses
// (including vectors if available on the CPU) and zero all the bytes of the
// block.
func AvMallocz[T Integer](size T) unsafe.Pointer {
	return C.av_mallocz((C.size_t)(size))
}

// AvMallocArray allocates a memory block for an array with AvMalloc().
func AvMallocArray[U, V Integer](nmemb U, size V) unsafe.Pointer {
	return C.av_malloc_array((C.size_t)(nmemb), (C.size_t)(size))
}

// AvMalloczArray allocates a memory block for an array with AvMallocz().
func AvMalloczArray[U, V Integer](nmemb U, size V) unsafe.Pointer {
	return C.av_mallocz_array((C.size_t)(nmemb), (C.size_t)(size))
}

// AvCalloc is non-inlined equivalent of AvMalloczArray().
func AvCalloc[U, V Integer](nmemb U, size V) unsafe.Pointer {
	return C.av_calloc((C.size_t)(nmemb), (C.size_t)(size))
}

// AvRealloc allocates, reallocates, or frees a block of memory.
func AvRealloc[T Integer](ptr CVoidPointer, size T) unsafe.Pointer {
	return C.av_realloc(VoidPointer(ptr), (C.size_t)(size))
}

// AvReallocp allocates, reallocates, or frees a block of memory through a pointer to a pointer.
func AvReallocp[T Integer](ptr CVoidPointer, size T) int32 {
	return (int32)(C.av_reallocp(VoidPointer(ptr), (C.size_t)(size)))
}

// AvReallocF allocates, reallocates, or frees a block of memory.
func AvReallocF[U, V Integer](ptr CVoidPointer, nelem U, elsize V) unsafe.Pointer {
	return C.av_realloc_f(VoidPointer(ptr), (C.size_t)(nelem), (C.size_t)(elsize))
}

// AvReallocpArray allocates, reallocates, or frees an array through a pointer to a pointer.
func AvReallocpArray[U, V Integer](ptr CVoidPointer, nmemb U, size V) int32 {
	return (int32)(C.av_reallocp_array(VoidPointer(ptr), (C.size_t)(nmemb), (C.size_t)(size)))
}

// AvFastRealloc reallocates the given buffer if it is not large enough, otherwise do nothing.
func AvFastRealloc[T Integer](ptr CVoidPointer, size *uint32, minSize T) unsafe.Pointer {
	return C.av_fast_realloc(VoidPointer(ptr), (*C.uint)(size), (C.size_t)(minSize))
}

// AvFastMalloc allocates a buffer, reusing the given one if large enough.
func AvFastMalloc[T Integer](ptr CVoidPointer, size *uint32, minSize T) {
	C.av_fast_malloc(VoidPointer(ptr), (*C.uint)(size), (C.size_t)(minSize))
}

// AvFastMallocz allocates and clear a buffer, reusing the given one if large enough.
func AvFastMallocz[T Integer](ptr CVoidPointer, size *uint32, minSize T) {
	C.av_fast_mallocz(VoidPointer(ptr), (*C.uint)(size), (C.size_t)(minSize))
}

// AvFree free a memory block which has been allocated with a function of AvMalloc()
// or AvRealloc() family.
func AvFree(ptr CVoidPointer) {
	C.av_free(VoidPointer(ptr))
}

// AvFreep frees a memory block which has been allocated with a function of AvMalloc()
// or AvRealloc() family, and set the pointer pointing to it to `NULL`.
func AvFreep(ptr CVoidPointer) {
	C.av_freep(VoidPointer(ptr))
}

// AvStrdup
func AvStrdup(s *int8) *int8 {
	return (*int8)(C.av_strdup((*C.char)(s)))
}

// AvStrndup
func AvStrndup[T Integer](s *int8, len T) *int8 {
	return (*int8)(C.av_strndup((*C.char)(s), (C.size_t)(len)))
}

// AvMemdup duplicates a buffer with av_malloc().
func AvMemdup[T Integer](p CVoidPointer, size T) unsafe.Pointer {
	return C.av_memdup(VoidPointer(p), (C.size_t)(size))
}

// AvMemcpyBackptr overlaps memcpy() implementation.
func AvMemcpyBackptr(dst *uint8, back, cnt int32) {
	C.av_memcpy_backptr((*C.uint8_t)(dst), (C.int)(back), (C.int)(cnt))
}

// AvDynarrayAdd adds the pointer to an element to a dynamic array.
func AvDynarrayAdd(tabPtr CVoidPointer, nbPtr *int32, elem CVoidPointer) {
	C.av_dynarray_add(VoidPointer(tabPtr), (*C.int)(nbPtr), VoidPointer(elem))
}

// AvDynarrayAddNofree adds an element to a dynamic array.
func AvDynarrayAddNofree(tabPtr CVoidPointer, nbPtr *int32, elem CVoidPointer) int32 {
	return (int32)(C.av_dynarray_add_nofree(VoidPointer(tabPtr), (*C.int)(nbPtr), VoidPointer(elem)))
}

// AvDynarray2Add adds an element of size `elem_size` to a dynamic array.
func AvDynarray2Add[T Integer](tabPtr CVoidPointerPointer, nbPtr *int32,
	elemSize T, elemData *uint8) unsafe.Pointer {
	return C.av_dynarray2_add(VoidPointerPointer(tabPtr),
		(*C.int)(nbPtr), (C.size_t)(elemSize), (*C.uint8_t)(elemData))
}

// AvSizeMult multiplies two `size_t` values checking for overflow.
func AvSizeMult[T Integer](a, b T, r *uintptr) int32 {
	return (int32)(C.av_size_mult((C.size_t)(a), (C.size_t)(b), (*C.size_t)(unsafe.Pointer(r))))
}

// AvMaxAlloc sets the maximum size that may be allocated in one block.
func AvMaxAlloc[T Integer](max T) {
	C.av_max_alloc((C.size_t)(max))
}
