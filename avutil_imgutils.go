// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/imgutils.h>
*/
import "C"
import (
	"unsafe"
)

// AvImageFillMaxPixsteps computes the max pixel step for each plane of an image with a
// format described by pixdesc.
func AvImageFillMaxPixsteps(maxPixsteps, maxPixstepComps []int32, pixdesc *AVPixFmtDescriptor) {
	if len(maxPixsteps) < 4 {
		panic("maxPixsteps len < 4")
	}
	if len(maxPixstepComps) < 4 {
		panic("maxPixstepComps len < 4")
	}
	C.av_image_fill_max_pixsteps((*C.int)(&maxPixsteps[0]), (*C.int)(&maxPixstepComps[0]),
		(*C.struct_AVPixFmtDescriptor)(pixdesc))
}

// AvImageGetLinesize computes the size of an image line with format pix_fmt and width
// width for the plane plane.
func AvImageGetLinesize(pixFmt AVPixelFormat, width, plane int32) int32 {
	return (int32)(C.av_image_get_linesize((C.enum_AVPixelFormat)(pixFmt),
		(C.int)(width), (C.int)(plane)))
}

// AvImageFillLinesizes fills plane linesizes for an image with pixel format pix_fmt and width width.
func AvImageFillLinesizes(linesizes []int32, pixFmt AVPixelFormat, width int32) int32 {
	if len(linesizes) < 4 {
		panic("linesizes len < 4")
	}
	return (int32)(C.av_image_fill_linesizes((*C.int)(&linesizes[0]),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width)))
}

// AvImageFillPlaneSizes fills plane sizes for an image with pixel format pix_fmt and height height.
func AvImageFillPlaneSizes(size []uintptr, pixFmt AVPixelFormat, height int32, linesizes []int) int32 {
	if len(size) < 4 {
		panic("size len < 4")
	}
	if len(linesizes) < 4 {
		panic("linesizes len < 4")
	}
	return (int32)(C.av_image_fill_plane_sizes((*C.size_t)(unsafe.Pointer(&size[0])),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(height), (*C.ptrdiff_t)(unsafe.Pointer(&linesizes[0]))))
}

// AvImageFillPointers fills plane data pointers for an image with pixel format pix_fmt and height height.
func AvImageFillPointers(data []*uint8, pixFmt AVPixelFormat,
	height int32, ptr *uint8, linesizes []int32) int32 {
	if len(data) < 4 {
		panic("data len < 4")
	}
	if len(linesizes) < 4 {
		panic("linesizes len < 4")
	}
	return (int32)(C.av_image_fill_pointers((**C.uint8_t)(unsafe.Pointer(&data[0])),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(height), (*C.uint8_t)(ptr),
		(*C.int)(&linesizes[0])))
}

// AvImageAlloc allocates an image with size w and h and pixel format pix_fmt, and
// fill pointers and linesizes accordingly.
func AvImageAlloc(pointers []*uint8, linesizes []int32, w, h int32,
	pixFmt AVPixelFormat, align int32) int32 {
	if len(pointers) < 4 {
		panic("pointers len < 4")
	}
	if len(linesizes) < 4 {
		panic("linesizes len < 4")
	}
	return (int32)(C.av_image_alloc((**C.uint8_t)(unsafe.Pointer(&pointers[0])),
		(*C.int)(&linesizes[0]),
		(C.int)(w), (C.int)(h), (C.enum_AVPixelFormat)(pixFmt), (C.int)(align)))
}

// AvImageCopyPlane copies image plane from src to dst.
func AvImageCopyPlane(dst *uint8, dstLinesize int32, src *uint8,
	srcLinesize int32, bytewidth, height int32) {
	C.av_image_copy_plane((*C.uint8_t)(dst), (C.int)(dstLinesize),
		(*C.uint8_t)(src), (C.int)(srcLinesize),
		(C.int)(bytewidth), (C.int)(height))
}

// AvImageCopyPlaneUcFrom copies image data located in uncacheable (e.g. GPU mapped) memory.
func AvImageCopyPlaneUcFrom(dst *uint8, dstLinesize int,
	src *uint8, srcLinesize int,
	bytewidth int, height int32) {
	C.av_image_copy_plane_uc_from((*C.uint8_t)(dst), (C.ptrdiff_t)(dstLinesize),
		(*C.uint8_t)(src), (C.ptrdiff_t)(srcLinesize),
		(C.ptrdiff_t)(bytewidth), (C.int)(height))
}

// AvImageCopy copies image in src_data to dst_data.
func AvImageCopy(dstData []*uint8, dstLinesizes []int32, srcData []*uint8, srcLinesizes []int32,
	pixFmt AVPixelFormat, width, height int32) {
	if len(dstData) < 4 {
		panic("dstData len < 4")
	}
	if len(dstLinesizes) < 4 {
		panic("dstLinesizes len < 4")
	}
	if len(srcData) < 4 {
		panic("srcData len < 4")
	}
	if len(srcLinesizes) < 4 {
		panic("srcLinesizes len < 4")
	}
	C.av_image_copy((**C.uint8_t)(unsafe.Pointer(&dstData[0])),
		(*C.int)(&dstLinesizes[0]),
		(**C.uint8_t)(unsafe.Pointer(&srcData[0])),
		(*C.int)(&srcLinesizes[0]),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height))
}

