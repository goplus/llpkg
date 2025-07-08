package freertos

import _ "unsafe"

type PpaDevT struct {
	Unused [8]uint8
}
type PpaSocHandleT *PpaDevT

/**
 * Context that should be maintained by both the driver and the HAL
 */

type PpaHalContextT struct {
	Dev PpaSocHandleT
}

/**
 * @brief Init the PPA hal. This function should be called first before other hal layer function is called
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*PpaHalContextT).PpaHalInit C.ppa_hal_init
func (recv_ *PpaHalContextT) PpaHalInit() {
}

/**
 * @brief De-init the PPA hal
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*PpaHalContextT).PpaHalDeinit C.ppa_hal_deinit
func (recv_ *PpaHalContextT) PpaHalDeinit() {
}
