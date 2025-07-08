package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Enable the clock of modem module
 *
 * Solve the clock dependency between modem modules, For example, the wifi
 * module depends on the wifi mac, wifi baseband and FE, when wifi module
 * clock is enabled, the wifi MAC, baseband and FE clocks will be enabled
 *
 * This interface and modem_clock_module_disable will jointly maintain the
 * ref_cnt of each device clock source. The ref_cnt indicates how many modules
 * are relying on the clock source. Each enable ops will add 1 to the ref_cnt of
 * the clock source that the module depends on, and only when the ref_cnt of
 * the module is from 0 to 1 will the clock enable be actually configured.
 *
 * !!! Do not use the hal/ll layer interface to configure the clock for the
 * consistency of the hardware state maintained in the driver and the hardware
 * actual state.
 *
 * @param module  modem module
 */
// llgo:link PeriphModuleT.ModemClockModuleEnable C.modem_clock_module_enable
func (recv_ PeriphModuleT) ModemClockModuleEnable() {
}

/**
 * @brief Disable the clock of modem module
 *
 * This interface and modem_clock_module_enable will jointly maintain the ref_cnt
 * of each device clock source. The ref_cnt indicates how many modules are relying
 * on the clock source. Each disable ops will minus 1 to the ref_cnt of the clock
 * source that the module depends on, and only when the ref_cnt of the module is
 * from 1 to 0 will the clock disable be actually configured.
 *
 * !!! Do not use the hal/ll layer interface to configure the clock for the
 * consistency of the hardware state maintained in the driver and the hardware
 * actual state.
 *
 * @param module  modem module
 */
// llgo:link PeriphModuleT.ModemClockModuleDisable C.modem_clock_module_disable
func (recv_ PeriphModuleT) ModemClockModuleDisable() {
}

/**
 * @brief Reset the mac of modem module
 *
 * @param module  modem module, must be one of
 *    PERIPH_WIFI_MODULE / PERIPH_BT_MODULE /PERIPH_IEEE802154_MODULE
 */
// llgo:link PeriphModuleT.ModemClockModuleMacReset C.modem_clock_module_mac_reset
func (recv_ PeriphModuleT) ModemClockModuleMacReset() {
}

/**
 * @brief Select the modem module lowpower clock source and configure the clock divider
 *
 * @param module  modem module
 * @param src     lowpower clock source
 * @param divider divider value to lowpower clock source
 */
// llgo:link PeriphModuleT.ModemClockSelectLpClockSource C.modem_clock_select_lp_clock_source
func (recv_ PeriphModuleT) ModemClockSelectLpClockSource(src ModemClockLpclkSrcT, divider c.Uint32T) {
}

/**
 * @brief Disable lowpower clock source selection
 * @param module  modem module
 */
// llgo:link PeriphModuleT.ModemClockDeselectLpClockSource C.modem_clock_deselect_lp_clock_source
func (recv_ PeriphModuleT) ModemClockDeselectLpClockSource() {
}

/**
* @brief Disable all modem module's lowpower clock source selection
 */
//go:linkname ModemClockDeselectAllModuleLpClockSource C.modem_clock_deselect_all_module_lp_clock_source
func ModemClockDeselectAllModuleLpClockSource()

/**
 * @brief Reset wifi mac
 */
//go:linkname ModemClockWifiMacReset C.modem_clock_wifi_mac_reset
func ModemClockWifiMacReset()

/**
 * @brief Enable clock registers which shared by both modem and ADC. Need a ref count to enable/disable them
 *
 * @param enable true: enable; false: disable
 */
//go:linkname ModemClockSharedEnable C.modem_clock_shared_enable
func ModemClockSharedEnable(enable bool)
