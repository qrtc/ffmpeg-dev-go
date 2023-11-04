// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/cpu.h>
*/
import "C"

const (
	AV_CPU_FLAG_FORCE    = C.AV_CPU_FLAG_FORCE
	AV_CPU_FLAG_MMX      = C.AV_CPU_FLAG_MMX
	AV_CPU_FLAG_MMXEXT   = C.AV_CPU_FLAG_MMXEXT
	AV_CPU_FLAG_MMX2     = C.AV_CPU_FLAG_MMX2
	AV_CPU_FLAG_3DNOW    = C.AV_CPU_FLAG_3DNOW
	AV_CPU_FLAG_SSE      = C.AV_CPU_FLAG_SSE
	AV_CPU_FLAG_SSE2     = C.AV_CPU_FLAG_SSE2
	AV_CPU_FLAG_SSE2SLOW = C.AV_CPU_FLAG_SSE2SLOW

	AV_CPU_FLAG_3DNOWEXT = C.AV_CPU_FLAG_3DNOWEXT
	AV_CPU_FLAG_SSE3     = C.AV_CPU_FLAG_SSE3
	AV_CPU_FLAG_SSE3SLOW = C.AV_CPU_FLAG_SSE3SLOW

	AV_CPU_FLAG_SSSE3     = C.AV_CPU_FLAG_SSSE3
	AV_CPU_FLAG_SSSE3SLOW = C.AV_CPU_FLAG_SSSE3SLOW
	AV_CPU_FLAG_ATOM      = C.AV_CPU_FLAG_ATOM
	AV_CPU_FLAG_SSE4      = C.AV_CPU_FLAG_SSE4
	AV_CPU_FLAG_SSE42     = C.AV_CPU_FLAG_SSE42
	AV_CPU_FLAG_AESNI     = C.AV_CPU_FLAG_AESNI
	AV_CPU_FLAG_AVX       = C.AV_CPU_FLAG_AVX
	AV_CPU_FLAG_AVXSLOW   = C.AV_CPU_FLAG_AVXSLOW
	AV_CPU_FLAG_XOP       = C.AV_CPU_FLAG_XOP
	AV_CPU_FLAG_FMA4      = C.AV_CPU_FLAG_FMA4
	AV_CPU_FLAG_CMOV      = C.AV_CPU_FLAG_CMOV
	AV_CPU_FLAG_AVX2      = C.AV_CPU_FLAG_AVX2
	AV_CPU_FLAG_FMA3      = C.AV_CPU_FLAG_FMA3
	AV_CPU_FLAG_BMI1      = C.AV_CPU_FLAG_BMI1
	AV_CPU_FLAG_BMI2      = C.AV_CPU_FLAG_BMI2
	AV_CPU_FLAG_AVX512    = C.AV_CPU_FLAG_AVX512

	AV_CPU_FLAG_ALTIVEC = C.AV_CPU_FLAG_ALTIVEC
	AV_CPU_FLAG_VSX     = C.AV_CPU_FLAG_VSX
	AV_CPU_FLAG_POWER8  = C.AV_CPU_FLAG_POWER8

	AV_CPU_FLAG_ARMV5TE = C.AV_CPU_FLAG_ARMV5TE
	AV_CPU_FLAG_ARMV6   = C.AV_CPU_FLAG_ARMV6
	AV_CPU_FLAG_ARMV6T2 = C.AV_CPU_FLAG_ARMV6T2
	AV_CPU_FLAG_VFP     = C.AV_CPU_FLAG_VFP
	AV_CPU_FLAG_VFPV3   = C.AV_CPU_FLAG_VFPV3
	AV_CPU_FLAG_NEON    = C.AV_CPU_FLAG_NEON
	AV_CPU_FLAG_ARMV8   = C.AV_CPU_FLAG_ARMV8
	AV_CPU_FLAG_VFP_VM  = C.AV_CPU_FLAG_VFP_VM
	AV_CPU_FLAG_SETEND  = C.AV_CPU_FLAG_SETEND
)

// AvGetCpuFlags returns the flags which specify extensions supported by the CPU.
func AvGetCpuFlags() int32 {
	return (int32)(C.av_get_cpu_flags())
}

// AvForceCpuFlags disables cpu detection and forces the specified flags.
func AvForceCpuFlags(flags int32) {
	C.av_force_cpu_flags((C.int)(flags))
}

// Deprecated: Use AvForceCpuFlags() and AvGetCpuFlags() instead.
//
// AvSetCpuFlagsMask set a mask on flags returned by AvGetCpuFlags().
func AvSetCpuFlagsMask(mask int32) {
	C.av_set_cpu_flags_mask((C.int)(mask))
}

// Deprecated: Use AvParseCpuCaps() when possible.
//
// AvParseCpuFlags parses CPU flags from a string.
func AvParseCpuFlags(str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.av_parse_cpu_flags((*C.char)(strPtr)))
}

// AvParseCpuCaps parses CPU caps from a string and update the given AV_CPU_* flags based on that.
func AvParseCpuCaps(flags *uint32, str string) int32 {
	strPtr, strFunc := StringCasting(str)
	defer strFunc()
	return (int32)(C.av_parse_cpu_caps((*C.uint)(flags), (*C.char)(strPtr)))
}

// AvCpuCount returns the number of logical CPU cores present.
func AvCpuCount() int32 {
	return (int32)(C.av_cpu_count())
}

// AvCpuMaxAlign gets the maximum data alignment that may be required by FFmpeg.
func AvCpuMaxAlign() int32 {
	return (int32)(C.av_cpu_max_align())
}
