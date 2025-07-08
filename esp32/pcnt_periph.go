package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PcntSignalConnT struct {
	Groups [1]struct {
		Units [4]struct {
			Channels [2]struct {
				PulseSig   c.Uint32T
				ControlSig c.Uint32T
			}
			ClearSig c.Uint32T
		}
		Irq c.Uint32T
	}
}
