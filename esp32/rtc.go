package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const OTHER_BLOCKS_POWERUP = 1
const OTHER_BLOCKS_WAIT = 1
const RTC_CNTL_DBIAS_SLP = 5
const RTC_CNTL_DBIAS_0V90 = 13
const RTC_CNTL_DBIAS_0V95 = 16
const RTC_CNTL_DBIAS_1V00 = 18
const RTC_CNTL_DBIAS_1V05 = 20
const RTC_CNTL_DBIAS_1V10 = 23
const RTC_CNTL_DBIAS_1V15 = 25
const RTC_CNTL_DBIAS_1V20 = 28
const RTC_CNTL_DBIAS_1V25 = 30
const RTC_CNTL_DBIAS_1V30 = 31
const SOC_DELAY_RTC_FAST_CLK_SWITCH = 3
const SOC_DELAY_RTC_SLOW_CLK_SWITCH = 300
const SOC_DELAY_RC_FAST_ENABLE = 50
const SOC_DELAY_RC_FAST_DIGI_SWITCH = 5
const RTC_CNTL_PLL_BUF_WAIT_DEFAULT = 20
const RTC_CNTL_XTL_BUF_WAIT_DEFAULT = 100
const RTC_CNTL_CK8M_WAIT_DEFAULT = 20
const RTC_CK8M_ENABLE_WAIT_DEFAULT = 5
const RTC_CNTL_CK8M_DFREQ_DEFAULT = 100
const RTC_CNTL_SCK_DCAP_DEFAULT = 255
const RTC_CNTL_DBG_ATTEN_LIGHTSLEEP_DEFAULT = 5
const RTC_CNTL_DBG_ATTEN_LIGHTSLEEP_NODROP = 0
const RTC_CNTL_DBG_ATTEN_DEEPSLEEP_DEFAULT = 14
const RTC_CNTL_DBG_ATTEN_DEEPSLEEP_ULTRA_LOW = 15
const RTC_CNTL_DBG_ATTEN_DEEPSLEEP_NODROP = 0
const RTC_CNTL_BIASSLP_SLEEP_DEFAULT = 1
const RTC_CNTL_BIASSLP_SLEEP_ON = 0
const RTC_CNTL_PD_CUR_SLEEP_DEFAULT = 1
const RTC_CNTL_PD_CUR_SLEEP_ON = 0
const RTC_CNTL_DG_VDD_DRV_B_SLP_DEFAULT = 0xf
const RTC_CNTL_DBG_ATTEN_MONITOR_DEFAULT = 0
const RTC_CNTL_BIASSLP_MONITOR_DEFAULT = 1
const RTC_CNTL_BIASSLP_MONITOR_ON = 0
const RTC_CNTL_PD_CUR_MONITOR_DEFAULT = 1
const RTC_CNTL_PD_CUR_MONITOR_ON = 0
const K_RTC_MID_MUL10000 = 198
const K_DIG_MID_MUL10000 = 211
const V_RTC_MID_MUL10000 = 10181
const V_DIG_MID_MUL10000 = 10841
const DEFAULT_LDO_SLAVE = 0x7
const RTC_CLK_CAL_FRACT = 19
const RTC_VDDSDIO_TIEH_1_8V = 0
const RTC_VDDSDIO_TIEH_3_3V = 1

/**
 * @brief CPU clock configuration structure
 */

type RtcCpuFreqConfigS struct {
	Source        SocCpuClkSrcT
	SourceFreqMhz c.Uint32T
	Div           c.Uint32T
	FreqMhz       c.Uint32T
}
type RtcCpuFreqConfigT RtcCpuFreqConfigS
type RtcCalSelT c.Int

const (
	RTC_CAL_RTC_MUX      RtcCalSelT = 0
	RTC_CAL_8MD256       RtcCalSelT = 1
	RTC_CAL_32K_XTAL     RtcCalSelT = 2
	RTC_CAL_INTERNAL_OSC RtcCalSelT = 3
)

/**
 * Initialization parameters for rtc_clk_init
 */

type RtcClkConfigT struct {
	XtalFreq     SocXtalFreqT
	CpuFreqMhz   c.Uint32T
	FastClkSrc   SocRtcFastClkSrcT
	SlowClkSrc   SocRtcSlowClkSrcT
	ClkRtcClkDiv c.Uint32T
	Clk8mClkDiv  c.Uint32T
	SlowClkDcap  c.Uint32T
	Clk8mDfreq   c.Uint32T
}

