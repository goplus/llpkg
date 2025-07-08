package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaChannelT struct {
	Unused [8]uint8
}
type GdmaChannelHandleT *GdmaChannelT

/**
 * @brief Collection of configuration items that used for allocating GDMA channel
 *
 */

type GdmaChannelAllocConfigT struct {
	SiblingChan GdmaChannelHandleT
	Direction   GdmaChannelDirectionT
	Flags       struct {
		ReserveSibling c.Int
		IsrCacheSafe   c.Int
	}
}

/**
 * @brief Type of GDMA event data
 */

type GdmaEventDataT struct {
	Flags struct {
		AbnormalEof c.Uint32T
		NormalEof   c.Uint32T
	}
}

// llgo:type C
type GdmaEventCallbackT func(GdmaChannelHandleT, *GdmaEventDataT, c.Pointer) bool

/**
 * @brief Group of supported GDMA TX callbacks
 * @note The callbacks are all running under ISR environment
 */

type GdmaTxEventCallbacksT struct {
	OnTransEof GdmaEventCallbackT
	OnDescrErr GdmaEventCallbackT
}

/**
 * @brief Group of supported GDMA RX callbacks
 * @note The callbacks are all running under ISR environment
 */

type GdmaRxEventCallbacksT struct {
	OnRecvEof  GdmaEventCallbackT
	OnDescrErr GdmaEventCallbackT
	OnRecvDone GdmaEventCallbackT
}

/**
 * @brief Type of GDMA engine trigger
 * @note It's recommended to initialize this structure with `GDMA_MAKE_TRIGGER`.
 *
 */

type GdmaTriggerT struct {
	Periph     GdmaTriggerPeripheralT
	InstanceId c.Int
	BusId      c.Int
}

/**
 * @brief A collection of strategy item that each GDMA channel could apply
 *
 */

type GdmaStrategyConfigT struct {
	OwnerCheck        bool
	AutoUpdateDesc    bool
	EofTillDataPopped bool
}

/**
 * @brief Create AHB-GDMA channel
 * @note This API won't install interrupt service for the allocated channel.
 *       If interrupt service is needed, user has to register GDMA event callback by `gdma_register_tx_event_callbacks` or `gdma_register_rx_event_callbacks`.
 *
 * @param[in] config Pointer to a collection of configurations for allocating GDMA channel
 * @param[out] ret_chan Returned channel handle
 * @return
 *      - ESP_OK: Create DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Create DMA channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create DMA channel failed because out of memory
 *      - ESP_FAIL: Create DMA channel failed because of other error
 */
// llgo:link (*GdmaChannelAllocConfigT).GdmaNewAhbChannel C.gdma_new_ahb_channel
func (recv_ *GdmaChannelAllocConfigT) GdmaNewAhbChannel(ret_chan *GdmaChannelHandleT) EspErrT {
	return 0
}

/**
 * @brief Create AXI-GDMA channel
 * @note This API won't install interrupt service for the allocated channel.
 *       If interrupt service is needed, user has to register GDMA event callback by `gdma_register_tx_event_callbacks` or `gdma_register_rx_event_callbacks`.
 *
 * @param[in] config Pointer to a collection of configurations for allocating GDMA channel
 * @param[out] ret_chan Returned channel handle
 * @return
 *      - ESP_OK: Create DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Create DMA channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create DMA channel failed because out of memory
 *      - ESP_FAIL: Create DMA channel failed because of other error
 */
// llgo:link (*GdmaChannelAllocConfigT).GdmaNewAxiChannel C.gdma_new_axi_channel
func (recv_ *GdmaChannelAllocConfigT) GdmaNewAxiChannel(ret_chan *GdmaChannelHandleT) EspErrT {
	return 0
}

/**
 * @brief Connect GDMA channel to trigger peripheral
 *
 * @note Suggest to use helper macro `GDMA_MAKE_TRIGGER` to construct parameter `trig_periph`. e.g. GDMA_MAKE_TRIGGER(GDMA_TRIG_PERIPH_SHA,0)
 * @note Connecting to a peripheral will also reset the DMA FIFO and FSM automatically
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[in] trig_periph GDMA trigger peripheral
 * @return
 *      - ESP_OK: Connect GDMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Connect GDMA channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Connect GDMA channel failed because DMA channel is working with another peripheral
 *      - ESP_FAIL: Connect GDMA channel failed because of other error
 */
