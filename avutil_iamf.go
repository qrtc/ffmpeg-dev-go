// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/iamf.h>
*/
import "C"
import "unsafe"

// AVIAMFAnimationType
type AVIAMFAnimationType = C.enum_AVIAMFAnimationType

const (
	AV_IAMF_ANIMATION_TYPE_STEP   = AVIAMFAnimationType(C.AV_IAMF_ANIMATION_TYPE_STEP)
	AV_IAMF_ANIMATION_TYPE_LINEAR = AVIAMFAnimationType(C.AV_IAMF_ANIMATION_TYPE_LINEAR)
	AV_IAMF_ANIMATION_TYPE_BEZIER = AVIAMFAnimationType(C.AV_IAMF_ANIMATION_TYPE_BEZIER)
)

// AVIAMFMixGain
type AVIAMFMixGain C.struct_AVIAMFMixGain

// GetAvClass gets `AVIAMFMixGain.*av_class` value.
func (mxga *AVIAMFMixGain) GetAvClass() *AVClass {
	return (*AVClass)(mxga.av_class)
}

// SetAvClass sets `AVIAMFMixGain.*av_class` value.
func (mxga *AVIAMFMixGain) SetAvClass(v *AVClass) {
	mxga.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFMixGain.*av_class` address.
func (mxga *AVIAMFMixGain) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&mxga.av_class))
}

// GetSubblockDuration gets `AVIAMFMixGain.subblock_duration` value.
func (mxga *AVIAMFMixGain) GetSubblockDuration() uint32 {
	return (uint32)(mxga.subblock_duration)
}

// SetSubblockDuration sets `AVIAMFMixGain.subblock_duration` value.
func (mxga *AVIAMFMixGain) SetSubblockDuration(v uint32) {
	mxga.subblock_duration = (C.uint)(v)
}

// GetSubblockDurationAddr gets `AVIAMFMixGain.subblock_duration` address.
func (mxga *AVIAMFMixGain) GetSubblockDurationAddr() *uint32 {
	return (*uint32)(&mxga.subblock_duration)
}

// GetAnimationType gets `AVIAMFMixGain.animation_type` value.
func (mxga *AVIAMFMixGain) GetAnimationType() AVIAMFAnimationType {
	return (AVIAMFAnimationType)(mxga.animation_type)
}

// SetAnimationType sets `AVIAMFMixGain.animation_type` value.
func (mxga *AVIAMFMixGain) SetAnimationType(v AVIAMFAnimationType) {
	mxga.animation_type = (C.enum_AVIAMFAnimationType)(v)
}

// GetAnimationTypeAddr gets `AVIAMFMixGain.animation_type` address.
func (mxga *AVIAMFMixGain) GetAnimationTypeAddr() *AVIAMFAnimationType {
	return (*AVIAMFAnimationType)(&mxga.animation_type)
}

// GetStartPointValue gets `AVIAMFMixGain.start_point_value` value.
func (mxga *AVIAMFMixGain) GetStartPointValue() AVRational {
	return (AVRational)(mxga.start_point_value)
}

// SetStartPointValue sets `AVIAMFMixGain.start_point_value` value.
func (mxga *AVIAMFMixGain) SetStartPointValue(v AVRational) {
	mxga.start_point_value = (C.struct_AVRational)(v)
}

// GetStartPointValueAddr gets `AVIAMFMixGain.start_point_value` address.
func (mxga *AVIAMFMixGain) GetStartPointValueAddr() *AVRational {
	return (*AVRational)(&mxga.start_point_value)
}

// GetEndPointValue gets `AVIAMFMixGain.end_point_value` value.
func (mxga *AVIAMFMixGain) GetEndPointValue() AVRational {
	return (AVRational)(mxga.end_point_value)
}

// SetEndPointValue sets `AVIAMFMixGain.end_point_value` value.
func (mxga *AVIAMFMixGain) SetEndPointValue(v AVRational) {
	mxga.end_point_value = (C.struct_AVRational)(v)
}

// GetEndPointValueAddr gets `AVIAMFMixGain.end_point_value` address.
func (mxga *AVIAMFMixGain) GetEndPointValueAddr() *AVRational {
	return (*AVRational)(&mxga.end_point_value)
}

// GetControlPointValue gets `AVIAMFMixGain.control_point_value` value.
func (mxga *AVIAMFMixGain) GetControlPointValue() AVRational {
	return (AVRational)(mxga.control_point_value)
}

// SetControlPointValue sets `AVIAMFMixGain.control_point_value` value.
func (mxga *AVIAMFMixGain) SetControlPointValue(v AVRational) {
	mxga.control_point_value = (C.struct_AVRational)(v)
}

// GetControlPointValueAddr gets `AVIAMFMixGain.control_point_value` address.
func (mxga *AVIAMFMixGain) GetControlPointValueAddr() *AVRational {
	return (*AVRational)(&mxga.control_point_value)
}

// GetControlPointRelativeTime gets `AVIAMFMixGain.control_point_relative_time` value.
func (mxga *AVIAMFMixGain) GetControlPointRelativeTime() AVRational {
	return (AVRational)(mxga.control_point_relative_time)
}

// SetControlPointRelativeTime sets `AVIAMFMixGain.control_point_relative_time` value.
func (mxga *AVIAMFMixGain) SetControlPointRelativeTime(v AVRational) {
	mxga.control_point_relative_time = (C.struct_AVRational)(v)
}

// GetControlPointRelativeTimeAddr gets `AVIAMFMixGain.control_point_relative_time` address.
func (mxga *AVIAMFMixGain) GetControlPointRelativeTimeAddr() *AVRational {
	return (*AVRational)(&mxga.control_point_relative_time)
}

// AVIAMFDemixingInfo
type AVIAMFDemixingInfo C.struct_AVIAMFDemixingInfo

// GetAvClass gets `AVIAMFDemixingInfo.av_class` value.
func (dmif *AVIAMFDemixingInfo) GetAvClass() *AVClass {
	return (*AVClass)(dmif.av_class)
}

// SetAvClass sets `AVIAMFDemixingInfo.av_class` value.
func (dmif *AVIAMFDemixingInfo) SetAvClass(v *AVClass) {
	dmif.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFDemixingInfo.av_class` address.
func (dmif *AVIAMFDemixingInfo) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&dmif.av_class))
}

// GetSubblockDuration gets `AVIAMFDemixingInfo.subblock_duration` value.
func (dmif *AVIAMFDemixingInfo) SubblockDurationGet() uint32 {
	return (uint32)(dmif.subblock_duration)
}

