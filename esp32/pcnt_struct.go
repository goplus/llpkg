package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of un_conf0 register
 *  Configuration register 0 for unit n
 */

type PcntUnConf0RegT struct {
	Val c.Uint32T
}

/** Type of un_conf1 register
 *  Configuration register 1 for unit n
 */

type PcntUnConf1RegT struct {
	Val c.Uint32T
}

/** Type of un_conf2 register
 *  Configuration register 2 for unit n
 */

type PcntUnConf2RegT struct {
	Val c.Uint32T
}

/** Type of ctrl register
 *  Control register for all counters
 */

type PcntCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Status Register */
/** Type of un_cnt register
 *  Counter value for unit n
 */

type PcntUnCntRegT struct {
	Val c.Uint32T
}

/** Type of un_status register
 *  PNCT UNITn status register
 */

type PcntUnStatusRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
/** Type of int_raw register
 *  Interrupt raw status register
 */

type PcntIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Interrupt status register
 */

type PcntIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable register
 */

type PcntIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear register
 */

type PcntIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  PCNT version control register
 */

type PcntDateRegT struct {
	Val c.Uint32T
}

type PcntDevT struct {
	ConfUnit [4]struct {
		Conf0 PcntUnConf0RegT
		Conf1 PcntUnConf1RegT
		Conf2 PcntUnConf2RegT
	}
	CntUnit     [4]PcntUnCntRegT
	IntRaw      PcntIntRawRegT
	IntSt       PcntIntStRegT
	IntEna      PcntIntEnaRegT
	IntClr      PcntIntClrRegT
	StatusUnit  [4]PcntUnStatusRegT
	Ctrl        PcntCtrlRegT
	Reserved064 [38]c.Uint32T
	Date        PcntDateRegT
}
