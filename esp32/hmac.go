package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

//go:linkname EtsHmacEnable C.ets_hmac_enable
func EtsHmacEnable()

//go:linkname EtsHmacDisable C.ets_hmac_disable
func EtsHmacDisable()

/* Use the "upstream" HMAC key (ETS_EFUSE_KEY_PURPOSE_HMAC_UP)
   to digest a message.
*/
// llgo:link EtsEfuseBlockT.EtsHmacCalculateMessage C.ets_hmac_calculate_message
func (recv_ EtsEfuseBlockT) EtsHmacCalculateMessage(message c.Pointer, message_len c.SizeT, hmac *c.Uint8T) c.Int {
	return 0
}

/* Calculate a downstream HMAC message to temporarily enable JTAG, or
   to generate a Digital Signature data decryption key.

   - purpose must be ETS_EFUSE_KEY_PURPOSE_HMAC_DOWN_DIGITAL_SIGNATURE
     or ETS_EFUSE_KEY_PURPOSE_HMAC_DOWN_JTAG

   - key_block must be in range ETS_EFUSE_BLOCK_KEY0 toETS_EFUSE_BLOCK_KEY6.
     This efuse block must have the corresponding purpose set in "purpose", or
     ETS_EFUSE_KEY_PURPOSE_HMAC_DOWN_ALL.

   The result of this HMAC calculation is only made available "downstream" to the
   corresponding hardware module, and cannot be accessed by software.
*/
// llgo:link EtsEfuseBlockT.EtsHmacCalculateDownstream C.ets_hmac_calculate_downstream
func (recv_ EtsEfuseBlockT) EtsHmacCalculateDownstream(purpose EtsEfusePurposeT) c.Int {
	return 0
}

/* Invalidate a downstream HMAC value previously calculated by ets_hmac_calculate_downstream().
 *
 * - purpose must match a previous call to ets_hmac_calculate_downstream().
 *
 * After this function is called, the corresponding internal operation (JTAG or DS) will no longer
 * have access to the generated key.
 */
// llgo:link EtsEfusePurposeT.EtsHmacInvalidateDownstream C.ets_hmac_invalidate_downstream
func (recv_ EtsEfusePurposeT) EtsHmacInvalidateDownstream() c.Int {
	return 0
}
