package ffmpeg

/*
#include <libswscale/swscale.h>
*/
import "C"
import "unsafe"

// SwScaleVersion returns the LIBSWSCALE_VERSION_INT constant.
func SwScaleVersion() uint32 {
	return (uint32)(C.swscale_version())
}

// SwScaleConfiguration returns the libswscale build-time configuration.
func SwScaleConfiguration() string {
	return C.GoString(C.swscale_configuration())
}

// SwScaleLicense returns the libswscale license.
func SwScaleLicense() string {
	return C.GoString(C.swscale_license())
}

const (
	SWS_FAST_BILINEAR = C.SWS_FAST_BILINEAR
	SWS_BILINEAR      = C.SWS_BILINEAR
	SWS_BICUBIC       = C.SWS_BICUBIC
	SWS_X             = C.SWS_X
	SWS_POINT         = C.SWS_POINT
	SWS_AREA          = C.SWS_AREA
	SWS_BICUBLIN      = C.SWS_BICUBLIN
	SWS_GAUSS         = C.SWS_GAUSS
	SWS_SINC          = C.SWS_SINC
	SWS_LANCZOS       = C.SWS_LANCZOS
	SWS_SPLINE        = C.SWS_SPLINE
)

const (
	SWS_SRC_V_CHR_DROP_MASK  = C.SWS_SRC_V_CHR_DROP_MASK
	SWS_SRC_V_CHR_DROP_SHIFT = C.SWS_SRC_V_CHR_DROP_SHIFT
)

const SWS_PARAM_DEFAULT = C.SWS_PARAM_DEFAULT

const SWS_PRINT_INFO = C.SWS_PRINT_INFO

const (
	SWS_FULL_CHR_H_INT  = C.SWS_FULL_CHR_H_INT
	SWS_FULL_CHR_H_INP  = C.SWS_FULL_CHR_H_INP
	SWS_DIRECT_BGR      = C.SWS_DIRECT_BGR
	SWS_ACCURATE_RND    = C.SWS_ACCURATE_RND
	SWS_BITEXACT        = C.SWS_BITEXACT
	SWS_ERROR_DIFFUSION = C.SWS_ERROR_DIFFUSION
)

const SWS_MAX_REDUCE_CUTOFF = C.SWS_MAX_REDUCE_CUTOFF

const (
	SWS_CS_ITU709    = C.SWS_CS_ITU709
	SWS_CS_FCC       = C.SWS_CS_FCC
	SWS_CS_ITU601    = C.SWS_CS_ITU601
	SWS_CS_ITU624    = C.SWS_CS_ITU624
	SWS_CS_SMPTE170M = C.SWS_CS_SMPTE170M
	SWS_CS_SMPTE240M = C.SWS_CS_SMPTE240M
	SWS_CS_DEFAULT   = C.SWS_CS_DEFAULT
	SWS_CS_BT2020    = C.SWS_CS_BT2020
)

// SwsGetCoefficients returns a pointer to yuv<->rgb coefficients for the given colorspace
// suitable for SwsSetcolorspacedetails().
func SwsGetCoefficients(colorspace int32) *int32 {
	return (*int32)(C.sws_getCoefficients((C.int)(colorspace)))
}

// SwsVector
type SwsVector C.struct_SwsVector

// Custom: GetCoeff gets `SwsVector.coeff` value.
func (sv *SwsVector) GetCoeff() *float64 {
	return (*float64)(sv.coeff)
}

// Custom: SetCoeff sets `SwsVector.coeff` value.
func (sv *SwsVector) SetCoeff(v *float64) {
	sv.coeff = (*C.double)(v)
}

// Custom: GetCoeffAddr gets `SwsVector.coeff` address.
func (sv *SwsVector) GetCoeffAddr() **float64 {
	return (**float64)(unsafe.Pointer(&sv.coeff))
}

// Custom: GetLength gets `SwsVector.length` value.
func (sv *SwsVector) GetLength() int32 {
	return (int32)(sv.length)
}

// Custom: SetLength sets `SwsVector.length` value.
func (sv *SwsVector) SetLength(v int32) {
	sv.length = (C.int)(v)
}

