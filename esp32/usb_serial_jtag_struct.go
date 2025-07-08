package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UsbSerialJtagDevS struct {
	Ep1 struct {
		Val c.Uint32T
	}
	Ep1Conf struct {
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
	Conf0 struct {
		Val c.Uint32T
	}
	Test struct {
		Val c.Uint32T
	}
	JfifoSt struct {
		Val c.Uint32T
	}
	FramNum struct {
		Val c.Uint32T
	}
	InEp0St struct {
		Val c.Uint32T
	}
	InEp1St struct {
		Val c.Uint32T
	}
	InEp2St struct {
		Val c.Uint32T
	}
	InEp3St struct {
		Val c.Uint32T
	}
	OutEp0St struct {
		Val c.Uint32T
	}
	OutEp1St struct {
		Val c.Uint32T
	}
	OutEp2St struct {
		Val c.Uint32T
	}
	MiscConf struct {
		Val c.Uint32T
	}
	MemConf struct {
		Val c.Uint32T
	}
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
	Date       c.Uint32T
}
type UsbSerialJtagDevT UsbSerialJtagDevS