// AvImageCopyUcFrom copies image data located in uncacheable (e.g. GPU mapped) memory.
func AvImageCopyUcFrom(dstData []*uint8, dstLinesizes []int, srcData []*uint8, srcLinesizes []int,
	pixFmt AVPixelFormat, width, height int32) {
	if len(dstData) < 4 {
		panic("dstData len < 4")
	}
	if len(dstLinesizes) < 4 {
		panic("dstLinesizes len < 4")
	}
	if len(srcData) < 4 {
		panic("srcData len < 4")
	}
	if len(srcLinesizes) < 4 {
		panic("srcLinesizes len < 4")
	}
	C.av_image_copy_uc_from((**C.uint8_t)(unsafe.Pointer(&dstData[0])),
		(*C.ptrdiff_t)(unsafe.Pointer(&dstLinesizes[0])),
		(**C.uint8_t)(unsafe.Pointer(&srcData[0])),
		(*C.ptrdiff_t)(unsafe.Pointer(&srcLinesizes[0])),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height))
}

// AvImageFillArrays setups the data pointers and linesizes based on the specified image
// parameters and the provided array.
func AvImageFillArrays(dstData []*uint8, dstLinesize []int32, src *uint8,
	pixFmt AVPixelFormat, width, height, align int32) {
	if len(dstData) < 4 {
		panic("dstData len < 4")
	}
	if len(dstLinesize) < 4 {
		panic("dstLinesize len < 4")
	}
	C.av_image_fill_arrays((**C.uint8_t)(unsafe.Pointer(&dstData[0])),
		(*C.int)(&dstLinesize[0]),
		(*C.uint8_t)(src),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height), (C.int)(align))
}

// AvImageGetBufferSize Return the size in bytes of the amount of data required to store an
// image with the given parameters.
func AvImageGetBufferSize(pixFmt AVPixelFormat, width, height, align int32) int32 {
	return (int32)(C.av_image_get_buffer_size((C.enum_AVPixelFormat)(pixFmt),
		(C.int)(width), (C.int)(height), (C.int)(align)))
}

// AvImageCopyToBuffer copies image data from an image into a buffer.
func AvImageCopyToBuffer(dst *uint8, dstSize int32, srcData []*uint8, srcLinesize []int32,
	pixFmt AVPixelFormat, width, height, align int32) int32 {
	if len(srcData) < 4 {
		panic("srcData len < 4")
	}
	if len(srcLinesize) < 4 {
		panic("srcLinesize len < 4")
	}
	return (int32)(C.av_image_copy_to_buffer((*C.uint8_t)(dst), (C.int)(dstSize),
		(**C.uint8_t)(unsafe.Pointer(&srcData[0])), (*C.int)(&srcLinesize[0]),
		(C.enum_AVPixelFormat)(pixFmt), (C.int)(width), (C.int)(height), (C.int)(align)))
}

// AvImageCheckSize checks if the given dimension of an image is valid, meaning that all
// bytes of the image can be addressed with a signed int.
func AvImageCheckSize(w, h uint32, logOffset int32, logCtx CVoidPointer) int32 {
	return (int32)(C.av_image_check_size((C.uint)(w), (C.uint)(h), (C.int)(logOffset), VoidPointer(logCtx)))
}

// AvImageCheckSize2 checks if the given dimension of an image is valid, meaning that all
// bytes of a plane of an image with the specified pix_fmt can be addressed with a signed int.
func AvImageCheckSize2(w, h uint32, maxPixels int64, pixFmt AVPixelFormat,
	logOffset int32, logCtx CVoidPointer) int32 {
	return (int32)(C.av_image_check_size2((C.uint)(w), (C.uint)(h),
		(C.int64_t)(maxPixels), (C.enum_AVPixelFormat)(pixFmt),
		(C.int)(logOffset), VoidPointer(logCtx)))
}

// AvImageCheckSar checks if the given sample aspect ratio of an image is valid.
func AvImageCheckSar(w, h uint32, sar AVRational) int32 {
	return (int32)(C.av_image_check_sar((C.uint)(w), (C.uint)(h), (C.struct_AVRational)(sar)))
}

// AvImageFillBlack overwrites the image data with black.
func AvImageFillBlack(dstData []*uint8, dstLinesize []int,
	pixFmt AVPixelFormat, _range AVColorRange, width, height int32) int32 {
	if len(dstData) < 4 {
		panic("dstData len < 4")
	}
	if len(dstLinesize) < 4 {
		panic("dstLinesize len < 4")
	}
	return (int32)(C.av_image_fill_black((**C.uint8_t)(unsafe.Pointer(&dstData[0])),
		(*C.ptrdiff_t)(unsafe.Pointer(&dstLinesize[0])),
		(C.enum_AVPixelFormat)(pixFmt), (C.enum_AVColorRange)(_range),
		(C.int)(width), (C.int)(height)))
}
