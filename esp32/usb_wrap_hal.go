package freertos

import _ "unsafe"

/**
 * @brief HAL context type of USB WRAP driver
 */

type UsbWrapHalContextT struct {
	Dev *UsbWrapDevT
}

/**
 * @brief Initialize the USB WRAP HAL driver
 *
 * @param hal USB WRAP HAL context
 */
// llgo:link (*UsbWrapHalContextT).X_usbWrapHalInit C._usb_wrap_hal_init
func (recv_ *UsbWrapHalContextT) X_usbWrapHalInit() {
}

/**
 * @brief Disable USB WRAP
 *
 * Disable clock to the peripheral
 */
//go:linkname X_usbWrapHalDisable C._usb_wrap_hal_disable
func X_usbWrapHalDisable()

/**
 * @brief Configure whether USB WRAP is routed to internal/external FSLS PHY
 *
 * @param hal USB WRAP HAL context
 * @param external True if external, False if internal
 */
// llgo:link (*UsbWrapHalContextT).UsbWrapHalPhySetExternal C.usb_wrap_hal_phy_set_external
func (recv_ *UsbWrapHalContextT) UsbWrapHalPhySetExternal(external bool) {
}
