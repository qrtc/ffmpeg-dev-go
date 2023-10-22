//go:build ffmpeg_hw_qsv

package ffmpeg

/*
#include <libavcodec/qsv.h>
*/
import "C"

// AVQSVContext
type AVQSVContext = C.struct_AVQSVContext

// AvQsvAllocContext allocates a new context.
// It must be freed by the caller with AvFree().
func AvQsvAllocContext() *AVQSVContext {
	return (*AVQSVContext)(C.av_qsv_alloc_context())
}
