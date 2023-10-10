package ffmpeg

/*
#include <libavutil/opt.h>
*/
import "C"
import "unsafe"

type AvOptionType int32

const (
	AV_OPT_TYPE_FLAGS          = int32(C.AV_OPT_TYPE_FLAGS)
	AV_OPT_TYPE_INT            = int32(C.AV_OPT_TYPE_INT)
	AV_OPT_TYPE_INT64          = int32(C.AV_OPT_TYPE_INT64)
	AV_OPT_TYPE_DOUBLE         = int32(C.AV_OPT_TYPE_DOUBLE)
	AV_OPT_TYPE_FLOAT          = int32(C.AV_OPT_TYPE_FLOAT)
	AV_OPT_TYPE_STRING         = int32(C.AV_OPT_TYPE_STRING)
	AV_OPT_TYPE_RATIONAL       = int32(C.AV_OPT_TYPE_RATIONAL)
	AV_OPT_TYPE_BINARY         = int32(C.AV_OPT_TYPE_BINARY)
	AV_OPT_TYPE_DICT           = int32(C.AV_OPT_TYPE_DICT)
	AV_OPT_TYPE_UINT64         = int32(C.AV_OPT_TYPE_UINT64)
	AV_OPT_TYPE_CONST          = int32(C.AV_OPT_TYPE_CONST)
	AV_OPT_TYPE_IMAGE_SIZE     = int32(C.AV_OPT_TYPE_IMAGE_SIZE)
	AV_OPT_TYPE_PIXEL_FMT      = int32(C.AV_OPT_TYPE_PIXEL_FMT)
	AV_OPT_TYPE_SAMPLE_FMT     = int32(C.AV_OPT_TYPE_SAMPLE_FMT)
	AV_OPT_TYPE_VIDEO_RATE     = int32(C.AV_OPT_TYPE_VIDEO_RATE)
	AV_OPT_TYPE_DURATION       = int32(C.AV_OPT_TYPE_DURATION)
	AV_OPT_TYPE_COLOR          = int32(C.AV_OPT_TYPE_COLOR)
	AV_OPT_TYPE_CHANNEL_LAYOUT = int32(C.AV_OPT_TYPE_CHANNEL_LAYOUT)
	AV_OPT_TYPE_BOOL           = int32(C.AV_OPT_TYPE_BOOL)
)

type AvOption C.struct_AVOption

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

type AvOptionRange C.struct_AVOptionRange

type AvOptionRanges C.struct_AVOptionRanges

// AvOptShow2 shows the obj options.
func AvOptShow2(obj, avLogObj unsafe.Pointer, reqFlags, rejFlags int32) int32 {
	return (int32)(C.av_opt_show2(obj, avLogObj, (C.int)(reqFlags), (C.int)(rejFlags)))
}

// AvOptSetDefaults sets the values of all AVOption fields to their default values.
func AvOptSetDefaults(s unsafe.Pointer) {
	C.av_opt_set_defaults(s)
}

// AvOptSetDefaults2 sets the values of all AVOption fields to their default values.
func AvOptSetDefaults2(s unsafe.Pointer, mask, flags int32) {
	C.av_opt_set_defaults2(s, (C.int)(mask), (C.int)(flags))
}

// AvSetOptionsString parses the key/value pairs list in opts. For each key/value pair
// found, stores the value in the field in ctx that is named like the
// key. ctx must be an AVClass context, storing is done using AVOptions.
func AvSetOptionsString(ctx unsafe.Pointer, opts, keyValSep, pairsSep string) int32 {
	optsPtr, optsFunc := StringCasting(opts)
	defer optsFunc()
	keyValSepPtr, keyValSepFunc := StringCasting(keyValSep)
	defer keyValSepFunc()
	pairsSepPtr, pairsSepFunc := StringCasting(pairsSep)
	defer pairsSepFunc()
	return (int32)(C.av_set_options_string(ctx, (*C.char)(optsPtr),
		(*C.char)(keyValSepPtr), (*C.char)(pairsSepPtr)))
}

// TODO. av_opt_set_from_string

// AvOptFree frees all allocated objects in obj.
func AvOptFree(obj unsafe.Pointer) {
	C.av_opt_free(obj)
}

// AvOptFlagIsSet checks whether a particular flag is set in a flags field.
func AvOptFlagIsSet(obj unsafe.Pointer, fieldName, flagName string) int32 {
	fieldNamePtr, fieldNameFunc := StringCasting(fieldName)
	defer fieldNameFunc()
	flagNamePtr, flagNameFunc := StringCasting(flagName)
	defer flagNameFunc()
	return (int32)(C.av_opt_flag_is_set(obj, (*C.char)(fieldNamePtr), (*C.char)(flagNamePtr)))
}

