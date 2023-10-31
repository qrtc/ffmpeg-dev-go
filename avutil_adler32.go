// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/adler32.h>
*/
import "C"

// AvAdler32Update calculates the Adler32 checksum of a buffer.
func AvAdler32Update(adler uint32, buf *uint8, len uint32) uint32 {
	return (uint32)(C.av_adler32_update((C.ulong)(adler), (*C.uint8_t)(buf), (C.uint)(len)))
}
