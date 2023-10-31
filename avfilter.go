// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavfilter/avfilter.h>
*/
import "C"
import "unsafe"

// AvFilterVersion returns the LIBAVFILTER_VERSION_INT constant.
func AvFilterVersion() uint32 {
	return (uint32)(C.avfilter_version())
}

// AvFilterConfiguration returns the libavfilter build-time configuration.
func AvFilterConfiguration() string {
	return C.GoString(C.avfilter_configuration())
}

// AvFilterLicense returns the libavfilter license.
func AvFilterLicense() string {
	return C.GoString(C.avfilter_license())
}

// AVFilterPad
type AVFilterPad C.struct_AVFilterPad

// AVFilterFormats
type AVFilterFormats C.struct_AVFilterFormats

// AVFilterChannelLayouts
type AVFilterChannelLayouts C.struct_AVFilterChannelLayouts

// Deprecated: Use AvfilterFilterPadCount() instead.
//
// AvFilterPadCount gets the number of elements in a NULL-terminated array of AVFilterPads (e.g.
// AvFilter.inputs/outputs).
func AvFilterPadCount(pads *AVFilterPad) int32 {
	return (int32)(C.avfilter_pad_count((*C.struct_AVFilterPad)(pads)))
}

// AvFilterPadGetName gets the name of an AVFilterPad.
func AvFilterPadGetName(pads *AVFilterPad, padIdx int32) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(pads), (C.int)(padIdx)))
}

// AvFilterPadGetType gets the type of an AVFilterPad.
func AvFilterPadGetType(pads *AVFilterPad, padIdx int32) AVMediaType {
	return (AVMediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(pads), (C.int)(padIdx)))
}

const (
	AVFILTER_FLAG_DYNAMIC_INPUTS            = C.AVFILTER_FLAG_DYNAMIC_INPUTS
	AVFILTER_FLAG_DYNAMIC_OUTPUTS           = C.AVFILTER_FLAG_DYNAMIC_OUTPUTS
	AVFILTER_FLAG_SLICE_THREADS             = C.AVFILTER_FLAG_SLICE_THREADS
	AVFILTER_FLAG_METADATA_ONLY             = C.AVFILTER_FLAG_METADATA_ONLY
	AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC  = C.AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC
	AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL = C.AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL
	AVFILTER_FLAG_SUPPORT_TIMELINE          = C.AVFILTER_FLAG_SUPPORT_TIMELINE
)

// Filter definition. This defines the pads a filter contains, and all the
// callback functions used to interact with the filter.
type AVFilter C.struct_AVFilter

// GetName gets `AVFilter.name` value.
func (flt *AVFilter) GetName() string {
	return C.GoString(flt.name)
}

// GetDescription gets `AVFilter.description` value.
func (flt *AVFilter) GetDescription() string {
	return C.GoString(flt.description)
}

// GetInputs gets `AVFilter.inputs` value.
func (flt *AVFilter) GetInputs() *AVFilterPad {
	return (*AVFilterPad)(flt.inputs)
}

// SetInputs sets `AVFilter.inputs` value.
func (flt *AVFilter) SetInputs(v *AVFilterPad) {
	flt.inputs = (*C.struct_AVFilterPad)(v)
}

// GetInputsAddr gets `AVFilter.inputs` address.
func (flt *AVFilter) GetInputsAddr() **AVFilterPad {
	return (**AVFilterPad)(unsafe.Pointer(&flt.inputs))
}

// GetOutputs gets `AVFilter.outputs` value.
func (flt *AVFilter) GetOutputs() *AVFilterPad {
	return (*AVFilterPad)(flt.outputs)
}

// SetOutputs sets `AVFilter.outputs` value.
func (flt *AVFilter) SetOutputs(v *AVFilterPad) {
	flt.outputs = (*C.struct_AVFilterPad)(v)
}

// GetOutputsAddr gets `AVFilter.outputs` address.
func (flt *AVFilter) GetOutputsAddr() **AVFilterPad {
	return (**AVFilterPad)(unsafe.Pointer(&flt.outputs))
}

// GetPrivClass gets `AVFilter.priv_class` value.
func (flt *AVFilter) GetPrivClass() *AVClass {
	return (*AVClass)(flt.priv_class)
}

// SetPrivClass sets `AVFilter.priv_class` value.
func (flt *AVFilter) SetPrivClass(v *AVClass) {
	flt.priv_class = (*C.struct_AVClass)(v)
}

// GetPrivClassAddr gets `AVFilter.priv_class` address.
func (flt *AVFilter) GetPrivClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&flt.priv_class))
}

// GetFlags gets `AVFilter.flags` value.
func (flt *AVFilter) GetFlags() int32 {
	return (int32)(flt.flags)
}

// SetFlags sets `AVFilter.flags` value.
func (flt *AVFilter) SetFlags(v int32) {
	flt.flags = (C.int)(v)
}

// GetFlagsAddr gets `AVFilter.flags` address.
func (flt *AVFilter) GetFlagsAddr() *int32 {
	return (*int32)(&flt.flags)
}

// AvfilterFilterPadCount
func AvfilterFilterPadCount(filter *AVFilter, isOutput int32) uint32 {
	return (uint32)(C.avfilter_filter_pad_count((*C.struct_AVFilter)(filter), (C.int)(isOutput)))
}

const (
	AVFILTER_THREAD_SLICE = C.AVFILTER_THREAD_SLICE
)

// AVFilterInternal
type AVFilterInternal C.struct_AVFilterInternal

// AVFilterContext
type AVFilterContext C.struct_AVFilterContext

// GetAvClass gets `AVFilterContext.av_class` value.
func (fltc *AVFilterContext) GetAvClass() *AVClass {
	return (*AVClass)(fltc.av_class)
}

