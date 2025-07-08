package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief BF configurations
 */

type IspHalBfCfgT struct {
	PaddingMode                    IspBfEdgePaddingModeT
	PaddingData                    c.Uint8T
	BfTemplate                     [0][0]c.Uint8T
	DenoisingLevel                 c.Uint8T
	PaddingLineTailValidStartPixel c.Uint8T
	PaddingLineTailValidEndPixel   c.Uint8T
}

/**
 * @brief Demosaic configurations
 */

type IspHalDemosaicCfgT struct {
	GradRatio                      IspDemosaicGradRatioT
	PaddingMode                    IspDemosaicEdgePaddingModeT
	PaddingData                    c.Uint8T
	PaddingLineTailValidStartPixel c.Uint8T
	PaddingLineTailValidEndPixel   c.Uint8T
}

/**
 * @brief Sharpen configurations
 */

type IspHalSharpenCfgT struct {
	HFreqCoeff                     IspSharpenHFreqCoeffT
	MFreqCoeff                     IspSharpenMFreqCoeff
	HThresh                        c.Uint8T
	LThresh                        c.Uint8T
	PaddingMode                    IspSharpenEdgePaddingModeT
	PaddingData                    c.Uint8T
	SharpenTemplate                [0][0]c.Uint8T
	PaddingLineTailValidStartPixel c.Uint8T
	PaddingLineTailValidEndPixel   c.Uint8T
}

/**
 * @brief Context that should be maintained by both the driver and the HAL
 */

type IspHalContextT struct {
	Hw    c.Pointer
	BfCfg IspHalBfCfgT
}

/**
 * @brief Init the ISP hal and set the ISP to the default configuration.
 *
 * @note This function should be called first before other hal layer function is called.
 *
 * @param[in] hal     Context of the HAL layer
 * @param[in] isp_id  ISP ID
 */
// llgo:link (*IspHalContextT).IspHalInit C.isp_hal_init
func (recv_ *IspHalContextT) IspHalInit(isp_id c.Int) {
}

/**
 * @brief Color configurations
 */

type IspHalColorCfgT struct {
	ColorContrast   IspColorContrastT
	ColorSaturation IspColorSaturationT
	ColorHue        c.Uint32T
	ColorBrightness c.Int
}

/*---------------------------------------------------------------
                      AF
---------------------------------------------------------------*/
/**
 * @brief Configure AF window
 *
 * @param[in] hal        Context of the HAL layer
 * @param[in] window_id  Window ID
 * @param[in] window     Window info, see `isp_window_t`
 */
// llgo:link (*IspHalContextT).IspHalAfWindowConfig C.isp_hal_af_window_config
func (recv_ *IspHalContextT) IspHalAfWindowConfig(window_id c.Int, window *IspWindowT) {
}

/*---------------------------------------------------------------
                      AE
---------------------------------------------------------------*/
/**
 * @brief Configure AE window
 *
 * @param[in] hal        Context of the HAL layer
 * @param[in] window     Window info, see `isp_window_t`
 */
// llgo:link (*IspHalContextT).IspHalAeWindowConfig C.isp_hal_ae_window_config
func (recv_ *IspHalContextT) IspHalAeWindowConfig(window *IspWindowT) {
}

/*---------------------------------------------------------------
                            AWB
---------------------------------------------------------------*/
/**
 * @brief  Set the window of the AWB
 *
 * @param[in] hal   Context of the HAL layer
 * @param[in] win   Pointer to the window of the AWB
 * @return
 *      - true      Set success
 *      - false     Invalid arg
 */
// llgo:link (*IspHalContextT).IspHalAwbSetWindowRange C.isp_hal_awb_set_window_range
func (recv_ *IspHalContextT) IspHalAwbSetWindowRange(win *IspWindowT) bool {
	return false
}

/**
 * @brief   Set the luminance range of the white patch
 *
 * @param[in] hal   Context of the HAL layer
 * @param[in] lum_min Minimum luminance
 * @param[in] lum_max Maximum luminance
 * @return
 *      - true      Set success
 *      - false     Invalid arg
 */
// llgo:link (*IspHalContextT).IspHalAwbSetLuminanceRange C.isp_hal_awb_set_luminance_range
func (recv_ *IspHalContextT) IspHalAwbSetLuminanceRange(lum_min c.Uint32T, lum_max c.Uint32T) bool {
	return false
}

