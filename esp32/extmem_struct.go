package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ExtmemDevS struct {
	DcacheCtrl struct {
		Val c.Uint32T
	}
	DcacheCtrl1 struct {
		Val c.Uint32T
	}
	DcacheTagPowerCtrl struct {
		Val c.Uint32T
	}
	DcachePrelockCtrl struct {
		Val c.Uint32T
	}
	DcachePrelockSct0Addr c.Uint32T
	DcachePrelockSct1Addr c.Uint32T
	DcachePrelockSctSize  struct {
		Val c.Uint32T
	}
	DcacheLockCtrl struct {
		Val c.Uint32T
	}
	DcacheLockAddr c.Uint32T
	DcacheLockSize struct {
		Val c.Uint32T
	}
	DcacheSyncCtrl struct {
		Val c.Uint32T
	}
	DcacheSyncAddr c.Uint32T
	DcacheSyncSize struct {
		Val c.Uint32T
	}
	DcacheOccupyCtrl struct {
		Val c.Uint32T
	}
	DcacheOccupyAddr c.Uint32T
	DcacheOccupySize struct {
		Val c.Uint32T
	}
	DcachePreloadCtrl struct {
		Val c.Uint32T
	}
	DcachePreloadAddr c.Uint32T
	DcachePreloadSize struct {
		Val c.Uint32T
	}
	DcacheAutoloadCtrl struct {
		Val c.Uint32T
	}
	DcacheAutoloadSct0Addr c.Uint32T
	DcacheAutoloadSct0Size struct {
		Val c.Uint32T
	}
	DcacheAutoloadSct1Addr c.Uint32T
	DcacheAutoloadSct1Size struct {
		Val c.Uint32T
	}
	IcacheCtrl struct {
		Val c.Uint32T
	}
	IcacheCtrl1 struct {
		Val c.Uint32T
	}
	IcacheTagPowerCtrl struct {
		Val c.Uint32T
	}
	IcachePrelockCtrl struct {
		Val c.Uint32T
	}
	IcachePrelockSct0Addr c.Uint32T
	IcachePrelockSct1Addr c.Uint32T
	IcachePrelockSctSize  struct {
		Val c.Uint32T
	}
	IcacheLockCtrl struct {
		Val c.Uint32T
	}
	IcacheLockAddr c.Uint32T
	IcacheLockSize struct {
		Val c.Uint32T
	}
	IcacheSyncCtrl struct {
		Val c.Uint32T
	}
	IcacheSyncAddr c.Uint32T
	IcacheSyncSize struct {
		Val c.Uint32T
	}
	IcachePreloadCtrl struct {
		Val c.Uint32T
	}
	IcachePreloadAddr c.Uint32T
	IcachePreloadSize struct {
		Val c.Uint32T
	}
	IcacheAutoloadCtrl struct {
		Val c.Uint32T
	}
	IcacheAutoloadSct0Addr c.Uint32T
	IcacheAutoloadSct0Size struct {
		Val c.Uint32T
	}
	IcacheAutoloadSct1Addr c.Uint32T
	IcacheAutoloadSct1Size struct {
		Val c.Uint32T
	}
	IbusToFlashStartVaddr c.Uint32T
	IbusToFlashEndVaddr   c.Uint32T
	DbusToFlashStartVaddr c.Uint32T
	DbusToFlashEndVaddr   c.Uint32T
	CacheAcsCntClr        struct {
		Val c.Uint32T
	}
	IbusAcsMissCnt       c.Uint32T
	IbusAcsCnt           c.Uint32T
	DbusAcsFlashMissCnt  c.Uint32T
	DbusAcsSpiramMissCnt c.Uint32T
	DbusAcsCnt           c.Uint32T
	CacheIlgIntEna       struct {
		Val c.Uint32T
	}
	CacheIlgIntClr struct {
		Val c.Uint32T
	}
	CacheIlgIntSt struct {
		Val c.Uint32T
	}
	Core0AcsCacheIntEna struct {
		Val c.Uint32T
	}
	Core0AcsCacheIntClr struct {
		Val c.Uint32T
	}
	Core0AcsCacheIntSt struct {
		Val c.Uint32T
	}
	Core1AcsCacheIntEna struct {
		Val c.Uint32T
	}
	Core1AcsCacheIntClr struct {
		Val c.Uint32T
	}
	Core1AcsCacheIntSt struct {
		Val c.Uint32T
	}
	Core0DbusRejectSt struct {
		Val c.Uint32T
	}
	Core0DbusRejectVaddr c.Uint32T
	Core0IbusRejectSt    struct {
		Val c.Uint32T
	}
	Core0IbusRejectVaddr c.Uint32T
	Core1DbusRejectSt    struct {
		Val c.Uint32T
	}
	Core1DbusRejectVaddr c.Uint32T
	Core1IbusRejectSt    struct {
		Val c.Uint32T
	}
	Core1IbusRejectVaddr c.Uint32T
	CacheMmuFaultContent struct {
		Val c.Uint32T
	}
	CacheMmuFaultVaddr  c.Uint32T
	CacheWrapAroundCtrl struct {
		Val c.Uint32T
	}
	CacheMmuPowerCtrl struct {
		Val c.Uint32T
	}
	CacheState struct {
		Val c.Uint32T
	}
	CacheEncryptDecryptRecordDisable struct {
		Val c.Uint32T
	}
	CacheEncryptDecryptClkForceOn struct {
		Val c.Uint32T
	}
	CacheBridgeArbiterCtrl struct {
		Val c.Uint32T
	}
	CachePreloadIntCtrl struct {
		Val c.Uint32T
	}
	CacheSyncIntCtrl struct {
		Val c.Uint32T
	}
	CacheMmuOwner struct {
		Val c.Uint32T
	}
	CacheConfMisc struct {
		Val c.Uint32T
	}
	DcacheFreeze struct {
		Val c.Uint32T
	}
	IcacheFreeze struct {
		Val c.Uint32T
	}
	IcacheAtomicOperateEna struct {
		Val c.Uint32T
	}
	DcacheAtomicOperateEna struct {
		Val c.Uint32T
	}
	CacheRequest struct {
		Val c.Uint32T
	}
	ClockGate struct {
		Val c.Uint32T
	}
	Reserved168        c.Uint32T
	Reserved16c        c.Uint32T
	Reserved170        c.Uint32T
	Reserved174        c.Uint32T
	Reserved178        c.Uint32T
	Reserved17c        c.Uint32T
	CacheTagObjectCtrl struct {
		Val c.Uint32T
	}
	CacheTagWayObject struct {
		Val c.Uint32T
	}
	CacheVaddr      c.Uint32T
	CacheTagContent c.Uint32T
	Reserved190     c.Uint32T
	Reserved194     c.Uint32T
	Reserved198     c.Uint32T
	Reserved19c     c.Uint32T
	Reserved1a0     c.Uint32T
	Reserved1a4     c.Uint32T
	Reserved1a8     c.Uint32T
	Reserved1ac     c.Uint32T
	Reserved1b0     c.Uint32T
	Reserved1b4     c.Uint32T
	Reserved1b8     c.Uint32T
	Reserved1bc     c.Uint32T
	Reserved1c0     c.Uint32T
	Reserved1c4     c.Uint32T
	Reserved1c8     c.Uint32T
	Reserved1cc     c.Uint32T
	Reserved1d0     c.Uint32T
	Reserved1d4     c.Uint32T
	Reserved1d8     c.Uint32T
	Reserved1dc     c.Uint32T
	Reserved1e0     c.Uint32T
	Reserved1e4     c.Uint32T
	Reserved1e8     c.Uint32T
	Reserved1ec     c.Uint32T
	Reserved1f0     c.Uint32T
	Reserved1f4     c.Uint32T
	Reserved1f8     c.Uint32T
	Reserved1fc     c.Uint32T
	Reserved200     c.Uint32T
	Reserved204     c.Uint32T
	Reserved208     c.Uint32T
	Reserved20c     c.Uint32T
	Reserved210     c.Uint32T
	Reserved214     c.Uint32T
	Reserved218     c.Uint32T
	Reserved21c     c.Uint32T
	Reserved220     c.Uint32T
	Reserved224     c.Uint32T
	Reserved228     c.Uint32T
	Reserved22c     c.Uint32T
	Reserved230     c.Uint32T
	Reserved234     c.Uint32T
	Reserved238     c.Uint32T
	Reserved23c     c.Uint32T
	Reserved240     c.Uint32T
	Reserved244     c.Uint32T
	Reserved248     c.Uint32T
	Reserved24c     c.Uint32T
	Reserved250     c.Uint32T
	Reserved254     c.Uint32T
	Reserved258     c.Uint32T
	Reserved25c     c.Uint32T
	Reserved260     c.Uint32T
	Reserved264     c.Uint32T
	Reserved268     c.Uint32T
	Reserved26c     c.Uint32T
	Reserved270     c.Uint32T
	Reserved274     c.Uint32T
	Reserved278     c.Uint32T
	Reserved27c     c.Uint32T
	Reserved280     c.Uint32T
	Reserved284     c.Uint32T
	Reserved288     c.Uint32T
	Reserved28c     c.Uint32T
	Reserved290     c.Uint32T
	Reserved294     c.Uint32T
	Reserved298     c.Uint32T
	Reserved29c     c.Uint32T
	Reserved2a0     c.Uint32T
	Reserved2a4     c.Uint32T
	Reserved2a8     c.Uint32T
	Reserved2ac     c.Uint32T
	Reserved2b0     c.Uint32T
	Reserved2b4     c.Uint32T
	Reserved2b8     c.Uint32T
	Reserved2bc     c.Uint32T
	Reserved2c0     c.Uint32T
	Reserved2c4     c.Uint32T
	Reserved2c8     c.Uint32T
	Reserved2cc     c.Uint32T
	Reserved2d0     c.Uint32T
	Reserved2d4     c.Uint32T
	Reserved2d8     c.Uint32T
	Reserved2dc     c.Uint32T
	Reserved2e0     c.Uint32T
	Reserved2e4     c.Uint32T
	Reserved2e8     c.Uint32T
	Reserved2ec     c.Uint32T
	Reserved2f0     c.Uint32T
	Reserved2f4     c.Uint32T
	Reserved2f8     c.Uint32T
	Reserved2fc     c.Uint32T
	Reserved300     c.Uint32T
	Reserved304     c.Uint32T
	Reserved308     c.Uint32T
	Reserved30c     c.Uint32T
	Reserved310     c.Uint32T
	Reserved314     c.Uint32T
	Reserved318     c.Uint32T
	Reserved31c     c.Uint32T
	Reserved320     c.Uint32T
	Reserved324     c.Uint32T
	Reserved328     c.Uint32T
	Reserved32c     c.Uint32T
	Reserved330     c.Uint32T
	Reserved334     c.Uint32T
	Reserved338     c.Uint32T
	Reserved33c     c.Uint32T
	Reserved340     c.Uint32T
	Reserved344     c.Uint32T
	Reserved348     c.Uint32T
	Reserved34c     c.Uint32T
	Reserved350     c.Uint32T
	Reserved354     c.Uint32T
	Reserved358     c.Uint32T
	Reserved35c     c.Uint32T
	Reserved360     c.Uint32T
	Reserved364     c.Uint32T
	Reserved368     c.Uint32T
	Reserved36c     c.Uint32T
	Reserved370     c.Uint32T
	Reserved374     c.Uint32T
	Reserved378     c.Uint32T
	Reserved37c     c.Uint32T
	Reserved380     c.Uint32T
	Reserved384     c.Uint32T
	Reserved388     c.Uint32T
	Reserved38c     c.Uint32T
	Reserved390     c.Uint32T
	Reserved394     c.Uint32T
	Reserved398     c.Uint32T
	Reserved39c     c.Uint32T
	Reserved3a0     c.Uint32T
	Reserved3a4     c.Uint32T
	Reserved3a8     c.Uint32T
	Reserved3ac     c.Uint32T
	Reserved3b0     c.Uint32T
	Reserved3b4     c.Uint32T
	Reserved3b8     c.Uint32T
	Reserved3bc     c.Uint32T
	Reserved3c0     c.Uint32T
	Reserved3c4     c.Uint32T
	Reserved3c8     c.Uint32T
	Reserved3cc     c.Uint32T
	Reserved3d0     c.Uint32T
	Reserved3d4     c.Uint32T
	Reserved3d8     c.Uint32T
	Reserved3dc     c.Uint32T
	Reserved3e0     c.Uint32T
	Reserved3e4     c.Uint32T
	Reserved3e8     c.Uint32T
	Reserved3ec     c.Uint32T
	Reserved3f0     c.Uint32T
	Reserved3f4     c.Uint32T
	Reserved3f8     c.Uint32T
	Date            struct {
		Val c.Uint32T
	}
}
type ExtmemDevT ExtmemDevS
