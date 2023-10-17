package ffmpeg

/*
#include <libavutil/pixdesc.h>
*/
import "C"
import "unsafe"

type AvComponentDescriptor C.struct_AVComponentDescriptor

type AvPixFmtDescriptor C.struct_AVPixFmtDescriptor

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
func AvGetBitsPerPixel(pixdesc *AvPixFmtDescriptor) int32 {
	return (int32)(C.av_get_bits_per_pixel((*C.struct_AVPixFmtDescriptor)(pixdesc)))
}

// AvGetPaddedBitsPerPixel returns the number of bits per pixel for the pixel format
// described by pixdesc, including any padding or unused bits.
func AvGetPaddedBitsPerPixel(pixdesc *AvPixFmtDescriptor) int32 {
	return (int32)(C.av_get_padded_bits_per_pixel((*C.struct_AVPixFmtDescriptor)(pixdesc)))
}

// AvPixFmtDescGet returns a pixel format descriptor for provided pixel format or NULL if
// this pixel format is unknown.
func AvPixFmtDescGet(pixFmt AvPixelFormat) *AvPixFmtDescriptor {
	return (*AvPixFmtDescriptor)(C.av_pix_fmt_desc_get((C.enum_AVPixelFormat)(pixFmt)))
}

// AvPixFmtDescNext iterates over all pixel format descriptors known to libavutil.
func AvPixFmtDescNext(prev *AvPixFmtDescriptor) *AvPixFmtDescriptor {
	return (*AvPixFmtDescriptor)(C.av_pix_fmt_desc_next((*C.struct_AVPixFmtDescriptor)(prev)))
}

// AvPixFmtDescGetId returns an AvPixelFormat id described by desc, or AV_PIX_FMT_NONE if desc
// is not a valid pointer to a pixel format descriptor.
func AvPixFmtDescGetId(desc *AvPixFmtDescriptor) AvPixelFormat {
	return (AvPixelFormat)(C.av_pix_fmt_desc_get_id((*C.struct_AVPixFmtDescriptor)(desc)))
}

// AvPixFmtGetChromaSubSample accesses log2_chroma_w log2_chroma_h from the pixel format AvPixFmtDescriptor.
func AvPixFmtGetChromaSubSample(pixFmt AvPixelFormat, hShift, vShift *int32) int32 {
	return (int32)(C.av_pix_fmt_get_chroma_sub_sample((C.enum_AVPixelFormat)(pixFmt),
		(*C.int)(hShift), (*C.int)(vShift)))
}

// AvPixFmtCountPlanes returns number of planes in pix_fmt, a negative AvERROR if pix_fmt is not a
// valid pixel format.
func AvPixFmtCountPlanes(pixFmt AvPixelFormat) int32 {
	return (int32)(C.av_pix_fmt_count_planes((C.enum_AVPixelFormat)(pixFmt)))
}

// AvColorRangeName returns the name for provided color range or NULL if unknown.
func AvColorRangeName(_range AvColorRange) string {
	return C.GoString(C.av_color_range_name((C.enum_AVColorRange)(_range)))
}

// AvColorRangeFromName returns the AvColorRange value for name or an AvError if not found.
func AvColorRangeFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_color_range_from_name((*C.char)(namePtr)))
}

// AvColorPrimariesName returns the name for provided color primaries or NULL if unknown.
func AvColorPrimariesName(primaries AvColorPrimaries) string {
	return C.GoString(C.av_color_primaries_name((C.enum_AVColorRange)(primaries)))
}

// AvColorPrimariesFromName returns the AvColorPrimaries value for name or an AVError if not found.
func AvColorPrimariesFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_color_primaries_from_name((*C.char)(namePtr)))
}

// AvColorTransferName returns the name for provided color transfer or NULL if unknown.
func AvColorTransferName(transfer AvColorTransferCharacteristic) string {
	return C.GoString(C.av_color_transfer_name((C.enum_AVColorTransferCharacteristic)(transfer)))
}

// AvColorTransferFromName returns the AvColorTransferCharacteristic value for name or an AvError if not found.
func AvColorTransferFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_color_transfer_from_name((*C.char)(namePtr)))
}

// AvColorSpaceName returns the name for provided color space or NULL if unknown.
func AvColorSpaceName(space AvColorSpace) string {
	return C.GoString(C.av_color_space_name((C.enum_AVColorSpace)(space)))
}

// AvColorSpaceFromName returns the AvColorSpace value for name or an AvError if not found.
func AvColorSpaceFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_color_space_from_name((*C.char)(namePtr)))
}

// AvChromaLocationName returns the name for provided chroma location or NULL if unknown.
func AvChromaLocationName(location AvChromaLocation) string {
	return C.GoString(C.av_chroma_location_name((C.enum_AVChromaLocation)(location)))
}

// AvChromaLocationFromName returns the AvChromaLocation value for name or an AVError if not found.
func AvChromaLocationFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_chroma_location_from_name((*C.char)(namePtr)))
}

// AvGetPixFmt returns the pixel format corresponding to name.
func AvGetPixFmt(name string) AvPixelFormat {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (AvPixelFormat)(C.av_get_pix_fmt((*C.char)(namePtr)))
}

