package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AESTYPE c.Int

const (
	AES_ENC AESTYPE = 0
	AES_DEC AESTYPE = 1
)

type AESBITS c.Int

const (
	AES128 AESBITS = 0
	AES192 AESBITS = 1
	AES256 AESBITS = 2
)

//go:linkname EtsAesEnable C.ets_aes_enable
func EtsAesEnable()

//go:linkname EtsAesDisable C.ets_aes_disable
func EtsAesDisable()

// llgo:link AESTYPE.EtsAesSetkey C.ets_aes_setkey
func (recv_ AESTYPE) EtsAesSetkey(key c.Pointer, bits AESBITS) c.Int {
	return 0
}

//go:linkname EtsAesSetkeyEnc C.ets_aes_setkey_enc
func EtsAesSetkeyEnc(key c.Pointer, bits AESBITS) c.Int

//go:linkname EtsAesSetkeyDec C.ets_aes_setkey_dec
func EtsAesSetkeyDec(key c.Pointer, bits AESBITS) c.Int

//go:linkname EtsAesBlock C.ets_aes_block
func EtsAesBlock(input c.Pointer, output c.Pointer)
