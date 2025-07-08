package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MEMP_HAL_CORE_X_IRAM0_DRAM0_DMA_SRAM_CATEGORY_BITS_BELOW_SA = 0x0
const MEMP_HAL_CORE_X_IRAM0_DRAM0_DMA_SRAM_CATEGORY_BITS_EQUAL_SA = 0x2
const MEMP_HAL_CORE_X_IRAM0_DRAM0_DMA_SRAM_CATEGORY_BITS_ABOVE_SA = 0x3

type MemprotHalErrT c.Int

const (
	MEMP_HAL_OK                          MemprotHalErrT = 0
	MEMP_HAL_ERR_SPLIT_ADDR_OUT_OF_RANGE MemprotHalErrT = 2
	MEMP_HAL_ERR_SPLIT_ADDR_INVALID      MemprotHalErrT = 2
	MEMP_HAL_ERR_SPLIT_ADDR_UNALIGNED    MemprotHalErrT = 3
	MEMP_HAL_ERR_UNI_BLOCK_INVALID       MemprotHalErrT = 4
	MEMP_HAL_ERR_AREA_INVALID            MemprotHalErrT = 5
	MEMP_HAL_ERR_WORLD_INVALID           MemprotHalErrT = 6
	MEMP_HAL_ERR_CORE_INVALID            MemprotHalErrT = 7
	MEMP_HAL_FAIL                        MemprotHalErrT = -1
)

type MemprotHalWorldT c.Int

const (
	MEMP_HAL_WORLD_NONE MemprotHalWorldT = 0
	MEMP_HAL_WORLD_0    MemprotHalWorldT = 1
	MEMP_HAL_WORLD_1    MemprotHalWorldT = 16
)

type MemprotHalAreaT c.Int

const (
	MEMP_HAL_AREA_NONE MemprotHalAreaT = 0
	MEMP_HAL_AREA_LOW  MemprotHalAreaT = 1
	MEMP_HAL_AREA_HIGH MemprotHalAreaT = 2
)
