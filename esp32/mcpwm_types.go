package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type McpwmTimerClockSourceT SocPeriphMcpwmTimerClkSrcT
type McpwmCaptureClockSourceT SocPeriphMcpwmCaptureClkSrcT
type McpwmCarrierClockSourceT SocPeriphMcpwmCarrierClkSrcT
type McpwmTimerDirectionT c.Int

const (
	MCPWM_TIMER_DIRECTION_UP   McpwmTimerDirectionT = 0
	MCPWM_TIMER_DIRECTION_DOWN McpwmTimerDirectionT = 1
)

type McpwmTimerEventT c.Int

const (
	MCPWM_TIMER_EVENT_EMPTY   McpwmTimerEventT = 0
	MCPWM_TIMER_EVENT_FULL    McpwmTimerEventT = 1
	MCPWM_TIMER_EVENT_INVALID McpwmTimerEventT = 2
)

type McpwmTimerCountModeT c.Int

const (
	MCPWM_TIMER_COUNT_MODE_PAUSE   McpwmTimerCountModeT = 0
	MCPWM_TIMER_COUNT_MODE_UP      McpwmTimerCountModeT = 1
	MCPWM_TIMER_COUNT_MODE_DOWN    McpwmTimerCountModeT = 2
	MCPWM_TIMER_COUNT_MODE_UP_DOWN McpwmTimerCountModeT = 3
)

type McpwmTimerStartStopCmdT c.Int

const (
	MCPWM_TIMER_STOP_EMPTY       McpwmTimerStartStopCmdT = 0
	MCPWM_TIMER_STOP_FULL        McpwmTimerStartStopCmdT = 1
	MCPWM_TIMER_START_NO_STOP    McpwmTimerStartStopCmdT = 2
	MCPWM_TIMER_START_STOP_EMPTY McpwmTimerStartStopCmdT = 3
	MCPWM_TIMER_START_STOP_FULL  McpwmTimerStartStopCmdT = 4
)

type McpwmGeneratorActionT c.Int

const (
	MCPWM_GEN_ACTION_KEEP   McpwmGeneratorActionT = 0
	MCPWM_GEN_ACTION_LOW    McpwmGeneratorActionT = 1
	MCPWM_GEN_ACTION_HIGH   McpwmGeneratorActionT = 2
	MCPWM_GEN_ACTION_TOGGLE McpwmGeneratorActionT = 3
)

type McpwmOperatorBrakeModeT c.Int

const (
	MCPWM_OPER_BRAKE_MODE_CBC     McpwmOperatorBrakeModeT = 0
	MCPWM_OPER_BRAKE_MODE_OST     McpwmOperatorBrakeModeT = 1
	MCPWM_OPER_BRAKE_MODE_INVALID McpwmOperatorBrakeModeT = 2
)

type McpwmCaptureEdgeT c.Int

const (
	MCPWM_CAP_EDGE_POS McpwmCaptureEdgeT = 0
	MCPWM_CAP_EDGE_NEG McpwmCaptureEdgeT = 1
)

type McpwmComparatorEtmEventTypeT c.Int

const (
	MCPWM_CMPR_ETM_EVENT_EQUAL McpwmComparatorEtmEventTypeT = 0
	MCPWM_CMPR_ETM_EVENT_MAX   McpwmComparatorEtmEventTypeT = 1
)
