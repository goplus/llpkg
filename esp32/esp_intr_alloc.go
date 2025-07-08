package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Mark an interrupt as a shared interrupt
 *
 * This will mark a certain interrupt on the specified CPU as
 * an interrupt that can be used to hook shared interrupt handlers
 * to.
 *
 * @param intno The number of the interrupt (0-31)
 * @param cpu CPU on which the interrupt should be marked as shared (0 or 1)
 * @param is_in_iram Shared interrupt is for handlers that reside in IRAM and
 *                   the int can be left enabled while the flash cache is disabled.
 *
 * @return ESP_ERR_INVALID_ARG if cpu or intno is invalid
 *         ESP_OK otherwise
 */
//go:linkname EspIntrMarkShared C.esp_intr_mark_shared
func EspIntrMarkShared(intno c.Int, cpu c.Int, is_in_iram bool) EspErrT

/**
 * @brief Reserve an interrupt to be used outside of this framework
 *
 * This will mark a certain interrupt on the specified CPU as
 * reserved, not to be allocated for any reason.
 *
 * @param intno The number of the interrupt (0-31)
 * @param cpu CPU on which the interrupt should be marked as shared (0 or 1)
 *
 * @return ESP_ERR_INVALID_ARG if cpu or intno is invalid
 *         ESP_OK otherwise
 */
//go:linkname EspIntrReserve C.esp_intr_reserve
func EspIntrReserve(intno c.Int, cpu c.Int) EspErrT

/**
 * @brief Allocate an interrupt with the given parameters.
 *
 * This finds an interrupt that matches the restrictions as given in the flags
 * parameter, maps the given interrupt source to it and hooks up the given
 * interrupt handler (with optional argument) as well. If needed, it can return
 * a handle for the interrupt as well.
 *
 * The interrupt will always be allocated on the core that runs this function.
 *
 * If ESP_INTR_FLAG_IRAM flag is used, and handler address is not in IRAM or
 * RTC_FAST_MEM, then ESP_ERR_INVALID_ARG is returned.
 *
 * @param source The interrupt source. One of the ETS_*_INTR_SOURCE interrupt mux
 *               sources, as defined in soc/soc.h, or one of the internal
 *               ETS_INTERNAL_*_INTR_SOURCE sources as defined in this header.
 * @param flags An ORred mask of the ESP_INTR_FLAG_* defines. These restrict the
 *               choice of interrupts that this routine can choose from. If this value
 *               is 0, it will default to allocating a non-shared interrupt of level
 *               1, 2 or 3. If this is ESP_INTR_FLAG_SHARED, it will allocate a shared
 *               interrupt of level 1. Setting ESP_INTR_FLAG_INTRDISABLED will return
 *               from this function with the interrupt disabled.
 * @param handler The interrupt handler. Must be NULL when an interrupt of level >3
 *               is requested, because these types of interrupts aren't C-callable.
 * @param arg    Optional argument for passed to the interrupt handler
 * @param ret_handle Pointer to an intr_handle_t to store a handle that can later be
 *               used to request details or free the interrupt. Can be NULL if no handle
 *               is required.
 *
 * @return ESP_ERR_INVALID_ARG if the combination of arguments is invalid.
 *         ESP_ERR_NOT_FOUND No free interrupt found with the specified flags
 *         ESP_OK otherwise
 */
//go:linkname EspIntrAlloc C.esp_intr_alloc
func EspIntrAlloc(source c.Int, flags c.Int, handler IntrHandlerT, arg c.Pointer, ret_handle *IntrHandleT) EspErrT

/**
 * @brief Allocate an interrupt with the given parameters.
 *
 *
 * This essentially does the same as esp_intr_alloc, but allows specifying a register and mask
 * combo. For shared interrupts, the handler is only called if a read from the specified
 * register, ANDed with the mask, returns non-zero. By passing an interrupt status register
 * address and a fitting mask, this can be used to accelerate interrupt handling in the case
 * a shared interrupt is triggered; by checking the interrupt statuses first, the code can
 * decide which ISRs can be skipped
 *
 * @param source The interrupt source. One of the ETS_*_INTR_SOURCE interrupt mux
 *               sources, as defined in soc/soc.h, or one of the internal
 *               ETS_INTERNAL_*_INTR_SOURCE sources as defined in this header.
 * @param flags An ORred mask of the ESP_INTR_FLAG_* defines. These restrict the
 *               choice of interrupts that this routine can choose from. If this value
 *               is 0, it will default to allocating a non-shared interrupt of level
 *               1, 2 or 3. If this is ESP_INTR_FLAG_SHARED, it will allocate a shared
 *               interrupt of level 1. Setting ESP_INTR_FLAG_INTRDISABLED will return
 *               from this function with the interrupt disabled.
 * @param intrstatusreg The address of an interrupt status register
 * @param intrstatusmask A mask. If a read of address intrstatusreg has any of the bits
 *               that are 1 in the mask set, the ISR will be called. If not, it will be
 *               skipped.
 * @param handler The interrupt handler. Must be NULL when an interrupt of level >3
 *               is requested, because these types of interrupts aren't C-callable.
 * @param arg    Optional argument for passed to the interrupt handler
 * @param ret_handle Pointer to an intr_handle_t to store a handle that can later be
 *               used to request details or free the interrupt. Can be NULL if no handle
 *               is required.
 *
 * @return ESP_ERR_INVALID_ARG if the combination of arguments is invalid.
 *         ESP_ERR_NOT_FOUND No free interrupt found with the specified flags
 *         ESP_OK otherwise
 */
