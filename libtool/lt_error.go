package libtool

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const LT_ERROR_H = 1
const (
	LT_ERROR_UNKNOWN               c.Int = 0
	LT_ERROR_DLOPEN_NOT_SUPPORTED  c.Int = 1
	LT_ERROR_INVALID_LOADER        c.Int = 2
	LT_ERROR_INIT_LOADER           c.Int = 3
	LT_ERROR_REMOVE_LOADER         c.Int = 4
	LT_ERROR_FILE_NOT_FOUND        c.Int = 5
	LT_ERROR_DEPLIB_NOT_FOUND      c.Int = 6
	LT_ERROR_NO_SYMBOLS            c.Int = 7
	LT_ERROR_CANNOT_OPEN           c.Int = 8
	LT_ERROR_CANNOT_CLOSE          c.Int = 9
	LT_ERROR_SYMBOL_NOT_FOUND      c.Int = 10
	LT_ERROR_NO_MEMORY             c.Int = 11
	LT_ERROR_INVALID_HANDLE        c.Int = 12
	LT_ERROR_BUFFER_OVERFLOW       c.Int = 13
	LT_ERROR_INVALID_ERRORCODE     c.Int = 14
	LT_ERROR_SHUTDOWN              c.Int = 15
	LT_ERROR_CLOSE_RESIDENT_MODULE c.Int = 16
	LT_ERROR_INVALID_MUTEX_ARGS    c.Int = 17
	LT_ERROR_INVALID_POSITION      c.Int = 18
	LT_ERROR_CONFLICTING_FLAGS     c.Int = 19
	LT_ERROR_MAX                   c.Int = 20
)

/* These functions are only useful from inside custom module loaders. */
//go:linkname Dladderror C.lt_dladderror
func Dladderror(diagnostic *c.Char) c.Int

//go:linkname Dlseterror C.lt_dlseterror
func Dlseterror(errorcode c.Int) c.Int
