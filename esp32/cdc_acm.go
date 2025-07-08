package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ACM_BYTES_PER_TX = 64

type CdcAcmDevice c.Void

// llgo:type C
type UartIrqCallbackT func(*CdcAcmDevice, c.Int)

/**
 * @brief Get amount of received characters in buffer
 *
 * @returns character count
 */
// llgo:link (*CdcAcmDevice).CdcAcmRxFifoCnt C.cdc_acm_rx_fifo_cnt
func (recv_ *CdcAcmDevice) CdcAcmRxFifoCnt() c.Int {
	return 0
}

/*
 * @brief Output a character in polled mode.
 *
 * The UART poll method for USB UART is simulated by waiting till
 * we get the next BULK In upcall from the USB device controller or 100 ms.
 *
 * @return the same character which is sent
 */
// llgo:link (*CdcAcmDevice).CdcAcmPollOut C.cdc_acm_poll_out
func (recv_ *CdcAcmDevice) CdcAcmPollOut(c c.Char) c.Char {
	return 0
}

/**
 * @brief Fill FIFO with data
 *
 * @param dev     CDC ACM device struct.
 * @param tx_data Data to transmit.
 * @param len     Number of bytes to send.
 *
 * @return Number of bytes sent.
 */
// llgo:link (*CdcAcmDevice).CdcAcmFifoFill C.cdc_acm_fifo_fill
func (recv_ *CdcAcmDevice) CdcAcmFifoFill(tx_data *c.Uint8T, len c.Int) c.Int {
	return 0
}

/**
 * @brief Read data from FIFO
 *
 * @param dev     CDC ACM device struct.
 * @param rx_data Pointer to data container.
 * @param size    Container size.
 *
 * @return Number of bytes read.
 */
// llgo:link (*CdcAcmDevice).CdcAcmFifoRead C.cdc_acm_fifo_read
func (recv_ *CdcAcmDevice) CdcAcmFifoRead(rx_data *c.Uint8T, size c.Int) c.Int {
	return 0
}

/**
 * @brief Enable TX interrupt
 *
 * @param dev CDC ACM device struct.
 *
 * @return N/A.
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqTxEnable C.cdc_acm_irq_tx_enable
func (recv_ *CdcAcmDevice) CdcAcmIrqTxEnable() {
}

/**
 * @brief Disable TX interrupt
 *
 * @param dev CDC ACM device struct.
 *
 * @return N/A.
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqTxDisable C.cdc_acm_irq_tx_disable
func (recv_ *CdcAcmDevice) CdcAcmIrqTxDisable() {
}

/**
 * @brief Check if Tx IRQ has been raised
 *
 * @param dev CDC ACM device struct.
 *
 * @return 1 if a Tx IRQ is pending, 0 otherwise.
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqTxReady C.cdc_acm_irq_tx_ready
func (recv_ *CdcAcmDevice) CdcAcmIrqTxReady() c.Int {
	return 0
}

/**
 * @brief Enable RX interrupt
 *
 * @param dev CDC ACM device struct.
 *
 * @return N/A
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqRxEnable C.cdc_acm_irq_rx_enable
func (recv_ *CdcAcmDevice) CdcAcmIrqRxEnable() {
}

/**
 * @brief Disable RX interrupt
 *
 * @param dev CDC ACM device struct.
 *
 * @return N/A.
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqRxDisable C.cdc_acm_irq_rx_disable
func (recv_ *CdcAcmDevice) CdcAcmIrqRxDisable() {
}

/**
 * @brief Enable line state interrupt
 *
 * @param dev CDC ACM device struct.
 *
 * @return N/A.
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqStateEnable C.cdc_acm_irq_state_enable
func (recv_ *CdcAcmDevice) CdcAcmIrqStateEnable() {
}

/**
 * @brief Disable line state interrupt
 *
 * @param dev CDC ACM device struct.
 *
 * @return N/A.
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqStateDisable C.cdc_acm_irq_state_disable
func (recv_ *CdcAcmDevice) CdcAcmIrqStateDisable() {
}

/**
 * @brief Check if Rx IRQ has been raised
 *
 * @param dev CDC ACM device struct.
 *
 * @return 1 if an IRQ is ready, 0 otherwise.
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqRxReady C.cdc_acm_irq_rx_ready
func (recv_ *CdcAcmDevice) CdcAcmIrqRxReady() c.Int {
	return 0
}

/**
 * @brief Check if Tx or Rx IRQ is pending
 *
 * @param dev CDC ACM device struct.
 *
 * @return 1 if a Tx or Rx IRQ is pending, 0 otherwise.
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqIsPending C.cdc_acm_irq_is_pending
func (recv_ *CdcAcmDevice) CdcAcmIrqIsPending() c.Int {
	return 0
}

/**
 * @brief Set the callback function pointer for IRQ.
 *
 * @param dev CDC ACM device struct.
 * @param cb  Callback function pointer.
 *
 * @return N/A
 */
// llgo:link (*CdcAcmDevice).CdcAcmIrqCallbackSet C.cdc_acm_irq_callback_set
func (recv_ *CdcAcmDevice) CdcAcmIrqCallbackSet(cb UartIrqCallbackT) {
}

/**
 * @brief Manipulate line control for UART.
 *
 * @param dev CDC ACM device struct
 * @param ctrl The line control to be manipulated
 * @param val Value to set the line control
 *
 * @return 0 if successful, failed otherwise.
 */
// llgo:link (*CdcAcmDevice).CdcAcmLineCtrlSet C.cdc_acm_line_ctrl_set
func (recv_ *CdcAcmDevice) CdcAcmLineCtrlSet(ctrl c.Uint32T, val c.Uint32T) c.Int {
	return 0
}

/**
 * @brief Manipulate line control for UART.
 *
 * @param dev CDC ACM device struct
 * @param ctrl The line control to be manipulated
 * @param val Value to set the line control
 *
 * @return 0 if successful, failed otherwise.
 */
// llgo:link (*CdcAcmDevice).CdcAcmLineCtrlGet C.cdc_acm_line_ctrl_get
func (recv_ *CdcAcmDevice) CdcAcmLineCtrlGet(ctrl c.Uint32T, val *c.Uint32T) c.Int {
	return 0
}

/**
 * @brief Initialize UART channel
 *
 * This routine is called to reset the chip in a quiescent state.
 * It is assumed that this function is called only once per UART.
 *
 * @param mem_chunk Memory chunk to use for internal use
 * @param mem_chunk_size Size of the memory chunk in bytes
 *
 * @return dev or NULL
 */
//go:linkname CdcAcmInit C.cdc_acm_init
func CdcAcmInit(mem_chunk c.Pointer, mem_chunk_size c.Int) *CdcAcmDevice
