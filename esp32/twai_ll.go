package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * The following frame structure has an NEARLY identical bit field layout to
 * each byte of the TX buffer. This allows for formatting and parsing frames to
 * be done outside of time critical regions (i.e., ISRs). All the ISR needs to
 * do is to copy byte by byte to/from the TX/RX buffer. The two reserved bits in
 * TX buffer are used in the frame structure to store the self_reception and
 * single_shot flags which in turn indicate the type of transmission to execute.
 */

type TwaiLlFrameBufferT struct {
	Bytes [13]c.Uint8T
}
