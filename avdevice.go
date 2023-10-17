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
func AvInputAudioDeviceNext(d *AvInputFormat) *AvInputFormat {
	return (*AvInputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(d)))
}

// AvInputVideoDeviceNext iterates video input devices.
func AvInputVideoDeviceNext(d *AvInputFormat) *AvInputFormat {
	return (*AvInputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(d)))
}

// AvOutputAudioDeviceNext iterates audio output devices.
func AvOutputAudioDeviceNext(d *AvOutputFormat) *AvOutputFormat {
	return (*AvOutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(d)))
}

// AvOutputVideoDeviceNext iterates video output devices.
func AvOutputVideoDeviceNext(d *AvOutputFormat) *AvOutputFormat {
	return (*AvOutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(d)))
}

// AvDeviceRect
type AvDeviceRect C.struct_AVDeviceRect

// AvAppToDevMessageType
type AvAppToDevMessageType = C.enum_AVAppToDevMessageType

const (
	AV_APP_TO_DEV_NONE           = AvAppToDevMessageType(C.AV_APP_TO_DEV_NONE)
	AV_APP_TO_DEV_WINDOW_SIZE    = AvAppToDevMessageType(C.AV_APP_TO_DEV_WINDOW_SIZE)
	AV_APP_TO_DEV_WINDOW_REPAINT = AvAppToDevMessageType(C.AV_APP_TO_DEV_WINDOW_REPAINT)
	AV_APP_TO_DEV_PAUSE          = AvAppToDevMessageType(C.AV_APP_TO_DEV_PAUSE)
	AV_APP_TO_DEV_PLAY           = AvAppToDevMessageType(C.AV_APP_TO_DEV_PLAY)
	AV_APP_TO_DEV_TOGGLE_PAUSE   = AvAppToDevMessageType(C.AV_APP_TO_DEV_TOGGLE_PAUSE)
	AV_APP_TO_DEV_SET_VOLUME     = AvAppToDevMessageType(C.AV_APP_TO_DEV_SET_VOLUME)
	AV_APP_TO_DEV_MUTE           = AvAppToDevMessageType(C.AV_APP_TO_DEV_MUTE)
	AV_APP_TO_DEV_UNMUTE         = AvAppToDevMessageType(C.AV_APP_TO_DEV_UNMUTE)
	AV_APP_TO_DEV_TOGGLE_MUTE    = AvAppToDevMessageType(C.AV_APP_TO_DEV_TOGGLE_MUTE)
	AV_APP_TO_DEV_GET_VOLUME     = AvAppToDevMessageType(C.AV_APP_TO_DEV_GET_VOLUME)
	AV_APP_TO_DEV_GET_MUTE       = AvAppToDevMessageType(C.AV_APP_TO_DEV_GET_MUTE)
)

// AvDevToAppMessageType
type AvDevToAppMessageType = C.enum_AVDevToAppMessageType

const (
	AV_DEV_TO_APP_NONE                  = AvDevToAppMessageType(C.AV_DEV_TO_APP_NONE)
	AV_DEV_TO_APP_CREATE_WINDOW_BUFFER  = AvDevToAppMessageType(C.AV_DEV_TO_APP_CREATE_WINDOW_BUFFER)
	AV_DEV_TO_APP_PREPARE_WINDOW_BUFFER = AvDevToAppMessageType(C.AV_DEV_TO_APP_PREPARE_WINDOW_BUFFER)
	AV_DEV_TO_APP_DISPLAY_WINDOW_BUFFER = AvDevToAppMessageType(C.AV_DEV_TO_APP_DISPLAY_WINDOW_BUFFER)
	AV_DEV_TO_APP_DESTROY_WINDOW_BUFFER = AvDevToAppMessageType(C.AV_DEV_TO_APP_DESTROY_WINDOW_BUFFER)
	AV_DEV_TO_APP_BUFFER_OVERFLOW       = AvDevToAppMessageType(C.AV_DEV_TO_APP_BUFFER_OVERFLOW)
	AV_DEV_TO_APP_BUFFER_UNDERFLOW      = AvDevToAppMessageType(C.AV_DEV_TO_APP_BUFFER_UNDERFLOW)
	AV_DEV_TO_APP_BUFFER_READABLE       = AvDevToAppMessageType(C.AV_DEV_TO_APP_BUFFER_READABLE)
	AV_DEV_TO_APP_BUFFER_WRITABLE       = AvDevToAppMessageType(C.AV_DEV_TO_APP_BUFFER_WRITABLE)
	AV_DEV_TO_APP_MUTE_STATE_CHANGED    = AvDevToAppMessageType(C.AV_DEV_TO_APP_MUTE_STATE_CHANGED)
	AV_DEV_TO_APP_VOLUME_LEVEL_CHANGED  = AvDevToAppMessageType(C.AV_DEV_TO_APP_VOLUME_LEVEL_CHANGED)
)

