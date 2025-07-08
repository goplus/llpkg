package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LcdClockSourceT SocPeriphLcdClkSrcT
type LcdRgbDataEndianT c.Int

const (
	LCD_RGB_DATA_ENDIAN_BIG    LcdRgbDataEndianT = 0
	LCD_RGB_DATA_ENDIAN_LITTLE LcdRgbDataEndianT = 1
)

type LcdColorSpaceT c.Int

const (
	LCD_COLOR_SPACE_RGB LcdColorSpaceT = 2
	LCD_COLOR_SPACE_YUV LcdColorSpaceT = 3
)

type LcdColorRgbPixelFormatT c.Int

const (
	LCD_COLOR_PIXEL_FORMAT_RGB565 LcdColorRgbPixelFormatT = 2
	LCD_COLOR_PIXEL_FORMAT_RGB666 LcdColorRgbPixelFormatT = 1
	LCD_COLOR_PIXEL_FORMAT_RGB888 LcdColorRgbPixelFormatT = 0
)

type LcdColorFormatT c.Int

const (
	LCD_COLOR_FMT_RGB565 LcdColorFormatT = 33554434
	LCD_COLOR_FMT_RGB666 LcdColorFormatT = 33554433
	LCD_COLOR_FMT_RGB888 LcdColorFormatT = 33554432
	LCD_COLOR_FMT_YUV422 LcdColorFormatT = 50331649
)

type LcdColorRangeT c.Int

const (
	LCD_COLOR_RANGE_LIMIT LcdColorRangeT = 0
	LCD_COLOR_RANGE_FULL  LcdColorRangeT = 1
)

type LcdYuvSampleT c.Int

const (
	LCD_YUV_SAMPLE_422 LcdYuvSampleT = 1
	LCD_YUV_SAMPLE_420 LcdYuvSampleT = 2
	LCD_YUV_SAMPLE_411 LcdYuvSampleT = 3
)

type LcdYuvConvStdT c.Int

const (
	LCD_YUV_CONV_STD_BT601 LcdYuvConvStdT = 0
	LCD_YUV_CONV_STD_BT709 LcdYuvConvStdT = 1
)

type LcdYuv422PackOrderT c.Int

const (
	LCD_YUV422_PACK_ORDER_YUYV LcdYuv422PackOrderT = 0
	LCD_YUV422_PACK_ORDER_YVYU LcdYuv422PackOrderT = 1
	LCD_YUV422_PACK_ORDER_UYVY LcdYuv422PackOrderT = 2
	LCD_YUV422_PACK_ORDER_VYUY LcdYuv422PackOrderT = 3
)
