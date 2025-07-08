package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief  Get one random 32-bit word from hardware RNG
 *
 * If Wi-Fi or Bluetooth are enabled, this function returns true random numbers. In other
 * situations, if true random numbers are required then consult the ESP-IDF Programming
 * Guide "Random Number Generation" section for necessary prerequisites.
 *
 * This function automatically busy-waits to ensure enough external entropy has been
 * introduced into the hardware RNG state, before returning a new random number.
 * This delay makes sure the reading frequency does not exceed 15-75 KHz. The actual
 * value is dependent on the specific chip. More information on this can be found in
 * components/esp_hw_support/hw_random.c.
 *
 * @return Random value between 0 and UINT32_MAX
 */
//go:linkname EspRandom C.esp_random
func EspRandom() c.Uint32T

/**
 * @brief Fill a buffer with random bytes from hardware RNG
 *
 * @note This function is implemented via calls to esp_random(), so the same
 * constraints apply.
 *
 * @param buf Pointer to buffer to fill with random numbers.
 * @param len Length of buffer in bytes
 */
//go:linkname EspFillRandom C.esp_fill_random
func EspFillRandom(buf c.Pointer, len c.SizeT)
