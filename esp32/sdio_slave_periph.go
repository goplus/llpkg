package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** pin and signal information of each slot */

type SdioSlaveSlotInfoT struct {
	ClkGpio c.Uint32T
	CmdGpio c.Uint32T
	D0Gpio  c.Uint32T
	D1Gpio  c.Uint32T
	D2Gpio  c.Uint32T
	D3Gpio  c.Uint32T
	Func    c.Int
}
