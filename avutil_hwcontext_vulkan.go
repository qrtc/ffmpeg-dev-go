// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_vulkan

package ffmpeg

/*
#include <libavutil/hwcontext_vulkan.h>
*/
import "C"

// AVVulkanDeviceContext
type AVVulkanDeviceContext C.struct_AVVulkanDeviceContext

// AVVulkanFramesContext
type AVVulkanFramesContext C.struct_AVVulkanFramesContext

// AVVkFrame
type AVVkFrame C.struct_AVVkFrame

// AvVkFrameAlloc allocates a single AVVkFrame and initializes everything as 0.
func AvVkFrameAlloc() *AVVkFrame {
	return (*AVVkFrame)(C.av_vk_frame_alloc())
}

// AvVkfmtFromPixfmt returns the format of each image up to the number of planes for a given sw_format.
// Returns NULL on unsupported formats.
func AvVkfmtFromPixfmt(p AVPixelFormat) *AVVkFrame {
	return (*AVVkFrame)(C.av_vkfmt_from_pixfmt((C.enum_AVPixelFormat)(p)))
}
