package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* Access internal registers, don't use in application */
//go:linkname Regi2cCtrlReadReg C.regi2c_ctrl_read_reg
func Regi2cCtrlReadReg(block c.Uint8T, host_id c.Uint8T, reg_add c.Uint8T) c.Uint8T

//go:linkname Regi2cCtrlReadRegMask C.regi2c_ctrl_read_reg_mask
func Regi2cCtrlReadRegMask(block c.Uint8T, host_id c.Uint8T, reg_add c.Uint8T, msb c.Uint8T, lsb c.Uint8T) c.Uint8T

//go:linkname Regi2cCtrlWriteReg C.regi2c_ctrl_write_reg
func Regi2cCtrlWriteReg(block c.Uint8T, host_id c.Uint8T, reg_add c.Uint8T, data c.Uint8T)

//go:linkname Regi2cCtrlWriteRegMask C.regi2c_ctrl_write_reg_mask
func Regi2cCtrlWriteRegMask(block c.Uint8T, host_id c.Uint8T, reg_add c.Uint8T, msb c.Uint8T, lsb c.Uint8T, data c.Uint8T)

/* enter the critical section that protects internal registers. Don't use it in SDK. Use the functions above. */
//go:linkname Regi2cEnterCritical C.regi2c_enter_critical
func Regi2cEnterCritical()

//go:linkname Regi2cExitCritical C.regi2c_exit_critical
func Regi2cExitCritical()

/**
 * Restore regi2c analog calibration related configuration registers.
 * This is a workaround, and is fixed on later chips
 */
//go:linkname Regi2cAnalogCaliRegRead C.regi2c_analog_cali_reg_read
func Regi2cAnalogCaliRegRead()

//go:linkname Regi2cAnalogCaliRegWrite C.regi2c_analog_cali_reg_write
func Regi2cAnalogCaliRegWrite()

/* Enable/Disable regi2c_saradc with calling these two functions.
   With reference count protection inside.
   Internal use only.
*/
//go:linkname Regi2cSaradcEnable C.regi2c_saradc_enable
func Regi2cSaradcEnable()

//go:linkname Regi2cSaradcDisable C.regi2c_saradc_disable
func Regi2cSaradcDisable()
