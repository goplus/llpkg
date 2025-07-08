package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* The TWAI peripheral's registers are 8bits, however the ESP32-S3 can only access
 * peripheral registers every 32bits. Therefore each TWAI register is mapped to
 * the least significant byte of every 32bits.
 */

type TwaiDevS struct {
	ModeReg struct {
		Val c.Uint32T
	}
	CommandReg struct {
		Val c.Uint32T
	}
	StatusReg struct {
		Val c.Uint32T
	}
	InterruptReg struct {
		Val c.Uint32T
	}
	InterruptEnableReg struct {
		Val c.Uint32T
	}
	Reserved14    c.Uint32T
	BusTiming0Reg struct {
		Val c.Uint32T
	}
	BusTiming1Reg struct {
		Val c.Uint32T
	}
	Reserved20               c.Uint32T
	Reserved24               c.Uint32T
	Reserved28               c.Uint32T
	ArbitrationLostCaptueReg struct {
		Val c.Uint32T
	}
	ErrorCodeCaptureReg struct {
		Val c.Uint32T
	}
	ErrorWarningLimitReg struct {
		Val c.Uint32T
	}
	RxErrorCounterReg struct {
		Val c.Uint32T
	}
	TxErrorCounterReg struct {
		Val c.Uint32T
	}
	RxMessageCounterReg struct {
		Val c.Uint32T
	}
	Reserved78      c.Uint32T
	ClockDividerReg struct {
		Val c.Uint32T
	}
}
type TwaiDevT TwaiDevS
