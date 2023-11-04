// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/channel_layout.h>
*/
import "C"
import "unsafe"

const (
	AV_CH_FRONT_LEFT            = uint64(C.AV_CH_FRONT_LEFT)
	AV_CH_FRONT_RIGHT           = uint64(C.AV_CH_FRONT_RIGHT)
	AV_CH_FRONT_CENTER          = uint64(C.AV_CH_FRONT_CENTER)
	AV_CH_LOW_FREQUENCY         = uint64(C.AV_CH_LOW_FREQUENCY)
	AV_CH_BACK_LEFT             = uint64(C.AV_CH_BACK_LEFT)
	AV_CH_BACK_RIGHT            = uint64(C.AV_CH_BACK_RIGHT)
	AV_CH_FRONT_LEFT_OF_CENTER  = uint64(C.AV_CH_FRONT_LEFT_OF_CENTER)
	AV_CH_FRONT_RIGHT_OF_CENTER = uint64(C.AV_CH_FRONT_RIGHT_OF_CENTER)
	AV_CH_BACK_CENTER           = uint64(C.AV_CH_BACK_CENTER)
	AV_CH_SIDE_LEFT             = uint64(C.AV_CH_SIDE_LEFT)
	AV_CH_SIDE_RIGHT            = uint64(C.AV_CH_SIDE_RIGHT)
	AV_CH_TOP_CENTER            = uint64(C.AV_CH_TOP_CENTER)
	AV_CH_TOP_FRONT_LEFT        = uint64(C.AV_CH_TOP_FRONT_LEFT)
	AV_CH_TOP_FRONT_CENTER      = uint64(C.AV_CH_TOP_FRONT_CENTER)
	AV_CH_TOP_FRONT_RIGHT       = uint64(C.AV_CH_TOP_FRONT_RIGHT)
	AV_CH_TOP_BACK_LEFT         = uint64(C.AV_CH_TOP_BACK_LEFT)
	AV_CH_TOP_BACK_CENTER       = uint64(C.AV_CH_TOP_BACK_CENTER)
	AV_CH_TOP_BACK_RIGHT        = uint64(C.AV_CH_TOP_BACK_RIGHT)
	AV_CH_STEREO_LEFT           = uint64(C.AV_CH_STEREO_LEFT)
	AV_CH_STEREO_RIGHT          = uint64(C.AV_CH_STEREO_RIGHT)
	AV_CH_WIDE_LEFT             = uint64(C.AV_CH_WIDE_LEFT)
	AV_CH_WIDE_RIGHT            = uint64(C.AV_CH_WIDE_RIGHT)
	AV_CH_SURROUND_DIRECT_LEFT  = uint64(C.AV_CH_SURROUND_DIRECT_LEFT)
	AV_CH_SURROUND_DIRECT_RIGHT = uint64(C.AV_CH_SURROUND_DIRECT_RIGHT)
	AV_CH_LOW_FREQUENCY_2       = uint64(C.AV_CH_LOW_FREQUENCY_2)

	AV_CH_LAYOUT_NATIVE = uint64(C.AV_CH_LAYOUT_NATIVE)
)

const (
	AV_CH_LAYOUT_MONO              = uint64(C.AV_CH_LAYOUT_MONO)
	AV_CH_LAYOUT_STEREO            = uint64(C.AV_CH_LAYOUT_STEREO)
	AV_CH_LAYOUT_2POINT1           = uint64(C.AV_CH_LAYOUT_2POINT1)
	AV_CH_LAYOUT_2_1               = uint64(C.AV_CH_LAYOUT_2_1)
	AV_CH_LAYOUT_SURROUND          = uint64(C.AV_CH_LAYOUT_SURROUND)
	AV_CH_LAYOUT_3POINT1           = uint64(C.AV_CH_LAYOUT_3POINT1)
	AV_CH_LAYOUT_4POINT0           = uint64(C.AV_CH_LAYOUT_4POINT0)
	AV_CH_LAYOUT_4POINT1           = uint64(C.AV_CH_LAYOUT_4POINT1)
	AV_CH_LAYOUT_2_2               = uint64(C.AV_CH_LAYOUT_2_2)
	AV_CH_LAYOUT_QUAD              = uint64(C.AV_CH_LAYOUT_QUAD)
	AV_CH_LAYOUT_5POINT0           = uint64(C.AV_CH_LAYOUT_5POINT0)
	AV_CH_LAYOUT_5POINT1           = uint64(C.AV_CH_LAYOUT_5POINT1)
	AV_CH_LAYOUT_5POINT0_BACK      = uint64(C.AV_CH_LAYOUT_5POINT0_BACK)
	AV_CH_LAYOUT_5POINT1_BACK      = uint64(C.AV_CH_LAYOUT_5POINT1_BACK)
	AV_CH_LAYOUT_6POINT0           = uint64(C.AV_CH_LAYOUT_6POINT0)
	AV_CH_LAYOUT_6POINT0_FRONT     = uint64(C.AV_CH_LAYOUT_6POINT0_FRONT)
	AV_CH_LAYOUT_HEXAGONAL         = uint64(C.AV_CH_LAYOUT_HEXAGONAL)
	AV_CH_LAYOUT_6POINT1           = uint64(C.AV_CH_LAYOUT_6POINT1)
	AV_CH_LAYOUT_6POINT1_BACK      = uint64(C.AV_CH_LAYOUT_6POINT1_BACK)
	AV_CH_LAYOUT_6POINT1_FRONT     = uint64(C.AV_CH_LAYOUT_6POINT1_FRONT)
	AV_CH_LAYOUT_7POINT0           = uint64(C.AV_CH_LAYOUT_7POINT0)
	AV_CH_LAYOUT_7POINT0_FRONT     = uint64(C.AV_CH_LAYOUT_7POINT0_FRONT)
	AV_CH_LAYOUT_7POINT1           = uint64(C.AV_CH_LAYOUT_7POINT1)
	AV_CH_LAYOUT_7POINT1_WIDE      = uint64(C.AV_CH_LAYOUT_7POINT1_WIDE)
	AV_CH_LAYOUT_7POINT1_WIDE_BACK = uint64(C.AV_CH_LAYOUT_7POINT1_WIDE_BACK)
	AV_CH_LAYOUT_OCTAGONAL         = uint64(C.AV_CH_LAYOUT_OCTAGONAL)
	AV_CH_LAYOUT_HEXADECAGONAL     = uint64(C.AV_CH_LAYOUT_HEXADECAGONAL)
	AV_CH_LAYOUT_STEREO_DOWNMIX    = uint64(C.AV_CH_LAYOUT_STEREO_DOWNMIX)
)

