package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Set the work mode of the operation
 *
 * @param mode Mode of operation
 */
// llgo:link EccModeT.EccHalSetMode C.ecc_hal_set_mode
func (recv_ EccModeT) EccHalSetMode() {
}

/**
 * @brief Set the ECC curve of operation
 *
 * @param curve Curve to use for operation
 */
// llgo:link EccCurveT.EccHalSetCurve C.ecc_hal_set_curve
func (recv_ EccCurveT) EccHalSetCurve() {
}

/**
 * @brief Start calculation
 *
 */
//go:linkname EccHalStartCalc C.ecc_hal_start_calc
func EccHalStartCalc()

/**
 * @brief Check whether the calculation has finished
 *
 * @return - 1 if the hardware has finished calculating
 *         - 0 otherwise
 */
//go:linkname EccHalIsCalcFinished C.ecc_hal_is_calc_finished
func EccHalIsCalcFinished() c.Int

/**
 * @brief Write parameters for point multiplication (K * (Px, Py))
 *
 * @param k Scalar value
 * @param px X coordinate of the ECC point
 * @param py Y coordinate of the ECC point
 * @param len Length (in bytes) of the ECC point
 *            - 32 bytes for SECP256R1
 *            - 24 bytes for SECP192R1
 */
//go:linkname EccHalWriteMulParam C.ecc_hal_write_mul_param
func EccHalWriteMulParam(k *c.Uint8T, px *c.Uint8T, py *c.Uint8T, len c.Uint16T)

/**
 * @brief Write parameters for point verification,
 *        i.e to check if the point lies on the curve
 *
 * @param px X coordinate of the ECC point
 * @param py Y coordinate of the ECC point
 * @param len Length (in bytes) of the ECC point
 *            - 32 for SECP256R1
 *            - 24 for SECP192R1
 */
//go:linkname EccHalWriteVerifyParam C.ecc_hal_write_verify_param
func EccHalWriteVerifyParam(px *c.Uint8T, py *c.Uint8T, len c.Uint16T)

/**
 * @brief Read point multiplication result
 *
 * @param rx X coordinate of the multiplication result
 * @param ry Y coordinate of the multiplication result
 * @param len Length (in bytes) of the ECC point
 *            - 32 for SECP256R1
 *            - 24 for SECP192R1
 *
 * @return - 0 if the operation was successful
 *         - -1 if the operation was not successful
 *
 * In case the operation is not successful, rx and ry will contain
 * all zeros
 */
//go:linkname EccHalReadMulResult C.ecc_hal_read_mul_result
func EccHalReadMulResult(rx *c.Uint8T, ry *c.Uint8T, len c.Uint16T) c.Int

/**
 * @brief Read point verification result
 *
 * @return - 1 if point lies on curve
 *         - 0 otherwise
 */
//go:linkname EccHalReadVerifyResult C.ecc_hal_read_verify_result
func EccHalReadVerifyResult() c.Int

/**
 * @brief Enable constant time multiplication operations
 *
 * @param true: enable; false: disable
 */
//go:linkname EccHalEnableConstantTimePointMul C.ecc_hal_enable_constant_time_point_mul
func EccHalEnableConstantTimePointMul(enable bool)
