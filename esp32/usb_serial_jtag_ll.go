package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_SERIAL_JTAG_LL_EXT_PHY_SUPPORTED = 1

type UsbSerialJtagIntrT c.Int

const (
	USB_SERIAL_JTAG_INTR_SOF                 UsbSerialJtagIntrT = 2
	USB_SERIAL_JTAG_INTR_SERIAL_OUT_RECV_PKT UsbSerialJtagIntrT = 4
	USB_SERIAL_JTAG_INTR_SERIAL_IN_EMPTY     UsbSerialJtagIntrT = 8
	USB_SERIAL_JTAG_INTR_TOKEN_REC_IN_EP1    UsbSerialJtagIntrT = 256
	USB_SERIAL_JTAG_INTR_BUS_RESET           UsbSerialJtagIntrT = 512
	USB_SERIAL_JTAG_INTR_EP1_ZERO_PAYLOAD    UsbSerialJtagIntrT = 1024
)
