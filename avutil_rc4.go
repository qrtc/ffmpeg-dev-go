package ffmpeg

/*
#include <libavutil/rc4.h>
*/
import "C"

type AVRC4 C.struct_AVRC4

// AvRc4Alloc allocates an AVRC4 context.
func AvRc4Alloc() *AVRC4 {
	return (*AVRC4)(C.av_rc4_alloc())
}

// AvRc4Init initializes an AVRC4 context.
func AvRc4Init(d *AVRC4, key *uint8, keyBits int32, decrypt int32) int32 {
	return (int32)(C.av_rc4_init((*C.struct_AVRC4)(d),
		(*C.uint8_t)(key), (C.int)(keyBits), (C.int)(decrypt)))
}

// AvRc4Crypt encrypts / decrypts using the RC4 algorithm.
func AvRc4Crypt(d *AVRC4, dst, src *uint8, count int32, iv *uint8, decrypt int32) {
	C.av_rc4_crypt((*C.struct_AVRC4)(d),
		(*C.uint8_t)(dst), (*C.uint8_t)(src),
		(C.int)(count), (*C.uint8_t)(iv), (C.int)(decrypt))
}
