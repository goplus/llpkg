package sqlite3

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

// llgo:type C
type LoadextEntry func(*Sqlite3, **int8, *ApiRoutines) c.Int
