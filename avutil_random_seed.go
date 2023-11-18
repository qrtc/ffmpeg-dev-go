// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/random_seed.h>
*/
import "C"

// AvGetRandomSeed gets a seed to use in conjunction with random functions.
func AvGetRandomSeed() uint32 {
	return (uint32)(C.av_get_random_seed())
}

// AvRandomBytes gets a seed to use in conjunction with random functions.
func AvRandomBytes[T Integer](buf *uint8, len T) int32 {
	return (int32)(C.av_random_bytes((*C.uint8_t)(buf), (C.size_t)(len)))
}
