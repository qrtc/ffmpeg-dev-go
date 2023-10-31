// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

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
	AV_FILM_GRAIN_PARAMS_H274 = AVFilmGrainParamsType(C.AV_FILM_GRAIN_PARAMS_H274)
)

// AVFilmGrainAOMParams
type AVFilmGrainAOMParams C.struct_AVFilmGrainAOMParams

// GetNumYPoints gets `AVFilmGrainAOMParams.num_y_points` value.
func (aomp *AVFilmGrainAOMParams) GetNumYPoints() int32 {
	return (int32)(aomp.num_y_points)
}

// SetNumYPoints sets `AVFilmGrainAOMParams.num_y_points` value.
func (aomp *AVFilmGrainAOMParams) SetNumYPoints(v int32) {
	aomp.num_y_points = (C.int)(v)
}

// GetNumYPointsAddr gets `AVFilmGrainAOMParams.num_y_points` address.
func (aomp *AVFilmGrainAOMParams) GetNumYPointsAddr() *int32 {
	return (*int32)(&aomp.num_y_points)
}

// GetYPoints gets `AVFilmGrainAOMParams.y_points` value.
func (aomp *AVFilmGrainAOMParams) GetYPoints() (v [][]uint8) {
	for i := 0; i < 14; i++ {
		v = append(v, unsafe.Slice((*uint8)(&aomp.y_points[i][0]), 2))
	}
	return v
}

// SetYPoints sets `AVFilmGrainAOMParams.y_points` value.
func (aomp *AVFilmGrainAOMParams) SetYPoints(v [][]uint8) {
	for i := 0; i < FFMIN(len(v), 14); i++ {
		for j := 0; j < FFMIN(len(v[i]), 2); j++ {
			aomp.y_points[i][j] = (C.uint8_t)(v[i][j])
		}
	}
}

// GetYPointsAddr gets `AVFilmGrainAOMParams.y_points` address.
func (aomp *AVFilmGrainAOMParams) GetYPointsAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&aomp.y_points))
}

// GetChromaScalingFromLuma gets `AVFilmGrainAOMParams.chroma_scaling_from_luma` value.
func (aomp *AVFilmGrainAOMParams) GetChromaScalingFromLuma() int32 {
	return (int32)(aomp.chroma_scaling_from_luma)
}

// SetChromaScalingFromLuma sets `AVFilmGrainAOMParams.chroma_scaling_from_luma` value.
func (aomp *AVFilmGrainAOMParams) SetChromaScalingFromLuma(v int32) {
	aomp.chroma_scaling_from_luma = (C.int)(v)
}

// GetChromaScalingFromLumaAddr gets `AVFilmGrainAOMParams.chroma_scaling_from_luma` address.
func (aomp *AVFilmGrainAOMParams) GetChromaScalingFromLumaAddr() *int32 {
	return (*int32)(&aomp.chroma_scaling_from_luma)
}

// GetNumUvPoints gets `AVFilmGrainAOMParams.num_uv_points` value.
func (aomp *AVFilmGrainAOMParams) GetNumUvPoints() []int32 {
	return unsafe.Slice((*int32)(&aomp.num_uv_points[0]), 2)
}

// SetNumUvPoints sets `AVFilmGrainAOMParams.num_uv_points` value.
func (aomp *AVFilmGrainAOMParams) SetNumUvPoints(v []int32) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		aomp.num_uv_points[i] = (C.int)(v[i])
	}
}

// GetNumUvPointsAddr gets `AVFilmGrainAOMParams.num_uv_points` address.
func (aomp *AVFilmGrainAOMParams) GetNumUvPointsAddr() **int32 {
	return (**int32)(unsafe.Pointer(&aomp.num_uv_points))
}

// GetUvPoints gets `AVFilmGrainAOMParams.uv_points` value.
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

// SetUvPoints sets `AVFilmGrainAOMParams.uv_points` value.
func (aomp *AVFilmGrainAOMParams) SetUvPoints(v [][][]uint8) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		for j := 0; j < FFMIN(len(v[i]), 10); j++ {
			for k := 0; k < FFMIN(len(v[i][j]), 2); k++ {
				aomp.uv_points[i][j][k] = (C.uint8_t)(v[i][j][k])
			}
		}
	}
}

