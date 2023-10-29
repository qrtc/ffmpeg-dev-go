// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/downmix_info.h>
*/
import "C"

// AVDownmixType
type AVDownmixType = C.enum_AVDownmixType

const (
	AV_DOWNMIX_TYPE_UNKNOWN = AVDownmixType(C.AV_DOWNMIX_TYPE_UNKNOWN)
	AV_DOWNMIX_TYPE_LORO    = AVDownmixType(C.AV_DOWNMIX_TYPE_LORO)
	AV_DOWNMIX_TYPE_LTRT    = AVDownmixType(C.AV_DOWNMIX_TYPE_LTRT)
	AV_DOWNMIX_TYPE_DPLII   = AVDownmixType(C.AV_DOWNMIX_TYPE_DPLII)
	AV_DOWNMIX_TYPE_NB      = AVDownmixType(C.AV_DOWNMIX_TYPE_NB)
)

// AVDownmixInfo
type AVDownmixInfo C.struct_AVDownmixInfo

// GetPreferredDownmixType gets `AVDownmixInfo.preferred_downmixtype` value.
func (di *AVDownmixInfo) GetPreferredDownmixType() AVDownmixType {
	return (AVDownmixType)(di.preferred_downmix_type)
}

// SetPreferredDownmixType sets `AVDownmixInfo.preferred_downmixtype` value.
func (di *AVDownmixInfo) SetPreferredDownmixType(v AVDownmixType) {
	di.preferred_downmix_type = (C.enum_AVDownmixType)(v)
}

// GetPreferredDownmixTypeAddr gets `AVDownmixInfo.preferred_downmixtype` address.
func (di *AVDownmixInfo) GetPreferredDownmixTypeAddr() *AVDownmixType {
	return (*AVDownmixType)(&di.preferred_downmix_type)
}

// GetCenterMixLevel gets `AVDownmixInfo.center_mix_level` value.
func (di *AVDownmixInfo) GetCenterMixLevel() float64 {
	return (float64)(di.center_mix_level)
}

// SetCenterMixLevel sets `AVDownmixInfo.center_mix_level` value.
func (di *AVDownmixInfo) SetCenterMixLevel(v float64) {
	di.center_mix_level = (C.double)(v)
}

// GetCenterMixLevelAddr gets `AVDownmixInfo.center_mix_level` address.
func (di *AVDownmixInfo) GetCenterMixLevelAddr() *float64 {
	return (*float64)(&di.center_mix_level)
}

// GetCenterMixLevelLtrt gets `AVDownmixInfo.center_mix_level_ltrt` value.
func (di *AVDownmixInfo) GetCenterMixLevelLtrt() float64 {
	return (float64)(di.center_mix_level_ltrt)
}

// SetCenterMixLevelLtrt sets `AVDownmixInfo.center_mix_level_ltrt` value.
func (di *AVDownmixInfo) SetCenterMixLevelLtrt(v float64) {
	di.center_mix_level_ltrt = (C.double)(v)
}

// GetCenterMixLevelLtrtAddr gets `AVDownmixInfo.center_mix_level_ltrt` address.
func (di *AVDownmixInfo) GetCenterMixLevelLtrtAddr() *float64 {
	return (*float64)(&di.center_mix_level_ltrt)
}

// GetSurroundMixLevel gets `AVDownmixInfo.surround_mix_level` value.
func (di *AVDownmixInfo) GetSurroundMixLevel() float64 {
	return (float64)(di.surround_mix_level)
}

// SetSurroundMixLevel sets `AVDownmixInfo.surround_mix_level` value.
func (di *AVDownmixInfo) SetSurroundMixLevel(v float64) {
	di.surround_mix_level = (C.double)(v)
}

// GetSurroundMixLevelAddr gets `AVDownmixInfo.surround_mix_level` address.
func (di *AVDownmixInfo) GetSurroundMixLevelAddr() *float64 {
	return (*float64)(&di.surround_mix_level)
}

// GetSurroundMixLevelLtrt gets `AVDownmixInfo.surround_mix_level_ltrt` value.
func (di *AVDownmixInfo) GetSurroundMixLevelLtrt() float64 {
	return (float64)(di.surround_mix_level_ltrt)
}

// SetSurroundMixLevelLtrt sets `AVDownmixInfo.surround_mix_level_ltrt` value.
func (di *AVDownmixInfo) SetSurroundMixLevelLtrt(v float64) {
	di.surround_mix_level_ltrt = (C.double)(v)
}

// GetSurroundMixLevelLtrtAddr gets `AVDownmixInfo.surround_mix_level_ltrt` address.
func (di *AVDownmixInfo) GetSurroundMixLevelLtrtAddr() *float64 {
	return (*float64)(&di.surround_mix_level_ltrt)
}

// GetLfeMixLevel gets `AVDownmixInfo.lfe_mix_level` value.
func (di *AVDownmixInfo) GetLfeMixLevel() float64 {
	return (float64)(di.lfe_mix_level)
}

// SetLfeMixLevel sets `AVDownmixInfo.lfe_mix_level` value.
func (di *AVDownmixInfo) SetLfeMixLevel(v float64) {
	di.lfe_mix_level = (C.double)(v)
}

// GetLfeMixLevelAddr gets `AVDownmixInfo.lfe_mix_level` address.
func (di *AVDownmixInfo) GetLfeMixLevelAddr() *float64 {
	return (*float64)(&di.lfe_mix_level)
}

// AvDownmixInfoUpdateSideData gets a frame's AV_FRAME_DATA_DOWNMIX_INFO side data for editing.
func AvDownmixInfoUpdateSideData(frame *AVFrame) *AVDownmixInfo {
	return (*AVDownmixInfo)(C.av_downmix_info_update_side_data((*C.struct_AVFrame)(frame)))
}
