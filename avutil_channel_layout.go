// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/channel_layout.h>

uint64_t get_av_channel_layout_u_mask(AVChannelLayout *cl) {
	return cl->u.mask;
}

void set_av_channel_layout_u_mask(AVChannelLayout *cl, uint64_t v) {
	cl->u.mask = v;
}

uint64_t* get_av_channel_layout_u_mask_addr(AVChannelLayout *cl) {
	return &cl->u.mask;
}

AVChannelCustom* get_av_channel_layout_u_map(AVChannelLayout *cl) {
	return cl->u.map;
}

void set_av_channel_layout_u_map(AVChannelLayout *cl, AVChannelCustom* v) {
	cl->u.map = v;
}

AVChannelCustom** get_av_channel_layout_u_map_addr(AVChannelLayout *cl) {
	return &cl->u.map;
}

const struct AVChannelLayout av_channel_layout_mono                  = AV_CHANNEL_LAYOUT_MONO;
const struct AVChannelLayout av_channel_layout_stereo                = AV_CHANNEL_LAYOUT_STEREO;
const struct AVChannelLayout av_channel_layout_2point1               = AV_CHANNEL_LAYOUT_2POINT1;
const struct AVChannelLayout av_channel_layout_2_1                   = AV_CHANNEL_LAYOUT_2_1;
const struct AVChannelLayout av_channel_layout_surround              = AV_CHANNEL_LAYOUT_SURROUND;
const struct AVChannelLayout av_channel_layout_3point1               = AV_CHANNEL_LAYOUT_3POINT1;
const struct AVChannelLayout av_channel_layout_4point0               = AV_CHANNEL_LAYOUT_4POINT0;
const struct AVChannelLayout av_channel_layout_4point1               = AV_CHANNEL_LAYOUT_4POINT1;
const struct AVChannelLayout av_channel_layout_2_2                   = AV_CHANNEL_LAYOUT_2_2;
const struct AVChannelLayout av_channel_layout_quad                  = AV_CHANNEL_LAYOUT_QUAD;
const struct AVChannelLayout av_channel_layout_5point0               = AV_CHANNEL_LAYOUT_5POINT0;
const struct AVChannelLayout av_channel_layout_5point1               = AV_CHANNEL_LAYOUT_5POINT1;
const struct AVChannelLayout av_channel_layout_5point0_back          = AV_CHANNEL_LAYOUT_5POINT0_BACK;
const struct AVChannelLayout av_channel_layout_5point1_back          = AV_CHANNEL_LAYOUT_5POINT1_BACK;
const struct AVChannelLayout av_channel_layout_6point0               = AV_CHANNEL_LAYOUT_6POINT0;
const struct AVChannelLayout av_channel_layout_6point0_front         = AV_CHANNEL_LAYOUT_6POINT0_FRONT;
const struct AVChannelLayout av_channel_layout_3point1point2         = AV_CHANNEL_LAYOUT_3POINT1POINT2;
const struct AVChannelLayout av_channel_layout_hexagonal             = AV_CHANNEL_LAYOUT_HEXAGONAL;
const struct AVChannelLayout av_channel_layout_6point1               = AV_CHANNEL_LAYOUT_6POINT1;
const struct AVChannelLayout av_channel_layout_6point1_back          = AV_CHANNEL_LAYOUT_6POINT1_BACK;
const struct AVChannelLayout av_channel_layout_6point1_front         = AV_CHANNEL_LAYOUT_6POINT1_FRONT;
const struct AVChannelLayout av_channel_layout_7point0               = AV_CHANNEL_LAYOUT_7POINT0;
const struct AVChannelLayout av_channel_layout_7point0_front         = AV_CHANNEL_LAYOUT_7POINT0_FRONT;
const struct AVChannelLayout av_channel_layout_7point1               = AV_CHANNEL_LAYOUT_7POINT1;
const struct AVChannelLayout av_channel_layout_7point1_wide          = AV_CHANNEL_LAYOUT_7POINT1_WIDE;
const struct AVChannelLayout av_channel_layout_7point1_wide_back     = AV_CHANNEL_LAYOUT_7POINT1_WIDE_BACK;
const struct AVChannelLayout av_channel_layout_5point1point2_back    = AV_CHANNEL_LAYOUT_5POINT1POINT2_BACK;
const struct AVChannelLayout av_channel_layout_octagonal             = AV_CHANNEL_LAYOUT_OCTAGONAL;
const struct AVChannelLayout av_channel_layout_cube                  = AV_CHANNEL_LAYOUT_CUBE;
const struct AVChannelLayout av_channel_layout_5point1point4_back    = AV_CHANNEL_LAYOUT_5POINT1POINT4_BACK;
const struct AVChannelLayout av_channel_layout_7point1point2         = AV_CHANNEL_LAYOUT_7POINT1POINT2;
const struct AVChannelLayout av_channel_layout_7point1point4_back    = AV_CHANNEL_LAYOUT_7POINT1POINT4_BACK;
const struct AVChannelLayout av_channel_layout_7point2point3         = AV_CHANNEL_LAYOUT_7POINT2POINT3;
const struct AVChannelLayout av_channel_layout_9point1point4_back    = AV_CHANNEL_LAYOUT_9POINT1POINT4_BACK;
const struct AVChannelLayout av_channel_layout_hexadecagonal         = AV_CHANNEL_LAYOUT_HEXADECAGONAL;
const struct AVChannelLayout av_channel_layout_stereo_downmix        = AV_CHANNEL_LAYOUT_STEREO_DOWNMIX;
const struct AVChannelLayout av_channel_layout_22point2              = AV_CHANNEL_LAYOUT_22POINT2;
const struct AVChannelLayout av_channel_layout_ambisonic_first_order = AV_CHANNEL_LAYOUT_AMBISONIC_FIRST_ORDER;