// GetUvPointsAddr gets `AVFilmGrainAOMParams.uv_points` address.
func (aomp *AVFilmGrainAOMParams) GetUvPointsAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&aomp.uv_points))
}

// GetScalingShift gets `AVFilmGrainAOMParams.scaling_shift` value.
func (aomp *AVFilmGrainAOMParams) GetScalingShift() int32 {
	return (int32)(aomp.scaling_shift)
}

// SetScalingShift sets `AVFilmGrainAOMParams.scaling_shift` value.
func (aomp *AVFilmGrainAOMParams) SetScalingShift(v int32) {
	aomp.scaling_shift = (C.int)(v)
}

// GetScalingShiftAddr gets `AVFilmGrainAOMParams.scaling_shift` address.
func (aomp *AVFilmGrainAOMParams) GetScalingShiftAddr() *int32 {
	return (*int32)(&aomp.scaling_shift)
}

// GetArCoeffLag gets `AVFilmGrainAOMParams.ar_coeff_lag` value.
func (aomp *AVFilmGrainAOMParams) GetArCoeffLag() int32 {
	return (int32)(aomp.ar_coeff_lag)
}

// SetArCoeffLag sets `AVFilmGrainAOMParams.ar_coeff_lag` value.
func (aomp *AVFilmGrainAOMParams) SetArCoeffLag(v int32) {
	aomp.ar_coeff_lag = (C.int)(v)
}

// GetArCoeffLagAddr gets `AVFilmGrainAOMParams.ar_coeff_lag` address.
func (aomp *AVFilmGrainAOMParams) GetArCoeffLagAddr() *int32 {
	return (*int32)(&aomp.ar_coeff_lag)
}

// GetArCoeffsY gets `AVFilmGrainAOMParams.ar_coeffs_y` value.
func (aomp *AVFilmGrainAOMParams) GetArCoeffsY() []int8 {
	return unsafe.Slice((*int8)(&aomp.ar_coeffs_y[0]), 24)
}

// SetArCoeffsY sets `AVFilmGrainAOMParams.ar_coeffs_y` value.
func (aomp *AVFilmGrainAOMParams) SetArCoeffsY(v []int8) {
	for i := 0; i < FFMIN(len(v), 24); i++ {
		aomp.ar_coeffs_y[i] = (C.int8_t)(v[i])
	}
}

// GetArCoeffsYAddr gets `AVFilmGrainAOMParams.ar_coeffs_y` address.
func (aomp *AVFilmGrainAOMParams) GetArCoeffsYAddr() **int8 {
	return (**int8)(unsafe.Pointer(&aomp.ar_coeffs_y))
}

// GetArCoeffsUv gets `AVFilmGrainAOMParams.ar_coeffs_uv` value.
func (aomp *AVFilmGrainAOMParams) GetArCoeffsUv() (v [][]int8) {
	for i := 0; i < 2; i++ {
		v = append(v, unsafe.Slice((*int8)(&aomp.ar_coeffs_uv[i][0]), 25))
	}
	return v
}

// SetArCoeffsUv sets `AVFilmGrainAOMParams.ar_coeffs_uv` value.
func (aomp *AVFilmGrainAOMParams) SetArCoeffsUv(v [][]int8) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		for j := 0; j < FFMIN(len(v[i]), 25); j++ {
			aomp.ar_coeffs_uv[i][j] = (C.int8_t)(v[i][j])
		}
	}
}

// GetArCoeffsUvAddr gets `AVFilmGrainAOMParams.ar_coeffs_uv` address.
func (aomp *AVFilmGrainAOMParams) GetArCoeffsUvAddr() **int8 {
	return (**int8)(unsafe.Pointer(&aomp.ar_coeffs_uv))
}

// GetArCoeffShift gets `AVFilmGrainAOMParams.ar_coeff_shift` value.
func (aomp *AVFilmGrainAOMParams) GetArCoeffShift() int32 {
	return (int32)(aomp.ar_coeff_shift)
}

