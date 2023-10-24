package ffmpeg

/*
#include <libavutil/des.h>
*/
import "C"
import "unsafe"

// AVDES
type AVDES C.struct_AVDES

// Custom: GetRoundKeys gets `AVDES.round_keys` value.
func (d *AVDES) GetRoundKeys() (v [][]uint64) {
	for i := 0; i < 3; i++ {
		v = append(v, unsafe.Slice((*uint64)(&d.round_keys[i][0]), 16))
	}
	return v
}

// Custom: SetRoundKeys sets `AVDES.round_keys` value.
func (d *AVDES) SetRoundKeys(v [][]uint64) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		for j := 0; j < FFMIN(len(v[i]), 16); j++ {
			d.round_keys[i][j] = (C.uint64_t)(v[i][j])
		}
	}
}

// Custom: GetRoundKeysAddr gets `AVDES.round_keys` address.
func (d *AVDES) GetRoundKeysAddr() **uint64 {
	return (**uint64)(unsafe.Pointer(&d.round_keys))
}

// Custom: GetTripleDes gets `AVDES.triple_des` value.
func (d *AVDES) GetTripleDes() int32 {
	return (int32)(d.triple_des)
}

// Custom: SetTripleDes sets `AVDES.triple_des` value.
func (d *AVDES) SetTripleDes(v int32) {
	d.triple_des = (C.int)(v)
}

// Custom: GetTripleDesAddr gets `AVDES.triple_des` address.
func (d *AVDES) GetTripleDesAddr() *int32 {
	return (*int32)(&d.triple_des)
}

// AvDesAlloc allocates an AVDES context.
func AvDesAlloc() *AVDES {
	return (*AVDES)(C.av_des_alloc())
}

// AvDesInit initializes an AVDES context.
func AvDesInit(d *AVDES, key *uint8, keyBits, decrypt int32) int32 {
	return (int32)(C.av_des_init((*C.struct_AVDES)(d), (*C.uint8_t)(key),
		(C.int)(keyBits), (C.int)(decrypt)))
}

// AvDesCrypt encrypts / decrypts using the DES algorithm.
func AvDesCrypt(d *AVDES, dst, src *uint8, count int32, iv *uint8, decrypt int32) {
	C.av_des_crypt((*C.struct_AVDES)(d), (*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}

// AvDesMac calculates CBC-MAC using the DES algorithm.
func AvDesMac(d *AVDES, dst, src *uint8, count int32) {
	C.av_des_mac((*C.struct_AVDES)(d), (*C.uint8_t)(dst), (*C.uint8_t)(src), (C.int)(count))
}
