package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LcdSocHandleT *LcdCamDevT

/**
 * @brief LCD HAL layer context
 */

type LcdHalContextT struct {
	Dev LcdSocHandleT
}

/**
 * @brief LCD HAL layer initialization
 *
 * @param hal LCD HAL layer context
 * @param id LCD peripheral ID
 */
// llgo:link (*LcdHalContextT).LcdHalInit C.lcd_hal_init
func (recv_ *LcdHalContextT) LcdHalInit(id c.Int) {
}

/**
 * @brief LCD PCLK clock calculation
 * @note Currently this function is only used by RGB LCD driver, I80 driver still uses a fixed clock division
 *
 * @param hal LCD HAL layer context
 * @param src_freq_hz LCD source clock frequency in Hz
 * @param expect_pclk_freq_hz Expected LCD PCLK frequency in Hz
 * @param lcd_clk_div Returned LCD clock divider parameter
 * @return Actual LCD PCLK frequency in Hz
 */
// llgo:link (*LcdHalContextT).LcdHalCalPclkFreq C.lcd_hal_cal_pclk_freq
func (recv_ *LcdHalContextT) LcdHalCalPclkFreq(src_freq_hz c.Uint32T, expect_pclk_freq_hz c.Uint32T, lcd_clk_div *HalUtilsClkDivT) c.Uint32T {
	return 0
}
