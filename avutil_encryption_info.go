// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/encryption_info.h>
*/
import "C"
import "unsafe"

// AVSubsampleEncryptionInfo
type AVSubsampleEncryptionInfo C.struct_AVSubsampleEncryptionInfo

// GetBytesOfClearData gets `AVSubsampleEncryptionInfo.bytes_of_clear_data` value.
func (sei *AVSubsampleEncryptionInfo) GetBytesOfClearData() uint32 {
	return (uint32)(sei.bytes_of_clear_data)
}

// SetBytesOfClearData sets `AVSubsampleEncryptionInfo.bytes_of_clear_data` value.
func (sei *AVSubsampleEncryptionInfo) SetBytesOfClearData(v uint32) {
	sei.bytes_of_clear_data = (C.uint32_t)(v)
}

// GetBytesOfClearDataAddr gets `AVSubsampleEncryptionInfo.bytes_of_clear_data` address.
func (sei *AVSubsampleEncryptionInfo) GetBytesOfClearDataAddr() *uint32 {
	return (*uint32)(&sei.bytes_of_clear_data)
}

// GetBytesOfProtectedData gets `AVSubsampleEncryptionInfo.bytes_of_protected_data` value.
func (sei *AVSubsampleEncryptionInfo) GetBytesOfProtectedData() uint32 {
	return (uint32)(sei.bytes_of_protected_data)
}

// SetBytesOfProtectedData sets `AVSubsampleEncryptionInfo.bytes_of_protected_data` value.
func (sei *AVSubsampleEncryptionInfo) SetBytesOfProtectedData(v uint32) {
	sei.bytes_of_protected_data = (C.uint32_t)(v)
}

// GetBytesOfProtectedDataAddr gets `AVSubsampleEncryptionInfo.bytes_of_protected_data` address.
func (sei *AVSubsampleEncryptionInfo) GetBytesOfProtectedDataAddr() *uint32 {
	return (*uint32)(&sei.bytes_of_protected_data)
}

// AVEncryptionInfo
type AVEncryptionInfo C.struct_AVEncryptionInfo

// GetScheme gets `AVEncryptionInfo.scheme` value.
func (ei *AVEncryptionInfo) GetScheme() uint32 {
	return (uint32)(ei.scheme)
}

// SetScheme sets `AVEncryptionInfo.scheme` value.
func (ei *AVEncryptionInfo) SetScheme(v uint32) {
	ei.scheme = (C.uint32_t)(v)
}

// GetSchemeAddr gets `AVEncryptionInfo.scheme` address.
func (ei *AVEncryptionInfo) GetSchemeAddr() *uint32 {
	return (*uint32)(&ei.scheme)
}

// GetCryptByteBlock gets `AVEncryptionInfo.crypt_byte_block` value.
func (ei *AVEncryptionInfo) GetCryptByteBlock() uint32 {
	return (uint32)(ei.crypt_byte_block)
}

// SetCryptByteBlock sets `AVEncryptionInfo.crypt_byte_block` value.
func (ei *AVEncryptionInfo) SetCryptByteBlock(v uint32) {
	ei.crypt_byte_block = (C.uint32_t)(v)
}

// GetCryptByteBlockAddr gets `AVEncryptionInfo.crypt_byte_block` address.
func (ei *AVEncryptionInfo) GetCryptByteBlockAddr() *uint32 {
	return (*uint32)(&ei.crypt_byte_block)
}

// GetSkipByteBlock gets `AVEncryptionInfo.skip_byte_block` value.
func (ei *AVEncryptionInfo) GetSkipByteBlock() uint32 {
	return (uint32)(ei.skip_byte_block)
}

// SetSkipByteBlock sets `AVEncryptionInfo.skip_byte_block` value.
func (ei *AVEncryptionInfo) SetSkipByteBlock(v uint32) {
	ei.skip_byte_block = (C.uint32_t)(v)
}

// GetSkipByteBlockAddr gets `AVEncryptionInfo.skip_byte_block` address.
func (ei *AVEncryptionInfo) GetSkipByteBlockAddr() *uint32 {
	return (*uint32)(&ei.skip_byte_block)
}

// GetKeyId gets `AVEncryptionInfo.key_id` value.
func (ei *AVEncryptionInfo) GetKeyId() *uint8 {
	return (*uint8)(ei.key_id)
}

