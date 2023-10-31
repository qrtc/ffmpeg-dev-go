// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/pixdesc.h>
*/
import "C"
import "unsafe"

// AVComponentDescriptor
type AVComponentDescriptor C.struct_AVComponentDescriptor

// GetPlane gets `AVComponentDescriptor.plane` value.
func (cd *AVComponentDescriptor) GetPlane() int32 {
	return (int32)(cd.plane)
}

// SetPlane sets `AVComponentDescriptor.plane` value.
func (cd *AVComponentDescriptor) SetPlane(v int32) {
	cd.plane = (C.int)(v)
}

// GetPlaneAddr gets `AVComponentDescriptor.plane` address.
func (cd *AVComponentDescriptor) GetPlaneAddr() *int32 {
	return (*int32)(&cd.plane)
}

// GetStep gets `AVComponentDescriptor.step` value.
func (cd *AVComponentDescriptor) GetStep() int32 {
	return (int32)(cd.step)
}

// SetStep sets `AVComponentDescriptor.step` value.
func (cd *AVComponentDescriptor) SetStep(v int32) {
	cd.step = (C.int)(v)
}

// GetStepAddr gets `AVComponentDescriptor.step` address.
func (cd *AVComponentDescriptor) GetStepAddr() *int32 {
	return (*int32)(&cd.step)
}

// GetOffset gets `AVComponentDescriptor.offset` value.
func (cd *AVComponentDescriptor) GetOffset() int32 {
	return (int32)(cd.offset)
}

// SetOffset sets `AVComponentDescriptor.offset` value.
func (cd *AVComponentDescriptor) SetOffset(v int32) {
	cd.offset = (C.int)(v)
}

// GetOffsetAddr gets `AVComponentDescriptor.offset` address.
func (cd *AVComponentDescriptor) GetOffsetAddr() *int32 {
	return (*int32)(&cd.offset)
}

// GetShift gets `AVComponentDescriptor.shift` value.
func (cd *AVComponentDescriptor) GetShift() int32 {
	return (int32)(cd.shift)
}

// SetShift sets `AVComponentDescriptor.shift` value.
func (cd *AVComponentDescriptor) SetShift(v int32) {
	cd.shift = (C.int)(v)
}

// GetShiftAddr gets `AVComponentDescriptor.shift` address.
func (cd *AVComponentDescriptor) GetShiftAddr() *int32 {
	return (*int32)(&cd.shift)
}

// GetDepth gets `AVComponentDescriptor.depth` value.
func (cd *AVComponentDescriptor) GetDepth() int32 {
	return (int32)(cd.depth)
}

// SetDepth sets `AVComponentDescriptor.depth` value.
func (cd *AVComponentDescriptor) SetDepth(v int32) {
	cd.depth = (C.int)(v)
}

// GetDepthAddr gets `AVComponentDescriptor.depth` address.
func (cd *AVComponentDescriptor) GetDepthAddr() *int32 {
	return (*int32)(&cd.depth)
}

// Deprecated: Use step instead.
//
// GetStepMinus1 gets `AVComponentDescriptor.step_minus1` value.
func (cd *AVComponentDescriptor) GetStepMinus1() int32 {
	return (int32)(cd.step_minus1)
}

// Deprecated: Use step instead.
//
// SetStepMinus1 sets `AVComponentDescriptor.step_minus1` value.
func (cd *AVComponentDescriptor) SetStepMinus1(v int32) {
	cd.step_minus1 = (C.int)(v)
}

// Deprecated: Use step instead.
//
// GetStepMinus1Addr gets `AVComponentDescriptor.step_minus1` address.
func (cd *AVComponentDescriptor) GetStepMinus1Addr() *int32 {
	return (*int32)(&cd.step_minus1)
}

// Deprecated: Use depth instead.
//
// GetDepthMinus1 gets `AVComponentDescriptor.depth_minus1` value.
func (cd *AVComponentDescriptor) GetDepthMinus1() int32 {
	return (int32)(cd.depth_minus1)
}