// SetArCoeffShift sets `AVFilmGrainAOMParams.ar_coeff_shift` value.
func (aomp *AVFilmGrainAOMParams) SetArCoeffShift(v int32) {
	aomp.ar_coeff_shift = (C.int)(v)
}

// GetArCoeffShiftAddr gets `AVFilmGrainAOMParams.ar_coeff_shift` address.
func (aomp *AVFilmGrainAOMParams) GetArCoeffShiftAddr() *int32 {
	return (*int32)(&aomp.ar_coeff_shift)
}

// GetGrainScaleShift gets `AVFilmGrainAOMParams.grain_scale_shift` value.
func (aomp *AVFilmGrainAOMParams) GetGrainScaleShift() int32 {
	return (int32)(aomp.grain_scale_shift)
}

// SetGrainScaleShift sets `AVFilmGrainAOMParams.grain_scale_shift` value.
func (aomp *AVFilmGrainAOMParams) SetGrainScaleShift(v int32) {
	aomp.grain_scale_shift = (C.int)(v)
}

// GetGrainScaleShiftAddr gets `AVFilmGrainAOMParams.grain_scale_shift` address.
func (aomp *AVFilmGrainAOMParams) GetGrainScaleShiftAddr() *int32 {
	return (*int32)(&aomp.grain_scale_shift)
}

// GetUvMult gets `AVFilmGrainAOMParams.uv_mult` value.
func (aomp *AVFilmGrainAOMParams) GetUvMult() []int32 {
	return unsafe.Slice((*int32)(&aomp.uv_mult[0]), 2)
}

// SetUvMult sets `AVFilmGrainAOMParams.uv_mult` value.
func (aomp *AVFilmGrainAOMParams) SetUvMult(v []int32) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		aomp.uv_mult[i] = (C.int)(v[i])
	}
}

// GetUvMultAddr gets `AVFilmGrainAOMParams.uv_mult` address.
func (aomp *AVFilmGrainAOMParams) GetUvMultAddr() **int32 {
	return (**int32)(unsafe.Pointer(&aomp.uv_mult))
}

// GetUvMultLuma gets `AVFilmGrainAOMParams.uv_mult_luma` value.
func (aomp *AVFilmGrainAOMParams) GetUvMultLuma() []int32 {
	return unsafe.Slice((*int32)(&aomp.uv_mult_luma[0]), 2)
}

// SetUvMultLuma sets `AVFilmGrainAOMParams.uv_mult_luma` value.
func (aomp *AVFilmGrainAOMParams) SetUvMultLuma(v []int32) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		aomp.uv_mult_luma[i] = (C.int)(v[i])
	}
}

// GetUvMultLumaAddr gets `AVFilmGrainAOMParams.uv_mult_luma` address.
func (aomp *AVFilmGrainAOMParams) GetUvMultLumaAddr() **int32 {
	return (**int32)(unsafe.Pointer(&aomp.uv_mult_luma))
}

// GetUvOffset gets `AVFilmGrainAOMParams.uv_offset` value.
func (aomp *AVFilmGrainAOMParams) GetUvOffset() []int32 {
	return unsafe.Slice((*int32)(&aomp.uv_offset[0]), 2)
}

// SetUvOffset sets `AVFilmGrainAOMParams.uv_offset` value.
func (aomp *AVFilmGrainAOMParams) SetUvOffset(v []int32) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		aomp.uv_offset[i] = (C.int)(v[i])
	}
}

// GetUvOffsetAddr gets `AVFilmGrainAOMParams.uv_offset` address.
func (aomp *AVFilmGrainAOMParams) GetUvOffsetAddr() **int32 {
	return (**int32)(unsafe.Pointer(&aomp.uv_offset))
}

// GetOverlapFlag gets `AVFilmGrainAOMParams.overlap_flag` value.
func (aomp *AVFilmGrainAOMParams) GetOverlapFlag() int32 {
	return (int32)(aomp.overlap_flag)
}

