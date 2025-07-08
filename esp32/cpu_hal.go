package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type WatchpointTriggerT c.Int

const (
	WATCHPOINT_TRIGGER_ON_RO WatchpointTriggerT = 0
	WATCHPOINT_TRIGGER_ON_WO WatchpointTriggerT = 1
	WATCHPOINT_TRIGGER_ON_RW WatchpointTriggerT = 2
)
