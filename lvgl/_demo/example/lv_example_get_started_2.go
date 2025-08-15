package main

import (
	"github.com/goplus/lib/c"
	"github.com/goplus/llpkg/lvgl"
)

//var cnt uint8 = 0

func btn_event_cb(e *lvgl.EventT) {
	code := e.EventGetCode()
	btn := e.EventGetTarget()
	//fmt.Println("code", code)
	if code == lvgl.EVENT_CLICKED {

		//cnt++
		/*Get the first child of the button which is the label and change its text*/
		label := (*lvgl.ObjT)(btn).ObjGetChild(0)

		// c.Str(btnstr)

		label.LabelSetTextFmt(c.Str("test"))
	}
}

/**
 * Create a button with a label and react on click event.
 */
func lv_example_get_started_2() {
	objT := lvgl.ScreenActive()
	btn := objT.ButtonCreate()                           /*Add a button the current screen*/
	btn.ObjSetPos(10, 10)                                /*Set its position*/
	btn.ObjSetSize(120, 50)                              /*Set its size*/
	btn.ObjAddEventCb(btn_event_cb, lvgl.EVENT_ALL, nil) /*Assign a callback to the button*/

	label := btn.LabelCreate()          /*Add a label to the button*/
	label.LabelSetText(c.Str("Button")) /*Set the labels text*/
	label.ObjCenter()
}
