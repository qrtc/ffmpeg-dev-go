// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/hdr_dynamic_vivid_metadata.h>
*/
import "C"
import "unsafe"

// AVHDRVivid3SplineParams
type AVHDRVivid3SplineParams C.struct_AVHDRVivid3SplineParams

// GetThMode gets `AVHDRVivid3SplineParams.th_mode` value.
func (d3sp *AVHDRVivid3SplineParams) GetThMode() int32 {
	return (int32)(d3sp.th_mode)
}

// SetThMode sets `AVHDRVivid3SplineParams.th_mode` value.
func (d3sp *AVHDRVivid3SplineParams) SetThMode(v int32) {
	d3sp.th_mode = (C.int)(v)
}

// GetThModeAddr gets `AVHDRVivid3SplineParams.th_mode` address.
func (d3sp *AVHDRVivid3SplineParams) GetThModeAddr() *int32 {
	return (*int32)(&d3sp.th_mode)
}

// GetThEnableMb gets `AVHDRVivid3SplineParams.th_enable_mb` value.
func (d3sp *AVHDRVivid3SplineParams) GetThEnableMb() AVRational {
	return (AVRational)(d3sp.th_enable_mb)
}

// SetThEnableMb sets `AVHDRVivid3SplineParams.th_enable_mb` value.
func (d3sp *AVHDRVivid3SplineParams) SetThEnableMb(v AVRational) {
	d3sp.th_enable_mb = (C.struct_AVRational)(v)
}

// GetThEnableMbAddr gets `AVHDRVivid3SplineParams.th_enable_mb` address.
func (d3sp *AVHDRVivid3SplineParams) GetThEnableMbAddr() *AVRational {
	return (*AVRational)(&d3sp.th_enable_mb)
}

// GetThEnable gets `AVHDRVivid3SplineParams.th_enable` value.
func (d3sp *AVHDRVivid3SplineParams) GetThEnable() AVRational {
	return (AVRational)(d3sp.th_enable)
}

// SetThEnable sets `AVHDRVivid3SplineParams.th_enable` value.
func (d3sp *AVHDRVivid3SplineParams) SetThEnable(v AVRational) {
	d3sp.th_enable = (C.struct_AVRational)(v)
}

// GetThEnableAddr gets `AVHDRVivid3SplineParams.th_enable` address.
func (d3sp *AVHDRVivid3SplineParams) GetThEnableAddr() *AVRational {
	return (*AVRational)(&d3sp.th_enable)
}

// GetThDelta1 gets `AVHDRVivid3SplineParams.th_delta1` value.
func (d3sp *AVHDRVivid3SplineParams) GetThDelta1() AVRational {
	return (AVRational)(d3sp.th_delta1)
}

// SetThDelta1 sets `AVHDRVivid3SplineParams.th_delta1` value.
func (d3sp *AVHDRVivid3SplineParams) SetThDelta1(v AVRational) {
	d3sp.th_delta1 = (C.struct_AVRational)(v)
}

// GetThDelta1Addr gets `AVHDRVivid3SplineParams.th_delta1` address.
func (d3sp *AVHDRVivid3SplineParams) GetThDelta1Addr() *AVRational {
	return (*AVRational)(&d3sp.th_delta1)
}

// GetThDelta2 gets `AVHDRVivid3SplineParams.th_delta2` value.
func (d3sp *AVHDRVivid3SplineParams) GetThDelta2() AVRational {
	return (AVRational)(d3sp.th_delta2)
}

// SetThDelta2 sets `AVHDRVivid3SplineParams.th_delta2` value.
func (d3sp *AVHDRVivid3SplineParams) SetThDelta2(v AVRational) {
	d3sp.th_delta2 = (C.struct_AVRational)(v)
}

// GetThDelta2Addr gets `AVHDRVivid3SplineParams.th_delta2` address.
func (d3sp *AVHDRVivid3SplineParams) GetThDelta2Addr() *AVRational {
	return (*AVRational)(&d3sp.th_delta2)
}

