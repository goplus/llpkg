package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MemTypeProtT c.Int

const (
	MEMPROT_NONE            MemTypeProtT = 0
	MEMPROT_IRAM0_SRAM      MemTypeProtT = 1
	MEMPROT_DRAM0_SRAM      MemTypeProtT = 2
	MEMPROT_IRAM0_RTCFAST   MemTypeProtT = 4
	MEMPROT_DRAM0_RTCFAST   MemTypeProtT = 8
	MEMPROT_PERI1_RTCSLOW   MemTypeProtT = 16
	MEMPROT_PERI2_RTCSLOW_0 MemTypeProtT = 32
	MEMPROT_PERI2_RTCSLOW_1 MemTypeProtT = 64
	MEMPROT_ALL             MemTypeProtT = -1
)

/**
 * @brief Returns splitting address for required memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 *
 * @return Splitting address for the memory region required.
 * The address is given by region-specific global symbol exported from linker script,
 * it is not read out from related configuration register.
 */
// llgo:link MemTypeProtT.EspMemprotGetSplitAddr C.esp_memprot_get_split_addr
func (recv_ MemTypeProtT) EspMemprotGetSplitAddr() *c.Uint32T {
	return nil
}

/**
 * @brief Initializes illegal memory access control for required memory section.
 *
 * All memory access interrupts share ETS_MEMACCESS_ERR_INUM input channel, it is caller's
 * responsibility to properly detect actual intr. source as well as possible prioritization in case
 * of multiple source reported during one intr.handling routine run
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)\
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotIntrInit C.esp_memprot_intr_init
func (recv_ MemTypeProtT) EspMemprotIntrInit() EspErrT {
	return 0
}

/**
 * @brief Enable/disable the memory protection interrupt
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param enable enable/disable
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotIntrEna C.esp_memprot_intr_ena
func (recv_ MemTypeProtT) EspMemprotIntrEna(enable bool) EspErrT {
	return 0
}

/**
 * @brief Sets a request for clearing interrupt-on flag for specified memory region (register write)
 *
 * @note When called without actual interrupt-on flag set, subsequent occurrence of related interrupt is ignored.
 * Should be used only after the real interrupt appears, typically as the last step in interrupt handler's routine.
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotClearIntr C.esp_memprot_clear_intr
func (recv_ MemTypeProtT) EspMemprotClearIntr() EspErrT {
	return 0
}

/**
 * @brief Detects which memory protection interrupt is active
 *
 * @note Check order
 *          MEMPROT_IRAM0_SRAM
 *          MEMPROT_IRAM0_RTCFAST
 *          MEMPROT_DRAM0_SRAM
 *          MEMPROT_DRAM0_RTCFAST
 *
 * @return Memory protection area type (see mem_type_prot_t enum)
 */
//go:linkname EspMemprotGetActiveIntrMemtype C.esp_memprot_get_active_intr_memtype
func EspMemprotGetActiveIntrMemtype() MemTypeProtT

/**
 * @brief Gets interrupt status register contents for specified memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param fault_reg_val Contents of status register
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetFaultReg C.esp_memprot_get_fault_reg
func (recv_ MemTypeProtT) EspMemprotGetFaultReg(fault_reg_val *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Get details of given interrupt status
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param faulting_address Faulting address causing the interrupt [out]
 * @param op_type Operation being processed at the faulting address [out]
 *               IRAM0: 0 - read, 1 - write
 *               DRAM0: 0 - read, 1 - write
 * @param op_subtype Additional info for op_type [out]
 *               IRAM0: 0 - instruction segment access, 1 - data segment access
 *               DRAM0: 0 - non-atomic operation, 1 - atomic operation
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetFaultStatus C.esp_memprot_get_fault_status
func (recv_ MemTypeProtT) EspMemprotGetFaultStatus(faulting_address **c.Uint32T, op_type *c.Uint32T, op_subtype *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Gets string representation of required memory region identifier
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 *
 * @return mem_type as string
 */
// llgo:link MemTypeProtT.EspMemprotTypeToStr C.esp_memprot_type_to_str
func (recv_ MemTypeProtT) EspMemprotTypeToStr() *c.Char {
	return nil
}

/**
 * @brief Detects whether any of the interrupt locks is active (requires digital system reset to unlock)
 *
 * @return true/false
 */
//go:linkname EspMemprotIsLockedAny C.esp_memprot_is_locked_any
func EspMemprotIsLockedAny() bool

