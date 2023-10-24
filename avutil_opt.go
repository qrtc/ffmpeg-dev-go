package ffmpeg

/*
#include <libavutil/opt.h>

int64_t get_av_option_default_val_i64(AVOption *opt) {
	return opt->default_val.i64;
}

void set_av_option_default_val_i64(AVOption *opt, int64_t v) {
	opt->default_val.i64 = v;
}

int64_t* get_av_option_default_val_i64_addr(AVOption *opt) {
	return &opt->default_val.i64;
}

double get_av_option_default_val_dbl(AVOption *opt) {
	return opt->default_val.dbl;
}

void set_av_option_default_val_dbl(AVOption *opt, double v) {
	opt->default_val.dbl = v;
}

double* get_av_option_default_val_dbl_addr(AVOption *opt) {
	return &opt->default_val.dbl;
}

const char* get_av_option_default_val_str(AVOption *opt) {
	return opt->default_val.str;
}

void set_av_option_default_val_str(AVOption *opt, const char* v) {
	opt->default_val.str = v;
}

const char** get_av_option_default_val_str_addr(AVOption *opt) {
	return &opt->default_val.str;
}

AVRational get_av_option_default_val_q(AVOption *opt) {
	return opt->default_val.q;
}

void set_av_option_default_val_q(AVOption *opt, AVRational v) {
	opt->default_val.q = v;
}

AVRational* get_av_option_default_val_q_addr(AVOption *opt) {
	return &opt->default_val.q;
}

int av_opt_set_int_list_wrap(void *obj, const char *name, void *val, uint64_t term, int flags, int size) {
	if (av_int_list_length(val, term) > INT_MAX / size) {
		return AVERROR(EINVAL);
	}
	return av_opt_set_bin(obj, name, (const uint8_t *)val, av_int_list_length(val, term) * size , flags);
}
*/
import "C"
import (
	"unsafe"
)

// AVOptionType
type AVOptionType = int32

const (
	AV_OPT_TYPE_FLAGS          = AVOptionType(C.AV_OPT_TYPE_FLAGS)
	AV_OPT_TYPE_INT            = AVOptionType(C.AV_OPT_TYPE_INT)
	AV_OPT_TYPE_INT64          = AVOptionType(C.AV_OPT_TYPE_INT64)
	AV_OPT_TYPE_DOUBLE         = AVOptionType(C.AV_OPT_TYPE_DOUBLE)
	AV_OPT_TYPE_FLOAT          = AVOptionType(C.AV_OPT_TYPE_FLOAT)
	AV_OPT_TYPE_STRING         = AVOptionType(C.AV_OPT_TYPE_STRING)
	AV_OPT_TYPE_RATIONAL       = AVOptionType(C.AV_OPT_TYPE_RATIONAL)
	AV_OPT_TYPE_BINARY         = AVOptionType(C.AV_OPT_TYPE_BINARY)
	AV_OPT_TYPE_DICT           = AVOptionType(C.AV_OPT_TYPE_DICT)
	AV_OPT_TYPE_UINT64         = AVOptionType(C.AV_OPT_TYPE_UINT64)
	AV_OPT_TYPE_CONST          = AVOptionType(C.AV_OPT_TYPE_CONST)
	AV_OPT_TYPE_IMAGE_SIZE     = AVOptionType(C.AV_OPT_TYPE_IMAGE_SIZE)
	AV_OPT_TYPE_PIXEL_FMT      = AVOptionType(C.AV_OPT_TYPE_PIXEL_FMT)
	AV_OPT_TYPE_SAMPLE_FMT     = AVOptionType(C.AV_OPT_TYPE_SAMPLE_FMT)
	AV_OPT_TYPE_VIDEO_RATE     = AVOptionType(C.AV_OPT_TYPE_VIDEO_RATE)
	AV_OPT_TYPE_DURATION       = AVOptionType(C.AV_OPT_TYPE_DURATION)
	AV_OPT_TYPE_COLOR          = AVOptionType(C.AV_OPT_TYPE_COLOR)
	AV_OPT_TYPE_CHANNEL_LAYOUT = AVOptionType(C.AV_OPT_TYPE_CHANNEL_LAYOUT)
	AV_OPT_TYPE_BOOL           = AVOptionType(C.AV_OPT_TYPE_BOOL)
)

