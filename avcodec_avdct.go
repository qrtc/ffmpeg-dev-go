package ffmpeg

/*
#include <libavcodec/avdct.h>
*/
import "C"

type AvDCT C.struct_AVDCT

// AvCodecDctAlloc allocates a AVDCT context.
func AvCodecDctAlloc() *AvDCT {
	return (*AvDCT)(C.avcodec_dct_alloc())
}

// AvCodecDctInit
func AvCodecDctInit(dct *AvDCT) int32 {
	return (int32)(C.avcodec_dct_init((*C.struct_AVDCT)(dct)))
}

// AvCodecDctGetClass
func AvCodecDctGetClass() *AvClass {
	return (*AvClass)(C.avcodec_dct_get_class())
}