// SetSubblockDuration sets `AVIAMFDemixingInfo.subblock_duration` value.
func (dmif *AVIAMFDemixingInfo) SetSubblockDuration(v uint32) {
	dmif.subblock_duration = (C.uint)(v)
}

// GetSubblockDurationAddr gets `AVIAMFDemixingInfo.subblock_duration` address.
func (dmif *AVIAMFDemixingInfo) GetSubblockDurationAddr() *uint32 {
	return (*uint32)(&dmif.subblock_duration)
}

// GetDmixpMode gets `AVIAMFDemixingInfo.dmixp_mode` value.
func (dmif *AVIAMFDemixingInfo) DmixpModeGet() uint32 {
	return (uint32)(dmif.dmixp_mode)
}

// SetDmixpMode sets `AVIAMFDemixingInfo.dmixp_mode` value.
func (dmif *AVIAMFDemixingInfo) SetDmixpMode(v uint32) {
	dmif.dmixp_mode = (C.uint)(v)
}

// GetDmixpModeAddr gets `AVIAMFDemixingInfo.dmixp_mode` address.
func (dmif *AVIAMFDemixingInfo) GetDmixpModeAddr() *uint32 {
	return (*uint32)(&dmif.dmixp_mode)
}

// AVIAMFReconGain
type AVIAMFReconGain C.struct_AVIAMFReconGain

// GetAvClass gets `AVIAMFReconGain.av_class` value.
func (rcgn *AVIAMFReconGain) GetAvClass() *AVClass {
	return (*AVClass)(rcgn.av_class)
}

// SetAvClass sets `AVIAMFReconGain.av_class` value.
func (rcgn *AVIAMFReconGain) SetAvClass(v *AVClass) {
	rcgn.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFReconGain.av_class` address.
func (rcgn *AVIAMFReconGain) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&rcgn.av_class))
}

// GetSubblockDuration gets `AVIAMFReconGain.subblock_duration` value.
func (rcgn *AVIAMFReconGain) GetSubblockDuration() uint32 {
	return (uint32)(rcgn.subblock_duration)
}

// SetSubblockDuration sets `AVIAMFReconGain.subblock_duration` value.
func (rcgn *AVIAMFReconGain) SetSubblockDuration(v uint32) {
	rcgn.subblock_duration = (C.uint)(v)
}

// GetSubblockDurationAddr gets `AVIAMFReconGain.subblock_duration` address.
func (rcgn *AVIAMFReconGain) GetSubblockDurationAddr() *uint32 {
	return (*uint32)(&rcgn.subblock_duration)
}

// GetReconGain gets `AVIAMFReconGain.recon_gain` value.
func (rcgn *AVIAMFReconGain) GetReconGain() (v [][]uint8) {
	for i := 0; i < 6; i++ {
		v = append(v, unsafe.Slice((*uint8)(&rcgn.recon_gain[i][0]), 12))
	}
	return v
}

// SetReconGain sets `AVIAMFReconGain.recon_gain` value.
func (rcgn *AVIAMFReconGain) SetReconGain(v [][]uint8) {
	for i := 0; i < FFMIN(len(v), 6); i++ {
		for j := 0; j < FFMIN(len(v[i]), 12); j++ {
			rcgn.recon_gain[i][j] = (C.uint8_t)(v[i][j])
		}
	}
}

// GetReconGainAddr gets `AVIAMFReconGain.recon_gain` address.
func (rcgn *AVIAMFReconGain) GetReconGainAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&rcgn.recon_gain))
}

// AVIAMFParamDefinitionType
type AVIAMFParamDefinitionType = C.enum_AVIAMFParamDefinitionType

const (
	AV_IAMF_PARAMETER_DEFINITION_MIX_GAIN   = AVIAMFParamDefinitionType(C.AV_IAMF_PARAMETER_DEFINITION_MIX_GAIN)
	AV_IAMF_PARAMETER_DEFINITION_DEMIXING   = AVIAMFParamDefinitionType(C.AV_IAMF_PARAMETER_DEFINITION_DEMIXING)
	AV_IAMF_PARAMETER_DEFINITION_RECON_GAIN = AVIAMFParamDefinitionType(C.AV_IAMF_PARAMETER_DEFINITION_RECON_GAIN)
)

// AVIAMFParamDefinition
type AVIAMFParamDefinition C.struct_AVIAMFParamDefinition

// GetAvClass gets `AVIAMFParamDefinition.av_class` value.
func (pdef *AVIAMFParamDefinition) GetAvClass() *AVClass {
	return (*AVClass)(pdef.av_class)
}

// SetAvClass sets `AVIAMFParamDefinition.av_class` value.
func (pdef *AVIAMFParamDefinition) SetAvClass(v *AVClass) {
	pdef.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFParamDefinition.av_class` address.
func (pdef *AVIAMFParamDefinition) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&pdef.av_class))
}

// GetSubblocksOffset gets `AVIAMFParamDefinition.subblocks_offset` value.
func (pdef *AVIAMFParamDefinition) SubblocksOffsetGet() uintptr {
	return (uintptr)(pdef.subblocks_offset)
}

// SetSubblocksOffset sets `AVIAMFParamDefinition.subblocks_offset` value.
func (pdef *AVIAMFParamDefinition) SetSubblocksOffset(v uintptr) {
	pdef.subblocks_offset = (C.size_t)(v)
}

// GetSubblocksOffsetAddr gets `AVIAMFParamDefinition.subblocks_offset` address.
func (pdef *AVIAMFParamDefinition) GetSubblocksOffsetAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&pdef.subblocks_offset))
}

// GetSubblockSize gets `AVIAMFParamDefinition.subblock_size` value.
func (pdef *AVIAMFParamDefinition) SubblockSizeGet() uintptr {
	return (uintptr)(pdef.subblock_size)
}

// SetSubblockSize sets `AVIAMFParamDefinition.subblock_size` value.
func (pdef *AVIAMFParamDefinition) SetSubblockSize(v uintptr) {
	pdef.subblock_size = (C.size_t)(v)
}

// GetSubblockSizeAddr gets `AVIAMFParamDefinition.subblock_size` address.
func (pdef *AVIAMFParamDefinition) GetSubblockSizeAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&pdef.subblock_size))
}

// GetNbSubblocks gets `AVIAMFParamDefinition.nb_subblocks` value.
func (pdef *AVIAMFParamDefinition) NbSubblocksGet() uint32 {
	return (uint32)(pdef.nb_subblocks)
}

// SetNbSubblocks sets `AVIAMFParamDefinition.nb_subblocks` value.
func (pdef *AVIAMFParamDefinition) SetNbSubblocks(v uint32) {
	pdef.nb_subblocks = (C.uint)(v)
}

