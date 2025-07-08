package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LedcDevS struct {
	ChannelGroup [1]struct {
		Channel [8]struct {
			Conf0 struct {
				Val c.Uint32T
			}
			Hpoint struct {
				Val c.Uint32T
			}
			Duty struct {
				Val c.Uint32T
			}
			Conf1 struct {
				Val c.Uint32T
			}
			DutyRd struct {
				Val c.Uint32T
			}
		}
	}
	TimerGroup [1]struct {
		Timer [4]struct {
			Conf struct {
				Val c.Uint32T
			}
			Value struct {
				Val c.Uint32T
			}
		}
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
	Conf struct {
		Val c.Uint32T
	}
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
type LedcDevT LedcDevS
