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

// AvFilterPad
type AvFilterPad C.struct_AVFilterPad

// AvFilterFormats
type AvFilterFormats C.struct_AVFilterFormats

// AvFilterChannelLayouts
type AvFilterChannelLayouts C.struct_AVFilterChannelLayouts

// AvFilterPadCount gets the number of elements in a NULL-terminated array of AVFilterPads (e.g.
// AvFilter.inputs/outputs).
func AvFilterPadCount(pads *AvFilterPad) int32 {
	return (int32)(C.avfilter_pad_count((*C.struct_AVFilterPad)(pads)))
}

// AvFilterPadGetName gets the name of an AvFilterPad.
func AvFilterPadGetName(pads *AvFilterPad, padIdx int32) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(pads), (C.int)(padIdx)))
}

// AvFilterPadGetType gets the type of an AvFilterPad.
func AvFilterPadGetType(pads *AvFilterPad, padIdx int32) AvMediaType {
	return (AvMediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(pads), (C.int)(padIdx)))
}

const (
	AVFILTER_FLAG_DYNAMIC_INPUTS            = C.AVFILTER_FLAG_DYNAMIC_INPUTS
	AVFILTER_FLAG_DYNAMIC_OUTPUTS           = C.AVFILTER_FLAG_DYNAMIC_OUTPUTS
	AVFILTER_FLAG_SLICE_THREADS             = C.AVFILTER_FLAG_SLICE_THREADS
	AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC  = C.AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC
	AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL = C.AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL
	AVFILTER_FLAG_SUPPORT_TIMELINE          = C.AVFILTER_FLAG_SUPPORT_TIMELINE
)

// Filter definition. This defines the pads a filter contains, and all the
// callback functions used to interact with the filter.
type AvFilter C.struct_AVFilter

const (
	AVFILTER_THREAD_SLICE = C.AVFILTER_THREAD_SLICE
)

// AvFilterInternal
type AvFilterInternal C.struct_AVFilterInternal

// AvFilterContext
type AvFilterContext C.struct_AVFilterContext

// AvFilterFormatsConfig
type AvFilterFormatsConfig C.struct_AVFilterFormatsConfig

// AvFilterLink
type AvFilterLink C.struct_AVFilterLink

// AvFilterContextLink links two filters together.
func AvFilterContextLink(src *AvFilterContext, srcpad uint32,
	dst *AvFilterContext, dstpad uint32) int32 {
	return (int32)(C.avfilter_link((*C.struct_AVFilterContext)(src), (C.uint)(srcpad),
		(*C.struct_AVFilterContext)(dst), (C.uint)(dstpad)))
}

// AvFilterLinkFree frees the link in *link, and set its pointer to NULL.
func AvFilterLinkFree(link **AvFilterLink) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(link)))
}

// Deprecated: Use av_buffersink_get_channels() instead.
func AvFilterLinkGetChannels(link *AvFilterLink) int32 {
	return (int32)(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(link)))
}

// Deprecated: No use
func AvFilterLinkSetClosed(link *AvFilterLink, closed int32) {
	C.avfilter_link_set_closed((*C.struct_AVFilterLink)(link), (C.int)(closed))
}

// AvFilterConfigLinks negotiates the media format, dimensions, etc of all inputs to a filter.
func AvFilterConfigLinks(filter *AvFilterContext) int32 {
	return (int32)(C.avfilter_config_links((*C.struct_AVFilterContext)(filter)))
}

const (
	AVFILTER_CMD_FLAG_ONE  = C.AVFILTER_CMD_FLAG_ONE
	AVFILTER_CMD_FLAG_FAST = C.AVFILTER_CMD_FLAG_FAST
)

// AvFilterProcessCommand  makes the filter instance process a command.
// It is recommended to use AvFilterGraphSendCommand().
func AvFilterProcessCommand(filter *AvFilterContext, cmd, arg string, resLen, flags int32) (res string, ret int32) {
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
func AvFilterIterate(opaque *unsafe.Pointer) *AvFilter {
	return (*AvFilter)(C.av_filter_iterate(opaque))
}

// Deprecated: No use
func AvFilterRegisterAll() {
	C.avfilter_register_all()
}

// Deprecated: No use
func AvFilterRegister(filter *AvFilter) {
	C.avfilter_register((*C.struct_AVFilter)(filter))
}

// Deprecated: No use
func AvFilterNext(filter *AvFilter) *AvFilter {
	return (*AvFilter)(C.avfilter_next((*C.struct_AVFilter)(filter)))
}

// AvFilterGetByName gets a filter definition matching the given name.
func AvFilterGetByName(name string) *AvFilter {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AvFilter)(C.avfilter_get_by_name((*C.char)(namePtr)))
}

// AvFilterInitStr initializes a filter with the supplied parameters.
func AvFilterInitStr(ctx *AvFilterContext, args string) int32 {
	argsPtr, argsFunc := StringCasting(args)
	defer argsFunc()
	return (int32)(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx),
		(*C.char)(argsPtr)))
}

