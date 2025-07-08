package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_DWC_QTD_LIST_MEM_ALIGN = 512
const USB_DWC_FRAME_LIST_MEM_ALIGN = 512
const USB_DWC_LL_QTD_STATUS_SUCCESS = 0x0
const USB_DWC_LL_QTD_STATUS_PKTERR = 0x1
const USB_DWC_LL_QTD_STATUS_BUFFER = 0x3
const USB_DWC_LL_QTD_STATUS_NOT_EXECUTED = 0x4

/*
 * QTD (Queue Transfer Descriptor) structure used in Scatter/Gather DMA mode.
 * Each QTD describes one transfer. Scatter gather mode will automatically split
 * a transfer into multiple MPS packets. Each QTD is 64bits in size
 *
 * Note: The status information part of the QTD is interpreted differently depending
 * on IN or OUT, and ISO or non-ISO
 */

type UsbDwcLlDmaQtdT struct {
	Buffer *c.Uint8T
}
