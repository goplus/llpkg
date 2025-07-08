package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type HmacHalOutputT c.Int

const (
	HMAC_OUTPUT_USER        HmacHalOutputT = 0
	HMAC_OUTPUT_DS          HmacHalOutputT = 1
	HMAC_OUTPUT_JTAG_ENABLE HmacHalOutputT = 2
	HMAC_OUTPUT_ALL         HmacHalOutputT = 3
)
