package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Enter deep-sleep mode from deep sleep wake stub code
 *
 * This should be called from the wake stub code.
 *
 * @param new_stub  new wake stub function will be set
 */
//go:linkname EspWakeStubSleep C.esp_wake_stub_sleep
func EspWakeStubSleep(new_stub EspDeepSleepWakeStubFnT)

/**
 * @brief Wait while uart transmission is in progress
 *
 * This function is waiting while uart transmission is not completed,
 * and this function should be called from the wake stub code.
 *
 * @param uart_no  UART port to wait idle
 */
//go:linkname EspWakeStubUartTxWaitIdle C.esp_wake_stub_uart_tx_wait_idle
func EspWakeStubUartTxWaitIdle(uart_no c.Uint8T)

/**
 * @brief Set wakeup time from deep sleep stub.
 *
 * This should be called from the wake stub code.
 *
 * @param time_in_us  wakeup time in us
 */
//go:linkname EspWakeStubSetWakeupTime C.esp_wake_stub_set_wakeup_time
func EspWakeStubSetWakeupTime(time_in_us c.Uint64T)

/**
 * @brief Get wakeup cause from deep sleep stub.
 *
 * This should be called from the wake stub code.
 *
 * @return wakeup casue value
 */
//go:linkname EspWakeStubGetWakeupCause C.esp_wake_stub_get_wakeup_cause
func EspWakeStubGetWakeupCause() c.Uint32T