// AvGetPixFmtName returns the short name for a pixel format, NULL in case pix_fmt is unknown.
func AvGetPixFmtName(pixFmt AvPixelFormat) string {
	return C.GoString(C.av_get_pix_fmt_name((C.enum_AVPixelFormat)(pixFmt)))
}

// AvGetPixFmtString prints in buf the string corresponding to the pixel format with
// number pix_fmt, or a header if pix_fmt is negative.
func AvGetPixFmtString(buf *int8, bufSize int32, pixFmt AvPixelFormat) string {
	return C.GoString(C.av_get_pix_fmt_string((*C.char)(buf), (C.int)(bufSize),
		(C.enum_AVPixelFormat)(pixFmt)))
}

// AvReadImageLine2 reads a line from an image, and write the values of the
// pixel format component c to dst.
func AvReadImageLine2(dst unsafe.Pointer, data []*uint8, linesize []int,
	desc *AvPixFmtDescriptor, x, y, c, w, readPalComponent, dstElementSize int32) {
	if len(data) != 4 {
		panic("data need len = 4")
	}
	if len(linesize) != 4 {
		panic("linesize need len = 4")
	}
	C.av_read_image_line2(dst,
		(**C.uint8_t)(unsafe.Pointer(&data[0])),
		(*C.int)(unsafe.Pointer(&linesize[0])),
		(*C.struct_AVPixFmtDescriptor)(desc),
		(C.int)(x), (C.int)(y), (C.int)(c), (C.int)(w),
		(C.int)(readPalComponent), (C.int)(dstElementSize))
}

// AvReadImageLine reads a line from an image, and write the values of the
// pixel format component c to dst.
func AvReadImageLine(dst *uint16, data []*uint8, linesize []int,
	desc *AvPixFmtDescriptor, x, y, c, w, readPalComponent int32) {
	if len(data) != 4 {
		panic("data need len = 4")
	}
	if len(linesize) != 4 {
		panic("linesize need len = 4")
	}
	C.av_read_image_line((*C.uint16_t)(dst),
		(**C.uint8_t)(unsafe.Pointer(&data[0])),
		(*C.int)(unsafe.Pointer(&linesize[0])),
		(*C.struct_AVPixFmtDescriptor)(desc),
		(C.int)(x), (C.int)(y), (C.int)(c), (C.int)(w),
		(C.int)(readPalComponent))
}

// AvWriteImageLine2 writes the values from src to the pixel format component c of an image line.
func AvWriteImageLine2(src unsafe.Pointer, data []*uint8, linesize []int,
	desc *AvPixFmtDescriptor, x, y, c, w, srcElementSize int32) {
	if len(data) != 4 {
		panic("data need len = 4")
	}
	if len(linesize) != 4 {
		panic("linesize need len = 4")
	}
	C.av_write_image_line2(src,
		(**C.uint8_t)(unsafe.Pointer(&data[0])),
		(*C.int)(unsafe.Pointer(&linesize[0])),
		(*C.struct_AVPixFmtDescriptor)(desc),
		(C.int)(x), (C.int)(y), (C.int)(c), (C.int)(w),
		(C.int)(srcElementSize))
}

// AvWriteImageLine writes the values from src to the pixel format component c of an image line.
func AvWriteImageLine(src *uint16, data []*uint8, linesize []int,
	desc *AvPixFmtDescriptor, x, y, c, w int32) {
	if len(data) != 4 {
		panic("data need len = 4")
	}
	if len(linesize) != 4 {
		panic("linesize need len = 4")
	}
	C.av_write_image_line((*C.uint16_t)(src),
		(**C.uint8_t)(unsafe.Pointer(&data[0])),
		(*C.int)(unsafe.Pointer(&linesize[0])),
		(*C.struct_AVPixFmtDescriptor)(desc),
		(C.int)(x), (C.int)(y), (C.int)(c), (C.int)(w))
}

// AvPixFmtSwapEndianness
func AvPixFmtSwapEndianness(pixFmt AvPixelFormat) AvPixelFormat {
	return (AvPixelFormat)(C.av_pix_fmt_swap_endianness((C.enum_AVPixelFormat)(pixFmt)))
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
func AvGetPixFmtLoss(dstPixFmt, srcPixFmt AvPixelFormat, hasAlpha int32) int32 {
	return (int32)(C.av_get_pix_fmt_loss((C.enum_AVPixelFormat)(dstPixFmt),
		(C.enum_AVPixelFormat)(srcPixFmt), (C.int)(hasAlpha)))
}

// AvFindBestPixFmtOf2 compute what kind of losses will occur when converting from one specific
// pixel format to another.
func AvFindBestPixFmtOf2(dstPixFmt1, dstPixFmt2, srcPixFmt AvPixelFormat,
	hasAlpha int32, lossPtr *int32) AvPixelFormat {
	return (AvPixelFormat)(C.av_find_best_pix_fmt_of_2((C.enum_AVPixelFormat)(dstPixFmt1),
		(C.enum_AVPixelFormat)(dstPixFmt2),
		(C.enum_AVPixelFormat)(srcPixFmt),
		(C.int)(hasAlpha), (*C.int)(lossPtr)))
}
