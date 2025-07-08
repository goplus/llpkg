package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Context that should be maintained by both the driver and the HAL.
 */

type SpiSlaveHalContextT struct {
	Hw        *SpiDevT
	DmadescRx *SpiDmaDescT
	DmadescTx *SpiDmaDescT
	DmadescN  c.Int
	Mode      c.Int
	Bitlen    c.Uint32T
	TxBuffer  c.Pointer
	RxBuffer  c.Pointer
	RcvBitlen c.Uint32T
}

type SpiSlaveHalConfigT struct {
	HostId c.Uint32T
}

/**
 * Init the peripheral and the context.
 *
 * @param hal        Context of the HAL layer.
 * @param hal_config Configuration of the HAL
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalInit C.spi_slave_hal_init
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalInit(hal_config *SpiSlaveHalConfigT) {
}

/**
 * Deinit the peripheral (and the context if needed).
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalDeinit C.spi_slave_hal_deinit
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalDeinit() {
}

/**
 * Setup device-related configurations according to the settings in the context.
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalSetupDevice C.spi_slave_hal_setup_device
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalSetupDevice() {
}

/**
 * Prepare rx hardware for a new DMA trans
 *
 * @param hw Beginning address of the peripheral registers.
 */
// llgo:link (*SpiDevT).SpiSlaveHalHwPrepareRx C.spi_slave_hal_hw_prepare_rx
func (recv_ *SpiDevT) SpiSlaveHalHwPrepareRx() {
}

/**
 * Prepare tx hardware for a new DMA trans
 *
 * @param hw Beginning address of the peripheral registers.
 */
// llgo:link (*SpiDevT).SpiSlaveHalHwPrepareTx C.spi_slave_hal_hw_prepare_tx
func (recv_ *SpiDevT) SpiSlaveHalHwPrepareTx() {
}

/**
 * Rest peripheral registers to default value
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalHwReset C.spi_slave_hal_hw_reset
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalHwReset() {
}

/**
 * Rest hw fifo in peripheral, for a CPU controlled trans
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalHwFifoReset C.spi_slave_hal_hw_fifo_reset
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalHwFifoReset(tx_rst bool, rx_rst bool) {
}

/**
 * Push data needed to be transmit into hw fifo
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalPushTxBuffer C.spi_slave_hal_push_tx_buffer
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalPushTxBuffer() {
}

/**
 * Config transaction bit length for slave
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalSetTransBitlen C.spi_slave_hal_set_trans_bitlen
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalSetTransBitlen() {
}

/**
 * Enable/Disable miso/mosi signals in peripheral
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalEnableDataLine C.spi_slave_hal_enable_data_line
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalEnableDataLine() {
}

/**
 * Trigger start a user-defined transaction.
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalUserStart C.spi_slave_hal_user_start
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalUserStart() {
}

/**
 * Check whether the transaction is done (trans_done is set).
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalUsrIsDone C.spi_slave_hal_usr_is_done
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalUsrIsDone() bool {
	return false
}

/**
 * Post transaction operations, fetch data from the buffer and recorded the length.
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalStoreResult C.spi_slave_hal_store_result
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalStoreResult() {
}

/**
 * Get the length of last transaction, in bits. Should be called after ``spi_slave_hal_store_result``.
 *
 * Note that if last transaction is longer than configured before, the return
 * value still the actual length.
 *
 * @param hal Context of the HAL layer.
 *
 * @return Length of the last transaction, in bits.
 */
// llgo:link (*SpiSlaveHalContextT).SpiSlaveHalGetRcvBitlen C.spi_slave_hal_get_rcv_bitlen
func (recv_ *SpiSlaveHalContextT) SpiSlaveHalGetRcvBitlen() c.Uint32T {
	return 0
}