// SetOverlapFlag sets `AVFilmGrainAOMParams.overlap_flag` value.
func (aomp *AVFilmGrainAOMParams) SetOverlapFlag(v int32) {
	aomp.overlap_flag = (C.int)(v)
}

// GetOverlapFlagAddr gets `AVFilmGrainAOMParams.overlap_flag` address.
func (aomp *AVFilmGrainAOMParams) GetOverlapFlagAddr() *int32 {
	return (*int32)(&aomp.overlap_flag)
}

// GetLimitOutputRange gets `AVFilmGrainAOMParams.limit_outputrange` value.
func (aomp *AVFilmGrainAOMParams) GetLimitOutputRange() int32 {
	return (int32)(aomp.limit_output_range)
}

// SetLimitOutputRange sets `AVFilmGrainAOMParams.limit_outputrange` value.
func (aomp *AVFilmGrainAOMParams) SetLimitOutputRange(v int32) {
	aomp.limit_output_range = (C.int)(v)
}

// GetLimitOutputRangeAddr gets `AVFilmGrainAOMParams.limit_outputrange` address.
func (aomp *AVFilmGrainAOMParams) GetLimitOutputRangeAddr() *int32 {
	return (*int32)(&aomp.limit_output_range)
}

// AVFilmGrainH274Params
type AVFilmGrainH274Params C.struct_AVFilmGrainH274Params

// GetModelId gets `AVFilmGrainH274Params.model_id` value.
func (h274p *AVFilmGrainH274Params) GetModelId() int32 {
	return (int32)(h274p.model_id)
}

// SetModelId sets `AVFilmGrainH274Params.model_id` value.
func (h274p *AVFilmGrainH274Params) SetModelId(v int32) {
	h274p.model_id = (C.int)(v)
}

// GetModelIdAddr gets `AVFilmGrainH274Params.model_id` address.
func (h274p *AVFilmGrainH274Params) GetModelIdAddr() *int32 {
	return (*int32)(&h274p.model_id)
}

// GetBitDepthLuma gets `AVFilmGrainH274Params.bit_depth_luma` value.
func (h274p *AVFilmGrainH274Params) GetBitDepthLuma() int32 {
	return (int32)(h274p.bit_depth_luma)
}

// SetBitDepthLuma sets `AVFilmGrainH274Params.bit_depth_luma` value.
func (h274p *AVFilmGrainH274Params) SetBitDepthLuma(v int32) {
	h274p.bit_depth_luma = (C.int)(v)
}

// GetBitDepthLumaAddr gets `AVFilmGrainH274Params.bit_depth_luma` address.
func (h274p *AVFilmGrainH274Params) GetBitDepthLumaAddr() *int32 {
	return (*int32)(&h274p.bit_depth_luma)
}

// GetBitDepthChroma gets `AVFilmGrainH274Params.bit_depth_chroma` value.
func (h274p *AVFilmGrainH274Params) GetBitDepthChroma() int32 {
	return (int32)(h274p.bit_depth_chroma)
}

// SetBitDepthChroma sets `AVFilmGrainH274Params.bit_depth_chroma` value.
func (h274p *AVFilmGrainH274Params) SetBitDepthChroma(v int32) {
	h274p.bit_depth_chroma = (C.int)(v)
}

// GetBitDepthChromaAddr gets `AVFilmGrainH274Params.bit_depth_chroma` address.
func (h274p *AVFilmGrainH274Params) GetBitDepthChromaAddr() *int32 {
	return (*int32)(&h274p.bit_depth_chroma)
}

// GetColorRange gets `AVFilmGrainH274Params.color_range` value.
func (h274p *AVFilmGrainH274Params) GetColorRange() AVColorRange {
	return (AVColorRange)(h274p.color_range)
}

// SetColorRange sets `AVFilmGrainH274Params.color_range` value.
func (h274p *AVFilmGrainH274Params) SetColorRange(v AVColorRange) {
	h274p.color_range = (C.enum_AVColorRange)(v)
}

