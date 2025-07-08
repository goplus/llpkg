package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
* @brief Convert Memprot low level errors to esp_err_t
 */
// llgo:link MemprotHalErrT.EspMprotLlErrToEspErr C.esp_mprot_ll_err_to_esp_err
func (recv_ MemprotHalErrT) EspMprotLlErrToEspErr() EspErrT {
	return 0
}

/**
 * @brief Convert Memprot low level PMS World IDs to esp_mprot_pms_world_t
 */
// llgo:link MemprotHalWorldT.EspMprotLlWorldToHlWorld C.esp_mprot_ll_world_to_hl_world
func (recv_ MemprotHalWorldT) EspMprotLlWorldToHlWorld() EspMprotPmsWorldT {
	return 0
}

/**
 * @brief Converts operation type to string, no combination of operations allowed
 *
 * @param oper_type PMS operation type
 */
//go:linkname EspMprotOperTypeToStr C.esp_mprot_oper_type_to_str
func EspMprotOperTypeToStr(oper_type c.Uint32T) *c.Char

/**
 * @brief Converts PMS World type to string
 *
 * @param area_type PMS World type
 */
// llgo:link EspMprotPmsWorldT.EspMprotPmsWorldToStr C.esp_mprot_pms_world_to_str
func (recv_ EspMprotPmsWorldT) EspMprotPmsWorldToStr() *c.Char {
	return nil
}

/**
 * @brief Sets splitting address for given line type in the target Memory type
 *
 * @param mem_type memory type
 * @param line_type split address type
 * @param line_addr target address from a memory range relevant to given line_addr
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 *         ESP_ERR_MEMPROT_SPLIT_ADDR_INVALID on invalid line_type
 *         ESP_ERR_MEMPROT_SPLIT_ADDR_OUT_OF_RANGE on splitting line out of given memory-type range
 *         ESP_ERR_MEMPROT_SPLIT_ADDR_UNALIGNED on splitting line not aligned to PMS-required boundaries
 */
