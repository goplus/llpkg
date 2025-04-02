package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type C14NMode c.Int

const (
	C14N10          C14NMode = 0
	C14NEXCLUSIVE10 C14NMode = 1
	C14N11          C14NMode = 2
)

// llgo:type C
type C14NIsVisibleCallback func(unsafe.Pointer, NodePtr, NodePtr) c.Int
