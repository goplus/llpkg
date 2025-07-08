package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
Stores a bunch of per-I2S-peripheral data.
*/
type I2sSignalConnT struct {
	MckOutSig c.Uint8T
	MckInSig  c.Uint8T
	MTxBckSig c.Uint8T
	MRxBckSig c.Uint8T
	MTxWsSig  c.Uint8T
	MRxWsSig  c.Uint8T
	STxBckSig c.Uint8T
	SRxBckSig c.Uint8T
	STxWsSig  c.Uint8T
	SRxWsSig  c.Uint8T
	Irq       c.Uint8T
}
