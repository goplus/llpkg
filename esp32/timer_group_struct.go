package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration and control registers */
/** Type of tnconfig register
 *  Timer n configuration register
 */

type TimgTnconfigRegT struct {
	Val c.Uint32T
}

/** Type of tnlo register
 *  Timer n current value, low 32 bits
 */

type TimgTnloRegT struct {
	Val c.Uint32T
}

/** Type of tnhi register
 *  Timer n current value, high 22 bits
 */

type TimgTnhiRegT struct {
	Val c.Uint32T
}

/** Type of tnupdate register
 *  Write to copy current timer value to TIMGn_Tn_(LO/HI)_REG
 */

type TimgTnupdateRegT struct {
	Val c.Uint32T
}

/** Type of tnalarmlo register
 *  Timer n alarm value, low 32 bits
 */

type TimgTnalarmloRegT struct {
	Val c.Uint32T
}

/** Type of tnalarmhi register
 *  Timer n alarm value, high bits
 */

type TimgTnalarmhiRegT struct {
	Val c.Uint32T
}

/** Type of tnloadlo register
 *  Timer n reload value, low 32 bits
 */

type TimgTnloadloRegT struct {
	Val c.Uint32T
}

/** Type of tnloadhi register
 *  Timer n reload value, high 22 bits
 */

type TimgTnloadhiRegT struct {
	Val c.Uint32T
}

/** Type of tnload register
 *  Write to reload timer from TIMG_Tn_(LOADLOLOADHI)_REG
 */

type TimgTnloadRegT struct {
	Val c.Uint32T
}

/** Group: Configuration and control registers for WDT */
/** Type of wdtconfig0 register
 *  Watchdog timer configuration register
 */

type TimgWdtconfig0RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig1 register
 *  Watchdog timer prescaler register
 */

type TimgWdtconfig1RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig2 register
 *  Watchdog timer stage 0 timeout value
 */

type TimgWdtconfig2RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig3 register
 *  Watchdog timer stage 1 timeout value
 */

type TimgWdtconfig3RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig4 register
 *  Watchdog timer stage 2 timeout value
 */

type TimgWdtconfig4RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig5 register
 *  Watchdog timer stage 3 timeout value
 */

type TimgWdtconfig5RegT struct {
	Val c.Uint32T
}

/** Type of wdtfeed register
 *  Write to feed the watchdog timer
 */

type TimgWdtfeedRegT struct {
	Val c.Uint32T
}

/** Type of wdtwprotect register
 *  Watchdog write protect register
 */

type TimgWdtwprotectRegT struct {
	Val c.Uint32T
}

/** Group: Configuration and control registers for RTC CALI */
/** Type of rtccalicfg register
 *  RTC calibration configure register
 */

type TimgRtccalicfgRegT struct {
	Val c.Uint32T
}

/** Type of rtccalicfg1 register
 *  RTC calibration configure1 register
 */

type TimgRtccalicfg1RegT struct {
	Val c.Uint32T
}

/** Type of rtccalicfg2 register
 *  Timer group calibration register
 */

type TimgRtccalicfg2RegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of int_ena_timers register
 *  Interrupt enable bits
 */

type TimgIntEnaTimersRegT struct {
	Val c.Uint32T
}

/** Type of int_raw_timers register
 *  Raw interrupt status
 */

type TimgIntRawTimersRegT struct {
	Val c.Uint32T
}

/** Type of int_st_timers register
 *  Masked interrupt status
 */

type TimgIntStTimersRegT struct {
	Val c.Uint32T
}

/** Type of int_clr_timers register
 *  Interrupt clear bits
 */

type TimgIntClrTimersRegT struct {
	Val c.Uint32T
}

/** Group: Configuration registers */
/** Type of ntimers_date register
 *  Timer version control register
 */

type TimgNtimersDateRegT struct {
	Val c.Uint32T
}

/** Type of regclk register
 *  Timer group clock gate register
 */

type TimgRegclkRegT struct {
	Val c.Uint32T
}

type TimgHwtimerRegT struct {
	Config  TimgTnconfigRegT
	Lo      TimgTnloRegT
	Hi      TimgTnhiRegT
	Update  TimgTnupdateRegT
	Alarmlo TimgTnalarmloRegT
	Alarmhi TimgTnalarmhiRegT
	Loadlo  TimgTnloadloRegT
	Loadhi  TimgTnloadhiRegT
	Load    TimgTnloadRegT
}

type TimgDevT struct {
	HwTimer      [2]TimgHwtimerRegT
	Wdtconfig0   TimgWdtconfig0RegT
	Wdtconfig1   TimgWdtconfig1RegT
	Wdtconfig2   TimgWdtconfig2RegT
	Wdtconfig3   TimgWdtconfig3RegT
	Wdtconfig4   TimgWdtconfig4RegT
	Wdtconfig5   TimgWdtconfig5RegT
	Wdtfeed      TimgWdtfeedRegT
	Wdtwprotect  TimgWdtwprotectRegT
	Rtccalicfg   TimgRtccalicfgRegT
	Rtccalicfg1  TimgRtccalicfg1RegT
	IntEnaTimers TimgIntEnaTimersRegT
	IntRawTimers TimgIntRawTimersRegT
	IntStTimers  TimgIntStTimersRegT
	IntClrTimers TimgIntClrTimersRegT
	Rtccalicfg2  TimgRtccalicfg2RegT
	Reserved084  [29]c.Uint32T
	NtimersDate  TimgNtimersDateRegT
	Regclk       TimgRegclkRegT
}
