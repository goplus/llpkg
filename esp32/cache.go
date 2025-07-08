package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MIN_ICACHE_SIZE = 16384
const MAX_ICACHE_SIZE = 32768
const MIN_DCACHE_SIZE = 32768
const MAX_DCACHE_SIZE = 65536
const MIN_ICACHE_WAYS = 4
const MAX_ICACHE_WAYS = 8
const MIN_DCACHE_WAYS = 4
const MAX_DCACHE_WAYS = 4
const MAX_CACHE_WAYS = 8
const MIN_CACHE_LINE_SIZE = 16
const TAG_SIZE = 4
const MIN_ICACHE_BANK_NUM = 1
const MAX_ICACHE_BANK_NUM = 2
const MIN_DCACHE_BANK_NUM = 1
const MAX_DCACHE_BANK_NUM = 2
const CACHE_MEMORY_BANK_NUM = 4
const CACHE_MEMORY_IBANK_SIZE = 0x4000
const CACHE_MEMORY_DBANK_SIZE = 0x8000
const ESP_ROM_ERR_INVALID_ARG = 1
const MMU_SET_ADDR_ALIGNED_ERROR = 2
const MMU_SET_PASE_SIZE_ERROR = 3
const MMU_SET_VADDR_OUT_RANGE = 4
const CACHE_OP_ICACHE_Y = 1
const CACHE_OP_ICACHE_N = 0

type CacheT c.Int

const (
	CACHE_DCACHE  CacheT = 0
	CACHE_ICACHE0 CacheT = 1
	CACHE_ICACHE1 CacheT = 2
)

type CacheArrayT c.Int

const (
	CACHE_MEMORY_INVALID CacheArrayT = 0
	CACHE_MEMORY_IBANK0  CacheArrayT = 1
	CACHE_MEMORY_IBANK1  CacheArrayT = 2
	CACHE_MEMORY_IBANK2  CacheArrayT = 4
	CACHE_MEMORY_IBANK3  CacheArrayT = 8
	CACHE_MEMORY_DBANK0  CacheArrayT = 1
	CACHE_MEMORY_DBANK1  CacheArrayT = 2
	CACHE_MEMORY_DBANK2  CacheArrayT = 4
	CACHE_MEMORY_DBANK3  CacheArrayT = 8
)

type CacheSizeT c.Int

const (
	CACHE_SIZE_HALF CacheSizeT = 0
	CACHE_SIZE_FULL CacheSizeT = 1
)

type CacheWaysT c.Int

const (
	CACHE_4WAYS_ASSOC CacheWaysT = 0
	CACHE_8WAYS_ASSOC CacheWaysT = 1
)

type CacheLineSizeT c.Int

const (
	CACHE_LINE_SIZE_16B CacheLineSizeT = 0
	CACHE_LINE_SIZE_32B CacheLineSizeT = 1
	CACHE_LINE_SIZE_64B CacheLineSizeT = 2
)

type CacheAutoloadOrderT c.Int

const (
	CACHE_AUTOLOAD_POSITIVE CacheAutoloadOrderT = 0
	CACHE_AUTOLOAD_NEGATIVE CacheAutoloadOrderT = 1
)

type CacheAutoloadRegionT c.Int

const (
	CACHE_AUTOLOAD_REGION0 CacheAutoloadRegionT = 0
	CACHE_AUTOLOAD_REGION1 CacheAutoloadRegionT = 1
)

type CacheAutoloadTriggerT c.Int

const (
	CACHE_AUTOLOAD_MISS_TRIGGER CacheAutoloadTriggerT = 0
	CACHE_AUTOLOAD_HIT_TRIGGER  CacheAutoloadTriggerT = 1
	CACHE_AUTOLOAD_BOTH_TRIGGER CacheAutoloadTriggerT = 2
)

type CacheFreezeModeT c.Int

const (
	CACHE_FREEZE_ACK_BUSY  CacheFreezeModeT = 0
	CACHE_FREEZE_ACK_ERROR CacheFreezeModeT = 1
)

type CacheMode struct {
	CacheSize     c.Uint32T
	CacheLineSize c.Uint16T
	CacheWays     c.Uint8T
	Icache        c.Uint8T
}

type IcacheTagItem struct {
	Valid    c.Uint32T
	Lock     c.Uint32T
	Attr     c.Uint32T
	FifoCnt  c.Uint32T
	Tag      c.Uint32T
	Reserved c.Uint32T
}

type DcacheTagItem struct {
	Dirty    c.Uint32T
	Valid    c.Uint32T
	Lock     c.Uint32T
	Occupy   c.Uint32T
	Attr     c.Uint32T
	FifoCnt  c.Uint32T
	Tag      c.Uint32T
	Reserved c.Uint32T
}

type AutoloadConfig struct {
	Ena     c.Uint8T
	Order   c.Uint8T
	Trigger c.Uint8T
	Size    c.Uint8T
}

type AutoloadRegionConfig struct {
	Region c.Uint8T
	Ena    c.Uint8T
	Addr   c.Uint32T
	Size   c.Uint32T
}

type TagGroupInfo struct {
	Mode              CacheMode
	FilterAddr        c.Uint32T
	VaddrOffset       c.Uint32T
	TagAddr           [8]c.Uint32T
	CacheMemoryOffset [8]c.Uint32T
	UseLegacy         c.Uint8T
}

