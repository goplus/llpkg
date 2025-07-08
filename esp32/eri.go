package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief  Perform an ERI write
 * @param  addr : ERI register to write to
 * @param  data : Value to write
 *
 * @return Value read
 */
//go:linkname EriWrite C.eri_write
func EriWrite(addr c.Int, data c.Uint32T)
