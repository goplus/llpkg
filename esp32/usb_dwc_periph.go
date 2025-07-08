package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UsbPhyInstT c.Int

const (
	USB_PHY_INST_FSLS_INTERN_0 UsbPhyInstT = 1
	USB_PHY_INST_FSLS_INTERN_1 UsbPhyInstT = 2
	USB_PHY_INST_UTMI_0        UsbPhyInstT = 4
	USB_PHY_INST_EXTERN        UsbPhyInstT = 8
)

/**
 * @brief USB PHY FSLS Serial Interface Signals
 *
 * Structure to store the GPIO matrix signal indexes for a USB PHY FSLS Serial
 * interface's signals.
 *
 * @note Refer to section "2.2.1.13 FsLsSerialMode" of the UTMI+ for more
 *       details regarding the FSLS Serial Interface.
 */

type UsbFslsSerialSignalConnT struct {
	RxDp      c.Int
	RxDm      c.Int
	RxRcv     c.Int
	SuspendN  c.Int
	TxEnableN c.Int
	TxDp      c.Int
	TxDm      c.Int
	FsEdgeSel c.Int
}

/**
 * @brief USB PHY UTMI OTG Interface Signal Index Type
 *
 * Structure to store the GPIO matrix signal indexes for a UTMI PHY interface's
 * OTG signals.
 */

type UsbOtgSignalConnT struct {
	Iddig       c.Int
	Avalid      c.Int
	Bvalid      c.Int
	Vbusvalid   c.Int
	Sessend     c.Int
	Idpullup    c.Int
	Dppulldown  c.Int
	Dmpulldown  c.Int
	Drvvbus     c.Int
	Chrgvbus    c.Int
	Dischrgvbus c.Int
}

/**
 * @brief Internal USB PHY IO
 *
 * Structure to store the IO numbers for a particular internal USB PHY
 */

type UsbInternalPhyIoT struct {
	Dp c.Int
	Dm c.Int
}

/**
 * @brief USB Controller Information
 *
 * Structure to store information for all USB-DWC instances
 *
 * For targets with multiple USB controllers, we support only fixed mapping of the PHYs.
 * This is a software limitation; the hardware supports swapping Controllers and PHYs.
 */

type UsbDwcInfoT struct {
	Controllers [1]struct {
		FslsSignals   *UsbFslsSerialSignalConnT
		OtgSignals    *UsbOtgSignalConnT
		InternalPhyIo *UsbInternalPhyIoT
		SupportedPhys UsbPhyInstT
		Irq           c.Int
		Irq2ndCpu     c.Int
	}
}

/*
Stores a bunch of USB-peripheral data.
*/
type UsbPhySignalConnT struct {
	ExtphyVpIn        c.Uint8T
	ExtphyVmIn        c.Uint8T
	ExtphyRcvIn       c.Uint8T
	ExtphyOenOut      c.Uint8T
	ExtphyVpoOut      c.Uint8T
	ExtphyVmoOut      c.Uint8T
	ExtphySuspendIn   c.Uint8T
	ExtphySpeedIn     c.Uint8T
	SrpBvalidIn       c.Uint8T
	SrpSessendIn      c.Uint8T
	SrpChrgvbusOut    c.Uint8T
	SrpDischrgvbusOut c.Uint8T
	OtgIddigIn        c.Uint8T
	OtgAvalidIn       c.Uint8T
	OtgVbusvalidIn    c.Uint8T
	OtgIdpullupOut    c.Uint8T
	OtgDppulldownOut  c.Uint8T
	OtgDmpulldownOut  c.Uint8T
	OtgDrvvbusOut     c.Uint8T
	Module            PeriphModuleT
}
