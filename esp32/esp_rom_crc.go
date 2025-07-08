package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief CRC32 value in little endian.
 *
 * @param crc: Initial CRC value (result of last calculation or 0 for the first time)
 * @param buf: Data buffer that used to calculate the CRC value
 * @param len: Length of the data buffer
 * @return CRC32 value
 */
//go:linkname EspRomCrc32Le C.esp_rom_crc32_le
func EspRomCrc32Le(crc c.Uint32T, buf *c.Uint8T, len c.Uint32T) c.Uint32T

/**
 * @brief CRC32 value in big endian.
 *
 * @param crc: Initial CRC value (result of last calculation or 0 for the first time)
 * @param buf: Data buffer that used to calculate the CRC value
 * @param len: Length of the data buffer
 * @return CRC32 value
 */
//go:linkname EspRomCrc32Be C.esp_rom_crc32_be
func EspRomCrc32Be(crc c.Uint32T, buf *c.Uint8T, len c.Uint32T) c.Uint32T

/**
 * @brief CRC16 value in little endian.
 *
 * @param crc: Initial CRC value (result of last calculation or 0 for the first time)
 * @param buf: Data buffer that used to calculate the CRC value
 * @param len: Length of the data buffer
 * @return CRC16 value
 */
//go:linkname EspRomCrc16Le C.esp_rom_crc16_le
func EspRomCrc16Le(crc c.Uint16T, buf *c.Uint8T, len c.Uint32T) c.Uint16T

/**
 * @brief CRC16 value in big endian.
 *
 * @param crc: Initial CRC value (result of last calculation or 0 for the first time)
 * @param buf: Data buffer that used to calculate the CRC value
 * @param len: Length of the data buffer
 * @return CRC16 value
 */
//go:linkname EspRomCrc16Be C.esp_rom_crc16_be
func EspRomCrc16Be(crc c.Uint16T, buf *c.Uint8T, len c.Uint32T) c.Uint16T

/**
 * @brief CRC8 value in little endian.
 *
 * @param crc: Initial CRC value (result of last calculation or 0 for the first time)
 * @param buf: Data buffer that used to calculate the CRC value
 * @param len: Length of the data buffer
 * @return CRC8 value
 */
//go:linkname EspRomCrc8Le C.esp_rom_crc8_le
func EspRomCrc8Le(crc c.Uint8T, buf *c.Uint8T, len c.Uint32T) c.Uint8T

/**
 * @brief CRC8 value in big endian.
 *
 * @param crc: Initial CRC value (result of last calculation or 0 for the first time)
 * @param buf: Data buffer that used to calculate the CRC value
 * @param len: Length of the data buffer
 * @return CRC8 value
 */
//go:linkname EspRomCrc8Be C.esp_rom_crc8_be
func EspRomCrc8Be(crc c.Uint8T, buf *c.Uint8T, len c.Uint32T) c.Uint8T
