package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Context that should be maintained by both the driver and the HAL
 */

type UartHalContextT struct {
	Dev *UartDevT
}

/**
 * @brief  Read data from the UART rxfifo
 *
 * @param[in] hal Context of the HAL layer
 * @param[in] buf Pointer to the buffer used to store the read data. The buffer size should be large than 128 byte
 * @param[inout] inout_rd_len As input, the size of output buffer to read (set to 0 to read all available data).
 *                            As output, returns the actual size written into the output buffer.
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalReadRxfifo C.uart_hal_read_rxfifo
func (recv_ *UartHalContextT) UartHalReadRxfifo(buf *c.Uint8T, inout_rd_len *c.Int) {
}

/**
 * @brief  Write data into the UART txfifo
 *
 * @param  hal Context of the HAL layer
 * @param  buf Pointer of the data buffer need to be written to txfifo
 * @param  data_size The data size(in byte) need to be written
 * @param  write_size The size has been written
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalWriteTxfifo C.uart_hal_write_txfifo
func (recv_ *UartHalContextT) UartHalWriteTxfifo(buf *c.Uint8T, data_size c.Uint32T, write_size *c.Uint32T) {
}

/**
 * @brief  Reset the UART txfifo
 * @note   On ESP32, this function is reserved for UART1 and UART2.
 *
 * @param  hal Context of the HAL layer
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalTxfifoRst C.uart_hal_txfifo_rst
func (recv_ *UartHalContextT) UartHalTxfifoRst() {
}

/**
 * @brief  Reset the UART rxfifo
 *
 * @param  hal Context of the HAL layer
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalRxfifoRst C.uart_hal_rxfifo_rst
func (recv_ *UartHalContextT) UartHalRxfifoRst() {
}

/**
 * @brief Init the UART hal and set the UART to the default configuration.
 *
 * @param  hal Context of the HAL layer
 * @param  uart_num The uart port number, the max port number is (UART_NUM_MAX -1)
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalInit C.uart_hal_init
func (recv_ *UartHalContextT) UartHalInit(uart_num UartPortT) {
}

/**
 * @brief Get the UART source clock type
 *
 * @param  hal Context of the HAL layer
 * @param  sclk The pointer to accept the UART source clock type
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalGetSclk C.uart_hal_get_sclk
func (recv_ *UartHalContextT) UartHalGetSclk(sclk *SocModuleClkT) {
}

/**
 * @brief  Configure the UART stop bit
 *
 * @param  hal Context of the HAL layer
 * @param  stop_bit The stop bit to be set
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetStopBits C.uart_hal_set_stop_bits
func (recv_ *UartHalContextT) UartHalSetStopBits(stop_bit UartStopBitsT) {
}

/**
 * @brief Configure the UART data bit
 *
 * @param  hal Context of the HAL layer
 * @param  data_bit The data bit to be set
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetDataBitNum C.uart_hal_set_data_bit_num
func (recv_ *UartHalContextT) UartHalSetDataBitNum(data_bit UartWordLengthT) {
}

/**
 * @brief Configure the UART parity mode
 *
 * @param  hal Context of the HAL layer
 * @param  parity_mode The UART parity mode to be set
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetParity C.uart_hal_set_parity
func (recv_ *UartHalContextT) UartHalSetParity(parity_mode UartParityT) {
}

/**
 * @brief Configure the UART hardware flow control
 *
 * @param  hal Context of the HAL layer
 * @param  flow_ctrl The flow control mode to be set
 * @param  rx_thresh The rts flow control signal will be active if the data length in rxfifo is large than this value
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetHwFlowCtrl C.uart_hal_set_hw_flow_ctrl
func (recv_ *UartHalContextT) UartHalSetHwFlowCtrl(flow_ctrl UartHwFlowcontrolT, rx_thresh c.Uint8T) {
}

/**
 * @brief Configure the UART AT cmd char detect function. When the receiver receives a continuous AT cmd char, it will produce a interrupt
 *
 * @param  hal Context of the HAL layer
 * @param  at_cmd The AT cmd char detect configuration
 *
 * @return None.
 */
