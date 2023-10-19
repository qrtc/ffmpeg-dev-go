package ffmpeg

/*
#include <libavutil/hwcontext.h>

typedef void (*av_hw_device_context_free_func)(struct AVHWDeviceContext *ctx);
typedef void (*av_hw_frames_context_free_func)(struct AVHWFramesContext *ctx);
*/
import "C"
import "unsafe"

// AVHWDeviceType
type AVHWDeviceType = C.enum_AVHWDeviceType

const (
	AV_HWDEVICE_TYPE_NONE         = AVHWDeviceType(C.AV_HWDEVICE_TYPE_NONE)
	AV_HWDEVICE_TYPE_VDPAU        = AVHWDeviceType(C.AV_HWDEVICE_TYPE_VDPAU)
	AV_HWDEVICE_TYPE_CUDA         = AVHWDeviceType(C.AV_HWDEVICE_TYPE_CUDA)
	AV_HWDEVICE_TYPE_VAAPI        = AVHWDeviceType(C.AV_HWDEVICE_TYPE_VAAPI)
	AV_HWDEVICE_TYPE_DXVA2        = AVHWDeviceType(C.AV_HWDEVICE_TYPE_DXVA2)
	AV_HWDEVICE_TYPE_QSV          = AVHWDeviceType(C.AV_HWDEVICE_TYPE_QSV)
	AV_HWDEVICE_TYPE_VIDEOTOOLBOX = AVHWDeviceType(C.AV_HWDEVICE_TYPE_VIDEOTOOLBOX)
	AV_HWDEVICE_TYPE_D3D11VA      = AVHWDeviceType(C.AV_HWDEVICE_TYPE_D3D11VA)
	AV_HWDEVICE_TYPE_DRM          = AVHWDeviceType(C.AV_HWDEVICE_TYPE_DRM)
	AV_HWDEVICE_TYPE_OPENCL       = AVHWDeviceType(C.AV_HWDEVICE_TYPE_OPENCL)
	AV_HWDEVICE_TYPE_MEDIACODEC   = AVHWDeviceType(C.AV_HWDEVICE_TYPE_MEDIACODEC)
	AV_HWDEVICE_TYPE_VULKAN       = AVHWDeviceType(C.AV_HWDEVICE_TYPE_VULKAN)
)

type AVHWDeviceInternal C.struct_AVHWDeviceInternal

type AVHWDeviceContext C.struct_AVHWDeviceContext

type AVHWDeviceContextFreeFunc = C.av_hw_device_context_free_func

// Custom: GetAvClass gets `AVHWDeviceContext.av_class` value.
func (dc *AVHWDeviceContext) GetAvClass() *AVClass {
	return (*AVClass)(dc.av_class)
}

// Custom: SetAvClass sets `AVHWDeviceContext.av_class` value.
func (dc *AVHWDeviceContext) SetAvClass(v *AVClass) {
	dc.av_class = (*C.struct_AVClass)(v)
}

// Custom: GetAvClassAddr gets `AVHWDeviceContext.av_class` address.
func (dc *AVHWDeviceContext) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&dc.av_class))
}

// Custom: GetInternal gets `AVHWDeviceContext.internal` value.
func (dc *AVHWDeviceContext) GetInternal() *AVHWDeviceInternal {
	return (*AVHWDeviceInternal)(dc.internal)
}

// Custom: SetInternal sets `AVHWDeviceContext.internal` value.
func (dc *AVHWDeviceContext) SetInternal(v *AVHWDeviceInternal) {
	dc.internal = (*C.struct_AVHWDeviceInternal)(v)
}

// Custom: GetInternalAddr gets `AVHWDeviceContext.internal` address.
func (dc *AVHWDeviceContext) GetInternalAddr() **AVHWDeviceInternal {
	return (**AVHWDeviceInternal)(unsafe.Pointer(&dc.internal))
}