// AVOption
type AVOption C.struct_AVOption

// Custom: GetName gets `AVOption.name` value.
func (opt *AVOption) GetName() string {
	return C.GoString(opt.name)
}

// Custom: GetHelp gets `AVOption.help` value.
func (opt *AVOption) GetHelp() string {
	return C.GoString(opt.help)
}

// Custom: GetOffset gets `AVOption.offset` value.
func (opt *AVOption) GetOffset() int32 {
	return (int32)(opt.offset)
}

// Custom: SetOffset sets `AVOption.offset` value.
func (opt *AVOption) SetOffset(v int32) {
	opt.offset = (C.int)(v)
}

// Custom: GetOffsetAddr gets `AVOption.offset` address.
func (opt *AVOption) GetOffsetAddr() *int32 {
	return (*int32)(&opt.offset)
}

// Custom: GetType gets `AVOption.type` value.
func (opt *AVOption) GetType() AVOptionType {
	return (AVOptionType)(opt._type)
}

// Custom: SetType sets `AVOption.type` value.
func (opt *AVOption) SetType(v AVOptionType) {
	opt._type = (C.enum_AVOptionType)(v)
}

// Custom: GetTypeAddr gets `AVOption.type` address.
func (opt *AVOption) GetTypeAddr() *AVOptionType {
	return (*AVOptionType)(unsafe.Pointer(&opt._type))
}

// Custom: GetDefaultValI64 gets `AVOption.default_val.i64` value.
func (opt *AVOption) GetDefaultValI64() int64 {
	return (int64)(C.get_av_option_default_val_i64((*C.struct_AVOption)(opt)))
}

// Custom: SetDefaultValI64 sets `AVOption.default_val.i64` value.
func (opt *AVOption) SetDefaultValI64(v int64) {
	C.set_av_option_default_val_i64((*C.struct_AVOption)(opt), (C.int64_t)(v))
}

// Custom: GetDefaultValI64Addr gets `AVOption.default_val.i64` address.
func (opt *AVOption) GetDefaultValI64Addr() *int64 {
	return (*int64)(C.get_av_option_default_val_i64_addr((*C.struct_AVOption)(opt)))
}

// Custom: GetDefaultValDbl gets `AVOption.default_val.dbl` value.
func (opt *AVOption) GetDefaultValDbl() float64 {
	return (float64)(C.get_av_option_default_val_dbl((*C.struct_AVOption)(opt)))
}

// Custom: SetDefaultValDbl sets `AVOption.default_val.dbl` value.
func (opt *AVOption) SetDefaultValDbl(v float64) {
	C.set_av_option_default_val_dbl((*C.struct_AVOption)(opt), (C.double)(v))
}

// Custom: GetDefaultValDblAddr gets `AVOption.default_val.dbl` address.
func (opt *AVOption) GetDefaultValDblAddr() *float64 {
	return (*float64)(C.get_av_option_default_val_dbl_addr((*C.struct_AVOption)(opt)))
}

// Custom: GetDefaultValStr gets `AVOption.default_val.str` value.
func (opt *AVOption) GetDefaultValStr() string {
	return C.GoString(C.get_av_option_default_val_str((*C.struct_AVOption)(opt)))
}

// Custom: SetDefaultValStr sets `AVOption.default_val.str` value.
func (opt *AVOption) SetDefaultValStr(v *int8) {
	C.set_av_option_default_val_str((*C.struct_AVOption)(opt), (*C.char)(v))
}

// Custom: GetDefaultValStrAddr gets `AVOption.default_val.str` address.
func (opt *AVOption) GetDefaultValStrAddr() **int8 {
	return (**int8)(unsafe.Pointer(C.get_av_option_default_val_str_addr((*C.struct_AVOption)(opt))))
}