const struct AVChannelLayout av_channel_layout_7point1_top_back      = AV_CHANNEL_LAYOUT_7POINT1_TOP_BACK;

*/
import "C"
import "unsafe"

const (
	AV_CH_LAYOUT_MAX_STRING_SIZE = 256
)

// AVChannel
type AVChannel = C.enum_AVChannel

const (
	AV_CHAN_NONE                  = AVChannel(C.AV_CHAN_NONE)
	AV_CHAN_FRONT_LEFT            = AVChannel(C.AV_CHAN_FRONT_LEFT)
	AV_CHAN_FRONT_RIGHT           = AVChannel(C.AV_CHAN_FRONT_RIGHT)
	AV_CHAN_FRONT_CENTER          = AVChannel(C.AV_CHAN_FRONT_CENTER)
	AV_CHAN_LOW_FREQUENCY         = AVChannel(C.AV_CHAN_LOW_FREQUENCY)
	AV_CHAN_BACK_LEFT             = AVChannel(C.AV_CHAN_BACK_LEFT)
	AV_CHAN_BACK_RIGHT            = AVChannel(C.AV_CHAN_BACK_RIGHT)
	AV_CHAN_FRONT_LEFT_OF_CENTER  = AVChannel(C.AV_CHAN_FRONT_LEFT_OF_CENTER)
	AV_CHAN_FRONT_RIGHT_OF_CENTER = AVChannel(C.AV_CHAN_FRONT_RIGHT_OF_CENTER)
	AV_CHAN_BACK_CENTER           = AVChannel(C.AV_CHAN_BACK_CENTER)
	AV_CHAN_SIDE_LEFT             = AVChannel(C.AV_CHAN_SIDE_LEFT)
	AV_CHAN_SIDE_RIGHT            = AVChannel(C.AV_CHAN_SIDE_RIGHT)
	AV_CHAN_TOP_CENTER            = AVChannel(C.AV_CHAN_TOP_CENTER)
	AV_CHAN_TOP_FRONT_LEFT        = AVChannel(C.AV_CHAN_TOP_FRONT_LEFT)
	AV_CHAN_TOP_FRONT_CENTER      = AVChannel(C.AV_CHAN_TOP_FRONT_CENTER)
	AV_CHAN_TOP_FRONT_RIGHT       = AVChannel(C.AV_CHAN_TOP_FRONT_RIGHT)
	AV_CHAN_TOP_BACK_LEFT         = AVChannel(C.AV_CHAN_TOP_BACK_LEFT)
	AV_CHAN_TOP_BACK_CENTER       = AVChannel(C.AV_CHAN_TOP_BACK_CENTER)
	AV_CHAN_TOP_BACK_RIGHT        = AVChannel(C.AV_CHAN_TOP_BACK_RIGHT)
	AV_CHAN_STEREO_LEFT           = AVChannel(C.AV_CHAN_STEREO_LEFT)
	AV_CHAN_STEREO_RIGHT          = AVChannel(C.AV_CHAN_STEREO_RIGHT)
	AV_CHAN_WIDE_LEFT             = AVChannel(C.AV_CHAN_WIDE_LEFT)
	AV_CHAN_WIDE_RIGHT            = AVChannel(C.AV_CHAN_WIDE_RIGHT)
	AV_CHAN_SURROUND_DIRECT_LEFT  = AVChannel(C.AV_CHAN_SURROUND_DIRECT_LEFT)
	AV_CHAN_SURROUND_DIRECT_RIGHT = AVChannel(C.AV_CHAN_SURROUND_DIRECT_RIGHT)
	AV_CHAN_LOW_FREQUENCY_2       = AVChannel(C.AV_CHAN_LOW_FREQUENCY_2)
	AV_CHAN_TOP_SIDE_LEFT         = AVChannel(C.AV_CHAN_TOP_SIDE_LEFT)
	AV_CHAN_TOP_SIDE_RIGHT        = AVChannel(C.AV_CHAN_TOP_SIDE_RIGHT)
	AV_CHAN_BOTTOM_FRONT_CENTER   = AVChannel(C.AV_CHAN_BOTTOM_FRONT_CENTER)
	AV_CHAN_BOTTOM_FRONT_LEFT     = AVChannel(C.AV_CHAN_BOTTOM_FRONT_LEFT)
	AV_CHAN_BOTTOM_FRONT_RIGHT    = AVChannel(C.AV_CHAN_BOTTOM_FRONT_RIGHT)
	AV_CHAN_UNUSED                = AVChannel(C.AV_CHAN_UNUSED)
	AV_CHAN_UNKNOWN               = AVChannel(C.AV_CHAN_UNKNOWN)
	AV_CHAN_AMBISONIC_BASE        = AVChannel(C.AV_CHAN_AMBISONIC_BASE)
	AV_CHAN_AMBISONIC_END         = AVChannel(C.AV_CHAN_AMBISONIC_END)
)

