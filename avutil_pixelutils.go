package ffmpeg

/*
#include <libavutil/pixelutils.h>

*/
import "C"
import "unsafe"

type AvPixelutilsSadFn C.av_pixelutils_sad_fn

// AvPixelutilsGetSadFn gets a potentially optimized pointer to a Sum-of-absolute-differences
// function (see the av_pixelutils_sad_fn prototype).
func AvPixelutilsGetSadFn(wBits, hBits, aligned int32, logCtx unsafe.Pointer) AvPixelutilsSadFn {
	return (AvPixelutilsSadFn)(C.av_pixelutils_get_sad_fn((C.int)(wBits), (C.int)(hBits),
		(C.int)(aligned), logCtx))
}
