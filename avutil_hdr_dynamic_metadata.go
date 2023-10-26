package ffmpeg

/*
#include <libavutil/hdr_dynamic_metadata.h>
*/
import "C"
import "unsafe"

// AVHDRPlusOverlapProcessOption
type AVHDRPlusOverlapProcessOption = C.enum_AVHDRPlusOverlapProcessOption

const (
	AV_HDR_PLUS_OVERLAP_PROCESS_WEIGHTED_AVERAGING = AVHDRPlusOverlapProcessOption(
		C.AV_HDR_PLUS_OVERLAP_PROCESS_WEIGHTED_AVERAGING)
	AV_HDR_PLUS_OVERLAP_PROCESS_LAYERING = AVHDRPlusOverlapProcessOption(
		C.AV_HDR_PLUS_OVERLAP_PROCESS_LAYERING)
)

// AVHDRPlusPercentile
type AVHDRPlusPercentile C.struct_AVHDRPlusPercentile

// Custom: GetPercentage gets `AVHDRPlusPercentile.percentage` value.
func (pct *AVHDRPlusPercentile) GetPercentage() uint8 {
	return (uint8)(pct.percentage)
}

// Custom: SetPercentage sets `AVHDRPlusPercentile.percentage` value.
func (pct *AVHDRPlusPercentile) SetPercentage(v uint8) {
	pct.percentage = (C.uint8_t)(v)
}

// Custom: GetPercentageAddr gets `AVHDRPlusPercentile.percentage` address.
func (pct *AVHDRPlusPercentile) GetPercentageAddr() *uint8 {
	return (*uint8)(&pct.percentage)
}

// Custom: GetPercentile gets `AVHDRPlusPercentile.percentile` value.
func (pct *AVHDRPlusPercentile) GetPercentile() AVRational {
	return (AVRational)(pct.percentile)
}

// Custom: SetPercentile sets `AVHDRPlusPercentile.percentile` value.
func (pct *AVHDRPlusPercentile) SetPercentile(v AVRational) {
	pct.percentile = (C.struct_AVRational)(v)
}

// Custom: GetPercentileAddr gets `AVHDRPlusPercentile.percentile` address.
func (pct *AVHDRPlusPercentile) GetPercentileAddr() *AVRational {
	return (*AVRational)(&pct.percentile)
}

// AVHDRPlusColorTransformParams
type AVHDRPlusColorTransformParams C.struct_AVHDRPlusColorTransformParams

// Custom: GetWindowUpperLeftCornerX gets `AVHDRPlusColorTransformParams.window_upper_left_corner_x` value.
func (ctp *AVHDRPlusColorTransformParams) GetWindowUpperLeftCornerX() AVRational {
	return (AVRational)(ctp.window_upper_left_corner_x)
}

// Custom: SetWindowUpperLeftCornerX sets `AVHDRPlusColorTransformParams.window_upper_left_corner_x` value.
func (ctp *AVHDRPlusColorTransformParams) SetWindowUpperLeftCornerX(v AVRational) {
	ctp.window_upper_left_corner_x = (C.struct_AVRational)(v)
}

// Custom: GetWindowUpperLeftCornerXAddr gets `AVHDRPlusColorTransformParams.window_upper_left_corner_x` address.
func (ctp *AVHDRPlusColorTransformParams) GetWindowUpperLeftCornerXAddr() *AVRational {
	return (*AVRational)(&ctp.window_upper_left_corner_x)
}

// Custom: GetWindowUpperLeftCornerY gets `AVHDRPlusColorTransformParams.window_upper_left_corner_y` value.
func (ctp *AVHDRPlusColorTransformParams) GetWindowUpperLeftCornerY() AVRational {
	return (AVRational)(ctp.window_upper_left_corner_y)
}

// Custom: SetWindowUpperLeftCornerY sets `AVHDRPlusColorTransformParams.window_upper_left_corner_y` value.
func (ctp *AVHDRPlusColorTransformParams) SetWindowUpperLeftCornerY(v AVRational) {
	ctp.window_upper_left_corner_y = (C.struct_AVRational)(v)
}

// Custom: GetWindowUpperLeftCornerYAddr gets `AVHDRPlusColorTransformParams.window_upper_left_corner_y` address.
func (ctp *AVHDRPlusColorTransformParams) GetWindowUpperLeftCornerYAddr() *AVRational {
	return (*AVRational)(&ctp.window_upper_left_corner_y)
}

// Custom: GetWindowLowerRightCornerX gets `AVHDRPlusColorTransformParams.window_lower_right_corner_x` value.
func (ctp *AVHDRPlusColorTransformParams) GetWindowLowerRightCornerX() AVRational {
	return (AVRational)(ctp.window_lower_right_corner_x)
}

// Custom: SetWindowLowerRightCornerX sets `AVHDRPlusColorTransformParams.window_lower_right_corner_x` value.
func (ctp *AVHDRPlusColorTransformParams) SetWindowLowerRightCornerX(v AVRational) {
	ctp.window_lower_right_corner_x = (C.struct_AVRational)(v)
}

