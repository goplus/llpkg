package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type EspUsbConsoleCbT func(c.Pointer)

/**
 * Initialize USB console output using ROM USB CDC driver.
 * This function is called by the early startup code if USB CDC is
 * selected as the console output option.
 * @return
 *  - ESP_OK on success
 *  - ESP_ERR_NO_MEM
 *  - other error codes from the interrupt allocator
 */
//go:linkname EspUsbConsoleInit C.esp_usb_console_init
func EspUsbConsoleInit() EspErrT

/**
 * Write a buffer to USB CDC
 * @param buf  data to write
 * @param size  size of the data, in bytes
 * @return -1 on error, otherwise the number of bytes
 */
//go:linkname EspUsbConsoleWriteBuf C.esp_usb_console_write_buf
func EspUsbConsoleWriteBuf(buf *c.Char, size c.SizeT) c.SsizeT

/**
 * @brief Wait until all buffered USB CDC output is written
 *
 * @return ssize_t  Number of bytes written, or -1 if the driver is not initialized
 */
//go:linkname EspUsbConsoleFlush C.esp_usb_console_flush
func EspUsbConsoleFlush() c.SsizeT

/**
 * @brief Read data from USB CDC
 *
 * May read less data than requested.
 *
 * @param buf  Buffer to read data into
 * @param buf_size  Size of the buffer
 * @return ssize_t  Number of bytes written into the buffer, or -1 if the driver is not initialized
 */
//go:linkname EspUsbConsoleReadBuf C.esp_usb_console_read_buf
func EspUsbConsoleReadBuf(buf *c.Char, buf_size c.SizeT) c.SsizeT

/**
 * @brief Get the number of bytes available for reading from USB CDC
 *
 * @return ssize_t Number of bytes available, or -1 if the driver is not initialized
 */
//go:linkname EspUsbConsoleAvailableForRead C.esp_usb_console_available_for_read
func EspUsbConsoleAvailableForRead() c.SsizeT

/**
 * @brief Check if data can be written into USB CDC
 *
 * @return true if data can be written now without blocking
 */
//go:linkname EspUsbConsoleWriteAvailable C.esp_usb_console_write_available
func EspUsbConsoleWriteAvailable() bool

/**
 * @brief Set RX/TX callback functions to be called from ISR
 *
 * @param rx_cb  RX callback function
 * @param tx_cb  TX callback function
 * @param arg  callback-specific context pointer
 * @return ESP_OK if the callbacks were set, ESP_ERR_INVALID_STATE if the driver is not initialized
 */
//go:linkname EspUsbConsoleSetCb C.esp_usb_console_set_cb
func EspUsbConsoleSetCb(rx_cb EspUsbConsoleCbT, tx_cb EspUsbConsoleCbT, arg c.Pointer) EspErrT

/**
 * @brief Call the USB interrupt handler while any interrupts
 * are pending, but not more than a few times. Non-static to
 * allow placement into IRAM by ldgen.
 *
 */
//go:linkname EspUsbConsolePollInterrupts C.esp_usb_console_poll_interrupts
func EspUsbConsolePollInterrupts()
