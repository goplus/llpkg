package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const BUS_LOCK_DEBUG = 0
const DEV_NUM_MAX = 6

type SpiBusLockT struct {
	Unused [8]uint8
}

type SpiBusLockDevT struct {
	Unused [8]uint8
}
type SpiBusLockHandleT *SpiBusLockT
type SpiBusLockDevHandleT *SpiBusLockDevT

// llgo:type C
type BgCtrlFuncT func(c.Pointer)

/**
 * @brief Try to claim a SPI peripheral
 *
 * Call this if your driver wants to manage a SPI peripheral.
 *
 * @param host Peripheral to claim
 * @param source The caller indentification string.
 *
 * @return True if peripheral is claimed successfully; false if peripheral already is claimed.
 */
// llgo:link SpiHostDeviceT.SpicommonPeriphClaim C.spicommon_periph_claim
func (recv_ SpiHostDeviceT) SpicommonPeriphClaim(source *c.Char) bool {
	return false
}

/**
 * @brief Check whether the spi periph is in use.
 *
 * @param host Peripheral to check.
 *
 * @return True if in use, otherwise false.
 */
// llgo:link SpiHostDeviceT.SpicommonPeriphInUse C.spicommon_periph_in_use
func (recv_ SpiHostDeviceT) SpicommonPeriphInUse() bool {
	return false
}

/**
 * @brief Return the SPI peripheral so another driver can claim it.
 *
 * @param host Peripheral to return
 *
 * @return True if peripheral is returned successfully; false if peripheral was free to claim already.
 */
// llgo:link SpiHostDeviceT.SpicommonPeriphFree C.spicommon_periph_free
func (recv_ SpiHostDeviceT) SpicommonPeriphFree() bool {
	return false
}

// / Lock configuration struct
type SpiBusLockConfigT struct {
	HostId c.Int
	CsNum  c.Int
}

// / Child-lock configuration struct
type SpiBusLockDevConfigT struct {
	Flags c.Uint32T
}

/************* Common *********************/
/**
 * Initialize a lock for an SPI bus.
 *
 * @param out_lock Output of the handle to the lock
 * @return
 *  - ESP_ERR_NO_MEM: if memory exhausted
 *  - ESP_OK: if success
 */
//go:linkname SpiBusInitLock C.spi_bus_init_lock
func SpiBusInitLock(out_lock *SpiBusLockHandleT, config *SpiBusLockConfigT) EspErrT

/**
 * Free the resources used by an SPI bus lock.
 *
 * @note All attached devices should have been unregistered before calling this
 *       funciton.
 *
 * @param lock Handle to the lock to free.
 */
//go:linkname SpiBusDeinitLock C.spi_bus_deinit_lock
func SpiBusDeinitLock(lock SpiBusLockHandleT)

/**
 * @brief Get the corresponding lock according to bus id.
 *
 * @param host_id The bus id to get the lock
 * @return The lock handle
 */
// llgo:link SpiHostDeviceT.SpiBusLockGetById C.spi_bus_lock_get_by_id
func (recv_ SpiHostDeviceT) SpiBusLockGetById() SpiBusLockHandleT {
	return nil
}

/**
 * @brief Configure how the SPI bus lock enable the background operation.
 *
 * @note The lock will not try to stop the background operations, but wait for
 *       The background operations finished indicated by `spi_bus_lock_bg_resume_acquired_dev`.
 *
 * @param lock Handle to the lock to set
 * @param bg_enable The enabling function
 * @param bg_disable The disabling function, set to NULL if not required
 * @param arg Argument to pass to the enabling/disabling function.
 */
//go:linkname SpiBusLockSetBgControl C.spi_bus_lock_set_bg_control
func SpiBusLockSetBgControl(lock SpiBusLockHandleT, bg_enable BgCtrlFuncT, bg_disable BgCtrlFuncT, arg c.Pointer)

/**
 * Attach a device onto an SPI bus lock. The returning handle is used to perform
 * following requests for the attached device.
 *
 * @param lock SPI bus lock to attach
 * @param out_dev_handle Output handle corresponding to the device
 * @param flags requirement of the device, bitwise OR of SPI_BUS_LOCK_FLAG_* flags
 *
 * @return
 *  - ESP_ERR_NOT_SUPPORTED: if there's no hardware resources for new devices.
 *  - ESP_ERR_NO_MEM: if memory exhausted
 *  - ESP_OK: if success
 */
//go:linkname SpiBusLockRegisterDev C.spi_bus_lock_register_dev
func SpiBusLockRegisterDev(lock SpiBusLockHandleT, config *SpiBusLockDevConfigT, out_dev_handle *SpiBusLockDevHandleT) EspErrT

/**
 * Detach a device from its bus and free the resources used
 *
 * @param dev_handle Handle to the device.
 */
//go:linkname SpiBusLockUnregisterDev C.spi_bus_lock_unregister_dev
func SpiBusLockUnregisterDev(dev_handle SpiBusLockDevHandleT)

