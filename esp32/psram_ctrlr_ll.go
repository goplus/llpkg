package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const PSRAM_CTRLR_LL_MSPI_ID_0 = 0
const PSRAM_CTRLR_LL_MSPI_ID_1 = 1

type PsramLlCsIdT c.Int

const (
	PSRAM_LL_CS_ID_0 PsramLlCsIdT = 0
	PSRAM_LL_CS_ID_1 PsramLlCsIdT = 1
)
