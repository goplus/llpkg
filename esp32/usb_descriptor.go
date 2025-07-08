package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_DESCRIPTOR_TYPE_ACM = 0
const USB_DESCRIPTOR_TYPE_DFU = 1

//go:linkname UsbSetCurrentDescriptor C.usb_set_current_descriptor
func UsbSetCurrentDescriptor(descriptor_type c.Int)

//go:linkname UsbGetDescriptor C.usb_get_descriptor
func UsbGetDescriptor(type_index c.Uint16T, lang_id c.Uint16T, len *c.Int32T, data **c.Uint8T) bool
