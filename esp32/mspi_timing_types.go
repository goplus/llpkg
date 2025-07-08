package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MspiTimingSpeedModeT c.Int

const (
	MSPI_TIMING_SPEED_MODE_LOW_PERF    MspiTimingSpeedModeT = 0
	MSPI_TIMING_SPEED_MODE_NORMAL_PERF MspiTimingSpeedModeT = 1
)
