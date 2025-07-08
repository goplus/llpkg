package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type EspAllocFailedHookT func(c.SizeT, c.Uint32T, *c.Char)

/**
 * @brief registers a callback function to be invoked if a memory allocation operation fails
 * @param callback caller defined callback to be invoked
 * @return ESP_OK if callback was registered.
 */
//go:linkname HeapCapsRegisterFailedAllocCallback C.heap_caps_register_failed_alloc_callback
func HeapCapsRegisterFailedAllocCallback(callback EspAllocFailedHookT) EspErrT

/**
 * @brief Allocate a chunk of memory which has the given capabilities
 *
 * Equivalent semantics to libc malloc(), for capability-aware memory.
 *
 * @param size Size, in bytes, of the amount of memory to allocate
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory to be returned
 *
 * @return A pointer to the memory allocated on success, NULL on failure
 */
//go:linkname HeapCapsMalloc C.heap_caps_malloc
func HeapCapsMalloc(size c.SizeT, caps c.Uint32T) c.Pointer

/**
 * @brief Free memory previously allocated via heap_caps_malloc() or heap_caps_realloc().
 *
 * Equivalent semantics to libc free(), for capability-aware memory.
 *
 *  In IDF, ``free(p)`` is equivalent to ``heap_caps_free(p)``.
 *
 * @param ptr Pointer to memory previously returned from heap_caps_malloc() or heap_caps_realloc(). Can be NULL.
 */
//go:linkname HeapCapsFree C.heap_caps_free
func HeapCapsFree(ptr c.Pointer)

/**
 * @brief Reallocate memory previously allocated via heap_caps_malloc() or heap_caps_realloc().
 *
 * Equivalent semantics to libc realloc(), for capability-aware memory.
 *
 * In IDF, ``realloc(p, s)`` is equivalent to ``heap_caps_realloc(p, s, MALLOC_CAP_8BIT)``.
 *
 * 'caps' parameter can be different to the capabilities that any original 'ptr' was allocated with. In this way,
 * realloc can be used to "move" a buffer if necessary to ensure it meets a new set of capabilities.
 *
 * @param ptr Pointer to previously allocated memory, or NULL for a new allocation.
 * @param size Size of the new buffer requested, or 0 to free the buffer.
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory desired for the new allocation.
 *
 * @return Pointer to a new buffer of size 'size' with capabilities 'caps', or NULL if allocation failed.
 */
//go:linkname HeapCapsRealloc C.heap_caps_realloc
func HeapCapsRealloc(ptr c.Pointer, size c.SizeT, caps c.Uint32T) c.Pointer

/**
 * @brief Allocate an aligned chunk of memory which has the given capabilities
 *
 * Equivalent semantics to libc aligned_alloc(), for capability-aware memory.
 * @param alignment  How the pointer received needs to be aligned
 *                   must be a power of two
 * @param size Size, in bytes, of the amount of memory to allocate
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory to be returned
 *
 * @return A pointer to the memory allocated on success, NULL on failure
 *
 *
 */
//go:linkname HeapCapsAlignedAlloc C.heap_caps_aligned_alloc
func HeapCapsAlignedAlloc(alignment c.SizeT, size c.SizeT, caps c.Uint32T) c.Pointer

/**
 * @brief Used to deallocate memory previously allocated with heap_caps_aligned_alloc
 *
 * @param ptr Pointer to the memory allocated
 * @note This function is deprecated, please consider using heap_caps_free() instead
 */
//go:linkname HeapCapsAlignedFree C.heap_caps_aligned_free
func HeapCapsAlignedFree(ptr c.Pointer)

/**
 * @brief Allocate an aligned chunk of memory which has the given capabilities. The initialized value in the memory is set to zero.
 *
 * @param alignment  How the pointer received needs to be aligned
 *                   must be a power of two
 * @param n    Number of continuing chunks of memory to allocate
 * @param size Size, in bytes, of a chunk of memory to allocate
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory to be returned
 *
 * @return A pointer to the memory allocated on success, NULL on failure
 *
 */
//go:linkname HeapCapsAlignedCalloc C.heap_caps_aligned_calloc
func HeapCapsAlignedCalloc(alignment c.SizeT, n c.SizeT, size c.SizeT, caps c.Uint32T) c.Pointer

