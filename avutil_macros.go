package ffmpeg

/*
#include <libavutil/macros.h>
*/
import "C"

func FFALIGN[T Integer](x, a T) T {
	return ((x) + (a) - 1) & ^((a) - 1)
}