type RtcInitConfigT struct {
	WifiPowerupCycles   c.Uint16T
	WifiWaitCycles      c.Uint16T
	BtPowerupCycles     c.Uint16T
	BtWaitCycles        c.Uint16T
	RtcPowerupCycles    c.Uint16T
	RtcWaitCycles       c.Uint16T
	CpuTopPowerupCycles c.Uint16T
	CpuTopWaitCycles    c.Uint16T
	DgWrapPowerupCycles c.Uint16T
	DgWrapWaitCycles    c.Uint16T
	DgPeriPowerupCycles c.Uint16T
	DgPeriWaitCycles    c.Uint16T
	RtcMemPowerupCycles c.Uint16T
	RtcMemWaitCycles    c.Uint16T
}

//go:linkname RtcClkDividerSet C.rtc_clk_divider_set
func RtcClkDividerSet(div c.Uint32T)

//go:linkname RtcClk8mDividerSet C.rtc_clk_8m_divider_set
func RtcClk8mDividerSet(div c.Uint32T)

/**
 * Initialize clocks and set CPU frequency
 *
 * @param cfg clock configuration as rtc_clk_config_t
 */
// llgo:link RtcClkConfigT.RtcClkInit C.rtc_clk_init
func (recv_ RtcClkConfigT) RtcClkInit() {
}

/**
 * @brief Get main XTAL frequency
 *
 * This is the value stored in RTC register RTC_XTAL_FREQ_REG by the bootloader. As passed to
 * rtc_clk_init function
 *
 * @return XTAL frequency, one of soc_xtal_freq_t
 */
//go:linkname RtcClkXtalFreqGet C.rtc_clk_xtal_freq_get
func RtcClkXtalFreqGet() SocXtalFreqT

/**
 * @brief Update XTAL frequency
 *
 * Updates the XTAL value stored in RTC_XTAL_FREQ_REG. Usually this value is ignored
 * after startup.
 *
 * @param xtal_freq New frequency value
 */
// llgo:link SocXtalFreqT.RtcClkXtalFreqUpdate C.rtc_clk_xtal_freq_update
func (recv_ SocXtalFreqT) RtcClkXtalFreqUpdate() {
}

/**
 * @brief Enable or disable 32 kHz XTAL oscillator
 * @param en  true to enable, false to disable
 */
//go:linkname RtcClk32kEnable C.rtc_clk_32k_enable
func RtcClk32kEnable(en bool)

/**
 * @brief Configure 32 kHz XTAL oscillator to accept external clock signal
 */
//go:linkname RtcClk32kEnableExternal C.rtc_clk_32k_enable_external
func RtcClk32kEnableExternal()

/**
 * @brief Disable 32 kHz XTAL oscillator input.
 */
//go:linkname RtcClk32kDisableExternal C.rtc_clk_32k_disable_external
func RtcClk32kDisableExternal()

/**
 * @brief Get the state of 32k XTAL oscillators
 * @return true if 32k XTAL oscillator has been enabled
 */
//go:linkname RtcClk32kEnabled C.rtc_clk_32k_enabled
func RtcClk32kEnabled() bool

/**
 * @brief Enable 32k oscillator, configuring it for fast startup time.
 * Note: to achieve higher frequency stability, rtc_clk_32k_enable function
 * must be called one the 32k XTAL oscillator has started up. This function
 * will initially disable the 32k XTAL oscillator, so it should not be called
 * when the system is using 32k XTAL as RTC_SLOW_CLK.
 *
 * @param cycle Number of 32kHz cycles to bootstrap external crystal.
 *              If 0, no square wave will be used to bootstrap crystal oscillation.
 */
//go:linkname RtcClk32kBootstrap C.rtc_clk_32k_bootstrap
func RtcClk32kBootstrap(cycle c.Uint32T)

/**
 * @brief Enable or disable 8 MHz internal oscillator
 *
 * Output from 8 MHz internal oscillator is passed into a configurable
 * divider, which by default divides the input clock frequency by 256.
 * Output of the divider may be used as RTC_SLOW_CLK source.
 * Output of the divider is referred to in register descriptions and code as
 * 8md256 or simply d256. Divider values other than 256 may be configured, but
 * this facility is not currently needed, so is not exposed in the code.
 *
 * When 8MHz/256 divided output is not needed, the divider should be disabled
 * to reduce power consumption.
 *
 * @param clk_8m_en true to enable 8MHz generator
 * @param d256_en true to enable /256 divider
 */
//go:linkname RtcClk8mEnable C.rtc_clk_8m_enable
func RtcClk8mEnable(clk_8m_en bool, d256_en bool)

/**
 * @brief Get the state of 8 MHz internal oscillator
 * @return true if the oscillator is enabled
 */