// Custom: GetWindowLowerRightCornerXAddr gets `AVHDRPlusColorTransformParams.window_lower_right_corner_x` address.
func (ctp *AVHDRPlusColorTransformParams) GetWindowLowerRightCornerXAddr() *AVRational {
	return (*AVRational)(&ctp.window_lower_right_corner_x)
}

// Custom: GetWindowLowerRightCornerY gets `AVHDRPlusColorTransformParams.window_lower_right_corner_y` value.
func (ctp *AVHDRPlusColorTransformParams) GetWindowLowerRightCornerY() AVRational {
	return (AVRational)(ctp.window_lower_right_corner_y)
}

// Custom: SetWindowLowerRightCornerY sets `AVHDRPlusColorTransformParams.window_lower_right_corner_y` value.
func (ctp *AVHDRPlusColorTransformParams) SetWindowLowerRightCornerY(v AVRational) {
	ctp.window_lower_right_corner_y = (C.struct_AVRational)(v)
}

// Custom: GetWindowLowerRightCornerYAddr gets `AVHDRPlusColorTransformParams.window_lower_right_corner_y` address.
func (ctp *AVHDRPlusColorTransformParams) GetWindowLowerRightCornerYAddr() *AVRational {
	return (*AVRational)(&ctp.window_lower_right_corner_y)
}

// Custom: GetCenterOfEllipseX gets `AVHDRPlusColorTransformParams.center_of_ellipse_x` value.
func (ctp *AVHDRPlusColorTransformParams) GetCenterOfEllipseX() uint16 {
	return (uint16)(ctp.center_of_ellipse_x)
}

// Custom: SetCenterOfEllipseX sets `AVHDRPlusColorTransformParams.center_of_ellipse_x` value.
func (ctp *AVHDRPlusColorTransformParams) SetCenterOfEllipseX(v uint16) {
	ctp.center_of_ellipse_x = (C.uint16_t)(v)
}

// Custom: GetCenterOfEllipseXAddr gets `AVHDRPlusColorTransformParams.center_of_ellipse_x` address.
func (ctp *AVHDRPlusColorTransformParams) GetCenterOfEllipseXAddr() *uint16 {
	return (*uint16)(&ctp.center_of_ellipse_x)
}

// Custom: GetCenterOfEllipseY gets `AVHDRPlusColorTransformParams.center_of_ellipse_y` value.
func (ctp *AVHDRPlusColorTransformParams) GetCenterOfEllipseY() uint16 {
	return (uint16)(ctp.center_of_ellipse_y)
}

// Custom: SetCenterOfEllipseY sets `AVHDRPlusColorTransformParams.center_of_ellipse_y` value.
func (ctp *AVHDRPlusColorTransformParams) SetCenterOfEllipseY(v uint16) {
	ctp.center_of_ellipse_y = (C.uint16_t)(v)
}

// Custom: GetCenterOfEllipseYAddr gets `AVHDRPlusColorTransformParams.center_of_ellipse_y` address.
func (ctp *AVHDRPlusColorTransformParams) GetCenterOfEllipseYAddr() *uint16 {
	return (*uint16)(&ctp.center_of_ellipse_y)
}

// Custom: GetRotationAngle gets `AVHDRPlusColorTransformParams.rotation_angle` value.
func (ctp *AVHDRPlusColorTransformParams) GetRotationAngle() uint8 {
	return (uint8)(ctp.rotation_angle)
}

// Custom: SetRotationAngle sets `AVHDRPlusColorTransformParams.rotation_angle` value.
func (ctp *AVHDRPlusColorTransformParams) SetRotationAngle(v uint8) {
	ctp.rotation_angle = (C.uint8_t)(v)
}

// Custom: GetRotationAngleAddr gets `AVHDRPlusColorTransformParams.rotation_angle` address.
func (ctp *AVHDRPlusColorTransformParams) GetRotationAngleAddr() *uint8 {
	return (*uint8)(&ctp.rotation_angle)
}

// Custom: GetSemimajorAxisInternalEllipse gets `AVHDRPlusColorTransformParams.semimajor_axis_internal_ellipse` value.
func (ctp *AVHDRPlusColorTransformParams) GetSemimajorAxisInternalEllipse() uint16 {
	return (uint16)(ctp.semimajor_axis_internal_ellipse)
}

// Custom: SetSemimajorAxisInternalEllipse sets `AVHDRPlusColorTransformParams.semimajor_axis_internal_ellipse` value.
func (ctp *AVHDRPlusColorTransformParams) SetSemimajorAxisInternalEllipse(v uint16) {
	ctp.semimajor_axis_internal_ellipse = (C.uint16_t)(v)
}

