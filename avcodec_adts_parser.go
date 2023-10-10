package ffmpeg

/*
#include <libavcodec/adts_parser.h>
*/
import "C"

const (
	AV_AAC_ADTS_HEADER_SIZE = C.AV_AAC_ADTS_HEADER_SIZE
)

// AvAdtsHeaderParse extracts the number of samples and frames from AAC data.
func AvAdtsHeaderParse(buf *uint8, sample *uint32, frame *uint8) int32 {
	return (int32)(C.av_adts_header_parse((*C.uint8_t)(buf),
		(*C.uint32_t)(sample), (*C.uint8_t)(frame)))
}
