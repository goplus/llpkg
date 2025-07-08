package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief UHCI escape sequence
 */

type UhciSeperChrT struct {
	SeperChr c.Uint8T
	SubChr1  c.Uint8T
	SubChr2  c.Uint8T
	SubChrEn bool
}

/**
 * @brief UHCI software flow control
 */

type UhciSwflowCtrlSubChrT struct {
	XonChr   c.Uint8T
	XonSub1  c.Uint8T
	XonSub2  c.Uint8T
	XoffChr  c.Uint8T
	XoffSub1 c.Uint8T
	XoffSub2 c.Uint8T
	FlowEn   c.Uint8T
}
