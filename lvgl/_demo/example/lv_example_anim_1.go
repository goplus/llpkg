package main

import (
	"log"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
	"github.com/goplus/llpkg/lvgl/widgets"
)

func anim_x_cb(var_ c.Pointer, v int32) {
	obj := (*core.LvObjT)(var_)
	core.LvObjSetX(obj, c.Int32T(v))
}

func sw_event_cb(e *core.LvEventT) {
	log.Println("sw_event_cb")
	sw := (*core.LvObjT)(core.LvEventGetTarget(e))
	label := (*core.LvObjT)(core.LvEventGetUserData(e))

	if core.LvObjHasState(sw, (core.LvStateT)(core.LV_STATE_CHECKED)) {
		a := core.LvAnimT{}
		core.LvAnimInit(&a)
		core.LvAnimSetVar(&a, label)
		core.LvAnimSetValues(&a, core.LvObjGetX(label), 100)
		core.LvAnimSetDuration(&a, 500)
		core.LvAnimSetExecCb(&a, anim_x_cb)
		core.LvAnimSetPathCb(&a, core.LvAnimPathOvershoot)
		core.LvAnimStart(&a)
	} else {
		a := core.LvAnimT{}
		core.LvAnimInit(&a)
		core.LvAnimSetVar(&a, label)
		core.LvAnimSetValues(&a, core.LvObjGetX(label), -core.LvObjGetWidth(label))
		core.LvAnimSetDuration(&a, 500)
		core.LvAnimSetExecCb(&a, anim_x_cb)
		core.LvAnimSetPathCb(&a, core.LvAnimPathEaseIn)
		core.LvAnimStart(&a)
	}
}

/**
 * Start animation on an event
 */
func lv_example_anim_1() {
	label := widgets.LvLabelCreate(display.LvScreenActive())
	widgets.LvLabelSetText(label, c.Str("Hello animations!"))
	core.LvObjSetPos(label, 100, 10)

	sw := widgets.LvSwitchCreate(display.LvScreenActive())
	core.LvObjCenter(sw)
	core.LvObjAddState(sw, core.LV_STATE_CHECKED)
	core.LvObjAddEventCb(sw, sw_event_cb, core.LV_EVENT_VALUE_CHANGED, (c.Pointer)(label))

}
