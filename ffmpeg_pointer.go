package ffmpeg

/*
#include <stdlib.h>
*/
import "C"
import (
	"sync"
	"unsafe"
)

var (
	pointerSyncMap sync.Map
)

func PointerStore(v interface{}) unsafe.Pointer {
	if v == nil {
		return nil
	}
	ptr := C.malloc(C.size_t(1))
	if ptr == nil {
		panic("allocate memory failed")
	}
	pointerSyncMap.Store(ptr, v)
	return ptr
}

func PointerLoad(ptr unsafe.Pointer) (v interface{}) {
	if ptr == nil {
		return nil
	}
	v, _ = pointerSyncMap.Load(ptr)
	return v
}

func PointerDelete(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}
	pointerSyncMap.Delete(ptr)
	C.free(ptr)
}
