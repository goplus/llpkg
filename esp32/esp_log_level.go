package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspLogLevelT c.Int

const (
	ESP_LOG_NONE    EspLogLevelT = 0
	ESP_LOG_ERROR   EspLogLevelT = 1
	ESP_LOG_WARN    EspLogLevelT = 2
	ESP_LOG_INFO    EspLogLevelT = 3
	ESP_LOG_DEBUG   EspLogLevelT = 4
	ESP_LOG_VERBOSE EspLogLevelT = 5
	ESP_LOG_MAX     EspLogLevelT = 6
)

/**
 * @brief Set log level for given tag
 *
 * If logging for given component has already been enabled, changes previous setting.
 *
 * @note Note that this function can not raise log level above the level set using
 *       CONFIG_LOG_MAXIMUM_LEVEL setting in menuconfig.
 *
 * To raise log level above the default one for a given file, define
 * LOG_LOCAL_LEVEL to one of the ESP_LOG_* values, before including esp_log.h in this file.
 *
 * If CONFIG_LOG_DYNAMIC_LEVEL_CONTROL is not selected the static (no-op) implementation of log level is used.
 * Changing the log level is not possible, esp_log_level_set does not work.
 *
 * @param tag   Tag of the log entries to enable. Must be a non-NULL zero terminated string.
 *              Value "*" resets log level for all tags to the given value.
 *              If the tag is NULL then a silent return happens.
 * @param level Selects log level to enable.
 *              Only logs at this and lower verbosity levels will be shown.
 */
//go:linkname EspLogLevelSet C.esp_log_level_set
func EspLogLevelSet(tag *c.Char, level EspLogLevelT)

/**
 * @brief Get log level for a given tag, can be used to avoid expensive log statements
 *
 * If CONFIG_LOG_DYNAMIC_LEVEL_CONTROL is not selected the static (no-op) implementation of log level is used.
 * Changing the log level is not possible, esp_log_level_set does not work. This function returns the default log level.
 *
 * @param tag   Tag of the log to query current level. Must be a zero terminated string.
 *              If tag is NULL then the default log level is returned (see esp_log_get_default_level()).
 * @return      The current log level for the given tag.
 */
//go:linkname EspLogLevelGet C.esp_log_level_get
func EspLogLevelGet(tag *c.Char) EspLogLevelT