// AVChannelOrder
type AVChannelOrder = C.enum_AVChannelOrder

const (
	AV_CHANNEL_ORDER_UNSPEC    = AVChannelOrder(C.AV_CHANNEL_ORDER_UNSPEC)
	AV_CHANNEL_ORDER_NATIVE    = AVChannelOrder(C.AV_CHANNEL_ORDER_NATIVE)
	AV_CHANNEL_ORDER_CUSTOM    = AVChannelOrder(C.AV_CHANNEL_ORDER_CUSTOM)
	AV_CHANNEL_ORDER_AMBISONIC = AVChannelOrder(C.AV_CHANNEL_ORDER_AMBISONIC)
)

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
	AV_CH_TOP_SIDE_LEFT         = uint64(C.AV_CH_TOP_SIDE_LEFT)
	AV_CH_TOP_SIDE_RIGHT        = uint64(C.AV_CH_TOP_SIDE_RIGHT)
	AV_CH_BOTTOM_FRONT_CENTER   = uint64(C.AV_CH_BOTTOM_FRONT_CENTER)
	AV_CH_BOTTOM_FRONT_LEFT     = uint64(C.AV_CH_BOTTOM_FRONT_LEFT)
	AV_CH_BOTTOM_FRONT_RIGHT    = uint64(C.AV_CH_BOTTOM_FRONT_RIGHT)
)

