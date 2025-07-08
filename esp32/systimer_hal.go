package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SystimerSocHandleT *SystimerDevT

// llgo:type C
type TicksToUsFuncT func(c.Uint64T) c.Uint64T

// llgo:type C
type UsToTicksFuncT func(c.Uint64T) c.Uint64T

/**
 * @brief systimer HAL context structure
 */

type SystimerHalContextT struct {
	Dev       SystimerSocHandleT
	TicksToUs TicksToUsFuncT
	UsToTicks UsToTicksFuncT
}

/**
 * @brief systimer HAL configuration structure
 */

type SystimerHalTickRateOpsT struct {
	TicksToUs TicksToUsFuncT
	UsToTicks UsToTicksFuncT
}
type SystimerClockSourceT SocPeriphSystimerClkSrcT

/**
 * @brief initialize systimer in HAL layer
 */
// llgo:link (*SystimerHalContextT).SystimerHalInit C.systimer_hal_init
func (recv_ *SystimerHalContextT) SystimerHalInit() {
}

/**
 * @brief Deinitialize systimer HAL layer
 */
// llgo:link (*SystimerHalContextT).SystimerHalDeinit C.systimer_hal_deinit
func (recv_ *SystimerHalContextT) SystimerHalDeinit() {
}

/**
 * @brief Set tick rate operation functions
 */
// llgo:link (*SystimerHalContextT).SystimerHalSetTickRateOps C.systimer_hal_set_tick_rate_ops
func (recv_ *SystimerHalContextT) SystimerHalSetTickRateOps(ops *SystimerHalTickRateOpsT) {
}

/**
 * @brief enable systimer counter
 */
// llgo:link (*SystimerHalContextT).SystimerHalEnableCounter C.systimer_hal_enable_counter
func (recv_ *SystimerHalContextT) SystimerHalEnableCounter(counter_id c.Uint32T) {
}

/**
 * @brief get current counter value
 */
// llgo:link (*SystimerHalContextT).SystimerHalGetCounterValue C.systimer_hal_get_counter_value
func (recv_ *SystimerHalContextT) SystimerHalGetCounterValue(counter_id c.Uint32T) c.Uint64T {
	return 0
}

/**
 * @brief get current time (in microseconds)
 */
// llgo:link (*SystimerHalContextT).SystimerHalGetTime C.systimer_hal_get_time
func (recv_ *SystimerHalContextT) SystimerHalGetTime(counter_id c.Uint32T) c.Uint64T {
	return 0
}

/*
 * @brief set alarm target value (used in one-shot mode)
 */
// llgo:link (*SystimerHalContextT).SystimerHalSetAlarmTarget C.systimer_hal_set_alarm_target
func (recv_ *SystimerHalContextT) SystimerHalSetAlarmTarget(alarm_id c.Uint32T, target c.Uint64T) {
}

/**
 * @brief set alarm period value (used in period mode)
 */
// llgo:link (*SystimerHalContextT).SystimerHalSetAlarmPeriod C.systimer_hal_set_alarm_period
func (recv_ *SystimerHalContextT) SystimerHalSetAlarmPeriod(alarm_id c.Uint32T, period c.Uint32T) {
}

/**
 * @brief get alarm time
 */
// llgo:link (*SystimerHalContextT).SystimerHalGetAlarmValue C.systimer_hal_get_alarm_value
func (recv_ *SystimerHalContextT) SystimerHalGetAlarmValue(alarm_id c.Uint32T) c.Uint64T {
	return 0
}

/**
 * @brief enable alarm interrupt
 */
// llgo:link (*SystimerHalContextT).SystimerHalEnableAlarmInt C.systimer_hal_enable_alarm_int
func (recv_ *SystimerHalContextT) SystimerHalEnableAlarmInt(alarm_id c.Uint32T) {
}

/**
 * @brief select alarm mode
 */
// llgo:link (*SystimerHalContextT).SystimerHalSelectAlarmMode C.systimer_hal_select_alarm_mode
func (recv_ *SystimerHalContextT) SystimerHalSelectAlarmMode(alarm_id c.Uint32T, mode SystimerAlarmModeT) {
}

/**
 * @brief update systimer step when apb clock gets changed
 */
// llgo:link (*SystimerHalContextT).SystimerHalOnApbFreqUpdate C.systimer_hal_on_apb_freq_update
func (recv_ *SystimerHalContextT) SystimerHalOnApbFreqUpdate(apb_ticks_per_us c.Uint32T) {
}

/**
 * @brief move systimer counter value forward or backward
 */
// llgo:link (*SystimerHalContextT).SystimerHalCounterValueAdvance C.systimer_hal_counter_value_advance
func (recv_ *SystimerHalContextT) SystimerHalCounterValueAdvance(counter_id c.Uint32T, time_us c.Int64T) {
}

/**
 * @brief  connect alarm unit to selected counter
 */
// llgo:link (*SystimerHalContextT).SystimerHalConnectAlarmCounter C.systimer_hal_connect_alarm_counter
func (recv_ *SystimerHalContextT) SystimerHalConnectAlarmCounter(alarm_id c.Uint32T, counter_id c.Uint32T) {
}

/**
 * @brief  set if a counter should be stalled when CPU is halted by the debugger
 */
// llgo:link (*SystimerHalContextT).SystimerHalCounterCanStallByCpu C.systimer_hal_counter_can_stall_by_cpu
func (recv_ *SystimerHalContextT) SystimerHalCounterCanStallByCpu(counter_id c.Uint32T, cpu_id c.Uint32T, can bool) {
}

/**
 * @brief Get Systimer clock source
 */
// llgo:link (*SystimerHalContextT).SystimerHalGetClockSource C.systimer_hal_get_clock_source
func (recv_ *SystimerHalContextT) SystimerHalGetClockSource() SystimerClockSourceT {
	return 0
}
