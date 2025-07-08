package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
Stores a bunch of per-ledc-peripheral data.
*/
type LedcSignalConnT struct {
	SigOut0Idx c.Uint8T
}
