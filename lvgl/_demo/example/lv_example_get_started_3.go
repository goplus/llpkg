package main

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
	"github.com/goplus/llpkg/lvgl/widgets"
)

var styleBtn core.LvStyleT
var styleButtonPressed core.LvStyleT
var styleButtonRed core.LvStyleT

func darken(dsc *core.LvColorFilterDscT, color core.LvColorT, opa core.LvOpaT) core.LvColorT {
	// LV_UNUSED(dsc) in Go we just don't use the parameter
	return core.LvColorDarken(color, opa)
}

func styleInit() {
	/*Create a simple button style*/
	core.LvStyleInit(&styleBtn)
	core.LvStyleSetRadius(&styleBtn, 10)
	core.LvStyleSetBgOpa(&styleBtn, core.LV_OPA_COVER)
	core.LvStyleSetBgColor(&styleBtn, core.LvPaletteLighten(core.LV_PALETTE_GREY, c.Uint8T(3)))
	core.LvStyleSetBgGradColor(&styleBtn, core.LvPaletteMain(core.LV_PALETTE_GREY))
	core.LvStyleSetBgGradDir(&styleBtn, core.LV_GRAD_DIR_VER)

	core.LvStyleSetBorderColor(&styleBtn, core.LvColorBlack())
	core.LvStyleSetBorderOpa(&styleBtn, core.LV_OPA_20)
	core.LvStyleSetBorderWidth(&styleBtn, c.Int(2))

	core.LvStyleSetTextColor(&styleBtn, core.LvColorBlack())

	/*Create a style for the pressed state.
	 *Use a color filter to simply modify all colors in this state*/
	var colorFilter core.LvColorFilterDscT
	core.LvColorFilterDscInit(&colorFilter, darken)
	core.LvStyleInit(&styleButtonPressed)
	core.LvStyleSetColorFilterDsc(&styleButtonPressed, &colorFilter)
	core.LvStyleSetColorFilterOpa(&styleButtonPressed, core.LV_OPA_20)

	/*Create a red style. Change only some colors.*/
	core.LvStyleInit(&styleButtonRed)
	core.LvStyleSetBgColor(&styleButtonRed, core.LvPaletteMain(core.LV_PALETTE_RED))
	core.LvStyleSetBgGradColor(&styleButtonRed, core.LvPaletteLighten(core.LV_PALETTE_RED, 3))
}

/**
 * Create core from scratch for buttons.
 */
func lv_example_get_started_3() {
	/*Initialize the style*/
	styleInit()

	/*Create a button and use the new core*/
	btn := widgets.LvButtonCreate(display.LvScreenActive())
	/* Remove the core coming from the theme
	 * Note that size and position are also stored as style properties
	 * so lv_obj_remove_style_all will remove the set size and position too */
	core.LvObjRemoveStyleAll(btn)
	core.LvObjSetPos(btn, 10, 10)
	core.LvObjSetSize(btn, 120, 50)
	core.LvObjAddStyle(btn, &styleBtn, 0)
	core.LvObjAddStyle(btn, &styleButtonPressed, core.LV_STATE_PRESSED)

	/*Add a label to the button*/
	label := widgets.LvLabelCreate(btn)
	widgets.LvLabelSetText(label, c.Str("Button"))
	core.LvObjCenter(label)

	/*Create another button and use the red style too*/
	btn2 := widgets.LvButtonCreate(display.LvScreenActive())
	core.LvObjRemoveStyleAll(btn2) /*Remove the core coming from the theme*/
	core.LvObjSetPos(btn2, 10, 80)
	core.LvObjSetSize(btn2, 120, 50)
	core.LvObjAddStyle(btn2, &styleBtn, 0)
	core.LvObjAddStyle(btn2, &styleButtonRed, 0)
	core.LvObjAddStyle(btn2, &styleButtonPressed, core.LV_STATE_PRESSED)
	//core.LvObjSetStyleRadius(btn2, core.LV_RADIUS_CIRCLE, 0) /*Add a local style too*/

	label = widgets.LvLabelCreate(btn2)
	widgets.LvLabelSetText(label, c.Str("Button 2"))
	core.LvObjCenter(label)
}
