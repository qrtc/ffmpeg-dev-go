package ffmpeg

/*
#include <libavutil/audio_fifo.h>
*/
import "C"
import "unsafe"

type AvAudioFifo C.struct_AVAudioFifo

// AvAudioFifoFree frees an AVAudioFifo.
func AvAudioFifoFree(af *AvAudioFifo) {
	C.av_audio_fifo_free((*C.struct_AVAudioFifo)(af))
}

// AvAudioFifoAlloc allocates an AVAudioFifo.
func AvAudioFifoAlloc(sampleFmt AvSampleFormat, channels, nbSamples int32) *AvAudioFifo {
	return (*AvAudioFifo)(C.av_audio_fifo_alloc((C.enum_AVSampleFormat)(sampleFmt),
		(C.int)(channels), (C.int)(nbSamples)))
}

// AvAudioFifoRealloc reallocate an AVAudioFifo.
func AvAudioFifoRealloc(af *AvAudioFifo, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_realloc((*C.struct_AVAudioFifo)(af), (C.int)(nbSamples)))
}

// AvAudioFifoWrite writes data to an AVAudioFifo.
func AvAudioFifoWrite(af *AvAudioFifo, data *unsafe.Pointer, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_write((*C.struct_AVAudioFifo)(af), data, (C.int)(nbSamples)))
}

// AvAudioFifoPeek peeks data from an AVAudioFifo.
func AvAudioFifoPeek(af *AvAudioFifo, data *unsafe.Pointer, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_peek((*C.struct_AVAudioFifo)(af), data, (C.int)(nbSamples)))
}

// AvAudioFifoPeekAt peeks data from an AVAudioFifo.
func AvAudioFifoPeekAt(af *AvAudioFifo, data *unsafe.Pointer, nbSamples, offset int32) int32 {
	return (int32)(C.av_audio_fifo_peek_at((*C.struct_AVAudioFifo)(af), data,
		(C.int)(nbSamples), (C.int)(offset)))
}

// AvAudioFifoRead reads data from an AVAudioFifo.
func AvAudioFifoRead(af *AvAudioFifo, data *unsafe.Pointer, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_read((*C.struct_AVAudioFifo)(af), data, (C.int)(nbSamples)))
}

// AvAudioFifoDrain drains data from an AVAudioFifo.
func AvAudioFifoDrain(af *AvAudioFifo, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_drain((*C.struct_AVAudioFifo)(af), (C.int)(nbSamples)))
}

// AvAudioFifoReset resets the AVAudioFifo buffer.
func AvAudioFifoReset(af *AvAudioFifo) {
	C.av_audio_fifo_reset((*C.struct_AVAudioFifo)(af))
}

// AvAudioFifoSize gets the current number of samples in the AVAudioFifo available for reading.
func AvAudioFifoSize(af *AvAudioFifo) int32 {
	return (int32)(C.av_audio_fifo_size((*C.struct_AVAudioFifo)(af)))
}

// AvAudioFifoSpace gets the current number of samples in the AVAudioFifo available for writing.
func AvAudioFifoSpace(af *AvAudioFifo) int32 {
	return (int32)(C.av_audio_fifo_space((*C.struct_AVAudioFifo)(af)))
}
