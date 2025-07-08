package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief set eFuse timings
 *
 * @param apb_freq_hz APB frequency in Hz
 */
//go:linkname EfuseHalSetTiming C.efuse_hal_set_timing
func EfuseHalSetTiming(apb_freq_hz c.Uint32T)

/**
 * @brief trigger eFuse read operation
 */
//go:linkname EfuseHalRead C.efuse_hal_read
func EfuseHalRead()

/**
 * @brief clear registers for programming eFuses
 */
//go:linkname EfuseHalClearProgramRegisters C.efuse_hal_clear_program_registers
func EfuseHalClearProgramRegisters()

/**
 * @brief burn eFuses written in programming registers (one block at once)
 *
 * @param block block number
 */
//go:linkname EfuseHalProgram C.efuse_hal_program
func EfuseHalProgram(block c.Uint32T)

/**
 * @brief Calculate Reed-Solomon Encoding values for a block of efuse data.
 *
 * @param data Pointer to data buffer (length 32 bytes)
 * @param rs_values Pointer to write encoded data to (length 12 bytes)
 */
//go:linkname EfuseHalRsCalculate C.efuse_hal_rs_calculate
func EfuseHalRsCalculate(data c.Pointer, rs_values c.Pointer)

/**
 * @brief Checks coding error in a block
 *
 * @param block Index of efuse block
 *
 * @return True  - block has an error.
 *         False - no error.
 */
//go:linkname EfuseHalIsCodingErrorInBlock C.efuse_hal_is_coding_error_in_block
func EfuseHalIsCodingErrorInBlock(block c.Uint) bool