// GetEnableStrength gets `AVHDRVivid3SplineParams.enable_strength` value.
func (d3sp *AVHDRVivid3SplineParams) GetEnableStrength() AVRational {
	return (AVRational)(d3sp.enable_strength)
}

// SetEnableStrength sets `AVHDRVivid3SplineParams.enable_strength` value.
func (d3sp *AVHDRVivid3SplineParams) SetEnableStrength(v AVRational) {
	d3sp.enable_strength = (C.struct_AVRational)(v)
}

// GetEnableStrengthAddr gets `AVHDRVivid3SplineParams.enable_strength` address.
func (d3sp *AVHDRVivid3SplineParams) GetEnableStrengthAddr() *AVRational {
	return (*AVRational)(&d3sp.enable_strength)
}

// AVHDRVividColorToneMappingParams
type AVHDRVividColorToneMappingParams C.struct_AVHDRVividColorToneMappingParams

// GetTargetedSystemDisplayMaximumLuminance gets `AVHDRVividColorToneMappingParams.targeted_system_display_maximum_luminance` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetTargetedSystemDisplayMaximumLuminance() AVRational {
	return (AVRational)(ctmp.targeted_system_display_maximum_luminance)
}

// SetTargetedSystemDisplayMaximumLuminance sets `AVHDRVividColorToneMappingParams.targeted_system_display_maximum_luminance` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetTargetedSystemDisplayMaximumLuminance(v AVRational) {
	ctmp.targeted_system_display_maximum_luminance = (C.struct_AVRational)(v)
}

// GetTargetedSystemDisplayMaximumLuminanceAddr gets `AVHDRVividColorToneMappingParams.targeted_system_display_maximum_luminance` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetTargetedSystemDisplayMaximumLuminanceAddr() *AVRational {
	return (*AVRational)(&ctmp.targeted_system_display_maximum_luminance)
}

// GetBaseEnableFlag gets `AVHDRVividColorToneMappingParams.base_enable_flag` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseEnableFlag() int32 {
	return (int32)(ctmp.base_enable_flag)
}

// SetBaseEnableFlag sets `AVHDRVividColorToneMappingParams.base_enable_flag` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseEnableFlag(v int32) {
	ctmp.base_enable_flag = (C.int)(v)
}

// GetBaseEnableFlagAddr gets `AVHDRVividColorToneMappingParams.base_enable_flag` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseEnableFlagAddr() *int32 {
	return (*int32)(&ctmp.base_enable_flag)
}

// GetBaseParamMP gets `AVHDRVividColorToneMappingParams.base_param_m_p` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMP() AVRational {
	return (AVRational)(ctmp.base_param_m_p)
}

// SetBaseParamMP sets `AVHDRVividColorToneMappingParams.base_param_m_p` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamMP(v AVRational) {
	ctmp.base_param_m_p = (C.struct_AVRational)(v)
}

// GetBaseParamMPAddr gets `AVHDRVividColorToneMappingParams.base_param_m_p` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMPAddr() *AVRational {
	return (*AVRational)(&ctmp.base_param_m_p)
}

// GetBaseParamMM gets `AVHDRVividColorToneMappingParams.base_param_m_m` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMM() AVRational {
	return (AVRational)(ctmp.base_param_m_m)
}

// SetBaseParamMM sets `AVHDRVividColorToneMappingParams.base_param_m_m` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamMM(v AVRational) {
	ctmp.base_param_m_m = (C.struct_AVRational)(v)
}

// GetBaseParamMMAddr gets `AVHDRVividColorToneMappingParams.base_param_m_m` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMMAddr() *AVRational {
	return (*AVRational)(&ctmp.base_param_m_m)
}

// GetBaseParamMA gets `AVHDRVividColorToneMappingParams.base_param_m_a` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMA() AVRational {
	return (AVRational)(ctmp.base_param_m_a)
}