// AvOptSetDict sets all the options from a given dictionary on an object.
func AvOptSetDict(obj unsafe.Pointer, options **AvDictionary) int32 {
	return (int32)(C.av_opt_set_dict(obj, (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvOptSetDict2 sets all the options from a given dictionary on an object.
func AvOptSetDict2(obj unsafe.Pointer, options **AvDictionary, searchFlags int32) int32 {
	return (int32)(C.av_opt_set_dict2(obj, (**C.struct_AVDictionary)(unsafe.Pointer(options)),
		(C.int)(searchFlags)))
}

// TODO. av_opt_get_key_value

const (
	AV_OPT_FLAG_IMPLICIT_KEY = int32(C.AV_OPT_FLAG_IMPLICIT_KEY)
)

// AvOptEvalFlags
func AvOptEvalFlags(obj unsafe.Pointer, o *AvOption, val string, flags_out *int32) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_flags(obj, (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.int)(flags_out)))
}

// AvOptEvalInt
func AvOptEvalInt(obj unsafe.Pointer, o *AvOption, val string, int_out *int32) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_int(obj, (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.int)(int_out)))
}

// AvOptEvalInt64
func AvOptEvalInt64(obj unsafe.Pointer, o *AvOption, val string, int64_out *int64) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_int64(obj, (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.int64_t)(int64_out)))
}

// AvOptEvalFloat
func AvOptEvalFloat(obj unsafe.Pointer, o *AvOption, val string, float_out *float32) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_float(obj, (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.float)(float_out)))
}

// AvOptEvalDouble
func AvOptEvalDouble(obj unsafe.Pointer, o *AvOption, val string, double_out *float64) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_double(obj, (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.double)(double_out)))
}

// AvOptEvalQ
func AvOptEvalQ(obj unsafe.Pointer, o *AvOption, val string, q_out *AvRational) int32 {
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_eval_q(obj, (*C.struct_AVOption)(o), (*C.char)(valPtr), (*C.struct_AVRational)(q_out)))
}

const (
	AV_OPT_SEARCH_CHILDREN       = C.AV_OPT_SEARCH_CHILDREN
	AV_OPT_SEARCH_FAKE_OBJ       = C.AV_OPT_SEARCH_FAKE_OBJ
	AV_OPT_ALLOW_NULL            = C.AV_OPT_ALLOW_NULL
	AV_OPT_MULTI_COMPONENT_RANGE = C.AV_OPT_MULTI_COMPONENT_RANGE
)

// AvOptFind looks for an option in an object. Consider only options which
// have all the specified flags set.
func AvOptFind(obj unsafe.Pointer, name, unit string, optFlags, searchFlags int32) *AvOption {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	unitPtr, unitFunc := StringCasting(unit)
	defer unitFunc()
	return (*AvOption)(C.av_opt_find(obj, namePtr, unitPtr, (C.int)(optFlags), (C.int)(searchFlags)))
}

// AvOptFind2 looks for an option in an object. Consider only options which
// have all the specified flags set.
func AvOptFind2(obj unsafe.Pointer, name, unit string, optFlags, searchFlags int32,
	targetObj *unsafe.Pointer) *AvOption {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	unitPtr, unitFunc := StringCasting(unit)
	defer unitFunc()
	return (*AvOption)(C.av_opt_find2(obj, namePtr, unitPtr,
		(C.int)(optFlags), (C.int)(searchFlags), targetObj))
}

// AvOptNext iterates over all AVOptions belonging to obj.
func AvOptNext(obj unsafe.Pointer, prev *AvOption) *AvOption {
	return (*AvOption)(C.av_opt_next(obj, (*C.struct_AVOption)(prev)))
}

// AvOptChildNext iterates over AVOptions-enabled children of obj.
func AvOptChildNext(obj, prev unsafe.Pointer) unsafe.Pointer {
	return C.av_opt_child_next(obj, prev)
}

// Deprecated: Use AvOptChildClassIterate instead.
func AvOptChildClassNext(parent, prev *AvClass) *AvClass {
	return (*AvClass)(C.av_opt_child_class_next((*C.struct_AVClass)(parent),
		(*C.struct_AVClass)(prev)))
}

// AvOptChildClassIterate iterates over potential AVOptions-enabled children of parent.
func AvOptChildClassIterate(parent *AvClass, iter *unsafe.Pointer) *AvClass {
	return (*AvClass)(C.av_opt_child_class_iterate((*C.struct_AVClass)(parent), iter))
}