// Custom: GetLengthAddr gets `SwsVector.length` address.
func (sv *SwsVector) GetLengthAddr() *int32 {
	return (*int32)(&sv.length)
}

// SwsFilter
type SwsFilter C.struct_SwsFilter

// Custom: GetLumH gets `SwsFilter.lumH` value.
func (sf *SwsFilter) GetLumH() *SwsVector {
	return (*SwsVector)(sf.lumH)
}

// Custom: SetLumH sets `SwsFilter.lumH` value.
func (sf *SwsFilter) SetLumH(v *SwsVector) {
	sf.lumH = (*C.struct_SwsVector)(v)
}

// Custom: GetLumHAddr gets `SwsFilter.lumH` address.
func (sf *SwsFilter) GetLumHAddr() **SwsVector {
	return (**SwsVector)(unsafe.Pointer(&sf.lumH))
}

// Custom: GetLumV gets `SwsFilter.lumV` value.
func (sf *SwsFilter) GetLumV() *SwsVector {
	return (*SwsVector)(sf.lumV)
}

// Custom: SetLumV sets `SwsFilter.lumV` value.
func (sf *SwsFilter) SetLumV(v *SwsVector) {
	sf.lumV = (*C.struct_SwsVector)(v)
}

// Custom: GetLumVAddr gets `SwsFilter.lumV` address.
func (sf *SwsFilter) GetLumVAddr() **SwsVector {
	return (**SwsVector)(unsafe.Pointer(&sf.lumV))
}

// Custom: GetChrH gets `SwsFilter.chrH` value.
func (sf *SwsFilter) GetChrH() *SwsVector {
	return (*SwsVector)(sf.chrH)
}

// Custom: SetChrH sets `SwsFilter.chrH` value.
func (sf *SwsFilter) SetChrH(v *SwsVector) {
	sf.chrH = (*C.struct_SwsVector)(v)
}

// Custom: GetChrHAddr gets `SwsFilter.chrH` address.
func (sf *SwsFilter) GetChrHAddr() **SwsVector {
	return (**SwsVector)(unsafe.Pointer(&sf.chrH))
}

// Custom: GetChrV gets `SwsFilter.chrV` value.
func (sf *SwsFilter) GetChrV() *SwsVector {
	return (*SwsVector)(sf.chrV)
}

// Custom: SetChrV sets `SwsFilter.chrV` value.
func (sf *SwsFilter) SetChrV(v *SwsVector) {
	sf.chrV = (*C.struct_SwsVector)(v)
}

// Custom: GetChrVAddr gets `SwsFilter.chrV` address.
func (sf *SwsFilter) GetChrVAddr() **SwsVector {
	return (**SwsVector)(unsafe.Pointer(&sf.chrV))
}

// SwsContext
type SwsContext C.struct_SwsContext

// SwsIsSupportedInput returns a positive value if pix_fmt is a supported input format.
func SwsIsSupportedInput(pixFmt AVPixelFormat) int32 {
	return (int32)(C.sws_isSupportedInput((C.enum_AVPixelFormat)(pixFmt)))
}

// SwsIsSupportedOutput returns a positive value if pix_fmt is a supported output format.
func SwsIsSupportedOutput(pixFmt AVPixelFormat) int32 {
	return (int32)(C.sws_isSupportedOutput((C.enum_AVPixelFormat)(pixFmt)))
}

// SwsIsSupportedEndiannessConversion returns a positive value
// if pix_fmt is a supported endianness conversion format.
func SwsIsSupportedEndiannessConversion(pixFmt AVPixelFormat) int32 {
	return (int32)(C.sws_isSupportedEndiannessConversion((C.enum_AVPixelFormat)(pixFmt)))
}

// SwsAllocContext allocates an empty SwsContext.
func SwsAllocContext() *SwsContext {
	return (*SwsContext)(C.sws_alloc_context())
}

// SwsInitContext initializes the swscaler context sws_context.
func SwsInitContext(sctx *SwsContext, srcFilter, dstFilter *SwsFilter) int32 {
	return (int32)(C.sws_init_context((*C.struct_SwsContext)(sctx),
		(*C.struct_SwsFilter)(srcFilter), (*C.struct_SwsFilter)(dstFilter)))
}

