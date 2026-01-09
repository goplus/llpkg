package tick

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// Function types for callbacks
type LvTickGetCbT func() c.Uint32T
type LvDelayCbT func(c.Uint32T)

//go:linkname LvTickInc C.lv_tick_inc
func LvTickInc(tickPeriod c.Uint32T)

//go:linkname LvTickGet C.lv_tick_get
func LvTickGet() c.Uint32T

//go:linkname LvTickElaps C.lv_tick_elaps
func LvTickElaps(prevTick c.Uint32T) c.Uint32T

//go:linkname LvDelayMs C.lv_delay_ms
func LvDelayMs(ms c.Uint32T)

//go:linkname LvTickSetCb C.lv_tick_set_cb
func LvTickSetCb(cb LvTickGetCbT)

//go:linkname LvDelaySetCb C.lv_delay_set_cb
func LvDelaySetCb(cb LvDelayCbT)
