package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_SYSTEM_INIT_STAGE_CORE = 0
const ESP_SYSTEM_INIT_STAGE_SECONDARY = 1

// llgo:type C
type SysStartupFnT func()

//go:linkname StartupResumeOtherCores C.startup_resume_other_cores
func StartupResumeOtherCores()

/**
 * Internal structure describing ESP_SYSTEM_INIT_FN startup functions
 */

type EspSystemInitFnT struct {
	Fn    c.Pointer
	Cores c.Uint16T
	Stage c.Uint16T
}
