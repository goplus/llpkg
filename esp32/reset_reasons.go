package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SocResetReasonT c.Int

const (
	RESET_REASON_CHIP_POWER_ON   SocResetReasonT = 1
	RESET_REASON_CHIP_BROWN_OUT  SocResetReasonT = 1
	RESET_REASON_CHIP_SUPER_WDT  SocResetReasonT = 1
	RESET_REASON_CORE_SW         SocResetReasonT = 3
	RESET_REASON_CORE_DEEP_SLEEP SocResetReasonT = 5
	RESET_REASON_CORE_MWDT0      SocResetReasonT = 7
	RESET_REASON_CORE_MWDT1      SocResetReasonT = 8
	RESET_REASON_CORE_RTC_WDT    SocResetReasonT = 9
	RESET_REASON_CPU0_MWDT0      SocResetReasonT = 11
	RESET_REASON_CPU1_MWDT0      SocResetReasonT = 11
	RESET_REASON_CPU0_SW         SocResetReasonT = 12
	RESET_REASON_CPU1_SW         SocResetReasonT = 12
	RESET_REASON_CPU0_RTC_WDT    SocResetReasonT = 13
	RESET_REASON_CPU1_RTC_WDT    SocResetReasonT = 13
	RESET_REASON_SYS_BROWN_OUT   SocResetReasonT = 15
	RESET_REASON_SYS_RTC_WDT     SocResetReasonT = 16
	RESET_REASON_CPU0_MWDT1      SocResetReasonT = 17
	RESET_REASON_CPU1_MWDT1      SocResetReasonT = 17
	RESET_REASON_SYS_SUPER_WDT   SocResetReasonT = 18
	RESET_REASON_SYS_CLK_GLITCH  SocResetReasonT = 19
	RESET_REASON_CORE_EFUSE_CRC  SocResetReasonT = 20
	RESET_REASON_CORE_USB_UART   SocResetReasonT = 21
	RESET_REASON_CORE_USB_JTAG   SocResetReasonT = 22
	RESET_REASON_CORE_PWR_GLITCH SocResetReasonT = 23
)