// llgo:link (*UartHalContextT).UartHalSetAtCmdChar C.uart_hal_set_at_cmd_char
func (recv_ *UartHalContextT) UartHalSetAtCmdChar(at_cmd *UartAtCmdT) {
}

/**
 * @brief Set the timeout value of the UART receiver
 *
 * @param  hal Context of the HAL layer
 * @param  tout The timeout value for receiver to receive a data
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetRxTimeout C.uart_hal_set_rx_timeout
func (recv_ *UartHalContextT) UartHalSetRxTimeout(tout c.Uint8T) {
}

/**
 * @brief Set the UART dtr signal active level
 *
 * @param  hal Context of the HAL layer
 * @param  active_level The dtr active level. The active level is low if set to 0. The active level is high if set to 1
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetDtr C.uart_hal_set_dtr
func (recv_ *UartHalContextT) UartHalSetDtr(active_level c.Int) {
}

/**
 * @brief Set the UART software flow control
 *
 * @param  hal Context of the HAL layer
 * @param  flow_ctrl The software flow control configuration
 * @param  sw_flow_ctrl_en Set true to enable the software flow control, otherwise set it false
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetSwFlowCtrl C.uart_hal_set_sw_flow_ctrl
func (recv_ *UartHalContextT) UartHalSetSwFlowCtrl(flow_ctrl *UartSwFlowctrlT, sw_flow_ctrl_en bool) {
}

/**
 * @brief Set the UART tx idle number
 *
 * @param  hal Context of the HAL layer
 * @param  idle_num The cycle number betwin the two transmission
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetTxIdleNum C.uart_hal_set_tx_idle_num
func (recv_ *UartHalContextT) UartHalSetTxIdleNum(idle_num c.Uint16T) {
}

/**
 * @brief Set the UART rxfifo full threshold
 *
 * @param  hal Context of the HAL layer
 * @param  full_thrhd The rxfifo full threshold. If the `UART_RXFIFO_FULL` interrupt is enabled and
 *         the data length in rxfifo is more than this value, it will generate `UART_RXFIFO_FULL` interrupt
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetRxfifoFullThr C.uart_hal_set_rxfifo_full_thr
func (recv_ *UartHalContextT) UartHalSetRxfifoFullThr(full_thrhd c.Uint32T) {
}

/**
 * @brief Set the UART txfifo empty threshold
 *
 * @param  hal Context of the HAL layer
 * @param  empty_thrhd The txfifo empty threshold to be set. If the `UART_TXFIFO_EMPTY` interrupt is enabled and
 *         the data length in txfifo is less than this value, it will generate `UART_TXFIFO_EMPTY` interrupt
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetTxfifoEmptyThr C.uart_hal_set_txfifo_empty_thr
func (recv_ *UartHalContextT) UartHalSetTxfifoEmptyThr(empty_thrhd c.Uint32T) {
}

/**
 * @brief Configure the UART to send a number of break(NULL) chars
 *
 * @param  hal Context of the HAL layer
 * @param  break_num The number of the break char need to be send
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalTxBreak C.uart_hal_tx_break
func (recv_ *UartHalContextT) UartHalTxBreak(break_num c.Uint32T) {
}

/**
 * @brief Configure the UART wake up function.
 *     Note that RXD cannot be input through GPIO Matrix but only through IO_MUX when use this function
 *
 * @param  hal Context of the HAL layer
 * @param  wakeup_thrd The wake up threshold to be set. The system will be woken up from light-sleep when the input RXD edge changes more times than `wakeup_thrd+2`
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetWakeupThrd C.uart_hal_set_wakeup_thrd
func (recv_ *UartHalContextT) UartHalSetWakeupThrd(wakeup_thrd c.Uint32T) {
}

/**
 * @brief Configure the UART mode
 *
 * @param  hal Context of the HAL layer
 * @param  mode The UART mode to be set
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetMode C.uart_hal_set_mode
func (recv_ *UartHalContextT) UartHalSetMode(mode UartModeT) {
}

/**
 * @brief Configure the UART hardware to inverse the signals
 *
 * @param  hal Context of the HAL layer
 * @param  inv_mask The signal mask needs to be inversed. Use the ORred mask of type `uart_signal_inv_t`
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalInverseSignal C.uart_hal_inverse_signal
func (recv_ *UartHalContextT) UartHalInverseSignal(inv_mask c.Uint32T) {
}

/**
 * @brief Get the UART wakeup threshold configuration
 *
 * @param  hal Context of the HAL layer
 * @param  wakeup_thrd Pointer to accept the value of UART wakeup threshold configuration
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalGetWakeupThrd C.uart_hal_get_wakeup_thrd
func (recv_ *UartHalContextT) UartHalGetWakeupThrd(wakeup_thrd *c.Uint32T) {
}

/**
 * @brief Get the UART data bit configuration
 *
 * @param  hal Context of the HAL layer
 * @param  data_bit Pointer to accept the value of UART data bit configuration
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalGetDataBitNum C.uart_hal_get_data_bit_num
func (recv_ *UartHalContextT) UartHalGetDataBitNum(data_bit *UartWordLengthT) {
}

/**
 * @brief Get the UART stop bit configuration
 *
 * @param  hal Context of the HAL layer
 * @param  stop_bit Pointer to accept the value of UART stop bit configuration
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalGetStopBits C.uart_hal_get_stop_bits
func (recv_ *UartHalContextT) UartHalGetStopBits(stop_bit *UartStopBitsT) {
}

/**
 * @brief Get the UART parity mode configuration
 *
 * @param  hal Context of the HAL layer
 * @param  parity_mode Pointer to accept the UART parity mode configuration
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalGetParity C.uart_hal_get_parity
func (recv_ *UartHalContextT) UartHalGetParity(parity_mode *UartParityT) {
}

/**
 * @brief Get the UART baud-rate configuration
 *
 * @param  hal Context of the HAL layer
 * @param  baud_rate Pointer to accept the current baud-rate
 * @param  sclk_freq Frequency of the clock source of UART, in Hz.
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalGetBaudrate C.uart_hal_get_baudrate
func (recv_ *UartHalContextT) UartHalGetBaudrate(baud_rate *c.Uint32T, sclk_freq c.Uint32T) {
}

/**
 * @brief Get the hw flow control configuration
 *
 * @param  hal Context of the HAL layer
 * @param  flow_ctrl Pointer to accept the UART flow control configuration
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalGetHwFlowCtrl C.uart_hal_get_hw_flow_ctrl
func (recv_ *UartHalContextT) UartHalGetHwFlowCtrl(flow_ctrl *UartHwFlowcontrolT) {
}

/**
 * @brief Check if the UART rts flow control is enabled
 *
 * @param  hal Context of the HAL layer
 *
 * @return True if rts flow control is enabled, otherwise false will be returned
 */
