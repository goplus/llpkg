package freertos

import _ "unsafe"

type ParlIoDevT struct {
	Unused [8]uint8
}
type ParlioSocHandleT *ParlIoDevT

/**
 * @brief HAL context type of Parallel IO driver
 */

type ParlioHalContextT struct {
	Regs ParlioSocHandleT
}

/**
 * @brief Initialize the Parallel IO HAL driver
 *
 * @param hal: Parallel IO HAL context
 */
// llgo:link (*ParlioHalContextT).ParlioHalInit C.parlio_hal_init
func (recv_ *ParlioHalContextT) ParlioHalInit() {
}

/**
 * @brief Deinitialize the Parallel IO HAL driver
 *
 * @param hal: Parallel IO HAL context
 */
// llgo:link (*ParlioHalContextT).ParlioHalDeinit C.parlio_hal_deinit
func (recv_ *ParlioHalContextT) ParlioHalDeinit() {
}