// SetBaseParamMA sets `AVHDRVividColorToneMappingParams.base_param_m_a` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamMA(v AVRational) {
	ctmp.base_param_m_a = (C.struct_AVRational)(v)
}

// GetBaseParamMAAddr gets `AVHDRVividColorToneMappingParams.base_param_m_a` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMAAddr() *AVRational {
	return (*AVRational)(&ctmp.base_param_m_a)
}

// GetBaseParamMB gets `AVHDRVividColorToneMappingParams.base_param_m_b` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMB() AVRational {
	return (AVRational)(ctmp.base_param_m_b)
}

// SetBaseParamMB sets `AVHDRVividColorToneMappingParams.base_param_m_b` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamMB(v AVRational) {
	ctmp.base_param_m_b = (C.struct_AVRational)(v)
}

// GetBaseParamMBAddr gets `AVHDRVividColorToneMappingParams.base_param_m_b` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMBAddr() *AVRational {
	return (*AVRational)(&ctmp.base_param_m_b)
}

// GetBaseParamMN gets `AVHDRVividColorToneMappingParams.base_param_m_n` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMN() AVRational {
	return (AVRational)(ctmp.base_param_m_n)
}

// SetBaseParamMN sets `AVHDRVividColorToneMappingParams.base_param_m_n` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamMN(v AVRational) {
	ctmp.base_param_m_n = (C.struct_AVRational)(v)
}

// GetBaseParamMNAddr gets `AVHDRVividColorToneMappingParams.base_param_m_n` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamMNAddr() *AVRational {
	return (*AVRational)(&ctmp.base_param_m_n)
}

// GetBaseParamK1 gets `AVHDRVividColorToneMappingParams.base_param_k1` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamK1() int32 {
	return (int32)(ctmp.base_param_k1)
}

// SetBaseParamK1 sets `AVHDRVividColorToneMappingParams.base_param_k1` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamK1(v int32) {
	ctmp.base_param_k1 = (C.int)(v)
}

// GetBaseParamK1Addr gets `AVHDRVividColorToneMappingParams.base_param_k1` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamK1Addr() *int32 {
	return (*int32)(&ctmp.base_param_k1)
}

// GetBaseParamK2 gets `AVHDRVividColorToneMappingParams.base_param_k2` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamK2() int32 {
	return (int32)(ctmp.base_param_k2)
}

// SetBaseParamK2 sets `AVHDRVividColorToneMappingParams.base_param_k2` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamK2(v int32) {
	ctmp.base_param_k2 = (C.int)(v)
}

// GetBaseParamK2Addr gets `AVHDRVividColorToneMappingParams.base_param_k2` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamK2Addr() *int32 {
	return (*int32)(&ctmp.base_param_k2)
}

// GetBaseParamK3 gets `AVHDRVividColorToneMappingParams.base_param_k3` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamK3() int32 {
	return (int32)(ctmp.base_param_k3)
}

// SetBaseParamK3 sets `AVHDRVividColorToneMappingParams.base_param_k3` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamK3(v int32) {
	ctmp.base_param_k3 = (C.int)(v)
}

// GetBaseParamK3Addr gets `AVHDRVividColorToneMappingParams.base_param_k3` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamK3Addr() *int32 {
	return (*int32)(&ctmp.base_param_k3)
}

// GetBaseParamDeltaEnableMode gets `AVHDRVividColorToneMappingParams.base_param_Delta_enable_mode` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamDeltaEnableMode() int32 {
	return (int32)(ctmp.base_param_Delta_enable_mode)
}

// SetBaseParamDeltaEnableMode sets `AVHDRVividColorToneMappingParams.base_param_Delta_enable_mode` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamDeltaEnableMode(v int32) {
	ctmp.base_param_Delta_enable_mode = (C.int)(v)
}

// GetBaseParamDeltaEnableModeAddr gets `AVHDRVividColorToneMappingParams.base_param_Delta_enable_mode` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamDeltaEnableModeAddr() *int32 {
	return (*int32)(&ctmp.base_param_Delta_enable_mode)
}