type LockConfig struct {
	Addr  c.Uint32T
	Size  c.Uint16T
	Group c.Uint16T
}

type CacheInternalStubTable struct {
	IcacheLineSize        c.Pointer
	DcacheLineSize        c.Pointer
	IcacheAddr            c.Pointer
	DcacheAddr            c.Pointer
	InvalidateIcacheItems c.Pointer
	InvalidateDcacheItems c.Pointer
	CleanItems            c.Pointer
	WritebackItems        c.Pointer
	LockIcacheItems       c.Pointer
	LockDcacheItems       c.Pointer
	UnlockIcacheItems     c.Pointer
	UnlockDcacheItems     c.Pointer
	OccupyItems           c.Pointer
	SuspendIcacheAutoload c.Pointer
	ResumeIcacheAutoload  c.Pointer
	SuspendDcacheAutoload c.Pointer
	ResumeDcacheAutoload  c.Pointer
	FreezeIcacheEnable    c.Pointer
	FreezeIcacheDisable   c.Pointer
	FreezeDcacheEnable    c.Pointer
	FreezeDcacheDisable   c.Pointer
	OpAddr                c.Pointer
}

// llgo:type C
type CacheOpStart func()

// llgo:type C
type CacheOpEnd func()

type CacheOpCbT struct {
	Start CacheOpStart
	End   CacheOpEnd
}

/**
 * @brief Initialise cache mmu, mark all entries as invalid.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname CacheMMUInit C.Cache_MMU_Init
func CacheMMUInit()

/**
 * @brief Set ICache mmu mapping.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t ext_ram : SOC_MMU_ACCESS_FLASH for flash, SOC_MMU_ACCESS_SPIRAM for spiram, SOC_MMU_INVALID for invalid.
 *
 * @param  uint32_t vaddr : virtual address in CPU address space.
 *                              Can be Iram0,Iram1,Irom0,Drom0 and AHB buses address.
 *                              Should be aligned by psize.
 *
 * @param  uint32_t paddr : physical address in external memory.
 *                              Should be aligned by psize.
 *
 * @param  uint32_t psize : page size of ICache, in kilobytes. Should be 64 here.
 *
 * @param  uint32_t num : pages to be set.
 *
 * @param  uint32_t fixed : 0 for physical pages grow with virtual pages, other for virtual pages map to same physical page.
 *
 * @return uint32_t: error status
 *                   0 : mmu set success
 *                   2 : vaddr or paddr is not aligned
 *                   3 : psize error
 *                   4 : vaddr is out of range
 */
//go:linkname CacheIbusMMUSet C.Cache_Ibus_MMU_Set
func CacheIbusMMUSet(ext_ram c.Uint32T, vaddr c.Uint32T, paddr c.Uint32T, psize c.Uint32T, num c.Uint32T, fixed c.Uint32T) c.Int

/**
 * @brief Set DCache mmu mapping.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t ext_ram : SOC_MMU_ACCESS_FLASH for flash, SOC_MMU_ACCESS_SPIRAM for spiram, SOC_MMU_INVALID for invalid.
 *
 * @param  uint32_t vaddr : virtual address in CPU address space.
 *                              Can be DRam0, DRam1, DRom0, DPort and AHB buses address.
 *                              Should be aligned by psize.
 *
 * @param  uint32_t paddr : physical address in external memory.
 *                              Should be aligned by psize.
 *
 * @param  uint32_t psize : page size of DCache, in kilobytes. Should be 64 here.
 *
 * @param  uint32_t num : pages to be set.

 * @param  uint32_t fixed : 0 for physical pages grow with virtual pages, other for virtual pages map to same physical page.
 *
 * @return uint32_t: error status
 *                   0 : mmu set success
 *                   2 : vaddr or paddr is not aligned
 *                   3 : psize error
 *                   4 : vaddr is out of range
 */
//go:linkname CacheDbusMMUSet C.Cache_Dbus_MMU_Set
func CacheDbusMMUSet(ext_ram c.Uint32T, vaddr c.Uint32T, paddr c.Uint32T, psize c.Uint32T, num c.Uint32T, fixed c.Uint32T) c.Int

/**
 * @brief Count the pages in the bus room address which map to Flash.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t bus : the bus to count with.
 *
 * @param uint32_t * page0_mapped : value should be initial by user, 0 for not mapped, other for mapped count.
 *
 * return uint32_t : the number of pages which map to Flash.
 */
//go:linkname CacheCountFlashPages C.Cache_Count_Flash_Pages
func CacheCountFlashPages(bus c.Uint32T, page0_mapped *c.Uint32T) c.Uint32T

/**
 * @brief Copy Instruction or rodata from Flash to SPIRAM, and remap to SPIRAM.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t bus : the bus which need to copy to SPIRAM.
 *
 * @param uint32_t bus_start_addr : the start virtual address for the bus.
 *
 * @param uint32_t start_page : the start (64KB) page number in SPIRAM.
 *
 * @param uint32_t * page0_page : the flash page0 in SPIRAM page number, 0xffff for invalid.
 *
 * return uint32_t : the next start page number for SPIRAM not mapped.
 */
//go:linkname CacheFlashToSPIRAMCopy C.Cache_Flash_To_SPIRAM_Copy
func CacheFlashToSPIRAMCopy(bus c.Uint32T, bus_start_addr c.Uint32T, start_page c.Uint32T, page0_page *c.Uint32T) c.Uint32T

