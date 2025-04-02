package libxml2

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const DOTTED_VERSION = "2.9.9"
const VERSION = 20909
const VERSION_STRING = "20909"
const VERSION_EXTRA = ""
const MODULE_EXTENSION = ".so"

/*
 * use those to be sure nothing nasty will happen if
 * your library and includes mismatch
 */
//go:linkname CheckVersion C.xmlCheckVersion
func CheckVersion(version c.Int)
