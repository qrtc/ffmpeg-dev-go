package ffmpeg

/*
#include <libavutil/random_seed.h>
*/
import "C"

// AvGetRandomSeed gets a seed to use in conjunction with random functions.
func AvGetRandomSeed() uint32 {
	return (uint32)(C.av_get_random_seed())
}