/**
 * @brief allocate memory to used by ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param cache_array_t icache_low : the data array bank used by icache low part. Due to timing constraint, can only be CACHE_MEMORY_INVALID, CACHE_MEMORY_IBANK0
 *
 * @param cache_array_t icache_high : the data array bank used by icache high part. Due to timing constraint, can only be CACHE_MEMORY_INVALID, or CACHE_MEMORY_IBANK1 only if icache_low and icache_high is CACHE_MEMORY_IBANK0
 *
 * return none
 */
// llgo:link CacheArrayT.CacheOccupyICacheMEMORY C.Cache_Occupy_ICache_MEMORY
func (recv_ CacheArrayT) CacheOccupyICacheMEMORY(icache_high CacheArrayT) {
}

/**
 * @brief allocate memory to used by DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param cache_array_t dcache_low : the data array bank used by dcache low part. Due to timing constraint, can only be CACHE_MEMORY_INVALID, CACHE_MEMORY_DBANK1
 *
 * @param cache_array_t dcache1_high : the data array bank used by dcache high part. Due to timing constraint, can only be CACHE_MEMORY_INVALID, or CACHE_MEMORY_DBANK0 only if dcache_low0 and dcache_low1 is CACHE_MEMORY_DBANK1
 *
 * return none
 */
// llgo:link CacheArrayT.CacheOccupyDCacheMEMORY C.Cache_Occupy_DCache_MEMORY
func (recv_ CacheArrayT) CacheOccupyDCacheMEMORY(dcache_high CacheArrayT) {
}

/**
 * @brief Get cache mode of ICache or DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param struct cache_mode * mode : the pointer of cache mode struct, caller should set the icache field
 *
 * return none
 */
// llgo:link (*CacheMode).CacheGetMode C.Cache_Get_Mode
func (recv_ *CacheMode) CacheGetMode() {
}

/**
 * @brief set ICache modes: cache size, associate ways and cache line size.
 *        Please do not call this function in your SDK application.
 *
 * @param cache_size_t cache_size : the cache size, can be CACHE_SIZE_HALF and CACHE_SIZE_FULL
 *
 * @param cache_ways_t ways : the associate ways of cache, can be CACHE_4WAYS_ASSOC and CACHE_8WAYS_ASSOC
 *
 * @param cache_line_size_t cache_line_size : the cache line size, can be CACHE_LINE_SIZE_16B and CACHE_LINE_SIZE_32B
 *
 * return none
 */
// llgo:link CacheSizeT.CacheSetICacheMode C.Cache_Set_ICache_Mode
func (recv_ CacheSizeT) CacheSetICacheMode(ways CacheWaysT, cache_line_size CacheLineSizeT) {
}

/**
 * @brief set DCache modes: cache size, associate ways and cache line size.
 *        Please do not call this function in your SDK application.
 *
 * @param cache_size_t cache_size : the cache size, can be CACHE_SIZE_HALF and CACHE_SIZE_FULL
 *
 * @param cache_ways_t ways : the associate ways of cache, can be CACHE_4WAYS_ASSOC and CACHE_8WAYS_ASSOC, only CACHE_4WAYS_ASSOC works
 *
 * @param cache_line_size_t cache_line_size : the cache line size, can be CACHE_LINE_SIZE_16B, CACHE_LINE_SIZE_32B and CACHE_LINE_SIZE_64B
 *
 * return none
 */
// llgo:link CacheSizeT.CacheSetDCacheMode C.Cache_Set_DCache_Mode
func (recv_ CacheSizeT) CacheSetDCacheMode(ways CacheWaysT, cache_line_size CacheLineSizeT) {
}

/**
 * @brief check if the address is accessed through ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr : the address to check.
 *
 * @return 1 if the address is accessed through ICache, 0 if not.
 */
//go:linkname CacheAddressThroughICache C.Cache_Address_Through_ICache
func CacheAddressThroughICache(addr c.Uint32T) c.Uint32T

/**
 * @brief check if the address is accessed through DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr : the address to check.
 *
 * @return 1 if the address is accessed through DCache, 0 if not.
 */
//go:linkname CacheAddressThroughDCache C.Cache_Address_Through_DCache
func CacheAddressThroughDCache(addr c.Uint32T) c.Uint32T

/**
 * @brief Init Cache for ROM boot, including resetting the Dcache, initializing Owner, MMU, setting DCache mode, Enabling DCache, unmasking bus.
 *
 * @param None
 *
 * @return None
 */
//go:linkname ROMBootCacheInit C.ROM_Boot_Cache_Init
func ROMBootCacheInit()

/**
 * @brief Init mmu owner register to make i/d cache use half mmu entries.
 *
 * @param None
 *
 * @return None
 */
//go:linkname CacheOwnerInit C.Cache_Owner_Init
func CacheOwnerInit()

/**
 * @brief Invalidate the cache items for ICache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in ICache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr: start address to invalidate
 *
 * @param  uint32_t items: cache lines to invalidate, items * cache_line_size should not exceed the bus address size(16MB/32MB/64MB)
 *
 * @return None
 */
//go:linkname CacheInvalidateICacheItems C.Cache_Invalidate_ICache_Items
func CacheInvalidateICacheItems(addr c.Uint32T, items c.Uint32T)

