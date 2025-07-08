package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2sDevS struct {
	Reserved0 c.Uint32T
	Reserved4 c.Uint32T
	Reserved8 c.Uint32T
	IntRaw    struct {
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
	Reserved1c c.Uint32T
	RxConf     struct {
		Val c.Uint32T
	}
	TxConf struct {
		Val c.Uint32T
	}
	RxConf1 struct {
		Val c.Uint32T
	}
	TxConf1 struct {
		Val c.Uint32T
	}
	RxClkmConf struct {
		Val c.Uint32T
	}
	TxClkmConf struct {
		Val c.Uint32T
	}
	RxClkmDivConf struct {
		Val c.Uint32T
	}
	TxClkmDivConf struct {
		Val c.Uint32T
	}
	TxPcm2pdmConf struct {
		Val c.Uint32T
	}
	TxPcm2pdmConf1 struct {
		Val c.Uint32T
	}
	Reserved48 c.Uint32T
	Reserved4c c.Uint32T
	RxTdmCtrl  struct {
		Val c.Uint32T
	}
	TxTdmCtrl struct {
		Val c.Uint32T
	}
	RxTiming struct {
		Val c.Uint32T
	}
	TxTiming struct {
		Val c.Uint32T
	}
	LcHungConf struct {
		Val c.Uint32T
	}
	RxEofNum struct {
		Val c.Uint32T
	}
	ConfSingleData c.Uint32T
	State          struct {
		Val c.Uint32T
	}
	Reserved70 c.Uint32T
	Reserved74 c.Uint32T
	Reserved78 c.Uint32T
	Reserved7c c.Uint32T
	Date       struct {
		Val c.Uint32T
	}
}
type I2sDevT I2sDevS
