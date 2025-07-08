package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type InterruptCore0DevS struct {
	Core0MacIntrMap struct {
		Val c.Uint32T
	}
	Core0MacNmiMap struct {
		Val c.Uint32T
	}
	Core0PwrIntrMap struct {
		Val c.Uint32T
	}
	Core0BbIntMap struct {
		Val c.Uint32T
	}
	Core0BtMacIntMap struct {
		Val c.Uint32T
	}
	Core0BtBbIntMap struct {
		Val c.Uint32T
	}
	Core0BtBbNmiMap struct {
		Val c.Uint32T
	}
	Core0RwbtIrqMap struct {
		Val c.Uint32T
	}
	Core0RwbleIrqMap struct {
		Val c.Uint32T
	}
	Core0RwbtNmiMap struct {
		Val c.Uint32T
	}
	Core0RwbleNmiMap struct {
		Val c.Uint32T
	}
	Core0I2cMstIntMap struct {
		Val c.Uint32T
	}
	Core0Slc0IntrMap struct {
		Val c.Uint32T
	}
	Core0Slc1IntrMap struct {
		Val c.Uint32T
	}
	Core0Uhci0IntrMap struct {
		Val c.Uint32T
	}
	Core0Uhci1IntrMap struct {
		Val c.Uint32T
	}
	Core0GpioInterruptProMap struct {
		Val c.Uint32T
	}
	Core0GpioInterruptProNmiMap struct {
		Val c.Uint32T
	}
	Core0GpioInterruptAppMap struct {
		Val c.Uint32T
	}
	Core0GpioInterruptAppNmiMap struct {
		Val c.Uint32T
	}
	Core0SpiIntr1Map struct {
		Val c.Uint32T
	}
	Core0SpiIntr2Map struct {
		Val c.Uint32T
	}
	Core0SpiIntr3Map struct {
		Val c.Uint32T
	}
	Core0SpiIntr4Map struct {
		Val c.Uint32T
	}
	Core0LcdCamIntMap struct {
		Val c.Uint32T
	}
	Core0I2s0IntMap struct {
		Val c.Uint32T
	}
	Core0I2s1IntMap struct {
		Val c.Uint32T
	}
	Core0UartIntrMap struct {
		Val c.Uint32T
	}
	Core0Uart1IntrMap struct {
		Val c.Uint32T
	}
	Core0Uart2IntrMap struct {
		Val c.Uint32T
	}
	Core0SdioHostInterruptMap struct {
		Val c.Uint32T
	}
	Core0Pwm0IntrMap struct {
		Val c.Uint32T
	}
	Core0Pwm1IntrMap struct {
		Val c.Uint32T
	}
	Core0Pwm2IntrMap struct {
		Val c.Uint32T
	}
	Core0Pwm3IntrMap struct {
		Val c.Uint32T
	}
	Core0LedcIntMap struct {
		Val c.Uint32T
	}
	Core0EfuseIntMap struct {
		Val c.Uint32T
	}
	Core0CanIntMap struct {
		Val c.Uint32T
	}
	Core0UsbIntrMap struct {
		Val c.Uint32T
	}
	Core0RtcCoreIntrMap struct {
		Val c.Uint32T
	}
	Core0RmtIntrMap struct {
		Val c.Uint32T
	}
	Core0PcntIntrMap struct {
		Val c.Uint32T
	}
	Core0I2cExt0IntrMap struct {
		Val c.Uint32T
	}
	Core0I2cExt1IntrMap struct {
		Val c.Uint32T
	}
	Core0Spi2DmaIntMap struct {
		Val c.Uint32T
	}
	Core0Spi3DmaIntMap struct {
		Val c.Uint32T
	}
	Core0Spi4DmaIntMap struct {
		Val c.Uint32T
	}
	Core0WdgIntMap struct {
		Val c.Uint32T
	}
	Core0TimerInt1Map struct {
		Val c.Uint32T
	}
	Core0TimerInt2Map struct {
		Val c.Uint32T
	}
	Core0TgT0IntMap struct {
		Val c.Uint32T
	}
	Core0TgT1IntMap struct {
		Val c.Uint32T
	}
	Core0TgWdtIntMap struct {
		Val c.Uint32T
	}
	Core0Tg1T0IntMap struct {
		Val c.Uint32T
	}
	Core0Tg1T1IntMap struct {
		Val c.Uint32T
	}
	Core0Tg1WdtIntMap struct {
		Val c.Uint32T
	}
	Core0CacheIaIntMap struct {
		Val c.Uint32T
	}
	Core0SystimerTarget0IntMap struct {
		Val c.Uint32T
	}
	Core0SystimerTarget1IntMap struct {
		Val c.Uint32T
	}
	Core0SystimerTarget2IntMap struct {
		Val c.Uint32T
	}
	Core0SpiMemRejectIntrMap struct {
		Val c.Uint32T
	}
	Core0DcachePreloadIntMap struct {
		Val c.Uint32T
	}
	Core0IcachePreloadIntMap struct {
		Val c.Uint32T
	}
	Core0DcacheSyncIntMap struct {
		Val c.Uint32T
	}
	Core0IcacheSyncIntMap struct {
		Val c.Uint32T
	}
	Core0ApbAdcIntMap struct {
		Val c.Uint32T
	}
	Core0DmaInCh0IntMap struct {
		Val c.Uint32T
	}
	Core0DmaInCh1IntMap struct {
		Val c.Uint32T
	}
	Core0DmaInCh2IntMap struct {
		Val c.Uint32T
	}
	Core0DmaInCh3IntMap struct {
		Val c.Uint32T
	}
	Core0DmaInCh4IntMap struct {
		Val c.Uint32T
	}
	Core0DmaOutCh0IntMap struct {
		Val c.Uint32T
	}
	Core0DmaOutCh1IntMap struct {
		Val c.Uint32T
	}
	Core0DmaOutCh2IntMap struct {
		Val c.Uint32T
	}
	Core0DmaOutCh3IntMap struct {
		Val c.Uint32T
	}
	Core0DmaOutCh4IntMap struct {
		Val c.Uint32T
	}
	Core0RsaIntMap struct {
		Val c.Uint32T
	}
	Core0AesIntMap struct {
		Val c.Uint32T
	}
	Core0ShaIntMap struct {
		Val c.Uint32T
	}
	Core0CpuIntrFromCpu0Map struct {
		Val c.Uint32T
	}
	Core0CpuIntrFromCpu1Map struct {
		Val c.Uint32T
	}
	Core0CpuIntrFromCpu2Map struct {
		Val c.Uint32T
	}
	Core0CpuIntrFromCpu3Map struct {
		Val c.Uint32T
	}
	Core0AssistDebugIntrMap struct {
		Val c.Uint32T
	}
	Core0DmaApbperiPmsMonitorViolateIntrMap struct {
		Val c.Uint32T
	}
	Core0Core0Iram0PmsMonitorViolateIntrMap struct {
		Val c.Uint32T
	}
	Core0Core0Dram0PmsMonitorViolateIntrMap struct {
		Val c.Uint32T
	}
	Core0Core0PifPmsMonitorViolateIntrMap struct {
		Val c.Uint32T
	}
	Core0Core0PifPmsMonitorViolateSizeIntrMap struct {
		Val c.Uint32T
	}
	Core0Core1Iram0PmsMonitorViolateIntrMap struct {
		Val c.Uint32T
	}
	Core0Core1Dram0PmsMonitorViolateIntrMap struct {
		Val c.Uint32T
	}
	Core0Core1PifPmsMonitorViolateIntrMap struct {
		Val c.Uint32T
	}
	Core0Core1PifPmsMonitorViolateSizeIntrMap struct {
		Val c.Uint32T
	}
	Core0BackupPmsViolateIntrMap struct {
		Val c.Uint32T
	}
	Core0CacheCore0AcsIntMap struct {
		Val c.Uint32T
	}
	Core0CacheCore1AcsIntMap struct {
		Val c.Uint32T
	}
	Core0UsbDeviceIntMap struct {
		Val c.Uint32T
	}
	Core0PeriBackupIntMap struct {
		Val c.Uint32T
	}
	Core0DmaExtmemRejectIntMap struct {
		Val c.Uint32T
	}
	Core0IntrStatus0 c.Uint32T
	Core0IntrStatus1 c.Uint32T
	Core0IntrStatus2 c.Uint32T
	Core0IntrStatus3 c.Uint32T
	Core0ClockGate   struct {
		Val c.Uint32T
	}
	Reserved1a0        c.Uint32T
	Reserved1a4        c.Uint32T
	Reserved1a8        c.Uint32T
	Reserved1ac        c.Uint32T
	Reserved1b0        c.Uint32T
	Reserved1b4        c.Uint32T
	Reserved1b8        c.Uint32T
	Reserved1bc        c.Uint32T
	Reserved1c0        c.Uint32T
	Reserved1c4        c.Uint32T
	Reserved1c8        c.Uint32T
	Reserved1cc        c.Uint32T
	Reserved1d0        c.Uint32T
	Reserved1d4        c.Uint32T
	Reserved1d8        c.Uint32T
	Reserved1dc        c.Uint32T
	Reserved1e0        c.Uint32T
	Reserved1e4        c.Uint32T
	Reserved1e8        c.Uint32T
	Reserved1ec        c.Uint32T
	Reserved1f0        c.Uint32T
	Reserved1f4        c.Uint32T
	Reserved1f8        c.Uint32T
	Reserved1fc        c.Uint32T
	Reserved200        c.Uint32T
	Reserved204        c.Uint32T
	Reserved208        c.Uint32T
	Reserved20c        c.Uint32T
	Reserved210        c.Uint32T
	Reserved214        c.Uint32T
	Reserved218        c.Uint32T
	Reserved21c        c.Uint32T
	Reserved220        c.Uint32T
	Reserved224        c.Uint32T
	Reserved228        c.Uint32T
	Reserved22c        c.Uint32T
	Reserved230        c.Uint32T
	Reserved234        c.Uint32T
	Reserved238        c.Uint32T
	Reserved23c        c.Uint32T
	Reserved240        c.Uint32T
	Reserved244        c.Uint32T
	Reserved248        c.Uint32T
	Reserved24c        c.Uint32T
	Reserved250        c.Uint32T
	Reserved254        c.Uint32T
	Reserved258        c.Uint32T
	Reserved25c        c.Uint32T
	Reserved260        c.Uint32T
	Reserved264        c.Uint32T
	Reserved268        c.Uint32T
	Reserved26c        c.Uint32T
	Reserved270        c.Uint32T
	Reserved274        c.Uint32T
	Reserved278        c.Uint32T
	Reserved27c        c.Uint32T
	Reserved280        c.Uint32T
	Reserved284        c.Uint32T
	Reserved288        c.Uint32T
	Reserved28c        c.Uint32T
	Reserved290        c.Uint32T
	Reserved294        c.Uint32T
	Reserved298        c.Uint32T
	Reserved29c        c.Uint32T
	Reserved2a0        c.Uint32T
	Reserved2a4        c.Uint32T
	Reserved2a8        c.Uint32T
	Reserved2ac        c.Uint32T
	Reserved2b0        c.Uint32T
	Reserved2b4        c.Uint32T
	Reserved2b8        c.Uint32T
	Reserved2bc        c.Uint32T
	Reserved2c0        c.Uint32T
	Reserved2c4        c.Uint32T
	Reserved2c8        c.Uint32T
	Reserved2cc        c.Uint32T
	Reserved2d0        c.Uint32T
	Reserved2d4        c.Uint32T
	Reserved2d8        c.Uint32T
	Reserved2dc        c.Uint32T
	Reserved2e0        c.Uint32T
	Reserved2e4        c.Uint32T
	Reserved2e8        c.Uint32T
	Reserved2ec        c.Uint32T
	Reserved2f0        c.Uint32T
	Reserved2f4        c.Uint32T
	Reserved2f8        c.Uint32T
	Reserved2fc        c.Uint32T
	Reserved300        c.Uint32T
	Reserved304        c.Uint32T
	Reserved308        c.Uint32T
	Reserved30c        c.Uint32T
	Reserved310        c.Uint32T
	Reserved314        c.Uint32T
	Reserved318        c.Uint32T
	Reserved31c        c.Uint32T
	Reserved320        c.Uint32T
	Reserved324        c.Uint32T
	Reserved328        c.Uint32T
	Reserved32c        c.Uint32T
	Reserved330        c.Uint32T
	Reserved334        c.Uint32T
	Reserved338        c.Uint32T
	Reserved33c        c.Uint32T
	Reserved340        c.Uint32T
	Reserved344        c.Uint32T
	Reserved348        c.Uint32T
	Reserved34c        c.Uint32T
	Reserved350        c.Uint32T
	Reserved354        c.Uint32T
	Reserved358        c.Uint32T
	Reserved35c        c.Uint32T
	Reserved360        c.Uint32T
	Reserved364        c.Uint32T
	Reserved368        c.Uint32T
	Reserved36c        c.Uint32T
	Reserved370        c.Uint32T
	Reserved374        c.Uint32T
	Reserved378        c.Uint32T
	Reserved37c        c.Uint32T
	Reserved380        c.Uint32T
	Reserved384        c.Uint32T
	Reserved388        c.Uint32T
	Reserved38c        c.Uint32T
	Reserved390        c.Uint32T
	Reserved394        c.Uint32T
	Reserved398        c.Uint32T
	Reserved39c        c.Uint32T
	Reserved3a0        c.Uint32T
	Reserved3a4        c.Uint32T
	Reserved3a8        c.Uint32T
	Reserved3ac        c.Uint32T
	Reserved3b0        c.Uint32T
	Reserved3b4        c.Uint32T
	Reserved3b8        c.Uint32T
	Reserved3bc        c.Uint32T
	Reserved3c0        c.Uint32T
	Reserved3c4        c.Uint32T
	Reserved3c8        c.Uint32T
	Reserved3cc        c.Uint32T
	Reserved3d0        c.Uint32T
	Reserved3d4        c.Uint32T
	Reserved3d8        c.Uint32T
	Reserved3dc        c.Uint32T
	Reserved3e0        c.Uint32T
	Reserved3e4        c.Uint32T
	Reserved3e8        c.Uint32T
	Reserved3ec        c.Uint32T
	Reserved3f0        c.Uint32T
	Reserved3f4        c.Uint32T
	Reserved3f8        c.Uint32T
	Reserved3fc        c.Uint32T
	Reserved400        c.Uint32T
	Reserved404        c.Uint32T
	Reserved408        c.Uint32T
	Reserved40c        c.Uint32T
	Reserved410        c.Uint32T
	Reserved414        c.Uint32T
	Reserved418        c.Uint32T
	Reserved41c        c.Uint32T
	Reserved420        c.Uint32T
	Reserved424        c.Uint32T
	Reserved428        c.Uint32T
	Reserved42c        c.Uint32T
	Reserved430        c.Uint32T
	Reserved434        c.Uint32T
	Reserved438        c.Uint32T
	Reserved43c        c.Uint32T
	Reserved440        c.Uint32T
	Reserved444        c.Uint32T
	Reserved448        c.Uint32T
	Reserved44c        c.Uint32T
	Reserved450        c.Uint32T
	Reserved454        c.Uint32T
	Reserved458        c.Uint32T
	Reserved45c        c.Uint32T
	Reserved460        c.Uint32T
	Reserved464        c.Uint32T
	Reserved468        c.Uint32T
	Reserved46c        c.Uint32T
	Reserved470        c.Uint32T
	Reserved474        c.Uint32T
	Reserved478        c.Uint32T
	Reserved47c        c.Uint32T
	Reserved480        c.Uint32T
	Reserved484        c.Uint32T
	Reserved488        c.Uint32T
	Reserved48c        c.Uint32T
	Reserved490        c.Uint32T
	Reserved494        c.Uint32T
	Reserved498        c.Uint32T
	Reserved49c        c.Uint32T
	Reserved4a0        c.Uint32T
	Reserved4a4        c.Uint32T
	Reserved4a8        c.Uint32T
	Reserved4ac        c.Uint32T
	Reserved4b0        c.Uint32T
	Reserved4b4        c.Uint32T
	Reserved4b8        c.Uint32T
	Reserved4bc        c.Uint32T
	Reserved4c0        c.Uint32T
	Reserved4c4        c.Uint32T
	Reserved4c8        c.Uint32T
	Reserved4cc        c.Uint32T
	Reserved4d0        c.Uint32T
	Reserved4d4        c.Uint32T
	Reserved4d8        c.Uint32T
	Reserved4dc        c.Uint32T
	Reserved4e0        c.Uint32T
	Reserved4e4        c.Uint32T
	Reserved4e8        c.Uint32T
	Reserved4ec        c.Uint32T
	Reserved4f0        c.Uint32T
	Reserved4f4        c.Uint32T
	Reserved4f8        c.Uint32T
	Reserved4fc        c.Uint32T
	Reserved500        c.Uint32T
	Reserved504        c.Uint32T
	Reserved508        c.Uint32T
	Reserved50c        c.Uint32T
	Reserved510        c.Uint32T
	Reserved514        c.Uint32T
	Reserved518        c.Uint32T
	Reserved51c        c.Uint32T
	Reserved520        c.Uint32T
	Reserved524        c.Uint32T
	Reserved528        c.Uint32T
	Reserved52c        c.Uint32T
	Reserved530        c.Uint32T
	Reserved534        c.Uint32T
	Reserved538        c.Uint32T
	Reserved53c        c.Uint32T
	Reserved540        c.Uint32T
	Reserved544        c.Uint32T
	Reserved548        c.Uint32T
	Reserved54c        c.Uint32T
	Reserved550        c.Uint32T
	Reserved554        c.Uint32T
	Reserved558        c.Uint32T
	Reserved55c        c.Uint32T
	Reserved560        c.Uint32T
	Reserved564        c.Uint32T
	Reserved568        c.Uint32T
	Reserved56c        c.Uint32T
	Reserved570        c.Uint32T
	Reserved574        c.Uint32T
	Reserved578        c.Uint32T
	Reserved57c        c.Uint32T
	Reserved580        c.Uint32T
	Reserved584        c.Uint32T
	Reserved588        c.Uint32T
	Reserved58c        c.Uint32T
	Reserved590        c.Uint32T
	Reserved594        c.Uint32T
	Reserved598        c.Uint32T
	Reserved59c        c.Uint32T
	Reserved5a0        c.Uint32T
	Reserved5a4        c.Uint32T
	Reserved5a8        c.Uint32T
	Reserved5ac        c.Uint32T
	Reserved5b0        c.Uint32T
	Reserved5b4        c.Uint32T
	Reserved5b8        c.Uint32T
	Reserved5bc        c.Uint32T
	Reserved5c0        c.Uint32T
	Reserved5c4        c.Uint32T
	Reserved5c8        c.Uint32T
	Reserved5cc        c.Uint32T
	Reserved5d0        c.Uint32T
	Reserved5d4        c.Uint32T
	Reserved5d8        c.Uint32T
	Reserved5dc        c.Uint32T
	Reserved5e0        c.Uint32T
	Reserved5e4        c.Uint32T
	Reserved5e8        c.Uint32T
	Reserved5ec        c.Uint32T
	Reserved5f0        c.Uint32T
	Reserved5f4        c.Uint32T
	Reserved5f8        c.Uint32T
	Reserved5fc        c.Uint32T
	Reserved600        c.Uint32T
	Reserved604        c.Uint32T
	Reserved608        c.Uint32T
	Reserved60c        c.Uint32T
	Reserved610        c.Uint32T
	Reserved614        c.Uint32T
	Reserved618        c.Uint32T
	Reserved61c        c.Uint32T
	Reserved620        c.Uint32T
	Reserved624        c.Uint32T
	Reserved628        c.Uint32T
	Reserved62c        c.Uint32T
	Reserved630        c.Uint32T
	Reserved634        c.Uint32T
	Reserved638        c.Uint32T
	Reserved63c        c.Uint32T
	Reserved640        c.Uint32T
	Reserved644        c.Uint32T
	Reserved648        c.Uint32T
	Reserved64c        c.Uint32T
	Reserved650        c.Uint32T
	Reserved654        c.Uint32T
	Reserved658        c.Uint32T
	Reserved65c        c.Uint32T
	Reserved660        c.Uint32T
	Reserved664        c.Uint32T
	Reserved668        c.Uint32T
	Reserved66c        c.Uint32T
	Reserved670        c.Uint32T
	Reserved674        c.Uint32T
	Reserved678        c.Uint32T
	Reserved67c        c.Uint32T
	Reserved680        c.Uint32T
	Reserved684        c.Uint32T
	Reserved688        c.Uint32T
	Reserved68c        c.Uint32T
	Reserved690        c.Uint32T
	Reserved694        c.Uint32T
	Reserved698        c.Uint32T
	Reserved69c        c.Uint32T
	Reserved6a0        c.Uint32T
	Reserved6a4        c.Uint32T
	Reserved6a8        c.Uint32T
	Reserved6ac        c.Uint32T
	Reserved6b0        c.Uint32T
	Reserved6b4        c.Uint32T
	Reserved6b8        c.Uint32T
	Reserved6bc        c.Uint32T
	Reserved6c0        c.Uint32T
	Reserved6c4        c.Uint32T
	Reserved6c8        c.Uint32T
	Reserved6cc        c.Uint32T
	Reserved6d0        c.Uint32T
	Reserved6d4        c.Uint32T
	Reserved6d8        c.Uint32T
	Reserved6dc        c.Uint32T
	Reserved6e0        c.Uint32T
	Reserved6e4        c.Uint32T
	Reserved6e8        c.Uint32T
	Reserved6ec        c.Uint32T
	Reserved6f0        c.Uint32T
	Reserved6f4        c.Uint32T
	Reserved6f8        c.Uint32T
	Reserved6fc        c.Uint32T
	Reserved700        c.Uint32T
	Reserved704        c.Uint32T
	Reserved708        c.Uint32T
	Reserved70c        c.Uint32T
	Reserved710        c.Uint32T
	Reserved714        c.Uint32T
	Reserved718        c.Uint32T
	Reserved71c        c.Uint32T
	Reserved720        c.Uint32T
	Reserved724        c.Uint32T
	Reserved728        c.Uint32T
	Reserved72c        c.Uint32T
	Reserved730        c.Uint32T
	Reserved734        c.Uint32T
	Reserved738        c.Uint32T
	Reserved73c        c.Uint32T
	Reserved740        c.Uint32T
	Reserved744        c.Uint32T
	Reserved748        c.Uint32T
	Reserved74c        c.Uint32T
	Reserved750        c.Uint32T
	Reserved754        c.Uint32T
	Reserved758        c.Uint32T
	Reserved75c        c.Uint32T
	Reserved760        c.Uint32T
	Reserved764        c.Uint32T
	Reserved768        c.Uint32T
	Reserved76c        c.Uint32T
	Reserved770        c.Uint32T
	Reserved774        c.Uint32T
	Reserved778        c.Uint32T
	Reserved77c        c.Uint32T
	Reserved780        c.Uint32T
	Reserved784        c.Uint32T
	Reserved788        c.Uint32T
	Reserved78c        c.Uint32T
	Reserved790        c.Uint32T
	Reserved794        c.Uint32T
	Reserved798        c.Uint32T
	Reserved79c        c.Uint32T
	Reserved7a0        c.Uint32T
	Reserved7a4        c.Uint32T
	Reserved7a8        c.Uint32T
	Reserved7ac        c.Uint32T
	Reserved7b0        c.Uint32T
	Reserved7b4        c.Uint32T
	Reserved7b8        c.Uint32T
	Reserved7bc        c.Uint32T
	Reserved7c0        c.Uint32T
	Reserved7c4        c.Uint32T
	Reserved7c8        c.Uint32T
	Reserved7cc        c.Uint32T
	Reserved7d0        c.Uint32T
	Reserved7d4        c.Uint32T
	Reserved7d8        c.Uint32T
	Reserved7dc        c.Uint32T
	Reserved7e0        c.Uint32T
	Reserved7e4        c.Uint32T
	Reserved7e8        c.Uint32T
	Reserved7ec        c.Uint32T
	Reserved7f0        c.Uint32T
	Reserved7f4        c.Uint32T
	Reserved7f8        c.Uint32T
	Core0InterruptDate struct {
		Val c.Uint32T
	}
}
type InterruptCore0DevT InterruptCore0DevS
