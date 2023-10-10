package ffmpeg

/*
#include <libavcodec/ac3_parser.h>
*/
import "C"

// AvAc3ParseHeader extracts the bitstream ID and the frame size from AC-3 data.
func AvAc3ParseHeader(buf *uint8, size uint, bitstreamID *uint8, frameSize *uint16) int32 {
	return (int32)(C.av_ac3_parse_header((*C.uint8_t)(buf), (C.size_t)(size),
		(*C.uint8_t)(bitstreamID), (*C.uint16_t)(frameSize)))
}
