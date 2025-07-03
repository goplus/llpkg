package main

import (
	"github.com/goplus/lib/c"
	"github.com/goplus/llpkg/lvgl"
)

//var cnt uint8 = 0

func btn_event_cb(e *lvgl.LvEventT) {
	code := e.LvEventGetCode()
	btn := e.LvEventGetTarget()
	//fmt.Println("code", code)
	if code == lvgl.LV_EVENT_CLICKED {

		//cnt++
		/*Get the first child of the button which is the label and change its text*/
		label := (*lvgl.LvObjT)(btn).LvObjGetChild(0)

		// c.Str(btnstr)

		label.LvLabelSetTextFmt(c.Str("test"))
	}
}

/**
 * Create a button with a label and react on click event.
 */
func lv_example_get_started_2() {
	objT := lvgl.LvScreenActive()
	btn := objT.LvButtonCreate()                                /*Add a button the current screen*/
	btn.LvObjSetPos(10, 10)                                     /*Set its position*/
	btn.LvObjSetSize(120, 50)                                   /*Set its size*/
	btn.LvObjAddEventCb(btn_event_cb, lvgl.LV_EVENT_ALL, nil) /*Assign a callback to the button*/

	label := btn.LvLabelCreate()          /*Add a label to the button*/
	label.LvLabelSetText(c.Str("Button")) /*Set the labels text*/
	label.LvObjCenter()
}