// SetKeyId sets `AVEncryptionInfo.key_id` value.
func (ei *AVEncryptionInfo) SetKeyId(v *uint8) {
	ei.key_id = (*C.uint8_t)(v)
}

// GetKeyIdAddr gets `AVEncryptionInfo.key_id` address.
func (ei *AVEncryptionInfo) GetKeyIdAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&ei.key_id))
}

// GetKeyIdSize gets `AVEncryptionInfo.key_id_size` value.
func (ei *AVEncryptionInfo) GetKeyIdSize() uint32 {
	return (uint32)(ei.key_id_size)
}

// SetKeyIdSize sets `AVEncryptionInfo.key_id_size` value.
func (ei *AVEncryptionInfo) SetKeyIdSize(v uint32) {
	ei.key_id_size = (C.uint32_t)(v)
}

// GetKeyIdSizeAddr gets `AVEncryptionInfo.key_id_size` address.
func (ei *AVEncryptionInfo) GetKeyIdSizeAddr() *uint32 {
	return (*uint32)(&ei.key_id_size)
}

// GetIv gets `AVEncryptionInfo.iv` value.
func (ei *AVEncryptionInfo) GetIv() *uint8 {
	return (*uint8)(ei.iv)
}

// SetIv sets `AVEncryptionInfo.iv` value.
func (ei *AVEncryptionInfo) SetIv(v *uint8) {
	ei.iv = (*C.uint8_t)(v)
}

// GetIvAddr gets `AVEncryptionInfo.iv` address.
func (ei *AVEncryptionInfo) GetIvAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&ei.iv))
}

// GetIvSize gets `AVEncryptionInfo.iv_size` value.
func (ei *AVEncryptionInfo) GetIvSize() uint32 {
	return (uint32)(ei.iv_size)
}

// SetIvSize sets `AVEncryptionInfo.iv_size` value.
func (ei *AVEncryptionInfo) SetIvSize(v uint32) {
	ei.iv_size = (C.uint32_t)(v)
}

// GetIvSizeAddr gets `AVEncryptionInfo.iv_size` address.
func (ei *AVEncryptionInfo) GetIvSizeAddr() *uint32 {
	return (*uint32)(&ei.iv_size)
}

// GetSubsamples gets `AVEncryptionInfo.subsamples` value.
func (ei *AVEncryptionInfo) GetSubsamples() *AVSubsampleEncryptionInfo {
	return (*AVSubsampleEncryptionInfo)(ei.subsamples)
}

// SetSubsamples sets `AVEncryptionInfo.subsamples` value.
func (ei *AVEncryptionInfo) SetSubsamples(v *AVSubsampleEncryptionInfo) {
	ei.subsamples = (*C.struct_AVSubsampleEncryptionInfo)(v)
}

// GetSubsamplesAddr gets `AVEncryptionInfo.subsamples` address.
func (ei *AVEncryptionInfo) GetSubsamplesAddr() **AVSubsampleEncryptionInfo {
	return (**AVSubsampleEncryptionInfo)(unsafe.Pointer(&ei.subsamples))
}

// GetSubsampleCount gets `AVEncryptionInfo.subsample_count` value.
func (ei *AVEncryptionInfo) GetSubsampleCount() uint32 {
	return (uint32)(ei.subsample_count)
}

// SetSubsampleCount sets `AVEncryptionInfo.subsample_count` value.
func (ei *AVEncryptionInfo) SetSubsampleCount(v uint32) {
	ei.subsample_count = (C.uint32_t)(v)
}

// GetSubsampleCountAddr gets `AVEncryptionInfo.subsample_count` address.
func (ei *AVEncryptionInfo) GetSubsampleCountAddr() *uint32 {
	return (*uint32)(&ei.subsample_count)
}

// AVEncryptionInitInfo
type AVEncryptionInitInfo C.struct_AVEncryptionInitInfo

// GetSystemId gets `AVEncryptionInitInfo.system_id` value.
func (eii *AVEncryptionInitInfo) GetSystemId() *uint8 {
	return (*uint8)(eii.system_id)
}

// SetSystemId sets `AVEncryptionInitInfo.system_id` value.
func (eii *AVEncryptionInitInfo) SetSystemId(v *uint8) {
	eii.system_id = (*C.uint8_t)(v)
}