/**
 * @brief Invalidate the cache items for DCache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in DCache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr: start address to invalidate
 *
 * @param  uint32_t items: cache lines to invalidate, items * cache_line_size should not exceed the bus address size(16MB/32MB/64MB)
 *
 * @return None
 */
//go:linkname CacheInvalidateDCacheItems C.Cache_Invalidate_DCache_Items
func CacheInvalidateDCacheItems(addr c.Uint32T, items c.Uint32T)

/**
 * @brief Clean the dirty bit of cache Items of DCache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in DCache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr: start address to Clean
 *
 * @param  uint32_t items: cache lines to invalidate, items * cache_line_size should not exceed the bus address size(16MB/32MB/64MB)
 *
 * @return None
 */
//go:linkname CacheCleanItems C.Cache_Clean_Items
func CacheCleanItems(addr c.Uint32T, items c.Uint32T)

/**
 * @brief Write back the cache items of DCache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in DCache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr: start address to write back
 *
 * @param  uint32_t items: cache lines to invalidate, items * cache_line_size should not exceed the bus address size(16MB/32MB/64MB)
 *
 * @return None
 */
//go:linkname CacheWriteBackItems C.Cache_WriteBack_Items
func CacheWriteBackItems(addr c.Uint32T, items c.Uint32T)

/**
 * @brief Invalidate the Cache items in the region from ICache or DCache.
 *        If the region is not in Cache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr : invalidated region start address.
 *
 * @param  uint32_t size : invalidated region size.
 *
 * @return 0 for success
 *         1 for invalid argument
 */
//go:linkname CacheInvalidateAddr C.Cache_Invalidate_Addr
func CacheInvalidateAddr(addr c.Uint32T, size c.Uint32T) c.Int

/**
 * @brief Clean the dirty bit of Cache items in the region from DCache.
 *        If the region is not in DCache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr : cleaned region start address.
 *
 * @param  uint32_t size : cleaned region size.
 *
 * @return 0 for success
 *         1 for invalid argument
 */
//go:linkname CacheCleanAddr C.Cache_Clean_Addr
func CacheCleanAddr(addr c.Uint32T, size c.Uint32T) c.Int

/**
 * @brief Writeback the Cache items(also clean the dirty bit) in the region from DCache.
 *        If the region is not in DCache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr : writeback region start address.
 *
 * @param  uint32_t size : writeback region size.
 *
 * @return 0 for success
 *         1 for invalid argument
 */
//go:linkname CacheWriteBackAddr C.Cache_WriteBack_Addr
func CacheWriteBackAddr(addr c.Uint32T, size c.Uint32T) c.Int

/**
 * @brief Invalidate all cache items in ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname CacheInvalidateICacheAll C.Cache_Invalidate_ICache_All
func CacheInvalidateICacheAll()

/**
 * @brief Invalidate all cache items in DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname CacheInvalidateDCacheAll C.Cache_Invalidate_DCache_All
func CacheInvalidateDCacheAll()

/**
 * @brief Clean the dirty bit of all cache items in DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname CacheCleanAll C.Cache_Clean_All
func CacheCleanAll()

/**
 * @brief WriteBack all cache items in DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname CacheWriteBackAll C.Cache_WriteBack_All
func CacheWriteBackAll()

/**
 * @brief Mask all buses through ICache and DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname CacheMaskAll C.Cache_Mask_All
func CacheMaskAll()

/**
 * @brief UnMask DRam0 bus through DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname CacheUnMaskDram0 C.Cache_UnMask_Dram0
func CacheUnMaskDram0()

/**
 * @brief Suspend ICache auto preload operation, then you can resume it after some ICache operations.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return uint32_t : 0 for ICache not auto preload before suspend.
 */
//go:linkname CacheSuspendICacheAutoload C.Cache_Suspend_ICache_Autoload
func CacheSuspendICacheAutoload() c.Uint32T

/**
 * @brief Resume ICache auto preload operation after some ICache operations.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t autoload : 0 for ICache not auto preload before suspend.
 *
 * @return None.
 */
//go:linkname CacheResumeICacheAutoload C.Cache_Resume_ICache_Autoload
func CacheResumeICacheAutoload(autoload c.Uint32T)

/**
 * @brief Suspend DCache auto preload operation, then you can resume it after some DCache operations.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return uint32_t : 0 for DCache not auto preload before suspend.
 */
//go:linkname CacheSuspendDCacheAutoload C.Cache_Suspend_DCache_Autoload
func CacheSuspendDCacheAutoload() c.Uint32T

/**
 * @brief Resume DCache auto preload operation after some DCache operations.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t autoload : 0 for DCache not auto preload before suspend.
 *
 * @return None.
 */
//go:linkname CacheResumeDCacheAutoload C.Cache_Resume_DCache_Autoload
func CacheResumeDCacheAutoload(autoload c.Uint32T)

/**
 * @brief Start an ICache manual preload, will suspend auto preload of ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t addr : start address of the preload region.
 *
 * @param uint32_t size : size of the preload region, should not exceed the size of ICache.
 *
 * @param uint32_t order : the preload order, 0 for positive, other for negative
 *
 * @return uint32_t : 0 for ICache not auto preload before manual preload.
 */
//go:linkname CacheStartICachePreload C.Cache_Start_ICache_Preload
func CacheStartICachePreload(addr c.Uint32T, size c.Uint32T, order c.Uint32T) c.Uint32T

