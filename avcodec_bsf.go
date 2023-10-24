package ffmpeg

/*
#include <libavcodec/bsf.h>
*/
import "C"
import "unsafe"

// AVBSFInternal
type AVBSFInternal C.struct_AVBSFInternal

// AVBSFContext
type AVBSFContext C.struct_AVBSFContext

// Custom: GetAvClass gets `AVBSFContext.av_class` value.
func (bsfc *AVBSFContext) GetAvClass() *AVClass {
	return (*AVClass)(bsfc.av_class)
}

// Custom: SetAvClass sets `AVBSFContext.av_class` value.
func (bsfc *AVBSFContext) SetAvClass(v *AVClass) {
	bsfc.av_class = (*C.struct_AVClass)(v)
}

// Custom: GetAvClassAddr gets `AVBSFContext.av_class` address.
func (bsfc *AVBSFContext) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&bsfc.av_class))
}

// Custom: GetFilter gets `AVBSFContext.filter` value.
func (bsfc *AVBSFContext) GetFilter() *AVBitStreamFilter {
	return (*AVBitStreamFilter)(bsfc.filter)
}

// Custom: SetFilter sets `AVBSFContext.filter` value.
func (bsfc *AVBSFContext) SetFilter(v *AVBitStreamFilter) {
	bsfc.filter = (*C.struct_AVBitStreamFilter)(v)
}

// Custom: GetFilterAddr gets `AVBSFContext.filter` address.
func (bsfc *AVBSFContext) GetFilterAddr() **AVBitStreamFilter {
	return (**AVBitStreamFilter)(unsafe.Pointer(&bsfc.filter))
}

// Custom: GetInternal gets `AVBSFContext.internal` value.
func (bsfc *AVBSFContext) GetInternal() *AVBSFInternal {
	return (*AVBSFInternal)(bsfc.internal)
}

// Custom: SetInternal sets `AVBSFContext.internal` value.
func (bsfc *AVBSFContext) SetInternal(v *AVBSFInternal) {
	bsfc.internal = (*C.struct_AVBSFInternal)(v)
}

// Custom: GetInternalAddr gets `AVBSFContext.internal` address.
func (bsfc *AVBSFContext) GetInternalAddr() **AVBSFInternal {
	return (**AVBSFInternal)(unsafe.Pointer(&bsfc.internal))
}

// Custom: GetPrivData gets `AVBSFContext.priv_data` value.
func (bsfc *AVBSFContext) GetPrivData() unsafe.Pointer {
	return bsfc.priv_data
}

// Custom: SetPrivData sets `AVBSFContext.priv_data` value.
func (bsfc *AVBSFContext) SetPrivData(v CVoidPointer) {
	bsfc.priv_data = VoidPointer(v)
}

// Custom: GetPrivDataAddr gets `AVBSFContext.priv_data` address.
func (bsfc *AVBSFContext) GetPrivDataAddr() unsafe.Pointer {
	return (unsafe.Pointer)(&bsfc.priv_data)
}

// Custom: GetParIn gets `AVBSFContext.par_in` value.
func (bsfc *AVBSFContext) GetParIn() *AVCodecParameters {
	return (*AVCodecParameters)(bsfc.par_in)
}

// Custom: SetParIn sets `AVBSFContext.par_in` value.
func (bsfc *AVBSFContext) SetParIn(v *AVCodecParameters) {
	bsfc.par_in = (*C.struct_AVCodecParameters)(v)
}

// Custom: GetParInAddr gets `AVBSFContext.par_in` address.
func (bsfc *AVBSFContext) GetParInAddr() **AVCodecParameters {
	return (**AVCodecParameters)(unsafe.Pointer(&bsfc.par_in))
}

// Custom: GetParOut gets `AVBSFContext.par_out` value.
func (bsfc *AVBSFContext) GetParOut() *AVCodecParameters {
	return (*AVCodecParameters)(bsfc.par_out)
}

