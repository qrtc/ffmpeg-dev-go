package ffmpeg

/*
#include <libavutil/lfg.h>
*/
import "C"

// AVLFG
type AVLFG C.struct_AVLFG

// AvLfgInit
func AvLfgInit(c *AVLFG, seed uint32) {
	C.av_lfg_init((*C.struct_AVLFG)(c), (C.uint)(seed))
}

// AvLfgInitFromData seeds the state of the ALFG using binary data.
func AvLfgInitFromData(c *AVLFG, data *uint8, length uint) int32 {
	return (int32)(C.av_lfg_init_from_data((*C.struct_AVLFG)(c),
		(*C.uint8_t)(data), (C.uint)(length)))
}

// AvLfgGet gets the next random unsigned 32-bit number using an ALFG.
func AvLfgGet(c *AVLFG) uint32 {
	return (uint32)(C.av_lfg_get((*C.struct_AVLFG)(c)))
}

// AvMlfgGet gets the next random unsigned 32-bit number using a MLFG.
func AvMlfgGet(c *AVLFG) uint32 {
	return (uint32)(C.av_mlfg_get((*C.struct_AVLFG)(c)))
}

// AvBmgGet gets the next two numbers generated by a Box-Muller Gaussian
// generator using the random numbers issued by lfg.
func AvBmgGet(lfg *AVLFG, out []float64) {
	if len(out) < 2 {
		panic("out len < 2")
	}
	C.av_bmg_get((*C.struct_AVLFG)(lfg), (*C.double)(&out[0]))
}
