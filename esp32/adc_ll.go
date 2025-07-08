package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ADC_LL_CLKM_DIV_NUM_DEFAULT = 15
const ADC_LL_CLKM_DIV_B_DEFAULT = 1
const ADC_LL_CLKM_DIV_A_DEFAULT = 0
const ADC_LL_DEFAULT_CONV_LIMIT_EN = 0
const ADC_LL_DEFAULT_CONV_LIMIT_NUM = 10

type AdcLlPowerT c.Int

const (
	ADC_LL_POWER_BY_FSM AdcLlPowerT = 0
	ADC_LL_POWER_SW_ON  AdcLlPowerT = 1
	ADC_LL_POWER_SW_OFF AdcLlPowerT = 2
)

type AdcLlRtcRawDataT c.Int

const (
	ADC_LL_RTC_DATA_OK         AdcLlRtcRawDataT = 0
	ADC_LL_RTC_CTRL_UNSELECTED AdcLlRtcRawDataT = 1
	ADC_LL_RTC_CTRL_BREAK      AdcLlRtcRawDataT = 2
	ADC_LL_RTC_DATA_FAIL       AdcLlRtcRawDataT = -1
)

type AdcLlControllerT c.Int

const (
	ADC_LL_CTRL_RTC AdcLlControllerT = 0
	ADC_LL_CTRL_ULP AdcLlControllerT = 1
	ADC_LL_CTRL_DIG AdcLlControllerT = 2
	ADC_LL_CTRL_ARB AdcLlControllerT = 3
)

type AdcLlDigiConvertModeT c.Int

const (
	ADC_LL_DIGI_CONV_ONLY_ADC1  AdcLlDigiConvertModeT = 0
	ADC_LL_DIGI_CONV_ONLY_ADC2  AdcLlDigiConvertModeT = 1
	ADC_LL_DIGI_CONV_BOTH_UNIT  AdcLlDigiConvertModeT = 2
	ADC_LL_DIGI_CONV_ALTER_UNIT AdcLlDigiConvertModeT = 3
)

type AdcLlDigiPatternTableT struct {
	Unused [8]uint8
}

/**
 * @brief Analyze whether the obtained raw data is correct.
 *        ADC2 use arbiter by default. The arbitration result can be judged by the flag bit in the original data.
 *
 */

type AdcLlRtcOutputDataT struct {
	Unused [8]uint8
}
