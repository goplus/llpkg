package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

/**********************
 * GLOBAL PROTOTYPES
 **********************/

/**
 * Get the pixel size of a color format in bits, bpp
 * @param cf        a color format (`LV_COLOR_FORMAT_...`)
 * @return          the pixel size in bits
 * @sa              LV_COLOR_FORMAT_GET_BPP
 */
//go:linkname LvColorFormatGetBpp C.lv_color_format_get_bpp
func LvColorFormatGetBpp(cf LvColorFormatT) c.Uint8T

/**
 * Get the pixel size of a color format in bytes
 * @param cf        a color format (`LV_COLOR_FORMAT_...`)
 * @return          the pixel size in bytes
 * @sa              LV_COLOR_FORMAT_GET_SIZE
 */
//go:linkname LvColorFormatGetSize C.lv_color_format_get_size
func LvColorFormatGetSize(cf LvColorFormatT) c.Uint8T

/**
 * Check if a color format has alpha channel or not
 * @param src_cf    a color format (`LV_COLOR_FORMAT_...`)
 * @return          true: has alpha channel; false: doesn't have alpha channel
 */
//go:linkname LvColorFormatHasAlpha C.lv_color_format_has_alpha
func LvColorFormatHasAlpha(srcCf LvColorFormatT) c.Char

/**
 * Create an ARGB8888 color from RGB888 + alpha
 * @param color     an RGB888 color
 * @param opa       the alpha value
 * @return          the ARGB8888 color
 */
//go:linkname LvColorTo32 C.lv_color_to_32
func LvColorTo32(color LvColorT, opa LvOpaT) LvColor32T

/**
 * Convert an RGB888 color to an integer
 * @param c     an RGB888 color
 * @return      `c` as an integer
 */
//go:linkname LvColorToInt C.lv_color_to_int
func LvColorToInt(c LvColorT) c.Uint32T

/**
 * Check if two RGB888 color are equal
 * @param c1    the first color
 * @param c2    the second color
 * @return      true: equal
 */
//go:linkname LvColorEq C.lv_color_eq
func LvColorEq(c1 LvColorT, c2 LvColorT) c.Char

/**
 * Check if two ARGB8888 color are equal
 * @param c1    the first color
 * @param c2    the second color
 * @return      true: equal
 */
//go:linkname LvColor32Eq C.lv_color32_eq
func LvColor32Eq(c1 LvColor32T, c2 LvColor32T) c.Char

/**
 * Create a color from 0x000000..0xffffff input
 * @param c     the hex input
 * @return      the color
 */
//go:linkname LvColorHex C.lv_color_hex
func LvColorHex(c c.Uint32T) LvColorT

/**
 * Create an RGB888 color
 * @param r     the red channel (0..255)
 * @param g     the green channel (0..255)
 * @param b     the blue channel (0..255)
 * @return      the color
 */
//go:linkname LvColorMake C.lv_color_make
func LvColorMake(r c.Uint8T, g c.Uint8T, b c.Uint8T) LvColorT

/**
 * Create an ARGB8888 color
 * @param r     the red channel (0..255)
 * @param g     the green channel (0..255)
 * @param b     the blue channel (0..255)
 * @param a     the alpha channel (0..255)
 * @return      the color
 */
//go:linkname LvColor32Make C.lv_color32_make
func LvColor32Make(r c.Uint8T, g c.Uint8T, b c.Uint8T, a c.Uint8T) LvColor32T

/**
 * Create a color from 0x000..0xfff input
 * @param c     the hex input (e.g. 0x123 will be 0x112233)
 * @return      the color
 */
//go:linkname LvColorHex3 C.lv_color_hex3
func LvColorHex3(c c.Uint32T) LvColorT

/**
 * Convert am RGB888 color to RGB565 stored in `uint16_t`
 * @param color     and RGB888 color
 * @return          `color` as RGB565 on `uin16_t`
 */
//go:linkname LvColorToU16 C.lv_color_to_u16
func LvColorToU16(color LvColorT) c.Uint16T

