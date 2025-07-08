package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type Dma2dDevT struct {
	Unused [8]uint8
}
type Dma2dSocHandleT *Dma2dDevT

/**
 * Context that should be maintained by both the driver and the HAL
 */

type Dma2dHalContextT struct {
	Dev Dma2dSocHandleT
}

/**
 * @brief Init the 2D-DMA hal. This function should be called first before other hal layer function is called
 *
 * @param hal Context of the HAL layer
 * @param group_id The 2D-DMA group number
 */
// llgo:link (*Dma2dHalContextT).Dma2dHalInit C.dma2d_hal_init
func (recv_ *Dma2dHalContextT) Dma2dHalInit(group_id c.Int) {
}

/**
 * @brief Reset 2D-DMA TX channel
 *
 * @param hal Context of the HAL layer
 * @param channel TX channel ID
 */
// llgo:link (*Dma2dHalContextT).Dma2dHalTxResetChannel C.dma2d_hal_tx_reset_channel
func (recv_ *Dma2dHalContextT) Dma2dHalTxResetChannel(channel c.Uint32T) {
}

/**
 * @brief Reset 2D-DMA RX channel
 *
 * @param hal Context of the HAL layer
 * @param channel RX channel ID
 */
// llgo:link (*Dma2dHalContextT).Dma2dHalRxResetChannel C.dma2d_hal_rx_reset_channel
func (recv_ *Dma2dHalContextT) Dma2dHalRxResetChannel(channel c.Uint32T) {
}
