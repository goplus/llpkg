package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ETSSTATUS c.Int

const (
	ETS_OK      ETSSTATUS = 0
	ETS_FAILED  ETSSTATUS = 1
	ETS_PENDING ETSSTATUS = 2
	ETS_BUSY    ETSSTATUS = 3
	ETS_CANCEL  ETSSTATUS = 4
)

type EtsStatusT ETSSTATUS
type ETSSignal c.Uint32T
type ETSParam c.Uint32T

type ETSEventTag struct {
	Sig ETSSignal
	Par ETSParam
}
type ETSEvent ETSEventTag

// llgo:type C
type ETSTask func(*ETSEvent)

// llgo:type C
type EtsIdleCbT func(c.Pointer)

/**
 * @brief  Set Pro cpu Entry code, code can be called in PRO CPU when booting is not completed.
 *         When Pro CPU booting is completed, Pro CPU will call the Entry code if not NULL.
 *
 * @param  uint32_t start : the PRO Entry code address value in uint32_t
 *
 * @return None
 */
//go:linkname EtsSetUserStart C.ets_set_user_start
func EtsSetUserStart(start c.Uint32T)

/**
 * @brief  Set App cpu Entry code, code can be called in PRO CPU.
 *         When APP booting is completed, APP CPU will call the Entry code if not NULL.
 *
 * @param  uint32_t start : the APP Entry code address value in uint32_t, stored in register APPCPU_CTRL_REG_D.
 *
 * @return None
 */
//go:linkname EtsSetAppcpuBootAddr C.ets_set_appcpu_boot_addr
func EtsSetAppcpuBootAddr(start c.Uint32T)

/**
 * @brief  Printf the strings to uart or other devices, similar with printf, simple than printf.
 *         Can not print float point data format, or longlong data format.
 *         So we maybe only use this in ROM.
 *
 * @param  const char *fmt : See printf.
 *
 * @param  ... : See printf.
 *
 * @return int : the length printed to the output device.
 */
//go:linkname EtsPrintf C.ets_printf
func EtsPrintf(fmt *c.Char, __llgo_va_list ...interface{}) c.Int

/**
 * @brief Get the uart channel of ets_printf(uart_tx_one_char).
 *
 * @return uint8_t uart channel used by ets_printf(uart_tx_one_char).
 */
//go:linkname EtsGetPrintfChannel C.ets_get_printf_channel
func EtsGetPrintfChannel() c.Uint8T

/**
 * @brief  Output a char to uart, which uart to output(which is in uart module in ROM) is not in scope of the function.
 *         Can not print float point data format, or longlong data format
 *
 * @param  char c : char to output.
 *
 * @return None
 */
//go:linkname EtsWriteCharUart C.ets_write_char_uart
func EtsWriteCharUart(c c.Char)

/**
 * @brief  Ets_printf have two output functions： putc1 and putc2, both of which will be called if need output.
 *         To install putc1, which is defaulted installed as ets_write_char_uart in none silent boot mode, as NULL in silent mode.
 *
 * @param  void (*)(char) p: Output function to install.
 *
 * @return None
 */
//go:linkname EtsInstallPutc1 C.ets_install_putc1
func EtsInstallPutc1(p func(c.Char))

/**
 * @brief  Ets_printf have two output functions： putc1 and putc2, both of which will be called if need output.
 *         To install putc2, which is defaulted installed as NULL.
 *
 * @param  void (*)(char) p: Output function to install.
 *
 * @return None
 */
//go:linkname EtsInstallPutc2 C.ets_install_putc2
func EtsInstallPutc2(p func(c.Char))

/**
 * @brief  Install putc1 as ets_write_char_uart.
 *         In silent boot mode(to void interfere the UART attached MCU), we can call this function, after booting ok.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname EtsInstallUartPrintf C.ets_install_uart_printf
func EtsInstallUartPrintf()

// llgo:type C
type ETSTimerFunc func(c.Pointer)

type X_ETSTIMER_ struct {
	TimerNext   *X_ETSTIMER_
	TimerExpire c.Uint32T
	TimerPeriod c.Uint32T
	TimerFunc   *ETSTimerFunc
	TimerArg    c.Pointer
}
type ETSTimer X_ETSTIMER_

/**
 * @brief  Init ets timer, this timer range is 640 us to 429496 ms
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname EtsTimerInit C.ets_timer_init
func EtsTimerInit()

/**
 * @brief  In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname EtsTimerDeinit C.ets_timer_deinit
func EtsTimerDeinit()

/**
 * @brief  Arm an ets timer, this timer range is 640 us to 429496 ms.
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  ETSTimer *timer : Timer struct pointer.
 *
 * @param  uint32_t tmout : Timer value in ms, range is 1 to 429496.
 *
 * @param  bool repeat : Timer is periodic repeated.
 *
 * @return None
 */
