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

// Custom: GetX gets `AVDeviceRect.x` value.
func (dr *AVDeviceRect) GetX() int32 {
	return (int32)(dr.x)
}

// Custom: SetX sets `AVDeviceRect.x` value.
func (dr *AVDeviceRect) SetX(v int32) {
	dr.x = (C.int)(v)
}

// Custom: GetXAddr gets `AVDeviceRect.x` address.
func (dr *AVDeviceRect) GetXAddr() *int32 {
	return (*int32)(&dr.x)
}

// Custom: GetY gets `AVDeviceRect.y` value.
func (dr *AVDeviceRect) GetY() int32 {
	return (int32)(dr.y)
}

// Custom: SetY sets `AVDeviceRect.y` value.
func (dr *AVDeviceRect) SetY(v int32) {
	dr.y = (C.int)(v)
}

// Custom: GetYAddr gets `AVDeviceRect.y` address.
func (dr *AVDeviceRect) GetYAddr() *int32 {
	return (*int32)(&dr.y)
}

// Custom: GetWidth gets `AVDeviceRect.width` value.
func (dr *AVDeviceRect) GetWidth() int32 {
	return (int32)(dr.width)
}

// Custom: SetWidth sets `AVDeviceRect.width` value.
func (dr *AVDeviceRect) SetWidth(v int32) {
	dr.width = (C.int)(v)
}

// Custom: GetWidthAddr gets `AVDeviceRect.width` address.
func (dr *AVDeviceRect) GetWidthAddr() *int32 {
	return (*int32)(&dr.width)
}

// Custom: GetHeight gets `AVDeviceRect.height` value.
func (dr *AVDeviceRect) GetHeight() int32 {
	return (int32)(dr.height)
}

// Custom: SetHeight sets `AVDeviceRect.height` value.
func (dr *AVDeviceRect) SetHeight(v int32) {
	dr.height = (C.int)(v)
}

// Custom: GetHeightAddr gets `AVDeviceRect.height` address.
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
	_type AVAppToDevMessageType, data CVoidPointer, dataSize uint) int32 {
	return (int32)(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(s),
		(C.enum_AVAppToDevMessageType)(_type), VoidPointer(data), (C.size_t)(dataSize)))
}

// AvDeviceDevToAppControlMessage sends control message from device to application.
func AvDeviceDevToAppControlMessage(s *AVFormatContext,
	_type AVDevToAppMessageType, data CVoidPointer, dataSize uint) int32 {
	return (int32)(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(s),
		(C.enum_AVDevToAppMessageType)(_type), VoidPointer(data), (C.size_t)(dataSize)))
}

// AVDeviceCapabilitiesQuery
type AVDeviceCapabilitiesQuery C.struct_AVDeviceCapabilitiesQuery

// Custom: GetAvClass gets `AVDeviceCapabilitiesQuery.av_class` value.
func (dc *AVDeviceCapabilitiesQuery) GetAvClass() *AVClass {
	return (*AVClass)(dc.av_class)
}

// Custom: SetAvClass sets `AVDeviceCapabilitiesQuery.av_class` value.
func (dc *AVDeviceCapabilitiesQuery) SetAvClass(v *AVClass) {
	dc.av_class = (*C.struct_AVClass)(v)
}

// Custom: GetAvClassAddr gets `AVDeviceCapabilitiesQuery.av_class` address.
func (dc *AVDeviceCapabilitiesQuery) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&dc.av_class))
}

// Custom: GetDeviceContext gets `AVDeviceCapabilitiesQuery.device_context` value.
func (dc *AVDeviceCapabilitiesQuery) GetDeviceContext() *AVFormatContext {
	return (*AVFormatContext)(dc.device_context)
}

// Custom: SetDeviceContext sets `AVDeviceCapabilitiesQuery.device_context` value.
func (dc *AVDeviceCapabilitiesQuery) SetDeviceContext(v *AVFormatContext) {
	dc.device_context = (*C.struct_AVFormatContext)(v)
}

// Custom: GetDeviceContextAddr gets `AVDeviceCapabilitiesQuery.device_context` address.
func (dc *AVDeviceCapabilitiesQuery) GetDeviceContextAddr() **AVFormatContext {
	return (**AVFormatContext)(unsafe.Pointer(&dc.device_context))
}

// Custom: GetCodec gets `AVDeviceCapabilitiesQuery.codec` value.
func (dc *AVDeviceCapabilitiesQuery) GetCodec() AVCodecID {
	return (AVCodecID)(dc.codec)
}

// Custom: SetCodec sets `AVDeviceCapabilitiesQuery.codec` value.
func (dc *AVDeviceCapabilitiesQuery) SetCodec(v AVCodecID) {
	dc.codec = (C.enum_AVCodecID)(v)
}