// Custom: GetType gets `AVHWDeviceContext.type` value.
func (dc *AVHWDeviceContext) GetType() AVHWDeviceType {
	return (AVHWDeviceType)(dc._type)
}

// Custom: SetType sets `AVHWDeviceContext.type` value.
func (dc *AVHWDeviceContext) SetType(v AVHWDeviceType) {
	dc._type = (C.enum_AVHWDeviceType)(v)
}

// Custom: GetTypeAddr gets `AVHWDeviceContext.type` address.
func (dc *AVHWDeviceContext) GetTypeAddr() *AVHWDeviceType {
	return (*AVHWDeviceType)(&dc._type)
}

// Custom: GetHwctx gets `AVHWDeviceContext.hwctx` value.
func (dc *AVHWDeviceContext) GetHwctx() unsafe.Pointer {
	return dc.hwctx
}

// Custom: SetHwctx sets `AVHWDeviceContext.hwctx` value.
func (dc *AVHWDeviceContext) SetHwctx(v CVoidPointer) {
	dc.hwctx = VoidPointer(v)
}

// Custom: GetHwctxAddr gets `AVHWDeviceContext.hwctx` address.
func (dc *AVHWDeviceContext) GetHwctxAddr() *unsafe.Pointer {
	return &dc.hwctx
}

// Custom: SetFree sets `AVHWDeviceContext.free` value.
func (dc *AVHWDeviceContext) SetFree(v AVHWDeviceContextFreeFunc) {
	dc.free = (C.av_hw_device_context_free_func)(v)
}

// Custom: GetUserOpaque gets `AVHWDeviceContext.user_opaque` value.
func (dc *AVHWDeviceContext) GetUserOpaque() unsafe.Pointer {
	return dc.user_opaque
}

// Custom: SetUserOpaque sets `AVHWDeviceContext.user_opaque` value.
func (dc *AVHWDeviceContext) SetUserOpaque(v CVoidPointer) {
	dc.user_opaque = VoidPointer(v)
}

// Custom: GetUserOpaqueAddr gets `AVHWDeviceContext.user_opaque` address.
func (dc *AVHWDeviceContext) GetUserOpaqueAddr() *unsafe.Pointer {
	return &dc.user_opaque
}

type AVHWFramesInternal C.struct_AVHWFramesInternal

type AVHWFramesContext C.struct_AVHWFramesContext

type AVHWFramesContextFreeFunc = C.av_hw_frames_context_free_func

// Custom: GetAvClass gets `AVHWFramesContext.av_class` value.
func (fc *AVHWFramesContext) GetAvClass() *AVClass {
	return (*AVClass)(fc.av_class)
}

// Custom: SetAvClass sets `AVHWFramesContext.av_class` value.
func (fc *AVHWFramesContext) SetAvClass(v *AVClass) {
	fc.av_class = (*C.struct_AVClass)(v)
}

// Custom: GetAvClassAddr gets `AVHWFramesContext.av_class` address.
func (fc *AVHWFramesContext) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&fc.av_class))
}

// Custom: GetInternal gets `AVHWFramesContext.internal` value.
func (fc *AVHWFramesContext) GetInternal() *AVHWFramesInternal {
	return (*AVHWFramesInternal)(fc.internal)
}

// Custom: SetInternal sets `AVHWFramesContext.internal` value.
func (fc *AVHWFramesContext) SetInternal(v *AVHWFramesInternal) {
	fc.internal = (*C.struct_AVHWFramesInternal)(v)
}

// Custom: GetInternalAddr gets `AVHWFramesContext.internal` address.
func (fc *AVHWFramesContext) GetInternalAddr() **AVHWFramesInternal {
	return (**AVHWFramesInternal)(unsafe.Pointer(&fc.internal))
}

// Custom: GetDeviceRef gets `AVHWFramesContext.device_ref` value.
func (fc *AVHWFramesContext) GetDeviceRef() *AVBufferRef {
	return (*AVBufferRef)(fc.device_ref)
}

