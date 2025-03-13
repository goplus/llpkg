package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

// llgo:type C
type FreeFunc func(unsafe.Pointer)

// llgo:type C
type MallocFunc func(uintptr) unsafe.Pointer

// llgo:type C
type ReallocFunc func(unsafe.Pointer, uintptr) unsafe.Pointer

// llgo:type C
type StrdupFunc func(*int8) *int8

/*
 * The way to overload the existing functions.
 * The xmlGc function have an extra entry for atomic block
 * allocations useful for garbage collected memory allocators
 */
//go:linkname MemSetup C.xmlMemSetup
func MemSetup(freeFunc FreeFunc, mallocFunc MallocFunc, reallocFunc ReallocFunc, strdupFunc StrdupFunc) c.Int

//go:linkname MemGet C.xmlMemGet
func MemGet(freeFunc FreeFunc, mallocFunc MallocFunc, reallocFunc ReallocFunc, strdupFunc StrdupFunc) c.Int

//go:linkname GcMemSetup C.xmlGcMemSetup
func GcMemSetup(freeFunc FreeFunc, mallocFunc MallocFunc, mallocAtomicFunc MallocFunc, reallocFunc ReallocFunc, strdupFunc StrdupFunc) c.Int

//go:linkname GcMemGet C.xmlGcMemGet
func GcMemGet(freeFunc FreeFunc, mallocFunc MallocFunc, mallocAtomicFunc MallocFunc, reallocFunc ReallocFunc, strdupFunc StrdupFunc) c.Int

/*
 * Initialization of the memory layer.
 */
//go:linkname InitMemory C.xmlInitMemory
func InitMemory() c.Int

/*
 * Cleanup of the memory layer.
 */
//go:linkname CleanupMemory C.xmlCleanupMemory
func CleanupMemory()

/*
 * These are specific to the XML debug memory wrapper.
 */
//go:linkname MemSize C.xmlMemSize
func MemSize(ptr unsafe.Pointer) uintptr

//go:linkname MemUsed C.xmlMemUsed
func MemUsed() c.Int

//go:linkname MemBlocks C.xmlMemBlocks
func MemBlocks() c.Int

//go:linkname MemDisplay C.xmlMemDisplay
func MemDisplay(fp *c.FILE)

//go:linkname MemDisplayLast C.xmlMemDisplayLast
func MemDisplayLast(fp *c.FILE, nbBytes c.Long)

//go:linkname MemShow C.xmlMemShow
func MemShow(fp *c.FILE, nr c.Int)

//go:linkname MemoryDump C.xmlMemoryDump
func MemoryDump()

//go:linkname MemMalloc C.xmlMemMalloc
func MemMalloc(size uintptr) unsafe.Pointer

//go:linkname MemRealloc C.xmlMemRealloc
func MemRealloc(ptr unsafe.Pointer, size uintptr) unsafe.Pointer

//go:linkname MemFree C.xmlMemFree
func MemFree(ptr unsafe.Pointer)

//go:linkname MemoryStrdup C.xmlMemoryStrdup
func MemoryStrdup(str *int8) *int8

//go:linkname MallocLoc C.xmlMallocLoc
func MallocLoc(size uintptr, file *int8, line c.Int) unsafe.Pointer

//go:linkname ReallocLoc C.xmlReallocLoc
func ReallocLoc(ptr unsafe.Pointer, size uintptr, file *int8, line c.Int) unsafe.Pointer

//go:linkname MallocAtomicLoc C.xmlMallocAtomicLoc
func MallocAtomicLoc(size uintptr, file *int8, line c.Int) unsafe.Pointer

//go:linkname MemStrdupLoc C.xmlMemStrdupLoc
func MemStrdupLoc(str *int8, file *int8, line c.Int) *int8
