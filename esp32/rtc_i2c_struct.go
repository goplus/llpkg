package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configure Registers */
/** Type of i2c_scl_low register
 *  configure low scl period
 */

type RtcI2cSclLowRegT struct {
	Val c.Uint32T
}

/** Type of i2c_ctrl register
 *  configure i2c ctrl
 */

type RtcI2cCtrlRegT struct {
	Val c.Uint32T
}

/** Type of i2c_to register
 *  configure time out
 */

type RtcI2cToRegT struct {
	Val c.Uint32T
}

/** Type of i2c_slave_addr register
 *  configure slave id
 */

type RtcI2cSlaveAddrRegT struct {
	Val c.Uint32T
}

/** Type of i2c_scl_high register
 *  configure high scl period
 */

type RtcI2cSclHighRegT struct {
	Val c.Uint32T
}

/** Type of i2c_sda_duty register
 *  configure sda duty
 */

type RtcI2cSdaDutyRegT struct {
	Val c.Uint32T
}

/** Type of i2c_scl_start_period register
 *  configure scl start period
 */

type RtcI2cSclStartPeriodRegT struct {
	Val c.Uint32T
}

/** Type of i2c_scl_stop_period register
 *  configure scl stop period
 */

type RtcI2cSclStopPeriodRegT struct {
	Val c.Uint32T
}

/** Type of i2c_data register
 *  get i2c data status
 */

type RtcI2cDataRegT struct {
	Val c.Uint32T
}

/** Type of i2c_cmd register
 *  i2c command register
 */

type RtcI2cCmdRegT struct {
	Val c.Uint32T
}

/** Group: status register */
/** Type of i2c_status register
 *  get i2c status
 */

type RtcI2cStatusRegT struct {
	Val c.Uint32T
}

/** Group: interrupt Register */
/** Type of i2c_int_clr register
 *  interrupt clear register
 */

type RtcI2cIntClrRegT struct {
	Val c.Uint32T
}

/** Type of i2c_int_raw register
 *  interrupt raw register
 */

type RtcI2cIntRawRegT struct {
	Val c.Uint32T
}

/** Type of i2c_int_st register
 *  interrupt state register
 */

type RtcI2cIntStRegT struct {
	Val c.Uint32T
}

/** Type of i2c_int_ena register
 *  interrupt enable register
 */

type RtcI2cIntEnaRegT struct {
	Val c.Uint32T
}

/** Group: version Registers */
/** Type of i2c_date register
 *  version register
 */

type RtcI2cDateRegT struct {
	Val c.Uint32T
}

type RtcI2cDevT struct {
	I2cSclLow         RtcI2cSclLowRegT
	I2cCtrl           RtcI2cCtrlRegT
	I2cStatus         RtcI2cStatusRegT
	I2cTo             RtcI2cToRegT
	I2cSlaveAddr      RtcI2cSlaveAddrRegT
	I2cSclHigh        RtcI2cSclHighRegT
	I2cSdaDuty        RtcI2cSdaDutyRegT
	I2cSclStartPeriod RtcI2cSclStartPeriodRegT
	I2cSclStopPeriod  RtcI2cSclStopPeriodRegT
	I2cIntClr         RtcI2cIntClrRegT
	I2cIntRaw         RtcI2cIntRawRegT
	I2cIntSt          RtcI2cIntStRegT
	I2cIntEna         RtcI2cIntEnaRegT
	I2cData           RtcI2cDataRegT
	I2cCmd            [16]RtcI2cCmdRegT
	Reserved078       [33]c.Uint32T
	I2cDate           RtcI2cDateRegT
}