// Custom: GetDefaultValQ gets `AVOption.default_val.q` value.
func (opt *AVOption) GetDefaultValQ() AVRational {
	return (AVRational)(C.get_av_option_default_val_q((*C.struct_AVOption)(opt)))
}

// Custom: SetDefaultValQ sets `AVOption.default_val.q` value.
func (opt *AVOption) SetDefaultValQ(v AVRational) {
	C.set_av_option_default_val_q((*C.struct_AVOption)(opt), (C.struct_AVRational)(v))
}

// Custom: GetDefaultValQAddr gets `AVOption.default_val.q` address.
func (opt *AVOption) GetDefaultValQAddr() *AVRational {
	return (*AVRational)(C.get_av_option_default_val_q_addr((*C.struct_AVOption)(opt)))
}

// Custom: GetMin gets `AVOption.min` value.
func (opt *AVOption) GetMin() float64 {
	return (float64)(opt.min)
}

// Custom: SetMin sets `AVOption.min` value.
func (opt *AVOption) SetMin(v float64) {
	opt.min = (C.double)(v)
}

// Custom: GetMinAddr gets `AVOption.min` address.
func (opt *AVOption) GetMinAddr() *float64 {
	return (*float64)(&opt.min)
}

// Custom: GetMax gets `AVOption.max` value.
func (opt *AVOption) GetMax() float64 {
	return (float64)(opt.max)
}

// Custom: SetMax sets `AVOption.max` value.
func (opt *AVOption) SetMax(v float64) {
	opt.max = (C.double)(v)
}

// Custom: GetMaxAddr gets `AVOption.max` address.
func (opt *AVOption) GetMaxAddr() *float64 {
	return (*float64)(&opt.max)
}

// Custom: GetFlags gets `AVOption.flags` value.
func (opt *AVOption) GetFlags() int32 {
	return (int32)(opt.flags)
}

// Custom: SetFlags sets `AVOption.flags` value.
func (opt *AVOption) SetFlags(v int32) {
	opt.flags = (C.int)(v)
}

// Custom: GetFlagsAddr gets `AVOption.flags` address.
func (opt *AVOption) GetFlagsAddr() *int32 {
	return (*int32)(&opt.flags)
}

const (
	AV_OPT_FLAG_ENCODING_PARAM = int32(C.AV_OPT_FLAG_ENCODING_PARAM)
	AV_OPT_FLAG_DECODING_PARAM = int32(C.AV_OPT_FLAG_DECODING_PARAM)
	AV_OPT_FLAG_AUDIO_PARAM    = int32(C.AV_OPT_FLAG_AUDIO_PARAM)
	AV_OPT_FLAG_VIDEO_PARAM    = int32(C.AV_OPT_FLAG_VIDEO_PARAM)
	AV_OPT_FLAG_SUBTITLE_PARAM = int32(C.AV_OPT_FLAG_SUBTITLE_PARAM)

	AV_OPT_FLAG_EXPORT = int32(C.AV_OPT_FLAG_EXPORT)

	AV_OPT_FLAG_READONLY        = int32(C.AV_OPT_FLAG_READONLY)
	AV_OPT_FLAG_BSF_PARAM       = int32(C.AV_OPT_FLAG_BSF_PARAM)
	AV_OPT_FLAG_RUNTIME_PARAM   = int32(C.AV_OPT_FLAG_RUNTIME_PARAM)
	AV_OPT_FLAG_FILTERING_PARAM = int32(C.AV_OPT_FLAG_FILTERING_PARAM)
	AV_OPT_FLAG_DEPRECATED      = int32(C.AV_OPT_FLAG_DEPRECATED)
	AV_OPT_FLAG_CHILD_CONSTS    = int32(C.AV_OPT_FLAG_CHILD_CONSTS)
)

// Custom: GetUnit gets `AVOption.unit` value.
func (opt *AVOption) GetUnit() string {
	return C.GoString(opt.unit)
}