// Custom: SetDeviceRef sets `AVHWFramesContext.device_ref` value.
func (fc *AVHWFramesContext) SetDeviceRef(v *AVBufferRef) {
	fc.device_ref = (*C.struct_AVBufferRef)(v)
}

// Custom: GetDeviceRefAddr gets `AVHWFramesContext.device_ref` address.
func (fc *AVHWFramesContext) GetDeviceRefAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&fc.device_ref))
}

// Custom: GetDeviceCtx gets `AVHWFramesContext.device_ctx` value.
func (fc *AVHWFramesContext) GetDeviceCtx() *AVHWDeviceContext {
	return (*AVHWDeviceContext)(fc.device_ctx)
}

// Custom: SetDeviceCtx sets `AVHWFramesContext.device_ctx` value.
func (fc *AVHWFramesContext) SetDeviceCtx(v *AVHWDeviceContext) {
	fc.device_ctx = (*C.struct_AVHWDeviceContext)(v)
}

// Custom: GetDeviceCtxAddr gets `AVHWFramesContext.device_ctx` address.
func (fc *AVHWFramesContext) GetDeviceCtxAddr() **AVHWDeviceContext {
	return (**AVHWDeviceContext)(unsafe.Pointer(&fc.device_ctx))
}

// Custom: GetHwctx gets `AVHWFramesContext.hwctx` value.
func (fc *AVHWFramesContext) GetHwctx() unsafe.Pointer {
	return (unsafe.Pointer)(fc.hwctx)
}

// Custom: SetHwctx sets `AVHWFramesContext.hwctx` value.
func (fc *AVHWFramesContext) SetHwctx(v CVoidPointer) {
	fc.hwctx = VoidPointer(v)
}

// Custom: GetHwctxAddr gets `AVHWFramesContext.hwctx` address.
func (fc *AVHWFramesContext) GetHwctxAddr() *unsafe.Pointer {
	return (*unsafe.Pointer)(&fc.hwctx)
}

// Custom: SetFree sets `AVHWFramesContext.free` value.
func (fc *AVHWFramesContext) SetFree(v AVHWFramesContextFreeFunc) {
	fc.free = (C.av_hw_frames_context_free_func)(v)
}

// Custom: GetUserOpaque gets `AVHWFramesContext.user_opaque` value.
func (fc *AVHWFramesContext) GetUserOpaque() unsafe.Pointer {
	return (unsafe.Pointer)(fc.user_opaque)
}

// Custom: SetUserOpaque sets `AVHWFramesContext.user_opaque` value.
func (fc *AVHWFramesContext) SetUserOpaque(v CVoidPointer) {
	fc.user_opaque = VoidPointer(v)
}

// Custom: GetUserOpaqueAddr gets `AVHWFramesContext.user_opaque` address.
func (fc *AVHWFramesContext) GetUserOpaqueAddr() *unsafe.Pointer {
	return (*unsafe.Pointer)(&fc.user_opaque)
}

// Custom: GetPool gets `AVHWFramesContext.pool` value.
func (fc *AVHWFramesContext) GetPool() *AVBufferPool {
	return (*AVBufferPool)(fc.pool)
}

// Custom: SetPool sets `AVHWFramesContext.pool` value.
func (fc *AVHWFramesContext) SetPool(v *AVBufferPool) {
	fc.pool = (*C.struct_AVBufferPool)(v)
}

// Custom: GetPoolAddr gets `AVHWFramesContext.pool` address.
func (fc *AVHWFramesContext) GetPoolAddr() **AVBufferPool {
	return (**AVBufferPool)(unsafe.Pointer(&fc.pool))
}

// Custom: GetInitialPoolSize gets `AVHWFramesContext.initial_pool_size` value.
func (fc *AVHWFramesContext) GetInitialPoolSize() int32 {
	return (int32)(fc.initial_pool_size)
}

// Custom: SetInitialPoolSize sets `AVHWFramesContext.initial_pool_size` value.
func (fc *AVHWFramesContext) SetInitialPoolSize(v int32) {
	fc.initial_pool_size = (C.int)(v)
}

