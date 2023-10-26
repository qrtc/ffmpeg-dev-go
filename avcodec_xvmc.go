//go:build ffmpeg_hw_xvmc

package ffmpeg

/*
#include <libavcodec/videotoolbox.h>
*/
import "C"

const (
	AV_XVMC_ID = C.AV_XVMC_ID
)

// Deprecated: No use.
//
// XvmcPixFmt
type XvmcPixFmt C.struct_xvmc_pix_fmt
