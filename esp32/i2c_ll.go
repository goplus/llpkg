package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const I2C_LL_CMD_RESTART = 6
const I2C_LL_CMD_WRITE = 1
const I2C_LL_CMD_READ = 3
const I2C_LL_CMD_STOP = 2
const I2C_LL_CMD_END = 4

/**
 * @brief I2C hardware cmd register fields.
 */

type I2cLlHwCmdT struct {
	Val c.Uint32T
}
type I2cLlMasterIntrT c.Int

const (
	I2C_INTR_MST_TXFIFO_WM   I2cLlMasterIntrT = 2
	I2C_INTR_MST_RXFIFO_WM   I2cLlMasterIntrT = 1
	I2C_LL_INTR_NACK         I2cLlMasterIntrT = 1024
	I2C_LL_INTR_TIMEOUT      I2cLlMasterIntrT = 256
	I2C_LL_INTR_MST_COMPLETE I2cLlMasterIntrT = 128
	I2C_LL_INTR_ARBITRATION  I2cLlMasterIntrT = 32
	I2C_LL_INTR_END_DETECT   I2cLlMasterIntrT = 8
	I2C_LL_INTR_ST_TO        I2cLlMasterIntrT = 8192
)

type I2cLlSlaveIntrT c.Int

const (
	I2C_INTR_SLV_TXFIFO_WM I2cLlSlaveIntrT = 2
	I2C_INTR_SLV_RXFIFO_WM I2cLlSlaveIntrT = 1
	I2C_INTR_SLV_COMPLETE  I2cLlSlaveIntrT = 128
	I2C_INTR_START         I2cLlSlaveIntrT = 32768
	I2C_INTR_STRETCH       I2cLlSlaveIntrT = 65536
)

type I2cIntrEventT c.Int

const (
	I2C_INTR_EVENT_ERR          I2cIntrEventT = 0
	I2C_INTR_EVENT_ARBIT_LOST   I2cIntrEventT = 1
	I2C_INTR_EVENT_NACK         I2cIntrEventT = 2
	I2C_INTR_EVENT_TOUT         I2cIntrEventT = 3
	I2C_INTR_EVENT_END_DET      I2cIntrEventT = 4
	I2C_INTR_EVENT_TRANS_DONE   I2cIntrEventT = 5
	I2C_INTR_EVENT_RXFIFO_FULL  I2cIntrEventT = 6
	I2C_INTR_EVENT_TXFIFO_EMPTY I2cIntrEventT = 7
)