// GetSystemIdAddr gets `AVEncryptionInitInfo.system_id` address.
func (eii *AVEncryptionInitInfo) GetSystemIdAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&eii.system_id))
}

// GetSystemIdSize gets `AVEncryptionInitInfo.system_id_size` value.
func (eii *AVEncryptionInitInfo) GetSystemIdSize() uint32 {
	return (uint32)(eii.system_id_size)
}

// SetSystemIdSize sets `AVEncryptionInitInfo.system_id_size` value.
func (eii *AVEncryptionInitInfo) SetSystemIdSize(v uint32) {
	eii.system_id_size = (C.uint32_t)(v)
}

// GetSystemIdSizeAddr gets `AVEncryptionInitInfo.system_id_size` address.
func (eii *AVEncryptionInitInfo) GetSystemIdSizeAddr() *uint32 {
	return (*uint32)(&eii.system_id_size)
}

// GetKeyIds gets `AVEncryptionInitInfo.key_ids` value.
func (eii *AVEncryptionInitInfo) GetKeyIds() []*uint8 {
	return unsafe.Slice((**uint8)(unsafe.Pointer(eii.key_ids)), eii.key_id_size)
}

// SetKeyIds sets `AVEncryptionInitInfo.key_ids` value.
func (eii *AVEncryptionInitInfo) SetKeyIds(v **uint8) {
	eii.key_ids = (**C.uint8_t)(unsafe.Pointer(v))
}

// GetKeyIdsAddr gets `AVEncryptionInitInfo.key_ids` address.
func (eii *AVEncryptionInitInfo) GetKeyIdsAddr() ***uint8 {
	return (***uint8)(unsafe.Pointer(&eii.key_ids))
}

// GetNumKeyIds gets `AVEncryptionInitInfo.num_key_ids` value.
func (eii *AVEncryptionInitInfo) GetNumKeyIds() uint32 {
	return (uint32)(eii.num_key_ids)
}

// SetNumKeyIds sets `AVEncryptionInitInfo.num_key_ids` value.
func (eii *AVEncryptionInitInfo) SetNumKeyIds(v uint32) {
	eii.num_key_ids = (C.uint32_t)(v)
}

// GetNumKeyIdsAddr gets `AVEncryptionInitInfo.num_key_ids` address.
func (eii *AVEncryptionInitInfo) GetNumKeyIdsAddr() *uint32 {
	return (*uint32)(&eii.num_key_ids)
}

// GetKeyIdSize gets `AVEncryptionInitInfo.key_id_size` value.
func (eii *AVEncryptionInitInfo) GetKeyIdSize() uint32 {
	return (uint32)(eii.key_id_size)
}

// SetKeyIdSize sets `AVEncryptionInitInfo.key_id_size` value.
func (eii *AVEncryptionInitInfo) SetKeyIdSize(v uint32) {
	eii.key_id_size = (C.uint32_t)(v)
}

// GetKeyIdSizeAddr gets `AVEncryptionInitInfo.key_id_size` address.
func (eii *AVEncryptionInitInfo) GetKeyIdSizeAddr() *uint32 {
	return (*uint32)(&eii.key_id_size)
}

// GetData gets `AVEncryptionInitInfo.data` value.
func (eii *AVEncryptionInitInfo) GetData() *uint8 {
	return (*uint8)(eii.data)
}

// SetData sets `AVEncryptionInitInfo.data` value.
func (eii *AVEncryptionInitInfo) SetData(v *uint8) {
	eii.data = (*C.uint8_t)(v)
}

// GetDataAddr gets `AVEncryptionInitInfo.data` address.
func (eii *AVEncryptionInitInfo) GetDataAddr() **uint8 {
	return (**uint8)(unsafe.Pointer(&eii.data))
}

// GetDataSize gets `AVEncryptionInitInfo.data_size` value.
func (eii *AVEncryptionInitInfo) GetDataSize() uint32 {
	return (uint32)(eii.data_size)
}

// SetDataSize sets `AVEncryptionInitInfo.data_size` value.
func (eii *AVEncryptionInitInfo) SetDataSize(v uint32) {
	eii.data_size = (C.uint32_t)(v)
}

// GetDataSizeAddr gets `AVEncryptionInitInfo.data_size` address.
func (eii *AVEncryptionInitInfo) GetDataSizeAddr() *uint32 {
	return (*uint32)(&eii.data_size)
}

