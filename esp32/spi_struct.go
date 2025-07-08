package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SpiDevS struct {
	Cmd struct {
		Val c.Uint32T
	}
	Addr c.Uint32T
	Ctrl struct {
		Val c.Uint32T
	}
	Clock struct {
		Val c.Uint32T
	}
	User struct {
		Val c.Uint32T
	}
	User1 struct {
		Val c.Uint32T
	}
	User2 struct {
		Val c.Uint32T
	}
	MsDlen struct {
		Val c.Uint32T
	}
	Misc struct {
		Val c.Uint32T
	}
	DinMode struct {
		Val c.Uint32T
	}
	DinNum struct {
		Val c.Uint32T
	}
	DoutMode struct {
		Val c.Uint32T
	}
	DmaConf struct {
		Val c.Uint32T
	}
	DmaIntEna struct {
		Val c.Uint32T
	}
	DmaIntClr struct {
		Val c.Uint32T
	}
	DmaIntRaw struct {
		Val c.Uint32T
	}
	DmaIntSt struct {
		Val c.Uint32T
	}
	DmaIntSet struct {
		Val c.Uint32T
	}
	Reserved48 c.Uint32T
	Reserved4c c.Uint32T
	Reserved50 c.Uint32T
	Reserved54 c.Uint32T
	Reserved58 c.Uint32T
	Reserved5c c.Uint32T
	Reserved60 c.Uint32T
	Reserved64 c.Uint32T
	Reserved68 c.Uint32T
	Reserved6c c.Uint32T
	Reserved70 c.Uint32T
	Reserved74 c.Uint32T
	Reserved78 c.Uint32T
	Reserved7c c.Uint32T
	Reserved80 c.Uint32T
	Reserved84 c.Uint32T
	Reserved88 c.Uint32T
	Reserved8c c.Uint32T
	Reserved90 c.Uint32T
	Reserved94 c.Uint32T
	DataBuf    [16]c.Uint32T
	ReservedD8 c.Uint32T
	ReservedDc c.Uint32T
	Slave      struct {
		Val c.Uint32T
	}
	Slave1 struct {
		Val c.Uint32T
	}
	ClkGate struct {
		Val c.Uint32T
	}
	ReservedEc c.Uint32T
	Date       struct {
		Val c.Uint32T
	}
}
type SpiDevT SpiDevS
