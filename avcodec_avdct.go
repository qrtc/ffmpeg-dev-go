package ffmpeg

/*
#include <libavcodec/avdct.h>

typedef void (*avdct_idct_func)(int16_t *block);
typedef void (*avdct_fdct_func)(int16_t *block);
typedef void (*avdct_get_pixels_func)(int16_t *block, const uint8_t *pixels, ptrdiff_t line_size);
typedef void (*avdct_get_pixels_unaligned_func)(int16_t *block, const uint8_t *pixels, ptrdiff_t line_size);
*/
import "C"
import "unsafe"

// AVDCT
type AVDCT C.struct_AVDCT

// typedef void (*avdct_idct_func)(int16_t *block);
type AVDCTIdctFunc = C.avdct_idct_func

// typedef void (*avdct_fdct_func)(int16_t *block);
type AVDCTFdctFunc = C.avdct_fdct_func

// typedef void (*avdct_get_pixels_func)(int16_t *block, const uint8_t *pixels, ptrdiff_t line_size);
type AVDCTGetPixelsFunc = C.avdct_get_pixels_func

// typedef void (*avdct_get_pixels_unaligned_func)(int16_t *block, const uint8_t *pixels, ptrdiff_t line_size);
type AVDCTGetPixelsUnalignedFunc = C.avdct_get_pixels_unaligned_func

// Custom: GetAvClass gets `AVDCT.av_class` value.
func (dct *AVDCT) GetAvClass() *AVClass {
	return (*AVClass)(dct.av_class)
}

// Custom: SetAvClass sets `AVDCT.av_class` value.
func (dct *AVDCT) SetAvClass(v *AVClass) {
	dct.av_class = (*C.struct_AVClass)(v)
}

// Custom: GetAvClassAddr gets `AVDCT.av_class` address.
func (dct *AVDCT) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&dct.av_class))
}

// Custom: GetIdct gets `AVDCT.idct` value.
func (dct *AVDCT) GetIdct() AVDCTIdctFunc {
	return (AVDCTIdctFunc)(dct.idct)
}

// Custom: SetIdct sets `AVDCT.idct` value.
func (dct *AVDCT) SetIdct(v AVDCTIdctFunc) {
	dct.idct = (C.avdct_idct_func)(v)
}

// Custom: GetIdctAddr gets `AVDCT.idct` address.
func (dct *AVDCT) GetIdctAddr() *AVDCTIdctFunc {
	return (*AVDCTIdctFunc)(&dct.idct)
}

// Custom: GetIdctPermutation gets `AVDCT.idct_permutation` value.
func (dct *AVDCT) GetIdctPermutation() []uint8 {
	return unsafe.Slice((*uint8)(&dct.idct_permutation[0]), 64)
}

// Custom: SetIdctPermutation sets `AVDCT.idct_permutation` value.
func (dct *AVDCT) SetIdctPermutation(v []uint8) {
	for i := 0; i < FFMIN(len(v), 64); i++ {
		dct.idct_permutation[i] = (C.uint8_t)(v[i])
	}
}

// Custom: GetIdctPermutationAddr gets `AVDCT.idct_permutation` address.
func (dct *AVDCT) GetIdctPermutationAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&dct.idct_permutation))
}

// Custom: GetFdct gets `AVDCT.fdct` value.
func (dct *AVDCT) GetFdct() AVDCTFdctFunc {
	return (AVDCTFdctFunc)(dct.fdct)
}

// Custom: SetFdct sets `AVDCT.fdct` value.
func (dct *AVDCT) SetFdct(v AVDCTFdctFunc) {
	dct.fdct = (C.avdct_fdct_func)(v)
}

// Custom: GetFdctAddr gets `AVDCT.fdct` address.
func (dct *AVDCT) GetFdctAddr() *AVDCTFdctFunc {
	return (*AVDCTFdctFunc)(&dct.fdct)
}

// Custom: GetDctAlgo gets `AVDCT.dct_algo` value.
func (dct *AVDCT) GetDctAlgo() int32 {
	return (int32)(dct.dct_algo)
}

