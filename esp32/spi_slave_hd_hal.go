package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SpiDmaDescT DmaDescriptorAlign4T

/**
 * @brief Type of dma descriptor with appended members
 *        this structure inherits DMA descriptor, with a pointer to the transaction descriptor passed from users.
 */

type SpiSlaveHdHalDescAppendT struct {
	Desc *SpiDmaDescT
	Arg  c.Pointer
}

// / Configuration of the HAL
type SpiSlaveHdHalConfigT struct {
	HostId      c.Uint32T
	DmaEnabled  bool
	AppendMode  bool
	SpicsIoNum  c.Uint32T
	Mode        c.Uint8T
	CommandBits c.Uint32T
	AddressBits c.Uint32T
	DummyBits   c.Uint32T
}

// / Context of the HAL, initialized by :cpp:func:`spi_slave_hd_hal_init`.
type SpiSlaveHdHalContextT struct {
	DmadescTx         *SpiSlaveHdHalDescAppendT
	DmadescRx         *SpiSlaveHdHalDescAppendT
	Dev               *SpiDevT
	DmaEnabled        bool
	AppendMode        bool
	DmaDescNum        c.Uint32T
	CurrentEofAddr    c.Uint32T
	TxCurDesc         *SpiSlaveHdHalDescAppendT
	TxDmaHead         *SpiSlaveHdHalDescAppendT
	TxDmaTail         *SpiSlaveHdHalDescAppendT
	TxUsedDescCnt     c.Uint32T
	TxRecycledDescCnt c.Uint32T
	RxCurDesc         *SpiSlaveHdHalDescAppendT
	RxDmaHead         *SpiSlaveHdHalDescAppendT
	RxDmaTail         *SpiSlaveHdHalDescAppendT
	RxUsedDescCnt     c.Uint32T
	RxRecycledDescCnt c.Uint32T
	IntrNotTriggered  c.Uint32T
	TxDmaStarted      bool
	RxDmaStarted      bool
}

