package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Sync MSPI speed mode before CPU frequency switching, only needed when frequency is decreasing.
 *
 * @param target_cpu_src_freq   Target clock source frequency for CPU frequency switching
 * @param target_cpu_freq       CPU frequency switching target frequency
 */
//go:linkname EspClkUtilsMspiSpeedModeSyncBeforeCpuFreqSwitching C.esp_clk_utils_mspi_speed_mode_sync_before_cpu_freq_switching
func EspClkUtilsMspiSpeedModeSyncBeforeCpuFreqSwitching(target_cpu_src_freq c.Uint32T, target_cpu_freq c.Uint32T)

/**
 * @brief Sync MSPI speed mode after CPU frequency switching, only needed when frequency is upcreasing.
 *
 * @param target_cpu_src_freq Target clock source frequency for CPU frequency switching
 * @param target_cpu_freq     CPU frequency switching target frequency
 */
//go:linkname EspClkUtilsMspiSpeedModeSyncAfterCpuFreqSwitching C.esp_clk_utils_mspi_speed_mode_sync_after_cpu_freq_switching
func EspClkUtilsMspiSpeedModeSyncAfterCpuFreqSwitching(target_cpu_src_freq c.Uint32T, target_cpu_freq c.Uint32T)
