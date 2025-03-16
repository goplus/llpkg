package themes

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvThemeSimpleInit C.lv_theme_simple_init
func LvThemeSimpleInit(disp *core.LvDisplayT) *core.LvThemeT

//go:linkname LvThemeSimpleIsInited C.lv_theme_simple_is_inited
func LvThemeSimpleIsInited() c.Char

//go:linkname LvThemeSimpleGet C.lv_theme_simple_get
func LvThemeSimpleGet() *core.LvThemeT

//go:linkname LvThemeSimpleDeinit C.lv_theme_simple_deinit
func LvThemeSimpleDeinit()
