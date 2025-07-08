package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MAX_PACKET_SIZE0 = 64

/** setup packet definitions */

type UsbSetupPacket struct {
	BmRequestType c.Uint8T
	BRequest      c.Uint8T
	WValue        c.Uint16T
	WIndex        c.Uint16T
	WLength       c.Uint16T
}

// llgo:type C
type UsbStatusCallback func(UsbDcStatusCode, *c.Uint8T)

// llgo:type C
type UsbEpCallback func(c.Uint8T, UsbDcEpCbStatusCode)

// llgo:type C
type UsbRequestHandler func(*UsbSetupPacket, *c.Int32T, **c.Uint8T) c.Int

// llgo:type C
type UsbInterfaceConfig func(c.Uint8T)

/*
 * USB Endpoint Configuration
 */

type UsbEpCfgData struct {
	EpCb   UsbEpCallback
	EpAddr c.Uint8T
}

/**
 * USB Interface Configuration
 */

type UsbInterfaceCfgData struct {
	ClassHandler  UsbRequestHandler
	VendorHandler UsbRequestHandler
	CustomHandler UsbRequestHandler
	PayloadData   *c.Uint8T
	VendorData    *c.Uint8T
}

/*
 * @brief USB device configuration
 *
 * The Application instantiates this with given parameters added
 * using the "usb_set_config" function. Once this function is called
 * changes to this structure will result in undefined behaviour. This structure
 * may only be updated after calls to usb_deconfig
 */

type UsbCfgData struct {
	UsbDeviceDescription *c.Uint8T
	InterfaceDescriptor  c.Pointer
	InterfaceConfig      UsbInterfaceConfig
	CbUsbStatus          UsbStatusCallback
	Interface            UsbInterfaceCfgData
	NumEndpoints         c.Uint8T
	Endpoint             *UsbEpCfgData
}

/*
 * @brief configure USB controller
 *
 * Function to configure USB controller.
 * Configuration parameters must be valid or an error is returned
 *
 * @param[in] config Pointer to configuration structure
 *
 * @return 0 on success, negative errno code on fail
 */
// llgo:link (*UsbCfgData).UsbSetConfig C.usb_set_config
func (recv_ *UsbCfgData) UsbSetConfig() c.Int {
	return 0
}

/*
 * @brief return the USB device to it's initial state
 *
 * @return 0 on success, negative errno code on fail
 */
//go:linkname UsbDeconfig C.usb_deconfig
func UsbDeconfig() c.Int

/*
 * @brief enable USB for host/device connection
 *
 * Function to enable USB for host/device connection.
 * Upon success, the USB module is no longer clock gated in hardware,
 * it is now capable of transmitting and receiving on the USB bus and
 * of generating interrupts.
 *
 * @return 0 on success, negative errno code on fail.
 */
// llgo:link (*UsbCfgData).UsbEnable C.usb_enable
func (recv_ *UsbCfgData) UsbEnable() c.Int {
	return 0
}

/*
 * @brief disable the USB device.
 *
 * Function to disable the USB device.
 * Upon success, the specified USB interface is clock gated in hardware,
 * it is no longer capable of generating interrupts.
 *
 * @return 0 on success, negative errno code on fail
 */
//go:linkname UsbDisable C.usb_disable
func UsbDisable() c.Int

/*
 * @brief Check if a write to an in ep would block until there is enough space
 * in the fifo
 *
 * @param[in]  ep        Endpoint address corresponding to the one listed in the
 *                       device configuration table
 *
 * @return 0 if free to write, 1 if a write would block, negative errno code on fail
 */
//go:linkname UsbWriteWouldBlock C.usb_write_would_block
func UsbWriteWouldBlock(ep c.Uint8T) c.Int

/*
 * @brief write data to the specified endpoint
 *
 * Function to write data to the specified endpoint. The supplied
 * usb_ep_callback will be called when transmission is done.
 *
 * @param[in]  ep        Endpoint address corresponding to the one listed in the
 *                       device configuration table
 * @param[in]  data      Pointer to data to write
 * @param[in]  data_len  Length of data requested to write. This may be zero for
 *                       a zero length status packet.
 * @param[out] bytes_ret Bytes written to the EP FIFO. This value may be NULL if
 *                       the application expects all bytes to be written
 *
 * @return 0 on success, negative errno code on fail
 */
//go:linkname UsbWrite C.usb_write
func UsbWrite(ep c.Uint8T, data *c.Uint8T, data_len c.Uint32T, bytes_ret *c.Uint32T) c.Int

/*
 * @brief read data from the specified endpoint
 *
 * This function is called by the Endpoint handler function, after an
 * OUT interrupt has been received for that EP. The application must
 * only call this function through the supplied usb_ep_callback function.
 *
 * @param[in]  ep           Endpoint address corresponding to the one listed in
 *                          the device configuration table
 * @param[in]  data         Pointer to data buffer to write to
 * @param[in]  max_data_len Max length of data to read
 * @param[out] ret_bytes    Number of bytes read. If data is NULL and
 *                          max_data_len is 0 the number of bytes available
 *                          for read is returned.
 *
 * @return  0 on success, negative errno code on fail
 */