// SetAvClass sets `AVFilterContext.av_class` value.
func (fltc *AVFilterContext) SetAvClass(v *AVClass) {
	fltc.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVFilterContext.av_class` address.
func (fltc *AVFilterContext) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&fltc.av_class))
}

// GetFilter gets `AVFilterContext.filter` value.
func (fltc *AVFilterContext) GetFilter() *AVFilter {
	return (*AVFilter)(fltc.filter)
}

// SetFilter sets `AVFilterContext.filter` value.
func (fltc *AVFilterContext) SetFilter(v *AVFilter) {
	fltc.filter = (*C.struct_AVFilter)(v)
}

// GetFilterAddr gets `AVFilterContext.filter` address.
func (fltc *AVFilterContext) GetFilterAddr() **AVFilter {
	return (**AVFilter)(unsafe.Pointer(&fltc.filter))
}

// GetName gets `AVFilterContext.name` value.
func (fltc *AVFilterContext) GetName() string {
	return C.GoString(fltc.name)
}

// SetName sets `AVFilterContext.name` value.
func (fltc *AVFilterContext) SetName(v string) {
	FreePointer(fltc.name)
	fltc.name, _ = StringCasting(v)
}

// GetInputPads gets `AVFilterContext.input_pads` value.
func (fltc *AVFilterContext) GetInputPads() *AVFilterPad {
	return (*AVFilterPad)(fltc.input_pads)
}

// SetInputPads sets `AVFilterContext.input_pads` value.
func (fltc *AVFilterContext) SetInputPads(v *AVFilterPad) {
	fltc.input_pads = (*C.struct_AVFilterPad)(v)
}

// GetInputPadsAddr gets `AVFilterContext.input_pads` address.
func (fltc *AVFilterContext) GetInputPadsAddr() **AVFilterPad {
	return (**AVFilterPad)(unsafe.Pointer(&fltc.input_pads))
}

// GetInputs gets `AVFilterContext.inputs` value.
func (fltc *AVFilterContext) GetInputs() []*AVFilterLink {
	if fltc.inputs == nil {
		return nil
	}
	return unsafe.Slice((**AVFilterLink)(unsafe.Pointer(fltc.inputs)), fltc.nb_inputs)
}

// SetInputs sets `AVFilterContext.inputs` value.
func (fltc *AVFilterContext) SetInputs(v **AVFilterLink) {
	fltc.inputs = (**C.struct_AVFilterLink)(unsafe.Pointer(v))
}

// GetInputsAddr gets `AVFilterContext.inputs` address.
func (fltc *AVFilterContext) GetInputsAddr() ***AVFilterLink {
	return (***AVFilterLink)(unsafe.Pointer(&fltc.inputs))
}

// GetNbInputs gets `AVFilterContext.nb_inputs` value.
func (fltc *AVFilterContext) GetNbInputs() uint32 {
	return (uint32)(fltc.nb_inputs)
}

// SetNbInputs sets `AVFilterContext.nb_inputs` value.
func (fltc *AVFilterContext) SetNbInputs(v uint32) {
	fltc.nb_inputs = (C.uint)(v)
}

// GetNbInputsAddr gets `AVFilterContext.nb_inputs` address.
func (fltc *AVFilterContext) GetNbInputsAddr() *uint32 {
	return (*uint32)(&fltc.nb_inputs)
}

// GetOutputPads gets `AVFilterContext.output_pads` value.
func (fltc *AVFilterContext) GetOutputPads() *AVFilterPad {
	return (*AVFilterPad)(unsafe.Pointer(fltc.output_pads))
}

// SetOutputPads sets `AVFilterContext.output_pads` value.
func (fltc *AVFilterContext) SetOutputPads(v *AVFilterPad) {
	fltc.output_pads = (*C.struct_AVFilterPad)(v)
}

// GetOutputPadsAddr gets `AVFilterContext.output_pads` address.
func (fltc *AVFilterContext) GetOutputPadsAddr() **AVFilterPad {
	return (**AVFilterPad)(unsafe.Pointer(&fltc.output_pads))
}

// GetOutputs gets `AVFilterContext.outputs` value.
func (fltc *AVFilterContext) GetOutputs() []*AVFilterLink {
	if fltc.outputs == nil {
		return nil
	}
	return unsafe.Slice((**AVFilterLink)(unsafe.Pointer(fltc.outputs)), fltc.nb_outputs)
}

// SetOutputs sets `AVFilterContext.outputs` value.
func (fltc *AVFilterContext) SetOutputs(v **AVFilterLink) {
	fltc.outputs = (**C.struct_AVFilterLink)(unsafe.Pointer(v))
}

// GetOutputsAddr gets `AVFilterContext.outputs` address.
func (fltc *AVFilterContext) GetOutputsAddr() ***AVFilterLink {
	return (***AVFilterLink)(unsafe.Pointer(&fltc.outputs))
}

// GetNbOutputs gets `AVFilterContext.nb_outputs` value.
func (fltc *AVFilterContext) GetNbOutputs() uint32 {
	return (uint32)(fltc.nb_outputs)
}

// SetNbOutputs sets `AVFilterContext.nb_outputs` value.
func (fltc *AVFilterContext) SetNbOutputs(v uint32) {
	fltc.nb_outputs = (C.uint)(v)
}

// GetNbOutputsAddr gets `AVFilterContext.nb_outputs` address.
func (fltc *AVFilterContext) GetNbOutputsAddr() *uint32 {
	return (*uint32)(&fltc.nb_outputs)
}

// GetPriv gets `AVFilterContext.priv` value.
func (fltc *AVFilterContext) GetPriv() unsafe.Pointer {
	return fltc.priv
}

// SetPriv sets `AVFilterContext.priv` value.
func (fltc *AVFilterContext) SetPriv(v CVoidPointer) {
	fltc.priv = VoidPointer(v)
}

// GetPrivAddr gets `AVFilterContext.priv` address.
func (fltc *AVFilterContext) GetPrivAddr() *unsafe.Pointer {
	return &fltc.priv
}

// GetGraph gets `AVFilterContext.graph` value.
func (fltc *AVFilterContext) GetGraph() *AVFilterGraph {
	return (*AVFilterGraph)(fltc.graph)
}

// SetGraph sets `AVFilterContext.graph` value.
func (fltc *AVFilterContext) SetGraph(v *AVFilterGraph) {
	fltc.graph = (*C.struct_AVFilterGraph)(v)
}

// GetGraphAddr gets `AVFilterContext.graph` address.
func (fltc *AVFilterContext) GetGraphAddr() **AVFilterGraph {
	return (**AVFilterGraph)(unsafe.Pointer(&fltc.graph))
}

// GetThreadType gets `AVFilterContext.threadtype` value.
func (fltc *AVFilterContext) GetThreadType() int32 {
	return (int32)(fltc.thread_type)
}

// SetThreadType sets `AVFilterContext.threadtype` value.
func (fltc *AVFilterContext) SetThreadType(v int32) {
	fltc.thread_type = (C.int)(v)
}

// GetThreadTypeAddr gets `AVFilterContext.threadtype` address.
func (fltc *AVFilterContext) GetThreadTypeAddr() *int32 {
	return (*int32)(&fltc.thread_type)
}

// GetInternal gets `AVFilterContext.internal` value.
func (fltc *AVFilterContext) GetInternal() *AVFilterInternal {
	return (*AVFilterInternal)(fltc.internal)
}

// SetInternal sets `AVFilterContext.internal` value.
func (fltc *AVFilterContext) SetInternal(v *AVFilterInternal) {
	fltc.internal = (*C.struct_AVFilterInternal)(v)
}

// GetInternalAddr gets `AVFilterContext.internal` address.
func (fltc *AVFilterContext) GetInternalAddr() **AVFilterInternal {
	return (**AVFilterInternal)(unsafe.Pointer(&fltc.internal))
}

// AVFilterCommand
type AVFilterCommand C.struct_AVFilterCommand

// GetCommandQueue gets `AVFilterContext.command_queue` value.
func (fltc *AVFilterContext) GetCommandQueue() *AVFilterCommand {
	return (*AVFilterCommand)(fltc.command_queue)
}

// SetCommandQueue sets `AVFilterContext.command_queue` value.
func (fltc *AVFilterContext) SetCommandQueue(v *AVFilterCommand) {
	fltc.command_queue = (*C.struct_AVFilterCommand)(v)
}

// GetCommandQueueAddr gets `AVFilterContext.command_queue` address.
func (fltc *AVFilterContext) GetCommandQueueAddr() **AVFilterCommand {
	return (**AVFilterCommand)(unsafe.Pointer(&fltc.command_queue))
}

// GetEnableStr gets `AVFilterContext.enable_str` value.
func (fltc *AVFilterContext) GetEnableStr() string {
	return C.GoString(fltc.enable_str)
}

// GetEnable gets `AVFilterContext.enable` value.
func (fltc *AVFilterContext) GetEnable() unsafe.Pointer {
	return fltc.enable
}

// SetEnable sets `AVFilterContext.enable` value.
func (fltc *AVFilterContext) SetEnable(v CVoidPointer) {
	fltc.enable = VoidPointer(v)
}

// GetEnableAddr gets `AVFilterContext.enable` address.
func (fltc *AVFilterContext) GetEnableAddr() *unsafe.Pointer {
	return &fltc.enable
}

// GetVarValues gets `AVFilterContext.var_values` value.
func (fltc *AVFilterContext) GetVarValues() *float64 {
	return (*float64)(fltc.var_values)
}

// SetVarValues sets `AVFilterContext.var_values` value.
func (fltc *AVFilterContext) SetVarValues(v *float64) {
	fltc.var_values = (*C.double)(v)
}

// GetVarValuesAddr gets `AVFilterContext.var_values` address.
func (fltc *AVFilterContext) GetVarValuesAddr() **float64 {
	return (**float64)(unsafe.Pointer(&fltc.var_values))
}

// GetIsDisabled gets `AVFilterContext.is_disabled` value.
func (fltc *AVFilterContext) GetIsDisabled() int32 {
	return (int32)(fltc.is_disabled)
}

// SetIsDisabled sets `AVFilterContext.is_disabled` value.
func (fltc *AVFilterContext) SetIsDisabled(v int32) {
	fltc.is_disabled = (C.int)(v)
}

// GetIsDisabledAddr gets `AVFilterContext.is_disabled` address.
func (fltc *AVFilterContext) GetIsDisabledAddr() *int32 {
	return (*int32)(&fltc.is_disabled)
}

// GetHwDeviceCtx gets `AVFilterContext.hw_device_ctx` value.
func (fltc *AVFilterContext) GetHwDeviceCtx() *AVBufferRef {
	return (*AVBufferRef)(fltc.hw_device_ctx)
}

// SetHwDeviceCtx sets `AVFilterContext.hw_device_ctx` value.
func (fltc *AVFilterContext) SetHwDeviceCtx(v *AVBufferRef) {
	fltc.hw_device_ctx = (*C.struct_AVBufferRef)(v)
}

// GetHwDeviceCtxAddr gets `AVFilterContext.hw_device_ctx` address.
func (fltc *AVFilterContext) GetHwDeviceCtxAddr() **AVBufferRef {
	return (**AVBufferRef)(unsafe.Pointer(&fltc.hw_device_ctx))
}

// GetNbThreads gets `AVFilterContext.nb_threads` value.
func (fltc *AVFilterContext) GetNbThreads() int32 {
	return (int32)(fltc.nb_threads)
}

// SetNbThreads sets `AVFilterContext.nb_threads` value.
func (fltc *AVFilterContext) SetNbThreads(v int32) {
	fltc.nb_threads = (C.int)(v)
}

// GetNbThreadsAddr gets `AVFilterContext.nb_threads` address.
func (fltc *AVFilterContext) GetNbThreadsAddr() *int32 {
	return (*int32)(&fltc.nb_threads)
}

// GetReady gets `AVFilterContext.ready` value.
func (fltc *AVFilterContext) GetReady() uint32 {
	return (uint32)(fltc.ready)
}

// SetReady sets `AVFilterContext.ready` value.
func (fltc *AVFilterContext) SetReady(v uint32) {
	fltc.ready = (C.uint)(v)
}

// GetReadyAddr gets `AVFilterContext.ready` address.
func (fltc *AVFilterContext) GetReadyAddr() *uint32 {
	return (*uint32)(&fltc.ready)
}

// GetExtraHwFrames gets `AVFilterContext.extra_hw_frames` value.
func (fltc *AVFilterContext) GetExtraHwFrames() int32 {
	return (int32)(fltc.extra_hw_frames)
}

// SetExtraHwFrames sets `AVFilterContext.extra_hw_frames` value.
func (fltc *AVFilterContext) SetExtraHwFrames(v int32) {
	fltc.extra_hw_frames = (C.int)(v)
}

// GetExtraHwFramesAddr gets `AVFilterContext.extra_hw_frames` address.
func (fltc *AVFilterContext) GetExtraHwFramesAddr() *int32 {
	return (*int32)(&fltc.extra_hw_frames)
}

// AVFilterFormatsConfig
type AVFilterFormatsConfig C.struct_AVFilterFormatsConfig

// GetFormats gets `AVFilterFormatsConfig.formats` value.
func (fltf *AVFilterFormatsConfig) GetFormats() *AVFilterFormats {
	return (*AVFilterFormats)(fltf.formats)
}

// SetFormats sets `AVFilterFormatsConfig.formats` value.
func (fltf *AVFilterFormatsConfig) SetFormats(v *AVFilterFormats) {
	fltf.formats = (*C.struct_AVFilterFormats)(v)
}

// GetFormatsAddr gets `AVFilterFormatsConfig.formats` address.
func (fltf *AVFilterFormatsConfig) GetFormatsAddr() **AVFilterFormats {
	return (**AVFilterFormats)(unsafe.Pointer(&fltf.formats))
}

// GetSamplerates gets `AVFilterFormatsConfig.samplerates` value.
func (fltf *AVFilterFormatsConfig) GetSamplerates() *AVFilterFormats {
	return (*AVFilterFormats)(fltf.samplerates)
}

// SetSamplerates sets `AVFilterFormatsConfig.samplerates` value.
func (fltf *AVFilterFormatsConfig) SetSamplerates(v *AVFilterFormats) {
	fltf.samplerates = (*C.struct_AVFilterFormats)(v)
}

// GetSampleratesAddr gets `AVFilterFormatsConfig.samplerates` address.
func (fltf *AVFilterFormatsConfig) GetSampleratesAddr() **AVFilterFormats {
	return (**AVFilterFormats)(unsafe.Pointer(&fltf.samplerates))
}

// GetChannelLayouts gets `AVFilterFormatsConfig.channel_layouts` value.
func (fltf *AVFilterFormatsConfig) GetChannelLayouts() *AVFilterChannelLayouts {
	return (*AVFilterChannelLayouts)(fltf.channel_layouts)
}

// SetChannelLayouts sets `AVFilterFormatsConfig.channel_layouts` value.
func (fltf *AVFilterFormatsConfig) SetChannelLayouts(v *AVFilterChannelLayouts) {
	fltf.channel_layouts = (*C.struct_AVFilterChannelLayouts)(v)
}

// GetChannelLayoutsAddr gets `AVFilterFormatsConfig.channel_layouts` address.
func (fltf *AVFilterFormatsConfig) GetChannelLayoutsAddr() **AVFilterChannelLayouts {
	return (**AVFilterChannelLayouts)(unsafe.Pointer(&fltf.channel_layouts))
}

// AVFilterLink
type AVFilterLink C.struct_AVFilterLink

// GetSrc gets `AVFilterLink.src` value.
func (fltl *AVFilterLink) GetSrc() *AVFilterContext {
	return (*AVFilterContext)(fltl.src)
}

// SetSrc sets `AVFilterLink.src` value.
func (fltl *AVFilterLink) SetSrc(v *AVFilterContext) {
	fltl.src = (*C.struct_AVFilterContext)(v)
}

// GetSrcAddr gets `AVFilterLink.src` address.
func (fltl *AVFilterLink) GetSrcAddr() **AVFilterContext {
	return (**AVFilterContext)(unsafe.Pointer(&fltl.src))
}

// GetSrcpad gets `AVFilterLink.srcpad` value.
func (fltl *AVFilterLink) GetSrcpad() *AVFilterPad {
	return (*AVFilterPad)(fltl.srcpad)
}

// SetSrcpad sets `AVFilterLink.srcpad` value.
func (fltl *AVFilterLink) SetSrcpad(v *AVFilterPad) {
	fltl.srcpad = (*C.struct_AVFilterPad)(v)
}

// GetSrcpadAddr gets `AVFilterLink.srcpad` address.
func (fltl *AVFilterLink) GetSrcpadAddr() **AVFilterPad {
	return (**AVFilterPad)(unsafe.Pointer(&fltl.srcpad))
}

// GetDst gets `AVFilterLink.dst` value.
func (fltl *AVFilterLink) GetDst() *AVFilterContext {
	return (*AVFilterContext)(fltl.dst)
}

// SetDst sets `AVFilterLink.dst` value.
func (fltl *AVFilterLink) SetDst(v *AVFilterContext) {
	fltl.dst = (*C.struct_AVFilterContext)(v)
}

// GetDstAddr gets `AVFilterLink.dst` address.
func (fltl *AVFilterLink) GetDstAddr() **AVFilterContext {
	return (**AVFilterContext)(unsafe.Pointer(&fltl.dst))
}

// GetDstpad gets `AVFilterLink.dstpad` value.
func (fltl *AVFilterLink) GetDstpad() *AVFilterPad {
	return (*AVFilterPad)(fltl.dstpad)
}

// SetDstpad sets `AVFilterLink.dstpad` value.
func (fltl *AVFilterLink) SetDstpad(v *AVFilterPad) {
	fltl.dstpad = (*C.struct_AVFilterPad)(v)
}

// GetDstpadAddr gets `AVFilterLink.dstpad` address.
func (fltl *AVFilterLink) GetDstpadAddr() **AVFilterPad {
	return (**AVFilterPad)(unsafe.Pointer(&fltl.dstpad))
}

// GetType gets `AVFilterLink.type` value.
func (fltl *AVFilterLink) GetType() AVMediaType {
	return (AVMediaType)(fltl._type)
}

// SetType sets `AVFilterLink.type` value.
func (fltl *AVFilterLink) SetType(v AVMediaType) {
	fltl._type = (C.enum_AVMediaType)(v)
}

// GetTypeAddr gets `AVFilterLink.type` address.
func (fltl *AVFilterLink) GetTypeAddr() *AVMediaType {
	return (*AVMediaType)(&fltl._type)
}

// GetW gets `AVFilterLink.w` value.
func (fltl *AVFilterLink) GetW() int32 {
	return (int32)(fltl.w)
}

// SetW sets `AVFilterLink.w` value.
func (fltl *AVFilterLink) SetW(v int32) {
	fltl.w = (C.int)(v)
}

// GetWAddr gets `AVFilterLink.w` address.
func (fltl *AVFilterLink) GetWAddr() *int32 {
	return (*int32)(&fltl.w)
}

// GetH gets `AVFilterLink.h` value.
func (fltl *AVFilterLink) GetH() int32 {
	return (int32)(fltl.h)
}

// SetH sets `AVFilterLink.h` value.
func (fltl *AVFilterLink) SetH(v int32) {
	fltl.h = (C.int)(v)
}

// GetHAddr gets `AVFilterLink.h` address.
func (fltl *AVFilterLink) GetHAddr() *int32 {
	return (*int32)(&fltl.h)
}

// GetSampleAspectRatio gets `AVFilterLink.sample_aspect_ratio` value.
func (fltl *AVFilterLink) GetSampleAspectRatio() AVRational {
	return (AVRational)(fltl.sample_aspect_ratio)
}

// SetSampleAspectRatio sets `AVFilterLink.sample_aspect_ratio` value.
func (fltl *AVFilterLink) SetSampleAspectRatio(v AVRational) {
	fltl.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// GetSampleAspectRatioAddr gets `AVFilterLink.sample_aspect_ratio` address.
func (fltl *AVFilterLink) GetSampleAspectRatioAddr() *AVRational {
	return (*AVRational)(&fltl.sample_aspect_ratio)
}

// GetChannelLayout gets `AVFilterLink.channel_layout` value.
func (fltl *AVFilterLink) GetChannelLayout() uint64 {
	return (uint64)(fltl.channel_layout)
}

// SetChannelLayout sets `AVFilterLink.channel_layout` value.
func (fltl *AVFilterLink) SetChannelLayout(v uint64) {
	fltl.channel_layout = (C.uint64_t)(v)
}

// GetChannelLayoutAddr gets `AVFilterLink.channel_layout` address.
func (fltl *AVFilterLink) GetChannelLayoutAddr() *uint64 {
	return (*uint64)(&fltl.channel_layout)
}

// GetSampleRate gets `AVFilterLink.sample_rate` value.
func (fltl *AVFilterLink) GetSampleRate() int32 {
	return (int32)(fltl.sample_rate)
}

// SetSampleRate sets `AVFilterLink.sample_rate` value.
func (fltl *AVFilterLink) SetSampleRate(v int32) {
	fltl.sample_rate = (C.int)(v)
}

// GetSampleRateAddr gets `AVFilterLink.sample_rate` address.
func (fltl *AVFilterLink) GetSampleRateAddr() *int32 {
	return (*int32)(&fltl.sample_rate)
}

// GetFormat gets `AVFilterLink.format` value.
func (fltl *AVFilterLink) GetFormat() int32 {
	return (int32)(fltl.format)
}

// SetFormat sets `AVFilterLink.format` value.
func (fltl *AVFilterLink) SetFormat(v int32) {
	fltl.format = (C.int)(v)
}

// GetFormatAddr gets `AVFilterLink.format` address.
func (fltl *AVFilterLink) GetFormatAddr() *int32 {
	return (*int32)(&fltl.format)
}

// GetTimeBase gets `AVFilterLink.time_base` value.
func (fltl *AVFilterLink) GetTimeBase() AVRational {
	return (AVRational)(fltl.time_base)
}

// SetTimeBase sets `AVFilterLink.time_base` value.
func (fltl *AVFilterLink) SetTimeBase(v AVRational) {
	fltl.time_base = (C.struct_AVRational)(v)
}

// GetTimeBaseAddr gets `AVFilterLink.time_base` address.
func (fltl *AVFilterLink) GetTimeBaseAddr() *AVRational {
	return (*AVRational)(&fltl.time_base)
}

// AvFilterLink links two filters together.
func AvFilterLink(src *AVFilterContext, srcpad uint32,
	dst *AVFilterContext, dstpad uint32) int32 {
	return (int32)(C.avfilter_link((*C.struct_AVFilterContext)(src), (C.uint)(srcpad),
		(*C.struct_AVFilterContext)(dst), (C.uint)(dstpad)))
}

// AvFilterLinkFree frees the link in *link, and set its pointer to NULL.
func AvFilterLinkFree(link **AVFilterLink) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(link)))
}

// AvFilterConfigLinks negotiates the media format, dimensions, etc of all inputs to a filter.
func AvFilterConfigLinks(filter *AVFilterContext) int32 {
	return (int32)(C.avfilter_config_links((*C.struct_AVFilterContext)(filter)))
}

const (
	AVFILTER_CMD_FLAG_ONE  = C.AVFILTER_CMD_FLAG_ONE
	AVFILTER_CMD_FLAG_FAST = C.AVFILTER_CMD_FLAG_FAST
)

// AvFilterProcessCommand makes the filter instance process a command.
// It is recommended to use AVFilterGraphSendCommand().
func AvFilterProcessCommand(filter *AVFilterContext, cmd, arg string, resLen, flags int32) (res string, ret int32) {
	cmdPtr, cmdFunc := StringCasting(cmd)
	defer cmdFunc()
	argPtr, argFunc := StringCasting(arg)
	defer argFunc()
	resBuf := make([]C.char, resLen)
	ret = (int32)(C.avfilter_process_command((*C.struct_AVFilterContext)(filter),
		(*C.char)(cmdPtr), (*C.char)(argPtr), (*C.char)(&resBuf[0]), (C.int)(resLen), (C.int)(flags)))
	return C.GoString(&resBuf[0]), ret
}

// AvFilterIterate iterates over all registered filters.
func AvFilterIterate(opaque CVoidPointerPointer) *AVFilter {
	return (*AVFilter)(C.av_filter_iterate(VoidPointerPointer(opaque)))
}

// AvFilterGetByName gets a filter definition matching the given name.
func AvFilterGetByName(name string) *AVFilter {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AVFilter)(C.avfilter_get_by_name((*C.char)(namePtr)))
}

// AvFilterInitStr initializes a filter with the supplied parameters.
func AvFilterInitStr(ctx *AVFilterContext, args string) int32 {
	argsPtr, argsFunc := StringCasting(args)
	defer argsFunc()
	return (int32)(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx),
		(*C.char)(argsPtr)))
}

// AvFilterInitDict initialize a filter with the supplied dictionary of options.
func AvFilterInitDict(ctx *AVFilterContext, options **AVDictionary) int32 {
	return (int32)(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvFilterFree frees a filter context. This will also remove the filter from its
// filtergraph's list of filters.
func AvFilterFree(ctx *AVFilterContext) {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

// AvFilterInsertFilter inserts a filter in the middle of an existing link.
func AvFilterInsertFilter(ctx *AVFilterContext, link *AVFilterLink,
	filtSrcpadIdx, filtDstpadIdx uint32) int32 {
	return (int32)(C.avfilter_insert_filter(
		(*C.struct_AVFilterLink)(link),
		(*C.struct_AVFilterContext)(ctx),
		(C.uint)(filtSrcpadIdx), (C.uint)(filtDstpadIdx)))
}

// AvFilterGetClass returns AVClass for AVFilterContext.
func AvFilterGetClass() *AVClass {
	return (*AVClass)(C.avfilter_get_class())
}

// AVFilterGraphInternal
type AVFilterGraphInternal C.struct_AVFilterGraphInternal

// typedef int (avfilter_action_func)(AVFilterContext *ctx, void *arg, int jobnr, int nb_jobs);
type AVFilterActionFunc = C.avfilter_action_func

// typedef int (avfilter_execute_func)(AVFilterContext *ctx, avfilter_action_func *func,
// void *arg, int *ret, int nb_jobs);
type AVFilterExecuteFunc = C.avfilter_execute_func

// AVFilterGraph
type AVFilterGraph C.struct_AVFilterGraph

// GetAvClass gets `AVFilterGraph.av_class` value.
func (fltg *AVFilterGraph) GetAvClass() *AVClass {
	return (*AVClass)(fltg.av_class)
}

// SetAvClass sets `AVFilterGraph.av_class` value.
func (fltg *AVFilterGraph) SetAvClass(v *AVClass) {
	fltg.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVFilterGraph.av_class` address.
func (fltg *AVFilterGraph) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&fltg.av_class))
}

