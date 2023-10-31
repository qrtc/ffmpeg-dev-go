// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavdevice/avdevice.h>
*/
import "C"
import "unsafe"

// AvDeviceVersion returns the LIBAVDEVICE_VERSION_INT constant.
func AvDeviceVersion() uint32 {
	return (uint32)(C.avdevice_version())
}

// AvDeviceConfiguration returns the libavdevice build-time configuration.
func AvDeviceConfiguration() string {
	return C.GoString(C.avdevice_configuration())
}

// AvDeviceLicense returns the libavdevice license.
func AvDeviceLicense() string {
	return C.GoString(C.avdevice_license())
}

// AvDeviceRegisterAll initializes libavdevice and register all the input and output devices.
func AvDeviceRegisterAll() {
	C.avdevice_register_all()
}

// AvInputAudioDeviceNext iterates audio input devices.
func AvInputAudioDeviceNext(d *AVInputFormat) *AVInputFormat {
	return (*AVInputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(d)))
}

// AvInputVideoDeviceNext iterates video input devices.
func AvInputVideoDeviceNext(d *AVInputFormat) *AVInputFormat {
	return (*AVInputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(d)))
}

// AvOutputAudioDeviceNext iterates audio output devices.
func AvOutputAudioDeviceNext(d *AVOutputFormat) *AVOutputFormat {
	return (*AVOutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(d)))
}

// AvOutputVideoDeviceNext iterates video output devices.
func AvOutputVideoDeviceNext(d *AVOutputFormat) *AVOutputFormat {
	return (*AVOutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(d)))
}

// AVDeviceRect
type AVDeviceRect C.struct_AVDeviceRect

// GetX gets `AVDeviceRect.x` value.
func (dr *AVDeviceRect) GetX() int32 {
	return (int32)(dr.x)
}

// SetX sets `AVDeviceRect.x` value.
func (dr *AVDeviceRect) SetX(v int32) {
	dr.x = (C.int)(v)
}

// GetXAddr gets `AVDeviceRect.x` address.
func (dr *AVDeviceRect) GetXAddr() *int32 {
	return (*int32)(&dr.x)
}

// GetY gets `AVDeviceRect.y` value.
func (dr *AVDeviceRect) GetY() int32 {
	return (int32)(dr.y)
}

// SetY sets `AVDeviceRect.y` value.
func (dr *AVDeviceRect) SetY(v int32) {
	dr.y = (C.int)(v)
}

// GetYAddr gets `AVDeviceRect.y` address.
func (dr *AVDeviceRect) GetYAddr() *int32 {
	return (*int32)(&dr.y)
}

// GetWidth gets `AVDeviceRect.width` value.
func (dr *AVDeviceRect) GetWidth() int32 {
	return (int32)(dr.width)
}

// SetWidth sets `AVDeviceRect.width` value.
func (dr *AVDeviceRect) SetWidth(v int32) {
	dr.width = (C.int)(v)
}

// GetWidthAddr gets `AVDeviceRect.width` address.
func (dr *AVDeviceRect) GetWidthAddr() *int32 {
	return (*int32)(&dr.width)
}

// GetHeight gets `AVDeviceRect.height` value.
func (dr *AVDeviceRect) GetHeight() int32 {
	return (int32)(dr.height)
}

// SetHeight sets `AVDeviceRect.height` value.
func (dr *AVDeviceRect) SetHeight(v int32) {
	dr.height = (C.int)(v)
}

// GetHeightAddr gets `AVDeviceRect.height` address.
func (dr *AVDeviceRect) GetHeightAddr() *int32 {
	return (*int32)(&dr.height)
}

// AVAppToDevMessageType
type AVAppToDevMessageType = C.enum_AVAppToDevMessageType

const (
	AV_APP_TO_DEV_NONE           = AVAppToDevMessageType(C.AV_APP_TO_DEV_NONE)
	AV_APP_TO_DEV_WINDOW_SIZE    = AVAppToDevMessageType(C.AV_APP_TO_DEV_WINDOW_SIZE)
	AV_APP_TO_DEV_WINDOW_REPAINT = AVAppToDevMessageType(C.AV_APP_TO_DEV_WINDOW_REPAINT)
	AV_APP_TO_DEV_PAUSE          = AVAppToDevMessageType(C.AV_APP_TO_DEV_PAUSE)
	AV_APP_TO_DEV_PLAY           = AVAppToDevMessageType(C.AV_APP_TO_DEV_PLAY)
	AV_APP_TO_DEV_TOGGLE_PAUSE   = AVAppToDevMessageType(C.AV_APP_TO_DEV_TOGGLE_PAUSE)
	AV_APP_TO_DEV_SET_VOLUME     = AVAppToDevMessageType(C.AV_APP_TO_DEV_SET_VOLUME)
	AV_APP_TO_DEV_MUTE           = AVAppToDevMessageType(C.AV_APP_TO_DEV_MUTE)
	AV_APP_TO_DEV_UNMUTE         = AVAppToDevMessageType(C.AV_APP_TO_DEV_UNMUTE)
	AV_APP_TO_DEV_TOGGLE_MUTE    = AVAppToDevMessageType(C.AV_APP_TO_DEV_TOGGLE_MUTE)
	AV_APP_TO_DEV_GET_VOLUME     = AVAppToDevMessageType(C.AV_APP_TO_DEV_GET_VOLUME)
	AV_APP_TO_DEV_GET_MUTE       = AVAppToDevMessageType(C.AV_APP_TO_DEV_GET_MUTE)
)

// AVDevToAppMessageType
type AVDevToAppMessageType = C.enum_AVDevToAppMessageType