// GetNbSubblocksAddr gets `AVIAMFParamDefinition.nb_subblocks` address.
func (pdef *AVIAMFParamDefinition) GetNbSubblocksAddr() *uint32 {
	return (*uint32)(&pdef.nb_subblocks)
}

// GetType gets `AVIAMFParamDefinition.type` value.
func (pdef *AVIAMFParamDefinition) TypeGet() AVIAMFParamDefinitionType {
	return (AVIAMFParamDefinitionType)(pdef._type)
}

// SetType sets `AVIAMFParamDefinition.type` value.
func (pdef *AVIAMFParamDefinition) SetType(v AVIAMFParamDefinitionType) {
	pdef._type = (C.enum_AVIAMFParamDefinitionType)(v)
}

// GetTypeAddr gets `AVIAMFParamDefinition.type` address.
func (pdef *AVIAMFParamDefinition) GetTypeAddr() *AVIAMFParamDefinitionType {
	return (*AVIAMFParamDefinitionType)(&pdef._type)
}

// GetParameterId gets `AVIAMFParamDefinition.parameter_id` value.
func (pdef *AVIAMFParamDefinition) ParameterIdGet() uint32 {
	return (uint32)(pdef.parameter_id)
}

// SetParameterId sets `AVIAMFParamDefinition.parameter_id` value.
func (pdef *AVIAMFParamDefinition) SetParameterId(v uint32) {
	pdef.parameter_id = (C.uint)(v)
}

// GetParameterIdAddr gets `AVIAMFParamDefinition.parameter_id` address.
func (pdef *AVIAMFParamDefinition) GetParameterIdAddr() *uint32 {
	return (*uint32)(&pdef.parameter_id)
}

// GetParameterRate gets `AVIAMFParamDefinition.parameter_rate` value.
func (pdef *AVIAMFParamDefinition) ParameterRateGet() uint32 {
	return (uint32)(pdef.parameter_rate)
}

// SetParameterRate sets `AVIAMFParamDefinition.parameter_rate` value.
func (pdef *AVIAMFParamDefinition) SetParameterRate(v uint32) {
	pdef.parameter_rate = (C.uint)(v)
}

// GetParameterRateAddr gets `AVIAMFParamDefinition.parameter_rate` address.
func (pdef *AVIAMFParamDefinition) GetParameterRateAddr() *uint32 {
	return (*uint32)(&pdef.parameter_rate)
}

// GetDuration gets `AVIAMFParamDefinition.duration` value.
func (pdef *AVIAMFParamDefinition) DurationGet() uint32 {
	return (uint32)(pdef.duration)
}

// SetDuration sets `AVIAMFParamDefinition.duration` value.
func (pdef *AVIAMFParamDefinition) SetDuration(v uint32) {
	pdef.duration = (C.uint)(v)
}

// GetDurationAddr gets `AVIAMFParamDefinition.duration` address.
func (pdef *AVIAMFParamDefinition) GetDurationAddr() *uint32 {
	return (*uint32)(&pdef.duration)
}

// GetConstantSubblockDuration gets `AVIAMFParamDefinition.constant_subblock_duration` value.
func (pdef *AVIAMFParamDefinition) ConstantSubblockDurationGet() uint32 {
	return (uint32)(pdef.constant_subblock_duration)
}

// SetConstantSubblockDuration sets `AVIAMFParamDefinition.constant_subblock_duration` value.
func (pdef *AVIAMFParamDefinition) SetConstantSubblockDuration(v uint32) {
	pdef.constant_subblock_duration = (C.uint)(v)
}

// GetConstantSubblockDurationAddr gets `AVIAMFParamDefinition.constant_subblock_duration` address.
func (pdef *AVIAMFParamDefinition) GetConstantSubblockDurationAddr() *uint32 {
	return (*uint32)(&pdef.constant_subblock_duration)
}

// AvIamfParamDefinitionGetClass
func AvIamfParamDefinitionGetClass() *AVClass {
	return (*AVClass)(C.av_iamf_param_definition_get_class())
}

// AvIamfParamDefinitionAlloc allocates memory for AVIAMFParamDefinition.
func AvIamfParamDefinitionAlloc(_type AVIAMFParamDefinitionType,
	nbSubblocks uint32, size *uintptr) *AVIAMFParamDefinition {
	return (*AVIAMFParamDefinition)(C.av_iamf_param_definition_alloc(
		(C.enum_AVIAMFParamDefinitionType)(_type),
		(C.uint)(nbSubblocks),
		(*C.size_t)(unsafe.Pointer(size))))
}

// AvIamfParamDefinitionGetSubblock gets the subblock at the specified idx.
func AvIamfParamDefinitionGetSubblock(par *AVIAMFParamDefinition, idx uint32) unsafe.Pointer {
	return VoidPointer(C.av_iamf_param_definition_get_subblock(
		(*C.struct_AVIAMFParamDefinition)(par),
		(C.uint)(idx)))
}

// AVIAMFAmbisonicsMode
type AVIAMFAmbisonicsMode = C.enum_AVIAMFAmbisonicsMode

const (
	AV_IAMF_AMBISONICS_MODE_MONO       = AVIAMFAmbisonicsMode(C.AV_IAMF_AMBISONICS_MODE_MONO)
	AV_IAMF_AMBISONICS_MODE_PROJECTION = AVIAMFAmbisonicsMode(C.AV_IAMF_AMBISONICS_MODE_PROJECTION)
)

const (
	AV_IAMF_LAYER_FLAG_RECON_GAIN = C.AV_IAMF_LAYER_FLAG_RECON_GAIN
)

// AVIAMFLayer
type AVIAMFLayer C.struct_AVIAMFLayer

// GetAvClass gets `AVIAMFLayer.av_class` value.
func (layer *AVIAMFLayer) GetAvClass() *AVClass {
	return (*AVClass)(layer.av_class)
}

