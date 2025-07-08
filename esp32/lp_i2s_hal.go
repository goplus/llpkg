package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LpI2sSocHandleT *c.Uint32T

/**
 * Context that should be maintained by both the driver and the HAL
 */

type LpI2sHalContextT struct {
	Dev LpI2sSocHandleT
}

/**
 * @brief Init the LP I2S hal and set the LP I2S to the default configuration.
 * @note This function should be called first before other hal layer function is called.
 *
 * @param hal Context of the HAL layer
 * @param group_id LP I2S group ID
 */
// llgo:link (*LpI2sHalContextT).LpI2sHalInit C.lp_i2s_hal_init
func (recv_ *LpI2sHalContextT) LpI2sHalInit(group_id c.Int) {
}
