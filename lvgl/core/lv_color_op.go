package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//llgo:type C
type LvColorFilterCbT func(dsc *LvColorFilterDscT, color LvColorT, opa LvOpaT) LvColorT

/**********************
 * GLOBAL PROTOTYPES
 **********************/

/**
 * Mix two colors with a given ratio.
 * @param c1 the first color to mix (usually the foreground)
 * @param c2 the second color to mix (usually the background)
 * @param mix The ratio of the colors. 0: full `c2`, 255: full `c1`, 127: half `c1` and half`c2`
 * @return the mixed color
 */
//go:linkname LvColorMix C.lv_color_mix
func LvColorMix(c1 LvColorT, c2 LvColorT, mix c.Uint8T) LvColorT

/**
 *
 * @param fg
 * @param bg
 * @return
 * @note Use bg.alpha in the return value
 * @note Use fg.alpha as mix ratio
 */
//go:linkname LvColorMix32 C.lv_color_mix32
func LvColorMix32(fg LvColorT, bg LvColorT) LvColorT

/**
 * Get the brightness of a color
 * @param c   a color
 * @return brightness in range [0..255]
 */
//go:linkname LvColorBrightness C.lv_color_brightness
func LvColorBrightness(c LvColorT) c.Uint8T

//go:linkname LvColorFilterDscInit C.lv_color_filter_dsc_init
func LvColorFilterDscInit(dsc *LvColorFilterDscT, cb LvColorFilterCbT)