// GetColorRangeAddr gets `AVFilmGrainH274Params.color_range` address.
func (h274p *AVFilmGrainH274Params) GetColorRangeAddr() *AVColorRange {
	return (*AVColorRange)(&h274p.color_range)
}

// GetColorPrimaries gets `AVFilmGrainH274Params.color_primaries` value.
func (h274p *AVFilmGrainH274Params) GetColorPrimaries() AVColorPrimaries {
	return (AVColorPrimaries)(h274p.color_primaries)
}

// SetColorPrimaries sets `AVFilmGrainH274Params.color_primaries` value.
func (h274p *AVFilmGrainH274Params) SetColorPrimaries(v AVColorPrimaries) {
	h274p.color_primaries = (C.enum_AVColorPrimaries)(v)
}

// GetColorPrimariesAddr gets `AVFilmGrainH274Params.color_primaries` address.
func (h274p *AVFilmGrainH274Params) GetColorPrimariesAddr() *AVColorPrimaries {
	return (*AVColorPrimaries)(&h274p.color_primaries)
}

// GetColorTrc gets `AVFilmGrainH274Params.color_trc` value.
func (h274p *AVFilmGrainH274Params) GetColorTrc() AVColorTransferCharacteristic {
	return (AVColorTransferCharacteristic)(h274p.color_trc)
}

// SetColorTrc sets `AVFilmGrainH274Params.color_trc` value.
func (h274p *AVFilmGrainH274Params) SetColorTrc(v AVColorTransferCharacteristic) {
	h274p.color_trc = (C.enum_AVColorTransferCharacteristic)(v)
}

// GetColorTrcAddr gets `AVFilmGrainH274Params.color_trc` address.
func (h274p *AVFilmGrainH274Params) GetColorTrcAddr() *AVColorTransferCharacteristic {
	return (*AVColorTransferCharacteristic)(&h274p.color_trc)
}

// GetColorSpace gets `AVFilmGrainH274Params.color_space` value.
func (h274p *AVFilmGrainH274Params) GetColorSpace() AVColorSpace {
	return (AVColorSpace)(h274p.color_space)
}

// SetColorSpace sets `AVFilmGrainH274Params.color_space` value.
func (h274p *AVFilmGrainH274Params) SetColorSpace(v AVColorSpace) {
	h274p.color_space = (C.enum_AVColorSpace)(v)
}

// GetColorSpaceAddr gets `AVFilmGrainH274Params.color_space` address.
func (h274p *AVFilmGrainH274Params) GetColorSpaceAddr() *AVColorSpace {
	return (*AVColorSpace)(&h274p.color_space)
}

// GetBlendingModeId gets `AVFilmGrainH274Params.blending_mode_id` value.
func (h274p *AVFilmGrainH274Params) GetBlendingModeId() int32 {
	return (int32)(h274p.blending_mode_id)
}

// SetBlendingModeId sets `AVFilmGrainH274Params.blending_mode_id` value.
func (h274p *AVFilmGrainH274Params) SetBlendingModeId(v int32) {
	h274p.blending_mode_id = (C.int)(v)
}

// GetBlendingModeIdAddr gets `AVFilmGrainH274Params.blending_mode_id` address.
func (h274p *AVFilmGrainH274Params) GetBlendingModeIdAddr() *int32 {
	return (*int32)(&h274p.blending_mode_id)
}

// GetLog2ScaleFactor gets `AVFilmGrainH274Params.log2_scale_factor` value.
func (h274p *AVFilmGrainH274Params) GetLog2ScaleFactor() int32 {
	return (int32)(h274p.log2_scale_factor)
}

// SetLog2ScaleFactor sets `AVFilmGrainH274Params.log2_scale_factor` value.
func (h274p *AVFilmGrainH274Params) SetLog2ScaleFactor(v int32) {
	h274p.log2_scale_factor = (C.int)(v)
}

// GetLog2ScaleFactorAddr gets `AVFilmGrainH274Params.log2_scale_factor` address.
func (h274p *AVFilmGrainH274Params) GetLog2ScaleFactorAddr() *int32 {
	return (*int32)(&h274p.log2_scale_factor)
}

