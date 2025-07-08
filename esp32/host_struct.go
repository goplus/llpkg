package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type HostDevS struct {
	Reserved0  c.Uint32T
	Reserved4  c.Uint32T
	Reserved8  c.Uint32T
	ReservedC  c.Uint32T
	Reserved10 c.Uint32T
	Reserved14 c.Uint32T
	Reserved18 c.Uint32T
	Reserved1c c.Uint32T
	Func22     struct {
		Val c.Uint32T
	}
	Reserved24  c.Uint32T
	Reserved28  c.Uint32T
	Reserved2c  c.Uint32T
	Reserved30  c.Uint32T
	GpioStatus0 c.Uint32T
	GpioStatus1 struct {
		Val c.Uint32T
	}
	GpioIn0 c.Uint32T
	GpioIn1 struct {
		Val c.Uint32T
	}
	Slc0TokenRdata struct {
		Val c.Uint32T
	}
	Slc0Pf     c.Uint32T
	Reserved4c c.Uint32T
	Slc0IntRaw struct {
		Val c.Uint32T
	}
	Reserved54 c.Uint32T
	Slc0IntSt  struct {
		Val c.Uint32T
	}
	Reserved5c c.Uint32T
	PktLen     struct {
		Val c.Uint32T
	}
	StateW0 struct {
		Val c.Uint32T
	}
	StateW1 struct {
		Val c.Uint32T
	}
	ConfW0 struct {
		Val c.Uint32T
	}
	ConfW1 struct {
		Val c.Uint32T
	}
	ConfW2 struct {
		Val c.Uint32T
	}
	ConfW3 struct {
		Val c.Uint32T
	}
	ConfW4 struct {
		Val c.Uint32T
	}
	ConfW5 struct {
		Val c.Uint32T
	}
	WinCmd struct {
		Val c.Uint32T
	}
	ConfW6 struct {
		Val c.Uint32T
	}
	ConfW7 struct {
		Val c.Uint32T
	}
	PktLen0 struct {
		Val c.Uint32T
	}
	PktLen1 struct {
		Val c.Uint32T
	}
	PktLen2 struct {
		Val c.Uint32T
	}
	ConfW8 struct {
		Val c.Uint32T
	}
	ConfW9 struct {
		Val c.Uint32T
	}
	ConfW10 struct {
		Val c.Uint32T
	}
	ConfW11 struct {
		Val c.Uint32T
	}
	ConfW12 struct {
		Val c.Uint32T
	}
	ConfW13 struct {
		Val c.Uint32T
	}
	ConfW14 struct {
		Val c.Uint32T
	}
	ConfW15 struct {
		Val c.Uint32T
	}
	CheckSum0      c.Uint32T
	CheckSum1      c.Uint32T
	ReservedC4     c.Uint32T
	Slc0TokenWdata struct {
		Val c.Uint32T
	}
	ReservedCc c.Uint32T
	TokenCon   struct {
		Val c.Uint32T
	}
	Slc0IntClr struct {
		Val c.Uint32T
	}
	ReservedD8      c.Uint32T
	Slc0Func1IntEna struct {
		Val c.Uint32T
	}
	ReservedE0 c.Uint32T
	ReservedE4 c.Uint32T
	ReservedE8 c.Uint32T
	Slc0IntEna struct {
		Val c.Uint32T
	}
	ReservedF0  c.Uint32T
	Slc0RxInfor struct {
		Val c.Uint32T
	}
	ReservedF8  c.Uint32T
	Slc0LenWd   c.Uint32T
	ApbwinWdata c.Uint32T
	ApbwinConf  struct {
		Val c.Uint32T
	}
	ApbwinRdata c.Uint32T
	Slc0Rdclr   struct {
		Val c.Uint32T
	}
	Reserved110 c.Uint32T
	Slc0IntEna1 struct {
		Val c.Uint32T
	}
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
	Date        c.Uint32T
	Id          c.Uint32T
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
	Conf        struct {
		Val c.Uint32T
	}
	InfSt struct {
		Val c.Uint32T
	}
}
type HostDevT HostDevS
