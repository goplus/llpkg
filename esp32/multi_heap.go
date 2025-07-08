package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MultiHeapInfo struct {
	Unused [8]uint8
}
type MultiHeapHandleT *MultiHeapInfo

/**
 * @brief allocate a chunk of memory with specific alignment
 *
 * @param heap  Handle to a registered heap.
 * @param size  size in bytes of memory chunk
 * @param alignment  how the memory must be aligned
 *
 * @return pointer to the memory allocated, NULL on failure
 */
//go:linkname MultiHeapAlignedAlloc C.multi_heap_aligned_alloc
func MultiHeapAlignedAlloc(heap MultiHeapHandleT, size c.SizeT, alignment c.SizeT) c.Pointer

/** @brief malloc() a buffer in a given heap
 *
 * Semantics are the same as standard malloc(), only the returned buffer will be allocated in the specified heap.
 *
 * @param heap Handle to a registered heap.
 * @param size Size of desired buffer.
 *
 * @return Pointer to new memory, or NULL if allocation fails.
 */
//go:linkname MultiHeapMalloc C.multi_heap_malloc
func MultiHeapMalloc(heap MultiHeapHandleT, size c.SizeT) c.Pointer

/** @brief free() a buffer aligned in a given heap.
 *
 * @param heap Handle to a registered heap.
 * @param p NULL, or a pointer previously returned from multi_heap_aligned_alloc() for the same heap.
 * @note This function is deprecated, consider using multi_heap_free() instead
 */
//go:linkname MultiHeapAlignedFree C.multi_heap_aligned_free
func MultiHeapAlignedFree(heap MultiHeapHandleT, p c.Pointer)

/** @brief free() a buffer in a given heap.
 *
 * Semantics are the same as standard free(), only the argument 'p' must be NULL or have been allocated in the specified heap.
 *
 * @param heap Handle to a registered heap.
 * @param p NULL, or a pointer previously returned from multi_heap_malloc() or multi_heap_realloc() for the same heap.
 */
//go:linkname MultiHeapFree C.multi_heap_free
func MultiHeapFree(heap MultiHeapHandleT, p c.Pointer)

/** @brief realloc() a buffer in a given heap.
 *
 * Semantics are the same as standard realloc(), only the argument 'p' must be NULL or have been allocated in the specified heap.
 *
 * @param heap Handle to a registered heap.
 * @param p NULL, or a pointer previously returned from multi_heap_malloc() or multi_heap_realloc() for the same heap.
 * @param size Desired new size for buffer.
 *
 * @return New buffer of 'size' containing contents of 'p', or NULL if reallocation failed.
 */
//go:linkname MultiHeapRealloc C.multi_heap_realloc
func MultiHeapRealloc(heap MultiHeapHandleT, p c.Pointer, size c.SizeT) c.Pointer

/** @brief Return the size that a particular pointer was allocated with.
 *
 * @param heap Handle to a registered heap.
 * @param p Pointer, must have been previously returned from multi_heap_malloc() or multi_heap_realloc() for the same heap.
 *
 * @return Size of the memory allocated at this block. May be more than the original size argument, due
 * to padding and minimum block sizes.
 */
//go:linkname MultiHeapGetAllocatedSize C.multi_heap_get_allocated_size
func MultiHeapGetAllocatedSize(heap MultiHeapHandleT, p c.Pointer) c.SizeT

/** @brief Register a new heap for use
 *
 * This function initialises a heap at the specified address, and returns a handle for future heap operations.
 *
 * There is no equivalent function for deregistering a heap - if all blocks in the heap are free, you can immediately start using the memory for other purposes.
 *
 * @param start Start address of the memory to use for a new heap.
 * @param size Size (in bytes) of the new heap.
 *
 * @return Handle of a new heap ready for use, or NULL if the heap region was too small to be initialised.
 */
//go:linkname MultiHeapRegister C.multi_heap_register
func MultiHeapRegister(start c.Pointer, size c.SizeT) MultiHeapHandleT

/** @brief Associate a private lock pointer with a heap
 *
 * The lock argument is supplied to the MULTI_HEAP_LOCK() and MULTI_HEAP_UNLOCK() macros, defined in multi_heap_platform.h.
 *
 * The lock in question must be recursive.
 *
 * When the heap is first registered, the associated lock is NULL.
 *
 * @param heap Handle to a registered heap.
 * @param lock Optional pointer to a locking structure to associate with this heap.
 */
//go:linkname MultiHeapSetLock C.multi_heap_set_lock
func MultiHeapSetLock(heap MultiHeapHandleT, lock c.Pointer)

/** @brief Dump heap information to stdout
 *
 * For debugging purposes, this function dumps information about every block in the heap to stdout.
 *
 * @param heap Handle to a registered heap.
 */
//go:linkname MultiHeapDump C.multi_heap_dump
func MultiHeapDump(heap MultiHeapHandleT)

