package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RtcIoDevS struct {
	Out struct {
		Val c.Uint32T
	}
	OutW1ts struct {
		Val c.Uint32T
	}
	OutW1tc struct {
		Val c.Uint32T
	}
	Enable struct {
		Val c.Uint32T
	}
	EnableW1ts struct {
		Val c.Uint32T
	}
	EnableW1tc struct {
		Val c.Uint32T
	}
	Status struct {
		Val c.Uint32T
	}
	StatusW1ts struct {
		Val c.Uint32T
	}
	StatusW1tc struct {
		Val c.Uint32T
	}
	InVal struct {
		Val c.Uint32T
	}
	Pin [22]struct {
		Val c.Uint32T
	}
	DebugSel struct {
		Val c.Uint32T
	}
	TouchPad [15]struct {
		Val c.Uint32T
	}
	Xtal32pPad struct {
		Val c.Uint32T
	}
	Xtal32nPad struct {
		Val c.Uint32T
	}
	PadDac [2]struct {
		Val c.Uint32T
	}
	RtcPad19 struct {
		Val c.Uint32T
	}
	RtcPad20 struct {
		Val c.Uint32T
	}
	RtcPad21 struct {
		Val c.Uint32T
	}
	ExtWakeup0 struct {
		Val c.Uint32T
	}
	XtlExtCtr struct {
		Val c.Uint32T
	}
	SarI2cIo struct {
		Val c.Uint32T
	}
	TouchCtrl struct {
		Val c.Uint32T
	}
	ReservedEc  c.Uint32T
	ReservedF0  c.Uint32T
	ReservedF4  c.Uint32T
	ReservedF8  c.Uint32T
	ReservedFc  c.Uint32T
	Reserved100 c.Uint32T
	Reserved104 c.Uint32T
	Reserved108 c.Uint32T
	Reserved10c c.Uint32T
	Reserved110 c.Uint32T
	Reserved114 c.Uint32T
	Reserved118 c.Uint32T
	Reserved11c c.Uint32T
	Reserved120 c.Uint32T
	Reserved124 c.Uint32T
	Reserved128 c.Uint32T
	Reserved12c c.Uint32T
	Reserved130 c.Uint32T
	Reserved134 c.Uint32T
	Reserved138 c.Uint32T
	Reserved13c c.Uint32T
	Reserved140 c.Uint32T
	Reserved144 c.Uint32T
	Reserved148 c.Uint32T
	Reserved14c c.Uint32T
	Reserved150 c.Uint32T
	Reserved154 c.Uint32T
	Reserved158 c.Uint32T
	Reserved15c c.Uint32T
	Reserved160 c.Uint32T
	Reserved164 c.Uint32T
	Reserved168 c.Uint32T
	Reserved16c c.Uint32T
	Reserved170 c.Uint32T
	Reserved174 c.Uint32T
	Reserved178 c.Uint32T
	Reserved17c c.Uint32T
	Reserved180 c.Uint32T
	Reserved184 c.Uint32T
	Reserved188 c.Uint32T
	Reserved18c c.Uint32T
	Reserved190 c.Uint32T
	Reserved194 c.Uint32T
	Reserved198 c.Uint32T
	Reserved19c c.Uint32T
	Reserved1a0 c.Uint32T
	Reserved1a4 c.Uint32T
	Reserved1a8 c.Uint32T
	Reserved1ac c.Uint32T
	Reserved1b0 c.Uint32T
	Reserved1b4 c.Uint32T
	Reserved1b8 c.Uint32T
	Reserved1bc c.Uint32T
	Reserved1c0 c.Uint32T
	Reserved1c4 c.Uint32T
	Reserved1c8 c.Uint32T
	Reserved1cc c.Uint32T
	Reserved1d0 c.Uint32T
	Reserved1d4 c.Uint32T
	Reserved1d8 c.Uint32T
	Reserved1dc c.Uint32T
	Reserved1e0 c.Uint32T
	Reserved1e4 c.Uint32T
	Reserved1e8 c.Uint32T
	Reserved1ec c.Uint32T
	Reserved1f0 c.Uint32T
	Reserved1f4 c.Uint32T
	Reserved1f8 c.Uint32T
	Date        struct {
		Val c.Uint32T
	}
}
type RtcIoDevT RtcIoDevS
