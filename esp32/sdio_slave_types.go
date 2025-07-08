package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SdioSlaveHostintT c.Int

const (
	SDIO_SLAVE_HOSTINT_BIT0            SdioSlaveHostintT = 1
	SDIO_SLAVE_HOSTINT_BIT1            SdioSlaveHostintT = 2
	SDIO_SLAVE_HOSTINT_BIT2            SdioSlaveHostintT = 4
	SDIO_SLAVE_HOSTINT_BIT3            SdioSlaveHostintT = 8
	SDIO_SLAVE_HOSTINT_BIT4            SdioSlaveHostintT = 16
	SDIO_SLAVE_HOSTINT_BIT5            SdioSlaveHostintT = 32
	SDIO_SLAVE_HOSTINT_BIT6            SdioSlaveHostintT = 64
	SDIO_SLAVE_HOSTINT_BIT7            SdioSlaveHostintT = 128
	SDIO_SLAVE_HOSTINT_SEND_NEW_PACKET SdioSlaveHostintT = 8388608
)

type SdioSlaveTimingT c.Int

const (
	SDIO_SLAVE_TIMING_PSEND_PSAMPLE SdioSlaveTimingT = 0
	SDIO_SLAVE_TIMING_NSEND_PSAMPLE SdioSlaveTimingT = 1
	SDIO_SLAVE_TIMING_PSEND_NSAMPLE SdioSlaveTimingT = 2
	SDIO_SLAVE_TIMING_NSEND_NSAMPLE SdioSlaveTimingT = 3
)

type SdioSlaveSendingModeT c.Int

const (
	SDIO_SLAVE_SEND_STREAM SdioSlaveSendingModeT = 0
	SDIO_SLAVE_SEND_PACKET SdioSlaveSendingModeT = 1
)
