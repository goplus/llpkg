package main

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
	"github.com/goplus/llpkg/lvgl/widgets"
)

var label *core.LvObjT

func slider_event_cb(e *core.LvEventT) {
	slider := core.LvEventGetTarget(e)

	/*Refresh the text*/
	widgets.LvLabelSetTextFmt(label, c.Str("%d"), widgets.LvSliderGetValue((*core.LvObjT)(slider)))
	core.LvObjAlignTo(label, (*core.LvObjT)(slider), core.LV_ALIGN_OUT_TOP_MID, 0, -15) /*Align top of the slider*/
}

/**
 * Create a slider and write its value on a label.
 */
func lv_example_get_started_4() {
	/*Create a slider in the center of the display*/
	slider := widgets.LvSliderCreate(display.LvScreenActive())
	core.LvObjSetWidth(slider, 200)                                                 /*Set the width*/
	core.LvObjCenter(slider)                                                        /*Align to the center of the parent (screen)*/
	core.LvObjAddEventCb(slider, slider_event_cb, core.LV_EVENT_VALUE_CHANGED, nil) /*Assign an event function*/

	/*Create a label above the slider*/
	label = widgets.LvLabelCreate(display.LvScreenActive())
	widgets.LvLabelSetText(label, c.Str("0"))
	core.LvObjAlignTo(label, slider, core.LV_ALIGN_OUT_TOP_MID, 0, -15) /*Align top of the slider*/
}
