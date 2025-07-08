package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SPINLOCK_FREE = 0xB33FFFFF
const SPINLOCK_NO_WAIT = 0
const SPINLOCK_OWNER_ID_0 = 0xCDCD
const SPINLOCK_OWNER_ID_1 = 0xABAB

type SpinlockT struct {
	Owner c.Uint32T
	Count c.Uint32T
}