/**
 * @brief   Set the R/G ratio of the white patch
 *
 * @param[in] hal   Context of the HAL layer
 * @param[in] rg_ratio_range Range of Red to Green ratio
 * @return
 *      - true      Set success
 *      - false     Invalid arg
 */
// llgo:link (*IspHalContextT).IspHalAwbSetRgRatioRange C.isp_hal_awb_set_rg_ratio_range
func (recv_ *IspHalContextT) IspHalAwbSetRgRatioRange(rg_min c.Float, rg_max c.Float) bool {
	return false
}

/**
 * @brief   Set the B/R ratio of the white patch
 *
 * @param[in] hal   Context of the HAL layer
 * @param[in] bg_ratio_range Range of Blue to Green ratio
 * @return
 *      - true      Set success
 *      - false     Invalid arg
 */
// llgo:link (*IspHalContextT).IspHalAwbSetBgRatioRange C.isp_hal_awb_set_bg_ratio_range
func (recv_ *IspHalContextT) IspHalAwbSetBgRatioRange(bg_min c.Float, bg_max c.Float) bool {
	return false
}

/*---------------------------------------------------------------
                      BF
---------------------------------------------------------------*/
/**
 * @brief Configure ISP BF
 *
 * @param[in] hal        Context of the HAL layer
 * @param[in] config     BF config, set NULL to de-config the ISP BF
 */
// llgo:link (*IspHalContextT).IspHalBfConfig C.isp_hal_bf_config
func (recv_ *IspHalContextT) IspHalBfConfig(config *IspHalBfCfgT) {
}

/*---------------------------------------------------------------
                      Color Correction Matrix
---------------------------------------------------------------*/
/**
 * @brief Set Color Correction Matrix
 *
 * @param[in] hal           Context of the HAL layer
 * @param[in] saturation    Whether to enable saturation when float data overflow
 * @param[in] flt_matrix    3x3 RGB correction matrix
 * @return
 *      - true      Set success
 *      - false     Invalid argument
 */
// llgo:link (*IspHalContextT).IspHalCcmSetMatrix C.isp_hal_ccm_set_matrix
func (recv_ *IspHalContextT) IspHalCcmSetMatrix(saturation bool, flt_matrix **c.Float) bool {
	return false
}

/*---------------------------------------------------------------
                      Demosaic
---------------------------------------------------------------*/
/**
 * @brief Configure ISP Demosaic
 *
 * @param[in] hal        Context of the HAL layer
 * @param[in] config     Demosaic config, set NULL to de-config the ISP Demosaic
 */
// llgo:link (*IspHalContextT).IspHalDemosaicConfig C.isp_hal_demosaic_config
func (recv_ *IspHalContextT) IspHalDemosaicConfig(config *IspHalDemosaicCfgT) {
}

/*---------------------------------------------------------------
                      INTR
---------------------------------------------------------------*/
/**
 * @brief Clear ISP HW intr event
 *
 * @param[in] hal   Context of the HAL layer
 * @param[in] mask  HW event mask
 */
// llgo:link (*IspHalContextT).IspHalCheckClearIntrEvent C.isp_hal_check_clear_intr_event
func (recv_ *IspHalContextT) IspHalCheckClearIntrEvent(mask c.Uint32T) c.Uint32T {
	return 0
}

/*---------------------------------------------------------------
                      Sharpness
---------------------------------------------------------------*/
/**
 * @brief Configure ISP sharpeness
 *
 * @param[in] hal        Context of the HAL layer
 * @param[in] config     Sharpness config, set NULL to de-config the ISP sharpness
 */
// llgo:link (*IspHalContextT).IspHalSharpenConfig C.isp_hal_sharpen_config
func (recv_ *IspHalContextT) IspHalSharpenConfig(config *IspHalSharpenCfgT) {
}

/*---------------------------------------------------------------
                      Histogram
---------------------------------------------------------------*/
/**
 * @brief Configure Histogram window
 *
 * @param[in] hal        Context of the HAL layer
 * @param[in] window     Window info, see `isp_window_t`
 */
// llgo:link (*IspHalContextT).IspHalHistWindowConfig C.isp_hal_hist_window_config
func (recv_ *IspHalContextT) IspHalHistWindowConfig(window *IspWindowT) {
}

/**
 * @brief   Set the color config
 *
 * @param[in] hal                       Context of the HAL layer
 * @param[in] config                    Color config, set NULL to de-config the ISP color
 */
// llgo:link (*IspHalContextT).IspHalColorConfig C.isp_hal_color_config
func (recv_ *IspHalContextT) IspHalColorConfig(config *IspHalColorCfgT) {
}
