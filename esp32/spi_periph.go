package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
Stores a bunch of per-spi-peripheral data.
*/
type SpiSignalConnT struct {
	SpiclkOut      c.Uint8T
	SpiclkIn       c.Uint8T
	SpidOut        c.Uint8T
	SpiqOut        c.Uint8T
	SpiwpOut       c.Uint8T
	SpihdOut       c.Uint8T
	Spid4Out       c.Uint8T
	Spid5Out       c.Uint8T
	Spid6Out       c.Uint8T
	Spid7Out       c.Uint8T
	SpidIn         c.Uint8T
	SpiqIn         c.Uint8T
	SpiwpIn        c.Uint8T
	SpihdIn        c.Uint8T
	Spid4In        c.Uint8T
	Spid5In        c.Uint8T
	Spid6In        c.Uint8T
	Spid7In        c.Uint8T
	SpicsOut       [6]c.Uint8T
	SpicsIn        c.Uint8T
	SpidqsOut      c.Uint8T
	SpicdOut       c.Uint8T
	SpiclkIomuxPin c.Uint8T
	SpidIomuxPin   c.Uint8T
	SpiqIomuxPin   c.Uint8T
	SpiwpIomuxPin  c.Uint8T
	SpihdIomuxPin  c.Uint8T
	Spics0IomuxPin c.Uint8T
	Irq            c.Uint8T
	IrqDma         c.Uint8T
	Func           c.Int
	Hw             *SpiDevT
}