// GetBaseParamDelta gets `AVHDRVividColorToneMappingParams.base_param_Delta` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamDelta() AVRational {
	return (AVRational)(ctmp.base_param_Delta)
}

// SetBaseParamDelta sets `AVHDRVividColorToneMappingParams.base_param_Delta` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetBaseParamDelta(v AVRational) {
	ctmp.base_param_Delta = (C.struct_AVRational)(v)
}

// GetBaseParamDeltaAddr gets `AVHDRVividColorToneMappingParams.base_param_Delta` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetBaseParamDeltaAddr() *AVRational {
	return (*AVRational)(&ctmp.base_param_Delta)
}

// GetThreeSplineEnableFlag gets `AVHDRVividColorToneMappingParams.three_Spline_enable_flag` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineEnableFlag() int32 {
	return (int32)(ctmp.three_Spline_enable_flag)
}

// SetThreeSplineEnableFlag sets `AVHDRVividColorToneMappingParams.three_Spline_enable_flag` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetThreeSplineEnableFlag(v int32) {
	ctmp.three_Spline_enable_flag = (C.int)(v)
}

// GetThreeSplineEnableFlagAddr gets `AVHDRVividColorToneMappingParams.three_Spline_enable_flag` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineEnableFlagAddr() *int32 {
	return (*int32)(&ctmp.three_Spline_enable_flag)
}

// GetThreeSplineNum gets `AVHDRVividColorToneMappingParams.three_Spline_num` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineNum() int32 {
	return (int32)(ctmp.three_Spline_num)
}

// SetThreeSplineNum sets `AVHDRVividColorToneMappingParams.three_Spline_num` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetThreeSplineNum(v int32) {
	ctmp.three_Spline_num = (C.int)(v)
}

// GetThreeSplineNumAddr gets `AVHDRVividColorToneMappingParams.three_Spline_num` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineNumAddr() *int32 {
	return (*int32)(&ctmp.three_Spline_num)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThMode gets `AVHDRVividColorToneMappingParams.three_Spline_TH_mode` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThMode() int32 {
	return (int32)(ctmp.three_Spline_TH_mode)
}

