package freertos

import _ "unsafe"

/**
 * @brief Switch CPU clock source to XTAL, and let cpu frequency equal to main XTAL frequency.
 *
 * This function does not disable CPU's source PLL. If the PLL requires to be disabled to save power, please call
 * `rtc_clk_cpu_freq_set_xtal` instead. It will always disable the corresponding PLL after switching the CPU clock
 * source to XTAL (except for S2).
 *
 * Currently, this function is only called in `esp_restart_noos` and `esp_restart_noos_dig` to switch the CPU
 * clock source back to XTAL (by default) before reset.
 */
//go:linkname RtcClkCpuSetToDefaultConfig C.rtc_clk_cpu_set_to_default_config
func RtcClkCpuSetToDefaultConfig()

/**
 * @brief Switch CPU clock source to XTAL, the PLL has different processing methods for different chips.
 *        1. For earlier chips without PMU, there is no PMU module that can turn off the CPU's PLL, so it has to be
 *           disabled at here to save the power consumption. Though ESP32C3/S3 has USB CDC device, it can not function
 *           properly during sleep due to the lack of APB clock (before C6, USJ relies on APB clock to work). Therefore,
 *           we will always disable CPU's PLL (i.e. BBPLL).
 *        2. For PMU supported chips, CPU's PLL power can be turned off by PMU, so no need to disable the PLL at here.
 *           Leaving PLL on at this stage also helps USJ keep connection and retention operation (if they rely on this PLL).
 *           For ESP32P4, if the APB frequency is configured as the hardware default value (10MHz), this will cause the
 *           regdma backup/restore to not achieve optimal performance. The MEM/APB frequency divider needs to be configured
 *           to 40MHz to speed up the retention speed.
 */
//go:linkname RtcClkCpuFreqSetXtalForSleep C.rtc_clk_cpu_freq_set_xtal_for_sleep
func RtcClkCpuFreqSetXtalForSleep()

/**
 * @brief Notify that the BBPLL has a new in-use consumer
 *
 * Currently, this function is only used for tracking whether USB Serial/JTAG is using the 48MHz PHY clock
 *
 * Note: Calling this function only helps to not disable the BBPLL clock in `rtc_clk_cpu_freq_set_config`.
 */
//go:linkname RtcClkBbpllAddConsumer C.rtc_clk_bbpll_add_consumer
func RtcClkBbpllAddConsumer()

/**
 * @brief Notify that the BBPLL has lost a consumer
 */
//go:linkname RtcClkBbpllRemoveConsumer C.rtc_clk_bbpll_remove_consumer
func RtcClkBbpllRemoveConsumer()
