package freertos

import _ "unsafe"

/**
 * @brief USB WRAP pull up/down resistor override values
 *
 * Specifies whether each pull up/down resistor should be enabled/disabled when
 * overriding connected USB PHY's pull resistors.
 */

type UsbWrapPullOverrideValsT struct {
	DpPu bool
	DmPu bool
	DpPd bool
	DmPd bool
}

/**
 * @brief USB WRAP test mode values
 *
 * Specifies the logic values of each of the USB FSLS Serial PHY interface
 * signals when in test mode.
 *
 * @note See section "2.2.1.13 FsLsSerialMode" of UTMI+ specification for more
 * details of each signal.
 */

type UsbWrapTestModeValsT struct {
	TxEnableN bool
	TxDp      bool
	TxDm      bool
	RxDp      bool
	RxDm      bool
	RxRcv     bool
}
