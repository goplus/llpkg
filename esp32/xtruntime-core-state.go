package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CORE_STATE_SIGNATURE = 0xB1C5AFED

/*
 *  Save area for saving entire core state, such as across Power Shut-Off (PSO).
 */

type XtosCoreState struct {
	Signature       c.Long
	RestoreLabel    c.Long
	AftersaveLabel  c.Long
	Areg            [64]c.Long
	CallerRegs      [16]c.Long
	CallerRegsSaved c.Long
	Windowbase      c.Long
	Windowstart     c.Long
	Sar             c.Long
	Epc1            c.Long
	Ps              c.Long
	Excsave1        c.Long
	Depc            c.Long
	Epc             [6]c.Long
	Eps             [6]c.Long
	Excsave         [6]c.Long
	Lcount          c.Long
	Lbeg            c.Long
	Lend            c.Long
	Vecbase         c.Long
	Atomctl         c.Long
	Memctl          c.Long
	Ccount          c.Long
	Ccompare        [3]c.Long
	Intenable       c.Long
	Interrupt       c.Long
	Icount          c.Long
	Icountlevel     c.Long
	Debugcause      c.Long
	Dbreakc         [2]c.Long
	Dbreaka         [2]c.Long
	Ibreaka         [2]c.Long
	Ibreakenable    c.Long
	Misc            [4]c.Long
	Cpenable        c.Long
	Tlbs            [16]c.Long
	Ncp             [36]c.Char
	Cp0             [72]c.Char
	Cp3             [208]c.Char
}
