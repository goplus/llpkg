package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SigmaDeltaSignalConnT struct {
	Channels [8]struct {
		SdSig c.Int
	}
}
