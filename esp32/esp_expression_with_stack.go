package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type SharedStackFunction func()

/**
 * @brief Calls function on user defined shared stack space
 *
 * After returning, the original stack is used again.
 *
 * @warning This function does minimal preparation of the provided piece of memory (\c stack).
 *          DO NOT do any of the following in \c function or any of its callees:
 *          * Use Thread-local storage
 *          * Use the Floating-point unit on ESP32-P4
 *          * Use the AI co-processor on ESP32-P4
 *          * Call vTaskDelete(NULL) (deleting the currently running task)
 *          Furthermore, backtraces will be wrong when called from \c function or any of its callees.
 *          The limitations are quite sever, so that we might deprecate this function in the future.
 *          If you have any use case which can only be implemented using this function, please open
 *          an issue on github.
 *
 * @param lock Mutex object to protect in case of shared stack
 * @param stack Pointer to user allocated stack
 * @param stack_size Size of current stack in bytes
 * @param function pointer to the shared stack function to be executed
 * @note  if either lock, stack or stack size is invalid, the expression will
 *          be called using the current stack.
 */
//go:linkname EspExecuteSharedStackFunction C.esp_execute_shared_stack_function
func EspExecuteSharedStackFunction(lock SemaphoreHandleT, stack c.Pointer, stack_size c.SizeT, function SharedStackFunction)
