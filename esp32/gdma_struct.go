package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaDevS struct {
	Channel [5]struct {
		In struct {
			Conf0 struct {
				Val c.Uint32T
			}
			Conf1 struct {
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
			InfifoStatus struct {
				Val c.Uint32T
			}
			Pop struct {
				Val c.Uint32T
			}
			Link struct {
				Val c.Uint32T
			}
			State struct {
				Val c.Uint32T
			}
			SucEofDesAddr c.Uint32T
			ErrEofDesAddr c.Uint32T
			Dscr          c.Uint32T
			DscrBf0       c.Uint32T
			DscrBf1       c.Uint32T
			Weight        struct {
				Val c.Uint32T
			}
			Reserved40 c.Uint32T
			Pri        struct {
				Val c.Uint32T
			}
			PeriSel struct {
				Val c.Uint32T
			}
			Reserved4c c.Uint32T
			Reserved50 c.Uint32T
			Reserved54 c.Uint32T
			Reserved58 c.Uint32T
			Reserved5c c.Uint32T
		}
		Out struct {
			Conf0 struct {
				Val c.Uint32T
			}
			Conf1 struct {
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
			OutfifoStatus struct {
				Val c.Uint32T
			}
			Push struct {
				Val c.Uint32T
			}
			Link struct {
				Val c.Uint32T
			}
			State struct {
				Val c.Uint32T
			}
			EofDesAddr    c.Uint32T
			EofBfrDesAddr c.Uint32T
			Dscr          c.Uint32T
			DscrBf0       c.Uint32T
			DscrBf1       c.Uint32T
			Weight        struct {
				Val c.Uint32T
			}
			ReservedA0 c.Uint32T
			Pri        struct {
				Val c.Uint32T
			}
			PeriSel struct {
				Val c.Uint32T
			}
			ReservedAc c.Uint32T
			ReservedB0 c.Uint32T
			ReservedB4 c.Uint32T
			ReservedB8 c.Uint32T
			ReservedBc c.Uint32T
		}
	}
	AhbTest struct {
		Val c.Uint32T
	}
	PdConf struct {
		Val c.Uint32T
	}
	MiscConf struct {
		Val c.Uint32T
	}
	SramSize [5]struct {
		In struct {
			Val c.Uint32T
		}
		Out struct {
			Val c.Uint32T
		}
	}
	ExtmemRejectAddr c.Uint32T
	ExtmemRejectSt   struct {
		Val c.Uint32T
	}
	ExtmemRejectIntRaw struct {
		Val c.Uint32T
	}
	ExtmemRejectIntSt struct {
		Val c.Uint32T
	}
	ExtmemRejectIntEna struct {
		Val c.Uint32T
	}
	ExtmemRejectIntClr struct {
		Val c.Uint32T
	}
	Date c.Uint32T
}
type GdmaDevT GdmaDevS
