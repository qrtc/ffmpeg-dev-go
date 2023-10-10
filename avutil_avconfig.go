package ffmpeg

/*
#include <libavutil/avconfig.h>
*/
import "C"

const (
	AV_HAVE_BIGENDIAN      = C.AV_HAVE_BIGENDIAN
	AV_HAVE_FAST_UNALIGNED = C.AV_HAVE_FAST_UNALIGNED
)
