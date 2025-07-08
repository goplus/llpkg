package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief This function is used to enable the digital RC_FAST clock,
 *        to support the peripherals.
 *
 * @note If this function is called a number of times, the `periph_rtc_dig_clk8m_disable`
 *       function needs to be called same times to disable.
 *
 * @return true: success for enable the RC_FAST clock, false: RC_FAST clock enable failed
 */
//go:linkname PeriphRtcDigClk8mEnable C.periph_rtc_dig_clk8m_enable
func PeriphRtcDigClk8mEnable() bool

/**
 * @brief This function is used to disable the digital RC_FAST clock, which should be called
 *        with the `periph_rtc_dig_clk8m_enable` pairedly
 *
 * @note If this function is called a number of times, the `periph_rtc_dig_clk8m_disable`
 *       function needs to be called same times to disable.
 */
//go:linkname PeriphRtcDigClk8mDisable C.periph_rtc_dig_clk8m_disable
func PeriphRtcDigClk8mDisable()

/**
 * @brief This function is used to get the real clock frequency value of RC_FAST clock
 *
 * @return The real clock value, in Hz
 */
//go:linkname PeriphRtcDigClk8mGetFreq C.periph_rtc_dig_clk8m_get_freq
func PeriphRtcDigClk8mGetFreq() c.Uint32T
