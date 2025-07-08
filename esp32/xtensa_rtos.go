package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const XT_BOARD = 1
const XTENSA_PORT_VERSION = 1.7
const XTENSA_PORT_VERSION_STRING = "1.7"

// llgo:type C
type XTINTEXCHOOK func(c.Uint) c.Uint
