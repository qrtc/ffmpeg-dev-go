package ffmpeg

/*
#include <libavutil/film_grain_params.h>

AVFilmGrainAOMParams get_av_film_grain_params_codec_aom(AVFilmGrainParams *v) {
	return v->codec.aom;
}

void set_av_film_grain_params_codec_aom(AVFilmGrainParams *v, AVFilmGrainAOMParams p) {
	v->codec.aom = p;
}

AVFilmGrainAOMParams* get_av_film_grain_params_codec_aom_addr(AVFilmGrainParams *v) {
	return &v->codec.aom;
}
*/
import "C"
import "unsafe"

// AVFilmGrainParamsType
type AVFilmGrainParamsType = C.enum_AVFilmGrainParamsType

const (
	AV_FILM_GRAIN_PARAMS_NONE = AVFilmGrainParamsType(C.AV_FILM_GRAIN_PARAMS_NONE)
	AV_FILM_GRAIN_PARAMS_AV1  = AVFilmGrainParamsType(C.AV_FILM_GRAIN_PARAMS_AV1)
)

// AVFilmGrainAOMParams
type AVFilmGrainAOMParams C.struct_AVFilmGrainAOMParams

// Custom: GetNumYPoints gets `AVFilmGrainAOMParams.num_y_points` value.
func (aomp *AVFilmGrainAOMParams) GetNumYPoints() int32 {
	return (int32)(aomp.num_y_points)
}

// Custom: SetNumYPoints sets `AVFilmGrainAOMParams.num_y_points` value.
func (aomp *AVFilmGrainAOMParams) SetNumYPoints(v int32) {
	aomp.num_y_points = (C.int)(v)
}

// Custom: GetNumYPointsAddr gets `AVFilmGrainAOMParams.num_y_points` address.
func (aomp *AVFilmGrainAOMParams) GetNumYPointsAddr() *int32 {
	return (*int32)(&aomp.num_y_points)
}

// Custom: GetYPoints gets `AVFilmGrainAOMParams.y_points` value.
func (aomp *AVFilmGrainAOMParams) GetYPoints() (v [][]uint8) {
	for i := 0; i < 14; i++ {
		v = append(v, unsafe.Slice((*uint8)(&aomp.y_points[i][0]), 2))
	}
	return v
}

// Custom: SetYPoints sets `AVFilmGrainAOMParams.y_points` value.
func (aomp *AVFilmGrainAOMParams) SetYPoints(v [][]uint8) {
	for i := 0; i < FFMIN(len(v), 14); i++ {
		for j := 0; j < FFMIN(len(v[i]), 2); j++ {
			aomp.y_points[i][j] = (C.uint8_t)(v[i][j])
		}
	}
}

// Custom: GetYPointsAddr gets `AVFilmGrainAOMParams.y_points` address.
func (aomp *AVFilmGrainAOMParams) GetYPointsAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&aomp.y_points))
}

// Custom: GetChromaScalingFromLuma gets `AVFilmGrainAOMParams.chroma_scaling_from_luma` value.
func (aomp *AVFilmGrainAOMParams) GetChromaScalingFromLuma() int32 {
	return (int32)(aomp.chroma_scaling_from_luma)
}

// Custom: SetChromaScalingFromLuma sets `AVFilmGrainAOMParams.chroma_scaling_from_luma` value.
func (aomp *AVFilmGrainAOMParams) SetChromaScalingFromLuma(v int32) {
	aomp.chroma_scaling_from_luma = (C.int)(v)
}

// Custom: GetChromaScalingFromLumaAddr gets `AVFilmGrainAOMParams.chroma_scaling_from_luma` address.
func (aomp *AVFilmGrainAOMParams) GetChromaScalingFromLumaAddr() *int32 {
	return (*int32)(&aomp.chroma_scaling_from_luma)
}

// Custom: GetNumUvPoints gets `AVFilmGrainAOMParams.num_uv_points` value.
func (aomp *AVFilmGrainAOMParams) GetNumUvPoints() []int32 {
	return unsafe.Slice((*int32)(&aomp.num_uv_points[0]), 2)
}

// Custom: SetNumUvPoints sets `AVFilmGrainAOMParams.num_uv_points` value.
func (aomp *AVFilmGrainAOMParams) SetNumUvPoints(v []int32) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		aomp.num_uv_points[i] = (C.int)(v[i])
	}
}

// Custom: GetNumUvPointsAddr gets `AVFilmGrainAOMParams.num_uv_points` address.
func (aomp *AVFilmGrainAOMParams) GetNumUvPointsAddr() **int32 {
	return (**int32)(unsafe.Pointer(&aomp.num_uv_points))
}

