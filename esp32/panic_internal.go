package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type PanicInfoDumpFnT func(c.Pointer)
type PanicExceptionT c.Int

const (
	PANIC_EXCEPTION_DEBUG PanicExceptionT = 0
	PANIC_EXCEPTION_IWDT  PanicExceptionT = 1
	PANIC_EXCEPTION_TWDT  PanicExceptionT = 2
	PANIC_EXCEPTION_ABORT PanicExceptionT = 3
	PANIC_EXCEPTION_FAULT PanicExceptionT = 4
)

type PanicInfoT struct {
	Core          c.Int
	Exception     PanicExceptionT
	Reason        *c.Char
	Description   *c.Char
	Details       PanicInfoDumpFnT
	State         PanicInfoDumpFnT
	Addr          c.Pointer
	Frame         c.Pointer
	PseudoExcause bool
}

// Create own print functions, since printf might be broken, and can be silenced
// when CONFIG_ESP_SYSTEM_PANIC_SILENT_REBOOT
//
//go:linkname PanicPrintChar C.panic_print_char
func PanicPrintChar(c c.Char)

//go:linkname PanicPrintStr C.panic_print_str
func PanicPrintStr(str *c.Char)

//go:linkname PanicPrintDec C.panic_print_dec
func PanicPrintDec(d c.Int)

//go:linkname PanicPrintHex C.panic_print_hex
func PanicPrintHex(h c.Int)

//go:linkname PanicAbort C.panic_abort
func PanicAbort(details *c.Char)

//go:linkname PanicArchFillInfo C.panic_arch_fill_info
func PanicArchFillInfo(frame c.Pointer, info *PanicInfoT)

//go:linkname PanicSocFillInfo C.panic_soc_fill_info
func PanicSocFillInfo(frame c.Pointer, info *PanicInfoT)

//go:linkname PanicSocCheckPseudoCause C.panic_soc_check_pseudo_cause
func PanicSocCheckPseudoCause(f c.Pointer, info *PanicInfoT) bool

//go:linkname PanicPrintRegisters C.panic_print_registers
func PanicPrintRegisters(frame c.Pointer, core c.Int)

//go:linkname PanicPrintBacktrace C.panic_print_backtrace
func PanicPrintBacktrace(frame c.Pointer, core c.Int)

//go:linkname PanicGetAddress C.panic_get_address
func PanicGetAddress(frame c.Pointer) c.Uint32T

//go:linkname PanicSetAddress C.panic_set_address
func PanicSetAddress(frame c.Pointer, addr c.Uint32T)

//go:linkname PanicGetCause C.panic_get_cause
func PanicGetCause(frame c.Pointer) c.Uint32T

//go:linkname PanicPrepareFrameFromCtx C.panic_prepare_frame_from_ctx
func PanicPrepareFrameFromCtx(frame c.Pointer)