// SwsFreecontext frees the swscaler context swsContext.
func SwsFreeContext(sctx *SwsContext) {
	C.sws_freeContext((*C.struct_SwsContext)(sctx))
}

// SwsGetcontext allocates and returns an SwsContext.
func SwsGetContext(srcW, srcH int32, srcFormat AVPixelFormat,
	dstW, dstH int32, dstFormat AVPixelFormat,
	flags int32, srcFilter, dstFilter *SwsFilter, param *float64) *SwsContext {
	return (*SwsContext)(C.sws_getContext((C.int)(srcW), (C.int)(srcH), (C.enum_AVPixelFormat)(srcFormat),
		(C.int)(dstW), (C.int)(dstH), (C.enum_AVPixelFormat)(dstFormat),
		(C.int)(flags), (*C.struct_SwsFilter)(srcFilter), (*C.struct_SwsFilter)(dstFilter),
		(*C.double)(param)))
}

// SwsScale scales the image slice in srcSlice and put the resulting scaled
// slice in the image in dst. A slice is a sequence of consecutive rows in an image.
func SwsScale(sctx *SwsContext, srcSlice []*uint8, srcStride []int32,
	srcSliceY, srcSliceH int32,
	dst []*uint8, dstStride []int32) int32 {
	return (int32)(C.sws_scale((*C.struct_SwsContext)(sctx),
		(**C.uint8_t)(unsafe.Pointer(&srcSlice[0])), (*C.int)(unsafe.Pointer(&srcStride[0])),
		(C.int)(srcSliceY), (C.int)(srcSliceH),
		(**C.uint8_t)(unsafe.Pointer(&dst[0])), (*C.int)(unsafe.Pointer(&dstStride[0]))))
}

// SwsSetColorspaceDetails
func SwsSetColorSpaceDetails(sctx *SwsContext, invTable []int32, srcRange int32,
	table []int32, dstRange int32, brightness, contrast, saturation int32) int32 {
	if len(invTable) < 4 {
		panic("invTable len < 4")
	}
	if len(table) < 4 {
		panic("table len < 4")
	}
	return (int32)(C.sws_setColorspaceDetails((*C.struct_SwsContext)(sctx),
		(*C.int)(unsafe.Pointer(&invTable[0])), (C.int)(srcRange),
		(*C.int)(unsafe.Pointer(&table[0])), (C.int)(dstRange),
		(C.int)(brightness), (C.int)(contrast), (C.int)(saturation)))
}

// SwsGetColorSpaceDetails
func SwsGetColorSpaceDetails(sctx *SwsContext, invTable []int32, srcRange *int32,
	table []int32, dstRange *int32, brightness, contrast, saturation *int32) int32 {
	if len(invTable) < 4 {
		panic("invTable len < 4")
	}
	if len(table) < 4 {
		panic("table len = 4")
	}
	invTablePtr := unsafe.Pointer(&invTable[0])
	tablePtr := unsafe.Pointer(&table[0])
	return (int32)(C.sws_getColorspaceDetails((*C.struct_SwsContext)(sctx),
		(**C.int)(unsafe.Pointer(&invTablePtr)), (*C.int)(srcRange),
		(**C.int)(unsafe.Pointer(&tablePtr)), (*C.int)(dstRange),
		(*C.int)(brightness), (*C.int)(contrast), (*C.int)(saturation)))
}

// SwsAllocVec allocates and returns an uninitialized vector with length coefficients.
func SwsAllocVec(length int32) *SwsVector {
	return (*SwsVector)(C.sws_allocVec((C.int)(length)))
}

// SwsGetGaussianVec Return a normalized Gaussian curve used to filter stuff.
func SwsGetGaussianVec(variance, quality float64) *SwsVector {
	return (*SwsVector)(C.sws_getGaussianVec((C.double)(variance), (C.double)(quality)))
}

// SwsScaleVec scales all the coefficients of a by the scalar value.
func SwsScaleVec(a *SwsVector, scalar float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(a), (C.double)(scalar))
}

