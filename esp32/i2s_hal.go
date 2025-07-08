package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief General slot configuration information
 * @note It is a general purpose struct, not supposed to be used directly by user
 */

type I2sHalSlotConfigT struct {
	DataBitWidth I2sDataBitWidthT
	SlotBitWidth I2sSlotBitWidthT
	SlotMode     I2sSlotModeT
}

/**
 * @brief I2S clock configuration
 */

type I2sHalClockInfoT struct {
	Sclk    c.Uint32T
	Mclk    c.Uint32T
	Bclk    c.Uint32T
	MclkDiv c.Uint16T
	BclkDiv c.Uint16T
}

/**
 * Context that should be maintained by both the driver and the HAL
 */

type I2sHalContextT struct {
	Dev *I2sDevT
}

/**
 * @brief Init I2S hal context
 *
 * @param hal Context of the HAL layer
 * @param port_id The I2S port number, the max port number is (SOC_I2S_NUM -1)
 */
// llgo:link (*I2sHalContextT).I2sHalInit C.i2s_hal_init
func (recv_ *I2sHalContextT) I2sHalInit(port_id c.Int) {
}

/**
 * @brief Helper function for calculating the precise mclk division by sclk and mclk
 *
 * @param sclk      system clock
 * @param mclk      module clock
 * @param mclk_div  mclk division coefficients, including integer part and decimal part
 */
//go:linkname I2sHalCalcMclkPreciseDivision C.i2s_hal_calc_mclk_precise_division
func I2sHalCalcMclkPreciseDivision(sclk c.Uint32T, mclk c.Uint32T, mclk_div *HalUtilsClkDivT)

/**
 * @brief Set tx channel clock
 *
 * @param hal Context of the HAL layer
 * @param clk_info clock information, if it is NULL, only set the clock source
 * @param clk_src clock source
 */
// llgo:link (*I2sHalContextT).I2sHalSetTxClock C.i2s_hal_set_tx_clock
func (recv_ *I2sHalContextT) I2sHalSetTxClock(clk_info *I2sHalClockInfoT, clk_src I2sClockSrcT) {
}

/**
 * @brief Set rx channel clock
 *
 * @param hal Context of the HAL layer
 * @param clk_info clock information, if it is NULL, only set the clock source
 * @param clk_src clock source
 */
// llgo:link (*I2sHalContextT).X_i2sHalSetRxClock C._i2s_hal_set_rx_clock
func (recv_ *I2sHalContextT) X_i2sHalSetRxClock(clk_info *I2sHalClockInfoT, clk_src I2sClockSrcT) {
}

/*-------------------------------------------------------------------------
|                           STD configuration                            |
-------------------------------------------------------------------------*/
/**
 * @brief Set tx slot to standard mode
 *
 * @param hal Context of the HAL layer
 * @param is_slave If is slave role
 * @param slot_config General slot configuration pointer, but will specified to i2s standard mode
 */
// llgo:link (*I2sHalContextT).I2sHalStdSetTxSlot C.i2s_hal_std_set_tx_slot
func (recv_ *I2sHalContextT) I2sHalStdSetTxSlot(is_slave bool, slot_cfg *I2sHalSlotConfigT) {
}

/**
 * @brief Set rx slot to standard mode
 *
 * @param hal Context of the HAL layer
 * @param is_slave If is slave role
 * @param slot_config General slot configuration pointer, but will specified to i2s standard mode
 */
// llgo:link (*I2sHalContextT).I2sHalStdSetRxSlot C.i2s_hal_std_set_rx_slot
func (recv_ *I2sHalContextT) I2sHalStdSetRxSlot(is_slave bool, slot_cfg *I2sHalSlotConfigT) {
}

/**
 * @brief Enable tx channel as standard mode
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*I2sHalContextT).I2sHalStdEnableTxChannel C.i2s_hal_std_enable_tx_channel
func (recv_ *I2sHalContextT) I2sHalStdEnableTxChannel() {
}

/**
 * @brief Enable rx channel as standard mode
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*I2sHalContextT).I2sHalStdEnableRxChannel C.i2s_hal_std_enable_rx_channel
func (recv_ *I2sHalContextT) I2sHalStdEnableRxChannel() {
}

/**
 * @brief Set tx slot to pdm mode
 *
 * @param hal Context of the HAL layer
 * @param is_slave If is slave role
 * @param slot_config General slot configuration pointer, but will specified to i2s pdm tx mode
 */
// llgo:link (*I2sHalContextT).I2sHalPdmSetTxSlot C.i2s_hal_pdm_set_tx_slot
func (recv_ *I2sHalContextT) I2sHalPdmSetTxSlot(is_slave bool, slot_cfg *I2sHalSlotConfigT) {
}

/**
 * @brief Enable tx channel as pdm mode
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*I2sHalContextT).I2sHalPdmEnableTxChannel C.i2s_hal_pdm_enable_tx_channel
func (recv_ *I2sHalContextT) I2sHalPdmEnableTxChannel() {
}

/**
 * @brief Set rx slot to pdm mode
 *
 * @param hal Context of the HAL layer
 * @param is_slave If is slave role
 * @param slot_config General slot configuration pointer, but will specified to i2s pdm rx mode
 */
// llgo:link (*I2sHalContextT).I2sHalPdmSetRxSlot C.i2s_hal_pdm_set_rx_slot
func (recv_ *I2sHalContextT) I2sHalPdmSetRxSlot(is_slave bool, slot_cfg *I2sHalSlotConfigT) {
}

/**
 * @brief Enable rx channel as pdm mode
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*I2sHalContextT).I2sHalPdmEnableRxChannel C.i2s_hal_pdm_enable_rx_channel
func (recv_ *I2sHalContextT) I2sHalPdmEnableRxChannel() {
}

/**
 * @brief Set tx slot to tdm mode
 *
 * @param hal Context of the HAL layer
 * @param is_slave If is slave role
 * @param slot_config General slot configuration pointer, but will specified to i2s tdm mode
 */
// llgo:link (*I2sHalContextT).I2sHalTdmSetTxSlot C.i2s_hal_tdm_set_tx_slot
func (recv_ *I2sHalContextT) I2sHalTdmSetTxSlot(is_slave bool, slot_cfg *I2sHalSlotConfigT) {
}

/**
 * @brief Set rx slot to tdm mode
 *
 * @param hal Context of the HAL layer
 * @param is_slave If is slave role
 * @param slot_config General slot configuration pointer, but will specified to i2s tdm mode
 */
// llgo:link (*I2sHalContextT).I2sHalTdmSetRxSlot C.i2s_hal_tdm_set_rx_slot
func (recv_ *I2sHalContextT) I2sHalTdmSetRxSlot(is_slave bool, slot_cfg *I2sHalSlotConfigT) {
}

/**
 * @brief Enable tx channel as tdm mode
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*I2sHalContextT).I2sHalTdmEnableTxChannel C.i2s_hal_tdm_enable_tx_channel
func (recv_ *I2sHalContextT) I2sHalTdmEnableTxChannel() {
}

/**
 * @brief Enable rx channel as tdm mode
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*I2sHalContextT).I2sHalTdmEnableRxChannel C.i2s_hal_tdm_enable_rx_channel
func (recv_ *I2sHalContextT) I2sHalTdmEnableRxChannel() {
}
