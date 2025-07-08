package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get the bit depth (bits/pixel) of each color space pixel format
 *
 * @param color_type_id Value constructed in color_space_pixel_format_t struct
 *
 * @return Number of bits per pixel
 */
// llgo:link ColorSpacePixelFormatT.ColorHalPixelFormatGetBitDepth C.color_hal_pixel_format_get_bit_depth
func (recv_ ColorSpacePixelFormatT) ColorHalPixelFormatGetBitDepth() c.Uint32T {
	return 0
}