// AVOptionRange
type AVOptionRange C.struct_AVOptionRange

// Custom: GetStr gets `AVOptionRange.str` value.
func (optr *AVOptionRange) GetStr() string {
	return C.GoString(optr.str)
}

// Custom: GetValueMin gets `AVOptionRange.value_min` value.
func (optr *AVOptionRange) GetValueMin() float64 {
	return (float64)(optr.value_min)
}

// Custom: SetValueMin sets `AVOptionRange.value_min` value.
func (optr *AVOptionRange) SetValueMin(v float64) {
	optr.value_min = (C.double)(v)
}

// Custom: GetValueMinAddr gets `AVOptionRange.value_min` address.
func (optr *AVOptionRange) GetValueMinAddr() *float64 {
	return (*float64)(&optr.value_min)
}

// Custom: GetValueMax gets `AVOptionRange.value_max` value.
func (optr *AVOptionRange) GetValueMax() float64 {
	return (float64)(optr.value_max)
}

// Custom: SetValueMax sets `AVOptionRange.value_max` value.
func (optr *AVOptionRange) SetValueMax(v float64) {
	optr.value_max = (C.double)(v)
}

// Custom: GetValueMaxAddr gets `AVOptionRange.value_max` address.
func (optr *AVOptionRange) GetValueMaxAddr() *float64 {
	return (*float64)(&optr.value_max)
}

// Custom: GetComponentMin gets `AVOptionRange.component_min` value.
func (optr *AVOptionRange) GetComponentMin() float64 {
	return (float64)(optr.component_min)
}

// Custom: SetComponentMin sets `AVOptionRange.component_min` value.
func (optr *AVOptionRange) SetComponentMin(v float64) {
	optr.component_min = (C.double)(v)
}

// Custom: GetComponentMinAddr gets `AVOptionRange.component_min` address.
func (optr *AVOptionRange) GetComponentMinAddr() *float64 {
	return (*float64)(&optr.component_min)
}

// Custom: GetComponentMax gets `AVOptionRange.component_max` value.
func (optr *AVOptionRange) GetComponentMax() float64 {
	return (float64)(optr.component_max)
}

// Custom: SetComponentMax sets `AVOptionRange.component_max` value.
func (optr *AVOptionRange) SetComponentMax(v float64) {
	optr.component_max = (C.double)(v)
}

// Custom: GetComponentMaxAddr gets `AVOptionRange.component_max` address.
func (optr *AVOptionRange) GetComponentMaxAddr() *float64 {
	return (*float64)(&optr.component_max)
}

// Custom: GetIsRange gets `AVOptionRange.isrange` value.
func (optr *AVOptionRange) GetIsRange() int32 {
	return (int32)(optr.is_range)
}

// Custom: SetIsRange sets `AVOptionRange.isrange` value.
func (optr *AVOptionRange) SetIsRange(v int32) {
	optr.is_range = (C.int)(v)
}

// Custom: GetIsRangeAddr gets `AVOptionRange.isrange` address.
func (optr *AVOptionRange) GetIsRangeAddr() *int32 {
	return (*int32)(&optr.is_range)
}

// AVOptionRanges
type AVOptionRanges C.struct_AVOptionRanges

// Custom: GetRange gets `AVOptionRanges.range` value.
func (optrs *AVOptionRanges) GetRange() []*AVOptionRange {
	if optrs._range == nil {
		return nil
	}
	return unsafe.Slice((**AVOptionRange)(unsafe.Pointer(optrs._range)),
		optrs.nb_components*optrs.nb_ranges)
}

// Custom: GetNbRanges gets `AVOptionRanges.nb_ranges` value.
func (optrs *AVOptionRanges) GetNbRanges() int32 {
	return (int32)(optrs.nb_ranges)
}

// Custom: GetNbComponents gets `AVOptionRanges.nb_components` value.
func (optrs *AVOptionRanges) GetNbComponents() int32 {
	return (int32)(optrs.nb_components)
}

