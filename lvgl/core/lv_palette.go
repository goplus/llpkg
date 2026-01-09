package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// LvPaletteT represents the color palette in LVGL
type LvPaletteT uint8

// Palette constants
const (
	LV_PALETTE_RED         LvPaletteT = iota // Red palette
	LV_PALETTE_PINK                          // Pink palette
	LV_PALETTE_PURPLE                        // Purple palette
	LV_PALETTE_DEEP_PURPLE                   // Deep purple palette
	LV_PALETTE_INDIGO                        // Indigo palette
	LV_PALETTE_BLUE                          // Blue palette
	LV_PALETTE_LIGHT_BLUE                    // Light blue palette
	LV_PALETTE_CYAN                          // Cyan palette
	LV_PALETTE_TEAL                          // Teal palette
	LV_PALETTE_GREEN                         // Green palette
	LV_PALETTE_LIGHT_GREEN                   // Light green palette
	LV_PALETTE_LIME                          // Lime palette
	LV_PALETTE_YELLOW                        // Yellow palette
	LV_PALETTE_AMBER                         // Amber palette
	LV_PALETTE_ORANGE                        // Orange palette
	LV_PALETTE_DEEP_ORANGE                   // Deep orange palette
	LV_PALETTE_BROWN                         // Brown palette
	LV_PALETTE_BLUE_GREY                     // Blue grey palette
	LV_PALETTE_GREY                          // Grey palette
	LV_PALETTE_LAST                          // Last palette (for iteration)
	LV_PALETTE_NONE        LvPaletteT = 0xff // No palette
)

//go:linkname LvPaletteMain C.lv_palette_main
func LvPaletteMain(p LvPaletteT) LvColorT

//go:linkname LvPaletteLighten C.lv_palette_lighten
func LvPaletteLighten(p LvPaletteT, lvl c.Uint8T) LvColorT

//go:linkname LvPaletteDarken C.lv_palette_darken
func LvPaletteDarken(p LvPaletteT, lvl c.Uint8T) LvColorT
