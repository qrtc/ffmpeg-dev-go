//go:build ffmpeg_hw_d3d11va

package ffmpeg

/*
#include <libavutil/hwcontext_d3d11va.h>
*/
import "C"

// AVD3D11VADeviceContext
type AVD3D11VADeviceContext C.struct_AVD3D11VADeviceContext

// AVD3D11FrameDescriptor
type AVD3D11FrameDescriptor C.struct_AVD3D11FrameDescriptor

// AVD3D11VAFramesContext
type AVD3D11VAFramesContext C.struct_AVD3D11VAFramesContext
