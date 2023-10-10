package ffmpeg

/*
#include <libpostproc/postprocess.h>
*/
import "C"
import "unsafe"

// PostprocVersion returns the LIBPOSTPROC_VERSION_INT constant.
func PostprocVersion() uint32 {
	return (uint32)(C.postproc_version())
}

// PostprocConfiguration returns the libpostproc build-time configuration.
func PostprocConfiguration() string {
	return C.GoString(C.postproc_configuration())
}

// PostprocLicense returns the libpostproc license.
func PostprocLicense() string {
	return C.GoString(C.postproc_license())
}

const (
	PP_QUALITY_MAX = C.PP_QUALITY_MAX
)

type PpContext C.pp_context

type PpMode C.pp_mode

// PpPostprocess
func PpPostprocess(src [3]*uint8, srcStride []int32,
	dst [3]*uint8, dstStride []int32,
	horizontalSize, verticalSize int32,
	QPStore *int8, QPStride int32,
	ppMode *PpMode, ppContext *PpContext, pictType int32) {
	C.pp_postprocess((**C.uint8_t)(unsafe.Pointer(&src[0])), (*C.int)(&srcStride[0]),
		(**C.uint8_t)(unsafe.Pointer(&dst[0])), (*C.int)(&dstStride[0]),
		(C.int)(horizontalSize), (C.int)(verticalSize),
		(*C.int8_t)(QPStore), (C.int)(QPStride),
		unsafe.Pointer(ppMode), unsafe.Pointer(ppContext), (C.int)(pictType))
}

// PpGetModeByNameAndQuality returns a pp_mode or NULL if an error occurred.
func PpGetModeByNameAndQuality(name string, quality int32) *PpMode {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (*PpMode)(C.pp_get_mode_by_name_and_quality((*C.char)(namePtr), (C.int)(quality)))
}

// PpFreeMode
func PpFreeMode(mode *PpMode) {
	C.pp_free_mode(unsafe.Pointer(mode))
}

// PpGetContext
func PpGetContext(width, height, flags int32) *PpContext {
	return (*PpContext)(C.pp_get_context((C.int)(width), (C.int)(height), (C.int)(flags)))
}

// PpFreeContext
func PpFreeContext(mode *PpContext) {
	C.pp_free_context(unsafe.Pointer(mode))
}

const (
	PP_CPU_CAPS_MMX     = C.PP_CPU_CAPS_MMX
	PP_CPU_CAPS_MMX2    = C.PP_CPU_CAPS_MMX2
	PP_CPU_CAPS_3DNOW   = C.PP_CPU_CAPS_3DNOW
	PP_CPU_CAPS_ALTIVEC = C.PP_CPU_CAPS_ALTIVEC
	PP_CPU_CAPS_AUTO    = C.PP_CPU_CAPS_AUTO
)

const (
	PP_FORMAT     = C.PP_FORMAT
	PP_FORMAT_420 = C.PP_FORMAT_420
	PP_FORMAT_422 = C.PP_FORMAT_422
	PP_FORMAT_411 = C.PP_FORMAT_411
	PP_FORMAT_444 = C.PP_FORMAT_444
	PP_FORMAT_440 = C.PP_FORMAT_440
)

const (
	PP_PICT_TYPE_QP2 = C.PP_PICT_TYPE_QP2
)