// GetFilters gets `AVFilterGraph.filters` value.
func (fltg *AVFilterGraph) GetFilters() []*AVFilterContext {
	if fltg.filters == nil {
		return nil
	}
	return unsafe.Slice((**AVFilterContext)(unsafe.Pointer(fltg.filters)), fltg.nb_filters)
}

// SetFilters sets `AVFilterGraph.filters` value.
func (fltg *AVFilterGraph) SetFilters(v **AVFilterContext) {
	fltg.filters = (**C.struct_AVFilterContext)(unsafe.Pointer(v))
}

// GetFiltersAddr gets `AVFilterGraph.filters` address.
func (fltg *AVFilterGraph) GetFiltersAddr() ***AVFilterContext {
	return (***AVFilterContext)(unsafe.Pointer(&fltg.filters))
}

// GetNbFilters gets `AVFilterGraph.nb_filters` value.
func (fltg *AVFilterGraph) GetNbFilters() uint32 {
	return (uint32)(fltg.nb_filters)
}

// SetNbFilters sets `AVFilterGraph.nb_filters` value.
func (fltg *AVFilterGraph) SetNbFilters(v uint32) {
	fltg.nb_filters = (C.uint)(v)
}

// GetNbFiltersAddr gets `AVFilterGraph.nb_filters` address.
func (fltg *AVFilterGraph) GetNbFiltersAddr() *uint32 {
	return (*uint32)(&fltg.nb_filters)
}

