package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DwGdmaChannelT struct {
	Unused [8]uint8
}
type DwGdmaChannelHandleT *DwGdmaChannelT

type DwGdmaLinkListT struct {
	Unused [8]uint8
}
type DwGdmaLinkListHandleT *DwGdmaLinkListT

type DwGdmaLinkListItemT struct {
	Unused [8]uint8
}
type DwGdmaLliHandleT *DwGdmaLinkListItemT

/**
 * @brief A group of channel's static configurations
 *
 * @note By static, we mean these channel end configurations shouldn't be changed after the DMA channel is created.
 */

type DwGdmaChannelStaticConfigT struct {
	BlockTransferType      DwGdmaBlockTransferTypeT
	Role                   DwGdmaRoleT
	HandshakeType          DwGdmaHandshakeTypeT
	NumOutstandingRequests c.Uint8T
	StatusFetchAddr        c.Uint32T
}

/**
 * @brief Configurations for allocating a DMA channel
 */

type DwGdmaChannelAllocConfigT struct {
	Src            DwGdmaChannelStaticConfigT
	Dst            DwGdmaChannelStaticConfigT
	FlowController DwGdmaFlowControllerT
	ChanPriority   c.Int
	IntrPriority   c.Int
}

/**
 * @brief Create a DMA channel
 *
 * @param[in] config Channel allocation configuration
 * @param[out] ret_chan Returned channel handle
 * @return
 *      - ESP_OK: Create DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Create DMA channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create DMA channel failed because out of memory
 *      - ESP_FAIL: Create DMA channel failed because of other error
 */