// Custom: GetSemimajorAxisInternalEllipseAddr gets `AVHDRPlusColorTransformParams.semimajor_axis_internal_ellipse` address.
func (ctp *AVHDRPlusColorTransformParams) GetSemimajorAxisInternalEllipseAddr() *uint16 {
	return (*uint16)(&ctp.semimajor_axis_internal_ellipse)
}

// Custom: GetSemimajorAxisExternalEllipse gets `AVHDRPlusColorTransformParams.semimajor_axis_external_ellipse` value.
func (ctp *AVHDRPlusColorTransformParams) GetSemimajorAxisExternalEllipse() uint16 {
	return (uint16)(ctp.semimajor_axis_external_ellipse)
}

// Custom: SetSemimajorAxisExternalEllipse sets `AVHDRPlusColorTransformParams.semimajor_axis_external_ellipse` value.
func (ctp *AVHDRPlusColorTransformParams) SetSemimajorAxisExternalEllipse(v uint16) {
	ctp.semimajor_axis_external_ellipse = (C.uint16_t)(v)
}

// Custom: GetSemimajorAxisExternalEllipseAddr gets `AVHDRPlusColorTransformParams.semimajor_axis_external_ellipse` address.
func (ctp *AVHDRPlusColorTransformParams) GetSemimajorAxisExternalEllipseAddr() *uint16 {
	return (*uint16)(&ctp.semimajor_axis_external_ellipse)
}

// Custom: GetSemiminorAxisExternalEllipse gets `AVHDRPlusColorTransformParams.semiminor_axis_external_ellipse` value.
func (ctp *AVHDRPlusColorTransformParams) GetSemiminorAxisExternalEllipse() uint16 {
	return (uint16)(ctp.semiminor_axis_external_ellipse)
}

// Custom: SetSemiminorAxisExternalEllipse sets `AVHDRPlusColorTransformParams.semiminor_axis_external_ellipse` value.
func (ctp *AVHDRPlusColorTransformParams) SetSemiminorAxisExternalEllipse(v uint16) {
	ctp.semiminor_axis_external_ellipse = (C.uint16_t)(v)
}

// Custom: GetSemiminorAxisExternalEllipseAddr gets `AVHDRPlusColorTransformParams.semiminor_axis_external_ellipse` address.
func (ctp *AVHDRPlusColorTransformParams) GetSemiminorAxisExternalEllipseAddr() *uint16 {
	return (*uint16)(&ctp.semiminor_axis_external_ellipse)
}

// Custom: GetOverlapProcessOption gets `AVHDRPlusColorTransformParams.overlap_process_option` value.
func (ctp *AVHDRPlusColorTransformParams) GetOverlapProcessOption() AVHDRPlusOverlapProcessOption {
	return (AVHDRPlusOverlapProcessOption)(ctp.overlap_process_option)
}

// Custom: SetOverlapProcessOption sets `AVHDRPlusColorTransformParams.overlap_process_option` value.
func (ctp *AVHDRPlusColorTransformParams) SetOverlapProcessOption(v AVHDRPlusOverlapProcessOption) {
	ctp.overlap_process_option = (C.enum_AVHDRPlusOverlapProcessOption)(v)
}

// Custom: GetOverlapProcessOptionAddr gets `AVHDRPlusColorTransformParams.overlap_process_option` address.
func (ctp *AVHDRPlusColorTransformParams) GetOverlapProcessOptionAddr() *AVHDRPlusOverlapProcessOption {
	return (*AVHDRPlusOverlapProcessOption)(&ctp.overlap_process_option)
}

// Custom: GetMaxscl gets `AVHDRPlusColorTransformParams.maxscl` value.
func (ctp *AVHDRPlusColorTransformParams) GetMaxscl() []AVRational {
	return unsafe.Slice((*AVRational)(&ctp.maxscl[0]), 3)
}

// Custom: SetMaxscl sets `AVHDRPlusColorTransformParams.maxscl` value.
func (ctp *AVHDRPlusColorTransformParams) SetMaxscl(v []AVRational) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		ctp.maxscl[i] = (C.struct_AVRational)(v[i])
	}
}

// Custom: GetMaxsclAddr gets `AVHDRPlusColorTransformParams.maxscl` address.
func (ctp *AVHDRPlusColorTransformParams) GetMaxsclAddr() **AVRational {
	return (**AVRational)(unsafe.Pointer(&ctp.maxscl))
}

// Custom: GetAverageMaxrgb gets `AVHDRPlusColorTransformParams.average_maxrgb` value.
func (ctp *AVHDRPlusColorTransformParams) GetAverageMaxrgb() AVRational {
	return (AVRational)(ctp.average_maxrgb)
}

// Custom: SetAverageMaxrgb sets `AVHDRPlusColorTransformParams.average_maxrgb` value.
func (ctp *AVHDRPlusColorTransformParams) SetAverageMaxrgb(v AVRational) {
	ctp.average_maxrgb = (C.struct_AVRational)(v)
}

