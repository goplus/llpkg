package freertos

import _ "unsafe"

/**
 * @brief Call once in startup to disable the wakeup IO pins and release their holding state after waking up from Deep-sleep
 */
//go:linkname EspDeepSleepWakeupIoReset C.esp_deep_sleep_wakeup_io_reset
func EspDeepSleepWakeupIoReset()
