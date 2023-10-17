package ffmpeg

/*
#include <libavutil/pixfmt.h>
*/
import "C"

// Pixel format.
type AvPixelFormat = C.enum_AVPixelFormat

const (
	AV_PIX_FMT_NONE      = AvPixelFormat(C.AV_PIX_FMT_NONE)
	AV_PIX_FMT_YUV420P   = AvPixelFormat(C.AV_PIX_FMT_YUV420P)
	AV_PIX_FMT_YUYV422   = AvPixelFormat(C.AV_PIX_FMT_YUYV422)
	AV_PIX_FMT_RGB24     = AvPixelFormat(C.AV_PIX_FMT_RGB24)
	AV_PIX_FMT_BGR24     = AvPixelFormat(C.AV_PIX_FMT_BGR24)
	AV_PIX_FMT_YUV422P   = AvPixelFormat(C.AV_PIX_FMT_YUV422P)
	AV_PIX_FMT_YUV444P   = AvPixelFormat(C.AV_PIX_FMT_YUV444P)
	AV_PIX_FMT_YUV410P   = AvPixelFormat(C.AV_PIX_FMT_YUV410P)
	AV_PIX_FMT_YUV411P   = AvPixelFormat(C.AV_PIX_FMT_YUV411P)
	AV_PIX_FMT_GRAY8     = AvPixelFormat(C.AV_PIX_FMT_GRAY8)
	AV_PIX_FMT_MONOWHITE = AvPixelFormat(C.AV_PIX_FMT_MONOWHITE)
	AV_PIX_FMT_MONOBLACK = AvPixelFormat(C.AV_PIX_FMT_MONOBLACK)
	AV_PIX_FMT_PAL8      = AvPixelFormat(C.AV_PIX_FMT_PAL8)
	AV_PIX_FMT_YUVJ420P  = AvPixelFormat(C.AV_PIX_FMT_YUVJ420P)
	AV_PIX_FMT_YUVJ422P  = AvPixelFormat(C.AV_PIX_FMT_YUVJ422P)
	AV_PIX_FMT_YUVJ444P  = AvPixelFormat(C.AV_PIX_FMT_YUVJ444P)
	AV_PIX_FMT_UYVY422   = AvPixelFormat(C.AV_PIX_FMT_UYVY422)
	AV_PIX_FMT_UYYVYY411 = AvPixelFormat(C.AV_PIX_FMT_UYYVYY411)
	AV_PIX_FMT_BGR8      = AvPixelFormat(C.AV_PIX_FMT_BGR8)
	AV_PIX_FMT_BGR4      = AvPixelFormat(C.AV_PIX_FMT_BGR4)
	AV_PIX_FMT_BGR4_BYTE = AvPixelFormat(C.AV_PIX_FMT_BGR4_BYTE)
	AV_PIX_FMT_RGB8      = AvPixelFormat(C.AV_PIX_FMT_RGB8)
	AV_PIX_FMT_RGB4      = AvPixelFormat(C.AV_PIX_FMT_RGB4)
	AV_PIX_FMT_RGB4_BYTE = AvPixelFormat(C.AV_PIX_FMT_RGB4_BYTE)
	AV_PIX_FMT_NV12      = AvPixelFormat(C.AV_PIX_FMT_NV12)
	AV_PIX_FMT_NV21      = AvPixelFormat(C.AV_PIX_FMT_NV21)

	AV_PIX_FMT_ARGB = AvPixelFormat(C.AV_PIX_FMT_ARGB)
	AV_PIX_FMT_RGBA = AvPixelFormat(C.AV_PIX_FMT_RGBA)
	AV_PIX_FMT_ABGR = AvPixelFormat(C.AV_PIX_FMT_ABGR)
	AV_PIX_FMT_BGRA = AvPixelFormat(C.AV_PIX_FMT_BGRA)

	AV_PIX_FMT_GRAY16BE = AvPixelFormat(C.AV_PIX_FMT_GRAY16BE)
	AV_PIX_FMT_GRAY16LE = AvPixelFormat(C.AV_PIX_FMT_GRAY16LE)
	AV_PIX_FMT_YUV440P  = AvPixelFormat(C.AV_PIX_FMT_YUV440P)
	AV_PIX_FMT_YUVJ440P = AvPixelFormat(C.AV_PIX_FMT_YUVJ440P)
	AV_PIX_FMT_YUVA420P = AvPixelFormat(C.AV_PIX_FMT_YUVA420P)
	AV_PIX_FMT_RGB48BE  = AvPixelFormat(C.AV_PIX_FMT_RGB48BE)
	AV_PIX_FMT_RGB48LE  = AvPixelFormat(C.AV_PIX_FMT_RGB48LE)

	AV_PIX_FMT_RGB565BE = AvPixelFormat(C.AV_PIX_FMT_RGB565BE)
	AV_PIX_FMT_RGB565LE = AvPixelFormat(C.AV_PIX_FMT_RGB565LE)
	AV_PIX_FMT_RGB555BE = AvPixelFormat(C.AV_PIX_FMT_RGB555BE)
	AV_PIX_FMT_RGB555LE = AvPixelFormat(C.AV_PIX_FMT_RGB555LE)

	AV_PIX_FMT_BGR565BE = AvPixelFormat(C.AV_PIX_FMT_BGR565BE)
	AV_PIX_FMT_BGR565LE = AvPixelFormat(C.AV_PIX_FMT_BGR565LE)
	AV_PIX_FMT_BGR555BE = AvPixelFormat(C.AV_PIX_FMT_BGR555BE)
	AV_PIX_FMT_BGR555LE = AvPixelFormat(C.AV_PIX_FMT_BGR555LE)

	AV_PIX_FMT_VAAPI_MOCO = AvPixelFormat(C.AV_PIX_FMT_VAAPI_MOCO)
	AV_PIX_FMT_VAAPI_IDCT = AvPixelFormat(C.AV_PIX_FMT_VAAPI_IDCT)
	AV_PIX_FMT_VAAPI_VLD  = AvPixelFormat(C.AV_PIX_FMT_VAAPI_VLD)
	AV_PIX_FMT_VAAPI      = AvPixelFormat(C.AV_PIX_FMT_VAAPI)

	AV_PIX_FMT_YUV420P16LE = AvPixelFormat(C.AV_PIX_FMT_YUV420P16LE)
	AV_PIX_FMT_YUV420P16BE = AvPixelFormat(C.AV_PIX_FMT_YUV420P16BE)
	AV_PIX_FMT_YUV422P16LE = AvPixelFormat(C.AV_PIX_FMT_YUV422P16LE)
	AV_PIX_FMT_YUV422P16BE = AvPixelFormat(C.AV_PIX_FMT_YUV422P16BE)
	AV_PIX_FMT_YUV444P16LE = AvPixelFormat(C.AV_PIX_FMT_YUV444P16LE)
	AV_PIX_FMT_YUV444P16BE = AvPixelFormat(C.AV_PIX_FMT_YUV444P16BE)
	AV_PIX_FMT_DXVA2_VLD   = AvPixelFormat(C.AV_PIX_FMT_DXVA2_VLD)

	AV_PIX_FMT_RGB444LE = AvPixelFormat(C.AV_PIX_FMT_RGB444LE)
	AV_PIX_FMT_RGB444BE = AvPixelFormat(C.AV_PIX_FMT_RGB444BE)
	AV_PIX_FMT_BGR444LE = AvPixelFormat(C.AV_PIX_FMT_BGR444LE)
	AV_PIX_FMT_BGR444BE = AvPixelFormat(C.AV_PIX_FMT_BGR444BE)
	AV_PIX_FMT_YA8      = AvPixelFormat(C.AV_PIX_FMT_YA8)

	AV_PIX_FMT_Y400A  = AvPixelFormat(C.AV_PIX_FMT_Y400A)
	AV_PIX_FMT_GRAY8A = AvPixelFormat(C.AV_PIX_FMT_GRAY8A)

	AV_PIX_FMT_BGR48BE = AvPixelFormat(C.AV_PIX_FMT_BGR48BE)
	AV_PIX_FMT_BGR48LE = AvPixelFormat(C.AV_PIX_FMT_BGR48LE)

	AV_PIX_FMT_YUV420P9BE   = AvPixelFormat(C.AV_PIX_FMT_YUV420P9BE)
	AV_PIX_FMT_YUV420P9LE   = AvPixelFormat(C.AV_PIX_FMT_YUV420P9LE)
	AV_PIX_FMT_YUV420P10BE  = AvPixelFormat(C.AV_PIX_FMT_YUV420P10BE)
	AV_PIX_FMT_YUV420P10LE  = AvPixelFormat(C.AV_PIX_FMT_YUV420P10LE)
	AV_PIX_FMT_YUV422P10BE  = AvPixelFormat(C.AV_PIX_FMT_YUV422P10BE)
	AV_PIX_FMT_YUV422P10LE  = AvPixelFormat(C.AV_PIX_FMT_YUV422P10LE)
	AV_PIX_FMT_YUV444P9BE   = AvPixelFormat(C.AV_PIX_FMT_YUV444P9BE)
	AV_PIX_FMT_YUV444P9LE   = AvPixelFormat(C.AV_PIX_FMT_YUV444P9LE)
	AV_PIX_FMT_YUV444P10BE  = AvPixelFormat(C.AV_PIX_FMT_YUV444P10BE)
	AV_PIX_FMT_YUV444P10LE  = AvPixelFormat(C.AV_PIX_FMT_YUV444P10LE)
	AV_PIX_FMT_YUV422P9BE   = AvPixelFormat(C.AV_PIX_FMT_YUV422P9BE)
	AV_PIX_FMT_YUV422P9LE   = AvPixelFormat(C.AV_PIX_FMT_YUV422P9LE)
	AV_PIX_FMT_GBRP         = AvPixelFormat(C.AV_PIX_FMT_GBRP)
	AV_PIX_FMT_GBR24P       = AvPixelFormat(C.AV_PIX_FMT_GBR24P)
	AV_PIX_FMT_GBRP9BE      = AvPixelFormat(C.AV_PIX_FMT_GBRP9BE)
	AV_PIX_FMT_GBRP9LE      = AvPixelFormat(C.AV_PIX_FMT_GBRP9LE)
	AV_PIX_FMT_GBRP10BE     = AvPixelFormat(C.AV_PIX_FMT_GBRP10BE)
	AV_PIX_FMT_GBRP10LE     = AvPixelFormat(C.AV_PIX_FMT_GBRP10LE)
	AV_PIX_FMT_GBRP16BE     = AvPixelFormat(C.AV_PIX_FMT_GBRP16BE)
	AV_PIX_FMT_GBRP16LE     = AvPixelFormat(C.AV_PIX_FMT_GBRP16LE)
	AV_PIX_FMT_YUVA422P     = AvPixelFormat(C.AV_PIX_FMT_YUVA422P)
	AV_PIX_FMT_YUVA444P     = AvPixelFormat(C.AV_PIX_FMT_YUVA444P)
	AV_PIX_FMT_YUVA420P9BE  = AvPixelFormat(C.AV_PIX_FMT_YUVA420P9BE)
	AV_PIX_FMT_YUVA420P9LE  = AvPixelFormat(C.AV_PIX_FMT_YUVA420P9LE)
	AV_PIX_FMT_YUVA422P9BE  = AvPixelFormat(C.AV_PIX_FMT_YUVA422P9BE)
	AV_PIX_FMT_YUVA422P9LE  = AvPixelFormat(C.AV_PIX_FMT_YUVA422P9LE)
	AV_PIX_FMT_YUVA444P9BE  = AvPixelFormat(C.AV_PIX_FMT_YUVA444P9BE)
	AV_PIX_FMT_YUVA444P9LE  = AvPixelFormat(C.AV_PIX_FMT_YUVA444P9LE)
	AV_PIX_FMT_YUVA420P10BE = AvPixelFormat(C.AV_PIX_FMT_YUVA420P10BE)
	AV_PIX_FMT_YUVA420P10LE = AvPixelFormat(C.AV_PIX_FMT_YUVA420P10LE)
	AV_PIX_FMT_YUVA422P10BE = AvPixelFormat(C.AV_PIX_FMT_YUVA422P10BE)
	AV_PIX_FMT_YUVA422P10LE = AvPixelFormat(C.AV_PIX_FMT_YUVA422P10LE)
	AV_PIX_FMT_YUVA444P10BE = AvPixelFormat(C.AV_PIX_FMT_YUVA444P10BE)
	AV_PIX_FMT_YUVA444P10LE = AvPixelFormat(C.AV_PIX_FMT_YUVA444P10LE)
	AV_PIX_FMT_YUVA420P16BE = AvPixelFormat(C.AV_PIX_FMT_YUVA420P16BE)
	AV_PIX_FMT_YUVA420P16LE = AvPixelFormat(C.AV_PIX_FMT_YUVA420P16LE)
	AV_PIX_FMT_YUVA422P16BE = AvPixelFormat(C.AV_PIX_FMT_YUVA422P16BE)
	AV_PIX_FMT_YUVA422P16LE = AvPixelFormat(C.AV_PIX_FMT_YUVA422P16LE)
	AV_PIX_FMT_YUVA444P16BE = AvPixelFormat(C.AV_PIX_FMT_YUVA444P16BE)
	AV_PIX_FMT_YUVA444P16LE = AvPixelFormat(C.AV_PIX_FMT_YUVA444P16LE)

	AV_PIX_FMT_VDPAU = AvPixelFormat(C.AV_PIX_FMT_VDPAU)

	AV_PIX_FMT_XYZ12LE  = AvPixelFormat(C.AV_PIX_FMT_XYZ12LE)
	AV_PIX_FMT_XYZ12BE  = AvPixelFormat(C.AV_PIX_FMT_XYZ12BE)
	AV_PIX_FMT_NV16     = AvPixelFormat(C.AV_PIX_FMT_NV16)
	AV_PIX_FMT_NV20LE   = AvPixelFormat(C.AV_PIX_FMT_NV20LE)
	AV_PIX_FMT_NV20BE   = AvPixelFormat(C.AV_PIX_FMT_NV20BE)
	AV_PIX_FMT_RGBA64BE = AvPixelFormat(C.AV_PIX_FMT_RGBA64BE)
	AV_PIX_FMT_RGBA64LE = AvPixelFormat(C.AV_PIX_FMT_RGBA64LE)
	AV_PIX_FMT_BGRA64BE = AvPixelFormat(C.AV_PIX_FMT_BGRA64BE)
	AV_PIX_FMT_BGRA64LE = AvPixelFormat(C.AV_PIX_FMT_BGRA64LE)

	AV_PIX_FMT_YVYU422 = AvPixelFormat(C.AV_PIX_FMT_YVYU422)

	AV_PIX_FMT_YA16BE = AvPixelFormat(C.AV_PIX_FMT_YA16BE)
	AV_PIX_FMT_YA16LE = AvPixelFormat(C.AV_PIX_FMT_YA16LE)

	AV_PIX_FMT_GBRAP     = AvPixelFormat(C.AV_PIX_FMT_GBRAP)
	AV_PIX_FMT_GBRAP16BE = AvPixelFormat(C.AV_PIX_FMT_GBRAP16BE)
	AV_PIX_FMT_GBRAP16LE = AvPixelFormat(C.AV_PIX_FMT_GBRAP16LE)

	AV_PIX_FMT_QSV = AvPixelFormat(C.AV_PIX_FMT_QSV)

	AV_PIX_FMT_MMAL = AvPixelFormat(C.AV_PIX_FMT_MMAL)

	AV_PIX_FMT_D3D11VA_VLD = AvPixelFormat(C.AV_PIX_FMT_D3D11VA_VLD)

	AV_PIX_FMT_CUDA = AvPixelFormat(C.AV_PIX_FMT_CUDA)

	AV_PIX_FMT_0RGB = AvPixelFormat(C.AV_PIX_FMT_0RGB)
	AV_PIX_FMT_RGB0 = AvPixelFormat(C.AV_PIX_FMT_RGB0)
	AV_PIX_FMT_0BGR = AvPixelFormat(C.AV_PIX_FMT_0BGR)
	AV_PIX_FMT_BGR0 = AvPixelFormat(C.AV_PIX_FMT_BGR0)

	AV_PIX_FMT_YUV420P12BE = AvPixelFormat(C.AV_PIX_FMT_YUV420P12BE)
	AV_PIX_FMT_YUV420P12LE = AvPixelFormat(C.AV_PIX_FMT_YUV420P12LE)
	AV_PIX_FMT_YUV420P14BE = AvPixelFormat(C.AV_PIX_FMT_YUV420P14BE)
	AV_PIX_FMT_YUV420P14LE = AvPixelFormat(C.AV_PIX_FMT_YUV420P14LE)
	AV_PIX_FMT_YUV422P12BE = AvPixelFormat(C.AV_PIX_FMT_YUV422P12BE)
	AV_PIX_FMT_YUV422P12LE = AvPixelFormat(C.AV_PIX_FMT_YUV422P12LE)
	AV_PIX_FMT_YUV422P14BE = AvPixelFormat(C.AV_PIX_FMT_YUV422P14BE)
	AV_PIX_FMT_YUV422P14LE = AvPixelFormat(C.AV_PIX_FMT_YUV422P14LE)
	AV_PIX_FMT_YUV444P12BE = AvPixelFormat(C.AV_PIX_FMT_YUV444P12BE)
	AV_PIX_FMT_YUV444P12LE = AvPixelFormat(C.AV_PIX_FMT_YUV444P12LE)
	AV_PIX_FMT_YUV444P14BE = AvPixelFormat(C.AV_PIX_FMT_YUV444P14BE)
	AV_PIX_FMT_YUV444P14LE = AvPixelFormat(C.AV_PIX_FMT_YUV444P14LE)
	AV_PIX_FMT_GBRP12BE    = AvPixelFormat(C.AV_PIX_FMT_GBRP12BE)
	AV_PIX_FMT_GBRP12LE    = AvPixelFormat(C.AV_PIX_FMT_GBRP12LE)
	AV_PIX_FMT_GBRP14BE    = AvPixelFormat(C.AV_PIX_FMT_GBRP14BE)
	AV_PIX_FMT_GBRP14LE    = AvPixelFormat(C.AV_PIX_FMT_GBRP14LE)
	AV_PIX_FMT_YUVJ411P    = AvPixelFormat(C.AV_PIX_FMT_YUVJ411P)

	AV_PIX_FMT_BAYER_BGGR8    = AvPixelFormat(C.AV_PIX_FMT_BAYER_BGGR8)
	AV_PIX_FMT_BAYER_RGGB8    = AvPixelFormat(C.AV_PIX_FMT_BAYER_RGGB8)
	AV_PIX_FMT_BAYER_GBRG8    = AvPixelFormat(C.AV_PIX_FMT_BAYER_GBRG8)
	AV_PIX_FMT_BAYER_GRBG8    = AvPixelFormat(C.AV_PIX_FMT_BAYER_GRBG8)
	AV_PIX_FMT_BAYER_BGGR16LE = AvPixelFormat(C.AV_PIX_FMT_BAYER_BGGR16LE)
	AV_PIX_FMT_BAYER_BGGR16BE = AvPixelFormat(C.AV_PIX_FMT_BAYER_BGGR16BE)
	AV_PIX_FMT_BAYER_RGGB16LE = AvPixelFormat(C.AV_PIX_FMT_BAYER_RGGB16LE)
	AV_PIX_FMT_BAYER_RGGB16BE = AvPixelFormat(C.AV_PIX_FMT_BAYER_RGGB16BE)
	AV_PIX_FMT_BAYER_GBRG16LE = AvPixelFormat(C.AV_PIX_FMT_BAYER_GBRG16LE)
	AV_PIX_FMT_BAYER_GBRG16BE = AvPixelFormat(C.AV_PIX_FMT_BAYER_GBRG16BE)
	AV_PIX_FMT_BAYER_GRBG16LE = AvPixelFormat(C.AV_PIX_FMT_BAYER_GRBG16LE)
	AV_PIX_FMT_BAYER_GRBG16BE = AvPixelFormat(C.AV_PIX_FMT_BAYER_GRBG16BE)

	AV_PIX_FMT_XVMC = AvPixelFormat(C.AV_PIX_FMT_XVMC)

	AV_PIX_FMT_YUV440P10LE = AvPixelFormat(C.AV_PIX_FMT_YUV440P10LE)
	AV_PIX_FMT_YUV440P10BE = AvPixelFormat(C.AV_PIX_FMT_YUV440P10BE)
	AV_PIX_FMT_YUV440P12LE = AvPixelFormat(C.AV_PIX_FMT_YUV440P12LE)
	AV_PIX_FMT_YUV440P12BE = AvPixelFormat(C.AV_PIX_FMT_YUV440P12BE)
	AV_PIX_FMT_AYUV64LE    = AvPixelFormat(C.AV_PIX_FMT_AYUV64LE)
	AV_PIX_FMT_AYUV64BE    = AvPixelFormat(C.AV_PIX_FMT_AYUV64BE)

	AV_PIX_FMT_VIDEOTOOLBOX = AvPixelFormat(C.AV_PIX_FMT_VIDEOTOOLBOX)

	AV_PIX_FMT_P010LE = AvPixelFormat(C.AV_PIX_FMT_P010LE)
	AV_PIX_FMT_P010BE = AvPixelFormat(C.AV_PIX_FMT_P010BE)

	AV_PIX_FMT_GBRAP12BE = AvPixelFormat(C.AV_PIX_FMT_GBRAP12BE)
	AV_PIX_FMT_GBRAP12LE = AvPixelFormat(C.AV_PIX_FMT_GBRAP12LE)

	AV_PIX_FMT_GBRAP10BE = AvPixelFormat(C.AV_PIX_FMT_GBRAP10BE)
	AV_PIX_FMT_GBRAP10LE = AvPixelFormat(C.AV_PIX_FMT_GBRAP10LE)

	AV_PIX_FMT_MEDIACODEC = AvPixelFormat(C.AV_PIX_FMT_MEDIACODEC)

	AV_PIX_FMT_GRAY12BE = AvPixelFormat(C.AV_PIX_FMT_GRAY12BE)
	AV_PIX_FMT_GRAY12LE = AvPixelFormat(C.AV_PIX_FMT_GRAY12LE)
	AV_PIX_FMT_GRAY10BE = AvPixelFormat(C.AV_PIX_FMT_GRAY10BE)
	AV_PIX_FMT_GRAY10LE = AvPixelFormat(C.AV_PIX_FMT_GRAY10LE)

	AV_PIX_FMT_P016LE = AvPixelFormat(C.AV_PIX_FMT_P016LE)
	AV_PIX_FMT_P016BE = AvPixelFormat(C.AV_PIX_FMT_P016BE)

	AV_PIX_FMT_D3D11 = AvPixelFormat(C.AV_PIX_FMT_D3D11)

	AV_PIX_FMT_GRAY9BE = AvPixelFormat(C.AV_PIX_FMT_GRAY9BE)
	AV_PIX_FMT_GRAY9LE = AvPixelFormat(C.AV_PIX_FMT_GRAY9LE)

	AV_PIX_FMT_GBRPF32BE  = AvPixelFormat(C.AV_PIX_FMT_GBRPF32BE)
	AV_PIX_FMT_GBRPF32LE  = AvPixelFormat(C.AV_PIX_FMT_GBRPF32LE)
	AV_PIX_FMT_GBRAPF32BE = AvPixelFormat(C.AV_PIX_FMT_GBRAPF32BE)
	AV_PIX_FMT_GBRAPF32LE = AvPixelFormat(C.AV_PIX_FMT_GBRAPF32LE)

	AV_PIX_FMT_DRM_PRIME = AvPixelFormat(C.AV_PIX_FMT_DRM_PRIME)

	AV_PIX_FMT_OPENCL = AvPixelFormat(C.AV_PIX_FMT_OPENCL)

	AV_PIX_FMT_GRAY14BE = AvPixelFormat(C.AV_PIX_FMT_GRAY14BE)
	AV_PIX_FMT_GRAY14LE = AvPixelFormat(C.AV_PIX_FMT_GRAY14LE)

	AV_PIX_FMT_GRAYF32BE = AvPixelFormat(C.AV_PIX_FMT_GRAYF32BE)
	AV_PIX_FMT_GRAYF32LE = AvPixelFormat(C.AV_PIX_FMT_GRAYF32LE)

	AV_PIX_FMT_YUVA422P12BE = AvPixelFormat(C.AV_PIX_FMT_YUVA422P12BE)
	AV_PIX_FMT_YUVA422P12LE = AvPixelFormat(C.AV_PIX_FMT_YUVA422P12LE)
	AV_PIX_FMT_YUVA444P12BE = AvPixelFormat(C.AV_PIX_FMT_YUVA444P12BE)
	AV_PIX_FMT_YUVA444P12LE = AvPixelFormat(C.AV_PIX_FMT_YUVA444P12LE)

	AV_PIX_FMT_NV24 = AvPixelFormat(C.AV_PIX_FMT_NV24)
	AV_PIX_FMT_NV42 = AvPixelFormat(C.AV_PIX_FMT_NV42)

	AV_PIX_FMT_VULKAN = AvPixelFormat(C.AV_PIX_FMT_VULKAN)

	AV_PIX_FMT_Y210BE = AvPixelFormat(C.AV_PIX_FMT_Y210BE)
	AV_PIX_FMT_Y210LE = AvPixelFormat(C.AV_PIX_FMT_Y210LE)

	AV_PIX_FMT_X2RGB10LE = AvPixelFormat(C.AV_PIX_FMT_X2RGB10LE)
	AV_PIX_FMT_X2RGB10BE = AvPixelFormat(C.AV_PIX_FMT_X2RGB10BE)
	AV_PIX_FMT_NB        = AvPixelFormat(C.AV_PIX_FMT_NB)
)

