package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief A pin descriptor for init (DEPRECATED)
 *
 * Todo: Remove in IDF v6.0 (IDF-9029)
 *
 * @note These IO pins are deprecated. When connecting USB OTG to an external
 * FSLS PHY, the FSLS Serial Interface signals can be routed to any GPIO via the
 * GPI0 matrix. Thus, this mapping of signals to IO pins is meaningless.
 */

type UsbIopinDscT struct {
	Pin        c.Int
	Func       c.Int
	IsOutput   bool
	ExtPhyOnly c.Int
}
