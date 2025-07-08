package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* The functions below are designed to be used in interrupt/panic handler
 * In case using them in user's code put them into critical section.*/
//go:linkname EspHwStackGuardMonitorStart C.esp_hw_stack_guard_monitor_start
func EspHwStackGuardMonitorStart()

//go:linkname EspHwStackGuardMonitorStop C.esp_hw_stack_guard_monitor_stop
func EspHwStackGuardMonitorStop()

//go:linkname EspHwStackGuardSetBounds C.esp_hw_stack_guard_set_bounds
func EspHwStackGuardSetBounds(sp_min c.Uint32T, sp_max c.Uint32T)

//go:linkname EspHwStackGuardGetBounds C.esp_hw_stack_guard_get_bounds
func EspHwStackGuardGetBounds(core_id c.Uint32T, sp_min *c.Uint32T, sp_max *c.Uint32T)

//go:linkname EspHwStackGuardGetFiredCpu C.esp_hw_stack_guard_get_fired_cpu
func EspHwStackGuardGetFiredCpu() c.Uint32T

//go:linkname EspHwStackGuardGetPc C.esp_hw_stack_guard_get_pc
func EspHwStackGuardGetPc(core_id c.Uint32T) c.Uint32T