// llgo:link (*UartHalContextT).UartHalIsHwRtsEn C.uart_hal_is_hw_rts_en
func (recv_ *UartHalContextT) UartHalIsHwRtsEn() bool {
	return false
}

/**
 * @brief Configure TX signal loop back to RX module, just for the testing purposes
 *
 * @param  hal Context of the HAL layer
 * @param  loop_back_en Set true to enable the loop back function, else set it false.
 *
 * @return None
 */
// llgo:link (*UartHalContextT).UartHalSetLoopBack C.uart_hal_set_loop_back
func (recv_ *UartHalContextT) UartHalSetLoopBack(loop_back_en bool) {
}

/**
 * @brief  Calculate uart symbol bit length, as defined in configuration.
 *
 * @param  hw Beginning address of the peripheral registers.
 *
 * @return number of bits per UART symbol.
 */
// llgo:link (*UartHalContextT).UartHalGetSymbLen C.uart_hal_get_symb_len
func (recv_ *UartHalContextT) UartHalGetSymbLen() c.Uint8T {
	return 0
}

/**
 * @brief  Get UART maximum timeout threshold.
 *
 * @param  hw Beginning address of the peripheral registers.
 *
 * @return maximum timeout threshold value for target.
 */
// llgo:link (*UartHalContextT).UartHalGetMaxRxTimeoutThrd C.uart_hal_get_max_rx_timeout_thrd
func (recv_ *UartHalContextT) UartHalGetMaxRxTimeoutThrd() c.Uint16T {
	return 0
}
