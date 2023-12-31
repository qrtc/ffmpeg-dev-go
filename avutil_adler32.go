// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/adler32.h>
*/
import "C"

// AVAdler
type AVAdler = C.AVAdler

// AvAdler32Update calculates the Adler32 checksum of a buffer.
func AvAdler32Update(adler AVAdler, buf *uint8, len uintptr) AVAdler {
	return (AVAdler)(C.av_adler32_update((C.AVAdler)(adler), (*C.uint8_t)(buf), (C.size_t)(len)))
}
