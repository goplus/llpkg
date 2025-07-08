package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UartPeriphSigT struct {
	DefaultGpio c.Int32T
	IomuxFunc   c.Int32T
	Input       c.Uint32T
	Signal      c.Uint32T
}

type UartSignalConnT struct {
	Pins [4]UartPeriphSigT
	Irq  c.Uint8T
}