// SwsNormalizeVec scales all the coefficients of a so that their sum equals height.
func SwsNormalizeVec(a *SwsVector, height float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(a), (C.double)(height))
}

// Deprecated: No use
func SwsGetConstVec(c float64, length int32) *SwsVector {
	return (*SwsVector)(C.sws_getConstVec((C.double)(c), (C.int)(length)))
}

// Deprecated: No use
func SwsGetIdentityVec() *SwsVector {
	return (*SwsVector)(C.sws_getIdentityVec())
}

// Deprecated: No use
func SwsConvVec(a, b *SwsVector) {
	C.sws_convVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

// Deprecated: No use
func SwsAddVec(a, b *SwsVector) {
	C.sws_addVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

// Deprecated: No use
func SwsSubVec(a, b *SwsVector) {
	C.sws_subVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

// Deprecated: No use
func SwsShiftVec(a *SwsVector, shift int32) {
	C.sws_shiftVec((*C.struct_SwsVector)(a), (C.int)(shift))
}

// Deprecated: No use
func SwsCloneVec(a *SwsVector) *SwsVector {
	return (*SwsVector)(C.sws_cloneVec((*C.struct_SwsVector)(a)))
}

// Deprecated: No use
func SwsPrintVec2(a *SwsVector, logCtx *AVClass, logLevel int32) {
	C.sws_printVec2((*C.struct_SwsVector)(a),
		(*C.struct_AVClass)(logCtx), (C.int)(logLevel))
}

// SwsFreeVec
func SwsFreeVec(a *SwsVector) {
	C.sws_freeVec((*C.struct_SwsVector)(a))
}

// SwsGetDefaultFilter
func SwsGetDefaultFilter(lumaGBlur, chromaGBlur float32,
	lumaSharpen, chromaSharpen float32,
	chromaHShift, chromaVShift float32, verbose int32) *SwsFilter {
	return (*SwsFilter)(C.sws_getDefaultFilter((C.float)(lumaGBlur), (C.float)(chromaGBlur),
		(C.float)(lumaSharpen), (C.float)(chromaSharpen),
		(C.float)(chromaHShift), (C.float)(chromaVShift), (C.int)(verbose)))
}

// SwsFreeFilter
func SwsFreeFilter(filter *SwsFilter) {
	C.sws_freeFilter((*C.struct_SwsFilter)(filter))
}

// SwsGetCachedContext check if context can be reused, otherwise reallocate a new one.
func SwsGetCachedContext(context *SwsContext,
	srcW, srcH int32, srcFormat AVPixelFormat,
	dstW, dstH int32, dstFormat AVPixelFormat,
	flags int32, srcFilter, dstFilter *SwsFilter, param *float64) *SwsContext {
	return (*SwsContext)(C.sws_getCachedContext((*C.struct_SwsContext)(context),
		(C.int)(srcW), (C.int)(srcH), (C.enum_AVPixelFormat)(srcFormat),
		(C.int)(dstW), (C.int)(dstH), (C.enum_AVPixelFormat)(dstFormat),
		(C.int)(flags), (*C.struct_SwsFilter)(srcFilter), (*C.struct_SwsFilter)(dstFilter),
		(*C.double)(param)))
}

// SwsConvertPalette8ToPacked32 converts an 8-bit paletted frame into a frame with a color depth of 32 bits.
func SwsConvertPalette8ToPacked32(src, dst *uint8, numPixels int32, palette *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(src), (*C.uint8_t)(dst),
		(C.int)(numPixels), (*C.uint8_t)(palette))
}

// SwsConvertPalette8ToPacked24 converts an 8-bit paletted frame into a frame with a color depth of 24 bits.
func SwsConvertPalette8ToPacked24(src, dst *uint8, numPixels int32, palette *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(src), (*C.uint8_t)(dst),
		(C.int)(numPixels), (*C.uint8_t)(palette))
}

// SwsGetClass gets the AVClass for swsContext.
func SwsGetClass() *AVClass {
	return (*AVClass)(C.sws_get_class())
}
