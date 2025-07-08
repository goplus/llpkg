package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type IntDescFlagT c.Int

const (
	INTDESC_NORMAL  IntDescFlagT = 0
	INTDESC_RESVD   IntDescFlagT = 1
	INTDESC_SPECIAL IntDescFlagT = 2
)

type IntTypeT c.Int

const (
	INTTP_LEVEL IntTypeT = 0
	INTTP_EDGE  IntTypeT = 1
	INTTP_NA    IntTypeT = 2
)

type IntDescT struct {
	Level    c.Int
	Type     IntTypeT
	Cpuflags [2]IntDescFlagT
}

// llgo:type C
type InterruptHandlerT func(c.Pointer)
