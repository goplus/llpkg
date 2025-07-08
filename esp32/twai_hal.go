package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TwaiHalFrameT TwaiLlFrameBufferT

type TwaiHalContextT struct {
	Dev           *TwaiDevT
	StateFlags    c.Uint32T
	ClockSourceHz c.Uint32T
}

/* ---------------------------- Init and Config ----------------------------- */

type TwaiHalConfigT struct {
	ControllerId  c.Int
	ClockSourceHz c.Uint32T
}

/**
 * @brief Initialize TWAI peripheral and HAL context
 *
 * Sets HAL context, puts TWAI peripheral into reset mode, then sets some
 * registers with default values.
 *
 * @param hal_ctx Context of the HAL layer
 * @param config HAL driver configuration
 * @return True if successfully initialized, false otherwise.
 */
// llgo:link (*TwaiHalContextT).TwaiHalInit C.twai_hal_init
func (recv_ *TwaiHalContextT) TwaiHalInit(config *TwaiHalConfigT) bool {
	return false
}

/**
 * @brief Deinitialize the TWAI peripheral and HAL context
 *
 * Clears any unhandled interrupts and unsets HAL context
 *
 * @param hal_ctx Context of the HAL layer
 */
// llgo:link (*TwaiHalContextT).TwaiHalDeinit C.twai_hal_deinit
func (recv_ *TwaiHalContextT) TwaiHalDeinit() {
}

/**
 * @brief Configure the TWAI peripheral
 *
 * @param hal_ctx Context of the HAL layer
 * @param t_config Pointer to timing configuration structure
 * @param f_config Pointer to filter configuration structure
 * @param intr_mask Mask of interrupts to enable
 * @param clkout_divider Clock divider value for CLKOUT. Set to -1 to disable CLKOUT
 */
// llgo:link (*TwaiHalContextT).TwaiHalConfigure C.twai_hal_configure
func (recv_ *TwaiHalContextT) TwaiHalConfigure(t_config *TwaiTimingConfigT, f_config *TwaiFilterConfigT, intr_mask c.Uint32T, clkout_divider c.Uint32T) {
}

/**
 * @brief Start the TWAI peripheral
 *
 * Start the TWAI peripheral by configuring its operating mode, then exiting
 * reset mode so that the TWAI peripheral can participate in bus activities.
 *
 * @param hal_ctx Context of the HAL layer
 * @param mode Operating mode
 */
// llgo:link (*TwaiHalContextT).TwaiHalStart C.twai_hal_start
func (recv_ *TwaiHalContextT) TwaiHalStart(mode TwaiModeT) {
}

/**
 * @brief Stop the TWAI peripheral
 *
 * Stop the TWAI peripheral by entering reset mode to stop any bus activity, then
 * setting the operating mode to Listen Only so that REC is frozen.
 *
 * @param hal_ctx Context of the HAL layer
 */
// llgo:link (*TwaiHalContextT).TwaiHalStop C.twai_hal_stop
func (recv_ *TwaiHalContextT) TwaiHalStop() {
}

/**
 * @brief Get a bit mask of the events that triggered that triggered an interrupt
 *
 * This function should be called at the beginning of an interrupt. This function will do the following:
 * - Read and clear interrupt register
 * - Calculate what events have triggered the interrupt
 * - Respond to low latency interrupt events
 *      - Bus off: Change to LOM to freeze TEC/REC. Errata 1 Fix
 *      - Recovery complete: Enter reset mode
 *      - Clear ECC and ALC so that their interrupts are re-armed
 * - Update HAL state flags based on interrupts that have occurred.
 * - For the ESP32, check for errata conditions. If a HW reset is required, this function
 *   will set the TWAI_HAL_EVENT_NEED_PERIPH_RESET event.
 *
 * @param hal_ctx Context of the HAL layer
 * @return Bit mask of events that have occurred
 */
// llgo:link (*TwaiHalContextT).TwaiHalGetEvents C.twai_hal_get_events
func (recv_ *TwaiHalContextT) TwaiHalGetEvents() c.Uint32T {
	return 0
}

/**
 * @brief Copy a frame into the TX buffer and transmit
 *
 * This function copies a formatted TX frame into the TX buffer, and the
 * transmit by setting the correct transmit command (e.g. normal, single shot,
 * self RX) in the command register.
 *
 * @param hal_ctx Context of the HAL layer
 * @param tx_frame Pointer to structure containing formatted TX frame
 */
// llgo:link (*TwaiHalContextT).TwaiHalSetTxBufferAndTransmit C.twai_hal_set_tx_buffer_and_transmit
func (recv_ *TwaiHalContextT) TwaiHalSetTxBufferAndTransmit(tx_frame *TwaiHalFrameT) {
}
