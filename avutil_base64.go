// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/base64.h>
*/
import "C"

// AvBase64Decode decodes a base64-encoded string.
func AvBase64Decode(out *uint8, in *int8, outSize int32) int32 {
	return (int32)(C.av_base64_decode((*C.uint8_t)(out), (*C.char)(in), (C.int)(outSize)))
}

// AV_BASE64_DECODE_SIZE
func AV_BASE64_DECODE_SIZE[T Integer](x T) T {
	return x * 3 / 4
}

// AvBase64Encode encodes data to base64 and null-terminate.
func AvBase64Encode(out *int8, outSize int32, in *uint8, inSize int32) *int8 {
	return (*int8)(C.av_base64_encode((*C.char)(out), (C.int)(outSize), (*C.uint8_t)(in), (C.int)(inSize)))
}

// AV_BASE64_SIZE
func AV_BASE64_SIZE[T Integer](x T) T {
	return (x+2)/3*4 + 1
}