// GetScaleSwsOpts gets `AVFilterGraph.scale_sws_opts` value.
func (fltg *AVFilterGraph) GetScaleSwsOpts() string {
	return C.GoString(fltg.scale_sws_opts)
}

// GetThreadType gets `AVFilterGraph.threadtype` value.
func (fltg *AVFilterGraph) GetThreadType() int32 {
	return (int32)(fltg.thread_type)
}

// SetThreadType sets `AVFilterGraph.threadtype` value.
func (fltg *AVFilterGraph) SetThreadType(v int32) {
	fltg.thread_type = (C.int)(v)
}

// GetThreadTypeAddr gets `AVFilterGraph.threadtype` address.
func (fltg *AVFilterGraph) GetThreadTypeAddr() *int32 {
	return (*int32)(&fltg.thread_type)
}

// GetNbThreads gets `AVFilterGraph.nb_threads` value.
func (fltg *AVFilterGraph) GetNbThreads() int32 {
	return (int32)(fltg.nb_threads)
}

// SetNbThreads sets `AVFilterGraph.nb_threads` value.
func (fltg *AVFilterGraph) SetNbThreads(v int32) {
	fltg.nb_threads = (C.int)(v)
}

// GetNbThreadsAddr gets `AVFilterGraph.nb_threads` address.
func (fltg *AVFilterGraph) GetNbThreadsAddr() *int32 {
	return (*int32)(&fltg.nb_threads)
}