const (
	AV_CH_LAYOUT_MONO               = uint64(C.AV_CH_LAYOUT_MONO)
	AV_CH_LAYOUT_STEREO             = uint64(C.AV_CH_LAYOUT_STEREO)
	AV_CH_LAYOUT_2POINT1            = uint64(C.AV_CH_LAYOUT_2POINT1)
	AV_CH_LAYOUT_2_1                = uint64(C.AV_CH_LAYOUT_2_1)
	AV_CH_LAYOUT_SURROUND           = uint64(C.AV_CH_LAYOUT_SURROUND)
	AV_CH_LAYOUT_3POINT1            = uint64(C.AV_CH_LAYOUT_3POINT1)
	AV_CH_LAYOUT_4POINT0            = uint64(C.AV_CH_LAYOUT_4POINT0)
	AV_CH_LAYOUT_4POINT1            = uint64(C.AV_CH_LAYOUT_4POINT1)
	AV_CH_LAYOUT_2_2                = uint64(C.AV_CH_LAYOUT_2_2)
	AV_CH_LAYOUT_QUAD               = uint64(C.AV_CH_LAYOUT_QUAD)
	AV_CH_LAYOUT_5POINT0            = uint64(C.AV_CH_LAYOUT_5POINT0)
	AV_CH_LAYOUT_5POINT1            = uint64(C.AV_CH_LAYOUT_5POINT1)
	AV_CH_LAYOUT_5POINT0_BACK       = uint64(C.AV_CH_LAYOUT_5POINT0_BACK)
	AV_CH_LAYOUT_5POINT1_BACK       = uint64(C.AV_CH_LAYOUT_5POINT1_BACK)
	AV_CH_LAYOUT_6POINT0            = uint64(C.AV_CH_LAYOUT_6POINT0)
	AV_CH_LAYOUT_6POINT0_FRONT      = uint64(C.AV_CH_LAYOUT_6POINT0_FRONT)
	AV_CH_LAYOUT_HEXAGONAL          = uint64(C.AV_CH_LAYOUT_HEXAGONAL)
	AV_CH_LAYOUT_3POINT1POINT2      = uint64(C.AV_CH_LAYOUT_3POINT1POINT2)
	AV_CH_LAYOUT_6POINT1            = uint64(C.AV_CH_LAYOUT_6POINT1)
	AV_CH_LAYOUT_6POINT1_BACK       = uint64(C.AV_CH_LAYOUT_6POINT1_BACK)
	AV_CH_LAYOUT_6POINT1_FRONT      = uint64(C.AV_CH_LAYOUT_6POINT1_FRONT)
	AV_CH_LAYOUT_7POINT0            = uint64(C.AV_CH_LAYOUT_7POINT0)
	AV_CH_LAYOUT_7POINT0_FRONT      = uint64(C.AV_CH_LAYOUT_7POINT0_FRONT)
	AV_CH_LAYOUT_7POINT1            = uint64(C.AV_CH_LAYOUT_7POINT1)
	AV_CH_LAYOUT_7POINT1_WIDE       = uint64(C.AV_CH_LAYOUT_7POINT1_WIDE)
	AV_CH_LAYOUT_7POINT1_WIDE_BACK  = uint64(C.AV_CH_LAYOUT_7POINT1_WIDE_BACK)
	AV_CH_LAYOUT_5POINT1POINT2_BACK = uint64(C.AV_CH_LAYOUT_5POINT1POINT2_BACK)
	AV_CH_LAYOUT_OCTAGONAL          = uint64(C.AV_CH_LAYOUT_OCTAGONAL)
	AV_CH_LAYOUT_CUBE               = uint64(C.AV_CH_LAYOUT_CUBE)
	AV_CH_LAYOUT_5POINT1POINT4_BACK = uint64(C.AV_CH_LAYOUT_5POINT1POINT4_BACK)
	AV_CH_LAYOUT_7POINT1POINT2      = uint64(C.AV_CH_LAYOUT_7POINT1POINT2)
	AV_CH_LAYOUT_7POINT1POINT4_BACK = uint64(C.AV_CH_LAYOUT_7POINT1POINT4_BACK)
	AV_CH_LAYOUT_7POINT2POINT3      = uint64(C.AV_CH_LAYOUT_7POINT2POINT3)
	AV_CH_LAYOUT_9POINT1POINT4_BACK = uint64(C.AV_CH_LAYOUT_9POINT1POINT4_BACK)
	AV_CH_LAYOUT_HEXADECAGONAL      = uint64(C.AV_CH_LAYOUT_HEXADECAGONAL)
	AV_CH_LAYOUT_STEREO_DOWNMIX     = uint64(C.AV_CH_LAYOUT_STEREO_DOWNMIX)
	AV_CH_LAYOUT_22POINT2           = uint64(C.AV_CH_LAYOUT_22POINT2)

	AV_CH_LAYOUT_7POINT1_TOP_BACK = uint64(C.AV_CH_LAYOUT_7POINT1_TOP_BACK)
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

// AVChannelCustom
type AVChannelCustom C.struct_AVChannelCustom

// GetId gets `AVChannelCustom.id` value.
func (cc *AVChannelCustom) GetId() AVChannel {
	return (AVChannel)(cc.id)
}

// SetId sets `AVChannelCustom.id` value.
func (cc *AVChannelCustom) SetId(v AVChannel) {
	cc.id = (C.enum_AVChannel)(v)
}

// GetIdAddr gets `AVChannelCustom.id` address.
func (cc *AVChannelCustom) GetIdAddr() *AVChannel {
	return (*AVChannel)(&cc.id)
}

// GetName gets `AVChannelCustom.name` value.
func (cc *AVChannelCustom) GetName() string {
	return C.GoString(&cc.name[0])
}

// GetOpaque gets `AVChannelCustom.opaque` value.
func (cc *AVChannelCustom) GetOpaque() unsafe.Pointer {
	return cc.opaque
}

// SetOpaque sets `AVChannelCustom.opaque` value.
func (cc *AVChannelCustom) SetOpaque(v CVoidPointer) {
	cc.opaque = VoidPointer(v)
}

// GetOpaqueAddr gets `AVChannelCustom.opaque` address.
func (cc *AVChannelCustom) GetOpaqueAddr() *unsafe.Pointer {
	return &cc.opaque
}

// AVChannelLayout
type AVChannelLayout C.struct_AVChannelLayout

// GetOrder gets `AVChannelLayout.order` value.
func (cl *AVChannelLayout) GetOrder() AVChannelOrder {
	return (AVChannelOrder)(cl.order)
}

// SetOrder sets `AVChannelLayout.order` value.
func (cl *AVChannelLayout) SetOrder(v AVChannelOrder) {
	cl.order = (C.enum_AVChannelOrder)(v)
}

// GetOrderAddr gets `AVChannelLayout.order` address.
func (cl *AVChannelLayout) GetOrderAddr() *AVChannelOrder {
	return (*AVChannelOrder)(&cl.order)
}

// GetNbChannels gets `AVChannelLayout.nb_channels` value.
func (cl *AVChannelLayout) GetNbChannels() int32 {
	return (int32)(cl.nb_channels)
}

// SetNbChannels sets `AVChannelLayout.nb_channels` value.
func (cl *AVChannelLayout) SetNbChannels(v int32) {
	cl.nb_channels = (C.int)(v)
}

// GetNbChannelsAddr gets `AVChannelLayout.nb_channels` address.
func (cl *AVChannelLayout) GetNbChannelsAddr() *int32 {
	return (*int32)(&cl.nb_channels)
}

// GetUMask gets `AVChannelLayout.u.mask` value.
func (cl *AVChannelLayout) GetUMask() uint64 {
	return (uint64)(C.get_av_channel_layout_u_mask((*C.struct_AVChannelLayout)(cl)))
}

// SetUMask sets `AVChannelLayout.u.mask` value.
func (cl *AVChannelLayout) SetUMask(v uint64) {
	C.set_av_channel_layout_u_mask((*C.struct_AVChannelLayout)(cl), (C.uint64_t)(v))
}

// GetUMaskAddr gets `AVChannelLayout.u.mask` address.
func (cl *AVChannelLayout) GetUMaskAddr() *uint64 {
	return (*uint64)(C.get_av_channel_layout_u_mask_addr((*C.struct_AVChannelLayout)(cl)))
}

// GetUMap gets `AVChannelLayout.u.map` value.
func (cl *AVChannelLayout) GetUMap() *AVChannelCustom {

	return (*AVChannelCustom)(C.get_av_channel_layout_u_map((*C.struct_AVChannelLayout)(cl)))
}

// SetUMap sets `AVChannelLayout.u.map` value.
func (cl *AVChannelLayout) SetUMap(v *AVChannelCustom) {
	C.set_av_channel_layout_u_map((*C.struct_AVChannelLayout)(cl), (*C.struct_AVChannelCustom)(v))

}

// GetUMapAddr gets `AVChannelLayout.u.map` address.
func (cl *AVChannelLayout) GetUMapAddr() **AVChannelCustom {
	return (**AVChannelCustom)(
		unsafe.Pointer(C.get_av_channel_layout_u_map_addr((*C.struct_AVChannelLayout)(cl))))
}

// GetOpaque gets `AVChannelLayout.opaque` value.
func (cl *AVChannelLayout) GetOpaque() CVoidPointer {
	return VoidPointer(cl.opaque)
}

// SetOpaque sets `AVChannelLayout.opaque` value.
func (cl *AVChannelLayout) SetOpaque(v unsafe.Pointer) {
	cl.opaque = v
}

// GetOpaqueAddr gets `AVChannelLayout.opaque` address.
func (cl *AVChannelLayout) GetOpaqueAddr() *unsafe.Pointer {
	return &cl.opaque
}

func AV_CHANNEL_LAYOUT_MONO() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_mono)
}
func AV_CHANNEL_LAYOUT_STEREO() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_stereo)
}
func AV_CHANNEL_LAYOUT_2POINT1() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_2point1)
}
func AV_CHANNEL_LAYOUT_2_1() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_2_1)
}
func AV_CHANNEL_LAYOUT_SURROUND() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_surround)
}
func AV_CHANNEL_LAYOUT_3POINT1() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_3point1)
}
func AV_CHANNEL_LAYOUT_4POINT0() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_4point0)
}
func AV_CHANNEL_LAYOUT_4POINT1() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_4point1)
}
func AV_CHANNEL_LAYOUT_2_2() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_2_2)
}
func AV_CHANNEL_LAYOUT_QUAD() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_quad)
}
func AV_CHANNEL_LAYOUT_5POINT0() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_5point0)
}
func AV_CHANNEL_LAYOUT_5POINT1() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_5point1)
}
func AV_CHANNEL_LAYOUT_5POINT0_BACK() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_5point0_back)
}
func AV_CHANNEL_LAYOUT_5POINT1_BACK() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_5point1_back)
}
func AV_CHANNEL_LAYOUT_6POINT0() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_6point0)
}
func AV_CHANNEL_LAYOUT_6POINT0_FRONT() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_6point0_front)
}
func AV_CHANNEL_LAYOUT_3POINT1POINT2() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_3point1point2)
}
func AV_CHANNEL_LAYOUT_HEXAGONAL() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_hexagonal)
}
func AV_CHANNEL_LAYOUT_6POINT1() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_6point1)
}
func AV_CHANNEL_LAYOUT_6POINT1_BACK() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_6point1_back)
}
func AV_CHANNEL_LAYOUT_6POINT1_FRONT() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_6point1_front)
}
func AV_CHANNEL_LAYOUT_7POINT0() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_7point0)
}
func AV_CHANNEL_LAYOUT_7POINT0_FRONT() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_7point0_front)
}
func AV_CHANNEL_LAYOUT_7POINT1() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_7point1)
}
func AV_CHANNEL_LAYOUT_7POINT1_WIDE() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_7point1_wide)
}
func AV_CHANNEL_LAYOUT_7POINT1_WIDE_BACK() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_7point1_wide_back)
}
func AV_CHANNEL_LAYOUT_5POINT1POINT2_BACK() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_5point1point2_back)
}
func AV_CHANNEL_LAYOUT_OCTAGONAL() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_octagonal)
}
func AV_CHANNEL_LAYOUT_CUBE() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_cube)
}
func AV_CHANNEL_LAYOUT_5POINT1POINT4_BACK() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_5point1point4_back)
}
func AV_CHANNEL_LAYOUT_7POINT1POINT2() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_7point1point2)
}
func AV_CHANNEL_LAYOUT_7POINT1POINT4_BACK() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_7point1point4_back)
}
func AV_CHANNEL_LAYOUT_7POINT2POINT3() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_7point2point3)
}
func AV_CHANNEL_LAYOUT_9POINT1POINT4_BACK() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_9point1point4_back)
}
func AV_CHANNEL_LAYOUT_HEXADECAGONAL() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_hexadecagonal)
}
func AV_CHANNEL_LAYOUT_STEREO_DOWNMIX() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_stereo_downmix)
}
func AV_CHANNEL_LAYOUT_22POINT2() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_22point2)
}
func AV_CHANNEL_LAYOUT_AMBISONIC_FIRST_ORDER() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_ambisonic_first_order)
}
func AV_CHANNEL_LAYOUT_7POINT1_TOP_BACK() *AVChannelLayout {
	return (*AVChannelLayout)(&C.av_channel_layout_7point1_top_back)
}

