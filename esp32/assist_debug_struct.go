package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AssistDebugDevS struct {
	Core0InterruptEna struct {
		Val c.Uint32T
	}
	Core0InterruptRaw struct {
		Val c.Uint32T
	}
	Core0InterruptRls struct {
		Val c.Uint32T
	}
	Core0InterruptClr struct {
		Val c.Uint32T
	}
	Core0AreaDram00Min c.Uint32T
	Core0AreaDram00Max c.Uint32T
	Core0AreaDram01Min c.Uint32T
	Core0AreaDram01Max c.Uint32T
	Core0AreaPif0Min   c.Uint32T
	Core0AreaPif0Max   c.Uint32T
	Core0AreaPif1Min   c.Uint32T
	Core0AreaPif1Max   c.Uint32T
	Core0AreaSp        c.Uint32T
	Core0AreaPc        c.Uint32T
	Core0SpUnstable    struct {
		Val c.Uint32T
	}
	Core0SpMin           c.Uint32T
	Core0SpMax           c.Uint32T
	Core0SpPc            c.Uint32T
	Core0RcdPdebugenable struct {
		Val c.Uint32T
	}
	Core0RcdRecording struct {
		Val c.Uint32T
	}
	Core0RcdPdebuginst   c.Uint32T
	Core0RcdPdebugstatus struct {
		Val c.Uint32T
	}
	Core0RcdPdebugdata          c.Uint32T
	Core0RcdPdebugpc            c.Uint32T
	Core0RcdPdebugls0stat       c.Uint32T
	Core0RcdPdebugls0addr       c.Uint32T
	Core0RcdPdebugls0data       c.Uint32T
	Core0RcdSp                  c.Uint32T
	Core0Iram0ExceptionMonitor0 struct {
		Val c.Uint32T
	}
	Core0Iram0ExceptionMonitor1 struct {
		Val c.Uint32T
	}
	Core0Dram0ExceptionMonitor0 struct {
		Val c.Uint32T
	}
	Core0Dram0ExceptionMonitor1 struct {
		Val c.Uint32T
	}
	Core0Dram0ExceptionMonitor2 c.Uint32T
	Core0Dram0ExceptionMonitor3 struct {
		Val c.Uint32T
	}
	Core0Dram0ExceptionMonitor4 struct {
		Val c.Uint32T
	}
	Core0Dram0ExceptionMonitor5 c.Uint32T
	Core1InterruptEna           struct {
		Val c.Uint32T
	}
	Core1InterruptRaw struct {
		Val c.Uint32T
	}
	Core1InterruptRls struct {
		Val c.Uint32T
	}
	Core1InterruptClr struct {
		Val c.Uint32T
	}
	Core1AreaDram00Min c.Uint32T
	Core1AreaDram00Max c.Uint32T
	Core1AreaDram01Min c.Uint32T
	Core1AreaDram01Max c.Uint32T
	Core1AreaPif0Min   c.Uint32T
	Core1AreaPif0Max   c.Uint32T
	Core1AreaPif1Min   c.Uint32T
	Core1AreaPif1Max   c.Uint32T
	Core1AreaPc        c.Uint32T
	Core1AreaSp        c.Uint32T
	Core1SpUnstable    struct {
		Val c.Uint32T
	}
	Core1SpMin           c.Uint32T
	Core1SpMax           c.Uint32T
	Core1SpPc            c.Uint32T
	Core1RcdPdebugenable struct {
		Val c.Uint32T
	}
	Core1RcdRecording struct {
		Val c.Uint32T
	}
	Core1RcdPdebuginst   c.Uint32T
	Core1RcdPdebugstatus struct {
		Val c.Uint32T
	}
	Core1RcdPdebugdata          c.Uint32T
	Core1RcdPdebugpc            c.Uint32T
	Core1RcdPdebugls0stat       c.Uint32T
	Core1RcdPdebugls0addr       c.Uint32T
	Core1RcdPdebugls0data       c.Uint32T
	Core1RcdSp                  c.Uint32T
	Core1Iram0ExceptionMonitor0 struct {
		Val c.Uint32T
	}
	Core1Iram0ExceptionMonitor1 struct {
		Val c.Uint32T
	}
	Core1Dram0ExceptionMonitor0 struct {
		Val c.Uint32T
	}
	Core1Dram0ExceptionMonitor1 struct {
		Val c.Uint32T
	}
	Core1Dram0ExceptionMonitor2 c.Uint32T
	Core1Dram0ExceptionMonitor3 struct {
		Val c.Uint32T
	}
	Core1Dram0ExceptionMonitor4 struct {
		Val c.Uint32T
	}
	Core1Dram0ExceptionMonitor5      c.Uint32T
	CoreXIram0Dram0ExceptionMonitor0 struct {
		Val c.Uint32T
	}
	CoreXIram0Dram0ExceptionMonitor1 struct {
		Val c.Uint32T
	}
	LogSetting struct {
		Val c.Uint32T
	}
	LogData0    c.Uint32T
	LogData1    c.Uint32T
	LogData2    c.Uint32T
	LogData3    c.Uint32T
	LogDataMask struct {
		Val c.Uint32T
	}
	LogMin            c.Uint32T
	LogMax            c.Uint32T
	LogMemStart       c.Uint32T
	LogMemEnd         c.Uint32T
	LogMemWritingAddr c.Uint32T
	LogMemFullFlag    struct {
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
	RegDate     struct {
		Val c.Uint32T
	}
}
type AssistDebugDevT AssistDebugDevS
