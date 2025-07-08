package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_DWC_HAL_XFER_DESC_FLAG_IN = 0x01
const USB_DWC_HAL_XFER_DESC_FLAG_SETUP = 0x02
const USB_DWC_HAL_XFER_DESC_FLAG_HOC = 0x04

type UsbHalFifoBiasT c.Int

const (
	USB_HAL_FIFO_BIAS_DEFAULT UsbHalFifoBiasT = 0
	USB_HAL_FIFO_BIAS_RX      UsbHalFifoBiasT = 1
	USB_HAL_FIFO_BIAS_PTX     UsbHalFifoBiasT = 2
)

/**
 * @brief MPS limits based on FIFO configuration
 *
 * In bytes
 *
 * The resulting values depend on
 * 1. FIFO total size (chip specific)
 * 2. Set FIFO bias
 */

type UsbHalFifoMpsLimitsT struct {
	InMps             c.Uint
	NonPeriodicOutMps c.Uint
	PeriodicOutMps    c.Uint
}

/**
 * @brief FIFO size configuration structure
 */

type UsbDwcHalFifoConfigT struct {
	RxFifoLines   c.Uint32T
	NptxFifoLines c.Uint32T
	PtxFifoLines  c.Uint32T
}
type UsbDwcHalPortEventT c.Int

const (
	USB_DWC_HAL_PORT_EVENT_NONE       UsbDwcHalPortEventT = 0
	USB_DWC_HAL_PORT_EVENT_CHAN       UsbDwcHalPortEventT = 1
	USB_DWC_HAL_PORT_EVENT_CONN       UsbDwcHalPortEventT = 2
	USB_DWC_HAL_PORT_EVENT_DISCONN    UsbDwcHalPortEventT = 3
	USB_DWC_HAL_PORT_EVENT_ENABLED    UsbDwcHalPortEventT = 4
	USB_DWC_HAL_PORT_EVENT_DISABLED   UsbDwcHalPortEventT = 5
	USB_DWC_HAL_PORT_EVENT_OVRCUR     UsbDwcHalPortEventT = 6
	USB_DWC_HAL_PORT_EVENT_OVRCUR_CLR UsbDwcHalPortEventT = 7
)

type UsbDwcHalChanEventT c.Int

const (
	USB_DWC_HAL_CHAN_EVENT_CPLT     UsbDwcHalChanEventT = 0
	USB_DWC_HAL_CHAN_EVENT_ERROR    UsbDwcHalChanEventT = 1
	USB_DWC_HAL_CHAN_EVENT_HALT_REQ UsbDwcHalChanEventT = 2
	USB_DWC_HAL_CHAN_EVENT_NONE     UsbDwcHalChanEventT = 3
)

type UsbDwcHalChanErrorT c.Int

const (
	USB_DWC_HAL_CHAN_ERROR_XCS_XACT UsbDwcHalChanErrorT = 0
	USB_DWC_HAL_CHAN_ERROR_BNA      UsbDwcHalChanErrorT = 1
	USB_DWC_HAL_CHAN_ERROR_PKT_BBL  UsbDwcHalChanErrorT = 2
	USB_DWC_HAL_CHAN_ERROR_STALL    UsbDwcHalChanErrorT = 3
)

/**
 * @brief Endpoint characteristics structure
 */

type UsbDwcHalEpCharT struct {
	Periodic struct {
		Interval c.Uint
		Offset   c.Uint32T
		IsHs     bool
	}
}

/**
 * @brief Channel object
 */

type UsbDwcHalChanT struct {
	Flags struct {
		Val c.Uint32T
	}
	Regs    *UsbDwcHostChanRegsT
	Error   UsbDwcHalChanErrorT
	Type    UsbDwcXferTypeT
	ChanCtx c.Pointer
}

/**
 * @brief HAL context structure
 */

type UsbDwcHalContextT struct {
	Dev               *UsbDwcDevT
	PeriodicFrameList *c.Uint32T
	FrameListLen      UsbHalFrameListLenT
	FifoConfig        UsbDwcHalFifoConfigT
	ConstantConfig    struct {
		ChanNumTotal c.Uint
		HsphyType    c.Uint
		FifoSize     c.Uint
	}
	Flags struct {
		Val c.Uint32T
	}
	Channels struct {
		NumAllocated     c.Int
		ChanPendIntrsMsk c.Uint32T
		Hdls             **UsbDwcHalChanT
	}
}

