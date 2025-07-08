package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type HalUtilsDivRoundOptT c.Int

const (
	HAL_DIV_ROUND_DOWN HalUtilsDivRoundOptT = 0
	HAL_DIV_ROUND_UP   HalUtilsDivRoundOptT = 1
	HAL_DIV_ROUND      HalUtilsDivRoundOptT = 2
)

/**
 * @brief Clock information
 *
 */

type HalUtilsClkInfoT struct {
	SrcFreqHz c.Uint32T
	ExpFreqHz c.Uint32T
	MaxInteg  c.Uint32T
	MinInteg  c.Uint32T
}

/**
 * @brief Members of clock division
 *
 */

type HalUtilsClkDivT struct {
	Integer     c.Uint32T
	Denominator c.Uint32T
	Numerator   c.Uint32T
}

/**
 * @brief Calculate the clock division with fractal part fast
 * @note  Speed first algorithm, Time complexity O(log n).
 *        About 8~10 times faster than the accurate algorithm
 *
 * @param[in]  clk_info     The clock information
 * @param[out] clk_div      The clock division with integral and fractal part
 * @return
 *      - 0: Failed to get the result because the division is out of range
 *      - others: The real output clock frequency
 */
// llgo:link (*HalUtilsClkInfoT).HalUtilsCalcClkDivFracFast C.hal_utils_calc_clk_div_frac_fast
func (recv_ *HalUtilsClkInfoT) HalUtilsCalcClkDivFracFast(clk_div *HalUtilsClkDivT) c.Uint32T {
	return 0
}

/**
 * @brief Calculate the clock division with fractal part accurately
 * @note  Accuracy first algorithm, Time complexity O(n).
 *        About 1~hundreds times more accurate than the fast algorithm
 *
 * @param[in]  clk_info     The clock information
 * @param[out] clk_div      The clock division with integral and fractal part
 * @return
 *      - 0: Failed to get the result because the division is out of range
 *      - others: The real output clock frequency
 */
// llgo:link (*HalUtilsClkInfoT).HalUtilsCalcClkDivFracAccurate C.hal_utils_calc_clk_div_frac_accurate
func (recv_ *HalUtilsClkInfoT) HalUtilsCalcClkDivFracAccurate(clk_div *HalUtilsClkDivT) c.Uint32T {
	return 0
}

/**
 * @brief Calculate the clock division without fractal part
 *
 * @param[in]  clk_info     The clock information
 * @param[out] int_div      The clock integral division
 * @return
 *      - 0: Failed to get the result because the division is out of range,
 *           but parameter `int_div` will still be assigned to min/max division that given in `clk_info`,
 *           in case the caller still want to use the min/max division in this case.
 *      - others: The real output clock frequency
 */
// llgo:link (*HalUtilsClkInfoT).HalUtilsCalcClkDivInteger C.hal_utils_calc_clk_div_integer
func (recv_ *HalUtilsClkInfoT) HalUtilsCalcClkDivInteger(int_div *c.Uint32T) c.Uint32T {
	return 0
}

/**
 * @brief Fixed-point data configuration
 *
 */

type HalUtilsFixedPointT struct {
	IntBit     c.Uint32T
	FracBit    c.Uint32T
	Saturation bool
}

/**
 * @brief Convert the float type to fixed point type
 * @note  The supported data format:
 *        - [input] float (IEEE 754):
 *          sign(1bit) + exponent(8bit) + mantissa(23bit)       (32 bit in total)
 *        - [output] fixed-point:
 *          sign(1bit) + integer(int_bit) + fraction(frac_bit)  (less or equal to 32 bit)
 *
 * @param[in]  flt          IEEE 754 float type data
 * @param[in]  fp_cfg       Fixed-point data configuration
 * @param[out] fp_out       The output fixed-point data
 * @return
 *      0:              Success
 *      -1:             Fixed point data overflow, `fp_out` will still be assigned
 *      -2:             Float is NaN
 *      -3:             Invalid configuration
 */
//go:linkname HalUtilsFloatToFixedPoint32b C.hal_utils_float_to_fixed_point_32b
func HalUtilsFloatToFixedPoint32b(flt c.Float, fp_cfg *HalUtilsFixedPointT, fp_out *c.Uint32T) c.Int
