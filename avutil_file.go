// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/file.h>
*/
import "C"
import "unsafe"

// AvFileMap reads the file with name filename, and put its content in a newly
// allocated buffer or map it with mmap() when available.
func AvFileMap(filename string, bufptr **uint8, size *uintptr, logOffset int32, logCtx CVoidPointer) int32 {
	filenamePtr, filenameFunc := StringCasting(filename)
	defer filenameFunc()
	return (int32)(C.av_file_map((*C.char)(filenamePtr), (**C.uint8_t)(unsafe.Pointer(bufptr)),
		(*C.size_t)(unsafe.Pointer(size)), (C.int)(logOffset), VoidPointer(logCtx)))
}

// AvFileUnmap unmaps or frees the buffer bufptr created by AvFileMap().
func AvFileUnmap(bufptr *uint8, size uintptr) {
	C.av_file_unmap((*C.uint8_t)(bufptr), (C.size_t)(size))
}
