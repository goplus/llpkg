package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const PortUSING_MPU_WRAPPERS = 0
const PortNUM_CONFIGURABLE_REGIONS = 1
const PortHAS_STACK_OVERFLOW_CHECKING = 0

// llgo:link (*StackTypeT).PxPortInitialiseStack C.pxPortInitialiseStack
func (recv_ *StackTypeT) PxPortInitialiseStack(pxCode TaskFunctionT, pvParameters c.Pointer) *StackTypeT {
	return nil
}

/* Used by heap_5.c to define the start address and size of each memory region
 * that together comprise the total FreeRTOS heap space. */

type HeapRegion struct {
	PucStartAddress *c.Uint8T
	XSizeInBytes    c.SizeT
}
type HeapRegionT HeapRegion

/* Used to pass information about the heap out of vPortGetHeapStats(). */

type XHeapStats struct {
	XAvailableHeapSpaceInBytes      c.SizeT
	XSizeOfLargestFreeBlockInBytes  c.SizeT
	XSizeOfSmallestFreeBlockInBytes c.SizeT
	XNumberOfFreeBlocks             c.SizeT
	XMinimumEverFreeBytesRemaining  c.SizeT
	XNumberOfSuccessfulAllocations  c.SizeT
	XNumberOfSuccessfulFrees        c.SizeT
}
type HeapStatsT XHeapStats

/*
 * Used to define multiple heap regions for use by heap_5.c.  This function
 * must be called before any calls to pvPortMalloc() - not creating a task,
 * queue, semaphore, mutex, software timer, event group, etc. will result in
 * pvPortMalloc being called.
 *
 * pxHeapRegions passes in an array of HeapRegion_t structures - each of which
 * defines a region of memory that can be used as the heap.  The array is
 * terminated by a HeapRegions_t structure that has a size of 0.  The region
 * with the lowest start address must appear first in the array.
 */
// llgo:link (*HeapRegionT).VPortDefineHeapRegions C.vPortDefineHeapRegions
func (recv_ *HeapRegionT) VPortDefineHeapRegions() {
}

/*
 * Returns a HeapStats_t structure filled with information about the current
 * heap state.
 */
// llgo:link (*HeapStatsT).VPortGetHeapStats C.vPortGetHeapStats
func (recv_ *HeapStatsT) VPortGetHeapStats() {
}

/*
 * Map to the memory management routines required for the port.
 */
//go:linkname PvPortMalloc C.pvPortMalloc
func PvPortMalloc(xSize c.SizeT) c.Pointer

//go:linkname PvPortCalloc C.pvPortCalloc
func PvPortCalloc(xNum c.SizeT, xSize c.SizeT) c.Pointer

//go:linkname VPortFree C.vPortFree
func VPortFree(pv c.Pointer)

//go:linkname VPortInitialiseBlocks C.vPortInitialiseBlocks
func VPortInitialiseBlocks()

//go:linkname XPortGetFreeHeapSize C.xPortGetFreeHeapSize
func XPortGetFreeHeapSize() c.SizeT

//go:linkname XPortGetMinimumEverFreeHeapSize C.xPortGetMinimumEverFreeHeapSize
func XPortGetMinimumEverFreeHeapSize() c.SizeT

/*
 * Setup the hardware ready for the scheduler to take control.  This generally
 * sets up a tick interrupt and sets timers for the correct tick frequency.
 */
//go:linkname XPortStartScheduler C.xPortStartScheduler
func XPortStartScheduler() BaseTypeT

/*
 * Undo any hardware/ISR setup that was performed by xPortStartScheduler() so
 * the hardware is left in its original condition after the scheduler stops
 * executing.
 */
//go:linkname VPortEndScheduler C.vPortEndScheduler
func VPortEndScheduler()
