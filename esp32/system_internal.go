package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief  Internal function to restart PRO and APP CPUs.
 *
 * @note This function should not be called from FreeRTOS applications.
 *       Use esp_restart instead.
 *
 * This function executes a CPU reset (see TRM).
 *
 * CPU resets do not reset digital peripherals, but this function will
 * manually reset a subset of digital peripherals (depending on target) before
 * carrying out the CPU reset.
 *
 * Memory protection is also cleared by a CPU reset.
 *
 * This is an internal function called by esp_restart. It is called directly
 * by the panic handler and brownout detector interrupt.
 */
//go:linkname EspRestartNoos C.esp_restart_noos
func EspRestartNoos()

/**
 * @brief Similar to esp_restart_noos, but resets all the digital peripherals.
 */
//go:linkname EspRestartNoosDig C.esp_restart_noos_dig
func EspRestartNoosDig()

/**
 * @brief  Internal function to set reset reason hint
 *
 * The hint is used do distinguish different reset reasons when software reset
 * is performed.
 *
 * The hint is stored in RTC store register, RTC_RESET_CAUSE_REG.
 *
 * @param hint  Desired esp_reset_reason_t value for the real reset reason
 */
// llgo:link EspResetReasonT.EspResetReasonSetHint C.esp_reset_reason_set_hint
func (recv_ EspResetReasonT) EspResetReasonSetHint() {
}

/**
 * @brief  Internal function to get the reset hint value
 * @return  - Reset hint value previously stored into RTC_RESET_CAUSE_REG using
 *          esp_reset_reason_set_hint function
 *          - ESP_RST_UNKNOWN if the value in RTC_RESET_CAUSE_REG is invalid
 */
//go:linkname EspResetReasonGetHint C.esp_reset_reason_get_hint
func EspResetReasonGetHint() EspResetReasonT

/**
 * @brief Get the time in microseconds since startup
 *
 * @returns time since g_startup_time; definition should be fixed by system time provider
 * no matter the underlying timer used.
 */
//go:linkname EspSystemGetTime C.esp_system_get_time
func EspSystemGetTime() c.Int64T

/**
 * @brief Get the resolution of the time returned by `esp_system_get_time`.
 *
 * @returns the resolution in nanoseconds
 */
//go:linkname EspSystemGetTimeResolution C.esp_system_get_time_resolution
func EspSystemGetTimeResolution() c.Uint32T

/**
 * @brief Before the system exit (e.g. panic, brownout, restart, etc.), this function is to be called to reset all necessary peripherals.
 */
//go:linkname EspSystemResetModulesOnExit C.esp_system_reset_modules_on_exit
func EspSystemResetModulesOnExit()
