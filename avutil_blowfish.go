// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/blowfish.h>
*/
import "C"
import "unsafe"

const (
	AV_BF_ROUNDS = C.AV_BF_ROUNDS
)

// AVBlowfish
type AVBlowfish C.struct_AVBlowfish

// GetP gets `AVBlowfish.p` value.
func (bf *AVBlowfish) GetP() []uint32 {
	return unsafe.Slice((*uint32)(&bf.p[0]), AV_BF_ROUNDS+2)
}

// SetP sets `AVBlowfish.p` value.
func (bf *AVBlowfish) SetP(v []uint32) {
	for i := 0; i < FFMIN(len(v), AV_BF_ROUNDS+2); i++ {
		bf.p[i] = (C.uint32_t)(v[i])
	}
}

// GetPAddr gets `AVBlowfish.p` address.
func (bf *AVBlowfish) GetPAddr() **uint32 {
	return (**uint32)(unsafe.Pointer(&bf.p))
}

// GetS gets `AVBlowfish.s` value.
func (bf *AVBlowfish) GetS() (v [][]uint32) {
	for i := 0; i < 4; i++ {
		v = append(v, unsafe.Slice((*uint32)(&bf.s[i][0]), 256))
	}
	return v
}

// SetS sets `AVBlowfish.s` value.
func (bf *AVBlowfish) SetS(v [][]uint32) {
	for i := 0; i < FFMIN(len(v), 4); i++ {
		for j := 0; j < FFMIN(len(v[i]), 256); j++ {
			bf.s[i][j] = (C.uint32_t)(v[i][j])
		}
	}
}

// GetSAddr gets `AVBlowfish.s` address.
func (bf *AVBlowfish) GetSAddr() **uint32 {
	return (**uint32)(unsafe.Pointer(&bf.s))
}

// AvBlowfishAlloc allocates an AVBlowfish context.
func AvBlowfishAlloc() *AVBlowfish {
	return (*AVBlowfish)(C.av_blowfish_alloc())
}

// AvBlowfishInit initializes an AVBlowfish context.
func AvBlowfishInit(ctx *AVBlowfish, key *uint8, keyLen int32) {
	C.av_blowfish_init((*C.struct_AVBlowfish)(ctx),
		(*C.uint8_t)(key), (C.int)(keyLen))
}

// AvBlowfishCryptEcb encrypts or decrypts a buffer using a previously initialized context.
func AvBlowfishCryptEcb(ctx *AVBlowfish, xl, xr *uint32, decrypt int32) {
	C.av_blowfish_crypt_ecb((*C.struct_AVBlowfish)(ctx),
		(*C.uint32_t)(xl), (*C.uint32_t)(xr), (C.int)(decrypt))
}

// AvBlowfishCrypt encrypts or decrypts a buffer using a previously initialized context.
func AvBlowfishCrypt(ctx *AVBlowfish, dst, src *uint8,
	count int32, iv *uint8, decrypt int32) {
	C.av_blowfish_crypt((*C.struct_AVBlowfish)(ctx),
		(*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}
