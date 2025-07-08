package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UsbDwcSpeedT c.Int

const (
	USB_DWC_SPEED_HIGH UsbDwcSpeedT = 0
	USB_DWC_SPEED_FULL UsbDwcSpeedT = 1
	USB_DWC_SPEED_LOW  UsbDwcSpeedT = 2
)

type UsbDwcXferTypeT c.Int

const (
	USB_DWC_XFER_TYPE_CTRL        UsbDwcXferTypeT = 0
	USB_DWC_XFER_TYPE_ISOCHRONOUS UsbDwcXferTypeT = 1
	USB_DWC_XFER_TYPE_BULK        UsbDwcXferTypeT = 2
	USB_DWC_XFER_TYPE_INTR        UsbDwcXferTypeT = 3
)

type UsbHalFrameListLenT c.Int

const (
	USB_HAL_FRAME_LIST_LEN_8  UsbHalFrameListLenT = 8
	USB_HAL_FRAME_LIST_LEN_16 UsbHalFrameListLenT = 16
	USB_HAL_FRAME_LIST_LEN_32 UsbHalFrameListLenT = 32
	USB_HAL_FRAME_LIST_LEN_64 UsbHalFrameListLenT = 64
)