// Deprecated: Use depth instead.
//
// SetDepthMinus1 sets `AVComponentDescriptor.depth_minus1` value.
func (cd *AVComponentDescriptor) SetDepthMinus1(v int32) {
	cd.depth_minus1 = (C.int)(v)
}

// Deprecated: Use depth instead.
//
// GetDepthMinus1Addr gets `AVComponentDescriptor.depth_minus1` address.
func (cd *AVComponentDescriptor) GetDepthMinus1Addr() *int32 {
	return (*int32)(&cd.depth_minus1)
}

// Deprecated: Use offset instead.
//
// GetOffsetPlus1 gets `AVComponentDescriptor.offset_plus1` value.
func (cd *AVComponentDescriptor) GetOffsetPlus1() int32 {
	return (int32)(cd.offset_plus1)
}

// Deprecated: Use offset instead.
//
// SetOffsetPlus1 sets `AVComponentDescriptor.offset_plus1` value.
func (cd *AVComponentDescriptor) SetOffsetPlus1(v int32) {
	cd.offset_plus1 = (C.int)(v)
}

// Deprecated: Use offset instead.
//
// GetOffsetPlus1Addr gets `AVComponentDescriptor.offset_plus1` address.
func (cd *AVComponentDescriptor) GetOffsetPlus1Addr() *int32 {
	return (*int32)(&cd.offset_plus1)
}

// AVPixFmtDescriptor
type AVPixFmtDescriptor C.struct_AVPixFmtDescriptor

// GetName gets `AVPixFmtDescriptor.name` value.
func (pfd *AVPixFmtDescriptor) GetName() string {
	return C.GoString(pfd.name)
}

// GetNbComponents gets `AVPixFmtDescriptor.nb_components` value.
func (pfd *AVPixFmtDescriptor) GetNbComponents() uint8 {
	return (uint8)(pfd.nb_components)
}

// SetNbComponents sets `AVPixFmtDescriptor.nb_components` value.
func (pfd *AVPixFmtDescriptor) SetNbComponents(v uint8) {
	pfd.nb_components = (C.uint8_t)(v)
}

// GetNbComponentsAddr gets `AVPixFmtDescriptor.nb_components` address.
func (pfd *AVPixFmtDescriptor) GetNbComponentsAddr() *uint8 {
	return (*uint8)(&pfd.nb_components)
}

// GetLog2ChromaW gets `AVPixFmtDescriptor.log2_chroma_w` value.
func (pfd *AVPixFmtDescriptor) GetLog2ChromaW() uint8 {
	return (uint8)(pfd.log2_chroma_w)
}

// SetLog2ChromaW sets `AVPixFmtDescriptor.log2_chroma_w` value.
func (pfd *AVPixFmtDescriptor) SetLog2ChromaW(v uint8) {
	pfd.log2_chroma_w = (C.uint8_t)(v)
}

// GetLog2ChromaWAddr gets `AVPixFmtDescriptor.log2_chroma_w` address.
func (pfd *AVPixFmtDescriptor) GetLog2ChromaWAddr() *uint8 {
	return (*uint8)(&pfd.log2_chroma_w)
}

// GetLog2ChromaH gets `AVPixFmtDescriptor.log2_chroma_h` value.
func (pfd *AVPixFmtDescriptor) GetLog2ChromaH() uint8 {
	return (uint8)(pfd.log2_chroma_h)
}

// SetLog2ChromaH sets `AVPixFmtDescriptor.log2_chroma_h` value.
func (pfd *AVPixFmtDescriptor) SetLog2ChromaH(v uint8) {
	pfd.log2_chroma_h = (C.uint8_t)(v)
}

// GetLog2ChromaHAddr gets `AVPixFmtDescriptor.log2_chroma_h` address.
func (pfd *AVPixFmtDescriptor) GetLog2ChromaHAddr() *uint8 {
	return (*uint8)(&pfd.log2_chroma_h)
}