/**
 * @brief Initialize the HAL context and check if DWC_OTG is alive
 *
 * Entry:
 * - The peripheral must have been reset and clock un-gated
 * - The USB PHY (internal or external) and associated GPIOs must already be configured
 * - GPIO pins configured
 * - Interrupt allocated but DISABLED (in case of an unknown interrupt state)
 * Exit:
 * - Checks to see if DWC_OTG is alive, and if HW version/config is correct
 * - HAL context initialized
 * - Read and save relevant USB-DWC configuration parameters
 * - Sets default values to some global and OTG registers (GAHBCFG and GUSBCFG)
 * - Umask global interrupt signal
 * - Put DWC_OTG into host mode. Require 25ms delay before this takes effect.
 * - State -> USB_DWC_HAL_PORT_STATE_OTG
 * - Interrupts cleared. Users can now enable their ISR
 *
 * @attention The user must allocate memory for channel handlers with
 *            `hal->channels.hdls = malloc(hal->constant_config.chan_num_total * sizeof(usb_dwc_hal_chan_t*))`
 * @param[inout] hal     Context of the HAL layer
 * @param[in]    port_id USB port ID
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalInit C.usb_dwc_hal_init
func (recv_ *UsbDwcHalContextT) UsbDwcHalInit(port_id c.Int) {
}

/**
 * @brief Deinitialize the HAL context
 *
 * Entry:
 * - All channels must be properly disabled, and any pending events handled
 * Exit:
 * - DWC_OTG global interrupt disabled
 * - HAL context deinitialized
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalDeinit C.usb_dwc_hal_deinit
func (recv_ *UsbDwcHalContextT) UsbDwcHalDeinit() {
}

/**
 * @brief Issue a soft reset to the controller
 *
 * This should be called when the host port encounters an error event or has been disconnected. Before calling this,
 * users are responsible for safely freeing all channels as a soft reset will wipe all host port and channel registers.
 * This function will result in the host port being put back into same state as after calling usb_dwc_hal_init().
 *
 * @note This has nothing to do with a USB bus reset. It simply resets the peripheral
 *
 * @param[in] hal Context of the HAL layer
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalCoreSoftReset C.usb_dwc_hal_core_soft_reset
func (recv_ *UsbDwcHalContextT) UsbDwcHalCoreSoftReset() {
}

/**
 * @brief Set FIFO bias
 *
 * This function will set the sizes of each of the FIFOs (RX FIFO, Non-periodic TX FIFO, Periodic TX FIFO) and must be
 * called at least once before allocating the channel. Based on the type of endpoints (and the endpoints' MPS), there
 * may be situations where this function may need to be called again to resize the FIFOs. If resizing FIFOs dynamically,
 * it is the user's responsibility to ensure there are no active channels when this function is called.
 *
 * @note After a port reset, the FIFO size registers will reset to their default values, so this function must be called
 *       again post reset.
 *
 * @param[in] hal       Context of the HAL layer
 * @param[in] fifo_bias FIFO bias configuration
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalSetFifoBias C.usb_dwc_hal_set_fifo_bias
func (recv_ *UsbDwcHalContextT) UsbDwcHalSetFifoBias(fifo_bias UsbHalFifoBiasT) {
}

/**
 * @brief Get MPS limits
 *
 * @param[in]  hal        Context of the HAL layer
 * @param[out] mps_limits MPS limits
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalGetMpsLimits C.usb_dwc_hal_get_mps_limits
func (recv_ *UsbDwcHalContextT) UsbDwcHalGetMpsLimits(mps_limits *UsbHalFifoMpsLimitsT) {
}

/**
 * @brief Enable the host port
 *
 * Entry:
 * - Host port enabled event triggered following a reset
 * Exit:
 * - Host port enabled to operate in scatter/gather DMA mode
 * - DMA fifo sizes configured
 *
 * @param hal Context of the HAL layer
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalPortEnable C.usb_dwc_hal_port_enable
func (recv_ *UsbDwcHalContextT) UsbDwcHalPortEnable() {
}

/**
 * @brief Allocate a channel
 *
 * @param[in] hal Context of the HAL layer
 * @param[inout] chan_obj Empty channel object
 * @param[in] chan_ctx Context variable for the allocator of the channel
 * @return true Channel successfully allocated
 * @return false Failed to allocate channel
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalChanAlloc C.usb_dwc_hal_chan_alloc
func (recv_ *UsbDwcHalContextT) UsbDwcHalChanAlloc(chan_obj *UsbDwcHalChanT, chan_ctx c.Pointer) bool {
	return false
}

/**
 * @brief Free a channel
 *
 * @param[in] hal Context of the HAL layer
 * @param[in] chan_obj Channel object
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalChanFree C.usb_dwc_hal_chan_free
func (recv_ *UsbDwcHalContextT) UsbDwcHalChanFree(chan_obj *UsbDwcHalChanT) {
}

/**
 * @brief Set the endpoint information for a particular channel
 *
 * This should be called when a channel switches target from one EP to another
 *
 * @note the channel must be in the disabled state in order to change its EP
 *       information
 *
 * @param hal Context of the HAL layer
 * @param chan_obj Channel object
 * @param ep_char Endpoint characteristics
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalChanSetEpChar C.usb_dwc_hal_chan_set_ep_char
func (recv_ *UsbDwcHalContextT) UsbDwcHalChanSetEpChar(chan_obj *UsbDwcHalChanT, ep_char *UsbDwcHalEpCharT) {
}

/**
 * @brief Activate a channel
 *
 * Activating a channel will cause the channel to start executing transfer descriptors.
 *
 * @note This function should only be called on channels that were previously halted
 * @note An event will be generated when the channel is halted
 *
 * @param chan_obj Channel object
 * @param xfer_desc_list A filled transfer descriptor list
 * @param desc_list_len Transfer descriptor list length
 * @param start_idx Index of the starting transfer descriptor in the list
 */
