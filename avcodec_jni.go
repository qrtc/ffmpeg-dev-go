// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavcodec/jni.h>
*/
import "C"

// AvJniSetJavaVm sets a Java virtual machine which will be used to retrieve the JNI environment.
func AvJniSetJavaVm(vm, logCtx CVoidPointer) int32 {
	return (int32)(C.av_jni_set_java_vm(VoidPointer(vm), VoidPointer(logCtx)))
}

// AvJniGetJavaVm gets the Java virtual machine which has been set with AvJniSetJavaVm.
func AvJniGetJavaVm(logCtx CVoidPointer) {
	C.av_jni_get_java_vm(VoidPointer(logCtx))
}
