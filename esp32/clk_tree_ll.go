package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CLK_LL_RC_FAST_WAIT_DEFAULT = 20
const CLK_LL_RC_FAST_ENABLE_WAIT_DEFAULT = 5

type ClkLlXtal32kEnableModeT c.Int

const (
	CLK_LL_XTAL32K_ENABLE_MODE_CRYSTAL   ClkLlXtal32kEnableModeT = 0
	CLK_LL_XTAL32K_ENABLE_MODE_EXTERNAL  ClkLlXtal32kEnableModeT = 1
	CLK_LL_XTAL32K_ENABLE_MODE_BOOTSTRAP ClkLlXtal32kEnableModeT = 2
)

/**
 * @brief XTAL32K_CLK configuration structure
 */

type ClkLlXtal32kConfigT struct {
	Dac  c.Uint32T
	Dres c.Uint32T
	Dgm  c.Uint32T
	Dbuf c.Uint32T
}
