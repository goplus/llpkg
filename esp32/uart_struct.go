package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: FIFO Configuration */
/** Type of fifo register
 *  FIFO data register
 */

type UartFifoRegT struct {
	Val c.Uint32T
}

/** Type of mem_conf register
 *  UART threshold and allocation configuration
 */

type UartMemConfRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
/** Type of int_raw register
 *  Raw interrupt status
 */

type UartIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Masked interrupt status
 */

type UartIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable bits
 */

type UartIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear bits
 */

type UartIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Configuration Register */
/** Type of clkdiv register
 *  Clock divider configuration
 */

type UartClkdivRegT struct {
	Val c.Uint32T
}

/** Type of rx_filt register
 *  Rx Filter configuration
 */

type UartRxFiltRegT struct {
	Val c.Uint32T
}

/** Type of conf0 register
 *  a
 */

type UartConf0RegT struct {
	Val c.Uint32T
}

/** Type of conf1 register
 *  Configuration register 1
 */

type UartConf1RegT struct {
	Val c.Uint32T
}

/** Type of flow_conf register
 *  Software flow-control configuration
 */

type UartFlowConfRegT struct {
	Val c.Uint32T
}

/** Type of sleep_conf register
 *  Sleep-mode configuration
 */

type UartSleepConfRegT struct {
	Val c.Uint32T
}

/** Type of swfc_conf0 register
 *  Software flow-control character configuration
 */

type UartSwfcConf0RegT struct {
	Val c.Uint32T
}

/** Type of swfc_conf1 register
 *  Software flow-control character configuration
 */

type UartSwfcConf1RegT struct {
	Val c.Uint32T
}

/** Type of txbrk_conf register
 *  Tx Break character configuration
 */

type UartTxbrkConfRegT struct {
	Val c.Uint32T
}

/** Type of idle_conf register
 *  Frame-end idle configuration
 */

type UartIdleConfRegT struct {
	Val c.Uint32T
}

/** Type of rs485_conf register
 *  RS485 mode configuration
 */

type UartRs485ConfRegT struct {
	Val c.Uint32T
}

/** Type of clk_conf register
 *  UART core clock configuration
 */

type UartClkConfRegT struct {
	Val c.Uint32T
}

/** Group: Status Register */
/** Type of status register
 *  UART status register
 */

type UartStatusRegT struct {
	Val c.Uint32T
}

/** Type of mem_tx_status register
 *  Tx-FIFO write and read offset address.
 */

type UartMemTxStatusRegT struct {
	Val c.Uint32T
}

/** Type of mem_rx_status register
 *  Rx-FIFO write and read offset address.
 */

type UartMemRxStatusRegT struct {
	Val c.Uint32T
}

/** Type of fsm_status register
 *  UART transmit and receive status.
 */

type UartFsmStatusRegT struct {
	Val c.Uint32T
}

/** Group: Autobaud Register */
/** Type of lowpulse register
 *  Autobaud minimum low pulse duration register
 */

type UartLowpulseRegT struct {
	Val c.Uint32T
}

/** Type of highpulse register
 *  Autobaud minimum high pulse duration register
 */

type UartHighpulseRegT struct {
	Val c.Uint32T
}

/** Type of rxd_cnt register
 *  Autobaud edge change count register
 */

type UartRxdCntRegT struct {
	Val c.Uint32T
}

/** Type of pospulse register
 *  Autobaud high pulse register
 */

type UartPospulseRegT struct {
	Val c.Uint32T
}

/** Type of negpulse register
 *  Autobaud low pulse register
 */

type UartNegpulseRegT struct {
	Val c.Uint32T
}

/** Group: AT Escape Sequence Selection Configuration */
/** Type of at_cmd_precnt register
 *  Pre-sequence timing configuration
 */

type UartAtCmdPrecntRegT struct {
	Val c.Uint32T
}

/** Type of at_cmd_postcnt register
 *  Post-sequence timing configuration
 */

type UartAtCmdPostcntRegT struct {
	Val c.Uint32T
}

/** Type of at_cmd_gaptout register
 *  Timeout configuration
 */

type UartAtCmdGaptoutRegT struct {
	Val c.Uint32T
}

/** Type of at_cmd_char register
 *  AT escape sequence detection configuration
 */

type UartAtCmdCharRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  UART Version register
 */

type UartDateRegT struct {
	Val c.Uint32T
}

/** Type of id register
 *  UART ID register
 */

type UartIdRegT struct {
	Val c.Uint32T
}

type UartDevT struct {
	Fifo         UartFifoRegT
	IntRaw       UartIntRawRegT
	IntSt        UartIntStRegT
	IntEna       UartIntEnaRegT
	IntClr       UartIntClrRegT
	Clkdiv       UartClkdivRegT
	RxFilt       UartRxFiltRegT
	Status       UartStatusRegT
	Conf0        UartConf0RegT
	Conf1        UartConf1RegT
	Lowpulse     UartLowpulseRegT
	Highpulse    UartHighpulseRegT
	RxdCnt       UartRxdCntRegT
	FlowConf     UartFlowConfRegT
	SleepConf    UartSleepConfRegT
	SwfcConf0    UartSwfcConf0RegT
	SwfcConf1    UartSwfcConf1RegT
	TxbrkConf    UartTxbrkConfRegT
	IdleConf     UartIdleConfRegT
	Rs485Conf    UartRs485ConfRegT
	AtCmdPrecnt  UartAtCmdPrecntRegT
	AtCmdPostcnt UartAtCmdPostcntRegT
	AtCmdGaptout UartAtCmdGaptoutRegT
	AtCmdChar    UartAtCmdCharRegT
	MemConf      UartMemConfRegT
	MemTxStatus  UartMemTxStatusRegT
	MemRxStatus  UartMemRxStatusRegT
	FsmStatus    UartFsmStatusRegT
	Pospulse     UartPospulseRegT
	Negpulse     UartNegpulseRegT
	ClkConf      UartClkConfRegT
	Date         UartDateRegT
	Id           UartIdRegT
}
