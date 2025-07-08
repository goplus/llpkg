package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * ECDSA peripheral config structure
 */

type EcdsaHalConfigT struct {
	Mode        EcdsaModeT
	Curve       EcdsaCurveT
	ShaMode     EcdsaShaModeT
	EfuseKeyBlk c.Int
	UseKmKey    bool
	SignType    EcdsaSignTypeT
	LoopNumber  c.Uint16T
}

/**
 * @brief Generate ECDSA signature
 *
 * @param conf Configuration for ECDSA operation, see ``ecdsa_hal_config_t``
 * @param hash Hash that is to be signed
 * @param r_out Buffer that will contain `R` component of ECDSA signature
 * @param s_out Buffer that will contain `S` component of ECDSA signature
 * @param len Length of the r_out and s_out buffer (32 bytes for SECP256R1, 24 for SECP192R1)
 */
// llgo:link (*EcdsaHalConfigT).EcdsaHalGenSignature C.ecdsa_hal_gen_signature
func (recv_ *EcdsaHalConfigT) EcdsaHalGenSignature(hash *c.Uint8T, r_out *c.Uint8T, s_out *c.Uint8T, len c.Uint16T) {
}

/**
 * @brief Verify given ECDSA signature
 *
 * @param conf Configuration for ECDSA operation, see ``ecdsa_hal_config_t``
 * @param hash Hash that was signed
 * @param r `R` component of ECDSA signature
 * @param s `S` component of ECDSA signature
 * @param pub_x X coordinate of public key
 * @param pub_y Y coordinate of public key
 * @param len Length of r and s buffer (32 bytes for SECP256R1, 24 for SECP192R1)
 *
 * @return - 0, if the signature matches
 *         - -1, if verification fails
 */
// llgo:link (*EcdsaHalConfigT).EcdsaHalVerifySignature C.ecdsa_hal_verify_signature
func (recv_ *EcdsaHalConfigT) EcdsaHalVerifySignature(hash *c.Uint8T, r *c.Uint8T, s *c.Uint8T, pub_x *c.Uint8T, pub_y *c.Uint8T, len c.Uint16T) c.Int {
	return 0
}

/**
 * @brief Check if the ECDSA operation is successful
 *
 * @return - true, if the ECDSA operation is successful
 *         - false, if the ECDSA operation fails
 */
//go:linkname EcdsaHalGetOperationResult C.ecdsa_hal_get_operation_result
func EcdsaHalGetOperationResult() bool
