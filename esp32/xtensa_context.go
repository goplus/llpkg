package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const XT_CP0_SA = 0
const XT_CPENABLE = 0
const XT_CPSTORED = 2
const XT_CP_CS_ST = 4
const XT_CP_ASA = 8
const CORE_ID_REGVAL_PRO = 0xCDCD
const CORE_ID_REGVAL_APP = 0xABAB

/*
-------------------------------------------------------------------------------

	INTERRUPT/EXCEPTION STACK FRAME FOR A THREAD OR NESTED INTERRUPT

	A stack frame of this structure is allocated for any interrupt or exception.
	It goes on the current stack. If the RTOS has a system stack for handling
	interrupts, every thread stack must allow space for just one interrupt stack
	frame, then nested interrupt stack frames go on the system stack.

	The frame includes basic registers (explicit) and "extra" registers introduced
	by user TIE or the use of the MAC16 option in the user's Xtensa config.
	The frame size is minimized by omitting regs not applicable to user's config.

	For Windowed ABI, this stack frame includes the interruptee's base save area,
	another base save area to manage gcc nested functions, and a little temporary
	space to help manage the spilling of the register windows.

-------------------------------------------------------------------------------
*/
type XtExcFrame struct {
	Exit     c.Long
	Pc       c.Long
	Ps       c.Long
	A0       c.Long
	A1       c.Long
	A2       c.Long
	A3       c.Long
	A4       c.Long
	A5       c.Long
	A6       c.Long
	A7       c.Long
	A8       c.Long
	A9       c.Long
	A10      c.Long
	A11      c.Long
	A12      c.Long
	A13      c.Long
	A14      c.Long
	A15      c.Long
	Sar      c.Long
	Exccause c.Long
	Excvaddr c.Long
	Lbeg     c.Long
	Lend     c.Long
	Lcount   c.Long
	Tmp0     c.Long
	Tmp1     c.Long
	Tmp2     c.Long
}

/*
-------------------------------------------------------------------------------

	SOLICITED STACK FRAME FOR A THREAD

	A stack frame of this structure is allocated whenever a thread enters the
	RTOS kernel intentionally (and synchronously) to submit to thread scheduling.
	It goes on the current thread's stack.

	The solicited frame only includes registers that are required to be preserved
	by the callee according to the compiler's ABI conventions, some space to save
	the return address for returning to the caller, and the caller's PS register.

	Note: Although the xtensa ABI considers the threadptr as "global" across
	functions (meanig it is neither caller or callee saved), it is treated as a
	callee-saved register in a solicited stack frame. This omits the need for the
	OS to include extra logic to save "global" registers on each context switch.
	Only the threadptr register is treated as callee-saved, as all other NCP
	(non-coprocessor extra) registers are caller-saved. See "tie.h" for more
	details.

	For Windowed ABI, this stack frame includes the caller's base save area.

	Note on XT_SOL_EXIT field:
	    It is necessary to distinguish a solicited from an interrupt stack frame.
	    This field corresponds to XT_STK_EXIT in the interrupt stack frame and is
	    always at the same offset (0). It can be written with a code (usually 0)
	    to distinguish a solicted frame from an interrupt frame. An RTOS port may
	    opt to ignore this field if it has another way of distinguishing frames.

-------------------------------------------------------------------------------
*/
type XtSolFrame struct {
	Exit      c.Long
	Pc        c.Long
	Ps        c.Long
	Threadptr c.Long
	A0        c.Long
	A1        c.Long
	A2        c.Long
	A3        c.Long
}
