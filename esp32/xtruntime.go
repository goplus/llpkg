package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const XTOS_KEEPON_MEM = 0x00000100
const XTOS_KEEPON_MEM_SHIFT = 8
const XTOS_KEEPON_DEBUG = 0x00001000
const XTOS_KEEPON_DEBUG_SHIFT = 12
const XTOS_IDMA_NO_WAIT = 0x00010000
const XTOS_IDMA_WAIT_STANDBY = 0x00020000
const XTOS_COREF_PSO = 0x00000001
const XTOS_COREF_PSO_SHIFT = 0

// llgo:type C
type X_xtosHandlerFunc func()
type X_xtosHandler *X_xtosHandlerFunc

/*  These two are deprecated. Use the newer functions below.  */
//go:linkname X_xtosIntsOff C._xtos_ints_off
func X_xtosIntsOff(mask c.Uint) c.Uint

//go:linkname X_xtosIntsOn C._xtos_ints_on
func X_xtosIntsOn(mask c.Uint) c.Uint

//go:linkname X_xtosSetIntlevel C._xtos_set_intlevel
func X_xtosSetIntlevel(intlevel c.Int) c.Uint

//go:linkname X_xtosSetMinIntlevel C._xtos_set_min_intlevel
func X_xtosSetMinIntlevel(intlevel c.Int) c.Uint

//go:linkname X_xtosRestoreIntlevel C._xtos_restore_intlevel
func X_xtosRestoreIntlevel(restoreval c.Uint) c.Uint

//go:linkname X_xtosRestoreJustIntlevel C._xtos_restore_just_intlevel
func X_xtosRestoreJustIntlevel(restoreval c.Uint) c.Uint

//go:linkname X_xtosSetInterruptHandler C._xtos_set_interrupt_handler
func X_xtosSetInterruptHandler(n c.Int, f X_xtosHandler) X_xtosHandler

//go:linkname X_xtosSetInterruptHandlerArg C._xtos_set_interrupt_handler_arg
func X_xtosSetInterruptHandlerArg(n c.Int, f X_xtosHandler, arg c.Pointer) X_xtosHandler

//go:linkname X_xtosSetExceptionHandler C._xtos_set_exception_handler
func X_xtosSetExceptionHandler(n c.Int, f X_xtosHandler) X_xtosHandler

//go:linkname X_xtosMemepInitrams C._xtos_memep_initrams
func X_xtosMemepInitrams()

//go:linkname X_xtosMemepEnable C._xtos_memep_enable
func X_xtosMemepEnable(flags c.Int)

/*  For use with the tiny LSP (see LSP reference manual).  */
//go:linkname X_xtosDispatchLevel1Interrupts C._xtos_dispatch_level1_interrupts
func X_xtosDispatchLevel1Interrupts()

//go:linkname X_xtosDispatchLevel2Interrupts C._xtos_dispatch_level2_interrupts
func X_xtosDispatchLevel2Interrupts()

//go:linkname X_xtosDispatchLevel3Interrupts C._xtos_dispatch_level3_interrupts
func X_xtosDispatchLevel3Interrupts()

//go:linkname X_xtosDispatchLevel4Interrupts C._xtos_dispatch_level4_interrupts
func X_xtosDispatchLevel4Interrupts()

//go:linkname X_xtosDispatchLevel5Interrupts C._xtos_dispatch_level5_interrupts
func X_xtosDispatchLevel5Interrupts()

//go:linkname X_xtosDispatchLevel6Interrupts C._xtos_dispatch_level6_interrupts
func X_xtosDispatchLevel6Interrupts()

/*  Deprecated (but kept because they were documented):  */
//go:linkname X_xtosReadInts C._xtos_read_ints
func X_xtosReadInts() c.Uint

//go:linkname X_xtosClearInts C._xtos_clear_ints
func X_xtosClearInts(mask c.Uint)

/*  Power shut-off related routines.  */
//go:linkname X_xtosCoreShutoff C._xtos_core_shutoff
func X_xtosCoreShutoff(flags c.Uint) c.Int

//go:linkname X_xtosCoreSave C._xtos_core_save
func X_xtosCoreSave(flags c.Uint, savearea *XtosCoreState, code c.Pointer) c.Int

//go:linkname X_xtosCoreRestore C._xtos_core_restore
func X_xtosCoreRestore(retvalue c.Uint, savearea *XtosCoreState)

/*  Deprecated:  */
//go:linkname X_xtosTimer0Delta C._xtos_timer_0_delta
func X_xtosTimer0Delta(cycles c.Int)

//go:linkname X_xtosTimer1Delta C._xtos_timer_1_delta
func X_xtosTimer1Delta(cycles c.Int)

//go:linkname X_xtosTimer2Delta C._xtos_timer_2_delta
func X_xtosTimer2Delta(cycles c.Int)
