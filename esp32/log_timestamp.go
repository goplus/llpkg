package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get the timestamp in milliseconds for logging.
 *
 * This function retrieves the timestamp in milliseconds for logging purposes.
 *
 * @param[in] critical Flag indicating whether the timestamp is needed for a critical log.
 * If this flag is true, it means the function is called in one of the following states:
 * - early stage, when the Freertos scheduler is not running,
 * - ISR,
 * - CACHE is disabled.
 *
 * If the critical flag is set then uint32 timestamp is returned due to cpu ticks being used for this case.
 * For some applications this uint32 timestamp may overflow after 4294967295/1000/86400 = 49 days of operation.
 *
 * @return The uint64 timestamp in milliseconds.
 */
//go:linkname EspLogTimestamp64 C.esp_log_timestamp64
func EspLogTimestamp64(critical bool) c.Uint64T

/**
 * @brief Convert the uint64 timestamp to a string representation.
 *
 * This function converts the uint64 timestamp in milliseconds to a string representation.
 * The string representation depends on Kconfig options:
 * - Milliseconds since boot,
 * - Date and time,
 * - Time.
 *
 * @param[in] critical Flag indicating whether the timestamp is critical. If this flag is true,
 * it means the function is called in one of the following states:
 * - early stage, when the Freertos scheduler is not running,
 * - ISR,
 * - CACHE is disabled.
 * @param[in] timestamp_ms The timestamp to convert, in milliseconds.
 * @param[out] buffer   Pointer to the buffer where the string representation will be stored.
 *
 * @return Pointer to the buffer containing the string representation of the timestamp.
 */
//go:linkname EspLogTimestampStr C.esp_log_timestamp_str
func EspLogTimestampStr(critical bool, timestamp_ms c.Uint64T, buffer *c.Char) *c.Char
