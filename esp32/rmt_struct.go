package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: FIFO R/W registers */
/** Type of chndata register
 *  The read and write  data register for CHANNELn by apb fifo access.
 */

type RmtChndataRegT struct {
	Val c.Uint32T
}

/** Type of chmdata register
 *  The read and write  data register for CHANNEL$n by apb fifo access.
 */

type RmtChmdataRegT struct {
	Val c.Uint32T
}

/** Group: Configuration registers */
/** Type of chnconf0 register
 *  Channel n configure register 0
 */

type RmtChnconf0RegT struct {
	Val c.Uint32T
}

/** Type of chmconf0 register
 *  Channel m configure register 0
 */

type RmtChmconf0RegT struct {
	Val c.Uint32T
}

/** Type of chmconf1 register
 *  Channel m configure register 1
 */

type RmtChmconf1RegT struct {
	Val c.Uint32T
}

/** Type of chm_rx_carrier_rm register
 *  Channel m carrier remove register
 */

type RmtChmRxCarrierRmRegT struct {
	Val c.Uint32T
}

/** Type of sys_conf register
 *  RMT apb configuration register
 */

type RmtSysConfRegT struct {
	Val c.Uint32T
}

/** Type of ref_cnt_rst register
 *  RMT clock divider reset register
 */

type RmtRefCntRstRegT struct {
	Val c.Uint32T
}

/** Group: Status registers */
/** Type of chnstatus register
 *  Channel n status register
 */

type RmtChnstatusRegT struct {
	Val c.Uint32T
}

/** Type of chmstatus register
 *  Channel m status register
 */

type RmtChmstatusRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of int_raw register
 *  Raw interrupt status
 */

type RmtIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Masked interrupt status
 */

type RmtIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable bits
 */

type RmtIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear bits
 */

type RmtIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Carrier wave duty cycle registers */
/** Type of chncarrier_duty register
 *  Channel n duty cycle configuration register
 */

type RmtChncarrierDutyRegT struct {
	Val c.Uint32T
}

/** Group: Tx event configuration registers */
/** Type of chn_tx_lim register
 *  Channel n Tx event configuration register
 */

type RmtChnTxLimRegT struct {
	Val c.Uint32T
}

/** Type of tx_sim register
 *  RMT TX synchronous register
 */

type RmtTxSimRegT struct {
	Val c.Uint32T
}

/** Group: Rx event configuration registers */
/** Type of chm_rx_lim register
 *  Channel m Rx event configuration register
 */

type RmtChmRxLimRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  RMT version register
 */

type RmtDateRegT struct {
	Val c.Uint32T
}

type RmtDevT struct {
	Chndata  [4]RmtChndataRegT
	Chmdata  [4]RmtChmdataRegT
	Chnconf0 [4]RmtChnconf0RegT
	Chmconf  [4]struct {
		Conf0 RmtChmconf0RegT
		Conf1 RmtChmconf1RegT
	}
	Chnstatus      [4]RmtChnstatusRegT
	Chmstatus      [4]RmtChmstatusRegT
	IntRaw         RmtIntRawRegT
	IntSt          RmtIntStRegT
	IntEna         RmtIntEnaRegT
	IntClr         RmtIntClrRegT
	ChncarrierDuty [4]RmtChncarrierDutyRegT
	ChmRxCarrierRm [4]RmtChmRxCarrierRmRegT
	ChnTxLim       [4]RmtChnTxLimRegT
	ChmRxLim       [4]RmtChmRxLimRegT
	SysConf        RmtSysConfRegT
	TxSim          RmtTxSimRegT
	RefCntRst      RmtRefCntRstRegT
	Date           RmtDateRegT
}
