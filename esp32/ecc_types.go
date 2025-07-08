package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EccModeT c.Int

const (
	ECC_MODE_POINT_MUL                 EccModeT = 0
	ECC_MODE_INVERSE_MUL               EccModeT = 1
	ECC_MODE_VERIFY                    EccModeT = 2
	ECC_MODE_VERIFY_THEN_POINT_MUL     EccModeT = 3
	ECC_MODE_JACOBIAN_POINT_MUL        EccModeT = 4
	ECC_MODE_POINT_ADD                 EccModeT = 5
	ECC_MODE_JACOBIAN_POINT_VERIFY     EccModeT = 6
	ECC_MODE_POINT_VERIFY_JACOBIAN_MUL EccModeT = 7
	ECC_MODE_MOD_ADD                   EccModeT = 8
	ECC_MODE_MOD_SUB                   EccModeT = 9
	ECC_MODE_MOD_MUL                   EccModeT = 10
)

type EccCurveT c.Int

const (
	ECC_CURVE_SECP192R1 EccCurveT = 0
	ECC_CURVE_SECP256R1 EccCurveT = 1
)

type EccModBaseT c.Int

const (
	ECC_MOD_N EccModBaseT = 0
	ECC_MOD_P EccModBaseT = 1
)
