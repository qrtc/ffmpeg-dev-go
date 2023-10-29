// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/hmac.h>
*/
import "C"

// AVHMACType
type AVHMACType = C.enum_AVHMACType

const (
	AV_HMAC_MD5    = AVHMACType(C.AV_HMAC_MD5)
	AV_HMAC_SHA1   = AVHMACType(C.AV_HMAC_SHA1)
	AV_HMAC_SHA224 = AVHMACType(C.AV_HMAC_SHA224)
	AV_HMAC_SHA256 = AVHMACType(C.AV_HMAC_SHA256)
	AV_HMAC_SHA384 = AVHMACType(C.AV_HMAC_SHA384)
	AV_HMAC_SHA512 = AVHMACType(C.AV_HMAC_SHA512)
)

// AVHMAC
type AVHMAC C.struct_AVHMAC

// AvHmacAlloc allocates an AVHMAC context.
func AvHmacAlloc(_type AVHMACType) *AVHMAC {
	return (*AVHMAC)(C.av_hmac_alloc((C.enum_AVHMACType)(_type)))
}

// AvHmacFree frees an AVHMAC context.
func AvHmacFree(ctx *AVHMAC) {
	C.av_hmac_free((*C.struct_AVHMAC)(ctx))
}

// AvHmacInit initializes an AVHMAC context with an authentication key.
func AvHmacInit(ctx *AVHMAC, key *uint8, keylen uint32) {
	C.av_hmac_init((*C.struct_AVHMAC)(ctx), (*C.uint8_t)(key), (C.uint)(keylen))
}

// AvHmacUpdate hashes data with the HMAC.
func AvHmacUpdate(ctx *AVHMAC, data *uint8, len uint32) {
	C.av_hmac_update((*C.struct_AVHMAC)(ctx), (*C.uint8_t)(data), (C.uint)(len))
}

// AvHmacFinal finishes hashing and output the HMAC digest.
func AvHmacFinal(ctx *AVHMAC, out *uint8, outlen uint32) int32 {
	return (int32)(C.av_hmac_final((*C.struct_AVHMAC)(ctx), (*C.uint8_t)(out), (C.uint)(outlen)))
}

// AvHmacCalc hashes an array of data with a key.
func AvHmacCalc(ctx *AVHMAC, data *uint8, len uint32,
	key *uint8, keylen uint32,
	out *uint8, outlen uint32) int32 {
	return (int32)(C.av_hmac_calc((*C.struct_AVHMAC)(ctx), (*C.uint8_t)(data), (C.uint)(len),
		(*C.uint8_t)(key), (C.uint)(keylen),
		(*C.uint8_t)(out), (C.uint)(outlen)))
}
