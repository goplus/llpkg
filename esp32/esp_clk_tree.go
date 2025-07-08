package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspClkTreeSrcFreqPrecisionT c.Int

const (
	ESP_CLK_TREE_SRC_FREQ_PRECISION_CACHED  EspClkTreeSrcFreqPrecisionT = 0
	ESP_CLK_TREE_SRC_FREQ_PRECISION_APPROX  EspClkTreeSrcFreqPrecisionT = 1
	ESP_CLK_TREE_SRC_FREQ_PRECISION_EXACT   EspClkTreeSrcFreqPrecisionT = 2
	ESP_CLK_TREE_SRC_FREQ_PRECISION_INVALID EspClkTreeSrcFreqPrecisionT = 3
)

/**
 * @brief Get frequency of module clock source
 *
 * @param[in] clk_src Clock source available to modules, in soc_module_clk_t
 * @param[in] precision Degree of precision, one of esp_clk_tree_src_freq_precision_t values
 *                      This arg only applies to the clock sources that their frequencies can vary:
 *                      SOC_MOD_CLK_RTC_FAST, SOC_MOD_CLK_RTC_SLOW, SOC_MOD_CLK_RC_FAST, SOC_MOD_CLK_RC_FAST_D256,
 *                      SOC_MOD_CLK_XTAL32K
 *                      For other clock sources, this field is ignored.
 * @param[out] freq_value Frequency of the clock source, in Hz
 *
 * @return
 *      - ESP_OK               Success
 *      - ESP_ERR_INVALID_ARG  Parameter error
 *      - ESP_FAIL             Calibration failed
 */
// llgo:link SocModuleClkT.EspClkTreeSrcGetFreqHz C.esp_clk_tree_src_get_freq_hz
func (recv_ SocModuleClkT) EspClkTreeSrcGetFreqHz(precision EspClkTreeSrcFreqPrecisionT, freq_value *c.Uint32T) EspErrT {
	return 0
}
