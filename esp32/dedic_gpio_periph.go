package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DedicGpioSignalConnT struct {
	Irq   c.Int
	Cores [2]struct {
		InSigPerChannel  [8]c.Int
		OutSigPerChannel [8]c.Int
	}
}
