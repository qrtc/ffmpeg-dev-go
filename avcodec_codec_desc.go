package ffmpeg

/*
#include <libavcodec/codec_desc.h>
*/
import "C"

// AvCodecDescriptor
type AvCodecDescriptor C.struct_AVCodecDescriptor

const (
	AV_CODEC_PROP_INTRA_ONLY = C.AV_CODEC_PROP_INTRA_ONLY
	AV_CODEC_PROP_LOSSY      = C.AV_CODEC_PROP_LOSSY
	AV_CODEC_PROP_LOSSLESS   = C.AV_CODEC_PROP_LOSSLESS
	AV_CODEC_PROP_REORDER    = C.AV_CODEC_PROP_REORDER
	AV_CODEC_PROP_BITMAP_SUB = C.AV_CODEC_PROP_BITMAP_SUB
	AV_CODEC_PROP_TEXT_SUB   = C.AV_CODEC_PROP_TEXT_SUB
)

// AvCodecDescriptorGet returns descriptor for given codec ID or NULL if no descriptor exists.
func AvCodecDescriptorGet(id AvCodecID) *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

// AvCodecDescriptorNext iterates over all codec descriptors known to libavcodec.
func AvCodecDescriptorNext(prev *AvCodecDescriptor) *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(prev)))
}

// AvCodecDescriptorGetByName returns codec descriptor with the given name or NULL
// if no such descriptor exists.
func AvCodecDescriptorGetByName(name string) *AvCodecDescriptor {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*AvCodecDescriptor)(C.avcodec_descriptor_get_by_name((*C.char)(namePtr)))
}