//go:linkname RtcClk8mEnabled C.rtc_clk_8m_enabled
func RtcClk8mEnabled() bool

/**
 * @brief Get the state of /256 divider which is applied to 8MHz clock
 * @return true if the divided output is enabled
 */
//go:linkname RtcClk8md256Enabled C.rtc_clk_8md256_enabled
func RtcClk8md256Enabled() bool

/**
 * @brief Select source for RTC_SLOW_CLK
 * @param clk_src clock source (one of soc_rtc_slow_clk_src_t values)
 */
// llgo:link SocRtcSlowClkSrcT.RtcClkSlowSrcSet C.rtc_clk_slow_src_set
func (recv_ SocRtcSlowClkSrcT) RtcClkSlowSrcSet() {
}

/**
 * @brief Get the RTC_SLOW_CLK source
 * @return currently selected clock source (one of soc_rtc_slow_clk_src_t values)
 */
//go:linkname RtcClkSlowSrcGet C.rtc_clk_slow_src_get
func RtcClkSlowSrcGet() SocRtcSlowClkSrcT

/**
 * @brief Get the approximate frequency of RTC_SLOW_CLK, in Hz
 *
 * - if SOC_RTC_SLOW_CLK_SRC_RC_SLOW is selected, returns ~150000
 * - if SOC_RTC_SLOW_CLK_SRC_XTAL32K is selected, returns 32768
 * - if SOC_RTC_SLOW_CLK_SRC_RC_FAST_D256 is selected, returns ~68000
 *
 * rtc_clk_cal function can be used to get more precise value by comparing
 * RTC_SLOW_CLK frequency to the frequency of main XTAL.
 *
 * @return RTC_SLOW_CLK frequency, in Hz
 */
//go:linkname RtcClkSlowFreqGetHz C.rtc_clk_slow_freq_get_hz
func RtcClkSlowFreqGetHz() c.Uint32T

/**
 * @brief Select source for RTC_FAST_CLK
 * @param clk_src clock source (one of soc_rtc_fast_clk_src_t values)
 */
// llgo:link SocRtcFastClkSrcT.RtcClkFastSrcSet C.rtc_clk_fast_src_set
func (recv_ SocRtcFastClkSrcT) RtcClkFastSrcSet() {
}

/**
 * @brief Get the RTC_FAST_CLK source
 * @return currently selected clock source (one of soc_rtc_fast_clk_src_t values)
 */
//go:linkname RtcClkFastSrcGet C.rtc_clk_fast_src_get
func RtcClkFastSrcGet() SocRtcFastClkSrcT

/**
 * @brief Get CPU frequency config for a given frequency
 * @param freq_mhz  Frequency in MHz
 * @param[out] out_config Output, CPU frequency configuration structure
 * @return true if frequency can be obtained, false otherwise
 */
//go:linkname RtcClkCpuFreqMhzToConfig C.rtc_clk_cpu_freq_mhz_to_config
func RtcClkCpuFreqMhzToConfig(freq_mhz c.Uint32T, out_config *RtcCpuFreqConfigT) bool

/**
 * @brief Switch CPU frequency
 *
 * This function sets CPU frequency according to the given configuration
 * structure. It enables PLLs, if necessary.
 *
 * @note This function in not intended to be called by applications in FreeRTOS
 * environment. This is because it does not adjust various timers based on the
 * new CPU frequency.
 *
 * @param config  CPU frequency configuration structure
 */
// llgo:link (*RtcCpuFreqConfigT).RtcClkCpuFreqSetConfig C.rtc_clk_cpu_freq_set_config
func (recv_ *RtcCpuFreqConfigT) RtcClkCpuFreqSetConfig() {
}

/**
 * @brief Switch CPU frequency (optimized for speed)
 *
 * This function is a faster equivalent of rtc_clk_cpu_freq_set_config.
 * It works faster because it does not disable PLLs when switching from PLL to
 * XTAL and does not enabled them when switching back. If PLL is not already
 * enabled when this function is called to switch from XTAL to PLL frequency,
 * or the PLL which is enabled is the wrong one, this function will fall back
 * to calling rtc_clk_cpu_freq_set_config.
 *
 * Unlike rtc_clk_cpu_freq_set_config, this function relies on static data,
 * so it is less safe to use it e.g. from a panic handler (when memory might
 * be corrupted).
 *
 * @note This function in not intended to be called by applications in FreeRTOS
 * environment. This is because it does not adjust various timers based on the
 * new CPU frequency.
 *
 * @param config  CPU frequency configuration structure
 */
