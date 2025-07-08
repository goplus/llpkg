package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_IDF_VERSION_MAJOR = 5
const ESP_IDF_VERSION_MINOR = 4
const ESP_IDF_VERSION_PATCH = 2

/**
 * Return full IDF version string, same as 'git describe' output.
 *
 * @note If you are printing the ESP-IDF version in a log file or other information,
 * this function provides more information than using the numerical version macros.
 * For example, numerical version macros don't differentiate between development,
 * pre-release and release versions, but the output of this function does.
 *
 * @return constant string from IDF_VER
 */
//go:linkname EspGetIdfVersion C.esp_get_idf_version
func EspGetIdfVersion() *c.Char
