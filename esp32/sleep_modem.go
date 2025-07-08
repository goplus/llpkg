package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Whether the current target allows Modem or the TOP power domain to be powered off during light sleep
 *
 * During light sleep on some targets, it is possible to power OFF the Modem or TOP
 * power domains in order to further lower power power consumption. However, this
 * can only occur on targets that support REGDMA for modem (WiFi, Bluetooth,
 * IEEE802.15.4) retention.
 */
//go:linkname ModemDomainPdAllowed C.modem_domain_pd_allowed
func ModemDomainPdAllowed() bool

/**
 * @brief Get the reject trigger signal of Modem system
 *
 * @return the reject trigger signal of Modem system.
 */
//go:linkname SleepModemRejectTriggers C.sleep_modem_reject_triggers
func SleepModemRejectTriggers() c.Uint32T

/**
 * @brief Configure the parameters of the modem subsystem during the sleep process
 *
 * In light sleep mode, the wake-up early time of the WiFi module and the TBTT
 * interrupt early time (trigger enabling RF) are determined by the maximum and
 * minimum frequency of system (higher system frequency means less time to wake
 * up and enable RF).
 * For the esp32c6 SOC, the modem state is strongly dependent on the light sleep
 * mode, and the modem state will be enabled only when light sleep is enabled
 * and the `CONFIG_ESP_WIFI_ENHANCED_LIGHT_SLEEP` is configured in menuconfig.
 *
 * @param max_freq_mhz       the maximum frequency of system
 * @param min_freq_mhz       the minimum frequency of system
 * @param light_sleep_enable true or false for enable or disable light sleep mode, respectively
 *
 * @return
 *  - ESP_OK on success
 */
//go:linkname SleepModemConfigure C.sleep_modem_configure
func SleepModemConfigure(max_freq_mhz c.Int, min_freq_mhz c.Int, light_sleep_enable bool) EspErrT

// llgo:type C
type InformOutLightSleepOverheadCbT func(c.Uint32T)

/**
 * @brief  Register informing peripherals light sleep wakeup overhead time callback
 *
 * This function allows you to register a callback that informs the peripherals of
 * the wakeup overhead time of light sleep.
 * @param cb function to inform time
 * @return
 *   - ESP_OK on success
 *   - ESP_ERR_NO_MEM if no more callback slots are available
 */
//go:linkname EspPmRegisterInformOutLightSleepOverheadCallback C.esp_pm_register_inform_out_light_sleep_overhead_callback
func EspPmRegisterInformOutLightSleepOverheadCallback(cb InformOutLightSleepOverheadCbT) EspErrT

/**
 * @brief  Unregister informing peripherals light sleep wakeup overhead time callback
 *
 * This function allows you to unregister a callback that informs the peripherals of
 * the wakeup overhead time of light sleep.
 * @param cb function to inform time
 * @return
 *   - ESP_OK on success
 *   - ESP_ERR_INVALID_STATE if the given callback hasn't been registered before
 */
//go:linkname EspPmUnregisterInformOutLightSleepOverheadCallback C.esp_pm_unregister_inform_out_light_sleep_overhead_callback
func EspPmUnregisterInformOutLightSleepOverheadCallback(cb InformOutLightSleepOverheadCbT) EspErrT

/**
 * @brief  A callback that informs the peripherals of the wakeup overhead time of light sleep
 *
 * @param out_light_sleep_time wakeup overhead time of light sleep
 */
//go:linkname PeriphInformOutLightSleepOverhead C.periph_inform_out_light_sleep_overhead
func PeriphInformOutLightSleepOverhead(out_light_sleep_time c.Uint32T)

// llgo:type C
type UpdateLightSleepDefaultParamsConfigCbT func(c.Int, c.Int)

/**
 * @brief  Register peripherals light sleep default parameters configure callback
 *
 * This function allows you to register a callback that configure the peripherals
 * of default parameters of light sleep
 * @param cb function to update default parameters
 */
//go:linkname EspPmRegisterLightSleepDefaultParamsConfigCallback C.esp_pm_register_light_sleep_default_params_config_callback
func EspPmRegisterLightSleepDefaultParamsConfigCallback(cb UpdateLightSleepDefaultParamsConfigCbT)

/**
 * @brief  Unregister peripherals light sleep default parameters configure Callback
 *
 * This function allows you to unregister a callback that configure the peripherals
 * of default parameters of light sleep
 */
//go:linkname EspPmUnregisterLightSleepDefaultParamsConfigCallback C.esp_pm_unregister_light_sleep_default_params_config_callback
func EspPmUnregisterLightSleepDefaultParamsConfigCallback()