/**
 * @brief Sets lock for specified memory region.
 *
 * Locks can be unlocked only by digital system reset
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotSetLock C.esp_memprot_set_lock
func (recv_ MemTypeProtT) EspMemprotSetLock() EspErrT {
	return 0
}

/**
 * @brief Gets lock status for required memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param locked Settings locked: true/false (locked/unlocked)
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetLock C.esp_memprot_get_lock
func (recv_ MemTypeProtT) EspMemprotGetLock(locked *bool) EspErrT {
	return 0
}

/**
 * @brief Gets permission control configuration register contents for required memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param conf_reg_val Permission control register contents
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetConfReg C.esp_memprot_get_conf_reg
func (recv_ MemTypeProtT) EspMemprotGetConfReg(conf_reg_val *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Gets interrupt permission settings for unified management block
 *
 * Gets interrupt permission settings register contents for required memory region, returns settings for unified management blocks
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param perm_reg Permission settings register contents
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetPermUniReg C.esp_memprot_get_perm_uni_reg
func (recv_ MemTypeProtT) EspMemprotGetPermUniReg(perm_reg *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Gets interrupt permission settings for split management block
 *
 * Gets interrupt permission settings register contents for required memory region (unified management blocks)
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @return split_reg Unified management settings register contents
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetPermSplitReg C.esp_memprot_get_perm_split_reg
func (recv_ MemTypeProtT) EspMemprotGetPermSplitReg(split_reg *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Detects whether any of the memory protection interrupts is enabled
 *
 * @return true/false
 */
//go:linkname EspMemprotIsIntrEnaAny C.esp_memprot_is_intr_ena_any
func EspMemprotIsIntrEnaAny() bool

