// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/pixfmt.h>
*/
import "C"

// Pixel format.
type AVPixelFormat = C.enum_AVPixelFormat

const (
	AV_PIX_FMT_NONE      = AVPixelFormat(C.AV_PIX_FMT_NONE)
	AV_PIX_FMT_YUV420P   = AVPixelFormat(C.AV_PIX_FMT_YUV420P)
	AV_PIX_FMT_YUYV422   = AVPixelFormat(C.AV_PIX_FMT_YUYV422)
	AV_PIX_FMT_RGB24     = AVPixelFormat(C.AV_PIX_FMT_RGB24)
	AV_PIX_FMT_BGR24     = AVPixelFormat(C.AV_PIX_FMT_BGR24)
	AV_PIX_FMT_YUV422P   = AVPixelFormat(C.AV_PIX_FMT_YUV422P)
	AV_PIX_FMT_YUV444P   = AVPixelFormat(C.AV_PIX_FMT_YUV444P)
	AV_PIX_FMT_YUV410P   = AVPixelFormat(C.AV_PIX_FMT_YUV410P)
	AV_PIX_FMT_YUV411P   = AVPixelFormat(C.AV_PIX_FMT_YUV411P)
	AV_PIX_FMT_GRAY8     = AVPixelFormat(C.AV_PIX_FMT_GRAY8)
	AV_PIX_FMT_MONOWHITE = AVPixelFormat(C.AV_PIX_FMT_MONOWHITE)
	AV_PIX_FMT_MONOBLACK = AVPixelFormat(C.AV_PIX_FMT_MONOBLACK)
	AV_PIX_FMT_PAL8      = AVPixelFormat(C.AV_PIX_FMT_PAL8)
	AV_PIX_FMT_YUVJ420P  = AVPixelFormat(C.AV_PIX_FMT_YUVJ420P)
	AV_PIX_FMT_YUVJ422P  = AVPixelFormat(C.AV_PIX_FMT_YUVJ422P)
	AV_PIX_FMT_YUVJ444P  = AVPixelFormat(C.AV_PIX_FMT_YUVJ444P)
	AV_PIX_FMT_UYVY422   = AVPixelFormat(C.AV_PIX_FMT_UYVY422)
	AV_PIX_FMT_UYYVYY411 = AVPixelFormat(C.AV_PIX_FMT_UYYVYY411)
	AV_PIX_FMT_BGR8      = AVPixelFormat(C.AV_PIX_FMT_BGR8)
	AV_PIX_FMT_BGR4      = AVPixelFormat(C.AV_PIX_FMT_BGR4)
	AV_PIX_FMT_BGR4_BYTE = AVPixelFormat(C.AV_PIX_FMT_BGR4_BYTE)
	AV_PIX_FMT_RGB8      = AVPixelFormat(C.AV_PIX_FMT_RGB8)
	AV_PIX_FMT_RGB4      = AVPixelFormat(C.AV_PIX_FMT_RGB4)
	AV_PIX_FMT_RGB4_BYTE = AVPixelFormat(C.AV_PIX_FMT_RGB4_BYTE)
	AV_PIX_FMT_NV12      = AVPixelFormat(C.AV_PIX_FMT_NV12)
	AV_PIX_FMT_NV21      = AVPixelFormat(C.AV_PIX_FMT_NV21)

	AV_PIX_FMT_ARGB = AVPixelFormat(C.AV_PIX_FMT_ARGB)
	AV_PIX_FMT_RGBA = AVPixelFormat(C.AV_PIX_FMT_RGBA)
	AV_PIX_FMT_ABGR = AVPixelFormat(C.AV_PIX_FMT_ABGR)
	AV_PIX_FMT_BGRA = AVPixelFormat(C.AV_PIX_FMT_BGRA)

	AV_PIX_FMT_GRAY16BE = AVPixelFormat(C.AV_PIX_FMT_GRAY16BE)
	AV_PIX_FMT_GRAY16LE = AVPixelFormat(C.AV_PIX_FMT_GRAY16LE)
	AV_PIX_FMT_YUV440P  = AVPixelFormat(C.AV_PIX_FMT_YUV440P)
	AV_PIX_FMT_YUVJ440P = AVPixelFormat(C.AV_PIX_FMT_YUVJ440P)
	AV_PIX_FMT_YUVA420P = AVPixelFormat(C.AV_PIX_FMT_YUVA420P)
	AV_PIX_FMT_RGB48BE  = AVPixelFormat(C.AV_PIX_FMT_RGB48BE)
	AV_PIX_FMT_RGB48LE  = AVPixelFormat(C.AV_PIX_FMT_RGB48LE)

	AV_PIX_FMT_RGB565BE = AVPixelFormat(C.AV_PIX_FMT_RGB565BE)
	AV_PIX_FMT_RGB565LE = AVPixelFormat(C.AV_PIX_FMT_RGB565LE)
	AV_PIX_FMT_RGB555BE = AVPixelFormat(C.AV_PIX_FMT_RGB555BE)
	AV_PIX_FMT_RGB555LE = AVPixelFormat(C.AV_PIX_FMT_RGB555LE)

	AV_PIX_FMT_BGR565BE = AVPixelFormat(C.AV_PIX_FMT_BGR565BE)
	AV_PIX_FMT_BGR565LE = AVPixelFormat(C.AV_PIX_FMT_BGR565LE)
	AV_PIX_FMT_BGR555BE = AVPixelFormat(C.AV_PIX_FMT_BGR555BE)
	AV_PIX_FMT_BGR555LE = AVPixelFormat(C.AV_PIX_FMT_BGR555LE)

	AV_PIX_FMT_VAAPI = AVPixelFormat(C.AV_PIX_FMT_VAAPI)

	AV_PIX_FMT_YUV420P16LE = AVPixelFormat(C.AV_PIX_FMT_YUV420P16LE)
	AV_PIX_FMT_YUV420P16BE = AVPixelFormat(C.AV_PIX_FMT_YUV420P16BE)
	AV_PIX_FMT_YUV422P16LE = AVPixelFormat(C.AV_PIX_FMT_YUV422P16LE)
	AV_PIX_FMT_YUV422P16BE = AVPixelFormat(C.AV_PIX_FMT_YUV422P16BE)
	AV_PIX_FMT_YUV444P16LE = AVPixelFormat(C.AV_PIX_FMT_YUV444P16LE)
	AV_PIX_FMT_YUV444P16BE = AVPixelFormat(C.AV_PIX_FMT_YUV444P16BE)
	AV_PIX_FMT_DXVA2_VLD   = AVPixelFormat(C.AV_PIX_FMT_DXVA2_VLD)

	AV_PIX_FMT_RGB444LE = AVPixelFormat(C.AV_PIX_FMT_RGB444LE)
	AV_PIX_FMT_RGB444BE = AVPixelFormat(C.AV_PIX_FMT_RGB444BE)
	AV_PIX_FMT_BGR444LE = AVPixelFormat(C.AV_PIX_FMT_BGR444LE)
	AV_PIX_FMT_BGR444BE = AVPixelFormat(C.AV_PIX_FMT_BGR444BE)
	AV_PIX_FMT_YA8      = AVPixelFormat(C.AV_PIX_FMT_YA8)

	AV_PIX_FMT_Y400A  = AVPixelFormat(C.AV_PIX_FMT_Y400A)
	AV_PIX_FMT_GRAY8A = AVPixelFormat(C.AV_PIX_FMT_GRAY8A)

	AV_PIX_FMT_BGR48BE = AVPixelFormat(C.AV_PIX_FMT_BGR48BE)
	AV_PIX_FMT_BGR48LE = AVPixelFormat(C.AV_PIX_FMT_BGR48LE)

	AV_PIX_FMT_YUV420P9BE   = AVPixelFormat(C.AV_PIX_FMT_YUV420P9BE)
	AV_PIX_FMT_YUV420P9LE   = AVPixelFormat(C.AV_PIX_FMT_YUV420P9LE)
	AV_PIX_FMT_YUV420P10BE  = AVPixelFormat(C.AV_PIX_FMT_YUV420P10BE)
	AV_PIX_FMT_YUV420P10LE  = AVPixelFormat(C.AV_PIX_FMT_YUV420P10LE)
	AV_PIX_FMT_YUV422P10BE  = AVPixelFormat(C.AV_PIX_FMT_YUV422P10BE)
	AV_PIX_FMT_YUV422P10LE  = AVPixelFormat(C.AV_PIX_FMT_YUV422P10LE)
	AV_PIX_FMT_YUV444P9BE   = AVPixelFormat(C.AV_PIX_FMT_YUV444P9BE)
	AV_PIX_FMT_YUV444P9LE   = AVPixelFormat(C.AV_PIX_FMT_YUV444P9LE)
	AV_PIX_FMT_YUV444P10BE  = AVPixelFormat(C.AV_PIX_FMT_YUV444P10BE)
	AV_PIX_FMT_YUV444P10LE  = AVPixelFormat(C.AV_PIX_FMT_YUV444P10LE)
	AV_PIX_FMT_YUV422P9BE   = AVPixelFormat(C.AV_PIX_FMT_YUV422P9BE)
	AV_PIX_FMT_YUV422P9LE   = AVPixelFormat(C.AV_PIX_FMT_YUV422P9LE)
	AV_PIX_FMT_GBRP         = AVPixelFormat(C.AV_PIX_FMT_GBRP)
	AV_PIX_FMT_GBR24P       = AVPixelFormat(C.AV_PIX_FMT_GBR24P)
	AV_PIX_FMT_GBRP9BE      = AVPixelFormat(C.AV_PIX_FMT_GBRP9BE)
	AV_PIX_FMT_GBRP9LE      = AVPixelFormat(C.AV_PIX_FMT_GBRP9LE)
	AV_PIX_FMT_GBRP10BE     = AVPixelFormat(C.AV_PIX_FMT_GBRP10BE)
	AV_PIX_FMT_GBRP10LE     = AVPixelFormat(C.AV_PIX_FMT_GBRP10LE)
	AV_PIX_FMT_GBRP16BE     = AVPixelFormat(C.AV_PIX_FMT_GBRP16BE)
	AV_PIX_FMT_GBRP16LE     = AVPixelFormat(C.AV_PIX_FMT_GBRP16LE)
	AV_PIX_FMT_YUVA422P     = AVPixelFormat(C.AV_PIX_FMT_YUVA422P)
	AV_PIX_FMT_YUVA444P     = AVPixelFormat(C.AV_PIX_FMT_YUVA444P)
	AV_PIX_FMT_YUVA420P9BE  = AVPixelFormat(C.AV_PIX_FMT_YUVA420P9BE)
	AV_PIX_FMT_YUVA420P9LE  = AVPixelFormat(C.AV_PIX_FMT_YUVA420P9LE)
	AV_PIX_FMT_YUVA422P9BE  = AVPixelFormat(C.AV_PIX_FMT_YUVA422P9BE)
	AV_PIX_FMT_YUVA422P9LE  = AVPixelFormat(C.AV_PIX_FMT_YUVA422P9LE)
	AV_PIX_FMT_YUVA444P9BE  = AVPixelFormat(C.AV_PIX_FMT_YUVA444P9BE)
	AV_PIX_FMT_YUVA444P9LE  = AVPixelFormat(C.AV_PIX_FMT_YUVA444P9LE)
	AV_PIX_FMT_YUVA420P10BE = AVPixelFormat(C.AV_PIX_FMT_YUVA420P10BE)
	AV_PIX_FMT_YUVA420P10LE = AVPixelFormat(C.AV_PIX_FMT_YUVA420P10LE)
	AV_PIX_FMT_YUVA422P10BE = AVPixelFormat(C.AV_PIX_FMT_YUVA422P10BE)
	AV_PIX_FMT_YUVA422P10LE = AVPixelFormat(C.AV_PIX_FMT_YUVA422P10LE)
	AV_PIX_FMT_YUVA444P10BE = AVPixelFormat(C.AV_PIX_FMT_YUVA444P10BE)
	AV_PIX_FMT_YUVA444P10LE = AVPixelFormat(C.AV_PIX_FMT_YUVA444P10LE)
	AV_PIX_FMT_YUVA420P16BE = AVPixelFormat(C.AV_PIX_FMT_YUVA420P16BE)
	AV_PIX_FMT_YUVA420P16LE = AVPixelFormat(C.AV_PIX_FMT_YUVA420P16LE)
	AV_PIX_FMT_YUVA422P16BE = AVPixelFormat(C.AV_PIX_FMT_YUVA422P16BE)
	AV_PIX_FMT_YUVA422P16LE = AVPixelFormat(C.AV_PIX_FMT_YUVA422P16LE)
	AV_PIX_FMT_YUVA444P16BE = AVPixelFormat(C.AV_PIX_FMT_YUVA444P16BE)
	AV_PIX_FMT_YUVA444P16LE = AVPixelFormat(C.AV_PIX_FMT_YUVA444P16LE)

	AV_PIX_FMT_VDPAU = AVPixelFormat(C.AV_PIX_FMT_VDPAU)

	AV_PIX_FMT_XYZ12LE  = AVPixelFormat(C.AV_PIX_FMT_XYZ12LE)
	AV_PIX_FMT_XYZ12BE  = AVPixelFormat(C.AV_PIX_FMT_XYZ12BE)
	AV_PIX_FMT_NV16     = AVPixelFormat(C.AV_PIX_FMT_NV16)
	AV_PIX_FMT_NV20LE   = AVPixelFormat(C.AV_PIX_FMT_NV20LE)
	AV_PIX_FMT_NV20BE   = AVPixelFormat(C.AV_PIX_FMT_NV20BE)
	AV_PIX_FMT_RGBA64BE = AVPixelFormat(C.AV_PIX_FMT_RGBA64BE)
	AV_PIX_FMT_RGBA64LE = AVPixelFormat(C.AV_PIX_FMT_RGBA64LE)
	AV_PIX_FMT_BGRA64BE = AVPixelFormat(C.AV_PIX_FMT_BGRA64BE)
	AV_PIX_FMT_BGRA64LE = AVPixelFormat(C.AV_PIX_FMT_BGRA64LE)

	AV_PIX_FMT_YVYU422 = AVPixelFormat(C.AV_PIX_FMT_YVYU422)

	AV_PIX_FMT_YA16BE = AVPixelFormat(C.AV_PIX_FMT_YA16BE)
	AV_PIX_FMT_YA16LE = AVPixelFormat(C.AV_PIX_FMT_YA16LE)

	AV_PIX_FMT_GBRAP     = AVPixelFormat(C.AV_PIX_FMT_GBRAP)
	AV_PIX_FMT_GBRAP16BE = AVPixelFormat(C.AV_PIX_FMT_GBRAP16BE)
	AV_PIX_FMT_GBRAP16LE = AVPixelFormat(C.AV_PIX_FMT_GBRAP16LE)

	AV_PIX_FMT_QSV = AVPixelFormat(C.AV_PIX_FMT_QSV)

	AV_PIX_FMT_MMAL = AVPixelFormat(C.AV_PIX_FMT_MMAL)

	AV_PIX_FMT_D3D11VA_VLD = AVPixelFormat(C.AV_PIX_FMT_D3D11VA_VLD)

	AV_PIX_FMT_CUDA = AVPixelFormat(C.AV_PIX_FMT_CUDA)

	AV_PIX_FMT_0RGB = AVPixelFormat(C.AV_PIX_FMT_0RGB)
	AV_PIX_FMT_RGB0 = AVPixelFormat(C.AV_PIX_FMT_RGB0)
	AV_PIX_FMT_0BGR = AVPixelFormat(C.AV_PIX_FMT_0BGR)
	AV_PIX_FMT_BGR0 = AVPixelFormat(C.AV_PIX_FMT_BGR0)

	AV_PIX_FMT_YUV420P12BE = AVPixelFormat(C.AV_PIX_FMT_YUV420P12BE)
	AV_PIX_FMT_YUV420P12LE = AVPixelFormat(C.AV_PIX_FMT_YUV420P12LE)
	AV_PIX_FMT_YUV420P14BE = AVPixelFormat(C.AV_PIX_FMT_YUV420P14BE)
	AV_PIX_FMT_YUV420P14LE = AVPixelFormat(C.AV_PIX_FMT_YUV420P14LE)
	AV_PIX_FMT_YUV422P12BE = AVPixelFormat(C.AV_PIX_FMT_YUV422P12BE)
	AV_PIX_FMT_YUV422P12LE = AVPixelFormat(C.AV_PIX_FMT_YUV422P12LE)
	AV_PIX_FMT_YUV422P14BE = AVPixelFormat(C.AV_PIX_FMT_YUV422P14BE)
	AV_PIX_FMT_YUV422P14LE = AVPixelFormat(C.AV_PIX_FMT_YUV422P14LE)
	AV_PIX_FMT_YUV444P12BE = AVPixelFormat(C.AV_PIX_FMT_YUV444P12BE)
	AV_PIX_FMT_YUV444P12LE = AVPixelFormat(C.AV_PIX_FMT_YUV444P12LE)
	AV_PIX_FMT_YUV444P14BE = AVPixelFormat(C.AV_PIX_FMT_YUV444P14BE)
	AV_PIX_FMT_YUV444P14LE = AVPixelFormat(C.AV_PIX_FMT_YUV444P14LE)
	AV_PIX_FMT_GBRP12BE    = AVPixelFormat(C.AV_PIX_FMT_GBRP12BE)
	AV_PIX_FMT_GBRP12LE    = AVPixelFormat(C.AV_PIX_FMT_GBRP12LE)
	AV_PIX_FMT_GBRP14BE    = AVPixelFormat(C.AV_PIX_FMT_GBRP14BE)
	AV_PIX_FMT_GBRP14LE    = AVPixelFormat(C.AV_PIX_FMT_GBRP14LE)
	AV_PIX_FMT_YUVJ411P    = AVPixelFormat(C.AV_PIX_FMT_YUVJ411P)

	AV_PIX_FMT_BAYER_BGGR8    = AVPixelFormat(C.AV_PIX_FMT_BAYER_BGGR8)
	AV_PIX_FMT_BAYER_RGGB8    = AVPixelFormat(C.AV_PIX_FMT_BAYER_RGGB8)
	AV_PIX_FMT_BAYER_GBRG8    = AVPixelFormat(C.AV_PIX_FMT_BAYER_GBRG8)
	AV_PIX_FMT_BAYER_GRBG8    = AVPixelFormat(C.AV_PIX_FMT_BAYER_GRBG8)
	AV_PIX_FMT_BAYER_BGGR16LE = AVPixelFormat(C.AV_PIX_FMT_BAYER_BGGR16LE)
	AV_PIX_FMT_BAYER_BGGR16BE = AVPixelFormat(C.AV_PIX_FMT_BAYER_BGGR16BE)
	AV_PIX_FMT_BAYER_RGGB16LE = AVPixelFormat(C.AV_PIX_FMT_BAYER_RGGB16LE)
	AV_PIX_FMT_BAYER_RGGB16BE = AVPixelFormat(C.AV_PIX_FMT_BAYER_RGGB16BE)
	AV_PIX_FMT_BAYER_GBRG16LE = AVPixelFormat(C.AV_PIX_FMT_BAYER_GBRG16LE)
	AV_PIX_FMT_BAYER_GBRG16BE = AVPixelFormat(C.AV_PIX_FMT_BAYER_GBRG16BE)
	AV_PIX_FMT_BAYER_GRBG16LE = AVPixelFormat(C.AV_PIX_FMT_BAYER_GRBG16LE)
	AV_PIX_FMT_BAYER_GRBG16BE = AVPixelFormat(C.AV_PIX_FMT_BAYER_GRBG16BE)

	AV_PIX_FMT_XVMC = AVPixelFormat(C.AV_PIX_FMT_XVMC)

	AV_PIX_FMT_YUV440P10LE = AVPixelFormat(C.AV_PIX_FMT_YUV440P10LE)
	AV_PIX_FMT_YUV440P10BE = AVPixelFormat(C.AV_PIX_FMT_YUV440P10BE)
	AV_PIX_FMT_YUV440P12LE = AVPixelFormat(C.AV_PIX_FMT_YUV440P12LE)
	AV_PIX_FMT_YUV440P12BE = AVPixelFormat(C.AV_PIX_FMT_YUV440P12BE)
	AV_PIX_FMT_AYUV64LE    = AVPixelFormat(C.AV_PIX_FMT_AYUV64LE)
	AV_PIX_FMT_AYUV64BE    = AVPixelFormat(C.AV_PIX_FMT_AYUV64BE)

	AV_PIX_FMT_VIDEOTOOLBOX = AVPixelFormat(C.AV_PIX_FMT_VIDEOTOOLBOX)

	AV_PIX_FMT_P010LE = AVPixelFormat(C.AV_PIX_FMT_P010LE)
	AV_PIX_FMT_P010BE = AVPixelFormat(C.AV_PIX_FMT_P010BE)

	AV_PIX_FMT_GBRAP12BE = AVPixelFormat(C.AV_PIX_FMT_GBRAP12BE)
	AV_PIX_FMT_GBRAP12LE = AVPixelFormat(C.AV_PIX_FMT_GBRAP12LE)

	AV_PIX_FMT_GBRAP10BE = AVPixelFormat(C.AV_PIX_FMT_GBRAP10BE)
	AV_PIX_FMT_GBRAP10LE = AVPixelFormat(C.AV_PIX_FMT_GBRAP10LE)

	AV_PIX_FMT_MEDIACODEC = AVPixelFormat(C.AV_PIX_FMT_MEDIACODEC)

	AV_PIX_FMT_GRAY12BE = AVPixelFormat(C.AV_PIX_FMT_GRAY12BE)
	AV_PIX_FMT_GRAY12LE = AVPixelFormat(C.AV_PIX_FMT_GRAY12LE)
	AV_PIX_FMT_GRAY10BE = AVPixelFormat(C.AV_PIX_FMT_GRAY10BE)
	AV_PIX_FMT_GRAY10LE = AVPixelFormat(C.AV_PIX_FMT_GRAY10LE)

	AV_PIX_FMT_P016LE = AVPixelFormat(C.AV_PIX_FMT_P016LE)
	AV_PIX_FMT_P016BE = AVPixelFormat(C.AV_PIX_FMT_P016BE)

	AV_PIX_FMT_D3D11 = AVPixelFormat(C.AV_PIX_FMT_D3D11)

	AV_PIX_FMT_GRAY9BE = AVPixelFormat(C.AV_PIX_FMT_GRAY9BE)
	AV_PIX_FMT_GRAY9LE = AVPixelFormat(C.AV_PIX_FMT_GRAY9LE)

	AV_PIX_FMT_GBRPF32BE  = AVPixelFormat(C.AV_PIX_FMT_GBRPF32BE)
	AV_PIX_FMT_GBRPF32LE  = AVPixelFormat(C.AV_PIX_FMT_GBRPF32LE)
	AV_PIX_FMT_GBRAPF32BE = AVPixelFormat(C.AV_PIX_FMT_GBRAPF32BE)
	AV_PIX_FMT_GBRAPF32LE = AVPixelFormat(C.AV_PIX_FMT_GBRAPF32LE)

	AV_PIX_FMT_DRM_PRIME = AVPixelFormat(C.AV_PIX_FMT_DRM_PRIME)

	AV_PIX_FMT_OPENCL = AVPixelFormat(C.AV_PIX_FMT_OPENCL)

	AV_PIX_FMT_GRAY14BE = AVPixelFormat(C.AV_PIX_FMT_GRAY14BE)
	AV_PIX_FMT_GRAY14LE = AVPixelFormat(C.AV_PIX_FMT_GRAY14LE)

	AV_PIX_FMT_GRAYF32BE = AVPixelFormat(C.AV_PIX_FMT_GRAYF32BE)
	AV_PIX_FMT_GRAYF32LE = AVPixelFormat(C.AV_PIX_FMT_GRAYF32LE)

	AV_PIX_FMT_YUVA422P12BE = AVPixelFormat(C.AV_PIX_FMT_YUVA422P12BE)
	AV_PIX_FMT_YUVA422P12LE = AVPixelFormat(C.AV_PIX_FMT_YUVA422P12LE)
	AV_PIX_FMT_YUVA444P12BE = AVPixelFormat(C.AV_PIX_FMT_YUVA444P12BE)
	AV_PIX_FMT_YUVA444P12LE = AVPixelFormat(C.AV_PIX_FMT_YUVA444P12LE)

	AV_PIX_FMT_NV24 = AVPixelFormat(C.AV_PIX_FMT_NV24)
	AV_PIX_FMT_NV42 = AVPixelFormat(C.AV_PIX_FMT_NV42)

	AV_PIX_FMT_VULKAN = AVPixelFormat(C.AV_PIX_FMT_VULKAN)

	AV_PIX_FMT_Y210BE = AVPixelFormat(C.AV_PIX_FMT_Y210BE)
	AV_PIX_FMT_Y210LE = AVPixelFormat(C.AV_PIX_FMT_Y210LE)

	AV_PIX_FMT_X2RGB10LE = AVPixelFormat(C.AV_PIX_FMT_X2RGB10LE)
	AV_PIX_FMT_X2RGB10BE = AVPixelFormat(C.AV_PIX_FMT_X2RGB10BE)
	AV_PIX_FMT_X2BGR10LE = AVPixelFormat(C.AV_PIX_FMT_X2BGR10LE)
	AV_PIX_FMT_X2BGR10BE = AVPixelFormat(C.AV_PIX_FMT_X2BGR10BE)

	AV_PIX_FMT_P210BE = AVPixelFormat(C.AV_PIX_FMT_P210BE)
	AV_PIX_FMT_P210LE = AVPixelFormat(C.AV_PIX_FMT_P210LE)

	AV_PIX_FMT_P410BE = AVPixelFormat(C.AV_PIX_FMT_P410BE)
	AV_PIX_FMT_P410LE = AVPixelFormat(C.AV_PIX_FMT_P410LE)

	AV_PIX_FMT_P216BE = AVPixelFormat(C.AV_PIX_FMT_P216BE)
	AV_PIX_FMT_P216LE = AVPixelFormat(C.AV_PIX_FMT_P216LE)

	AV_PIX_FMT_P416BE = AVPixelFormat(C.AV_PIX_FMT_P416BE)
	AV_PIX_FMT_P416LE = AVPixelFormat(C.AV_PIX_FMT_P416LE)

	AV_PIX_FMT_VUYA = AVPixelFormat(C.AV_PIX_FMT_VUYA)

	AV_PIX_FMT_RGBAF16BE = AVPixelFormat(C.AV_PIX_FMT_RGBAF16BE)
	AV_PIX_FMT_RGBAF16LE = AVPixelFormat(C.AV_PIX_FMT_RGBAF16LE)

	AV_PIX_FMT_VUYX = AVPixelFormat(C.AV_PIX_FMT_VUYX)

	AV_PIX_FMT_P012LE = AVPixelFormat(C.AV_PIX_FMT_P012LE)
	AV_PIX_FMT_P012BE = AVPixelFormat(C.AV_PIX_FMT_P012BE)

	AV_PIX_FMT_Y212BE = AVPixelFormat(C.AV_PIX_FMT_Y212BE)
	AV_PIX_FMT_Y212LE = AVPixelFormat(C.AV_PIX_FMT_Y212LE)

	AV_PIX_FMT_XV30BE = AVPixelFormat(C.AV_PIX_FMT_XV30BE)
	AV_PIX_FMT_XV30LE = AVPixelFormat(C.AV_PIX_FMT_XV30LE)

	AV_PIX_FMT_XV36BE = AVPixelFormat(C.AV_PIX_FMT_XV36BE)
	AV_PIX_FMT_XV36LE = AVPixelFormat(C.AV_PIX_FMT_XV36LE)

	AV_PIX_FMT_RGBF32BE = AVPixelFormat(C.AV_PIX_FMT_RGBF32BE)
	AV_PIX_FMT_RGBF32LE = AVPixelFormat(C.AV_PIX_FMT_RGBF32LE)

	AV_PIX_FMT_RGBAF32BE = AVPixelFormat(C.AV_PIX_FMT_RGBAF32BE)
	AV_PIX_FMT_RGBAF32LE = AVPixelFormat(C.AV_PIX_FMT_RGBAF32LE)

	AV_PIX_FMT_NB = AVPixelFormat(C.AV_PIX_FMT_NB)
)

