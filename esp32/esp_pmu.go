package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PmuHpIcgModemModeT c.Int

const (
	PMU_HP_ICG_MODEM_CODE_SLEEP  PmuHpIcgModemModeT = 0
	PMU_HP_ICG_MODEM_CODE_MODEM  PmuHpIcgModemModeT = 1
	PMU_HP_ICG_MODEM_CODE_ACTIVE PmuHpIcgModemModeT = 2
)
