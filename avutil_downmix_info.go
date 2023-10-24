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

// Custom: GetPreferredDownmixType gets `AVDownmixInfo.preferred_downmixtype` value.
func (di *AVDownmixInfo) GetPreferredDownmixType() AVDownmixType {
	return (AVDownmixType)(di.preferred_downmix_type)
}

// Custom: SetPreferredDownmixType sets `AVDownmixInfo.preferred_downmixtype` value.
func (di *AVDownmixInfo) SetPreferredDownmixType(v AVDownmixType) {
	di.preferred_downmix_type = (C.enum_AVDownmixType)(v)
}

// Custom: GetPreferredDownmixTypeAddr gets `AVDownmixInfo.preferred_downmixtype` address.
func (di *AVDownmixInfo) GetPreferredDownmixTypeAddr() *AVDownmixType {
	return (*AVDownmixType)(&di.preferred_downmix_type)
}

// Custom: GetCenterMixLevel gets `AVDownmixInfo.center_mix_level` value.
func (di *AVDownmixInfo) GetCenterMixLevel() float64 {
	return (float64)(di.center_mix_level)
}

// Custom: SetCenterMixLevel sets `AVDownmixInfo.center_mix_level` value.
func (di *AVDownmixInfo) SetCenterMixLevel(v float64) {
	di.center_mix_level = (C.double)(v)
}

// Custom: GetCenterMixLevelAddr gets `AVDownmixInfo.center_mix_level` address.
func (di *AVDownmixInfo) GetCenterMixLevelAddr() *float64 {
	return (*float64)(&di.center_mix_level)
}

// Custom: GetCenterMixLevelLtrt gets `AVDownmixInfo.center_mix_level_ltrt` value.
func (di *AVDownmixInfo) GetCenterMixLevelLtrt() float64 {
	return (float64)(di.center_mix_level_ltrt)
}

// Custom: SetCenterMixLevelLtrt sets `AVDownmixInfo.center_mix_level_ltrt` value.
func (di *AVDownmixInfo) SetCenterMixLevelLtrt(v float64) {
	di.center_mix_level_ltrt = (C.double)(v)
}

// Custom: GetCenterMixLevelLtrtAddr gets `AVDownmixInfo.center_mix_level_ltrt` address.
func (di *AVDownmixInfo) GetCenterMixLevelLtrtAddr() *float64 {
	return (*float64)(&di.center_mix_level_ltrt)
}

// Custom: GetSurroundMixLevel gets `AVDownmixInfo.surround_mix_level` value.
func (di *AVDownmixInfo) GetSurroundMixLevel() float64 {
	return (float64)(di.surround_mix_level)
}

// Custom: SetSurroundMixLevel sets `AVDownmixInfo.surround_mix_level` value.
func (di *AVDownmixInfo) SetSurroundMixLevel(v float64) {
	di.surround_mix_level = (C.double)(v)
}

// Custom: GetSurroundMixLevelAddr gets `AVDownmixInfo.surround_mix_level` address.
func (di *AVDownmixInfo) GetSurroundMixLevelAddr() *float64 {
	return (*float64)(&di.surround_mix_level)
}

// Custom: GetSurroundMixLevelLtrt gets `AVDownmixInfo.surround_mix_level_ltrt` value.
func (di *AVDownmixInfo) GetSurroundMixLevelLtrt() float64 {
	return (float64)(di.surround_mix_level_ltrt)
}

// Custom: SetSurroundMixLevelLtrt sets `AVDownmixInfo.surround_mix_level_ltrt` value.
func (di *AVDownmixInfo) SetSurroundMixLevelLtrt(v float64) {
	di.surround_mix_level_ltrt = (C.double)(v)
}

// Custom: GetSurroundMixLevelLtrtAddr gets `AVDownmixInfo.surround_mix_level_ltrt` address.
func (di *AVDownmixInfo) GetSurroundMixLevelLtrtAddr() *float64 {
	return (*float64)(&di.surround_mix_level_ltrt)
}

// Custom: GetLfeMixLevel gets `AVDownmixInfo.lfe_mix_level` value.
func (di *AVDownmixInfo) GetLfeMixLevel() float64 {
	return (float64)(di.lfe_mix_level)
}

// Custom: SetLfeMixLevel sets `AVDownmixInfo.lfe_mix_level` value.
func (di *AVDownmixInfo) SetLfeMixLevel(v float64) {
	di.lfe_mix_level = (C.double)(v)
}

// Custom: GetLfeMixLevelAddr gets `AVDownmixInfo.lfe_mix_level` address.
func (di *AVDownmixInfo) GetLfeMixLevelAddr() *float64 {
	return (*float64)(&di.lfe_mix_level)
}

// AvDownmixInfoUpdateSideData gets a frame's AV_FRAME_DATA_DOWNMIX_INFO side data for editing.
func AvDownmixInfoUpdateSideData(frame *AVFrame) *AVDownmixInfo {
	return (*AVDownmixInfo)(C.av_downmix_info_update_side_data((*C.struct_AVFrame)(frame)))
}