// Custom: GetAverageMaxrgbAddr gets `AVHDRPlusColorTransformParams.average_maxrgb` address.
func (ctp *AVHDRPlusColorTransformParams) GetAverageMaxrgbAddr() *AVRational {
	return (*AVRational)(&ctp.average_maxrgb)
}

// Custom: GetNumDistributionMaxrgbPercentiles gets `AVHDRPlusColorTransformParams.num_distribution_maxrgb_percentiles` value.
func (ctp *AVHDRPlusColorTransformParams) GetNumDistributionMaxrgbPercentiles() uint8 {
	return (uint8)(ctp.num_distribution_maxrgb_percentiles)
}

// Custom: SetNumDistributionMaxrgbPercentiles sets `AVHDRPlusColorTransformParams.num_distribution_maxrgb_percentiles` value.
func (ctp *AVHDRPlusColorTransformParams) SetNumDistributionMaxrgbPercentiles(v uint8) {
	ctp.num_distribution_maxrgb_percentiles = (C.uint8_t)(v)
}

// Custom: GetNumDistributionMaxrgbPercentilesAddr gets `AVHDRPlusColorTransformParams.num_distribution_maxrgb_percentiles` address.
func (ctp *AVHDRPlusColorTransformParams) GetNumDistributionMaxrgbPercentilesAddr() *uint8 {
	return (*uint8)(&ctp.num_distribution_maxrgb_percentiles)
}

// Custom: GetDistributionMaxrgb gets `AVHDRPlusColorTransformParams.distribution_maxrgb` value.
func (ctp *AVHDRPlusColorTransformParams) GetDistributionMaxrgb() []AVHDRPlusPercentile {
	return unsafe.Slice((*AVHDRPlusPercentile)(&ctp.distribution_maxrgb[0]), 15)
}

// Custom: SetDistributionMaxrgb sets `AVHDRPlusColorTransformParams.distribution_maxrgb` value.
func (ctp *AVHDRPlusColorTransformParams) SetDistributionMaxrgb(v []AVHDRPlusPercentile) {
	for i := 0; i < FFMIN(len(v), 15); i++ {
		ctp.distribution_maxrgb[i] = (C.struct_AVHDRPlusPercentile)(v[i])
	}
}

// Custom: GetDistributionMaxrgbAddr gets `AVHDRPlusColorTransformParams.distribution_maxrgb` address.
func (ctp *AVHDRPlusColorTransformParams) GetDistributionMaxrgbAddr() **AVHDRPlusPercentile {
	return (**AVHDRPlusPercentile)(unsafe.Pointer(&ctp.distribution_maxrgb))
}

// Custom: GetFractionBrightPixels gets `AVHDRPlusColorTransformParams.fraction_bright_pixels` value.
func (ctp *AVHDRPlusColorTransformParams) GetFractionBrightPixels() AVRational {
	return (AVRational)(ctp.fraction_bright_pixels)
}

// Custom: SetFractionBrightPixels sets `AVHDRPlusColorTransformParams.fraction_bright_pixels` value.
func (ctp *AVHDRPlusColorTransformParams) SetFractionBrightPixels(v AVRational) {
	ctp.fraction_bright_pixels = (C.struct_AVRational)(v)
}

// Custom: GetFractionBrightPixelsAddr gets `AVHDRPlusColorTransformParams.fraction_bright_pixels` address.
func (ctp *AVHDRPlusColorTransformParams) GetFractionBrightPixelsAddr() *AVRational {
	return (*AVRational)(&ctp.fraction_bright_pixels)
}

// Custom: GetToneMappingFlag gets `AVHDRPlusColorTransformParams.tone_mapping_flag` value.
func (ctp *AVHDRPlusColorTransformParams) GetToneMappingFlag() uint8 {
	return (uint8)(ctp.tone_mapping_flag)
}

// Custom: SetToneMappingFlag sets `AVHDRPlusColorTransformParams.tone_mapping_flag` value.
func (ctp *AVHDRPlusColorTransformParams) SetToneMappingFlag(v uint8) {
	ctp.tone_mapping_flag = (C.uint8_t)(v)
}

// Custom: GetToneMappingFlagAddr gets `AVHDRPlusColorTransformParams.tone_mapping_flag` address.
func (ctp *AVHDRPlusColorTransformParams) GetToneMappingFlagAddr() *uint8 {
	return (*uint8)(&ctp.tone_mapping_flag)
}

// Custom: GetKneePointX gets `AVHDRPlusColorTransformParams.knee_point_x` value.
func (ctp *AVHDRPlusColorTransformParams) GetKneePointX() AVRational {
	return (AVRational)(ctp.knee_point_x)
}

// Custom: SetKneePointX sets `AVHDRPlusColorTransformParams.knee_point_x` value.
func (ctp *AVHDRPlusColorTransformParams) SetKneePointX(v AVRational) {
	ctp.knee_point_x = (C.struct_AVRational)(v)
}