/**
 * @brief Allocate a chunk of memory which has the given capabilities. The initialized value in the memory is set to zero.
 *
 * Equivalent semantics to libc calloc(), for capability-aware memory.
 *
 * In IDF, ``calloc(p)`` is equivalent to ``heap_caps_calloc(p, MALLOC_CAP_8BIT)``.
 *
 * @param n    Number of continuing chunks of memory to allocate
 * @param size Size, in bytes, of a chunk of memory to allocate
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory to be returned
 *
 * @return A pointer to the memory allocated on success, NULL on failure
 */
//go:linkname HeapCapsCalloc C.heap_caps_calloc
func HeapCapsCalloc(n c.SizeT, size c.SizeT, caps c.Uint32T) c.Pointer

/**
 * @brief Get the total size of all the regions that have the given capabilities
 *
 * This function takes all regions capable of having the given capabilities allocated in them
 * and adds up the total space they have.
 *
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory
 *
 * @return total size in bytes
 */
//go:linkname HeapCapsGetTotalSize C.heap_caps_get_total_size
func HeapCapsGetTotalSize(caps c.Uint32T) c.SizeT

/**
 * @brief Get the total free size of all the regions that have the given capabilities
 *
 * This function takes all regions capable of having the given capabilities allocated in them
 * and adds up the free space they have.
 *
 * @note Note that because of heap fragmentation it is probably not possible to allocate a single block of memory
 * of this size. Use heap_caps_get_largest_free_block() for this purpose.

 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory
 *
 * @return Amount of free bytes in the regions
 */
//go:linkname HeapCapsGetFreeSize C.heap_caps_get_free_size
func HeapCapsGetFreeSize(caps c.Uint32T) c.SizeT

/**
 * @brief Get the total minimum free memory of all regions with the given capabilities
 *
 * This adds all the low watermarks of the regions capable of delivering the memory
 * with the given capabilities.
 *
 * @note Note the result may be less than the global all-time minimum available heap of this kind, as "low watermarks" are
 * tracked per-region. Individual regions' heaps may have reached their "low watermarks" at different points in time. However,
 * this result still gives a "worst case" indication for all-time minimum free heap.
 *
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory
 *
 * @return Amount of free bytes in the regions
 */
//go:linkname HeapCapsGetMinimumFreeSize C.heap_caps_get_minimum_free_size
func HeapCapsGetMinimumFreeSize(caps c.Uint32T) c.SizeT

/**
 * @brief Get the largest free block of memory able to be allocated with the given capabilities.
 *
 * Returns the largest value of ``s`` for which ``heap_caps_malloc(s, caps)`` will succeed.
 *
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory
 *
 * @return Size of the largest free block in bytes.
 */
//go:linkname HeapCapsGetLargestFreeBlock C.heap_caps_get_largest_free_block
func HeapCapsGetLargestFreeBlock(caps c.Uint32T) c.SizeT

/**
 * @brief Start monitoring the value of minimum_free_bytes from the moment this
 * function is called instead of from startup.
 *
 * @note This allows to detect local lows of the minimum_free_bytes value
 * that wouldn't be detected otherwise.
 *
 * @return esp_err_t ESP_OK if the function executed properly
 *                   ESP_FAIL if called when monitoring already active
 */
//go:linkname HeapCapsMonitorLocalMinimumFreeSizeStart C.heap_caps_monitor_local_minimum_free_size_start
func HeapCapsMonitorLocalMinimumFreeSizeStart() EspErrT

/**
 * @brief Stop monitoring the value of minimum_free_bytes. After this call
 * the minimum_free_bytes value calculated from startup will be returned in
 * heap_caps_get_info and heap_caps_get_minimum_free_size.
 *
 * @return esp_err_t ESP_OK if the function executed properly
 *                   ESP_FAIL if called when monitoring not active
 */
//go:linkname HeapCapsMonitorLocalMinimumFreeSizeStop C.heap_caps_monitor_local_minimum_free_size_stop
func HeapCapsMonitorLocalMinimumFreeSizeStop() EspErrT

/**
 * @brief Get heap info for all regions with the given capabilities.
 *
 * Calls multi_heap_info() on all heaps which share the given capabilities. The information returned is an aggregate
 * across all matching heaps. The meanings of fields are the same as defined for multi_heap_info_t, except that
 * ``minimum_free_bytes`` has the same caveats described in heap_caps_get_minimum_free_size().
 *
 * @param info        Pointer to a structure which will be filled with relevant
 *                    heap metadata.
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory
 *
 */
