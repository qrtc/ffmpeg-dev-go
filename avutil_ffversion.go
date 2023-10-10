package ffmpeg

/*
#include <libavutil/ffversion.h>
*/
import "C"

const (
	FFMPEG_VERSION = C.FFMPEG_VERSION
)