// AvFilterInitDict initialize a filter with the supplied dictionary of options.
func AvFilterInitDict(ctx *AvFilterContext, options **AvDictionary) int32 {
	return (int32)(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvFilterFree frees a filter context. This will also remove the filter from its
// filtergraph's list of filters.
func AvFilterFree(ctx *AvFilterContext) {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

// AvFilterInsertFilter inserts a filter in the middle of an existing link.
func AvFilterInsertFilter(ctx *AvFilterContext, link *AvFilterLink,
	filtSrcpadIdx, filtDstpadIdx uint32) int32 {
	return (int32)(C.avfilter_insert_filter(
		(*C.struct_AVFilterLink)(link),
		(*C.struct_AVFilterContext)(ctx),
		(C.uint)(filtSrcpadIdx), (C.uint)(filtDstpadIdx)))
}

// AvFilterGetClass returns AvClass for AvFilterContext.
func AvFilterGetClass() *AvClass {
	return (*AvClass)(C.avfilter_get_class())
}

// typedef int (avfilter_action_func)(AVFilterContext *ctx, void *arg, int jobnr, int nb_jobs)
type AvfilterActionFunc C.avfilter_action_func

// typedef int (avfilter_execute_func)(AVFilterContext *ctx, avfilter_action_func *func,
// void *arg, int *ret, int nb_jobs)
type AvfilterExecuteFunc C.avfilter_execute_func

type AvFilterGraph C.struct_AVFilterGraph

// AvFilterGraphAlloc allocates a filter graph.
func AvFilterGraphAlloc() *AvFilterGraph {
	return (*AvFilterGraph)(C.avfilter_graph_alloc())
}

// AvFilterGraphAllocFilter creates a new filter instance in a filter graph.
func AvFilterGraphAllocFilter(graph *AvFilterGraph, filter *AvFilter, name string) *AvFilterContext {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AvFilterContext)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(graph),
		(*C.struct_AVFilter)(filter), (*C.char)(namePtr)))
}

// AvFilterGraphGetFilter gets a filter instance identified by instance name from graph.
func AvFilterGraphGetFilter(graph *AvFilterGraph, name string) *AvFilterContext {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AvFilterContext)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(graph),
		(*C.char)(namePtr)))
}

// AvFilterGraphCreateFilter creates and adds a filter instance into an existing graph.
func AvFilterGraphCreateFilter(graph *AvFilterGraph, filtCtx **AvFilterContext, filter *AvFilter,
	name, args string, opaque unsafe.Pointer) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	argsPtr, argsFunc := StringCasting(args)
	defer argsFunc()
	return (int32)(C.avfilter_graph_create_filter(
		(**C.struct_AVFilterContext)(unsafe.Pointer(filtCtx)),
		(*C.struct_AVFilter)(filter), (*C.char)(namePtr), (*C.char)(argsPtr),
		opaque, (*C.struct_AVFilterGraph)(graph)))
}

// AvFilterGraphSetAutoConvert enables or disables automatic format conversion inside the graph.
func AvFilterGraphSetAutoConvert(graph *AvFilterGraph, flags uint32) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(graph), (C.uint)(flags))
}

const (
	AVFILTER_AUTO_CONVERT_ALL  = int32(C.AVFILTER_AUTO_CONVERT_ALL)
	AVFILTER_AUTO_CONVERT_NONE = int32(C.AVFILTER_AUTO_CONVERT_NONE)
)

// AvFilterGraphConfig checks validity and configure all the links and formats in the graph.
func AvFilterGraphConfig(graph *AvFilterGraph, logCtx unsafe.Pointer) int32 {
	return (int32)(C.avfilter_graph_config((*C.struct_AVFilterGraph)(graph), logCtx))
}

// AvFilterGraphFree frees a graph, destroy its links, and set *graph to NULL.
func AvFilterGraphFree(graph **AvFilterGraph) {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(graph)))
}

type AvFilterInOut C.struct_AVFilterInOut

// AvFilterInoutAlloc allocates a single AVFilterInOut entry.
func AvFilterInoutAlloc() *AvFilterInOut {
	return (*AvFilterInOut)(C.avfilter_inout_alloc())
}

// AvFilterInoutFree frees the supplied list of AVFilterInOut and set *inout to NULL.
func AvFilterInoutFree(inout **AvFilterInOut) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(inout)))
}

// AvFilterGraphParse adds a graph described by a string to a graph.
func AvFilterGraphParse(graph *AvFilterGraph, filters string, inputs, outputs *AvFilterInOut,
	logCtx unsafe.Pointer) int32 {
	filtersPtr, filtersFunc := StringCasting(filters)
	defer filtersFunc()
	return (int32)(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(graph),
		(*C.char)(filtersPtr),
		(*C.struct_AVFilterInOut)(inputs),
		(*C.struct_AVFilterInOut)(outputs), logCtx))
}

// AvFilterGraphParsePtr adds a graph described by a string to a graph.
func AvFilterGraphParsePtr(graph *AvFilterGraph, filters string, inputs, outputs *AvFilterInOut,
	logCtx unsafe.Pointer) int32 {
	filtersPtr, filtersFunc := StringCasting(filters)
	defer filtersFunc()
	return (int32)(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(graph),
		(*C.char)(filtersPtr),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(outputs)), logCtx))
}

// AvFilterGraphParse2 adds a graph described by a string to a graph.
func AvFilterGraphParse2(graph *AvFilterGraph, filters string,
	inputs, outputs *AvFilterInOut) int32 {
	filtersPtr, filtersFunc := StringCasting(filters)
	defer filtersFunc()
	return (int32)(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(graph),
		(*C.char)(filtersPtr),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(outputs))))
}

// AvFilterGraphSendCommand sends a command to one or more filter instances.
func AvFilterGraphSendCommand(graph *AvFilterGraph, target, cmd, arg string,
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
func AvFilterGraphQueueCommand(graph *AvFilterGraph, target, cmd, arg string,
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
func AvFilterGraphDump(graph *AvFilterGraph, options string) string {
	optionsPtr, optionsFunc := StringCasting(options)
	defer optionsFunc()
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(graph),
		(*C.char)(optionsPtr)))
}

// AvFilterGraphRequestOldest requests a frame on the oldest sink link.
func AvFilterGraphRequestOldest(graph *AvFilterGraph) int32 {
	return (int32)(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(graph)))
}
