package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Make the peripheral ready for use.
 *
 * This triggers any further steps necessary after enabling the device
 */
//go:linkname HmacHalStart C.hmac_hal_start
func HmacHalStart()

/**
 * @brief Configure which hardware key slot should be used and configure the target of the HMAC output.
 *
 * @note Writing out-of-range values is undefined behavior. The user has to ensure that the parameters are in range.
 *
 * @param config The target of the HMAC. Possible targets are described in \c hmac_hal_output_t.
 * See the TRM of your target chip for more details.
 * @param key_id The ID of the hardware key slot to be used.
 *
 * @return 0 if the configuration was successful, non-zero if not.
 * An unsuccessful configuration means that the purpose value in the eFuse of the corresponding key slot
 * doesn't match to supplied value of \c config.
 */
// llgo:link HmacHalOutputT.HmacHalConfigure C.hmac_hal_configure
func (recv_ HmacHalOutputT) HmacHalConfigure(key_id c.Uint32T) c.Uint32T {
	return 0
}

/**
 * @brief Write a padded single-block message of 512 bits to the HMAC peripheral.
 *
 * The message must not be longer than one block (512 bits) and the padding has to be applied by software before
 * writing. The padding has to be able to fit into the block after the message.
 * For more information on HMAC padding, see the TRM of your target chip.
 */
//go:linkname HmacHalWriteOneBlock512 C.hmac_hal_write_one_block_512
func HmacHalWriteOneBlock512(block c.Pointer)

/**
 * @brief Write a message block of 512 bits to the HMAC peripheral.
 *
 * This function must be used incombination with \c hmac_hal_next_block_normal() or \c hmac_hal_next_block_padding().
 * The first message block is written without any prerequisite.
 * All message blocks which are not the last one, need a call to \c hmac_hal_next_block_normal() before, indicating
 * to the hardware that a "normal", i.e. non-padded block will follow. This is even the case for a block which begins
 * padding already but where the padding doesn't fit in (remaining message size > (block size - padding size)).
 * Before writing the last block which contains the padding, a call to \c hmac_hal_next_block_padding() is necessary
 * to indicate to the hardware that a block with padding will be written.
 *
 * For more information on HMAC padding, see the TRM of your target chip for more details.
 */
//go:linkname HmacHalWriteBlock512 C.hmac_hal_write_block_512
func HmacHalWriteBlock512(block c.Pointer)

/**
 * @brief Indicate to the hardware that a normal block will be written.
 */
//go:linkname HmacHalNextBlockNormal C.hmac_hal_next_block_normal
func HmacHalNextBlockNormal()

/**
 * @brief Indicate to the hardware that a block with padding will be written.
 */
//go:linkname HmacHalNextBlockPadding C.hmac_hal_next_block_padding
func HmacHalNextBlockPadding()

/**
 * @brief Read the 256 bit HMAC result from the hardware.
 */
//go:linkname HmacHalReadResult256 C.hmac_hal_read_result_256
func HmacHalReadResult256(result c.Pointer)

/**
 * @brief Clear (invalidate) the HMAC result provided to other hardware.
 */
//go:linkname HmacHalClean C.hmac_hal_clean
func HmacHalClean()
