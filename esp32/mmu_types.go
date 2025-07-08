package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MmuMemCapsT c.Int

const (
	MMU_MEM_CAP_EXEC  MmuMemCapsT = 1
	MMU_MEM_CAP_READ  MmuMemCapsT = 2
	MMU_MEM_CAP_WRITE MmuMemCapsT = 4
	MMU_MEM_CAP_32BIT MmuMemCapsT = 8
	MMU_MEM_CAP_8BIT  MmuMemCapsT = 16
)

type MmuPageSizeT c.Int

const (
	MMU_PAGE_8KB  MmuPageSizeT = 8192
	MMU_PAGE_16KB MmuPageSizeT = 16384
	MMU_PAGE_32KB MmuPageSizeT = 32768
	MMU_PAGE_64KB MmuPageSizeT = 65536
)

type MmuVaddrT c.Int

const (
	MMU_VADDR_DATA        MmuVaddrT = 1
	MMU_VADDR_INSTRUCTION MmuVaddrT = 2
)

type MmuTargetT c.Int

const (
	MMU_TARGET_FLASH0 MmuTargetT = 1
	MMU_TARGET_PSRAM0 MmuTargetT = 2
)

type MmuTableIdT c.Int

const (
	MMU_TABLE_CORE0 MmuTableIdT = 0
	MMU_TABLE_CORE1 MmuTableIdT = 1
)
