package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type EspIpcIsrFuncT func(c.Pointer)

/**
 * @brief Execute an ISR callback on the other CPU
 *
 * Execute a given callback on the other CPU in the context of a High Priority Interrupt.
 *
 * - This function will busy-wait in a critical section until the other CPU has started execution of the callback
 * - The callback must be written:
 *    - in assembly for XTENSA chips (such as ESP32, ESP32S3).
 *      The function is invoked using a CALLX0 instruction and can use only a2, a3, a4 registers.
 *      See :doc:`IPC in Interrupt Context </api-reference/system/ipc>` doc for more details.
 *    - in C or assembly for RISCV chips (such as ESP32P4).
 *
 * @note This function is not available in single-core mode.
 *
 * @param[in]   func    Pointer to a function of type void func(void* arg) to be executed
 * @param[in]   arg     Arbitrary argument of type void* to be passed into the function
 */
//go:linkname EspIpcIsrCall C.esp_ipc_isr_call
func EspIpcIsrCall(func_ EspIpcIsrFuncT, arg c.Pointer)

/**
 * @brief Execute an ISR callback on the other CPU and busy-wait until it completes
 *
 * This function is identical to esp_ipc_isr_call() except that this function will busy-wait until the execution of
 * the callback completes.
 *
 * @note This function is not available in single-core mode.
 *
 * @param[in]   func    Pointer to a function of type void func(void* arg) to be executed
 * @param[in]   arg     Arbitrary argument of type void* to be passed into the function
 */
//go:linkname EspIpcIsrCallBlocking C.esp_ipc_isr_call_blocking
func EspIpcIsrCallBlocking(func_ EspIpcIsrFuncT, arg c.Pointer)

/**
 * @brief Stall the other CPU
 *
 * This function will stall the other CPU. The other CPU is stalled by busy-waiting in the context of a High Priority
 * Interrupt. The other CPU will not be resumed until esp_ipc_isr_release_other_cpu() is called.
 *
 * - This function is internally implemented using IPC ISR
 * - This function is used for DPORT workaround.
 * - If the stall feature is paused using esp_ipc_isr_stall_pause(), this function will have no effect
 *
 * @note This function is not available in single-core mode.
 * @note It is the caller's responsibility to avoid deadlocking on spinlocks
 */
//go:linkname EspIpcIsrStallOtherCpu C.esp_ipc_isr_stall_other_cpu
func EspIpcIsrStallOtherCpu()

/**
 * @brief Release the other CPU
 *
 * This function will release the other CPU that was previously stalled from calling esp_ipc_isr_stall_other_cpu()
 *
 * - This function is used for DPORT workaround.
 * - If the stall feature is paused using esp_ipc_isr_stall_pause(), this function will have no effect
 *
 * @note This function is not available in single-core mode.
 */
//go:linkname EspIpcIsrReleaseOtherCpu C.esp_ipc_isr_release_other_cpu
func EspIpcIsrReleaseOtherCpu()

/**
 * @brief Puase the CPU stall feature
 *
 * This function will pause the CPU stall feature. Once paused, calls to esp_ipc_isr_stall_other_cpu() and
 * esp_ipc_isr_release_other_cpu() will have no effect. If a IPC ISR call is already in progress, this function will
 * busy-wait until the call completes before pausing the CPU stall feature.
 */
//go:linkname EspIpcIsrStallPause C.esp_ipc_isr_stall_pause
func EspIpcIsrStallPause()

/**
 * @brief Abort a CPU stall
 *
 * This function will abort any stalling routine of the other CPU due to a pervious call to
 * esp_ipc_isr_stall_other_cpu(). This function aborts the stall in a non-recoverable manner, thus should only be called
 * in case of a panic().
 *
 * - This function is used in panic handling code
 */
//go:linkname EspIpcIsrStallAbort C.esp_ipc_isr_stall_abort
func EspIpcIsrStallAbort()

/**
 * @brief Resume the CPU stall feature
 *
 * This function will resume the CPU stall feature that was previously paused by calling esp_ipc_isr_stall_pause(). Once
 * resumed, calls to esp_ipc_isr_stall_other_cpu() and esp_ipc_isr_release_other_cpu() will have effect again.
 */
//go:linkname EspIpcIsrStallResume C.esp_ipc_isr_stall_resume
func EspIpcIsrStallResume()

/**
 * @brief Initialize the IPC ISR feature, must be called for each CPU
 *
 * @note This function is called from ipc_task().
 *
 * This function initializes the IPC ISR feature and must be called before any other esp_ipc_isr...() functions.
 * The IPC ISR feature allows for callbacks (written in assembly) to be run on a particular CPU in the context of a
 * High Priority Interrupt.
 *
 * - This function will register a High Priority Interrupt for a CPU where it is called. The priority of the interrupts is dependent on
 *   the CONFIG_ESP_SYSTEM_CHECK_INT_LEVEL option.
 * - Callbacks written in assembly can then run in context of the registered High Priority Interrupts
 * - Callbacks can be executed by calling esp_ipc_isr_call() or esp_ipc_isr_call_blocking()
 */
//go:linkname EspIpcIsrInit C.esp_ipc_isr_init
func EspIpcIsrInit()