// llgo:link (*RtcCpuFreqConfigT).RtcClkCpuFreqSetConfigFast C.rtc_clk_cpu_freq_set_config_fast
func (recv_ *RtcCpuFreqConfigT) RtcClkCpuFreqSetConfigFast() {
}

/**
 * @brief Get the currently used CPU frequency configuration
 * @param[out] out_config  Output, CPU frequency configuration structure
 */
// llgo:link (*RtcCpuFreqConfigT).RtcClkCpuFreqGetConfig C.rtc_clk_cpu_freq_get_config
func (recv_ *RtcCpuFreqConfigT) RtcClkCpuFreqGetConfig() {
}

/**
 * @brief Switch CPU clock source to XTAL
 *
 * Short form for filling in rtc_cpu_freq_config_t structure and calling
 * rtc_clk_cpu_freq_set_config when a switch to XTAL is needed.
 * Assumes that XTAL frequency has been determined — don't call in startup code.
 *
 * @note This function always disables BBPLL after switching the CPU clock source to XTAL for power saving purpose.
 * If this is unwanted, please use rtc_clk_cpu_freq_set_config. It helps to check whether USB Serial JTAG is in use,
 * if so, then BBPLL will not be turned off.
 */
//go:linkname RtcClkCpuFreqSetXtal C.rtc_clk_cpu_freq_set_xtal
func RtcClkCpuFreqSetXtal()

/**
 * @brief Store new APB frequency value in RAM
 *
 * This function doesn't change any hardware clocks.
 *
 * Functions which perform frequency switching and change APB frequency call
 * this function to update the value of APB frequency stored in RAM.
 * (This should not normally be called from application code.)
 *
 * @param apb_freq  new APB frequency, in Hz
 */
//go:linkname RtcClkApbFreqUpdate C.rtc_clk_apb_freq_update
func RtcClkApbFreqUpdate(apb_freq c.Uint32T)

/**
 * @brief Get the current stored APB frequency.
 * @return The APB frequency value as last set via rtc_clk_apb_freq_update(), in Hz.
 */
//go:linkname RtcClkApbFreqGet C.rtc_clk_apb_freq_get
func RtcClkApbFreqGet() c.Uint32T

// llgo:link RtcCalSelT.RtcClkCalInternal C.rtc_clk_cal_internal
func (recv_ RtcCalSelT) RtcClkCalInternal(slowclk_cycles c.Uint32T) c.Uint32T {
	return 0
}

/**
 * @brief Measure RTC slow clock's period, based on main XTAL frequency
 *
 * This function will time out and return 0 if the time for the given number
 * of cycles to be counted exceeds the expected time twice. This may happen if
 * 32k XTAL is being calibrated, but the oscillator has not started up (due to
 * incorrect loading capacitance, board design issue, or lack of 32 XTAL on board).
 *
 * @note When 32k CLK is being calibrated, this function will check the accuracy
 * of the clock. Since the xtal 32k or ext osc 32k is generally very stable, if
 * the check fails, then consider this an invalid 32k clock and return 0. This
 * check can filter some jamming signal.
 *
 * @param cal_clk  clock to be measured
 * @param slow_clk_cycles  number of slow clock cycles to average
 * @return average slow clock period in microseconds, Q13.19 fixed point format,
 *         or 0 if calibration has timed out
 */
// llgo:link RtcCalSelT.RtcClkCal C.rtc_clk_cal
func (recv_ RtcCalSelT) RtcClkCal(slow_clk_cycles c.Uint32T) c.Uint32T {
	return 0
}

/**
 * @brief Measure ratio between XTAL frequency and RTC slow clock frequency
 * @param cal_clk slow clock to be measured
 * @param slow_clk_cycles number of slow clock cycles to average
 * @return average ratio between XTAL frequency and slow clock frequency,
 *         Q13.19 fixed point format, or 0 if calibration has timed out.
 */
// llgo:link RtcCalSelT.RtcClkCalRatio C.rtc_clk_cal_ratio
func (recv_ RtcCalSelT) RtcClkCalRatio(slow_clk_cycles c.Uint32T) c.Uint32T {
	return 0
}

/**
 * @brief Convert time interval from microseconds to RTC_SLOW_CLK cycles
 * @param time_in_us Time interval in microseconds
 * @param slow_clk_period  Period of slow clock in microseconds, Q13.19
 *                         fixed point format (as returned by rtc_slowck_cali).
 * @return number of slow clock cycles
 */
//go:linkname RtcTimeUsToSlowclk C.rtc_time_us_to_slowclk
func RtcTimeUsToSlowclk(time_in_us c.Uint64T, period c.Uint32T) c.Uint64T

