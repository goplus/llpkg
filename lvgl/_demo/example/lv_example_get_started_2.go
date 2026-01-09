package main

import (
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
	"github.com/goplus/llpkg/lvgl/widgets"
)

//var cnt uint8 = 0

func btn_event_cb(e *core.LvEventT) {
	code := core.LvEventGetCode(e)
	btn := core.LvEventGetTarget(e)
	//fmt.Println("code", code)
	if code == core.LV_EVENT_CLICKED {

		//cnt++
		/*Get the first child of the button which is the label and change its text*/
		label := core.LvObjGetChild((*core.LvObjT)(unsafe.Pointer(btn)), 0)

		// btnstr := fmt.Sprintf("Button: %d", cnt)
		// c.Str(btnstr)

		widgets.LvLabelSetTextFmt(label, c.Str("test"))
	}
}

/**
 * Create a button with a label and react on click event.
 */
func lv_example_get_started_2() {
	btn := widgets.LvButtonCreate(display.LvScreenActive())         /*Add a button the current screen*/
	core.LvObjSetPos(btn, 10, 10)                                   /*Set its position*/
	core.LvObjSetSize(btn, 120, 50)                                 /*Set its size*/
	core.LvObjAddEventCb(btn, btn_event_cb, core.LV_EVENT_ALL, nil) /*Assign a callback to the button*/

	label := widgets.LvLabelCreate(btn)            /*Add a label to the button*/
	widgets.LvLabelSetText(label, c.Str("Button")) /*Set the labels text*/
	core.LvObjCenter(label)
}