//go:linkname EspIntrAllocIntrstatus C.esp_intr_alloc_intrstatus
func EspIntrAllocIntrstatus(source c.Int, flags c.Int, intrstatusreg c.Uint32T, intrstatusmask c.Uint32T, handler IntrHandlerT, arg c.Pointer, ret_handle *IntrHandleT) EspErrT

/**
 * @brief Disable and free an interrupt.
 *
 * Use an interrupt handle to disable the interrupt and release the resources associated with it.
 * If the current core is not the core that registered this interrupt, this routine will be assigned to
 * the core that allocated this interrupt, blocking and waiting until the resource is successfully released.
 *
 * @note
 * When the handler shares its source with other handlers, the interrupt status
 * bits it's responsible for should be managed properly before freeing it. see
 * ``esp_intr_disable`` for more details. Please do not call this function in ``esp_ipc_call_blocking``.
 *
 * @param handle The handle, as obtained by esp_intr_alloc or esp_intr_alloc_intrstatus
 *
 * @return ESP_ERR_INVALID_ARG the handle is NULL
 *         ESP_FAIL failed to release this handle
 *         ESP_OK otherwise
 */
//go:linkname EspIntrFree C.esp_intr_free
func EspIntrFree(handle IntrHandleT) EspErrT

/**
 * @brief Get CPU number an interrupt is tied to
 *
 * @param handle The handle, as obtained by esp_intr_alloc or esp_intr_alloc_intrstatus
 *
 * @return The core number where the interrupt is allocated
 */
//go:linkname EspIntrGetCpu C.esp_intr_get_cpu
func EspIntrGetCpu(handle IntrHandleT) c.Int

/**
 * @brief Get the allocated interrupt for a certain handle
 *
 * @param handle The handle, as obtained by esp_intr_alloc or esp_intr_alloc_intrstatus
 *
 * @return The interrupt number
 */
//go:linkname EspIntrGetIntno C.esp_intr_get_intno
func EspIntrGetIntno(handle IntrHandleT) c.Int

/**
 * @brief Disable the interrupt associated with the handle
 *
 * @note
 * 1. For local interrupts (ESP_INTERNAL_* sources), this function has to be called on the
 * CPU the interrupt is allocated on. Other interrupts have no such restriction.
 * 2. When several handlers sharing a same interrupt source, interrupt status bits, which are
 * handled in the handler to be disabled, should be masked before the disabling, or handled
 * in other enabled interrupts properly. Miss of interrupt status handling will cause infinite
 * interrupt calls and finally system crash.
 *
 * @param handle The handle, as obtained by esp_intr_alloc or esp_intr_alloc_intrstatus
 *
 * @return ESP_ERR_INVALID_ARG if the combination of arguments is invalid.
 *         ESP_OK otherwise
 */
//go:linkname EspIntrDisable C.esp_intr_disable
func EspIntrDisable(handle IntrHandleT) EspErrT

/**
 * @brief Enable the interrupt associated with the handle
 *
 * @note For local interrupts (ESP_INTERNAL_* sources), this function has to be called on the
 *       CPU the interrupt is allocated on. Other interrupts have no such restriction.
 *
 * @param handle The handle, as obtained by esp_intr_alloc or esp_intr_alloc_intrstatus
 *
 * @return ESP_ERR_INVALID_ARG if the combination of arguments is invalid.
 *         ESP_OK otherwise
 */
//go:linkname EspIntrEnable C.esp_intr_enable
func EspIntrEnable(handle IntrHandleT) EspErrT

/**
 * @brief Set the "in IRAM" status of the handler.
 *
 * @note Does not work on shared interrupts.
 *
 * @param handle The handle, as obtained by esp_intr_alloc or esp_intr_alloc_intrstatus
 * @param is_in_iram Whether the handler associated with this handle resides in IRAM.
 *                   Handlers residing in IRAM can be called when cache is disabled.
 *
 * @return ESP_ERR_INVALID_ARG if the combination of arguments is invalid.
 *         ESP_OK otherwise
 */
//go:linkname EspIntrSetInIram C.esp_intr_set_in_iram
func EspIntrSetInIram(handle IntrHandleT, is_in_iram bool) EspErrT

/**
 * @brief Disable interrupts that aren't specifically marked as running from IRAM
 */
//go:linkname EspIntrNoniramDisable C.esp_intr_noniram_disable
func EspIntrNoniramDisable()

/**
 * @brief Re-enable interrupts disabled by esp_intr_noniram_disable
 */
//go:linkname EspIntrNoniramEnable C.esp_intr_noniram_enable
func EspIntrNoniramEnable()

/**
 * @brief enable the interrupt source based on its number
 * @param inum interrupt number from 0 to 31
 */
//go:linkname EspIntrEnableSource C.esp_intr_enable_source
func EspIntrEnableSource(inum c.Int)

/**
 * @brief disable the interrupt source based on its number
 * @param inum interrupt number from 0 to 31
 */
//go:linkname EspIntrDisableSource C.esp_intr_disable_source
func EspIntrDisableSource(inum c.Int)

/**
 * @brief Dump the status of allocated interrupts
 * @param stream  The stream to dump to, if NULL then stdout is used
 * @return ESP_OK on success
 */
//go:linkname EspIntrDump C.esp_intr_dump
func EspIntrDump(stream *c.FILE) EspErrT

/**
 * @brief Check if the given pointer is in the safe ISR area.
 * In other words, make sure that the pointer's content is accessible at
 * any time, regardless of the cache status
 *
 * @param ptr Pointer to check
 *
 * @return true if `ptr` points to ISR area, false else
 */
//go:linkname EspIntrPtrInIsrRegion C.esp_intr_ptr_in_isr_region
func EspIntrPtrInIsrRegion(ptr c.Pointer) bool
