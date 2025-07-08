package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

//go:linkname ChipUsbDwInit C.chip_usb_dw_init
func ChipUsbDwInit() c.Int

//go:linkname ChipUsbDwDidPersist C.chip_usb_dw_did_persist
func ChipUsbDwDidPersist() c.Int

//go:linkname ChipUsbDwPreparePersist C.chip_usb_dw_prepare_persist
func ChipUsbDwPreparePersist()

//go:linkname ChipUsbGetPersistFlags C.chip_usb_get_persist_flags
func ChipUsbGetPersistFlags() c.Uint32T

//go:linkname ChipUsbSetPersistFlags C.chip_usb_set_persist_flags
func ChipUsbSetPersistFlags(flags c.Uint32T)
