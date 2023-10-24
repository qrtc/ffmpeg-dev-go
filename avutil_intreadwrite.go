package ffmpeg

/*
#include <libavutil/intreadwrite.h>
*/
import "C"

// See https://pkg.go.dev/encoding/binary

// Arch-specific headers can provide any combination of
// AV_[RW][BLN](16|24|32|48|64) and AV_(COPY|SWAP|ZERO)(64|128) macros.
// Preprocessor symbols must be defined, even if these are implemented
// as inline functions.
//
// R/W means read/write, B/L/N means big/little/native endianness.
// The following macros require aligned access, compared to their
// unaligned variants: AV_(COPY|SWAP|ZERO)(64|128), AV_[RW]N[8-64]A.
// Incorrect usage may range from abysmal performance to crash
// depending on the platform.
//
// The unaligned variants are AV_[RW][BLN][8-64] and AV_COPY*U.

// Define AV_[RW]N helper macros to simplify definitions not provided
// by per-arch headers.

// NONEED: AV_RN

// NONEED: AV_WN

// NONEED: AV_RN16

// NONEED: AV_RN32

// NONEED: AV_RN64

// NONEED: AV_WN16

// NONEED: AV_WN32

// NONEED: AV_WN64

// NONEED: AV_RB

// NONEED: AV_WB

// NONEED: AV_RL

// NONEED: AV_WL

// NONEED: AV_RB8

// NONEED: AV_WB8

// NONEED: AV_RL8

// NONEED: AV_WL8

// NONEED: AV_RB16

// NONEED: AV_WB16

// NONEED: AV_RL16

// NONEED: AV_WL16

// NONEED: AV_RB32

// NONEED: AV_WB32

// NONEED: AV_RL32

// NONEED: AV_WL32

// NONEED: AV_RB64

// NONEED: AV_WB64

// NONEED: AV_RL64

// NONEED: AV_WL64

// NONEED: AV_RB24

// NONEED: AV_WB24

// NONEED: AV_RL24

// NONEED: AV_WL24

// NONEED: AV_RB48

// NONEED: AV_WB48

// NONEED: AV_RL48

// NONEED: AV_WL48

// The AV_[RW]NA macros access naturally aligned data
// in a type-safe way.

// NONEED: AV_RNA

// NONEED: AV_WNA

// NONEED: AV_RN16A

// NONEED: AV_RN32A

// NONEED: AV_RN64A

// NONEED: AV_WN16A

// NONEED: AV_WN32A

// NONEED: AV_WN64A

// NONEED: AV_RLA

// NONEED: AV_WLA

// NONEED: AV_RL64A

// NONEED: AV_WL64A

// The AV_COPYxxU macros are suitable for copying data to/from unaligned
// memory locations.

// NONEED: AV_COPYU

// NONEED: AV_COPY16U

// NONEED: AV_COPY32U

// NONEED: AV_COPY64U

// NONEED: AV_COPY128U

// Parameters for AV_COPY*, AV_SWAP*, AV_ZERO* must be
// naturally aligned. They may be implemented using MMX,
// so emms_c() must be called before using any float code
// afterwards.

// NONEED: AV_COPY

// NONEED: AV_COPY16

// NONEED: AV_COPY32

// NONEED: AV_COPY64

// NONEED: AV_COPY128

// NONEED: AV_SWAP

// NONEED: AV_SWAP64

// NONEED: AV_ZERO

// NONEED: AV_ZERO16

// NONEED: AV_ZERO32

// NONEED: AV_ZERO64

// NONEED: AV_ZERO128