// GetInternal gets `AVFilterGraph.internal` value.
func (fltg *AVFilterGraph) GetInternal() *AVFilterGraphInternal {
	return (*AVFilterGraphInternal)(fltg.internal)
}

// SetInternal sets `AVFilterGraph.internal` value.
func (fltg *AVFilterGraph) SetInternal(v *AVFilterGraphInternal) {
	fltg.internal = (*C.struct_AVFilterGraphInternal)(v)
}

// GetInternalAddr gets `AVFilterGraph.internal` address.
func (fltg *AVFilterGraph) GetInternalAddr() **AVFilterGraphInternal {
	return (**AVFilterGraphInternal)(unsafe.Pointer(&fltg.internal))
}

// GetOpaque gets `AVFilterGraph.opaque` value.
func (fltg *AVFilterGraph) GetOpaque() unsafe.Pointer {
	return fltg.opaque
}

// SetOpaque sets `AVFilterGraph.opaque` value.
func (fltg *AVFilterGraph) SetOpaque(v CVoidPointer) {
	fltg.opaque = VoidPointer(v)
}

// GetOpaqueAddr gets `AVFilterGraph.opaque` address.
func (fltg *AVFilterGraph) GetOpaqueAddr() *unsafe.Pointer {
	return (*unsafe.Pointer)(&fltg.opaque)
}

