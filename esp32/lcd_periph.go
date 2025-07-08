package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LcdI80SignalConnT struct {
	Buses [1]struct {
		Module   PeriphModuleT
		IrqId    c.Int
		DataSigs [16]c.Int
		CsSig    c.Int
		DcSig    c.Int
		WrSig    c.Int
	}
}

type LcdRgbSignalConnT struct {
	Panels [1]struct {
		Module   PeriphModuleT
		IrqId    c.Int
		DataSigs [16]c.Int
		HsyncSig c.Int
		VsyncSig c.Int
		PclkSig  c.Int
		DeSig    c.Int
	}
}