// GetFlags gets `AVPixFmtDescriptor.flags` value.
func (pfd *AVPixFmtDescriptor) GetFlags() uint64 {
	return (uint64)(pfd.flags)
}

// SetFlags sets `AVPixFmtDescriptor.flags` value.
func (pfd *AVPixFmtDescriptor) SetFlags(v uint64) {
	pfd.flags = (C.uint64_t)(v)
}

// GetFlagsAddr gets `AVPixFmtDescriptor.flags` address.
func (pfd *AVPixFmtDescriptor) GetFlagsAddr() *uint64 {
	return (*uint64)(&pfd.flags)
}

// GetComp gets `AVPixFmtDescriptor.comp` value.
func (pfd *AVPixFmtDescriptor) GetComp() []AVComponentDescriptor {
	return unsafe.Slice((*AVComponentDescriptor)(&pfd.comp[0]), 4)
}

// SetComp sets `AVPixFmtDescriptor.comp` value.
func (pfd *AVPixFmtDescriptor) SetComp(v []AVComponentDescriptor) {
	for i := 0; i < FFMIN(len(v), 4); i++ {
		pfd.comp[i] = (C.struct_AVComponentDescriptor)(v[i])
	}
}

// GetCompAddr gets `AVPixFmtDescriptor.comp` address.
func (pfd *AVPixFmtDescriptor) GetCompAddr() **AVComponentDescriptor {
	return (**AVComponentDescriptor)(unsafe.Pointer(&pfd.comp))
}

// GetAlias gets `AVPixFmtDescriptor.alias` value.
func (pfd *AVPixFmtDescriptor) GetAlias() string {
	return C.GoString(pfd.alias)
}

const (
	AV_PIX_FMT_FLAG_BE        = C.AV_PIX_FMT_FLAG_BE
	AV_PIX_FMT_FLAG_PAL       = C.AV_PIX_FMT_FLAG_PAL
	AV_PIX_FMT_FLAG_BITSTREAM = C.AV_PIX_FMT_FLAG_BITSTREAM
	AV_PIX_FMT_FLAG_HWACCEL   = C.AV_PIX_FMT_FLAG_HWACCEL
	AV_PIX_FMT_FLAG_PLANAR    = C.AV_PIX_FMT_FLAG_PLANAR
	AV_PIX_FMT_FLAG_RGB       = C.AV_PIX_FMT_FLAG_RGB
	AV_PIX_FMT_FLAG_PSEUDOPAL = C.AV_PIX_FMT_FLAG_PSEUDOPAL
	AV_PIX_FMT_FLAG_ALPHA     = C.AV_PIX_FMT_FLAG_ALPHA
	AV_PIX_FMT_FLAG_BAYER     = C.AV_PIX_FMT_FLAG_BAYER
	AV_PIX_FMT_FLAG_FLOAT     = C.AV_PIX_FMT_FLAG_FLOAT
)

// AvGetBitsPerPixel returns the number of bits per pixel used by the pixel format
// described by pixdesc.
func AvGetBitsPerPixel(pixdesc *AVPixFmtDescriptor) int32 {
	return (int32)(C.av_get_bits_per_pixel((*C.struct_AVPixFmtDescriptor)(pixdesc)))
}

// AvGetPaddedBitsPerPixel returns the number of bits per pixel for the pixel format
// described by pixdesc, including any padding or unused bits.
func AvGetPaddedBitsPerPixel(pixdesc *AVPixFmtDescriptor) int32 {
	return (int32)(C.av_get_padded_bits_per_pixel((*C.struct_AVPixFmtDescriptor)(pixdesc)))
}

// AvPixFmtDescGet returns a pixel format descriptor for provided pixel format or NULL if
// this pixel format is unknown.
func AvPixFmtDescGet(pixFmt AVPixelFormat) *AVPixFmtDescriptor {
	return (*AVPixFmtDescriptor)(C.av_pix_fmt_desc_get((C.enum_AVPixelFormat)(pixFmt)))
}

