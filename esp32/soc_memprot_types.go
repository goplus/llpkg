package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspMprotMemT c.Int

const (
	MEMPROT_TYPE_NONE          EspMprotMemT = 0
	MEMPROT_TYPE_IRAM0_SRAM    EspMprotMemT = 1
	MEMPROT_TYPE_DRAM0_SRAM    EspMprotMemT = 2
	MEMPROT_TYPE_IRAM0_RTCFAST EspMprotMemT = 4
	MEMPROT_TYPE_ALL           EspMprotMemT = 2147483647
	MEMPROT_TYPE_INVALID       EspMprotMemT = -2147483648
	MEMPROT_TYPE_IRAM0_ANY     EspMprotMemT = 5
)

type EspMprotSplitAddrT c.Int

const (
	MEMPROT_SPLIT_ADDR_NONE             EspMprotSplitAddrT = 0
	MEMPROT_SPLIT_ADDR_IRAM0_DRAM0      EspMprotSplitAddrT = 1
	MEMPROT_SPLIT_ADDR_IRAM0_LINE_0     EspMprotSplitAddrT = 2
	MEMPROT_SPLIT_ADDR_IRAM0_LINE_1     EspMprotSplitAddrT = 4
	MEMPROT_SPLIT_ADDR_DRAM0_DMA_LINE_0 EspMprotSplitAddrT = 8
	MEMPROT_SPLIT_ADDR_DRAM0_DMA_LINE_1 EspMprotSplitAddrT = 16
	MEMPROT_SPLIT_ADDR_ALL              EspMprotSplitAddrT = 2147483647
	MEMPROT_SPLIT_ADDR_INVALID          EspMprotSplitAddrT = -2147483648
	MEMPROT_SPLIT_ADDR_MAIN             EspMprotSplitAddrT = 1
)

type EspMprotPmsAreaT c.Int

const (
	MEMPROT_PMS_AREA_NONE             EspMprotPmsAreaT = 0
	MEMPROT_PMS_AREA_IRAM0_0          EspMprotPmsAreaT = 1
	MEMPROT_PMS_AREA_IRAM0_1          EspMprotPmsAreaT = 2
	MEMPROT_PMS_AREA_IRAM0_2          EspMprotPmsAreaT = 4
	MEMPROT_PMS_AREA_IRAM0_3          EspMprotPmsAreaT = 8
	MEMPROT_PMS_AREA_DRAM0_0          EspMprotPmsAreaT = 16
	MEMPROT_PMS_AREA_DRAM0_1          EspMprotPmsAreaT = 32
	MEMPROT_PMS_AREA_DRAM0_2          EspMprotPmsAreaT = 64
	MEMPROT_PMS_AREA_DRAM0_3          EspMprotPmsAreaT = 128
	MEMPROT_PMS_AREA_IRAM0_RTCFAST_LO EspMprotPmsAreaT = 256
	MEMPROT_PMS_AREA_IRAM0_RTCFAST_HI EspMprotPmsAreaT = 512
	MEMPROT_PMS_AREA_ICACHE_0         EspMprotPmsAreaT = 1024
	MEMPROT_PMS_AREA_ICACHE_1         EspMprotPmsAreaT = 2048
	MEMPROT_PMS_AREA_ALL              EspMprotPmsAreaT = 2147483647
	MEMPROT_PMS_AREA_INVALID          EspMprotPmsAreaT = -2147483648
)

/**
* @brief Memory protection configuration
 */
type EspMempConfigT struct {
	InvokePanicHandler bool
	LockFeature        bool
	SplitAddr          c.Pointer
	MemTypeMask        c.Uint32T
	TargetCpuCount     c.SizeT
	TargetCpu          [2]c.Int
}
