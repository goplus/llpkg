package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
* @brief Basic PMS interrupt source info
 */
type EspMempIntrSourceT struct {
	MemType EspMprotMemT
	Core    c.Int
}

/**
 * @brief Clears current interrupt ON flag for given Memory type and CPU/Core ID
 *
 * This operation is non-atomic for some chips by PMS module design
 * In such a case the interrupt clearing happens in two steps:
 *      1. Interrupt CLR flag is set (clears interrupt-ON status and inhibits linked interrupt processing)
 *      2. Interrupt CLR flag is reset (resumes the interrupt monitoring)
 *
 * @param mem_type Memory type (see esp_mprot_mem_t enum)
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on passing invalid pointer
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 */
// llgo:link EspMprotMemT.EspMprotMonitorClearIntr C.esp_mprot_monitor_clear_intr
func (recv_ EspMprotMemT) EspMprotMonitorClearIntr(core c.Int) EspErrT {
	return 0
}

/**
 * @brief Checks whether any of the PMS settings is locked
 *
 * @param[out] locked Any lock on? (true/false)
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on invalid locked ptr
 *         Other failures: error code of any failing esp_mprot_get_*_lock() routine (called internally)
 */
//go:linkname EspMprotIsConfLockedAny C.esp_mprot_is_conf_locked_any
func EspMprotIsConfLockedAny(locked *bool) EspErrT

/**
 * @brief Checks whether any PMS violation-interrupt monitoring is enabled
 *
 * @param[out] locked Any PMS violation interrupt monitor is enabled (true/false)
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on invalid enabled ptr
 *         Other failures: error code of esp_mprot_get_monitor_en() routine (called internally for all Memory types)
 */
//go:linkname EspMprotIsIntrEnaAny C.esp_mprot_is_intr_ena_any
func EspMprotIsIntrEnaAny(enabled *bool) EspErrT

/**
 * @brief Returns active PMS violation-interrupt Memory type if any (MEMPROT_TYPE_NONE when none detected)
 * and the CPU/CoreID which was running the faulty code (-1 when no interrupt available)
 *
 * If there are more interrupts indicated on (shouldn't happen), the order of precedence is given by 'esp_mprot_mem_t' enum definition (low->high)
 *
 * @param[out] mem_type Out-pointer for Memory type given by the faulting address (see esp_mprot_mem_t enum)
 * @param[out] core Out-pointer for CPU/Core ID (see *_CPU_NUM defs in soc.h)
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on passing invalid pointer(s)
 */
// llgo:link (*EspMempIntrSourceT).EspMprotGetActiveIntr C.esp_mprot_get_active_intr
func (recv_ *EspMempIntrSourceT) EspMprotGetActiveIntr() EspErrT {
	return 0
}

/**
 * @brief Returns the address which caused the violation interrupt for given Memory type and CPU/Core ID.
 * This function is to be called after a basic resolving of (current) interrupt's parameters (ie corresponding
 * Memory type and CPU ID see esp_mprot_get_active_intr()). This is to minimize processing time of actual exception
 * as this API is typicaly used in a panic-handling code.
 * If there is no active interrupt available for the Memory type/CPU ID required, fault_addr is set to NULL.
 *
 * @param mem_type memory type
 * @param[out] fault_addr Address of the operation which caused the PMS violation interrupt
 * @param core Faulting instruction CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 *         ESP_ERR_INVALID_ARG on invalid fault_addr pointer
 */