// llgo:link (*ETSTimer).EtsTimerArm C.ets_timer_arm
func (recv_ *ETSTimer) EtsTimerArm(tmout c.Uint32T, repeat bool) {
}

/**
 * @brief  Arm an ets timer, this timer range is 640 us to 429496 ms.
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  ETSTimer *timer : Timer struct pointer.
 *
 * @param  uint32_t tmout : Timer value in us, range is 1 to 429496729.
 *
 * @param  bool repeat : Timer is periodic repeated.
 *
 * @return None
 */
// llgo:link (*ETSTimer).EtsTimerArmUs C.ets_timer_arm_us
func (recv_ *ETSTimer) EtsTimerArmUs(us c.Uint32T, repeat bool) {
}

/**
 * @brief  Disarm an ets timer.
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  ETSTimer *timer : Timer struct pointer.
 *
 * @return None
 */
// llgo:link (*ETSTimer).EtsTimerDisarm C.ets_timer_disarm
func (recv_ *ETSTimer) EtsTimerDisarm() {
}

/**
 * @brief  Set timer callback and argument.
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  ETSTimer *timer : Timer struct pointer.
 *
 * @param  ETSTimerFunc *pfunction : Timer callback.
 *
 * @param  void *parg : Timer callback argument.
 *
 * @return None
 */
// llgo:link (*ETSTimer).EtsTimerSetfn C.ets_timer_setfn
func (recv_ *ETSTimer) EtsTimerSetfn(pfunction ETSTimerFunc, parg c.Pointer) {
}

/**
 * @brief  Unset timer callback and argument to NULL.
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  ETSTimer *timer : Timer struct pointer.
 *
 * @return None
 */
// llgo:link (*ETSTimer).EtsTimerDone C.ets_timer_done
func (recv_ *ETSTimer) EtsTimerDone() {
}

/**
 * @brief  CPU do while loop for some time.
 *         In FreeRTOS task, please call FreeRTOS apis.
 *
 * @param  uint32_t us : Delay time in us.
 *
 * @return None
 */
//go:linkname EtsDelayUs C.ets_delay_us
func EtsDelayUs(us c.Uint32T)

/**
 * @brief  Set the real CPU ticks per us to the ets, so that ets_delay_us will be accurate.
 *         Call this function when CPU frequency is changed.
 *
 * @param  uint32_t ticks_per_us : CPU ticks per us.
 *
 * @return None
 */
//go:linkname EtsUpdateCpuFrequency C.ets_update_cpu_frequency
func EtsUpdateCpuFrequency(ticks_per_us c.Uint32T)

/**
 * @brief  Get the real CPU ticks per us to the ets.
 *         This function do not return real CPU ticks per us, just the record in ets. It can be used to check with the real CPU frequency.
 *
 * @param  None
 *
 * @return uint32_t : CPU ticks per us record in ets.
 */
//go:linkname EtsGetCpuFrequency C.ets_get_cpu_frequency
func EtsGetCpuFrequency() c.Uint32T

/**
 * @brief  Get xtal_freq value, If value not stored in RTC_STORE5, than store.
 *
 * @param  None
 *
 * @return uint32_t : if stored in efuse(not 0)
 *                         clock = ets_efuse_get_xtal_freq() * 1000000;
 *                    else if analog_8M in efuse
 *                         clock = ets_get_xtal_scale() * 625 / 16 * ets_efuse_get_8M_clock();
 *                    else clock = 40M.
 */
//go:linkname EtsGetXtalFreq C.ets_get_xtal_freq
func EtsGetXtalFreq() c.Uint32T

