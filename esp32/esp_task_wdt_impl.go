package freertos

import _ "unsafe"

/**
 * @brief   Allocate and initialize the Task Watchdog Timer (TWDT) with the given configuration.
 *
 * @param[in] config Pointer to the configuration structure
 * @param[out] obj Abstract context for the current timer, this will be passed to all the other functions
 *
 * @return
 *  - ESP_OK: Successfully initialized and configured the timer
 *  - Other: Failed to initialize the timer
 */
// llgo:link (*EspTaskWdtConfigT).EspTaskWdtImplTimerAllocate C.esp_task_wdt_impl_timer_allocate
func (recv_ *EspTaskWdtConfigT) EspTaskWdtImplTimerAllocate(callback TwdtIsrCallback, obj *TwdtCtxT) EspErrT {
	return 0
}

/**
 * @brief   Reconfigure a timer.
 *
 * The timer must be stopped when calling this function. The timer will not be restarted at the end of this
 * function.
 *
 * @param[in] config Pointer to the configuration structure
 *
 * @return
 *  - ESP_OK: Successfully reconfigured the timer
 *  - Other: Failed to reconfigure the timer
 */
//go:linkname EspTaskWdtImplTimerReconfigure C.esp_task_wdt_impl_timer_reconfigure
func EspTaskWdtImplTimerReconfigure(obj TwdtCtxT, config *EspTaskWdtConfigT) EspErrT

/**
 * @brief   Free the Task Watchdog Timer (TWDT).
 *
 * @param[in] obj Abstract implementation context
 *
 */
//go:linkname EspTaskWdtImplTimerFree C.esp_task_wdt_impl_timer_free
func EspTaskWdtImplTimerFree(obj TwdtCtxT)

/**
 * @brief   Feed the Task Watchdog Timer (TWDT)
 *
 * Feed the timer underneath to prevent it from triggering for the next period (configured at initialization).
 *
 * @param[in] obj Abstract implementation context
 * @return
 *  - ESP_OK: timer successfully feeded
 *  - Other: failed to feed the timer
 */
//go:linkname EspTaskWdtImplTimerFeed C.esp_task_wdt_impl_timer_feed
func EspTaskWdtImplTimerFeed(obj TwdtCtxT) EspErrT

/**
 * @brief   Function invoked as soon as the Task Watchdog Timer (TWDT) ISR callback is called.
 *
 * @param[in] obj Abstract implementation context
 */
//go:linkname EspTaskWdtImplTimeoutTriggered C.esp_task_wdt_impl_timeout_triggered
func EspTaskWdtImplTimeoutTriggered(obj TwdtCtxT)

/**
 * @brief   Stop the Task Watchdog Timer (TWDT).
 *
 * @param[in] obj Abstract implementation context
 *
 */
//go:linkname EspTaskWdtImplTimerStop C.esp_task_wdt_impl_timer_stop
func EspTaskWdtImplTimerStop(obj TwdtCtxT) EspErrT

/**
 * @brief   Restart the Task Watchdog Timer (TWDT)
 *
 * This function will restart/resume the timer after it has been stopped.
 *
 * @param[in] obj Abstract implementation context
 * @return
 *  - ESP_OK: timer successfully stopped
 *  - Other: failed to stop the timer
 */
//go:linkname EspTaskWdtImplTimerRestart C.esp_task_wdt_impl_timer_restart
func EspTaskWdtImplTimerRestart(obj TwdtCtxT) EspErrT
