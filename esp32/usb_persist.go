package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USBDC_TESTAMENT_LOC_MASK = 0x7FFFF

// The testament is a FIFO. The ROM will output all data between textstart and textend; if textend is lower than textstart it will
// output everything from textstart to memend, then memstart to textend.
type UsbdcTestamentT struct {
	Memstart  *c.Char
	Memend    *c.Char
	Textstart *c.Char
	Textend   *c.Char
}