/**
 * @brief Return if the ICache manual preload done.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return uint32_t : 0 for ICache manual preload not done.
 */
//go:linkname CacheICachePreloadDone C.Cache_ICache_Preload_Done
func CacheICachePreloadDone() c.Uint32T

/**
 * @brief End the ICache manual preload to resume auto preload of ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t autoload : 0 for ICache not auto preload before manual preload.
 *
 * @return None
 */
//go:linkname CacheEndICachePreload C.Cache_End_ICache_Preload
func CacheEndICachePreload(autoload c.Uint32T)

/**
 * @brief Start an DCache manual preload, will suspend auto preload of DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t addr : start address of the preload region.
 *
 * @param uint32_t size : size of the preload region, should not exceed the size of DCache.
 *
 * @param uint32_t order : the preload order, 0 for positive, other for negative
 *
 * @return uint32_t : 0 for DCache not auto preload before manual preload.
 */
//go:linkname CacheStartDCachePreload C.Cache_Start_DCache_Preload
func CacheStartDCachePreload(addr c.Uint32T, size c.Uint32T, order c.Uint32T) c.Uint32T

/**
 * @brief Return if the DCache manual preload done.
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return uint32_t : 0 for DCache manual preload not done.
 */
//go:linkname CacheDCachePreloadDone C.Cache_DCache_Preload_Done
func CacheDCachePreloadDone() c.Uint32T

/**
 * @brief End the DCache manual preload to resume auto preload of DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t autoload : 0 for DCache not auto preload before manual preload.
 *
 * @return None
 */
//go:linkname CacheEndDCachePreload C.Cache_End_DCache_Preload
func CacheEndDCachePreload(autoload c.Uint32T)

/**
 * @brief Config autoload parameters of ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param struct autoload_config * config : autoload parameters.
 *
 * @return None
 */
// llgo:link (*AutoloadConfig).CacheConfigICacheAutoload C.Cache_Config_ICache_Autoload
func (recv_ *AutoloadConfig) CacheConfigICacheAutoload() {
}

/**
 * @brief Config region autoload parameters of ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param struct autoload_region_config * config : region autoload parameters.
 *
 * @return ESP_ROM_ERR_INVALID_ARG : invalid param, 0 : success
 */
// llgo:link (*AutoloadRegionConfig).CacheConfigICacheRegionAutoload C.Cache_Config_ICache_Region_Autoload
func (recv_ *AutoloadRegionConfig) CacheConfigICacheRegionAutoload() c.Int {
	return 0
}

/**
 * @brief Enable auto preload for ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param None
 *
 * @return None
 */
//go:linkname CacheEnableICacheAutoload C.Cache_Enable_ICache_Autoload
func CacheEnableICacheAutoload()

/**
 * @brief Disable auto preload for ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param None
 *
 * @return None
 */
//go:linkname CacheDisableICacheAutoload C.Cache_Disable_ICache_Autoload
func CacheDisableICacheAutoload()

/**
 * @brief Config autoload parameters of DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param struct autoload_config * config : autoload parameters.
 *
 * @return None
 */
// llgo:link (*AutoloadConfig).CacheConfigDCacheAutoload C.Cache_Config_DCache_Autoload
func (recv_ *AutoloadConfig) CacheConfigDCacheAutoload() {
}

/**
 * @brief Config region autoload parameters of DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param struct autoload_region_config * config : region autoload parameters.
 *
 * @return ESP_ROM_ERR_INVALID_ARG : invalid param, 0 : success
 */
// llgo:link (*AutoloadRegionConfig).CacheConfigDCacheRegionAutoload C.Cache_Config_DCache_Region_Autoload
func (recv_ *AutoloadRegionConfig) CacheConfigDCacheRegionAutoload() c.Int {
	return 0
}

/**
 * @brief Enable auto preload for DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param None
 *
 * @return None
 */
//go:linkname CacheEnableDCacheAutoload C.Cache_Enable_DCache_Autoload
func CacheEnableDCacheAutoload()

/**
 * @brief Disable auto preload for DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param None
 *
 * @return None
 */
//go:linkname CacheDisableDCacheAutoload C.Cache_Disable_DCache_Autoload
func CacheDisableDCacheAutoload()

/**
 * @brief Config a group of prelock parameters of ICache.
 *        Please do not call this function in your SDK application.
 *
 * @param struct lock_config * config : a group of lock parameters.
 *
 * @return None
 */
// llgo:link (*LockConfig).CacheEnableICachePreLock C.Cache_Enable_ICache_PreLock
func (recv_ *LockConfig) CacheEnableICachePreLock() {
}

/**
 * @brief Disable a group of prelock parameters for ICache.
 *        However, the locked data will not be released.
 *        Please do not call this function in your SDK application.
 *
 * @param uint16_t group : 0 for group0, 1 for group1.
 *
 * @return None
 */
//go:linkname CacheDisableICachePreLock C.Cache_Disable_ICache_PreLock
func CacheDisableICachePreLock(group c.Uint16T)

/**
 * @brief Lock the cache items for ICache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in ICache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr: start address to lock
 *
 * @param  uint32_t items: cache lines to lock, items * cache_line_size should not exceed the bus address size(16MB/32MB/64MB)
 *
 * @return None
 */
//go:linkname CacheLockICacheItems C.Cache_Lock_ICache_Items
func CacheLockICacheItems(addr c.Uint32T, items c.Uint32T)

