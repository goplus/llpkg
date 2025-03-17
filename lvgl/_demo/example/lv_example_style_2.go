package main

import (
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
)

func lv_example_style_2() {
	var style core.LvStyleT
	core.LvStyleInit(&style)
	core.LvStyleSetRadius(&style, 5)

	/*Make a gradient*/
	core.LvStyleSetBgOpa(&style, core.LV_OPA_COVER)
	var grad core.LvGradDscT
	// 设置渐变方向
	grad.SetDir(core.LV_GRAD_DIR_VER)

	// 设置渐变停止点数量
	grad.StopsCount = 2

	// 设置第一个渐变停止点
	grad.Stops[0].Color = core.LvPaletteLighten(core.LV_PALETTE_GREY, 1)
	grad.Stops[0].Opa = core.LV_OPA_COVER
	grad.Stops[0].Frac = 128 // 移动渐变到下方

	// 设置第二个渐变停止点
	grad.Stops[1].Color = core.LvPaletteMain(core.LV_PALETTE_BLUE)
	grad.Stops[1].Opa = core.LV_OPA_COVER
	grad.Stops[1].Frac = 192 // 移动渐变到下方

	core.LvStyleSetBgGrad(&style, &grad)

	/*Create an object with the new style*/
	obj := core.LvObjCreate(display.LvScreenActive())
	core.LvObjAddStyle(obj, &style, 0)
	core.LvObjCenter(obj)
}
