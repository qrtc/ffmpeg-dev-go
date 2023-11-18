// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

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
	AV_CUDA_USE_CURRENT_CONTEXT = C.AV_CUDA_USE_CURRENT_CONTEXT
)