// Custom: GetKneePointXAddr gets `AVHDRPlusColorTransformParams.knee_point_x` address.
func (ctp *AVHDRPlusColorTransformParams) GetKneePointXAddr() *AVRational {
	return (*AVRational)(&ctp.knee_point_x)
}

// Custom: GetKneePointY gets `AVHDRPlusColorTransformParams.knee_point_y` value.
func (ctp *AVHDRPlusColorTransformParams) GetKneePointY() AVRational {
	return (AVRational)(ctp.knee_point_y)
}

// Custom: SetKneePointY sets `AVHDRPlusColorTransformParams.knee_point_y` value.
func (ctp *AVHDRPlusColorTransformParams) SetKneePointY(v AVRational) {
	ctp.knee_point_y = (C.struct_AVRational)(v)
}

// Custom: GetKneePointYAddr gets `AVHDRPlusColorTransformParams.knee_point_y` address.
func (ctp *AVHDRPlusColorTransformParams) GetKneePointYAddr() *AVRational {
	return (*AVRational)(&ctp.knee_point_y)
}

// Custom: GetNumBezierCurveAnchors gets `AVHDRPlusColorTransformParams.num_bezier_curve_anchors` value.
func (ctp *AVHDRPlusColorTransformParams) GetNumBezierCurveAnchors() uint8 {
	return (uint8)(ctp.num_bezier_curve_anchors)
}

// Custom: SetNumBezierCurveAnchors sets `AVHDRPlusColorTransformParams.num_bezier_curve_anchors` value.
func (ctp *AVHDRPlusColorTransformParams) SetNumBezierCurveAnchors(v uint8) {
	ctp.num_bezier_curve_anchors = (C.uint8_t)(v)
}

// Custom: GetNumBezierCurveAnchorsAddr gets `AVHDRPlusColorTransformParams.num_bezier_curve_anchors` address.
func (ctp *AVHDRPlusColorTransformParams) GetNumBezierCurveAnchorsAddr() *uint8 {
	return (*uint8)(&ctp.num_bezier_curve_anchors)
}

// Custom: GetBezierCurveAnchors gets `AVHDRPlusColorTransformParams.bezier_curve_anchors` value.
func (ctp *AVHDRPlusColorTransformParams) GetBezierCurveAnchors() []AVRational {
	return unsafe.Slice((*AVRational)(&ctp.bezier_curve_anchors[0]), 15)
}

// Custom: SetBezierCurveAnchors sets `AVHDRPlusColorTransformParams.bezier_curve_anchors` value.
func (ctp *AVHDRPlusColorTransformParams) SetBezierCurveAnchors(v []AVRational) {
	for i := 0; i < FFMIN(len(v), 15); i++ {
		ctp.bezier_curve_anchors[i] = (C.struct_AVRational)(v[i])
	}
}

// Custom: GetBezierCurveAnchorsAddr gets `AVHDRPlusColorTransformParams.bezier_curve_anchors` address.
func (ctp *AVHDRPlusColorTransformParams) GetBezierCurveAnchorsAddr() **AVRational {
	return (**AVRational)(unsafe.Pointer(&ctp.bezier_curve_anchors))
}

// Custom: GetColorSaturationMappingFlag gets `AVHDRPlusColorTransformParams.color_saturation_mapping_flag` value.
func (ctp *AVHDRPlusColorTransformParams) GetColorSaturationMappingFlag() uint8 {
	return (uint8)(ctp.color_saturation_mapping_flag)
}

// Custom: SetColorSaturationMappingFlag sets `AVHDRPlusColorTransformParams.color_saturation_mapping_flag` value.
func (ctp *AVHDRPlusColorTransformParams) SetColorSaturationMappingFlag(v uint8) {
	ctp.color_saturation_mapping_flag = (C.uint8_t)(v)
}

// Custom: GetColorSaturationMappingFlagAddr gets `AVHDRPlusColorTransformParams.color_saturation_mapping_flag` address.
func (ctp *AVHDRPlusColorTransformParams) GetColorSaturationMappingFlagAddr() *uint8 {
	return (*uint8)(&ctp.color_saturation_mapping_flag)
}

// Custom: GetColorSaturationWeight gets `AVHDRPlusColorTransformParams.color_saturation_weight` value.
func (ctp *AVHDRPlusColorTransformParams) GetColorSaturationWeight() AVRational {
	return (AVRational)(ctp.color_saturation_weight)
}

// Custom: SetColorSaturationWeight sets `AVHDRPlusColorTransformParams.color_saturation_weight` value.
func (ctp *AVHDRPlusColorTransformParams) SetColorSaturationWeight(v AVRational) {
	ctp.color_saturation_weight = (C.struct_AVRational)(v)
}

// Custom: GetColorSaturationWeightAddr gets `AVHDRPlusColorTransformParams.color_saturation_weight` address.
func (ctp *AVHDRPlusColorTransformParams) GetColorSaturationWeightAddr() *AVRational {
	return (*AVRational)(&ctp.color_saturation_weight)
}