// Custom: GetCodecAddr gets `AVDeviceCapabilitiesQuery.codec` address.
func (dc *AVDeviceCapabilitiesQuery) GetCodecAddr() *AVCodecID {
	return (*AVCodecID)(&dc.codec)
}

// Custom: GetSampleFormat gets `AVDeviceCapabilitiesQuery.sample_format` value.
func (dc *AVDeviceCapabilitiesQuery) GetSampleFormat() AVSampleFormat {
	return (AVSampleFormat)(dc.sample_format)
}

// Custom: SetSampleFormat sets `AVDeviceCapabilitiesQuery.sample_format` value.
func (dc *AVDeviceCapabilitiesQuery) SetSampleFormat(v AVSampleFormat) {
	dc.sample_format = (C.enum_AVSampleFormat)(v)
}

// Custom: GetSampleFormatAddr gets `AVDeviceCapabilitiesQuery.sample_format` address.
func (dc *AVDeviceCapabilitiesQuery) GetSampleFormatAddr() *AVSampleFormat {
	return (*AVSampleFormat)(&dc.sample_format)
}

// Custom: GetPixelFormat gets `AVDeviceCapabilitiesQuery.pixel_format` value.
func (dc *AVDeviceCapabilitiesQuery) GetPixelFormat() AVPixelFormat {
	return (AVPixelFormat)(dc.pixel_format)
}

// Custom: SetPixelFormat sets `AVDeviceCapabilitiesQuery.pixel_format` value.
func (dc *AVDeviceCapabilitiesQuery) SetPixelFormat(v AVPixelFormat) {
	dc.pixel_format = (C.enum_AVPixelFormat)(v)
}

// Custom: GetPixelFormatAddr gets `AVDeviceCapabilitiesQuery.pixel_format` address.
func (dc *AVDeviceCapabilitiesQuery) GetPixelFormatAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&dc.pixel_format)
}

// Custom: GetSampleRate gets `AVDeviceCapabilitiesQuery.sample_rate` value.
func (dc *AVDeviceCapabilitiesQuery) GetSampleRate() int32 {
	return (int32)(dc.sample_rate)
}

// Custom: SetSampleRate sets `AVDeviceCapabilitiesQuery.sample_rate` value.
func (dc *AVDeviceCapabilitiesQuery) SetSampleRate(v int32) {
	dc.sample_rate = (C.int)(v)
}

// Custom: GetSampleRateAddr gets `AVDeviceCapabilitiesQuery.sample_rate` address.
func (dc *AVDeviceCapabilitiesQuery) GetSampleRateAddr() *int32 {
	return (*int32)(&dc.sample_rate)
}

// Custom: GetChannels gets `AVDeviceCapabilitiesQuery.channels` value.
func (dc *AVDeviceCapabilitiesQuery) GetChannels() int32 {
	return (int32)(dc.channels)
}

// Custom: SetChannels sets `AVDeviceCapabilitiesQuery.channels` value.
func (dc *AVDeviceCapabilitiesQuery) SetChannels(v int32) {
	dc.channels = (C.int)(v)
}

// Custom: GetChannelsAddr gets `AVDeviceCapabilitiesQuery.channels` address.
func (dc *AVDeviceCapabilitiesQuery) GetChannelsAddr() *int32 {
	return (*int32)(&dc.channels)
}

// Custom: GetChannelLayout gets `AVDeviceCapabilitiesQuery.channel_layout` value.
func (dc *AVDeviceCapabilitiesQuery) GetChannelLayout() int64 {
	return (int64)(dc.channel_layout)
}

// Custom: SetChannelLayout sets `AVDeviceCapabilitiesQuery.channel_layout` value.
func (dc *AVDeviceCapabilitiesQuery) SetChannelLayout(v int64) {
	dc.channel_layout = (C.int64_t)(v)
}

// Custom: GetChannelLayoutAddr gets `AVDeviceCapabilitiesQuery.channel_layout` address.
func (dc *AVDeviceCapabilitiesQuery) GetChannelLayoutAddr() *int64 {
	return (*int64)(&dc.channel_layout)
}

// Custom: GetWindowWidth gets `AVDeviceCapabilitiesQuery.window_width` value.
func (dc *AVDeviceCapabilitiesQuery) GetWindowWidth() int32 {
	return (int32)(dc.window_width)
}

// Custom: SetWindowWidth sets `AVDeviceCapabilitiesQuery.window_width` value.
func (dc *AVDeviceCapabilitiesQuery) SetWindowWidth(v int32) {
	dc.window_width = (C.int)(v)
}

// Custom: GetWindowWidthAddr gets `AVDeviceCapabilitiesQuery.window_width` address.
func (dc *AVDeviceCapabilitiesQuery) GetWindowWidthAddr() *int32 {
	return (*int32)(&dc.window_width)
}

// Custom: GetWindowHeight gets `AVDeviceCapabilitiesQuery.window_height` value.
func (dc *AVDeviceCapabilitiesQuery) GetWindowHeight() int32 {
	return (int32)(dc.window_height)
}

