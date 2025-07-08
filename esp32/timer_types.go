package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GptimerClockSourceT SocPeriphGptimerClkSrcT
type GptimerCountDirectionT c.Int

const (
	GPTIMER_COUNT_DOWN GptimerCountDirectionT = 0
	GPTIMER_COUNT_UP   GptimerCountDirectionT = 1
)
