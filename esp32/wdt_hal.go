package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Context that should be maintained by both the driver and the HAL
 */

type WdtHalContextT struct {
	Inst WdtInstT
}

/**
 * @brief Initialize one of the WDTs associated HAL context
 *
 * This function initializes one of the WDTs (MWDT0, MWDT1, or RWDT) hardware by
 * doing the following:
 * - Disables the WDT and all of its stages
 * - Sets some registers with default values
 * - Sets the WDTs source clock prescaler (not applicable to RWDT)
 * - Optionally enables the level interrupt
 *
 * The HAL context is initialized by storing the type (i.e. MWDT or RWDT) of
 * this WDT instance, and a pointer to the associated registers.
 *
 * @param hal Context of HAL layer
 * @param wdt_inst Which WDT instance to initialize (MWDT0, MWDT1, or RWDT)
 * @param prescaler MWDT source clock prescaler. Unused for RWDT
 * @param enable_intr True to enable level interrupt. False to disable
 *
 * @note Although the WDTs on the ESP32 have an edge interrupt, this HAL does
 *       not utilize it and will always disables it.
 * @note RWDT does not have a prescaler. Its tick rate is equal to the
 *       frequency of its source clock (RTC slow clock).
 */
// llgo:link (*WdtHalContextT).WdtHalInit C.wdt_hal_init
func (recv_ *WdtHalContextT) WdtHalInit(wdt_inst WdtInstT, prescaler c.Uint32T, enable_intr bool) {
}

/**
 * @brief Deinitialize a WDT and its HAL context
 *
 * This function deinitializes a WDT by feeding then disabling it. The WDT's
 * interrupt is also cleared and disabled. The HAL context is cleared.
 *
 * @param hal Context of HAL layer
 */
// llgo:link (*WdtHalContextT).WdtHalDeinit C.wdt_hal_deinit
func (recv_ *WdtHalContextT) WdtHalDeinit() {
}

/**
 * @brief Configure a particular stage of a WDT
 *
 * @param hal Context of HAL layer
 * @param stage Stage to configure (0 to 3)
 * @param timeout Number of WDT ticks for the stage to time out
 * @param behavior What action to take when the stage times out. Note that only
 *                 the RWDT supports the RTC reset action.
 *
 * @note This function can only be called when the WDT is unlocked. Call
 *       wdt_hal_write_protect_disable() first.
 */
// llgo:link (*WdtHalContextT).WdtHalConfigStage C.wdt_hal_config_stage
func (recv_ *WdtHalContextT) WdtHalConfigStage(stage WdtStageT, timeout c.Uint32T, behavior WdtStageActionT) {
}

/**
 * @brief Disable write protection of the WDT registers
 *
 * @param hal Context of HAL layer
 */
// llgo:link (*WdtHalContextT).WdtHalWriteProtectDisable C.wdt_hal_write_protect_disable
func (recv_ *WdtHalContextT) WdtHalWriteProtectDisable() {
}

/**
 * @brief Enable write protection of the WDT registers
 *
 * @param hal Context of HAL layer
 */
// llgo:link (*WdtHalContextT).WdtHalWriteProtectEnable C.wdt_hal_write_protect_enable
func (recv_ *WdtHalContextT) WdtHalWriteProtectEnable() {
}

/**
 * @brief Enable the WDT
 *
 * The WDT will start counting when enabled. This function also feeds the WDT
 * before enabling it.
 *
 * @param hal Context of HAL layer
 *
 * @note This function can only be called when the WDT is unlocked. Call
 *       wdt_hal_write_protect_disable() first.
 */
// llgo:link (*WdtHalContextT).WdtHalEnable C.wdt_hal_enable
func (recv_ *WdtHalContextT) WdtHalEnable() {
}

/**
 * @brief Disable the WDT
 *
 * @param hal Context of HAL layer
 *
 * @note This function can only be called when the WDT is unlocked. Call
 *       wdt_hal_write_protect_disable() first.
 */
// llgo:link (*WdtHalContextT).WdtHalDisable C.wdt_hal_disable
func (recv_ *WdtHalContextT) WdtHalDisable() {
}

/**
 * @brief Handle WDT interrupt
 *
 * Clears the interrupt status bit and feeds the WDT
 *
 * @param hal Context of HAL layer
 *
 * @note This function can only be called when the WDT is unlocked. Call
 *       wdt_hal_write_protect_disable() first.
 */
// llgo:link (*WdtHalContextT).WdtHalHandleIntr C.wdt_hal_handle_intr
func (recv_ *WdtHalContextT) WdtHalHandleIntr() {
}

/**
 * @brief Feed the WDT
 *
 * Feeding the WDT will reset the internal count and current stage.
 *
 * @param hal Context of HAL layer
 *
 * @note This function can only be called when the WDT is unlocked. Call
 *       wdt_hal_write_protect_disable() first.
 */
// llgo:link (*WdtHalContextT).WdtHalFeed C.wdt_hal_feed
func (recv_ *WdtHalContextT) WdtHalFeed() {
}

/**
 * @brief Enable/Disable the WDT flash boot mode
 *
 * @param hal Context of HAL layer
 * @param enable True to enable flash boot mode, false to disable.
 *
 * @note Flash boot mode can trigger a time out even if the WDT is disabled.
 * @note This function can only be called when the WDT is unlocked. Call
 *       wdt_hal_write_protect_disable() first.
 */
// llgo:link (*WdtHalContextT).WdtHalSetFlashbootEn C.wdt_hal_set_flashboot_en
func (recv_ *WdtHalContextT) WdtHalSetFlashbootEn(enable bool) {
}

/**
 * @brief Check if the WDT is enabled
 *
 * @param hal Context of HAL layer
 * @return True if enabled, false otherwise
 */
// llgo:link (*WdtHalContextT).WdtHalIsEnabled C.wdt_hal_is_enabled
func (recv_ *WdtHalContextT) WdtHalIsEnabled() bool {
	return false
}