/**
 * Convert am RGB888 color to XRGB8888 stored in `uint32_t`
 * @param color     and RGB888 color
 * @return          `color` as XRGB8888 on `uin32_t` (the alpha channel is always set to 0xFF)
 */
//go:linkname LvColorToU32 C.lv_color_to_u32
func LvColorToU32(color LvColorT) c.Uint32T

/**
 * Mix two RGB565 colors
 * @param c1        the first color (typically the foreground color)
 * @param c2        the second color  (typically the background color)
 * @param mix       0..255, or LV_OPA_0/10/20...
 * @return          mix == 0: c2
 *                  mix == 255: c1
 *                  mix == 128: 0.5 x c1 + 0.5 x c2
 */
//go:linkname LvColor1616Mix C.lv_color_16_16_mix
func LvColor1616Mix(c1 c.Uint16T, c2 c.Uint16T, mix c.Uint8T) c.Uint16T

/**
 * Mix white to a color
 * @param c     the base color
 * @param lvl   the intensity of white (0: no change, 255: fully white)
 * @return      the mixed color
 */
//go:linkname LvColorLighten C.lv_color_lighten
func LvColorLighten(c LvColorT, lvl LvOpaT) LvColorT

/**
 * Mix black to a color
 * @param c     the base color
 * @param lvl   the intensity of black (0: no change, 255: fully black)
 * @return      the mixed color
 */
//go:linkname LvColorDarken C.lv_color_darken
func LvColorDarken(c LvColorT, lvl LvOpaT) LvColorT

/**
 * Convert a HSV color to RGB
 * @param h hue [0..359]
 * @param s saturation [0..100]
 * @param v value [0..100]
 * @return the given RGB color in RGB (with LV_COLOR_DEPTH depth)
 */
//go:linkname LvColorHsvToRgb C.lv_color_hsv_to_rgb
func LvColorHsvToRgb(h c.Uint16T, s c.Uint8T, v c.Uint8T) LvColorT

/**
 * Convert a 32-bit RGB color to HSV
 * @param r8 8-bit red
 * @param g8 8-bit green
 * @param b8 8-bit blue
 * @return the given RGB color in HSV
 */
//go:linkname LvColorRgbToHsv C.lv_color_rgb_to_hsv
func LvColorRgbToHsv(r8 c.Uint8T, g8 c.Uint8T, b8 c.Uint8T) LvColorHsvT

/**
 * Convert a color to HSV
 * @param color color
 * @return the given color in HSV
 */
//go:linkname LvColorToHsv C.lv_color_to_hsv
func LvColorToHsv(color LvColorT) LvColorHsvT

/*Source: https://vuetifyjs.com/en/styles/colors/#material-colors*/

/**
 * A helper for white color
 * @return      a white color
 */
//go:linkname LvColorWhite C.lv_color_white
func LvColorWhite() LvColorT

/**
 * A helper for black color
 * @return      a black color
 */
//go:linkname LvColorBlack C.lv_color_black
func LvColorBlack() LvColorT

//go:linkname LvColorPremultiply C.lv_color_premultiply
func LvColorPremultiply(c *LvColor32T)

//go:linkname LvColor16Premultiply C.lv_color16_premultiply
func LvColor16Premultiply(c *LvColor16T, a LvOpaT)

/**
 * Get the luminance of a color: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
//go:linkname LvColorLuminance C.lv_color_luminance
func LvColorLuminance(c LvColorT) c.Uint8T

/**
 * Get the luminance of a color16: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
//go:linkname LvColor16Luminance C.lv_color16_luminance
func LvColor16Luminance(c LvColor16T) c.Uint8T

/**
 * Get the luminance of a color24: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
//go:linkname LvColor24Luminance C.lv_color24_luminance
func LvColor24Luminance(c *c.Uint8T) c.Uint8T

/**
 * Get the luminance of a color32: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
//go:linkname LvColor32Luminance C.lv_color32_luminance
func LvColor32Luminance(c LvColor32T) c.Uint8T
