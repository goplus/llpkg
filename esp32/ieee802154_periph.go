package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type Ieee802154ConnT struct {
	Module PeriphModuleT
	IrqId  c.Int
}
