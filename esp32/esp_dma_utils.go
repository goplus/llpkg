package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief DMA Mem info
 */

type EspDmaMemInfoT struct {
	ExtraHeapCaps     c.Int
	DmaAlignmentBytes c.SizeT
}

/**
 * @brief Helper function for malloc a DMA capable memory buffer
 *
 * @note This API will take care of the cache alignment internally,
 *       you will need to set `esp_dma_mem_info_t: dma_alignment_bytes`
 *       with either the custom alignment or DMA alignment of used peripheral driver.
 *
 * @param[in]  size          Size in bytes, the amount of memory to allocate
 * @param[in]  dma_mem_info  DMA and memory info, see `esp_dma_mem_info_t`
 * @param[out] out_ptr       A pointer to the memory allocated successfully
 * @param[out] actual_size   Actual size for allocation in bytes, when the size you specified doesn't meet the DMA alignment requirements, this value might be bigger than the size you specified. Set null if you don't care this value.
 *
 * @return
 *        - ESP_OK:
 *        - ESP_ERR_INVALID_ARG: Invalid argument
 *        - ESP_ERR_NO_MEM:      No enough memory for allocation
 */
//go:linkname EspDmaCapableMalloc C.esp_dma_capable_malloc
func EspDmaCapableMalloc(size c.SizeT, dma_mem_info *EspDmaMemInfoT, out_ptr *c.Pointer, actual_size *c.SizeT) EspErrT

/**
 * @brief Helper function for calloc a DMA capable memory buffer
 *
 * @param[in]  calloc_num    Number of elements to allocate
 * @param[in]  size          Size in bytes, the amount of memory to allocate
 * @param[in]  dma_mem_info  DMA and memory info, see `esp_dma_mem_info_t`
 * @param[out] out_ptr       A pointer to the memory allocated successfully
 * @param[out] actual_size   Actual size for allocation in bytes, when the size you specified doesn't meet the DMA alignment requirements, this value might be bigger than the size you specified. Set null if you don't care this value.
 *
 * @return
 *        - ESP_OK:
 *        - ESP_ERR_INVALID_ARG: Invalid argument
 *        - ESP_ERR_NO_MEM:      No enough memory for allocation
 */
//go:linkname EspDmaCapableCalloc C.esp_dma_capable_calloc
func EspDmaCapableCalloc(calloc_num c.SizeT, size c.SizeT, dma_mem_info *EspDmaMemInfoT, out_ptr *c.Pointer, actual_size *c.SizeT) EspErrT

/**
 * @brief Helper function to check if a DMA buffer pointer and size meet both hardware alignment requirements and custom alignment requirements
 *
 * @param[in]  ptr           Pointer to the buffer
 * @param[in]  size          Size of the buffer
 * @param[in]  dma_mem_info  DMA and memory info, see `esp_dma_mem_info_t`
 *
 * @return
 *        - True:  Buffer is aligned
 *        - False: Buffer is not aligned, or buffer is not DMA capable
 */
//go:linkname EspDmaIsBufferAlignmentSatisfied C.esp_dma_is_buffer_alignment_satisfied
func EspDmaIsBufferAlignmentSatisfied(ptr c.Pointer, size c.SizeT, dma_mem_info EspDmaMemInfoT) bool

/**
 * @brief Needed info to get GDMA alignment
 */

type DmaAlignmentInfoT struct {
	IsDesc  bool
	OnPsram bool
}

/**
 * @note This API will use MAX alignment requirement
 */
//go:linkname EspDmaMalloc C.esp_dma_malloc
func EspDmaMalloc(size c.SizeT, flags c.Uint32T, out_ptr *c.Pointer, actual_size *c.SizeT) EspErrT

/**
 * @note This API will use MAX alignment requirement
 */
//go:linkname EspDmaCalloc C.esp_dma_calloc
func EspDmaCalloc(n c.SizeT, size c.SizeT, flags c.Uint32T, out_ptr *c.Pointer, actual_size *c.SizeT) EspErrT

type EspDmaBufLocationT c.Int

