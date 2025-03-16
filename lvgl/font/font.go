package font

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Font glyph format constants
const (
	LvFontGlyphFormatNone      = 0x00
	LvFontGlyphFormatA1        = 0x01
	LvFontGlyphFormatA2        = 0x02
	LvFontGlyphFormatA3        = 0x03
	LvFontGlyphFormatA4        = 0x04
	LvFontGlyphFormatA8        = 0x08
	LvFontGlyphFormatA1Aligned = 0x011
	LvFontGlyphFormatA2Aligned = 0x012
	LvFontGlyphFormatA4Aligned = 0x014
	LvFontGlyphFormatA8Aligned = 0x018
	LvFontGlyphFormatImage     = 0x19
	LvFontGlyphFormatVector    = 0x1A
	LvFontGlyphFormatSvg       = 0x1B
	LvFontGlyphFormatCustom    = 0xFF
)

// Font subpixel rendering modes
const (
	LvFontSubpxNone = iota
	LvFontSubpxHor
	LvFontSubpxVer
	LvFontSubpxBoth
)

// Font kerning modes
const (
	LvFontKerningNormal = iota
	LvFontKerningNone
)

//llgo:type C
type lv_font_t struct {
	Unused [0]byte
}
type LvFontT lv_font_t

//llgo:type C
type lv_font_glyph_dsc_t struct {
	Unused [0]byte
}
type LvFontGlyphDscT lv_font_glyph_dsc_t

//go:linkname LvFontGetGlyphBitmap C.lv_font_get_glyph_bitmap
func LvFontGetGlyphBitmap(gDsc *LvFontGlyphDscT, drawBuf *core.LvDrawBufT) unsafe.Pointer

//go:linkname LvFontGetGlyphDsc C.lv_font_get_glyph_dsc
func LvFontGetGlyphDsc(font *LvFontT, dscOut *LvFontGlyphDscT, letter c.Uint32T, letterNext c.Uint32T) c.Char

//go:linkname LvFontGlyphReleaseDrawData C.lv_font_glyph_release_draw_data
func LvFontGlyphReleaseDrawData(gDsc *LvFontGlyphDscT)

//go:linkname LvFontGetGlyphWidth C.lv_font_get_glyph_width
func LvFontGetGlyphWidth(font *LvFontT, letter c.Uint32T, letterNext c.Uint32T) c.Uint16T

//go:linkname LvFontGetLineHeight C.lv_font_get_line_height
func LvFontGetLineHeight(font *LvFontT) c.Int32T

//go:linkname LvFontSetKerning C.lv_font_set_kerning
func LvFontSetKerning(font *LvFontT, kerning c.Int32T)

//go:linkname LvFontGetDefault C.lv_font_get_default
func LvFontGetDefault() *LvFontT

// Built-in font declarations
//
//go:linkname LvFontMontserrat8 C.lv_font_montserrat_8
var LvFontMontserrat8 LvFontT

//go:linkname LvFontMontserrat10 C.lv_font_montserrat_10
var LvFontMontserrat10 LvFontT

//go:linkname LvFontMontserrat12 C.lv_font_montserrat_12
var LvFontMontserrat12 LvFontT

//go:linkname LvFontMontserrat14 C.lv_font_montserrat_14
var LvFontMontserrat14 LvFontT

//go:linkname LvFontMontserrat16 C.lv_font_montserrat_16
var LvFontMontserrat16 LvFontT

//go:linkname LvFontMontserrat18 C.lv_font_montserrat_18
var LvFontMontserrat18 LvFontT

//go:linkname LvFontMontserrat20 C.lv_font_montserrat_20
var LvFontMontserrat20 LvFontT

//go:linkname LvFontMontserrat22 C.lv_font_montserrat_22
var LvFontMontserrat22 LvFontT

//go:linkname LvFontMontserrat24 C.lv_font_montserrat_24
var LvFontMontserrat24 LvFontT
