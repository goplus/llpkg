package freertos

import _ "unsafe"

/**
 * @brief USJ test mode values
 *
 * Specifies the logic values of each of the USB FSLS Serial PHY interface
 * signals when in test mode.
 */

type UsbSerialJtagPullOverrideValsT struct {
	DpPu bool
	DmPu bool
	DpPd bool
	DmPd bool
}
