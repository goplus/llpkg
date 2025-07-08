package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LedcModeT c.Int

const (
	LEDC_LOW_SPEED_MODE LedcModeT = 0
	LEDC_SPEED_MODE_MAX LedcModeT = 1
)

type LedcIntrTypeT c.Int

const (
	LEDC_INTR_DISABLE  LedcIntrTypeT = 0
	LEDC_INTR_FADE_END LedcIntrTypeT = 1
	LEDC_INTR_MAX      LedcIntrTypeT = 2
)

type LedcDutyDirectionT c.Int

const (
	LEDC_DUTY_DIR_DECREASE LedcDutyDirectionT = 0
	LEDC_DUTY_DIR_INCREASE LedcDutyDirectionT = 1
	LEDC_DUTY_DIR_MAX      LedcDutyDirectionT = 2
)

type LedcSlowClkSelT c.Int

const (
	LEDC_SLOW_CLK_RC_FAST LedcSlowClkSelT = 9
	LEDC_SLOW_CLK_APB     LedcSlowClkSelT = 4
	LEDC_SLOW_CLK_XTAL    LedcSlowClkSelT = 11
	LEDC_SLOW_CLK_RTC8M   LedcSlowClkSelT = 9
)

type LedcClkCfgT SocPeriphLedcClkSrcLegacyT
type LedcClkSrcT c.Int

const (
	LEDC_APB_CLK LedcClkSrcT = 4
	LEDC_SCLK    LedcClkSrcT = 4
)

type LedcTimerT c.Int

const (
	LEDC_TIMER_0   LedcTimerT = 0
	LEDC_TIMER_1   LedcTimerT = 1
	LEDC_TIMER_2   LedcTimerT = 2
	LEDC_TIMER_3   LedcTimerT = 3
	LEDC_TIMER_MAX LedcTimerT = 4
)

type LedcChannelT c.Int

const (
	LEDC_CHANNEL_0   LedcChannelT = 0
	LEDC_CHANNEL_1   LedcChannelT = 1
	LEDC_CHANNEL_2   LedcChannelT = 2
	LEDC_CHANNEL_3   LedcChannelT = 3
	LEDC_CHANNEL_4   LedcChannelT = 4
	LEDC_CHANNEL_5   LedcChannelT = 5
	LEDC_CHANNEL_6   LedcChannelT = 6
	LEDC_CHANNEL_7   LedcChannelT = 7
	LEDC_CHANNEL_MAX LedcChannelT = 8
)

type LedcTimerBitT c.Int

const (
	LEDC_TIMER_1_BIT   LedcTimerBitT = 1
	LEDC_TIMER_2_BIT   LedcTimerBitT = 2
	LEDC_TIMER_3_BIT   LedcTimerBitT = 3
	LEDC_TIMER_4_BIT   LedcTimerBitT = 4
	LEDC_TIMER_5_BIT   LedcTimerBitT = 5
	LEDC_TIMER_6_BIT   LedcTimerBitT = 6
	LEDC_TIMER_7_BIT   LedcTimerBitT = 7
	LEDC_TIMER_8_BIT   LedcTimerBitT = 8
	LEDC_TIMER_9_BIT   LedcTimerBitT = 9
	LEDC_TIMER_10_BIT  LedcTimerBitT = 10
	LEDC_TIMER_11_BIT  LedcTimerBitT = 11
	LEDC_TIMER_12_BIT  LedcTimerBitT = 12
	LEDC_TIMER_13_BIT  LedcTimerBitT = 13
	LEDC_TIMER_14_BIT  LedcTimerBitT = 14
	LEDC_TIMER_BIT_MAX LedcTimerBitT = 15
)

type LedcFadeModeT c.Int

const (
	LEDC_FADE_NO_WAIT   LedcFadeModeT = 0
	LEDC_FADE_WAIT_DONE LedcFadeModeT = 1
	LEDC_FADE_MAX       LedcFadeModeT = 2
)
