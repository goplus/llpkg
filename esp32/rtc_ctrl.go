package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Register a handler for specific RTC_CNTL interrupts
 *
 * Multiple handlers can be registered using this function. Whenever an
 * RTC interrupt happens, all handlers with matching rtc_intr_mask values
 * will be called.
 *
 * @param handler  handler function to call
 * @param handler_arg  argument to be passed to the handler
 * @param rtc_intr_mask  combination of RTC_CNTL_*_INT_ENA bits indicating the
 *                       sources to call the handler for
 * @param flags  An ORred mask of the RTC_INTR_FLAG_* defines. You can pass different
 *               flags to it to realize different purpose. If 0, the interrupt will
 *               not handle anything special. If you pass `RTC_INTR_FLAG_IRAM`, means
 *               the interrupt can be triggered with cache disabled.
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NO_MEM not enough memory to allocate handler structure
 *      - other errors returned by esp_intr_alloc
 */
//go:linkname RtcIsrRegister C.rtc_isr_register
func RtcIsrRegister(handler IntrHandlerT, handler_arg c.Pointer, rtc_intr_mask c.Uint32T, flags c.Uint32T) EspErrT

/**
 * @brief Deregister the handler previously registered using rtc_isr_register
 * @param handler  handler function to call (as passed to rtc_isr_register)
 * @param handler_arg  argument of the handler (as passed to rtc_isr_register)
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if a handler matching both handler and
 *        handler_arg isn't registered
 */
//go:linkname RtcIsrDeregister C.rtc_isr_deregister
func RtcIsrDeregister(handler IntrHandlerT, handler_arg c.Pointer) EspErrT

/**
 * @brief Disable the RTC interrupt that is allowed to be executed when cache is disabled.
 * cache disabled. Internal interrupt handle function will call this function in interrupt
 * handler function. Disable bits when `esp_intr_noniram_disable` is called.
 *
 * @param cpu CPU number.
 */
//go:linkname RtcIsrNoniramDisable C.rtc_isr_noniram_disable
func RtcIsrNoniramDisable(cpu c.Uint32T)

/**
 * @brief Enable the RTC interrupt that is allowed to be executed when cache is disabled.
 * cache disabled. Internal interrupt handle function will call this function in interrupt
 * handler function. Enable bits when `esp_intr_noniram_enable` is called.
 *
 * @param cpu CPU number.
 */
//go:linkname RtcIsrNoniramEnable C.rtc_isr_noniram_enable
func RtcIsrNoniramEnable(cpu c.Uint32T)