// Custom: GetInitialPoolSizeAddr gets `AVHWFramesContext.initial_pool_size` address.
func (fc *AVHWFramesContext) GetInitialPoolSizeAddr() *int32 {
	return (*int32)(&fc.initial_pool_size)
}

// Custom: GetFormat gets `AVHWFramesContext.format` value.
func (fc *AVHWFramesContext) GetFormat() AVPixelFormat {
	return (AVPixelFormat)(fc.format)
}

// Custom: SetFormat sets `AVHWFramesContext.format` value.
func (fc *AVHWFramesContext) SetFormat(v AVPixelFormat) {
	fc.format = (C.enum_AVPixelFormat)(v)
}

// Custom: GetFormatAddr gets `AVHWFramesContext.format` address.
func (fc *AVHWFramesContext) GetFormatAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&fc.format)
}

// Custom: GetSwFormat gets `AVHWFramesContext.sw_format` value.
func (fc *AVHWFramesContext) GetSwFormat() AVPixelFormat {
	return (AVPixelFormat)(fc.sw_format)
}

// Custom: SetSwFormat sets `AVHWFramesContext.sw_format` value.
func (fc *AVHWFramesContext) SetSwFormat(v AVPixelFormat) {
	fc.sw_format = (C.enum_AVPixelFormat)(v)
}

// Custom: GetSwFormatAddr gets `AVHWFramesContext.sw_format` address.
func (fc *AVHWFramesContext) GetSwFormatAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&fc.sw_format)
}

// Custom: GetWidth gets `AVHWFramesContext.width` value.
func (fc *AVHWFramesContext) GetWidth() int32 {
	return (int32)(fc.width)
}

// Custom: SetWidth sets `AVHWFramesContext.width` value.
func (fc *AVHWFramesContext) SetWidth(v int32) {
	fc.width = (C.int)(v)
}

// Custom: GetWidthAddr gets `AVHWFramesContext.width` address.
func (fc *AVHWFramesContext) GetWidthAddr() *int32 {
	return (*int32)(&fc.width)
}

// Custom: GetHeight gets `AVHWFramesContext.height` value.
func (fc *AVHWFramesContext) GetHeight() int32 {
	return (int32)(fc.height)
}

// Custom: SetHeight sets `AVHWFramesContext.height` value.
func (fc *AVHWFramesContext) SetHeight(v int32) {
	fc.height = (C.int)(v)
}

// Custom: GetHeightAddr gets `AVHWFramesContext.height` address.
func (fc *AVHWFramesContext) GetHeightAddr() *int32 {
	return (*int32)(&fc.height)
}

// AvHWDeviceFindTypeByName looks up an AVHWDeviceType by name.
func AvHWDeviceFindTypeByName(name string) AVHWDeviceType {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (AVHWDeviceType)(C.av_hwdevice_find_type_by_name((*C.char)(namePtr)))
}

// AvHWDeviceGetTypeName gets the string name of an AVHWDeviceType.
func AvHWDeviceGetTypeName(_type AVHWDeviceType) string {
	return C.GoString(C.av_hwdevice_get_type_name((C.enum_AVHWDeviceType)(_type)))
}

// AvHWDeviceIterateTypes iterates over supported device types.
func AvHWDeviceIterateTypes(prev AVHWDeviceType) AVHWDeviceType {
	return (AVHWDeviceType)(C.av_hwdevice_iterate_types((C.enum_AVHWDeviceType)(prev)))
}

// AvHWDeviceCtxAlloc allocates an AVHWDeviceContext for a given hardware type.
func AvHWDeviceCtxAlloc(prev AVHWDeviceType) *AVBufferRef {
	return (*AVBufferRef)(C.av_hwdevice_ctx_alloc((C.enum_AVHWDeviceType)(prev)))
}