// AvOptShow2 shows the obj options.
func AvOptShow2(obj, avLogObj CVoidPointer, reqFlags, rejFlags int32) int32 {
	return (int32)(C.av_opt_show2(VoidPointer(obj), VoidPointer(avLogObj), (C.int)(reqFlags), (C.int)(rejFlags)))
}

// AvOptSetDefaults sets the values of all AVOption fields to their default values.
func AvOptSetDefaults(s CVoidPointer) {
	C.av_opt_set_defaults(VoidPointer(s))
}

// AvOptSetDefaults2 sets the values of all AVOption fields to their default values.
func AvOptSetDefaults2(s CVoidPointer, mask, flags int32) {
	C.av_opt_set_defaults2(VoidPointer(s), (C.int)(mask), (C.int)(flags))
}

// AvSetOptionsString parses the key/value pairs list in opts. For each key/value pair
// found, stores the value in the field in ctx that is named like the
// key. ctx must be an AVClass context, storing is done using AVOptions.
func AvSetOptionsString(ctx CVoidPointer, opts, keyValSep, pairsSep string) int32 {
	optsPtr, optsFunc := StringCasting(opts)
	defer optsFunc()
	keyValSepPtr, keyValSepFunc := StringCasting(keyValSep)
	defer keyValSepFunc()
	pairsSepPtr, pairsSepFunc := StringCasting(pairsSep)
	defer pairsSepFunc()
	return (int32)(C.av_set_options_string(VoidPointer(ctx), (*C.char)(optsPtr),
		(*C.char)(keyValSepPtr), (*C.char)(pairsSepPtr)))
}

// NONEED: av_opt_set_from_string

// AvOptFree frees all allocated objects in obj.
func AvOptFree(obj CVoidPointer) {
	C.av_opt_free(VoidPointer(obj))
}

// AvOptFlagIsSet checks whether a particular flag is set in a flags field.
func AvOptFlagIsSet(obj CVoidPointer, fieldName, flagName string) int32 {
	fieldNamePtr, fieldNameFunc := StringCasting(fieldName)
	defer fieldNameFunc()
	flagNamePtr, flagNameFunc := StringCasting(flagName)
	defer flagNameFunc()
	return (int32)(C.av_opt_flag_is_set(VoidPointer(obj), (*C.char)(fieldNamePtr), (*C.char)(flagNamePtr)))
}

// AvOptSetDict sets all the options from a given dictionary on an object.
func AvOptSetDict(obj CVoidPointer, options **AVDictionary) int32 {
	return (int32)(C.av_opt_set_dict(VoidPointer(obj), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvOptSetDict2 sets all the options from a given dictionary on an object.
func AvOptSetDict2(obj CVoidPointer, options **AVDictionary, searchFlags int32) int32 {
	return (int32)(C.av_opt_set_dict2(VoidPointer(obj), (**C.struct_AVDictionary)(unsafe.Pointer(options)),
		(C.int)(searchFlags)))
}

// NONEED: av_opt_get_key_value

const (
	AV_OPT_FLAG_IMPLICIT_KEY = int32(C.AV_OPT_FLAG_IMPLICIT_KEY)
)

// AvOptEvalFlags
func AvOptEvalFlags(obj CVoidPointer, o *AVOption, val string, flags_out *int32) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_flags(VoidPointer(obj), (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.int)(flags_out)))
}

// AvOptEvalInt
func AvOptEvalInt(obj CVoidPointer, o *AVOption, val string, int_out *int32) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_int(VoidPointer(obj), (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.int)(int_out)))
}

// AvOptEvalInt64
func AvOptEvalInt64(obj CVoidPointer, o *AVOption, val string, int64_out *int64) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_int64(VoidPointer(obj), (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.int64_t)(int64_out)))
}

// AvOptEvalFloat
func AvOptEvalFloat(obj CVoidPointer, o *AVOption, val string, float_out *float32) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_float(VoidPointer(obj), (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.float)(float_out)))
}

