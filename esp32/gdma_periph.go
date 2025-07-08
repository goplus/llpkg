package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaSignalConnT struct {
	Groups [1]struct {
		Module PeriphModuleT
		Pairs  [5]struct {
			RxIrqId c.Int
			TxIrqId c.Int
		}
	}
}
