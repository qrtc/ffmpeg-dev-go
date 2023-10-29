// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/aes_ctr.h>
*/
import "C"

const (
	AES_CTR_KEY_SIZE = C.AES_CTR_KEY_SIZE
	AES_CTR_IV_SIZE  = C.AES_CTR_IV_SIZE
)

// AVAESCTR
type AVAESCTR C.struct_AVAESCTR

// AvAesCtrAlloc allocates an AVAESCTR context.
func AvAesCtrAlloc() *AVAESCTR {
	return (*AVAESCTR)(C.av_aes_ctr_alloc())
}

// AvAesCtrInit initializes an AVAESCTR context.
func AvAesCtrInit(a *AVAESCTR, key *uint8) int32 {
	return (int32)(C.av_aes_ctr_init((*C.struct_AVAESCTR)(a), (*C.uint8_t)(key)))
}

// AvAesCtrFree releases an AVAESCTR context.
func AvAesCtrFree(a *AVAESCTR) {
	C.av_aes_ctr_free((*C.struct_AVAESCTR)(a))
}

// AvAesCtrCrypt processes a buffer using a previously initialized context.
func AvAesCtrCrypt(a *AVAESCTR, dst, src *uint8, size int32) {
	C.av_aes_ctr_crypt((*C.struct_AVAESCTR)(a),
		(*C.uint8_t)(dst), (*C.uint8_t)(src), (C.int)(size))
}

// AvAesCtrGetIv gets the current iv.
func AvAesCtrGetIv(a *AVAESCTR) *uint8 {
	return (*uint8)(C.av_aes_ctr_get_iv((*C.struct_AVAESCTR)(a)))
}

// AvAesCtrSetRandomIv generates a random iv.
func AvAesCtrSetRandomIv(a *AVAESCTR) {
	C.av_aes_ctr_set_random_iv((*C.struct_AVAESCTR)(a))
}

// AvAesCtrSetIv changes the 8-byte iv.
func AvAesCtrSetIv(a *AVAESCTR, iv *uint8) {
	C.av_aes_ctr_set_iv((*C.struct_AVAESCTR)(a), (*C.uint8_t)(iv))
}

// AvAesCtrSetFullIv changes the "full" 16-byte iv, including the counter.
func AvAesCtrSetFullIv(a *AVAESCTR, iv *uint8) {
	C.av_aes_ctr_set_full_iv((*C.struct_AVAESCTR)(a), (*C.uint8_t)(iv))
}

// AvAesCtrIncrementIv increments the top 64 bit of the iv.
func AvAesCtrIncrementIv(a *AVAESCTR) {
	C.av_aes_ctr_increment_iv((*C.struct_AVAESCTR)(a))
}