// Custom: SetParOut sets `AVBSFContext.par_out` value.
func (bsfc *AVBSFContext) SetParOut(v *AVCodecParameters) {
	bsfc.par_out = (*C.struct_AVCodecParameters)(v)
}

// Custom: GetParOutAddr gets `AVBSFContext.par_out` address.
func (bsfc *AVBSFContext) GetParOutAddr() **AVCodecParameters {
	return (**AVCodecParameters)(unsafe.Pointer(&bsfc.par_out))
}

// Custom: GetTimeBaseIn gets `AVBSFContext.time_base_in` value.
func (bsfc *AVBSFContext) GetTimeBaseIn() AVRational {
	return (AVRational)(bsfc.time_base_in)
}

// Custom: SetTimeBaseIn sets `AVBSFContext.time_base_in` value.
func (bsfc *AVBSFContext) SetTimeBaseIn(v AVRational) {
	bsfc.time_base_in = (C.struct_AVRational)(v)
}

// Custom: GetTimeBaseInAddr gets `AVBSFContext.time_base_in` address.
func (bsfc *AVBSFContext) GetTimeBaseInAddr() *AVRational {
	return (*AVRational)(&bsfc.time_base_in)
}

// Custom: GetTimeBaseOut gets `AVBSFContext.time_base_out` value.
func (bsfc *AVBSFContext) GetTimeBaseOut() AVRational {
	return (AVRational)(bsfc.time_base_out)
}

// Custom: SetTimeBaseOut sets `AVBSFContext.time_base_out` value.
func (bsfc *AVBSFContext) SetTimeBaseOut(v AVRational) {
	bsfc.time_base_out = (C.struct_AVRational)(v)
}

// Custom: GetTimeBaseOutAddr gets `AVBSFContext.time_base_out` address.
func (bsfc *AVBSFContext) GetTimeBaseOutAddr() *AVRational {
	return (*AVRational)(&bsfc.time_base_out)
}

// AVBitStreamFilter
type AVBitStreamFilter C.struct_AVBitStreamFilter

// Custom: GetName gets `AVBitStreamFilter.name` value.
func (bsf *AVBitStreamFilter) GetName() string {
	return C.GoString(bsf.name)
}

// Custom: GetCodecIds gets `AVBitStreamFilter.codec_ids` value.
func (bsf *AVBitStreamFilter) GetCodecIds() []AVCodecID {
	return SliceTrunc((*AVCodecID)(bsf.codec_ids), func(ac AVCodecID) bool {
		return ac == AV_CODEC_ID_NONE
	})
}

// Custom: GetPrivClass gets `AVBitStreamFilter.priv_class` value.
func (bsf *AVBitStreamFilter) GetPrivClass() *AVClass {
	return (*AVClass)(bsf.priv_class)
}

// Custom: SetPrivClass sets `AVBitStreamFilter.priv_class` value.
func (bsf *AVBitStreamFilter) SetPrivClass(v *AVClass) {
	bsf.priv_class = (*C.struct_AVClass)(v)
}

// Custom: GetPrivClassAddr gets `AVBitStreamFilter.priv_class` address.
func (bsf *AVBitStreamFilter) GetPrivClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&bsf.priv_class))
}

// AvBsfGetByName returns a bitstream filter with the specified name or NULL if no such
// bitstream filter exists.
func AvBsfGetByName(name string) *AVBitStreamFilter {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AVBitStreamFilter)(C.av_bsf_get_by_name((*C.char)(namePtr)))
}

// AvBsfIterate iterates over all registered bitstream filters.
func AvBsfIterate(opaque CVoidPointerPointer) *AVBitStreamFilter {
	return (*AVBitStreamFilter)(C.av_bsf_iterate(VoidPointerPointer(opaque)))
}

// AvBsfAlloc allocates a context for a given bitstream filter.
func AvBsfAlloc(filter *AVBitStreamFilter, ctx **AVBSFContext) int32 {
	return (int32)(C.av_bsf_alloc((*C.struct_AVBitStreamFilter)(filter),
		(**C.struct_AVBSFContext)(unsafe.Pointer(ctx))))
}

