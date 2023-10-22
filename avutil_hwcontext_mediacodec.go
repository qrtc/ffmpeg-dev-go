//go:build ffmpeg_hw_mediacodec

package ffmpeg

/*
#include <libavutil/hwcontext_mediacodec.h>
*/
import "C"

// AVMediaCodecDeviceContext
type AVMediaCodecDeviceContext C.struct_AVMediaCodecDeviceContext
