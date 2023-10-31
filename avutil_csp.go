// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/csp.h>
*/
import "C"

// AVLumaCoefficients
type AVLumaCoefficients C.struct_AVLumaCoefficients

// GetCr gets `AVLumaCoefficients.cr` value.
func (lc *AVLumaCoefficients) GetCr() AVRational {
	return (AVRational)(lc.cr)
}

// SetCr sets `AVLumaCoefficients.cr` value.
func (lc *AVLumaCoefficients) SetCr(v AVRational) {
	lc.cr = (C.struct_AVRational)(v)
}

// GetCrAddr gets `AVLumaCoefficients.cr` address.
func (lc *AVLumaCoefficients) GetCrAddr() *AVRational {
	return (*AVRational)(&lc.cr)
}

// GetCg gets `AVLumaCoefficients.cg` value.
func (lc *AVLumaCoefficients) GetCg() AVRational {
	return (AVRational)(lc.cg)
}

// SetCg sets `AVLumaCoefficients.cg` value.
func (lc *AVLumaCoefficients) SetCg(v AVRational) {
	lc.cg = (C.struct_AVRational)(v)
}

// GetCgAddr gets `AVLumaCoefficients.cg` address.
func (lc *AVLumaCoefficients) GetCgAddr() *AVRational {
	return (*AVRational)(&lc.cg)
}

// GetCb gets `AVLumaCoefficients.cb` value.
func (lc *AVLumaCoefficients) GetCb() AVRational {
	return (AVRational)(lc.cb)
}

// SetCb sets `AVLumaCoefficients.cb` value.
func (lc *AVLumaCoefficients) SetCb(v AVRational) {
	lc.cb = (C.struct_AVRational)(v)
}

// GetCbAddr gets `AVLumaCoefficients.cb` address.
func (lc *AVLumaCoefficients) GetCbAddr() *AVRational {
	return (*AVRational)(&lc.cb)
}

// AVCIExy
type AVCIExy C.struct_AVCIExy

// GetX gets `AVCIExy.x` value.
func (cxy *AVCIExy) GetX() AVRational {
	return (AVRational)(cxy.x)
}

// SetX sets `AVCIExy.x` value.
func (cxy *AVCIExy) SetX(v AVRational) {
	cxy.x = (C.struct_AVRational)(v)
}

// GetXAddr gets `AVCIExy.x` address.
func (cxy *AVCIExy) GetXAddr() *AVRational {
	return (*AVRational)(&cxy.x)
}

// GetY gets `AVCIExy.y` value.
func (cxy *AVCIExy) GetY() AVRational {
	return (AVRational)(cxy.y)
}

// SetY sets `AVCIExy.y` value.
func (cxy *AVCIExy) SetY(v AVRational) {
	cxy.y = (C.struct_AVRational)(v)
}

// GetYAddr gets `AVCIExy.y` address.
func (cxy *AVCIExy) GetYAddr() *AVRational {
	return (*AVRational)(&cxy.y)
}

// AVPrimaryCoefficients
type AVPrimaryCoefficients C.struct_AVPrimaryCoefficients

// GetR gets `AVPrimaryCoefficients.r` value.
func (pc *AVPrimaryCoefficients) GetR() AVCIExy {
	return (AVCIExy)(pc.r)
}

// SetR sets `AVPrimaryCoefficients.r` value.
func (pc *AVPrimaryCoefficients) SetR(v AVCIExy) {
	pc.r = (C.struct_AVCIExy)(v)
}

// GetRAddr gets `AVPrimaryCoefficients.r` address.
func (pc *AVPrimaryCoefficients) GetRAddr() *AVCIExy {
	return (*AVCIExy)(&pc.r)
}

// GetG gets `AVPrimaryCoefficients.g` value.
func (pc *AVPrimaryCoefficients) GetG() AVCIExy {
	return (AVCIExy)(pc.g)
}