const (
	AV_PIX_FMT_RGB32   = AVPixelFormat(C.AV_PIX_FMT_RGB32)
	AV_PIX_FMT_RGB32_1 = AVPixelFormat(C.AV_PIX_FMT_RGB32_1)
	AV_PIX_FMT_BGR32   = AVPixelFormat(C.AV_PIX_FMT_BGR32)
	AV_PIX_FMT_BGR32_1 = AVPixelFormat(C.AV_PIX_FMT_BGR32_1)
	AV_PIX_FMT_0RGB32  = AVPixelFormat(C.AV_PIX_FMT_0RGB32)
	AV_PIX_FMT_0BGR32  = AVPixelFormat(C.AV_PIX_FMT_0BGR32)

	AV_PIX_FMT_GRAY9  = AVPixelFormat(C.AV_PIX_FMT_GRAY9)
	AV_PIX_FMT_GRAY10 = AVPixelFormat(C.AV_PIX_FMT_GRAY10)
	AV_PIX_FMT_GRAY12 = AVPixelFormat(C.AV_PIX_FMT_GRAY12)
	AV_PIX_FMT_GRAY14 = AVPixelFormat(C.AV_PIX_FMT_GRAY14)
	AV_PIX_FMT_GRAY16 = AVPixelFormat(C.AV_PIX_FMT_GRAY16)
	AV_PIX_FMT_YA16   = AVPixelFormat(C.AV_PIX_FMT_YA16)
	AV_PIX_FMT_RGB48  = AVPixelFormat(C.AV_PIX_FMT_RGB48)
	AV_PIX_FMT_RGB565 = AVPixelFormat(C.AV_PIX_FMT_RGB565)
	AV_PIX_FMT_RGB555 = AVPixelFormat(C.AV_PIX_FMT_RGB555)
	AV_PIX_FMT_RGB444 = AVPixelFormat(C.AV_PIX_FMT_RGB444)
	AV_PIX_FMT_RGBA64 = AVPixelFormat(C.AV_PIX_FMT_RGBA64)
	AV_PIX_FMT_BGR48  = AVPixelFormat(C.AV_PIX_FMT_BGR48)
	AV_PIX_FMT_BGR565 = AVPixelFormat(C.AV_PIX_FMT_BGR565)
	AV_PIX_FMT_BGR555 = AVPixelFormat(C.AV_PIX_FMT_BGR555)
	AV_PIX_FMT_BGR444 = AVPixelFormat(C.AV_PIX_FMT_BGR444)
	AV_PIX_FMT_BGRA64 = AVPixelFormat(C.AV_PIX_FMT_BGRA64)

	AV_PIX_FMT_YUV420P9  = AVPixelFormat(C.AV_PIX_FMT_YUV420P9)
	AV_PIX_FMT_YUV422P9  = AVPixelFormat(C.AV_PIX_FMT_YUV422P9)
	AV_PIX_FMT_YUV444P9  = AVPixelFormat(C.AV_PIX_FMT_YUV444P9)
	AV_PIX_FMT_YUV420P10 = AVPixelFormat(C.AV_PIX_FMT_YUV420P10)
	AV_PIX_FMT_YUV422P10 = AVPixelFormat(C.AV_PIX_FMT_YUV422P10)
	AV_PIX_FMT_YUV440P10 = AVPixelFormat(C.AV_PIX_FMT_YUV440P10)
	AV_PIX_FMT_YUV444P10 = AVPixelFormat(C.AV_PIX_FMT_YUV444P10)
	AV_PIX_FMT_YUV420P12 = AVPixelFormat(C.AV_PIX_FMT_YUV420P12)
	AV_PIX_FMT_YUV422P12 = AVPixelFormat(C.AV_PIX_FMT_YUV422P12)
	AV_PIX_FMT_YUV440P12 = AVPixelFormat(C.AV_PIX_FMT_YUV440P12)
	AV_PIX_FMT_YUV444P12 = AVPixelFormat(C.AV_PIX_FMT_YUV444P12)
	AV_PIX_FMT_YUV420P14 = AVPixelFormat(C.AV_PIX_FMT_YUV420P14)
	AV_PIX_FMT_YUV422P14 = AVPixelFormat(C.AV_PIX_FMT_YUV422P14)
	AV_PIX_FMT_YUV444P14 = AVPixelFormat(C.AV_PIX_FMT_YUV444P14)
	AV_PIX_FMT_YUV420P16 = AVPixelFormat(C.AV_PIX_FMT_YUV420P16)
	AV_PIX_FMT_YUV422P16 = AVPixelFormat(C.AV_PIX_FMT_YUV422P16)
	AV_PIX_FMT_YUV444P16 = AVPixelFormat(C.AV_PIX_FMT_YUV444P16)

	AV_PIX_FMT_GBRP9   = AVPixelFormat(C.AV_PIX_FMT_GBRP9)
	AV_PIX_FMT_GBRP10  = AVPixelFormat(C.AV_PIX_FMT_GBRP10)
	AV_PIX_FMT_GBRP12  = AVPixelFormat(C.AV_PIX_FMT_GBRP12)
	AV_PIX_FMT_GBRP14  = AVPixelFormat(C.AV_PIX_FMT_GBRP14)
	AV_PIX_FMT_GBRP16  = AVPixelFormat(C.AV_PIX_FMT_GBRP16)
	AV_PIX_FMT_GBRAP10 = AVPixelFormat(C.AV_PIX_FMT_GBRAP10)
	AV_PIX_FMT_GBRAP12 = AVPixelFormat(C.AV_PIX_FMT_GBRAP12)
	AV_PIX_FMT_GBRAP16 = AVPixelFormat(C.AV_PIX_FMT_GBRAP16)

	AV_PIX_FMT_BAYER_BGGR16 = AVPixelFormat(C.AV_PIX_FMT_BAYER_BGGR16)
	AV_PIX_FMT_BAYER_RGGB16 = AVPixelFormat(C.AV_PIX_FMT_BAYER_RGGB16)
	AV_PIX_FMT_BAYER_GBRG16 = AVPixelFormat(C.AV_PIX_FMT_BAYER_GBRG16)
	AV_PIX_FMT_BAYER_GRBG16 = AVPixelFormat(C.AV_PIX_FMT_BAYER_GRBG16)

	AV_PIX_FMT_GBRPF32  = AVPixelFormat(C.AV_PIX_FMT_GBRPF32)
	AV_PIX_FMT_GBRAPF32 = AVPixelFormat(C.AV_PIX_FMT_GBRAPF32)

	AV_PIX_FMT_GRAYF32 = AVPixelFormat(C.AV_PIX_FMT_GRAYF32)

	AV_PIX_FMT_YUVA420P9  = AVPixelFormat(C.AV_PIX_FMT_YUVA420P9)
	AV_PIX_FMT_YUVA422P9  = AVPixelFormat(C.AV_PIX_FMT_YUVA422P9)
	AV_PIX_FMT_YUVA444P9  = AVPixelFormat(C.AV_PIX_FMT_YUVA444P9)
	AV_PIX_FMT_YUVA420P10 = AVPixelFormat(C.AV_PIX_FMT_YUVA420P10)
	AV_PIX_FMT_YUVA422P10 = AVPixelFormat(C.AV_PIX_FMT_YUVA422P10)
	AV_PIX_FMT_YUVA444P10 = AVPixelFormat(C.AV_PIX_FMT_YUVA444P10)
	AV_PIX_FMT_YUVA422P12 = AVPixelFormat(C.AV_PIX_FMT_YUVA422P12)
	AV_PIX_FMT_YUVA444P12 = AVPixelFormat(C.AV_PIX_FMT_YUVA444P12)
	AV_PIX_FMT_YUVA420P16 = AVPixelFormat(C.AV_PIX_FMT_YUVA420P16)
	AV_PIX_FMT_YUVA422P16 = AVPixelFormat(C.AV_PIX_FMT_YUVA422P16)
	AV_PIX_FMT_YUVA444P16 = AVPixelFormat(C.AV_PIX_FMT_YUVA444P16)

	AV_PIX_FMT_XYZ12  = AVPixelFormat(C.AV_PIX_FMT_XYZ12)
	AV_PIX_FMT_NV20   = AVPixelFormat(C.AV_PIX_FMT_NV20)
	AV_PIX_FMT_AYUV64 = AVPixelFormat(C.AV_PIX_FMT_AYUV64)
	AV_PIX_FMT_P010   = AVPixelFormat(C.AV_PIX_FMT_P010)
	AV_PIX_FMT_P016   = AVPixelFormat(C.AV_PIX_FMT_P016)

	AV_PIX_FMT_Y210    = AVPixelFormat(C.AV_PIX_FMT_Y210)
	AV_PIX_FMT_X2RGB10 = AVPixelFormat(C.AV_PIX_FMT_X2RGB10)
	AV_PIX_FMT_X2BGR10 = AVPixelFormat(C.AV_PIX_FMT_X2BGR10)

	AV_PIX_FMT_P210 = AVPixelFormat(C.AV_PIX_FMT_P210)
	AV_PIX_FMT_P410 = AVPixelFormat(C.AV_PIX_FMT_P410)
	AV_PIX_FMT_P216 = AVPixelFormat(C.AV_PIX_FMT_P216)
	AV_PIX_FMT_P416 = AVPixelFormat(C.AV_PIX_FMT_P416)
)