// llgo:link EspMprotMemT.EspMprotSetSplitAddr C.esp_mprot_set_split_addr
func (recv_ EspMprotMemT) EspMprotSetSplitAddr(line_type EspMprotSplitAddrT, line_addr c.Pointer, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Gets PMS splitting address for given split_addr type
 *
 * The value is read from the PMS configuration registers
 *
 * @param mem_type memory type
 * @param line_type Split line type (see esp_mprot_split_addr_t enum)
 * @param[out] line_addr Split line address from the configuration register
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on line_addr is pointer
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 *         ESP_ERR_MEMPROT_SPLIT_ADDR_INVALID on invalid line_type
 */
// llgo:link EspMprotMemT.EspMprotGetSplitAddr C.esp_mprot_get_split_addr
func (recv_ EspMprotMemT) EspMprotGetSplitAddr(line_type EspMprotSplitAddrT, line_addr *c.Pointer, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Returns default main I/D splitting address for given Memory type
 *
 * @param mem_type memory type
 * @param[out] def_split_addr Main I/D splitting address of required mem_type
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on invalid def_split_addr pointer
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 */
// llgo:link EspMprotMemT.EspMprotGetDefaultMainSplitAddr C.esp_mprot_get_default_main_split_addr
func (recv_ EspMprotMemT) EspMprotGetDefaultMainSplitAddr(def_split_addr *c.Pointer) EspErrT {
	return 0
}

/**
 * @brief Sets a lock for the main IRAM/DRAM splitting addresses
 * Locks can be unlocked only by digital system reset
 *
 * @param mem_type memory type
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 */
// llgo:link EspMprotMemT.EspMprotSetSplitAddrLock C.esp_mprot_set_split_addr_lock
func (recv_ EspMprotMemT) EspMprotSetSplitAddrLock(core c.Int) EspErrT {
	return 0
}

/**
 * @brief Gets a lock status for the splitting address configuration of given Memory type
 *
 * @param mem_type memory type
 * @param[out] locked mem_type related lock status
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARGUMENT on invalid locked pointer
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 */
// llgo:link EspMprotMemT.EspMprotGetSplitAddrLock C.esp_mprot_get_split_addr_lock
func (recv_ EspMprotMemT) EspMprotGetSplitAddrLock(locked *bool, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Sets a lock for PMS Area settings of required Memory type
 * Locks can be unlocked only by digital system reset
 *
 * @param mem_type memory type
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 */
// llgo:link EspMprotMemT.EspMprotSetPmsLock C.esp_mprot_set_pms_lock
func (recv_ EspMprotMemT) EspMprotSetPmsLock(core c.Int) EspErrT {
	return 0
}

/**
 * @brief Gets a lock status for PMS Area settings of required Memory type
 *
 * @param mem_type memory type
 * @param[out] locked mem_type related lock status
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 *         ESP_ERR_INVALID_ARGUMENT on invalid locked pointer
 */
// llgo:link EspMprotMemT.EspMprotGetPmsLock C.esp_mprot_get_pms_lock
func (recv_ EspMprotMemT) EspMprotGetPmsLock(locked *bool, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Sets permissions for given PMS Area
 *
 * @param area_type PMS area type
 * @param flags combination of MEMPROT_OP_* defines
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 */
// llgo:link EspMprotPmsAreaT.EspMprotSetPmsArea C.esp_mprot_set_pms_area
func (recv_ EspMprotPmsAreaT) EspMprotSetPmsArea(flags c.Uint32T, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Gets current permissions for given PMS Area
 *
 * @param area_type PMS area type
 * @param[out] flags combination of MEMPROT_OP_* defines
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 *         ESP_ERR_INVALID_ARG on invalid flags pointer
 */
// llgo:link EspMprotPmsAreaT.EspMprotGetPmsArea C.esp_mprot_get_pms_area
func (recv_ EspMprotPmsAreaT) EspMprotGetPmsArea(flags *c.Uint32T, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Sets a lock for PMS interrupt monitor settings of required Memory type
 *
 * Locks can be unlocked only by digital system reset
 *
 * @param mem_type memory type (see esp_mprot_mem_t enum)
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 */
// llgo:link EspMprotMemT.EspMprotSetMonitorLock C.esp_mprot_set_monitor_lock
func (recv_ EspMprotMemT) EspMprotSetMonitorLock(core c.Int) EspErrT {
	return 0
}

/**
 * @brief Gets a lock status for PMS interrupt monitor settings of required Memory type
 *
 * @param mem_type memory type
 * @param[out] locked mem_type related lock status
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 *         ESP_ERR_INVALID_ARG on invalid locked pointer
 */
// llgo:link EspMprotMemT.EspMprotGetMonitorLock C.esp_mprot_get_monitor_lock
func (recv_ EspMprotMemT) EspMprotGetMonitorLock(locked *bool, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Enable PMS violation interrupt monitoring of required Memory type
 *
 * @param mem_type memory type
 * @param enable enable/disable violation interrupt monitoring
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 */
// llgo:link EspMprotMemT.EspMprotSetMonitorEn C.esp_mprot_set_monitor_en
func (recv_ EspMprotMemT) EspMprotSetMonitorEn(enable bool, core c.Int) EspErrT {
	return 0
}

/**
 * @brief Gets PMS violation-monitoring-enabled flag for required Memory type
 *
 * @param mem_type memory type
 * @param[out] enabled violation interrupt monitoring enable flag
 * @param core Target CPU/Core ID (see *_CPU_NUM defs in soc.h). Can be NULL on 1-CPU systems
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG on invalid enabled pointer
 *         ESP_ERR_MEMPROT_MEMORY_TYPE_INVALID on invalid mem_type
 */
// llgo:link EspMprotMemT.EspMprotGetMonitorEn C.esp_mprot_get_monitor_en
func (recv_ EspMprotMemT) EspMprotGetMonitorEn(enabled *bool, core c.Int) EspErrT {
	return 0
}