// AvBsfInit prepares the filter for use, after all the parameters and options have been set.
func AvBsfInit(ctx *AVBSFContext) int32 {
	return (int32)(C.av_bsf_init((*C.struct_AVBSFContext)(ctx)))
}

// AvBsfSendPacket submits a packet for filtering.
func AvBsfSendPacket(ctx *AVBSFContext, pkt *AVPacket) int32 {
	return (int32)(C.av_bsf_send_packet((*C.struct_AVBSFContext)(ctx),
		(*C.struct_AVPacket)(pkt)))
}

// AvBsfReceivePacket retrieves a filtered packet.
func AvBsfReceivePacket(ctx *AVBSFContext, pkt *AVPacket) int32 {
	return (int32)(C.av_bsf_receive_packet((*C.struct_AVBSFContext)(ctx),
		(*C.struct_AVPacket)(pkt)))
}

// AvBsfFlush resets the internal bitstream filter state. Should be called e.g. when seeking.
func AvBsfFlush(ctx *AVBSFContext) {
	C.av_bsf_flush((*C.struct_AVBSFContext)(ctx))
}

// AvBsfFree frees a bitstream filter context.
func AvBsfFree(ctx **AVBSFContext) {
	C.av_bsf_free((**C.struct_AVBSFContext)(unsafe.Pointer(ctx)))
}

// AvBsfGetClass gets the AVClass for AVBSFContext.
func AvBsfGetClass() *AVClass {
	return (*AVClass)(C.av_bsf_get_class())
}

// AVBSFList
type AVBSFList C.struct_AVBSFList

// AvBsfListAlloc allocates empty list of bitstream filters.
func AvBsfListAlloc() *AVBSFList {
	return (*AVBSFList)(C.av_bsf_list_alloc())
}

// AvBsfListFree frees list of bitstream filters.
func AvBsfListFree(lst **AVBSFList) {
	C.av_bsf_list_free((**C.struct_AVBSFList)(unsafe.Pointer(lst)))
}

// AvBsfListAppend appends bitstream filter to the list of bitstream filters.
func AvBsfListAppend(lst *AVBSFList, bsf *AVBSFContext) {
	C.av_bsf_list_append((*C.struct_AVBSFList)(lst),
		(*C.struct_AVBSFContext)(bsf))
}

// AvBsfListAppend2
func AvBsfListAppend2(lst *AVBSFList, bsfName string, options **AVDictionary) {
	bsfNamePtr, bsfNameFunc := StringCasting(bsfName)
	defer bsfNameFunc()
	C.av_bsf_list_append2((*C.struct_AVBSFList)(lst),
		(*C.char)(bsfNamePtr), (**C.struct_AVDictionary)(unsafe.Pointer(options)))
}

// AvBsfListFinalize finalizes list of bitstream filters.
func AvBsfListFinalize(lst **AVBSFList, bsf **AVBSFContext) int32 {
	return (int32)(C.av_bsf_list_finalize((**C.struct_AVBSFList)(unsafe.Pointer(lst)),
		(**C.struct_AVBSFContext)(unsafe.Pointer(bsf))))
}

// AvBsfListParseStr parses string describing list of bitstream filters and creates single
// AvBSFContext describing the whole chain of bitstream filters.
func AvBsfListParseStr(str string, bsf **AVBSFContext) {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	C.av_bsf_list_parse_str((*C.char)(strPtr), (**C.struct_AVBSFContext)(unsafe.Pointer(bsf)))
}

// AvBsfGetNullFilter gets null/pass-through bitstream filter.
func AvBsfGetNullFilter(bsf **AVBSFContext) int32 {
	return (int32)(C.av_bsf_get_null_filter((**C.struct_AVBSFContext)(unsafe.Pointer(bsf))))
}