/**
 * @brief  Get the apb divisor. The xtal frequency gets divided
 *         by this value to generate the APB clock.
 *         When any types of reset happens, the default value is 2.
 *
 * @param  None
 *
 * @return uint32_t : 1 or 2.
 */
//go:linkname EtsGetXtalDiv C.ets_get_xtal_div
func EtsGetXtalDiv() c.Uint32T

/**
 * @brief  Modifies the apb divisor. The xtal frequency gets divided by this to
 *         generate the APB clock.
 *
 * @note The xtal frequency divisor is 2 by default as the glitch detector
 *       doesn't properly stop glitches when it is 1. Please do not set the
 *       divisor to 1 before the PLL is active without being aware that you
 *       may be introducing a security risk.
 *
 * @param  div Divisor. 1 = xtal freq, 2 = 1/2th xtal freq.
 */
//go:linkname EtsSetXtalDiv C.ets_set_xtal_div
func EtsSetXtalDiv(div c.Int)

/**
 * @brief  Get apb_freq value, If value not stored in RTC_STORE5, than store.
 *
 * @param  None
 *
 * @return uint32_t : if rtc store the value (RTC_STORE5 high 16 bits and low 16 bits with same value), read from rtc register.
 *                         clock = (REG_READ(RTC_STORE5) & 0xffff) << 12;
 *                    else store ets_get_detected_xtal_freq() in.
 */
//go:linkname EtsGetApbFreq C.ets_get_apb_freq
func EtsGetApbFreq() c.Uint32T

// llgo:type C
type EtsIsrT func(c.Pointer)

/**
 * @brief  Attach a interrupt handler to a CPU interrupt number.
 *         This function equals to _xtos_set_interrupt_handler_arg(i, func, arg).
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  int i : CPU interrupt number.
 *
 * @param  ets_isr_t func : Interrupt handler.
 *
 * @param  void *arg : argument of the handler.
 *
 * @return None
 */
//go:linkname EtsIsrAttach C.ets_isr_attach
func EtsIsrAttach(i c.Int, func_ EtsIsrT, arg c.Pointer)

/**
 * @brief  Mask the interrupts which show in mask bits.
 *         This function equals to _xtos_ints_off(mask).
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  uint32_t mask : BIT(i) means mask CPU interrupt number i.
 *
 * @return None
 */
//go:linkname EtsIsrMask C.ets_isr_mask
func EtsIsrMask(mask c.Uint32T)

/**
 * @brief  Unmask the interrupts which show in mask bits.
 *         This function equals to _xtos_ints_on(mask).
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  uint32_t mask : BIT(i) means mask CPU interrupt number i.
 *
 * @return None
 */
//go:linkname EtsIsrUnmask C.ets_isr_unmask
func EtsIsrUnmask(unmask c.Uint32T)

/**
 * @brief  Lock the interrupt to level 2.
 *         This function direct set the CPU registers.
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname EtsIntrLock C.ets_intr_lock
func EtsIntrLock()

/**
 * @brief  Unlock the interrupt to level 0.
 *         This function direct set the CPU registers.
 *         In FreeRTOS, please call FreeRTOS apis, never call this api.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname EtsIntrUnlock C.ets_intr_unlock
func EtsIntrUnlock()

/**
 * @brief  Attach an CPU interrupt to a hardware source.
 *         We have 4 steps to use an interrupt:
 *         1.Attach hardware interrupt source to CPU.  intr_matrix_set(0, ETS_WIFI_MAC_INTR_SOURCE, ETS_WMAC_INUM);
 *         2.Set interrupt handler.                    xt_set_interrupt_handler(ETS_WMAC_INUM, func, NULL);
 *         3.Enable interrupt for CPU.                 xt_ints_on(1 << ETS_WMAC_INUM);
 *         4.Enable interrupt in the module.
 *
 * @param  int cpu_no : The CPU which the interrupt number belongs.
 *
 * @param  uint32_t model_num : The interrupt hardware source number, please see the interrupt hardware source table.
 *
 * @param  uint32_t intr_num : The interrupt number CPU, please see the interrupt cpu using table.
 *
 * @return None
 */
//go:linkname IntrMatrixSet C.intr_matrix_set
func IntrMatrixSet(cpu_no c.Int, model_num c.Uint32T, intr_num c.Uint32T)
