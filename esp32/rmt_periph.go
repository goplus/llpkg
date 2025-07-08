package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RmtSignalConnT struct {
	Groups [1]struct {
		Irq      c.Int
		Channels [8]struct {
		}
	}
}
