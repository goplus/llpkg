package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GptimerSocHandleT *TimgDevT

/**
 * Context that should be maintained by both the driver and the HAL
 */

type TimerHalContextT struct {
	Dev     GptimerSocHandleT
	TimerId c.Uint32T
}

/**
 * @brief Init the timer hal. This function should be called first before other hal layer function is called
 *
 * @param hal Context of the HAL layer
 * @param group_num The timer group number
 * @param timer_num The timer number
 */
// llgo:link (*TimerHalContextT).TimerHalInit C.timer_hal_init
func (recv_ *TimerHalContextT) TimerHalInit(group_num c.Uint32T, timer_num c.Uint32T) {
}

/**
 * @brief Deinit timer hal context.
 *
 * @param hal Context of HAL layer
 */
// llgo:link (*TimerHalContextT).TimerHalDeinit C.timer_hal_deinit
func (recv_ *TimerHalContextT) TimerHalDeinit() {
}

/**
 * @brief Load counter value into time-base counter
 *
 * @param hal Context of the HAL layer
 * @param load_val Counter value
 */
// llgo:link (*TimerHalContextT).TimerHalSetCounterValue C.timer_hal_set_counter_value
func (recv_ *TimerHalContextT) TimerHalSetCounterValue(load_val c.Uint64T) {
}

/**
 * @brief Trigger a software capture event and then return the captured count value
 *
 * @param hal Context of the HAL layer
 * @return Counter value
 */
// llgo:link (*TimerHalContextT).TimerHalCaptureAndGetCounterValue C.timer_hal_capture_and_get_counter_value
func (recv_ *TimerHalContextT) TimerHalCaptureAndGetCounterValue() c.Uint64T {
	return 0
}
