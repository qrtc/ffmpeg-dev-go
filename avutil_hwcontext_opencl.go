//go:build ffmpeg_hw_opencl

package ffmpeg

/*
#include <libavutil/hwcontext_opencl.h>
*/
import "C"

// AVOpenCLFrameDescriptor
type AVOpenCLFrameDescriptor C.struct_AVOpenCLFrameDescriptor

// AVOpenCLDeviceContext
type AVOpenCLDeviceContext C.struct_AVOpenCLDeviceContext

// AVOpenCLFramesContext
type AVOpenCLFramesContext C.struct_AVOpenCLFramesContext
