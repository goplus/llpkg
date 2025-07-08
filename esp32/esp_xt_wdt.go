package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief esp_xt_wdt configuration struct
 *
 */

type EspXtWdtConfigT struct {
	Timeout             c.Uint8T
	AutoBackupClkEnable bool
}

// llgo:type C
type EspXtCallbackT func(c.Pointer)

/**
 * @brief Initializes the xtal32k watchdog timer
 *
 * @param cfg Pointer to configuration struct
 * @return esp_err_t
 *      - ESP_OK: XTWDT was successfully enabled
 *      - ESP_ERR_NO_MEM: Failed to allocate ISR
 */
// llgo:link (*EspXtWdtConfigT).EspXtWdtInit C.esp_xt_wdt_init
func (recv_ *EspXtWdtConfigT) EspXtWdtInit() EspErrT {
	return 0
}

/**
 * @brief Register a callback function that will be called when the watchdog
 *        times out.
 *
 * @note This function will be called from an interrupt context where the cache might be disabled.
 *       Thus the function should be placed in IRAM and must not perform any blocking operations.
 *
 *       Only one callback function can be registered, any call to esp_xt_wdt_register_callback
 *       will override the previous callback function.
 *
 * @param func The callback function to register
 * @param arg  Pointer to argument that will be passed to the callback function
 */
//go:linkname EspXtWdtRegisterCallback C.esp_xt_wdt_register_callback
func EspXtWdtRegisterCallback(func_ EspXtCallbackT, arg c.Pointer)

/**
 * @brief Restores the xtal32k clock and re-enables the WDT
 *
 */
//go:linkname EspXtWdtRestoreClk C.esp_xt_wdt_restore_clk
func EspXtWdtRestoreClk()
