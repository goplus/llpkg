package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SDMMC_DMA_MAX_BUF_LEN = 4096

type SdmmcDescS struct {
	Buffer1Ptr c.Pointer
}
type SdmmcDescT SdmmcDescS

type SdmmcHwCmdS struct {
	CmdIndex         c.Uint32T
	ResponseExpect   c.Uint32T
	ResponseLong     c.Uint32T
	CheckResponseCrc c.Uint32T
	DataExpected     c.Uint32T
	Rw               c.Uint32T
	StreamMode       c.Uint32T
	SendAutoStop     c.Uint32T
	WaitComplete     c.Uint32T
	StopAbortCmd     c.Uint32T
	SendInit         c.Uint32T
	CardNum          c.Uint32T
	UpdateClkReg     c.Uint32T
	ReadCeata        c.Uint32T
	CcsExpected      c.Uint32T
	EnableBoot       c.Uint32T
	ExpectBootAck    c.Uint32T
	DisableBoot      c.Uint32T
	BootMode         c.Uint32T
	VoltSwitch       c.Uint32T
	UseHoldReg       c.Uint32T
	Reserved         c.Uint32T
	StartCommand     c.Uint32T
}
type SdmmcHwCmdT SdmmcHwCmdS

type SdmmcDevT struct {
	Ctrl struct {
		Val c.Uint32T
	}
	Pwren  c.Uint32T
	Clkdiv struct {
		Val c.Uint32T
	}
	Clksrc struct {
		Val c.Uint32T
	}
	Clkena struct {
		Val c.Uint32T
	}
	Tmout struct {
		Val c.Uint32T
	}
	Ctype struct {
		Val c.Uint32T
	}
	Blksiz struct {
		Val c.Uint32T
	}
	Bytcnt  c.Uint32T
	Intmask struct {
		Val c.Uint32T
	}
	Cmdarg  c.Uint32T
	Cmd     SdmmcHwCmdT
	Resp    [4]c.Uint32T
	Mintsts struct {
		Val c.Uint32T
	}
	Rintsts struct {
		Val c.Uint32T
	}
	Status struct {
		Val c.Uint32T
	}
	Fifoth struct {
		Val c.Uint32T
	}
	Cdetect struct {
		Val c.Uint32T
	}
	Wrtprt struct {
		Val c.Uint32T
	}
	Gpio   c.Uint32T
	Tcbcnt c.Uint32T
	Tbbcnt c.Uint32T
	Debnce struct {
	}
	Usrid c.Uint32T
	Verid c.Uint32T
	Hcon  struct {
		Val c.Uint32T
	}
	Uhs struct {
		Val c.Uint32T
	}
	RstN struct {
	}
	Reserved7c c.Uint32T
	Bmod       struct {
		Val c.Uint32T
	}
	Pldmnd c.Uint32T
	Dbaddr *SdmmcDescT
	Idsts  struct {
		Val c.Uint32T
	}
	Idinten struct {
		Val c.Uint32T
	}
	Dscaddr    c.Uint32T
	Dscaddrl   c.Uint32T
	Dscaddru   c.Uint32T
	Bufaddrl   c.Uint32T
	Bufaddru   c.Uint32T
	ReservedA8 [22]c.Uint32T
	Cardthrctl struct {
		Val c.Uint32T
	}
	BackEndPower c.Uint32T
	UhsRegExt    c.Uint32T
	EmmcDdrReg   c.Uint32T
	EnableShift  c.Uint32T
	Reserved114  [443]c.Uint32T
	Clock        struct {
		Val c.Uint32T
	}
}
