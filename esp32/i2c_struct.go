package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Timing registers */
/** Type of scl_low_period register
 *  Configures the low level width of the SCL
 *  Clock
 */

type I2cSclLowPeriodRegT struct {
	Val c.Uint32T
}

/** Type of sda_hold register
 *  Configures the hold time after a negative SCL edge.
 */

type I2cSdaHoldRegT struct {
	Val c.Uint32T
}

/** Type of sda_sample register
 *  Configures the sample time after a positive SCL edge.
 */

type I2cSdaSampleRegT struct {
	Val c.Uint32T
}

/** Type of scl_high_period register
 *  Configures the high level width of SCL
 */

type I2cSclHighPeriodRegT struct {
	Val c.Uint32T
}

/** Type of scl_start_hold register
 *  Configures the delay between the SDA and SCL negative edge for a start condition
 */

type I2cSclStartHoldRegT struct {
	Val c.Uint32T
}

/** Type of scl_rstart_setup register
 *  Configures the delay between the positive
 *  edge of SCL and the negative edge of SDA
 */

type I2cSclRstartSetupRegT struct {
	Val c.Uint32T
}

/** Type of scl_stop_hold register
 *  Configures the delay after the SCL clock
 *  edge for a stop condition
 */

type I2cSclStopHoldRegT struct {
	Val c.Uint32T
}

/** Type of scl_stop_setup register
 *  Configures the delay between the SDA and
 *  SCL positive edge for a stop condition
 */

type I2cSclStopSetupRegT struct {
	Val c.Uint32T
}

/** Type of scl_st_time_out register
 *  SCL status time out register
 */

type I2cSclStTimeOutRegT struct {
	Val c.Uint32T
}

/** Type of scl_main_st_time_out register
 *  SCL main status time out register
 */

type I2cSclMainStTimeOutRegT struct {
	Val c.Uint32T
}

/** Group: Configuration registers */
/** Type of ctr register
 *  Transmission setting
 */

type I2cCtrRegT struct {
	Val c.Uint32T
}

/** Type of to register
 *  Setting time out control for receiving data.
 */

type I2cToRegT struct {
	Val c.Uint32T
}

/** Type of slave_addr register
 *  Local slave address setting
 */

type I2cSlaveAddrRegT struct {
	Val c.Uint32T
}

/** Type of fifo_conf register
 *  FIFO configuration register.
 */

type I2cFifoConfRegT struct {
	Val c.Uint32T
}

/** Type of filter_cfg register
 *  SCL and SDA filter configuration register
 */

type I2cFilterCfgRegT struct {
	Val c.Uint32T
}

/** Type of clk_conf register
 *  I2C CLK configuration register
 */

type I2cClkConfRegT struct {
	Val c.Uint32T
}

/** Type of scl_sp_conf register
 *  Power configuration register
 */

type I2cSclSpConfRegT struct {
	Val c.Uint32T
}

/** Type of scl_stretch_conf register
 *  Set SCL stretch of I2C slave
 */

type I2cSclStretchConfRegT struct {
	Val c.Uint32T
}

/** Group: Status registers */
/** Type of sr register
 *  Describe I2C work status.
 */

type I2cSrRegT struct {
	Val c.Uint32T
}

/** Type of fifo_st register
 *  FIFO status register.
 */

type I2cFifoStRegT struct {
	Val c.Uint32T
}

/** Type of data register
 *  Rx FIFO read data.
 */

type I2cDataRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of int_raw register
 *  Raw interrupt status
 */

type I2cIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear bits
 */

type I2cIntClrRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable bits
 */

type I2cIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_status register
 *  Status of captured I2C communication events
 */

type I2cIntStatusRegT struct {
	Val c.Uint32T
}

/** Group: Command registers */
/** Type of command register
 *  I2C command register
 */

type I2cComdRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version register
 */

type I2cDateRegT struct {
	Val c.Uint32T
}

type I2cDevT struct {
	SclLowPeriod     I2cSclLowPeriodRegT
	Ctr              I2cCtrRegT
	Sr               I2cSrRegT
	To               I2cToRegT
	SlaveAddr        I2cSlaveAddrRegT
	FifoSt           I2cFifoStRegT
	FifoConf         I2cFifoConfRegT
	Data             I2cDataRegT
	IntRaw           I2cIntRawRegT
	IntClr           I2cIntClrRegT
	IntEna           I2cIntEnaRegT
	IntStatus        I2cIntStatusRegT
	SdaHold          I2cSdaHoldRegT
	SdaSample        I2cSdaSampleRegT
	SclHighPeriod    I2cSclHighPeriodRegT
	Reserved03c      c.Uint32T
	SclStartHold     I2cSclStartHoldRegT
	SclRstartSetup   I2cSclRstartSetupRegT
	SclStopHold      I2cSclStopHoldRegT
	SclStopSetup     I2cSclStopSetupRegT
	FilterCfg        I2cFilterCfgRegT
	ClkConf          I2cClkConfRegT
	Comd             [8]I2cComdRegT
	SclStTimeOut     I2cSclStTimeOutRegT
	SclMainStTimeOut I2cSclMainStTimeOutRegT
	SclSpConf        I2cSclSpConfRegT
	SclStretchConf   I2cSclStretchConfRegT
	Reserved088      [28]c.Uint32T
	Date             I2cDateRegT
	Reserved0fc      c.Uint32T
	TxfifoMem        [32]c.Uint32T
	RxfifoMem        [32]c.Uint32T
}
