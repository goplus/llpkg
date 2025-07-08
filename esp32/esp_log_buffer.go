package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Logs a buffer of hexadecimal bytes at the specified log level.
 *
 * This function logs a buffer of hexadecimal bytes with 16 bytes per line. The
 * log level determines the severity of the log message.
 *
 * @note This function does not check the log level against the ESP_LOCAL_LEVEL.
 * The log level comparison should be done in esp_log.h.
 *
 * @param tag       Description tag to identify the log.
 * @param buffer    Pointer to the buffer array containing the data to be logged.
 * @param buff_len  Length of the buffer in bytes.
 * @param level     Log level indicating the severity of the log message.
 */
//go:linkname EspLogBufferHexInternal C.esp_log_buffer_hex_internal
func EspLogBufferHexInternal(tag *c.Char, buffer c.Pointer, buff_len c.Uint16T, level EspLogLevelT)

/**
 * @brief This function logs a buffer of characters with 16 characters per line.
 * The buffer should contain only printable characters. The log level determines
 * the severity of the log message.
 *
 * @note This function does not check the log level against the ESP_LOCAL_LEVEL.
 * The log level comparison should be done in esp_log.h.
 *
 * @param tag       Description tag to identify the log.
 * @param buffer    Pointer to the buffer array containing the data to be logged.
 * @param buff_len  Length of the buffer in bytes.
 * @param level     Log level indicating the severity of the log message.
 */
//go:linkname EspLogBufferCharInternal C.esp_log_buffer_char_internal
func EspLogBufferCharInternal(tag *c.Char, buffer c.Pointer, buff_len c.Uint16T, level EspLogLevelT)

/**
 * @brief This function dumps a buffer to the log in a formatted hex dump style,
 * displaying both the memory address and the corresponding hex and ASCII values
 * of the bytes. The log level determines the severity of the log message.
 *
 * @note This function does not check the log level against the ESP_LOCAL_LEVEL.
 * The log level comparison should be done in esp_log.h.
 * @note It is recommended to use terminals with a width of at least 102
 * characters to display the log dump properly.
 *
 * @param tag       Description tag to identify the log.
 * @param buffer    Pointer to the buffer array containing the data to be logged.
 * @param buff_len  Length of the buffer in bytes.
 * @param log_level Log level indicating the severity of the log message.
 */
//go:linkname EspLogBufferHexdumpInternal C.esp_log_buffer_hexdump_internal
func EspLogBufferHexdumpInternal(tag *c.Char, buffer c.Pointer, buff_len c.Uint16T, log_level EspLogLevelT)