// llgo:link (*MultiHeapInfoT).HeapCapsGetInfo C.heap_caps_get_info
func (recv_ *MultiHeapInfoT) HeapCapsGetInfo(caps c.Uint32T) {
}

/**
 * @brief Print a summary of all memory with the given capabilities.
 *
 * Calls multi_heap_info on all heaps which share the given capabilities, and
 * prints a two-line summary for each, then a total summary.
 *
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory
 *
 */
//go:linkname HeapCapsPrintHeapInfo C.heap_caps_print_heap_info
func HeapCapsPrintHeapInfo(caps c.Uint32T)

/**
 * @brief Check integrity of all heap memory in the system.
 *
 * Calls multi_heap_check on all heaps. Optionally print errors if heaps are corrupt.
 *
 * Calling this function is equivalent to calling heap_caps_check_integrity
 * with the caps argument set to MALLOC_CAP_INVALID.
 *
 * @param print_errors Print specific errors if heap corruption is found.
 *
 * @note Please increase the value of `CONFIG_ESP_INT_WDT_TIMEOUT_MS` when using this API
 * with PSRAM enabled.
 *
 * @return True if all heaps are valid, False if at least one heap is corrupt.
 */
//go:linkname HeapCapsCheckIntegrityAll C.heap_caps_check_integrity_all
func HeapCapsCheckIntegrityAll(print_errors bool) bool

/**
 * @brief Check integrity of all heaps with the given capabilities.
 *
 * Calls multi_heap_check on all heaps which share the given capabilities. Optionally
 * print errors if the heaps are corrupt.
 *
 * See also heap_caps_check_integrity_all to check all heap memory
 * in the system and heap_caps_check_integrity_addr to check memory
 * around a single address.
 *
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory
 * @param print_errors Print specific errors if heap corruption is found.
 *
 * @note Please increase the value of `CONFIG_ESP_INT_WDT_TIMEOUT_MS` when using this API
 * with PSRAM capability flag.
 *
 * @return True if all heaps are valid, False if at least one heap is corrupt.
 */
//go:linkname HeapCapsCheckIntegrity C.heap_caps_check_integrity
func HeapCapsCheckIntegrity(caps c.Uint32T, print_errors bool) bool

/**
 * @brief Check integrity of heap memory around a given address.
 *
 * This function can be used to check the integrity of a single region of heap memory,
 * which contains the given address.
 *
 * This can be useful if debugging heap integrity for corruption at a known address,
 * as it has a lower overhead than checking all heap regions. Note that if the corrupt
 * address moves around between runs (due to timing or other factors) then this approach
 * won't work, and you should call heap_caps_check_integrity or
 * heap_caps_check_integrity_all instead.
 *
 * @note The entire heap region around the address is checked, not only the adjacent
 * heap blocks.
 *
 * @param addr Address in memory. Check for corruption in region containing this address.
 * @param print_errors Print specific errors if heap corruption is found.
 *
 * @return True if the heap containing the specified address is valid,
 * False if at least one heap is corrupt or the address doesn't belong to a heap region.
 */
//go:linkname HeapCapsCheckIntegrityAddr C.heap_caps_check_integrity_addr
func HeapCapsCheckIntegrityAddr(addr c.IntptrT, print_errors bool) bool

/**
 * @brief Enable malloc() in external memory and set limit below which
 *        malloc() attempts are placed in internal memory.
 *
 * When external memory is in use, the allocation strategy is to initially try to
 * satisfy smaller allocation requests with internal memory and larger requests
 * with external memory. This sets the limit between the two, as well as generally
 * enabling allocation in external memory.
 *
 * @param limit       Limit, in bytes.
 */
//go:linkname HeapCapsMallocExtmemEnable C.heap_caps_malloc_extmem_enable
func HeapCapsMallocExtmemEnable(limit c.SizeT)

/**
 * @brief Allocate a chunk of memory as preference in decreasing order.
 *
 * @attention The variable parameters are bitwise OR of MALLOC_CAP_* flags indicating the type of memory.
 *            This API prefers to allocate memory with the first parameter. If failed, allocate memory with
 *            the next parameter. It will try in this order until allocating a chunk of memory successfully
 *            or fail to allocate memories with any of the parameters.
 *
 * @param size Size, in bytes, of the amount of memory to allocate
 * @param num Number of variable parameters
 *
 * @return A pointer to the memory allocated on success, NULL on failure
 */
//go:linkname HeapCapsMallocPrefer C.heap_caps_malloc_prefer
func HeapCapsMallocPrefer(size c.SizeT, num c.SizeT, __llgo_va_list ...interface{}) c.Pointer

