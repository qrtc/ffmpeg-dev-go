package ffmpeg

/*
#include <libavutil/crc.h>
*/
import "C"

// AVCRC
type AVCRC C.AVCRC

// AVCRCId
type AVCRCId = C.AVCRCId

const (
	AV_CRC_8_ATM      = AVCRCId(C.AV_CRC_8_ATM)
	AV_CRC_16_ANSI    = AVCRCId(C.AV_CRC_16_ANSI)
	AV_CRC_16_CCITT   = AVCRCId(C.AV_CRC_16_CCITT)
	AV_CRC_32_IEEE    = AVCRCId(C.AV_CRC_32_IEEE)
	AV_CRC_32_IEEE_LE = AVCRCId(C.AV_CRC_32_IEEE_LE)
	AV_CRC_16_ANSI_LE = AVCRCId(C.AV_CRC_16_ANSI_LE)
	AV_CRC_24_IEEE    = AVCRCId(C.AV_CRC_24_IEEE)
	AV_CRC_8_EBU      = AVCRCId(C.AV_CRC_8_EBU)
	AV_CRC_MAX        = AVCRCId(C.AV_CRC_MAX)
)

// AvCrcInit initializes a CRC table.
func AvCrcInit(ctx *AVCRC, le, bits int32, poly uint32, ctxSize int32) int32 {
	return (int32)(C.av_crc_init((*C.AVCRC)(ctx),
		(C.int)(le), (C.int)(bits), (C.uint32_t)(poly), (C.int)(ctxSize)))
}

// AvCrcGetTable gets an initialized standard CRC table.
func AvCrcGetTable(crcId AVCRCId) *AVCRC {
	return (*AVCRC)(C.av_crc_get_table((C.AVCRCId)(crcId)))
}

// AvCrc calculates the CRC of a block.
func AvCrc(ctx *AVCRC, crc uint32, buffer *uint8, length uintptr) int32 {
	return (int32)(C.av_crc((*C.AVCRC)(ctx),
		(C.uint32_t)(crc), (*C.uint8_t)(buffer), (C.size_t)(length)))
}