// SetG sets `AVPrimaryCoefficients.g` value.
func (pc *AVPrimaryCoefficients) SetG(v AVCIExy) {
	pc.g = (C.struct_AVCIExy)(v)
}

// GetGAddr gets `AVPrimaryCoefficients.g` address.
func (pc *AVPrimaryCoefficients) GetGAddr() *AVCIExy {
	return (*AVCIExy)(&pc.g)
}

// GetB gets `AVPrimaryCoefficients.b` value.
func (pc *AVPrimaryCoefficients) GetB() AVCIExy {
	return (AVCIExy)(pc.b)
}

// SetB sets `AVPrimaryCoefficients.b` value.
func (pc *AVPrimaryCoefficients) SetB(v AVCIExy) {
	pc.b = (C.struct_AVCIExy)(v)
}

// GetBAddr gets `AVPrimaryCoefficients.b` address.
func (pc *AVPrimaryCoefficients) GetBAddr() *AVCIExy {
	return (*AVCIExy)(&pc.b)
}

// AVWhitepointCoefficients
type AVWhitepointCoefficients = AVCIExy

// AVColorPrimariesDesc
type AVColorPrimariesDesc C.struct_AVColorPrimariesDesc

// GetWp gets `AVColorPrimariesDesc.wp` value.
func (cpd *AVColorPrimariesDesc) GetWp() AVWhitepointCoefficients {
	return (AVWhitepointCoefficients)(cpd.wp)
}

// SetWp sets `AVColorPrimariesDesc.wp` value.
func (cpd *AVColorPrimariesDesc) SetWp(v AVWhitepointCoefficients) {
	cpd.wp = (C.AVWhitepointCoefficients)(v)
}

// GetWpAddr gets `AVColorPrimariesDesc.wp` address.
func (cpd *AVColorPrimariesDesc) GetWpAddr() *AVWhitepointCoefficients {
	return (*AVWhitepointCoefficients)(&cpd.wp)
}

// GetPrim gets `AVColorPrimariesDesc.prim` value.
func (cpd *AVColorPrimariesDesc) GetPrim() AVPrimaryCoefficients {
	return (AVPrimaryCoefficients)(cpd.prim)
}

// SetPrim sets `AVColorPrimariesDesc.prim` value.
func (cpd *AVColorPrimariesDesc) SetPrim(v AVPrimaryCoefficients) {
	cpd.prim = (C.struct_AVPrimaryCoefficients)(v)
}

// GetPrimAddr gets `AVColorPrimariesDesc.prim` address.
func (cpd *AVColorPrimariesDesc) GetPrimAddr() *AVPrimaryCoefficients {
	return (*AVPrimaryCoefficients)(&cpd.prim)
}

// AvCspLumaCoeffsFromAvcsp retrieves the Luma coefficients necessary to construct a conversion matrix
// from an enum constant describing the colorspace.
func AvCspLumaCoeffsFromAvcsp(csp AVColorSpace) *AVLumaCoefficients {
	return (*AVLumaCoefficients)(C.av_csp_luma_coeffs_from_avcsp((C.enum_AVColorSpace)(csp)))
}

// AvCspPrimariesDescFromId retrieves a complete gamut description from an enum constant describing the
// color primaries.
func AvCspPrimariesDescFromId(prm AVColorPrimaries) *AVColorPrimariesDesc {
	return (*AVColorPrimariesDesc)(C.av_csp_primaries_desc_from_id((C.enum_AVColorPrimaries)(prm)))
}

// AvCspPrimariesIdFromDesc detects which enum AVColorPrimaries constant corresponds to the given complete
// gamut description.
func AvCspPrimariesIdFromDesc(prm *AVColorPrimariesDesc) AVColorPrimaries {
	return (AVColorPrimaries)(C.av_csp_primaries_id_from_desc((*C.struct_AVColorPrimariesDesc)(prm)))
}