// AvOptEvalDouble
func AvOptEvalDouble(obj CVoidPointer, o *AVOption, val string, double_out *float64) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_double(VoidPointer(obj), (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.double)(double_out)))
}

// AvOptEvalQ
func AvOptEvalQ(obj CVoidPointer, o *AVOption, val string, q_out *AVRational) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_q(VoidPointer(obj), (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.struct_AVRational)(q_out)))
}

const (
	AV_OPT_SEARCH_CHILDREN       = C.AV_OPT_SEARCH_CHILDREN
	AV_OPT_SEARCH_FAKE_OBJ       = C.AV_OPT_SEARCH_FAKE_OBJ
	AV_OPT_ALLOW_NULL            = C.AV_OPT_ALLOW_NULL
	AV_OPT_MULTI_COMPONENT_RANGE = C.AV_OPT_MULTI_COMPONENT_RANGE
)

// AvOptFind looks for an option in an object. Consider only options which
// have all the specified flags set.
func AvOptFind(obj CVoidPointer, name, unit string, optFlags, searchFlags int32) *AVOption {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	unitPtr, unitFunc := StringCasting(unit)
	defer unitFunc()
	return (*AVOption)(C.av_opt_find(VoidPointer(obj), (*C.char)(namePtr), (*C.char)(unitPtr),
		(C.int)(optFlags), (C.int)(searchFlags)))
}

// AvOptFind2 looks for an option in an object. Consider only options which
// have all the specified flags set.
func AvOptFind2(obj CVoidPointer, name, unit string, optFlags, searchFlags int32,
	targetObj CVoidPointerPointer) *AVOption {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	unitPtr, unitFunc := StringCasting(unit)
	defer unitFunc()
	return (*AVOption)(C.av_opt_find2(VoidPointer(obj), (*C.char)(namePtr), (*C.char)(unitPtr),
		(C.int)(optFlags), (C.int)(searchFlags), VoidPointerPointer(targetObj)))
}

// AvOptNext iterates over all AVOptions belonging to obj.
func AvOptNext(obj CVoidPointer, prev *AVOption) *AVOption {
	return (*AVOption)(C.av_opt_next(VoidPointer(obj), (*C.struct_AVOption)(prev)))
}

// AvOptChildNext iterates over AVOptions-enabled children of obj.
func AvOptChildNext(obj, prev CVoidPointer) unsafe.Pointer {
	return C.av_opt_child_next(VoidPointer(obj), VoidPointer(prev))
}

// Deprecated: Use AvOptChildClassIterate instead.
func AvOptChildClassNext(parent, prev *AVClass) *AVClass {
	return (*AVClass)(C.av_opt_child_class_next((*C.struct_AVClass)(parent),
		(*C.struct_AVClass)(prev)))
}

// AvOptChildClassIterate iterates over potential AVOptions-enabled children of parent.
func AvOptChildClassIterate(parent *AVClass, iter CVoidPointerPointer) *AVClass {
	return (*AVClass)(C.av_opt_child_class_iterate((*C.struct_AVClass)(parent), VoidPointerPointer(iter)))
}

// AvOptSet
func AvOptSet(obj CVoidPointer, name string, val string, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_set(VoidPointer(obj), (*C.char)(namePtr), (*C.char)(valPtr), (C.int)(searchFlags)))
}

// AvOptSetInt
func AvOptSetInt(obj CVoidPointer, name string, val int64, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_int(VoidPointer(obj), (*C.char)(namePtr), (C.int64_t)(val), (C.int)(searchFlags)))
}

// AvOptSetDouble
func AvOptSetDouble(obj CVoidPointer, name string, val float64, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_double(VoidPointer(obj), (*C.char)(namePtr), (C.double)(val), (C.int)(searchFlags)))
}

// AvOptSetQ
func AvOptSetQ(obj CVoidPointer, name string, val AVRational, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_q(VoidPointer(obj), (*C.char)(namePtr), (C.struct_AVRational)(val), (C.int)(searchFlags)))
}

