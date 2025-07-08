package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SUPPORT_WIFI = 1
const SUPPORT_BTDM = 1
const SUPPORT_USB_DWCOTG = 1

/* Structure and functions for returning ROM global layout
 *
 * This is for address symbols defined in the linker script, which may change during ECOs.
 */

type EtsRomLayoutT struct {
	Dram0StackSharedMemStart   c.Pointer
	Dram0RtosReservedStart     c.Pointer
	StackSentry                c.Pointer
	Stack                      c.Pointer
	StackSentryApp             c.Pointer
	StackApp                   c.Pointer
	DataStartBtdm              c.Pointer
	DataEndBtdm                c.Pointer
	BssStartBtdm               c.Pointer
	BssEndBtdm                 c.Pointer
	DataStartBtdmRom           c.Pointer
	DataEndBtdmRom             c.Pointer
	DataStartInterfaceBtdm     c.Pointer
	DataEndInterfaceBtdm       c.Pointer
	BssStartInterfaceBtdm      c.Pointer
	BssEndInterfaceBtdm        c.Pointer
	DramStartPhyrom            c.Pointer
	DramEndPhyrom              c.Pointer
	DramStartCoexist           c.Pointer
	DramEndCoexist             c.Pointer
	DramStartNet80211          c.Pointer
	DramEndNet80211            c.Pointer
	DramStartPp                c.Pointer
	DramEndPp                  c.Pointer
	DataStartInterfaceCoexist  c.Pointer
	DataEndInterfaceCoexist    c.Pointer
	BssStartInterfaceCoexist   c.Pointer
	BssEndInterfaceCoexist     c.Pointer
	DataStartInterfaceNet80211 c.Pointer
	DataEndInterfaceNet80211   c.Pointer
	BssStartInterfaceNet80211  c.Pointer
	BssEndInterfaceNet80211    c.Pointer
	DataStartInterfacePp       c.Pointer
	DataEndInterfacePp         c.Pointer
	BssStartInterfacePp        c.Pointer
	BssEndInterfacePp          c.Pointer
	DramStartUsbDwcotgRom      c.Pointer
	DramEndUsbDwcotgRom        c.Pointer
	DramStartUartRom           c.Pointer
	DramEndUartRom             c.Pointer
}
