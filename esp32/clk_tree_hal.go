package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get SOC_ROOT_CLK frequency
 *
 * @param cpu_clk_src One of the clock sources in soc_cpu_clk_src_t
 *
 * @return SOC ROOT clock frequency, in MHz. Returns 0 if input argument is invalid.
 */
// llgo:link SocCpuClkSrcT.ClkHalSocRootGetFreqMhz C.clk_hal_soc_root_get_freq_mhz
func (recv_ SocCpuClkSrcT) ClkHalSocRootGetFreqMhz() c.Uint32T {
	return 0
}

/**
 * @brief Get CPU_CLK frequency
 *
 * @return CPU clock frequency, in Hz. Returns 0 if internal clock configuration is invalid.
 */
//go:linkname ClkHalCpuGetFreqHz C.clk_hal_cpu_get_freq_hz
func ClkHalCpuGetFreqHz() c.Uint32T

/**
 * @brief Get APB_CLK frequency
 *
 * @return APB clock frequency, in Hz. Returns 0 if internal clock configuration is invalid.
 */
//go:linkname ClkHalApbGetFreqHz C.clk_hal_apb_get_freq_hz
func ClkHalApbGetFreqHz() c.Uint32T

/**
 * @brief Get LP_SLOW_CLK (i.e. RTC_SLOW_CLK) approximate frequency
 *
 * @return LP Slow clock frequency, in Hz. Returns 0 if LP_SLOW clock source is invalid.
 */
//go:linkname ClkHalLpSlowGetFreqHz C.clk_hal_lp_slow_get_freq_hz
func ClkHalLpSlowGetFreqHz() c.Uint32T

/**
 * @brief Get XTAL_CLK frequency
 *
 * @return XTAL clock frequency, in MHz
 */
//go:linkname ClkHalXtalGetFreqMhz C.clk_hal_xtal_get_freq_mhz
func ClkHalXtalGetFreqMhz() c.Uint32T

/**
 * @brief Set up clock output channel
 * @param clk_sig    The clock signal source to be mapped to GPIOs
 * @param channel_id The clock output channel to setup
 */
// llgo:link SocClkoutSigIdT.ClkHalClockOutputSetup C.clk_hal_clock_output_setup
func (recv_ SocClkoutSigIdT) ClkHalClockOutputSetup(channel_id ClockOutChannelT) {
}

/**
 * @brief Teardown clock output channel configuration
 * @param channel_id The clock output channel to teardown
 */
// llgo:link ClockOutChannelT.ClkHalClockOutputTeardown C.clk_hal_clock_output_teardown
func (recv_ ClockOutChannelT) ClkHalClockOutputTeardown() {
}
