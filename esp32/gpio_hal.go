package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Context that should be maintained by both the driver and the HAL
 */

type GpioHalContextT struct {
	Dev *GpioDevT
}

/**
 * @brief  Enable GPIO module interrupt signal
 *
 * @param  hal Context of the HAL layer
 * @param  gpio_num GPIO number. If you want to enable the interrupt of e.g. GPIO16, gpio_num should be GPIO_NUM_16 (16);
 * @param  core_id Interrupt enabled CPU to corresponding ID
 */
// llgo:link (*GpioHalContextT).GpioHalIntrEnableOnCore C.gpio_hal_intr_enable_on_core
func (recv_ *GpioHalContextT) GpioHalIntrEnableOnCore(gpio_num c.Uint32T, core_id c.Uint32T) {
}

/**
 * @brief  Disable GPIO module interrupt signal
 *
 * @param  hal Context of the HAL layer
 * @param  gpio_num GPIO number. If you want to disable the interrupt of e.g. GPIO16, gpio_num should be GPIO_NUM_16 (16);
 */
// llgo:link (*GpioHalContextT).GpioHalIntrDisable C.gpio_hal_intr_disable
func (recv_ *GpioHalContextT) GpioHalIntrDisable(gpio_num c.Uint32T) {
}

/**
 * @brief Helper function to disconnect internal circuits from an GPIO in sleep mode.
 *        This function disables input, output, pullup, pulldown for an GPIO in sleep mode.
 *
 * @param hal Context of the HAL layer
 * @param gpio_num GPIO number
 */
// llgo:link (*GpioHalContextT).GpioHalIsolateInSleep C.gpio_hal_isolate_in_sleep
func (recv_ *GpioHalContextT) GpioHalIsolateInSleep(gpio_num c.Uint32T) {
}
