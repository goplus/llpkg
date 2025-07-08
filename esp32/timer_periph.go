package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TimerGroupSignalConnT struct {
	Groups [2]struct {
		Module     PeriphModuleT
		TimerIrqId [2]c.Int
	}
}