// Deprecated: Use three_spline instead.
//
// SetThreeSplineThMode sets `AVHDRVividColorToneMappingParams.three_Spline_TH_mode` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetThreeSplineThMode(v int32) {
	ctmp.three_Spline_TH_mode = (C.int)(v)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThModeAddr gets `AVHDRVividColorToneMappingParams.three_Spline_TH_mode` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThModeAddr() *int32 {
	return (*int32)(&ctmp.three_Spline_TH_mode)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThEnableMb gets `AVHDRVividColorToneMappingParams.three_Spline_TH_enable_MB` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThEnableMb() AVRational {
	return (AVRational)(ctmp.three_Spline_TH_enable_MB)
}

// Deprecated: Use three_spline instead.
//
// SetThreeSplineThEnableMb sets `AVHDRVividColorToneMappingParams.three_Spline_TH_enable_MB` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetThreeSplineThEnableMb(v AVRational) {
	ctmp.three_Spline_TH_enable_MB = (C.struct_AVRational)(v)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThEnableMbAddr gets `AVHDRVividColorToneMappingParams.three_Spline_TH_enable_MB` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThEnableMbAddr() *AVRational {
	return (*AVRational)(&ctmp.three_Spline_TH_enable_MB)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThEnable gets `AVHDRVividColorToneMappingParams.three_Spline_TH_enable` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThEnable() AVRational {
	return (AVRational)(ctmp.three_Spline_TH_enable)
}

// Deprecated: Use three_spline instead.
//
// SetThreeSplineThEnable sets `AVHDRVividColorToneMappingParams.three_Spline_TH_enable` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetThreeSplineThEnable(v AVRational) {
	ctmp.three_Spline_TH_enable = (C.struct_AVRational)(v)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThEnableAddr gets `AVHDRVividColorToneMappingParams.three_Spline_TH_enable` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThEnableAddr() *AVRational {
	return (*AVRational)(&ctmp.three_Spline_TH_enable)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThDelta1 gets `AVHDRVividColorToneMappingParams.three_Spline_TH_Delta1` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThDelta1() AVRational {
	return (AVRational)(ctmp.three_Spline_TH_Delta1)
}

// Deprecated: Use three_spline instead.
//
// SetThreeSplineThDelta1 sets `AVHDRVividColorToneMappingParams.three_Spline_TH_Delta1` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetThreeSplineThDelta1(v AVRational) {
	ctmp.three_Spline_TH_Delta1 = (C.struct_AVRational)(v)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThDelta1Addr gets `AVHDRVividColorToneMappingParams.three_Spline_TH_Delta1` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThDelta1Addr() *AVRational {
	return (*AVRational)(&ctmp.three_Spline_TH_Delta1)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThDelta2 gets `AVHDRVividColorToneMappingParams.three_Spline_TH_Delta2` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThDelta2() AVRational {
	return (AVRational)(ctmp.three_Spline_TH_Delta2)
}

// Deprecated: Use three_spline instead.
//
// SetThreeSplineThDelta2 sets `AVHDRVividColorToneMappingParams.three_Spline_TH_Delta2` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetThreeSplineThDelta2(v AVRational) {
	ctmp.three_Spline_TH_Delta2 = (C.struct_AVRational)(v)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineThDelta2Addr gets `AVHDRVividColorToneMappingParams.three_Spline_TH_Delta2` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineThDelta2Addr() *AVRational {
	return (*AVRational)(&ctmp.three_Spline_TH_Delta2)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineEnableStrength gets `AVHDRVividColorToneMappingParams.three_Spline_enable_Strength` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineEnableStrength() AVRational {
	return (AVRational)(ctmp.three_Spline_enable_Strength)
}

// Deprecated: Use three_spline instead.
//
// SetThreeSplineEnableStrength sets `AVHDRVividColorToneMappingParams.three_Spline_enable_Strength` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetThreeSplineEnableStrength(v AVRational) {
	ctmp.three_Spline_enable_Strength = (C.struct_AVRational)(v)
}

// Deprecated: Use three_spline instead.
//
// GetThreeSplineEnableStrengthAddr gets `AVHDRVividColorToneMappingParams.three_Spline_enable_Strength` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineEnableStrengthAddr() *AVRational {
	return (*AVRational)(&ctmp.three_Spline_enable_Strength)
}

// GetThreeSpline gets `AVHDRVividColorToneMappingParams.three_spline` value.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSpline() []AVHDRVivid3SplineParams {
	return unsafe.Slice((*AVHDRVivid3SplineParams)(&ctmp.three_spline[0]), 2)
}

// SetThreeSpline sets `AVHDRVividColorToneMappingParams.three_spline` value.
func (ctmp *AVHDRVividColorToneMappingParams) SetThreeSpline(v []AVHDRVivid3SplineParams) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		ctmp.three_spline[i] = (C.struct_AVHDRVivid3SplineParams)(v[i])
	}
}

// GetThreeSplineAddr gets `AVHDRVividColorToneMappingParams.three_spline` address.
func (ctmp *AVHDRVividColorToneMappingParams) GetThreeSplineAddr() **AVHDRVivid3SplineParams {
	return (**AVHDRVivid3SplineParams)(unsafe.Pointer(&ctmp.three_spline))
}

// AVHDRVividColorTransformParams
type AVHDRVividColorTransformParams C.struct_AVHDRVividColorTransformParams

// GetMinimumMaxrgb gets `AVHDRVividColorTransformParams.minimum_maxrgb` value.
func (ctfp *AVHDRVividColorTransformParams) GetMinimumMaxrgb() AVRational {
	return (AVRational)(ctfp.minimum_maxrgb)
}

