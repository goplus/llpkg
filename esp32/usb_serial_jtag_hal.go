package freertos

import _ "unsafe"

/**
 * @brief HAL context type of USJ driver
 */

type UsbSerialJtagHalContextT struct {
	Dev *UsbSerialJtagDevT
}

/**
 * @brief Initialize the USJ HAL driver
 *
 * @param hal USJ HAL context
 */
// llgo:link (*UsbSerialJtagHalContextT).UsbSerialJtagHalInit C.usb_serial_jtag_hal_init
func (recv_ *UsbSerialJtagHalContextT) UsbSerialJtagHalInit() {
}

/**
 * @brief Configure whether USJ is routed to internal/external FSLS PHY
 *
 * @param hal USJ HAL context
 * @param external True if external, False if internal
 */
// llgo:link (*UsbSerialJtagHalContextT).UsbSerialJtagHalPhySetExternal C.usb_serial_jtag_hal_phy_set_external
func (recv_ *UsbSerialJtagHalContextT) UsbSerialJtagHalPhySetExternal(external bool) {
}