/** @brief Check heap integrity
 *
 * Walks the heap and checks all heap data structures are valid. If any errors are detected, an error-specific message
 * can be optionally printed to stderr. Print behaviour can be overridden at compile time by defining
 * MULTI_CHECK_FAIL_PRINTF in multi_heap_platform.h.
 *
 * @note This function is not thread-safe as it sets a global variable with the value of print_errors.
 *
 * @param heap Handle to a registered heap.
 * @param print_errors If true, errors will be printed to stderr.
 * @return true if heap is valid, false otherwise.
 */
//go:linkname MultiHeapCheck C.multi_heap_check
func MultiHeapCheck(heap MultiHeapHandleT, print_errors bool) bool

/** @brief Return free heap size
 *
 * Returns the number of bytes available in the heap.
 *
 * Equivalent to the total_free_bytes member returned by multi_heap_get_heap_info().
 *
 * Note that the heap may be fragmented, so the actual maximum size for a single malloc() may be lower. To know this
 * size, see the largest_free_block member returned by multi_heap_get_heap_info().
 *
 * @param heap Handle to a registered heap.
 * @return Number of free bytes.
 */
//go:linkname MultiHeapFreeSize C.multi_heap_free_size
func MultiHeapFreeSize(heap MultiHeapHandleT) c.SizeT

/** @brief Return the lifetime minimum free heap size
 *
 * Equivalent to the minimum_free_bytes member returned by multi_heap_get_info().
 *
 * Returns the lifetime "low watermark" of possible values returned from multi_free_heap_size(), for the specified
 * heap.
 *
 * @param heap Handle to a registered heap.
 * @return Number of free bytes.
 */
//go:linkname MultiHeapMinimumFreeSize C.multi_heap_minimum_free_size
func MultiHeapMinimumFreeSize(heap MultiHeapHandleT) c.SizeT

/** @brief Structure to access heap metadata via multi_heap_get_info */

type MultiHeapInfoT struct {
	TotalFreeBytes      c.SizeT
	TotalAllocatedBytes c.SizeT
	LargestFreeBlock    c.SizeT
	MinimumFreeBytes    c.SizeT
	AllocatedBlocks     c.SizeT
	FreeBlocks          c.SizeT
	TotalBlocks         c.SizeT
}

/** @brief Return metadata about a given heap
 *
 * Fills a multi_heap_info_t structure with information about the specified heap.
 *
 * @param heap Handle to a registered heap.
 * @param info Pointer to a structure to fill with heap metadata.
 */
//go:linkname MultiHeapGetInfo C.multi_heap_get_info
func MultiHeapGetInfo(heap MultiHeapHandleT, info *MultiHeapInfoT)

/**
 * @brief Perform an aligned allocation from the provided offset
 *
 * @param heap The heap in which to perform the allocation
 * @param size The size of the allocation
 * @param alignment How the memory must be aligned
 * @param offset The offset at which the alignment should start
 * @return void* The ptr to the allocated memory
 */
//go:linkname MultiHeapAlignedAllocOffs C.multi_heap_aligned_alloc_offs
func MultiHeapAlignedAllocOffs(heap MultiHeapHandleT, size c.SizeT, alignment c.SizeT, offset c.SizeT) c.Pointer

/**
 * @brief Reset the minimum_free_bytes value (setting it to free_bytes) and return the former value
 *
 * @param heap The heap in which the reset is taking place
 * @return size_t the value of minimum_free_bytes before it is reset
 */
//go:linkname MultiHeapResetMinimumFreeBytes C.multi_heap_reset_minimum_free_bytes
func MultiHeapResetMinimumFreeBytes(heap MultiHeapHandleT) c.SizeT

/**
 * @brief Set the value of minimum_free_bytes to new_minimum_free_bytes_value or keep
 * the current value of minimum_free_bytes if it is smaller than new_minimum_free_bytes_value
 *
 * @param heap The heap in which the restore is taking place
 * @param new_minimum_free_bytes_value The value to restore the minimum_free_bytes to
 */
//go:linkname MultiHeapRestoreMinimumFreeBytes C.multi_heap_restore_minimum_free_bytes
func MultiHeapRestoreMinimumFreeBytes(heap MultiHeapHandleT, new_minimum_free_bytes_value c.SizeT)

// llgo:type C
type MultiHeapWalkerCbT func(c.Pointer, c.SizeT, c.Int, c.Pointer) bool

/**
 * @brief Call the tlsf_walk_pool function of the heap given as parameter with
 * the walker function passed as parameter
 *
 * @param heap The heap to traverse
 * @param walker_func The walker to trigger on each block of the heap
 * @param user_data Opaque pointer to user defined data
 */
//go:linkname MultiHeapWalk C.multi_heap_walk
func MultiHeapWalk(heap MultiHeapHandleT, walker_func MultiHeapWalkerCbT, user_data c.Pointer)
