package lvgl

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

/**********************
 * GLOBAL PROTOTYPES
 **********************/

//! @cond Doxygen_Suppress

/**
 * Call it periodically to handle lv_timers.
 * @return time till it needs to be run next (in ms)
 */
// LV_ATTRIBUTE_TIMER_HANDLER uint32_t lv_timer_handler(void);

//go:linkname LvTimerHandler C.lv_timer_handler
func LvTimerHandler() c.Uint