// AVDynamicHDRPlus
type AVDynamicHDRPlus C.struct_AVDynamicHDRPlus

// Custom: GetItuTT35CountryCode gets `AVDynamicHDRPlus.itu_t_t35_country_code` value.
func (dhp *AVDynamicHDRPlus) GetItuTT35CountryCode() uint8 {
	return (uint8)(dhp.itu_t_t35_country_code)
}

// Custom: SetItuTT35CountryCode sets `AVDynamicHDRPlus.itu_t_t35_country_code` value.
func (dhp *AVDynamicHDRPlus) SetItuTT35CountryCode(v uint8) {
	dhp.itu_t_t35_country_code = (C.uint8_t)(v)
}

// Custom: GetItuTT35CountryCodeAddr gets `AVDynamicHDRPlus.itu_t_t35_country_code` address.
func (dhp *AVDynamicHDRPlus) GetItuTT35CountryCodeAddr() *uint8 {
	return (*uint8)(&dhp.itu_t_t35_country_code)
}

// Custom: GetApplicationVersion gets `AVDynamicHDRPlus.application_version` value.
func (dhp *AVDynamicHDRPlus) GetApplicationVersion() uint8 {
	return (uint8)(dhp.application_version)
}

// Custom: SetApplicationVersion sets `AVDynamicHDRPlus.application_version` value.
func (dhp *AVDynamicHDRPlus) SetApplicationVersion(v uint8) {
	dhp.application_version = (C.uint8_t)(v)
}

// Custom: GetApplicationVersionAddr gets `AVDynamicHDRPlus.application_version` address.
func (dhp *AVDynamicHDRPlus) GetApplicationVersionAddr() *uint8 {
	return (*uint8)(&dhp.application_version)
}

// Custom: GetNumWindows gets `AVDynamicHDRPlus.num_windows` value.
func (dhp *AVDynamicHDRPlus) GetNumWindows() uint8 {
	return (uint8)(dhp.num_windows)
}

// Custom: SetNumWindows sets `AVDynamicHDRPlus.num_windows` value.
func (dhp *AVDynamicHDRPlus) SetNumWindows(v uint8) {
	dhp.num_windows = (C.uint8_t)(v)
}

// Custom: GetNumWindowsAddr gets `AVDynamicHDRPlus.num_windows` address.
func (dhp *AVDynamicHDRPlus) GetNumWindowsAddr() *uint8 {
	return (*uint8)(&dhp.num_windows)
}

// Custom: GetParams gets `AVDynamicHDRPlus.params` value.
func (dhp *AVDynamicHDRPlus) GetParams() []AVHDRPlusColorTransformParams {
	return unsafe.Slice((*AVHDRPlusColorTransformParams)(&dhp.params[0]), 3)
}

// Custom: SetParams sets `AVDynamicHDRPlus.params` value.
func (dhp *AVDynamicHDRPlus) SetParams(v []AVHDRPlusColorTransformParams) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		dhp.params[i] = (C.struct_AVHDRPlusColorTransformParams)(v[i])
	}
}

// Custom: GetParamsAddr gets `AVDynamicHDRPlus.params` address.
func (dhp *AVDynamicHDRPlus) GetParamsAddr() **AVHDRPlusColorTransformParams {
	return (**AVHDRPlusColorTransformParams)(unsafe.Pointer(&dhp.params))
}

// Custom: GetTargetedSystemDisplayMaximumLuminance gets `AVDynamicHDRPlus.targeted_system_display_maximum_luminance` value.
func (dhp *AVDynamicHDRPlus) GetTargetedSystemDisplayMaximumLuminance() AVRational {
	return (AVRational)(dhp.targeted_system_display_maximum_luminance)
}

// Custom: SetTargetedSystemDisplayMaximumLuminance sets `AVDynamicHDRPlus.targeted_system_display_maximum_luminance` value.
func (dhp *AVDynamicHDRPlus) SetTargetedSystemDisplayMaximumLuminance(v AVRational) {
	dhp.targeted_system_display_maximum_luminance = (C.struct_AVRational)(v)
}

// Custom: GetTargetedSystemDisplayMaximumLuminanceAddr gets `AVDynamicHDRPlus.targeted_system_display_maximum_luminance` address.
func (dhp *AVDynamicHDRPlus) GetTargetedSystemDisplayMaximumLuminanceAddr() *AVRational {
	return (*AVRational)(&dhp.targeted_system_display_maximum_luminance)
}

// Custom: GetTargetedSystemDisplayActualPeakLuminanceFlag gets `AVDynamicHDRPlus.targeted_system_display_actual_peak_luminance_flag` value.
func (dhp *AVDynamicHDRPlus) GetTargetedSystemDisplayActualPeakLuminanceFlag() uint8 {
	return (uint8)(dhp.targeted_system_display_actual_peak_luminance_flag)
}

