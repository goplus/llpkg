package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const I_D_SRAM_SEGMENT_SIZE = 0x10000
const I_D_SPLIT_LINE_ALIGN = 0x100
const I_D_SPLIT_LINE_SHIFT = 0x8
const I_FAULT_ADDR_SHIFT = 0x2
const D_FAULT_ADDR_SHIFT = 0x4
const IRAM0_VIOLATE_STATUS_ADDR_OFFSET = 0x40000000
const DRAM0_VIOLATE_STATUS_ADDR_OFFSET = 0x3C000000
const SENSITIVE_CORE_X_ICACHE_PMS_CONSTRAIN_SRAM_WORLD_X_R = 0x1
const SENSITIVE_CORE_X_ICACHE_PMS_CONSTRAIN_SRAM_WORLD_X_W = 0x2
const SENSITIVE_CORE_X_ICACHE_PMS_CONSTRAIN_SRAM_WORLD_X_F = 0x4
const SENSITIVE_CORE_X_IRAM0_PMS_CONSTRAIN_SRAM_WORLD_X_R = 0x1
const SENSITIVE_CORE_X_IRAM0_PMS_CONSTRAIN_SRAM_WORLD_X_W = 0x2
const SENSITIVE_CORE_X_IRAM0_PMS_CONSTRAIN_SRAM_WORLD_X_F = 0x4
const SENSITIVE_CORE_X_DRAM0_PMS_CONSTRAIN_SRAM_WORLD_X_R = 0x1
const SENSITIVE_CORE_X_DRAM0_PMS_CONSTRAIN_SRAM_WORLD_X_W = 0x2
const SENSITIVE_CORE_X_PIF_PMS_CONSTRAIN_RTCFAST_WORLD_X_W = 0x1
const SENSITIVE_CORE_X_PIF_PMS_CONSTRAIN_RTCFAST_WORLD_X_R = 0x2
const SENSITIVE_CORE_X_PIF_PMS_CONSTRAIN_RTCFAST_WORLD_X_F = 0x4

/*
 * PMS register configuration structure for I/D splitting address.
 * Category bits define the splitting address being below, inside or above specific memory level range:
 *  - for details of ESP32S3 memory layout, see 725_mem_map.* documents
 *  - for category bits settings, see MEMP_HAL_CORE_X_IRAM0_DRAM0_DMA_SRAM_CATEGORY_BITS*
 *    (components/hal/include/hal/memprot_types.h)
 *  - for details on assembling full splitting address
 *    see function memprot_ll_get_split_addr_from_reg() (components/hal/esp32s3/include/hal/memprot_ll.h)
 */

type ConstrainRegFieldsT struct {
	Val c.Uint32T
}
