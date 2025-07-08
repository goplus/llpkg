package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief MIPI CSI PHY PLL frequency range
 */

type SocMipiCsiPhyPllFreqRangeT struct {
	StartMbps      c.Uint32T
	EndMbps        c.Uint32T
	HsFreqRangeSel c.Uint8T
}
