package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SdBusWidthT c.Int

const (
	SD_BUS_WIDTH_1_BIT SdBusWidthT = 0
	SD_BUS_WIDTH_4_BIT SdBusWidthT = 1
	SD_BUS_WIDTH_8_BIT SdBusWidthT = 2
)
