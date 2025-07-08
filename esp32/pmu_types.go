package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PmuHpModeT c.Int

const (
	PMU_MODE_HP_ACTIVE PmuHpModeT = 0
	PMU_MODE_HP_MODEM  PmuHpModeT = 1
	PMU_MODE_HP_SLEEP  PmuHpModeT = 2
	PMU_MODE_HP_MAX    PmuHpModeT = 3
)

type PmuLpModeT c.Int

const (
	PMU_MODE_LP_ACTIVE PmuLpModeT = 0
	PMU_MODE_LP_SLEEP  PmuLpModeT = 1
	PMU_MODE_LP_MAX    PmuLpModeT = 2
)

type PmuHpPowerDomainT c.Int

const (
	PMU_HP_PD_TOP      PmuHpPowerDomainT = 0
	PMU_HP_PD_CPU      PmuHpPowerDomainT = 2
	PMU_HP_PD_RESERVED PmuHpPowerDomainT = 3
	PMU_HP_PD_WIFI     PmuHpPowerDomainT = 4
)