const (
	AV_PIX_FMT_RGB32   = AvPixelFormat(C.AV_PIX_FMT_RGB32)
	AV_PIX_FMT_RGB32_1 = AvPixelFormat(C.AV_PIX_FMT_RGB32_1)
	AV_PIX_FMT_BGR32   = AvPixelFormat(C.AV_PIX_FMT_BGR32)
	AV_PIX_FMT_BGR32_1 = AvPixelFormat(C.AV_PIX_FMT_BGR32_1)
	AV_PIX_FMT_0RGB32  = AvPixelFormat(C.AV_PIX_FMT_0RGB32)
	AV_PIX_FMT_0BGR32  = AvPixelFormat(C.AV_PIX_FMT_0BGR32)

	AV_PIX_FMT_GRAY9  = AvPixelFormat(C.AV_PIX_FMT_GRAY9)
	AV_PIX_FMT_GRAY10 = AvPixelFormat(C.AV_PIX_FMT_GRAY10)
	AV_PIX_FMT_GRAY12 = AvPixelFormat(C.AV_PIX_FMT_GRAY12)
	AV_PIX_FMT_GRAY14 = AvPixelFormat(C.AV_PIX_FMT_GRAY14)
	AV_PIX_FMT_GRAY16 = AvPixelFormat(C.AV_PIX_FMT_GRAY16)
	AV_PIX_FMT_YA16   = AvPixelFormat(C.AV_PIX_FMT_YA16)
	AV_PIX_FMT_RGB48  = AvPixelFormat(C.AV_PIX_FMT_RGB48)
	AV_PIX_FMT_RGB565 = AvPixelFormat(C.AV_PIX_FMT_RGB565)
	AV_PIX_FMT_RGB555 = AvPixelFormat(C.AV_PIX_FMT_RGB555)
	AV_PIX_FMT_RGB444 = AvPixelFormat(C.AV_PIX_FMT_RGB444)
	AV_PIX_FMT_RGBA64 = AvPixelFormat(C.AV_PIX_FMT_RGBA64)
	AV_PIX_FMT_BGR48  = AvPixelFormat(C.AV_PIX_FMT_BGR48)
	AV_PIX_FMT_BGR565 = AvPixelFormat(C.AV_PIX_FMT_BGR565)
	AV_PIX_FMT_BGR555 = AvPixelFormat(C.AV_PIX_FMT_BGR555)
	AV_PIX_FMT_BGR444 = AvPixelFormat(C.AV_PIX_FMT_BGR444)
	AV_PIX_FMT_BGRA64 = AvPixelFormat(C.AV_PIX_FMT_BGRA64)

	AV_PIX_FMT_YUV420P9  = AvPixelFormat(C.AV_PIX_FMT_YUV420P9)
	AV_PIX_FMT_YUV422P9  = AvPixelFormat(C.AV_PIX_FMT_YUV422P9)
	AV_PIX_FMT_YUV444P9  = AvPixelFormat(C.AV_PIX_FMT_YUV444P9)
	AV_PIX_FMT_YUV420P10 = AvPixelFormat(C.AV_PIX_FMT_YUV420P10)
	AV_PIX_FMT_YUV422P10 = AvPixelFormat(C.AV_PIX_FMT_YUV422P10)
	AV_PIX_FMT_YUV440P10 = AvPixelFormat(C.AV_PIX_FMT_YUV440P10)
	AV_PIX_FMT_YUV444P10 = AvPixelFormat(C.AV_PIX_FMT_YUV444P10)
	AV_PIX_FMT_YUV420P12 = AvPixelFormat(C.AV_PIX_FMT_YUV420P12)
	AV_PIX_FMT_YUV422P12 = AvPixelFormat(C.AV_PIX_FMT_YUV422P12)
	AV_PIX_FMT_YUV440P12 = AvPixelFormat(C.AV_PIX_FMT_YUV440P12)
	AV_PIX_FMT_YUV444P12 = AvPixelFormat(C.AV_PIX_FMT_YUV444P12)
	AV_PIX_FMT_YUV420P14 = AvPixelFormat(C.AV_PIX_FMT_YUV420P14)
	AV_PIX_FMT_YUV422P14 = AvPixelFormat(C.AV_PIX_FMT_YUV422P14)
	AV_PIX_FMT_YUV444P14 = AvPixelFormat(C.AV_PIX_FMT_YUV444P14)
	AV_PIX_FMT_YUV420P16 = AvPixelFormat(C.AV_PIX_FMT_YUV420P16)
	AV_PIX_FMT_YUV422P16 = AvPixelFormat(C.AV_PIX_FMT_YUV422P16)
	AV_PIX_FMT_YUV444P16 = AvPixelFormat(C.AV_PIX_FMT_YUV444P16)

	AV_PIX_FMT_GBRP9   = AvPixelFormat(C.AV_PIX_FMT_GBRP9)
	AV_PIX_FMT_GBRP10  = AvPixelFormat(C.AV_PIX_FMT_GBRP10)
	AV_PIX_FMT_GBRP12  = AvPixelFormat(C.AV_PIX_FMT_GBRP12)
	AV_PIX_FMT_GBRP14  = AvPixelFormat(C.AV_PIX_FMT_GBRP14)
	AV_PIX_FMT_GBRP16  = AvPixelFormat(C.AV_PIX_FMT_GBRP16)
	AV_PIX_FMT_GBRAP10 = AvPixelFormat(C.AV_PIX_FMT_GBRAP10)
	AV_PIX_FMT_GBRAP12 = AvPixelFormat(C.AV_PIX_FMT_GBRAP12)
	AV_PIX_FMT_GBRAP16 = AvPixelFormat(C.AV_PIX_FMT_GBRAP16)

	AV_PIX_FMT_BAYER_BGGR16 = AvPixelFormat(C.AV_PIX_FMT_BAYER_BGGR16)
	AV_PIX_FMT_BAYER_RGGB16 = AvPixelFormat(C.AV_PIX_FMT_BAYER_RGGB16)
	AV_PIX_FMT_BAYER_GBRG16 = AvPixelFormat(C.AV_PIX_FMT_BAYER_GBRG16)
	AV_PIX_FMT_BAYER_GRBG16 = AvPixelFormat(C.AV_PIX_FMT_BAYER_GRBG16)

	AV_PIX_FMT_GBRPF32  = AvPixelFormat(C.AV_PIX_FMT_GBRPF32)
	AV_PIX_FMT_GBRAPF32 = AvPixelFormat(C.AV_PIX_FMT_GBRAPF32)

	AV_PIX_FMT_GRAYF32 = AvPixelFormat(C.AV_PIX_FMT_GRAYF32)

	AV_PIX_FMT_YUVA420P9  = AvPixelFormat(C.AV_PIX_FMT_YUVA420P9)
	AV_PIX_FMT_YUVA422P9  = AvPixelFormat(C.AV_PIX_FMT_YUVA422P9)
	AV_PIX_FMT_YUVA444P9  = AvPixelFormat(C.AV_PIX_FMT_YUVA444P9)
	AV_PIX_FMT_YUVA420P10 = AvPixelFormat(C.AV_PIX_FMT_YUVA420P10)
	AV_PIX_FMT_YUVA422P10 = AvPixelFormat(C.AV_PIX_FMT_YUVA422P10)
	AV_PIX_FMT_YUVA444P10 = AvPixelFormat(C.AV_PIX_FMT_YUVA444P10)
	AV_PIX_FMT_YUVA422P12 = AvPixelFormat(C.AV_PIX_FMT_YUVA422P12)
	AV_PIX_FMT_YUVA444P12 = AvPixelFormat(C.AV_PIX_FMT_YUVA444P12)
	AV_PIX_FMT_YUVA420P16 = AvPixelFormat(C.AV_PIX_FMT_YUVA420P16)
	AV_PIX_FMT_YUVA422P16 = AvPixelFormat(C.AV_PIX_FMT_YUVA422P16)
	AV_PIX_FMT_YUVA444P16 = AvPixelFormat(C.AV_PIX_FMT_YUVA444P16)

	AV_PIX_FMT_XYZ12  = AvPixelFormat(C.AV_PIX_FMT_XYZ12)
	AV_PIX_FMT_NV20   = AvPixelFormat(C.AV_PIX_FMT_NV20)
	AV_PIX_FMT_AYUV64 = AvPixelFormat(C.AV_PIX_FMT_AYUV64)
	AV_PIX_FMT_P010   = AvPixelFormat(C.AV_PIX_FMT_P010)
	AV_PIX_FMT_P016   = AvPixelFormat(C.AV_PIX_FMT_P016)

	AV_PIX_FMT_Y210    = AvPixelFormat(C.AV_PIX_FMT_Y210)
	AV_PIX_FMT_X2RGB10 = AvPixelFormat(C.AV_PIX_FMT_X2RGB10)
)