/**
 * @brief Convert time interval from RTC_SLOW_CLK to microseconds
 * @param time_in_us Time interval in RTC_SLOW_CLK cycles
 * @param slow_clk_period  Period of slow clock in microseconds, Q13.19
 *                         fixed point format (as returned by rtc_slowck_cali).
 * @return time interval in microseconds
 */
//go:linkname RtcTimeSlowclkToUs C.rtc_time_slowclk_to_us
func RtcTimeSlowclkToUs(rtc_cycles c.Uint64T, period c.Uint32T) c.Uint64T

/**
 * @brief Get current value of RTC counter
 *
 * RTC has a 48-bit counter which is incremented by 2 every 2 RTC_SLOW_CLK
 * cycles. Counter value is not writable by software. The value is not adjusted
 * when switching to a different RTC_SLOW_CLK source.
 *
 * Note: this function may take up to 1 RTC_SLOW_CLK cycle to execute
 *
 * @return current value of RTC counter
 */
//go:linkname RtcTimeGet C.rtc_time_get
func RtcTimeGet() c.Uint64T

/**
 * @brief Busy loop until next RTC_SLOW_CLK cycle
 *
 * This function returns not earlier than the next RTC_SLOW_CLK clock cycle.
 * In some cases (e.g. when RTC_SLOW_CLK cycle is very close), it may return
 * one RTC_SLOW_CLK cycle later.
 */
//go:linkname RtcClkWaitForSlowCycle C.rtc_clk_wait_for_slow_cycle
func RtcClkWaitForSlowCycle()

/**
 * @brief Enable the rtc digital 8M clock
 *
 * This function is used to enable the digital rtc 8M clock to support peripherals.
 * For enabling the analog 8M clock, using `rtc_clk_8M_enable` function above.
 */
//go:linkname RtcDigClk8mEnable C.rtc_dig_clk8m_enable
func RtcDigClk8mEnable()

/**
 * @brief Disable the rtc digital 8M clock
 *
 * This function is used to disable the digital rtc 8M clock, which is only used to support peripherals.
 */
//go:linkname RtcDigClk8mDisable C.rtc_dig_clk8m_disable
func RtcDigClk8mDisable()

/**
 * @brief Get whether the rtc digital 8M clock is enabled
 */
//go:linkname RtcDig8mEnabled C.rtc_dig_8m_enabled
func RtcDig8mEnabled() bool

/**
 * @brief Calculate the real clock value after the clock calibration
 *
 * @param cal_val Average slow clock period in microseconds, fixed point value as returned from `rtc_clk_cal`
 * @return Frequency of the clock in Hz
 */
//go:linkname RtcClkFreqCal C.rtc_clk_freq_cal
func RtcClkFreqCal(cal_val c.Uint32T) c.Uint32T

/**
 * @brief Calculate the slow clock period value by slow clock frequency
 *
 * @param freq_hz Frequency of the slow clock in Hz
 * @return Fixed point value of slow clock period in microseconds
 */
//go:linkname RtcClkFreqToPeriod C.rtc_clk_freq_to_period
func RtcClkFreqToPeriod(freq_hz c.Uint32T) c.Uint32T

/**
 * @brief Power up flags for rtc_sleep_pd function
 */

type RtcSleepPuConfigT struct {
	DigFpu    c.Uint32T
	RtcFpu    c.Uint32T
	CpuFpu    c.Uint32T
	I2sFpu    c.Uint32T
	BbFpu     c.Uint32T
	NrxFpu    c.Uint32T
	FeFpu     c.Uint32T
	SramFpu   c.Uint32T
	RomRamFpu c.Uint32T
}

// llgo:link RtcSleepPuConfigT.RtcSleepPu C.rtc_sleep_pu
func (recv_ RtcSleepPuConfigT) RtcSleepPu() {
}

/**
 * @brief sleep configuration for rtc_sleep_init function
 */

type RtcSleepConfigT struct {
	LslpMemInfFpu      c.Uint32T
	RtcMemInfFollowCpu c.Uint32T
	RtcFastmemPdEn     c.Uint32T
	RtcSlowmemPdEn     c.Uint32T
	RtcPeriPdEn        c.Uint32T
	ModemPdEn          c.Uint32T
	CpuPdEn            c.Uint32T
	Int8mPdEn          c.Uint32T
	DigPeriPdEn        c.Uint32T
	DeepSlp            c.Uint32T
	WdtFlashbootModEn  c.Uint32T
	DigDbiasSlp        c.Uint32T
	RtcDbiasSlp        c.Uint32T
	BiasSleepMonitor   c.Uint32T
	DbgAttenSlp        c.Uint32T
	BiasSleepSlp       c.Uint32T
	PdCurMonitor       c.Uint32T
	PdCurSlp           c.Uint32T
	VddsdioPdEn        c.Uint32T
	XtalFpu            c.Uint32T
	RtcRegulatorFpu    c.Uint32T
	DeepSlpReject      c.Uint32T
	LightSlpReject     c.Uint32T
}