// GetExecute gets `AVFilterGraph.execute` value.
func (fltg *AVFilterGraph) GetExecute() *AVFilterExecuteFunc {
	return (*AVFilterExecuteFunc)(fltg.execute)
}

// SetExecute sets `AVFilterGraph.execute` value.
func (fltg *AVFilterGraph) SetExecute(v *AVFilterExecuteFunc) {
	fltg.execute = (*C.avfilter_execute_func)(v)
}

// GetExecuteAddr gets `AVFilterGraph.execute` address.
func (fltg *AVFilterGraph) GetExecuteAddr() **AVFilterExecuteFunc {
	return (**AVFilterExecuteFunc)(unsafe.Pointer(&fltg.execute))
}

// GetAresampleSwrOpts gets `AVFilterGraph.aresample_swr_opts` value.
func (fltg *AVFilterGraph) GetAresampleSwrOpts() string {
	return C.GoString(fltg.aresample_swr_opts)
}

// AvFilterGraphAlloc allocates a filter graph.
func AvFilterGraphAlloc() *AVFilterGraph {
	return (*AVFilterGraph)(C.avfilter_graph_alloc())
}

// AvFilterGraphAllocFilter creates a new filter instance in a filter graph.
func AvFilterGraphAllocFilter(graph *AVFilterGraph, filter *AVFilter, name string) *AVFilterContext {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AVFilterContext)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(graph),
		(*C.struct_AVFilter)(filter), (*C.char)(namePtr)))
}

