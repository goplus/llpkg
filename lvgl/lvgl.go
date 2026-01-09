package lvgl

import (
	_ "unsafe"
)

const (
	LLGoPackage = "link: $(pkg-config --libs sdl2 lvgl); -llvgl -lSDL2"
)

//go:linkname LvglInit C.lv_init
func LvglInit()

//go:linkname LvglDeinit C.lv_deinit
func LvglDeinit()

//go:linkname LvglRun C.lv_run
func LvglRun()
