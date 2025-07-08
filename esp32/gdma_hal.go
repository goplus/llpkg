package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaHalContextT struct {
	PrivData               *GdmaHalPrivDataT
	StartWithDesc          c.Pointer
	Stop                   c.Pointer
	Append                 c.Pointer
	Reset                  c.Pointer
	SetPriority            c.Pointer
	ConnectPeri            c.Pointer
	DisconnectPeri         c.Pointer
	EnableBurst            c.Pointer
	SetBurstSize           c.Pointer
	SetStrategy            c.Pointer
	GetIntrStatusReg       c.Pointer
	EnableIntr             c.Pointer
	ClearIntr              c.Pointer
	ReadIntrStatus         c.Pointer
	GetEofDescAddr         c.Pointer
	EnableAccessEncryptMem c.Pointer
}

/**
 * @brief GDMA HAL configuration
 */

type GdmaHalConfigT struct {
	GroupId c.Int
}

type GdmaHalCrcConfigT struct {
	InitValue       c.Uint32T
	CrcBitWidth     c.Uint32T
	PolyHex         c.Uint32T
	ReverseDataMask bool
}

/**
 * @brief GDMA HAL private data
 */

type GdmaHalPrivDataT struct {
	M2mFreePeriphMask c.Uint32T
}

// llgo:link (*GdmaHalContextT).GdmaHalDeinit C.gdma_hal_deinit
func (recv_ *GdmaHalContextT) GdmaHalDeinit() {
}

// llgo:link (*GdmaHalContextT).GdmaHalStartWithDesc C.gdma_hal_start_with_desc
func (recv_ *GdmaHalContextT) GdmaHalStartWithDesc(chan_id c.Int, dir GdmaChannelDirectionT, desc_base_addr c.IntptrT) {
}

// llgo:link (*GdmaHalContextT).GdmaHalStop C.gdma_hal_stop
func (recv_ *GdmaHalContextT) GdmaHalStop(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaHalAppend C.gdma_hal_append
func (recv_ *GdmaHalContextT) GdmaHalAppend(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaHalReset C.gdma_hal_reset
func (recv_ *GdmaHalContextT) GdmaHalReset(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaHalSetPriority C.gdma_hal_set_priority
func (recv_ *GdmaHalContextT) GdmaHalSetPriority(chan_id c.Int, dir GdmaChannelDirectionT, priority c.Uint32T) {
}

// llgo:link (*GdmaHalContextT).GdmaHalConnectPeri C.gdma_hal_connect_peri
func (recv_ *GdmaHalContextT) GdmaHalConnectPeri(chan_id c.Int, dir GdmaChannelDirectionT, periph GdmaTriggerPeripheralT, periph_sub_id c.Int) {
}

// llgo:link (*GdmaHalContextT).GdmaHalDisconnectPeri C.gdma_hal_disconnect_peri
func (recv_ *GdmaHalContextT) GdmaHalDisconnectPeri(chan_id c.Int, dir GdmaChannelDirectionT) {
}

// llgo:link (*GdmaHalContextT).GdmaHalEnableBurst C.gdma_hal_enable_burst
func (recv_ *GdmaHalContextT) GdmaHalEnableBurst(chan_id c.Int, dir GdmaChannelDirectionT, en_data_burst bool, en_desc_burst bool) {
}

// llgo:link (*GdmaHalContextT).GdmaHalSetBurstSize C.gdma_hal_set_burst_size
func (recv_ *GdmaHalContextT) GdmaHalSetBurstSize(chan_id c.Int, dir GdmaChannelDirectionT, burst_sz c.Uint32T) {
}

// llgo:link (*GdmaHalContextT).GdmaHalSetStrategy C.gdma_hal_set_strategy
func (recv_ *GdmaHalContextT) GdmaHalSetStrategy(chan_id c.Int, dir GdmaChannelDirectionT, en_owner_check bool, en_desc_write_back bool, eof_till_popped bool) {
}

// llgo:link (*GdmaHalContextT).GdmaHalEnableIntr C.gdma_hal_enable_intr
func (recv_ *GdmaHalContextT) GdmaHalEnableIntr(chan_id c.Int, dir GdmaChannelDirectionT, intr_event_mask c.Uint32T, en_or_dis bool) {
}

// llgo:link (*GdmaHalContextT).GdmaHalClearIntr C.gdma_hal_clear_intr
func (recv_ *GdmaHalContextT) GdmaHalClearIntr(chan_id c.Int, dir GdmaChannelDirectionT, intr_event_mask c.Uint32T) {
}

// llgo:link (*GdmaHalContextT).GdmaHalGetIntrStatusReg C.gdma_hal_get_intr_status_reg
func (recv_ *GdmaHalContextT) GdmaHalGetIntrStatusReg(chan_id c.Int, dir GdmaChannelDirectionT) c.Uint32T {
	return 0
}

// llgo:link (*GdmaHalContextT).GdmaHalReadIntrStatus C.gdma_hal_read_intr_status
func (recv_ *GdmaHalContextT) GdmaHalReadIntrStatus(chan_id c.Int, dir GdmaChannelDirectionT, raw bool) c.Uint32T {
	return 0
}

// llgo:link (*GdmaHalContextT).GdmaHalGetEofDescAddr C.gdma_hal_get_eof_desc_addr
func (recv_ *GdmaHalContextT) GdmaHalGetEofDescAddr(chan_id c.Int, dir GdmaChannelDirectionT, is_success bool) c.Uint32T {
	return 0
}

// llgo:link (*GdmaHalContextT).GdmaHalEnableAccessEncryptMem C.gdma_hal_enable_access_encrypt_mem
func (recv_ *GdmaHalContextT) GdmaHalEnableAccessEncryptMem(chan_id c.Int, dir GdmaChannelDirectionT, en_or_dis bool) {
}
