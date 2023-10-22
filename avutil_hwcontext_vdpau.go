//go:build ffmpeg_hw_vdpau

package ffmpeg

/*
#include <libavutil/hwcontext_vdpau.h>
*/
import "C"

// AVVDPAUDeviceContext
type AVVDPAUDeviceContext C.struct_AVVDPAUDeviceContext
