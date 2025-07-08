package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DsKeyCheckT c.Int

const (
	DS_KEY_INPUT_OK DsKeyCheckT = 0
	DS_NO_KEY_INPUT DsKeyCheckT = 1
	DS_OTHER_WRONG  DsKeyCheckT = 2
)

type DsSignatureCheckT c.Int

const (
	DS_SIGNATURE_OK                  DsSignatureCheckT = 0
	DS_SIGNATURE_PADDING_FAIL        DsSignatureCheckT = 1
	DS_SIGNATURE_MD_FAIL             DsSignatureCheckT = 2
	DS_SIGNATURE_PADDING_AND_MD_FAIL DsSignatureCheckT = 3
)