// Chromaticity coordinates of the source primaries.
type AVColorPrimaries = C.enum_AVColorPrimaries

const (
	AVCOL_PRI_RESERVED0   = AVColorPrimaries(C.AVCOL_PRI_RESERVED0)
	AVCOL_PRI_BT709       = AVColorPrimaries(C.AVCOL_PRI_BT709)
	AVCOL_PRI_UNSPECIFIED = AVColorPrimaries(C.AVCOL_PRI_UNSPECIFIED)
	AVCOL_PRI_RESERVED    = AVColorPrimaries(C.AVCOL_PRI_RESERVED)
	AVCOL_PRI_BT470M      = AVColorPrimaries(C.AVCOL_PRI_BT470M)

	AVCOL_PRI_BT470BG      = AVColorPrimaries(C.AVCOL_PRI_BT470BG)
	AVCOL_PRI_SMPTE170M    = AVColorPrimaries(C.AVCOL_PRI_SMPTE170M)
	AVCOL_PRI_SMPTE240M    = AVColorPrimaries(C.AVCOL_PRI_SMPTE240M)
	AVCOL_PRI_FILM         = AVColorPrimaries(C.AVCOL_PRI_FILM)
	AVCOL_PRI_BT2020       = AVColorPrimaries(C.AVCOL_PRI_BT2020)
	AVCOL_PRI_SMPTE428     = AVColorPrimaries(C.AVCOL_PRI_SMPTE428)
	AVCOL_PRI_SMPTEST428_1 = AVColorPrimaries(C.AVCOL_PRI_SMPTEST428_1)
	AVCOL_PRI_SMPTE431     = AVColorPrimaries(C.AVCOL_PRI_SMPTE431)
	AVCOL_PRI_SMPTE432     = AVColorPrimaries(C.AVCOL_PRI_SMPTE432)
	AVCOL_PRI_EBU3213      = AVColorPrimaries(C.AVCOL_PRI_EBU3213)
	AVCOL_PRI_JEDEC_P22    = AVColorPrimaries(C.AVCOL_PRI_JEDEC_P22)
	AVCOL_PRI_NB           = AVColorPrimaries(C.AVCOL_PRI_NB)
)