/**
 * @brief Unlock the cache items for ICache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in ICache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr: start address to unlock
 *
 * @param  uint32_t items: cache lines to unlock, items * cache_line_size should not exceed the bus address size(16MB/32MB/64MB)
 *
 * @return None
 */
//go:linkname CacheUnlockICacheItems C.Cache_Unlock_ICache_Items
func CacheUnlockICacheItems(addr c.Uint32T, items c.Uint32T)

/**
 * @brief Config a group of prelock parameters of DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param struct lock_config * config : a group of lock parameters.
 *
 * @return None
 */
// llgo:link (*LockConfig).CacheEnableDCachePreLock C.Cache_Enable_DCache_PreLock
func (recv_ *LockConfig) CacheEnableDCachePreLock() {
}

/**
 * @brief Disable a group of prelock parameters for DCache.
 *        However, the locked data will not be released.
 *        Please do not call this function in your SDK application.
 *
 * @param uint16_t group : 0 for group0, 1 for group1.
 *
 * @return None
 */
//go:linkname CacheDisableDCachePreLock C.Cache_Disable_DCache_PreLock
func CacheDisableDCachePreLock(group c.Uint16T)

/**
 * @brief Lock the cache items for DCache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in DCache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr: start address to lock
 *
 * @param  uint32_t items: cache lines to lock, items * cache_line_size should not exceed the bus address size(16MB/32MB/64MB)
 *
 * @return None
 */
//go:linkname CacheLockDCacheItems C.Cache_Lock_DCache_Items
func CacheLockDCacheItems(addr c.Uint32T, items c.Uint32T)

/**
 * @brief Unlock the cache items for DCache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in DCache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t addr: start address to unlock
 *
 * @param  uint32_t items: cache lines to unlock, items * cache_line_size should not exceed the bus address size(16MB/32MB/64MB)
 *
 * @return None
 */
//go:linkname CacheUnlockDCacheItems C.Cache_Unlock_DCache_Items
func CacheUnlockDCacheItems(addr c.Uint32T, items c.Uint32T)

/**
 * @brief Lock the cache items in tag memory for ICache or DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t addr : start address of lock region.
 *
 * @param uint32_t size : size of lock region.
 *
 * @return 0 for success
 *         1 for invalid argument
 */
//go:linkname CacheLockAddr C.Cache_Lock_Addr
func CacheLockAddr(addr c.Uint32T, size c.Uint32T) c.Int

/**
 * @brief Unlock the cache items in tag memory for ICache or DCache.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t addr : start address of unlock region.
 *
 * @param uint32_t size : size of unlock region.
 *
 * @return 0 for success
 *         1 for invalid argument
 */
//go:linkname CacheUnlockAddr C.Cache_Unlock_Addr
func CacheUnlockAddr(addr c.Uint32T, size c.Uint32T) c.Int

/**
 * @brief Disable ICache access for the cpu.
 *        This operation will make all ICache tag memory invalid, CPU can't access ICache, ICache will keep idle.
 *        Please do not call this function in your SDK application.
 *
 * @return uint32_t : auto preload enabled before
 */
//go:linkname CacheDisableICache C.Cache_Disable_ICache
func CacheDisableICache() c.Uint32T

/**
 * @brief Enable ICache access for the cpu.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t autoload : ICache will preload then.
 *
 * @return None
 */
//go:linkname CacheEnableICache C.Cache_Enable_ICache
func CacheEnableICache(autoload c.Uint32T)

/**
 * @brief Disable DCache access for the cpu.
 *        This operation will make all DCache tag memory invalid, CPU can't access DCache, DCache will keep idle
 *        Please do not call this function in your SDK application.
 *
 * @return uint32_t : auto preload enabled before
 */
//go:linkname CacheDisableDCache C.Cache_Disable_DCache
func CacheDisableDCache() c.Uint32T

/**
 * @brief Enable DCache access for the cpu.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t autoload : DCache will preload then.
 *
 * @return None
 */
//go:linkname CacheEnableDCache C.Cache_Enable_DCache
func CacheEnableDCache(autoload c.Uint32T)

/**
 * @brief Suspend ICache access for the cpu.
 *        The ICache tag memory is still there, CPU can't access ICache, ICache will keep idle.
 *        Please do not change MMU, cache mode or tag memory(tag memory can be changed in some special case).
 *        Please do not call this function in your SDK application.
 *
 * @param  None
 *
 * @return uint32_t : auto preload enabled before
 */
//go:linkname CacheSuspendICache C.Cache_Suspend_ICache
func CacheSuspendICache() c.Uint32T

/**
 * @brief Resume ICache access for the cpu.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t autoload : ICache will preload then.
 *
 * @return None
 */
//go:linkname CacheResumeICache C.Cache_Resume_ICache
func CacheResumeICache(autoload c.Uint32T)

/**
  * @brief Suspend DCache access for the cpu.
  *        The ICache tag memory is still there, CPU can't access DCache, DCache will keep idle.
  ×        Please do not change MMU, cache mode or tag memory(tag memory can be changed in some special case).
  *        Please do not call this function in your SDK application.
  *
  * @param  None
  *
  * @return uint32_t : auto preload enabled before
*/
//go:linkname CacheSuspendDCache C.Cache_Suspend_DCache
func CacheSuspendDCache() c.Uint32T

