package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SYSTIMER_COUNTER_ESPTIMER = 0
const SYSTIMER_COUNTER_OS_TICK = 1
const SYSTIMER_ALARM_OS_TICK_CORE0 = 0
const SYSTIMER_ALARM_OS_TICK_CORE1 = 1
const SYSTIMER_ALARM_ESPTIMER = 2

/**
 * @brief Convert ticks to microseconds
 *
 * @param ticks ticks to convert
 * @return microseconds
 */
//go:linkname SystimerTicksToUs C.systimer_ticks_to_us
func SystimerTicksToUs(ticks c.Uint64T) c.Uint64T

/**
 * @brief Convert microseconds to ticks
 *
 * @param us microseconds to convert
 * @return ticks
 */
//go:linkname SystimerUsToTicks C.systimer_us_to_ticks
func SystimerUsToTicks(us c.Uint64T) c.Uint64T
