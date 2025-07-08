package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const S16_F = "d"
const U16_F = "d"
const X16_F = "x"

type U8T c.Uint8T
type S8T c.Int8T
type U16T c.Uint16T
type S16T c.Int16T
type U32T c.Uint32T
type S32T c.Int32T
type SysProtT c.Int
