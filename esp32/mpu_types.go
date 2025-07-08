package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MpuAccessT c.Int

const (
	MPU_REGION_ILLEGAL MpuAccessT = 0
	MPU_REGION_RW      MpuAccessT = 1
	MPU_REGION_X       MpuAccessT = 2
	MPU_REGION_RWX     MpuAccessT = 3
)
