package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UhciDevS struct {
	Conf0 struct {
		Val c.Uint32T
	}
	IntRaw struct {
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
	AppIntSet struct {
		Val c.Uint32T
	}
	Conf1 struct {
		Val c.Uint32T
	}
	State0 struct {
		Val c.Uint32T
	}
	State1 struct {
		Val c.Uint32T
	}
	EscapeConf struct {
		Val c.Uint32T
	}
	HungConf struct {
		Val c.Uint32T
	}
	AckNum struct {
		Val c.Uint32T
	}
	RxHead    c.Uint32T
	QuickSent struct {
		Val c.Uint32T
	}
	QData [7]struct {
		Word [2]c.Uint32T
	}
	EscConf0 struct {
		Val c.Uint32T
	}
	EscConf1 struct {
		Val c.Uint32T
	}
	EscConf2 struct {
		Val c.Uint32T
	}
	EscConf3 struct {
		Val c.Uint32T
	}
	PktThres struct {
		Val c.Uint32T
	}
	Date c.Uint32T
}
type UhciDevT UhciDevS
