package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PsramHalCmdModeT c.Int

const (
	PSRAM_HAL_CMD_QPI PsramHalCmdModeT = 0
	PSRAM_HAL_CMD_SPI PsramHalCmdModeT = 1
)
