package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_EP_DIR_MASK = 0x80
const USB_EP_DIR_IN = 0x80
const USB_EP_DIR_OUT = 0x00

type UsbDcStatusCode c.Int

const (
	USB_DC_ERROR        UsbDcStatusCode = 0
	USB_DC_RESET        UsbDcStatusCode = 1
	USB_DC_CONNECTED    UsbDcStatusCode = 2
	USB_DC_CONFIGURED   UsbDcStatusCode = 3
	USB_DC_DISCONNECTED UsbDcStatusCode = 4
	USB_DC_SUSPEND      UsbDcStatusCode = 5
	USB_DC_RESUME       UsbDcStatusCode = 6
	USB_DC_INTERFACE    UsbDcStatusCode = 7
	USB_DC_SET_HALT     UsbDcStatusCode = 8
	USB_DC_CLEAR_HALT   UsbDcStatusCode = 9
	USB_DC_UNKNOWN      UsbDcStatusCode = 10
)

type UsbDcEpCbStatusCode c.Int

const (
	USB_DC_EP_SETUP    UsbDcEpCbStatusCode = 0
	USB_DC_EP_DATA_OUT UsbDcEpCbStatusCode = 1
	USB_DC_EP_DATA_IN  UsbDcEpCbStatusCode = 2
)

type UsbDcEpType c.Int

const (
	USB_DC_EP_CONTROL     UsbDcEpType = 0
	USB_DC_EP_ISOCHRONOUS UsbDcEpType = 1
	USB_DC_EP_BULK        UsbDcEpType = 2
	USB_DC_EP_INTERRUPT   UsbDcEpType = 3
)

/**
 * USB Endpoint Configuration.
 */

type UsbDcEpCfgData struct {
	EpAddr c.Uint8T
	EpMps  c.Uint16T
	EpType UsbDcEpType
}

// llgo:type C
type UsbDcEpCallback func(c.Uint8T, UsbDcEpCbStatusCode)

// llgo:type C
type UsbDcStatusCallback func(UsbDcStatusCode, *c.Uint8T)

/**
 * @brief attach USB for device connection
 *
 * Function to attach USB for device connection. Upon success, the USB PLL
 * is enabled, and the USB device is now capable of transmitting and receiving
 * on the USB bus and of generating interrupts.
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcAttach C.usb_dc_attach
func UsbDcAttach() c.Int

/**
 * @brief detach the USB device
 *
 * Function to detach the USB device. Upon success, the USB hardware PLL
 * is powered down and USB communication is disabled.
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcDetach C.usb_dc_detach
func UsbDcDetach() c.Int

/**
 * @brief reset the USB device
 *
 * This function returns the USB device and firmware back to it's initial state.
 * N.B. the USB PLL is handled by the usb_detach function
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcReset C.usb_dc_reset
func UsbDcReset() c.Int

/**
 * @brief set USB device address
 *
 * @param[in] addr device address
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcSetAddress C.usb_dc_set_address
func UsbDcSetAddress(addr c.Uint8T) c.Int

/**
 * @brief set USB device controller status callback
 *
 * Function to set USB device controller status callback. The registered
 * callback is used to report changes in the status of the device controller.
 *
 * @param[in] cb callback function
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcSetStatusCallback C.usb_dc_set_status_callback
func UsbDcSetStatusCallback(cb UsbDcStatusCallback) c.Int

/**
 * @brief check endpoint capabilities
 *
 * Function to check capabilities of an endpoint. usb_dc_ep_cfg_data structure
 * provides the endpoint configuration parameters: endpoint address,
 * endpoint maximum packet size and endpoint type.
 * The driver should check endpoint capabilities and return 0 if the
 * endpoint configuration is possible.
 *
 * @param[in] cfg Endpoint config
 *
 * @return 0 on success, negative errno code on fail.
 */
// llgo:link (*UsbDcEpCfgData).UsbDcEpCheckCap C.usb_dc_ep_check_cap
func (recv_ *UsbDcEpCfgData) UsbDcEpCheckCap() c.Int {
	return 0
}

/**
 * @brief configure endpoint
 *
 * Function to configure an endpoint. usb_dc_ep_cfg_data structure provides
 * the endpoint configuration parameters: endpoint address, endpoint maximum
 * packet size and endpoint type.
 *
 * @param[in] cfg Endpoint config
 *
 * @return 0 on success, negative errno code on fail.
 */