//go:linkname GdmaConnect C.gdma_connect
func GdmaConnect(dma_chan GdmaChannelHandleT, trig_periph GdmaTriggerT) EspErrT

/**
 * @brief Disconnect GMA channel from peripheral
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @return
 *      - ESP_OK: Disconnect GDMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Disconnect GDMA channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disconnect GDMA channel failed because DMA channel is not connected to any peripheral
 *      - ESP_FAIL: Disconnect DMA channel failed because of other error
 */
//go:linkname GdmaDisconnect C.gdma_disconnect
func GdmaDisconnect(dma_chan GdmaChannelHandleT) EspErrT

/**
 * @brief Channel transfer configurations
 */

type GdmaTransferConfigT struct {
	MaxDataBurstSize c.Uint32T
	AccessExtMem     bool
}

/**
 * @brief Configure transfer parameters for a DMA channel
 *
 * @note It's highly recommended to enable the burst mode and set proper burst size for the DMA channel,
 *       which can improve the performance in accessing external memory by a lot.
 *
 * @param[in] chan DMA channel handle, allocated by `gdma_new_channel`
 * @param[in] config Transfer configurations
 * @return
 *      - ESP_OK: Configure DMA transfer parameters successfully
 *      - ESP_ERR_INVALID_ARG: Configure DMA transfer parameters failed because of invalid argument
 *      - ESP_FAIL: Configure DMA transfer parameters failed because of other error
 */
//go:linkname GdmaConfigTransfer C.gdma_config_transfer
func GdmaConfigTransfer(dma_chan GdmaChannelHandleT, config *GdmaTransferConfigT) EspErrT

/**
 * @brief Get the alignment constraints for internal and external memory
 *
 * @note You should call this function after `gdma_config_transfer`, the later one can
 *       adjust the alignment constraints based on various conditions, e.g. burst size, memory encryption, etc.
 * @note You can use returned alignment value to validate if a DMA buffer provided by the upper layer meets the constraints.
 * @note The returned alignment doesn't take the cache line size into account, if you want to do aligned memory allocation,
 *       you should align the buffer size to the cache line size by yourself if the DMA buffer is behind a cache.
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[out] int_mem_alignment Internal memory alignment
 * @param[out] ext_mem_alignment External memory alignment
 * @return
 *      - ESP_OK: Get alignment constraints successfully
 *      - ESP_ERR_INVALID_ARG: Get alignment constraints failed because of invalid argument
 *      - ESP_FAIL: Get alignment constraints failed because of other error
 */
//go:linkname GdmaGetAlignmentConstraints C.gdma_get_alignment_constraints
func GdmaGetAlignmentConstraints(dma_chan GdmaChannelHandleT, int_mem_alignment *c.SizeT, ext_mem_alignment *c.SizeT) EspErrT

/**
 * @brief Apply channel strategy for GDMA channel
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[in] config Configuration of GDMA channel strategy
 *      - ESP_OK: Apply channel strategy successfully
 *      - ESP_ERR_INVALID_ARG: Apply channel strategy failed because of invalid argument
 *      - ESP_FAIL: Apply channel strategy failed because of other error
 */
//go:linkname GdmaApplyStrategy C.gdma_apply_strategy
func GdmaApplyStrategy(dma_chan GdmaChannelHandleT, config *GdmaStrategyConfigT) EspErrT

/**
 * @brief Set GDMA channel priority
 *
 * @note By default, all GDMA channels are with the same priority: 0. Channels with the same priority are served in round-robin manner.
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[in] priority Priority of GDMA channel, higher value means higher priority
 * @return
 *      - ESP_OK: Set GDMA channel priority successfully
 *      - ESP_ERR_INVALID_ARG: Set GDMA channel priority failed because of invalid argument, e.g. priority out of range [0,GDMA_LL_CHANNEL_MAX_PRIORITY]
 *      - ESP_FAIL: Set GDMA channel priority failed because of other error
 */
