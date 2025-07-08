package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const PWDET_CONF_REG = 0x6000E060

type SarCtrlLlPowerT c.Int

const (
	SAR_CTRL_LL_POWER_FSM SarCtrlLlPowerT = 0
	SAR_CTRL_LL_POWER_ON  SarCtrlLlPowerT = 1
	SAR_CTRL_LL_POWER_OFF SarCtrlLlPowerT = 2
)
