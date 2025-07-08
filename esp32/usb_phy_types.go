package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UsbPhyTargetT c.Int

const (
	USB_PHY_TARGET_INT  UsbPhyTargetT = 0
	USB_PHY_TARGET_UTMI UsbPhyTargetT = 1
	USB_PHY_TARGET_EXT  UsbPhyTargetT = 2
	USB_PHY_TARGET_MAX  UsbPhyTargetT = 3
)

type UsbPhyControllerT c.Int

const (
	USB_PHY_CTRL_OTG         UsbPhyControllerT = 0
	USB_PHY_CTRL_SERIAL_JTAG UsbPhyControllerT = 1
	USB_PHY_CTRL_MAX         UsbPhyControllerT = 2
)

type UsbOtgModeT c.Int

const (
	USB_PHY_MODE_DEFAULT UsbOtgModeT = 0
	USB_OTG_MODE_HOST    UsbOtgModeT = 1
	USB_OTG_MODE_DEVICE  UsbOtgModeT = 2
	USB_OTG_MODE_MAX     UsbOtgModeT = 3
)

type UsbPhySpeedT c.Int

const (
	USB_PHY_SPEED_UNDEFINED UsbPhySpeedT = 0
	USB_PHY_SPEED_LOW       UsbPhySpeedT = 1
	USB_PHY_SPEED_FULL      UsbPhySpeedT = 2
	USB_PHY_SPEED_HIGH      UsbPhySpeedT = 3
	USB_PHY_SPEED_MAX       UsbPhySpeedT = 4
)
