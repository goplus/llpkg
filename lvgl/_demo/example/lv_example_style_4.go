package main

import (
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
)

func lv_example_style_4() {
	var style core.LvStyleT
	core.LvStyleInit(&style)

	/*Set a background color and a radius*/
	core.LvStyleSetRadius(&style, 5)
	core.LvStyleSetBgOpa(&style, core.LV_OPA_COVER)
	core.LvStyleSetBgColor(&style, core.LvPaletteLighten(core.LV_PALETTE_GREY, 1))

	/*Add outline*/
	core.LvStyleSetOutlineWidth(&style, 2)
	core.LvStyleSetOutlineColor(&style, core.LvPaletteMain(core.LV_PALETTE_BLUE))
	core.LvStyleSetOutlinePad(&style, 8)

	/*Create an object with the new style*/
	obj := core.LvObjCreate(display.LvScreenActive())
	core.LvObjAddStyle(obj, &style, 0)
	core.LvObjCenter(obj)
}