// SetMinimumMaxrgb sets `AVHDRVividColorTransformParams.minimum_maxrgb` value.
func (ctfp *AVHDRVividColorTransformParams) SetMinimumMaxrgb(v AVRational) {
	ctfp.minimum_maxrgb = (C.struct_AVRational)(v)
}

// GetMinimumMaxrgbAddr gets `AVHDRVividColorTransformParams.minimum_maxrgb` address.
func (ctfp *AVHDRVividColorTransformParams) GetMinimumMaxrgbAddr() *AVRational {
	return (*AVRational)(&ctfp.minimum_maxrgb)
}

// GetAverageMaxrgb gets `AVHDRVividColorTransformParams.average_maxrgb` value.
func (ctfp *AVHDRVividColorTransformParams) GetAverageMaxrgb() AVRational {
	return (AVRational)(ctfp.average_maxrgb)
}

// SetAverageMaxrgb sets `AVHDRVividColorTransformParams.average_maxrgb` value.
func (ctfp *AVHDRVividColorTransformParams) SetAverageMaxrgb(v AVRational) {
	ctfp.average_maxrgb = (C.struct_AVRational)(v)
}

// GetAverageMaxrgbAddr gets `AVHDRVividColorTransformParams.average_maxrgb` address.
func (ctfp *AVHDRVividColorTransformParams) GetAverageMaxrgbAddr() *AVRational {
	return (*AVRational)(&ctfp.average_maxrgb)
}

// GetVarianceMaxrgb gets `AVHDRVividColorTransformParams.variance_maxrgb` value.
func (ctfp *AVHDRVividColorTransformParams) GetVarianceMaxrgb() AVRational {
	return (AVRational)(ctfp.variance_maxrgb)
}

// SetVarianceMaxrgb sets `AVHDRVividColorTransformParams.variance_maxrgb` value.
func (ctfp *AVHDRVividColorTransformParams) SetVarianceMaxrgb(v AVRational) {
	ctfp.variance_maxrgb = (C.struct_AVRational)(v)
}

// GetVarianceMaxrgbAddr gets `AVHDRVividColorTransformParams.variance_maxrgb` address.
func (ctfp *AVHDRVividColorTransformParams) GetVarianceMaxrgbAddr() *AVRational {
	return (*AVRational)(&ctfp.variance_maxrgb)
}

// GetMaximumMaxrgb gets `AVHDRVividColorTransformParams.maximum_maxrgb` value.
func (ctfp *AVHDRVividColorTransformParams) GetMaximumMaxrgb() AVRational {
	return (AVRational)(ctfp.maximum_maxrgb)
}

// SetMaximumMaxrgb sets `AVHDRVividColorTransformParams.maximum_maxrgb` value.
func (ctfp *AVHDRVividColorTransformParams) SetMaximumMaxrgb(v AVRational) {
	ctfp.maximum_maxrgb = (C.struct_AVRational)(v)
}

// GetMaximumMaxrgbAddr gets `AVHDRVividColorTransformParams.maximum_maxrgb` address.
func (ctfp *AVHDRVividColorTransformParams) GetMaximumMaxrgbAddr() *AVRational {
	return (*AVRational)(&ctfp.maximum_maxrgb)
}

// GetToneMappingModeFlag gets `AVHDRVividColorTransformParams.tone_mapping_mode_flag` value.
func (ctfp *AVHDRVividColorTransformParams) GetToneMappingModeFlag() int32 {
	return (int32)(ctfp.tone_mapping_mode_flag)
}

// SetToneMappingModeFlag sets `AVHDRVividColorTransformParams.tone_mapping_mode_flag` value.
func (ctfp *AVHDRVividColorTransformParams) SetToneMappingModeFlag(v int32) {
	ctfp.tone_mapping_mode_flag = (C.int)(v)
}

