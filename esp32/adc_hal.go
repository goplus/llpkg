package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcHalDmaDescStatusT c.Int

const (
	ADC_HAL_DMA_DESC_VALID   AdcHalDmaDescStatusT = 0
	ADC_HAL_DMA_DESC_WAITING AdcHalDmaDescStatusT = 1
	ADC_HAL_DMA_DESC_NULL    AdcHalDmaDescStatusT = 2
)

/**
 * @brief Configuration of the HAL
 */

type AdcHalDmaConfigT struct {
	EofDescNum c.Uint32T
	EofStep    c.Uint32T
	EofNum     c.Uint32T
}

/**
 * @brief Context of the HAL
 */

type AdcHalDmaCtxT struct {
	RxDesc        *DmaDescriptorT
	DescDummyHead DmaDescriptorT
	CurDescPtr    *DmaDescriptorT
	EofDescNum    c.Uint32T
	EofStep       c.Uint32T
	EofNum        c.Uint32T
}

type AdcHalDigiCtrlrCfgT struct {
	AdcPatternLen c.Uint32T
	AdcPattern    *AdcDigiPatternConfigT
	SampleFreqHz  c.Uint32T
	ConvMode      AdcDigiConvertModeT
	BitWidth      c.Uint32T
	ClkSrc        AdcContinuousClkSrcT
	ClkSrcFreqHz  c.Uint32T
}

/*---------------------------------------------------------------
                    Digital controller setting
---------------------------------------------------------------*/
/**
 * @brief Initialize the HW
 *
 * @param hal Context of the HAL
 */
// llgo:link (*AdcHalDmaCtxT).AdcHalDigiInit C.adc_hal_digi_init
func (recv_ *AdcHalDmaCtxT) AdcHalDigiInit() {
}

/**
 * Digital controller deinitialization.
 *
 */
//go:linkname AdcHalDigiDeinit C.adc_hal_digi_deinit
func AdcHalDigiDeinit()

/**
 * @brief Initialize the hal context
 *
 * @param hal    Context of the HAL
 * @param config Configuration of the HAL
 */
// llgo:link (*AdcHalDmaCtxT).AdcHalDmaCtxConfig C.adc_hal_dma_ctx_config
func (recv_ *AdcHalDmaCtxT) AdcHalDmaCtxConfig(config *AdcHalDmaConfigT) {
}

/**
 * Setting the digital controller.
 *
 * @param hal    Context of the HAL
 * @param cfg    Pointer to digital controller parameter.
 */
// llgo:link (*AdcHalDmaCtxT).AdcHalDigiControllerConfig C.adc_hal_digi_controller_config
func (recv_ *AdcHalDmaCtxT) AdcHalDigiControllerConfig(cfg *AdcHalDigiCtrlrCfgT) {
}

/**
 * @brief Link DMA descriptor
 *
 * @param hal Context of the HAL
 * @param data_buf Pointer to the data buffer, the length should be multiple of ``desc_max_num`` and ``eof_num`` in ``adc_hal_dma_ctx_t``
 */
// llgo:link (*AdcHalDmaCtxT).AdcHalDigiDmaLink C.adc_hal_digi_dma_link
func (recv_ *AdcHalDmaCtxT) AdcHalDigiDmaLink(data_buf *c.Uint8T) {
}

/**
 * @brief Get the ADC reading result
 *
 * @param      hal           Context of the HAL
 * @param      eof_desc_addr The last descriptor that is finished by HW. Should be got from DMA
 * @param[out] buffer        ADC reading result buffer
 * @param[out] len           ADC reading result len
 *
 * @return                   See ``adc_hal_dma_desc_status_t``
 */
// llgo:link (*AdcHalDmaCtxT).AdcHalGetReadingResult C.adc_hal_get_reading_result
func (recv_ *AdcHalDmaCtxT) AdcHalGetReadingResult(eof_desc_addr c.IntptrT, buffer **c.Uint8T, len *c.Uint32T) AdcHalDmaDescStatusT {
	return 0
}

/**
 * @brief Enable or disable ADC digital controller
 *
 * @param enable true to enable, false to disable
 */
//go:linkname AdcHalDigiEnable C.adc_hal_digi_enable
func AdcHalDigiEnable(enable bool)

/**
 * @brief Enable pr disable output data to DMA from adc digital controller.
 *
 * @param enable true to enable, false to disable
 */
//go:linkname AdcHalDigiConnect C.adc_hal_digi_connect
func AdcHalDigiConnect(enable bool)

/**
 * @brief Reset adc digital controller.
 */
//go:linkname AdcHalDigiReset C.adc_hal_digi_reset
func AdcHalDigiReset()

/**
 * @brief Set ADC monitor with high and low thresholds, and will enable the interrupts accordingly
 *
 * @param monitor_id Monitor to configure
 * @param adc_n Which ADC unit will be monitored
 * @param adc_ch Which ADC channel will be monitored
 * @param h_thres High threshold (disable if < 0)
 * @param l_thres Low threshold (disable if < 0)
 */
// llgo:link AdcMonitorIdT.AdcHalDigiMonitorSetThres C.adc_hal_digi_monitor_set_thres
func (recv_ AdcMonitorIdT) AdcHalDigiMonitorSetThres(adc_n AdcUnitT, adc_ch c.Uint8T, h_thres c.Int32T, l_thres c.Int32T) {
}