// Color Transfer Characteristic.
type AVColorTransferCharacteristic = C.enum_AVColorTransferCharacteristic

const (
	AVCOL_TRC_RESERVED0    = AVColorTransferCharacteristic(C.AVCOL_TRC_RESERVED0)
	AVCOL_TRC_BT709        = AVColorTransferCharacteristic(C.AVCOL_TRC_BT709)
	AVCOL_TRC_UNSPECIFIED  = AVColorTransferCharacteristic(C.AVCOL_TRC_UNSPECIFIED)
	AVCOL_TRC_RESERVED     = AVColorTransferCharacteristic(C.AVCOL_TRC_RESERVED)
	AVCOL_TRC_GAMMA22      = AVColorTransferCharacteristic(C.AVCOL_TRC_GAMMA22)
	AVCOL_TRC_GAMMA28      = AVColorTransferCharacteristic(C.AVCOL_TRC_GAMMA28)
	AVCOL_TRC_SMPTE170M    = AVColorTransferCharacteristic(C.AVCOL_TRC_SMPTE170M)
	AVCOL_TRC_SMPTE240M    = AVColorTransferCharacteristic(C.AVCOL_TRC_SMPTE240M)
	AVCOL_TRC_LINEAR       = AVColorTransferCharacteristic(C.AVCOL_TRC_LINEAR)
	AVCOL_TRC_LOG          = AVColorTransferCharacteristic(C.AVCOL_TRC_LOG)
	AVCOL_TRC_LOG_SQRT     = AVColorTransferCharacteristic(C.AVCOL_TRC_LOG_SQRT)
	AVCOL_TRC_IEC61966_2_4 = AVColorTransferCharacteristic(C.AVCOL_TRC_IEC61966_2_4)
	AVCOL_TRC_BT1361_ECG   = AVColorTransferCharacteristic(C.AVCOL_TRC_BT1361_ECG)
	AVCOL_TRC_IEC61966_2_1 = AVColorTransferCharacteristic(C.AVCOL_TRC_IEC61966_2_1)
	AVCOL_TRC_BT2020_10    = AVColorTransferCharacteristic(C.AVCOL_TRC_BT2020_10)
	AVCOL_TRC_BT2020_12    = AVColorTransferCharacteristic(C.AVCOL_TRC_BT2020_12)
	AVCOL_TRC_SMPTE2084    = AVColorTransferCharacteristic(C.AVCOL_TRC_SMPTE2084)
	AVCOL_TRC_SMPTEST2084  = AVColorTransferCharacteristic(C.AVCOL_TRC_SMPTEST2084)
	AVCOL_TRC_SMPTE428     = AVColorTransferCharacteristic(C.AVCOL_TRC_SMPTE428)
	AVCOL_TRC_SMPTEST428_1 = AVColorTransferCharacteristic(C.AVCOL_TRC_SMPTEST428_1)
	AVCOL_TRC_ARIB_STD_B67 = AVColorTransferCharacteristic(C.AVCOL_TRC_ARIB_STD_B67)
	AVCOL_TRC_NB           = AVColorTransferCharacteristic(C.AVCOL_TRC_NB)
)

