package ffmpeg

/*
#include <libavutil/spherical.h>
*/
import "C"
import "unsafe"

type AVSphericalProjection = C.enum_AVSphericalProjection

const (
	AV_SPHERICAL_EQUIRECTANGULAR      = AVSphericalProjection(C.AV_SPHERICAL_EQUIRECTANGULAR)
	AV_SPHERICAL_CUBEMAP              = AVSphericalProjection(C.AV_SPHERICAL_CUBEMAP)
	AV_SPHERICAL_EQUIRECTANGULAR_TILE = AVSphericalProjection(C.AV_SPHERICAL_EQUIRECTANGULAR_TILE)
)

type AVSphericalMapping C.struct_AVSphericalMapping

// Custom: GetProjection gets `AVSphericalMapping.projection` value.
func (smp *AVSphericalMapping) GetProjection() AVSphericalProjection {
	return (AVSphericalProjection)(smp.projection)
}

// Custom: SetProjection sets `AVSphericalMapping.projection` value.
func (smp *AVSphericalMapping) SetProjection(v AVSphericalProjection) {
	smp.projection = (C.enum_AVSphericalProjection)(v)
}

// Custom: GetProjectionAddr gets `AVSphericalMapping.projection` address.
func (smp *AVSphericalMapping) GetProjectionAddr() *AVSphericalProjection {
	return (*AVSphericalProjection)(&smp.projection)
}

// Custom: GetYaw gets `AVSphericalMapping.yaw` value.
func (smp *AVSphericalMapping) GetYaw() int32 {
	return (int32)(smp.yaw)
}

// Custom: SetYaw sets `AVSphericalMapping.yaw` value.
func (smp *AVSphericalMapping) SetYaw(v int32) {
	smp.yaw = (C.int32_t)(v)
}

// Custom: GetYawAddr gets `AVSphericalMapping.yaw` address.
func (smp *AVSphericalMapping) GetYawAddr() *int32 {
	return (*int32)(&smp.yaw)
}

// Custom: GetPitch gets `AVSphericalMapping.pitch` value.
func (smp *AVSphericalMapping) GetPitch() int32 {
	return (int32)(smp.pitch)
}

// Custom: SetPitch sets `AVSphericalMapping.pitch` value.
func (smp *AVSphericalMapping) SetPitch(v int32) {
	smp.pitch = (C.int32_t)(v)
}

// Custom: GetPitchAddr gets `AVSphericalMapping.pitch` address.
func (smp *AVSphericalMapping) GetPitchAddr() *int32 {
	return (*int32)(&smp.pitch)
}

// Custom: GetRoll gets `AVSphericalMapping.roll` value.
func (smp *AVSphericalMapping) GetRoll() int32 {
	return (int32)(smp.roll)
}

// Custom: SetRoll sets `AVSphericalMapping.roll` value.
func (smp *AVSphericalMapping) SetRoll(v int32) {
	smp.roll = (C.int32_t)(v)
}

// Custom: GetRollAddr gets `AVSphericalMapping.roll` address.
func (smp *AVSphericalMapping) GetRollAddr() *int32 {
	return (*int32)(&smp.roll)
}

// Custom: GetBoundLeft gets `AVSphericalMapping.bound_left` value.
func (smp *AVSphericalMapping) GetBoundLeft() uint32 {
	return (uint32)(smp.bound_left)
}

// Custom: SetBoundLeft sets `AVSphericalMapping.bound_left` value.
func (smp *AVSphericalMapping) SetBoundLeft(v uint32) {
	smp.bound_left = (C.uint32_t)(v)
}

// Custom: GetBoundLeftAddr gets `AVSphericalMapping.bound_left` address.
func (smp *AVSphericalMapping) GetBoundLeftAddr() *uint32 {
	return (*uint32)(&smp.bound_left)
}

// Custom: GetBoundTop gets `AVSphericalMapping.bound_top` value.
func (smp *AVSphericalMapping) GetBoundTop() uint32 {
	return (uint32)(smp.bound_top)
}

