package ffmpeg

/*
#include <stdlib.h>
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type HelperInteger interface {
	HelperSingedInteger | HelperUnsingedInteger
}

type HelperSingedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type HelperUnsingedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

const NIL = "\\'nil'\\"

// StringCasting casts go string to c world char* with free function
// Note: if input is a NIL string will return a nil pointer.
func StringCasting(str string) (allocPtr *C.char, freeFunc func()) {
	if str == NIL {
		return nil, func() {}
	}
	allocPtr = C.CString(str)
	freeFunc = func() { C.free(unsafe.Pointer(allocPtr)) }
	return allocPtr, freeFunc
}

// SliceWithOffset returns a []byte slice from a porinter with offset and size.
func ByteSliceWithOffset[X any, Y, Z HelperInteger](data *X, offset Y, size Z) []byte {
	return unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(uintptr(unsafe.Pointer(data))), offset)), size)
}

// Slice returns a []byte slice from a porinter with size.
func ByteSlice[U any, V HelperInteger](data *U, size V) []byte {
	return ByteSliceWithOffset(data, 0, size)
}

// PointerOffset offset the pointer point.
func PointerOffset[U any, V HelperInteger](ptr *U, offset V) *U {
	if ptr == nil {
		return nil
	}
	return (*U)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
		uintptr(unsafe.Sizeof(*ptr))*(uintptr(offset))))
}

// TruncSlice return a slice from a sign-terminated array.
func TruncSlice[T any](ptr *T, fn func(T) bool) []T {
	if ptr == nil {
		return nil
	}
	for i := 0; ; i++ {
		if fn(*(*T)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))*uintptr(i)))) {
			return unsafe.Slice(ptr, i)
		}
	}
}

// TruncStringSlice returns a string slice from a NULL-terminated *C.char array.
func TruncStringSlice(ptr **C.char) (v []string) {
	if ptr == nil {
		return nil
	}
	for *ptr != nil && **ptr != C.char(0x00) {
		v = append(v, C.GoString(*ptr))
		ptr = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// CVoidPointer represents a (void*) type pointer in the C world.
type CVoidPointer any

// VoidPointer returns a unsafe.Pointer from CVoidPointer.
func VoidPointer(a CVoidPointer) unsafe.Pointer {
	if a == nil {
		return nil
	}
	return unsafe.Pointer(reflect.ValueOf(a).Pointer())
}

// CVoidPointer represents a (void**) type pointer in the C world.
type CVoidPointerPointer any

// VoidPointerPointer returns a *unsafe.Pointer from CVoidPointerPointer.
func VoidPointerPointer(a CVoidPointerPointer) *unsafe.Pointer {
	if a == nil {
		return nil
	}
	return (*unsafe.Pointer)(unsafe.Pointer(reflect.ValueOf(a).Pointer()))
}
