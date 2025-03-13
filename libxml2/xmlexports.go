package libxml2

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

/*
 * Originally declared in xmlversion.h which is generated
 */
//go:linkname CheckVersion C.xmlCheckVersion
func CheckVersion(version c.Int)