// AvHWDeviceCtxInit finalizes the device context before use.
func AvHWDeviceCtxInit(ref *AVBufferRef) int32 {
	return (int32)(C.av_hwdevice_ctx_init((*C.struct_AVBufferRef)(ref)))
}

// AvHWDeviceCtxCreate opens a device of the specified type and create an AVHWDeviceContext for it.
func AvHWDeviceCtxCreate(deviceCtx **AVBufferRef, _type AVHWDeviceType,
	device string, opts *AVDictionary, flags int32) int32 {
	devicePtr, deviceFunc := StringCasting(device)
	defer deviceFunc()
	return (int32)(C.av_hwdevice_ctx_create(
		(**C.struct_AVBufferRef)(unsafe.Pointer(deviceCtx)),
		(C.enum_AVHWDeviceType)(_type),
		(*C.char)(devicePtr),
		(*C.struct_AVDictionary)(opts),
		(C.int)(flags)))
}

// AvHWDeviceCtxCreateDerived creates a new device of the specified type from an existing device.
func AvHWDeviceCtxCreateDerived(dstCtx **AVBufferRef, _type AVHWDeviceType,
	srcCtx *AVBufferRef, flags int32) int32 {
	return (int32)(C.av_hwdevice_ctx_create_derived(
		(**C.struct_AVBufferRef)(unsafe.Pointer(dstCtx)),
		(C.enum_AVHWDeviceType)(_type),
		(*C.struct_AVBufferRef)(srcCtx),
		(C.int)(flags)))
}

// AvHWDeviceCtxCreateDerivedOpts creates a new device of the specified type from an existing device.
func AvHWDeviceCtxCreateDerivedOpts(dstCtx **AVBufferRef, _type AVHWDeviceType,
	srcCtx *AVBufferRef, options *AVDictionary, flags int32) int32 {
	return (int32)(C.av_hwdevice_ctx_create_derived_opts(
		(**C.struct_AVBufferRef)(unsafe.Pointer(dstCtx)),
		(C.enum_AVHWDeviceType)(_type),
		(*C.struct_AVBufferRef)(srcCtx),
		(*C.struct_AVDictionary)(options),
		(C.int)(flags)))
}

// AvHWFrameCtxAlloc allocates an AVHWFramesContext tied to a given device context.
func AvHWFrameCtxAlloc(deviceCtx *AVBufferRef) *AVBufferRef {
	return (*AVBufferRef)(C.av_hwframe_ctx_alloc((*C.struct_AVBufferRef)(deviceCtx)))
}

// AvHWFrameCtxInit finalizes the context before use.
func AvHWFrameCtxInit(ref *AVBufferRef) int32 {
	return (int32)(C.av_hwframe_ctx_init((*C.struct_AVBufferRef)(ref)))
}

// AvHWFrameGetBuffer allocates a new frame attached to the given AVHWFramesContext.
func AvHWFrameGetBuffer(hwframeCtx *AVBufferRef, frame *AVFrame, flags int32) int32 {
	return (int32)(C.av_hwframe_get_buffer((*C.struct_AVBufferRef)(hwframeCtx),
		(*C.struct_AVFrame)(frame), (C.int)(flags)))
}

// AvHWFrameTransferData copies data to or from a hw surface.
func AvHWFrameTransferData(dst, src *AVFrame, flags int32) int32 {
	return (int32)(C.av_hwframe_transfer_data(
		(*C.struct_AVFrame)(dst),
		(*C.struct_AVFrame)(src),
		(C.int)(flags)))
}

type AVHWFrameTransferDirection = C.enum_AVHWFrameTransferDirection

const (
	AV_HWFRAME_TRANSFER_DIRECTION_FROM = AVHWFrameTransferDirection(C.AV_HWFRAME_TRANSFER_DIRECTION_FROM)
	AV_HWFRAME_TRANSFER_DIRECTION_TO   = AVHWFrameTransferDirection(C.AV_HWFRAME_TRANSFER_DIRECTION_TO)
)

