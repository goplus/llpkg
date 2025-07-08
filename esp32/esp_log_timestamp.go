package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Function which returns timestamp to be used in log output
 *
 * This function is used in expansion of ESP_LOGx macros.
 * In the 2nd stage bootloader, and at early application startup stage
 * this function uses CPU cycle counter as time source. Later when
 * FreeRTOS scheduler start running, it switches to FreeRTOS tick count.
 *
 * For now, we ignore millisecond counter overflow.
 *
 * @return timestamp, in milliseconds
 */
//go:linkname EspLogTimestamp C.esp_log_timestamp
func EspLogTimestamp() c.Uint32T

/**
 * @brief Function which returns system timestamp to be used in log output
 *
 * This function is used in expansion of ESP_LOGx macros to print
 * the system time as "HH:MM:SS.sss". The system time is initialized to
 * 0 on startup, this can be set to the correct time with an SNTP sync,
 * or manually with standard POSIX time functions.
 *
 * Currently, this will not get used in logging from binary blobs
 * (i.e. Wi-Fi & Bluetooth libraries), these will still print the RTOS tick time.
 *
 * @return timestamp, in "HH:MM:SS.sss"
 */
//go:linkname EspLogSystemTimestamp C.esp_log_system_timestamp
func EspLogSystemTimestamp() *c.Char

/**
 * @brief Function which returns timestamp to be used in log output
 *
 * This function uses HW cycle counter and does not depend on OS,
 * so it can be safely used after application crash.
 *
 * @return timestamp, in milliseconds
 */
//go:linkname EspLogEarlyTimestamp C.esp_log_early_timestamp
func EspLogEarlyTimestamp() c.Uint32T
