package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Configure IO Pad as General Purpose IO,
 *        so that it can be connected to internal Matrix,
 *        then combined with one or more peripheral signals.
 *
 * @param iopad_num IO Pad number
 */
//go:linkname EspRomGpioPadSelectGpio C.esp_rom_gpio_pad_select_gpio
func EspRomGpioPadSelectGpio(iopad_num c.Uint32T)

/**
 * @brief Enable internal pull up, and disable internal pull down.
 *
 * @param iopad_num IO Pad number
 */
//go:linkname EspRomGpioPadPullupOnly C.esp_rom_gpio_pad_pullup_only
func EspRomGpioPadPullupOnly(iopad_num c.Uint32T)

/**
 * @brief Unhold the IO Pad.
 * @note When the Pad is set to hold, the state is latched at that moment and won't get changed.
 *
 * @param iopad_num IP Pad number
 */
//go:linkname EspRomGpioPadUnhold C.esp_rom_gpio_pad_unhold
func EspRomGpioPadUnhold(gpio_num c.Uint32T)

/**
 * @brief Set IO Pad current drive capability.
 *
 * @param iopad_num IO Pad number
 * @param drv Numeric to indicate the capability of current drive
 *      - 0: 5mA
 *      - 1: 10mA
 *      - 2: 20mA
 *      - 3: 40mA
 */
//go:linkname EspRomGpioPadSetDrv C.esp_rom_gpio_pad_set_drv
func EspRomGpioPadSetDrv(iopad_num c.Uint32T, drv c.Uint32T)

/**
 * @brief Combine a GPIO input with a peripheral signal, which tagged as input attribute.
 *
 * @note There's no limitation on the number of signals that a GPIO can combine with.
 *
 * @param gpio_num GPIO number, especially, `GPIO_MATRIX_CONST_ZERO_INPUT` means connect logic 0 to signal
 *                                          `GPIO_MATRIX_CONST_ONE_INPUT` means connect logic 1 to signal
 * @param signal_idx Peripheral signal index (tagged as input attribute)
 * @param inv  Whether the GPIO input to be inverted or not
 */
//go:linkname EspRomGpioConnectInSignal C.esp_rom_gpio_connect_in_signal
func EspRomGpioConnectInSignal(gpio_num c.Uint32T, signal_idx c.Uint32T, inv bool)

/**
 * @brief Combine a peripheral signal which tagged as output attribute with a GPIO.
 *
 * @note There's no limitation on the number of signals that a GPIO can combine with.
 * @note Internally, the signal will be connected first, then output will be enabled on the pad.
 *
 * @param gpio_num GPIO number
 * @param signal_idx Peripheral signal index (tagged as output attribute). Particularly, `SIG_GPIO_OUT_IDX` means disconnect GPIO and other peripherals. Only the GPIO driver can control the output level.
 * @param out_inv Whether to signal to be inverted or not
 * @param oen_inv Whether the output enable control is inverted or not
 */
//go:linkname EspRomGpioConnectOutSignal C.esp_rom_gpio_connect_out_signal
func EspRomGpioConnectOutSignal(gpio_num c.Uint32T, signal_idx c.Uint32T, out_inv bool, oen_inv bool)
