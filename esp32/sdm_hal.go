package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SdmSocHandleT *GpioSdDevT

/**
 * HAL context type of Sigma-Delta driver
 */

type SdmHalContextT struct {
	Dev SdmSocHandleT
}

/**
 * @brief Initialize Sigma-Delta hal driver
 *
 * @param hal Context of the HAL layer
 * @param group_id Sigma-Delta group number
 */
// llgo:link (*SdmHalContextT).SdmHalInit C.sdm_hal_init
func (recv_ *SdmHalContextT) SdmHalInit(group_id c.Int) {
}
