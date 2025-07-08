package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioDevS struct {
	BtSelect c.Uint32T
	Out      c.Uint32T
	OutW1ts  c.Uint32T
	OutW1tc  c.Uint32T
	Out1     struct {
		Val c.Uint32T
	}
	Out1W1ts struct {
		Val c.Uint32T
	}
	Out1W1tc struct {
		Val c.Uint32T
	}
	SdioSelect struct {
		Val c.Uint32T
	}
	Enable     c.Uint32T
	EnableW1ts c.Uint32T
	EnableW1tc c.Uint32T
	Enable1    struct {
		Val c.Uint32T
	}
	Enable1W1ts struct {
		Val c.Uint32T
	}
	Enable1W1tc struct {
		Val c.Uint32T
	}
	Strap struct {
		Val c.Uint32T
	}
	In  c.Uint32T
	In1 struct {
		Val c.Uint32T
	}
	Status     c.Uint32T
	StatusW1ts c.Uint32T
	StatusW1tc c.Uint32T
	Status1    struct {
		Val c.Uint32T
	}
	Status1W1ts struct {
		Val c.Uint32T
	}
	Status1W1tc struct {
		Val c.Uint32T
	}
	PcpuInt    c.Uint32T
	PcpuNmiInt c.Uint32T
	CpusdioInt c.Uint32T
	PcpuInt1   struct {
		Val c.Uint32T
	}
	PcpuNmiInt1 struct {
		Val c.Uint32T
	}
	CpusdioInt1 struct {
		Val c.Uint32T
	}
	Pin [54]struct {
		Val c.Uint32T
	}
	StatusNext  c.Uint32T
	StatusNext1 struct {
		Val c.Uint32T
	}
	FuncInSelCfg [256]struct {
		Val c.Uint32T
	}
	FuncOutSelCfg [54]struct {
		Val c.Uint32T
	}
	ClockGate struct {
		Val c.Uint32T
	}
	Reserved630 c.Uint32T
	Reserved634 c.Uint32T
	Reserved638 c.Uint32T
	Reserved63c c.Uint32T
	Reserved640 c.Uint32T
	Reserved644 c.Uint32T
	Reserved648 c.Uint32T
	Reserved64c c.Uint32T
	Reserved650 c.Uint32T
	Reserved654 c.Uint32T
	Reserved658 c.Uint32T
	Reserved65c c.Uint32T
	Reserved660 c.Uint32T
	Reserved664 c.Uint32T
	Reserved668 c.Uint32T
	Reserved66c c.Uint32T
	Reserved670 c.Uint32T
	Reserved674 c.Uint32T
	Reserved678 c.Uint32T
	Reserved67c c.Uint32T
	Reserved680 c.Uint32T
	Reserved684 c.Uint32T
	Reserved688 c.Uint32T
	Reserved68c c.Uint32T
	Reserved690 c.Uint32T
	Reserved694 c.Uint32T
	Reserved698 c.Uint32T
	Reserved69c c.Uint32T
	Reserved6a0 c.Uint32T
	Reserved6a4 c.Uint32T
	Reserved6a8 c.Uint32T
	Reserved6ac c.Uint32T
	Reserved6b0 c.Uint32T
	Reserved6b4 c.Uint32T
	Reserved6b8 c.Uint32T
	Reserved6bc c.Uint32T
	Reserved6c0 c.Uint32T
	Reserved6c4 c.Uint32T
	Reserved6c8 c.Uint32T
	Reserved6cc c.Uint32T
	Reserved6d0 c.Uint32T
	Reserved6d4 c.Uint32T
	Reserved6d8 c.Uint32T
	Reserved6dc c.Uint32T
	Reserved6e0 c.Uint32T
	Reserved6e4 c.Uint32T
	Reserved6e8 c.Uint32T
	Reserved6ec c.Uint32T
	Reserved6f0 c.Uint32T
	Reserved6f4 c.Uint32T
	Reserved6f8 c.Uint32T
	Date        struct {
		Val c.Uint32T
	}
}
type GpioDevT GpioDevS