// Custom: SetWindowHeight sets `AVDeviceCapabilitiesQuery.window_height` value.
func (dc *AVDeviceCapabilitiesQuery) SetWindowHeight(v int32) {
	dc.window_height = (C.int)(v)
}

// Custom: GetWindowHeightAddr gets `AVDeviceCapabilitiesQuery.window_height` address.
func (dc *AVDeviceCapabilitiesQuery) GetWindowHeightAddr() *int32 {
	return (*int32)(&dc.window_height)
}

// Custom: GetFrameWidth gets `AVDeviceCapabilitiesQuery.frame_width` value.
func (dc *AVDeviceCapabilitiesQuery) GetFrameWidth() int32 {
	return (int32)(dc.frame_width)
}

// Custom: SetFrameWidth sets `AVDeviceCapabilitiesQuery.frame_width` value.
func (dc *AVDeviceCapabilitiesQuery) SetFrameWidth(v int32) {
	dc.frame_width = (C.int)(v)
}

// Custom: GetFrameWidthAddr gets `AVDeviceCapabilitiesQuery.frame_width` address.
func (dc *AVDeviceCapabilitiesQuery) GetFrameWidthAddr() *int32 {
	return (*int32)(&dc.frame_width)
}

// Custom: GetFrameHeight gets `AVDeviceCapabilitiesQuery.frame_height` value.
func (dc *AVDeviceCapabilitiesQuery) GetFrameHeight() int32 {
	return (int32)(dc.frame_height)
}

// Custom: SetFrameHeight sets `AVDeviceCapabilitiesQuery.frame_height` value.
func (dc *AVDeviceCapabilitiesQuery) SetFrameHeight(v int32) {
	dc.frame_height = (C.int)(v)
}

// Custom: GetFrameHeightAddr gets `AVDeviceCapabilitiesQuery.frame_height` address.
func (dc *AVDeviceCapabilitiesQuery) GetFrameHeightAddr() *int32 {
	return (*int32)(&dc.frame_height)
}

// Custom: GetFps gets `AVDeviceCapabilitiesQuery.fps` value.
func (dc *AVDeviceCapabilitiesQuery) GetFps() AVRational {
	return (AVRational)(dc.fps)
}

// Custom: SetFps sets `AVDeviceCapabilitiesQuery.fps` value.
func (dc *AVDeviceCapabilitiesQuery) SetFps(v AVRational) {
	dc.fps = (C.struct_AVRational)(v)
}

// Custom: GetFpsAddr gets `AVDeviceCapabilitiesQuery.fps` address.
func (dc *AVDeviceCapabilitiesQuery) GetFpsAddr() *AVRational {
	return (*AVRational)(&dc.fps)
}

// Deprecated: No use
func AvDeviceCapabilitiesCreate(caps **AVDeviceCapabilitiesQuery,
	s *AVFormatContext, deviceOptions **AVDictionary) int32 {
	return (int32)(C.avdevice_capabilities_create((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(caps)),
		(*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(deviceOptions))))
}

// Deprecated: No use
func AvDeviceCapabilitiesFree(caps **AVDeviceCapabilitiesQuery, s *AVFormatContext) {
	C.avdevice_capabilities_free((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(caps)),
		(*C.struct_AVFormatContext)(s))
}

// AVDeviceInfo
type AVDeviceInfo C.struct_AVDeviceInfo

// Custom: GetDeviceName gets `AVDeviceInfo.device_name` value.
func (di *AVDeviceInfo) GetDeviceName() string {
	return C.GoString(di.device_name)
}

// Custom: GetDeviceDescription gets `AVDeviceInfo.device_description` value.
func (di *AVDeviceInfo) GetDeviceDescription() string {
	return C.GoString(di.device_description)
}

// AVDeviceInfoList
type AVDeviceInfoList C.struct_AVDeviceInfoList

// Custom: GetDevices gets `AVDeviceInfoList.devices` value.
func (dcl *AVDeviceInfoList) GetDevices() []*AVDeviceInfo {
	if dcl.devices == nil {
		return nil
	}
	return unsafe.Slice((**AVDeviceInfo)(unsafe.Pointer(dcl.devices)), dcl.nb_devices)
}

// Custom: GetDevicesIndex gets `AVDeviceInfoList.devices` index value.
func (dcl *AVDeviceInfoList) GetDevicesIdx(idx int) *AVDeviceInfo {
	if idx >= int(dcl.nb_devices) {
		return nil
	}
	return PointerOffset((*AVDeviceInfo)(*dcl.devices), idx)
}

// Custom: GetNbDevices gets `AVDeviceInfoList.nb_devices` value.
func (dcl *AVDeviceInfoList) GetNbDevices() int32 {
	return (int32)(dcl.nb_devices)
}

// Custom: GetDefaultDevice gets `AVDeviceInfoList.default_device` value.
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
