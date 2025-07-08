package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TouchPadT c.Int

const (
	TOUCH_PAD_NUM0  TouchPadT = 0
	TOUCH_PAD_NUM1  TouchPadT = 1
	TOUCH_PAD_NUM2  TouchPadT = 2
	TOUCH_PAD_NUM3  TouchPadT = 3
	TOUCH_PAD_NUM4  TouchPadT = 4
	TOUCH_PAD_NUM5  TouchPadT = 5
	TOUCH_PAD_NUM6  TouchPadT = 6
	TOUCH_PAD_NUM7  TouchPadT = 7
	TOUCH_PAD_NUM8  TouchPadT = 8
	TOUCH_PAD_NUM9  TouchPadT = 9
	TOUCH_PAD_NUM10 TouchPadT = 10
	TOUCH_PAD_NUM11 TouchPadT = 11
	TOUCH_PAD_NUM12 TouchPadT = 12
	TOUCH_PAD_NUM13 TouchPadT = 13
	TOUCH_PAD_NUM14 TouchPadT = 14
	TOUCH_PAD_MAX   TouchPadT = 15
)

type TouchHighVoltT c.Int

const (
	TOUCH_HVOLT_KEEP TouchHighVoltT = -1
	TOUCH_HVOLT_2V4  TouchHighVoltT = 0
	TOUCH_HVOLT_2V5  TouchHighVoltT = 1
	TOUCH_HVOLT_2V6  TouchHighVoltT = 2
	TOUCH_HVOLT_2V7  TouchHighVoltT = 3
	TOUCH_HVOLT_MAX  TouchHighVoltT = 4
)

type TouchLowVoltT c.Int

const (
	TOUCH_LVOLT_KEEP TouchLowVoltT = -1
	TOUCH_LVOLT_0V5  TouchLowVoltT = 0
	TOUCH_LVOLT_0V6  TouchLowVoltT = 1
	TOUCH_LVOLT_0V7  TouchLowVoltT = 2
	TOUCH_LVOLT_0V8  TouchLowVoltT = 3
	TOUCH_LVOLT_MAX  TouchLowVoltT = 4
)

type TouchVoltAttenT c.Int

const (
	TOUCH_HVOLT_ATTEN_KEEP TouchVoltAttenT = -1
	TOUCH_HVOLT_ATTEN_1V5  TouchVoltAttenT = 0
	TOUCH_HVOLT_ATTEN_1V   TouchVoltAttenT = 1
	TOUCH_HVOLT_ATTEN_0V5  TouchVoltAttenT = 2
	TOUCH_HVOLT_ATTEN_0V   TouchVoltAttenT = 3
	TOUCH_HVOLT_ATTEN_MAX  TouchVoltAttenT = 4
)

type TouchCntSlopeT c.Int

const (
	TOUCH_PAD_SLOPE_0   TouchCntSlopeT = 0
	TOUCH_PAD_SLOPE_1   TouchCntSlopeT = 1
	TOUCH_PAD_SLOPE_2   TouchCntSlopeT = 2
	TOUCH_PAD_SLOPE_3   TouchCntSlopeT = 3
	TOUCH_PAD_SLOPE_4   TouchCntSlopeT = 4
	TOUCH_PAD_SLOPE_5   TouchCntSlopeT = 5
	TOUCH_PAD_SLOPE_6   TouchCntSlopeT = 6
	TOUCH_PAD_SLOPE_7   TouchCntSlopeT = 7
	TOUCH_PAD_SLOPE_MAX TouchCntSlopeT = 8
)

type TouchTieOptT c.Int

const (
	TOUCH_PAD_TIE_OPT_LOW   TouchTieOptT = 0
	TOUCH_PAD_TIE_OPT_HIGH  TouchTieOptT = 1
	TOUCH_PAD_TIE_OPT_FLOAT TouchTieOptT = 2
	TOUCH_PAD_TIE_OPT_MAX   TouchTieOptT = 3
)

type TouchFsmModeT c.Int

const (
	TOUCH_FSM_MODE_TIMER TouchFsmModeT = 0
	TOUCH_FSM_MODE_SW    TouchFsmModeT = 1
	TOUCH_FSM_MODE_MAX   TouchFsmModeT = 2
)

type TouchTriggerModeT c.Int

const (
	TOUCH_TRIGGER_BELOW TouchTriggerModeT = 0
	TOUCH_TRIGGER_ABOVE TouchTriggerModeT = 1
	TOUCH_TRIGGER_MAX   TouchTriggerModeT = 2
)

type TouchTriggerSrcT c.Int

const (
	TOUCH_TRIGGER_SOURCE_BOTH TouchTriggerSrcT = 0
	TOUCH_TRIGGER_SOURCE_SET1 TouchTriggerSrcT = 1
	TOUCH_TRIGGER_SOURCE_MAX  TouchTriggerSrcT = 2
)

type TouchPadIntrMaskT c.Int

const (
	TOUCH_PAD_INTR_MASK_DONE            TouchPadIntrMaskT = 1
	TOUCH_PAD_INTR_MASK_ACTIVE          TouchPadIntrMaskT = 2
	TOUCH_PAD_INTR_MASK_INACTIVE        TouchPadIntrMaskT = 4
	TOUCH_PAD_INTR_MASK_SCAN_DONE       TouchPadIntrMaskT = 8
	TOUCH_PAD_INTR_MASK_TIMEOUT         TouchPadIntrMaskT = 16
	TOUCH_PAD_INTR_MASK_PROXI_MEAS_DONE TouchPadIntrMaskT = 32
	TOUCH_PAD_INTR_MASK_MAX             TouchPadIntrMaskT = 33
)

