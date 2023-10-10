package ffmpeg

/*
#include <libavcodec/bsf.h>
*/
import "C"
import "unsafe"

// AvBSFInternal
type AvBSFInternal C.struct_AVBSFInternal

// AvBSFContext
type AvBSFContext C.struct_AVBSFContext

// AvBitStreamFilter
type AvBitStreamFilter C.struct_AVBitStreamFilter

// AvBsfGetByName returns a bitstream filter with the specified name or NULL if no such
// bitstream filter exists.
func AvBsfGetByName(name string) *AvBitStreamFilter {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AvBitStreamFilter)(C.av_bsf_get_by_name((*C.char)(namePtr)))
}

// AvBsfIterate iterates over all registered bitstream filters.
func AvBsfIterate(opaque *unsafe.Pointer) *AvBitStreamFilter {
	return (*AvBitStreamFilter)(C.av_bsf_iterate(opaque))
}

// AvBsfAlloc allocates a context for a given bitstream filter.
func AvBsfAlloc(filter *AvBitStreamFilter, ctx **AvBSFContext) int32 {
	return (int32)(C.av_bsf_alloc((*C.struct_AVBitStreamFilter)(filter),
		(**C.struct_AVBSFContext)(unsafe.Pointer(ctx))))
}

// AvBsfInit prepares the filter for use, after all the parameters and options have been set.
func AvBsfInit(ctx *AvBSFContext) int32 {
	return (int32)(C.av_bsf_init((*C.struct_AVBSFContext)(ctx)))
}

// AvBsfSendPacket submits a packet for filtering.
func AvBsfSendPacket(ctx *AvBSFContext, pkt *AvPacket) int32 {
	return (int32)(C.av_bsf_send_packet((*C.struct_AVBSFContext)(ctx),
		(*C.struct_AVPacket)(pkt)))
}

// AvBsfReceivePacket retrieves a filtered packet.
func AvBsfReceivePacket(ctx *AvBSFContext, pkt *AvPacket) int32 {
	return (int32)(C.av_bsf_receive_packet((*C.struct_AVBSFContext)(ctx),
		(*C.struct_AVPacket)(pkt)))
}

// AvBsfFlush resets the internal bitstream filter state. Should be called e.g. when seeking.
func AvBsfFlush(ctx *AvBSFContext) {
	C.av_bsf_flush((*C.struct_AVBSFContext)(ctx))
}

// AvBsfFree frees a bitstream filter context.
func AvBsfFree(ctx **AvBSFContext) {
	C.av_bsf_free((**C.struct_AVBSFContext)(unsafe.Pointer(ctx)))
}

// AvBsfGetClass gets the AVClass for AVBSFContext.
func AvBsfGetClass() *AvClass {
	return (*AvClass)(C.av_bsf_get_class())
}

type AvBSFList C.struct_AVBSFList

// AvBsfListAlloc allocates empty list of bitstream filters.
func AvBsfListAlloc() *AvBSFList {
	return (*AvBSFList)(C.av_bsf_list_alloc())
}

// AvBsfListFree frees list of bitstream filters.
func AvBsfListFree(lst **AvBSFList) {
	C.av_bsf_list_free((**C.struct_AVBSFList)(unsafe.Pointer(lst)))
}

// AvBsfListAppend appends bitstream filter to the list of bitstream filters.
func AvBsfListAppend(lst *AvBSFList, bsf *AvBSFContext) {
	C.av_bsf_list_append((*C.struct_AVBSFList)(lst),
		(*C.struct_AVBSFContext)(bsf))
}

// AvBsfListAppend2
func AvBsfListAppend2(lst *AvBSFList, bsfName string, options **AvDictionary) {
	bsfNamePtr, bsfNameFunc := StringCasting(bsfName)
	defer bsfNameFunc()
	C.av_bsf_list_append2((*C.struct_AVBSFList)(lst),
		(*C.char)(bsfNamePtr), (**C.struct_AVDictionary)(unsafe.Pointer(options)))
}

// AvBsfListFinalize finalizes list of bitstream filters.
func AvBsfListFinalize(lst **AvBSFList, bsf **AvBSFContext) int32 {
	return (int32)(C.av_bsf_list_finalize((**C.struct_AVBSFList)(unsafe.Pointer(lst)),
		(**C.struct_AVBSFContext)(unsafe.Pointer(bsf))))
}

// AvBsfListParseStr parses string describing list of bitstream filters and creates single
// AVBSFContext describing the whole chain of bitstream filters.
func AvBsfListParseStr(str string, bsf **AvBSFContext) {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	C.av_bsf_list_parse_str((*C.char)(strPtr), (**C.struct_AVBSFContext)(unsafe.Pointer(bsf)))
}

// AvBsfGetNullFilter gets null/pass-through bitstream filter.
func AvBsfGetNullFilter(bsf **AvBSFContext) int32 {
	return (int32)(C.av_bsf_get_null_filter((**C.struct_AVBSFContext)(unsafe.Pointer(bsf))))
}
