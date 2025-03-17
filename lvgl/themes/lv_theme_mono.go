package themes

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
	font "github.com/goplus/llpkg/lvgl/font"
)

//go:linkname LvThemeMonoInit C.lv_theme_mono_init
func LvThemeMonoInit(disp *core.LvDisplayT, dark c.Char, font *font.LvFontT) *core.LvThemeT

//go:linkname LvThemeMonoGet C.lv_theme_mono_get
func LvThemeMonoGet() *core.LvThemeT

//go:linkname LvThemeMonoIsInited C.lv_theme_mono_is_inited
func LvThemeMonoIsInited() c.Char

//go:linkname LvThemeMonoDeinit C.lv_theme_mono_deinit
func LvThemeMonoDeinit()