// AvHWFrameTransferGetFormats gets a list of possible source or target formats usable in
// AvHWFrameTransferData().
func AvHWFrameTransferGetFormats(hwframeCtx *AVBufferRef, dir AVHWFrameTransferDirection,
	formats **AVPixelFormat, flags int32) int32 {
	return (int32)(C.av_hwframe_transfer_get_formats((*C.struct_AVBufferRef)(hwframeCtx),
		(C.enum_AVHWFrameTransferDirection)(dir),
		(**C.enum_AVPixelFormat)(unsafe.Pointer(formats)),
		(C.int)(flags)))
}

type AVHWFramesConstraints C.struct_AVHWFramesConstraints

// Custom: GetValidHwFormats gets `AVHWFramesConstraints.valid_hw_formats` value.
func (fcs *AVHWFramesConstraints) GetValidHwFormats() []AVPixelFormat {
	return TruncSlice((*AVPixelFormat)(fcs.valid_hw_formats), func(pf AVPixelFormat) bool {
		return pf == AV_PIX_FMT_NONE
	})
}

// Custom: SetValidHwFormats sets `AVHWFramesConstraints.valid_hw_formats` value.
func (fcs *AVHWFramesConstraints) SetValidHwFormats(v *AVPixelFormat) {
	fcs.valid_hw_formats = (*C.enum_AVPixelFormat)(v)
}

// Custom: GetValidHwFormatsAddr gets `AVHWFramesConstraints.valid_hw_formats` address.
func (fcs *AVHWFramesConstraints) GetValidHwFormatsAddr() **AVPixelFormat {
	return (**AVPixelFormat)(&fcs.valid_hw_formats)
}

// Custom: GetValidSwFormats gets `AVHWFramesConstraints.valid_sw_formats` value.
func (fcs *AVHWFramesConstraints) GetValidSwFormats() []AVPixelFormat {
	return TruncSlice((*AVPixelFormat)(fcs.valid_sw_formats), func(pf AVPixelFormat) bool {
		return pf == AV_PIX_FMT_NONE
	})
}

// Custom: SetValidSwFormats sets `AVHWFramesConstraints.valid_sw_formats` value.
func (fcs *AVHWFramesConstraints) SetValidSwFormats(v *AVPixelFormat) {
	fcs.valid_sw_formats = (*C.enum_AVPixelFormat)(v)
}

// Custom: GetValidSwFormatsAddr gets `AVHWFramesConstraints.valid_sw_formats` address.
func (fcs *AVHWFramesConstraints) GetValidSwFormatsAddr() **AVPixelFormat {
	return (**AVPixelFormat)(&fcs.valid_sw_formats)
}

// Custom: GetMinWidth gets `AVHWFramesConstraints.min_width` value.
func (fcs *AVHWFramesConstraints) GetMinWidth() int32 {
	return (int32)(fcs.min_width)
}

// Custom: SetMinWidth sets `AVHWFramesConstraints.min_width` value.
func (fcs *AVHWFramesConstraints) SetMinWidth(v int32) {
	fcs.min_width = (C.int)(v)
}

// Custom: GetMinWidthAddr gets `AVHWFramesConstraints.min_width` address.
func (fcs *AVHWFramesConstraints) GetMinWidthAddr() *int32 {
	return (*int32)(&fcs.min_width)
}

// Custom: GetMinHeight gets `AVHWFramesConstraints.min_height` value.
func (fcs *AVHWFramesConstraints) GetMinHeight() int32 {
	return (int32)(fcs.min_height)
}

// Custom: SetMinHeight sets `AVHWFramesConstraints.min_height` value.
func (fcs *AVHWFramesConstraints) SetMinHeight(v int32) {
	fcs.min_height = (C.int)(v)
}

// Custom: GetMinHeightAddr gets `AVHWFramesConstraints.min_height` address.
func (fcs *AVHWFramesConstraints) GetMinHeightAddr() *int32 {
	return (*int32)(&fcs.min_height)
}

