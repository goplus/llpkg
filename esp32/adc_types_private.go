package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcArbiterModeT c.Int

const (
	ADC_ARB_MODE_SHIELD AdcArbiterModeT = 0
	ADC_ARB_MODE_FIX    AdcArbiterModeT = 1
	ADC_ARB_MODE_LOOP   AdcArbiterModeT = 2
)

/**
 * @brief ADC arbiter work mode and priority setting.
 *
 * @note Only ADC2 support arbiter.
 */

type AdcArbiterT struct {
	Mode     AdcArbiterModeT
	RtcPri   c.Uint8T
	DigPri   c.Uint8T
	PwdetPri c.Uint8T
}
