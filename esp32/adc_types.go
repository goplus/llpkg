package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcUnitT c.Int

const (
	ADC_UNIT_1 AdcUnitT = 0
	ADC_UNIT_2 AdcUnitT = 1
)

type AdcChannelT c.Int

const (
	ADC_CHANNEL_0 AdcChannelT = 0
	ADC_CHANNEL_1 AdcChannelT = 1
	ADC_CHANNEL_2 AdcChannelT = 2
	ADC_CHANNEL_3 AdcChannelT = 3
	ADC_CHANNEL_4 AdcChannelT = 4
	ADC_CHANNEL_5 AdcChannelT = 5
	ADC_CHANNEL_6 AdcChannelT = 6
	ADC_CHANNEL_7 AdcChannelT = 7
	ADC_CHANNEL_8 AdcChannelT = 8
	ADC_CHANNEL_9 AdcChannelT = 9
)

type AdcAttenT c.Int

const (
	ADC_ATTEN_DB_0   AdcAttenT = 0
	ADC_ATTEN_DB_2_5 AdcAttenT = 1
	ADC_ATTEN_DB_6   AdcAttenT = 2
	ADC_ATTEN_DB_12  AdcAttenT = 3
	ADC_ATTEN_DB_11  AdcAttenT = 3
)

type AdcBitwidthT c.Int

const (
	ADC_BITWIDTH_DEFAULT AdcBitwidthT = 0
	ADC_BITWIDTH_9       AdcBitwidthT = 9
	ADC_BITWIDTH_10      AdcBitwidthT = 10
	ADC_BITWIDTH_11      AdcBitwidthT = 11
	ADC_BITWIDTH_12      AdcBitwidthT = 12
	ADC_BITWIDTH_13      AdcBitwidthT = 13
)

type AdcUlpModeT c.Int

const (
	ADC_ULP_MODE_DISABLE AdcUlpModeT = 0
	ADC_ULP_MODE_FSM     AdcUlpModeT = 1
	ADC_ULP_MODE_RISCV   AdcUlpModeT = 2
)

type AdcDigiConvertModeT c.Int

const (
	ADC_CONV_SINGLE_UNIT_1 AdcDigiConvertModeT = 1
	ADC_CONV_SINGLE_UNIT_2 AdcDigiConvertModeT = 2
	ADC_CONV_BOTH_UNIT     AdcDigiConvertModeT = 3
	ADC_CONV_ALTER_UNIT    AdcDigiConvertModeT = 7
)

type AdcDigiOutputFormatT c.Int

const (
	ADC_DIGI_OUTPUT_FORMAT_TYPE1 AdcDigiOutputFormatT = 0
	ADC_DIGI_OUTPUT_FORMAT_TYPE2 AdcDigiOutputFormatT = 1
)

type AdcOneshotClkSrcT SocPeriphAdcRtcClkSrcT
type AdcContinuousClkSrcT SocPeriphAdcDigiClkSrcT

/**
 * @brief ADC digital controller pattern configuration
 */

type AdcDigiPatternConfigT struct {
	Atten    c.Uint8T
	Channel  c.Uint8T
	Unit     c.Uint8T
	BitWidth c.Uint8T
}
type AdcDigiIirFilterT c.Int

const (
	ADC_DIGI_IIR_FILTER_0 AdcDigiIirFilterT = 0
	ADC_DIGI_IIR_FILTER_1 AdcDigiIirFilterT = 1
)

type AdcDigiIirFilterCoeffT c.Int

const (
	ADC_DIGI_IIR_FILTER_COEFF_2  AdcDigiIirFilterCoeffT = 0
	ADC_DIGI_IIR_FILTER_COEFF_4  AdcDigiIirFilterCoeffT = 1
	ADC_DIGI_IIR_FILTER_COEFF_8  AdcDigiIirFilterCoeffT = 2
	ADC_DIGI_IIR_FILTER_COEFF_16 AdcDigiIirFilterCoeffT = 3
	ADC_DIGI_IIR_FILTER_COEFF_32 AdcDigiIirFilterCoeffT = 4
	ADC_DIGI_IIR_FILTER_COEFF_64 AdcDigiIirFilterCoeffT = 5
)

type AdcMonitorIdT c.Int

const (
	ADC_MONITOR_0 AdcMonitorIdT = 0
	ADC_MONITOR_1 AdcMonitorIdT = 1
)

type AdcMonitorModeT c.Int

const (
	ADC_MONITOR_MODE_HIGH AdcMonitorModeT = 0
	ADC_MONITOR_MODE_LOW  AdcMonitorModeT = 1
)

/**
 * @brief ADC digital controller (DMA mode) output data format.
 *        Used to analyze the acquired ADC (DMA) data.
 */

type AdcDigiOutputDataT struct {
	Unused [8]uint8
}