// AVColorSpace
type AVColorSpace = C.enum_AVColorSpace

const (
	AVCOL_SPC_RGB                = AVColorSpace(C.AVCOL_SPC_RGB)
	AVCOL_SPC_BT709              = AVColorSpace(C.AVCOL_SPC_BT709)
	AVCOL_SPC_UNSPECIFIED        = AVColorSpace(C.AVCOL_SPC_UNSPECIFIED)
	AVCOL_SPC_RESERVED           = AVColorSpace(C.AVCOL_SPC_RESERVED)
	AVCOL_SPC_FCC                = AVColorSpace(C.AVCOL_SPC_FCC)
	AVCOL_SPC_BT470BG            = AVColorSpace(C.AVCOL_SPC_BT470BG)
	AVCOL_SPC_SMPTE170M          = AVColorSpace(C.AVCOL_SPC_SMPTE170M)
	AVCOL_SPC_SMPTE240M          = AVColorSpace(C.AVCOL_SPC_SMPTE240M)
	AVCOL_SPC_YCGCO              = AVColorSpace(C.AVCOL_SPC_YCGCO)
	AVCOL_SPC_YCOCG              = AVColorSpace(C.AVCOL_SPC_YCOCG)
	AVCOL_SPC_BT2020_NCL         = AVColorSpace(C.AVCOL_SPC_BT2020_NCL)
	AVCOL_SPC_BT2020_CL          = AVColorSpace(C.AVCOL_SPC_BT2020_CL)
	AVCOL_SPC_SMPTE2085          = AVColorSpace(C.AVCOL_SPC_SMPTE2085)
	AVCOL_SPC_CHROMA_DERIVED_NCL = AVColorSpace(C.AVCOL_SPC_CHROMA_DERIVED_NCL)
	AVCOL_SPC_CHROMA_DERIVED_CL  = AVColorSpace(C.AVCOL_SPC_CHROMA_DERIVED_CL)
	AVCOL_SPC_ICTCP              = AVColorSpace(C.AVCOL_SPC_ICTCP)
	AVCOL_SPC_NB                 = AVColorSpace(C.AVCOL_SPC_NB)
)

