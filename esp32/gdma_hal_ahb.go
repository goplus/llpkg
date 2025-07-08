package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:link (*GdmaHalContextT).GdmaAhbHalStartWithDesc C.gdma_ahb_hal_start_with_desc
func (recv_ *GdmaHalContextT) GdmaAhbHalStartWithDesc(chan_id c.Int, dir GdmaChannelDirectionT, desc_base_addr c.IntptrT) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalStop C.gdma_ahb_hal_stop
func (recv_ *GdmaHalContextT) GdmaAhbHalStop(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalAppend C.gdma_ahb_hal_append
func (recv_ *GdmaHalContextT) GdmaAhbHalAppend(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalReset C.gdma_ahb_hal_reset
func (recv_ *GdmaHalContextT) GdmaAhbHalReset(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalSetPriority C.gdma_ahb_hal_set_priority
func (recv_ *GdmaHalContextT) GdmaAhbHalSetPriority(chan_id c.Int, dir GdmaChannelDirectionT, priority c.Uint32T) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalConnectPeri C.gdma_ahb_hal_connect_peri
func (recv_ *GdmaHalContextT) GdmaAhbHalConnectPeri(chan_id c.Int, dir GdmaChannelDirectionT, periph GdmaTriggerPeripheralT, periph_sub_id c.Int) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalDisconnectPeri C.gdma_ahb_hal_disconnect_peri
func (recv_ *GdmaHalContextT) GdmaAhbHalDisconnectPeri(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalEnableBurst C.gdma_ahb_hal_enable_burst
func (recv_ *GdmaHalContextT) GdmaAhbHalEnableBurst(chan_id c.Int, dir GdmaChannelDirectionT, en_data_burst bool, en_desc_burst bool) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalSetBurstSize C.gdma_ahb_hal_set_burst_size
func (recv_ *GdmaHalContextT) GdmaAhbHalSetBurstSize(chan_id c.Int, dir GdmaChannelDirectionT, burst_sz c.Uint32T) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalSetStrategy C.gdma_ahb_hal_set_strategy
func (recv_ *GdmaHalContextT) GdmaAhbHalSetStrategy(chan_id c.Int, dir GdmaChannelDirectionT, en_owner_check bool, en_desc_write_back bool, eof_till_popped bool) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalEnableIntr C.gdma_ahb_hal_enable_intr
func (recv_ *GdmaHalContextT) GdmaAhbHalEnableIntr(chan_id c.Int, dir GdmaChannelDirectionT, intr_event_mask c.Uint32T, en_or_dis bool) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalClearIntr C.gdma_ahb_hal_clear_intr
func (recv_ *GdmaHalContextT) GdmaAhbHalClearIntr(chan_id c.Int, dir GdmaChannelDirectionT, intr_event_mask c.Uint32T) {
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalReadIntrStatus C.gdma_ahb_hal_read_intr_status
func (recv_ *GdmaHalContextT) GdmaAhbHalReadIntrStatus(chan_id c.Int, dir GdmaChannelDirectionT, raw bool) c.Uint32T {
	return 0
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalGetIntrStatusReg C.gdma_ahb_hal_get_intr_status_reg
func (recv_ *GdmaHalContextT) GdmaAhbHalGetIntrStatusReg(chan_id c.Int, dir GdmaChannelDirectionT) c.Uint32T {
	return 0
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalGetEofDescAddr C.gdma_ahb_hal_get_eof_desc_addr
func (recv_ *GdmaHalContextT) GdmaAhbHalGetEofDescAddr(chan_id c.Int, dir GdmaChannelDirectionT, is_success bool) c.Uint32T {
	return 0
}

// llgo:link (*GdmaHalContextT).GdmaAhbHalInit C.gdma_ahb_hal_init
func (recv_ *GdmaHalContextT) GdmaAhbHalInit(config *GdmaHalConfigT) {
}
