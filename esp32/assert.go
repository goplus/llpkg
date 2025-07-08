package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

//go:linkname X__assertFunc C.__assert_func
func X__assertFunc(file *c.Char, line c.Int, func_ *c.Char, expr *c.Char)

//go:linkname Abort C.abort
func Abort()