/**
 * Default initializer for rtc_sleep_config_t
 *
 * This initializer sets all fields to "reasonable" values (e.g. suggested for
 * production use) based on a combination of RTC_SLEEP_PD_x flags.
 *
 * @param RTC_SLEEP_PD_x flags combined using bitwise OR
 */
//go:linkname RtcSleepGetDefaultConfig C.rtc_sleep_get_default_config
func RtcSleepGetDefaultConfig(sleep_flags c.Uint32T, out_config *RtcSleepConfigT)

/**
 * @brief Prepare the chip to enter sleep mode
 *
 * This function configures various power control state machines to handle
 * entry into light sleep or deep sleep mode, switches APB and CPU clock source
 * (usually to XTAL), and sets bias voltages for digital and RTC power domains.
 *
 * This function does not actually enter sleep mode; this is done using
 * rtc_sleep_start function. Software may do some other actions between
 * rtc_sleep_init and rtc_sleep_start, such as set wakeup timer and configure
 * wakeup sources.
 * @param cfg sleep mode configuration
 */
// llgo:link RtcSleepConfigT.RtcSleepInit C.rtc_sleep_init
func (recv_ RtcSleepConfigT) RtcSleepInit() {
}

/**
 * @brief Low level initialize for rtc state machine waiting cycles after waking up
 *
 * This function configures the cycles chip need to wait for internal 8MHz
 * oscillator and external 40MHz crystal. As we configure fixed time for waiting
 * crystal, we need to pass period to calculate cycles. Now this function only
 * used in lightsleep mode.
 *
 * @param slowclk_period re-calibrated slow clock period
 * @param dslp true if initialize for deepsleep request
 */
//go:linkname RtcSleepLowInit C.rtc_sleep_low_init
func RtcSleepLowInit(slowclk_period c.Uint32T, dslp bool)

/**
 * @brief Enter deep or light sleep mode
 *
 * This function enters the sleep mode previously configured using rtc_sleep_init
 * function. Before entering sleep, software should configure wake up sources
 * appropriately (set up GPIO wakeup registers, timer wakeup registers,
 * and so on).
 *
 * If deep sleep mode was configured using rtc_sleep_init, and sleep is not
 * rejected by hardware (based on reject_opt flags), this function never returns.
 * When the chip wakes up from deep sleep, CPU is reset and execution starts
 * from ROM bootloader.
 *
 * If light sleep mode was configured using rtc_sleep_init, this function
 * returns on wakeup, or if sleep is rejected by hardware.
 *
 * @param wakeup_opt  bit mask wake up reasons to enable (RTC_xxx_TRIG_EN flags
 *                    combined with OR)
 * @param reject_opt  bit mask of sleep reject reasons:
 *                      - RTC_CNTL_GPIO_REJECT_EN
 *                      - RTC_CNTL_SDIO_REJECT_EN
 *                    These flags are used to prevent entering sleep when e.g.
 *                    an external host is communicating via SDIO slave
 * @param lslp_mem_inf_fpu If non-zero then the low power config is restored
 *                         immediately on wake. Recommended for light sleep,
 *                         has no effect if the system goes into deep sleep.
 *
 * @return non-zero if sleep was rejected by hardware
 */
//go:linkname RtcSleepStart C.rtc_sleep_start
func RtcSleepStart(wakeup_opt c.Uint32T, reject_opt c.Uint32T, lslp_mem_inf_fpu c.Uint32T) c.Uint32T

/**
 * @brief Enter deep sleep mode
 *
 * Similar to rtc_sleep_start(), but additionally uses hardware to calculate the CRC value
 * of RTC FAST memory. On wake, this CRC is used to determine if a deep sleep wake
 * stub is valid to execute (if a wake address is set).
 *
 * No RAM is accessed while calculating the CRC and going into deep sleep, which makes
 * this function safe to use even if the caller's stack is in RTC FAST memory.
 *
 * @note If no deep sleep wake stub address is set then calling rtc_sleep_start() will
 * have the same effect and takes less time as CRC calculation is skipped.
 *
 * @note This function should only be called after rtc_sleep_init() has been called to
 * configure the system for deep sleep.
 *
 * @param wakeup_opt - same as for rtc_sleep_start
 * @param reject_opt - same as for rtc_sleep_start
 *
 * @return non-zero if sleep was rejected by hardware
 */
