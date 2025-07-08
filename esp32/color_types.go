package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const COLOR_SPACE_BITWIDTH = 8
const COLOR_PIXEL_FORMAT_BITWIDTH = 24

type ColorSpaceT c.Int

const (
	COLOR_SPACE_RAW   ColorSpaceT = 1
	COLOR_SPACE_RGB   ColorSpaceT = 2
	COLOR_SPACE_YUV   ColorSpaceT = 3
	COLOR_SPACE_GRAY  ColorSpaceT = 4
	COLOR_SPACE_ARGB  ColorSpaceT = 5
	COLOR_SPACE_ALPHA ColorSpaceT = 6
	COLOR_SPACE_CLUT  ColorSpaceT = 7
)

type ColorPixelRawFormatT c.Int

const (
	COLOR_PIXEL_RAW8  ColorPixelRawFormatT = 0
	COLOR_PIXEL_RAW10 ColorPixelRawFormatT = 1
	COLOR_PIXEL_RAW12 ColorPixelRawFormatT = 2
)

type ColorPixelRgbFormatT c.Int

const (
	COLOR_PIXEL_RGB888 ColorPixelRgbFormatT = 0
	COLOR_PIXEL_RGB666 ColorPixelRgbFormatT = 1
	COLOR_PIXEL_RGB565 ColorPixelRgbFormatT = 2
)

type ColorPixelYuvFormatT c.Int

const (
	COLOR_PIXEL_YUV444 ColorPixelYuvFormatT = 0
	COLOR_PIXEL_YUV422 ColorPixelYuvFormatT = 1
	COLOR_PIXEL_YUV420 ColorPixelYuvFormatT = 2
	COLOR_PIXEL_YUV411 ColorPixelYuvFormatT = 3
)

type ColorPixelGrayFormatT c.Int

const (
	COLOR_PIXEL_GRAY4 ColorPixelGrayFormatT = 0
	COLOR_PIXEL_GRAY8 ColorPixelGrayFormatT = 1
)

type ColorPixelArgbFormatT c.Int

const COLOR_PIXEL_ARGB8888 ColorPixelArgbFormatT = 0

type ColorPixelAlphaFormatT c.Int

const (
	COLOR_PIXEL_A4 ColorPixelAlphaFormatT = 0
	COLOR_PIXEL_A8 ColorPixelAlphaFormatT = 1
)

type ColorPixelClutFormatT c.Int

const (
	COLOR_PIXEL_L4 ColorPixelClutFormatT = 0
	COLOR_PIXEL_L8 ColorPixelClutFormatT = 1
)

/**
 * @brief Color Space Info Structure
 */

type ColorSpacePixelFormatT struct {
	ColorTypeId c.Uint32T
}
type ColorRangeT c.Int

const (
	COLOR_RANGE_LIMIT ColorRangeT = 0
	COLOR_RANGE_FULL  ColorRangeT = 1
)

type ColorConvStdRgbYuvT c.Int

const (
	COLOR_CONV_STD_RGB_YUV_BT601 ColorConvStdRgbYuvT = 0
	COLOR_CONV_STD_RGB_YUV_BT709 ColorConvStdRgbYuvT = 1
)

type ColorRawElementOrderT c.Int

const (
	COLOR_RAW_ELEMENT_ORDER_BGGR ColorRawElementOrderT = 0
	COLOR_RAW_ELEMENT_ORDER_GBRG ColorRawElementOrderT = 1
	COLOR_RAW_ELEMENT_ORDER_GRBG ColorRawElementOrderT = 2
	COLOR_RAW_ELEMENT_ORDER_RGGB ColorRawElementOrderT = 3
)

type ColorRgbElementOrderT c.Int

const (
	COLOR_RGB_ELEMENT_ORDER_RGB ColorRgbElementOrderT = 0
	COLOR_RGB_ELEMENT_ORDER_BGR ColorRgbElementOrderT = 1
)

/**
 * @brief Data structure for ARGB8888 pixel unit
 */

type ColorPixelArgb8888DataT struct {
	Val c.Uint32T
}

/**
 * @brief Data structure for RGB888 pixel unit
 */

type ColorPixelRgb888DataT struct {
	B c.Uint8T
	G c.Uint8T
	R c.Uint8T
}

/**
 * @brief Data structure for RGB565 pixel unit
 */

type ColorPixelRgb565DataT struct {
	Val c.Uint16T
}
type ColorComponentT c.Int

const (
	COLOR_COMPONENT_R       ColorComponentT = 0
	COLOR_COMPONENT_G       ColorComponentT = 1
	COLOR_COMPONENT_B       ColorComponentT = 2
	COLOR_COMPONENT_INVALID ColorComponentT = 3
)

type ColorYuv422PackOrderT c.Int

const (
	COLOR_YUV422_PACK_ORDER_YUYV ColorYuv422PackOrderT = 0
	COLOR_YUV422_PACK_ORDER_YVYU ColorYuv422PackOrderT = 1
	COLOR_YUV422_PACK_ORDER_UYVY ColorYuv422PackOrderT = 2
	COLOR_YUV422_PACK_ORDER_VYUY ColorYuv422PackOrderT = 3
)
