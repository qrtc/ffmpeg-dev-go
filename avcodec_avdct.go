package ffmpeg

/*
#include <libavcodec/avdct.h>
*/
import "C"

// AVDCT
type AVDCT C.struct_AVDCT

// AvCodecDctAlloc allocates a AVDCT context.
func AvCodecDctAlloc() *AVDCT {
	return (*AVDCT)(C.avcodec_dct_alloc())
}

// AvCodecDctInit
func AvCodecDctInit(dct *AVDCT) int32 {
	return (int32)(C.avcodec_dct_init((*C.struct_AVDCT)(dct)))
}

// AvCodecDctGetClass
func AvCodecDctGetClass() *AVClass {
	return (*AVClass)(C.avcodec_dct_get_class())
}