//go:linkname RtcDeepSleepStart C.rtc_deep_sleep_start
func RtcDeepSleepStart(wakeup_opt c.Uint32T, reject_opt c.Uint32T) c.Uint32T

/**
 * RTC power and clock control initialization settings
 */

type RtcConfigT struct {
	Ck8mWait        c.Uint32T
	XtalWait        c.Uint32T
	PllWait         c.Uint32T
	ClkctlInit      c.Uint32T
	PwrctlInit      c.Uint32T
	RtcDboostFpd    c.Uint32T
	XtalFpu         c.Uint32T
	BbpllFpu        c.Uint32T
	CpuWaitiClkGate c.Uint32T
	CaliOcode       c.Uint32T
}

/**
 * Initialize RTC clock and power control related functions
 * @param cfg configuration options as rtc_config_t
 */
// llgo:link RtcConfigT.RtcInit C.rtc_init
func (recv_ RtcConfigT) RtcInit() {
}

/**
 * Structure describing vddsdio configuration
 */

type RtcVddsdioConfigT struct {
	Force  c.Uint32T
	Enable c.Uint32T
	Tieh   c.Uint32T
	Drefh  c.Uint32T
	Drefm  c.Uint32T
	Drefl  c.Uint32T
}

/**
 * Get current VDDSDIO configuration
 * If VDDSDIO configuration is overridden by RTC, get values from RTC
 * Otherwise, if VDDSDIO is configured by EFUSE, get values from EFUSE
 * Otherwise, use default values and the level of MTDI bootstrapping pin.
 * @return currently used VDDSDIO configuration
 */
//go:linkname RtcVddsdioGetConfig C.rtc_vddsdio_get_config
func RtcVddsdioGetConfig() RtcVddsdioConfigT

/**
 * Set new VDDSDIO configuration using RTC registers.
 * If config.force == 1, this overrides configuration done using bootstrapping
 * pins and EFUSE.
 *
 * @param config new VDDSDIO configuration
 */
// llgo:link RtcVddsdioConfigT.RtcVddsdioSetConfig C.rtc_vddsdio_set_config
func (recv_ RtcVddsdioConfigT) RtcVddsdioSetConfig() {
}

type RtcCpuFreqSrcT SocCpuClkSrcT
type RtcSlowFreqT SocRtcSlowClkSrcT
type RtcFastFreqT SocRtcFastClkSrcT
type RtcXtalFreqT SocXtalFreqT

/**
 * @brief Get current value of RTC counter in microseconds
 *
 * Note: this function may take up to 1 RTC_SLOW_CLK cycle to execute
 *
 * @return current value of RTC counter in microseconds
 */
//go:linkname EspRtcGetTimeUs C.esp_rtc_get_time_us
func EspRtcGetTimeUs() c.Uint64T

type SLEEPMODE c.Int

const (
	AWAKE       SLEEPMODE = 0
	LIGHT_SLEEP SLEEPMODE = 1
	DEEP_SLEEP  SLEEPMODE = 2
)

type RESETREASON c.Int

const (
	NO_MEAN                RESETREASON = 0
	POWERON_RESET          RESETREASON = 1
	RTC_SW_SYS_RESET       RESETREASON = 3
	DEEPSLEEP_RESET        RESETREASON = 5
	TG0WDT_SYS_RESET       RESETREASON = 7
	TG1WDT_SYS_RESET       RESETREASON = 8
	RTCWDT_SYS_RESET       RESETREASON = 9
	INTRUSION_RESET        RESETREASON = 10
	TG0WDT_CPU_RESET       RESETREASON = 11
	RTC_SW_CPU_RESET       RESETREASON = 12
	RTCWDT_CPU_RESET       RESETREASON = 13
	RTCWDT_BROWN_OUT_RESET RESETREASON = 15
	RTCWDT_RTC_RESET       RESETREASON = 16
	TG1WDT_CPU_RESET       RESETREASON = 17
	SUPER_WDT_RESET        RESETREASON = 18
	GLITCH_RTC_RESET       RESETREASON = 19
	EFUSE_RESET            RESETREASON = 20
	USB_UART_CHIP_RESET    RESETREASON = 21
	USB_JTAG_CHIP_RESET    RESETREASON = 22
	POWER_GLITCH_RESET     RESETREASON = 23
)

type WAKEUPREASON c.Int

