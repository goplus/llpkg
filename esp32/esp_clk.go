package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get the calibration value of RTC slow clock
 *
 * The value is in the same format as returned by rtc_clk_cal (microseconds,
 * in Q13.19 fixed-point format).
 *
 * @return the calibration value obtained using rtc_clk_cal, at startup time
 */
//go:linkname EspClkSlowclkCalGet C.esp_clk_slowclk_cal_get
func EspClkSlowclkCalGet() c.Uint32T

/**
 * @brief Update the calibration value of RTC slow clock
 *
 * The value has to be in the same format as returned by rtc_clk_cal (microseconds,
 * in Q13.19 fixed-point format).
 * This value is used by timekeeping functions (such as gettimeofday) to
 * calculate current time based on RTC counter value.
 * @param value calibration value obtained using rtc_clk_cal
 */
//go:linkname EspClkSlowclkCalSet C.esp_clk_slowclk_cal_set
func EspClkSlowclkCalSet(value c.Uint32T)

/**
 * @brief Return current CPU clock frequency
 * When frequency switching is performed, this frequency may change.
 * However it is guaranteed that the frequency never changes with a critical
 * section.
 *
 * @return CPU clock frequency, in Hz
 */
//go:linkname EspClkCpuFreq C.esp_clk_cpu_freq
func EspClkCpuFreq() c.Int

/**
 * @brief Return current APB clock frequency
 *
 * When frequency switching is performed, this frequency may change.
 * However it is guaranteed that the frequency never changes with a critical
 * section.
 *
 * @return APB clock frequency, in Hz
 */
//go:linkname EspClkApbFreq C.esp_clk_apb_freq
func EspClkApbFreq() c.Int

/**
 * @brief Return frequency of the main XTAL
 *
 * Frequency of the main XTAL can be either auto-detected or set at compile
 * time (see CONFIG_XTAL_FREQ_SEL sdkconfig option). In both cases, this
 * function returns the actual value at run time.
 *
 * @return XTAL frequency, in Hz
 */
//go:linkname EspClkXtalFreq C.esp_clk_xtal_freq
func EspClkXtalFreq() c.Int

/**
 * @brief Read value of RTC counter, converting it to microseconds
 * @attention The value returned by this function may change abruptly when
 * calibration value of RTC counter is updated via esp_clk_slowclk_cal_set
 * function. This should not happen unless application calls esp_clk_slowclk_cal_set.
 * In ESP-IDF, esp_clk_slowclk_cal_set is only called in startup code.
 *
 * @return Value or RTC counter, expressed in microseconds
 */
//go:linkname EspClkRtcTime C.esp_clk_rtc_time
func EspClkRtcTime() c.Uint64T

/**
 * @brief obtain internal critical section used esp_clk implementation.
 *
 * This is used by the esp_light_sleep_start() to avoid deadlocking when it
 * calls esp_clk related API after stalling the other CPU.
 */
//go:linkname EspClkPrivateLock C.esp_clk_private_lock
func EspClkPrivateLock()

/**
 * @brief counterpart of esp_clk_private_lock
 */
//go:linkname EspClkPrivateUnlock C.esp_clk_private_unlock
func EspClkPrivateUnlock()
