package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type Dma2dGroupT struct {
	Unused [8]uint8
}
type Dma2dPoolHandleT *Dma2dGroupT

type Dma2dTransS struct {
	Unused [8]uint8
}
type Dma2dTransT Dma2dTransS

/**
 * @brief Get the size for struct `dma2d_trans_t`
 *
 * @return size_t Size of struct `dma2d_trans_t`
 */
//go:linkname Dma2dGetTransElmSize C.dma2d_get_trans_elm_size
func Dma2dGetTransElmSize() c.SizeT

/**
 * @brief A collection of configuration items that used for allocating a 2D-DMA pool
 */

type Dma2dPoolConfigT struct {
	PoolId       c.Uint32T
	IntrPriority c.Uint32T
}

/**
 * @brief Acquire a 2D-DMA pool
 *
 * @param[in] config Pointer to a collection of configurations for the 2D-DMA pool
 * @param[out] ret_pool Returned pool handle
 * @return
 *      - ESP_OK: Acquire the 2D-DMA pool successfully
 *      - ESP_ERR_INVALID_ARG: Acquire the 2D-DMA pool failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Acquire the 2D-DMA pool failed because out of memory
 *      - ESP_FAIL: Acquire the 2D-DMA pool failed because of other error
 */
// llgo:link (*Dma2dPoolConfigT).Dma2dAcquirePool C.dma2d_acquire_pool
func (recv_ *Dma2dPoolConfigT) Dma2dAcquirePool(ret_pool *Dma2dPoolHandleT) EspErrT {
	return 0
}

/**
 * @brief Release a 2D-DMA pool
 *
 * @warning Upper driver should make sure there is no pending transaction (enqueued by the driver, but haven't be
 *          processed) before calling this function.
 *
 * @param[in] dma2d_pool 2D-DMA pool handle, allocated by `dma2d_acquire_pool`
 * @return
 *      - ESP_OK: Release the 2D-DMA pool successfully
 *      - ESP_ERR_INVALID_ARG: Release the 2D-DMA pool failed because of invalid argument
 *      - ESP_ERR_NOT_ALLOWED: Release the 2D-DMA pool failed because there is pending transactions in the pool,
 *                             pool can not be destroyed
 */
//go:linkname Dma2dReleasePool C.dma2d_release_pool
func Dma2dReleasePool(dma2d_pool Dma2dPoolHandleT) EspErrT

type Dma2dChannelT struct {
	Unused [8]uint8
}
type Dma2dChannelHandleT *Dma2dChannelT

/**
 * @brief Struct to save the necessary information of a 2D-DMA channel for upper drivers to configure the channels
 */

type Dma2dTransChannelInfoT struct {
	Dir  Dma2dChannelDirectionT
	Chan Dma2dChannelHandleT
}

// llgo:type C
type Dma2dTransOnPickedCallbackT func(c.Uint32T, *Dma2dTransChannelInfoT, c.Pointer) bool

/**
 * @brief A collection of configuration items for a 2D-DMA transaction
 */

type Dma2dTransConfigT struct {
	TxChannelNum           c.Uint32T
	RxChannelNum           c.Uint32T
	ChannelFlags           c.Uint32T
	SpecifiedTxChannelMask c.Uint32T
	SpecifiedRxChannelMask c.Uint32T
	OnJobPicked            Dma2dTransOnPickedCallbackT
	UserConfig             c.Pointer
}

/**
 * @brief Enqueue a 2D-DMA transaction to be picked up by a certain 2D-DMA pool
 *
 * @param[in] dma2d_pool 2D-DMA pool handle, allocated by `dma2d_acquire_pool`
 * @param[in] trans_desc Pointer to a collection of configurations for a transaction
 *                       The context must exist at least until `on_job_picked` callback function is called.
 * @param[in] trans_placeholder Address to the memory for storing this transaction context
 *                              Caller must malloc a placeholder for storing the 2D-DMA transaction, and pass it into the function.
 *                              Size of the placeholder can be get from `SIZEOF_DMA2D_TRANS_T` macro.
 *                              Freeing the 2D-DMA transaction placeholder should also be taken care by the upper driver.
 *                              It can be freed when `on_job_picked` callback function is called or anytime later.
 * @return
 *      - ESP_OK: Enqueue the 2D-DMA transaction successfully
 *      - ESP_ERR_INVALID_ARG: Enqueue the 2D-DMA transaction failed because of invalid argument
 */
