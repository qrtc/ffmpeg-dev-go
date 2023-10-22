//go:build ffmpeg_hw_qsv

package ffmpeg

/*
#include <libavutil/hwcontext_qsv.h>
*/
import "C"

// AVQSVDeviceContext
type AVQSVDeviceContext C.struct_AVQSVDeviceContext

// AVQSVFramesContext
type AVQSVFramesContext C.struct_AVQSVFramesContext
