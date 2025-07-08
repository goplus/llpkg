package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CRC_SIGN_BLOCK_LEN = 1196
const SIG_BLOCK_PADDING = 4096
const ETS_SECURE_BOOT_V2_SIGNATURE_MAGIC = 0xE7
const SECURE_BOOT_NUM_BLOCKS = 3
const MAX_KEY_DIGESTS = 3

type EtsSecureBootSigBlock struct {
	MagicByte   c.Uint8T
	Version     c.Uint8T
	X_reserved1 c.Uint8T
	X_reserved2 c.Uint8T
	ImageDigest [32]c.Uint8T
	Key         EtsRsaPubkeyT
	Signature   [384]c.Uint8T
	BlockCrc    c.Uint32T
	X_padding   [16]c.Uint8T
}
type EtsSecureBootSigBlockT EtsSecureBootSigBlock

type EtsSecureBootSignature struct {
	Block     [3]EtsSecureBootSigBlockT
	X_padding [448]c.Uint8T
}
type EtsSecureBootSignatureT EtsSecureBootSignature

type EtsSecureBootKeyDigests struct {
	KeyDigests     [3]c.Pointer
	AllowKeyRevoke bool
}
type EtsSecureBootKeyDigestsT EtsSecureBootKeyDigests
type SecureBootStatusT c.Int

const (
	SB_SUCCESS SecureBootStatusT = 978999973
	SB_FAILED  SecureBootStatusT = 1966311518
)

/* Verify bootloader image (reconfigures cache to map), with
   key digests provided as parameters.)

   Can be used to verify secure boot status before enabling
   secure boot permanently.

   If result is ETS_OK, the "simple hash" of the bootloader is
   copied into verified_hash.
*/
//go:linkname EtsSecureBootVerifyBootloaderWithKeys C.ets_secure_boot_verify_bootloader_with_keys
func EtsSecureBootVerifyBootloaderWithKeys(verified_hash *c.Uint8T, trusted_keys *EtsSecureBootKeyDigestsT, stage_load bool) SecureBootStatusT

/* Verify supplied signature against supplied digest, using
   supplied trusted key digests.

   Doesn't reconfigure cache or any other hardware access.
*/
// llgo:link (*EtsSecureBootSignatureT).EtsSecureBootVerifySignature C.ets_secure_boot_verify_signature
func (recv_ *EtsSecureBootSignatureT) EtsSecureBootVerifySignature(image_digest *c.Uint8T, trusted_keys *EtsSecureBootKeyDigestsT, verified_digest *c.Uint8T) SecureBootStatusT {
	return 0
}

/* Read key digests from efuse. Any revoked/missing digests will be
   marked as NULL

   Returns 0 if at least one valid digest was found.
*/
// llgo:link (*EtsSecureBootKeyDigestsT).EtsSecureBootReadKeyDigests C.ets_secure_boot_read_key_digests
func (recv_ *EtsSecureBootKeyDigestsT) EtsSecureBootReadKeyDigests() ETSSTATUS {
	return 0
}