// SetAvClass sets `AVIAMFLayer.av_class` value.
func (layer *AVIAMFLayer) SetAvClass(v *AVClass) {
	layer.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFLayer.av_class` address.
func (layer *AVIAMFLayer) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&layer.av_class))
}

// GetChLayout gets `AVIAMFLayer.ch_layout` value.
func (layer *AVIAMFLayer) GetChLayout() AVChannelLayout {
	return (AVChannelLayout)(layer.ch_layout)
}

// SetChLayout sets `AVIAMFLayer.ch_layout` value.
func (layer *AVIAMFLayer) SetChLayout(v AVChannelLayout) {
	layer.ch_layout = (C.struct_AVChannelLayout)(v)
}

// GetChLayoutAddr gets `AVIAMFLayer.ch_layout` address.
func (layer *AVIAMFLayer) GetChLayoutAddr() *AVChannelLayout {
	return (*AVChannelLayout)(&layer.ch_layout)
}

// GetFlags gets `AVIAMFLayer.flags` value.
func (layer *AVIAMFLayer) GetFlags() uint32 {
	return (uint32)(layer.flags)
}

// SetFlags sets `AVIAMFLayer.flags` value.
func (layer *AVIAMFLayer) SetFlags(v uint32) {
	layer.flags = (C.uint)(v)
}

// GetFlagsAddr gets `AVIAMFLayer.flags` address.
func (layer *AVIAMFLayer) GetFlagsAddr() *uint32 {
	return (*uint32)(&layer.flags)
}

// GetOutputGainFlags gets `AVIAMFLayer.output_gain_flags` value.
func (layer *AVIAMFLayer) GetOutputGainFlags() uint32 {
	return (uint32)(layer.output_gain_flags)
}

// SetOutputGainFlags sets `AVIAMFLayer.output_gain_flags` value.
func (layer *AVIAMFLayer) SetOutputGainFlags(v uint32) {
	layer.output_gain_flags = (C.uint)(v)
}

// GetOutputGainFlagsAddr gets `AVIAMFLayer.output_gain_flags` address.
func (layer *AVIAMFLayer) GetOutputGainFlagsAddr() *uint32 {
	return (*uint32)(&layer.output_gain_flags)
}

// GetOutputGain gets `AVIAMFLayer.output_gain` value.
func (layer *AVIAMFLayer) GetOutputGain() AVRational {
	return (AVRational)(layer.output_gain)
}

// SetOutputGain sets `AVIAMFLayer.output_gain` value.
func (layer *AVIAMFLayer) SetOutputGain(v AVRational) {
	layer.output_gain = (C.struct_AVRational)(v)
}

// GetOutputGainAddr gets `AVIAMFLayer.output_gain` address.
func (layer *AVIAMFLayer) GetOutputGainAddr() *AVRational {
	return (*AVRational)(&layer.output_gain)
}

// GetAmbisonicsMode gets `AVIAMFLayer.ambisonics_mode` value.
func (layer *AVIAMFLayer) GetAmbisonicsMode() AVIAMFAmbisonicsMode {
	return (AVIAMFAmbisonicsMode)(layer.ambisonics_mode)
}

// SetAmbisonicsMode sets `AVIAMFLayer.ambisonics_mode` value.
func (layer *AVIAMFLayer) SetAmbisonicsMode(v AVIAMFAmbisonicsMode) {
	layer.ambisonics_mode = (C.enum_AVIAMFAmbisonicsMode)(v)
}

// GetAmbisonicsModeAddr gets `AVIAMFLayer.ambisonics_mode` address.
func (layer *AVIAMFLayer) GetAmbisonicsModeAddr() *AVIAMFAmbisonicsMode {
	return (*AVIAMFAmbisonicsMode)(&layer.ambisonics_mode)
}

// GetDemixingMatrix gets `AVIAMFLayer.demixing_matrix` value.
func (layer *AVIAMFLayer) GetDemixingMatrix() *AVRational {
	return (*AVRational)(layer.demixing_matrix)
}

// SetDemixingMatrix sets `AVIAMFLayer.demixing_matrix` value.
func (layer *AVIAMFLayer) SetDemixingMatrix(v *AVRational) {
	layer.demixing_matrix = (*C.struct_AVRational)(v)
}

// GetDemixingMatrixAddr gets `AVIAMFLayer.demixing_matrix` address.
func (layer *AVIAMFLayer) GetDemixingMatrixAddr() **AVRational {
	return (**AVRational)(unsafe.Pointer(&layer.demixing_matrix))
}

// AVIAMFAudioElementType
type AVIAMFAudioElementType = C.enum_AVIAMFAudioElementType

const (
	AV_IAMF_AUDIO_ELEMENT_TYPE_CHANNEL = AVIAMFAudioElementType(C.AV_IAMF_AUDIO_ELEMENT_TYPE_CHANNEL)
	AV_IAMF_AUDIO_ELEMENT_TYPE_SCENE   = AVIAMFAudioElementType(C.AV_IAMF_AUDIO_ELEMENT_TYPE_SCENE)
)

// AVIAMFAudioElement
type AVIAMFAudioElement C.struct_AVIAMFAudioElement

// GetAvClass gets `AVIAMFAudioElement.av_class` value.
func (aelm *AVIAMFAudioElement) GetAvClass() *AVClass {
	return (*AVClass)(aelm.av_class)
}

// SetAvClass sets `AVIAMFAudioElement.av_class` value.
func (aelm *AVIAMFAudioElement) SetAvClass(v *AVClass) {
	aelm.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFAudioElement.av_class` address.
func (aelm *AVIAMFAudioElement) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&aelm.av_class))
}

// GetLayers gets `AVIAMFAudioElement.layers` value.
func (aelm *AVIAMFAudioElement) GetLayers() []*AVIAMFLayer {
	return unsafe.Slice((**AVIAMFLayer)(unsafe.Pointer(aelm.layers)), aelm.nb_layers)
}

// SetLayers sets `AVIAMFAudioElement.layers` value.
func (aelm *AVIAMFAudioElement) SetLayers(v **AVIAMFLayer) {
	aelm.layers = (**C.struct_AVIAMFLayer)(unsafe.Pointer(v))
}

// GetLayersAddr gets `AVIAMFAudioElement.layers` address.
func (aelm *AVIAMFAudioElement) GetLayersAddr() ***AVIAMFLayer {
	return (***AVIAMFLayer)(unsafe.Pointer(&aelm.layers))
}

// GetNbLayers gets `AVIAMFAudioElement.nb_layers` value.
func (aelm *AVIAMFAudioElement) GetNbLayers() uint32 {
	return (uint32)(aelm.nb_layers)
}

// SetNbLayers sets `AVIAMFAudioElement.nb_layers` value.
func (aelm *AVIAMFAudioElement) SetNbLayers(v uint32) {
	aelm.nb_layers = (C.uint)(v)
}

// GetNbLayersAddr gets `AVIAMFAudioElement.nb_layers` address.
func (aelm *AVIAMFAudioElement) GetNbLayersAddr() *uint32 {
	return (*uint32)(&aelm.nb_layers)
}

// GetDemixingInfo gets `AVIAMFAudioElement.demixing_info` value.
func (aelm *AVIAMFAudioElement) GetDemixingInfo() *AVIAMFParamDefinition {
	return (*AVIAMFParamDefinition)(aelm.demixing_info)
}

