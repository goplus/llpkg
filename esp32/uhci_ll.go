package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UhciRxeofCfgT c.Int

const (
	UHCI_RX_BREAK_CHR_EOF UhciRxeofCfgT = 1
	UHCI_RX_IDLE_EOF      UhciRxeofCfgT = 2
	UHCI_RX_LEN_EOF       UhciRxeofCfgT = 4
	UHCI_RX_EOF_MAX       UhciRxeofCfgT = 7
)