//go:linkname GdmaSetPriority C.gdma_set_priority
func GdmaSetPriority(dma_chan GdmaChannelHandleT, priority c.Uint32T) EspErrT

/**
 * @brief Delete GDMA channel
 * @note If you call `gdma_new_channel` several times for a same peripheral, make sure you call this API the same times.
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @return
 *      - ESP_OK: Delete GDMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Delete GDMA channel failed because of invalid argument
 *      - ESP_FAIL: Delete GDMA channel failed because of other error
 */
//go:linkname GdmaDelChannel C.gdma_del_channel
func GdmaDelChannel(dma_chan GdmaChannelHandleT) EspErrT

/**
 * @brief Get the group ID and the channel ID
 *
 * @note This API breaks the encapsulation of GDMA Channel Object.
 *       With the returned group/channel ID, you can even bypass all other GDMA driver API and access Low Level API directly.
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[out] group_id Returned group ID
 * @param[out] channel_id Returned channel ID
 * @return
 *      - ESP_OK: Get GDMA channel/group ID successfully
 *      - ESP_ERR_INVALID_ARG: Get GDMA channel/group ID failed because of invalid argument
 *      - ESP_FAIL: Get GDMA channel ID failed because of other error
 */
//go:linkname GdmaGetGroupChannelId C.gdma_get_group_channel_id
func GdmaGetGroupChannelId(dma_chan GdmaChannelHandleT, group_id *c.Int, channel_id *c.Int) EspErrT

/**
 * @brief Set GDMA event callbacks for TX channel
 * @note This API will install GDMA interrupt service for the channel internally
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname GdmaRegisterTxEventCallbacks C.gdma_register_tx_event_callbacks
func GdmaRegisterTxEventCallbacks(dma_chan GdmaChannelHandleT, cbs *GdmaTxEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Set GDMA event callbacks for RX channel
 * @note This API will install GDMA interrupt service for the channel internally
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname GdmaRegisterRxEventCallbacks C.gdma_register_rx_event_callbacks
func GdmaRegisterRxEventCallbacks(dma_chan GdmaChannelHandleT, cbs *GdmaRxEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Set DMA descriptor address and start engine
 *
 * @note This function is allowed to run within ISR context
 * @note This function is also allowed to run when Cache is disabled, if `CONFIG_GDMA_CTRL_FUNC_IN_IRAM` is enabled
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[in] desc_base_addr Base address of descriptors (usually the descriptors are chained into a link or ring)
 * @return
 *      - ESP_OK: Start DMA engine successfully
 *      - ESP_ERR_INVALID_ARG: Start DMA engine failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Start DMA engine failed because of invalid state, e.g. the channel is controlled by ETM, so can't start it manually
 *      - ESP_FAIL: Start DMA engine failed because of other error
 */
//go:linkname GdmaStart C.gdma_start
func GdmaStart(dma_chan GdmaChannelHandleT, desc_base_addr c.IntptrT) EspErrT

/**
 * @brief Stop DMA engine
 *
 * @note This function is allowed to run within ISR context
 * @note This function is also allowed to run when Cache is disabled, if `CONFIG_GDMA_CTRL_FUNC_IN_IRAM` is enabled
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @return
 *      - ESP_OK: Stop DMA engine successfully
 *      - ESP_ERR_INVALID_ARG: Stop DMA engine failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Stop DMA engine failed because of invalid state, e.g. the channel is controlled by ETM, so can't stop it manually
 *      - ESP_FAIL: Stop DMA engine failed because of other error
 */
//go:linkname GdmaStop C.gdma_stop
func GdmaStop(dma_chan GdmaChannelHandleT) EspErrT

