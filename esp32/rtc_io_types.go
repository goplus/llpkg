package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RtcGpioModeT c.Int

const (
	RTC_GPIO_MODE_INPUT_ONLY      RtcGpioModeT = 0
	RTC_GPIO_MODE_OUTPUT_ONLY     RtcGpioModeT = 1
	RTC_GPIO_MODE_INPUT_OUTPUT    RtcGpioModeT = 2
	RTC_GPIO_MODE_DISABLED        RtcGpioModeT = 3
	RTC_GPIO_MODE_OUTPUT_OD       RtcGpioModeT = 4
	RTC_GPIO_MODE_INPUT_OUTPUT_OD RtcGpioModeT = 5
)