// GetComponentModelPresent gets `AVFilmGrainH274Params.component_model_present` value.
func (h274p *AVFilmGrainH274Params) GetComponentModelPresent() []int32 {
	return unsafe.Slice((*int32)(&h274p.component_model_present[0]), 3)
}

// SetComponentModelPresent sets `AVFilmGrainH274Params.component_model_present` value.
func (h274p *AVFilmGrainH274Params) SetComponentModelPresent(v []int32) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		h274p.component_model_present[i] = (C.int)(v[i])
	}
}

// GetComponentModelPresentAddr gets `AVFilmGrainH274Params.component_model_present` address.
func (h274p *AVFilmGrainH274Params) GetComponentModelPresentAddr() **int32 {
	return (**int32)(unsafe.Pointer(&h274p.component_model_present))
}

// GetNumIntensityIntervals gets `AVFilmGrainH274Params.num_intensity_intervals` value.
func (h274p *AVFilmGrainH274Params) GetNumIntensityIntervals() []uint16 {
	return unsafe.Slice((*uint16)(&h274p.num_intensity_intervals[0]), 3)
}

// SetNumIntensityIntervals sets `AVFilmGrainH274Params.num_intensity_intervals` value.
func (h274p *AVFilmGrainH274Params) SetNumIntensityIntervals(v []uint16) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		h274p.num_intensity_intervals[i] = (C.uint16_t)(v[i])
	}
}

// GetNumIntensityIntervalsAddr gets `AVFilmGrainH274Params.num_intensity_intervals` address.
func (h274p *AVFilmGrainH274Params) GetNumIntensityIntervalsAddr() **uint16 {
	return (**uint16)(unsafe.Pointer(&h274p.num_intensity_intervals))
}

// GetNumModelValues gets `AVFilmGrainH274Params.num_model_values` value.
func (h274p *AVFilmGrainH274Params) GetNumModelValues() []uint8 {
	return unsafe.Slice((*uint8)(&h274p.num_model_values[0]), 3)
}

// SetNumModelValues sets `AVFilmGrainH274Params.num_model_values` value.
func (h274p *AVFilmGrainH274Params) SetNumModelValues(v []uint8) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		h274p.num_model_values[i] = (C.uint8_t)(v[i])
	}
}

// GetNumModelValuesAddr gets `AVFilmGrainH274Params.num_model_values` address.
func (h274p *AVFilmGrainH274Params) GetNumModelValuesAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&h274p.num_model_values))
}

// GetIntensityIntervalLowerBound gets `AVFilmGrainH274Params.intensity_interval_lower_bound` value.
func (h274p *AVFilmGrainH274Params) GetIntensityIntervalLowerBound() (v [][]uint8) {
	for i := 0; i < 3; i++ {
		v = append(v, unsafe.Slice((*uint8)(&h274p.intensity_interval_lower_bound[i][0]), 256))
	}
	return v
}

// SetIntensityIntervalLowerBound sets `AVFilmGrainH274Params.intensity_interval_lower_bound` value.
func (h274p *AVFilmGrainH274Params) SetIntensityIntervalLowerBound(v [][]uint8) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		for j := 0; j < FFMIN(len(v[i]), 256); j++ {
			h274p.intensity_interval_lower_bound[i][j] = (C.uint8_t)(v[i][j])
		}
	}
}

// GetIntensityIntervalLowerBoundAddr gets `AVFilmGrainH274Params.intensity_interval_lower_bound` address.
func (h274p *AVFilmGrainH274Params) GetIntensityIntervalLowerBoundAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&h274p.intensity_interval_lower_bound))
}

// GetIntensityIntervalUpperBound gets `AVFilmGrainH274Params.intensity_interval_upper_bound` value.
func (h274p *AVFilmGrainH274Params) GetIntensityIntervalUpperBound() (v [][]uint8) {
	for i := 0; i < 3; i++ {
		v = append(v, unsafe.Slice((*uint8)(&h274p.intensity_interval_upper_bound[i][0]), 256))
	}
	return v
}