// llgo:link EspMprotMemT.EspMprotGetViolateAddr C.esp_mprot_get_violate_addr
func (recv_ EspMprotMemT) EspMprotGetViolateAddr(fault_addr *c.Pointer, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Returns PMS World identifier of the code causing the violation interrupt
 *
 * The value is read from appropriate PMS violation status register and thus might be 0 if the interrupt is not currently active.
 *
 * @param mem_type Memory type
 * @param[out] world PMS World type (see esp_mprot_pms_world_t)
 * @param core Faulting instruction CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 *         ESP_ERR_INVALID_ARG on passing invalid pointer(s)
 *         ESP_ERR_MEMPROT_WORLD_INVALID on invalid World identifier fetched from the register
 */
// llgo:link EspMprotMemT.EspMprotGetViolateWorld C.esp_mprot_get_violate_world
func (recv_ EspMprotMemT) EspMprotGetViolateWorld(world *EspMprotPmsWorldT, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Returns an operation type which caused the violation interrupt
 *
 * The operation resolving is processed over various PMS status register flags, according to given Memory type argument.
 * If the interrupt is not active the result returned is irrelevant (likely evaluated to MEMPROT_OP_READ).
 *
 * @param mem_type Memory type
 * @param[out] oper Operation type (see MEMPROT_OP_* defines)
 * @param core Faulting instruction CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 *         ESP_ERR_INVALID_ARG on invalid oper pointer
 */
// llgo:link EspMprotMemT.EspMprotGetViolateOperation C.esp_mprot_get_violate_operation
func (recv_ EspMprotMemT) EspMprotGetViolateOperation(oper *c.Uint32T, core c.Int) EspErrT {
	return 0
}

/**
* @brief Checks whether given memory type supports byte-enables info
*
* Byte-enables status is available only for DMA/DRAM operations
 *
* @param mem_type memory type
*
* @return byte-enables info available true/false
*/
// llgo:link EspMprotMemT.EspMprotHasByteEnables C.esp_mprot_has_byte_enables
func (recv_ EspMprotMemT) EspMprotHasByteEnables() bool {
	return false
}

/**
 * @brief Returns byte-enables for the address which caused the violation interrupt
 *
 * The value is taken from appropriate PMS violation status register, based on given Memory type
 *
 * @param mem_type Memory type (MEMPROT_TYPE_DRAM0_SRAM)
 * @param[out] byte_en Byte-enables bits
 * @param core Faulting instruction CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 *         ESP_ERR_INVALID_ARGUMENT on invalid byte_en pointer
 */
// llgo:link EspMprotMemT.EspMprotGetViolateByteEnables C.esp_mprot_get_violate_byte_enables
func (recv_ EspMprotMemT) EspMprotGetViolateByteEnables(byte_en *c.Uint32T, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Convenient routine for setting the PMS defaults
 *
 * Called on system startup, depending on ESP_SYSTEM_MEMPROT_FEATURE Kconfig value
 *
 * @param memp_config pointer to Memprot configuration structure (esp_memp_config_t). The structure si chip-specific,
 * for details and defaults see appropriate [target-chip]/soc_memprot_types.h
 *
 * @return ESP_OK on success
 *         Other failures: error code of the failing routine called internally. No specific error processing provided in such a case
 *         due to large number of embedded calls (ie no global unique error table is provided and thus one error code can have different meanings,
 *         depending on the routine issuing the error)
 */
// llgo:link (*EspMempConfigT).EspMprotSetProt C.esp_mprot_set_prot
func (recv_ *EspMempConfigT) EspMprotSetProt() EspErrT {
	return 0
}

/**
 * @brief Generates PMS configuration string of actual device (diagnostics)
 *
 * The functions generates a string from current configuration, control and status registers of the PMS (or similar) module of actual device.
 * The values are fetched using HAL LL calls to help finding possible errors in the Memprot API implementation
 *
 * @param[out] dump_info_string configuration string buffer pointer. The string is allocated by the callee and must be freed by the caller.
 *
 * @return ESP_OK on success
 *         ESP_ERR_NO_MEM on buffer allocation failure
 *         ESP_ERR_INVALID_ARGUMENT on invalid dump_info_string pointer
 */
//go:linkname EspMprotDumpConfiguration C.esp_mprot_dump_configuration
func EspMprotDumpConfiguration(dump_info_string **c.Char) EspErrT