// AvChannelName gets a human readable string in an abbreviated form describing a given channel.
func AvChannelName(channel AVChannel) string {
	buf := make([]C.char, AV_CH_LAYOUT_MAX_STRING_SIZE)
	if ret := C.av_channel_name((*C.char)(&buf[0]), (C.size_t)(AV_CH_LAYOUT_MAX_STRING_SIZE),
		(C.enum_AVChannel)(channel)); ret < 0 {
		return ""
	}
	return C.GoString(&buf[0])
}

// AvChannelNameBprint is bprint variant of AvChannelName().
func AvChannelNameBprint(bp *AVBPrint, channelId AVChannel) {
	C.av_channel_name_bprint((*C.struct_AVBPrint)(bp), (C.enum_AVChannel)(channelId))
}

// AvChannelDescription gets a human readable string describing a given channel.
func AvChannelDescription(channel AVChannel) string {
	buf := make([]C.char, AV_CH_LAYOUT_MAX_STRING_SIZE)
	if ret := C.av_channel_description((*C.char)(&buf[0]), (C.size_t)(AV_CH_LAYOUT_MAX_STRING_SIZE),
		(C.enum_AVChannel)(channel)); ret < 0 {
		return ""
	}
	return C.GoString(&buf[0])
}

// AvChannelDescriptionBprint is bprint variant of AvChannelDescription().
func AvChannelDescriptionBprint(bp *AVBPrint, channelId AVChannel) {
	C.av_channel_description_bprint((*C.struct_AVBPrint)(bp), (C.enum_AVChannel)(channelId))
}