// AVMatrixEncoding
type AVMatrixEncoding = C.enum_AVMatrixEncoding

const (
	AV_MATRIX_ENCODING_NONE           = AVMatrixEncoding(C.AV_MATRIX_ENCODING_NONE)
	AV_MATRIX_ENCODING_DOLBY          = AVMatrixEncoding(C.AV_MATRIX_ENCODING_DOLBY)
	AV_MATRIX_ENCODING_DPLII          = AVMatrixEncoding(C.AV_MATRIX_ENCODING_DPLII)
	AV_MATRIX_ENCODING_DPLIIX         = AVMatrixEncoding(C.AV_MATRIX_ENCODING_DPLIIX)
	AV_MATRIX_ENCODING_DPLIIZ         = AVMatrixEncoding(C.AV_MATRIX_ENCODING_DPLIIZ)
	AV_MATRIX_ENCODING_DOLBYEX        = AVMatrixEncoding(C.AV_MATRIX_ENCODING_DOLBYEX)
	AV_MATRIX_ENCODING_DOLBYHEADPHONE = AVMatrixEncoding(C.AV_MATRIX_ENCODING_DOLBYHEADPHONE)
	AV_MATRIX_ENCODING_NB             = AVMatrixEncoding(C.AV_MATRIX_ENCODING_NB)
)

// AvGetChannelLayout returns a channel layout id that matches name, or 0 if no match is found.
func AvGetChannelLayout(name string) uint64 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (uint64)(C.av_get_channel_layout((*C.char)(namePtr)))
}

// AvGetExtendedChannelLayout returns a channel layout and the number of channels based on the specified name.
func AvGetExtendedChannelLayout(name string, channelLayout *uint64, nbChannels *int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_get_extended_channel_layout((*C.char)(namePtr),
		(*C.uint64_t)(channelLayout), (*C.int32_t)(nbChannels)))
}

const AV_CH_LAYOUT_MAX_STRING_SIZE = 256

// AvGetChannelLayoutString returns a description of a channel layout.
func AvGetChannelLayoutString(nbChannels int32, channelLayout uint64) string {
	buf := make([]C.char, AV_CH_LAYOUT_MAX_STRING_SIZE)
	C.av_get_channel_layout_string((*C.char)(&buf[0]), (C.int)(AV_CH_LAYOUT_MAX_STRING_SIZE),
		(C.int)(nbChannels), (C.uint64_t)(channelLayout))
	return C.GoString((*C.char)(&buf[0]))
}

// AvBPrintChannelLayout appends a description of a channel layout to a bprint buffer.
func AvBPrintChannelLayout(bp *AVBPrint, nbChannels int32, channelLayout uint64) {
	C.av_bprint_channel_layout((*C.struct_AVBPrint)(bp),
		(C.int)(nbChannels), (C.uint64_t)(channelLayout))
}

// AvGetChannelLayoutNbChannels returns the number of channels in the channel layout.
func AvGetChannelLayoutNbChannels(channelLayout uint64) int32 {
	return (int32)(C.av_get_channel_layout_nb_channels((C.uint64_t)(channelLayout)))
}

// AvGetDefaultChannelLayout returns default channel layout for a given number of channels.
func AvGetDefaultChannelLayout(nbChannels int32) int64 {
	return (int64)(C.av_get_default_channel_layout((C.int)(nbChannels)))
}

// AvGetChannelLayoutChannelIndex gets the index of a channel in channel_layout.
func AvGetChannelLayoutChannelIndex(channelLayout, channel uint64) int32 {
	return (int32)(C.av_get_channel_layout_channel_index((C.uint64_t)(channelLayout),
		(C.uint64_t)(channel)))
}

// AvChannelLayoutExtractChannel gets the channel with the given index in channel_layout.
func AvChannelLayoutExtractChannel(channelLayout uint64, index int32) uint64 {
	return (uint64)(C.av_channel_layout_extract_channel((C.uint64_t)(channelLayout),
		(C.int)(index)))
}

// AvGetChannelName gets the name of a given channel.
func AvGetChannelName(channel uint64) string {
	return C.GoString(C.av_get_channel_name((C.uint64_t)(channel)))
}

// AvGetChannelDescription gets the value and name of a standard channel layout.
func AvGetChannelDescription(channel uint64) string {
	return C.GoString(C.av_get_channel_description((C.uint64_t)(channel)))
}

// AvGetStandardChannelLayout Get the value and name of a standard channel layout.
func AvGetStandardChannelLayout(index uint32, layout *uint64, name **int8) int32 {
	return (int32)(C.av_get_standard_channel_layout((C.uint)(index),
		(*C.uint64_t)(layout), (**C.char)(unsafe.Pointer(name))))
}
