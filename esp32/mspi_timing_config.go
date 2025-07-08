package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*-------------------------------------------------------------------------------------------------
 * Generic Config APIs
 *-------------------------------------------------------------------------------------------------*/
/**
 * @brief Set Flash module clock
 *
 * @param flash_freq_mhz     Flash clock frequency in MHz
 * @param speed_mode         Speed mode
 * @param control_both_mspi  Control SPI1 as well
 */
//go:linkname MspiTimingConfigSetFlashClock C.mspi_timing_config_set_flash_clock
func MspiTimingConfigSetFlashClock(flash_freq_mhz c.Uint32T, speed_mode MspiTimingSpeedModeT, control_both_mspi bool)

/**
 * @brief Set PSRAM module clock
 *
 * @param psram_freq_mhz     PSRAM clock frequency in MHz
 * @param speed_mode         Speed mode
 * @param control_both_mspi  Not used, for compatibility
 */
//go:linkname MspiTimingConfigSetPsramClock C.mspi_timing_config_set_psram_clock
func MspiTimingConfigSetPsramClock(psram_freq_mhz c.Uint32T, speed_mode MspiTimingSpeedModeT, control_both_mspi bool)
