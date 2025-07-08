package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2cPortT c.Int

const (
	I2C_NUM_0   I2cPortT = 0
	I2C_NUM_1   I2cPortT = 1
	I2C_NUM_MAX I2cPortT = 2
)

type I2cAddrBitLenT c.Int

const (
	I2C_ADDR_BIT_LEN_7  I2cAddrBitLenT = 0
	I2C_ADDR_BIT_LEN_10 I2cAddrBitLenT = 1
)

/**
 * @brief Data structure for calculating I2C bus timing.
 */

type I2cHalClkConfigT struct {
	ClkmDiv     c.Uint16T
	SclLow      c.Uint16T
	SclHigh     c.Uint16T
	SclWaitHigh c.Uint16T
	SdaHold     c.Uint16T
	SdaSample   c.Uint16T
	Setup       c.Uint16T
	Hold        c.Uint16T
	Tout        c.Uint16T
}
type I2cModeT c.Int

const (
	I2C_MODE_SLAVE  I2cModeT = 0
	I2C_MODE_MASTER I2cModeT = 1
	I2C_MODE_MAX    I2cModeT = 2
)

type I2cRwT c.Int

const (
	I2C_MASTER_WRITE I2cRwT = 0
	I2C_MASTER_READ  I2cRwT = 1
)

type I2cTransModeT c.Int

const (
	I2C_DATA_MODE_MSB_FIRST I2cTransModeT = 0
	I2C_DATA_MODE_LSB_FIRST I2cTransModeT = 1
	I2C_DATA_MODE_MAX       I2cTransModeT = 2
)

type I2cAddrModeT c.Int

const (
	I2C_ADDR_BIT_7   I2cAddrModeT = 0
	I2C_ADDR_BIT_10  I2cAddrModeT = 1
	I2C_ADDR_BIT_MAX I2cAddrModeT = 2
)

type I2cAckTypeT c.Int

const (
	I2C_MASTER_ACK       I2cAckTypeT = 0
	I2C_MASTER_NACK      I2cAckTypeT = 1
	I2C_MASTER_LAST_NACK I2cAckTypeT = 2
	I2C_MASTER_ACK_MAX   I2cAckTypeT = 3
)

type I2cSlaveStretchCauseT c.Int

const (
	I2C_SLAVE_STRETCH_CAUSE_ADDRESS_MATCH I2cSlaveStretchCauseT = 0
	I2C_SLAVE_STRETCH_CAUSE_TX_EMPTY      I2cSlaveStretchCauseT = 1
	I2C_SLAVE_STRETCH_CAUSE_RX_FULL       I2cSlaveStretchCauseT = 2
	I2C_SLAVE_STRETCH_CAUSE_SENDING_ACK   I2cSlaveStretchCauseT = 3
)

type I2cSlaveReadWriteStatusT c.Int

const (
	I2C_SLAVE_WRITE_BY_MASTER I2cSlaveReadWriteStatusT = 0
	I2C_SLAVE_READ_BY_MASTER  I2cSlaveReadWriteStatusT = 1
)

type I2cBusModeT c.Int

const (
	I2C_BUS_MODE_MASTER I2cBusModeT = 0
	I2C_BUS_MODE_SLAVE  I2cBusModeT = 1
)

type I2cClockSourceT SocPeriphI2cClkSrcT
