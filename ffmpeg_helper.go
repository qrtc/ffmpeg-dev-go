// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <stdlib.h>
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type Integer interface {
	SingedInteger | UnsingedInteger
}

type SingedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsingedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

const NIL = "\\'<nil>'\\"

// StringCasting casts go string to c world char* with free function.
// Note: if input is a NIL string will return a nil pointer.
func StringCasting(s string) (allocPtr *C.char, freeFunc func()) {
	if s == NIL {
		return nil, func() {}
	}
	allocPtr = C.CString(s)
	freeFunc = func() { C.free(unsafe.Pointer(allocPtr)) }
	return allocPtr, freeFunc
}

// StringSliceCasting casts go string slice to c world char* slice with free function.
func StringSliceCasting(ss []string) (allocPtrs []*C.char, freeFunc func()) {
	var freeFns []func()
	for _, s := range ss {
		ptr, fn := StringCasting(s)
		allocPtrs = append(allocPtrs, ptr)
		freeFns = append(freeFns, fn)
	}
	return allocPtrs, func() {
		for _, fn := range freeFns {
			fn()
		}
	}
}

// SliceSlice returns a slice of slice from a pointer to pointer.
func SliceSlice[T any, X, Y Integer](data **T, x X, y Y) (v [][]T) {
	for i := 0; i < int(x); i++ {
		v = append(v, unsafe.Slice(*PointerOffset(data, i), y))
	}
	return v
}

// SliceTrunc return a slice from a sign-terminated array.
func SliceTrunc[T any](ptr *T, truncFunc func(T) bool) []T {
	if ptr == nil {
		return nil
	}
	for i := 0; ; i++ {
		if truncFunc(*(*T)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))*uintptr(i)))) {
			return unsafe.Slice(ptr, i)
		}
	}
}

// SliceTruncString returns a string slice from a NULL-terminated *C.char array.
func SliceTruncString(ptr **C.char) (v []string) {
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

// PointerOffset offset the pointer point.
func PointerOffset[U any, V Integer](ptr *U, offset V) *U {
	if ptr == nil {
		return nil
	}
	return (*U)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
		uintptr(unsafe.Sizeof(*ptr))*(uintptr(offset))))
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

// CondExpr is Conditional Operator like Ternary Operator in the C world.
func CondExpr[T any](cond bool, x, y T) T {
	if cond {
		return x
	}
	return y
}

func PlusPlus[T Integer](x *T) T {
	defer func() { *x++ }()
	return *x
}
