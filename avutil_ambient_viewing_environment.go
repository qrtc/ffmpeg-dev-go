// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/ambient_viewing_environment.h>
*/
import "C"
import "unsafe"

type AVAmbientViewingEnvironment C.struct_AVAmbientViewingEnvironment

// GetAmbientIlluminance gets `AVAmbientViewingEnvironment.ambient_illuminance` value.
func (ave *AVAmbientViewingEnvironment) GetAmbientIlluminance() AVRational {
	return (AVRational)(ave.ambient_illuminance)
}

// SetAmbientIlluminance sets `AVAmbientViewingEnvironment.ambient_illuminance` value.
func (ave *AVAmbientViewingEnvironment) SetAmbientIlluminance(v AVRational) {
	ave.ambient_illuminance = (C.struct_AVRational)(v)
}

// GetAmbientIlluminanceAddr gets `AVAmbientViewingEnvironment.ambient_illuminance` address.
func (ave *AVAmbientViewingEnvironment) GetAmbientIlluminanceAddr() *AVRational {
	return (*AVRational)(&ave.ambient_illuminance)
}

// GetAmbientLightX gets `AVAmbientViewingEnvironment.ambient_light_x` value.
func (ave *AVAmbientViewingEnvironment) GetAmbientLightX() AVRational {
	return (AVRational)(ave.ambient_light_x)
}

// SetAmbientLightX sets `AVAmbientViewingEnvironment.ambient_light_x` value.
func (ave *AVAmbientViewingEnvironment) SetAmbientLightX(v AVRational) {
	ave.ambient_light_x = (C.struct_AVRational)(v)
}

// GetAmbientLightXAddr gets `AVAmbientViewingEnvironment.ambient_light_x` address.
func (ave *AVAmbientViewingEnvironment) GetAmbientLightXAddr() *AVRational {
	return (*AVRational)(&ave.ambient_light_x)
}

// GetAmbientLightY gets `AVAmbientViewingEnvironment.ambient_light_y` value.
func (ave *AVAmbientViewingEnvironment) GetAmbientLightY() AVRational {
	return (AVRational)(ave.ambient_light_y)
}

// SetAmbientLightY sets `AVAmbientViewingEnvironment.ambient_light_y` value.
func (ave *AVAmbientViewingEnvironment) SetAmbientLightY(v AVRational) {
	ave.ambient_light_y = (C.struct_AVRational)(v)
}

// GetAmbientLightYAddr gets `AVAmbientViewingEnvironment.ambient_light_y` address.
func (ave *AVAmbientViewingEnvironment) GetAmbientLightYAddr() *AVRational {
	return (*AVRational)(&ave.ambient_light_y)
}

// AvAmbientViewingEnvironmentAlloc allocates an AVAmbientViewingEnvironment structure.
func AvAmbientViewingEnvironmentAlloc(size *uintptr) *AVAmbientViewingEnvironment {
	return (*AVAmbientViewingEnvironment)(C.av_ambient_viewing_environment_alloc(
		(*C.size_t)(unsafe.Pointer(size))))
}

// AvAmbientViewingEnvironmentCreateSideData
func AvAmbientViewingEnvironmentCreateSideData(frame *AVFrame) *AVAmbientViewingEnvironment {
	return (*AVAmbientViewingEnvironment)(C.av_ambient_viewing_environment_create_side_data(
		(*C.struct_AVFrame)(frame)))
}
