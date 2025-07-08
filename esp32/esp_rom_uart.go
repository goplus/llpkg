package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_ROM_CDC_ACM_WORK_BUF_MIN = 128

type EspRomUartNumT c.Int

const (
	ESP_ROM_UART_0   EspRomUartNumT = 0
	ESP_ROM_UART_1   EspRomUartNumT = 1
	ESP_ROM_UART_USB EspRomUartNumT = 2
)

/**
 * @brief Wait for UART TX FIFO is empty and all data has been sent out.
 *
 * @param serial_num The serial number defined in ROM, including UART_x, USB_OTG, USB_SERIAL_JTAG..
 */
//go:linkname EspRomOutputTxWaitIdle C.esp_rom_output_tx_wait_idle
func EspRomOutputTxWaitIdle(serial_num c.Uint8T)

//go:linkname EspRomUartTxWaitIdle C.esp_rom_uart_tx_wait_idle
func EspRomUartTxWaitIdle(serial_num c.Uint8T)

/**
 * @brief Wait until UART TX FIFO is empty (i.e. flush TX FIFO)
 *
 * @param serial_num UART port number
 */
//go:linkname EspRomOutputFlushTx C.esp_rom_output_flush_tx
func EspRomOutputFlushTx(serial_num c.Uint8T)

//go:linkname EspRomUartFlushTx C.esp_rom_uart_flush_tx
func EspRomUartFlushTx(serial_num c.Uint8T)

/**
 * @brief Transmit one character to the console channel.
 *
 * @param c Character to send
 * @return
 *      - 0 on success
 *      - 1 on failure
 */
//go:linkname EspRomOutputTxOneChar C.esp_rom_output_tx_one_char
func EspRomOutputTxOneChar(c c.Uint8T) c.Int

//go:linkname EspRomUartTxOneChar C.esp_rom_uart_tx_one_char
func EspRomUartTxOneChar(c c.Uint8T) c.Int

/**
 * @brief Transmit one character to the console channel.
 * @note This function is a wrapper over esp_rom_uart_tx_one_char, it can help handle line ending issue by replacing '\n' with '\r\n'.
 *
 * @param c Character to send
 */
//go:linkname EspRomOutputPutc C.esp_rom_output_putc
func EspRomOutputPutc(c c.Char)

//go:linkname EspRomUartPutc C.esp_rom_uart_putc
func EspRomUartPutc(c c.Char)

/**
 * @brief Get one character from the console channel.
 *
 * @param c Where to store the character
 * @return
 *      - 0 on success
 *      - 1 on failure or no data available
 */
//go:linkname EspRomOutputRxOneChar C.esp_rom_output_rx_one_char
func EspRomOutputRxOneChar(c *c.Uint8T) c.Int

//go:linkname EspRomUartRxOneChar C.esp_rom_uart_rx_one_char
func EspRomUartRxOneChar(c *c.Uint8T) c.Int

/**
 * @brief Get one line of string from console channel (line ending won't be stored in the buffer).
 *
 * @param str Where to store the string
 * @param max_len Maximum length of the buffer (including the NULL delimiter)
 * @return always return 0 when on success or wait in a loop for rx data
 */
//go:linkname EspRomOutputRxString C.esp_rom_output_rx_string
func EspRomOutputRxString(str *c.Uint8T, max_len c.Uint8T) c.Int

//go:linkname EspRomUartRxString C.esp_rom_uart_rx_string
func EspRomUartRxString(str *c.Uint8T, max_len c.Uint8T) c.Int

/**
 * @brief Set the UART port used by ets_printf.
 *
 * @note USB-CDC port is also treated as "UART" port in the ROM code.
 *       Use ESP_ROM_USB_SERIAL_DEVICE_NUM or ESP_ROM_USB_OTG_NUM to identify USB_SERIAL_JTAG and USB_OTG, respectively.
 *
 * @param serial_num UART port number
 */
//go:linkname EspRomOutputSetAsConsole C.esp_rom_output_set_as_console
func EspRomOutputSetAsConsole(serial_num c.Uint8T)

//go:linkname EspRomUartSetAsConsole C.esp_rom_uart_set_as_console
func EspRomUartSetAsConsole(serial_num c.Uint8T)

/**
 * @brief Switch the UART port that will use a buffer for TX and RX.
 *
 * @note USB-CDC port is also treated as "UART" port in the ROM code.
 *       Use ESP_ROM_USB_SERIAL_DEVICE_NUM or ESP_ROM_USB_OTG_NUM to identify USB_SERIAL_JTAG and USB_OTG, respectively.
 *
 * @param serial_num UART port number
 */
//go:linkname EspRomOutputSwitchBuffer C.esp_rom_output_switch_buffer
func EspRomOutputSwitchBuffer(serial_num c.Uint8T)

//go:linkname EspRomUartSwitchBuffer C.esp_rom_uart_switch_buffer
func EspRomUartSwitchBuffer(serial_num c.Uint8T)

/**
 * @brief Initialize the USB ACM UART
 * @note The ACM working memroy should be at least 128 bytes (ESP_ROM_CDC_ACM_WORK_BUF_MIN) in size.
 *
 * @param cdc_acm_work_mem Pointer to the work memroy used for CDC-ACM
 * @param cdc_acm_work_mem_len Length of work memory
 */
//go:linkname EspRomOutputUsbAcmInit C.esp_rom_output_usb_acm_init
func EspRomOutputUsbAcmInit(cdc_acm_work_mem c.Pointer, cdc_acm_work_mem_len c.Int)

//go:linkname EspRomUartUsbAcmInit C.esp_rom_uart_usb_acm_init
func EspRomUartUsbAcmInit(serial_num c.Uint8T)