/**
 * @brief Get the parent bus lock of the device
 *
 * @param dev_handle Handle to the device to get bus lock
 * @return The bus lock handle
 */
//go:linkname SpiBusLockGetParent C.spi_bus_lock_get_parent
func SpiBusLockGetParent(dev_handle SpiBusLockDevHandleT) SpiBusLockHandleT

/**
 * @brief Get the device ID of a lock.
 *
 * The callers should allocate CS pins according to this ID.
 *
 * @param dev_handle Handle to the device to get ID
 * @return ID of the device
 */
//go:linkname SpiBusLockGetDevId C.spi_bus_lock_get_dev_id
func SpiBusLockGetDevId(dev_handle SpiBusLockDevHandleT) c.Int

/**
 * @brief The device request to touch bus registers. Can only be called by the acquiring processor.
 *
 * Also check if the registers has been touched by other devices.
 *
 * @param dev_handle Handle to the device to operate the registers
 * @return true if there has been other devices touching SPI registers.
 *     The caller may need to do a full-configuration. Otherwise return
 *     false.
 */
//go:linkname SpiBusLockTouch C.spi_bus_lock_touch
func SpiBusLockTouch(dev_handle SpiBusLockDevHandleT) bool

/************* Acquiring service *********************/
/**
 * Acquiring the SPI bus for exclusive use. Will also wait for the BG to finish all requests of
 * this device before it returns.
 *
 * After successfully return, the caller becomes the acquiring processor.
 *
 * @note For the main flash bus, `bg_disable` will be called to disable the cache.
 *
 * @param dev_handle Handle to the device request for acquiring.
 * @param wait Time to wait until timeout or succeed, must be `portMAX_DELAY` for now.
 * @return
 *  - ESP_OK: on success
 *  - ESP_ERR_INVALID_ARG: timeout is not portMAX_DELAY
 */
//go:linkname SpiBusLockAcquireStart C.spi_bus_lock_acquire_start
func SpiBusLockAcquireStart(dev_handle SpiBusLockDevHandleT, wait TickTypeT) EspErrT

/**
 * Release the bus acquired. Will pass the acquiring processor to other blocked
 * processors (tasks or ISR), and cause them to be unblocked or invoked.
 *
 * The acquiring device may also become NULL if no device is asking for acquiring.
 * In this case, the BG may be invoked if there is any BG requests.
 *
 * If the new acquiring device has BG requests, the BG will be invoked before the
 * task is resumed later after the BG finishes all requests of the new acquiring
 * device. Otherwise the task of the new acquiring device will be resumed immediately.
 *
 * @param dev_handle Handle to the device releasing the bus.
 * @return
 *  - ESP_OK: on success
 *  - ESP_ERR_INVALID_STATE: the device hasn't acquired the lock yet
 */
//go:linkname SpiBusLockAcquireEnd C.spi_bus_lock_acquire_end
func SpiBusLockAcquireEnd(dev_handle SpiBusLockDevHandleT) EspErrT

/**
 * Get the device acquiring the bus.
 *
 * @note Return value is not stable as the acquiring processor may change
 *       when this function is called.
 *
 * @param lock Lock of SPI bus to get the acquiring device.
 * @return The argument corresponding to the acquiring device, see
 *         `spi_bus_lock_register_dev`.
 */
//go:linkname SpiBusLockGetAcquiringDev C.spi_bus_lock_get_acquiring_dev
func SpiBusLockGetAcquiringDev(lock SpiBusLockHandleT) SpiBusLockDevHandleT

/************* BG (Background, for ISR or cache) service *********************/
/**
 * Call by a device to request a BG operation.
 *
 * Depending on the bus lock state, the BG operations may be resumed by this
 * call, or pending until BG operations allowed.
 *
 * Cleared by `spi_bus_lock_bg_clear_req` in the BG.
 *
 * @param dev_handle The device requesting BG operations.
 * @return always ESP_OK
 */
//go:linkname SpiBusLockBgRequest C.spi_bus_lock_bg_request
func SpiBusLockBgRequest(dev_handle SpiBusLockDevHandleT) EspErrT

/**
 * Wait until the ISR has finished all the BG operations for the acquiring device.
 * If any `spi_bus_lock_bg_request` for this device has been called after
 * `spi_bus_lock_acquire_start`, this function must be called before any operation
 * in the task.
 *
 * @note Can only be called when bus acquired by this device.
 *
 * @param dev_handle Handle to the device acquiring the bus.
 * @param wait Time to wait until timeout or succeed, must be `portMAX_DELAY` for now.
 * @return
 *  - ESP_OK: on success
 *  - ESP_ERR_INVALID_STATE: The device is not the acquiring bus.
 *  - ESP_ERR_INVALID_ARG: Timeout is not portMAX_DELAY.
 */
