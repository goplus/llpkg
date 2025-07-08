package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MSPI_TIMING_LL_FLASH_SLOW_MODE_MASK = 0
const MSPI_TIMING_LL_CORE_CLOCK_MHZ_DEFAULT = 80
const MSPI_TIMING_LL_FLASH_CPU_CLK_SRC_BINDED = 1

type MspiTimingLlFlashModeT c.Int

const (
	MSPI_TIMING_LL_FLASH_OPI_MODE  MspiTimingLlFlashModeT = 1
	MSPI_TIMING_LL_FLASH_QIO_MODE  MspiTimingLlFlashModeT = 2
	MSPI_TIMING_LL_FLASH_QUAD_MODE MspiTimingLlFlashModeT = 4
	MSPI_TIMING_LL_FLASH_DIO_MODE  MspiTimingLlFlashModeT = 8
	MSPI_TIMING_LL_FLASH_DUAL_MODE MspiTimingLlFlashModeT = 16
	MSPI_TIMING_LL_FLASH_FAST_MODE MspiTimingLlFlashModeT = 32
	MSPI_TIMING_LL_FLASH_SLOW_MODE MspiTimingLlFlashModeT = 64
)
