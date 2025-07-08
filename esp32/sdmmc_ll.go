package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SDMMC_LL_EVENT_DMA_MASK = 0x1f

type SdmmcLlDelayPhaseT c.Int

const (
	SDMMC_LL_DELAY_PHASE_0 SdmmcLlDelayPhaseT = 0
	SDMMC_LL_DELAY_PHASE_1 SdmmcLlDelayPhaseT = 1
	SDMMC_LL_DELAY_PHASE_2 SdmmcLlDelayPhaseT = 2
	SDMMC_LL_DELAY_PHASE_3 SdmmcLlDelayPhaseT = 3
)
