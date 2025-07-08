package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SOC_MEMORY_TYPE_NO_PRIOS = 3

/* Type descriptor holds a description for a particular type of memory on a particular SoC.
 */

type SocMemoryTypeDescT struct {
	Name *c.Char
	Caps [3]c.Uint32T
}

/* Region descriptor holds a description for a particular region of memory on a particular SoC.
 */

type SocMemoryRegionT struct {
	Start        c.IntptrT
	Size         c.SizeT
	Type         c.SizeT
	IramAddress  c.IntptrT
	StartupStack bool
}

/*
Region descriptor holds a description for a particular region of

	memory reserved on this SoC for a particular use (ie not available
	for stack/heap usage.)
*/
type SocReservedRegionT struct {
	Start c.IntptrT
	End   c.IntptrT
}

/* Return available memory regions for this SoC. Each available memory
 * region is a contiguous piece of memory which is not being used by
 * static data, used by ROM code, or reserved by a component using
 * the SOC_RESERVE_MEMORY_REGION() macro.
 *
 * This result is soc_memory_regions[] minus all regions reserved
 * via the SOC_RESERVE_MEMORY_REGION() macro (which may also split
 * some regions up.)
 *
 * At startup, all available memory returned by this function is
 * registered as heap space.
 *
 * @note OS-level startup function only, not recommended to call from
 * app code.
 *
 * @param regions Pointer to an array for reading available regions into.
 * Size of the array should be at least the result of
 * soc_get_available_memory_region_max_count(). Entries in the array
 * will be ordered by memory address.
 *
 * @return Number of entries copied to 'regions'. Will be no greater than
 * the result of soc_get_available_memory_region_max_count().
 */
// llgo:link (*SocMemoryRegionT).SocGetAvailableMemoryRegions C.soc_get_available_memory_regions
func (recv_ *SocMemoryRegionT) SocGetAvailableMemoryRegions() c.SizeT {
	return 0
}

/* Return the maximum number of available memory regions which could be
 * returned by soc_get_available_memory_regions(). Used to size the
 * array passed to that function.
 */
//go:linkname SocGetAvailableMemoryRegionMaxCount C.soc_get_available_memory_region_max_count
func SocGetAvailableMemoryRegionMaxCount() c.SizeT
