package ffmpeg

/*
#include <libavcodec/dv_profile.h>
*/
import "C"
import "unsafe"

const (
	DV_PROFILE_BYTES = C.DV_PROFILE_BYTES
)

// AVDVProfile
type AVDVProfile C.struct_AVDVProfile

// Custom: GetDsf gets `AVDVProfile.dsf` value.
func (dvp *AVDVProfile) GetDsf() int32 {
	return (int32)(dvp.dsf)
}

// Custom: SetDsf sets `AVDVProfile.dsf` value.
func (dvp *AVDVProfile) SetDsf(v int32) {
	dvp.dsf = (C.int)(v)
}

// Custom: GetDsfAddr gets `AVDVProfile.dsf` address.
func (dvp *AVDVProfile) GetDsfAddr() *int32 {
	return (*int32)(&dvp.dsf)
}

// Custom: GetVideoStype gets `AVDVProfile.video_stype` value.
func (dvp *AVDVProfile) GetVideoStype() int32 {
	return (int32)(dvp.video_stype)
}

// Custom: SetVideoStype sets `AVDVProfile.video_stype` value.
func (dvp *AVDVProfile) SetVideoStype(v int32) {
	dvp.video_stype = (C.int)(v)
}

// Custom: GetVideoStypeAddr gets `AVDVProfile.video_stype` address.
func (dvp *AVDVProfile) GetVideoStypeAddr() *int32 {
	return (*int32)(&dvp.video_stype)
}

// Custom: GetFrameSize gets `AVDVProfile.frame_size` value.
func (dvp *AVDVProfile) GetFrameSize() int32 {
	return (int32)(dvp.frame_size)
}

// Custom: SetFrameSize sets `AVDVProfile.frame_size` value.
func (dvp *AVDVProfile) SetFrameSize(v int32) {
	dvp.frame_size = (C.int)(v)
}

// Custom: GetFrameSizeAddr gets `AVDVProfile.frame_size` address.
func (dvp *AVDVProfile) GetFrameSizeAddr() *int32 {
	return (*int32)(&dvp.frame_size)
}

// Custom: GetDifsegSize gets `AVDVProfile.difseg_size` value.
func (dvp *AVDVProfile) GetDifsegSize() int32 {
	return (int32)(dvp.difseg_size)
}

// Custom: SetDifsegSize sets `AVDVProfile.difseg_size` value.
func (dvp *AVDVProfile) SetDifsegSize(v int32) {
	dvp.difseg_size = (C.int)(v)
}

// Custom: GetDifsegSizeAddr gets `AVDVProfile.difseg_size` address.
func (dvp *AVDVProfile) GetDifsegSizeAddr() *int32 {
	return (*int32)(&dvp.difseg_size)
}

// Custom: GetNDifchan gets `AVDVProfile.n_difchan` value.
func (dvp *AVDVProfile) GetNDifchan() int32 {
	return (int32)(dvp.n_difchan)
}

// Custom: SetNDifchan sets `AVDVProfile.n_difchan` value.
func (dvp *AVDVProfile) SetNDifchan(v int32) {
	dvp.n_difchan = (C.int)(v)
}

// Custom: GetNDifchanAddr gets `AVDVProfile.n_difchan` address.
func (dvp *AVDVProfile) GetNDifchanAddr() *int32 {
	return (*int32)(&dvp.n_difchan)
}

// Custom: GetTimeBase gets `AVDVProfile.time_base` value.
func (dvp *AVDVProfile) GetTimeBase() AVRational {
	return (AVRational)(dvp.time_base)
}

// Custom: SetTimeBase sets `AVDVProfile.time_base` value.
func (dvp *AVDVProfile) SetTimeBase(v AVRational) {
	dvp.time_base = (C.struct_AVRational)(v)
}

// Custom: GetTimeBaseAddr gets `AVDVProfile.time_base` address.
func (dvp *AVDVProfile) GetTimeBaseAddr() *AVRational {
	return (*AVRational)(&dvp.time_base)
}

// Custom: GetLtcDivisor gets `AVDVProfile.ltc_divisor` value.
func (dvp *AVDVProfile) GetLtcDivisor() int32 {
	return (int32)(dvp.ltc_divisor)
}

