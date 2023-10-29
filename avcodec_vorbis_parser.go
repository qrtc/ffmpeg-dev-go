// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavcodec/vorbis_parser.h>
*/
import "C"
import "unsafe"

// AVVorbisParseContext
type AVVorbisParseContext C.struct_AVVorbisParseContext

// AvVorbisParseInit allocates and initialize the Vorbis parser using headers in the extradata.
func AvVorbisParseInit(extradata *uint8, extradataSize int32) *AVVorbisParseContext {
	return (*AVVorbisParseContext)(C.av_vorbis_parse_init((*C.uint8_t)(extradata),
		(C.int)(extradataSize)))
}

// AvVorbisParseFree frees the parser and everything associated with it.
func AvVorbisParseFree(s **AVVorbisParseContext) {
	C.av_vorbis_parse_free((**C.struct_AVVorbisParseContext)(unsafe.Pointer(s)))
}

const (
	VORBIS_FLAG_HEADER  = C.VORBIS_FLAG_HEADER
	VORBIS_FLAG_COMMENT = C.VORBIS_FLAG_COMMENT
	VORBIS_FLAG_SETUP   = C.VORBIS_FLAG_SETUP
)

// AvVorbisParseFrameFlags gets the duration for a Vorbis packet.
func AvVorbisParseFrameFlags(s *AVVorbisParseContext, buf *uint8, bufSize int32, flags *int32) int32 {
	return (int32)(C.av_vorbis_parse_frame_flags((*C.struct_AVVorbisParseContext)(s),
		(*C.uint8_t)(buf), (C.int)(bufSize), (*C.int)(flags)))
}

// AvVorbisParseFrame gets the duration for a Vorbis packet.
func AvVorbisParseFrame(s *AVVorbisParseContext, buf *uint8, bufSize int32) int32 {
	return (int32)(C.av_vorbis_parse_frame((*C.struct_AVVorbisParseContext)(s),
		(*C.uint8_t)(buf), (C.int)(bufSize)))
}

// AVVorbisParseReset
func AvVorbisParseReset(s *AVVorbisParseContext) {
	C.av_vorbis_parse_reset((*C.struct_AVVorbisParseContext)(s))
}