//go:linkname Dma2dEnqueue C.dma2d_enqueue
func Dma2dEnqueue(dma2d_pool Dma2dPoolHandleT, trans_desc *Dma2dTransConfigT, trans_placeholder *Dma2dTransT) EspErrT

/**
 * @brief Force end an in-flight 2D-DMA transaction
 *
 * This API is useful when the error was caused by the DMA consumer (such as JPEG). The error can only be detected
 * by the consumer module, and the error info will only be propagated to the consumer driver. The 2D-DMA channels being
 * involved to transfer the data has no way to be informed about the error at its upstream, it will keep waiting for
 * the data.
 *
 * Therefore, when the consumer driver is doing the error handling, it is required to call this API to end the on-going
 * transaction and release the taken TX and RX channels. It will stop and free the TX and RX channels that are bundled
 * together to process the transaction.
 *
 * @param[in] trans Pointer to the 2D-DMA transaction context
 * @param[out] need_yield Pointer to a status flag to record whether a task switch is needed if this API is being called in an ISR
 * @return
 *      - ESP_OK: Force end an in-flight transaction successfully
 *      - ESP_ERR_INVALID_ARG: Force end failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Force end failed because the transaction is not yet in-flight
 */
// llgo:link (*Dma2dTransT).Dma2dForceEnd C.dma2d_force_end
func (recv_ *Dma2dTransT) Dma2dForceEnd(need_yield *bool) EspErrT {
	return 0
}

/**
 * @brief Type of 2D-DMA engine trigger
 */

type Dma2dTriggerT struct {
	Periph      Dma2dTriggerPeripheralT
	PeriphSelId c.Int
}

/**
 * @brief Connect 2D-DMA channel to trigger peripheral, and configure all other channel settings to a certain state (the channel will be reset first)
 *
 * Usually only to be called in `on_job_picked` callback, and is the first step to do inside the callback, since it resets other configurations to a default mode.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle, get from the `on_job_picked` callback input argument `dma2d_chans`
 * @param[in] trig_periph 2D-DMA trigger peripheral
 * @return
 *      - ESP_OK: Connect 2D-DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Connect 2D-DMA channel failed because of invalid argument
 */
//go:linkname Dma2dConnect C.dma2d_connect
func Dma2dConnect(dma2d_chan Dma2dChannelHandleT, trig_periph *Dma2dTriggerT) EspErrT

/**
 * @brief A collection of strategy items that each 2D-DMA channel could apply
 */

type Dma2dStrategyConfigT struct {
	OwnerCheck        bool
	AutoUpdateDesc    bool
	EofTillDataPopped bool
}

/**
 * @brief Apply channel strategy for 2D-DMA channel
 *
 * Usually only to be called in `on_job_picked` callback.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle, get from the `on_job_picked` callback input argument `dma2d_chans`
 * @param[in] config Configuration of 2D-DMA channel strategy
 * @return
 *      - ESP_OK: Apply channel strategy successfully
 *      - ESP_ERR_INVALID_ARG: Apply channel strategy failed because of invalid argument
 */
//go:linkname Dma2dApplyStrategy C.dma2d_apply_strategy
func Dma2dApplyStrategy(dma2d_chan Dma2dChannelHandleT, config *Dma2dStrategyConfigT) EspErrT

