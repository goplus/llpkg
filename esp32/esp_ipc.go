package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type EspIpcFuncT func(c.Pointer)

/**
 * @brief Execute a callback on a given CPU
 *
 * Execute a given callback on a particular CPU. The callback must be of type "esp_ipc_func_t" and will be invoked in
 * the context of the target CPU's IPC task.
 *
 * - This function will block the target CPU's IPC task has begun execution of the callback
 * - If another IPC call is ongoing, this function will block until the ongoing IPC call completes
 * - The stack size of the IPC task can be configured via the CONFIG_ESP_IPC_TASK_STACK_SIZE option
 *
 * @note In single-core mode, returns ESP_ERR_INVALID_ARG for cpu_id 1.
 *
 * @param[in]   cpu_id  CPU where the given function should be executed (0 or 1)
 * @param[in]   func    Pointer to a function of type void func(void* arg) to be executed
 * @param[in]   arg     Arbitrary argument of type void* to be passed into the function
 *
 * @return
 *      - ESP_ERR_INVALID_ARG if cpu_id is invalid
 *      - ESP_ERR_INVALID_STATE if the FreeRTOS scheduler is not running
 *      - ESP_OK otherwise
 */
//go:linkname EspIpcCall C.esp_ipc_call
func EspIpcCall(cpu_id c.Uint32T, func_ EspIpcFuncT, arg c.Pointer) EspErrT

/**
 * @brief Execute a callback on a given CPU until and block until it completes
 *
 * This function is identical to esp_ipc_call() except that this function will block until the execution of the callback
 * completes.
 *
 * @note    In single-core mode, returns ESP_ERR_INVALID_ARG for cpu_id 1.
 *
 * @param[in]   cpu_id  CPU where the given function should be executed (0 or 1)
 * @param[in]   func    Pointer to a function of type void func(void* arg) to be executed
 * @param[in]   arg     Arbitrary argument of type void* to be passed into the function
 *
 * @return
 *      - ESP_ERR_INVALID_ARG if cpu_id is invalid
 *      - ESP_ERR_INVALID_STATE if the FreeRTOS scheduler is not running
 *      - ESP_OK otherwise
 */
//go:linkname EspIpcCallBlocking C.esp_ipc_call_blocking
func EspIpcCallBlocking(cpu_id c.Uint32T, func_ EspIpcFuncT, arg c.Pointer) EspErrT

/**
 * @brief Execute a callback on a given CPU without any blocking operations for the caller
 *
 * Since it does not have any blocking operations it is suitable to be run from interrupts
 * or even when the Scheduler on the current core is suspended.
 *
 * The function:
 * - does not wait for the callback to begin or complete execution,
 * - does not change the IPC priority.
 * The function returns after sending a notification to the IPC task to execute the callback.
 *
 * @param[in]   cpu_id  CPU where the given function should be executed (0 or 1)
 * @param[in]   func    Pointer to a function of type void func(void* arg) to be executed
 * @param[in]   arg     Arbitrary argument of type void* to be passed into the function
 *
 * @return
 *      - ESP_ERR_INVALID_ARG if cpu_id is invalid
 *      - ESP_ERR_INVALID_STATE 1. IPC tasks have not been initialized yet,
 *                              2. cpu_id requests IPC on the current core, but the FreeRTOS scheduler is not running on it
 *                              (the IPC task cannot be executed).
 *      - ESP_FAIL IPC is busy due to a previous call was not completed.
 *      - ESP_OK otherwise
 */
//go:linkname EspIpcCallNonblocking C.esp_ipc_call_nonblocking
func EspIpcCallNonblocking(cpu_id c.Uint32T, func_ EspIpcFuncT, arg c.Pointer) EspErrT