// AvChannelFromString is the inverse function of AvChannelName().
func AvChannelFromString(name string) AVChannel {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (AVChannel)(C.av_channel_from_string((*C.char)(namePtr)))
}

// AvChannelLayoutCustomInit
func AvChannelLayoutCustomInit(channelLayout *AVChannelLayout, nbChannels int32) int32 {
	return (int32)(C.av_channel_layout_custom_init((*C.struct_AVChannelLayout)(channelLayout), (C.int)(nbChannels)))
}

// AvChannelLayoutFromMask initializes a native channel layout from a bitmask indicating which channels
// are present.
func AvChannelLayoutFromMask(channelLayout *AVChannelLayout, mask uint64) int32 {
	return (int32)(C.av_channel_layout_from_mask((*C.struct_AVChannelLayout)(channelLayout),
		(C.uint64_t)(mask)))
}

// AvChannelLayoutFromString initializes a channel layout from a given string description.
func AvChannelLayoutFromString(channelLayout *AVChannelLayout, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.av_channel_layout_from_string((*C.struct_AVChannelLayout)(channelLayout),
		(*C.char)(strPtr)))
}

// AvChannelLayoutDefault gets the default channel layout for a given number of channels.
func AvChannelLayoutDefault(chLayout *AVChannelLayout, nbChannels int32) {
	C.av_channel_layout_default((*C.struct_AVChannelLayout)(chLayout),
		(C.int)(nbChannels))
}

