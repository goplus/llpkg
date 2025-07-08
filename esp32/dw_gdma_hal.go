package freertos

import _ "unsafe"

type DwGdmaDevT struct {
	Unused [8]uint8
}
type DwGdmaSocHandleT *DwGdmaDevT

/**
 * @brief DW_GDMA HAL driver context
 */

type DwGdmaHalContextT struct {
	Dev DwGdmaSocHandleT
}

/**
 * @brief DW_GDMA HAL driver configuration
 */

type DwGdmaHalConfigT struct {
	Unused [8]uint8
}

/**
 * @brief DW_GDMA HAL driver initialization
 *
 * @note Caller should malloc the memory for the hal context
 *
 * @param hal Pointer to the HAL driver context
 * @param config Pointer to the HAL driver configuration
 */
// llgo:link (*DwGdmaHalContextT).DwGdmaHalInit C.dw_gdma_hal_init
func (recv_ *DwGdmaHalContextT) DwGdmaHalInit(config *DwGdmaHalConfigT) {
}

/**
 * @brief DW_GDMA HAL driver deinitialization
 *
 * @param hal Pointer to the HAL driver context
 */
// llgo:link (*DwGdmaHalContextT).DwGdmaHalDeinit C.dw_gdma_hal_deinit
func (recv_ *DwGdmaHalContextT) DwGdmaHalDeinit() {
}
