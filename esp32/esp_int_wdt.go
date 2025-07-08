package freertos

import _ "unsafe"

/**
 * @brief Initialize the non-CPU-specific parts of interrupt watchdog.
 *
 * This function is automatically called during application startup if the
 * interrupt watchdog is enabled in menuconfig.
 */
//go:linkname EspIntWdtInit C.esp_int_wdt_init
func EspIntWdtInit()

/**
 * @brief Enable the interrupt watchdog on the current CPU.
 *
 * This function is automatically called during application startup for each CPU
 * that has enabled the interrupt watchdog in menuconfig.
 *
 * @note esp_int_wdt_init() must be called first before calling this function
 */
//go:linkname EspIntWdtCpuInit C.esp_int_wdt_cpu_init
func EspIntWdtCpuInit()