// GetToneMappingModeFlagAddr gets `AVHDRVividColorTransformParams.tone_mapping_mode_flag` address.
func (ctfp *AVHDRVividColorTransformParams) GetToneMappingModeFlagAddr() *int32 {
	return (*int32)(&ctfp.tone_mapping_mode_flag)
}

// GetToneMappingParamNum gets `AVHDRVividColorTransformParams.tone_mapping_param_num` value.
func (ctfp *AVHDRVividColorTransformParams) GetToneMappingParamNum() int32 {
	return (int32)(ctfp.tone_mapping_param_num)
}

// SetToneMappingParamNum sets `AVHDRVividColorTransformParams.tone_mapping_param_num` value.
func (ctfp *AVHDRVividColorTransformParams) SetToneMappingParamNum(v int32) {
	ctfp.tone_mapping_param_num = (C.int)(v)
}

// GetToneMappingParamNumAddr gets `AVHDRVividColorTransformParams.tone_mapping_param_num` address.
func (ctfp *AVHDRVividColorTransformParams) GetToneMappingParamNumAddr() *int32 {
	return (*int32)(&ctfp.tone_mapping_param_num)
}

// GetTmParams gets `AVHDRVividColorTransformParams.tm_params` value.
func (ctfp *AVHDRVividColorTransformParams) GetTmParams() []AVHDRVividColorToneMappingParams {
	return unsafe.Slice((*AVHDRVividColorToneMappingParams)(&ctfp.tm_params[0]), 2)
}

// SetTmParams sets `AVHDRVividColorTransformParams.tm_params` value.
func (ctfp *AVHDRVividColorTransformParams) SetTmParams(v []AVHDRVividColorToneMappingParams) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		ctfp.tm_params[i] = (C.struct_AVHDRVividColorToneMappingParams)(v[i])
	}
}

// GetTmParamsAddr gets `AVHDRVividColorTransformParams.tm_params` address.
func (ctfp *AVHDRVividColorTransformParams) GetTmParamsAddr() **AVHDRVividColorToneMappingParams {
	return (**AVHDRVividColorToneMappingParams)(unsafe.Pointer(&ctfp.tm_params))
}

// GetColorSaturationMappingFlag gets `AVHDRVividColorTransformParams.color_saturation_mapping_flag` value.
func (ctfp *AVHDRVividColorTransformParams) GetColorSaturationMappingFlag() int32 {
	return (int32)(ctfp.color_saturation_mapping_flag)
}

// SetColorSaturationMappingFlag sets `AVHDRVividColorTransformParams.color_saturation_mapping_flag` value.
func (ctfp *AVHDRVividColorTransformParams) SetColorSaturationMappingFlag(v int32) {
	ctfp.color_saturation_mapping_flag = (C.int)(v)
}

// GetColorSaturationMappingFlagAddr gets `AVHDRVividColorTransformParams.color_saturation_mapping_flag` address.
func (ctfp *AVHDRVividColorTransformParams) GetColorSaturationMappingFlagAddr() *int32 {
	return (*int32)(&ctfp.color_saturation_mapping_flag)
}

// GetColorSaturationNum gets `AVHDRVividColorTransformParams.color_saturation_num` value.
func (ctfp *AVHDRVividColorTransformParams) GetColorSaturationNum() int32 {
	return (int32)(ctfp.color_saturation_num)
}

// SetColorSaturationNum sets `AVHDRVividColorTransformParams.color_saturation_num` value.
func (ctfp *AVHDRVividColorTransformParams) SetColorSaturationNum(v int32) {
	ctfp.color_saturation_num = (C.int)(v)
}

// GetColorSaturationNumAddr gets `AVHDRVividColorTransformParams.color_saturation_num` address.
func (ctfp *AVHDRVividColorTransformParams) GetColorSaturationNumAddr() *int32 {
	return (*int32)(&ctfp.color_saturation_num)
}

// GetColorSaturationGain gets `AVHDRVividColorTransformParams.color_saturation_gain` value.
func (ctfp *AVHDRVividColorTransformParams) GetColorSaturationGain() []AVRational {
	return unsafe.Slice((*AVRational)(&ctfp.color_saturation_gain[0]), 8)
}

