package ffmpeg

/*
#include <libavutil/md5.h>
*/
import "C"

var (
	AvMd5Size = C.av_md5_size
)

// AVMD5
type AVMD5 C.struct_AVMD5

// AvMd5Alloc allocates an AVMD5 context.
func AvMd5Alloc() *AVMD5 {
	return (*AVMD5)(C.av_md5_alloc())
}

// AvMd5Init initializes MD5 hashing.
func AvMd5Init(ctx *AVMD5) {
	C.av_md5_init((*C.struct_AVMD5)(ctx))
}

// AvMd5Update updates hash value.
func AvMd5Update(ctx *AVMD5, src *uint8, len int32) {
	C.av_md5_update((*C.struct_AVMD5)(ctx), (*C.uint8_t)(src), (C.int)(len))
}

// AvMd5Final finishs hashing and output digest value.
func AvMd5Final(ctx *AVMD5, dst *uint8) {
	C.av_md5_final((*C.struct_AVMD5)(ctx), (*C.uint8_t)(dst))
}

// AvMd5Sum hashes an array of data.
func AvMd5Sum(dst, src *uint8, len int32) {
	C.av_md5_sum((*C.uint8_t)(dst), (*C.uint8_t)(src), (C.int)(len))
}
