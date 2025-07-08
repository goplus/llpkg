package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Common SDMMC slot info, doesn't depend on SOC_SDMMC_USE_{IOMUX,GPIO_MATRIX}
 */

type SdmmcSlotInfoT struct {
	Width        c.Uint8T
	CardDetect   c.Uint8T
	WriteProtect c.Uint8T
	CardInt      c.Uint8T
}

/**
 * This structure lists pin numbers (if SOC_SDMMC_USE_IOMUX is set)
 * or GPIO Matrix signal numbers (if SOC_SDMMC_USE_GPIO_MATRIX is set)
 * for the SD bus signals. Field names match SD bus signal names.
 */

type SdmmcSlotIoInfoT struct {
	Val [12]GpioNumT
}