// llgo:link (*UsbDcEpCfgData).UsbDcEpConfigure C.usb_dc_ep_configure
func (recv_ *UsbDcEpCfgData) UsbDcEpConfigure() c.Int {
	return 0
}

/**
 * @brief set stall condition for the selected endpoint
 *
 * @param[in] ep Endpoint address corresponding to the one
 *               listed in the device configuration table
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpSetStall C.usb_dc_ep_set_stall
func UsbDcEpSetStall(ep c.Uint8T) c.Int

/**
 * @brief clear stall condition for the selected endpoint
 *
 * @param[in] ep Endpoint address corresponding to the one
 *               listed in the device configuration table
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpClearStall C.usb_dc_ep_clear_stall
func UsbDcEpClearStall(ep c.Uint8T) c.Int

/**
 * @brief check if selected endpoint is stalled
 *
 * @param[in]  ep       Endpoint address corresponding to the one
 *                      listed in the device configuration table
 * @param[out] stalled  Endpoint stall status
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpIsStalled C.usb_dc_ep_is_stalled
func UsbDcEpIsStalled(ep c.Uint8T, stalled *c.Uint8T) c.Int

/**
 * @brief halt the selected endpoint
 *
 * @param[in] ep Endpoint address corresponding to the one
 *               listed in the device configuration table
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpHalt C.usb_dc_ep_halt
func UsbDcEpHalt(ep c.Uint8T) c.Int

/**
 * @brief enable the selected endpoint
 *
 * Function to enable the selected endpoint. Upon success interrupts are
 * enabled for the corresponding endpoint and the endpoint is ready for
 * transmitting/receiving data.
 *
 * @param[in] ep Endpoint address corresponding to the one
 *               listed in the device configuration table
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpEnable C.usb_dc_ep_enable
func UsbDcEpEnable(ep c.Uint8T) c.Int

/**
 * @brief disable the selected endpoint
 *
 * Function to disable the selected endpoint. Upon success interrupts are
 * disabled for the corresponding endpoint and the endpoint is no longer able
 * for transmitting/receiving data.
 *
 * @param[in] ep Endpoint address corresponding to the one
 *               listed in the device configuration table
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpDisable C.usb_dc_ep_disable
func UsbDcEpDisable(ep c.Uint8T) c.Int

/**
 * @brief flush the selected endpoint
 *
 * @param[in] ep Endpoint address corresponding to the one
 *               listed in the device configuration table
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpFlush C.usb_dc_ep_flush
func UsbDcEpFlush(ep c.Uint8T) c.Int

/**
 * @brief write data to the specified endpoint
 *
 * This function is called to write data to the specified endpoint. The supplied
 * usb_ep_callback function will be called when data is transmitted out.
 *
 * @param[in]  ep        Endpoint address corresponding to the one
 *                       listed in the device configuration table
 * @param[in]  data      pointer to data to write
 * @param[in]  data_len  length of data requested to write. This may
 *                       be zero for a zero length status packet.
 * @param[out] ret_bytes bytes scheduled for transmission. This value
 *                       may be NULL if the application expects all
 *                       bytes to be written
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpWrite C.usb_dc_ep_write
func UsbDcEpWrite(ep c.Uint8T, data *c.Uint8T, data_len c.Uint32T, ret_bytes *c.Uint32T) c.Int

/**
 * @brief Indicate if the write to an IN endpoint (using usb_dc_ep_write) would block
 * to wait until the endpoint has enoug space
 *
 * @param[in]  ep        Endpoint address corresponding to the one
 *                       listed in the device configuration table
 *
 * @return 0 when writable, 0 when not, negative errno code on fail.
 */
//go:linkname UsbDcEpWriteWouldBlock C.usb_dc_ep_write_would_block
func UsbDcEpWriteWouldBlock(ep c.Uint8T) c.Int