/**
 * @brief A collection of transfer ability items that each 2D-DMA channel could apply to improve transfer efficiency
 *
 * @note The 2D-DMA driver has no knowledge about the DMA buffer (address and size) used by upper layer.
 *       So it's the responsibility of the **upper layer** to take care of the buffer address and size.
 *       Usually RX buffer at least requires 4-byte alignment to avoid overwriting other data by DMA write PSRAM process
 *       or its data being overwritten.
 */

type Dma2dTransferAbilityT struct {
	DescBurstEn     bool
	DataBurstLength Dma2dDataBurstLengthT
	MbSize          Dma2dMacroBlockSizeT
}

/**
 * @brief Configure 2D-DMA channel transfer ability for transfer efficiency
 *
 * Usually only to be called in `on_job_picked` callback.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle, get from the `on_job_picked` callback input argument `dma2d_chans`
 * @param[in] ability Configuration of 2D-DMA channel transfer ability
 * @return
 *      - ESP_OK: Set channel transfer ability successfully
 *      - ESP_ERR_INVALID_ARG: Set channel transfer ability failed because of invalid argument
 */
//go:linkname Dma2dSetTransferAbility C.dma2d_set_transfer_ability
func Dma2dSetTransferAbility(dma2d_chan Dma2dChannelHandleT, ability *Dma2dTransferAbilityT) EspErrT

/**
 * @brief A collection of color space conversion (CSC) items that each 2D-DMA channel could apply
 */

type Dma2dCscConfigT struct {
	PreScramble  Dma2dScrambleOrderT
	PostScramble Dma2dScrambleOrderT
}

/**
 * @brief Configure color space conversion setting for 2D-DMA channel
 *
 * Usually only to be called in `on_job_picked` callback.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle, get from the `on_job_picked` callback input argument `dma2d_chans`
 * @param[in] config Configuration of 2D-DMA channel color space conversion
 * @return
 *      - ESP_OK: Configure DMA color space conversion successfully
 *      - ESP_ERR_INVALID_ARG: Configure DMA color space conversion failed because of invalid argument
 */
//go:linkname Dma2dConfigureColorSpaceConversion C.dma2d_configure_color_space_conversion
func Dma2dConfigureColorSpaceConversion(dma2d_chan Dma2dChannelHandleT, config *Dma2dCscConfigT) EspErrT

/**
 * @brief A collection of configurations apply to 2D-DMA channel DSCR-PORT mode
 */

type Dma2dDscrPortModeConfigT struct {
	BlockH c.Uint32T
	BlockV c.Uint32T
}

/**
 * @brief Configure 2D-DMA channel DSCR-PORT mode
 *
 * @note This API only targets PPA SRM, which uses 2D-DMA DSCR-PORT mode.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle, get from the `on_job_picked` callback input argument `dma2d_chans`
 * @param[in] config Configuration of 2D-DMA channel DSCR-PORT mode
 * @return
 *      - ESP_OK: Configure 2D-DMA dscr-port mode successfully
 *      - ESP_ERR_INVALID_ARG: Configure 2D-DMA dscr-port mode failed because of invalid argument
 */
//go:linkname Dma2dConfigureDscrPortMode C.dma2d_configure_dscr_port_mode
func Dma2dConfigureDscrPortMode(dma2d_chan Dma2dChannelHandleT, config *Dma2dDscrPortModeConfigT) EspErrT

/**
 * @brief Type of 2D-DMA event data
 */

type Dma2dEventDataT struct {
	Transaction *Dma2dTransT
}

// llgo:type C
type Dma2dEventCallbackT func(Dma2dChannelHandleT, *Dma2dEventDataT, c.Pointer) bool

/**
 * @brief Group of supported 2D-DMA TX callbacks
 * @note The callbacks are all running under ISR environment
 */

type Dma2dTxEventCallbacksT struct {
	OnDescDone Dma2dEventCallbackT
}

/**
 * @brief Group of supported 2D-DMA RX callbacks
 * @note The callbacks are all running under ISR environment
 *
 * Users should be clear on the unique responsibility of each callback when writing the callback functions, such as
 * where to free the transaction memory.
 */

