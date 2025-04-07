package main

import (
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
)

func lv_example_style_3() {
	var style core.LvStyleT
	core.LvStyleInit(&style)

	/*Set a background color and a radius*/
	core.LvStyleSetRadius(&style, 10)
	core.LvStyleSetBgOpa(&style, core.LV_OPA_COVER)
	core.LvStyleSetBgColor(&style, core.LvPaletteLighten(core.LV_PALETTE_GREY, 1))

	/*Add border to the bottom+right*/
	core.LvStyleSetBorderColor(&style, core.LvPaletteMain(core.LV_PALETTE_BLUE))
	core.LvStyleSetBorderWidth(&style, 5)
	core.LvStyleSetBorderOpa(&style, core.LV_OPA_50)
	core.LvStyleSetBorderSide(&style, core.LV_BORDER_SIDE_BOTTOM|core.LV_BORDER_SIDE_RIGHT)

	/*Create an object with the new style*/
	obj := core.LvObjCreate(display.LvScreenActive())
	core.LvObjAddStyle(obj, &style, 0)
	core.LvObjCenter(obj)
}