// Custom: SetBoundTop sets `AVSphericalMapping.bound_top` value.
func (smp *AVSphericalMapping) SetBoundTop(v uint32) {
	smp.bound_top = (C.uint32_t)(v)
}

// Custom: GetBoundTopAddr gets `AVSphericalMapping.bound_top` address.
func (smp *AVSphericalMapping) GetBoundTopAddr() *uint32 {
	return (*uint32)(&smp.bound_top)
}

// Custom: GetBoundRight gets `AVSphericalMapping.bound_right` value.
func (smp *AVSphericalMapping) GetBoundRight() uint32 {
	return (uint32)(smp.bound_right)
}

// Custom: SetBoundRight sets `AVSphericalMapping.bound_right` value.
func (smp *AVSphericalMapping) SetBoundRight(v uint32) {
	smp.bound_right = (C.uint32_t)(v)
}

// Custom: GetBoundRightAddr gets `AVSphericalMapping.bound_right` address.
func (smp *AVSphericalMapping) GetBoundRightAddr() *uint32 {
	return (*uint32)(&smp.bound_right)
}

// Custom: GetBoundBottom gets `AVSphericalMapping.bound_bottom` value.
func (smp *AVSphericalMapping) GetBoundBottom() uint32 {
	return (uint32)(smp.bound_bottom)
}

// Custom: SetBoundBottom sets `AVSphericalMapping.bound_bottom` value.
func (smp *AVSphericalMapping) SetBoundBottom(v uint32) {
	smp.bound_bottom = (C.uint32_t)(v)
}

// Custom: GetBoundBottomAddr gets `AVSphericalMapping.bound_bottom` address.
func (smp *AVSphericalMapping) GetBoundBottomAddr() *uint32 {
	return (*uint32)(&smp.bound_bottom)
}

// Custom: GetPadding gets `AVSphericalMapping.padding` value.
func (smp *AVSphericalMapping) GetPadding() uint32 {
	return (uint32)(smp.padding)
}

// Custom: SetPadding sets `AVSphericalMapping.padding` value.
func (smp *AVSphericalMapping) SetPadding(v uint32) {
	smp.padding = (C.uint32_t)(v)
}

// Custom: GetPaddingAddr gets `AVSphericalMapping.padding` address.
func (smp *AVSphericalMapping) GetPaddingAddr() *uint32 {
	return (*uint32)(&smp.padding)
}

// AvSphericalAlloc allocate a AVSphericalVideo structure and initialize its fields to default
// values.
func AvSphericalAlloc(size *uintptr) *AVSphericalMapping {
	return (*AVSphericalMapping)(C.av_spherical_alloc((*C.size_t)(unsafe.Pointer(size))))
}

// AvSphericalTileBounds converts the bounding fields from an AVSphericalVideo
// from 0.32 fixed point to pixels.
func AvSphericalTileBounds(_map *AVSphericalMapping,
	width, height uintptr,
	left, top, right, bottom *uintptr) {
	C.av_spherical_tile_bounds((*C.struct_AVSphericalMapping)(_map),
		(C.size_t)(width), (C.size_t)(height),
		(*C.size_t)(unsafe.Pointer(left)), (*C.size_t)(unsafe.Pointer(top)),
		(*C.size_t)(unsafe.Pointer(right)), (*C.size_t)(unsafe.Pointer(bottom)))
}

// AvSphericalProjectionName provides a human-readable name of a given AVSphericalProjection.
func AvSphericalProjectionName(projection AVSphericalProjection) string {
	return C.GoString(C.av_spherical_projection_name((C.enum_AVSphericalProjection)(projection)))
}

// AvSphericalFromName gets the AVSphericalProjection form a human-readable name.
func AvSphericalFromName(name string) int32 {
	namePtr, nameFunc := StringCasting(name)
	defer nameFunc()
	return (int32)(C.av_spherical_from_name((*C.char)(namePtr)))
}