type TouchPadDenoiseGradeT c.Int

const (
	TOUCH_PAD_DENOISE_BIT12 TouchPadDenoiseGradeT = 0
	TOUCH_PAD_DENOISE_BIT10 TouchPadDenoiseGradeT = 1
	TOUCH_PAD_DENOISE_BIT8  TouchPadDenoiseGradeT = 2
	TOUCH_PAD_DENOISE_BIT4  TouchPadDenoiseGradeT = 3
	TOUCH_PAD_DENOISE_MAX   TouchPadDenoiseGradeT = 4
)

type TouchPadDenoiseCapT c.Int

const (
	TOUCH_PAD_DENOISE_CAP_L0  TouchPadDenoiseCapT = 0
	TOUCH_PAD_DENOISE_CAP_L1  TouchPadDenoiseCapT = 1
	TOUCH_PAD_DENOISE_CAP_L2  TouchPadDenoiseCapT = 2
	TOUCH_PAD_DENOISE_CAP_L3  TouchPadDenoiseCapT = 3
	TOUCH_PAD_DENOISE_CAP_L4  TouchPadDenoiseCapT = 4
	TOUCH_PAD_DENOISE_CAP_L5  TouchPadDenoiseCapT = 5
	TOUCH_PAD_DENOISE_CAP_L6  TouchPadDenoiseCapT = 6
	TOUCH_PAD_DENOISE_CAP_L7  TouchPadDenoiseCapT = 7
	TOUCH_PAD_DENOISE_CAP_MAX TouchPadDenoiseCapT = 8
)

/** Touch sensor denoise configuration */

type TouchPadDenoise struct {
	Grade    TouchPadDenoiseGradeT
	CapLevel TouchPadDenoiseCapT
}
type TouchPadDenoiseT TouchPadDenoise
type TouchPadShieldDriverT c.Int

const (
	TOUCH_PAD_SHIELD_DRV_L0  TouchPadShieldDriverT = 0
	TOUCH_PAD_SHIELD_DRV_L1  TouchPadShieldDriverT = 1
	TOUCH_PAD_SHIELD_DRV_L2  TouchPadShieldDriverT = 2
	TOUCH_PAD_SHIELD_DRV_L3  TouchPadShieldDriverT = 3
	TOUCH_PAD_SHIELD_DRV_L4  TouchPadShieldDriverT = 4
	TOUCH_PAD_SHIELD_DRV_L5  TouchPadShieldDriverT = 5
	TOUCH_PAD_SHIELD_DRV_L6  TouchPadShieldDriverT = 6
	TOUCH_PAD_SHIELD_DRV_L7  TouchPadShieldDriverT = 7
	TOUCH_PAD_SHIELD_DRV_MAX TouchPadShieldDriverT = 8
)

/** Touch sensor waterproof configuration */

type TouchPadWaterproof struct {
	GuardRingPad TouchPadT
	ShieldDriver TouchPadShieldDriverT
}
type TouchPadWaterproofT TouchPadWaterproof
type TouchPadConnTypeT c.Int

const (
	TOUCH_PAD_CONN_HIGHZ TouchPadConnTypeT = 0
	TOUCH_PAD_CONN_GND   TouchPadConnTypeT = 1
	TOUCH_PAD_CONN_MAX   TouchPadConnTypeT = 2
)

type TouchFilterModeT c.Int

const (
	TOUCH_PAD_FILTER_IIR_4   TouchFilterModeT = 0
	TOUCH_PAD_FILTER_IIR_8   TouchFilterModeT = 1
	TOUCH_PAD_FILTER_IIR_16  TouchFilterModeT = 2
	TOUCH_PAD_FILTER_IIR_32  TouchFilterModeT = 3
	TOUCH_PAD_FILTER_IIR_64  TouchFilterModeT = 4
	TOUCH_PAD_FILTER_IIR_128 TouchFilterModeT = 5
	TOUCH_PAD_FILTER_IIR_256 TouchFilterModeT = 6
	TOUCH_PAD_FILTER_JITTER  TouchFilterModeT = 7
	TOUCH_PAD_FILTER_MAX     TouchFilterModeT = 8
)

type TouchSmoothModeT c.Int

const (
	TOUCH_PAD_SMOOTH_OFF   TouchSmoothModeT = 0
	TOUCH_PAD_SMOOTH_IIR_2 TouchSmoothModeT = 1
	TOUCH_PAD_SMOOTH_IIR_4 TouchSmoothModeT = 2
	TOUCH_PAD_SMOOTH_IIR_8 TouchSmoothModeT = 3
	TOUCH_PAD_SMOOTH_MAX   TouchSmoothModeT = 4
)

/** Touch sensor filter configuration */

type TouchFilterConfig struct {
	Mode        TouchFilterModeT
	DebounceCnt c.Uint32T
	NoiseThr    c.Uint32T
	JitterStep  c.Uint32T
	SmhLvl      TouchSmoothModeT
}
type TouchFilterConfigT TouchFilterConfig

/** Touch sensor channel sleep configuration */

type TouchPadSleepChannelT struct {
	TouchNum    TouchPadT
	EnProximity bool
}
