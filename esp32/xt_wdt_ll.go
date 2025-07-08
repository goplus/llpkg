package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Enable the XT_WDT
 *
 * @param hw Start address of the peripheral registers.
 */
// llgo:link (*RtcCntlDevT).XtWdtLlEnable C.xt_wdt_ll_enable
func (recv_ *RtcCntlDevT) XtWdtLlEnable(enable bool) {
}

/**
 * @brief Check if the XT_WDT is enabled
 *
 * @param hw Start address of the peripheral registers.
 * @return True if XT WDT is enabled
 */
// llgo:link (*RtcCntlDevT).XtWdtLlCheckIfEnabled C.xt_wdt_ll_check_if_enabled
func (recv_ *RtcCntlDevT) XtWdtLlCheckIfEnabled() bool {
	return false
}

/**
 * @brief Set the watchdog timeout value
 *
 * @param hw Start address of the peripheral registers.
 * @param timeout timeout value in RTC_CLK cycles
 */
// llgo:link (*RtcCntlDevT).XtWdtLlSetTimeout C.xt_wdt_ll_set_timeout
func (recv_ *RtcCntlDevT) XtWdtLlSetTimeout(timeout c.Uint8T) {
}

/**
 * @brief Reset the XT_WDT
 *
 * @param hw Start address of the peripheral registers.
 */
// llgo:link (*RtcCntlDevT).XtWdtLlReset C.xt_wdt_ll_reset
func (recv_ *RtcCntlDevT) XtWdtLlReset() {
}

/**
 * @brief Set the backup clock value
 *
 * @param hw Start address of the peripheral registers.
 * @param backup_clk_val Backup clock value, see TRM for definition
 */
// llgo:link (*RtcCntlDevT).XtWdtLlSetBackupClkFactor C.xt_wdt_ll_set_backup_clk_factor
func (recv_ *RtcCntlDevT) XtWdtLlSetBackupClkFactor(backup_clk_val c.Uint32T) {
}

/**
 * @brief Enable the auto-backup clock feature
 *
 * @param hw Start address of the peripheral registers.
 * @param enable True - enable, False - disable
 */
// llgo:link (*RtcCntlDevT).XtWdtLlAutoBackupEnable C.xt_wdt_ll_auto_backup_enable
func (recv_ *RtcCntlDevT) XtWdtLlAutoBackupEnable(enable bool) {
}

/**
 * @brief Enable the timeout interrupt
 *
 * @param hw Start address of the peripheral registers.
 * @param enable True - enable, False - disable
 */
// llgo:link (*RtcCntlDevT).XtWdtLlIntrEnable C.xt_wdt_ll_intr_enable
func (recv_ *RtcCntlDevT) XtWdtLlIntrEnable(enable bool) {
}