// Custom: GetUvPoints gets `AVFilmGrainAOMParams.uv_points` value.
func (aomp *AVFilmGrainAOMParams) GetUvPoints() (v [][][]uint8) {
	for i := 0; i < 2; i++ {
		tmp := [][]uint8{}
		for j := 0; j < 10; j++ {
			tmp = append(tmp, unsafe.Slice((*uint8)(&aomp.uv_points[i][j][0]), 2))
		}
		v = append(v, tmp)
	}
	return v
}

// Custom: SetUvPoints sets `AVFilmGrainAOMParams.uv_points` value.
func (aomp *AVFilmGrainAOMParams) SetUvPoints(v [][][]uint8) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		for j := 0; j < FFMIN(len(v[i]), 10); j++ {
			for k := 0; k < FFMIN(len(v[i][j]), 2); k++ {
				aomp.uv_points[i][j][k] = (C.uint8_t)(v[i][j][k])
			}
		}
	}
}

// Custom: GetUvPointsAddr gets `AVFilmGrainAOMParams.uv_points` address.
func (aomp *AVFilmGrainAOMParams) GetUvPointsAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&aomp.uv_points))
}

// Custom: GetScalingShift gets `AVFilmGrainAOMParams.scaling_shift` value.
func (aomp *AVFilmGrainAOMParams) GetScalingShift() int32 {
	return (int32)(aomp.scaling_shift)
}

// Custom: SetScalingShift sets `AVFilmGrainAOMParams.scaling_shift` value.
func (aomp *AVFilmGrainAOMParams) SetScalingShift(v int32) {
	aomp.scaling_shift = (C.int)(v)
}

// Custom: GetScalingShiftAddr gets `AVFilmGrainAOMParams.scaling_shift` address.
func (aomp *AVFilmGrainAOMParams) GetScalingShiftAddr() *int32 {
	return (*int32)(&aomp.scaling_shift)
}

// Custom: GetArCoeffLag gets `AVFilmGrainAOMParams.ar_coeff_lag` value.
func (aomp *AVFilmGrainAOMParams) GetArCoeffLag() int32 {
	return (int32)(aomp.ar_coeff_lag)
}

// Custom: SetArCoeffLag sets `AVFilmGrainAOMParams.ar_coeff_lag` value.
func (aomp *AVFilmGrainAOMParams) SetArCoeffLag(v int32) {
	aomp.ar_coeff_lag = (C.int)(v)
}

// Custom: GetArCoeffLagAddr gets `AVFilmGrainAOMParams.ar_coeff_lag` address.
func (aomp *AVFilmGrainAOMParams) GetArCoeffLagAddr() *int32 {
	return (*int32)(&aomp.ar_coeff_lag)
}

// Custom: GetArCoeffsY gets `AVFilmGrainAOMParams.ar_coeffs_y` value.
func (aomp *AVFilmGrainAOMParams) GetArCoeffsY() []int8 {
	return unsafe.Slice((*int8)(&aomp.ar_coeffs_y[0]), 24)
}

// Custom: SetArCoeffsY sets `AVFilmGrainAOMParams.ar_coeffs_y` value.
func (aomp *AVFilmGrainAOMParams) SetArCoeffsY(v []int8) {
	for i := 0; i < FFMIN(len(v), 24); i++ {
		aomp.ar_coeffs_y[i] = (C.int8_t)(v[i])
	}
}

// Custom: GetArCoeffsYAddr gets `AVFilmGrainAOMParams.ar_coeffs_y` address.
func (aomp *AVFilmGrainAOMParams) GetArCoeffsYAddr() **int8 {
	return (**int8)(unsafe.Pointer(&aomp.ar_coeffs_y))
}

// Custom: GetArCoeffsUv gets `AVFilmGrainAOMParams.ar_coeffs_uv` value.
func (aomp *AVFilmGrainAOMParams) GetArCoeffsUv() (v [][]int8) {
	for i := 0; i < 2; i++ {
		v = append(v, unsafe.Slice((*int8)(&aomp.ar_coeffs_uv[i][0]), 25))
	}
	return v
}

// Custom: SetArCoeffsUv sets `AVFilmGrainAOMParams.ar_coeffs_uv` value.
func (aomp *AVFilmGrainAOMParams) SetArCoeffsUv(v [][]int8) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		for j := 0; j < FFMIN(len(v[i]), 25); j++ {
			aomp.ar_coeffs_uv[i][j] = (C.int8_t)(v[i][j])
		}
	}
}

