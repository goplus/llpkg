package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Specify the type of access allowed on a memory region.
 *
 * @param id index to the region table; on targets not SOC_MPU_CONFIGURABLE_REGIONS_SUPPORTED,
 * the region divisions is predefined in hardware which is likely reflected in LL implementation.
 * @param access type of access allowed
 */
//go:linkname MpuHalSetRegionAccess C.mpu_hal_set_region_access
func MpuHalSetRegionAccess(id c.Int, access MpuAccessT)