/**
 * @brief Make the appended descriptors be aware to the DMA engine
 *
 * @note This function is allowed to run within ISR context
 * @note This function is also allowed to run when Cache is disabled, if `CONFIG_GDMA_CTRL_FUNC_IN_IRAM` is enabled
 * @note This API could also resume a paused DMA engine, make sure new descriptors have been appended to the descriptor chain before calling it.
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @return
 *      - ESP_OK: Send append command to DMA engine successfully
 *      - ESP_ERR_INVALID_ARG: Send append command to DMA engine failed because of invalid argument
 *      - ESP_FAIL: Send append command to DMA engine failed because of other error
 */
//go:linkname GdmaAppend C.gdma_append
func GdmaAppend(dma_chan GdmaChannelHandleT) EspErrT

/**
 * @brief Reset DMA channel FIFO and internal finite state machine
 *
 * @note This function is allowed to run within ISR context
 * @note This function is also allowed to run when Cache is disabled, if `CONFIG_GDMA_CTRL_FUNC_IN_IRAM` is enabled
 * @note Resetting a DMA channel won't break the connection with the target peripheral
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @return
 *      - ESP_OK: DMA channel reset successfully
 *      - ESP_ERR_INVALID_ARG: DMA channel reset failed due to invalid arguments
 *      - ESP_FAIL: DMA channel reset failed due to other errors
 */
//go:linkname GdmaReset C.gdma_reset
func GdmaReset(dma_chan GdmaChannelHandleT) EspErrT

/**
 * @brief Get the mask of free M2M trigger IDs
 *
 * @note On some ESP targets (e.g. ESP32C3/S3), DMA trigger used for memory copy can be any of valid peripheral's trigger ID,
 *       which can bring conflict if the peripheral is also using the same trigger ID. This function can return the free IDs
 *       for memory copy, at the runtime.
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[out] mask Returned mask of free M2M trigger IDs
 * @return
 *      - ESP_OK: Get free M2M trigger IDs successfully
 *      - ESP_ERR_INVALID_ARG: Get free M2M trigger IDs failed because of invalid argument
 *      - ESP_FAIL: Get free M2M trigger IDs failed because of other error
 */
//go:linkname GdmaGetFreeM2mTrigIdMask C.gdma_get_free_m2m_trig_id_mask
func GdmaGetFreeM2mTrigIdMask(dma_chan GdmaChannelHandleT, mask *c.Uint32T) EspErrT

/** @cond */
/**
 * @brief Create GDMA channel (Legacy API)
 *
 * @param[in] config Pointer to a collection of configurations for allocating GDMA channel
 * @param[out] ret_chan Returned channel handle
 * @return
 *      - ESP_OK: Create DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Create DMA channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create DMA channel failed because out of memory
 *      - ESP_FAIL: Create DMA channel failed because of other error
 */
// llgo:link (*GdmaChannelAllocConfigT).GdmaNewChannel C.gdma_new_channel
func (recv_ *GdmaChannelAllocConfigT) GdmaNewChannel(ret_chan *GdmaChannelHandleT) EspErrT {
	return 0
}

/**
 * @brief GDMA transfer ability
 *
 * @note The alignment set in this structure is **not** a guarantee that gdma driver will take care of the nonalignment cases.
 *       Actually the GDMA driver has no knowledge about the DMA buffer (address and size) used by upper layer.
 *       So it's the responsibility of the **upper layer** to take care of the buffer address and size.
 *
 */

type GdmaTransferAbilityT struct {
	SramTransAlign  c.SizeT
	PsramTransAlign c.SizeT
}

/**
 * @brief Set DMA channel transfer ability
 *
 * @param[in] dma_chan GDMA channel handle, allocated by `gdma_new_channel`
 * @param[in] ability Transfer ability, e.g. alignment
 * @return
 *      - ESP_OK: Set DMA channel transfer ability successfully
 *      - ESP_ERR_INVALID_ARG: Set DMA channel transfer ability failed because of invalid argument
 *      - ESP_FAIL: Set DMA channel transfer ability failed because of other error
 */
//go:linkname GdmaSetTransferAbility C.gdma_set_transfer_ability
func GdmaSetTransferAbility(dma_chan GdmaChannelHandleT, ability *GdmaTransferAbilityT) EspErrT