// AvPixFmtDescNext iterates over all pixel format descriptors known to libavutil.
func AvPixFmtDescNext(prev *AVPixFmtDescriptor) *AVPixFmtDescriptor {
	return (*AVPixFmtDescriptor)(C.av_pix_fmt_desc_next((*C.struct_AVPixFmtDescriptor)(prev)))
}

// AvPixFmtDescGetId returns an AVPixelFormat id described by desc, or AV_PIX_FMT_NONE if desc
// is not a valid pointer to a pixel format descriptor.
func AvPixFmtDescGetId(desc *AVPixFmtDescriptor) AVPixelFormat {
	return (AVPixelFormat)(C.av_pix_fmt_desc_get_id((*C.struct_AVPixFmtDescriptor)(desc)))
}

// AvPixFmtGetChromaSubSample accesses log2_chroma_w log2_chroma_h from the pixel format AVPixFmtDescriptor.
func AvPixFmtGetChromaSubSample(pixFmt AVPixelFormat, hShift, vShift *int32) int32 {
	return (int32)(C.av_pix_fmt_get_chroma_sub_sample((C.enum_AVPixelFormat)(pixFmt),
		(*C.int)(hShift), (*C.int)(vShift)))
}

// AvPixFmtCountPlanes returns number of planes in pix_fmt, a negative AvERROR if pix_fmt is not a
// valid pixel format.
func AvPixFmtCountPlanes(pixFmt AVPixelFormat) int32 {
	return (int32)(C.av_pix_fmt_count_planes((C.enum_AVPixelFormat)(pixFmt)))
}

// AvColorRangeName returns the name for provided color range or NULL if unknown.
func AvColorRangeName(_range AVColorRange) string {
	return C.GoString(C.av_color_range_name((C.enum_AVColorRange)(_range)))
}

// AvColorRangeFromName returns the AVColorRange value for name or an AvError if not found.
func AvColorRangeFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_color_range_from_name((*C.char)(namePtr)))
}

// AvColorPrimariesName returns the name for provided color primaries or NULL if unknown.
func AvColorPrimariesName(primaries AVColorPrimaries) string {
	return C.GoString(C.av_color_primaries_name((C.enum_AVColorRange)(primaries)))
}

// AvColorPrimariesFromName returns the AVColorPrimaries value for name or an AVError if not found.
func AvColorPrimariesFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_color_primaries_from_name((*C.char)(namePtr)))
}

// AvColorTransferName returns the name for provided color transfer or NULL if unknown.
func AvColorTransferName(transfer AVColorTransferCharacteristic) string {
	return C.GoString(C.av_color_transfer_name((C.enum_AVColorTransferCharacteristic)(transfer)))
}

// AvColorTransferFromName returns the AVColorTransferCharacteristic value for name or an AvError if not found.
func AvColorTransferFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_color_transfer_from_name((*C.char)(namePtr)))
}

// AvColorSpaceName returns the name for provided color space or NULL if unknown.
func AvColorSpaceName(space AVColorSpace) string {
	return C.GoString(C.av_color_space_name((C.enum_AVColorSpace)(space)))
}

// AvColorSpaceFromName returns the AVColorSpace value for name or an AvError if not found.
func AvColorSpaceFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_color_space_from_name((*C.char)(namePtr)))
}

// AvChromaLocationName returns the name for provided chroma location or NULL if unknown.
func AvChromaLocationName(location AVChromaLocation) string {
	return C.GoString(C.av_chroma_location_name((C.enum_AVChromaLocation)(location)))
}

// AvChromaLocationFromName returns the AVChromaLocation value for name or an AVError if not found.
func AvChromaLocationFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_chroma_location_from_name((*C.char)(namePtr)))
}

// AvGetPixFmt returns the pixel format corresponding to name.
func AvGetPixFmt(name string) AVPixelFormat {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (AVPixelFormat)(C.av_get_pix_fmt((*C.char)(namePtr)))
}