// Chromaticity coordinates of the source primaries.
type AvColorPrimaries = C.enum_AVColorPrimaries

const (
	AVCOL_PRI_RESERVED0   = AvColorPrimaries(C.AVCOL_PRI_RESERVED0)
	AVCOL_PRI_BT709       = AvColorPrimaries(C.AVCOL_PRI_BT709)
	AVCOL_PRI_UNSPECIFIED = AvColorPrimaries(C.AVCOL_PRI_UNSPECIFIED)
	AVCOL_PRI_RESERVED    = AvColorPrimaries(C.AVCOL_PRI_RESERVED)
	AVCOL_PRI_BT470M      = AvColorPrimaries(C.AVCOL_PRI_BT470M)

	AVCOL_PRI_BT470BG      = AvColorPrimaries(C.AVCOL_PRI_BT470BG)
	AVCOL_PRI_SMPTE170M    = AvColorPrimaries(C.AVCOL_PRI_SMPTE170M)
	AVCOL_PRI_SMPTE240M    = AvColorPrimaries(C.AVCOL_PRI_SMPTE240M)
	AVCOL_PRI_FILM         = AvColorPrimaries(C.AVCOL_PRI_FILM)
	AVCOL_PRI_BT2020       = AvColorPrimaries(C.AVCOL_PRI_BT2020)
	AVCOL_PRI_SMPTE428     = AvColorPrimaries(C.AVCOL_PRI_SMPTE428)
	AVCOL_PRI_SMPTEST428_1 = AvColorPrimaries(C.AVCOL_PRI_SMPTEST428_1)
	AVCOL_PRI_SMPTE431     = AvColorPrimaries(C.AVCOL_PRI_SMPTE431)
	AVCOL_PRI_SMPTE432     = AvColorPrimaries(C.AVCOL_PRI_SMPTE432)
	AVCOL_PRI_EBU3213      = AvColorPrimaries(C.AVCOL_PRI_EBU3213)
	AVCOL_PRI_JEDEC_P22    = AvColorPrimaries(C.AVCOL_PRI_JEDEC_P22)
	AVCOL_PRI_NB           = AvColorPrimaries(C.AVCOL_PRI_NB)
)