// SetDemixingInfo sets `AVIAMFAudioElement.demixing_info` value.
func (aelm *AVIAMFAudioElement) SetDemixingInfo(v *AVIAMFParamDefinition) {
	aelm.demixing_info = (*C.struct_AVIAMFParamDefinition)(v)
}

// GetDemixingInfoAddr gets `AVIAMFAudioElement.demixing_info` address.
func (aelm *AVIAMFAudioElement) GetDemixingInfoAddr() **AVIAMFParamDefinition {
	return (**AVIAMFParamDefinition)(unsafe.Pointer(&aelm.demixing_info))
}

// GetReconGainInfo gets `AVIAMFAudioElement.recon_gain_info` value.
func (aelm *AVIAMFAudioElement) GetReconGainInfo() *AVIAMFParamDefinition {
	return (*AVIAMFParamDefinition)(aelm.recon_gain_info)
}

// SetReconGainInfo sets `AVIAMFAudioElement.recon_gain_info` value.
func (aelm *AVIAMFAudioElement) SetReconGainInfo(v *AVIAMFParamDefinition) {
	aelm.recon_gain_info = (*C.struct_AVIAMFParamDefinition)(v)
}

// GetReconGainInfoAddr gets `AVIAMFAudioElement.recon_gain_info` address.
func (aelm *AVIAMFAudioElement) GetReconGainInfoAddr() **AVIAMFParamDefinition {
	return (**AVIAMFParamDefinition)(unsafe.Pointer(&aelm.recon_gain_info))
}

// GetAudioElementType gets `AVIAMFAudioElement.audio_element_type` value.
func (aelm *AVIAMFAudioElement) GetAudioElementType() AVIAMFAudioElementType {
	return (AVIAMFAudioElementType)(aelm.audio_element_type)
}

// SetAudioElementType sets `AVIAMFAudioElement.audio_element_type` value.
func (aelm *AVIAMFAudioElement) SetAudioElementType(v AVIAMFAudioElementType) {
	aelm.audio_element_type = (C.enum_AVIAMFAudioElementType)(v)
}

// GetAudioElementTypeAddr gets `AVIAMFAudioElement.audio_element_type` address.
func (aelm *AVIAMFAudioElement) GetAudioElementTypeAddr() *AVIAMFAudioElementType {
	return (*AVIAMFAudioElementType)(&aelm.audio_element_type)
}

// GetDefaultW gets `AVIAMFAudioElement.default_w` value.
func (aelm *AVIAMFAudioElement) GetDefaultW() uint32 {
	return (uint32)(aelm.default_w)
}

// SetDefaultW sets `AVIAMFAudioElement.default_w` value.
func (aelm *AVIAMFAudioElement) SetDefaultW(v uint32) {
	aelm.default_w = (C.uint)(v)
}

// GetDefaultWAddr gets `AVIAMFAudioElement.default_w` address.
func (aelm *AVIAMFAudioElement) GetDefaultWAddr() *uint32 {
	return (*uint32)(&aelm.default_w)
}

// AvIamfAudioElementGetClass
func AvIamfAudioElementGetClass() *AVClass {
	return (*AVClass)(C.av_iamf_audio_element_get_class())
}

// AvIamfAudioElementAlloc
func AvIamfAudioElementAlloc() *AVIAMFAudioElement {
	return (*AVIAMFAudioElement)(C.av_iamf_audio_element_alloc())
}

// AvIamfAudioElementAddLayer
func AvIamfAudioElementAddLayer(audioElement *AVIAMFAudioElement) *AVIAMFLayer {
	return (*AVIAMFLayer)(C.av_iamf_audio_element_add_layer((*C.struct_AVIAMFAudioElement)(audioElement)))
}

// AvIamfAudioElementFree
func AvIamfAudioElementFree(audioElement **AVIAMFAudioElement) {
	C.av_iamf_audio_element_free((**C.struct_AVIAMFAudioElement)(unsafe.Pointer(audioElement)))
}

// AVIAMFHeadphonesMode
type AVIAMFHeadphonesMode = C.enum_AVIAMFHeadphonesMode

const (
	AV_IAMF_HEADPHONES_MODE_STEREO   = AVIAMFHeadphonesMode(C.AV_IAMF_HEADPHONES_MODE_STEREO)
	AV_IAMF_HEADPHONES_MODE_BINAURAL = AVIAMFHeadphonesMode(C.AV_IAMF_HEADPHONES_MODE_BINAURAL)
)

// AVIAMFSubmixElement
type AVIAMFSubmixElement C.struct_AVIAMFSubmixElement

// GetAvClass gets `AVIAMFSubmixElement.av_class` value.
func (smel *AVIAMFSubmixElement) GetAvClass() *AVClass {
	return (*AVClass)(smel.av_class)
}