// llgo:link (*DwGdmaChannelAllocConfigT).DwGdmaNewChannel C.dw_gdma_new_channel
func (recv_ *DwGdmaChannelAllocConfigT) DwGdmaNewChannel(ret_chan *DwGdmaChannelHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete DMA channel
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @return
 *      - ESP_OK: Delete DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Delete DMA channel failed because of invalid argument
 *      - ESP_FAIL: Delete DMA channel failed because of other error
 */
//go:linkname DwGdmaDelChannel C.dw_gdma_del_channel
func DwGdmaDelChannel(chan_ DwGdmaChannelHandleT) EspErrT

/**
 * @brief Get the DMA channel ID
 *
 * @note This API breaks the encapsulation of DW_GDMA Channel Object.
 *       With the returned channel ID, you can even bypass all other driver API and access Low Level API directly.
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @param[out] channel_id Returned channel ID
 * @return
 *      - ESP_OK: Get DW_GDMA channel ID successfully
 *      - ESP_ERR_INVALID_ARG: Get DW_GDMA channel ID failed because of invalid argument
 *      - ESP_FAIL: Get DW_GDMA channel ID failed because of other error
 */
//go:linkname DwGdmaChannelGetId C.dw_gdma_channel_get_id
func DwGdmaChannelGetId(chan_ DwGdmaChannelHandleT, channel_id *c.Int) EspErrT

/**
 * @brief A group of channel's dynamic configurations
 *
 * @note By dynamic, we mean these channel end configurations can be changed in each transfer.
 */

type DwGdmaChannelDynamicConfigT struct {
	Addr       c.Uint32T
	Width      DwGdmaTransferWidthT
	BurstMode  DwGdmaBurstModeT
	BurstItems DwGdmaBurstItemsT
	BurstLen   c.Uint8T
	Flags      struct {
		EnStatusWriteBack c.Uint32T
	}
}

/**
 * @brief Channel block transfer configurations
 */

type DwGdmaBlockTransferConfigT struct {
	Src  DwGdmaChannelDynamicConfigT
	Dst  DwGdmaChannelDynamicConfigT
	Size c.SizeT
}

/**
 * @brief Configure transfer parameters for a DMA channel
 *
 * @note This is an "all-in-one" function for set up the block transfer.
 * @note This function can't work with Link-List transfer type. For Link-List transfer, please use `dw_gdma_lli_config_transfer` instead.
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @param[in] config Block transfer configurations
 * @return
 *      - ESP_OK: Configure DMA channel block transfer successfully
 *      - ESP_ERR_INVALID_ARG: Configure DMA channel block transfer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Configure DMA channel block transfer failed because the channel has Link-List transfer type
 *      - ESP_FAIL: Configure DMA channel block transfer failed because of other error
 */
//go:linkname DwGdmaChannelConfigTransfer C.dw_gdma_channel_config_transfer
func DwGdmaChannelConfigTransfer(chan_ DwGdmaChannelHandleT, config *DwGdmaBlockTransferConfigT) EspErrT

/**
 * @brief Enable or disable a DMA channel
 *
 * @note Before enabling a channel, you need to setup the channel transfer by either `dw_gdma_channel_config_transfer` or `dw_gdma_lli_config_transfer`
 * @note When a DMA channel is disabled, the DMA engine will stop working. You need to reconfigure the channel before enabling it again.
 * @note After all block transfers are completed, the DMA channel will be disabled automatically.
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @param[in] en_or_dis True to enable, false to disable the DMA channel
 * @return
 *      - ESP_OK: Enable or disable DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Enable or disable DMA channel failed because of invalid argument
 *      - ESP_FAIL: Enable or disable DMA channel failed because of other error
 */
//go:linkname DwGdmaChannelEnableCtrl C.dw_gdma_channel_enable_ctrl
func DwGdmaChannelEnableCtrl(chan_ DwGdmaChannelHandleT, en_or_dis bool) EspErrT

/**
 * @brief Suspend or resume a DMA channel
 *
 * @note When a DMA channel is suspended, the DMA engine will stop working gracefully and the channel's status will be saved.
 * @note The channel will exit the suspend state automatically if it is disabled.
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @param[in] enter_or_exit True to suspend, false to resume the DMA channel
 * @return
 *      - ESP_OK: Suspend or resume DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Suspend or resume DMA channel failed because of invalid argument
 *      - ESP_FAIL: Suspend or resume DMA channel failed because of other error
 */
//go:linkname DwGdmaChannelSuspendCtrl C.dw_gdma_channel_suspend_ctrl
func DwGdmaChannelSuspendCtrl(chan_ DwGdmaChannelHandleT, enter_or_exit bool) EspErrT

/**
 * @brief Abort the DMA channel
 *
 * @note If the channel is aborted, it will be disabled immediately, which may cause AXI bus protocol violation.
 * @note This function is recommended to only be used when the channel hangs. Recommend to try `dw_gdma_channel_enable_ctrl` first, then opt for aborting.
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @return
 *      - ESP_OK: Abort DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Abort DMA channel failed because of invalid argument
 *      - ESP_FAIL: Abort DMA channel failed because of other error
 */
//go:linkname DwGdmaChannelAbort C.dw_gdma_channel_abort
func DwGdmaChannelAbort(chan_ DwGdmaChannelHandleT) EspErrT

/**
 * @brief Lock the DMA channel at specific transfer level
 *
 * @note When a DMA channel is locked, no other channels are granted control of the master bus for the duration specified by the lock level.
 * @note Only lock the channel if you want to exclusive access to the master bus.
 * @note Channel locking feature is only for M2M transfer.
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @param[in] level Transfer level
 * @return
 *      - ESP_OK: Lock DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Lock DMA channel failed because of invalid argument
 *      - ESP_FAIL: Lock DMA channel failed because of other error
 */
//go:linkname DwGdmaChannelLock C.dw_gdma_channel_lock
func DwGdmaChannelLock(chan_ DwGdmaChannelHandleT, level DwGdmaLockLevelT) EspErrT

/**
 * @brief Unlock the DMA channel
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @return
 *      - ESP_OK: Unlock DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Unlock DMA channel failed because of invalid argument
 *      - ESP_FAIL: Unlock DMA channel failed because of other error
 */
//go:linkname DwGdmaChannelUnlock C.dw_gdma_channel_unlock
func DwGdmaChannelUnlock(chan_ DwGdmaChannelHandleT) EspErrT

/**
 * @brief Continue the temporarily stopped DMA transfer because of invalid block
 *
 * @note You should only call this API when the block becomes valid again,
 *       by calling `dw_gdma_lli_set_block_markers`/`dw_gdma_channel_set_block_markers` with `is_valid` set to true.
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @return
 *      - ESP_OK: Continue DMA transfer successfully
 *      - ESP_ERR_INVALID_ARG: Continue DMA transfer failed because of invalid argument
 *      - ESP_FAIL: Continue DMA transfer failed because of other error
 */
//go:linkname DwGdmaChannelContinue C.dw_gdma_channel_continue
func DwGdmaChannelContinue(chan_ DwGdmaChannelHandleT) EspErrT

/**
 * @brief Type of DW_GDMA trans done event data
 */

type DwGdmaTransDoneEventDataT struct {
	Unused [8]uint8
}

// llgo:type C
type DwGdmaTransDoneEventCallbackT func(DwGdmaChannelHandleT, *DwGdmaTransDoneEventDataT, c.Pointer) bool

/**
 * @brief Type of DW_GDMA break event data
 */

type DwGdmaBreakEventDataT struct {
	InvalidLli DwGdmaLliHandleT
}

// llgo:type C
type DwGdmaBreakEventCallbackT func(DwGdmaChannelHandleT, *DwGdmaBreakEventDataT, c.Pointer) bool

/**
 * @brief Group of supported DW_GDMA callbacks
 * @note The callbacks are all running under ISR environment
 */

type DwGdmaEventCallbacksT struct {
	OnBlockTransDone DwGdmaTransDoneEventCallbackT
	OnFullTransDone  DwGdmaTransDoneEventCallbackT
	OnInvalidBlock   DwGdmaBreakEventCallbackT
}

/**
 * @brief Set DW_GDMA event callbacks for a channel
 * @note This API will lazy install the DW_GDMA interrupt service
 *
 * @param[in] chan DW_GDMA channel handle, allocated by `dw_gdma_new_channel`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname DwGdmaChannelRegisterEventCallbacks C.dw_gdma_channel_register_event_callbacks
func DwGdmaChannelRegisterEventCallbacks(chan_ DwGdmaChannelHandleT, cbs *DwGdmaEventCallbacksT, user_data c.Pointer) EspErrT

type DwGdmaLinkListTypeT c.Int

const (
	DW_GDMA_LINKED_LIST_TYPE_SINGLY   DwGdmaLinkListTypeT = 0
	DW_GDMA_LINKED_LIST_TYPE_CIRCULAR DwGdmaLinkListTypeT = 1
)

/**
 * @brief DMA link list configurations
 */

type DwGdmaLinkListConfigT struct {
	NumItems c.Uint32T
	LinkType DwGdmaLinkListTypeT
}

/**
 * @brief Create a DMA link list
 *
 * @param[in] config Link list configurations
 * @param[out] ret_list Returned link list handle
 * @return
 *      - ESP_OK: Create DMA link list successfully
 *      - ESP_ERR_INVALID_ARG: Create DMA link list failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create DMA link list failed because out of memory
 *      - ESP_FAIL: Create DMA link list failed because of other error
 */
// llgo:link (*DwGdmaLinkListConfigT).DwGdmaNewLinkList C.dw_gdma_new_link_list
func (recv_ *DwGdmaLinkListConfigT) DwGdmaNewLinkList(ret_list *DwGdmaLinkListHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete a DMA link list
 *
 * @param[in] list Link list handle, allocated by `dw_gdma_new_link_list`
 * @return
 *      - ESP_OK: Delete DMA link list successfully
 *      - ESP_ERR_INVALID_ARG: Delete DMA link list failed because of invalid argument
 *      - ESP_FAIL: Delete DMA link list failed because of other error
 */
//go:linkname DwGdmaDelLinkList C.dw_gdma_del_link_list
func DwGdmaDelLinkList(list DwGdmaLinkListHandleT) EspErrT

/**
 * @brief Apply a link list to a DMA channel
 *
 * @note This function can only work with Link-List transfer type.
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @param[in] list Link list handle, allocated by `dw_gdma_new_link_list`
 * @return
 *      - ESP_OK: Apply link list to DMA channel successfully
 *      - ESP_ERR_INVALID_ARG: Apply link list to DMA channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Apply link list to DMA channel failed because the channel is not with Link-List transfer type
 *      - ESP_FAIL: Apply link list to DMA channel failed because of other error
 */
//go:linkname DwGdmaChannelUseLinkList C.dw_gdma_channel_use_link_list
func DwGdmaChannelUseLinkList(chan_ DwGdmaChannelHandleT, list DwGdmaLinkListHandleT) EspErrT

/**
 * @brief A helper function to return an item from a given link list, by index
 *
 * @note The address of the returned item is not behind the cache
 *
 * @param[in] list Link list handle, allocated by `dw_gdma_new_link_list`
 * @param[in] item_index Index of the item
 * @return
 *      - NULL: Invalid argument
 *      - Others: Link list item handle
 */
//go:linkname DwGdmaLinkListGetItem C.dw_gdma_link_list_get_item
func DwGdmaLinkListGetItem(list DwGdmaLinkListHandleT, item_index c.Int) DwGdmaLliHandleT

/**
 * @brief Configure transfer parameters for a DMA link list item
 *
 * @note This is an "all-in-one" function for set up the link list item.
 * @note This function can only work with Link-List transfer type. For other transfer types, please use `dw_gdma_channel_config_transfer` instead.
 *
 * @param[in] lli Link list item
 * @param[in] config Block transfer configurations
 * @return
 *      - ESP_OK: Configure link list item block transfer successfully
 *      - ESP_ERR_INVALID_ARG: Configure link list item block transfer failed because of invalid argument
 *      - ESP_FAIL: Configure link list item block transfer failed because of other error
 */
//go:linkname DwGdmaLliConfigTransfer C.dw_gdma_lli_config_transfer
func DwGdmaLliConfigTransfer(lli DwGdmaLliHandleT, config *DwGdmaBlockTransferConfigT) EspErrT

/**
 * @brief Set the next link list item for a given DMA link list item
 *
 * @param[in] lli Current link list item, can be obtained from `dw_gdma_link_list_get_item`
 * @param[in] next Next link list item, can be obtained from `dw_gdma_link_list_get_item`
 * @return
 *      - ESP_OK: Set next link list item successfully
 *      - ESP_ERR_INVALID_ARG: Set next link list item failed because of invalid argument
 *      - ESP_FAIL: Set next link list item failed because of other error
 */
//go:linkname DwGdmaLliSetNext C.dw_gdma_lli_set_next
func DwGdmaLliSetNext(lli DwGdmaLliHandleT, next DwGdmaLliHandleT) EspErrT

/**
 * @brief Markers of a DW_GDMA block
 */

type DwGdmaBlockMarkersT struct {
	IsLast          c.Uint32T
	IsValid         c.Uint32T
	EnTransDoneIntr c.Uint32T
}

/**
 * @brief Set block markers for a DMA channel
 *
 * @note This function doesn't work for Link-List transfer type. For Link-List transfer, please use `dw_gdma_lli_set_block_markers` instead.
 * @note Setting the markers should always be the last step of configuring a block transfer, before enabling/continuing the channel.
 *
 * @param[in] chan DMA channel handle, allocated by `dw_gdma_new_channel`
 * @param[in] markers Block markers
 * @return
 *      - ESP_OK: Set block markers successfully
 *      - ESP_ERR_INVALID_ARG: Set block markers failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Set block markers failed because the channel has Link-List transfer type
 *      - ESP_FAIL: Set block markers failed because of other error
 */
//go:linkname DwGdmaChannelSetBlockMarkers C.dw_gdma_channel_set_block_markers
func DwGdmaChannelSetBlockMarkers(chan_ DwGdmaChannelHandleT, markers DwGdmaBlockMarkersT) EspErrT

/**
 * @brief Set block markers for a DMA link list item
 *
 * @note Setting the markers should always be the last step of configuring a block transfer, before enabling/continuing the channel.
 *
 * @param[in] lli Link list item
 * @param[in] markers Block markers
 * @return
 *      - ESP_OK: Set block markers successfully
 *      - ESP_ERR_INVALID_ARG: Set block markers failed because of invalid argument
 *      - ESP_FAIL: Set block markers failed because of other error
 */
//go:linkname DwGdmaLliSetBlockMarkers C.dw_gdma_lli_set_block_markers
func DwGdmaLliSetBlockMarkers(lli DwGdmaLliHandleT, markers DwGdmaBlockMarkersT) EspErrT
