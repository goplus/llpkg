package libxml2

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

/*
 * Originally declared in xmlversion.h which is generated
 */
//go:linkname CheckVersion C.xmlCheckVersion
func CheckVersion(version c.Int)
