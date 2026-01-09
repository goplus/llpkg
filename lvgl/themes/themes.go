package themes

import (
	_ "unsafe"

	core "github.com/goplus/llpkg/lvgl/core"
	font "github.com/goplus/llpkg/lvgl/font"
)

// LvThemeApplyCbT is a callback type for theme application
//
//llgo:type C
type LvThemeApplyCbT func(*core.LvThemeT, *core.LvObjT)

//go:linkname LvThemeGetFromObj C.lv_theme_get_from_obj
func LvThemeGetFromObj(obj *core.LvObjT) *core.LvThemeT

//go:linkname LvThemeApply C.lv_theme_apply
func LvThemeApply(obj *core.LvObjT)

//go:linkname LvThemeSetParent C.lv_theme_set_parent
func LvThemeSetParent(newTheme *core.LvThemeT, parent *core.LvThemeT)

//go:linkname LvThemeSetApplyCb C.lv_theme_set_apply_cb
func LvThemeSetApplyCb(theme *core.LvThemeT, applyCb LvThemeApplyCbT)

//go:linkname LvThemeGetFontSmall C.lv_theme_get_font_small
func LvThemeGetFontSmall(obj *core.LvObjT) *font.LvFontT

//go:linkname LvThemeGetFontNormal C.lv_theme_get_font_normal
func LvThemeGetFontNormal(obj *core.LvObjT) *font.LvFontT

//go:linkname LvThemeGetFontLarge C.lv_theme_get_font_large
func LvThemeGetFontLarge(obj *core.LvObjT) *font.LvFontT

//go:linkname LvThemeGetColorPrimary C.lv_theme_get_color_primary
func LvThemeGetColorPrimary(obj *core.LvObjT) core.LvColorT

//go:linkname LvThemeGetColorSecondary C.lv_theme_get_color_secondary
func LvThemeGetColorSecondary(obj *core.LvObjT) core.LvColorT