// SetAvClass sets `AVIAMFSubmixElement.av_class` value.
func (smel *AVIAMFSubmixElement) SetAvClass(v *AVClass) {
	smel.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFSubmixElement.av_class` address.
func (smel *AVIAMFSubmixElement) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&smel.av_class))
}

// GetAudioElementId gets `AVIAMFSubmixElement.audio_element_id` value.
func (smel *AVIAMFSubmixElement) GetAudioElementId() uint32 {
	return (uint32)(smel.audio_element_id)
}

// SetAudioElementId sets `AVIAMFSubmixElement.audio_element_id` value.
func (smel *AVIAMFSubmixElement) SetAudioElementId(v uint32) {
	smel.audio_element_id = (C.uint)(v)
}

// GetAudioElementIdAddr gets `AVIAMFSubmixElement.audio_element_id` address.
func (smel *AVIAMFSubmixElement) GetAudioElementIdAddr() *uint32 {
	return (*uint32)(&smel.audio_element_id)
}

// GetElementMixConfig gets `AVIAMFSubmixElement.element_mix_config` value.
func (smel *AVIAMFSubmixElement) GetElementMixConfig() *AVIAMFParamDefinition {
	return (*AVIAMFParamDefinition)(smel.element_mix_config)
}

// SetElementMixConfig sets `AVIAMFSubmixElement.element_mix_config` value.
func (smel *AVIAMFSubmixElement) SetElementMixConfig(v *AVIAMFParamDefinition) {
	smel.element_mix_config = (*C.struct_AVIAMFParamDefinition)(v)
}

// GetElementMixConfigAddr gets `AVIAMFSubmixElement.element_mix_config` address.
func (smel *AVIAMFSubmixElement) GetElementMixConfigAddr() **AVIAMFParamDefinition {
	return (**AVIAMFParamDefinition)(unsafe.Pointer(&smel.element_mix_config))
}

// GetDefaultMixGain gets `AVIAMFSubmixElement.default_mix_gain` value.
func (smel *AVIAMFSubmixElement) GetDefaultMixGain() AVRational {
	return (AVRational)(smel.default_mix_gain)
}

// SetDefaultMixGain sets `AVIAMFSubmixElement.default_mix_gain` value.
func (smel *AVIAMFSubmixElement) SetDefaultMixGain(v AVRational) {
	smel.default_mix_gain = (C.struct_AVRational)(v)
}

// GetDefaultMixGainAddr gets `AVIAMFSubmixElement.default_mix_gain` address.
func (smel *AVIAMFSubmixElement) GetDefaultMixGainAddr() *AVRational {
	return (*AVRational)(&smel.default_mix_gain)
}

// GetHeadphonesRenderingMode gets `AVIAMFSubmixElement.headphones_rendering_mode` value.
func (smel *AVIAMFSubmixElement) GetHeadphonesRenderingMode() AVIAMFHeadphonesMode {
	return (AVIAMFHeadphonesMode)(smel.headphones_rendering_mode)
}

// SetHeadphonesRenderingMode sets `AVIAMFSubmixElement.headphones_rendering_mode` value.
func (smel *AVIAMFSubmixElement) SetHeadphonesRenderingMode(v AVIAMFHeadphonesMode) {
	smel.headphones_rendering_mode = (C.enum_AVIAMFHeadphonesMode)(v)
}

// GetHeadphonesRenderingModeAddr gets `AVIAMFSubmixElement.headphones_rendering_mode` address.
func (smel *AVIAMFSubmixElement) GetHeadphonesRenderingModeAddr() *AVIAMFHeadphonesMode {
	return (*AVIAMFHeadphonesMode)(&smel.headphones_rendering_mode)
}

// GetAnnotations gets `AVIAMFSubmixElement.annotations` value.
func (smel *AVIAMFSubmixElement) GetAnnotations() *AVDictionary {
	return (*AVDictionary)(smel.annotations)
}

// SetAnnotations sets `AVIAMFSubmixElement.annotations` value.
func (smel *AVIAMFSubmixElement) SetAnnotations(v *AVDictionary) {
	smel.annotations = (*C.struct_AVDictionary)(v)
}

// GetAnnotationsAddr gets `AVIAMFSubmixElement.annotations` address.
func (smel *AVIAMFSubmixElement) GetAnnotationsAddr() **AVDictionary {
	return (**AVDictionary)(unsafe.Pointer(&smel.annotations))
}

// AVIAMFSubmixLayoutType
type AVIAMFSubmixLayoutType = C.enum_AVIAMFSubmixLayoutType

const (
	AV_IAMF_SUBMIX_LAYOUT_TYPE_LOUDSPEAKERS = AVIAMFSubmixLayoutType(C.AV_IAMF_SUBMIX_LAYOUT_TYPE_LOUDSPEAKERS)
	AV_IAMF_SUBMIX_LAYOUT_TYPE_BINAURAL     = AVIAMFSubmixLayoutType(C.AV_IAMF_SUBMIX_LAYOUT_TYPE_BINAURAL)
)

// AVIAMFSubmixLayout
type AVIAMFSubmixLayout C.struct_AVIAMFSubmixLayout

// GetAvClass gets `AVIAMFSubmixLayout.av_class` value.
func (smlo *AVIAMFSubmixLayout) GetAvClass() *AVClass {
	return (*AVClass)(smlo.av_class)
}

// SetAvClass sets `AVIAMFSubmixLayout.av_class` value.
func (smlo *AVIAMFSubmixLayout) SetAvClass(v *AVClass) {
	smlo.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFSubmixLayout.av_class` address.
func (smlo *AVIAMFSubmixLayout) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&smlo.av_class))
}

// GetLayoutType gets `AVIAMFSubmixLayout.layout_type` value.
func (smlo *AVIAMFSubmixLayout) GetLayoutType() AVIAMFSubmixLayoutType {
	return (AVIAMFSubmixLayoutType)(smlo.layout_type)
}

// SetLayoutType sets `AVIAMFSubmixLayout.layout_type` value.
func (smlo *AVIAMFSubmixLayout) SetLayoutType(v AVIAMFSubmixLayoutType) {
	smlo.layout_type = (C.enum_AVIAMFSubmixLayoutType)(v)
}

// GetLayoutTypeAddr gets `AVIAMFSubmixLayout.layout_type` address.
func (smlo *AVIAMFSubmixLayout) GetLayoutTypeAddr() *AVIAMFSubmixLayoutType {
	return (*AVIAMFSubmixLayoutType)(&smlo.layout_type)
}

// GetSoundSystem gets `AVIAMFSubmixLayout.sound_system` value.
func (smlo *AVIAMFSubmixLayout) GetSoundSystem() AVChannelLayout {
	return (AVChannelLayout)(smlo.sound_system)
}

// SetSoundSystem sets `AVIAMFSubmixLayout.sound_system` value.
func (smlo *AVIAMFSubmixLayout) SetSoundSystem(v AVChannelLayout) {
	smlo.sound_system = (C.struct_AVChannelLayout)(v)
}

// GetSoundSystemAddr gets `AVIAMFSubmixLayout.sound_system` address.
func (smlo *AVIAMFSubmixLayout) GetSoundSystemAddr() *AVChannelLayout {
	return (*AVChannelLayout)(&smlo.sound_system)
}

// GetIntegratedLoudness gets `AVIAMFSubmixLayout.integrated_loudness` value.
func (smlo *AVIAMFSubmixLayout) GetIntegratedLoudness() AVRational {
	return (AVRational)(smlo.integrated_loudness)
}

// SetIntegratedLoudness sets `AVIAMFSubmixLayout.integrated_loudness` value.
func (smlo *AVIAMFSubmixLayout) SetIntegratedLoudness(v AVRational) {
	smlo.integrated_loudness = (C.struct_AVRational)(v)
}

// GetIntegratedLoudnessAddr gets `AVIAMFSubmixLayout.integrated_loudness` address.
func (smlo *AVIAMFSubmixLayout) GetIntegratedLoudnessAddr() *AVRational {
	return (*AVRational)(&smlo.integrated_loudness)
}

// GetDigitalPeak gets `AVIAMFSubmixLayout.digital_peak` value.
func (smlo *AVIAMFSubmixLayout) GetDigitalPeak() AVRational {
	return (AVRational)(smlo.digital_peak)
}