// AvFilterGraphGetFilter gets a filter instance identified by instance name from graph.
func AvFilterGraphGetFilter(graph *AVFilterGraph, name string) *AVFilterContext {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AVFilterContext)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(graph),
		(*C.char)(namePtr)))
}

// AvFilterGraphCreateFilter creates and adds a filter instance into an existing graph.
func AvFilterGraphCreateFilter(filtCtx **AVFilterContext, filter *AVFilter,
	name, args string, opaque CVoidPointer, graph *AVFilterGraph) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	argsPtr, argsFunc := StringCasting(args)
	defer argsFunc()
	return (int32)(C.avfilter_graph_create_filter(
		(**C.struct_AVFilterContext)(unsafe.Pointer(filtCtx)),
		(*C.struct_AVFilter)(filter), (*C.char)(namePtr), (*C.char)(argsPtr),
		VoidPointer(opaque), (*C.struct_AVFilterGraph)(graph)))
}

// AvFilterGraphSetAutoConvert enables or disables automatic format conversion inside the graph.
func AvFilterGraphSetAutoConvert(graph *AVFilterGraph, flags uint32) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(graph), (C.uint)(flags))
}

const (
	AVFILTER_AUTO_CONVERT_ALL  = int32(C.AVFILTER_AUTO_CONVERT_ALL)
	AVFILTER_AUTO_CONVERT_NONE = int32(C.AVFILTER_AUTO_CONVERT_NONE)
)

// AvFilterGraphConfig checks validity and configure all the links and formats in the graph.
func AvFilterGraphConfig(graph *AVFilterGraph, logCtx CVoidPointer) int32 {
	return (int32)(C.avfilter_graph_config((*C.struct_AVFilterGraph)(graph), VoidPointer(logCtx)))
}

// AvFilterGraphFree frees a graph, destroy its links, and set *graph to NULL.
func AvFilterGraphFree(graph **AVFilterGraph) {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(graph)))
}

// AVFilterInOut
type AVFilterInOut C.struct_AVFilterInOut

// GetName gets `AVFilterInOut.name` value.
func (fltio *AVFilterInOut) GetName() string {
	return C.GoString(fltio.name)
}

