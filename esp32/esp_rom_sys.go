package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Software Reset digital core include RTC.
 *
 * It is not recommended to use this function in esp-idf, use
 * esp_restart() instead.
 */
//go:linkname EspRomSoftwareResetSystem C.esp_rom_software_reset_system
func EspRomSoftwareResetSystem()

/**
 * @brief Software Reset cpu core.
 *
 * It is not recommended to use this function in esp-idf, use
 * esp_restart() instead.
 *
 * @param  cpu_no : The CPU to reset, 0 for PRO CPU, 1 for APP CPU.
 */
//go:linkname EspRomSoftwareResetCpu C.esp_rom_software_reset_cpu
func EspRomSoftwareResetCpu(cpu_no c.Int)

/**
 * @brief Print formatted string to console device
 * @note float and long long data are not supported!
 *
 * @param fmt Format string
 * @param ... Additional arguments, depending on the format string
 * @return int: Total number of characters written on success; A negative number on failure.
 */
//go:linkname EspRomPrintf C.esp_rom_printf
func EspRomPrintf(fmt *c.Char, __llgo_va_list ...interface{}) c.Int

/**
 * @brief Print formatted string to console device
 * @note float and long long data are not supported!
 *
 * @param fmt Format string
 * @param ap List of arguments.
 * @return int: Total number of characters written on success; A negative number on failure.
 */
//go:linkname EspRomVprintf C.esp_rom_vprintf
func EspRomVprintf(fmt *c.Char, ap c.VaList) c.Int

/**
 * @brief Pauses execution for us microseconds
 *
 * @param us Number of microseconds to pause
 */
//go:linkname EspRomDelayUs C.esp_rom_delay_us
func EspRomDelayUs(us c.Uint32T)

/**
 * @brief esp_rom_printf can print message to different channels simultaneously.
 *        This function can help install the low level putc function for esp_rom_printf.
 *
 * @param channel Channel number (starting from 1)
 * @param putc Function pointer to the putc implementation. Set NULL can disconnect esp_rom_printf with putc.
 */
//go:linkname EspRomInstallChannelPutc C.esp_rom_install_channel_putc
func EspRomInstallChannelPutc(channel c.Int, putc func(c.Char))

/**
 * @brief It outputs a character to different channels simultaneously.
 *        This function is used by esp_rom_printf/esp_rom_vprintf.
 *
 * @param c Char to output.
 */
//go:linkname EspRomOutputToChannels C.esp_rom_output_to_channels
func EspRomOutputToChannels(c c.Char)

/**
 * @brief Install UART1 as the default console channel, equivalent to `esp_rom_install_channel_putc(1, esp_rom_output_putc)`
 */
//go:linkname EspRomInstallUartPrintf C.esp_rom_install_uart_printf
func EspRomInstallUartPrintf()

/**
 * @brief Get reset reason of CPU
 *
 * @param cpu_no CPU number
 * @return Reset reason code (see in soc/reset_reasons.h)
 */
//go:linkname EspRomGetResetReason C.esp_rom_get_reset_reason
func EspRomGetResetReason(cpu_no c.Int) SocResetReasonT

/**
 * @brief Route peripheral interrupt sources to CPU's interrupt port by matrix
 *
 * Usually there're 4 steps to use an interrupt:
 * 1. Route peripheral interrupt source to CPU. e.g.  esp_rom_route_intr_matrix(0, ETS_WIFI_MAC_INTR_SOURCE, ETS_WMAC_INUM)
 * 2. Set interrupt handler for CPU
 * 3. Enable CPU interrupt
 * 4. Enable peripheral interrupt
 *
 * @param cpu_core The CPU number, which the peripheral interrupt will inform to
 * @param periph_intr_id The peripheral interrupt source number
 * @param cpu_intr_num The CPU (external) interrupt number. On targets that use CLIC as their interrupt controller,
 *                     this number represents the external interrupt number. For example, passing `cpu_intr_num = i`
 *                     to this function would in fact bind peripheral source to CPU interrupt `CLIC_EXT_INTR_NUM_OFFSET + i`.
 */
//go:linkname EspRomRouteIntrMatrix C.esp_rom_route_intr_matrix
func EspRomRouteIntrMatrix(cpu_core c.Int, periph_intr_id c.Uint32T, cpu_intr_num c.Uint32T)

/**
 * @brief Get the real CPU ticks per us
 *
 * @return CPU ticks per us
 */
//go:linkname EspRomGetCpuTicksPerUs C.esp_rom_get_cpu_ticks_per_us
func EspRomGetCpuTicksPerUs() c.Uint32T

/**
 * @brief Set the real CPU tick rate
 *
 * @note Call this function when CPU frequency is changed, otherwise the `esp_rom_delay_us` can be inaccurate.
 *
 * @param ticks_per_us CPU ticks per us
 */
//go:linkname EspRomSetCpuTicksPerUs C.esp_rom_set_cpu_ticks_per_us
func EspRomSetCpuTicksPerUs(ticks_per_us c.Uint32T)