/**
 * @brief Reallocate a chunk of memory as preference in decreasing order.
 *
 * @param ptr Pointer to previously allocated memory, or NULL for a new allocation.
 * @param size Size of the new buffer requested, or 0 to free the buffer.
 * @param num Number of variable parameters
 *
 * @return Pointer to a new buffer of size 'size', or NULL if allocation failed.
 */
//go:linkname HeapCapsReallocPrefer C.heap_caps_realloc_prefer
func HeapCapsReallocPrefer(ptr c.Pointer, size c.SizeT, num c.SizeT, __llgo_va_list ...interface{}) c.Pointer

/**
 * @brief Allocate a chunk of memory as preference in decreasing order.
 *
 * @param n    Number of continuing chunks of memory to allocate
 * @param size Size, in bytes, of a chunk of memory to allocate
 * @param num  Number of variable parameters
 *
 * @return A pointer to the memory allocated on success, NULL on failure
 */
//go:linkname HeapCapsCallocPrefer C.heap_caps_calloc_prefer
func HeapCapsCallocPrefer(n c.SizeT, size c.SizeT, num c.SizeT, __llgo_va_list ...interface{}) c.Pointer

/**
 * @brief Dump the full structure of all heaps with matching capabilities.
 *
 * Prints a large amount of output to serial (because of locking limitations,
 * the output bypasses stdout/stderr). For each (variable sized) block
 * in each matching heap, the following output is printed on a single line:
 *
 * - Block address (the data buffer returned by malloc is 4 bytes after this
 *   if heap debugging is set to Basic, or 8 bytes otherwise).
 * - Data size (the data size may be larger than the size requested by malloc,
 *   either due to heap fragmentation or because of heap debugging level).
 * - Address of next block in the heap.
 * - If the block is free, the address of the next free block is also printed.
 *
 * @param caps        Bitwise OR of MALLOC_CAP_* flags indicating the type
 *                    of memory
 */
//go:linkname HeapCapsDump C.heap_caps_dump
func HeapCapsDump(caps c.Uint32T)

/**
 * @brief Dump the full structure of all heaps.
 *
 * Covers all registered heaps. Prints a large amount of output to serial.
 *
 * Output is the same as for heap_caps_dump.
 *
 */
//go:linkname HeapCapsDumpAll C.heap_caps_dump_all
func HeapCapsDumpAll()

/**
 * @brief Return the size that a particular pointer was allocated with.
 *
 * @param ptr Pointer to currently allocated heap memory. Must be a pointer value previously
 * returned by heap_caps_malloc, malloc, calloc, etc. and not yet freed.
 *
 * @note The app will crash with an assertion failure if the pointer is not valid.
 *
 * @return Size of the memory allocated at this block.
 *
 */
//go:linkname HeapCapsGetAllocatedSize C.heap_caps_get_allocated_size
func HeapCapsGetAllocatedSize(ptr c.Pointer) c.SizeT

/**
 * @brief Structure used to store heap related data passed to
 * the walker callback function
 */

type WalkerHeapInfo struct {
	Start c.IntptrT
	End   c.IntptrT
}
type WalkerHeapIntoT WalkerHeapInfo

/**
 * @brief Structure used to store block related data passed to
 * the walker callback function
 */

type WalkerBlockInfo struct {
	Ptr  c.Pointer
	Size c.SizeT
	Used bool
}
type WalkerBlockInfoT WalkerBlockInfo

// llgo:type C
type HeapCapsWalkerCbT func(WalkerHeapIntoT, WalkerBlockInfoT, c.Pointer) bool

/**
 * @brief Function called to walk through the heaps with the given set of capabilities
 *
 * @param caps The set of capabilities assigned to the heaps to walk through
 * @param walker_func Callback called for each block of the heaps being traversed
 * @param user_data Opaque pointer to user defined data
 */
//go:linkname HeapCapsWalk C.heap_caps_walk
func HeapCapsWalk(caps c.Uint32T, walker_func HeapCapsWalkerCbT, user_data c.Pointer)

/**
 * @brief Function called to walk through all heaps defined by the heap component
 *
 * @param walker_func Callback called for each block of the heaps being traversed
 * @param user_data Opaque pointer to user defined data
 */
//go:linkname HeapCapsWalkAll C.heap_caps_walk_all
func HeapCapsWalkAll(walker_func HeapCapsWalkerCbT, user_data c.Pointer)
