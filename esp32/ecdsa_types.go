package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EcdsaModeT c.Int

const (
	ECDSA_MODE_SIGN_VERIFY   EcdsaModeT = 0
	ECDSA_MODE_SIGN_GEN      EcdsaModeT = 1
	ECDSA_MODE_EXPORT_PUBKEY EcdsaModeT = 2
)

type EcdsaCurveT c.Int

const (
	ECDSA_CURVE_SECP192R1 EcdsaCurveT = 0
	ECDSA_CURVE_SECP256R1 EcdsaCurveT = 1
)

type EcdsaShaModeT c.Int

const (
	ECDSA_Z_USE_SHA_PERI  EcdsaShaModeT = 0
	ECDSA_Z_USER_PROVIDED EcdsaShaModeT = 1
)

type EcdsaSignTypeT c.Int

const ECDSA_K_TYPE_TRNG EcdsaSignTypeT = 0
