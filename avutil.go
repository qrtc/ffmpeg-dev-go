package ffmpeg

/*
#include <libavutil/avutil.h>
*/
import "C"
import "unsafe"

// AvutilVersion returns the LIBAVUTIL_VERSION_INT constant.
func AvutilVersion() uint32 {
	return (uint32)(C.avutil_version())
}

// AvVersionInfo returns an informative version string.
func AvVersionInfo() string {
	return C.GoString(C.av_version_info())
}

// AvutilConfiguration returns the libavutil build-time configuration.
func AvutilConfiguration() string {
	return C.GoString(C.avutil_configuration())
}

// AvutilLicense returns the libavutil license.
func AvutilLicense() string {
	return C.GoString(C.avutil_license())
}

// Media Type
type AvMediaType int32

const (
	AVMEDIA_TYPE_UNKNOWN    = AvMediaType(C.AVMEDIA_TYPE_UNKNOWN)
	AVMEDIA_TYPE_VIDEO      = AvMediaType(C.AVMEDIA_TYPE_VIDEO)
	AVMEDIA_TYPE_AUDIO      = AvMediaType(C.AVMEDIA_TYPE_AUDIO)
	AVMEDIA_TYPE_DATA       = AvMediaType(C.AVMEDIA_TYPE_DATA)
	AVMEDIA_TYPE_SUBTITLE   = AvMediaType(C.AVMEDIA_TYPE_SUBTITLE)
	AVMEDIA_TYPE_ATTACHMENT = AvMediaType(C.AVMEDIA_TYPE_ATTACHMENT)
	AVMEDIA_TYPE_NB         = AvMediaType(C.AVMEDIA_TYPE_NB)
)

// AvGetMediaTypeString returns a string describing the MediaType,
// Empty string if media_type is unknown.
func AvGetMediaTypeString(mt AvMediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

const (
	FF_LAMBDA_SHIFT = C.FF_LAMBDA_SHIFT
	FF_LAMBDA_SCALE = C.FF_LAMBDA_SCALE
	FF_QP2LAMBDA    = C.FF_QP2LAMBDA
	FF_LAMBDA_MAX   = C.FF_LAMBDA_MAX

	FF_QUALITY_SCALE = C.FF_QUALITY_SCALE
)

const (
	AV_NOPTS_VALUE = C.AV_NOPTS_VALUE
	AV_TIME_BASE   = C.AV_TIME_BASE
)

var (
	AV_TIME_BASE_Q = AvRational(C.AV_TIME_BASE_Q)
)

// AvPictureType, pixel formats and basic image planes manipulation.
type AvPictureType int32

const (
	AV_PICTURE_TYPE_NONE = AvPictureType(C.AV_PICTURE_TYPE_NONE)
	AV_PICTURE_TYPE_I    = AvPictureType(C.AV_PICTURE_TYPE_I)
	AV_PICTURE_TYPE_P    = AvPictureType(C.AV_PICTURE_TYPE_P)
	AV_PICTURE_TYPE_B    = AvPictureType(C.AV_PICTURE_TYPE_B)
	AV_PICTURE_TYPE_S    = AvPictureType(C.AV_PICTURE_TYPE_S)
	AV_PICTURE_TYPE_SI   = AvPictureType(C.AV_PICTURE_TYPE_SI)
	AV_PICTURE_TYPE_SP   = AvPictureType(C.AV_PICTURE_TYPE_SP)
	AV_PICTURE_TYPE_BI   = AvPictureType(C.AV_PICTURE_TYPE_BI)
)

// AvGetPictureTypeChar returns a single letter to describe the given picture type.
func AvGetPictureTypeChar(pictType AvPictureType) string {
	c := C.av_get_picture_type_char((C.enum_AVPictureType)(pictType))
	return C.GoStringN(&c, 1)
}

// AvXIfNull returns x default pointer in case p is NULL.
func AvXIfNull(p, x unsafe.Pointer) unsafe.Pointer {
	return C.av_x_if_null(p, x)
}

// AvIntListLengthForSize computes the length of an integer list.
func AvIntListLengthForSize(elsize uint32, list unsafe.Pointer, term uint64) uint32 {
	return (uint32)(C.av_int_list_length_for_size((C.uint)(elsize), list, (C.uint64_t)(term)))
}

// TODO. av_int_list_length

// TODO. av_fopen_utf8

// AvGetTimeBaseQ returns the fractional representation of the internal time base.
func AvGetTimeBaseQ() AvRational {
	return (AvRational)(C.av_get_time_base_q())
}

const (
	AV_FOURCC_MAX_STRING_SIZE = C.AV_FOURCC_MAX_STRING_SIZE
)

// AvFourcc2str
func AvFourcc2str(fourcc uint32) string {
	buf := make([]C.char, AV_FOURCC_MAX_STRING_SIZE+1)
	return C.GoString(C.av_fourcc_make_string((*C.char)(&buf[0]), (C.uint32_t)(fourcc)))
}

// AvFourccMakeString fill the provided buffer with a string containing a FourCC
// (four-character code) representation.
func AvFourccMakeString(buf *int8, fourcc uint32) *int8 {
	return (*int8)(C.av_fourcc_make_string((*C.char)(buf), (C.uint32_t)(fourcc)))
}