// Color Transfer Characteristic.
type AvColorTransferCharacteristic = C.enum_AVColorTransferCharacteristic

const (
	AVCOL_TRC_RESERVED0    = AvColorTransferCharacteristic(C.AVCOL_TRC_RESERVED0)
	AVCOL_TRC_BT709        = AvColorTransferCharacteristic(C.AVCOL_TRC_BT709)
	AVCOL_TRC_UNSPECIFIED  = AvColorTransferCharacteristic(C.AVCOL_TRC_UNSPECIFIED)
	AVCOL_TRC_RESERVED     = AvColorTransferCharacteristic(C.AVCOL_TRC_RESERVED)
	AVCOL_TRC_GAMMA22      = AvColorTransferCharacteristic(C.AVCOL_TRC_GAMMA22)
	AVCOL_TRC_GAMMA28      = AvColorTransferCharacteristic(C.AVCOL_TRC_GAMMA28)
	AVCOL_TRC_SMPTE170M    = AvColorTransferCharacteristic(C.AVCOL_TRC_SMPTE170M)
	AVCOL_TRC_SMPTE240M    = AvColorTransferCharacteristic(C.AVCOL_TRC_SMPTE240M)
	AVCOL_TRC_LINEAR       = AvColorTransferCharacteristic(C.AVCOL_TRC_LINEAR)
	AVCOL_TRC_LOG          = AvColorTransferCharacteristic(C.AVCOL_TRC_LOG)
	AVCOL_TRC_LOG_SQRT     = AvColorTransferCharacteristic(C.AVCOL_TRC_LOG_SQRT)
	AVCOL_TRC_IEC61966_2_4 = AvColorTransferCharacteristic(C.AVCOL_TRC_IEC61966_2_4)
	AVCOL_TRC_BT1361_ECG   = AvColorTransferCharacteristic(C.AVCOL_TRC_BT1361_ECG)
	AVCOL_TRC_IEC61966_2_1 = AvColorTransferCharacteristic(C.AVCOL_TRC_IEC61966_2_1)
	AVCOL_TRC_BT2020_10    = AvColorTransferCharacteristic(C.AVCOL_TRC_BT2020_10)
	AVCOL_TRC_BT2020_12    = AvColorTransferCharacteristic(C.AVCOL_TRC_BT2020_12)
	AVCOL_TRC_SMPTE2084    = AvColorTransferCharacteristic(C.AVCOL_TRC_SMPTE2084)
	AVCOL_TRC_SMPTEST2084  = AvColorTransferCharacteristic(C.AVCOL_TRC_SMPTEST2084)
	AVCOL_TRC_SMPTE428     = AvColorTransferCharacteristic(C.AVCOL_TRC_SMPTE428)
	AVCOL_TRC_SMPTEST428_1 = AvColorTransferCharacteristic(C.AVCOL_TRC_SMPTEST428_1)
	AVCOL_TRC_ARIB_STD_B67 = AvColorTransferCharacteristic(C.AVCOL_TRC_ARIB_STD_B67)
	AVCOL_TRC_NB           = AvColorTransferCharacteristic(C.AVCOL_TRC_NB)
)