// AvGetPixFmtName returns the short name for a pixel format, NULL in case pix_fmt is unknown.
func AvGetPixFmtName(pixFmt AVPixelFormat) string {
	return C.GoString(C.av_get_pix_fmt_name((C.enum_AVPixelFormat)(pixFmt)))
}

// AvGetPixFmtString prints in buf the string corresponding to the pixel format with
// number pix_fmt, or a header if pix_fmt is negative.
func AvGetPixFmtString(buf *int8, bufSize int32, pixFmt AVPixelFormat) string {
	return C.GoString(C.av_get_pix_fmt_string((*C.char)(buf), (C.int)(bufSize),
		(C.enum_AVPixelFormat)(pixFmt)))
}

// AvReadImageLine reads a line from an image, and write the values of the
// pixel format component c to dst.
func AvReadImageLine(dst *uint16, data []*uint8, linesize []int,
	desc *AVPixFmtDescriptor, x, y, c, w, readPalComponent int32) {
	if len(data) < 4 {
		panic("data len < 4")
	}
	if len(linesize) < 4 {
		panic("linesize len < 4")
	}
	C.av_read_image_line((*C.uint16_t)(dst),
		(**C.uint8_t)(unsafe.Pointer(&data[0])),
		(*C.int)(unsafe.Pointer(&linesize[0])),
		(*C.struct_AVPixFmtDescriptor)(desc),
		(C.int)(x), (C.int)(y), (C.int)(c), (C.int)(w),
		(C.int)(readPalComponent))
}

// AvWriteImageLine writes the values from src to the pixel format component c of an image line.
func AvWriteImageLine(src *uint16, data []*uint8, linesize []int,
	desc *AVPixFmtDescriptor, x, y, c, w int32) {
	if len(data) < 4 {
		panic("data len < 4")
	}
	if len(linesize) < 4 {
		panic("linesize len < 4")
	}
	C.av_write_image_line((*C.uint16_t)(src),
		(**C.uint8_t)(unsafe.Pointer(&data[0])),
		(*C.int)(unsafe.Pointer(&linesize[0])),
		(*C.struct_AVPixFmtDescriptor)(desc),
		(C.int)(x), (C.int)(y), (C.int)(c), (C.int)(w))
}

// AvPixFmtSwapEndianness
func AvPixFmtSwapEndianness(pixFmt AVPixelFormat) AVPixelFormat {
	return (AVPixelFormat)(C.av_pix_fmt_swap_endianness((C.enum_AVPixelFormat)(pixFmt)))
}

const (
	FF_LOSS_RESOLUTION = C.FF_LOSS_RESOLUTION
	FF_LOSS_DEPTH      = C.FF_LOSS_DEPTH
	FF_LOSS_COLORSPACE = C.FF_LOSS_COLORSPACE
	FF_LOSS_ALPHA      = C.FF_LOSS_ALPHA
	FF_LOSS_COLORQUANT = C.FF_LOSS_COLORQUANT
	FF_LOSS_CHROMA     = C.FF_LOSS_CHROMA
)

// AvGetPixFmtLoss computes what kind of losses will occur when converting from one specific
// pixel format to another.
func AvGetPixFmtLoss(dstPixFmt, srcPixFmt AVPixelFormat, hasAlpha int32) int32 {
	return (int32)(C.av_get_pix_fmt_loss((C.enum_AVPixelFormat)(dstPixFmt),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha)))
}

// AvFindBestPixFmtOf2 compute what kind of losses will occur when converting from one specific
// pixel format to another.
func AvFindBestPixFmtOf2(dstPixFmt1, dstPixFmt2, srcPixFmt AVPixelFormat,
	hasAlpha int32, lossPtr *int32) AVPixelFormat {
	return (AVPixelFormat)(C.av_find_best_pix_fmt_of_2((C.enum_AVPixelFormat)(dstPixFmt1),
		(C.enum_AVPixelFormat)(dstPixFmt2),
		(C.enum_AVPixelFormat)(srcPixFmt),
		(C.int)(hasAlpha), (*C.int)(lossPtr)))
}
