package freertos

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

const PortCRITICAL_NESTING_IN_TCB = 0
const PortBYTE_ALIGNMENT = 16
const PortTICK_TYPE_IS_ATOMIC = 1

type StackTypeT c.Uint8T
type BaseTypeT c.Int
type UBaseTypeT c.Uint
type TickTypeT c.Uint32T

/**
 * @brief Checks if the current core is in an ISR context
 *
 * - ISR context consist of Low/Mid priority ISR, or time tick ISR
 * - High priority ISRs aren't detected here, but they normally cannot call C code, so that should not be an issue anyway.
 *
 * @note [refactor-todo] Check if this should be inlined
 * @return
 *  - pdTRUE if in ISR
 *  - pdFALSE otherwise
 */
//go:linkname XPortInIsrContext C.xPortInIsrContext
func XPortInIsrContext() BaseTypeT

/**
 * @brief Assert if in ISR context
 *
 * - Asserts on xPortInIsrContext() internally
 */
//go:linkname VPortAssertIfInISR C.vPortAssertIfInISR
func VPortAssertIfInISR()

/**
 * @brief Check if in ISR context from High priority ISRs
 *
 * - Called from High priority ISR
 * - Checks if the previous context (before high priority interrupt) was in ISR context (meaning low/med priority)
 *
 * @note [refactor-todo] Check if this should be inlined
 * @return
 *  - pdTRUE if in previous in ISR context
 *  - pdFALSE otherwise
 */
//go:linkname XPortInterruptedFromISRContext C.xPortInterruptedFromISRContext
func XPortInterruptedFromISRContext() BaseTypeT

type PortMUXTYPE SpinlockT

/**
 * @brief Enter a SMP critical section with a timeout
 *
 * This function enters an SMP critical section by disabling interrupts then
 * taking a spinlock with a specified timeout.
 *
 * This function can be called in a nested manner.
 *
 * @note This function is made non-inline on purpose to reduce code size
 * @param mux Spinlock
 * @param timeout Timeout to wait for spinlock in number of CPU cycles.
 *                Use portMUX_NO_TIMEOUT to wait indefinitely
 *                Use portMUX_TRY_LOCK to only getting the spinlock a single time
 * @retval pdPASS Critical section entered (spinlock taken)
 * @retval pdFAIL If timed out waiting for spinlock (will not occur if using portMUX_NO_TIMEOUT)
 */
// llgo:link (*PortMUXTYPE).XPortEnterCriticalTimeout C.xPortEnterCriticalTimeout
func (recv_ *PortMUXTYPE) XPortEnterCriticalTimeout(timeout BaseTypeT) BaseTypeT {
	return 0
}

/**
 * @brief Exit a SMP critical section
 *
 * This function can be called in a nested manner. On the outer most level of nesting, this function will:
 *
 * - Release the spinlock
 * - Restore the previous interrupt level before the critical section was entered
 *
 * If still nesting, this function simply decrements a critical nesting count
 *
 * @note This function is made non-inline on purpose to reduce code size
 * @param[in] mux Spinlock
 */
// llgo:link (*PortMUXTYPE).VPortExitCritical C.vPortExitCritical
func (recv_ *PortMUXTYPE) VPortExitCritical() {
}

/**
 * @brief FreeRTOS Compliant version of xPortEnterCriticalTimeout()
 *
 * Compliant version of xPortEnterCriticalTimeout() will ensure that this is
 * called from a task context only. An abort is called otherwise.
 *
 * @note This function is made non-inline on purpose to reduce code size
 *
 * @param mux Spinlock
 * @param timeout Timeout
 * @return BaseType_t
 */
// llgo:link (*PortMUXTYPE).XPortEnterCriticalTimeoutCompliance C.xPortEnterCriticalTimeoutCompliance
func (recv_ *PortMUXTYPE) XPortEnterCriticalTimeoutCompliance(timeout BaseTypeT) BaseTypeT {
	return 0
}

/**
 * @brief FreeRTOS compliant version of vPortExitCritical()
 *
 * Compliant version of vPortExitCritical() will ensure that this is
 * called from a task context only. An abort is called otherwise.
 *
 * @note This function is made non-inline on purpose to reduce code size
 * @param[in] mux Spinlock
 */
