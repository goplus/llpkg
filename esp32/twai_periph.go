package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TwaiControllerSignalConnT struct {
	Controllers [1]struct {
		Module     PeriphModuleT
		IrqId      c.Int
		TxSig      c.Int
		RxSig      c.Int
		ClkOutSig  c.Int
		BusOffSig  c.Int
		StandBySig c.Int
	}
}
