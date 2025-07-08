package freertos

import _ "unsafe"

const XT_TIMER_INDEX = 1

//go:linkname X_xtTickDivisorInit C._xt_tick_divisor_init
func X_xtTickDivisorInit()
