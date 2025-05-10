package main

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
	"github.com/goplus/llpkg/lvgl/widgets"
)

func lv_example_style_1() {
	style := core.LvStyleT{}
	core.LvStyleInit(&style)
	core.LvStyleSetRadius(&style, 5)

	/*Make a gradient*/
	core.LvStyleSetWidth(&style, 150)
	core.LvStyleSetHeight(&style, core.LV_SIZE_CONTENT)

	core.LvStyleSetPadVer(&style, 20)
	core.LvStyleSetPadLeft(&style, 5)

	core.LvStyleSetX(&style, c.Int(core.LvPct(c.Int32T(50))))
	core.LvStyleSetY(&style, c.Int(80))

	/*Create an object with the new style*/
	obj := core.LvObjCreate(display.LvScreenActive())
	core.LvObjAddStyle(obj, &style, 0)

	label := widgets.LvLabelCreate(obj)
	widgets.LvLabelSetText(label, c.Str("Hello"))
}