// Custom: GetArCoeffsUvAddr gets `AVFilmGrainAOMParams.ar_coeffs_uv` address.
func (aomp *AVFilmGrainAOMParams) GetArCoeffsUvAddr() **int8 {
	return (**int8)(unsafe.Pointer(&aomp.ar_coeffs_uv))
}

// Custom: GetArCoeffShift gets `AVFilmGrainAOMParams.ar_coeff_shift` value.
func (aomp *AVFilmGrainAOMParams) GetArCoeffShift() int32 {
	return (int32)(aomp.ar_coeff_shift)
}

// Custom: SetArCoeffShift sets `AVFilmGrainAOMParams.ar_coeff_shift` value.
func (aomp *AVFilmGrainAOMParams) SetArCoeffShift(v int32) {
	aomp.ar_coeff_shift = (C.int)(v)
}

// Custom: GetArCoeffShiftAddr gets `AVFilmGrainAOMParams.ar_coeff_shift` address.
func (aomp *AVFilmGrainAOMParams) GetArCoeffShiftAddr() *int32 {
	return (*int32)(&aomp.ar_coeff_shift)
}

// Custom: GetGrainScaleShift gets `AVFilmGrainAOMParams.grain_scale_shift` value.
func (aomp *AVFilmGrainAOMParams) GetGrainScaleShift() int32 {
	return (int32)(aomp.grain_scale_shift)
}

// Custom: SetGrainScaleShift sets `AVFilmGrainAOMParams.grain_scale_shift` value.
func (aomp *AVFilmGrainAOMParams) SetGrainScaleShift(v int32) {
	aomp.grain_scale_shift = (C.int)(v)
}

// Custom: GetGrainScaleShiftAddr gets `AVFilmGrainAOMParams.grain_scale_shift` address.
func (aomp *AVFilmGrainAOMParams) GetGrainScaleShiftAddr() *int32 {
	return (*int32)(&aomp.grain_scale_shift)
}

// Custom: GetUvMult gets `AVFilmGrainAOMParams.uv_mult` value.
func (aomp *AVFilmGrainAOMParams) GetUvMult() []int32 {
	return unsafe.Slice((*int32)(&aomp.uv_mult[0]), 2)
}

// Custom: SetUvMult sets `AVFilmGrainAOMParams.uv_mult` value.
func (aomp *AVFilmGrainAOMParams) SetUvMult(v []int32) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		aomp.uv_mult[i] = (C.int)(v[i])
	}
}

// Custom: GetUvMultAddr gets `AVFilmGrainAOMParams.uv_mult` address.
func (aomp *AVFilmGrainAOMParams) GetUvMultAddr() **int32 {
	return (**int32)(unsafe.Pointer(&aomp.uv_mult))
}

// Custom: GetUvMultLuma gets `AVFilmGrainAOMParams.uv_mult_luma` value.
func (aomp *AVFilmGrainAOMParams) GetUvMultLuma() []int32 {
	return unsafe.Slice((*int32)(&aomp.uv_mult_luma[0]), 2)
}

// Custom: SetUvMultLuma sets `AVFilmGrainAOMParams.uv_mult_luma` value.
func (aomp *AVFilmGrainAOMParams) SetUvMultLuma(v []int32) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		aomp.uv_mult_luma[i] = (C.int)(v[i])
	}
}

// Custom: GetUvMultLumaAddr gets `AVFilmGrainAOMParams.uv_mult_luma` address.
func (aomp *AVFilmGrainAOMParams) GetUvMultLumaAddr() **int32 {
	return (**int32)(unsafe.Pointer(&aomp.uv_mult_luma))
}

// Custom: GetUvOffset gets `AVFilmGrainAOMParams.uv_offset` value.
func (aomp *AVFilmGrainAOMParams) GetUvOffset() []int32 {
	return unsafe.Slice((*int32)(&aomp.uv_offset[0]), 2)
}

// Custom: SetUvOffset sets `AVFilmGrainAOMParams.uv_offset` value.
func (aomp *AVFilmGrainAOMParams) SetUvOffset(v []int32) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		aomp.uv_offset[i] = (C.int)(v[i])
	}
}

// Custom: GetUvOffsetAddr gets `AVFilmGrainAOMParams.uv_offset` address.
func (aomp *AVFilmGrainAOMParams) GetUvOffsetAddr() **int32 {
	return (**int32)(unsafe.Pointer(&aomp.uv_offset))
}

// Custom: GetOverlapFlag gets `AVFilmGrainAOMParams.overlap_flag` value.
func (aomp *AVFilmGrainAOMParams) GetOverlapFlag() int32 {
	return (int32)(aomp.overlap_flag)
}

