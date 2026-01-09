package themes

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
	font "github.com/goplus/llpkg/lvgl/font"
)

//go:linkname LvThemeDefaultInit C.lv_theme_default_init
func LvThemeDefaultInit(disp *core.LvDisplayT, colorPrimary core.LvColorT, colorSecondary core.LvColorT, dark c.Char, font *font.LvFontT) *core.LvThemeT

//go:linkname LvThemeDefaultGet C.lv_theme_default_get
func LvThemeDefaultGet() *core.LvThemeT

//go:linkname LvThemeDefaultIsInited C.lv_theme_default_is_inited
func LvThemeDefaultIsInited() c.Char

//go:linkname LvThemeDefaultDeinit C.lv_theme_default_deinit
func LvThemeDefaultDeinit()
