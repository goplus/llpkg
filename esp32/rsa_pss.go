package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ETS_SIG_LEN = 384
const ETS_DIGEST_LEN = 32

type EtsRsaPubkeyT struct {
	N     [384]c.Uint8T
	E     c.Uint32T
	Rinv  [384]c.Uint8T
	Mdash c.Uint32T
}

// llgo:link (*EtsRsaPubkeyT).EtsRsaPssVerify C.ets_rsa_pss_verify
func (recv_ *EtsRsaPubkeyT) EtsRsaPssVerify(sig *c.Uint8T, digest *c.Uint8T, verified_digest *c.Uint8T) bool {
	return false
}

//go:linkname EtsMgf1Sha256 C.ets_mgf1_sha256
func EtsMgf1Sha256(mgfSeed *c.Uint8T, seedLen c.SizeT, maskLen c.SizeT, mask *c.Uint8T)

//go:linkname EtsEmsaPssVerify C.ets_emsa_pss_verify
func EtsEmsaPssVerify(encoded_message *c.Uint8T, mhash *c.Uint8T) bool