/**
 * @brief Gets interrupt-enabled flag for given memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param enable_bit Interrupt-enabled flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetIntrEnaBit C.esp_memprot_get_intr_ena_bit
func (recv_ MemTypeProtT) EspMemprotGetIntrEnaBit(enable_bit *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Gets interrupt-active flag for given memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param intr_on_bit Interrupt-active flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure */
// llgo:link MemTypeProtT.EspMemprotGetIntrOnBit C.esp_memprot_get_intr_on_bit
func (recv_ MemTypeProtT) EspMemprotGetIntrOnBit(intr_on_bit *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Gets interrupt-clear request flag for given memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param clear_bit Interrupt-clear request flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetIntrClrBit C.esp_memprot_get_intr_clr_bit
func (recv_ MemTypeProtT) EspMemprotGetIntrClrBit(clear_bit *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Gets read permission value for specified block and memory region
 *
 * Returns read permission bit value for required unified-management block (0-3) in given memory region.
 * Applicable to all memory types.
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param block Memory block identifier (0-3)
 * @param read_bit Read permission value for required block
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetUniBlockReadBit C.esp_memprot_get_uni_block_read_bit
func (recv_ MemTypeProtT) EspMemprotGetUniBlockReadBit(block c.Uint32T, read_bit *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Gets write permission value for specified block and memory region
 *
 * Returns write permission bit value for required unified-management block (0-3) in given memory region.
 * Applicable to all memory types.
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param block Memory block identifier (0-3)
 * @param write_bit Write permission value for required block
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetUniBlockWriteBit C.esp_memprot_get_uni_block_write_bit
func (recv_ MemTypeProtT) EspMemprotGetUniBlockWriteBit(block c.Uint32T, write_bit *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Gets execute permission value for specified block and memory region
 *
 * Returns execute permission bit value for required unified-management block (0-3) in given memory region.
 * Applicable only to IRAM memory types
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param block Memory block identifier (0-3)
 * @param exec_bit Execute permission value for required block
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetUniBlockExecBit C.esp_memprot_get_uni_block_exec_bit
func (recv_ MemTypeProtT) EspMemprotGetUniBlockExecBit(block c.Uint32T, exec_bit *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Sets permissions for specified block in DRAM region
 *
 * Sets Read and Write permission for specified unified-management block (0-3) in given memory region.
 * Applicable only to DRAM memory types
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param block Memory block identifier (0-3)
 * @param write_perm Write permission flag
 * @param read_perm Read permission flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotSetUniBlockPermDram C.esp_memprot_set_uni_block_perm_dram
func (recv_ MemTypeProtT) EspMemprotSetUniBlockPermDram(block c.Uint32T, write_perm bool, read_perm bool) EspErrT {
	return 0
}

/**
 * @brief Sets permissions for high and low memory segment in DRAM region
 *
 * Sets Read and Write permission for both low and high memory segments given by splitting address.
 * The splitting address must be equal to or higher then beginning of block 5
 * Applicable only to DRAM memory types
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param split_addr Address to split the memory region to lower and higher segment
 * @param lw Low segment Write permission flag
 * @param lr Low segment Read permission flag
 * @param hw High segment Write permission flag
 * @param hr High segment Read permission flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotSetProtDram C.esp_memprot_set_prot_dram
func (recv_ MemTypeProtT) EspMemprotSetProtDram(split_addr *c.Uint32T, lw bool, lr bool, hw bool, hr bool) EspErrT {
	return 0
}

/**
 * @brief Sets permissions for specified block in IRAM region
 *
 * Sets Read, Write and Execute permission for specified unified-management block (0-3) in given memory region.
 * Applicable only to IRAM memory types
 *
 * @param mem_type Memory protection area type (MEMPROT_IRAM0_SRAM)
 * @param block Memory block identifier (0-3)
 * @param write_perm Write permission flag
 * @param read_perm Read permission flag
 * @param exec_perm Execute permission flag
 *
 * @return ESP_OK on success
 *         ESP_ERR_NOT_SUPPORTED on invalid mem_type
 *         ESP_ERR_INVALID_ARG on incorrect block number
 */
// llgo:link MemTypeProtT.EspMemprotSetUniBlockPermIram C.esp_memprot_set_uni_block_perm_iram
func (recv_ MemTypeProtT) EspMemprotSetUniBlockPermIram(block c.Uint32T, write_perm bool, read_perm bool, exec_perm bool) EspErrT {
	return 0
}

/**
 * @brief Sets permissions for high and low memory segment in IRAM region
 *
 * Sets Read, Write and Execute permission for both low and high memory segments given by splitting address.
 * The splitting address must be equal to or higher then beginning of block 5
 * Applicable only to IRAM memory types
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param split_addr Address to split the memory region to lower and higher segment
 * @param lw Low segment Write permission flag
 * @param lr Low segment Read permission flag
 * @param lx Low segment Execute permission flag
 * @param hw High segment Write permission flag
 * @param hr High segment Read permission flag
 * @param hx High segment Execute permission flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotSetProtIram C.esp_memprot_set_prot_iram
func (recv_ MemTypeProtT) EspMemprotSetProtIram(split_addr *c.Uint32T, lw bool, lr bool, lx bool, hw bool, hr bool, hx bool) EspErrT {
	return 0
}

/**
 * @brief Activates memory protection for all supported memory region types
 *
 * @note The feature is disabled when JTAG interface is connected
 *
 * @param invoke_panic_handler map mem.prot interrupt to ETS_MEMACCESS_ERR_INUM and thus invokes panic handler when fired ('true' not suitable for testing)
 * @param lock_feature sets LOCK bit, see esp_memprot_set_lock() ('true' not suitable for testing)
 * @param mem_type_mask holds a set of required memory protection types (bitmask built of mem_type_prot_t). NULL means default (MEMPROT_ALL in this version)
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
//go:linkname EspMemprotSetProt C.esp_memprot_set_prot
func EspMemprotSetProt(invoke_panic_handler bool, lock_feature bool, mem_type_mask *c.Uint32T) EspErrT

/**
 * @brief Get permission settings bits for IRAM0 split mgmt. Only IRAM0 memory types allowed
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param lw Low segment Write permission flag
 * @param lr Low segment Read permission flag
 * @param lx Low segment Execute permission flag
 * @param hw High segment Write permission flag
 * @param hr High segment Read permission flag
 * @param hx High segment Execute permission flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetPermSplitBitsIram C.esp_memprot_get_perm_split_bits_iram
func (recv_ MemTypeProtT) EspMemprotGetPermSplitBitsIram(lw *bool, lr *bool, lx *bool, hw *bool, hr *bool, hx *bool) EspErrT {
	return 0
}

/**
 * @brief Get permission settings bits for DRAM0 split mgmt. Only DRAM0 memory types allowed
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param lw Low segment Write permission flag
 * @param lr Low segment Read permission flag
 * @param hw High segment Write permission flag
 * @param hr High segment Read permission flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetPermSplitBitsDram C.esp_memprot_get_perm_split_bits_dram
func (recv_ MemTypeProtT) EspMemprotGetPermSplitBitsDram(lw *bool, lr *bool, hw *bool, hr *bool) EspErrT {
	return 0
}

/**
 * @brief Sets permissions for high and low memory segment in PERIBUS1 region
 *
 * Sets Read and Write permission for both low and high memory segments given by splitting address.
 * Applicable only to PERIBUS1 memory types
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param split_addr Address to split the memory region to lower and higher segment
 * @param lw Low segment Write permission flag
 * @param lr Low segment Read permission flag
 * @param hw High segment Write permission flag
 * @param hr High segment Read permission flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotSetProtPeri1 C.esp_memprot_set_prot_peri1
func (recv_ MemTypeProtT) EspMemprotSetProtPeri1(split_addr *c.Uint32T, lw bool, lr bool, hw bool, hr bool) EspErrT {
	return 0
}

/**
 * @brief Get permission settings bits for PERIBUS1 split mgmt. Only PERIBUS1 memory types allowed
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param lw Low segment Write permission flag
 * @param lr Low segment Read permission flag
 * @param hw High segment Write permission flag
 * @param hr High segment Read permission flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetPermSplitBitsPeri1 C.esp_memprot_get_perm_split_bits_peri1
func (recv_ MemTypeProtT) EspMemprotGetPermSplitBitsPeri1(lw *bool, lr *bool, hw *bool, hr *bool) EspErrT {
	return 0
}

/**
 * @brief Get permission settings bits for PERIBUS2 split mgmt. Only PERIBUS2 memory types allowed
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param lw Low segment Write permission flag
 * @param lr Low segment Read permission flag
 * @param lx Low segment Execute permission flag
 * @param hw High segment Write permission flag
 * @param hr High segment Read permission flag
 * @param hx High segment Execute permission flag
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_ARG on failure
 */
// llgo:link MemTypeProtT.EspMemprotGetPermSplitBitsPeri2 C.esp_memprot_get_perm_split_bits_peri2
func (recv_ MemTypeProtT) EspMemprotGetPermSplitBitsPeri2(lw *bool, lr *bool, lx *bool, hw *bool, hr *bool, hx *bool) EspErrT {
	return 0
}

/**
 * @brief Configures the memory protection for high and low segment in PERIBUS2 region
 *
 * Sets Read Write permission for both low and high memory segments given by splitting address.
 * Applicable only to PERIBUS2 memory types
 *
 * @param mem_type Memory protection area type (MEMPROT_PERI2_RTCSLOW_0, MEMPROT_PERI2_RTCSLOW_1)
 * @param split_addr Address to split the memory region to lower and higher segment (32bit aligned)
 * @param lw Low segment Write permission flag
 * @param lr Low segment Read permission flag
 * @param lx Low segment Execute permission flag
 * @param hw High segment Write permission flag
 * @param hr High segment Read permission flag
 * @param hx High segment Execute permission flag
 *
 * @return ESP_OK on success
 *         ESP_ERR_NOT_SUPPORTED on invalid mem_type
 *         ESP_ERR_INVALID_STATE on splitting address out of PERIBUS2 range
 *         ESP_ERR_INVALID_SIZE on splitting address not 32-bit aligned
 */
// llgo:link MemTypeProtT.EspMemprotSetProtPeri2 C.esp_memprot_set_prot_peri2
func (recv_ MemTypeProtT) EspMemprotSetProtPeri2(split_addr *c.Uint32T, lw bool, lr bool, lx bool, hw bool, hr bool, hx bool) EspErrT {
	return 0
}

/**
 * @brief Get permissions for specified memory type. Irrelevant bits are ignored
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param lw Low segment Write permission flag
 * @param lr Low segment Read permission flag
 * @param lx Low segment Execute permission flag
 * @param hw High segment Write permission flag
 * @param hr High segment Read permission flag
 * @param hx High segment Execute permission flag
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on NULL lw/lr/lx/hw/hr/hx args
 *         ESP_ERR_NOT_SUPPORTED on invalid mem_type
 */
// llgo:link MemTypeProtT.EspMemprotGetPermissions C.esp_memprot_get_permissions
func (recv_ MemTypeProtT) EspMemprotGetPermissions(lw *bool, lr *bool, lx *bool, hw *bool, hr *bool, hx *bool) EspErrT {
	return 0
}

/**
 * @brief Get Read permission settings for low and high regions of given memory type
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param lr Low segment Read permission flag
 * @param hr High segment Read permission flag
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on NULL lr/hr args
 *         ESP_ERR_NOT_SUPPORTED on invalid mem_type
 */
// llgo:link MemTypeProtT.EspMemprotGetPermRead C.esp_memprot_get_perm_read
func (recv_ MemTypeProtT) EspMemprotGetPermRead(lr *bool, hr *bool) EspErrT {
	return 0
}

/**
 * @brief Get Write permission settings for low and high regions of given memory type
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param lr Low segment Write permission flag
 * @param hr High segment Write permission flag
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on NULL lw/hw args
 *         ESP_ERR_NOT_SUPPORTED on invalid mem_type
 */
// llgo:link MemTypeProtT.EspMemprotGetPermWrite C.esp_memprot_get_perm_write
func (recv_ MemTypeProtT) EspMemprotGetPermWrite(lw *bool, hw *bool) EspErrT {
	return 0
}

/**
 * @brief Get Execute permission settings for low and high regions of given memory type
 * Applicable only to IBUS-compatible memory types
 *
 * @param mem_type Memory protection area type (MEMPROT_IRAM0_SRAM, MEMPROT_IRAM0_RTCFAST, MEMPROT_PERI2_RTCSLOW_0, MEMPROT_PERI2_RTCSLOW_1)
 * @param lx Low segment Exec permission flag
 * @param hx High segment Exec permission flag
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on NULL lx/hx args
 *         ESP_ERR_NOT_SUPPORTED on invalid mem_type
 */
// llgo:link MemTypeProtT.EspMemprotGetPermExec C.esp_memprot_get_perm_exec
func (recv_ MemTypeProtT) EspMemprotGetPermExec(lx *bool, hx *bool) EspErrT {
	return 0
}

/**
 * @brief Returns the lowest address in required memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 *
 * @return Required address or MEMPROT_INVALID_ADDRESS for invalid mem_type
 */
// llgo:link MemTypeProtT.EspMemprotGetLowLimit C.esp_memprot_get_low_limit
func (recv_ MemTypeProtT) EspMemprotGetLowLimit() c.Uint32T {
	return 0
}

/**
 * @brief Returns the highest address in required memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 *
 * @return Required address or MEMPROT_INVALID_ADDRESS for invalid mem_type
 */
// llgo:link MemTypeProtT.EspMemprotGetHighLimit C.esp_memprot_get_high_limit
func (recv_ MemTypeProtT) EspMemprotGetHighLimit() c.Uint32T {
	return 0
}

/**
 * @brief Sets READ permission bit for required memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param lr Low segment Read permission flag
 * @param hr High segment Read permission flag
 *
 * @return ESP_OK on success
 *         ESP_ERR_NOT_SUPPORTED on invalid mem_type
 */
// llgo:link MemTypeProtT.EspMemprotSetReadPerm C.esp_memprot_set_read_perm
func (recv_ MemTypeProtT) EspMemprotSetReadPerm(lr bool, hr bool) EspErrT {
	return 0
}

/**
 * @brief Sets WRITE permission bit for required memory region
 *
 * @param mem_type Memory protection area type (see mem_type_prot_t enum)
 * @param lr Low segment Write permission flag
 * @param hr High segment Write permission flag
 *
 * @return ESP_OK on success
 *         ESP_ERR_NOT_SUPPORTED on invalid mem_type
 */
// llgo:link MemTypeProtT.EspMemprotSetWritePerm C.esp_memprot_set_write_perm
func (recv_ MemTypeProtT) EspMemprotSetWritePerm(lw bool, hw bool) EspErrT {
	return 0
}

/**
 * @brief Sets EXECUTE permission bit for required memory region
 *
 * @param mem_type Memory protection area type (MEMPROT_IRAM0_SRAM, MEMPROT_IRAM0_RTCFAST, MEMPROT_PERI2_RTCSLOW_0, MEMPROT_PERI2_RTCSLOW_1)
 * @param lr Low segment Exec permission flag
 * @param hr High segment Exec permission flag
 *
 * @return ESP_OK on success
 *         ESP_ERR_NOT_SUPPORTED on invalid mem_type
 */
// llgo:link MemTypeProtT.EspMemprotSetExecPerm C.esp_memprot_set_exec_perm
func (recv_ MemTypeProtT) EspMemprotSetExecPerm(lx bool, hx bool) EspErrT {
	return 0
}