// Custom: GetMaxWidth gets `AVHWFramesConstraints.max_width` value.
func (fcs *AVHWFramesConstraints) GetMaxWidth() int32 {
	return (int32)(fcs.max_width)
}

// Custom: SetMaxWidth sets `AVHWFramesConstraints.max_width` value.
func (fcs *AVHWFramesConstraints) SetMaxWidth(v int32) {
	fcs.max_width = (C.int)(v)
}

// Custom: GetMaxWidthAddr gets `AVHWFramesConstraints.max_width` address.
func (fcs *AVHWFramesConstraints) GetMaxWidthAddr() *int32 {
	return (*int32)(&fcs.max_width)
}

// Custom: GetMaxHeight gets `AVHWFramesConstraints.max_height` value.
func (fcs *AVHWFramesConstraints) GetMaxHeight() int32 {
	return (int32)(fcs.max_height)
}

// Custom: SetMaxHeight sets `AVHWFramesConstraints.max_height` value.
func (fcs *AVHWFramesConstraints) SetMaxHeight(v int32) {
	fcs.max_height = (C.int)(v)
}

// Custom: GetMaxHeightAddr gets `AVHWFramesConstraints.max_height` address.
func (fcs *AVHWFramesConstraints) GetMaxHeightAddr() *int32 {
	return (*int32)(&fcs.max_height)
}

// AvHWDeviceHwconfigAlloc allocates a HW-specific configuration structure for a given HW device.
func AvHWDeviceHwconfigAlloc(deviceCtx *AVBufferRef) unsafe.Pointer {
	return C.av_hwdevice_hwconfig_alloc((*C.struct_AVBufferRef)(deviceCtx))
}

// AvHWDeviceGetHwframeConstraints
func AvHWDeviceGetHwframeConstraints(ref *AVBufferRef, hwconfig CVoidPointer) *AVHWFramesConstraints {
	return (*AVHWFramesConstraints)(C.av_hwdevice_get_hwframe_constraints(
		(*C.struct_AVBufferRef)(ref), VoidPointer(hwconfig)))
}

// AvHWFrameConstraintsFree frees an AVHWFrameConstraints structure.
func AvHWFrameConstraintsFree(constraints **AVHWFramesConstraints) {
	C.av_hwframe_constraints_free((**C.struct_AVHWFramesConstraints)(unsafe.Pointer(constraints)))
}

const (
	AV_HWFRAME_MAP_READ      = int32(C.AV_HWFRAME_MAP_READ)
	AV_HWFRAME_MAP_WRITE     = int32(C.AV_HWFRAME_MAP_WRITE)
	AV_HWFRAME_MAP_OVERWRITE = int32(C.AV_HWFRAME_MAP_OVERWRITE)
	AV_HWFRAME_MAP_DIRECT    = int32(C.AV_HWFRAME_MAP_DIRECT)
)

// AvHWFrameMap maps a hardware frame.
func AvHWFrameMap(dst, src *AVFrame, flags int32) int32 {
	return (int32)(C.av_hwframe_map(
		(*C.struct_AVFrame)(dst),
		(*C.struct_AVFrame)(src),
		(C.int)(flags)))
}

// AvHWFrameCtxCreateDerived creates and initialises an AVHWFramesContext as a mapping of another existing
// AvHWFramesContext on a different device.
func AvHWFrameCtxCreateDerived(derivedFrameCtx **AVBufferRef, format AVPixelFormat,
	derivedDeviceCtx, sourceFrameCtx *AVBufferRef, flags int32) int32 {
	return (int32)(C.av_hwframe_ctx_create_derived(
		(**C.struct_AVBufferRef)(unsafe.Pointer(derivedFrameCtx)),
		(C.enum_AVPixelFormat)(format),
		(*C.struct_AVBufferRef)(derivedDeviceCtx),
		(*C.struct_AVBufferRef)(sourceFrameCtx),
		(C.int)(flags)))
}