// GetNext gets `AVEncryptionInitInfo.next` value.
func (eii *AVEncryptionInitInfo) GetNext() *AVEncryptionInitInfo {
	return (*AVEncryptionInitInfo)(eii.next)
}

// SetNext sets `AVEncryptionInitInfo.next` value.
func (eii *AVEncryptionInitInfo) SetNext(v *AVEncryptionInitInfo) {
	eii.next = (*C.struct_AVEncryptionInitInfo)(v)
}

// GetNextAddr gets `AVEncryptionInitInfo.next` address.
func (eii *AVEncryptionInitInfo) GetNextAddr() **AVEncryptionInitInfo {
	return (**AVEncryptionInitInfo)(unsafe.Pointer(&eii.next))
}

// AvEncryptionInfoAlloc allocates an AVEncryptionInfo structure and sub-pointers to hold the given
// number of subsamples.
func AvEncryptionInfoAlloc(subsampleCount, keyIdSize, ivSize uint32) *AVEncryptionInfo {
	return (*AVEncryptionInfo)(C.av_encryption_info_alloc(
		(C.uint32_t)(subsampleCount), (C.uint32_t)(keyIdSize), (C.uint32_t)(ivSize)))
}

// AvEncryptionInfoClone allocates an AVEncryptionInfo structure with a copy of the given data.
func AvEncryptionInfoClone(info *AVEncryptionInfo) *AVEncryptionInfo {
	return (*AVEncryptionInfo)(C.av_encryption_info_clone((*C.struct_AVEncryptionInfo)(info)))
}

// AvEncryptionInfoFree frees the given encryption info object.
func AvEncryptionInfoFree(info *AVEncryptionInfo) {
	C.av_encryption_info_free((*C.struct_AVEncryptionInfo)(info))
}

// AvEncryptionInfoGetSideData creates a copy of the AVEncryptionInfo that is contained
// in the given side data.
func AvEncryptionInfoGetSideData(sideData *uint8, sideDataSize uintptr) *AVEncryptionInfo {
	return (*AVEncryptionInfo)(C.av_encryption_info_get_side_data(
		(*C.uint8_t)(sideData), (C.size_t)(sideDataSize)))
}

// AvEncryptionInfoAddSideData allocates and initializes side data that holds a copy
// of the given encryption info.
func AvEncryptionInfoAddSideData(info *AVEncryptionInfo, sideDataSize *uintptr) *uint8 {
	return (*uint8)(C.av_encryption_info_add_side_data(
		(*C.struct_AVEncryptionInfo)(info), (*C.size_t)(unsafe.Pointer(sideDataSize))))
}

// AvEncryptionInitInfoAlloc allocates an AVEncryptionInitInfo structure and sub-pointers to hold the
// given sizes.
func AvEncryptionInitInfoAlloc(systemIdSize, numKeyIds, keyIdSize, dataSize uint32) *AVEncryptionInitInfo {
	return (*AVEncryptionInitInfo)(C.av_encryption_init_info_alloc(
		(C.uint32_t)(systemIdSize), (C.uint32_t)(numKeyIds), (C.uint32_t)(keyIdSize), (C.uint32_t)(dataSize)))
}

// AvEncryptionInitInfoFree frees the given encryption init info object.
func AvEncryptionInitInfoFree(info *AVEncryptionInitInfo) {
	C.av_encryption_init_info_free((*C.struct_AVEncryptionInitInfo)(info))
}

// AvEncryptionInitInfoGetSideData creates a copy of the AVEncryptionInitInfo that is contained in the given
// side data.
func AvEncryptionInitInfoGetSideData(sideData *uint8, sideDataSize uintptr) *AVEncryptionInitInfo {
	return (*AVEncryptionInitInfo)(C.av_encryption_init_info_get_side_data(
		(*C.uint8_t)(sideData), (C.size_t)(sideDataSize)))
}

// AvEncryptionInitInfoAddSideData allocates and initializes side data that holds a copy
// of the given encryption init info.
func AvEncryptionInitInfoAddSideData(info *AVEncryptionInitInfo, sideDataSize *uintptr) *uint8 {
	return (*uint8)(C.av_encryption_init_info_add_side_data(
		(*C.struct_AVEncryptionInitInfo)(info), (*C.size_t)(unsafe.Pointer(sideDataSize))))
}
