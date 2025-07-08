package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspSleepSubModeT c.Int

const (
	ESP_SLEEP_RTC_USE_RC_FAST_MODE      EspSleepSubModeT = 0
	ESP_SLEEP_DIG_USE_RC_FAST_MODE      EspSleepSubModeT = 1
	ESP_SLEEP_USE_ADC_TSEN_MONITOR_MODE EspSleepSubModeT = 2
	ESP_SLEEP_ULTRA_LOW_MODE            EspSleepSubModeT = 3
	ESP_SLEEP_RTC_FAST_USE_XTAL_MODE    EspSleepSubModeT = 4
	ESP_SLEEP_DIG_USE_XTAL_MODE         EspSleepSubModeT = 5
	ESP_SLEEP_LP_USE_XTAL_MODE          EspSleepSubModeT = 6
	ESP_SLEEP_MODE_MAX                  EspSleepSubModeT = 7
)

/**
 * @brief Set sub-sleep power mode in sleep, mode enabled status is maintained by reference count.
 *        The caller should ensure that the enabling and disabling behavior is symmetric.
 *        This submode configuration will kept after deep sleep wakeup.
 *
 * @param  mode     sub-sleep mode type
 * @param  activate Activate or deactivate the sleep sub mode
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if either of the arguments is out of range
 */
// llgo:link EspSleepSubModeT.EspSleepSubModeConfig C.esp_sleep_sub_mode_config
func (recv_ EspSleepSubModeT) EspSleepSubModeConfig(activate bool) EspErrT {
	return 0
}

/**
 * @brief Force disable sub-sleep power mode in sleep, usually used during initialization.
 *
 * @param  mode     sub-sleep mode type
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if either of the arguments is out of range
 */
// llgo:link EspSleepSubModeT.EspSleepSubModeForceDisable C.esp_sleep_sub_mode_force_disable
func (recv_ EspSleepSubModeT) EspSleepSubModeForceDisable() EspErrT {
	return 0
}

/**
 * Dump the sub-sleep power mode enable status
 * @param  stream     The stream to dump to, if NULL then nothing will be dumped
 * @return            return the reference count array pointer
 */
//go:linkname EspSleepSubModeDumpConfig C.esp_sleep_sub_mode_dump_config
func EspSleepSubModeDumpConfig(stream *c.FILE) *c.Int32T

/**
 * @brief Isolate all digital IOs except those that are held during deep sleep
 *
 * Reduce digital IOs current leakage during deep sleep.
 */
//go:linkname EspSleepIsolateDigitalGpio C.esp_sleep_isolate_digital_gpio
func EspSleepIsolateDigitalGpio()

/**
 * Register a callback to be called from the deep sleep prepare for maintain the PHY state
 *          CPU is equal to min_freq_mhz (if DFS is enabled) when running this callback,
 *          and PLL clock is exists)
 *
 * @warning deepsleep PHY callbacks should without parameters, and MUST NOT,
 *          UNDER ANY CIRCUMSTANCES, CALL A FUNCTION THAT MIGHT BLOCK.
 *
 * @param new_dslp_cb     Callback to be called to close PHY related modules
 *
 * @return
 *     - ESP_OK:         PHY callback registered to the phy modules deepsleep prepare
 *     - ESP_ERR_NO_MEM: No more hook space for register the callback
 */
//go:linkname EspDeepSleepRegisterPhyHook C.esp_deep_sleep_register_phy_hook
func EspDeepSleepRegisterPhyHook(new_dslp_cb EspDeepSleepCbT) EspErrT

/**
 * @brief Unregister an PHY deepsleep callback
 *
 * @param old_dslp_cb     Callback to be unregistered
 */
//go:linkname EspDeepSleepDeregisterPhyHook C.esp_deep_sleep_deregister_phy_hook
func EspDeepSleepDeregisterPhyHook(old_dslp_cb EspDeepSleepCbT)

/**
 * @brief Notify the sleep process that `sleep_time_overhead_out` needs to be remeasured, which must be called
 *        in the following scenarios:
 *        1. When the CPU frequency changes to below the crystal oscillator frequency.
 *        2. When a new callback function is registered in the sleep process.
 *        3. Other events occur that affect the execution time of the CPU sleep process.
 */
//go:linkname EspSleepOverheadOutTimeRefresh C.esp_sleep_overhead_out_time_refresh
func EspSleepOverheadOutTimeRefresh()
