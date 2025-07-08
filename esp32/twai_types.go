package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const TWAI_EXTD_ID_MASK = 0x1FFFFFFF
const TWAI_STD_ID_MASK = 0x7FF
const TWAI_FRAME_MAX_DLC = 8
const TWAI_FRAME_EXTD_ID_LEN_BYTES = 4
const TWAI_FRAME_STD_ID_LEN_BYTES = 2
const TWAI_ERR_PASS_THRESH = 128
const TWAI_MSG_FLAG_NONE = 0x00
const TWAI_MSG_FLAG_EXTD = 0x01
const TWAI_MSG_FLAG_RTR = 0x02
const TWAI_MSG_FLAG_SS = 0x04
const TWAI_MSG_FLAG_SELF = 0x08
const TWAI_MSG_FLAG_DLC_NON_COMP = 0x10

type TwaiModeT c.Int

const (
	TWAI_MODE_NORMAL      TwaiModeT = 0
	TWAI_MODE_NO_ACK      TwaiModeT = 1
	TWAI_MODE_LISTEN_ONLY TwaiModeT = 2
)

/**
 * @brief   Structure to store a TWAI message
 *
 * @note    The flags member is deprecated
 */

type TwaiMessageT struct {
	Identifier     c.Uint32T
	DataLengthCode c.Uint8T
	Data           [8]c.Uint8T
}
type TwaiClockSourceT SocPeriphTwaiClkSrcT

/**
 * @brief   Structure for bit timing configuration of the TWAI driver
 *
 * @note    Macro initializers are available for this structure
 */

type TwaiTimingConfigT struct {
	ClkSrc             TwaiClockSourceT
	QuantaResolutionHz c.Uint32T
	Brp                c.Uint32T
	Tseg1              c.Uint8T
	Tseg2              c.Uint8T
	Sjw                c.Uint8T
	TripleSampling     bool
}

/**
 * @brief   Structure for acceptance filter configuration of the TWAI driver (see documentation)
 *
 * @note    Macro initializers are available for this structure
 */

type TwaiFilterConfigT struct {
	AcceptanceCode c.Uint32T
	AcceptanceMask c.Uint32T
	SingleFilter   bool
}