// llgo:link (*PortMUXTYPE).VPortExitCriticalCompliance C.vPortExitCriticalCompliance
func (recv_ *PortMUXTYPE) VPortExitCriticalCompliance() {
}

/**
 * @brief Perform a solicited context switch
 *
 * - Defined in portasm.S
 *
 * @note [refactor-todo] The rest of ESP-IDF should call taskYield() instead
 */
//go:linkname VPortYield C.vPortYield
func VPortYield()

/**
 * @brief Yields the other core
 *
 * - Send an interrupt to another core in order to make the task running on it yield for a higher-priority task.
 * - Can be used to yield current core as well
 *
 * @note [refactor-todo] Put this into private macros as its only called from task.c and is not public API
 * @param coreid ID of core to yield
 */
// llgo:link BaseTypeT.VPortYieldOtherCore C.vPortYieldOtherCore
func (recv_ BaseTypeT) VPortYieldOtherCore() {
}

/**
 * @brief Hook function called on entry to tickless idle
 *
 * - Implemented in pm_impl.c
 *
 * @param xExpectedIdleTime Expected idle time
 */
// llgo:link TickTypeT.VApplicationSleep C.vApplicationSleep
func (recv_ TickTypeT) VApplicationSleep() {
}

/**
 * @brief Get the tick rate per second
 *
 * @note [refactor-todo] make this inline
 * @return uint32_t Tick rate in Hz
 */
//go:linkname XPortGetTickRateHz C.xPortGetTickRateHz
func XPortGetTickRateHz() c.Uint32T

/**
 * @brief Set a watchpoint to watch the last 32 bytes of the stack
 *
 * Callback to set a watchpoint on the end of the stack. Called every context switch to change the stack watchpoint
 * around.
 *
 * @param pxStackStart Pointer to the start of the stack
 */
//go:linkname VPortSetStackWatchpoint C.vPortSetStackWatchpoint
func VPortSetStackWatchpoint(pxStackStart c.Pointer)

/**
 * @brief TCB cleanup hook
 *
 * The portCLEAN_UP_TCB() macro is called in prvDeleteTCB() right before a
 * deleted task's memory is freed. We map that macro to this internal function
 * so that IDF FreeRTOS ports can inject some task pre-deletion operations.
 *
 * @note We can't use vPortCleanUpTCB() due to API compatibility issues. See
 * CONFIG_FREERTOS_ENABLE_STATIC_TASK_CLEAN_UP. Todo: IDF-8097
 */
//go:linkname VPortTCBPreDeleteHook C.vPortTCBPreDeleteHook
func VPortTCBPreDeleteHook(pxTCB c.Pointer)

//go:linkname X_frxtSetupSwitch C._frxt_setup_switch
func X_frxtSetupSwitch()

/**
 * @brief Checks if a given piece of memory can be used to store a FreeRTOS list
 *
 * - Defined in heap_idf.c
 *
 * @param ptr Pointer to memory
 * @return true Memory can be used to store a List
 * @return false Otherwise
 */
//go:linkname XPortCheckValidListMem C.xPortCheckValidListMem
func XPortCheckValidListMem(ptr c.Pointer) bool

/**
 * @brief Checks if a given piece of memory can be used to store a task's TCB
 *
 * - Defined in heap_idf.c
 *
 * @param ptr Pointer to memory
 * @return true Memory can be used to store a TCB
 * @return false Otherwise
 */
//go:linkname XPortCheckValidTCBMem C.xPortCheckValidTCBMem
func XPortCheckValidTCBMem(ptr c.Pointer) bool

/**
 * @brief Checks if a given piece of memory can be used to store a task's stack
 *
 * - Defined in heap_idf.c
 *
 * @param ptr Pointer to memory
 * @return true Memory can be used to store a task stack
 * @return false Otherwise
 */
//go:linkname XPortcheckValidStackMem C.xPortcheckValidStackMem
func XPortcheckValidStackMem(ptr c.Pointer) bool
