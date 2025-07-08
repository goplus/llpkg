package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LdoUnitT c.Int

const (
	LDO_UNIT_1 LdoUnitT = 1
	LDO_UNIT_2 LdoUnitT = 2
	LDO_UNIT_3 LdoUnitT = 3
	LDO_UNIT_4 LdoUnitT = 4
)
