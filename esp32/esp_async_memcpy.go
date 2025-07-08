package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AsyncMemcpyContextT struct {
	Unused [8]uint8
}
type AsyncMemcpyHandleT *AsyncMemcpyContextT
type AsyncMemcpyT AsyncMemcpyHandleT

/**
 * @brief Async memory copy event data
 */

type AsyncMemcpyEventT struct {
	Data c.Pointer
}

// llgo:type C
type AsyncMemcpyIsrCbT func(AsyncMemcpyHandleT, *AsyncMemcpyEventT, c.Pointer) bool

/**
 * @brief Type of async memcpy configuration
 */

type AsyncMemcpyConfigT struct {
	Backlog        c.Uint32T
	SramTransAlign c.SizeT
	Flags          c.Uint32T
}

/**
 * @brief Install async memcpy driver, with AHB-GDMA as the backend
 *
 * @param[in] config Configuration of async memcpy
 * @param[out] mcp Returned driver handle
 * @return
 *      - ESP_OK: Install async memcpy driver successfully
 *      - ESP_ERR_INVALID_ARG: Install async memcpy driver failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Install async memcpy driver failed because out of memory
 *      - ESP_FAIL: Install async memcpy driver failed because of other error
 */
// llgo:link (*AsyncMemcpyConfigT).EspAsyncMemcpyInstallGdmaAhb C.esp_async_memcpy_install_gdma_ahb
func (recv_ *AsyncMemcpyConfigT) EspAsyncMemcpyInstallGdmaAhb(mcp *AsyncMemcpyHandleT) EspErrT {
	return 0
}

/**
 * @brief Install async memcpy driver with the default DMA backend
 *
 * @note On chip with CPDMA support, CPDMA is the default choice.
 *       On chip with AHB-GDMA support, AHB-GDMA is the default choice.
 *
 * @param[in] config Configuration of async memcpy
 * @param[out] mcp Returned driver handle
 * @return
 *      - ESP_OK: Install async memcpy driver successfully
 *      - ESP_ERR_INVALID_ARG: Install async memcpy driver failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Install async memcpy driver failed because out of memory
 *      - ESP_FAIL: Install async memcpy driver failed because of other error
 */
// llgo:link (*AsyncMemcpyConfigT).EspAsyncMemcpyInstall C.esp_async_memcpy_install
func (recv_ *AsyncMemcpyConfigT) EspAsyncMemcpyInstall(mcp *AsyncMemcpyHandleT) EspErrT {
	return 0
}

/**
 * @brief Uninstall async memcpy driver
 *
 * @param[in] mcp Handle of async memcpy driver that returned from `esp_async_memcpy_install`
 * @return
 *      - ESP_OK: Uninstall async memcpy driver successfully
 *      - ESP_ERR_INVALID_ARG: Uninstall async memcpy driver failed because of invalid argument
 *      - ESP_FAIL: Uninstall async memcpy driver failed because of other error
 */
//go:linkname EspAsyncMemcpyUninstall C.esp_async_memcpy_uninstall
func EspAsyncMemcpyUninstall(mcp AsyncMemcpyHandleT) EspErrT

/**
 * @brief Send an asynchronous memory copy request
 *
 * @note The callback function is invoked in interrupt context, never do blocking jobs in the callback.
 *
 * @param[in] mcp Handle of async memcpy driver that returned from `esp_async_memcpy_install`
 * @param[in] dst Destination address (copy to)
 * @param[in] src Source address (copy from)
 * @param[in] n Number of bytes to copy
 * @param[in] cb_isr Callback function, which got invoked in interrupt context. Set to NULL can bypass the callback.
 * @param[in] cb_args User defined argument to be passed to the callback function
 * @return
 *      - ESP_OK: Send memory copy request successfully
 *      - ESP_ERR_INVALID_ARG: Send memory copy request failed because of invalid argument
 *      - ESP_FAIL: Send memory copy request failed because of other error
 */
//go:linkname EspAsyncMemcpy C.esp_async_memcpy
func EspAsyncMemcpy(mcp AsyncMemcpyHandleT, dst c.Pointer, src c.Pointer, n c.SizeT, cb_isr AsyncMemcpyIsrCbT, cb_args c.Pointer) EspErrT
