package freertos

import _ "unsafe"

type SocEtmDevT struct {
	Unused [8]uint8
}
type EtmSocHandleT *SocEtmDevT

/**
 * @brief HAL context type of ETM driver
 */

type EtmHalContextT struct {
	Regs EtmSocHandleT
}

/**
 * @brief Initialize the ETM HAL driver
 *
 * @param hal: ETM HAL context
 */
// llgo:link (*EtmHalContextT).EtmHalInit C.etm_hal_init
func (recv_ *EtmHalContextT) EtmHalInit() {
}

/**
 * @brief Deinitialize the ETM HAL driver
 *
 * @param hal: ETM HAL context
 */
// llgo:link (*EtmHalContextT).EtmHalDeinit C.etm_hal_deinit
func (recv_ *EtmHalContextT) EtmHalDeinit() {
}
