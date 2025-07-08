package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RmtClockSourceT SocPeriphRmtClkSrcT

/**
 * @brief The layout of RMT symbol stored in memory, which is decided by the hardware design
 */

type RmtSymbolWordT struct {
	Val c.Uint32T
}