// AvChannelLayoutStandard iterates over all standard channel layouts.
func AvChannelLayoutStandard(opaque CVoidPointerPointer) *AVChannelLayout {
	return (*AVChannelLayout)(C.av_channel_layout_standard(VoidPointerPointer(opaque)))
}

// AvChannelLayoutUninit frees any allocated data in the channel layout and reset the channel
// count to 0.
func AvChannelLayoutUninit(channelLayout *AVChannelLayout) {
	C.av_channel_layout_uninit((*C.struct_AVChannelLayout)(channelLayout))
}

// AvChannelLayoutCopy makes a copy of a channel layout.
func AvChannelLayoutCopy(dst, src *AVChannelLayout) int32 {
	return (int32)(C.av_channel_layout_copy((*C.struct_AVChannelLayout)(dst),
		(*C.struct_AVChannelLayout)(src)))
}

// AvChannelLayoutDescribe gets a human-readable string describing the channel layout properties.
func AvChannelLayoutDescribe(channelLayout *AVChannelLayout) string {
	buf := make([]C.char, AV_CH_LAYOUT_MAX_STRING_SIZE)
	if ret := C.av_channel_layout_describe((*C.struct_AVChannelLayout)(channelLayout),
		(*C.char)(&buf[0]), (C.size_t)(AV_CH_LAYOUT_MAX_STRING_SIZE)); ret < 0 {
		return ""
	}
	return C.GoString(&buf[0])
}