// Custom: SetOverlapFlag sets `AVFilmGrainAOMParams.overlap_flag` value.
func (aomp *AVFilmGrainAOMParams) SetOverlapFlag(v int32) {
	aomp.overlap_flag = (C.int)(v)
}

// Custom: GetOverlapFlagAddr gets `AVFilmGrainAOMParams.overlap_flag` address.
func (aomp *AVFilmGrainAOMParams) GetOverlapFlagAddr() *int32 {
	return (*int32)(&aomp.overlap_flag)
}

// Custom: GetLimitOutputRange gets `AVFilmGrainAOMParams.limit_outputrange` value.
func (aomp *AVFilmGrainAOMParams) GetLimitOutputRange() int32 {
	return (int32)(aomp.limit_output_range)
}

// Custom: SetLimitOutputRange sets `AVFilmGrainAOMParams.limit_outputrange` value.
func (aomp *AVFilmGrainAOMParams) SetLimitOutputRange(v int32) {
	aomp.limit_output_range = (C.int)(v)
}

// Custom: GetLimitOutputRangeAddr gets `AVFilmGrainAOMParams.limit_outputrange` address.
func (aomp *AVFilmGrainAOMParams) GetLimitOutputRangeAddr() *int32 {
	return (*int32)(&aomp.limit_output_range)
}

// AVFilmGrainParams
type AVFilmGrainParams C.struct_AVFilmGrainParams

// Custom: GetType gets `AVFilmGrainParams.type` value.
func (fgp *AVFilmGrainParams) GetType() AVFilmGrainParamsType {
	return (AVFilmGrainParamsType)(fgp._type)
}

// Custom: SetType sets `AVFilmGrainParams.type` value.
func (fgp *AVFilmGrainParams) SetType(v AVFilmGrainParamsType) {
	fgp._type = (C.enum_AVFilmGrainParamsType)(v)
}

// Custom: GetTypeAddr gets `AVFilmGrainParams.type` address.
func (fgp *AVFilmGrainParams) GetTypeAddr() *AVFilmGrainParamsType {
	return (*AVFilmGrainParamsType)(&fgp._type)
}

// Custom: GetSeed gets `AVFilmGrainParams.seed` value.
func (fgp *AVFilmGrainParams) GetSeed() uint64 {
	return (uint64)(fgp.seed)
}

// Custom: SetSeed sets `AVFilmGrainParams.seed` value.
func (fgp *AVFilmGrainParams) SetSeed(v uint64) {
	fgp.seed = (C.uint64_t)(v)
}

// Custom: GetSeedAddr gets `AVFilmGrainParams.seed` address.
func (fgp *AVFilmGrainParams) GetSeedAddr() *uint64 {
	return (*uint64)(&fgp.seed)
}

// Custom: GetCodecAom gets `AVFilmGrainParams.codec_aom` value.
func (fgp *AVFilmGrainParams) GetCodecAom() AVFilmGrainAOMParams {
	return (AVFilmGrainAOMParams)(
		C.get_av_film_grain_params_codec_aom((*C.struct_AVFilmGrainParams)(fgp)))
}

// Custom: SetCodecAom sets `AVFilmGrainParams.codec_aom` value.
func (fgp *AVFilmGrainParams) SetCodecAom(v AVFilmGrainAOMParams) {
	C.set_av_film_grain_params_codec_aom((*C.struct_AVFilmGrainParams)(fgp),
		(C.struct_AVFilmGrainAOMParams)(v))
}

// Custom: GetCodecAomAddr gets `AVFilmGrainParams.codec_aom` address.
func (fgp *AVFilmGrainParams) GetCodecAomAddr() *AVFilmGrainAOMParams {
	return (*AVFilmGrainAOMParams)(C.get_av_film_grain_params_codec_aom_addr(
		(*C.struct_AVFilmGrainParams)(fgp)))
}

// AvFilmGrainParamsAlloc allocates an AVFilmGrainParams structure and set its fields to
// default values.
func AvFilmGrainParamsAlloc(size *uintptr) *AVFilmGrainParams {
	return (*AVFilmGrainParams)(C.av_film_grain_params_alloc((*C.size_t)(unsafe.Pointer(size))))
}

// AvFilmGrainParamsCreateSideData allocates a complete AVFilmGrainParams and add it to the frame.
func AvFilmGrainParamsCreateSideData(frame *AVFrame) *AVFilmGrainParams {
	return (*AVFilmGrainParams)(C.av_film_grain_params_create_side_data((*C.struct_AVFrame)(frame)))
}
