package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TemperatureSensorAttributeT struct {
	Offset   c.Int
	RegVal   c.Int
	RangeMin c.Int
	RangeMax c.Int
	ErrorMax c.Int
}
