package ffmpeg

/*
#include <libavutil/time.h>
*/
import "C"

// AvGettime gets the current time in microseconds.
func AvGettime() int64 {
	return (int64)(C.av_gettime())
}

// AvGettimeRelative gets the current time in microseconds since some unspecified starting point.
func AvGettimeRelative() int64 {
	return (int64)(C.av_gettime_relative())
}

// AvGettimeRelativeIsMonotonic indicates with a boolean result if the AvGettimeRelative() time source
// is monotonic.
func AvGettimeRelativeIsMonotonic() int32 {
	return (int32)(C.av_gettime_relative_is_monotonic())
}

// AvUsleep sleeps for a period of time.
func AvUsleep(usec uint32) int32 {
	return (int32)(C.av_usleep((C.uint)(usec)))
}
