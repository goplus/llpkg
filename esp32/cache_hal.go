package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Cache init and cache hal context init
 */
//go:linkname CacheHalInit C.cache_hal_init
func CacheHalInit()

/**
 * @brief Disable Cache
 *
 * Disable the ICache or DCache or both, of a certain level or all levels.
 * All the items in the corresponding Cache(s) will be invalideated.
 * Next request to these items will trigger a transaction to the physical memory
 *
 * @note If the autoload feature is enabled, this API will return until the ICache autoload is disabled.
 *
 * @param cache_level  Level of the Cache(s)
 * @param type         see `cache_type_t`
 */
//go:linkname CacheHalDisable C.cache_hal_disable
func CacheHalDisable(cache_level c.Uint32T, type_ CacheTypeT)

/**
 * @brief Enable Cache
 *
 * Enable the ICache or DCache or both, of a certain level or all levels.
 *
 * @param cache_level  Level of the Cache(s)
 * @param type         see `cache_type_t`
 */
//go:linkname CacheHalEnable C.cache_hal_enable
func CacheHalEnable(cache_level c.Uint32T, type_ CacheTypeT)

/**
 * @brief Suspend Cache
 *
 * Suspend the ICache or DCache or both, of a certain level or all levels.
 * This API suspends the CPU access to cache for a while, without invalidation.
 *
 * @param cache_level  Level of the Cache(s)
 * @param type         see `cache_type_t`
 */
//go:linkname CacheHalSuspend C.cache_hal_suspend
func CacheHalSuspend(cache_level c.Uint32T, type_ CacheTypeT)

/**
 * @brief Resume Cache
 *
 * Resume the ICache or DCache or both, of a certain level or all levels.
 *
 * @param cache_level  Level of the Cache(s)
 * @param type         see `cache_type_t`
 */
//go:linkname CacheHalResume C.cache_hal_resume
func CacheHalResume(cache_level c.Uint32T, type_ CacheTypeT)

/**
 * @brief Check if corresponding cache is enabled or not
 *
 * @param cache_level  Level of the Cache(s)
 * @param type         see `cache_type_t`
 *
 * @return true: enabled; false: disabled
 */
//go:linkname CacheHalIsCacheEnabled C.cache_hal_is_cache_enabled
func CacheHalIsCacheEnabled(cache_level c.Uint32T, type_ CacheTypeT) bool

/**
 * @brief Invalidate Cache supported addr
 *
 * Invalidate a Cache item for either ICache or DCache.
 *
 * @param vaddr  Start address of the region to be invalidated
 * @param size   Size of the region to be invalidated
 *
 * @return       True for valid address. No operation if invalid
 */
//go:linkname CacheHalInvalidateAddr C.cache_hal_invalidate_addr
func CacheHalInvalidateAddr(vaddr c.Uint32T, size c.Uint32T) bool

/**
 * @brief Writeback Cache supported addr
 *
 * Writeback the DCache item to external memory
 *
 * @param vaddr  Start address of the region to writeback
 * @param size   Size of the region to writeback
 *
 * @return       True for valid address. No operation if invalid
 */
//go:linkname CacheHalWritebackAddr C.cache_hal_writeback_addr
func CacheHalWritebackAddr(vaddr c.Uint32T, size c.Uint32T) bool

/**
 * @brief Freeze Cache
 *
 * Freeze cache, CPU access to cache will be suspended, until the cache is unfrozen.
 *
 * @param cache_level  Level of the Cache(s)
 * @param type         see `cache_type_t`
 */
//go:linkname CacheHalFreeze C.cache_hal_freeze
func CacheHalFreeze(cache_level c.Uint32T, type_ CacheTypeT)

/**
 * @brief Unfreeze cache
 *
 * Unfreeze cache, CPU access to cache will be restored
 *
 * @param cache_level  Level of the Cache(s)
 * @param type         see `cache_type_t`
 */
//go:linkname CacheHalUnfreeze C.cache_hal_unfreeze
func CacheHalUnfreeze(cache_level c.Uint32T, type_ CacheTypeT)

/**
 * @brief Get cache line size, in bytes
 *
 * @param cache_level  Level of the Cache(s)
 * @param type         see `cache_type_t`
 *
 * @return cache line size, in bytes. 0 stands for no such cache in this type or level
 */
//go:linkname CacheHalGetCacheLineSize C.cache_hal_get_cache_line_size
func CacheHalGetCacheLineSize(cache_level c.Uint32T, type_ CacheTypeT) c.Uint32T

/**
 * @brief Get Cache level and the ID of the vaddr
 *
 * @param vaddr_start       virtual address start
 * @param len               vaddr length
 * @param out_level         cache level
 * @param out_id            cache id
 *
 * @return true for valid, false for invalid addr or null pointer
 */
//go:linkname CacheHalVaddrToCacheLevelId C.cache_hal_vaddr_to_cache_level_id
func CacheHalVaddrToCacheLevelId(vaddr_start c.Uint32T, len c.Uint32T, out_level *c.Uint32T, out_id *c.Uint32T) bool
