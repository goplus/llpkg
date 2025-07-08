package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ISP_AE_BLOCK_X_NUM = 0
const ISP_AE_BLOCK_Y_NUM = 0
const ISP_AF_WINDOW_NUM = 0
const ISP_BF_TEMPLATE_X_NUMS = 0
const ISP_BF_TEMPLATE_Y_NUMS = 0
const ISP_CCM_DIMENSION = 0
const ISP_DEMOSAIC_GRAD_RATIO_INT_BITS = 8
const ISP_DEMOSAIC_GRAD_RATIO_DEC_BITS = 8
const ISP_DEMOSAIC_GRAD_RATIO_RES_BITS = 16
const ISP_DVP_DATA_SIG_NUM = 0
const ISP_SHARPEN_TEMPLATE_X_NUMS = 0
const ISP_SHARPEN_TEMPLATE_Y_NUMS = 0
const ISP_SHARPEN_H_FREQ_COEF_INT_BITS = 8
const ISP_SHARPEN_H_FREQ_COEF_DEC_BITS = 8
const ISP_SHARPEN_H_FREQ_COEF_RES_BITS = 16
const ISP_SHARPEN_M_FREQ_COEF_INT_BITS = 8
const ISP_SHARPEN_M_FREQ_COEF_DEC_BITS = 8
const ISP_SHARPEN_M_FREQ_COEF_RES_BITS = 16
const ISP_GAMMA_CURVE_POINTS_NUM = 16
const ISP_HIST_BLOCK_X_NUM = 0
const ISP_HIST_BLOCK_Y_NUM = 0
const ISP_HIST_SEGMENT_NUMS = 0
const ISP_HIST_INTERVAL_NUMS = 0
const ISP_HIST_WEIGHT_INT_BITS = 8
const ISP_HIST_WEIGHT_DEC_BITS = 7
const ISP_HIST_WEIGHT_RES_BITS = 17
const ISP_HIST_COEFF_INT_BITS = 8
const ISP_HIST_COEFF_DEC_BITS = 7
const ISP_HIST_COEFF_RES_BITS = 17
const ISP_COLOR_CONTRAST_INT_BITS = 1
const ISP_COLOR_CONTRAST_DEC_BITS = 7
const ISP_COLOR_CONTRAST_RES_BITS = 24
const ISP_COLOR_SATURATION_INT_BITS = 1
const ISP_COLOR_SATURATION_DEC_BITS = 7
const ISP_COLOR_SATURATION_RES_BITS = 24
const ISP_LSC_GRAD_RATIO_INT_BITS = 2
const ISP_LSC_GRAD_RATIO_DEC_BITS = 8
const ISP_LSC_GRAD_RATIO_RES_BITS = 22

type IspClkSrcT c.Int

/**
 * @brief ISP coordinate type
 *
 */

type IspCoordinateT struct {
	X c.Uint32T
	Y c.Uint32T
}

/**
 * @brief  The top left and bottom right coordinates of ISP full window
 */

type IspWindowT struct {
	TopLeft  IspCoordinateT
	BtmRight IspCoordinateT
}
type IspInputDataSourceT c.Int

const (
	ISP_INPUT_DATA_SOURCE_CSI    IspInputDataSourceT = 0
	ISP_INPUT_DATA_SOURCE_DVP    IspInputDataSourceT = 1
	ISP_INPUT_DATA_SOURCE_DWGDMA IspInputDataSourceT = 2
)

type IspColorT c.Int

const (
	ISP_COLOR_RAW8   IspColorT = 16777216
	ISP_COLOR_RAW10  IspColorT = 16777217
	ISP_COLOR_RAW12  IspColorT = 16777218
	ISP_COLOR_RGB888 IspColorT = 33554432
	ISP_COLOR_RGB565 IspColorT = 33554434
	ISP_COLOR_YUV422 IspColorT = 50331649
	ISP_COLOR_YUV420 IspColorT = 50331650
)

type IspColorRangeT c.Int

const (
	ISP_COLOR_RANGE_LIMIT IspColorRangeT = 0
	ISP_COLOR_RANGE_FULL  IspColorRangeT = 1
)

type IspYuvConvStdT c.Int

const (
	ISP_YUV_CONV_STD_BT601 IspYuvConvStdT = 0
	ISP_YUV_CONV_STD_BT709 IspYuvConvStdT = 1
)

type IspAeSamplePointT c.Int

