package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioPortT c.Int

const (
	GPIO_PORT_0   GpioPortT = 0
	GPIO_PORT_MAX GpioPortT = 1
)

type GpioIntTypeT c.Int

const (
	GPIO_INTR_DISABLE    GpioIntTypeT = 0
	GPIO_INTR_POSEDGE    GpioIntTypeT = 1
	GPIO_INTR_NEGEDGE    GpioIntTypeT = 2
	GPIO_INTR_ANYEDGE    GpioIntTypeT = 3
	GPIO_INTR_LOW_LEVEL  GpioIntTypeT = 4
	GPIO_INTR_HIGH_LEVEL GpioIntTypeT = 5
	GPIO_INTR_MAX        GpioIntTypeT = 6
)

type GpioModeT c.Int

const (
	GPIO_MODE_DISABLE         GpioModeT = 0
	GPIO_MODE_INPUT           GpioModeT = 1
	GPIO_MODE_OUTPUT          GpioModeT = 2
	GPIO_MODE_OUTPUT_OD       GpioModeT = 6
	GPIO_MODE_INPUT_OUTPUT_OD GpioModeT = 7
	GPIO_MODE_INPUT_OUTPUT    GpioModeT = 3
)

type GpioPullupT c.Int

const (
	GPIO_PULLUP_DISABLE GpioPullupT = 0
	GPIO_PULLUP_ENABLE  GpioPullupT = 1
)

type GpioPulldownT c.Int

const (
	GPIO_PULLDOWN_DISABLE GpioPulldownT = 0
	GPIO_PULLDOWN_ENABLE  GpioPulldownT = 1
)

type GpioPullModeT c.Int

const (
	GPIO_PULLUP_ONLY     GpioPullModeT = 0
	GPIO_PULLDOWN_ONLY   GpioPullModeT = 1
	GPIO_PULLUP_PULLDOWN GpioPullModeT = 2
	GPIO_FLOATING        GpioPullModeT = 3
)

type GpioDriveCapT c.Int

const (
	GPIO_DRIVE_CAP_0       GpioDriveCapT = 0
	GPIO_DRIVE_CAP_1       GpioDriveCapT = 1
	GPIO_DRIVE_CAP_2       GpioDriveCapT = 2
	GPIO_DRIVE_CAP_DEFAULT GpioDriveCapT = 2
	GPIO_DRIVE_CAP_3       GpioDriveCapT = 3
	GPIO_DRIVE_CAP_MAX     GpioDriveCapT = 4
)
