package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_CPU_INTR_DESC_FLAG_SPECIAL = 0x01
const ESP_CPU_INTR_DESC_FLAG_RESVD = 0x02

type EspCpuCycleCountT c.Uint32T
type EspCpuIntrTypeT c.Int

const (
	ESP_CPU_INTR_TYPE_LEVEL EspCpuIntrTypeT = 0
	ESP_CPU_INTR_TYPE_EDGE  EspCpuIntrTypeT = 1
	ESP_CPU_INTR_TYPE_NA    EspCpuIntrTypeT = 2
)

/**
 * @brief CPU interrupt descriptor
 *
 * Each particular CPU interrupt has an associated descriptor describing that
 * particular interrupt's characteristics. Call esp_cpu_intr_get_desc() to get
 * the descriptors of a particular interrupt.
 */

type EspCpuIntrDescT struct {
	Priority c.Int
	Type     EspCpuIntrTypeT
	Flags    c.Uint32T
}

// llgo:type C
type EspCpuIntrHandlerT func(c.Pointer)
type EspCpuWatchpointTriggerT c.Int

const (
	ESP_CPU_WATCHPOINT_LOAD   EspCpuWatchpointTriggerT = 0
	ESP_CPU_WATCHPOINT_STORE  EspCpuWatchpointTriggerT = 1
	ESP_CPU_WATCHPOINT_ACCESS EspCpuWatchpointTriggerT = 2
)

/**
 * @brief Stall a CPU core
 *
 * @param core_id  The core's ID
 */
//go:linkname EspCpuStall C.esp_cpu_stall
func EspCpuStall(core_id c.Int)

/**
 * @brief Resume a previously stalled CPU core
 *
 * @param core_id The core's ID
 */
//go:linkname EspCpuUnstall C.esp_cpu_unstall
func EspCpuUnstall(core_id c.Int)

/**
 * @brief Reset a CPU core
 *
 * @param core_id The core's ID
 */
//go:linkname EspCpuReset C.esp_cpu_reset
func EspCpuReset(core_id c.Int)

/**
 * @brief Wait for Interrupt
 *
 * This function causes the current CPU core to execute its Wait For Interrupt
 * (WFI or equivalent) instruction. After executing this function, the CPU core
 * will stop execution until an interrupt occurs.
 */
//go:linkname EspCpuWaitForIntr C.esp_cpu_wait_for_intr
func EspCpuWaitForIntr()

/**
 * @brief Get a CPU interrupt's descriptor
 *
 * Each CPU interrupt has a descriptor describing the interrupt's capabilities
 * and restrictions. This function gets the descriptor of a particular interrupt
 * on a particular CPU.
 *
 * @param[in] core_id The core's ID
 * @param[in] intr_num Interrupt number
 * @param[out] intr_desc_ret The interrupt's descriptor
 */
//go:linkname EspCpuIntrGetDesc C.esp_cpu_intr_get_desc
func EspCpuIntrGetDesc(core_id c.Int, intr_num c.Int, intr_desc_ret *EspCpuIntrDescT)

/**
 * @brief Configure the CPU to disable access to invalid memory regions
 */
//go:linkname EspCpuConfigureRegionProtection C.esp_cpu_configure_region_protection
func EspCpuConfigureRegionProtection()

/**
 * @brief Set and enable a hardware breakpoint on the current CPU
 *
 * @note This function is meant to be called by the panic handler to set a
 * breakpoint for an attached debugger during a panic.
 * @note Overwrites previously set breakpoint with same breakpoint number.
 * @param bp_num Hardware breakpoint number [0..SOC_CPU_BREAKPOINTS_NUM - 1]
 * @param bp_addr Address to set a breakpoint on
 * @return ESP_OK if breakpoint is set. Failure otherwise
 */
//go:linkname EspCpuSetBreakpoint C.esp_cpu_set_breakpoint
func EspCpuSetBreakpoint(bp_num c.Int, bp_addr c.Pointer) EspErrT

/**
 * @brief Clear a hardware breakpoint on the current CPU
 *
 * @note Clears a breakpoint regardless of whether it was previously set
 * @param bp_num Hardware breakpoint number [0..SOC_CPU_BREAKPOINTS_NUM - 1]
 * @return ESP_OK if breakpoint is cleared. Failure otherwise
 */
//go:linkname EspCpuClearBreakpoint C.esp_cpu_clear_breakpoint
func EspCpuClearBreakpoint(bp_num c.Int) EspErrT

/**
 * @brief Set and enable a hardware watchpoint on the current CPU
 *
 * Set and enable a hardware watchpoint on the current CPU, specifying the
 * memory range and trigger operation. Watchpoints will break/panic the CPU when
 * the CPU accesses (according to the trigger type) on a certain memory range.
 *
 * @note Overwrites previously set watchpoint with same watchpoint number.
 *       On RISC-V chips, this API uses method0(Exact matching) and method1(NAPOT matching) according to the
 *       riscv-debug-spec-0.13 specification for address matching.
 *       If the watch region size is 1byte, it uses exact matching (method 0).
 *       If the watch region size is larger than 1byte, it uses NAPOT matching (method 1). This mode requires
 *       the watching region start address to be aligned to the watching region size.
 *
 * @param wp_num Hardware watchpoint number [0..SOC_CPU_WATCHPOINTS_NUM - 1]
 * @param wp_addr Watchpoint's base address, must be naturally aligned to the size of the region
 * @param size Size of the region to watch. Must be one of 2^n and in the range of [1 ... SOC_CPU_WATCHPOINT_MAX_REGION_SIZE]
 * @param trigger Trigger type
 * @return ESP_ERR_INVALID_ARG on invalid arg, ESP_OK otherwise
 */
//go:linkname EspCpuSetWatchpoint C.esp_cpu_set_watchpoint
func EspCpuSetWatchpoint(wp_num c.Int, wp_addr c.Pointer, size c.SizeT, trigger EspCpuWatchpointTriggerT) EspErrT

/**
 * @brief Clear a hardware watchpoint on the current CPU
 *
 * @note Clears a watchpoint regardless of whether it was previously set
 * @param wp_num Hardware watchpoint number [0..SOC_CPU_WATCHPOINTS_NUM - 1]
 * @return ESP_OK if watchpoint was cleared. Failure otherwise.
 */
//go:linkname EspCpuClearWatchpoint C.esp_cpu_clear_watchpoint
func EspCpuClearWatchpoint(wp_num c.Int) EspErrT

/**
 * @brief Atomic compare-and-set operation
 *
 * @param addr Address of atomic variable
 * @param compare_value Value to compare the atomic variable to
 * @param new_value New value to set the atomic variable to
 * @return Whether the atomic variable was set or not
 */
//go:linkname EspCpuCompareAndSet C.esp_cpu_compare_and_set
func EspCpuCompareAndSet(addr *c.Uint32T, compare_value c.Uint32T, new_value c.Uint32T) bool