// AvOptSet
func AvOptSet(obj unsafe.Pointer, name string, val string, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	valPtr, valFunc := StringCasting(val)
	defer valFunc()
	return (int32)(C.av_opt_set(obj, (*C.char)(namePtr), (*C.char)(valPtr), (C.int)(searchFlags)))
}

// AvOptSetInt
func AvOptSetInt(obj unsafe.Pointer, name string, val int64, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_int(obj, (*C.char)(namePtr), (C.int64_t)(val), (C.int)(searchFlags)))
}

// AvOptSetDouble
func AvOptSetDouble(obj unsafe.Pointer, name string, val float64, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_double(obj, (*C.char)(namePtr), (C.double)(val), (C.int)(searchFlags)))
}

// AvOptSetQ
func AvOptSetQ(obj unsafe.Pointer, name string, val AvRational, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_q(obj, (*C.char)(namePtr), (C.AVRational)(val), (C.int)(searchFlags)))
}

// AvOptSetBin
func AvOptSetBin(obj unsafe.Pointer, name string, val *uint8, size int32, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_bin(obj, (*C.char)(namePtr), (*C.uint8_t)(val), (C.int)(size), (C.int)(searchFlags)))
}

// AvOptSetImageSize
func AvOptSetImageSize(obj unsafe.Pointer, name string, w, h int32, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_image_size(obj, (*C.char)(namePtr), (C.int)(w), (C.int)(h), (C.int)(searchFlags)))
}

// AvOptSetPixelFmt
func AvOptSetPixelFmt(obj unsafe.Pointer, name string, fmt AvPixelFormat, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_pixel_fmt(obj, (*C.char)(namePtr), (C.enum_AVPixelFormat)(fmt), (C.int)(searchFlags)))
}

// AvOptSetSampleFmt
func AvOptSetSampleFmt(obj unsafe.Pointer, name string, fmt AvSampleFormat, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_sample_fmt(obj, (*C.char)(namePtr), (C.enum_AVSampleFormat)(fmt), (C.int)(searchFlags)))
}

// AvOptSetVideoRate
func AvOptSetVideoRate(obj unsafe.Pointer, name string, val AvRational, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_video_rate(obj, (*C.char)(namePtr), (C.struct_AVRational)(val), (C.int)(searchFlags)))
}

// AvOptSetChannelLayout
func AvOptSetChannelLayout(obj unsafe.Pointer, name string, chLayout int64, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_channel_layout(obj, (*C.char)(namePtr), (C.int64_t)(chLayout), (C.int)(searchFlags)))
}

// AvOptSetDictVal
func AvOptSetDictVal(obj unsafe.Pointer, name string, val *AvDictionary, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_set_dict_val(obj, (*C.char)(namePtr), (*C.struct_AVDictionary)(val), (C.int)(searchFlags)))
}

// TODO. av_opt_set_int_list

// AvOptGet
func AvOptGet(obj unsafe.Pointer, name string, searchFlags int32, outVal **uint8) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get(obj, (*C.char)(namePtr), (C.int)(searchFlags),
		(**C.uint8_t)(unsafe.Pointer(outVal))))
}

// AvOptGetInt
func AvOptGetInt(obj unsafe.Pointer, name string, searchFlags int32, outVal *int64) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_int(obj, (*C.char)(namePtr), (C.int)(searchFlags), (*C.int64_t)(outVal)))
}

// AvOptGetDouble
func AvOptGetDouble(obj unsafe.Pointer, name string, searchFlags int32, outVal *float64) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_double(obj, (*C.char)(namePtr), (C.int)(searchFlags), (*C.double)(outVal)))
}

// AvOptGetQ
func AvOptGetQ(obj unsafe.Pointer, name string, searchFlags int32, outVal *AvRational) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_q(obj, (*C.char)(namePtr), (C.int)(searchFlags), (*C.struct_AVRational)(outVal)))
}

// AvOptGetImageSize
func AvOptGetImageSize(obj unsafe.Pointer, name string, searchFlags int32, wOut, hOut *int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_image_size(obj, (*C.char)(namePtr), (C.int)(searchFlags), (*C.int)(wOut), (*C.int)(hOut)))
}

// AvOptGetPixelFmt
func AvOptGetPixelFmt(obj unsafe.Pointer, name string, searchFlags int32, outFmt *AvPixelFormat) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_pixel_fmt(obj, (*C.char)(namePtr), (C.int)(searchFlags), (*C.enum_AVPixelFormat)(outFmt)))
}

