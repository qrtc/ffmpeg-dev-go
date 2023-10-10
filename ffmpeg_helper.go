package ffmpeg

/*
#include <stdlib.h>
*/
import "C"
import (
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

// StringCasting casts go string to c world char* with free function
// Note: if input is a empty string will return a nil pointer.
func StringCasting(str string) (allocPtr *C.char, freeFunc func()) {
	if len(str) == 0 {
		return nil, func() {}
	}
	allocPtr = C.CString(str)
	freeFunc = func() { C.free(unsafe.Pointer(allocPtr)) }
	return allocPtr, freeFunc
}

// SliceWithOffset returns a []byte slice from a porinter with size.
func Slice[T HelperInteger](data *uint8, size T) []byte {
	return unsafe.Slice(data, size)
}

// SliceWithOffset returns a []byte slice from a porinter with offset and size.
func SliceWithOffset[U, V HelperInteger](data *uint8, offset U, size V) []byte {
	return unsafe.Slice((*uint8)(unsafe.Add(unsafe.Pointer(uintptr(unsafe.Pointer(data))), offset)), size)
}
