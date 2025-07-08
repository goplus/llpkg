package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CONFIG_HEAP_TRACING_STACK_DEPTH = 0

type HeapTraceModeT c.Int

const (
	HEAP_TRACE_ALL   HeapTraceModeT = 0
	HEAP_TRACE_LEAKS HeapTraceModeT = 1
)

/**
 * @brief Trace record data type. Stores information about an allocated region of memory.
 */

type HeapTraceRecordT struct {
	Ccount    c.Uint32T
	Address   c.Pointer
	Size      c.SizeT
	AllocedBy [0]c.Pointer
	FreedBy   [0]c.Pointer
}

/**
 * @brief Stores information about the result of a heap trace.
 */

type HeapTraceSummaryT struct {
	Mode             HeapTraceModeT
	TotalAllocations c.SizeT
	TotalFrees       c.SizeT
	Count            c.SizeT
	Capacity         c.SizeT
	HighWaterMark    c.SizeT
	HasOverflowed    c.SizeT
}

/**
 * @brief Initialise heap tracing in standalone mode.
 *
 * This function must be called before any other heap tracing functions.
 *
 * To disable heap tracing and allow the buffer to be freed, stop tracing and then call heap_trace_init_standalone(NULL, 0);
 *
 * @param record_buffer Provide a buffer to use for heap trace data.
 * Note: External RAM is allowed, but it prevents recording allocations made from ISR's.
 * @param num_records Size of the heap trace buffer, as number of record structures.
 * @return
 *  - ESP_ERR_NOT_SUPPORTED Project was compiled without heap tracing enabled in menuconfig.
 *  - ESP_ERR_INVALID_STATE Heap tracing is currently in progress.
 *  - ESP_OK Heap tracing initialised successfully.
 */
// llgo:link (*HeapTraceRecordT).HeapTraceInitStandalone C.heap_trace_init_standalone
func (recv_ *HeapTraceRecordT) HeapTraceInitStandalone(num_records c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Initialise heap tracing in host-based mode.
 *
 * This function must be called before any other heap tracing functions.
 *
 * @return
 *  - ESP_ERR_INVALID_STATE Heap tracing is currently in progress.
 *  - ESP_OK Heap tracing initialised successfully.
 */
//go:linkname HeapTraceInitTohost C.heap_trace_init_tohost
func HeapTraceInitTohost() EspErrT

/**
 * @brief Start heap tracing. All heap allocations & frees will be traced, until heap_trace_stop() is called.
 *
 * @note heap_trace_init_standalone() must be called to provide a valid buffer, before this function is called.
 *
 * @note Calling this function while heap tracing is running will reset the heap trace state and continue tracing.
 *
 * @param mode Mode for tracing.
 * - HEAP_TRACE_ALL means all heap allocations and frees are traced.
 * - HEAP_TRACE_LEAKS means only suspected memory leaks are traced. (When memory is freed, the record is removed from the trace buffer.)
 * @return
 * - ESP_ERR_NOT_SUPPORTED Project was compiled without heap tracing enabled in menuconfig.
 * - ESP_ERR_INVALID_STATE A non-zero-length buffer has not been set via heap_trace_init_standalone().
 * - ESP_OK Tracing is started.
 */
// llgo:link HeapTraceModeT.HeapTraceStart C.heap_trace_start
func (recv_ HeapTraceModeT) HeapTraceStart() EspErrT {
	return 0
}

/**
 * @brief Stop heap tracing.
 *
 * @return
 * - ESP_ERR_NOT_SUPPORTED Project was compiled without heap tracing enabled in menuconfig.
 * - ESP_ERR_INVALID_STATE Heap tracing was not in progress.
 * - ESP_OK Heap tracing stopped.
 */
//go:linkname HeapTraceStop C.heap_trace_stop
func HeapTraceStop() EspErrT

/**
 * @brief Pause heap tracing of allocations.
 *
 * @note This function puts the heap tracing in the state where the new allocations
 * will no longer be traced but the free will still be. This can be used to e.g.,
 * strategically monitor a set of allocations to make sure each of them will get freed
 * without polluting the list of records with unwanted allocations.
 *
 * @return
 * - ESP_ERR_NOT_SUPPORTED Project was compiled without heap tracing enabled in menuconfig.
 * - ESP_ERR_INVALID_STATE Heap tracing was not in progress.
 * - ESP_OK Heap tracing paused.
 */
//go:linkname HeapTraceAllocPause C.heap_trace_alloc_pause
func HeapTraceAllocPause() EspErrT

/**
 * @brief Resume heap tracing which was previously stopped.
 *
 * Unlike heap_trace_start(), this function does not clear the
 * buffer of any pre-existing trace records.
 *
 * The heap trace mode is the same as when heap_trace_start() was
 * last called (or HEAP_TRACE_ALL if heap_trace_start() was never called).
 *
 * @return
 * - ESP_ERR_NOT_SUPPORTED Project was compiled without heap tracing enabled in menuconfig.
 * - ESP_ERR_INVALID_STATE Heap tracing was already started.
 * - ESP_OK Heap tracing resumed.
 */
//go:linkname HeapTraceResume C.heap_trace_resume
func HeapTraceResume() EspErrT

/**
 * @brief Return number of records in the heap trace buffer
 *
 * It is safe to call this function while heap tracing is running.
 */
//go:linkname HeapTraceGetCount C.heap_trace_get_count
func HeapTraceGetCount() c.SizeT

/**
 * @brief Return a raw record from the heap trace buffer
 *
 * @note It is safe to call this function while heap tracing is
 * running, however in HEAP_TRACE_LEAK mode record indexing may
 * skip entries unless heap tracing is stopped first.
 *
 * @param index Index (zero-based) of the record to return.
 * @param[out] record Record where the heap trace record will be copied.
 * @return
 * - ESP_ERR_NOT_SUPPORTED Project was compiled without heap tracing enabled in menuconfig.
 * - ESP_ERR_INVALID_STATE Heap tracing was not initialised.
 * - ESP_ERR_INVALID_ARG Index is out of bounds for current heap trace record count.
 * - ESP_OK Record returned successfully.
 */
//go:linkname HeapTraceGet C.heap_trace_get
func HeapTraceGet(index c.SizeT, record *HeapTraceRecordT) EspErrT

/**
 * @brief Dump heap trace record data to stdout
 *
 * @note It is safe to call this function while heap tracing is
 * running, however in HEAP_TRACE_LEAK mode the dump may skip
 * entries unless heap tracing is stopped first.
 */
//go:linkname HeapTraceDump C.heap_trace_dump
func HeapTraceDump()

/**
 * @brief Dump heap trace from the memory of the capabilities passed as parameter.
 *
 * @param caps Capability(ies) of the memory from which to dump the trace.
 * Set MALLOC_CAP_INTERNAL to dump heap trace data from internal memory.
 * Set MALLOC_CAP_SPIRAM to dump heap trace data from PSRAM.
 * Set both to dump both heap trace data.
 */
//go:linkname HeapTraceDumpCaps C.heap_trace_dump_caps
func HeapTraceDumpCaps(caps c.Uint32T)

/**
 * @brief Get summary information about the result of a heap trace
 *
 *  @note It is safe to call this function while heap tracing is running.
 */
// llgo:link (*HeapTraceSummaryT).HeapTraceSummary C.heap_trace_summary
func (recv_ *HeapTraceSummaryT) HeapTraceSummary() EspErrT {
	return 0
}
