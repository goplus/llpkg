package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_DS_IV_BIT_LEN = 128
const ESP_DS_SIGNATURE_MD_BIT_LEN = 256
const ESP_DS_SIGNATURE_M_PRIME_BIT_LEN = 32
const ESP_DS_SIGNATURE_L_BIT_LEN = 32
const ESP_DS_SIGNATURE_PADDING_BIT_LEN = 64

type EspDsContext struct {
	Unused [8]uint8
}
type EspDsContextT EspDsContext
type EspDigitalSignatureLengthT c.Int

const (
	ESP_DS_RSA_1024 EspDigitalSignatureLengthT = 31
	ESP_DS_RSA_2048 EspDigitalSignatureLengthT = 63
	ESP_DS_RSA_3072 EspDigitalSignatureLengthT = 95
	ESP_DS_RSA_4096 EspDigitalSignatureLengthT = 127
)

/**
 * Encrypted private key data. Recommended to store in flash in this format.
 *
 * @note This struct has to match to one from the ROM code! This documentation is mostly taken from there.
 */

type EspDigitalSignatureData struct {
	RsaLength EspDigitalSignatureLengthT
	Iv        [4]c.Uint32T
	C         [1584]c.Uint8T
}
type EspDsDataT EspDigitalSignatureData

/**
 * Plaintext parameters used by Digital Signature.
 *
 * This is only used for encrypting the RSA parameters by calling esp_ds_encrypt_params().
 * Afterwards, the result can be stored in flash or in other persistent memory.
 * The encryption is a prerequisite step before any signature operation can be done.
 *
 * @note
 * Y, M, Rb, & M_Prime must all be in little endian format.
 */

type EspDsPDataT struct {
	Y      [128]c.Uint32T
	M      [128]c.Uint32T
	Rb     [128]c.Uint32T
	MPrime c.Uint32T
	Length c.Uint32T
}

/**
* @brief Sign the message with a hardware key from specific key slot.
* The function calculates a plain RSA signature with help of the DS peripheral.
* The RSA encryption operation is as follows:
* Z = XY mod M where,
* Z is the signature, X is the input message,
* Y and M are the RSA private key parameters.
*
* This function is a wrapper around \c esp_ds_finish_sign() and \c esp_ds_start_sign(), so do not use them
* in parallel.
* It blocks until the signing is finished and then returns the signature.
*
* @note
* Please see note section of \c esp_ds_start_sign() for more details about the input parameters.
*
* @param message the message to be signed; its length should be (data->rsa_length + 1)*4 bytes, and those
         bytes must be in little endian format. It is your responsibility to apply your hash function
         and padding before calling this function, if required. (e.g. message = padding(hash(inputMsg)))
* @param data the encrypted signing key data (AES encrypted RSA key + IV)
* @param key_id the HMAC key ID determining the HMAC key of the HMAC which will be used to decrypt the
*        signing key data
* @param signature the destination of the signature, should be (data->rsa_length + 1)*4 bytes long
*
* @return
*      - ESP_OK if successful, the signature was written to the parameter \c signature.
*      - ESP_ERR_INVALID_ARG if one of the parameters is NULL or data->rsa_length is too long or 0
*      - ESP_ERR_HW_CRYPTO_DS_HMAC_FAIL if there was an HMAC failure during retrieval of the decryption key
*      - ESP_ERR_NO_MEM if there hasn't been enough memory to allocate the context object
*      - ESP_ERR_HW_CRYPTO_DS_INVALID_KEY if there's a problem with passing the HMAC key to the DS component
*      - ESP_ERR_HW_CRYPTO_DS_INVALID_DIGEST if the message digest didn't match; the signature is invalid.
*      - ESP_ERR_HW_CRYPTO_DS_INVALID_PADDING if the message padding is incorrect, the signature can be read though
*        since the message digest matches.
*/
//go:linkname EspDsSign C.esp_ds_sign
func EspDsSign(message c.Pointer, data *EspDsDataT, key_id HmacKeyIdT, signature c.Pointer) EspErrT