// SetDigitalPeak sets `AVIAMFSubmixLayout.digital_peak` value.
func (smlo *AVIAMFSubmixLayout) SetDigitalPeak(v AVRational) {
	smlo.digital_peak = (C.struct_AVRational)(v)
}

// GetDigitalPeakAddr gets `AVIAMFSubmixLayout.digital_peak` address.
func (smlo *AVIAMFSubmixLayout) GetDigitalPeakAddr() *AVRational {
	return (*AVRational)(&smlo.digital_peak)
}

// GetTruePeak gets `AVIAMFSubmixLayout.true_peak` value.
func (smlo *AVIAMFSubmixLayout) GetTruePeak() AVRational {
	return (AVRational)(smlo.true_peak)
}

// SetTruePeak sets `AVIAMFSubmixLayout.true_peak` value.
func (smlo *AVIAMFSubmixLayout) SetTruePeak(v AVRational) {
	smlo.true_peak = (C.struct_AVRational)(v)
}

// GetTruePeakAddr gets `AVIAMFSubmixLayout.true_peak` address.
func (smlo *AVIAMFSubmixLayout) GetTruePeakAddr() *AVRational {
	return (*AVRational)(&smlo.true_peak)
}

// GetDialogueAnchoredLoudness gets `AVIAMFSubmixLayout.dialogue_anchored_loudness` value.
func (smlo *AVIAMFSubmixLayout) GetDialogueAnchoredLoudness() AVRational {
	return (AVRational)(smlo.dialogue_anchored_loudness)
}

// SetDialogueAnchoredLoudness sets `AVIAMFSubmixLayout.dialogue_anchored_loudness` value.
func (smlo *AVIAMFSubmixLayout) SetDialogueAnchoredLoudness(v AVRational) {
	smlo.dialogue_anchored_loudness = (C.struct_AVRational)(v)
}

// GetDialogueAnchoredLoudnessAddr gets `AVIAMFSubmixLayout.dialogue_anchored_loudness` address.
func (smlo *AVIAMFSubmixLayout) GetDialogueAnchoredLoudnessAddr() *AVRational {
	return (*AVRational)(&smlo.dialogue_anchored_loudness)
}

// GetAlbumAnchoredLoudness gets `AVIAMFSubmixLayout.album_anchored_loudness` value.
func (smlo *AVIAMFSubmixLayout) GetAlbumAnchoredLoudness() AVRational {
	return (AVRational)(smlo.album_anchored_loudness)
}

// SetAlbumAnchoredLoudness sets `AVIAMFSubmixLayout.album_anchored_loudness` value.
func (smlo *AVIAMFSubmixLayout) SetAlbumAnchoredLoudness(v AVRational) {
	smlo.album_anchored_loudness = (C.struct_AVRational)(v)
}

// GetAlbumAnchoredLoudnessAddr gets `AVIAMFSubmixLayout.album_anchored_loudness` address.
func (smlo *AVIAMFSubmixLayout) GetAlbumAnchoredLoudnessAddr() *AVRational {
	return (*AVRational)(&smlo.album_anchored_loudness)
}

// AVIAMFSubmix
type AVIAMFSubmix C.struct_AVIAMFSubmix

// GetAvClass gets `AVIAMFSubmix.av_class` value.
func (smix *AVIAMFSubmix) GetAvClass() *AVClass {
	return (*AVClass)(smix.av_class)
}