// SetIntensityIntervalUpperBound sets `AVFilmGrainH274Params.intensity_interval_upper_bound` value.
func (h274p *AVFilmGrainH274Params) SetIntensityIntervalUpperBound(v [][]uint8) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		for j := 0; j < FFMIN(len(v[i]), 256); j++ {
			h274p.intensity_interval_upper_bound[i][j] = (C.uint8_t)(v[i][j])
		}
	}
}

// GetIntensityIntervalUpperBoundAddr gets `AVFilmGrainH274Params.intensity_interval_upper_bound` address.
func (h274p *AVFilmGrainH274Params) GetIntensityIntervalUpperBoundAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&h274p.intensity_interval_upper_bound))
}

// GetCompModelValue gets `AVFilmGrainH274Params.comp_model_value` value.
func (h274p *AVFilmGrainH274Params) GetCompModelValue() (v [][][]int16) {
	for i := 0; i < 3; i++ {
		tmp := [][]int16{}
		for j := 0; j < 256; j++ {
			tmp = append(tmp, unsafe.Slice((*int16)(&h274p.comp_model_value[i][j][0]), 6))
		}
		v = append(v, tmp)
	}
	return v
}

// SetCompModelValue sets `AVFilmGrainH274Params.comp_model_value` value.
func (h274p *AVFilmGrainH274Params) SetCompModelValue(v [][][]int16) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		for j := 0; j < FFMIN(len(v[i]), 256); j++ {
			for k := 0; k < FFMIN(len(v[i][j]), 6); k++ {
				h274p.comp_model_value[i][j][k] = (C.int16_t)(v[i][j][k])
			}
		}
	}
}

// GetCompModelValueAddr gets `AVFilmGrainH274Params.comp_model_value` address.
func (h274p *AVFilmGrainH274Params) GetCompModelValueAddr() **int16 {
	return (**int16)(unsafe.Pointer(&h274p.comp_model_value))
}

// AVFilmGrainParams
type AVFilmGrainParams C.struct_AVFilmGrainParams

// GetType gets `AVFilmGrainParams.type` value.
func (fgp *AVFilmGrainParams) GetType() AVFilmGrainParamsType {
	return (AVFilmGrainParamsType)(fgp._type)
}

// SetType sets `AVFilmGrainParams.type` value.
func (fgp *AVFilmGrainParams) SetType(v AVFilmGrainParamsType) {
	fgp._type = (C.enum_AVFilmGrainParamsType)(v)
}

// GetTypeAddr gets `AVFilmGrainParams.type` address.
func (fgp *AVFilmGrainParams) GetTypeAddr() *AVFilmGrainParamsType {
	return (*AVFilmGrainParamsType)(&fgp._type)
}

// GetSeed gets `AVFilmGrainParams.seed` value.
func (fgp *AVFilmGrainParams) GetSeed() uint64 {
	return (uint64)(fgp.seed)
}

// SetSeed sets `AVFilmGrainParams.seed` value.
func (fgp *AVFilmGrainParams) SetSeed(v uint64) {
	fgp.seed = (C.uint64_t)(v)
}

// GetSeedAddr gets `AVFilmGrainParams.seed` address.
func (fgp *AVFilmGrainParams) GetSeedAddr() *uint64 {
	return (*uint64)(&fgp.seed)
}

// GetCodecAom gets `AVFilmGrainParams.codec_aom` value.
func (fgp *AVFilmGrainParams) GetCodecAom() AVFilmGrainAOMParams {
	return (AVFilmGrainAOMParams)(
		C.get_av_film_grain_params_codec_aom((*C.struct_AVFilmGrainParams)(fgp)))
}

// SetCodecAom sets `AVFilmGrainParams.codec_aom` value.
func (fgp *AVFilmGrainParams) SetCodecAom(v AVFilmGrainAOMParams) {
	C.set_av_film_grain_params_codec_aom((*C.struct_AVFilmGrainParams)(fgp),
		(C.struct_AVFilmGrainAOMParams)(v))
}

// GetCodecAomAddr gets `AVFilmGrainParams.codec_aom` address.
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