// Custom: SetLtcDivisor sets `AVDVProfile.ltc_divisor` value.
func (dvp *AVDVProfile) SetLtcDivisor(v int32) {
	dvp.ltc_divisor = (C.int)(v)
}

// Custom: GetLtcDivisorAddr gets `AVDVProfile.ltc_divisor` address.
func (dvp *AVDVProfile) GetLtcDivisorAddr() *int32 {
	return (*int32)(&dvp.ltc_divisor)
}

// Custom: GetHeight gets `AVDVProfile.height` value.
func (dvp *AVDVProfile) GetHeight() int32 {
	return (int32)(dvp.height)
}

// Custom: SetHeight sets `AVDVProfile.height` value.
func (dvp *AVDVProfile) SetHeight(v int32) {
	dvp.height = (C.int)(v)
}

// Custom: GetHeightAddr gets `AVDVProfile.height` address.
func (dvp *AVDVProfile) GetHeightAddr() *int32 {
	return (*int32)(&dvp.height)
}

// Custom: GetWidth gets `AVDVProfile.width` value.
func (dvp *AVDVProfile) GetWidth() int32 {
	return (int32)(dvp.width)
}

// Custom: SetWidth sets `AVDVProfile.width` value.
func (dvp *AVDVProfile) SetWidth(v int32) {
	dvp.width = (C.int)(v)
}

// Custom: GetWidthAddr gets `AVDVProfile.width` address.
func (dvp *AVDVProfile) GetWidthAddr() *int32 {
	return (*int32)(&dvp.width)
}

// Custom: GetSar gets `AVDVProfile.sar` value.
func (dvp *AVDVProfile) GetSar() []AVRational {
	return unsafe.Slice((*AVRational)(&dvp.sar[0]), 2)
}

// Custom: SetSar sets `AVDVProfile.sar` value.
func (dvp *AVDVProfile) SetSar(v []AVRational) {
	for i := 0; i < FFMIN(len(v), 2); i++ {
		dvp.sar[i] = (C.struct_AVRational)(v[i])
	}
}

// Custom: GetSarAddr gets `AVDVProfile.sar` address.
func (dvp *AVDVProfile) GetSarAddr() **AVRational {
	return (**AVRational)(unsafe.Pointer(&dvp.sar))
}

// Custom: GetPixFmt gets `AVDVProfile.pix_fmt` value.
func (dvp *AVDVProfile) GetPixFmt() AVPixelFormat {
	return (AVPixelFormat)(dvp.pix_fmt)
}

// Custom: SetPixFmt sets `AVDVProfile.pix_fmt` value.
func (dvp *AVDVProfile) SetPixFmt(v AVPixelFormat) {
	dvp.pix_fmt = (C.enum_AVPixelFormat)(v)
}

// Custom: GetPixFmtAddr gets `AVDVProfile.pix_fmt` address.
func (dvp *AVDVProfile) GetPixFmtAddr() *AVPixelFormat {
	return (*AVPixelFormat)(&dvp.pix_fmt)
}

// Custom: GetBpm gets `AVDVProfile.bpm` value.
func (dvp *AVDVProfile) GetBpm() int32 {
	return (int32)(dvp.bpm)
}

// Custom: SetBpm sets `AVDVProfile.bpm` value.
func (dvp *AVDVProfile) SetBpm(v int32) {
	dvp.bpm = (C.int)(v)
}

// Custom: GetBpmAddr gets `AVDVProfile.bpm` address.
func (dvp *AVDVProfile) GetBpmAddr() *int32 {
	return (*int32)(&dvp.bpm)
}

// Custom: GetBlockSizes gets `AVDVProfile.block_sizes` value.
func (dvp *AVDVProfile) GetBlockSizes() *uint8 {
	return (*uint8)(dvp.block_sizes)
}

// Custom: SetBlockSizes sets `AVDVProfile.block_sizes` value.
func (dvp *AVDVProfile) SetBlockSizes(v *uint8) {
	dvp.block_sizes = (*C.uint8_t)(v)
}

// Custom: GetBlockSizesAddr gets `AVDVProfile.block_sizes` address.
func (dvp *AVDVProfile) GetBlockSizesAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&dvp.block_sizes))
}

// Custom: GetAudioStride gets `AVDVProfile.audio_stride` value.
func (dvp *AVDVProfile) GetAudioStride() int32 {
	return (int32)(dvp.audio_stride)
}