// SetAvClass sets `AVIAMFSubmix.av_class` value.
func (smix *AVIAMFSubmix) SetAvClass(v *AVClass) {
	smix.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFSubmix.av_class` address.
func (smix *AVIAMFSubmix) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&smix.av_class))
}

// GetElements gets `AVIAMFSubmix.elements` value.
func (smix *AVIAMFSubmix) GetElements() []*AVIAMFSubmixElement {
	return unsafe.Slice((**AVIAMFSubmixElement)(unsafe.Pointer(smix.elements)), smix.nb_elements)
}

// SetElements sets `AVIAMFSubmix.elements` value.
func (smix *AVIAMFSubmix) SetElements(v **AVIAMFSubmixElement) {
	smix.elements = (**C.struct_AVIAMFSubmixElement)(unsafe.Pointer(v))
}

// GetElementsAddr gets `AVIAMFSubmix.elements` address.
func (smix *AVIAMFSubmix) GetElementsAddr() ***AVIAMFSubmixElement {
	return (***AVIAMFSubmixElement)(unsafe.Pointer(&smix.elements))
}

// GetNbElements gets `AVIAMFSubmix.nb_elements` value.
func (smix *AVIAMFSubmix) GetNbElements() uint32 {
	return (uint32)(smix.nb_elements)
}

// SetNbElements sets `AVIAMFSubmix.nb_elements` value.
func (smix *AVIAMFSubmix) SetNbElements(v uint32) {
	smix.nb_elements = (C.uint)(v)
}

// GetNbElementsAddr gets `AVIAMFSubmix.nb_elements` address.
func (smix *AVIAMFSubmix) GetNbElementsAddr() *uint32 {
	return (*uint32)(&smix.nb_elements)
}

// GetLayouts gets `AVIAMFSubmix.layouts` value.
func (smix *AVIAMFSubmix) GetLayouts() []*AVIAMFSubmixLayout {
	return unsafe.Slice((**AVIAMFSubmixLayout)(unsafe.Pointer(smix.layouts)), smix.nb_layouts)
}

// SetLayouts sets `AVIAMFSubmix.layouts` value.
func (smix *AVIAMFSubmix) SetLayouts(v **AVIAMFSubmixLayout) {
	smix.layouts = (**C.struct_AVIAMFSubmixLayout)(unsafe.Pointer(v))
}

// GetLayoutsAddr gets `AVIAMFSubmix.layouts` address.
func (smix *AVIAMFSubmix) GetLayoutsAddr() ***AVIAMFSubmixLayout {
	return (***AVIAMFSubmixLayout)(unsafe.Pointer(&smix.layouts))
}

// GetNbLayouts gets `AVIAMFSubmix.nb_layouts` value.
func (smix *AVIAMFSubmix) GetNbLayouts() uint32 {
	return (uint32)(smix.nb_layouts)
}

// SetNbLayouts sets `AVIAMFSubmix.nb_layouts` value.
func (smix *AVIAMFSubmix) SetNbLayouts(v uint32) {
	smix.nb_layouts = (C.uint)(v)
}

// GetNbLayoutsAddr gets `AVIAMFSubmix.nb_layouts` address.
func (smix *AVIAMFSubmix) GetNbLayoutsAddr() *uint32 {
	return (*uint32)(&smix.nb_layouts)
}

// GetOutputMixConfig gets `AVIAMFSubmix.output_mix_config` value.
func (smix *AVIAMFSubmix) GetOutputMixConfig() *AVIAMFParamDefinition {
	return (*AVIAMFParamDefinition)(smix.output_mix_config)
}

// SetOutputMixConfig sets `AVIAMFSubmix.output_mix_config` value.
func (smix *AVIAMFSubmix) SetOutputMixConfig(v *AVIAMFParamDefinition) {
	smix.output_mix_config = (*C.struct_AVIAMFParamDefinition)(v)
}

// GetOutputMixConfigAddr gets `AVIAMFSubmix.output_mix_config` address.
func (smix *AVIAMFSubmix) GetOutputMixConfigAddr() **AVIAMFParamDefinition {
	return (**AVIAMFParamDefinition)(unsafe.Pointer(&smix.output_mix_config))
}

// GetDefaultMixGain gets `AVIAMFSubmix.default_mix_gain` value.
func (smix *AVIAMFSubmix) GetDefaultMixGain() AVRational {
	return (AVRational)(smix.default_mix_gain)
}

// SetDefaultMixGain sets `AVIAMFSubmix.default_mix_gain` value.
func (smix *AVIAMFSubmix) SetDefaultMixGain(v AVRational) {
	smix.default_mix_gain = (C.struct_AVRational)(v)
}

// GetDefaultMixGainAddr gets `AVIAMFSubmix.default_mix_gain` address.
func (smix *AVIAMFSubmix) GetDefaultMixGainAddr() *AVRational {
	return (*AVRational)(&smix.default_mix_gain)
}

// AVIAMFMixPresentation
type AVIAMFMixPresentation C.struct_AVIAMFMixPresentation

// GetAvClass gets `AVIAMFMixPresentation.av_class` value.
func (mpst *AVIAMFMixPresentation) GetAvClass() *AVClass {
	return (*AVClass)(mpst.av_class)
}

// SetAvClass sets `AVIAMFMixPresentation.av_class` value.
func (mpst *AVIAMFMixPresentation) SetAvClass(v *AVClass) {
	mpst.av_class = (*C.struct_AVClass)(v)
}

// GetAvClassAddr gets `AVIAMFMixPresentation.av_class` address.
func (mpst *AVIAMFMixPresentation) GetAvClassAddr() **AVClass {
	return (**AVClass)(unsafe.Pointer(&mpst.av_class))
}

// GetSubmixes gets `AVIAMFMixPresentation.submixes` value.
func (mpst *AVIAMFMixPresentation) GetSubmixes() []*AVIAMFSubmix {
	return unsafe.Slice((**AVIAMFSubmix)(unsafe.Pointer(mpst.submixes)), mpst.nb_submixes)
}

// SetSubmixes sets `AVIAMFMixPresentation.submixes` value.
func (mpst *AVIAMFMixPresentation) SetSubmixes(v **AVIAMFSubmix) {
	mpst.submixes = (**C.struct_AVIAMFSubmix)(unsafe.Pointer(v))
}

// GetSubmixesAddr gets `AVIAMFMixPresentation.submixes` address.
func (mpst *AVIAMFMixPresentation) GetSubmixesAddr() ***AVIAMFSubmix {
	return (***AVIAMFSubmix)(unsafe.Pointer(&mpst.submixes))
}

// GetNbSubmixes gets `AVIAMFMixPresentation.nb_submixes` value.
func (mpst *AVIAMFMixPresentation) GetNbSubmixes() uint32 {
	return (uint32)(mpst.nb_submixes)
}

// SetNbSubmixes sets `AVIAMFMixPresentation.nb_submixes` value.
func (mpst *AVIAMFMixPresentation) SetNbSubmixes(v uint32) {
	mpst.nb_submixes = (C.uint)(v)
}

// GetNbSubmixesAddr gets `AVIAMFMixPresentation.nb_submixes` address.
func (mpst *AVIAMFMixPresentation) GetNbSubmixesAddr() *uint32 {
	return (*uint32)(&mpst.nb_submixes)
}

// GetAnnotations gets `AVIAMFMixPresentation.annotations` value.
func (mpst *AVIAMFMixPresentation) GetAnnotations() *AVDictionary {
	return (*AVDictionary)(mpst.annotations)
}

// SetAnnotations sets `AVIAMFMixPresentation.annotations` value.
func (mpst *AVIAMFMixPresentation) SetAnnotations(v *AVDictionary) {
	mpst.annotations = (*C.struct_AVDictionary)(v)
}

// GetAnnotationsAddr gets `AVIAMFMixPresentation.annotations` address.
func (mpst *AVIAMFMixPresentation) GetAnnotationsAddr() **AVDictionary {
	return (**AVDictionary)(unsafe.Pointer(&mpst.annotations))
}

// AvIamfMixPresentationGetClass
func AvIamfMixPresentationGetClass() *AVClass {
	return (*AVClass)(C.av_iamf_mix_presentation_get_class())
}

// AvIamfMixPresentationAlloc allocates a AVIAMFMixPresentation, and initializes its fields with default values.
func AvIamfMixPresentationAlloc() *AVIAMFMixPresentation {
	return (*AVIAMFMixPresentation)(C.av_iamf_mix_presentation_alloc())
}

// AvIamfMixPresentationAddSubmix allocates a submix and add it to a given AVIAMFMixPresentation.
func AvIamfMixPresentationAddSubmix(mixPresentation *AVIAMFMixPresentation) *AVIAMFSubmix {
	return (*AVIAMFSubmix)(C.av_iamf_mix_presentation_add_submix((*C.struct_AVIAMFMixPresentation)(mixPresentation)))
}

// AvIamfSubmixAddElement allocates a submix element and add it to a given AVIAMFSubmix.
func AvIamfSubmixAddElement(submix *AVIAMFSubmix) *AVIAMFSubmixElement {
	return (*AVIAMFSubmixElement)(C.av_iamf_submix_add_element((*C.struct_AVIAMFSubmix)(submix)))
}

// AvIamfSubmixAddLayout allocates a submix layout and add it to a given AVIAMFSubmix.
func AvIamfSubmixAddLayout(submix *AVIAMFSubmix) *AVIAMFSubmixLayout {
	return (*AVIAMFSubmixLayout)(C.av_iamf_submix_add_layout((*C.struct_AVIAMFSubmix)(submix)))
}

// AvIamfMixPresentationFree frees an AVIAMFMixPresentation and all its contents.
func AvIamfMixPresentationFree(mixPresentation **AVIAMFMixPresentation) {
	C.av_iamf_mix_presentation_free((**C.struct_AVIAMFMixPresentation)(unsafe.Pointer(mixPresentation)))
}
