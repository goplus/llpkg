package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type BrownoutHalConfigT struct {
	Threshold      c.Uint8T
	Enabled        bool
	ResetEnabled   bool
	FlashPowerDown bool
	RfPowerDown    bool
}

/**
 * @brief Config brown out hal.
 *
 * @param cfg Pointer of brown out configuration structure.
 */
// llgo:link (*BrownoutHalConfigT).BrownoutHalConfig C.brownout_hal_config
func (recv_ *BrownoutHalConfigT) BrownoutHalConfig() {
}