const (
	AV_DEV_TO_APP_NONE                  = AVDevToAppMessageType(C.AV_DEV_TO_APP_NONE)
	AV_DEV_TO_APP_CREATE_WINDOW_BUFFER  = AVDevToAppMessageType(C.AV_DEV_TO_APP_CREATE_WINDOW_BUFFER)
	AV_DEV_TO_APP_PREPARE_WINDOW_BUFFER = AVDevToAppMessageType(C.AV_DEV_TO_APP_PREPARE_WINDOW_BUFFER)
	AV_DEV_TO_APP_DISPLAY_WINDOW_BUFFER = AVDevToAppMessageType(C.AV_DEV_TO_APP_DISPLAY_WINDOW_BUFFER)
	AV_DEV_TO_APP_DESTROY_WINDOW_BUFFER = AVDevToAppMessageType(C.AV_DEV_TO_APP_DESTROY_WINDOW_BUFFER)
	AV_DEV_TO_APP_BUFFER_OVERFLOW       = AVDevToAppMessageType(C.AV_DEV_TO_APP_BUFFER_OVERFLOW)
	AV_DEV_TO_APP_BUFFER_UNDERFLOW      = AVDevToAppMessageType(C.AV_DEV_TO_APP_BUFFER_UNDERFLOW)
	AV_DEV_TO_APP_BUFFER_READABLE       = AVDevToAppMessageType(C.AV_DEV_TO_APP_BUFFER_READABLE)
	AV_DEV_TO_APP_BUFFER_WRITABLE       = AVDevToAppMessageType(C.AV_DEV_TO_APP_BUFFER_WRITABLE)
	AV_DEV_TO_APP_MUTE_STATE_CHANGED    = AVDevToAppMessageType(C.AV_DEV_TO_APP_MUTE_STATE_CHANGED)
	AV_DEV_TO_APP_VOLUME_LEVEL_CHANGED  = AVDevToAppMessageType(C.AV_DEV_TO_APP_VOLUME_LEVEL_CHANGED)
)

// AvDeviceAppToDevControlMessage sends control message from application to device.
func AvDeviceAppToDevControlMessage(s *AVFormatContext,
	_type AVAppToDevMessageType, data CVoidPointer, dataSize uintptr) int32 {
	return (int32)(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(s),
		(C.enum_AVAppToDevMessageType)(_type), VoidPointer(data), (C.size_t)(dataSize)))
}

// AvDeviceDevToAppControlMessage sends control message from device to application.
func AvDeviceDevToAppControlMessage(s *AVFormatContext,
	_type AVDevToAppMessageType, data CVoidPointer, dataSize uintptr) int32 {
	return (int32)(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(s),
		(C.enum_AVDevToAppMessageType)(_type), VoidPointer(data), (C.size_t)(dataSize)))
}

// AVDeviceCapabilitiesQuery
type AVDeviceCapabilitiesQuery C.struct_AVDeviceCapabilitiesQuery

// NONEED: av_device_capabilities

// AVDeviceInfo
type AVDeviceInfo C.struct_AVDeviceInfo

// GetDeviceName gets `AVDeviceInfo.device_name` value.
func (di *AVDeviceInfo) GetDeviceName() string {
	return C.GoString(di.device_name)
}

// GetDeviceDescription gets `AVDeviceInfo.device_description` value.
func (di *AVDeviceInfo) GetDeviceDescription() string {
	return C.GoString(di.device_description)
}

// AVDeviceInfoList
type AVDeviceInfoList C.struct_AVDeviceInfoList

// GetDevices gets `AVDeviceInfoList.devices` value.
func (dcl *AVDeviceInfoList) GetDevices() []*AVDeviceInfo {
	if dcl.devices == nil {
		return nil
	}
	return unsafe.Slice((**AVDeviceInfo)(unsafe.Pointer(dcl.devices)), dcl.nb_devices)
}

// GetNbDevices gets `AVDeviceInfoList.nb_devices` value.
func (dcl *AVDeviceInfoList) GetNbDevices() int32 {
	return (int32)(dcl.nb_devices)
}

// GetDefaultDevice gets `AVDeviceInfoList.default_device` value.
func (dcl *AVDeviceInfoList) GetDefaultDevice() int32 {
	return (int32)(dcl.default_device)
}

// AvDeviceListDevices returns available device names and their parameters.
func AvDeviceListDevices(s *AVFormatContext, deviceList **AVDeviceInfoList) int32 {
	return (int32)(C.avdevice_list_devices((*C.struct_AVFormatContext)(s),
		(**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList))))
}

// AvDeviceFreeListDevices frees result of AvDeviceListDevices().
func AvDeviceFreeListDevices(deviceList **AVDeviceInfoList) {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList)))
}

// AvDeviceListInputSources lists input devices.
func AvDeviceListInputSources(device *AVInputFormat, deviceName string,
	deviceOptions *AVDictionary, deviceList **AVDeviceInfoList) int32 {
	deviceNamePtr, deviceNameFunc := StringCasting(deviceName)
	defer deviceNameFunc()
	return (int32)(C.avdevice_list_input_sources((*C.struct_AVInputFormat)(device),
		(*C.char)(deviceNamePtr), (*C.struct_AVDictionary)(deviceOptions),
		(**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList))))
}

// AvDeviceListOutputSinks lists output devices.
func AvDeviceListOutputSinks(device *AVOutputFormat, deviceName string,
	deviceOptions *AVDictionary, deviceList **AVDeviceInfoList) int32 {
	deviceNamePtr, deviceNameFunc := StringCasting(deviceName)
	defer deviceNameFunc()
	return (int32)(C.avdevice_list_output_sinks((*C.struct_AVOutputFormat)(device),
		(*C.char)(deviceNamePtr), (*C.struct_AVDictionary)(deviceOptions),
		(**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList))))
}
