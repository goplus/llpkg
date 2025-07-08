package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief A CRC8 algorithm used for MAC addresses stored in eFuse
 *
 * @param data Pointer to the original data
 * @param len Data length in byte
 * @return uint8_t CRC value
 */
//go:linkname EspRomEfuseMacAddressCrc8 C.esp_rom_efuse_mac_address_crc8
func EspRomEfuseMacAddressCrc8(data *c.Uint8T, len c.Uint32T) c.Uint8T

/**
 * @brief Get SPI Flash GPIO pin configurations from eFuse
 *
 * @return uint32_t
 *          - 0: default SPI pins (ESP_ROM_EFUSE_FLASH_DEFAULT_SPI)
 *          - 1: default HSPI pins (ESP_ROM_EFUSE_FLASH_DEFAULT_HSPI)
 *          - Others: Customized pin configuration mask. Pins are encoded as per the
 *                    EFUSE_SPICONFIG_RET_SPICLK, EFUSE_SPICONFIG_RET_SPIQ, EFUSE_SPICONFIG_RET_SPID,
 *                    EFUSE_SPICONFIG_RET_SPICS0, EFUSE_SPICONFIG_RET_SPIHD macros.
 *
 * @note WP pin (for quad I/O modes) is not saved in eFuse and not returned by this function.
 */
//go:linkname EspRomEfuseGetFlashGpioInfo C.esp_rom_efuse_get_flash_gpio_info
func EspRomEfuseGetFlashGpioInfo() c.Uint32T

/**
 * @brief Get SPI Flash WP pin information from eFuse
 *
 * @return uint32_t
 *      - 0x3F: invalid GPIO number
 *      - 0~46: valid GPIO number
 */
//go:linkname EspRomEfuseGetFlashWpGpio C.esp_rom_efuse_get_flash_wp_gpio
func EspRomEfuseGetFlashWpGpio() c.Uint32T

/**
 * @brief Read opi flash pads configuration from Efuse
 *
 * @return
 * - 0 for default SPI pins.
 * - Other values define a custom pin configuration mask. From the LSB, every 6 bits represent a GPIO number which stand for:
 *   DQS, D4, D5, D6, D7 accordingly.
 */
//go:linkname EspRomEfuseGetOpiconfig C.esp_rom_efuse_get_opiconfig
func EspRomEfuseGetOpiconfig() c.Uint32T

/**
 * @brief Read eFuse to check whether secure boot has been enabled or not
 *
 * @return true if secure boot is enabled, otherwise false
 */
//go:linkname EspRomEfuseIsSecureBootEnabled C.esp_rom_efuse_is_secure_boot_enabled
func EspRomEfuseIsSecureBootEnabled() bool
