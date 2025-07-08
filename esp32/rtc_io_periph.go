package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Pin function information for a single RTCIO pad's.
 *
 * This is an internal function of the driver, and is not usually useful
 * for external use.
 */

type RtcIoDescT struct {
	Reg       c.Uint32T
	Mux       c.Uint32T
	Func      c.Uint32T
	Ie        c.Uint32T
	Pullup    c.Uint32T
	Pulldown  c.Uint32T
	Slpsel    c.Uint32T
	Slpie     c.Uint32T
	Slpoe     c.Uint32T
	Hold      c.Uint32T
	HoldForce c.Uint32T
	DrvV      c.Uint32T
	DrvS      c.Uint32T
	RtcNum    c.Int
}
