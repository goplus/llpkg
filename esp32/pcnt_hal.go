package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PcntSocHandleT *PcntDevT

/**
 * Context that should be maintained by both the driver and the HAL
 */

type PcntHalContextT struct {
	Dev PcntSocHandleT
}

/**
 * @brief Init the PCNT hal and set the PCNT to the default configuration.
 * @note This function should be called first before other hal layer function is called.
 *
 * @param hal Context of the HAL layer
 * @param group_id PCNT group ID
 */
// llgo:link (*PcntHalContextT).PcntHalInit C.pcnt_hal_init
func (recv_ *PcntHalContextT) PcntHalInit(group_id c.Int) {
}