// AvOptSetBin
func AvOptSetBin(obj CVoidPointer, name string, val *uint8, size int32, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_bin(VoidPointer(obj), (*C.char)(namePtr), (*C.uint8_t)(val), (C.int)(size), (C.int)(searchFlags)))
}

// AvOptSetImageSize
func AvOptSetImageSize(obj CVoidPointer, name string, w, h int32, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_image_size(VoidPointer(obj), (*C.char)(namePtr), (C.int)(w), (C.int)(h), (C.int)(searchFlags)))
}

// AvOptSetPixelFmt
func AvOptSetPixelFmt(obj CVoidPointer, name string, fmt AVPixelFormat, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_pixel_fmt(VoidPointer(obj), (*C.char)(namePtr), (C.enum_AVPixelFormat)(fmt), (C.int)(searchFlags)))
}

// AvOptSetSampleFmt
func AvOptSetSampleFmt(obj CVoidPointer, name string, fmt AVSampleFormat, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_sample_fmt(VoidPointer(obj), (*C.char)(namePtr), (C.enum_AVSampleFormat)(fmt), (C.int)(searchFlags)))
}

// AvOptSetVideoRate
func AvOptSetVideoRate(obj CVoidPointer, name string, val AVRational, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_video_rate(VoidPointer(obj), (*C.char)(namePtr), (C.struct_AVRational)(val), (C.int)(searchFlags)))
}

// AvOptSetChannelLayout
func AvOptSetChannelLayout(obj CVoidPointer, name string, chLayout int64, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_channel_layout(VoidPointer(obj), (*C.char)(namePtr), (C.int64_t)(chLayout), (C.int)(searchFlags)))
}

// AvOptSetDictVal
func AvOptSetDictVal(obj CVoidPointer, name string, val *AVDictionary, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_dict_val(VoidPointer(obj), (*C.char)(namePtr), (*C.struct_AVDictionary)(val), (C.int)(searchFlags)))
}

// AvOptSetIntList sets a binary option to an integer list.
func AvOptSetIntList[T HelperInteger](obj CVoidPointer, name string,
	val *T, term uint64, flags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	size := (int32)(unsafe.Sizeof(*val))
	return (int32)(C.av_opt_set_int_list_wrap(VoidPointer(obj), (*C.char)(namePtr),
		unsafe.Pointer(val), (C.uint64_t)(term), (C.int)(flags), (C.int)(size)))
}

// AvOptGet
func AvOptGet(obj CVoidPointer, name string, searchFlags int32, outVal **uint8) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags),
		(**C.uint8_t)(unsafe.Pointer(outVal))))
}

// AvOptGetInt
func AvOptGetInt(obj CVoidPointer, name string, searchFlags int32, outVal *int64) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_int(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags), (*C.int64_t)(outVal)))
}

// AvOptGetDouble
func AvOptGetDouble(obj CVoidPointer, name string, searchFlags int32, outVal *float64) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_double(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags), (*C.double)(outVal)))
}

// AvOptGetQ
func AvOptGetQ(obj CVoidPointer, name string, searchFlags int32, outVal *AVRational) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_q(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags), (*C.struct_AVRational)(outVal)))
}

// AvOptGetImageSize
func AvOptGetImageSize(obj CVoidPointer, name string, searchFlags int32, wOut, hOut *int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_image_size(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags), (*C.int)(wOut), (*C.int)(hOut)))
}

// AvOptGetPixelFmt
func AvOptGetPixelFmt(obj CVoidPointer, name string, searchFlags int32, outFmt *AVPixelFormat) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_pixel_fmt(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags), (*C.enum_AVPixelFormat)(outFmt)))
}

// AvOptGetSampleFmt
func AvOptGetSampleFmt(obj CVoidPointer, name string, searchFlags int32, outFmt *AVSampleFormat) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_sample_fmt(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags), (*C.enum_AVSampleFormat)(outFmt)))
}

// AvOptGetVideoRate
func AvOptGetVideoRate(obj CVoidPointer, name string, searchFlags int32, outVal *AVRational) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_video_rate(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags), (*C.struct_AVRational)(outVal)))
}