// AvOptGetSampleFmt
func AvOptGetSampleFmt(obj unsafe.Pointer, name string, searchFlags int32, outFmt *AvSampleFormat) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_sample_fmt(obj, (*C.char)(namePtr), (C.int)(searchFlags), (*C.enum_AVSampleFormat)(outFmt)))
}

// AvOptGetVideoRate
func AvOptGetVideoRate(obj unsafe.Pointer, name string, searchFlags int32, outVal *AvRational) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_video_rate(obj, (*C.char)(namePtr), (C.int)(searchFlags), (*C.struct_AVRational)(outVal)))
}

// AvOptGetChannelLayout
func AvOptGetChannelLayout(obj unsafe.Pointer, name string, searchFlags int32, outVal *int64) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_channel_layout(obj, (*C.char)(namePtr), (C.int)(searchFlags), (*C.int64_t)(outVal)))
}

// AvOptGetDictVal
func AvOptGetDictVal(obj unsafe.Pointer, name string, searchFlags int32, outVal **AvDictionary) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_get_dict_val(obj, (*C.char)(namePtr), (C.int)(searchFlags),
		(**C.struct_AVDictionary)(unsafe.Pointer(outVal))))
}

// AvOptPtr gets a pointer to the requested field in a struct.
func AvOptPtr(avclass *AvClass, obj unsafe.Pointer, name string) {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	C.av_opt_ptr((*C.struct_AVClass)(avclass), obj, (*C.char)(namePtr))
}

// AvOptFreepRanges frees an AvOptionRanges struct and set it to NULL.
func AvOptFreepRanges(ranges **AvOptionRanges) {
	C.av_opt_freep_ranges((**C.struct_AVOptionRanges)(unsafe.Pointer(ranges)))
}

// AvOptQueryRanges gets a list of allowed ranges for the given option.
func AvOptQueryRanges(ranges **AvOptionRanges, obj unsafe.Pointer, key string, flags int32) int32 {
	keyPtr, keyFunc := StringCasting(key)
	defer keyFunc()
	return (int32)(C.av_opt_query_ranges((**C.struct_AVOptionRanges)(unsafe.Pointer(ranges)),
		obj, (*C.char)(keyPtr), (C.int)(flags)))
}

// AvOptCopy copies options from src object into dest object.
func AvOptCopy(dest, src unsafe.Pointer) int32 {
	return (int32)(C.av_opt_copy(dest, src))
}

// AvOptQueryRangesDefault gets a default list of allowed ranges for the given option.
func AvOptQueryRangesDefault(ranges **AvOptionRanges, obj unsafe.Pointer, key string, flags int32) int32 {
	keyPtr, keyFunc := StringCasting(key)
	defer keyFunc()
	return (int32)(C.av_opt_query_ranges_default((**C.struct_AVOptionRanges)(unsafe.Pointer(ranges)),
		obj, (*C.char)(keyPtr), (C.int)(flags)))
}

// AvOptIsSetToDefault checks if given option is set to its default value.
func AvOptIsSetToDefault(obj unsafe.Pointer, o *AvOption) int32 {
	return (int32)(C.av_opt_is_set_to_default(obj, (*C.struct_AVOption)(o)))
}

// AvOptIsSetToDefaultByName checks if given option is set to its default value.
func AvOptIsSetToDefaultByName(obj unsafe.Pointer, name string, searchFlags int32) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_opt_is_set_to_default_by_name(obj, (*C.char)(namePtr), (C.int)(searchFlags)))
}

const (
	AV_OPT_SERIALIZE_SKIP_DEFAULTS   = int32(C.AV_OPT_SERIALIZE_SKIP_DEFAULTS)
	AV_OPT_SERIALIZE_OPT_FLAGS_EXACT = int32(C.AV_OPT_SERIALIZE_OPT_FLAGS_EXACT)
)

// AvOptSerialize serializes object's options.
func AvOptSerialize(obj unsafe.Pointer, optFlags, flags int32, keyValSep, pairsSep string) (output string, ret int32) {
	var buffer *C.char
	keyValSepPtr, keyValSepFunc := StringCasting(keyValSep)
	defer keyValSepFunc()
	pairsSepPtr, pairsSepFunc := StringCasting(pairsSep)
	defer pairsSepFunc()
	ret = (int32)(C.av_opt_serialize(obj, (C.int)(optFlags), (C.int)(flags),
		(**C.char)(unsafe.Pointer(&buffer)), (C.char)(*keyValSepPtr), (C.char)(*pairsSepPtr)))
	return C.GoString(buffer), ret
}
