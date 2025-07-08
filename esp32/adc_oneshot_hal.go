package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type Sens_Dev_________ struct {
	Unused [8]uint8
}
type AdcOneshotSocHandleT *Sens_Dev_________

type AdcOneshotHalCfgT struct {
	Unit         AdcUnitT
	WorkMode     AdcHalWorkModeT
	ClkSrc       AdcOneshotClkSrcT
	ClkSrcFreqHz c.Uint32T
}

/**
 * ADC channel configuration
 */

type AdcOneshotHalChanCfgT struct {
	Atten    AdcAttenT
	Bitwidth AdcBitwidthT
}

/**
 * Context of the ADC unit, should be maintained by both the driver and the HAL.
 */

type AdcOneshotHalCtxT struct {
	Dev          AdcOneshotSocHandleT
	Unit         AdcUnitT
	WorkMode     AdcHalWorkModeT
	ChanConfigs  [10]AdcOneshotHalChanCfgT
	ClkSrc       AdcOneshotClkSrcT
	ClkSrcFreqHz c.Uint32T
}

/**
 * Initialise the context
 *
 * @param hal    ADC Oneshot Hal context
 * @param config ADC Oneshot Hal init config
 */
// llgo:link (*AdcOneshotHalCtxT).AdcOneshotHalInit C.adc_oneshot_hal_init
func (recv_ *AdcOneshotHalCtxT) AdcOneshotHalInit(config *AdcOneshotHalCfgT) {
}

/**
 * Prepare ADC Oneshot hal context
 *
 * @param hal    ADC Oneshot Hal context
 * @param config ADC Oneshot Hal configuration
 * @param chan   ADC Channel
 */
// llgo:link (*AdcOneshotHalCtxT).AdcOneshotHalChannelConfig C.adc_oneshot_hal_channel_config
func (recv_ *AdcOneshotHalCtxT) AdcOneshotHalChannelConfig(config *AdcOneshotHalChanCfgT, chan_ AdcChannelT) {
}

/**
 * Set ADC Oneshot mode required registers
 *
 * @param hal     ADC Oneshot Hal context
 * @param channel ADC Channel
 */
// llgo:link (*AdcOneshotHalCtxT).AdcOneshotHalSetup C.adc_oneshot_hal_setup
func (recv_ *AdcOneshotHalCtxT) AdcOneshotHalSetup(channel AdcChannelT) {
}

/**
 * Start ADC conversion in Oneshot mode and get the raw result
 *
 * @param hal          ADC Oneshot Hal context
 * @param[out] out_raw ADC oneshot conversion raw result
 *
 * @return
 *        - true: ADC raw result is valid
 *        - false: ADC raw result is invalid
 */
// llgo:link (*AdcOneshotHalCtxT).AdcOneshotHalConvert C.adc_oneshot_hal_convert
func (recv_ *AdcOneshotHalCtxT) AdcOneshotHalConvert(out_raw *c.Int) bool {
	return false
}