// SetName sets `AVFilterInOut.name` value.
func (fltio *AVFilterInOut) SetName(v string) {
	FreePointer(fltio.name)
	fltio.name, _ = StringCasting(v)
}

// GetFilterCtx gets `AVFilterInOut.filter_ctx` value.
func (fltio *AVFilterInOut) GetFilterCtx() *AVFilterContext {
	return (*AVFilterContext)(fltio.filter_ctx)
}

// SetFilterCtx sets `AVFilterInOut.filter_ctx` value.
func (fltio *AVFilterInOut) SetFilterCtx(v *AVFilterContext) {
	fltio.filter_ctx = (*C.struct_AVFilterContext)(v)
}

// GetFilterCtxAddr gets `AVFilterInOut.filter_ctx` address.
func (fltio *AVFilterInOut) GetFilterCtxAddr() **AVFilterContext {
	return (**AVFilterContext)(unsafe.Pointer(&fltio.filter_ctx))
}

// GetPadIdx gets `AVFilterInOut.pad_idx` value.
func (fltio *AVFilterInOut) GetPadIdx() int32 {
	return (int32)(fltio.pad_idx)
}

// SetPadIdx sets `AVFilterInOut.pad_idx` value.
func (fltio *AVFilterInOut) SetPadIdx(v int32) {
	fltio.pad_idx = (C.int)(v)
}

// GetPadIdxAddr gets `AVFilterInOut.pad_idx` address.
func (fltio *AVFilterInOut) GetPadIdxAddr() *int32 {
	return (*int32)(&fltio.pad_idx)
}

// GetNext gets `AVFilterInOut.next` value.
func (fltio *AVFilterInOut) GetNext() *AVFilterInOut {
	return (*AVFilterInOut)(fltio.next)
}

// SetNext sets `AVFilterInOut.next` value.
func (fltio *AVFilterInOut) SetNext(v *AVFilterInOut) {
	fltio.next = (*C.struct_AVFilterInOut)(v)
}

// GetNextAddr gets `AVFilterInOut.next` address.
func (fltio *AVFilterInOut) GetNextAddr() **AVFilterInOut {
	return (**AVFilterInOut)(unsafe.Pointer(&fltio.next))
}

// AvFilterInoutAlloc allocates a single AVFilterInOut entry.
func AvFilterInoutAlloc() *AVFilterInOut {
	return (*AVFilterInOut)(C.avfilter_inout_alloc())
}

// AvFilterInoutFree frees the supplied list of AVFilterInOut and set *inout to NULL.
func AvFilterInoutFree(inout **AVFilterInOut) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(inout)))
}

// AvFilterGraphParse adds a graph described by a string to a graph.
func AvFilterGraphParse(graph *AVFilterGraph, filters string, inputs, outputs *AVFilterInOut,
	logCtx CVoidPointer) int32 {
	filtersPtr, filtersFunc := StringCasting(filters)
	defer filtersFunc()
	return (int32)(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(graph),
		(*C.char)(filtersPtr),
		(*C.struct_AVFilterInOut)(inputs),
		(*C.struct_AVFilterInOut)(outputs), VoidPointer(logCtx)))
}

// AvFilterGraphParsePtr adds a graph described by a string to a graph.
func AvFilterGraphParsePtr(graph *AVFilterGraph, filters string, inputs, outputs **AVFilterInOut,
	logCtx CVoidPointer) int32 {
	filtersPtr, filtersFunc := StringCasting(filters)
	defer filtersFunc()
	return (int32)(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(graph),
		(*C.char)(filtersPtr),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(outputs)), VoidPointer(logCtx)))
}

// AvFilterGraphParse2 adds a graph described by a string to a graph.
func AvFilterGraphParse2(graph *AVFilterGraph, filters string,
	inputs, outputs **AVFilterInOut) int32 {
	filtersPtr, filtersFunc := StringCasting(filters)
	defer filtersFunc()
	return (int32)(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(graph),
		(*C.char)(filtersPtr),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(outputs))))
}

// AvFilterGraphSendCommand sends a command to one or more filter instances.
func AvFilterGraphSendCommand(graph *AVFilterGraph, target, cmd, arg string,
	resLen, flags int32) (res string, ret int32) {
	targetPtr, targetFunc := StringCasting(target)
	defer targetFunc()
	cmdPtr, cmdFunc := StringCasting(cmd)
	defer cmdFunc()
	argPtr, argFunc := StringCasting(arg)
	defer argFunc()
	resBuf := make([]C.char, resLen)
	ret = (int32)(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(graph),
		(*C.char)(targetPtr), (*C.char)(cmdPtr), (*C.char)(argPtr),
		(*C.char)(&resBuf[0]), (C.int)(resLen), (C.int)(flags)))
	return C.GoString(&resBuf[0]), ret
}

// AvFilterGraphQueueCommand queues a command for one or more filter instances.
func AvFilterGraphQueueCommand(graph *AVFilterGraph, target, cmd, arg string,
	flags int32, ts float64) int32 {
	targetPtr, targetFunc := StringCasting(target)
	defer targetFunc()
	cmdPtr, cmdFunc := StringCasting(cmd)
	defer cmdFunc()
	argPtr, argFunc := StringCasting(arg)
	defer argFunc()
	return (int32)(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(graph),
		(*C.char)(targetPtr), (*C.char)(cmdPtr), (*C.char)(argPtr), (C.int)(flags), (C.double)(ts)))
}

// AvFilterGraphDump dumps a graph into a human-readable string representation.
func AvFilterGraphDump(graph *AVFilterGraph, options string) string {
	optionsPtr, optionsFunc := StringCasting(options)
	defer optionsFunc()
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(graph),
		(*C.char)(optionsPtr)))
}

// AvFilterGraphRequestOldest requests a frame on the oldest sink link.
func AvFilterGraphRequestOldest(graph *AVFilterGraph) int32 {
	return (int32)(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(graph)))
}
