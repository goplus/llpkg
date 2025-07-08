package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PeriBackupDevS struct {
	Config struct {
		Val c.Uint32T
	}
	ApbAddr c.Uint32T
	MemAddr c.Uint32T
	RegMap0 c.Uint32T
	RegMap1 c.Uint32T
	RegMap2 c.Uint32T
	RegMap3 c.Uint32T
	IntRaw  struct {
		Val c.Uint32T
	}
	IntSt struct {
		Val c.Uint32T
	}
	IntEna struct {
		Val c.Uint32T
	}
	IntClr struct {
		Val c.Uint32T
	}
	Reserved2c c.Uint32T
	Reserved30 c.Uint32T
	Reserved34 c.Uint32T
	Reserved38 c.Uint32T
	Reserved3c c.Uint32T
	Reserved40 c.Uint32T
	Reserved44 c.Uint32T
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
	Reserved98 c.Uint32T
	Reserved9c c.Uint32T
	ReservedA0 c.Uint32T
	ReservedA4 c.Uint32T
	ReservedA8 c.Uint32T
	ReservedAc c.Uint32T
	ReservedB0 c.Uint32T
	ReservedB4 c.Uint32T
	ReservedB8 c.Uint32T
	ReservedBc c.Uint32T
	ReservedC0 c.Uint32T
	ReservedC4 c.Uint32T
	ReservedC8 c.Uint32T
	ReservedCc c.Uint32T
	ReservedD0 c.Uint32T
	ReservedD4 c.Uint32T
	ReservedD8 c.Uint32T
	ReservedDc c.Uint32T
	ReservedE0 c.Uint32T
	ReservedE4 c.Uint32T
	ReservedE8 c.Uint32T
	ReservedEc c.Uint32T
	ReservedF0 c.Uint32T
	ReservedF4 c.Uint32T
	ReservedF8 c.Uint32T
	Date       struct {
		Val c.Uint32T
	}
}
type PeriBackupDevT PeriBackupDevS
