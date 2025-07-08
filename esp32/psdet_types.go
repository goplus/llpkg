package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type BrownoutResetLevelT c.Int

const (
	BROWNOUT_RESET_LEVEL_CHIP   BrownoutResetLevelT = 0
	BROWNOUT_RESET_LEVEL_SYSTEM BrownoutResetLevelT = 1
)
