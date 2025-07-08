package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type VprintfLikeT func(*c.Char, c.VaList) c.Int

/**
 * @brief Set function used to output log entries
 *
 * By default, log output goes to UART0. This function can be used to redirect log
 * output to some other destination, such as file or network. Returns the original
 * log handler, which may be necessary to return output to the previous destination.
 *
 * @note Please note that function callback here must be re-entrant as it can be
 * invoked in parallel from multiple thread context.
 *
 * @param func new Function used for output. Must have same signature as vprintf.
 *
 * @return func old Function used for output.
 */
//go:linkname EspLogSetVprintf C.esp_log_set_vprintf
func EspLogSetVprintf(func_ VprintfLikeT) VprintfLikeT

/**
 * @brief Write message into the log
 *
 * This function is not intended to be used directly. Instead, use one of
 * ESP_LOGE, ESP_LOGW, ESP_LOGI, ESP_LOGD, ESP_LOGV macros.
 *
 * This function or these macros should not be used from an interrupt.
 */
// llgo:link EspLogLevelT.EspLogWrite C.esp_log_write
func (recv_ EspLogLevelT) EspLogWrite(tag *c.Char, format *c.Char, __llgo_va_list ...interface{}) {
}

/**
 * @brief Write message into the log, va_list variant
 * @see esp_log_write()
 *
 * This function is provided to ease integration toward other logging framework,
 * so that esp_log can be used as a log sink.
 */
// llgo:link EspLogLevelT.EspLogWritev C.esp_log_writev
func (recv_ EspLogLevelT) EspLogWritev(tag *c.Char, format *c.Char, args c.VaList) {
}
