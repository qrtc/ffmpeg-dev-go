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

// Custom: GetAvClass gets `AVBSFContext.av_class` value.
func (bsfc *AvBSFContext) GetAvClass() *AvClass {
	return (*AvClass)(bsfc.av_class)
}

// Custom: SetAvClass sets `AVBSFContext.av_class` value.
func (bsfc *AvBSFContext) SetAvClass(v *AvClass) {
	bsfc.av_class = (*C.struct_AVClass)(v)
}

// Custom: GetAvClassAddr gets `AVBSFContext.av_class` address.
func (bsfc *AvBSFContext) GetAvClassAddr() **AvClass {
	return (**AvClass)(unsafe.Pointer(&bsfc.av_class))
}

// Custom: GetFilter gets `AVBSFContext.filter` value.
func (bsfc *AvBSFContext) GetFilter() *AvBitStreamFilter {
	return (*AvBitStreamFilter)(bsfc.filter)
}

// Custom: SetFilter sets `AVBSFContext.filter` value.
func (bsfc *AvBSFContext) SetFilter(v *AvBitStreamFilter) {
	bsfc.filter = (*C.struct_AVBitStreamFilter)(v)
}

// Custom: GetFilterAddr gets `AVBSFContext.filter` address.
func (bsfc *AvBSFContext) GetFilterAddr() **AvBitStreamFilter {
	return (**AvBitStreamFilter)(unsafe.Pointer(&bsfc.filter))
}

// Custom: GetInternal gets `AVBSFContext.internal` value.
func (bsfc *AvBSFContext) GetInternal() *AvBSFInternal {
	return (*AvBSFInternal)(bsfc.internal)
}

// Custom: SetInternal sets `AVBSFContext.internal` value.
func (bsfc *AvBSFContext) SetInternal(v *AvBSFInternal) {
	bsfc.internal = (*C.struct_AVBSFInternal)(v)
}

// Custom: GetInternalAddr gets `AVBSFContext.internal` address.
func (bsfc *AvBSFContext) GetInternalAddr() **AvBSFInternal {
	return (**AvBSFInternal)(unsafe.Pointer(&bsfc.internal))
}

// Custom: GetPrivData gets `AVBSFContext.priv_data` value.
func (bsfc *AvBSFContext) GetPrivData() unsafe.Pointer {
	return bsfc.priv_data
}

// Custom: SetPrivData sets `AVBSFContext.priv_data` value.
func (bsfc *AvBSFContext) SetPrivData(v unsafe.Pointer) {
	bsfc.priv_data = v
}

// Custom: GetPrivDataAddr gets `AVBSFContext.priv_data` address.
func (bsfc *AvBSFContext) GetPrivDataAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&bsfc.priv_data)
}

// Custom: GetParIn gets `AVBSFContext.par_in` value.
func (bsfc *AvBSFContext) GetParIn() *AvCodecParameters {
	return (*AvCodecParameters)(bsfc.par_in)
}

// Custom: SetParIn sets `AVBSFContext.par_in` value.
func (bsfc *AvBSFContext) SetParIn(v *AvCodecParameters) {
	bsfc.par_in = (*C.struct_AVCodecParameters)(v)
}

// Custom: GetParInAddr gets `AVBSFContext.par_in` address.
func (bsfc *AvBSFContext) GetParInAddr() **AvCodecParameters {
	return (**AvCodecParameters)(unsafe.Pointer(&bsfc.par_in))
}

// Custom: GetParOut gets `AVBSFContext.par_out` value.
func (bsfc *AvBSFContext) GetParOut() *AvCodecParameters {
	return (*AvCodecParameters)(bsfc.par_out)
}

// Custom: SetParOut sets `AVBSFContext.par_out` value.
func (bsfc *AvBSFContext) SetParOut(v *AvCodecParameters) {
	bsfc.par_out = (*C.struct_AVCodecParameters)(v)
}

// Custom: GetParOutAddr gets `AVBSFContext.par_out` address.
func (bsfc *AvBSFContext) GetParOutAddr() **AvCodecParameters {
	return (**AvCodecParameters)(unsafe.Pointer(&bsfc.par_out))
}

// Custom: GetTimeBaseIn gets `AVBSFContext.time_base_in` value.
func (bsfc *AvBSFContext) GetTimeBaseIn() AvRational {
	return (AvRational)(bsfc.time_base_in)
}

// Custom: SetTimeBaseIn sets `AVBSFContext.time_base_in` value.
func (bsfc *AvBSFContext) SetTimeBaseIn(v AvRational) {
	bsfc.time_base_in = (C.AVRational)(v)
}

// Custom: GetTimeBaseInAddr gets `AVBSFContext.time_base_in` address.
func (bsfc *AvBSFContext) GetTimeBaseInAddr() *AvRational {
	return (*AvRational)(&bsfc.time_base_in)
}

// Custom: GetTimeBaseOut gets `AVBSFContext.time_base_out` value.
func (bsfc *AvBSFContext) GetTimeBaseOut() AvRational {
	return (AvRational)(bsfc.time_base_out)
}

// Custom: SetTimeBaseOut sets `AVBSFContext.time_base_out` value.
func (bsfc *AvBSFContext) SetTimeBaseOut(v AvRational) {
	bsfc.time_base_out = (C.AVRational)(v)
}

// Custom: GetTimeBaseOutAddr gets `AVBSFContext.time_base_out` address.
func (bsfc *AvBSFContext) GetTimeBaseOutAddr() *AvRational {
	return (*AvRational)(&bsfc.time_base_out)
}

// AvBitStreamFilter
type AvBitStreamFilter C.struct_AVBitStreamFilter

// Custom: GetName gets `AVBitStreamFilter.name` value.
func (bsf *AvBitStreamFilter) GetName() string {
	return C.GoString(bsf.name)
}

// Custom: GetCodecIds gets `AVBitStreamFilter.codec_ids` value.
func (bsf *AvBitStreamFilter) GetCodecIds() (v []AvCodecID) {
	if bsf.codec_ids == nil {
		return v
	}
	ptr := (*AvCodecID)(bsf.codec_ids)
	for *ptr != AV_CODEC_ID_NONE {
		v = append(v, *ptr)
		ptr = (*AvCodecID)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) +
			uintptr(unsafe.Sizeof(*ptr))))
	}
	return v
}

// Custom: GetPrivClass gets `AVBitStreamFilter.priv_class` value.
func (bsf *AvBitStreamFilter) GetPrivClass() *AvClass {
	return (*AvClass)(bsf.priv_class)
}

// Custom: SetPrivClass sets `AVBitStreamFilter.priv_class` value.
func (bsf *AvBitStreamFilter) SetPrivClass(v *AvClass) {
	bsf.priv_class = (*C.struct_AVClass)(v)
}

// Custom: GetPrivClassAddr gets `AVBitStreamFilter.priv_class` address.
func (bsf *AvBitStreamFilter) GetPrivClassAddr() **AvClass {
	return (**AvClass)(unsafe.Pointer(&bsf.priv_class))
}

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
