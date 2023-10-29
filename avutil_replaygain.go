// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/replaygain.h>
*/
import "C"

// AVReplayGain
type AVReplayGain C.struct_AVReplayGain

// GetTrackGain gets `AVReplayGain.track_gain` value.
func (rg *AVReplayGain) GetTrackGain() int32 {
	return (int32)(rg.track_gain)
}

// SetTrackGain sets `AVReplayGain.track_gain` value.
func (rg *AVReplayGain) SetTrackGain(v int32) {
	rg.track_gain = (C.int32_t)(v)
}

// GetTrackGainAddr gets `AVReplayGain.track_gain` address.
func (rg *AVReplayGain) GetTrackGainAddr() *int32 {
	return (*int32)(&rg.track_gain)
}

// GetTrackPeak gets `AVReplayGain.track_peak` value.
func (rg *AVReplayGain) GetTrackPeak() uint32 {
	return (uint32)(rg.track_peak)
}

// SetTrackPeak sets `AVReplayGain.track_peak` value.
func (rg *AVReplayGain) SetTrackPeak(v uint32) {
	rg.track_peak = (C.uint32_t)(v)
}

// GetTrackPeakAddr gets `AVReplayGain.track_peak` address.
func (rg *AVReplayGain) GetTrackPeakAddr() *uint32 {
	return (*uint32)(&rg.track_peak)
}

// GetAlbumGain gets `AVReplayGain.album_gain` value.
func (rg *AVReplayGain) GetAlbumGain() int32 {
	return (int32)(rg.album_gain)
}

// SetAlbumGain sets `AVReplayGain.album_gain` value.
func (rg *AVReplayGain) SetAlbumGain(v int32) {
	rg.album_gain = (C.int32_t)(v)
}

// GetAlbumGainAddr gets `AVReplayGain.album_gain` address.
func (rg *AVReplayGain) GetAlbumGainAddr() *int32 {
	return (*int32)(&rg.album_gain)
}

// GetAlbumPeak gets `AVReplayGain.album_peak` value.
func (rg *AVReplayGain) GetAlbumPeak() uint32 {
	return (uint32)(rg.album_peak)
}

// SetAlbumPeak sets `AVReplayGain.album_peak` value.
func (rg *AVReplayGain) SetAlbumPeak(v uint32) {
	rg.album_peak = (C.uint32_t)(v)
}

// GetAlbumPeakAddr gets `AVReplayGain.album_peak` address.
func (rg *AVReplayGain) GetAlbumPeakAddr() *uint32 {
	return (*uint32)(&rg.album_peak)
}