//go:linkname UsbRead C.usb_read
func UsbRead(ep c.Uint8T, data *c.Uint8T, max_data_len c.Uint32T, ret_bytes *c.Uint32T) c.Int

/*
 * @brief set STALL condition on the specified endpoint
 *
 * This function is called by USB device class handler code to set stall
 * condition on endpoint.
 *
 * @param[in]  ep           Endpoint address corresponding to the one listed in
 *                          the device configuration table
 *
 * @return  0 on success, negative errno code on fail
 */
//go:linkname UsbEpSetStall C.usb_ep_set_stall
func UsbEpSetStall(ep c.Uint8T) c.Int

/*
 * @brief clears STALL condition on the specified endpoint
 *
 * This function is called by USB device class handler code to clear stall
 * condition on endpoint.
 *
 * @param[in]  ep           Endpoint address corresponding to the one listed in
 *                          the device configuration table
 *
 * @return  0 on success, negative errno code on fail
 */
//go:linkname UsbEpClearStall C.usb_ep_clear_stall
func UsbEpClearStall(ep c.Uint8T) c.Int

/**
 * @brief read data from the specified endpoint
 *
 * This is similar to usb_ep_read, the difference being that, it doesn't
 * clear the endpoint NAKs so that the consumer is not bogged down by further
 * upcalls till he is done with the processing of the data. The caller should
 * reactivate ep by invoking usb_ep_read_continue() do so.
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
//go:linkname UsbEpReadWait C.usb_ep_read_wait
func UsbEpReadWait(ep c.Uint8T, data *c.Uint8T, max_data_len c.Uint32T, read_bytes *c.Uint32T) c.Int

/**
 * @brief Continue reading data from the endpoint
 *
 * Clear the endpoint NAK and enable the endpoint to accept more data
 * from the host. Usually called after usb_ep_read_wait() when the consumer
 * is fine to accept more data. Thus these calls together acts as flow control
 * mechanism.
 *
 * @param[in]  ep           Endpoint address corresponding to the one
 *                          listed in the device configuration table
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbEpReadContinue C.usb_ep_read_continue
func UsbEpReadContinue(ep c.Uint8T) c.Int

// llgo:type C
type UsbTransferCallback func(c.Uint8T, c.Int, c.Pointer)

/**
 * @brief Transfer management endpoint callback
 *
 * If a USB class driver wants to use high-level transfer functions, driver
 * needs to register this callback as usb endpoint callback.
 */
//go:linkname UsbTransferEpCallback C.usb_transfer_ep_callback
func UsbTransferEpCallback(ep c.Uint8T, __llgo_arg_1 UsbDcEpCbStatusCode)

/**
 * @brief Start a transfer
 *
 * Start a usb transfer to/from the data buffer. This function is asynchronous
 * and can be executed in IRQ context. The provided callback will be called
 * on transfer completion (or error) in thread context.
 *
 * @param[in]  ep           Endpoint address corresponding to the one
 *                          listed in the device configuration table
 * @param[in]  data         Pointer to data buffer to write-to/read-from
 * @param[in]  dlen         Size of data buffer
 * @param[in]  flags        Transfer flags (USB_TRANS_READ, USB_TRANS_WRITE...)
 * @param[in]  cb           Function called on transfer completion/failure
 * @param[in]  priv         Data passed back to the transfer completion callback
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbTransfer C.usb_transfer
func UsbTransfer(ep c.Uint8T, data *c.Uint8T, dlen c.SizeT, flags c.Uint, cb UsbTransferCallback, priv c.Pointer) c.Int

/**
 * @brief Start a transfer and block-wait for completion
 *
 * Synchronous version of usb_transfer, wait for transfer completion before
 * returning.
 *
 * @param[in]  ep           Endpoint address corresponding to the one
 *                          listed in the device configuration table
 * @param[in]  data         Pointer to data buffer to write-to/read-from
 * @param[in]  dlen         Size of data buffer
 * @param[in]  flags        Transfer flags

 *
 * @return number of bytes transferred on success, negative errno code on fail.
 */
//go:linkname UsbTransferSync C.usb_transfer_sync
func UsbTransferSync(ep c.Uint8T, data *c.Uint8T, dlen c.SizeT, flags c.Uint) c.Int

/**
 * @brief Cancel any ongoing transfer on the specified endpoint
 *
 * @param[in]  ep           Endpoint address corresponding to the one
 *                          listed in the device configuration table
 *
 * @return 0 on success, negative errno code on fail.
 */
//go:linkname UsbCancelTransfer C.usb_cancel_transfer
func UsbCancelTransfer(ep c.Uint8T)

/**
 * @brief Provide IDF with an interface to clear the static variable usb_dev
 *
 *
 */
//go:linkname UsbDevDeinit C.usb_dev_deinit
func UsbDevDeinit()

//go:linkname UsbDevResume C.usb_dev_resume
func UsbDevResume(configuration c.Int)

//go:linkname UsbDevGetConfiguration C.usb_dev_get_configuration
func UsbDevGetConfiguration() c.Int
