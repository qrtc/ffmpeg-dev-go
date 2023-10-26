//go:build ffmpeg_hw_vaapi

package ffmpeg

/*
#include <libavcodec/vappi.h>
*/
import "C"

// Deprecated: No use.
//
// VaapiContext
type VaapiContext C.struct_vaapi_context
