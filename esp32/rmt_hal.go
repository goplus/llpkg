package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RmtSocHandleT *RmtDevT

/**
 * @brief HAL context type of RMT driver
 */

type RmtHalContextT struct {
	Regs RmtSocHandleT
}

/**
 * @brief Initialize the RMT HAL driver
 *
 * @param hal: RMT HAL context
 */
// llgo:link (*RmtHalContextT).RmtHalInit C.rmt_hal_init
func (recv_ *RmtHalContextT) RmtHalInit() {
}

/**
 * @brief Deinitialize the RMT HAL driver
 *
 * @param hal: RMT HAL context
 */
// llgo:link (*RmtHalContextT).RmtHalDeinit C.rmt_hal_deinit
func (recv_ *RmtHalContextT) RmtHalDeinit() {
}

/**
 * @brief Reset RMT TX Channel
 *
 * @param hal: RMT HAL context
 * @param channel: RMT channel number
 */
// llgo:link (*RmtHalContextT).RmtHalTxChannelReset C.rmt_hal_tx_channel_reset
func (recv_ *RmtHalContextT) RmtHalTxChannelReset(channel c.Uint32T) {
}

/**
 * @brief Reset RMT TX Channel
 *
 * @param hal: RMT HAL context
 * @param channel: RMT channel number
 */
// llgo:link (*RmtHalContextT).RmtHalRxChannelReset C.rmt_hal_rx_channel_reset
func (recv_ *RmtHalContextT) RmtHalRxChannelReset(channel c.Uint32T) {
}
