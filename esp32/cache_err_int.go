package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief initialize cache invalid access interrupt
 *
 * This function enables cache invalid access interrupt source and connects it
 * to interrupt input number. It is called from the startup code.
 *
 * On ESP32, the interrupt input number is ETS_MEMACCESS_ERR_INUM. On other targets
 * it is ETS_CACHEERR_INUM. See soc/soc.h for more information.
 */
//go:linkname EspCacheErrIntInit C.esp_cache_err_int_init
func EspCacheErrIntInit()

/**
 * @brief get the CPU which caused cache invalid access interrupt. Helper function in
 * panic handling.
 * @return
 *  - PRO_CPU_NUM, if PRO_CPU has caused cache IA interrupt
 *  - APP_CPU_NUM, if APP_CPU has caused cache IA interrupt
 *  - (-1) otherwise
 */
//go:linkname EspCacheErrGetCpuid C.esp_cache_err_get_cpuid
func EspCacheErrGetCpuid() c.Int

/**
 * @brief Returns a pointer to the cache error message
 *
 * @return const char* Pointer to the error message
 */
//go:linkname EspCacheErrPanicString C.esp_cache_err_panic_string
func EspCacheErrPanicString() *c.Char

/**
 * @brief Checks if any cache errors are active
 *
 * @return true
 * @return false
 */
//go:linkname EspCacheErrHasActiveErr C.esp_cache_err_has_active_err
func EspCacheErrHasActiveErr() bool
