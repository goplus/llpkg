package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Read a sequence of DPORT registers to the buffer.
 *
 * @param[out] buff_out  Contains the read data.
 * @param[in]  address   Initial address for reading registers.
 * @param[in]  num_words The number of words.
 */
//go:linkname EspDportAccessReadBuffer C.esp_dport_access_read_buffer
func EspDportAccessReadBuffer(buff_out *c.Uint32T, address c.Uint32T, num_words c.Uint32T)
