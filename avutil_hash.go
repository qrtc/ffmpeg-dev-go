// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/hash.h>
*/
import "C"
import "unsafe"

// AVHashContext
type AVHashContext C.struct_AVHashContext

// AvHashAlloc allocates a hash context for the algorithm specified by name.
func AvHashAlloc(ctx **AVHashContext, name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_hash_alloc(
		(**C.struct_AVHashContext)(unsafe.Pointer(ctx)), (*C.char)(namePtr)))
}

// AvHashNames gets the names of available hash algorithms.
func AvHashNames(i int32) string {
	return C.GoString(C.av_hash_names((C.int)(i)))
}

// AvHashGetName gets the name of the algorithm corresponding to the given hash context.
func AvHashGetName(ctx *AVHashContext) string {
	return C.GoString(C.av_hash_get_name((*C.struct_AVHashContext)(ctx)))
}

const (
	AV_HASH_MAX_SIZE = C.AV_HASH_MAX_SIZE
)

// AvHashGetSize gets the size of the resulting hash value in bytes.
func AvHashGetSize(ctx *AVHashContext) int32 {
	return (int32)(C.av_hash_get_size((*C.struct_AVHashContext)(ctx)))
}

// AvHashInit initializes or reset a hash context.
func AvHashInit(ctx *AVHashContext) {
	C.av_hash_init((*C.struct_AVHashContext)(ctx))
}

// AvHashUpdate updates a hash context with additional data.
func AvHashUpdate(ctx *AVHashContext, src *uint8, len int32) {
	C.av_hash_update((*C.struct_AVHashContext)(ctx), (*C.uint8_t)(src), (C.int)(len))
}

// AvHashFinal finalizes a hash context and compute the actual hash value.
func AvHashFinal(ctx *AVHashContext, dst *uint8) {
	C.av_hash_final((*C.struct_AVHashContext)(ctx), (*C.uint8_t)(dst))
}

// AvHashFinalBin finalizes a hash context and store the actual hash value in a buffer.
func AvHashFinalBin(ctx *AVHashContext, dst *uint8, size int32) {
	C.av_hash_final_bin((*C.struct_AVHashContext)(ctx), (*C.uint8_t)(dst), (C.int)(size))
}

// AvHashFinalHex finalizes a hash context and store the hexadecimal representation of the
// actual hash value as a string.
func AvHashFinalHex(ctx *AVHashContext, dst *uint8, size int32) {
	C.av_hash_final_hex((*C.struct_AVHashContext)(ctx), (*C.uint8_t)(dst), (C.int)(size))
}

// AvHashFinalB64 finalizes a hash context and store the Base64 representation of the
// actual hash value as a string.
func AvHashFinalB64(ctx *AVHashContext, dst *uint8, size int32) {
	C.av_hash_final_b64((*C.struct_AVHashContext)(ctx), (*C.uint8_t)(dst), (C.int)(size))
}

// AvHashFreep frees hash context and set hash context pointer to `NULL`.
func AvHashFreep(ctx **AVHashContext) {
	C.av_hash_freep((**C.struct_AVHashContext)(unsafe.Pointer(ctx)))
}
