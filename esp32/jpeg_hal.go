package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type JpegDevT struct {
	Unused [8]uint8
}
type JpegSocHandleT *JpegDevT

/**
 * Context that should be maintained by both the driver and the HAL
 */

type JpegHalContextT struct {
	Dev JpegSocHandleT
}

/**
 * @brief Initialize the JPEG codec HAL driver
 *
 * @param hal: JPEG codec HAL context
 */
// llgo:link (*JpegHalContextT).JpegHalInit C.jpeg_hal_init
func (recv_ *JpegHalContextT) JpegHalInit() {
}

/**
 * @brief Deinitialize the JPEG codec HAL driver
 *
 * @param hal: JPEG codec HAL context
 */
// llgo:link (*JpegHalContextT).JpegHalDeinit C.jpeg_hal_deinit
func (recv_ *JpegHalContextT) JpegHalDeinit() {
}

// llgo:type C
type JpegConfigDhtTableT func(*JpegHalContextT, *c.Uint8T, *c.Uint8T, *c.Uint32T)

// llgo:type C
type JpegConfigFrameInfoT func(JpegSocHandleT, c.Uint8T, c.Uint8T, c.Uint8T, c.Uint8T)

// llgo:type C
type JpegConfigQuantizationCoefficientT func(JpegSocHandleT, *c.Uint32T)

/**
 * Set the quantization coefficients for luminance and chrominance in the JPEG hardware accelerator context.
 *
 * @param hal Pointer to the JPEG hardware accelerator context.
 * @param lqnr Pointer to an array of luminance quantization coefficients.
 * @param cqnr Pointer to an array of chrominance quantization coefficients.
 */
// llgo:link (*JpegHalContextT).JpegHalSetQuantizationCoefficient C.jpeg_hal_set_quantization_coefficient
func (recv_ *JpegHalContextT) JpegHalSetQuantizationCoefficient(lqnr *c.Uint32T, cqnr *c.Uint32T) {
}
