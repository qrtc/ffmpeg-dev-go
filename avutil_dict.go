package ffmpeg

/*
#include <libavutil/dict.h>
*/
import "C"
import "unsafe"

const (
	AV_DICT_MATCH_CASE      = C.AV_DICT_MATCH_CASE
	AV_DICT_IGNORE_SUFFIX   = C.AV_DICT_IGNORE_SUFFIX
	AV_DICT_DONT_STRDUP_KEY = C.AV_DICT_DONT_STRDUP_KEY
	AV_DICT_DONT_STRDUP_VAL = C.AV_DICT_DONT_STRDUP_VAL
	AV_DICT_DONT_OVERWRITE  = C.AV_DICT_DONT_OVERWRITE
	AV_DICT_APPEND          = C.AV_DICT_APPEND
	AV_DICT_MULTIKEY        = C.AV_DICT_MULTIKEY
)

type AVDictionaryEntry C.struct_AVDictionaryEntry

// GetKey gets `AVDictionaryEntry.key` value.
func (e *AVDictionaryEntry) GetKey() string {
	return C.GoString(e.key)
}

// GetValue gets `AVDictionaryEntry.value` value.
func (e *AVDictionaryEntry) GetValue() string {
	return C.GoString(e.value)
}

// AvDictionary gets a dictionary entry with matching key.
type AVDictionary C.struct_AVDictionary

// AvDictGet
func AvDictGet(m *AVDictionary, key string, prev *AVDictionaryEntry, flags int32) *AVDictionaryEntry {
	keyPtr, keyFunc := StringCasting(key)
	defer keyFunc()
	return (*AVDictionaryEntry)(C.av_dict_get((*C.struct_AVDictionary)(m),
		(*C.char)(keyPtr), (*C.struct_AVDictionaryEntry)(prev), (C.int)(flags)))
}

// AvDictCount gets number of entries in dictionary.
func AvDictCount(m *AVDictionary) int32 {
	return (int32)(C.av_dict_count((*C.struct_AVDictionary)(m)))
}

// AvDictSet sets the given entry in *pm, overwriting an existing entry.
func AvDictSet(pm **AVDictionary, key, value string, flags int32) int32 {
	keyPtr, keyFunc := StringCasting(key)
	defer keyFunc()
	valuePtr, valueFunc := StringCasting(value)
	defer valueFunc()
	return (int32)(C.av_dict_set((**C.struct_AVDictionary)(unsafe.Pointer(pm)),
		(*C.char)(keyPtr), (*C.char)(valuePtr), (C.int)(flags)))
}

// AvDictSetInt sets the given entry in *pm, overwriting an existing entry.
func AvDictSetInt(pm **AVDictionary, key string, value int64, flags int32) int32 {
	keyPtr, keyFunc := StringCasting(key)
	defer keyFunc()
	return (int32)(C.av_dict_set_int((**C.struct_AVDictionary)(unsafe.Pointer(pm)),
		(*C.char)(keyPtr), (C.int64_t)(value), (C.int)(flags)))
}

// AvDictParseString parses the key/value pairs list and add the parsed entries to a dictionary.
func AvDictParseString(pm **AVDictionary, str, keyValSep, pairsSep string, flags int32) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	keyValSepPtr, keyValSepFunc := StringCasting(keyValSep)
	defer keyValSepFunc()
	pairsSepPtr, pairsSepFunc := StringCasting(pairsSep)
	defer pairsSepFunc()
	return (int32)(C.av_dict_parse_string((**C.struct_AVDictionary)(unsafe.Pointer(pm)),
		(*C.char)(strPtr), (*C.char)(keyValSepPtr), (*C.char)(pairsSepPtr), (C.int)(flags)))
}

// AvDictCopy copies entries from one AVDictionary struct into another.
func AvDictCopy(dst **AVDictionary, src *AVDictionary, flags int32) int32 {
	return (int32)(C.av_dict_copy((**C.struct_AVDictionary)(unsafe.Pointer(dst)),
		(*C.struct_AVDictionary)(src), (C.int)(flags)))
}

// AvDictFree frees all the memory allocated for an AVDictionary struct and all keys and values.
func AvDictFree(m **AVDictionary) {
	C.av_dict_free((**C.struct_AVDictionary)(unsafe.Pointer(m)))
}

// AvDictGetString get dictionary entries as a string.
func AvDictGetString(m *AVDictionary, buffer **int8, keyValSep, pairsSep string) int32 {
	keyValSepPtr, keyValSepFunc := StringCasting(keyValSep)
	defer keyValSepFunc()
	pairsSepPtr, pairsSepFunc := StringCasting(pairsSep)
	defer pairsSepFunc()
	return (int32)(C.av_dict_get_string((*C.struct_AVDictionary)(m),
		(**C.char)(unsafe.Pointer(buffer)), (C.char)(*keyValSepPtr), (C.char)(*pairsSepPtr)))
}