// AvChannelLayoutDescribeBprint is bprint variant of AvChannelLayoutDescribe().
func AvChannelLayoutDescribeBprint(channelLayout *AVChannelLayout, bp *AVBPrint) int32 {
	return (int32)(C.av_channel_layout_describe_bprint(
		(*C.struct_AVChannelLayout)(channelLayout), (*C.struct_AVBPrint)(bp)))
}

// AvChannelLayoutChannelFromIndex gets the channel with the given index in a channel layout.
func AvChannelLayoutChannelFromIndex(channelLayout *AVChannelLayout, idx uint32) AVChannel {
	return (AVChannel)(C.av_channel_layout_channel_from_index(
		(*C.struct_AVChannelLayout)(channelLayout), (C.uint)(idx)))
}

// AvChannelLayoutIndexFromChannel gets the index of a given channel in a channel layout.
func AvChannelLayoutIndexFromChannel(channelLayout *AVChannelLayout, channel AVChannel) int32 {
	return (int32)(C.av_channel_layout_index_from_channel(
		(*C.struct_AVChannelLayout)(channelLayout), (C.enum_AVChannel)(channel)))
}

// AvChannelLayoutIndexFromString gets the index in a channel layout of a channel described by the given string.
func AvChannelLayoutIndexFromString(channelLayout *AVChannelLayout, name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_channel_layout_index_from_string(
		(*C.struct_AVChannelLayout)(channelLayout), (*C.char)(namePtr)))
}

// AvChannelLayoutChannelFromString gets a channel described by the given string.
func AvChannelLayoutChannelFromString(channelLayout *AVChannelLayout, name string) AVChannel {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (AVChannel)(C.av_channel_layout_channel_from_string(
		(*C.struct_AVChannelLayout)(channelLayout), (*C.char)(namePtr)))
}

// AvChannelLayoutSubset finds out what channels from a given set are present in a channel layout,
// without regard for their positions.
func AvChannelLayoutSubset(channelLayout *AVChannelLayout, mask uint64) uint64 {
	return (uint64)(C.av_channel_layout_subset(
		(*C.struct_AVChannelLayout)(channelLayout), (C.uint64_t)(mask)))
}

// AvChannelLayoutCheck checks whether a channel layout is valid, i.e. can possibly describe audio
// data.
func AvChannelLayoutCheck(channelLayout *AVChannelLayout) int32 {
	return (int32)(C.av_channel_layout_check(
		(*C.struct_AVChannelLayout)(channelLayout)))
}

// AvChannelLayoutCompare check whether two channel layouts are semantically the same, i.e. the same
// channels are present on the same positions in both.
func AvChannelLayoutCompare(chl, chl1 *AVChannelLayout) int32 {
	return (int32)(C.av_channel_layout_compare((*C.struct_AVChannelLayout)(chl),
		(*C.struct_AVChannelLayout)(chl1)))
}

const (
	AV_CHANNEL_LAYOUT_RETYPE_FLAG_LOSSLESS  = int32(C.AV_CHANNEL_LAYOUT_RETYPE_FLAG_LOSSLESS)
	AV_CHANNEL_LAYOUT_RETYPE_FLAG_CANONICAL = int32(C.AV_CHANNEL_LAYOUT_RETYPE_FLAG_CANONICAL)
)

// AvChannelLayoutRetype changes the AVChannelOrder of a channel layout.
func AvChannelLayoutRetype(channelLayout *AVChannelLayout, order AVChannelOrder, flags int32) int32 {
	return (int32)(C.av_channel_layout_retype((*C.struct_AVChannelLayout)(channelLayout),
		(C.enum_AVChannelOrder)(order), (C.int)(flags)))
}
