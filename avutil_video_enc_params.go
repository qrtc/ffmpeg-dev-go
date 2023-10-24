package ffmpeg

/*
#include <libavutil/video_enc_params.h>
*/
import "C"
import "unsafe"

// AVVideoEncParamsType
type AVVideoEncParamsType = C.enum_AVVideoEncParamsType

const (
	AV_VIDEO_ENC_PARAMS_NONE  = AVVideoEncParamsType(C.AV_VIDEO_ENC_PARAMS_NONE)
	AV_VIDEO_ENC_PARAMS_VP9   = AVVideoEncParamsType(C.AV_VIDEO_ENC_PARAMS_VP9)
	AV_VIDEO_ENC_PARAMS_H264  = AVVideoEncParamsType(C.AV_VIDEO_ENC_PARAMS_H264)
	AV_VIDEO_ENC_PARAMS_MPEG2 = AVVideoEncParamsType(C.AV_VIDEO_ENC_PARAMS_MPEG2)
)

// AVVideoEncParams
type AVVideoEncParams C.struct_AVVideoEncParams

// Custom: GetNbBlocks gets `AVVideoEncParams.nb_blocks` value.
func (vep *AVVideoEncParams) GetNbBlocks() uint32 {
	return (uint32)(vep.nb_blocks)
}

// Custom: SetNbBlocks sets `AVVideoEncParams.nb_blocks` value.
func (vep *AVVideoEncParams) SetNbBlocks(v uint32) {
	vep.nb_blocks = (C.uint)(v)
}

// Custom: GetNbBlocksAddr gets `AVVideoEncParams.nb_blocks` address.
func (vep *AVVideoEncParams) GetNbBlocksAddr() *uint32 {
	return (*uint32)(&vep.nb_blocks)
}

// Custom: GetBlocksOffset gets `AVVideoEncParams.blocks_offset` value.
func (vep *AVVideoEncParams) GetBlocksOffset() uintptr {
	return (uintptr)(vep.blocks_offset)
}

// Custom: SetBlocksOffset sets `AVVideoEncParams.blocks_offset` value.
func (vep *AVVideoEncParams) SetBlocksOffset(v uintptr) {
	vep.blocks_offset = (C.size_t)(v)
}

// Custom: GetBlocksOffsetAddr gets `AVVideoEncParams.blocks_offset` address.
func (vep *AVVideoEncParams) GetBlocksOffsetAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&vep.blocks_offset))
}

// Custom: GetBlockSize gets `AVVideoEncParams.block_size` value.
func (vep *AVVideoEncParams) GetBlockSize() uintptr {
	return (uintptr)(vep.block_size)
}

// Custom: SetBlockSize sets `AVVideoEncParams.block_size` value.
func (vep *AVVideoEncParams) SetBlockSize(v uintptr) {
	vep.block_size = (C.size_t)(v)
}

// Custom: GetBlockSizeAddr gets `AVVideoEncParams.block_size` address.
func (vep *AVVideoEncParams) GetBlockSizeAddr() *uintptr {
	return (*uintptr)(unsafe.Pointer(&vep.block_size))
}

// Custom: GetType gets `AVVideoEncParams.type` value.
func (vep *AVVideoEncParams) GetType() AVVideoEncParamsType {
	return (AVVideoEncParamsType)(vep._type)
}

// Custom: SetType sets `AVVideoEncParams.type` value.
func (vep *AVVideoEncParams) SetType(v AVVideoEncParamsType) {
	vep._type = (C.enum_AVVideoEncParamsType)(v)
}

// Custom: GetTypeAddr gets `AVVideoEncParams.type` address.
func (vep *AVVideoEncParams) GetTypeAddr() *AVVideoEncParamsType {
	return (*AVVideoEncParamsType)(&vep._type)
}

// Custom: GetQp gets `AVVideoEncParams.qp` value.
func (vep *AVVideoEncParams) GetQp() int32 {
	return (int32)(vep.qp)
}

// Custom: SetQp sets `AVVideoEncParams.qp` value.
func (vep *AVVideoEncParams) SetQp(v int32) {
	vep.qp = (C.int32_t)(v)
}

// Custom: GetQpAddr gets `AVVideoEncParams.qp` address.
func (vep *AVVideoEncParams) GetQpAddr() *int32 {
	return (*int32)(&vep.qp)
}

// Custom: GetDeltaQp gets `AVVideoEncParams.delta_qp` value.
func (vep *AVVideoEncParams) GetDeltaQp() (v [][]int32) {
	for i := 0; i < 4; i++ {
		v = append(v, unsafe.Slice((*int32)(&vep.delta_qp[i][0]), 2))
	}
	return v
}

// Custom: SetDeltaQp sets `AVVideoEncParams.delta_qp` value.
func (vep *AVVideoEncParams) SetDeltaQp(v [][]int32) {
	for i := 0; i < FFMIN(len(v), 4); i++ {
		for j := 0; j < FFMIN(len(v[i]), 2); j++ {
			vep.delta_qp[i][j] = (C.int)(v[i][j])
		}
	}
}

