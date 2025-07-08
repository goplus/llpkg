package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * MMU Hal layer initialisation
 */
//go:linkname MmuHalInit C.mmu_hal_init
func MmuHalInit()

/**
 * Unmap all the MMU table. After this all external memory vaddr are not available
 */
//go:linkname MmuHalUnmapAll C.mmu_hal_unmap_all
func MmuHalUnmapAll()

/**
 * Helper functions to convert the MMU page numbers into bytes. e.g.:
 * - When MMU page size is 16KB, page_num = 2 will be converted into 32KB
 * - When MMU page size is 32KB, page_num = 2 will be converted into 64KB
 *
 * @param mmu_id    MMU ID
 * @param page_num  page numbers
 *
 * @return
 *         length in byte
 */
//go:linkname MmuHalPagesToBytes C.mmu_hal_pages_to_bytes
func MmuHalPagesToBytes(mmu_id c.Uint32T, page_num c.Uint32T) c.Uint32T

/**
 * Helper functions to convert bytes into MMU page numbers. e.g.:
 * - When MMU page size is 16KB, bytes = 64KB will be converted into 4 pages
 * - When MMU page size is 32KB, bytes = 64KB will be converted into 2 pages
 *
 * @param mmu_id    MMU ID
 * @param bytes     length in byte
 *
 * @return
 *         length in CONFIG_MMU_PAGE_SIZE
 */
//go:linkname MmuHalBytesToPages C.mmu_hal_bytes_to_pages
func MmuHalBytesToPages(mmu_id c.Uint32T, bytes c.Uint32T) c.Uint32T

/**
 * To map a virtual address block to a physical memory block
 *
 * @param mmu_id       MMU ID
 * @param mem_type     physical memory type, see `mmu_target_t`
 * @param vaddr        start virtual address to be mapped
 * @param paddr        start physical address to be mapped
 * @param len          length to be mapped, in bytes
 * @param[out] out_len actual mapped length
 *
 * @note vaddr and paddr should be aligned with the mmu page size, see CONFIG_MMU_PAGE_SIZE
 */
//go:linkname MmuHalMapRegion C.mmu_hal_map_region
func MmuHalMapRegion(mmu_id c.Uint32T, mem_type MmuTargetT, vaddr c.Uint32T, paddr c.Uint32T, len c.Uint32T, out_len *c.Uint32T)

/**
 * To unmap a virtual address block that is mapped to a physical memory block previously
 *
 * @param[in] mmu_id  MMU ID
 * @param[in] vaddr   start virtual address
 * @param[in] len     length to be unmapped, in bytes
 */
//go:linkname MmuHalUnmapRegion C.mmu_hal_unmap_region
func MmuHalUnmapRegion(mmu_id c.Uint32T, vaddr c.Uint32T, len c.Uint32T)

/**
 * Convert virtual address to physical address
 *
 * @param mmu_id           MMU ID
 * @param vaddr            virtual address
 * @param[out] out_paddr   physical address
 * @param[out] out_target  Indicating the vaddr/paddr is mapped on which target, see `mmu_target_t`
 *
 * @return
 *        - true: virtual address is valid
 *        - false: virtual address isn't valid
 */
//go:linkname MmuHalVaddrToPaddr C.mmu_hal_vaddr_to_paddr
func MmuHalVaddrToPaddr(mmu_id c.Uint32T, vaddr c.Uint32T, out_paddr *c.Uint32T, out_target *MmuTargetT) bool

/**
 * Convert physical address to virtual address
 *
 * @note This function can only find the first match virtual address.
 *       However it is possible that a physical address is mapped to multiple virtual addresses.
 *
 * @param mmu_id          MMU ID
 * @param paddr           physical address
 * @param target          physical memory target, see `mmu_target_t`
 * @param type            virtual address type, could be instruction or data
 * @param[out] out_vaddr  virtual address
 *
 * @return
 *        - true: found a matched vaddr
 *        - false: not found a matched vaddr
 */
//go:linkname MmuHalPaddrToVaddr C.mmu_hal_paddr_to_vaddr
func MmuHalPaddrToVaddr(mmu_id c.Uint32T, paddr c.Uint32T, target MmuTargetT, type_ MmuVaddrT, out_vaddr *c.Uint32T) bool

/**
 * Check if the vaddr region is valid
 *
 * @param mmu_id      MMU ID
 * @param vaddr_start start of the virtual address
 * @param len         length, in bytes
 * @param type        virtual address type, could be instruction type or data type. See `mmu_vaddr_t`
 *
 * @return
 *         True for valid
 */
//go:linkname MmuHalCheckValidExtVaddrRegion C.mmu_hal_check_valid_ext_vaddr_region
func MmuHalCheckValidExtVaddrRegion(mmu_id c.Uint32T, vaddr_start c.Uint32T, len c.Uint32T, type_ MmuVaddrT) bool