/**
* @brief Start the signing process.
*
* This function yields a context object which needs to be passed to \c esp_ds_finish_sign() to finish the signing
* process.
* The function calculates a plain RSA signature with help of the DS peripheral.
* The RSA encryption operation is as follows:
* Z = XY mod M where,
* Z is the signature, X is the input message,
* Y and M are the RSA private key parameters.
*
* @note
* This function locks the HMAC, SHA, AES and RSA components, so the user has to ensure to call
* \c esp_ds_finish_sign() in a timely manner.
* The numbers Y, M, Rb which are a part of esp_ds_data_t should be provided in little endian format
* and should be of length equal to the RSA private key bit length
* The message length in bits should also be equal to the RSA private key bit length.
* No padding is applied to the message automatically, Please ensure the message is appropriate padded before
* calling the API.
*
* @param message the message to be signed; its length should be (data->rsa_length + 1)*4 bytes, and those
         bytes must be in little endian format. It is your responsibility to apply your hash function
         and padding before calling this function, if required. (e.g. message = padding(hash(inputMsg)))
* @param data the encrypted signing key data (AES encrypted RSA key + IV)
* @param key_id the HMAC key ID determining the HMAC key of the HMAC which will be used to decrypt the
*        signing key data
* @param esp_ds_ctx the context object which is needed for finishing the signing process later
*
* @return
*      - ESP_OK if successful, the ds operation was started now and has to be finished with \c esp_ds_finish_sign()
*      - ESP_ERR_INVALID_ARG if one of the parameters is NULL or data->rsa_length is too long or 0
*      - ESP_ERR_HW_CRYPTO_DS_HMAC_FAIL if there was an HMAC failure during retrieval of the decryption key
*      - ESP_ERR_NO_MEM if there hasn't been enough memory to allocate the context object
*      - ESP_ERR_HW_CRYPTO_DS_INVALID_KEY if there's a problem with passing the HMAC key to the DS component
*/
//go:linkname EspDsStartSign C.esp_ds_start_sign
func EspDsStartSign(message c.Pointer, data *EspDsDataT, key_id HmacKeyIdT, esp_ds_ctx **EspDsContextT) EspErrT

/**
 * Return true if the DS peripheral is busy, otherwise false.
 *
 * @note Only valid if \c esp_ds_start_sign() was called before.
 */
//go:linkname EspDsIsBusy C.esp_ds_is_busy
func EspDsIsBusy() bool

/**
* @brief Finish the signing process.
*
* @param signature the destination of the signature, should be (data->rsa_length + 1)*4 bytes long,
         the resultant signature bytes shall be written in little endian format.
* @param esp_ds_ctx the context object retrieved by \c esp_ds_start_sign()
*
* @return
*      - ESP_OK if successful, the ds operation has been finished and the result is written to signature.
*      - ESP_ERR_INVALID_ARG if one of the parameters is NULL
*      - ESP_ERR_HW_CRYPTO_DS_INVALID_DIGEST if the message digest didn't match; the signature is invalid.
*        This means that the encrypted RSA key parameters are invalid, indicating that they may have been tampered
*        with or indicating a flash error, etc.
*      - ESP_ERR_HW_CRYPTO_DS_INVALID_PADDING if the message padding is incorrect, the signature can be read though
*        since the message digest matches (see TRM for more details).
*/
//go:linkname EspDsFinishSign C.esp_ds_finish_sign
func EspDsFinishSign(signature c.Pointer, esp_ds_ctx *EspDsContextT) EspErrT

/**
 * @brief Encrypt the private key parameters.
 *
 * The encryption is a prerequisite step before any signature operation can be done.
 * It is not strictly necessary to use this encryption function, the encryption could also happen on an external
 * device.
 *
 * @param data Output buffer to store encrypted data, suitable for later use generating signatures.
 * @param iv Pointer to 16 byte IV buffer, will be copied into 'data'. Should be randomly generated bytes each time.
 * @param p_data Pointer to input plaintext key data. The expectation is this data will be deleted after this process
 *        is done and 'data' is stored.
 * @param key Pointer to 32 bytes of key data. Type determined by key_type parameter. The expectation is the
 *        corresponding HMAC key will be stored to efuse and then permanently erased.
 *
 * @note
 * The numbers Y, M, Rb which are a part of esp_ds_data_t should be provided in little endian format
 * and should be of length equal to the RSA private key bit length
 * The message length in bits should also be equal to the RSA private key bit length.
 * No padding is applied to the message automatically, Please ensure the message is appropriate padded before
 * calling the API.
 *
 * @return
 *      - ESP_OK if successful, the ds operation has been finished and the result is written to signature.
 *      - ESP_ERR_INVALID_ARG if one of the parameters is NULL or p_data->rsa_length is too long
 */
// llgo:link (*EspDsDataT).EspDsEncryptParams C.esp_ds_encrypt_params
func (recv_ *EspDsDataT) EspDsEncryptParams(iv c.Pointer, p_data *EspDsPDataT, key c.Pointer) EspErrT {
	return 0
}
