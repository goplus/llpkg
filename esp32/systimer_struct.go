package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** SYSTEM TIMER CLK CONTROL REGISTER */
/** Type of conf register
 *  Configure system timer clock
 */

type SystimerConfRegT struct {
	Val c.Uint32T
}

/** Type of unit_op register
 *  SYSTIMER_UNIT_OP.
 */

type SystimerUnitOpRegT struct {
	Val c.Uint32T
}

/** Type of unit_load register
 *  SYSTIMER_UNIT_LOAD
 */

type SystimerUnitLoadValRegT struct {
	Hi struct {
		Val c.Uint32T
	}
	Lo struct {
		Val c.Uint32T
	}
}

/** Type of target register
 *  SYSTIMER_TARGET.
 */

type SystimerTargetValRegT struct {
	Hi struct {
		Val c.Uint32T
	}
	Lo struct {
		Val c.Uint32T
	}
}

/** Type of target_conf register
 *  SYSTIMER_TARGET_CONF.
 */

type SystimerTargetConfRegT struct {
	Val c.Uint32T
}

/** Type of unit_value_hi register
 *  SYSTIMER_UNIT_VALUE_HI.
 */

type SystimerUnitValueRegT struct {
	Hi struct {
		Val c.Uint32T
	}
	Lo struct {
		Val c.Uint32T
	}
}

/** Type of comp_load register
 *  SYSTIMER_COMP_LOAD.
 */

type SystimerCompLoadRegT struct {
	Val c.Uint32T
}

/** Type of unit_load register
 *  SYSTIMER_UNIT_LOAD.
 */

type SystimerUnitLoadRegT struct {
	Val c.Uint32T
}

/** SYSTEM TIMER INTERRUPT REGISTER */
/** Type of int_ena register
 *  systimer interrupt enable register
 */

type SystimerIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_raw register
 *  systimer interrupt raw register
 */

type SystimerIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  systimer interrupt clear register
 */

type SystimerIntClrRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  systimer interrupt status register
 */

type SystimerIntStRegT struct {
	Val c.Uint32T
}

/** SYSTEM TIMER COMP STATUS REGISTER
 *  systimer comp actual target value low register
 */

type SystimerRealTargetValRegT struct {
	Lo struct {
		Val c.Uint32T
	}
	Hi struct {
		Val c.Uint32T
	}
}

/** VERSION REGISTER */
/** Type of date register
 *  system timer version control register
 */

type SystimerDateRegT struct {
	Val c.Uint32T
}

type SystimerDevT struct {
	Conf          SystimerConfRegT
	UnitOp        [2]SystimerUnitOpRegT
	UnitLoadVal   [2]SystimerUnitLoadValRegT
	TargetVal     [3]SystimerTargetValRegT
	TargetConf    [3]SystimerTargetConfRegT
	UnitVal       [2]SystimerUnitValueRegT
	CompLoad      [3]SystimerCompLoadRegT
	UnitLoad      [2]SystimerUnitLoadRegT
	IntEna        SystimerIntEnaRegT
	IntRaw        SystimerIntRawRegT
	IntClr        SystimerIntClrRegT
	IntSt         SystimerIntStRegT
	RealTargetVal [3]SystimerRealTargetValRegT
	Reserved08c   c.Uint32T
	Reserved090   c.Uint32T
	Reserved094   c.Uint32T
	Reserved098   c.Uint32T
	Reserved09c   c.Uint32T
	Reserved0a0   c.Uint32T
	Reserved0a4   c.Uint32T
	Reserved0a8   c.Uint32T
	Reserved0ac   c.Uint32T
	Reserved0b0   c.Uint32T
	Reserved0b4   c.Uint32T
	Reserved0b8   c.Uint32T
	Reserved0bc   c.Uint32T
	Reserved0c0   c.Uint32T
	Reserved0c4   c.Uint32T
	Reserved0c8   c.Uint32T
	Reserved0cc   c.Uint32T
	Reserved0d0   c.Uint32T
	Reserved0d4   c.Uint32T
	Reserved0d8   c.Uint32T
	Reserved0dc   c.Uint32T
	Reserved0e0   c.Uint32T
	Reserved0e4   c.Uint32T
	Reserved0e8   c.Uint32T
	Reserved0ec   c.Uint32T
	Reserved0f0   c.Uint32T
	Reserved0f4   c.Uint32T
	Reserved0f8   c.Uint32T
	Date          SystimerDateRegT
}
