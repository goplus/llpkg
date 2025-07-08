package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RtcCntlSleepCacheTagRetent struct {
	LinkAddr c.Pointer
	Icache   struct {
		StartPoint c.Uint32T
		VldSize    c.Uint32T
		Size       c.Uint32T
		Enable     c.Uint32T
	}
	Dcache struct {
		StartPoint c.Uint32T
		VldSize    c.Uint32T
		Size       c.Uint32T
		Enable     c.Uint32T
	}
}
type RtcCntlSleepCacheTagRetentT RtcCntlSleepCacheTagRetent

type RtcCntlSleepRetent struct {
	CpuPdMem c.Pointer
	Tagmem   RtcCntlSleepCacheTagRetentT
}
type RtcCntlSleepRetentT RtcCntlSleepRetent

//go:linkname RtcCntlHalDmaLinkInit C.rtc_cntl_hal_dma_link_init
func RtcCntlHalDmaLinkInit(elem c.Pointer, buff c.Pointer, size c.Int, next c.Pointer) c.Pointer

//go:linkname RtcCntlHalEnableCpuRetention C.rtc_cntl_hal_enable_cpu_retention
func RtcCntlHalEnableCpuRetention(addr c.Pointer)

//go:linkname RtcCntlHalDisableCpuRetention C.rtc_cntl_hal_disable_cpu_retention
func RtcCntlHalDisableCpuRetention(addr c.Pointer)

//go:linkname RtcCntlHalEnableTagmemRetention C.rtc_cntl_hal_enable_tagmem_retention
func RtcCntlHalEnableTagmemRetention(addr c.Pointer)

//go:linkname RtcCntlHalDisableTagmemRetention C.rtc_cntl_hal_disable_tagmem_retention
func RtcCntlHalDisableTagmemRetention(addr c.Pointer)
