package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type IntrHandlerT func(c.Pointer)

type IntrHandleDataT struct {
	Unused [8]uint8
}
type IntrHandleT *IntrHandleDataT
type EspIntrCpuAffinityT c.Int

const (
	ESP_INTR_CPU_AFFINITY_AUTO EspIntrCpuAffinityT = 0
	ESP_INTR_CPU_AFFINITY_0    EspIntrCpuAffinityT = 1
	ESP_INTR_CPU_AFFINITY_1    EspIntrCpuAffinityT = 2
)
