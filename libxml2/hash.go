package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_xmlHashTable struct {
	Unused [8]uint8
}
type HashTable X_xmlHashTable
type HashTablePtr *HashTable

// llgo:type C
type HashDeallocator func(unsafe.Pointer, *Char)

// llgo:type C
type HashCopier func(unsafe.Pointer, *Char) unsafe.Pointer

// llgo:type C
type HashScanner func(unsafe.Pointer, unsafe.Pointer, *Char)

// llgo:type C
type HashScannerFull func(unsafe.Pointer, unsafe.Pointer, *Char, *Char, *Char)

/*
 * Constructor and destructor.
 */
//go:linkname HashCreate C.xmlHashCreate
func HashCreate(size c.Int) HashTablePtr

//go:linkname HashCreateDict C.xmlHashCreateDict
func HashCreateDict(size c.Int, dict DictPtr) HashTablePtr

//go:linkname HashFree C.xmlHashFree
func HashFree(table HashTablePtr, f HashDeallocator)

//go:linkname HashDefaultDeallocator C.xmlHashDefaultDeallocator
func HashDefaultDeallocator(entry unsafe.Pointer, name *Char)

/*
 * Add a new entry to the hash table.
 */
//go:linkname HashAddEntry C.xmlHashAddEntry
func HashAddEntry(table HashTablePtr, name *Char, userdata unsafe.Pointer) c.Int

//go:linkname HashUpdateEntry C.xmlHashUpdateEntry
func HashUpdateEntry(table HashTablePtr, name *Char, userdata unsafe.Pointer, f HashDeallocator) c.Int

//go:linkname HashAddEntry2 C.xmlHashAddEntry2
func HashAddEntry2(table HashTablePtr, name *Char, name2 *Char, userdata unsafe.Pointer) c.Int

//go:linkname HashUpdateEntry2 C.xmlHashUpdateEntry2
func HashUpdateEntry2(table HashTablePtr, name *Char, name2 *Char, userdata unsafe.Pointer, f HashDeallocator) c.Int

//go:linkname HashAddEntry3 C.xmlHashAddEntry3
func HashAddEntry3(table HashTablePtr, name *Char, name2 *Char, name3 *Char, userdata unsafe.Pointer) c.Int

//go:linkname HashUpdateEntry3 C.xmlHashUpdateEntry3
func HashUpdateEntry3(table HashTablePtr, name *Char, name2 *Char, name3 *Char, userdata unsafe.Pointer, f HashDeallocator) c.Int

/*
 * Remove an entry from the hash table.
 */
//go:linkname HashRemoveEntry C.xmlHashRemoveEntry
func HashRemoveEntry(table HashTablePtr, name *Char, f HashDeallocator) c.Int

//go:linkname HashRemoveEntry2 C.xmlHashRemoveEntry2
func HashRemoveEntry2(table HashTablePtr, name *Char, name2 *Char, f HashDeallocator) c.Int

//go:linkname HashRemoveEntry3 C.xmlHashRemoveEntry3
func HashRemoveEntry3(table HashTablePtr, name *Char, name2 *Char, name3 *Char, f HashDeallocator) c.Int

/*
 * Retrieve the userdata.
 */
//go:linkname HashLookup C.xmlHashLookup
func HashLookup(table HashTablePtr, name *Char) unsafe.Pointer

//go:linkname HashLookup2 C.xmlHashLookup2
func HashLookup2(table HashTablePtr, name *Char, name2 *Char) unsafe.Pointer

//go:linkname HashLookup3 C.xmlHashLookup3
func HashLookup3(table HashTablePtr, name *Char, name2 *Char, name3 *Char) unsafe.Pointer

//go:linkname HashQLookup C.xmlHashQLookup
func HashQLookup(table HashTablePtr, name *Char, prefix *Char) unsafe.Pointer

//go:linkname HashQLookup2 C.xmlHashQLookup2
func HashQLookup2(table HashTablePtr, name *Char, prefix *Char, name2 *Char, prefix2 *Char) unsafe.Pointer

//go:linkname HashQLookup3 C.xmlHashQLookup3
func HashQLookup3(table HashTablePtr, name *Char, prefix *Char, name2 *Char, prefix2 *Char, name3 *Char, prefix3 *Char) unsafe.Pointer

/*
 * Helpers.
 */
//go:linkname HashCopy C.xmlHashCopy
func HashCopy(table HashTablePtr, f HashCopier) HashTablePtr

//go:linkname HashSize C.xmlHashSize
func HashSize(table HashTablePtr) c.Int

//go:linkname HashScan C.xmlHashScan
func HashScan(table HashTablePtr, f HashScanner, data unsafe.Pointer)

//go:linkname HashScan3 C.xmlHashScan3
func HashScan3(table HashTablePtr, name *Char, name2 *Char, name3 *Char, f HashScanner, data unsafe.Pointer)

//go:linkname HashScanFull C.xmlHashScanFull
func HashScanFull(table HashTablePtr, f HashScannerFull, data unsafe.Pointer)

//go:linkname HashScanFull3 C.xmlHashScanFull3
func HashScanFull3(table HashTablePtr, name *Char, name2 *Char, name3 *Char, f HashScannerFull, data unsafe.Pointer)
