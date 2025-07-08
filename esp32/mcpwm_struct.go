package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Prescaler configuration */
/** Type of clk_cfg register
 *  PWM clock prescaler register.
 */

type McpwmClkCfgRegT struct {
	Val c.Uint32T
}

/** Group: MCPWM Timer Configuration and status */
/** Type of timer_cfg0 register
 *  PWM timer period and update method configuration register.
 */

type McpwmTimerCfg0RegT struct {
	Val c.Uint32T
}

/** Type of timer_cfg1 register
 *  PWM timer working mode and start/stop control configuration register.
 */

type McpwmTimerCfg1RegT struct {
	Val c.Uint32T
}

/** Type of timer_sync register
 *  PWM timer sync function configuration register.
 */

type McpwmTimerSyncRegT struct {
	Val c.Uint32T
}

/** Type of timer_status register
 *  PWM timer status register.
 */

type McpwmTimerStatusRegT struct {
	Val c.Uint32T
}

/** Group: Common configuration for MCPWM timers */
/** Type of timer_synci_cfg register
 *  Synchronization input selection for three PWM timers.
 */

type McpwmTimerSynciCfgRegT struct {
	Val c.Uint32T
}

/** Type of operator_timersel register
 *  Select specific timer for PWM operators.
 */

type McpwmOperatorTimerselRegT struct {
	Val c.Uint32T
}

/** Group: MCPWM Operator Configuration and Status */
/** Type of gen_stmp_cfg register
 *  Transfer status and update method for time stamp registers A and B
 */

type McpwmGenStmpCfgRegT struct {
	Val c.Uint32T
}

/** Type of gen_tstmp register
 *  PWM generator shadow register for timer stamp
 */

type McpwmGenTstmpRegT struct {
	Val c.Uint32T
}

/** Type of gen_cfg0 register
 *  PWM generator event T0 and T1 handling
 */

type McpwmGenCfg0RegT struct {
	Val c.Uint32T
}

/** Type of gen_force register
 *  Permissives to force PWM0A and PWM0B outputs by software
 */

type McpwmGenForceRegT struct {
	Val c.Uint32T
}

/** Type of generator register
 *  Actions triggered by events on PWM0A
 */

type McpwmGenRegT struct {
	Val c.Uint32T
}

/** Type of dt_cfg register
 *  PWM generator dead time type selection and configuration
 */

type McpwmDtCfgRegT struct {
	Val c.Uint32T
}

/** Type of dt_fed_cfg register
 *  PWM generator shadow register for falling edge delay (FED).
 */

type McpwmDtFedCfgRegT struct {
	Val c.Uint32T
}

/** Type of dt_red_cfg register
 *  PWM generator shadow register for rising edge delay (RED).
 */

type McpwmDtRedCfgRegT struct {
	Val c.Uint32T
}

/** Type of carrier_cfg register
 *  PWM generator carrier enable and configuration
 */

type McpwmCarrierCfgRegT struct {
	Val c.Uint32T
}

/** Type of fh_cfg0 register
 *  Actions on PWM0A and PWM0B trip events
 */

type McpwmFhCfg0RegT struct {
	Val c.Uint32T
}

/** Type of fh_cfg1 register
 *  Software triggers for fault handler actions
 */

type McpwmFhCfg1RegT struct {
	Val c.Uint32T
}

/** Type of fh_status register
 *  Status of fault events.
 */

type McpwmFhStatusRegT struct {
	Val c.Uint32T
}

/** Group: Fault Detection Configuration and Status */
/** Type of fault_detect register
 *  Fault detection configuration and status
 */

type McpwmFaultDetectRegT struct {
	Val c.Uint32T
}

/** Group: Capture Configuration and Status */
/** Type of cap_timer_cfg register
 *  Configure capture timer
 */

type McpwmCapTimerCfgRegT struct {
	Val c.Uint32T
}

/** Type of cap_timer_phase register
 *  Phase for capture timer sync
 */

type McpwmCapTimerPhaseRegT struct {
	Val c.Uint32T
}

/** Type of cap_chn_cfg register
 *  Capture channel 0 configuration and enable
 */

type McpwmCapChnCfgRegT struct {
	Val c.Uint32T
}

/** Type of cap_chn register
 *  chn capture value status register
 */

type McpwmCapChnRegT struct {
	Val c.Uint32T
}

/** Type of cap_status register
 *  Edge of last capture trigger
 */

type McpwmCapStatusRegT struct {
	Val c.Uint32T
}

/** Group: Enable update of active registers */
/** Type of update_cfg register
 *  Enable update.
 */

type McpwmUpdateCfgRegT struct {
	Val c.Uint32T
}

/** Group: Manage Interrupts */
/** Type of int_ena register
 *  Interrupt enable bits
 */

type McpwmIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_raw register
 *  Raw interrupt status
 */

type McpwmIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Masked interrupt status
 */

type McpwmIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear bits
 */

type McpwmIntClrRegT struct {
	Val c.Uint32T
}

/** Group: MCMCPWM APB configuration register */
/** Type of clk register
 *  MCPWM APB configuration register
 */

type McpwmClkRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of version register
 *  Version register.
 */

type McpwmVersionRegT struct {
	Val c.Uint32T
}

type McpwmTimerRegsT struct {
	TimerCfg0   McpwmTimerCfg0RegT
	TimerCfg1   McpwmTimerCfg1RegT
	TimerSync   McpwmTimerSyncRegT
	TimerStatus McpwmTimerStatusRegT
}

type McpwmOperatorRegT struct {
	GenStmpCfg McpwmGenStmpCfgRegT
	Timestamp  [2]McpwmGenTstmpRegT
	GenCfg0    McpwmGenCfg0RegT
	GenForce   McpwmGenForceRegT
	Generator  [2]McpwmGenRegT
	DtCfg      McpwmDtCfgRegT
	DtFedCfg   McpwmDtFedCfgRegT
	DtRedCfg   McpwmDtRedCfgRegT
	CarrierCfg McpwmCarrierCfgRegT
	FhCfg0     McpwmFhCfg0RegT
	FhCfg1     McpwmFhCfg1RegT
	FhStatus   McpwmFhStatusRegT
}

type McpwmDevT struct {
	ClkCfg           McpwmClkCfgRegT
	Timer            [3]McpwmTimerRegsT
	TimerSynciCfg    McpwmTimerSynciCfgRegT
	OperatorTimersel McpwmOperatorTimerselRegT
	Operators        [3]McpwmOperatorRegT
	FaultDetect      McpwmFaultDetectRegT
	CapTimerCfg      McpwmCapTimerCfgRegT
	CapTimerPhase    McpwmCapTimerPhaseRegT
	CapChnCfg        [3]McpwmCapChnCfgRegT
	CapChn           [3]McpwmCapChnRegT
	CapStatus        McpwmCapStatusRegT
	UpdateCfg        McpwmUpdateCfgRegT
	IntEna           McpwmIntEnaRegT
	IntRaw           McpwmIntRawRegT
	IntSt            McpwmIntStRegT
	IntClr           McpwmIntClrRegT
	Clk              McpwmClkRegT
	Version          McpwmVersionRegT
}