// AVColorRange
type AVColorRange = C.enum_AVColorRange

const (
	AVCOL_RANGE_UNSPECIFIED = AVColorRange(C.AVCOL_RANGE_UNSPECIFIED)
	AVCOL_RANGE_MPEG        = AVColorRange(C.AVCOL_RANGE_MPEG)
	AVCOL_RANGE_JPEG        = AVColorRange(C.AVCOL_RANGE_JPEG)
	AVCOL_RANGE_NB          = AVColorRange(C.AVCOL_RANGE_NB)
)

// AVChromaLocation
type AVChromaLocation = C.enum_AVChromaLocation

const (
	AVCHROMA_LOC_UNSPECIFIED = AVChromaLocation(C.AVCHROMA_LOC_UNSPECIFIED)
	AVCHROMA_LOC_LEFT        = AVChromaLocation(C.AVCHROMA_LOC_LEFT)
	AVCHROMA_LOC_CENTER      = AVChromaLocation(C.AVCHROMA_LOC_CENTER)
	AVCHROMA_LOC_TOPLEFT     = AVChromaLocation(C.AVCHROMA_LOC_TOPLEFT)
	AVCHROMA_LOC_TOP         = AVChromaLocation(C.AVCHROMA_LOC_TOP)
	AVCHROMA_LOC_BOTTOMLEFT  = AVChromaLocation(C.AVCHROMA_LOC_BOTTOMLEFT)
	AVCHROMA_LOC_BOTTOM      = AVChromaLocation(C.AVCHROMA_LOC_BOTTOM)
	AVCHROMA_LOC_NB          = AVChromaLocation(C.AVCHROMA_LOC_NB)
)