type Dma2dRxEventCallbacksT struct {
	OnRecvEof   Dma2dEventCallbackT
	OnDescDone  Dma2dEventCallbackT
	OnDescEmpty Dma2dEventCallbackT
}

/**
 * @brief Set 2D-DMA event callbacks for TX channel
 *
 * Usually only to be called in `on_job_picked` callback.
 *
 * @param[in] dma2d_chan 2D-DMA TX channel handle, get from the `on_job_picked` callback input argument `dma2d_chans`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 */
//go:linkname Dma2dRegisterTxEventCallbacks C.dma2d_register_tx_event_callbacks
func Dma2dRegisterTxEventCallbacks(dma2d_chan Dma2dChannelHandleT, cbs *Dma2dTxEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Set 2D-DMA event callbacks for RX channel
 *
 * Usually only to be called in `on_job_picked` callback.
 *
 * @param[in] dma2d_chan 2D-DMA RX channel handle, get from the `on_job_picked` callback input argument `dma2d_chans`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 */
//go:linkname Dma2dRegisterRxEventCallbacks C.dma2d_register_rx_event_callbacks
func Dma2dRegisterRxEventCallbacks(dma2d_chan Dma2dChannelHandleT, cbs *Dma2dRxEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Set descriptor address for 2D-DMA channel
 *
 * Usually only to be called in `on_job_picked` callback.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle, get from the `on_job_picked` callback input argument `dma2d_chans`
 * @param[in] desc_base_addr Base address of descriptors
 * @return
 *      - ESP_OK: Set 2D-DMA descriptor addr successfully
 *      - ESP_ERR_INVALID_ARG: Set 2D-DMA descriptor addr failed because of invalid argument
 */
//go:linkname Dma2dSetDescAddr C.dma2d_set_desc_addr
func Dma2dSetDescAddr(dma2d_chan Dma2dChannelHandleT, desc_base_addr c.IntptrT) EspErrT

/**
 * @brief Start engine for 2D-DMA channel
 *
 * Usually only to be called in `on_job_picked` callback.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle, get from the `on_job_picked` callback input argument `dma2d_chans`
 * @return
 *      - ESP_OK: Start 2D-DMA engine successfully
 *      - ESP_ERR_INVALID_ARG: Start 2D-DMA engine failed because of invalid argument
 */
//go:linkname Dma2dStart C.dma2d_start
func Dma2dStart(dma2d_chan Dma2dChannelHandleT) EspErrT

/**
 * @brief Stop engine for 2D-DMA channel
 *
 * Usually to be called in ISR context.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle
 * @return
 *      - ESP_OK: Stop 2D-DMA engine successfully
 *      - ESP_ERR_INVALID_ARG: Stop 2D-DMA engine failed because of invalid argument
 */
//go:linkname Dma2dStop C.dma2d_stop
func Dma2dStop(dma2d_chan Dma2dChannelHandleT) EspErrT

/**
 * @brief Make the appended descriptors be aware to the 2D-DMA engine
 *
 * Usually to be called in ISR context.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle
 * @return
 *      - ESP_OK: Send append command to 2D-DMA engine successfully
 *      - ESP_ERR_INVALID_ARG: Send append command to 2D-DMA engine failed because of invalid argument
 */
//go:linkname Dma2dAppend C.dma2d_append
func Dma2dAppend(dma2d_chan Dma2dChannelHandleT) EspErrT

/**
 * @brief Reset engine for 2D-DMA channel
 *
 * Usually to be called in ISR context.
 *
 * @param[in] dma2d_chan 2D-DMA channel handle
 * @return
 *      - ESP_OK: Reset 2D-DMA engine successfully
 *      - ESP_ERR_INVALID_ARG: Reset 2D-DMA engine failed because of invalid argument
 */
//go:linkname Dma2dReset C.dma2d_reset
func Dma2dReset(dma2d_chan Dma2dChannelHandleT) EspErrT
