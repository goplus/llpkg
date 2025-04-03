package libtool

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const LT_ERROR_H = 1
const (
	LTERRORUNKNOWN             c.Int = 0
	LTERRORDLOPENNOTSUPPORTED  c.Int = 1
	LTERRORINVALIDLOADER       c.Int = 2
	LTERRORINITLOADER          c.Int = 3
	LTERRORREMOVELOADER        c.Int = 4
	LTERRORFILENOTFOUND        c.Int = 5
	LTERRORDEPLIBNOTFOUND      c.Int = 6
	LTERRORNOSYMBOLS           c.Int = 7
	LTERRORCANNOTOPEN          c.Int = 8
	LTERRORCANNOTCLOSE         c.Int = 9
	LTERRORSYMBOLNOTFOUND      c.Int = 10
	LTERRORNOMEMORY            c.Int = 11
	LTERRORINVALIDHANDLE       c.Int = 12
	LTERRORBUFFEROVERFLOW      c.Int = 13
	LTERRORINVALIDERRORCODE    c.Int = 14
	LTERRORSHUTDOWN            c.Int = 15
	LTERRORCLOSERESIDENTMODULE c.Int = 16
	LTERRORINVALIDMUTEXARGS    c.Int = 17
	LTERRORINVALIDPOSITION     c.Int = 18
	LTERRORCONFLICTINGFLAGS    c.Int = 19
	LTERRORMAX                 c.Int = 20
)

/* These functions are only useful from inside custom module loaders. */
//go:linkname Dladderror C.lt_dladderror
func Dladderror(diagnostic *int8) c.Int

//go:linkname Dlseterror C.lt_dlseterror
func Dlseterror(errorcode c.Int) c.Int
