package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:link (*GdmaHalContextT).GdmaAxiHalStartWithDesc C.gdma_axi_hal_start_with_desc
func (recv_ *GdmaHalContextT) GdmaAxiHalStartWithDesc(chan_id c.Int, dir GdmaChannelDirectionT, desc_base_addr c.IntptrT) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalStop C.gdma_axi_hal_stop
func (recv_ *GdmaHalContextT) GdmaAxiHalStop(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalAppend C.gdma_axi_hal_append
func (recv_ *GdmaHalContextT) GdmaAxiHalAppend(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalReset C.gdma_axi_hal_reset
func (recv_ *GdmaHalContextT) GdmaAxiHalReset(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalSetPriority C.gdma_axi_hal_set_priority
func (recv_ *GdmaHalContextT) GdmaAxiHalSetPriority(chan_id c.Int, dir GdmaChannelDirectionT, priority c.Uint32T) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalConnectPeri C.gdma_axi_hal_connect_peri
func (recv_ *GdmaHalContextT) GdmaAxiHalConnectPeri(chan_id c.Int, dir GdmaChannelDirectionT, periph GdmaTriggerPeripheralT, periph_sub_id c.Int) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalDisconnectPeri C.gdma_axi_hal_disconnect_peri
func (recv_ *GdmaHalContextT) GdmaAxiHalDisconnectPeri(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalEnableBurst C.gdma_axi_hal_enable_burst
func (recv_ *GdmaHalContextT) GdmaAxiHalEnableBurst(chan_id c.Int, dir GdmaChannelDirectionT, en_data_burst bool, en_desc_burst bool) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalSetBurstSize C.gdma_axi_hal_set_burst_size
func (recv_ *GdmaHalContextT) GdmaAxiHalSetBurstSize(chan_id c.Int, dir GdmaChannelDirectionT, burst_sz c.Uint32T) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalSetStrategy C.gdma_axi_hal_set_strategy
func (recv_ *GdmaHalContextT) GdmaAxiHalSetStrategy(chan_id c.Int, dir GdmaChannelDirectionT, en_owner_check bool, en_desc_write_back bool, eof_till_popped bool) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalEnableIntr C.gdma_axi_hal_enable_intr
func (recv_ *GdmaHalContextT) GdmaAxiHalEnableIntr(chan_id c.Int, dir GdmaChannelDirectionT, intr_event_mask c.Uint32T, en_or_dis bool) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalClearIntr C.gdma_axi_hal_clear_intr
func (recv_ *GdmaHalContextT) GdmaAxiHalClearIntr(chan_id c.Int, dir GdmaChannelDirectionT, intr_event_mask c.Uint32T) {
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalReadIntrStatus C.gdma_axi_hal_read_intr_status
func (recv_ *GdmaHalContextT) GdmaAxiHalReadIntrStatus(chan_id c.Int, dir GdmaChannelDirectionT, raw bool) c.Uint32T {
	return 0
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalGetIntrStatusReg C.gdma_axi_hal_get_intr_status_reg
func (recv_ *GdmaHalContextT) GdmaAxiHalGetIntrStatusReg(chan_id c.Int, dir GdmaChannelDirectionT) c.Uint32T {
	return 0
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalGetEofDescAddr C.gdma_axi_hal_get_eof_desc_addr
func (recv_ *GdmaHalContextT) GdmaAxiHalGetEofDescAddr(chan_id c.Int, dir GdmaChannelDirectionT, is_success bool) c.Uint32T {
	return 0
}

// llgo:link (*GdmaHalContextT).GdmaAxiHalInit C.gdma_axi_hal_init
func (recv_ *GdmaHalContextT) GdmaAxiHalInit(config *GdmaHalConfigT) {
}