// AvColorSpace
type AvColorSpace = C.enum_AVColorSpace

const (
	AVCOL_SPC_RGB                = AvColorSpace(C.AVCOL_SPC_RGB)
	AVCOL_SPC_BT709              = AvColorSpace(C.AVCOL_SPC_BT709)
	AVCOL_SPC_UNSPECIFIED        = AvColorSpace(C.AVCOL_SPC_UNSPECIFIED)
	AVCOL_SPC_RESERVED           = AvColorSpace(C.AVCOL_SPC_RESERVED)
	AVCOL_SPC_FCC                = AvColorSpace(C.AVCOL_SPC_FCC)
	AVCOL_SPC_BT470BG            = AvColorSpace(C.AVCOL_SPC_BT470BG)
	AVCOL_SPC_SMPTE170M          = AvColorSpace(C.AVCOL_SPC_SMPTE170M)
	AVCOL_SPC_SMPTE240M          = AvColorSpace(C.AVCOL_SPC_SMPTE240M)
	AVCOL_SPC_YCGCO              = AvColorSpace(C.AVCOL_SPC_YCGCO)
	AVCOL_SPC_YCOCG              = AvColorSpace(C.AVCOL_SPC_YCOCG)
	AVCOL_SPC_BT2020_NCL         = AvColorSpace(C.AVCOL_SPC_BT2020_NCL)
	AVCOL_SPC_BT2020_CL          = AvColorSpace(C.AVCOL_SPC_BT2020_CL)
	AVCOL_SPC_SMPTE2085          = AvColorSpace(C.AVCOL_SPC_SMPTE2085)
	AVCOL_SPC_CHROMA_DERIVED_NCL = AvColorSpace(C.AVCOL_SPC_CHROMA_DERIVED_NCL)
	AVCOL_SPC_CHROMA_DERIVED_CL  = AvColorSpace(C.AVCOL_SPC_CHROMA_DERIVED_CL)
	AVCOL_SPC_ICTCP              = AvColorSpace(C.AVCOL_SPC_ICTCP)
	AVCOL_SPC_NB                 = AvColorSpace(C.AVCOL_SPC_NB)
)