const (
	ESP_DMA_BUF_LOCATION_INTERNAL EspDmaBufLocationT = 0
	ESP_DMA_BUF_LOCATION_PSRAM    EspDmaBufLocationT = 1
	ESP_DMA_BUF_LOCATION_AUTO     EspDmaBufLocationT = 2
)

/**
 * @note This API will use MAX alignment requirement
 */
//go:linkname EspDmaIsBufferAligned C.esp_dma_is_buffer_aligned
func EspDmaIsBufferAligned(ptr c.Pointer, size c.SizeT, location EspDmaBufLocationT) bool

/**
 * @brief DMA buffer information
 */

type DmaBufferSplitInfoT struct {
	AlignedBuffer   c.Pointer
	RecoveryAddress c.Pointer
	Length          c.SizeT
}

/**
 * @brief DMA buffer aligned array
 * The array contains three parts: head, body and tail.
 * Length of each part will be >=0, especially, length=0 means that there is no such part.
 */

type DmaBufferSplitArrayT struct {
	Unused [8]uint8
}

/**
 * @brief Split DMA RX buffer to cache aligned buffers
 *
 * @note After the original RX buffer is split into an array, caller should mount the buffer array to the DMA controller in scatter-gather mode.
 *       Don't read/write the aligned buffers before the DMA finished using them.
 *
 * @param[in]   rx_buffer        The origin DMA buffer used for receiving data
 * @param[in]   buffer_len       rx_buffer length
 * @param[out]  align_buf_array  Aligned DMA buffer array
 * @param[out]  ret_stash_buffer Allocated stash buffer (caller should free it after use)
 * @return
 *      - ESP_OK: Split to aligned buffer successfully
 *      - ESP_ERR_INVALID_ARG: Split to aligned buffer failed because of invalid argument
 *
 *  brief sketch:
 *                  cache alignment delimiter    cache alignment delimiter
 *                              │                             │
 *     Origin Buffer            │        Origin Buffer        │
 *           │                  │              │              │
 *           │                  ▼              ▼              ▼
 *           │       ...---xxxxx|xxxxxxxxxxxxxxxxxxxxxxxxxxxxx|xxxxx----...
 *           │               │                 │                 │
 *           │               │                 ▼                 │
 *           │               │  |xxxxxxxxxxxxxxxxxxxxxxxxxxxxx|  │
 *           │               │                 ▲                 │
 *           ▼               │                 │                 │
 *     Aligned buffers       └──► Head        Body   Tail ◄──────┘
 *                                 │                  │
 *                                 ▼                  ▼
 *                              |xxxxx......|     |xxxxx......|
 */
//go:linkname EspDmaSplitRxBufferToCacheAligned C.esp_dma_split_rx_buffer_to_cache_aligned
func EspDmaSplitRxBufferToCacheAligned(rx_buffer c.Pointer, buffer_len c.SizeT, align_buf_array *DmaBufferSplitArrayT, ret_stash_buffer **c.Uint8T) EspErrT

/**
 * @brief Merge aligned RX buffer array to origin buffer
 *
 * @note This function can be used in the ISR context.
 *
 * @param[in] align_buf_array Aligned DMA buffer array
 * @return
 *      - ESP_OK: Merge aligned buffer to origin buffer successfully
 *      - ESP_ERR_INVALID_ARG: Merge aligned buffer to origin buffer failed because of invalid argument
 */
// llgo:link (*DmaBufferSplitArrayT).EspDmaMergeAlignedRxBuffers C.esp_dma_merge_aligned_rx_buffers
func (recv_ *DmaBufferSplitArrayT) EspDmaMergeAlignedRxBuffers() EspErrT {
	return 0
}

/**
 * @brief Calculate the number of DMA linked list nodes required for a given buffer size
 *
 * @param[in] buffer_size Total size of the buffer
 * @param[in] buffer_alignment Alignment requirement for the buffer
 * @param[in] max_buffer_size_per_node Maximum buffer size that each node can handle
 * @return Number of DMA linked list nodes required
 */
//go:linkname EspDmaCalculateNodeCount C.esp_dma_calculate_node_count
func EspDmaCalculateNodeCount(buffer_size c.SizeT, buffer_alignment c.SizeT, max_buffer_size_per_node c.SizeT) c.SizeT
