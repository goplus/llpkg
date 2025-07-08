package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ApmRegionPmsT c.Int

const (
	APM_REGION_PMS_X ApmRegionPmsT = 1
	APM_REGION_PMS_W ApmRegionPmsT = 2
	APM_REGION_PMS_R ApmRegionPmsT = 4
)
