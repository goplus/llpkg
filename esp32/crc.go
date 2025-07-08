package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Crc32 value that is in little endian.
 *
 * @param  uint32_t crc : init crc value, use 0 at the first use.
 *
 * @param  uint8_t const *buf : buffer to start calculate crc.
 *
 * @param  uint32_t len : buffer length in byte.
 *
 * @return None
 */
//go:linkname Crc32Le C.crc32_le
func Crc32Le(crc c.Uint32T, buf *c.Uint8T, len c.Uint32T) c.Uint32T

/**
 * @brief Crc32 value that is in big endian.
 *
 * @param  uint32_t crc : init crc value, use 0 at the first use.
 *
 * @param  uint8_t const *buf : buffer to start calculate crc.
 *
 * @param  uint32_t len : buffer length in byte.
 *
 * @return None
 */
//go:linkname Crc32Be C.crc32_be
func Crc32Be(crc c.Uint32T, buf *c.Uint8T, len c.Uint32T) c.Uint32T

/**
 * @brief Crc16 value that is in little endian.
 *
 * @param  uint16_t crc : init crc value, use 0 at the first use.
 *
 * @param  uint8_t const *buf : buffer to start calculate crc.
 *
 * @param  uint32_t len : buffer length in byte.
 *
 * @return None
 */
//go:linkname Crc16Le C.crc16_le
func Crc16Le(crc c.Uint16T, buf *c.Uint8T, len c.Uint32T) c.Uint16T

/**
 * @brief Crc16 value that is in big endian.
 *
 * @param  uint16_t crc : init crc value, use 0 at the first use.
 *
 * @param  uint8_t const *buf : buffer to start calculate crc.
 *
 * @param  uint32_t len : buffer length in byte.
 *
 * @return None
 */
//go:linkname Crc16Be C.crc16_be
func Crc16Be(crc c.Uint16T, buf *c.Uint8T, len c.Uint32T) c.Uint16T

/**
 * @brief Crc8 value that is in little endian.
 *
 * @param  uint8_t crc : init crc value, use 0 at the first use.
 *
 * @param  uint8_t const *buf : buffer to start calculate crc.
 *
 * @param  uint32_t len : buffer length in byte.
 *
 * @return None
 */
//go:linkname Crc8Le C.crc8_le
func Crc8Le(crc c.Uint8T, buf *c.Uint8T, len c.Uint32T) c.Uint8T

/**
 * @brief Crc8 value that is in big endian.
 *
 * @param  uint32_t crc : init crc value, use 0 at the first use.
 *
 * @param  uint8_t const *buf : buffer to start calculate crc.
 *
 * @param  uint32_t len : buffer length in byte.
 *
 * @return None
 */
//go:linkname Crc8Be C.crc8_be
func Crc8Be(crc c.Uint8T, buf *c.Uint8T, len c.Uint32T) c.Uint8T