const (
	NO_SLEEP        WAKEUPREASON = 0
	EXT_EVENT0_TRIG WAKEUPREASON = 1
	EXT_EVENT1_TRIG WAKEUPREASON = 2
	GPIO_TRIG       WAKEUPREASON = 4
	TIMER_EXPIRE    WAKEUPREASON = 8
	SDIO_TRIG       WAKEUPREASON = 16
	MAC_TRIG        WAKEUPREASON = 32
	UART0_TRIG      WAKEUPREASON = 64
	UART1_TRIG      WAKEUPREASON = 128
	TOUCH_TRIG      WAKEUPREASON = 256
	SAR_TRIG        WAKEUPREASON = 512
	BT_TRIG         WAKEUPREASON = 1024
	RISCV_TRIG      WAKEUPREASON = 2048
	XTAL_DEAD_TRIG  WAKEUPREASON = 4096
	RISCV_TRAP_TRIG WAKEUPREASON = 8192
	USB_TRIG        WAKEUPREASON = 16384
)

type WAKEUPENABLE c.Int

const (
	DISEN_WAKEUP       WAKEUPENABLE = 0
	EXT_EVENT0_TRIG_EN WAKEUPENABLE = 1
	EXT_EVENT1_TRIG_EN WAKEUPENABLE = 2
	GPIO_TRIG_EN       WAKEUPENABLE = 4
	TIMER_EXPIRE_EN    WAKEUPENABLE = 8
	SDIO_TRIG_EN       WAKEUPENABLE = 16
	MAC_TRIG_EN        WAKEUPENABLE = 32
	UART0_TRIG_EN      WAKEUPENABLE = 64
	UART1_TRIG_EN      WAKEUPENABLE = 128
	TOUCH_TRIG_EN      WAKEUPENABLE = 256
	SAR_TRIG_EN        WAKEUPENABLE = 512
	BT_TRIG_EN         WAKEUPENABLE = 1024
	RISCV_TRIG_EN      WAKEUPENABLE = 2048
	XTAL_DEAD_TRIG_EN  WAKEUPENABLE = 4096
	RISCV_TRAP_TRIG_EN WAKEUPENABLE = 8192
	USB_TRIG_EN        WAKEUPENABLE = 16384
)

/**
 * @brief  Get the reset reason for CPU.
 *
 * @param  int cpu_no : CPU no.
 *
 * @return RESET_REASON
 */
//go:linkname RtcGetResetReason C.rtc_get_reset_reason
func RtcGetResetReason(cpu_no c.Int) RESETREASON

/**
 * @brief  Get the wakeup cause for CPU.
 *
 * @param  int cpu_no : CPU no.
 *
 * @return WAKEUP_REASON
 */
//go:linkname RtcGetWakeupCause C.rtc_get_wakeup_cause
func RtcGetWakeupCause() WAKEUPREASON

// llgo:type C
type EspRomWakeFuncT func()

/**
 * @brief Read stored RTC wake function address
 *
 * Returns pointer to wake address if a value is set in RTC registers, and stored length & CRC all valid.
 * valid means that both stored stub length and stored wake function address are four-byte aligned non-zero values
 * and the crc check passes
 *
 * @param  None
 *
 * @return esp_rom_wake_func_t : Returns pointer to wake address if a value is set in RTC registers
 */
//go:linkname EspRomGetRtcWakeAddr C.esp_rom_get_rtc_wake_addr
func EspRomGetRtcWakeAddr() EspRomWakeFuncT

/**
 * @brief Store new RTC wake function address
 *
 * Set a new RTC wake address function. If a non-NULL function pointer is set then the function
 * memory is calculated and stored also.
 *
 * @param entry_addr Address of function. should be 4-bytes aligned otherwise it will not start from the stub after wake from deepsleep，
 *                   if NULL length will be ignored and all registers are cleared to 0.
 *
 * @param length length of function in RTC fast memory. should be less than RTC Fast memory size and aligned to 4-bytes.
 *               otherwise all registers are cleared to 0.
 *
 * @return None
 */
//go:linkname EspRomSetRtcWakeAddr C.esp_rom_set_rtc_wake_addr
func EspRomSetRtcWakeAddr(entry_addr EspRomWakeFuncT, length c.SizeT)

/**
 * @brief Software Reset digital core.
 *
 * It is not recommended to use this function in esp-idf, use
 * esp_restart() instead.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname SoftwareReset C.software_reset
func SoftwareReset()

/**
 * @brief Software Reset digital core.
 *
 * It is not recommended to use this function in esp-idf, use
 * esp_restart() instead.
 *
 * @param  int cpu_no : The CPU to reset, 0 for PRO CPU, 1 for APP CPU.
 *
 * @return None
 */
//go:linkname SoftwareResetCpu C.software_reset_cpu
func SoftwareResetCpu(cpu_no c.Int)
