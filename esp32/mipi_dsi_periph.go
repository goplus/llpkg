package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief MIPI DSI PHY PLL frequency range
 */

type SocMipiDsiPhyPllFreqRangeT struct {
	StartMbps      c.Uint32T
	EndMbps        c.Uint32T
	HsFreqRangeSel c.Uint8T
}