// Custom: GetDeltaQpAddr gets `AVVideoEncParams.delta_qp` address.
func (vep *AVVideoEncParams) GetDeltaQpAddr() **int32 {
	return (**int32)(unsafe.Pointer(&vep.delta_qp))
}

// AVVideoBlockParams
type AVVideoBlockParams C.struct_AVVideoBlockParams

// Custom: GetSrcX gets `AVVideoBlockParams.src_x` value.
func (vbp *AVVideoBlockParams) GetSrcX() int32 {
	return (int32)(vbp.src_x)
}

// Custom: SetSrcX sets `AVVideoBlockParams.src_x` value.
func (vbp *AVVideoBlockParams) SetSrcX(v int32) {
	vbp.src_x = (C.int)(v)
}

// Custom: GetSrcXAddr gets `AVVideoBlockParams.src_x` address.
func (vbp *AVVideoBlockParams) GetSrcXAddr() *int32 {
	return (*int32)(&vbp.src_x)
}

// Custom: GetSrcY gets `AVVideoBlockParams.src_y` value.
func (vbp *AVVideoBlockParams) GetSrcY() int32 {
	return (int32)(vbp.src_y)
}

// Custom: SetSrcY sets `AVVideoBlockParams.src_y` value.
func (vbp *AVVideoBlockParams) SetSrcY(v int32) {
	vbp.src_y = (C.int)(v)
}

// Custom: GetSrcYAddr gets `AVVideoBlockParams.src_y` address.
func (vbp *AVVideoBlockParams) GetSrcYAddr() *int32 {
	return (*int32)(&vbp.src_y)
}

// Custom: GetW gets `AVVideoBlockParams.w` value.
func (vbp *AVVideoBlockParams) GetW() int32 {
	return (int32)(vbp.w)
}

// Custom: SetW sets `AVVideoBlockParams.w` value.
func (vbp *AVVideoBlockParams) SetW(v int32) {
	vbp.w = (C.int)(v)
}

// Custom: GetWAddr gets `AVVideoBlockParams.w` address.
func (vbp *AVVideoBlockParams) GetWAddr() *int32 {
	return (*int32)(&vbp.w)
}

// Custom: GetH gets `AVVideoBlockParams.h` value.
func (vbp *AVVideoBlockParams) GetH() int32 {
	return (int32)(vbp.h)
}

// Custom: SetH sets `AVVideoBlockParams.h` value.
func (vbp *AVVideoBlockParams) SetH(v int32) {
	vbp.h = (C.int)(v)
}

// Custom: GetHAddr gets `AVVideoBlockParams.h` address.
func (vbp *AVVideoBlockParams) GetHAddr() *int32 {
	return (*int32)(&vbp.h)
}

// Custom: GetDeltaQp gets `AVVideoBlockParams.delta_qp` value.
func (vbp *AVVideoBlockParams) GetDeltaQp() int32 {
	return (int32)(vbp.delta_qp)
}

// Custom: SetDeltaQp sets `AVVideoBlockParams.delta_qp` value.
func (vbp *AVVideoBlockParams) SetDeltaQp(v int32) {
	vbp.delta_qp = (C.int32_t)(v)
}

// Custom: GetDeltaQpAddr gets `AVVideoBlockParams.delta_qp` address.
func (vbp *AVVideoBlockParams) GetDeltaQpAddr() *int32 {
	return (*int32)(&vbp.delta_qp)
}

// AvVideoEncParamsBlock gets the block at the specified idx.
func AvVideoEncParamsBlock(par *AVVideoEncParams, idx uint32) *AVVideoBlockParams {
	return (*AVVideoBlockParams)(C.av_video_enc_params_block((*C.struct_AVVideoEncParams)(par), (C.uint)(idx)))
}

// AvVideoEncParamsAlloc allocates memory for AVVideoEncParams of the given type, plus an array of
// nbBlocks AVVideoBlockParams and initializes the variables.
func AvVideoEncParamsAlloc(_type AVVideoEncParamsType, nbBlocks uint32, outSize *uintptr) *AVVideoEncParams {
	return (*AVVideoEncParams)(C.av_video_enc_params_alloc((C.enum_AVVideoEncParamsType)(_type),
		(C.uint)(nbBlocks), (*C.size_t)(unsafe.Pointer(outSize))))
}

// AvVideoEncParamsCreateSideData sllocates memory for AVEncodeInfoFrame plus an array of
// nbBlocks AVEncodeInfoBlock in the given AVFrame frame
// as AVFrameSideData of type AV_FRAME_DATA_VIDEO_ENC_PARAMS
// and initializes the variables.
func AvVideoEncParamsCreateSideData(frame *AVFrame, _type AVVideoEncParamsType, nbBlocks uint32) *AVVideoEncParams {
	return (*AVVideoEncParams)(C.av_video_enc_params_create_side_data((*C.struct_AVFrame)(frame),
		(C.enum_AVVideoEncParamsType)(_type), (C.uint)(nbBlocks)))
}
