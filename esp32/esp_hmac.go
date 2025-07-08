package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type HmacKeyIdT c.Int

const (
	HMAC_KEY0    HmacKeyIdT = 0
	HMAC_KEY1    HmacKeyIdT = 1
	HMAC_KEY2    HmacKeyIdT = 2
	HMAC_KEY3    HmacKeyIdT = 3
	HMAC_KEY4    HmacKeyIdT = 4
	HMAC_KEY5    HmacKeyIdT = 5
	HMAC_KEY_MAX HmacKeyIdT = 6
)

/**
 * @brief
 * Calculate the HMAC of a given message.
 *
 * Calculate the HMAC \c hmac of a given message \c message with length \c message_len.
 * SHA256 is used for the calculation.
 *
 * @note Uses the HMAC peripheral in "upstream" mode.
 *
 * @param key_id Determines which of the 6 key blocks in the efuses should be used for the HMAC calculation.
 *        The corresponding purpose field of the key block in the efuse must be set to the HMAC upstream purpose value.
 * @param message the message for which to calculate the HMAC
 * @param message_len message length
 *                         return ESP_ERR_INVALID_STATE if unsuccessful
 * @param [out] hmac the hmac result; the buffer behind the provided pointer must be a writeable buffer of 32 bytes
 *
 * @return
 *      * ESP_OK, if the calculation was successful,
 *      * ESP_ERR_INVALID_ARG if message or hmac is a nullptr or if key_id out of range
 *      * ESP_FAIL, if the hmac calculation failed
 */
// llgo:link HmacKeyIdT.EspHmacCalculate C.esp_hmac_calculate
func (recv_ HmacKeyIdT) EspHmacCalculate(message c.Pointer, message_len c.SizeT, hmac *c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Use HMAC peripheral in Downstream mode to re-enable the JTAG, if it is not permanently disabled by HW.
 *        In downstream mode, HMAC calculations performed by peripheral are used internally and not provided back to user.
 *
 * @param key_id Determines which of the 6 key blocks in the efuses should be used for the HMAC calculation.
 *        The corresponding purpose field of the key block in the efuse must be set to HMAC downstream purpose.
 *
 * @param token Pre calculated HMAC value of the 32-byte 0x00 using SHA-256 and the known private HMAC key. The key is already
 *        programmed to a eFuse key block. The key block number is provided as the first parameter to this function.
 *
 * @return
 *      * ESP_OK, if the key_purpose of the key_id matches to HMAC downstread mode,
 *                The API returns success even if calculated HMAC does not match with the provided token.
 *                However, The JTAG will be re-enabled only if the calculated HMAC value matches with provided token,
 *                otherwise JTAG will remain disabled.
 *      * ESP_FAIL, if the key_purpose of the key_id is not set to HMAC downstream purpose
 *                  or JTAG is permanently disabled by EFUSE_HARD_DIS_JTAG eFuse parameter.
 *      * ESP_ERR_INVALID_ARG, invalid input arguments
 *
 * @note  Return value of the API does not indicate the JTAG status.
 */
// llgo:link HmacKeyIdT.EspHmacJtagEnable C.esp_hmac_jtag_enable
func (recv_ HmacKeyIdT) EspHmacJtagEnable(token *c.Uint8T) EspErrT {
	return 0
}

/**
 *  @brief Disable the JTAG which might be enabled using the HMAC downstream mode. This function just clears the result generated
 *         by calling esp_hmac_jtag_enable() API.
 *
 *  @return
 *       * ESP_OK return ESP_OK after writing the HMAC_SET_INVALIDATE_JTAG_REG with value 1.
 */
//go:linkname EspHmacJtagDisable C.esp_hmac_jtag_disable
func EspHmacJtagDisable() EspErrT
