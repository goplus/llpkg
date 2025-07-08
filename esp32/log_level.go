package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Set the default log level.
 *
 * This function sets the default log level used by the ESP-IDF log library.
 * The default log level is used by the definition of ESP_LOGx macros and
 * can be overridden for specific tags using `esp_log_level_set("*", level)`.
 * This is for internal use only.
 *
 * If CONFIG_LOG_DYNAMIC_LEVEL_CONTROL is not set, this function does not change the default log level.
 *
 * @param level The new default log level to set.
 */
// llgo:link EspLogLevelT.EspLogSetDefaultLevel C.esp_log_set_default_level
func (recv_ EspLogLevelT) EspLogSetDefaultLevel() {
}

/**
 * @brief Get log level for a given tag, can be used to avoid expensive log statements
 *
 * The function takes the lock before checking the tag level,
 * if the lock wait time exceeds 10 ms (default), then ESP_LOG_NONE is returned.
 *
 * @param tag   Tag of the log to query current level. Must be a zero terminated string.
 *              If tag is NULL then the default log level is returned (see esp_log_get_default_level()).
 * @return      The current log level for the given tag.
 */
//go:linkname EspLogLevelGetTimeout C.esp_log_level_get_timeout
func EspLogLevelGetTimeout(tag *c.Char) EspLogLevelT