/**
 * @brief Initialize the hardware and part of the context
 *
 * @param hal        Context of the HAL layer
 * @param hal_config Configuration of the HAL
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalInit C.spi_slave_hd_hal_init
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalInit(hal_config *SpiSlaveHdHalConfigT) {
}

/**
 * @brief Check and clear signal of one event
 *
 * @param hal       Context of the HAL layer
 * @param ev        Event to check
 * @return          True if event triggered, otherwise false
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalCheckClearEvent C.spi_slave_hd_hal_check_clear_event
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalCheckClearEvent(ev SpiEventT) bool {
	return false
}

/**
 * @brief Check and clear the interrupt of one event.
 *
 * @note The event source will be kept, so that the interrupt can be invoked by
 *       :cpp:func:`spi_slave_hd_hal_invoke_event_intr`. If event not triggered, its interrupt source
 *       will not be disabled either.
 *
 * @param hal       Context of the HAL layer
 * @param ev        Event to check and disable
 * @return          True if event triggered, otherwise false
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalCheckDisableEvent C.spi_slave_hd_hal_check_disable_event
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalCheckDisableEvent(ev SpiEventT) bool {
	return false
}

/**
 * @brief Enable to involve the ISR of corresponding event.
 *
 * @note The function, compared with :cpp:func:`spi_slave_hd_hal_enable_event_intr`, contains a
 *       workaround to force trigger the interrupt, even if the interrupt source cannot be initialized
 *       correctly.
 *
 * @param hal       Context of the HAL layer
 * @param ev        Event (reason) to invoke the ISR
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalInvokeEventIntr C.spi_slave_hd_hal_invoke_event_intr
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalInvokeEventIntr(ev SpiEventT) {
}

/**
 * @brief Enable the interrupt source of corresponding event.
 *
 * @param hal       Context of the HAL layer
 * @param ev        Event whose corresponding interrupt source should be enabled.
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalEnableEventIntr C.spi_slave_hd_hal_enable_event_intr
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalEnableEventIntr(ev SpiEventT) {
}

////////////////////////////////////////////////////////////////////////////////
// RX DMA
////////////////////////////////////////////////////////////////////////////////
/**
 * @brief Start the RX DMA operation to the specified buffer.
 *
 * @param hal       Context of the HAL layer
 * @param[out] out_buf  Buffer to receive the data
 * @param len       Maximul length to receive
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalRxdma C.spi_slave_hd_hal_rxdma
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalRxdma() {
}

/**
 * @brief Get the length of total received data
 *
 * @param hal       Context of the HAL layer
 * @return          The received length
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalRxdmaSegGetLen C.spi_slave_hd_hal_rxdma_seg_get_len
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalRxdmaSegGetLen() c.Int {
	return 0
}

/**
 * @brief Prepare hardware for a new dma rx trans
 *
 * @param hal       Context of the HAL layer
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalHwPrepareRx C.spi_slave_hd_hal_hw_prepare_rx
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalHwPrepareRx() {
}

////////////////////////////////////////////////////////////////////////////////
// TX DMA
////////////////////////////////////////////////////////////////////////////////
/**
 * @brief Start the TX DMA operation with the specified buffer
 *
 * @param hal       Context of the HAL layer
 * @param data      Buffer of data to send
 * @param len       Size of the buffer, also the maximum length to send
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalTxdma C.spi_slave_hd_hal_txdma
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalTxdma() {
}

/**
 * @brief Prepare hardware for a new dma tx trans
 *
 * @param hal       Context of the HAL layer
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalHwPrepareTx C.spi_slave_hd_hal_hw_prepare_tx
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalHwPrepareTx() {
}

////////////////////////////////////////////////////////////////////////////////
// Shared buffer
////////////////////////////////////////////////////////////////////////////////
/**
 * @brief Read from the shared register buffer
 *
 * @param hal       Context of the HAL layer
 * @param addr      Address of the shared register to read
 * @param out_data  Buffer to store the read data
 * @param len       Length to read from the shared buffer
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalReadBuffer C.spi_slave_hd_hal_read_buffer
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalReadBuffer(addr c.Int, out_data *c.Uint8T, len c.SizeT) {
}

/**
 * @brief Write the shared register buffer
 *
 * @param hal       Context of the HAL layer
 * @param addr      Address of the shared register to write
 * @param data      Buffer of the data to write
 * @param len       Length to write into the shared buffer
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalWriteBuffer C.spi_slave_hd_hal_write_buffer
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalWriteBuffer(addr c.Int, data *c.Uint8T, len c.SizeT) {
}

/**
 * @brief Get the length of previous transaction.
 *
 * @param hal       Context of the HAL layer
 * @return          The length of previous transaction
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalGetRxlen C.spi_slave_hd_hal_get_rxlen
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalGetRxlen() c.Int {
	return 0
}

/**
 * @brief Get the address of last transaction
 *
 * @param hal       Context of the HAL layer
 * @return          The address of last transaction
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalGetLastAddr C.spi_slave_hd_hal_get_last_addr
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalGetLastAddr() c.Int {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// Append Mode
////////////////////////////////////////////////////////////////////////////////
/**
 * @brief Return the finished TX transaction
 *
 * @note This API is based on this assumption: the hardware behaviour of current transaction completion is only modified by the its own caller layer.
 * This means if some other code changed the hardware behaviour (e.g. clear intr raw bit), or the caller call this API without noticing the HW behaviour,
 * this API will go wrong.
 *
 * @param hal            Context of the HAL layer
 * @param out_trans      Pointer to the caller-defined transaction
 * @param real_buff_addr Actually data buffer head the HW used
 * @return               1: Transaction is finished; 0: Transaction is not finished
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalGetTxFinishedTrans C.spi_slave_hd_hal_get_tx_finished_trans
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalGetTxFinishedTrans(out_trans *c.Pointer, real_buff_addr *c.Pointer) bool {
	return false
}

/**
 * @brief Return the finished RX transaction
 *
 * @note This API is based on this assumption: the hardware behaviour of current transaction completion is only modified by the its own caller layer.
 * This means if some other code changed the hardware behaviour (e.g. clear intr raw bit), or the caller call this API without noticing the HW behaviour,
 * this API will go wrong.
 *
 * @param hal            Context of the HAL layer
 * @param out_trans      Pointer to the caller-defined transaction
 * @param real_buff_addr Actually data buffer head the HW used
 * @param out_len        Actual number of bytes of received data
 * @return               1: Transaction is finished; 0: Transaction is not finished
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalGetRxFinishedTrans C.spi_slave_hd_hal_get_rx_finished_trans
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalGetRxFinishedTrans(out_trans *c.Pointer, real_buff_addr *c.Pointer, out_len *c.SizeT) bool {
	return false
}

/**
 * @brief Load the TX DMA descriptors without stopping the DMA
 *
 * @param hal            Context of the HAL layer
 * @param data           Buffer of the transaction data
 * @param len            Length of the data
 * @param arg            Pointer used by the caller to indicate the transaction. Will be returned by ``spi_slave_hd_hal_get_tx_finished_trans`` when transaction is finished
 * @return
 *        - ESP_OK: on success
 *        - ESP_ERR_INVALID_STATE: Function called in invalid state.
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalTxdmaAppend C.spi_slave_hd_hal_txdma_append
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalTxdmaAppend(data *c.Uint8T, len c.SizeT, arg c.Pointer) EspErrT {
	return 0
}

/**
 * @brief Load the RX DMA descriptors without stopping the DMA
 *
 * @param hal            Context of the HAL layer
 * @param data           Buffer of the transaction data
 * @param len            Length of the data
 * @param arg            Pointer used by the caller to indicate the transaction. Will be returned by ``spi_slave_hd_hal_get_rx_finished_trans`` when transaction is finished
 * @return
 *        - ESP_OK: on success
 *        - ESP_ERR_INVALID_STATE: Function called in invalid state.
 */
// llgo:link (*SpiSlaveHdHalContextT).SpiSlaveHdHalRxdmaAppend C.spi_slave_hd_hal_rxdma_append
func (recv_ *SpiSlaveHdHalContextT) SpiSlaveHdHalRxdmaAppend(data *c.Uint8T, len c.SizeT, arg c.Pointer) EspErrT {
	return 0
}