//go:linkname SpiBusLockWaitBgDone C.spi_bus_lock_wait_bg_done
func SpiBusLockWaitBgDone(dev_handle SpiBusLockDevHandleT, wait TickTypeT) EspErrT

/**
 * Handle interrupt and closure of last operation. Should be called at the beginning of the ISR,
 * when the ISR is acting as the acquiring processor.
 *
 * @param lock The SPI bus lock
 *
 * @return false if the ISR has already touched the HW, should run closure of the
 *         last operation first; otherwise true if the ISR just start operating
 *         on the HW, closure should be skipped.
 */
//go:linkname SpiBusLockBgEntry C.spi_bus_lock_bg_entry
func SpiBusLockBgEntry(lock SpiBusLockHandleT) bool

/**
 * Handle the scheduling of other acquiring devices, and control of HW operation
 * status.
 *
 * If no BG request is found, call with `wip=false`. This function will return false,
 * indicating there is incoming BG requests for the current acquiring device (or
 * for all devices if there is no acquiring device) and the ISR needs retry.
 * Otherwise may schedule a new acquiring processor (unblock the task) if there
 * is, and return true.
 *
 * Otherwise if a BG request is started in this ISR, call with `wip=true` and the
 * function will enable the interrupt to make the ISR be called again when the
 * request is done.
 *
 * This function is safe and should still be called when the ISR just lost its acquiring processor
 * role, but hasn't quit.
 *
 * @note This function will not change acquiring device. The ISR call
 *       `spi_bus_lock_bg_update_acquiring` to check for new acquiring device,
 *       when acquiring devices need to be served before other devices.
 *
 * @param lock The SPI bus lock.
 * @param wip Whether an operation is being executed when quitting the ISR.
 * @param do_yield[out] Not touched when no yielding required, otherwise set
 *                      to pdTRUE.
 * @return false if retry is required, indicating that there is pending BG request.
 *         otherwise true and quit ISR is allowed.
 */
//go:linkname SpiBusLockBgExit C.spi_bus_lock_bg_exit
func SpiBusLockBgExit(lock SpiBusLockHandleT, wip bool, do_yield *BaseTypeT) bool

/**
 * Check whether there is device asking for the acquiring device, and the desired
 * device for the next operation is also recommended.
 *
 * @note Must be called when the ISR is acting as the acquiring processor, and
 *        there is no acquiring device.
 *
 * @param lock The SPI bus lock.
 * @param out_dev_lock The recommended device for hte next operation. It's the new
 *        acquiring device when found, otherwise a device that has active BG request.
 *
 * @return true if the ISR need to quit (new acquiring device has no active BG
 *         request, or no active BG requests for all devices when there is no
 *         acquiring device), otherwise false.
 */
//go:linkname SpiBusLockBgCheckDevAcq C.spi_bus_lock_bg_check_dev_acq
func SpiBusLockBgCheckDevAcq(lock SpiBusLockHandleT, out_dev_lock *SpiBusLockDevHandleT) bool

/**
 * Check if the device has BG requests. Must be called when the ISR is acting as
 * the acquiring processor.
 *
 * @note This is not stable, may become true again when a task request for BG
 *       operation (by `spi_bus_lock_bg_request`).
 *
 * @param dev_lock The device to check.
 * @return true if the device has BG requests, otherwise false.
 */
//go:linkname SpiBusLockBgCheckDevReq C.spi_bus_lock_bg_check_dev_req
func SpiBusLockBgCheckDevReq(dev_lock SpiBusLockDevHandleT) bool

/**
 * Clear the pending BG operation request of a device after served. Must be
 * called when the ISR is acting as the acquiring processor.
 *
 * @note When the return value is true, the ISR will lost the acquiring processor role. Then
 *       `spi_bus_lock_bg_exit` must be called and checked before calling all other functions that
 *       require to be called when the ISR is the acquiring processor again.
 *
 * @param dev_handle The device whose request is served.
 * @return True if no pending requests for the acquiring device, or for all devices
 *         if there is no acquiring device. Otherwise false. When the return value is
 *         true, the ISR is no longer the acquiring processor.
 */
//go:linkname SpiBusLockBgClearReq C.spi_bus_lock_bg_clear_req
func SpiBusLockBgClearReq(dev_lock SpiBusLockDevHandleT) bool

/**
 * Check if there is any active BG requests.
 *
 * @param lock The SPI bus lock.
 * @return true if any device has active BG requst, otherwise false.
 */
//go:linkname SpiBusLockBgReqExist C.spi_bus_lock_bg_req_exist
func SpiBusLockBgReqExist(lock SpiBusLockHandleT) bool

/**
 * @brief Initialize the main flash device, called during chip startup.
 *
 * @return
 *      - ESP_OK: if success
 *      - ESP_ERR_NO_MEM: memory exhausted
 */
//go:linkname SpiBusLockInitMainDev C.spi_bus_lock_init_main_dev
func SpiBusLockInitMainDev() EspErrT