// AvOptGetChannelLayout
func AvOptGetChannelLayout(obj CVoidPointer, name string, searchFlags int32, outVal *int64) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_channel_layout(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags), (*C.int64_t)(outVal)))
}

// AvOptGetDictVal
func AvOptGetDictVal(obj CVoidPointer, name string, searchFlags int32, outVal **AVDictionary) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_dict_val(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags),
		(**C.struct_AVDictionary)(unsafe.Pointer(outVal))))
}

// AvOptPtr gets a pointer to the requested field in a struct.
func AvOptPtr(avclass *AVClass, obj CVoidPointer, name string) {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	C.av_opt_ptr((*C.struct_AVClass)(avclass), VoidPointer(obj), (*C.char)(namePtr))
}

// AvOptFreepRanges frees an AVOptionRanges struct and set it to NULL.
func AvOptFreepRanges(ranges **AVOptionRanges) {
	C.av_opt_freep_ranges((**C.struct_AVOptionRanges)(unsafe.Pointer(ranges)))
}

// AvOptQueryRanges gets a list of allowed ranges for the given option.
func AvOptQueryRanges(ranges **AVOptionRanges, obj CVoidPointer, key string, flags int32) int32 {
	keyPtr, keyFunc := StringCasting(key)
	defer keyFunc()
	return (int32)(C.av_opt_query_ranges((**C.struct_AVOptionRanges)(unsafe.Pointer(ranges)),
		VoidPointer(obj), (*C.char)(keyPtr), (C.int)(flags)))
}

// AvOptCopy copies options from src object into dest object.
func AvOptCopy(dest, src CVoidPointer) int32 {
	return (int32)(C.av_opt_copy(VoidPointer(dest), VoidPointer(src)))
}

// AvOptQueryRangesDefault gets a default list of allowed ranges for the given option.
func AvOptQueryRangesDefault(ranges **AVOptionRanges, obj CVoidPointer, key string, flags int32) int32 {
	keyPtr, keyFunc := StringCasting(key)
	defer keyFunc()
	return (int32)(C.av_opt_query_ranges_default((**C.struct_AVOptionRanges)(unsafe.Pointer(ranges)),
		VoidPointer(obj), (*C.char)(keyPtr), (C.int)(flags)))
}

// AvOptIsSetToDefault checks if given option is set to its default value.
func AvOptIsSetToDefault(obj CVoidPointer, o *AVOption) int32 {
	return (int32)(C.av_opt_is_set_to_default(VoidPointer(obj), (*C.struct_AVOption)(o)))
}

// AvOptIsSetToDefaultByName checks if given option is set to its default value.
func AvOptIsSetToDefaultByName(obj CVoidPointer, name string, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_is_set_to_default_by_name(VoidPointer(obj), (*C.char)(namePtr), (C.int)(searchFlags)))
}

const (
	AV_OPT_SERIALIZE_SKIP_DEFAULTS   = int32(C.AV_OPT_SERIALIZE_SKIP_DEFAULTS)
	AV_OPT_SERIALIZE_OPT_FLAGS_EXACT = int32(C.AV_OPT_SERIALIZE_OPT_FLAGS_EXACT)
)

// AvOptSerialize serializes object's options.
func AvOptSerialize(obj CVoidPointer, optFlags, flags int32, keyValSep, pairsSep string) (output string, ret int32) {
	var buffer *C.char
	keyValSepPtr, keyValSepFunc := StringCasting(keyValSep)
	defer keyValSepFunc()
	pairsSepPtr, pairsSepFunc := StringCasting(pairsSep)
	defer pairsSepFunc()
	ret = (int32)(C.av_opt_serialize(VoidPointer(obj), (C.int)(optFlags), (C.int)(flags),
		(**C.char)(unsafe.Pointer(&buffer)), (C.char)(*keyValSepPtr), (C.char)(*pairsSepPtr)))
	return C.GoString(buffer), ret
}
