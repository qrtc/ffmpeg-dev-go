package ffmpeg

/*
#include <libavutil/adler32.h>
*/
import "C"

type AVAdler = C.AVAdler

// AvAdler32Update calculates the Adler32 checksum of a buffer.
func AvAdler32Update(adler AVAdler, buf *uint8, len uint32) AVAdler {
	return (AVAdler)(C.av_adler32_update((C.AVAdler)(adler), (*C.uint8_t)(buf), (C.uint)(len)))
}
