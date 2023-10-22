//go:build ffmpeg_hw_cuda

package ffmpeg

/*
#include <libavutil/hwcontext_cuda.h>
*/
import "C"

// AVCUDADeviceContextInternal
type AVCUDADeviceContextInternal C.struct_AVCUDADeviceContextInternal

const (
	AV_CUDA_USE_PRIMARY_CONTEXT = C.AV_CUDA_USE_PRIMARY_CONTEXT
)
