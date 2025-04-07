package sqlite3

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

// llgo:type C
type LoadextEntry func(*Sqlite3, **int8, *ApiRoutines) c.Int