// Custom: SetTargetedSystemDisplayActualPeakLuminanceFlag sets `AVDynamicHDRPlus.targeted_system_display_actual_peak_luminance_flag` value.
func (dhp *AVDynamicHDRPlus) SetTargetedSystemDisplayActualPeakLuminanceFlag(v uint8) {
	dhp.targeted_system_display_actual_peak_luminance_flag = (C.uint8_t)(v)
}

// Custom: GetTargetedSystemDisplayActualPeakLuminanceFlagAddr gets `AVDynamicHDRPlus.targeted_system_display_actual_peak_luminance_flag` address.
func (dhp *AVDynamicHDRPlus) GetTargetedSystemDisplayActualPeakLuminanceFlagAddr() *uint8 {
	return (*uint8)(&dhp.targeted_system_display_actual_peak_luminance_flag)
}

// Custom: GetNumRowsTargetedSystemDisplayActualPeakLuminance gets `AVDynamicHDRPlus.num_rows_targeted_system_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) GetNumRowsTargetedSystemDisplayActualPeakLuminance() uint8 {
	return (uint8)(dhp.num_rows_targeted_system_display_actual_peak_luminance)
}

// Custom: SetNumRowsTargetedSystemDisplayActualPeakLuminance sets `AVDynamicHDRPlus.num_rows_targeted_system_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) SetNumRowsTargetedSystemDisplayActualPeakLuminance(v uint8) {
	dhp.num_rows_targeted_system_display_actual_peak_luminance = (C.uint8_t)(v)
}

// Custom: GetNumRowsTargetedSystemDisplayActualPeakLuminanceAddr gets `AVDynamicHDRPlus.num_rows_targeted_system_display_actual_peak_luminance` address.
func (dhp *AVDynamicHDRPlus) GetNumRowsTargetedSystemDisplayActualPeakLuminanceAddr() *uint8 {
	return (*uint8)(&dhp.num_rows_targeted_system_display_actual_peak_luminance)
}

// Custom: GetNumColsTargetedSystemDisplayActualPeakLuminance gets `AVDynamicHDRPlus.num_cols_targeted_system_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) GetNumColsTargetedSystemDisplayActualPeakLuminance() uint8 {
	return (uint8)(dhp.num_cols_targeted_system_display_actual_peak_luminance)
}

// Custom: SetNumColsTargetedSystemDisplayActualPeakLuminance sets `AVDynamicHDRPlus.num_cols_targeted_system_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) SetNumColsTargetedSystemDisplayActualPeakLuminance(v uint8) {
	dhp.num_cols_targeted_system_display_actual_peak_luminance = (C.uint8_t)(v)
}

// Custom: GetNumColsTargetedSystemDisplayActualPeakLuminanceAddr gets `AVDynamicHDRPlus.num_cols_targeted_system_display_actual_peak_luminance` address.
func (dhp *AVDynamicHDRPlus) GetNumColsTargetedSystemDisplayActualPeakLuminanceAddr() *uint8 {
	return (*uint8)(&dhp.num_cols_targeted_system_display_actual_peak_luminance)
}

// Custom: GetTargetedSystemDisplayActualPeakLuminance gets `AVDynamicHDRPlus.targeted_system_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) GetTargetedSystemDisplayActualPeakLuminance() (v [][]AVRational) {
	for i := 0; i < 25; i++ {
		v = append(v, unsafe.Slice((*AVRational)(&dhp.targeted_system_display_actual_peak_luminance[i][0]), 25))
	}
	return v
}

// Custom: SetTargetedSystemDisplayActualPeakLuminance sets `AVDynamicHDRPlus.targeted_system_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) SetTargetedSystemDisplayActualPeakLuminance(v [][]AVRational) {
	for i := 0; i < FFMIN(len(v), 25); i++ {
		for j := 0; j < FFMIN(len(v[i]), 25); j++ {
			dhp.targeted_system_display_actual_peak_luminance[i][j] = (C.struct_AVRational)(v[i][j])
		}
	}
}

// Custom: GetTargetedSystemDisplayActualPeakLuminanceAddr gets `AVDynamicHDRPlus.targeted_system_display_actual_peak_luminance` address.
func (dhp *AVDynamicHDRPlus) GetTargetedSystemDisplayActualPeakLuminanceAddr() **AVRational {
	return (**AVRational)(unsafe.Pointer(&dhp.targeted_system_display_actual_peak_luminance))
}

// Custom: GetMasteringDisplayActualPeakLuminanceFlag gets `AVDynamicHDRPlus.mastering_display_actual_peak_luminance_flag` value.
func (dhp *AVDynamicHDRPlus) GetMasteringDisplayActualPeakLuminanceFlag() uint8 {
	return (uint8)(dhp.mastering_display_actual_peak_luminance_flag)
}

