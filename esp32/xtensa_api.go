package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type XtExcHandler func(*XtExcFrame)

/*
-------------------------------------------------------------------------------
  Call this function to set a handler for the specified exception. The handler
  will be installed on the core that calls this function.

    n        - Exception number (type)
    f        - Handler function address, NULL to uninstall handler.

  The handler will be passed a pointer to the exception frame, which is created
  on the stack of the thread that caused the exception.

  If the handler returns, the thread context will be restored and the faulting
  instruction will be retried. Any values in the exception frame that are
  modified by the handler will be restored as part of the context. For details
  of the exception frame structure see xtensa_context.h.
-------------------------------------------------------------------------------
*/
//go:linkname XtSetExceptionHandler C.xt_set_exception_handler
func XtSetExceptionHandler(n c.Int, f XtExcHandler) XtExcHandler

/*
-------------------------------------------------------------------------------
  Call this function to set a handler for the specified interrupt. The handler
  will be installed on the core that calls this function.

    n        - Interrupt number.
    f        - Handler function address, NULL to uninstall handler.
    arg      - Argument to be passed to handler.
-------------------------------------------------------------------------------
*/
//go:linkname XtSetInterruptHandler C.xt_set_interrupt_handler
func XtSetInterruptHandler(n c.Int, f c.Int, arg c.Pointer) c.Int

/*
-------------------------------------------------------------------------------
  Call this function to enable the specified interrupts on the core that runs
  this code.

    mask     - Bit mask of interrupts to be enabled.
-------------------------------------------------------------------------------
*/
//go:linkname XtIntsOn C.xt_ints_on
func XtIntsOn(mask c.Uint)

/*
-------------------------------------------------------------------------------
  Call this function to disable the specified interrupts on the core that runs
  this code.

    mask     - Bit mask of interrupts to be disabled.
-------------------------------------------------------------------------------
*/
//go:linkname XtIntsOff C.xt_ints_off
func XtIntsOff(mask c.Uint)

/*
-------------------------------------------------------------------------------
  Call this function to get handler's argument for the specified interrupt.

    n        - Interrupt number.
-------------------------------------------------------------------------------
*/
//go:linkname XtGetInterruptHandlerArg C.xt_get_interrupt_handler_arg
func XtGetInterruptHandlerArg(n c.Int) c.Pointer

/*
-------------------------------------------------------------------------------
  Call this function to check if the specified interrupt is free to use.

    intr       - Interrupt number.
    cpu        - cpu number.
-------------------------------------------------------------------------------
*/
//go:linkname XtIntHasHandler C.xt_int_has_handler
func XtIntHasHandler(intr c.Int, cpu c.Int) bool
