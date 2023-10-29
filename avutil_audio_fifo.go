// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/audio_fifo.h>
*/
import "C"

type AVAudioFifo C.struct_AVAudioFifo

// AvAudioFifoFree frees an AVAudioFifo.
func AvAudioFifoFree(af *AVAudioFifo) {
	C.av_audio_fifo_free((*C.struct_AVAudioFifo)(af))
}

// AvAudioFifoAlloc allocates an AVAudioFifo.
func AvAudioFifoAlloc(sampleFmt AVSampleFormat, channels, nbSamples int32) *AVAudioFifo {
	return (*AVAudioFifo)(C.av_audio_fifo_alloc((C.enum_AVSampleFormat)(sampleFmt),
		(C.int)(channels), (C.int)(nbSamples)))
}

// AvAudioFifoRealloc reallocate an AVAudioFifo.
func AvAudioFifoRealloc(af *AVAudioFifo, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_realloc((*C.struct_AVAudioFifo)(af), (C.int)(nbSamples)))
}

// AvAudioFifoWrite writes data to an AVAudioFifo.
func AvAudioFifoWrite(af *AVAudioFifo, data CVoidPointerPointer, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_write((*C.struct_AVAudioFifo)(af), VoidPointerPointer(data), (C.int)(nbSamples)))
}

// AvAudioFifoPeek peeks data from an AVAudioFifo.
func AvAudioFifoPeek(af *AVAudioFifo, data CVoidPointerPointer, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_peek((*C.struct_AVAudioFifo)(af), VoidPointerPointer(data), (C.int)(nbSamples)))
}

// AvAudioFifoPeekAt peeks data from an AVAudioFifo.
func AvAudioFifoPeekAt(af *AVAudioFifo, data CVoidPointerPointer, nbSamples, offset int32) int32 {
	return (int32)(C.av_audio_fifo_peek_at((*C.struct_AVAudioFifo)(af), VoidPointerPointer(data),
		(C.int)(nbSamples), (C.int)(offset)))
}

// AvAudioFifoRead reads data from an AVAudioFifo.
func AvAudioFifoRead(af *AVAudioFifo, data CVoidPointerPointer, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_read((*C.struct_AVAudioFifo)(af), VoidPointerPointer(data), (C.int)(nbSamples)))
}

// AvAudioFifoDrain drains data from an AVAudioFifo.
func AvAudioFifoDrain(af *AVAudioFifo, nbSamples int32) int32 {
	return (int32)(C.av_audio_fifo_drain((*C.struct_AVAudioFifo)(af), (C.int)(nbSamples)))
}

// AvAudioFifoReset resets the AVAudioFifo buffer.
func AvAudioFifoReset(af *AVAudioFifo) {
	C.av_audio_fifo_reset((*C.struct_AVAudioFifo)(af))
}

// AvAudioFifoSize gets the current number of samples in the AVAudioFifo available for reading.
func AvAudioFifoSize(af *AVAudioFifo) int32 {
	return (int32)(C.av_audio_fifo_size((*C.struct_AVAudioFifo)(af)))
}

// AvAudioFifoSpace gets the current number of samples in the AVAudioFifo available for writing.
func AvAudioFifoSpace(af *AVAudioFifo) int32 {
	return (int32)(C.av_audio_fifo_space((*C.struct_AVAudioFifo)(af)))
}