/**
 * @brief Resume DCache access for the cpu.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t autoload : DCache will preload then.
 *
 * @return None
 */
//go:linkname CacheResumeDCache C.Cache_Resume_DCache
func CacheResumeDCache(autoload c.Uint32T)

/**
 * @brief Get ICache cache line size
 *
 * @param  None
 *
 * @return uint32_t: 16, 32, 64 Byte
 */
//go:linkname CacheGetICacheLineSize C.Cache_Get_ICache_Line_Size
func CacheGetICacheLineSize() c.Uint32T

/**
 * @brief Get DCache cache line size
 *
 * @param  None
 *
 * @return uint32_t: 16, 32, 64 Byte
 */
//go:linkname CacheGetDCacheLineSize C.Cache_Get_DCache_Line_Size
func CacheGetDCacheLineSize() c.Uint32T

/**
 * @brief Set default mode from boot, 8KB ICache, 16Byte cache line size.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname CacheSetDefaultMode C.Cache_Set_Default_Mode
func CacheSetDefaultMode()

/**
 * @brief Set default mode from boot, 8KB ICache, 16Byte cache line size.
 *
 * @param None
 *
 * @return None
 */
//go:linkname CacheEnableDefalutICacheMode C.Cache_Enable_Defalut_ICache_Mode
func CacheEnableDefalutICacheMode()

/**
 * @brief Occupy the cache items for DCache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in DCache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t addr : start address of occupy region
 *
 * @param uint32_t items : cache lines to occupy, items * cache_line_size should not exceed the cache_size
 *
 * @return None
 */
//go:linkname CacheOccupyItems C.Cache_Occupy_Items
func CacheOccupyItems(addr c.Uint32T, items c.Uint32T)

/**
 * @brief Occupy the cache addr for DCache.
 *        Operation will be done CACHE_LINE_SIZE aligned.
 *        If the region is not in DCache addr room, nothing will be done.
 *        Please do not call this function in your SDK application.
 *
 * @param uint32_t addr : start address of occupy region
 *
 * @param uint32_t size : size of occupy region, size should not exceed the cache_size
 */
//go:linkname CacheOccupyAddr C.Cache_Occupy_Addr
func CacheOccupyAddr(addr c.Uint32T, size c.Uint32T) c.Int

/**
 * @brief Enable freeze for ICache.
 *        Any miss request will be rejected, including cpu miss and preload/autoload miss.
 *        Please do not call this function in your SDK application.
 *
 * @param cache_freeze_mode_t mode : 0 for assert busy 1 for assert hit
 *
 * @return None
 */
// llgo:link CacheFreezeModeT.CacheFreezeICacheEnable C.Cache_Freeze_ICache_Enable
func (recv_ CacheFreezeModeT) CacheFreezeICacheEnable() {
}

/**
 * @brief Disable freeze for ICache.
 *        Please do not call this function in your SDK application.
 *
 * @return None
 */
//go:linkname CacheFreezeICacheDisable C.Cache_Freeze_ICache_Disable
func CacheFreezeICacheDisable()

/**
 * @brief Enable freeze for DCache.
 *        Any miss request will be rejected, including cpu miss and preload/autoload miss.
 *        Please do not call this function in your SDK application.
 *
 * @param cache_freeze_mode_t mode : 0 for assert busy 1 for assert hit
 *
 * @return None
 */
// llgo:link CacheFreezeModeT.CacheFreezeDCacheEnable C.Cache_Freeze_DCache_Enable
func (recv_ CacheFreezeModeT) CacheFreezeDCacheEnable() {
}

/**
 * @brief Disable freeze for DCache.
 *        Please do not call this function in your SDK application.
 *
 * @return None
 */
//go:linkname CacheFreezeDCacheDisable C.Cache_Freeze_DCache_Disable
func CacheFreezeDCacheDisable()

/**
 * @brief Travel tag memory to run a call back function.
 *        ICache and DCache are suspend when doing this.
 *        The callback will get the parameter tag_group_info, which will include a group of tag memory addresses and cache memory addresses.
 *        Please do not call this function in your SDK application.
 *
 * @param  struct cache_mode * mode : the cache to check and the cache mode.
 *
 * @param  uint32_t filter_addr : only the cache lines which may include the filter_address will be returned to the call back function.
 *                                0 for do not filter, all cache lines will be returned.
 *
 * @param  void (* process)(struct tag_group_info *) : call back function, which may be called many times, a group(the addresses in the group are in the same position in the cache ways) a time.
 *
 * @return None
 */
// llgo:link (*CacheMode).CacheTravelTagMemory C.Cache_Travel_Tag_Memory
func (recv_ *CacheMode) CacheTravelTagMemory(filter_addr c.Uint32T, process func(*TagGroupInfo)) {
}

/**
 * @brief Travel tag memory to run a call back function, using 2nd tag registers.
 *        ICache and DCache are suspend when doing this.
 *        The callback will get the parameter tag_group_info, which will include a group of tag memory addresses and cache memory addresses.
 *        Please do not call this function in your SDK application.
 *
 * @param  struct cache_mode * mode : the cache to check and the cache mode.
 *
 * @param  uint32_t filter_addr : only the cache lines which may include the filter_address will be returned to the call back function.
 *                                0 for do not filter, all cache lines will be returned.
 *
 * @param  void (* process)(struct tag_group_info *) : call back function, which may be called many times, a group(the addresses in the group are in the same position in the cache ways) a time.
 *
 * @return None
 */