// AvColorRange
type AvColorRange = C.enum_AVColorRange

const (
	AVCOL_RANGE_UNSPECIFIED = AvColorRange(C.AVCOL_RANGE_UNSPECIFIED)
	AVCOL_RANGE_MPEG        = AvColorRange(C.AVCOL_RANGE_MPEG)
	AVCOL_RANGE_JPEG        = AvColorRange(C.AVCOL_RANGE_JPEG)
	AVCOL_RANGE_NB          = AvColorRange(C.AVCOL_RANGE_NB)
)

// AvChromaLocation
type AvChromaLocation = C.enum_AVChromaLocation

const (
	AVCHROMA_LOC_UNSPECIFIED = AvChromaLocation(C.AVCHROMA_LOC_UNSPECIFIED)
	AVCHROMA_LOC_LEFT        = AvChromaLocation(C.AVCHROMA_LOC_LEFT)
	AVCHROMA_LOC_CENTER      = AvChromaLocation(C.AVCHROMA_LOC_CENTER)
	AVCHROMA_LOC_TOPLEFT     = AvChromaLocation(C.AVCHROMA_LOC_TOPLEFT)
	AVCHROMA_LOC_TOP         = AvChromaLocation(C.AVCHROMA_LOC_TOP)
	AVCHROMA_LOC_BOTTOMLEFT  = AvChromaLocation(C.AVCHROMA_LOC_BOTTOMLEFT)
	AVCHROMA_LOC_BOTTOM      = AvChromaLocation(C.AVCHROMA_LOC_BOTTOM)
	AVCHROMA_LOC_NB          = AvChromaLocation(C.AVCHROMA_LOC_NB)
)
