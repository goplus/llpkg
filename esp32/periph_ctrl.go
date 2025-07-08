package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** @cond */
// The following functions are not intended to be used directly by the developers
// llgo:link PeriphModuleT.PeriphRccAcquireEnter C.periph_rcc_acquire_enter
func (recv_ PeriphModuleT) PeriphRccAcquireEnter() c.Uint8T {
	return 0
}

// llgo:link PeriphModuleT.PeriphRccAcquireExit C.periph_rcc_acquire_exit
func (recv_ PeriphModuleT) PeriphRccAcquireExit(ref_count c.Uint8T) {
}

// llgo:link PeriphModuleT.PeriphRccReleaseEnter C.periph_rcc_release_enter
func (recv_ PeriphModuleT) PeriphRccReleaseEnter() c.Uint8T {
	return 0
}

// llgo:link PeriphModuleT.PeriphRccReleaseExit C.periph_rcc_release_exit
func (recv_ PeriphModuleT) PeriphRccReleaseExit(ref_count c.Uint8T) {
}

//go:linkname PeriphRccEnter C.periph_rcc_enter
func PeriphRccEnter()

//go:linkname PeriphRccExit C.periph_rcc_exit
func PeriphRccExit()

/**
 * @brief Enable peripheral module by un-gating the clock and de-asserting the reset signal.
 *
 * @param[in] periph Peripheral module
 *
 * @note If @c periph_module_enable() is called a number of times,
 *       @c periph_module_disable() has to be called the same number of times,
 *       in order to put the peripheral into disabled state.
 */
// llgo:link PeriphModuleT.PeriphModuleEnable C.periph_module_enable
func (recv_ PeriphModuleT) PeriphModuleEnable() {
}

/**
 * @brief Disable peripheral module by gating the clock and asserting the reset signal.
 *
 * @param[in] periph Peripheral module
 *
 * @note If @c periph_module_enable() is called a number of times,
 *       @c periph_module_disable() has to be called the same number of times,
 *       in order to put the peripheral into disabled state.
 */
// llgo:link PeriphModuleT.PeriphModuleDisable C.periph_module_disable
func (recv_ PeriphModuleT) PeriphModuleDisable() {
}

/**
 * @brief Reset peripheral module by asserting and de-asserting the reset signal.
 *
 * @param[in] periph Peripheral module
 *
 * @note Calling this function does not enable or disable the clock for the module.
 */
// llgo:link PeriphModuleT.PeriphModuleReset C.periph_module_reset
func (recv_ PeriphModuleT) PeriphModuleReset() {
}

/**
 * @brief Enable Wi-Fi and BT common module
 *
 * @note If @c wifi_bt_common_module_enable() is called a number of times,
 *       @c wifi_bt_common_module_disable() has to be called the same number of times,
 *       in order to put the peripheral into disabled state.
 */
//go:linkname WifiBtCommonModuleEnable C.wifi_bt_common_module_enable
func WifiBtCommonModuleEnable()

/**
 * @brief Disable Wi-Fi and BT common module
 *
 * @note If @c wifi_bt_common_module_enable() is called a number of times,
 *       @c wifi_bt_common_module_disable() has to be called the same number of times,
 *       in order to put the peripheral into disabled state.
 */
//go:linkname WifiBtCommonModuleDisable C.wifi_bt_common_module_disable
func WifiBtCommonModuleDisable()

/**
 * @brief Enable Wi-Fi module
 *
 * @note Calling this function will only enable Wi-Fi module.
 */
//go:linkname WifiModuleEnable C.wifi_module_enable
func WifiModuleEnable()

/**
 * @brief Disable Wi-Fi module
 *
 * @note Calling this function will only disable Wi-Fi module.
 */
//go:linkname WifiModuleDisable C.wifi_module_disable
func WifiModuleDisable()
