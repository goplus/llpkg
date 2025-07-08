package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Start the whole signing process after the input key is ready.
 *
 * Call this before using any of the functions below. The input key is ready must be ready at this point.
 */
//go:linkname DsHalStart C.ds_hal_start
func DsHalStart()

/**
 * @brief Finish the whole signing process. Call this after the signature is read or in case of an error.
 */
//go:linkname DsHalFinish C.ds_hal_finish
func DsHalFinish()

/**
 * @brief Write the initialization vector.
 */
//go:linkname DsHalConfigureIv C.ds_hal_configure_iv
func DsHalConfigureIv(iv *c.Uint32T)

/**
 * @brief Write the message which should be signed.
 *
 * @param msg Pointer to the message.
 * @param size Length of signature result in bytes. It is the RSA signature length in bytes.
 */
//go:linkname DsHalWriteMessage C.ds_hal_write_message
func DsHalWriteMessage(msg *c.Uint8T, size c.SizeT)

/**
 * @brief Write the encrypted private key parameters.
 */
//go:linkname DsHalWritePrivateKeyParams C.ds_hal_write_private_key_params
func DsHalWritePrivateKeyParams(block *c.Uint8T)

/**
 * @brief Begin signing procedure.
 */
//go:linkname DsHalStartSign C.ds_hal_start_sign
func DsHalStartSign()

/**
 * @brief Check whether the hardware is busy with an operation.
 *
 * @return True if the hardware has finished the signing procedure, otherwise false.
 */
//go:linkname DsHalBusy C.ds_hal_busy
func DsHalBusy() bool

/**
 * @brief Check and read the signature from the hardware.
 *
 * @return
 * - DS_SIGNATURE_OK if no issue is detected with the signature.
 * - DS_SIGNATURE_PADDING_FAIL if the padding of the private key parameters is wrong.
 * - DS_SIGNATURE_MD_FAIL if the message digest check failed. This means that the message digest calculated using
 *      the private key parameters fails, i.e., the integrity of the private key parameters is not protected.
 * - DS_SIGNATURE_PADDING_AND_MD_FAIL if both padding and message digest check fail.
 */
//go:linkname DsHalReadResult C.ds_hal_read_result
func DsHalReadResult(result *c.Uint8T, size c.SizeT) DsSignatureCheckT
