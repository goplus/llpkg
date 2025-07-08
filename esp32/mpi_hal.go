package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Calculate the number of words needed to represent the input word in hardware.
 *
 * @param words The number of words to be represented.
 * @return size_t Number of words required.
 */
//go:linkname MpiHalCalcHardwareWords C.mpi_hal_calc_hardware_words
func MpiHalCalcHardwareWords(words c.SizeT) c.SizeT

/**
 * @brief Clear the MPI power control bit and intitialise the MPI hardware.
 *
 */
//go:linkname MpiHalEnableHardwareHwOp C.mpi_hal_enable_hardware_hw_op
func MpiHalEnableHardwareHwOp()

/**
 * @brief Set the MPI power control bit to disable the MPI hardware.
 *
 */
//go:linkname MpiHalDisableHardwareHwOp C.mpi_hal_disable_hardware_hw_op
func MpiHalDisableHardwareHwOp()

/**
 * @brief Enable/disables MPI operation complete interrupt.
 *
 * @param enable true: enable, false: disable.
 */
//go:linkname MpiHalInterruptEnable C.mpi_hal_interrupt_enable
func MpiHalInterruptEnable(enable bool)

/**
 * @brief Clears the MPI operation complete interrupt status.
 *
 */
//go:linkname MpiHalClearInterrupt C.mpi_hal_clear_interrupt
func MpiHalClearInterrupt()

/**
 * @brief Configure RSA length.
 *
 * @param num_words Number of words representing the RSA length.
 */
//go:linkname MpiHalSetMode C.mpi_hal_set_mode
func MpiHalSetMode(num_words c.SizeT)

/**
 * @brief Copy the large number (array of words) representation of the parameter 'param' to hardware memory block.
 *
 * @param param Type of parameter (enum).
 * @param offset Offset to copy in the memory from the base address of the parameter.
 * @param p Pointer to large number (array of words) representation of the parameter.
 * @param n Number of words needed to represent the large number as an array of words.
 * @param num_words Maximum hardware words needed.
 */
// llgo:link MpiParamT.MpiHalWriteToMemBlock C.mpi_hal_write_to_mem_block
func (recv_ MpiParamT) MpiHalWriteToMemBlock(offset c.SizeT, p *c.Uint32T, n c.SizeT, num_words c.SizeT) {
}

/**
 * @brief Write a word-sized value to hardware memory block of a parameter.
 *
 * @param param Type of parameter (enum).
 * @param offset Offset to copy in the memory from the base address of the parameter.
 * @param value Value to be written in the memory.
 */
// llgo:link MpiParamT.MpiHalWriteAtOffset C.mpi_hal_write_at_offset
func (recv_ MpiParamT) MpiHalWriteAtOffset(offset c.Int, value c.Uint32T) {
}

/**
 * @brief Write the modular multiplicative inverse of M.
 *
 * @param Mprime Modular multiplicative inverse of M.
 */
//go:linkname MpiHalWriteMPrime C.mpi_hal_write_m_prime
func MpiHalWriteMPrime(Mprime c.Uint32T)

/**
 * @brief Write first word of the parametr Rinv.
 *
 * @param rinv Value of first word of rinv.
 */
//go:linkname MpiHalWriteRinv C.mpi_hal_write_rinv
func MpiHalWriteRinv(rinv c.Uint32T)

/**
 * @brief Enable/Disable constant time acceleration option.
 *
 * @param enable true: enable, false: disable.
 */
//go:linkname MpiHalEnableConstantTime C.mpi_hal_enable_constant_time
func MpiHalEnableConstantTime(enable bool)

/**
 * @brief Enable/Disable search time acceleration option.
 *
 * @param enable
 */
//go:linkname MpiHalEnableSearch C.mpi_hal_enable_search
func MpiHalEnableSearch(enable bool)

/**
 * @brief Configures the starting address to start search.
 *
 * @param position Address to start search.
 */
//go:linkname MpiHalSetSearchPosition C.mpi_hal_set_search_position
func MpiHalSetSearchPosition(position c.SizeT)

/**
 * @brief Begin an MPI operation.
 *
 * @param op Operation type (enum).
 */
// llgo:link MpiOpT.MpiHalStartOp C.mpi_hal_start_op
func (recv_ MpiOpT) MpiHalStartOp() {
}

/**
 * @brief  Wait for an MPI operation to complete.
 *
 */
//go:linkname MpiHalWaitOpComplete C.mpi_hal_wait_op_complete
func MpiHalWaitOpComplete()

/**
 * @brief Wait for an MPI operation to complete and Read result from last MPI operation into parameter Z.
 *
 * @param p Pointer to large number (array of words) representation of the parameter.
 * @param n Number of words needed to represent the large number as an array of words.
 * @param z_words Calculated number of words of parameter Z.
 */
//go:linkname MpiHalReadResultHwOp C.mpi_hal_read_result_hw_op
func MpiHalReadResultHwOp(p *c.Uint32T, n c.SizeT, z_words c.SizeT)