/**
 * @brief read data from the specified endpoint
 *
 * This function is called by the Endpoint handler function, after an OUT
 * interrupt has been received for that EP. The application must only call this
 * function through the supplied usb_ep_callback function. This function clears
 * the ENDPOINT NAK, if all data in the endpoint FIFO has been read,
 * so as to accept more data from host.
 *
 * @param[in]  ep           Endpoint address corresponding to the one
 *                          listed in the device configuration table
 * @param[in]  data         pointer to data buffer to write to
 * @param[in]  max_data_len max length of data to read
 * @param[out] read_bytes   Number of bytes read. If data is NULL and
 *                          max_data_len is 0 the number of bytes
 *                          available for read should be returned.
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpRead C.usb_dc_ep_read
func UsbDcEpRead(ep c.Uint8T, data *c.Uint8T, max_data_len c.Uint32T, read_bytes *c.Uint32T) c.Int

/**
 * @brief set callback function for the specified endpoint
 *
 * Function to set callback function for notification of data received and
 * available to application or transmit done on the selected endpoint,
 * NULL if callback not required by application code.
 *
 * @param[in] ep Endpoint address corresponding to the one
 *               listed in the device configuration table
 * @param[in] cb callback function
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpSetCallback C.usb_dc_ep_set_callback
func UsbDcEpSetCallback(ep c.Uint8T, cb UsbDcEpCallback) c.Int

/**
 * @brief read data from the specified endpoint
 *
 * This is similar to usb_dc_ep_read, the difference being that, it doesn't
 * clear the endpoint NAKs so that the consumer is not bogged down by further
 * upcalls till he is done with the processing of the data. The caller should
 * reactivate ep by invoking usb_dc_ep_read_continue() do so.
 *
 * @param[in]  ep           Endpoint address corresponding to the one
 *                          listed in the device configuration table
 * @param[in]  data         pointer to data buffer to write to
 * @param[in]  max_data_len max length of data to read
 * @param[out] read_bytes   Number of bytes read. If data is NULL and
 *                          max_data_len is 0 the number of bytes
 *                          available for read should be returned.
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpReadWait C.usb_dc_ep_read_wait
func UsbDcEpReadWait(ep c.Uint8T, data *c.Uint8T, max_data_len c.Uint32T, read_bytes *c.Uint32T) c.Int

/**
 * @brief Continue reading data from the endpoint
 *
 * Clear the endpoint NAK and enable the endpoint to accept more data
 * from the host. Usually called after usb_dc_ep_read_wait() when the consumer
 * is fine to accept more data. Thus these calls together acts as flow control
 * mechanism.
 *
 * @param[in]  ep           Endpoint address corresponding to the one
 *                          listed in the device configuration table
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbDcEpReadContinue C.usb_dc_ep_read_continue
func UsbDcEpReadContinue(ep c.Uint8T) c.Int

/**
 * @brief Get endpoint max packet size
 *
 * @param[in]  ep           Endpoint address corresponding to the one
 *                          listed in the device configuration table
 *
 * @return endpoint max packet size (mps)
 */
//go:linkname UsbDcEpMps C.usb_dc_ep_mps
func UsbDcEpMps(ep c.Uint8T) c.Int

/**
 * @brief Poll for interrupts that need to be handled
 *
 * When the USB interrupt is not hooked up to an actual CPU interrupt, you
 * can call this periodically to handle the USB events that need handling.
 */
//go:linkname UsbDcCheckPollForInterrupts C.usb_dc_check_poll_for_interrupts
func UsbDcCheckPollForInterrupts()

/*
 * @brief Prepare for USB persist
 *
 * This takes the USB peripheral offline in such a way that it seems 'just busy' to the
 * host. This way, the chip can reboot (e.g. into bootloader mode) and pick up the USB
 * configuration again, without the connection to the host being interrupted.
 *
 * @note Actual persistence is depending on USBDC_PERSIST_ENA being set in flags, as this
 * is also used to e.g. reboot into DFU mode.
 *
 * @note Please reboot soon after calling this.
 */
//go:linkname UsbDcPreparePersist C.usb_dc_prepare_persist
func UsbDcPreparePersist() c.Int

/*
 * @brief USB interrupt handler
 *
 * This can be hooked up by the OS to the USB peripheral interrupt.
 */
//go:linkname UsbDwIsrHandler C.usb_dw_isr_handler
func UsbDwIsrHandler()

/**
 * @brief Provide IDF with an interface to clear the static variable usb_dw_ctrl
 *
 *
 */
//go:linkname UsbDwCtrlDeinit C.usb_dw_ctrl_deinit
func UsbDwCtrlDeinit()
