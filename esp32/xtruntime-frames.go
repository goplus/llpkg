package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CALL0_ABI = 0
const KERNELSTACKSIZE = 1024

/*
 *  Kernel vector mode exception stack frame.
 *
 *  NOTE:  due to the limited range of addi used in the current
 *  kernel exception vector, and the fact that historically
 *  the vector is limited to 12 bytes, the size of this
 *  stack frame is limited to 128 bytes (currently at 64).
 */

type KernelFrame struct {
	Pc     c.Long
	Ps     c.Long
	Areg   [4]c.Long
	Sar    c.Long
	Lcount c.Long
	Lbeg   c.Long
	Lend   c.Long
	Acclo  c.Long
	Acchi  c.Long
	Mr     [4]c.Long
}

/*
 *  User vector mode exception stack frame:
 *
 *  WARNING:  if you modify this structure, you MUST modify the
 *  computation of the pad size (ALIGNPAD) accordingly.
 */

type UserFrame struct {
	Pc       c.Long
	Ps       c.Long
	Sar      c.Long
	Vpri     c.Long
	A2       c.Long
	A3       c.Long
	A4       c.Long
	A5       c.Long
	Exccause c.Long
	Lcount   c.Long
	Lbeg     c.Long
	Lend     c.Long
	Acclo    c.Long
	Acchi    c.Long
	Mr       [4]c.Long
	Pad      [2]c.Long
}
