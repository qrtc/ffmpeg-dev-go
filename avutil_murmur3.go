package ffmpeg

/*
#include <libavutil/murmur3.h>
*/
import "C"

type AVMurMur3 C.struct_AVMurMur3

// AvMurmur3Alloc allocates an AVMurMur3 hash context.
func AvMurmur3Alloc() *AVMurMur3 {
	return (*AVMurMur3)(C.av_murmur3_alloc())
}

// AvMurmur3InitSeeded initializes or reinitializes an AVMurMur3 hash context with a seed.
func AvMurmur3InitSeeded(c *AVMurMur3, seed uint64) {
	C.av_murmur3_init_seeded((*C.struct_AVMurMur3)(c), (C.uint64_t)(seed))
}

// AvMurmur3Init initializes or reinitializes an AVMurMur3 hash context.
func AvMurmur3Init(c *AVMurMur3) {
	C.av_murmur3_init((*C.struct_AVMurMur3)(c))
}

// AvMurmur3Update updates hash context with new data.
func AvMurmur3Update(c *AVMurMur3, src *uint8, len int32) {
	C.av_murmur3_update((*C.struct_AVMurMur3)(c), (*C.uint8_t)(src), (C.int)(len))
}

// AvMurmur3Final finishes hashing and output digest value.
func AvMurmur3Final(c *AVMurMur3, dst []uint8) {
	if len(dst) < 16 {
		panic("dst len < 16")
	}
	C.av_murmur3_final((*C.struct_AVMurMur3)(c), (*C.uint8_t)(&dst[0]))
}
