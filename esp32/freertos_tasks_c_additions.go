package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get reentrancy structure of the current task
 *
 * - This function is required by newlib (when __DYNAMIC_REENT__ is enabled)
 * - It will return a pointer to the current task's reent struct
 * - If FreeRTOS is not running, it will return the global reent struct
 *
 * @return Pointer to a the (current taks's)/(global) reent struct
 */
//go:linkname X__getreent C.__getreent
func X__getreent() *X_reent

const (
	ESP_FREERTOS_DEBUG_TABLE_SIZE          c.Int = 0
	ESP_FREERTOS_DEBUG_TABLE_VERSION       c.Int = 1
	ESP_FREERTOS_DEBUG_KERNEL_VER_MAJOR    c.Int = 2
	ESP_FREERTOS_DEBUG_KERNEL_VER_MINOR    c.Int = 3
	ESP_FREERTOS_DEBUG_KERNEL_VER_BUILD    c.Int = 4
	ESP_FREERTOS_DEBUG_UX_TOP_USED_PIORITY c.Int = 5
	ESP_FREERTOS_DEBUG_PX_TOP_OF_STACK     c.Int = 6
	ESP_FREERTOS_DEBUG_PC_TASK_NAME        c.Int = 7
	ESP_FREERTOS_DEBUG_TABLE_END           c.Int = 8
)