// llgo:link (*CacheMode).CacheTravelTagMemory2 C.Cache_Travel_Tag_Memory2
func (recv_ *CacheMode) CacheTravelTagMemory2(filter_addr c.Uint32T, process func(*TagGroupInfo)) {
}

/**
 * @brief Get the virtual address from cache mode, cache tag and the virtual address offset of cache ways.
 *        Please do not call this function in your SDK application.
 *
 * @param  struct cache_mode * mode : the cache to calculate the virtual address and the cache mode.
 *
 * @param  uint32_t tag : the tag part of a tag item, 12-14 bits.
 *
 * @param  uint32_t addr_offset : the virtual address offset of the cache ways.
 *
 * @return uint32_t : the virtual address.
 */
// llgo:link (*CacheMode).CacheGetVirtualAddr C.Cache_Get_Virtual_Addr
func (recv_ *CacheMode) CacheGetVirtualAddr(tag c.Uint32T, vaddr_offset c.Uint32T) c.Uint32T {
	return 0
}

/**
 * @brief Get cache memory block base address.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t icache : 0 for dcache, other for icache.
 *
 * @param  uint32_t bank_no : 0 ~ 3 bank.
 *
 * @return uint32_t : the cache memory block base address, 0 if the block not used.
 */
//go:linkname CacheGetMemoryBaseAddr C.Cache_Get_Memory_BaseAddr
func CacheGetMemoryBaseAddr(icache c.Uint32T, bank_no c.Uint32T) c.Uint32T

/**
 * @brief Get the cache memory address from cache mode, cache memory offset and the virtual address offset of cache ways.
 *        Please do not call this function in your SDK application.
 *
 * @param  struct cache_mode * mode : the cache to calculate the virtual address and the cache mode.
 *
 * @param  uint32_t cache_memory_offset : the cache memory offset of the whole cache (ICache or DCache) for the cache line.
 *
 * @param  uint32_t addr_offset : the virtual address offset of the cache ways.
 *
 * @return uint32_t : the virtual address.
 */
// llgo:link (*CacheMode).CacheGetMemoryAddr C.Cache_Get_Memory_Addr
func (recv_ *CacheMode) CacheGetMemoryAddr(cache_memory_offset c.Uint32T, vaddr_offset c.Uint32T) c.Uint32T {
	return 0
}

/**
 * @brief Get the cache memory value by DRAM address.
 *        Please do not call this function in your SDK application.
 *
 * @param  uint32_t cache_memory_addr : DRAM address for the cache memory, should be 4 byte aligned for IBus address.
 *
 * @return uint32_t : the word value of the address.
 */
//go:linkname CacheGetMemoryValue C.Cache_Get_Memory_value
func CacheGetMemoryValue(cache_memory_addr c.Uint32T) c.Uint32T

/**
 * @brief Get the cache MMU IROM end address.
 *        Please do not call this function in your SDK application.
 *
 * @param  void
 *
 * @return uint32_t : the word value of the address.
 */
//go:linkname CacheGetIROMMMUEnd C.Cache_Get_IROM_MMU_End
func CacheGetIROMMMUEnd() c.Uint32T

/**
 * @brief Get the cache MMU DROM end address.
 *        Please do not call this function in your SDK application.
 *
 * @param  void
 *
 * @return uint32_t : the word value of the address.
 */
//go:linkname CacheGetDROMMMUEnd C.Cache_Get_DROM_MMU_End
func CacheGetDROMMMUEnd() c.Uint32T

/**
 * @brief Configure cache MMU page size according to instruction and rodata size
 *
 * @param irom_size The instruction cache MMU page size
 * @param drom_size The rodata data cache MMU page size
 */
//go:linkname CacheSetIDROMMMUSize C.Cache_Set_IDROM_MMU_Size
func CacheSetIDROMMMUSize(irom_size c.Uint32T, drom_size c.Uint32T)

/**
 * @brief Configure cache MMU page information
 *
 * @param instr_page_num The instruction cache MMU page num
 * @param rodata_page_num The rodata cache MMU page num
 * @param rodata_start The rodata start cache address
 * @param rodata_end The rodata end cache address
 * @param i_off The offset of instruction when instruction copied from flash to xip_psram
 * @param ro_off The offset of rodata when rodata copied from flash to xip_psram
 */
//go:linkname CacheSetIDROMMMUInfo C.Cache_Set_IDROM_MMU_Info
func CacheSetIDROMMMUInfo(instr_page_num c.Uint32T, rodata_page_num c.Uint32T, rodata_start c.Uint32T, rodata_end c.Uint32T, i_off c.Int, ro_off c.Int)

/**
 * @brief Used by SPI flash mmap
 *
 */
//go:linkname Flash2spiramInstructionOffset C.flash2spiram_instruction_offset
func Flash2spiramInstructionOffset() c.Int

//go:linkname Flash2spiramRodataOffset C.flash2spiram_rodata_offset
func Flash2spiramRodataOffset() c.Int

//go:linkname FlashInstrRodataStartPage C.flash_instr_rodata_start_page
func FlashInstrRodataStartPage(bus c.Uint32T) c.Uint32T

//go:linkname FlashInstrRodataEndPage C.flash_instr_rodata_end_page
func FlashInstrRodataEndPage(bus c.Uint32T) c.Uint32T