// Custom: SetDctAlgo sets `AVDCT.dct_algo` value.
func (dct *AVDCT) SetDctAlgo(v int32) {
	dct.dct_algo = (C.int)(v)
}

// Custom: GetDctAlgoAddr gets `AVDCT.dct_algo` address.
func (dct *AVDCT) GetDctAlgoAddr() *int32 {
	return (*int32)(&dct.dct_algo)
}

// Custom: GetIdctAlgo gets `AVDCT.idct_algo` value.
func (dct *AVDCT) GetIdctAlgo() int32 {
	return (int32)(dct.idct_algo)
}

// Custom: SetIdctAlgo sets `AVDCT.idct_algo` value.
func (dct *AVDCT) SetIdctAlgo(v int32) {
	dct.idct_algo = (C.int)(v)
}

// Custom: GetIdctAlgoAddr gets `AVDCT.idct_algo` address.
func (dct *AVDCT) GetIdctAlgoAddr() *int32 {
	return (*int32)(&dct.idct_algo)
}

// Custom: GetGetPixels gets `AVDCT.get_pixels` value.
func (dct *AVDCT) GetGetPixels() AVDCTGetPixelsFunc {
	return (AVDCTGetPixelsFunc)(dct.get_pixels)
}

// Custom: SetGetPixels sets `AVDCT.get_pixels` value.
func (dct *AVDCT) SetGetPixels(v AVDCTGetPixelsFunc) {
	dct.get_pixels = (C.avdct_get_pixels_func)(v)
}

// Custom: GetGetPixelsAddr gets `AVDCT.get_pixels` address.
func (dct *AVDCT) GetGetPixelsAddr() *AVDCTGetPixelsFunc {
	return (*AVDCTGetPixelsFunc)(&dct.get_pixels)
}

// Custom: GetBitsPerSample gets `AVDCT.bits_per_sample` value.
func (dct *AVDCT) GetBitsPerSample() int32 {
	return (int32)(dct.bits_per_sample)
}

// Custom: SetBitsPerSample sets `AVDCT.bits_per_sample` value.
func (dct *AVDCT) SetBitsPerSample(v int32) {
	dct.bits_per_sample = (C.int)(v)
}

// Custom: GetBitsPerSampleAddr gets `AVDCT.bits_per_sample` address.
func (dct *AVDCT) GetBitsPerSampleAddr() *int32 {
	return (*int32)(&dct.bits_per_sample)
}

// Custom: GetGetPixelsUnaligned gets `AVDCT.get_pixels_unaligned` value.
func (dct *AVDCT) GetGetPixelsUnaligned() AVDCTGetPixelsUnalignedFunc {
	return (AVDCTGetPixelsUnalignedFunc)(dct.get_pixels_unaligned)
}

// Custom: SetGetPixelsUnaligned sets `AVDCT.get_pixels_unaligned` value.
func (dct *AVDCT) SetGetPixelsUnaligned(v AVDCTGetPixelsUnalignedFunc) {
	dct.get_pixels_unaligned = (C.avdct_get_pixels_unaligned_func)(v)
}

// Custom: GetGetPixelsUnalignedAddr gets `AVDCT.get_pixels_unaligned` address.
func (dct *AVDCT) GetGetPixelsUnalignedAddr() *AVDCTGetPixelsUnalignedFunc {
	return (*AVDCTGetPixelsUnalignedFunc)(&dct.get_pixels_unaligned)
}

// AvCodecDctAlloc allocates a AVDCT context.
func AvCodecDctAlloc() *AVDCT {
	return (*AVDCT)(C.avcodec_dct_alloc())
}

// AvCodecDctInit
func AvCodecDctInit(dct *AVDCT) int32 {
	return (int32)(C.avcodec_dct_init((*C.struct_AVDCT)(dct)))
}

// AvCodecDctGetClass
func AvCodecDctGetClass() *AVClass {
	return (*AVClass)(C.avcodec_dct_get_class())
}