// AvDeviceAppToDevControlMessage sends control message from application to device.
func AvDeviceAppToDevControlMessage(s *AvFormatContext,
	_type AvAppToDevMessageType, data unsafe.Pointer, dataSize uint) int32 {
	return (int32)(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(s),
		(C.enum_AVAppToDevMessageType)(_type), data, (C.size_t)(dataSize)))
}

// AvDeviceDevToAppControlMessage sends control message from device to application.
func AvDeviceDevToAppControlMessage(s *AvFormatContext,
	_type AvDevToAppMessageType, data unsafe.Pointer, dataSize uint) int32 {
	return (int32)(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(s),
		(C.enum_AVDevToAppMessageType)(_type), data, (C.size_t)(dataSize)))
}

// AvDeviceCapabilitiesQuery
type AvDeviceCapabilitiesQuery C.struct_AVDeviceCapabilitiesQuery

// Deprecated: No use
func AvDeviceCapabilitiesCreate(caps **AvDeviceCapabilitiesQuery,
	s *AvFormatContext, deviceOptions **AvDictionary) int32 {
	return (int32)(C.avdevice_capabilities_create((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(caps)),
		(*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(deviceOptions))))
}

// Deprecated: No use
func AvDeviceCapabilitiesFree(caps **AvDeviceCapabilitiesQuery, s *AvFormatContext) {
	C.avdevice_capabilities_free((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(caps)),
		(*C.struct_AVFormatContext)(s))
}

// AvDeviceInfo
type AvDeviceInfo C.struct_AVDeviceInfo

// AvDeviceInfoList
type AvDeviceInfoList C.struct_AVDeviceInfoList

// AvDeviceListDevices returns available device names and their parameters.
func AvDeviceListDevices(s *AvFormatContext, deviceList **AvDeviceInfoList) int32 {
	return (int32)(C.avdevice_list_devices((*C.struct_AVFormatContext)(s),
		(**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList))))
}

// AvDeviceFreeListDevices frees result of AvDeviceListDevices().
func AvDeviceFreeListDevices(deviceList **AvDeviceInfoList) {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList)))
}

// AvDeviceListInputSources lists input devices.
func AvDeviceListInputSources(device *AvInputFormat, deviceName string,
	deviceOptions *AvDictionary, deviceList **AvDeviceInfoList) int32 {
	deviceNamePtr, deviceNameFunc := StringCasting(deviceName)
	defer deviceNameFunc()
	return (int32)(C.avdevice_list_input_sources((*C.struct_AVInputFormat)(device),
		(*C.char)(deviceNamePtr), (*C.struct_AVDictionary)(deviceOptions),
		(**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList))))
}

// AvDeviceListOutputSinks lists output devices.
func AvDeviceListOutputSinks(device *AvOutputFormat, deviceName string,
	deviceOptions *AvDictionary, deviceList **AvDeviceInfoList) int32 {
	deviceNamePtr, deviceNameFunc := StringCasting(deviceName)
	defer deviceNameFunc()
	return (int32)(C.avdevice_list_output_sinks((*C.struct_AVOutputFormat)(device),
		(*C.char)(deviceNamePtr), (*C.struct_AVDictionary)(deviceOptions),
		(**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList))))
}
