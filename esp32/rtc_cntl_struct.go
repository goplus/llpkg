package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RtcCntlDevS struct {
	Options0 struct {
		Val c.Uint32T
	}
	SlpTimer0 c.Uint32T
	SlpTimer1 struct {
		Val c.Uint32T
	}
	TimeUpdate struct {
		Val c.Uint32T
	}
	TimeLow0  c.Uint32T
	TimeHigh0 struct {
		Val c.Uint32T
	}
	State0 struct {
		Val c.Uint32T
	}
	Timer1 struct {
		Val c.Uint32T
	}
	Timer2 struct {
		Val c.Uint32T
	}
	Timer3 struct {
		Val c.Uint32T
	}
	Timer4 struct {
		Val c.Uint32T
	}
	Timer5 struct {
		Val c.Uint32T
	}
	Timer6 struct {
		Val c.Uint32T
	}
	AnaConf struct {
		Val c.Uint32T
	}
	ResetState struct {
		Val c.Uint32T
	}
	WakeupState struct {
		Val c.Uint32T
	}
	IntEna struct {
		Val c.Uint32T
	}
	IntRaw struct {
		Val c.Uint32T
	}
	IntSt struct {
		Val c.Uint32T
	}
	IntClr struct {
		Val c.Uint32T
	}
	Store      [4]c.Uint32T
	ExtXtlConf struct {
		Val c.Uint32T
	}
	ExtWakeupConf struct {
		Val c.Uint32T
	}
	SlpRejectConf struct {
		Val c.Uint32T
	}
	CpuPeriodConf struct {
		Val c.Uint32T
	}
	SdioActConf struct {
		Val c.Uint32T
	}
	ClkConf struct {
		Val c.Uint32T
	}
	SlowClkConf struct {
		Val c.Uint32T
	}
	SdioConf struct {
		Val c.Uint32T
	}
	BiasConf struct {
		Val c.Uint32T
	}
	Rtc struct {
		Val c.Uint32T
	}
	Pwc struct {
		Val c.Uint32T
	}
	RegulatorDrvCtrl struct {
		Val c.Uint32T
	}
	DigPwc struct {
		Val c.Uint32T
	}
	DigIso struct {
		Val c.Uint32T
	}
	WdtConfig0 struct {
		Val c.Uint32T
	}
	WdtConfig1 c.Uint32T
	WdtConfig2 c.Uint32T
	WdtConfig3 c.Uint32T
	WdtConfig4 c.Uint32T
	WdtFeed    struct {
		Val c.Uint32T
	}
	WdtWprotect c.Uint32T
	SwdConf     struct {
		Val c.Uint32T
	}
	SwdWprotect c.Uint32T
	SwCpuStall  struct {
		Val c.Uint32T
	}
	Store4     c.Uint32T
	Store5     c.Uint32T
	Store6     c.Uint32T
	Store7     c.Uint32T
	LowPowerSt struct {
		Val c.Uint32T
	}
	Diag0   c.Uint32T
	PadHold struct {
		Val c.Uint32T
	}
	DigPadHold c.Uint32T
	ExtWakeup1 struct {
		Val c.Uint32T
	}
	ExtWakeup1Status struct {
		Val c.Uint32T
	}
	BrownOut struct {
		Val c.Uint32T
	}
	TimeLow1  c.Uint32T
	TimeHigh1 struct {
		Val c.Uint32T
	}
	Xtal32kClkFactor c.Uint32T
	Xtal32kConf      struct {
		Val c.Uint32T
	}
	UlpCpTimer struct {
		Val c.Uint32T
	}
	UlpCpCtrl struct {
		Val c.Uint32T
	}
	CocpuCtrl struct {
		Val c.Uint32T
	}
	TouchCtrl1 struct {
		Val c.Uint32T
	}
	TouchCtrl2 struct {
		Val c.Uint32T
	}
	TouchScanCtrl struct {
		Val c.Uint32T
	}
	TouchSlpThres struct {
		Val c.Uint32T
	}
	TouchApproach struct {
		Val c.Uint32T
	}
	TouchFilterCtrl struct {
		Val c.Uint32T
	}
	UsbConf struct {
		Val c.Uint32T
	}
	TouchTimeoutCtrl struct {
		Val c.Uint32T
	}
	SlpRejectCause struct {
		Val c.Uint32T
	}
	Option1 struct {
		Val c.Uint32T
	}
	SlpWakeupCause struct {
		Val c.Uint32T
	}
	UlpCpTimer1 struct {
		Val c.Uint32T
	}
	IntEnaW1ts struct {
		Val c.Uint32T
	}
	IntEnaW1tc struct {
		Val c.Uint32T
	}
	RetentionCtrl struct {
		Val c.Uint32T
	}
	PgCtrl struct {
		Val c.Uint32T
	}
	FibSel struct {
		Val c.Uint32T
	}
	TouchDac struct {
		Val c.Uint32T
	}
	TouchDac1 struct {
		Val c.Uint32T
	}
	CocpuDisable struct {
		Val c.Uint32T
	}
	Reserved158 c.Uint32T
	Reserved15c c.Uint32T
	Reserved160 c.Uint32T
	Reserved164 c.Uint32T
	Reserved168 c.Uint32T
	Reserved16c c.Uint32T
	Reserved170 c.Uint32T
	Reserved174 c.Uint32T
	Reserved178 c.Uint32T
	Reserved17c c.Uint32T
	Reserved180 c.Uint32T
	Reserved184 c.Uint32T
	Reserved188 c.Uint32T
	Reserved18c c.Uint32T
	Reserved190 c.Uint32T
	Reserved194 c.Uint32T
	Reserved198 c.Uint32T
	Reserved19c c.Uint32T
	Reserved1a0 c.Uint32T
	Reserved1a4 c.Uint32T
	Reserved1a8 c.Uint32T
	Reserved1ac c.Uint32T
	Reserved1b0 c.Uint32T
	Reserved1b4 c.Uint32T
	Reserved1b8 c.Uint32T
	Reserved1bc c.Uint32T
	Reserved1c0 c.Uint32T
	Reserved1c4 c.Uint32T
	Reserved1c8 c.Uint32T
	Reserved1cc c.Uint32T
	Reserved1d0 c.Uint32T
	Reserved1d4 c.Uint32T
	Reserved1d8 c.Uint32T
	Reserved1dc c.Uint32T
	Reserved1e0 c.Uint32T
	Reserved1e4 c.Uint32T
	Reserved1e8 c.Uint32T
	Reserved1ec c.Uint32T
	Reserved1f0 c.Uint32T
	Reserved1f4 c.Uint32T
	Reserved1f8 c.Uint32T
	Date        struct {
		Val c.Uint32T
	}
}
type RtcCntlDevT RtcCntlDevS
