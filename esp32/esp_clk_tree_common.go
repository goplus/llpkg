package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get frequency of RC_FAST_D256_CLK
 *
 * @param precision Degree of precision of the returned frequency value, one of esp_clk_tree_src_freq_precision_t values
 *
 * @return RC_FAST_D256 clock frequency, in Hz. Returns 0 if degree of precision is invalid or calibration failed.
 */
// llgo:link EspClkTreeSrcFreqPrecisionT.EspClkTreeRcFastD256GetFreqHz C.esp_clk_tree_rc_fast_d256_get_freq_hz
func (recv_ EspClkTreeSrcFreqPrecisionT) EspClkTreeRcFastD256GetFreqHz() c.Uint32T {
	return 0
}

/**
 * @brief Get frequency of XTAL32K_CLK
 *
 * @param precision Degree of precision of the returned frequency value, one of esp_clk_tree_src_freq_precision_t values
 *
 * @return XTAL32K clock frequency, in Hz. Returns 0 if degree of precision is invalid or calibration failed.
 */
// llgo:link EspClkTreeSrcFreqPrecisionT.EspClkTreeXtal32kGetFreqHz C.esp_clk_tree_xtal32k_get_freq_hz
func (recv_ EspClkTreeSrcFreqPrecisionT) EspClkTreeXtal32kGetFreqHz() c.Uint32T {
	return 0
}

/**
 * @brief Get frequency of RC_FAST_CLK
 *
 * @param precision Degree of precision of the returned frequency value, one of esp_clk_tree_src_freq_precision_t values
 *
 * @return RC_FAST clock frequency, in Hz. Returns 0 if degree of precision is invalid or calibration failed.
 */
// llgo:link EspClkTreeSrcFreqPrecisionT.EspClkTreeRcFastGetFreqHz C.esp_clk_tree_rc_fast_get_freq_hz
func (recv_ EspClkTreeSrcFreqPrecisionT) EspClkTreeRcFastGetFreqHz() c.Uint32T {
	return 0
}

/**
 * @brief Get frequency of LP_SLOW_CLK (i.e. RTC_SLOW_CLK)
 *
 * @param precision Degree of precision of the returned frequency value, one of esp_clk_tree_src_freq_precision_t values
 *
 * @return LP_SLOW clock frequency, in Hz. Returns 0 if degree of precision is invalid or calibration failed.
 */
// llgo:link EspClkTreeSrcFreqPrecisionT.EspClkTreeLpSlowGetFreqHz C.esp_clk_tree_lp_slow_get_freq_hz
func (recv_ EspClkTreeSrcFreqPrecisionT) EspClkTreeLpSlowGetFreqHz() c.Uint32T {
	return 0
}

/**
 * @brief Get frequency of LP_FAST_CLK (i.e. RTC_FAST_CLK)
 *
 * @param precision Degree of precision of the returned frequency value, one of esp_clk_tree_src_freq_precision_t values
 *
 * @return LP_FAST clock frequency, in Hz. Returns 0 if degree of precision is invalid or calibration failed.
 */
// llgo:link EspClkTreeSrcFreqPrecisionT.EspClkTreeLpFastGetFreqHz C.esp_clk_tree_lp_fast_get_freq_hz
func (recv_ EspClkTreeSrcFreqPrecisionT) EspClkTreeLpFastGetFreqHz() c.Uint32T {
	return 0
}

/**
 * @brief Enable / Disable the clock gate of the clock source
 *
 * @param[in] clk_src Clock source available to modules, in soc_module_clk_t
 * @param[in] enable  Enable / Disable the clock gate
 *
 * @note !!! WARNING !!!
 *       There's no reference counter to protect the clock source status, the caller should use the interface
 *       with CAUTION to disable the clock source to avoid damaging other peripherals that are dependent on
 *       the clock source.
 *
 * @return
 *      - ESP_OK               Success
 *      - ESP_ERR_INVALID_ARG  Parameter error
 */
// llgo:link SocModuleClkT.EspClkTreeEnableSrc C.esp_clk_tree_enable_src
func (recv_ SocModuleClkT) EspClkTreeEnableSrc(enable bool) EspErrT {
	return 0
}
