package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioSdDevT struct {
	Channel [8]struct {
		Val c.Uint32T
	}
	Cg struct {
		Val c.Uint32T
	}
	Misc struct {
		Val c.Uint32T
	}
	Version struct {
		Val c.Uint32T
	}
}
