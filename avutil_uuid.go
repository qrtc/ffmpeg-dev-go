// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/uuid.h>
*/
import "C"

const (
	AV_UUID_LEN      = C.AV_UUID_LEN
	AV_UUID_STR_SIZE = 37
)

type AVUUID = []uint8

// AvMakeUUID makes a new AVUUID.
func AvMakeUUID() AVUUID {
	return make([]uint8, AV_UUID_LEN)
}

// AvUuidParse parses a string representation of a UUID formatted according to IETF RFC 4122
// into an AVUUID.
func AvUuidParse(in string, uu AVUUID) int32 {
	inPtr, inFunc := StringCasting(in)
	defer inFunc()
	return (int32)(C.av_uuid_parse((*C.char)(inPtr), (*C.uint8_t)(&uu[0])))
}

// AvUuidUrnParse parses a URN representation of a UUID, as specified at IETF RFC 4122,
// into an AVUUID.
func AvUuidUrnParse(in string, uu AVUUID) int32 {
	inPtr, inFunc := StringCasting(in)
	defer inFunc()
	return (int32)(C.av_uuid_urn_parse((*C.char)(inPtr), (*C.uint8_t)(&uu[0])))
}

// NONEED: av_uuid_parse_range

// AvUuidUnparse serializes a AVUUID into a string representation according to IETF RFC 4122.
func AvUuidUnparse(uu AVUUID) string {
	buf := make([]C.char, AV_UUID_STR_SIZE)
	C.av_uuid_unparse((*C.uint8_t)(&uu[0]), (*C.char)(&buf[0]))
	return C.GoString(&buf[0])
}

// AvUuidEqual compares two UUIDs for equality.
func AvUuidEqual(uu1, uu2 AVUUID) int32 {
	return (int32)(C.av_uuid_equal((*C.uint8_t)(&uu1[0]), (*C.uint8_t)(&uu2[0])))
}

// AvUuidCopy copies the bytes of src into dest.
func AvUuidCopy(dest, src AVUUID) {
	C.av_uuid_copy((*C.uint8_t)(&dest[0]), (*C.uint8_t)(&src[0]))
}

// AvUuidNil sets a UUID to the nil UUID.
func AvUuidNil(uu AVUUID) {
	C.av_uuid_nil((*C.uint8_t)(&uu[0]))
}