const (
	ISP_AE_SAMPLE_POINT_AFTER_DEMOSAIC IspAeSamplePointT = 0
	ISP_AE_SAMPLE_POINT_AFTER_GAMMA    IspAeSamplePointT = 1
)

type IspAwbSamplePointT c.Int

const (
	ISP_AWB_SAMPLE_POINT_BEFORE_CCM IspAwbSamplePointT = 0
	ISP_AWB_SAMPLE_POINT_AFTER_CCM  IspAwbSamplePointT = 1
)

type IspBfEdgePaddingModeT c.Int

const (
	ISP_BF_EDGE_PADDING_MODE_SRND_DATA   IspBfEdgePaddingModeT = 0
	ISP_BF_EDGE_PADDING_MODE_CUSTOM_DATA IspBfEdgePaddingModeT = 1
)

/**
 * @brief Gradient ratio
 */

type IspDemosaicGradRatioT struct {
	Val c.Uint32T
}
type IspDemosaicEdgePaddingModeT c.Int

const (
	ISP_DEMOSAIC_EDGE_PADDING_MODE_SRND_DATA   IspDemosaicEdgePaddingModeT = 0
	ISP_DEMOSAIC_EDGE_PADDING_MODE_CUSTOM_DATA IspDemosaicEdgePaddingModeT = 1
)

/**
 * @brief High freq pixel sharpeness coeff
 */

type IspSharpenHFreqCoeffT struct {
	Val c.Uint32T
}

/**
 * @brief Medium freq pixel sharpeness coeff
 */

type IspSharpenMFreqCoeff struct {
	Val c.Uint32T
}
type IspSharpenEdgePaddingModeT c.Int

const (
	ISP_SHARPEN_EDGE_PADDING_MODE_SRND_DATA   IspSharpenEdgePaddingModeT = 0
	ISP_SHARPEN_EDGE_PADDING_MODE_CUSTOM_DATA IspSharpenEdgePaddingModeT = 1
)

/**
 * @brief Structure that declares the points on an ISP gamma curve
 *
 * Constraint on pt[n].x:
 *  When n = 0, pt[n].x = 2 ^ a[n]
 *  When 0 < n < ISP_GAMMA_CURVE_POINTS_NUM-1, pt[n].x - pt[n-1].x = 2 ^ a[n]
 *  When n = ISP_GAMMA_CURVE_POINTS_NUM-1, pt[n].x = 255, (pt[n].x + 1) - pt[n-1].x = 2 ^ a[n]
 *  a[n] within [0, 7]
 */

type IspGammaCurvePointsT struct {
	Pt [16]struct {
		X c.Uint8T
		Y c.Uint8T
	}
}
type IspHistSamplingModeT c.Int

const (
	ISP_HIST_SAMPLING_RAW_B  IspHistSamplingModeT = 0
	ISP_HIST_SAMPLING_RAW_GB IspHistSamplingModeT = 1
	ISP_HIST_SAMPLING_RAW_GR IspHistSamplingModeT = 2
	ISP_HIST_SAMPLING_RAW_R  IspHistSamplingModeT = 3
	ISP_HIST_SAMPLING_RGB    IspHistSamplingModeT = 4
	ISP_HIST_SAMPLING_YUV_Y  IspHistSamplingModeT = 5
	ISP_HIST_SAMPLING_YUV_U  IspHistSamplingModeT = 6
	ISP_HIST_SAMPLING_YUV_V  IspHistSamplingModeT = 7
)

/**
 * @brief ISP histogram weight value
 */
type IspHistWeightT struct {
	Val c.Uint32T
}

/**
 * @brief ISP histogram coefficient value
 */
type IspHistCoeffT struct {
	Val c.Uint32T
}

/**
 * @brief ISP histogram r,g,b coefficient
 */
type IspHistRgbCoefficientT struct {
	CoeffR IspHistCoeffT
	CoeffG IspHistCoeffT
	CoeffB IspHistCoeffT
}

/**
 * @brief ISP histogram result.
 */
type IspHistResultT struct {
	HistValue [0]c.Uint32T
}

/**
 * @brief Color contrast value
 */

type IspColorContrastT struct {
	Val c.Uint32T
}

/**
 * @brief Color saturation value
 */

type IspColorSaturationT struct {
	Val c.Uint32T
}

/**
 * @brief LSC gain
 */

type IspLscGainT struct {
	Val c.Uint32T
}