// Custom: SetMasteringDisplayActualPeakLuminanceFlag sets `AVDynamicHDRPlus.mastering_display_actual_peak_luminance_flag` value.
func (dhp *AVDynamicHDRPlus) SetMasteringDisplayActualPeakLuminanceFlag(v uint8) {
	dhp.mastering_display_actual_peak_luminance_flag = (C.uint8_t)(v)
}

// Custom: GetMasteringDisplayActualPeakLuminanceFlagAddr gets `AVDynamicHDRPlus.mastering_display_actual_peak_luminance_flag` address.
func (dhp *AVDynamicHDRPlus) GetMasteringDisplayActualPeakLuminanceFlagAddr() *uint8 {
	return (*uint8)(&dhp.mastering_display_actual_peak_luminance_flag)
}

// Custom: GetNumRowsMasteringDisplayActualPeakLuminance gets `AVDynamicHDRPlus.num_rows_mastering_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) GetNumRowsMasteringDisplayActualPeakLuminance() uint8 {
	return (uint8)(dhp.num_rows_mastering_display_actual_peak_luminance)
}

// Custom: SetNumRowsMasteringDisplayActualPeakLuminance sets `AVDynamicHDRPlus.num_rows_mastering_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) SetNumRowsMasteringDisplayActualPeakLuminance(v uint8) {
	dhp.num_rows_mastering_display_actual_peak_luminance = (C.uint8_t)(v)
}

// Custom: GetNumRowsMasteringDisplayActualPeakLuminanceAddr gets `AVDynamicHDRPlus.num_rows_mastering_display_actual_peak_luminance` address.
func (dhp *AVDynamicHDRPlus) GetNumRowsMasteringDisplayActualPeakLuminanceAddr() *uint8 {
	return (*uint8)(&dhp.num_rows_mastering_display_actual_peak_luminance)
}

// Custom: GetNumColsMasteringDisplayActualPeakLuminance gets `AVDynamicHDRPlus.num_cols_mastering_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) GetNumColsMasteringDisplayActualPeakLuminance() uint8 {
	return (uint8)(dhp.num_cols_mastering_display_actual_peak_luminance)
}

// Custom: SetNumColsMasteringDisplayActualPeakLuminance sets `AVDynamicHDRPlus.num_cols_mastering_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) SetNumColsMasteringDisplayActualPeakLuminance(v uint8) {
	dhp.num_cols_mastering_display_actual_peak_luminance = (C.uint8_t)(v)
}

// Custom: GetNumColsMasteringDisplayActualPeakLuminanceAddr gets `AVDynamicHDRPlus.num_cols_mastering_display_actual_peak_luminance` address.
func (dhp *AVDynamicHDRPlus) GetNumColsMasteringDisplayActualPeakLuminanceAddr() *uint8 {
	return (*uint8)(&dhp.num_cols_mastering_display_actual_peak_luminance)
}

// Custom: GetMasteringDisplayActualPeakLuminance gets `AVDynamicHDRPlus.mastering_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) GetMasteringDisplayActualPeakLuminance() (v [][]AVRational) {
	for i := 0; i < 25; i++ {
		v = append(v, unsafe.Slice((*AVRational)(&dhp.mastering_display_actual_peak_luminance[i][0]), 25))
	}
	return v
}

// Custom: SetMasteringDisplayActualPeakLuminance sets `AVDynamicHDRPlus.mastering_display_actual_peak_luminance` value.
func (dhp *AVDynamicHDRPlus) SetMasteringDisplayActualPeakLuminance(v [][]AVRational) {
	for i := 0; i < FFMIN(len(v), 25); i++ {
		for j := 0; j < FFMIN(len(v[i]), 25); j++ {
			dhp.mastering_display_actual_peak_luminance[i][j] = (C.struct_AVRational)(v[i][j])
		}
	}
}

// Custom: GetMasteringDisplayActualPeakLuminanceAddr gets `AVDynamicHDRPlus.mastering_display_actual_peak_luminance` address.
func (dhp *AVDynamicHDRPlus) GetMasteringDisplayActualPeakLuminanceAddr() **AVRational {
	return (**AVRational)(unsafe.Pointer(&dhp.mastering_display_actual_peak_luminance))
}

// AvDynamicHdrPlusAlloc allocates an AVDynamicHDRPlus structure and set its fields to
// default values.
func AvDynamicHdrPlusAlloc(size *uintptr) *AVDynamicHDRPlus {
	return (*AVDynamicHDRPlus)(C.av_dynamic_hdr_plus_alloc((*C.size_t)(unsafe.Pointer(size))))
}

// AvDynamicHdrPlusCreateSideData allocates a complete AVDynamicHDRPlus and add it to the frame.
func AvDynamicHdrPlusCreateSideData(frame *AVFrame) *AVDynamicHDRPlus {
	return (*AVDynamicHDRPlus)(C.av_dynamic_hdr_plus_create_side_data((*C.struct_AVFrame)(frame)))
}
