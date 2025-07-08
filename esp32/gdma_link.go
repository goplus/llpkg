package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaLinkListT struct {
	Unused [8]uint8
}
type GdmaLinkListHandleT *GdmaLinkListT

/**
 * @brief DMA link list configurations
 */

type GdmaLinkListConfigT struct {
	NumItems        c.Uint32T
	ItemAlignment   c.SizeT
	BufferAlignment c.SizeT
	Flags           GdmaLinkListFlags
}

type GdmaLinkListFlags struct {
	Unused [8]uint8
}

/**
 * @brief Create a DMA link list
 *
 * @note This function will allocate memory for the link list.
 *
 * @param[in] config Link list configurations
 * @param[out] ret_list Returned link list handle
 * @return
 *      - ESP_OK: Create DMA link list successfully
 *      - ESP_ERR_INVALID_ARG: Create DMA link list failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create DMA link list failed because out of memory
 *      - ESP_FAIL: Create DMA link list failed because of other error
 */
// llgo:link (*GdmaLinkListConfigT).GdmaNewLinkList C.gdma_new_link_list
func (recv_ *GdmaLinkListConfigT) GdmaNewLinkList(ret_list *GdmaLinkListHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete a DMA link list
 *
 * @param[in] list Link list handle, allocated by `gdma_new_link_list`
 * @return
 *      - ESP_OK: Delete DMA link list successfully
 *      - ESP_ERR_INVALID_ARG: Delete DMA link list failed because of invalid argument
 *      - ESP_FAIL: Delete DMA link list failed because of other error
 */
//go:linkname GdmaDelLinkList C.gdma_del_link_list
func GdmaDelLinkList(list GdmaLinkListHandleT) EspErrT

/**
 * @brief DMA buffer mount configurations
 */

type GdmaBufferMountConfigT struct {
	Buffer c.Pointer
	Length c.SizeT
	Flags  GdmaBufferMountFlags
}

type GdmaBufferMountFlags struct {
	Unused [8]uint8
}

/**
 * @brief Mount one or more buffers to a given link list
 *
 * @note Different buffers won't be mounted to the same DMA link list item
 * @note After mount to the last list item, the next list item will be the head item (wrap around)
 *
 * @param[in] list Link list handle, allocated by `gdma_new_link_list`
 * @param[in] start_item_index Index of the first item in the link list to be mounted
 * @param[in] buf_config_array Array of buffer mount configurations
 * @param[in] num_buf Number of buffers to be mounted
 * @param[out] end_item_index Index of the last item in the link list that has been mounted
 * @return
 *      - ESP_OK: Mount the buffer successfully
 *      - ESP_ERR_INVALID_ARG: Mount the buffer failed because of invalid argument
 *      - ESP_FAIL: Mount the buffer failed because of other error
 */
//go:linkname GdmaLinkMountBuffers C.gdma_link_mount_buffers
func GdmaLinkMountBuffers(list GdmaLinkListHandleT, start_item_index c.Int, buf_config_array *GdmaBufferMountConfigT, num_buf c.SizeT, end_item_index *c.Int) EspErrT

/**
 * @brief Get the address of the head item in the link list
 *
 * @note The head address can be used to start a DMA transfer
 *
 * @param[in] list Link list handle, allocated by `gdma_new_link_list`
 * @return
 *      - Address of the head item in the link list
 *      - NULL: Get the address failed
 */
//go:linkname GdmaLinkGetHeadAddr C.gdma_link_get_head_addr
func GdmaLinkGetHeadAddr(list GdmaLinkListHandleT) c.UintptrT

/**
 * @brief Concatenate two link lists as follows:
 *
 *        Link A: A1 --> A2 --> A3 --> A4
 *                    | item_index
 *                    +-----+
 *                          |
 *                          v item_index
 *        Link B: B1 --> B2 --> B3 --> B4
 *
 *      After concatenation:
 *       Link A: A1 --> B3 --> B4
 *       Link B: B1 --> B2 --> B3 --> B4
 *
 * @param[in] first_link First link list handle, allocated by `gdma_new_link_list`
 * @param[in] first_link_item_index Index of the item in the first link list (-1 means the last item)
 * @param[in] second_link Second link list handle, allocated by `gdma_new_link_list`
 * @param[in] second_link_item_index Index of the item in the second link list (-1 means the last item)
 * @return
 *      - ESP_OK: Concatenate the link lists successfully
 *      - ESP_ERR_INVALID_ARG: Concatenate the link lists failed because of invalid argument
 *      - ESP_FAIL: Concatenate the link lists failed because of other error
 */
//go:linkname GdmaLinkConcat C.gdma_link_concat
func GdmaLinkConcat(first_link GdmaLinkListHandleT, first_link_item_index c.Int, second_link GdmaLinkListHandleT, second_link_item_index c.Int) EspErrT

type GdmaLliOwnerT c.Int

const (
	GDMA_LLI_OWNER_CPU GdmaLliOwnerT = 0
	GDMA_LLI_OWNER_DMA GdmaLliOwnerT = 1
)

/**
 * @brief Set the ownership for a DMA link list item
 *
 * @param[in] list Link list handle, allocated by `gdma_new_link_list`
 * @param[in] item_index Index of the link list item (-1 means the last item)
 * @param[in] owner Ownership
 * @return
 *      - ESP_OK: Set the ownership successfully
 *      - ESP_ERR_INVALID_ARG: Set the ownership failed because of invalid argument
 *      - ESP_FAIL: Set the ownership failed because of other error
 */
//go:linkname GdmaLinkSetOwner C.gdma_link_set_owner
func GdmaLinkSetOwner(list GdmaLinkListHandleT, item_index c.Int, owner GdmaLliOwnerT) EspErrT

/**
 * @brief Get the ownership of a DMA link list item
 *
 * @param[in] list Link list handle, allocated by `gdma_new_link_list`
 * @param[in] item_index Index of the link list item (-1 means the last item)
 * @param[out] owner Ownership
 * @return
 *      - ESP_OK: Get the ownership successfully
 *      - ESP_ERR_INVALID_ARG: Get the ownership failed because of invalid argument
 *      - ESP_FAIL: Get the ownership failed because of other error
 */
//go:linkname GdmaLinkGetOwner C.gdma_link_get_owner
func GdmaLinkGetOwner(list GdmaLinkListHandleT, item_index c.Int, owner *GdmaLliOwnerT) EspErrT

/**
 * @brief Get the size of the buffer that is mounted to the link list until the eof item (inclusive)
 *
 * @param[in] list Link list handle, allocated by `gdma_new_link_list`
 * @param[in] start_item_index Index of the first item in the link list to be calculated
 * @return Size of the buffer that is mounted to the link list until the eof item (inclusive).
 *         If the link list is empty or invalid, return 0.
 */
//go:linkname GdmaLinkCountBufferSizeTillEof C.gdma_link_count_buffer_size_till_eof
func GdmaLinkCountBufferSizeTillEof(list GdmaLinkListHandleT, start_item_index c.Int) c.SizeT
