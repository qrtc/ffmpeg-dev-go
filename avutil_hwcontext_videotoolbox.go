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
