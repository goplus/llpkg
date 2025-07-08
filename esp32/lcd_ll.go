package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const LCD_LL_CLK_FRAC_DIV_N_MAX = 256
const LCD_LL_CLK_FRAC_DIV_AB_MAX = 64
const LCD_LL_PCLK_DIV_MAX = 64
const LCD_LL_FIFO_DEPTH = 16

type LcdLlSwizzleModeT c.Int

const LCD_LL_SWIZZLE_AB2BA LcdLlSwizzleModeT = 0
