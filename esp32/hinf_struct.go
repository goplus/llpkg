package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type HinfDevS struct {
	CfgData0 struct {
		Val c.Uint32T
	}
	CfgData1 struct {
		Val c.Uint32T
	}
	Reserved8  c.Uint32T
	ReservedC  c.Uint32T
	Reserved10 c.Uint32T
	Reserved14 c.Uint32T
	Reserved18 c.Uint32T
	CfgData7   struct {
		Val c.Uint32T
	}
	CisConf0  c.Uint32T
	CisConf1  c.Uint32T
	CisConf2  c.Uint32T
	CisConf3  c.Uint32T
	CisConf4  c.Uint32T
	CisConf5  c.Uint32T
	CisConf6  c.Uint32T
	CisConf7  c.Uint32T
	CfgData16 struct {
		Val c.Uint32T
	}
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
	Date       c.Uint32T
}
type HinfDevT HinfDevS
