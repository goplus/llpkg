package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Print backtrace for the given execution frame.
 *
 * @param frame_or Snapshot of the CPU registers when the program stopped its
 *                 normal execution. This frame is usually generated on the
 *                 stack when an exception or an interrupt occurs.
 */
//go:linkname EspEhFramePrintBacktrace C.esp_eh_frame_print_backtrace
func EspEhFramePrintBacktrace(frame_or c.Pointer)