// SetColorSaturationGain sets `AVHDRVividColorTransformParams.color_saturation_gain` value.
func (ctfp *AVHDRVividColorTransformParams) SetColorSaturationGain(v []AVRational) {
	for i := 0; i < FFMIN(len(v), 8); i++ {
		ctfp.color_saturation_gain[i] = (C.struct_AVRational)(v[i])
	}
}

// GetColorSaturationGainAddr gets `AVHDRVividColorTransformParams.color_saturation_gain` address.
func (ctfp *AVHDRVividColorTransformParams) GetColorSaturationGainAddr() **AVRational {
	return (**AVRational)(unsafe.Pointer(&ctfp.color_saturation_gain))
}

// AVDynamicHDRVivid
type AVDynamicHDRVivid = C.struct_AVDynamicHDRVivid

// GetSystemStartCode gets `AVDynamicHDRVivid.system_start_code` value.
func (dhv *AVDynamicHDRVivid) GetSystemStartCode() uint8 {
	return (uint8)(dhv.system_start_code)
}

// SetSystemStartCode sets `AVDynamicHDRVivid.system_start_code` value.
func (dhv *AVDynamicHDRVivid) SetSystemStartCode(v uint8) {
	dhv.system_start_code = (C.uint8_t)(v)
}

// GetSystemStartCodeAddr gets `AVDynamicHDRVivid.system_start_code` address.
func (dhv *AVDynamicHDRVivid) GetSystemStartCodeAddr() *uint8 {
	return (*uint8)(&dhv.system_start_code)
}

// GetNumWindows gets `AVDynamicHDRVivid.num_windows` value.
func (dhv *AVDynamicHDRVivid) GetNumWindows() uint8 {
	return (uint8)(dhv.num_windows)
}

// SetNumWindows sets `AVDynamicHDRVivid.num_windows` value.
func (dhv *AVDynamicHDRVivid) SetNumWindows(v uint8) {
	dhv.num_windows = (C.uint8_t)(v)
}

// GetNumWindowsAddr gets `AVDynamicHDRVivid.num_windows` address.
func (dhv *AVDynamicHDRVivid) GetNumWindowsAddr() *uint8 {
	return (*uint8)(&dhv.num_windows)
}

// GetParams gets `AVDynamicHDRVivid.params` value.
func (dhv *AVDynamicHDRVivid) GetParams() []AVHDRVividColorTransformParams {
	return unsafe.Slice((*AVHDRVividColorTransformParams)(&dhv.params[0]), 3)
}

// SetParams sets `AVDynamicHDRVivid.params` value.
func (dhv *AVDynamicHDRVivid) SetParams(v []AVHDRVividColorTransformParams) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		dhv.params[i] = (C.struct_AVHDRVividColorTransformParams)(v[i])
	}
}

// GetParamsAddr gets `AVDynamicHDRVivid.params` address.
func (dhv *AVDynamicHDRVivid) GetParamsAddr() **AVHDRVividColorTransformParams {
	return (**AVHDRVividColorTransformParams)(unsafe.Pointer(&dhv.params))
}

// AvDynamicHdrVividAlloc allocates an AVDynamicHDRVivid structure and set its fields to default values.
func AvDynamicHdrVividAlloc(size *uintptr) *AVDynamicHDRVivid {
	return (*AVDynamicHDRVivid)(C.av_dynamic_hdr_vivid_alloc((*C.size_t)(unsafe.Pointer(size))))
}

// AvDynamicHdrVividCreateSideData allocates a complete AVDynamicHDRVivid and add it to the frame.
func AvDynamicHdrVividCreateSideData(frame *AVFrame) *AVDynamicHDRVivid {
	return (*AVDynamicHDRVivid)(C.av_dynamic_hdr_vivid_create_side_data((*C.struct_AVFrame)(frame)))
}