// llgo:link (*UsbDwcHalChanT).UsbDwcHalChanActivate C.usb_dwc_hal_chan_activate
func (recv_ *UsbDwcHalChanT) UsbDwcHalChanActivate(xfer_desc_list c.Pointer, desc_list_len c.Int, start_idx c.Int) {
}

/**
 * @brief Request to halt a channel
 *
 * This function should be called in order to halt a channel. If the channel is already halted, this function will
 * return true. If the channel is still active, this function will return false and users must wait for the
 * USB_DWC_HAL_CHAN_EVENT_HALT_REQ event before treating the channel as halted.
 *
 * @note When a transfer is in progress (i.e., the channel is active) and a halt is requested, the channel will halt
 *       after the next USB packet is completed. If the transfer has more pending packets, the transfer will just be
 *       marked as USB_DWC_HAL_XFER_DESC_STS_NOT_EXECUTED.
 *
 * @param chan_obj Channel object
 * @return true The channel is already halted
 * @return false The halt was requested, wait for USB_DWC_HAL_CHAN_EVENT_HALT_REQ
 */
// llgo:link (*UsbDwcHalChanT).UsbDwcHalChanRequestHalt C.usb_dwc_hal_chan_request_halt
func (recv_ *UsbDwcHalChanT) UsbDwcHalChanRequestHalt() bool {
	return false
}

/**
 * @brief Decode global and host port interrupts
 *
 * - Reads and clears global and host port interrupt registers
 * - Decodes the interrupt bits to determine what host port event occurred
 *
 * @note This should be the first interrupt decode function to be run
 *
 * @param hal Context of the HAL layer
 * @return usb_dwc_hal_port_event_t Host port event
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalDecodeIntr C.usb_dwc_hal_decode_intr
func (recv_ *UsbDwcHalContextT) UsbDwcHalDecodeIntr() UsbDwcHalPortEventT {
	return 0
}

/**
 * @brief Gets the next channel with a pending interrupt
 *
 * If no channel is pending an interrupt, this function will return NULL. If one or more channels are pending an
 * interrupt, this function returns one of the channel's objects. Call this function repeatedly until it returns NULL.
 *
 * @param hal Context of the HAL layer
 * @return usb_dwc_hal_chan_t* Channel object. NULL if no channel are pending an interrupt.
 */
// llgo:link (*UsbDwcHalContextT).UsbDwcHalGetChanPendingIntr C.usb_dwc_hal_get_chan_pending_intr
func (recv_ *UsbDwcHalContextT) UsbDwcHalGetChanPendingIntr() *UsbDwcHalChanT {
	return nil
}

/**
 * @brief Decode a particular channel's interrupt
 *
 * - Reads and clears the interrupt register of the channel
 * - Returns the corresponding event for that channel
 *
 * @param chan_obj Channel object
 * @note If the host port has an error (e.g., a sudden disconnect or an port error), any active channels will not
 *       receive an interrupt. Each active channel must be manually halted.
 * @return usb_dwc_hal_chan_event_t Channel event
 */
// llgo:link (*UsbDwcHalChanT).UsbDwcHalChanDecodeIntr C.usb_dwc_hal_chan_decode_intr
func (recv_ *UsbDwcHalChanT) UsbDwcHalChanDecodeIntr() UsbDwcHalChanEventT {
	return 0
}
