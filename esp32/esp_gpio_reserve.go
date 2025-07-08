package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Reserve the given GPIOs by mask, so they can't be used by others
 *
 * @param gpio_mask Mask of the GPIOs to be reserved
 * @return The mask of the GPIOs that were already reserved before this call
 */
//go:linkname EspGpioReserve C.esp_gpio_reserve
func EspGpioReserve(gpio_mask c.Uint64T) c.Uint64T

/**
 * @brief Revoke the given GPIOs by mask, so they can be reused again by others
 *
 * @param gpio_mask Mask of the GPIOs to be revoked
 * @return The mask of the GPIOs that were already reserved before this call
 */
//go:linkname EspGpioRevoke C.esp_gpio_revoke
func EspGpioRevoke(gpio_mask c.Uint64T) c.Uint64T

/**
 * @brief Check whether the given GPIOs are reserved
 *
 * @param gpio_mask Mask of the GPIOs to be checked
 * @return
 *      - true  Any of the given GPIO(s) is reserved
 *      - false Any of the given GPIO(s) is not reserved
 */
//go:linkname EspGpioIsReserved C.esp_gpio_is_reserved
func EspGpioIsReserved(gpio_mask c.Uint64T) bool
