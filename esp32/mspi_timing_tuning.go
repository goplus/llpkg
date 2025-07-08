package freertos

import _ "unsafe"

/**
 * @brief Make MSPI work under 20Mhz, remove the timing tuning required delays.
 * @param control_spi1  Select whether to control SPI1. For tuning, we need to use SPI1. After tuning (during startup stage), let the flash driver to control SPI1
 */
//go:linkname MspiTimingEnterLowSpeedMode C.mspi_timing_enter_low_speed_mode
func MspiTimingEnterLowSpeedMode(control_spi1 bool)

/**
 * @brief Make MSPI work under the frequency as users set, may add certain delays to MSPI RX direction to meet timing requirements.
 * @param control_spi1  Select whether to control SPI1. For tuning, we need to use SPI1. After tuning (during startup stage), let the flash driver to control SPI1
 */
//go:linkname MspiTimingEnterHighSpeedMode C.mspi_timing_enter_high_speed_mode
func MspiTimingEnterHighSpeedMode(control_spi1 bool)

/**
 * @brief Switch MSPI into low speed mode / high speed mode.
 * @note This API is cache safe, it will freeze both D$ and I$ and restore them after MSPI is switched
 * @note For some of the MSPI high frequency settings (e.g. 80M DDR mode Flash or PSRAM), timing tuning is required.
 *       Certain delays will be added to the MSPI RX direction. When CPU clock switches from PLL to XTAL, should call
 *       this API first to enter MSPI low speed mode to remove the delays, and vice versa.
 */
//go:linkname MspiTimingChangeSpeedModeCacheSafe C.mspi_timing_change_speed_mode_cache_safe
func MspiTimingChangeSpeedModeCacheSafe(switch_down bool)

/**
 * @brief Tune MSPI flash timing to make it work under high frequency
 */
//go:linkname MspiTimingFlashTuning C.mspi_timing_flash_tuning
func MspiTimingFlashTuning()

/**
 * @brief Tune MSPI psram timing to make it work under high frequency
 */
//go:linkname MspiTimingPsramTuning C.mspi_timing_psram_tuning
func MspiTimingPsramTuning()

/**
 * @brief Set MSPI pin default pin drive
 */
//go:linkname MspiTimingSetPinDriveStrength C.mspi_timing_set_pin_drive_strength
func MspiTimingSetPinDriveStrength()
