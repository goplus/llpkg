package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* this bitfield is start from the LSB!!! */

type LldescS struct {
	Size   c.Uint32T
	Length c.Uint32T
	Offset c.Uint32T
	Sosf   c.Uint32T
	Eof    c.Uint32T
	Owner  c.Uint32T
	Buf    *c.Uint8T
}
type LldescT LldescS
