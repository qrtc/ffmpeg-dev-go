package ffmpeg

/*
#include <libavutil/timestamp.h>
*/
import "C"

const (
	AV_TS_MAX_STRING_SIZE = C.AV_TS_MAX_STRING_SIZE
)

// AvTsMakeString makes a timestamp string.
func AvTsMakeString(ts int64) string {
	b := make([]C.char, AV_TS_MAX_STRING_SIZE+1)
	return C.GoString(C.av_ts_make_string((*C.char)(&b[0]), (C.int64_t)(ts)))
}

// AvTs2str
func AvTs2str[T HelperInteger](ts T) string {
	return AvTsMakeString(int64(ts))
}

// AvTsMakeTimeString makes a timestamp string.
func AvTsMakeTimeString(ts int64, tb *AVRational) string {
	b := make([]C.char, AV_TS_MAX_STRING_SIZE+1)
	return C.GoString(C.av_ts_make_time_string((*C.char)(&b[0]),
		(C.int64_t)(ts), (*C.struct_AVRational)(tb)))
}

// AvTs2timestr
func AvTs2timestr[T HelperInteger](ts T, tb *AVRational) string {
	return AvTsMakeTimeString((int64)(ts), tb)
}
