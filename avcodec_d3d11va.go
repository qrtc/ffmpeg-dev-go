//go:build ffmpeg_hw_d3d11va

package ffmpeg

/*
#include <libavcodec/d3d11va.h>
*/
import "C"

const (
	FF_DXVA2_WORKAROUND_SCALING_LIST_ZIGZAG = C.FF_DXVA2_WORKAROUND_SCALING_LIST_ZIGZAG
	FF_DXVA2_WORKAROUND_INTEL_CLEARVIDEO    = C.FF_DXVA2_WORKAROUND_INTEL_CLEARVIDEO
)

// AVD3D11VAContext
type AVD3D11VAContext C.struct_AVD3D11VAContext

// AvD3d11vaAllocContext allocates an AVD3D11VAContext.
func AvD3d11vaAllocContext() *AVD3D11VAContext {
	return (*AVD3D11VAContext)(C.av_d3d11va_alloc_context())
}
