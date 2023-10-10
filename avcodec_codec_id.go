package ffmpeg

/*
#include <libavcodec/codec_id.h>
*/
import "C"

// AvCodecID
type AvCodecID int32

const (
	AV_CODEC_ID_NONE = AvCodecID(C.AV_CODEC_ID_NONE)

	// video codecs
	AV_CODEC_ID_MPEG1VIDEO      = AvCodecID(C.AV_CODEC_ID_MPEG1VIDEO)
	AV_CODEC_ID_MPEG2VIDEO      = AvCodecID(C.AV_CODEC_ID_MPEG2VIDEO)
	AV_CODEC_ID_H261            = AvCodecID(C.AV_CODEC_ID_H261)
	AV_CODEC_ID_H263            = AvCodecID(C.AV_CODEC_ID_H263)
	AV_CODEC_ID_RV10            = AvCodecID(C.AV_CODEC_ID_RV10)
	AV_CODEC_ID_RV20            = AvCodecID(C.AV_CODEC_ID_RV20)
	AV_CODEC_ID_MJPEG           = AvCodecID(C.AV_CODEC_ID_MJPEG)
	AV_CODEC_ID_MJPEGB          = AvCodecID(C.AV_CODEC_ID_MJPEGB)
	AV_CODEC_ID_LJPEG           = AvCodecID(C.AV_CODEC_ID_LJPEG)
	AV_CODEC_ID_SP5X            = AvCodecID(C.AV_CODEC_ID_SP5X)
	AV_CODEC_ID_JPEGLS          = AvCodecID(C.AV_CODEC_ID_JPEGLS)
	AV_CODEC_ID_MPEG4           = AvCodecID(C.AV_CODEC_ID_MPEG4)
	AV_CODEC_ID_RAWVIDEO        = AvCodecID(C.AV_CODEC_ID_RAWVIDEO)
	AV_CODEC_ID_MSMPEG4V1       = AvCodecID(C.AV_CODEC_ID_MSMPEG4V1)
	AV_CODEC_ID_MSMPEG4V2       = AvCodecID(C.AV_CODEC_ID_MSMPEG4V2)
	AV_CODEC_ID_MSMPEG4V3       = AvCodecID(C.AV_CODEC_ID_MSMPEG4V3)
	AV_CODEC_ID_WMV1            = AvCodecID(C.AV_CODEC_ID_WMV1)
	AV_CODEC_ID_WMV2            = AvCodecID(C.AV_CODEC_ID_WMV2)
	AV_CODEC_ID_H263P           = AvCodecID(C.AV_CODEC_ID_H263P)
	AV_CODEC_ID_H263I           = AvCodecID(C.AV_CODEC_ID_H263I)
	AV_CODEC_ID_FLV1            = AvCodecID(C.AV_CODEC_ID_FLV1)
	AV_CODEC_ID_SVQ1            = AvCodecID(C.AV_CODEC_ID_SVQ1)
	AV_CODEC_ID_SVQ3            = AvCodecID(C.AV_CODEC_ID_SVQ3)
	AV_CODEC_ID_DVVIDEO         = AvCodecID(C.AV_CODEC_ID_DVVIDEO)
	AV_CODEC_ID_HUFFYUV         = AvCodecID(C.AV_CODEC_ID_HUFFYUV)
	AV_CODEC_ID_CYUV            = AvCodecID(C.AV_CODEC_ID_CYUV)
	AV_CODEC_ID_H264            = AvCodecID(C.AV_CODEC_ID_H264)
	AV_CODEC_ID_INDEO3          = AvCodecID(C.AV_CODEC_ID_INDEO3)
	AV_CODEC_ID_VP3             = AvCodecID(C.AV_CODEC_ID_VP3)
	AV_CODEC_ID_THEORA          = AvCodecID(C.AV_CODEC_ID_THEORA)
	AV_CODEC_ID_ASV1            = AvCodecID(C.AV_CODEC_ID_ASV1)
	AV_CODEC_ID_ASV2            = AvCodecID(C.AV_CODEC_ID_ASV2)
	AV_CODEC_ID_FFV1            = AvCodecID(C.AV_CODEC_ID_FFV1)
	AV_CODEC_ID_4XM             = AvCodecID(C.AV_CODEC_ID_4XM)
	AV_CODEC_ID_VCR1            = AvCodecID(C.AV_CODEC_ID_VCR1)
	AV_CODEC_ID_CLJR            = AvCodecID(C.AV_CODEC_ID_CLJR)
	AV_CODEC_ID_MDEC            = AvCodecID(C.AV_CODEC_ID_MDEC)
	AV_CODEC_ID_ROQ             = AvCodecID(C.AV_CODEC_ID_ROQ)
	AV_CODEC_ID_INTERPLAY_VIDEO = AvCodecID(C.AV_CODEC_ID_INTERPLAY_VIDEO)
	AV_CODEC_ID_XAN_WC3         = AvCodecID(C.AV_CODEC_ID_XAN_WC3)
	AV_CODEC_ID_XAN_WC4         = AvCodecID(C.AV_CODEC_ID_XAN_WC4)
	AV_CODEC_ID_RPZA            = AvCodecID(C.AV_CODEC_ID_RPZA)
	AV_CODEC_ID_CINEPAK         = AvCodecID(C.AV_CODEC_ID_CINEPAK)
	AV_CODEC_ID_WS_VQA          = AvCodecID(C.AV_CODEC_ID_WS_VQA)
	AV_CODEC_ID_MSRLE           = AvCodecID(C.AV_CODEC_ID_MSRLE)
	AV_CODEC_ID_MSVIDEO1        = AvCodecID(C.AV_CODEC_ID_MSVIDEO1)
	AV_CODEC_ID_IDCIN           = AvCodecID(C.AV_CODEC_ID_IDCIN)
	AV_CODEC_ID_8BPS            = AvCodecID(C.AV_CODEC_ID_8BPS)
	AV_CODEC_ID_SMC             = AvCodecID(C.AV_CODEC_ID_SMC)
	AV_CODEC_ID_FLIC            = AvCodecID(C.AV_CODEC_ID_FLIC)
	AV_CODEC_ID_TRUEMOTION1     = AvCodecID(C.AV_CODEC_ID_TRUEMOTION1)
	AV_CODEC_ID_VMDVIDEO        = AvCodecID(C.AV_CODEC_ID_VMDVIDEO)
	AV_CODEC_ID_MSZH            = AvCodecID(C.AV_CODEC_ID_MSZH)
	AV_CODEC_ID_ZLIB            = AvCodecID(C.AV_CODEC_ID_ZLIB)
	AV_CODEC_ID_QTRLE           = AvCodecID(C.AV_CODEC_ID_QTRLE)
	AV_CODEC_ID_TSCC            = AvCodecID(C.AV_CODEC_ID_TSCC)
	AV_CODEC_ID_ULTI            = AvCodecID(C.AV_CODEC_ID_ULTI)
	AV_CODEC_ID_QDRAW           = AvCodecID(C.AV_CODEC_ID_QDRAW)
	AV_CODEC_ID_VIXL            = AvCodecID(C.AV_CODEC_ID_VIXL)
	AV_CODEC_ID_QPEG            = AvCodecID(C.AV_CODEC_ID_QPEG)
	AV_CODEC_ID_PNG             = AvCodecID(C.AV_CODEC_ID_PNG)
	AV_CODEC_ID_PPM             = AvCodecID(C.AV_CODEC_ID_PPM)
	AV_CODEC_ID_PBM             = AvCodecID(C.AV_CODEC_ID_PBM)
	AV_CODEC_ID_PGM             = AvCodecID(C.AV_CODEC_ID_PGM)
	AV_CODEC_ID_PGMYUV          = AvCodecID(C.AV_CODEC_ID_PGMYUV)
	AV_CODEC_ID_PAM             = AvCodecID(C.AV_CODEC_ID_PAM)
	AV_CODEC_ID_FFVHUFF         = AvCodecID(C.AV_CODEC_ID_FFVHUFF)
	AV_CODEC_ID_RV30            = AvCodecID(C.AV_CODEC_ID_RV30)
	AV_CODEC_ID_RV40            = AvCodecID(C.AV_CODEC_ID_RV40)
	AV_CODEC_ID_VC1             = AvCodecID(C.AV_CODEC_ID_VC1)
	AV_CODEC_ID_WMV3            = AvCodecID(C.AV_CODEC_ID_WMV3)
	AV_CODEC_ID_LOCO            = AvCodecID(C.AV_CODEC_ID_LOCO)
	AV_CODEC_ID_WNV1            = AvCodecID(C.AV_CODEC_ID_WNV1)
	AV_CODEC_ID_AASC            = AvCodecID(C.AV_CODEC_ID_AASC)
	AV_CODEC_ID_INDEO2          = AvCodecID(C.AV_CODEC_ID_INDEO2)
	AV_CODEC_ID_FRAPS           = AvCodecID(C.AV_CODEC_ID_FRAPS)
	AV_CODEC_ID_TRUEMOTION2     = AvCodecID(C.AV_CODEC_ID_TRUEMOTION2)
	AV_CODEC_ID_BMP             = AvCodecID(C.AV_CODEC_ID_BMP)
	AV_CODEC_ID_CSCD            = AvCodecID(C.AV_CODEC_ID_CSCD)
	AV_CODEC_ID_MMVIDEO         = AvCodecID(C.AV_CODEC_ID_MMVIDEO)
	AV_CODEC_ID_ZMBV            = AvCodecID(C.AV_CODEC_ID_ZMBV)
	AV_CODEC_ID_AVS             = AvCodecID(C.AV_CODEC_ID_AVS)
	AV_CODEC_ID_SMACKVIDEO      = AvCodecID(C.AV_CODEC_ID_SMACKVIDEO)
	AV_CODEC_ID_NUV             = AvCodecID(C.AV_CODEC_ID_NUV)
	AV_CODEC_ID_KMVC            = AvCodecID(C.AV_CODEC_ID_KMVC)
	AV_CODEC_ID_FLASHSV         = AvCodecID(C.AV_CODEC_ID_FLASHSV)
	AV_CODEC_ID_CAVS            = AvCodecID(C.AV_CODEC_ID_CAVS)
	AV_CODEC_ID_JPEG2000        = AvCodecID(C.AV_CODEC_ID_JPEG2000)
	AV_CODEC_ID_VMNC            = AvCodecID(C.AV_CODEC_ID_VMNC)
	AV_CODEC_ID_VP5             = AvCodecID(C.AV_CODEC_ID_VP5)
	AV_CODEC_ID_VP6             = AvCodecID(C.AV_CODEC_ID_VP6)
	AV_CODEC_ID_VP6F            = AvCodecID(C.AV_CODEC_ID_VP6F)
	AV_CODEC_ID_TARGA           = AvCodecID(C.AV_CODEC_ID_TARGA)
	AV_CODEC_ID_DSICINVIDEO     = AvCodecID(C.AV_CODEC_ID_DSICINVIDEO)
	AV_CODEC_ID_TIERTEXSEQVIDEO = AvCodecID(C.AV_CODEC_ID_TIERTEXSEQVIDEO)
	AV_CODEC_ID_TIFF            = AvCodecID(C.AV_CODEC_ID_TIFF)
	AV_CODEC_ID_GIF             = AvCodecID(C.AV_CODEC_ID_GIF)
	AV_CODEC_ID_DXA             = AvCodecID(C.AV_CODEC_ID_DXA)
	AV_CODEC_ID_DNXHD           = AvCodecID(C.AV_CODEC_ID_DNXHD)
	AV_CODEC_ID_THP             = AvCodecID(C.AV_CODEC_ID_THP)
	AV_CODEC_ID_SGI             = AvCodecID(C.AV_CODEC_ID_SGI)
	AV_CODEC_ID_C93             = AvCodecID(C.AV_CODEC_ID_C93)
	AV_CODEC_ID_BETHSOFTVID     = AvCodecID(C.AV_CODEC_ID_BETHSOFTVID)
	AV_CODEC_ID_PTX             = AvCodecID(C.AV_CODEC_ID_PTX)
	AV_CODEC_ID_TXD             = AvCodecID(C.AV_CODEC_ID_TXD)
	AV_CODEC_ID_VP6A            = AvCodecID(C.AV_CODEC_ID_VP6A)
	AV_CODEC_ID_AMV             = AvCodecID(C.AV_CODEC_ID_AMV)
	AV_CODEC_ID_VB              = AvCodecID(C.AV_CODEC_ID_VB)
	AV_CODEC_ID_PCX             = AvCodecID(C.AV_CODEC_ID_PCX)
	AV_CODEC_ID_SUNRAST         = AvCodecID(C.AV_CODEC_ID_SUNRAST)
	AV_CODEC_ID_INDEO4          = AvCodecID(C.AV_CODEC_ID_INDEO4)
	AV_CODEC_ID_INDEO5          = AvCodecID(C.AV_CODEC_ID_INDEO5)
	AV_CODEC_ID_MIMIC           = AvCodecID(C.AV_CODEC_ID_MIMIC)
	AV_CODEC_ID_RL2             = AvCodecID(C.AV_CODEC_ID_RL2)
	AV_CODEC_ID_ESCAPE124       = AvCodecID(C.AV_CODEC_ID_ESCAPE124)
	AV_CODEC_ID_DIRAC           = AvCodecID(C.AV_CODEC_ID_DIRAC)
	AV_CODEC_ID_BFI             = AvCodecID(C.AV_CODEC_ID_BFI)
	AV_CODEC_ID_CMV             = AvCodecID(C.AV_CODEC_ID_CMV)
	AV_CODEC_ID_MOTIONPIXELS    = AvCodecID(C.AV_CODEC_ID_MOTIONPIXELS)
	AV_CODEC_ID_TGV             = AvCodecID(C.AV_CODEC_ID_TGV)
	AV_CODEC_ID_TGQ             = AvCodecID(C.AV_CODEC_ID_TGQ)
	AV_CODEC_ID_TQI             = AvCodecID(C.AV_CODEC_ID_TQI)
	AV_CODEC_ID_AURA            = AvCodecID(C.AV_CODEC_ID_AURA)
	AV_CODEC_ID_AURA2           = AvCodecID(C.AV_CODEC_ID_AURA2)
	AV_CODEC_ID_V210X           = AvCodecID(C.AV_CODEC_ID_V210X)
	AV_CODEC_ID_TMV             = AvCodecID(C.AV_CODEC_ID_TMV)
	AV_CODEC_ID_V210            = AvCodecID(C.AV_CODEC_ID_V210)
	AV_CODEC_ID_DPX             = AvCodecID(C.AV_CODEC_ID_DPX)
	AV_CODEC_ID_MAD             = AvCodecID(C.AV_CODEC_ID_MAD)
	AV_CODEC_ID_FRWU            = AvCodecID(C.AV_CODEC_ID_FRWU)
	AV_CODEC_ID_FLASHSV2        = AvCodecID(C.AV_CODEC_ID_FLASHSV2)
	AV_CODEC_ID_CDGRAPHICS      = AvCodecID(C.AV_CODEC_ID_CDGRAPHICS)
	AV_CODEC_ID_R210            = AvCodecID(C.AV_CODEC_ID_R210)
	AV_CODEC_ID_ANM             = AvCodecID(C.AV_CODEC_ID_ANM)
	AV_CODEC_ID_BINKVIDEO       = AvCodecID(C.AV_CODEC_ID_BINKVIDEO)
	AV_CODEC_ID_IFF_ILBM        = AvCodecID(C.AV_CODEC_ID_IFF_ILBM)
	AV_CODEC_ID_IFF_BYTERUN1    = AvCodecID(C.AV_CODEC_ID_IFF_BYTERUN1)
	AV_CODEC_ID_KGV1            = AvCodecID(C.AV_CODEC_ID_KGV1)
	AV_CODEC_ID_YOP             = AvCodecID(C.AV_CODEC_ID_YOP)
	AV_CODEC_ID_VP8             = AvCodecID(C.AV_CODEC_ID_VP8)
	AV_CODEC_ID_PICTOR          = AvCodecID(C.AV_CODEC_ID_PICTOR)
	AV_CODEC_ID_ANSI            = AvCodecID(C.AV_CODEC_ID_ANSI)
	AV_CODEC_ID_A64_MULTI       = AvCodecID(C.AV_CODEC_ID_A64_MULTI)
	AV_CODEC_ID_A64_MULTI5      = AvCodecID(C.AV_CODEC_ID_A64_MULTI5)
	AV_CODEC_ID_R10K            = AvCodecID(C.AV_CODEC_ID_R10K)
	AV_CODEC_ID_MXPEG           = AvCodecID(C.AV_CODEC_ID_MXPEG)
	AV_CODEC_ID_LAGARITH        = AvCodecID(C.AV_CODEC_ID_LAGARITH)
	AV_CODEC_ID_PRORES          = AvCodecID(C.AV_CODEC_ID_PRORES)
	AV_CODEC_ID_JV              = AvCodecID(C.AV_CODEC_ID_JV)
	AV_CODEC_ID_DFA             = AvCodecID(C.AV_CODEC_ID_DFA)
	AV_CODEC_ID_WMV3IMAGE       = AvCodecID(C.AV_CODEC_ID_WMV3IMAGE)
	AV_CODEC_ID_VC1IMAGE        = AvCodecID(C.AV_CODEC_ID_VC1IMAGE)
	AV_CODEC_ID_UTVIDEO         = AvCodecID(C.AV_CODEC_ID_UTVIDEO)
	AV_CODEC_ID_BMV_VIDEO       = AvCodecID(C.AV_CODEC_ID_BMV_VIDEO)
	AV_CODEC_ID_VBLE            = AvCodecID(C.AV_CODEC_ID_VBLE)
	AV_CODEC_ID_DXTORY          = AvCodecID(C.AV_CODEC_ID_DXTORY)
	AV_CODEC_ID_V410            = AvCodecID(C.AV_CODEC_ID_V410)
	AV_CODEC_ID_XWD             = AvCodecID(C.AV_CODEC_ID_XWD)
	AV_CODEC_ID_CDXL            = AvCodecID(C.AV_CODEC_ID_CDXL)
	AV_CODEC_ID_XBM             = AvCodecID(C.AV_CODEC_ID_XBM)
	AV_CODEC_ID_ZEROCODEC       = AvCodecID(C.AV_CODEC_ID_ZEROCODEC)
	AV_CODEC_ID_MSS1            = AvCodecID(C.AV_CODEC_ID_MSS1)
	AV_CODEC_ID_MSA1            = AvCodecID(C.AV_CODEC_ID_MSA1)
	AV_CODEC_ID_TSCC2           = AvCodecID(C.AV_CODEC_ID_TSCC2)
	AV_CODEC_ID_MTS2            = AvCodecID(C.AV_CODEC_ID_MTS2)
	AV_CODEC_ID_CLLC            = AvCodecID(C.AV_CODEC_ID_CLLC)
	AV_CODEC_ID_MSS2            = AvCodecID(C.AV_CODEC_ID_MSS2)
	AV_CODEC_ID_VP9             = AvCodecID(C.AV_CODEC_ID_VP9)
	AV_CODEC_ID_AIC             = AvCodecID(C.AV_CODEC_ID_AIC)
	AV_CODEC_ID_ESCAPE130       = AvCodecID(C.AV_CODEC_ID_ESCAPE130)
	AV_CODEC_ID_G2M             = AvCodecID(C.AV_CODEC_ID_G2M)
	AV_CODEC_ID_WEBP            = AvCodecID(C.AV_CODEC_ID_WEBP)
	AV_CODEC_ID_HNM4_VIDEO      = AvCodecID(C.AV_CODEC_ID_HNM4_VIDEO)
	AV_CODEC_ID_HEVC            = AvCodecID(C.AV_CODEC_ID_HEVC)
	AV_CODEC_ID_H265            = AvCodecID(C.AV_CODEC_ID_H265)
	AV_CODEC_ID_FIC             = AvCodecID(C.AV_CODEC_ID_FIC)
	AV_CODEC_ID_ALIAS_PIX       = AvCodecID(C.AV_CODEC_ID_ALIAS_PIX)
	AV_CODEC_ID_BRENDER_PIX     = AvCodecID(C.AV_CODEC_ID_BRENDER_PIX)
	AV_CODEC_ID_PAF_VIDEO       = AvCodecID(C.AV_CODEC_ID_PAF_VIDEO)
	AV_CODEC_ID_EXR             = AvCodecID(C.AV_CODEC_ID_EXR)
	AV_CODEC_ID_VP7             = AvCodecID(C.AV_CODEC_ID_VP7)
	AV_CODEC_ID_SANM            = AvCodecID(C.AV_CODEC_ID_SANM)
	AV_CODEC_ID_SGIRLE          = AvCodecID(C.AV_CODEC_ID_SGIRLE)
	AV_CODEC_ID_MVC1            = AvCodecID(C.AV_CODEC_ID_MVC1)
	AV_CODEC_ID_MVC2            = AvCodecID(C.AV_CODEC_ID_MVC2)
	AV_CODEC_ID_HQX             = AvCodecID(C.AV_CODEC_ID_HQX)
	AV_CODEC_ID_TDSC            = AvCodecID(C.AV_CODEC_ID_TDSC)
	AV_CODEC_ID_HQ_HQA          = AvCodecID(C.AV_CODEC_ID_HQ_HQA)
	AV_CODEC_ID_HAP             = AvCodecID(C.AV_CODEC_ID_HAP)
	AV_CODEC_ID_DDS             = AvCodecID(C.AV_CODEC_ID_DDS)
	AV_CODEC_ID_DXV             = AvCodecID(C.AV_CODEC_ID_DXV)
	AV_CODEC_ID_SCREENPRESSO    = AvCodecID(C.AV_CODEC_ID_SCREENPRESSO)
	AV_CODEC_ID_RSCC            = AvCodecID(C.AV_CODEC_ID_RSCC)
	AV_CODEC_ID_AVS2            = AvCodecID(C.AV_CODEC_ID_AVS2)
	AV_CODEC_ID_PGX             = AvCodecID(C.AV_CODEC_ID_PGX)
	AV_CODEC_ID_AVS3            = AvCodecID(C.AV_CODEC_ID_AVS3)
	AV_CODEC_ID_MSP2            = AvCodecID(C.AV_CODEC_ID_MSP2)
	AV_CODEC_ID_VVC             = AvCodecID(C.AV_CODEC_ID_VVC)
	AV_CODEC_ID_H266            = AvCodecID(C.AV_CODEC_ID_H266)
	AV_CODEC_ID_Y41P            = AvCodecID(C.AV_CODEC_ID_Y41P)
	AV_CODEC_ID_AVRP            = AvCodecID(C.AV_CODEC_ID_AVRP)
	AV_CODEC_ID_012V            = AvCodecID(C.AV_CODEC_ID_012V)
	AV_CODEC_ID_AVUI            = AvCodecID(C.AV_CODEC_ID_AVUI)
	AV_CODEC_ID_AYUV            = AvCodecID(C.AV_CODEC_ID_AYUV)
	AV_CODEC_ID_TARGA_Y216      = AvCodecID(C.AV_CODEC_ID_TARGA_Y216)
	AV_CODEC_ID_V308            = AvCodecID(C.AV_CODEC_ID_V308)
	AV_CODEC_ID_V408            = AvCodecID(C.AV_CODEC_ID_V408)
	AV_CODEC_ID_YUV4            = AvCodecID(C.AV_CODEC_ID_YUV4)
	AV_CODEC_ID_AVRN            = AvCodecID(C.AV_CODEC_ID_AVRN)
	AV_CODEC_ID_CPIA            = AvCodecID(C.AV_CODEC_ID_CPIA)
	AV_CODEC_ID_XFACE           = AvCodecID(C.AV_CODEC_ID_XFACE)
	AV_CODEC_ID_SNOW            = AvCodecID(C.AV_CODEC_ID_SNOW)
	AV_CODEC_ID_SMVJPEG         = AvCodecID(C.AV_CODEC_ID_SMVJPEG)
	AV_CODEC_ID_APNG            = AvCodecID(C.AV_CODEC_ID_APNG)
	AV_CODEC_ID_DAALA           = AvCodecID(C.AV_CODEC_ID_DAALA)
	AV_CODEC_ID_CFHD            = AvCodecID(C.AV_CODEC_ID_CFHD)
	AV_CODEC_ID_TRUEMOTION2RT   = AvCodecID(C.AV_CODEC_ID_TRUEMOTION2RT)
	AV_CODEC_ID_M101            = AvCodecID(C.AV_CODEC_ID_M101)
	AV_CODEC_ID_MAGICYUV        = AvCodecID(C.AV_CODEC_ID_MAGICYUV)
	AV_CODEC_ID_SHEERVIDEO      = AvCodecID(C.AV_CODEC_ID_SHEERVIDEO)
	AV_CODEC_ID_YLC             = AvCodecID(C.AV_CODEC_ID_YLC)
	AV_CODEC_ID_PSD             = AvCodecID(C.AV_CODEC_ID_PSD)
	AV_CODEC_ID_PIXLET          = AvCodecID(C.AV_CODEC_ID_PIXLET)
	AV_CODEC_ID_SPEEDHQ         = AvCodecID(C.AV_CODEC_ID_SPEEDHQ)
	AV_CODEC_ID_FMVC            = AvCodecID(C.AV_CODEC_ID_FMVC)
	AV_CODEC_ID_SCPR            = AvCodecID(C.AV_CODEC_ID_SCPR)
	AV_CODEC_ID_CLEARVIDEO      = AvCodecID(C.AV_CODEC_ID_CLEARVIDEO)
	AV_CODEC_ID_XPM             = AvCodecID(C.AV_CODEC_ID_XPM)
	AV_CODEC_ID_AV1             = AvCodecID(C.AV_CODEC_ID_AV1)
	AV_CODEC_ID_BITPACKED       = AvCodecID(C.AV_CODEC_ID_BITPACKED)
	AV_CODEC_ID_MSCC            = AvCodecID(C.AV_CODEC_ID_MSCC)
	AV_CODEC_ID_SRGC            = AvCodecID(C.AV_CODEC_ID_SRGC)
	AV_CODEC_ID_SVG             = AvCodecID(C.AV_CODEC_ID_SVG)
	AV_CODEC_ID_GDV             = AvCodecID(C.AV_CODEC_ID_GDV)
	AV_CODEC_ID_FITS            = AvCodecID(C.AV_CODEC_ID_FITS)
	AV_CODEC_ID_IMM4            = AvCodecID(C.AV_CODEC_ID_IMM4)
	AV_CODEC_ID_PROSUMER        = AvCodecID(C.AV_CODEC_ID_PROSUMER)
	AV_CODEC_ID_MWSC            = AvCodecID(C.AV_CODEC_ID_MWSC)
	AV_CODEC_ID_WCMV            = AvCodecID(C.AV_CODEC_ID_WCMV)
	AV_CODEC_ID_RASC            = AvCodecID(C.AV_CODEC_ID_RASC)
	AV_CODEC_ID_HYMT            = AvCodecID(C.AV_CODEC_ID_HYMT)
	AV_CODEC_ID_ARBC            = AvCodecID(C.AV_CODEC_ID_ARBC)
	AV_CODEC_ID_AGM             = AvCodecID(C.AV_CODEC_ID_AGM)
	AV_CODEC_ID_LSCR            = AvCodecID(C.AV_CODEC_ID_LSCR)
	AV_CODEC_ID_VP4             = AvCodecID(C.AV_CODEC_ID_VP4)
	AV_CODEC_ID_IMM5            = AvCodecID(C.AV_CODEC_ID_IMM5)
	AV_CODEC_ID_MVDV            = AvCodecID(C.AV_CODEC_ID_MVDV)
	AV_CODEC_ID_MVHA            = AvCodecID(C.AV_CODEC_ID_MVHA)
	AV_CODEC_ID_CDTOONS         = AvCodecID(C.AV_CODEC_ID_CDTOONS)
	AV_CODEC_ID_MV30            = AvCodecID(C.AV_CODEC_ID_MV30)
	AV_CODEC_ID_NOTCHLC         = AvCodecID(C.AV_CODEC_ID_NOTCHLC)
	AV_CODEC_ID_PFM             = AvCodecID(C.AV_CODEC_ID_PFM)
	AV_CODEC_ID_MOBICLIP        = AvCodecID(C.AV_CODEC_ID_MOBICLIP)
	AV_CODEC_ID_PHOTOCD         = AvCodecID(C.AV_CODEC_ID_PHOTOCD)
	AV_CODEC_ID_IPU             = AvCodecID(C.AV_CODEC_ID_IPU)
	AV_CODEC_ID_ARGO            = AvCodecID(C.AV_CODEC_ID_ARGO)
	AV_CODEC_ID_CRI             = AvCodecID(C.AV_CODEC_ID_CRI)
	AV_CODEC_ID_SIMBIOSIS_IMX   = AvCodecID(C.AV_CODEC_ID_SIMBIOSIS_IMX)
	AV_CODEC_ID_SGA_VIDEO       = AvCodecID(C.AV_CODEC_ID_SGA_VIDEO)

	// various PCM "codecs"
	AV_CODEC_ID_FIRST_AUDIO      = AvCodecID(C.AV_CODEC_ID_FIRST_AUDIO)
	AV_CODEC_ID_PCM_S16LE        = AvCodecID(C.AV_CODEC_ID_PCM_S16LE)
	AV_CODEC_ID_PCM_S16BE        = AvCodecID(C.AV_CODEC_ID_PCM_S16BE)
	AV_CODEC_ID_PCM_U16LE        = AvCodecID(C.AV_CODEC_ID_PCM_U16LE)
	AV_CODEC_ID_PCM_U16BE        = AvCodecID(C.AV_CODEC_ID_PCM_U16BE)
	AV_CODEC_ID_PCM_S8           = AvCodecID(C.AV_CODEC_ID_PCM_S8)
	AV_CODEC_ID_PCM_U8           = AvCodecID(C.AV_CODEC_ID_PCM_U8)
	AV_CODEC_ID_PCM_MULAW        = AvCodecID(C.AV_CODEC_ID_PCM_MULAW)
	AV_CODEC_ID_PCM_ALAW         = AvCodecID(C.AV_CODEC_ID_PCM_ALAW)
	AV_CODEC_ID_PCM_S32LE        = AvCodecID(C.AV_CODEC_ID_PCM_S32LE)
	AV_CODEC_ID_PCM_S32BE        = AvCodecID(C.AV_CODEC_ID_PCM_S32BE)
	AV_CODEC_ID_PCM_U32LE        = AvCodecID(C.AV_CODEC_ID_PCM_U32LE)
	AV_CODEC_ID_PCM_U32BE        = AvCodecID(C.AV_CODEC_ID_PCM_U32BE)
	AV_CODEC_ID_PCM_S24LE        = AvCodecID(C.AV_CODEC_ID_PCM_S24LE)
	AV_CODEC_ID_PCM_S24BE        = AvCodecID(C.AV_CODEC_ID_PCM_S24BE)
	AV_CODEC_ID_PCM_U24LE        = AvCodecID(C.AV_CODEC_ID_PCM_U24LE)
	AV_CODEC_ID_PCM_U24BE        = AvCodecID(C.AV_CODEC_ID_PCM_U24BE)
	AV_CODEC_ID_PCM_S24DAUD      = AvCodecID(C.AV_CODEC_ID_PCM_S24DAUD)
	AV_CODEC_ID_PCM_ZORK         = AvCodecID(C.AV_CODEC_ID_PCM_ZORK)
	AV_CODEC_ID_PCM_S16LE_PLANAR = AvCodecID(C.AV_CODEC_ID_PCM_S16LE_PLANAR)
	AV_CODEC_ID_PCM_DVD          = AvCodecID(C.AV_CODEC_ID_PCM_DVD)
	AV_CODEC_ID_PCM_F32BE        = AvCodecID(C.AV_CODEC_ID_PCM_F32BE)
	AV_CODEC_ID_PCM_F32LE        = AvCodecID(C.AV_CODEC_ID_PCM_F32LE)
	AV_CODEC_ID_PCM_F64BE        = AvCodecID(C.AV_CODEC_ID_PCM_F64BE)
	AV_CODEC_ID_PCM_F64LE        = AvCodecID(C.AV_CODEC_ID_PCM_F64LE)
	AV_CODEC_ID_PCM_BLURAY       = AvCodecID(C.AV_CODEC_ID_PCM_BLURAY)
	AV_CODEC_ID_PCM_LXF          = AvCodecID(C.AV_CODEC_ID_PCM_LXF)
	AV_CODEC_ID_S302M            = AvCodecID(C.AV_CODEC_ID_S302M)
	AV_CODEC_ID_PCM_S8_PLANAR    = AvCodecID(C.AV_CODEC_ID_PCM_S8_PLANAR)
	AV_CODEC_ID_PCM_S24LE_PLANAR = AvCodecID(C.AV_CODEC_ID_PCM_S24LE_PLANAR)
	AV_CODEC_ID_PCM_S32LE_PLANAR = AvCodecID(C.AV_CODEC_ID_PCM_S32LE_PLANAR)
	AV_CODEC_ID_PCM_S16BE_PLANAR = AvCodecID(C.AV_CODEC_ID_PCM_S16BE_PLANAR)
	AV_CODEC_ID_PCM_S64LE        = AvCodecID(C.AV_CODEC_ID_PCM_S64LE)
	AV_CODEC_ID_PCM_S64BE        = AvCodecID(C.AV_CODEC_ID_PCM_S64BE)
	AV_CODEC_ID_PCM_F16LE        = AvCodecID(C.AV_CODEC_ID_PCM_F16LE)
	AV_CODEC_ID_PCM_F24LE        = AvCodecID(C.AV_CODEC_ID_PCM_F24LE)
	AV_CODEC_ID_PCM_VIDC         = AvCodecID(C.AV_CODEC_ID_PCM_VIDC)
	AV_CODEC_ID_PCM_SGA          = AvCodecID(C.AV_CODEC_ID_PCM_SGA)

	// various ADPCM codecs
	AV_CODEC_ID_ADPCM_IMA_QT      = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_QT)
	AV_CODEC_ID_ADPCM_IMA_WAV     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_WAV)
	AV_CODEC_ID_ADPCM_IMA_DK3     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_DK3)
	AV_CODEC_ID_ADPCM_IMA_DK4     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_DK4)
	AV_CODEC_ID_ADPCM_IMA_WS      = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_WS)
	AV_CODEC_ID_ADPCM_IMA_SMJPEG  = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_SMJPEG)
	AV_CODEC_ID_ADPCM_MS          = AvCodecID(C.AV_CODEC_ID_ADPCM_MS)
	AV_CODEC_ID_ADPCM_4XM         = AvCodecID(C.AV_CODEC_ID_ADPCM_4XM)
	AV_CODEC_ID_ADPCM_XA          = AvCodecID(C.AV_CODEC_ID_ADPCM_XA)
	AV_CODEC_ID_ADPCM_ADX         = AvCodecID(C.AV_CODEC_ID_ADPCM_ADX)
	AV_CODEC_ID_ADPCM_EA          = AvCodecID(C.AV_CODEC_ID_ADPCM_EA)
	AV_CODEC_ID_ADPCM_G726        = AvCodecID(C.AV_CODEC_ID_ADPCM_G726)
	AV_CODEC_ID_ADPCM_CT          = AvCodecID(C.AV_CODEC_ID_ADPCM_CT)
	AV_CODEC_ID_ADPCM_SWF         = AvCodecID(C.AV_CODEC_ID_ADPCM_SWF)
	AV_CODEC_ID_ADPCM_YAMAHA      = AvCodecID(C.AV_CODEC_ID_ADPCM_YAMAHA)
	AV_CODEC_ID_ADPCM_SBPRO_4     = AvCodecID(C.AV_CODEC_ID_ADPCM_SBPRO_4)
	AV_CODEC_ID_ADPCM_SBPRO_3     = AvCodecID(C.AV_CODEC_ID_ADPCM_SBPRO_3)
	AV_CODEC_ID_ADPCM_SBPRO_2     = AvCodecID(C.AV_CODEC_ID_ADPCM_SBPRO_2)
	AV_CODEC_ID_ADPCM_THP         = AvCodecID(C.AV_CODEC_ID_ADPCM_THP)
	AV_CODEC_ID_ADPCM_IMA_AMV     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_AMV)
	AV_CODEC_ID_ADPCM_EA_R1       = AvCodecID(C.AV_CODEC_ID_ADPCM_EA_R1)
	AV_CODEC_ID_ADPCM_EA_R3       = AvCodecID(C.AV_CODEC_ID_ADPCM_EA_R3)
	AV_CODEC_ID_ADPCM_EA_R2       = AvCodecID(C.AV_CODEC_ID_ADPCM_EA_R2)
	AV_CODEC_ID_ADPCM_IMA_EA_SEAD = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD)
	AV_CODEC_ID_ADPCM_IMA_EA_EACS = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_EA_EACS)
	AV_CODEC_ID_ADPCM_EA_XAS      = AvCodecID(C.AV_CODEC_ID_ADPCM_EA_XAS)
	AV_CODEC_ID_ADPCM_EA_MAXIS_XA = AvCodecID(C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA)
	AV_CODEC_ID_ADPCM_IMA_ISS     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_ISS)
	AV_CODEC_ID_ADPCM_G722        = AvCodecID(C.AV_CODEC_ID_ADPCM_G722)
	AV_CODEC_ID_ADPCM_IMA_APC     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_APC)
	AV_CODEC_ID_ADPCM_VIMA        = AvCodecID(C.AV_CODEC_ID_ADPCM_VIMA)
	AV_CODEC_ID_ADPCM_AFC         = AvCodecID(C.AV_CODEC_ID_ADPCM_AFC)
	AV_CODEC_ID_ADPCM_IMA_OKI     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_OKI)
	AV_CODEC_ID_ADPCM_DTK         = AvCodecID(C.AV_CODEC_ID_ADPCM_DTK)
	AV_CODEC_ID_ADPCM_IMA_RAD     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_RAD)
	AV_CODEC_ID_ADPCM_G726LE      = AvCodecID(C.AV_CODEC_ID_ADPCM_G726LE)
	AV_CODEC_ID_ADPCM_THP_LE      = AvCodecID(C.AV_CODEC_ID_ADPCM_THP_LE)
	AV_CODEC_ID_ADPCM_PSX         = AvCodecID(C.AV_CODEC_ID_ADPCM_PSX)
	AV_CODEC_ID_ADPCM_AICA        = AvCodecID(C.AV_CODEC_ID_ADPCM_AICA)
	AV_CODEC_ID_ADPCM_IMA_DAT4    = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_DAT4)
	AV_CODEC_ID_ADPCM_MTAF        = AvCodecID(C.AV_CODEC_ID_ADPCM_MTAF)
	AV_CODEC_ID_ADPCM_AGM         = AvCodecID(C.AV_CODEC_ID_ADPCM_AGM)
	AV_CODEC_ID_ADPCM_ARGO        = AvCodecID(C.AV_CODEC_ID_ADPCM_ARGO)
	AV_CODEC_ID_ADPCM_IMA_SSI     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_SSI)
	AV_CODEC_ID_ADPCM_ZORK        = AvCodecID(C.AV_CODEC_ID_ADPCM_ZORK)
	AV_CODEC_ID_ADPCM_IMA_APM     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_APM)
	AV_CODEC_ID_ADPCM_IMA_ALP     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_ALP)
	AV_CODEC_ID_ADPCM_IMA_MTF     = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_MTF)
	AV_CODEC_ID_ADPCM_IMA_CUNNING = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_CUNNING)
	AV_CODEC_ID_ADPCM_IMA_MOFLEX  = AvCodecID(C.AV_CODEC_ID_ADPCM_IMA_MOFLEX)

	// AMR
	AV_CODEC_ID_AMR_NB = AvCodecID(C.AV_CODEC_ID_AMR_NB)
	AV_CODEC_ID_AMR_WB = AvCodecID(C.AV_CODEC_ID_AMR_WB)

	// RealAudio codecs
	AV_CODEC_ID_RA_144 = AvCodecID(C.AV_CODEC_ID_RA_144)
	AV_CODEC_ID_RA_288 = AvCodecID(C.AV_CODEC_ID_RA_288)

	// various DPCM codecs
	AV_CODEC_ID_ROQ_DPCM       = AvCodecID(C.AV_CODEC_ID_ROQ_DPCM)
	AV_CODEC_ID_INTERPLAY_DPCM = AvCodecID(C.AV_CODEC_ID_INTERPLAY_DPCM)
	AV_CODEC_ID_XAN_DPCM       = AvCodecID(C.AV_CODEC_ID_XAN_DPCM)
	AV_CODEC_ID_SOL_DPCM       = AvCodecID(C.AV_CODEC_ID_SOL_DPCM)

	AV_CODEC_ID_SDX2_DPCM    = AvCodecID(C.AV_CODEC_ID_SDX2_DPCM)
	AV_CODEC_ID_GREMLIN_DPCM = AvCodecID(C.AV_CODEC_ID_GREMLIN_DPCM)
	AV_CODEC_ID_DERF_DPCM    = AvCodecID(C.AV_CODEC_ID_DERF_DPCM)

	// audio codecs
	AV_CODEC_ID_MP2            = AvCodecID(C.AV_CODEC_ID_MP2)
	AV_CODEC_ID_MP3            = AvCodecID(C.AV_CODEC_ID_MP3)
	AV_CODEC_ID_AAC            = AvCodecID(C.AV_CODEC_ID_AAC)
	AV_CODEC_ID_AC3            = AvCodecID(C.AV_CODEC_ID_AC3)
	AV_CODEC_ID_DTS            = AvCodecID(C.AV_CODEC_ID_DTS)
	AV_CODEC_ID_VORBIS         = AvCodecID(C.AV_CODEC_ID_VORBIS)
	AV_CODEC_ID_DVAUDIO        = AvCodecID(C.AV_CODEC_ID_DVAUDIO)
	AV_CODEC_ID_WMAV1          = AvCodecID(C.AV_CODEC_ID_WMAV1)
	AV_CODEC_ID_WMAV2          = AvCodecID(C.AV_CODEC_ID_WMAV2)
	AV_CODEC_ID_MACE3          = AvCodecID(C.AV_CODEC_ID_MACE3)
	AV_CODEC_ID_MACE6          = AvCodecID(C.AV_CODEC_ID_MACE6)
	AV_CODEC_ID_VMDAUDIO       = AvCodecID(C.AV_CODEC_ID_VMDAUDIO)
	AV_CODEC_ID_FLAC           = AvCodecID(C.AV_CODEC_ID_FLAC)
	AV_CODEC_ID_MP3ADU         = AvCodecID(C.AV_CODEC_ID_MP3ADU)
	AV_CODEC_ID_MP3ON4         = AvCodecID(C.AV_CODEC_ID_MP3ON4)
	AV_CODEC_ID_SHORTEN        = AvCodecID(C.AV_CODEC_ID_SHORTEN)
	AV_CODEC_ID_ALAC           = AvCodecID(C.AV_CODEC_ID_ALAC)
	AV_CODEC_ID_WESTWOOD_SND1  = AvCodecID(C.AV_CODEC_ID_WESTWOOD_SND1)
	AV_CODEC_ID_GSM            = AvCodecID(C.AV_CODEC_ID_GSM)
	AV_CODEC_ID_QDM2           = AvCodecID(C.AV_CODEC_ID_QDM2)
	AV_CODEC_ID_COOK           = AvCodecID(C.AV_CODEC_ID_COOK)
	AV_CODEC_ID_TRUESPEECH     = AvCodecID(C.AV_CODEC_ID_TRUESPEECH)
	AV_CODEC_ID_TTA            = AvCodecID(C.AV_CODEC_ID_TTA)
	AV_CODEC_ID_SMACKAUDIO     = AvCodecID(C.AV_CODEC_ID_SMACKAUDIO)
	AV_CODEC_ID_QCELP          = AvCodecID(C.AV_CODEC_ID_QCELP)
	AV_CODEC_ID_WAVPACK        = AvCodecID(C.AV_CODEC_ID_WAVPACK)
	AV_CODEC_ID_DSICINAUDIO    = AvCodecID(C.AV_CODEC_ID_DSICINAUDIO)
	AV_CODEC_ID_IMC            = AvCodecID(C.AV_CODEC_ID_IMC)
	AV_CODEC_ID_MUSEPACK7      = AvCodecID(C.AV_CODEC_ID_MUSEPACK7)
	AV_CODEC_ID_MLP            = AvCodecID(C.AV_CODEC_ID_MLP)
	AV_CODEC_ID_GSM_MS         = AvCodecID(C.AV_CODEC_ID_GSM_MS)
	AV_CODEC_ID_ATRAC3         = AvCodecID(C.AV_CODEC_ID_ATRAC3)
	AV_CODEC_ID_APE            = AvCodecID(C.AV_CODEC_ID_APE)
	AV_CODEC_ID_NELLYMOSER     = AvCodecID(C.AV_CODEC_ID_NELLYMOSER)
	AV_CODEC_ID_MUSEPACK8      = AvCodecID(C.AV_CODEC_ID_MUSEPACK8)
	AV_CODEC_ID_SPEEX          = AvCodecID(C.AV_CODEC_ID_SPEEX)
	AV_CODEC_ID_WMAVOICE       = AvCodecID(C.AV_CODEC_ID_WMAVOICE)
	AV_CODEC_ID_WMAPRO         = AvCodecID(C.AV_CODEC_ID_WMAPRO)
	AV_CODEC_ID_WMALOSSLESS    = AvCodecID(C.AV_CODEC_ID_WMALOSSLESS)
	AV_CODEC_ID_ATRAC3P        = AvCodecID(C.AV_CODEC_ID_ATRAC3P)
	AV_CODEC_ID_EAC3           = AvCodecID(C.AV_CODEC_ID_EAC3)
	AV_CODEC_ID_SIPR           = AvCodecID(C.AV_CODEC_ID_SIPR)
	AV_CODEC_ID_MP1            = AvCodecID(C.AV_CODEC_ID_MP1)
	AV_CODEC_ID_TWINVQ         = AvCodecID(C.AV_CODEC_ID_TWINVQ)
	AV_CODEC_ID_TRUEHD         = AvCodecID(C.AV_CODEC_ID_TRUEHD)
	AV_CODEC_ID_MP4ALS         = AvCodecID(C.AV_CODEC_ID_MP4ALS)
	AV_CODEC_ID_ATRAC1         = AvCodecID(C.AV_CODEC_ID_ATRAC1)
	AV_CODEC_ID_BINKAUDIO_RDFT = AvCodecID(C.AV_CODEC_ID_BINKAUDIO_RDFT)
	AV_CODEC_ID_BINKAUDIO_DCT  = AvCodecID(C.AV_CODEC_ID_BINKAUDIO_DCT)
	AV_CODEC_ID_AAC_LATM       = AvCodecID(C.AV_CODEC_ID_AAC_LATM)
	AV_CODEC_ID_QDMC           = AvCodecID(C.AV_CODEC_ID_QDMC)
	AV_CODEC_ID_CELT           = AvCodecID(C.AV_CODEC_ID_CELT)
	AV_CODEC_ID_G723_1         = AvCodecID(C.AV_CODEC_ID_G723_1)
	AV_CODEC_ID_G729           = AvCodecID(C.AV_CODEC_ID_G729)
	AV_CODEC_ID_8SVX_EXP       = AvCodecID(C.AV_CODEC_ID_8SVX_EXP)
	AV_CODEC_ID_8SVX_FIB       = AvCodecID(C.AV_CODEC_ID_8SVX_FIB)
	AV_CODEC_ID_BMV_AUDIO      = AvCodecID(C.AV_CODEC_ID_BMV_AUDIO)
	AV_CODEC_ID_RALF           = AvCodecID(C.AV_CODEC_ID_RALF)
	AV_CODEC_ID_IAC            = AvCodecID(C.AV_CODEC_ID_IAC)
	AV_CODEC_ID_ILBC           = AvCodecID(C.AV_CODEC_ID_ILBC)
	AV_CODEC_ID_OPUS           = AvCodecID(C.AV_CODEC_ID_OPUS)
	AV_CODEC_ID_COMFORT_NOISE  = AvCodecID(C.AV_CODEC_ID_COMFORT_NOISE)
	AV_CODEC_ID_TAK            = AvCodecID(C.AV_CODEC_ID_TAK)
	AV_CODEC_ID_METASOUND      = AvCodecID(C.AV_CODEC_ID_METASOUND)
	AV_CODEC_ID_PAF_AUDIO      = AvCodecID(C.AV_CODEC_ID_PAF_AUDIO)
	AV_CODEC_ID_ON2AVC         = AvCodecID(C.AV_CODEC_ID_ON2AVC)
	AV_CODEC_ID_DSS_SP         = AvCodecID(C.AV_CODEC_ID_DSS_SP)
	AV_CODEC_ID_CODEC2         = AvCodecID(C.AV_CODEC_ID_CODEC2)

	AV_CODEC_ID_FFWAVESYNTH     = AvCodecID(C.AV_CODEC_ID_FFWAVESYNTH)
	AV_CODEC_ID_SONIC           = AvCodecID(C.AV_CODEC_ID_SONIC)
	AV_CODEC_ID_SONIC_LS        = AvCodecID(C.AV_CODEC_ID_SONIC_LS)
	AV_CODEC_ID_EVRC            = AvCodecID(C.AV_CODEC_ID_EVRC)
	AV_CODEC_ID_SMV             = AvCodecID(C.AV_CODEC_ID_SMV)
	AV_CODEC_ID_DSD_LSBF        = AvCodecID(C.AV_CODEC_ID_DSD_LSBF)
	AV_CODEC_ID_DSD_MSBF        = AvCodecID(C.AV_CODEC_ID_DSD_MSBF)
	AV_CODEC_ID_DSD_LSBF_PLANAR = AvCodecID(C.AV_CODEC_ID_DSD_LSBF_PLANAR)
	AV_CODEC_ID_DSD_MSBF_PLANAR = AvCodecID(C.AV_CODEC_ID_DSD_MSBF_PLANAR)
	AV_CODEC_ID_4GV             = AvCodecID(C.AV_CODEC_ID_4GV)
	AV_CODEC_ID_INTERPLAY_ACM   = AvCodecID(C.AV_CODEC_ID_INTERPLAY_ACM)
	AV_CODEC_ID_XMA1            = AvCodecID(C.AV_CODEC_ID_XMA1)
	AV_CODEC_ID_XMA2            = AvCodecID(C.AV_CODEC_ID_XMA2)
	AV_CODEC_ID_DST             = AvCodecID(C.AV_CODEC_ID_DST)
	AV_CODEC_ID_ATRAC3AL        = AvCodecID(C.AV_CODEC_ID_ATRAC3AL)
	AV_CODEC_ID_ATRAC3PAL       = AvCodecID(C.AV_CODEC_ID_ATRAC3PAL)
	AV_CODEC_ID_DOLBY_E         = AvCodecID(C.AV_CODEC_ID_DOLBY_E)
	AV_CODEC_ID_APTX            = AvCodecID(C.AV_CODEC_ID_APTX)
	AV_CODEC_ID_APTX_HD         = AvCodecID(C.AV_CODEC_ID_APTX_HD)
	AV_CODEC_ID_SBC             = AvCodecID(C.AV_CODEC_ID_SBC)
	AV_CODEC_ID_ATRAC9          = AvCodecID(C.AV_CODEC_ID_ATRAC9)
	AV_CODEC_ID_HCOM            = AvCodecID(C.AV_CODEC_ID_HCOM)
	AV_CODEC_ID_ACELP_KELVIN    = AvCodecID(C.AV_CODEC_ID_ACELP_KELVIN)
	AV_CODEC_ID_MPEGH_3D_AUDIO  = AvCodecID(C.AV_CODEC_ID_MPEGH_3D_AUDIO)
	AV_CODEC_ID_SIREN           = AvCodecID(C.AV_CODEC_ID_SIREN)
	AV_CODEC_ID_HCA             = AvCodecID(C.AV_CODEC_ID_HCA)
	AV_CODEC_ID_FASTAUDIO       = AvCodecID(C.AV_CODEC_ID_FASTAUDIO)

	// subtitle codecs
	AV_CODEC_ID_FIRST_SUBTITLE    = AvCodecID(C.AV_CODEC_ID_FIRST_SUBTITLE)
	AV_CODEC_ID_DVD_SUBTITLE      = AvCodecID(C.AV_CODEC_ID_DVD_SUBTITLE)
	AV_CODEC_ID_DVB_SUBTITLE      = AvCodecID(C.AV_CODEC_ID_DVB_SUBTITLE)
	AV_CODEC_ID_TEXT              = AvCodecID(C.AV_CODEC_ID_TEXT)
	AV_CODEC_ID_XSUB              = AvCodecID(C.AV_CODEC_ID_XSUB)
	AV_CODEC_ID_SSA               = AvCodecID(C.AV_CODEC_ID_SSA)
	AV_CODEC_ID_MOV_TEXT          = AvCodecID(C.AV_CODEC_ID_MOV_TEXT)
	AV_CODEC_ID_HDMV_PGS_SUBTITLE = AvCodecID(C.AV_CODEC_ID_HDMV_PGS_SUBTITLE)
	AV_CODEC_ID_DVB_TELETEXT      = AvCodecID(C.AV_CODEC_ID_DVB_TELETEXT)
	AV_CODEC_ID_SRT               = AvCodecID(C.AV_CODEC_ID_SRT)

	AV_CODEC_ID_MICRODVD           = AvCodecID(C.AV_CODEC_ID_MICRODVD)
	AV_CODEC_ID_EIA_608            = AvCodecID(C.AV_CODEC_ID_EIA_608)
	AV_CODEC_ID_JACOSUB            = AvCodecID(C.AV_CODEC_ID_JACOSUB)
	AV_CODEC_ID_SAMI               = AvCodecID(C.AV_CODEC_ID_SAMI)
	AV_CODEC_ID_REALTEXT           = AvCodecID(C.AV_CODEC_ID_REALTEXT)
	AV_CODEC_ID_STL                = AvCodecID(C.AV_CODEC_ID_STL)
	AV_CODEC_ID_SUBVIEWER1         = AvCodecID(C.AV_CODEC_ID_SUBVIEWER1)
	AV_CODEC_ID_SUBVIEWER          = AvCodecID(C.AV_CODEC_ID_SUBVIEWER)
	AV_CODEC_ID_SUBRIP             = AvCodecID(C.AV_CODEC_ID_SUBRIP)
	AV_CODEC_ID_WEBVTT             = AvCodecID(C.AV_CODEC_ID_WEBVTT)
	AV_CODEC_ID_MPL2               = AvCodecID(C.AV_CODEC_ID_MPL2)
	AV_CODEC_ID_VPLAYER            = AvCodecID(C.AV_CODEC_ID_VPLAYER)
	AV_CODEC_ID_PJS                = AvCodecID(C.AV_CODEC_ID_PJS)
	AV_CODEC_ID_ASS                = AvCodecID(C.AV_CODEC_ID_ASS)
	AV_CODEC_ID_HDMV_TEXT_SUBTITLE = AvCodecID(C.AV_CODEC_ID_HDMV_TEXT_SUBTITLE)
	AV_CODEC_ID_TTML               = AvCodecID(C.AV_CODEC_ID_TTML)
	AV_CODEC_ID_ARIB_CAPTION       = AvCodecID(C.AV_CODEC_ID_ARIB_CAPTION)

	// other specific kind of codecs (generally used for attachments)
	AV_CODEC_ID_FIRST_UNKNOWN = AvCodecID(C.AV_CODEC_ID_FIRST_UNKNOWN)
	AV_CODEC_ID_TTF           = AvCodecID(C.AV_CODEC_ID_TTF)

	AV_CODEC_ID_SCTE_35   = AvCodecID(C.AV_CODEC_ID_SCTE_35)
	AV_CODEC_ID_EPG       = AvCodecID(C.AV_CODEC_ID_EPG)
	AV_CODEC_ID_BINTEXT   = AvCodecID(C.AV_CODEC_ID_BINTEXT)
	AV_CODEC_ID_XBIN      = AvCodecID(C.AV_CODEC_ID_XBIN)
	AV_CODEC_ID_IDF       = AvCodecID(C.AV_CODEC_ID_IDF)
	AV_CODEC_ID_OTF       = AvCodecID(C.AV_CODEC_ID_OTF)
	AV_CODEC_ID_SMPTE_KLV = AvCodecID(C.AV_CODEC_ID_SMPTE_KLV)
	AV_CODEC_ID_DVD_NAV   = AvCodecID(C.AV_CODEC_ID_DVD_NAV)
	AV_CODEC_ID_TIMED_ID3 = AvCodecID(C.AV_CODEC_ID_TIMED_ID3)
	AV_CODEC_ID_BIN_DATA  = AvCodecID(C.AV_CODEC_ID_BIN_DATA)

	// codec_id is not known (like AV_CODEC_ID_NONE) but lavf should attempt to identify it
	AV_CODEC_ID_PROBE = AvCodecID(C.AV_CODEC_ID_PROBE)

	// Fake codec to indicate a raw MPEG-2 TS stream (only used by libavformat)
	AV_CODEC_ID_MPEG2TS = AvCodecID(C.AV_CODEC_ID_MPEG2TS)

	// Fake codec to indicate a MPEG-4 Systems stream (only used by libavformat)
	AV_CODEC_ID_MPEG4SYSTEMS = AvCodecID(C.AV_CODEC_ID_MPEG4SYSTEMS)

	// Dummy codec for streams containing only metadata information.
	AV_CODEC_ID_FFMETADATA = AvCodecID(C.AV_CODEC_ID_FFMETADATA)
	// Passthrough codec, AVFrames wrapped in AVPacket.
	AV_CODEC_ID_WRAPPED_AVFRAME = AvCodecID(C.AV_CODEC_ID_WRAPPED_AVFRAME)
)

// AvCodecGetType gets the type of the given codec.
func AvCodecGetType(codecID AvCodecID) AvMediaType {
	return (AvMediaType)(C.avcodec_get_type((C.enum_AVCodecID)(codecID)))
}

// AvCodecGetName gets the name of a codec.
func AvCodecGetName(codecID AvCodecID) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(codecID)))
}