// Custom: SetAudioStride sets `AVDVProfile.audio_stride` value.
func (dvp *AVDVProfile) SetAudioStride(v int32) {
	dvp.audio_stride = (C.int)(v)
}

// Custom: GetAudioStrideAddr gets `AVDVProfile.audio_stride` address.
func (dvp *AVDVProfile) GetAudioStrideAddr() *int32 {
	return (*int32)(&dvp.audio_stride)
}

// Custom: GetAudioMinSamples gets `AVDVProfile.audio_min_samples` value.
func (dvp *AVDVProfile) GetAudioMinSamples() []int32 {
	return unsafe.Slice((*int32)(&dvp.audio_min_samples[0]), 3)
}

// Custom: SetAudioMinSamples sets `AVDVProfile.audio_min_samples` value.
func (dvp *AVDVProfile) SetAudioMinSamples(v []int32) {
	for i := 0; i < FFMIN(len(v), 3); i++ {
		dvp.audio_min_samples[i] = (C.int)(v[i])
	}
}

// Custom: GetAudioMinSamplesAddr gets `AVDVProfile.audio_min_samples` address.
func (dvp *AVDVProfile) GetAudioMinSamplesAddr() **int32 {
	return (**int32)(unsafe.Pointer(&dvp.audio_min_samples))
}

// Custom: GetAudioSamplesDist gets `AVDVProfile.audio_samples_dist` value.
func (dvp *AVDVProfile) GetAudioSamplesDist() []int32 {
	return unsafe.Slice((*int32)(&dvp.audio_samples_dist[0]), 5)
}

// Custom: SetAudioSamplesDist sets `AVDVProfile.audio_samples_dist` value.
func (dvp *AVDVProfile) SetAudioSamplesDist(v []int32) {
	for i := 0; i < FFMIN(len(v), 5); i++ {
		dvp.audio_samples_dist[i] = (C.int)(v[i])
	}
}

// Custom: GetAudioSamplesDistAddr gets `AVDVProfile.audio_samples_dist` address.
func (dvp *AVDVProfile) GetAudioSamplesDistAddr() **int32 {
	return (**int32)(unsafe.Pointer(&dvp.audio_samples_dist))
}

// Custom: GetAudioShuffle gets `AVDVProfile.audio_shuffle` value.
func (dvp *AVDVProfile) GetAudioShuffle() []uint8 {
	return unsafe.Slice((*uint8)(&dvp.audio_shuffle[0]), 9)
}

// Custom: SetAudioShuffle sets `AVDVProfile.audio_shuffle` value.
func (dvp *AVDVProfile) SetAudioShuffle(v []uint8) {
	for i := 0; i < FFMIN(len(v), 9); i++ {
		dvp.audio_shuffle[i] = (C.uint8_t)(v[i])
	}
}

// Custom: GetAudioShuffleAddr gets `AVDVProfile.audio_shuffle` address.
func (dvp *AVDVProfile) GetAudioShuffleAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&dvp.audio_shuffle))
}

// AvDvFrameProfile gets a DV profile for the provided compressed frame.
func AvDvFrameProfile(sys *AVDVProfile, frame *uint8, bufSize uint32) *AVDVProfile {
	return (*AVDVProfile)(C.av_dv_frame_profile((*C.struct_AVDVProfile)(sys),
		(*C.uint8_t)(frame), (C.uint)(bufSize)))
}

// AvDvCodecProfile gets a DV profile for the provided stream parameters.
func AvDvCodecProfile(width, height int32, pixFmt AVPixelFormat) *AVDVProfile {
	return (*AVDVProfile)(C.av_dv_codec_profile((C.int)(width), (C.int)(height),
		(C.enum_AVPixelFormat)(pixFmt)))
}

// AvDvCodecProfile2 gets a DV profile for the provided stream parameters.
func AvDvCodecProfile2(width, height int32, pixFmt AVPixelFormat,
	frameRate AVRational) *AVDVProfile {
	return (*AVDVProfile)(C.av_dv_codec_profile2((C.int)(width), (C.int)(height),
		(C.enum_AVPixelFormat)(pixFmt), (C.struct_AVRational)(frameRate)))
}
