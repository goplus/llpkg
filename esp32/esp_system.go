package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspResetReasonT c.Int

const (
	ESP_RST_UNKNOWN    EspResetReasonT = 0
	ESP_RST_POWERON    EspResetReasonT = 1
	ESP_RST_EXT        EspResetReasonT = 2
	ESP_RST_SW         EspResetReasonT = 3
	ESP_RST_PANIC      EspResetReasonT = 4
	ESP_RST_INT_WDT    EspResetReasonT = 5
	ESP_RST_TASK_WDT   EspResetReasonT = 6
	ESP_RST_WDT        EspResetReasonT = 7
	ESP_RST_DEEPSLEEP  EspResetReasonT = 8
	ESP_RST_BROWNOUT   EspResetReasonT = 9
	ESP_RST_SDIO       EspResetReasonT = 10
	ESP_RST_USB        EspResetReasonT = 11
	ESP_RST_JTAG       EspResetReasonT = 12
	ESP_RST_EFUSE      EspResetReasonT = 13
	ESP_RST_PWR_GLITCH EspResetReasonT = 14
	ESP_RST_CPU_LOCKUP EspResetReasonT = 15
)

// llgo:type C
type ShutdownHandlerT func()

/**
 * @brief  Register shutdown handler
 *
 * This function allows you to register a handler that gets invoked before
 * the application is restarted using esp_restart function.
 * @param handle function to execute on restart
 * @return
 *   - ESP_OK on success
 *   - ESP_ERR_INVALID_STATE if the handler has already been registered
 *   - ESP_ERR_NO_MEM if no more shutdown handler slots are available
 */
//go:linkname EspRegisterShutdownHandler C.esp_register_shutdown_handler
func EspRegisterShutdownHandler(handle ShutdownHandlerT) EspErrT

/**
 * @brief  Unregister shutdown handler
 *
 * This function allows you to unregister a handler which was previously
 * registered using esp_register_shutdown_handler function.
 *   - ESP_OK on success
 *   - ESP_ERR_INVALID_STATE if the given handler hasn't been registered before
 */
//go:linkname EspUnregisterShutdownHandler C.esp_unregister_shutdown_handler
func EspUnregisterShutdownHandler(handle ShutdownHandlerT) EspErrT

/**
 * @brief  Restart PRO and APP CPUs.
 *
 * This function can be called both from PRO and APP CPUs.
 * After successful restart, CPU reset reason will be SW_CPU_RESET.
 * Peripherals (except for Wi-Fi, BT, UART0, SPI1, and legacy timers) are not reset.
 * This function does not return.
 */
//go:linkname EspRestart C.esp_restart
func EspRestart()

/**
 * @brief  Get reason of last reset
 * @return See description of esp_reset_reason_t for explanation of each value.
 */
//go:linkname EspResetReason C.esp_reset_reason
func EspResetReason() EspResetReasonT

/**
 * @brief  Get the size of available heap.
 *
 * @note Note that the returned value may be larger than the maximum contiguous block
 * which can be allocated.
 *
 * @return Available heap size, in bytes.
 */
//go:linkname EspGetFreeHeapSize C.esp_get_free_heap_size
func EspGetFreeHeapSize() c.Uint32T

/**
 * @brief  Get the size of available internal heap.
 *
 * @note Note that the returned value may be larger than the maximum contiguous block
 * which can be allocated.
 *
 * @return Available internal heap size, in bytes.
 */
//go:linkname EspGetFreeInternalHeapSize C.esp_get_free_internal_heap_size
func EspGetFreeInternalHeapSize() c.Uint32T

/**
 * @brief Get the minimum heap that has ever been available
 *
 * @return Minimum free heap ever available
 */
//go:linkname EspGetMinimumFreeHeapSize C.esp_get_minimum_free_heap_size
func EspGetMinimumFreeHeapSize() c.Uint32T

/**
 * @brief Trigger a software abort
 *
 * @param details Details that will be displayed during panic handling.
 */
//go:linkname EspSystemAbort C.esp_system_abort
func EspSystemAbort(details *c.Char)
