// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build ffmpeg_hw_videotoolbox

package ffmpeg

/*
#include <libavutil/hwcontext_videotoolbox.h>
*/
import "C"

// AvMapVideotoolboxFormatToPixfmt converts a VideoToolbox (actually CoreVideo) format to AVPixelFormat.
func AvMapVideotoolboxFormatToPixfmt(cvFmt uint32) AVPixelFormat {
	return (AVPixelFormat)(C.av_map_videotoolbox_format_to_pixfmt((C.uint32_t)(cvFmt)))
}

// AvMapVideotoolboxFormatFromPixfmt converts an AVPixelFormat to a VideoToolbox (actually CoreVideo) format.
func AvMapVideotoolboxFormatFromPixfmt(pixFmt AVPixelFormat) uint32 {
	return (uint32)(C.av_map_videotoolbox_format_from_pixfmt((C.enum_AVPixelFormat)(pixFmt)))
}

// AvMapVideotoolboxFormatFromPixfmt2 same as av_map_videotoolbox_format_from_pixfmt function, but can map and
// return full range pixel formats via a flag.
func AvMapVideotoolboxFormatFromPixfmt2(pixFmt AVPixelFormat, fullRange bool) uint32 {
	return (uint32)(C.av_map_videotoolbox_format_from_pixfmt2((C.enum_AVPixelFormat)(pixFmt), C.bool(fullRange)))
}

// NONEED: av_map_videotoolbox_chroma_loc_from_av

// NONEED: av_map_videotoolbox_color_matrix_from_av

// NONEED: av_map_videotoolbox_color_primaries_from_av

// NONEED: av_map_videotoolbox_color_trc_from_av

// NONEED: av_vt_pixbuf_set_attachments
