package core

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvObjSendEvent C.lv_obj_send_event
func LvObjSendEvent(obj *LvObjT, event_code LvEventCodeT, param unsafe.Pointer) LvResultT

//go:linkname LvObjEventBase C.lv_obj_event_base
func LvObjEventBase(class_p *LvObjClassT, e *LvEventT) LvResultT

//go:linkname LvEventGetCurrentTargetObj C.lv_event_get_current_target_obj
func LvEventGetCurrentTargetObj(e *LvEventT) *LvObjT

//go:linkname LvEventGetTargetObj C.lv_event_get_target_obj
func LvEventGetTargetObj(e *LvEventT) *LvObjT

//go:linkname LvObjAddEventCb C.lv_obj_add_event_cb
func LvObjAddEventCb(obj *LvObjT, event_cb LvEventCbT, filter LvEventCodeT, user_data unsafe.Pointer) *lv_event_dsc_t

//go:linkname LvObjGetEventCount C.lv_obj_get_event_count
func LvObjGetEventCount(obj *LvObjT) c.Uint

//go:linkname LvObjGetEventDsc C.lv_obj_get_event_dsc
func LvObjGetEventDsc(obj *LvObjT, index c.Uint) *lv_event_dsc_t

//go:linkname LvObjRemoveEvent C.lv_obj_remove_event
func LvObjRemoveEvent(obj *LvObjT, index c.Uint) bool

//go:linkname LvObjRemoveEventCb C.lv_obj_remove_event_cb
func LvObjRemoveEventCb(obj *LvObjT, event_cb LvEventCbT) bool

//go:linkname LvObjRemoveEventDsc C.lv_obj_remove_event_dsc
func LvObjRemoveEventDsc(obj *LvObjT, dsc *LvEventDscT) bool

//go:linkname LvObjRemoveEventCbWithUserData C.lv_obj_remove_event_cb_with_user_data
func LvObjRemoveEventCbWithUserData(obj *LvObjT, event_cb LvEventCbT, user_data unsafe.Pointer) c.Uint

//go:linkname LvEventGetIndev C.lv_event_get_indev
func LvEventGetIndev(e *LvEventT) *LvIndevT

//go:linkname LvEventGetLayer C.lv_event_get_layer
func LvEventGetLayer(e *LvEventT) *LvLayerT

//go:linkname LvEventGetOldSize C.lv_event_get_old_size
func LvEventGetOldSize(e *LvEventT) *LvAreaT

//go:linkname LvEventGetKey C.lv_event_get_key
func LvEventGetKey(e *LvEventT) c.Uint

//go:linkname LvEventGetRotaryDiff C.lv_event_get_rotary_diff
func LvEventGetRotaryDiff(e *LvEventT) c.Int

//go:linkname LvEventGetScrollAnim C.lv_event_get_scroll_anim
func LvEventGetScrollAnim(e *LvEventT) *LvAnimT

//go:linkname LvEventSetExtDrawSize C.lv_event_set_ext_draw_size
func LvEventSetExtDrawSize(e *LvEventT, size c.Int)

//go:linkname LvEventGetSelfSizeInfo C.lv_event_get_self_size_info
func LvEventGetSelfSizeInfo(e *LvEventT) *LvPointT

//go:linkname LvEventGetHitTestInfo C.lv_event_get_hit_test_info
func LvEventGetHitTestInfo(e *LvEventT) *LvHitTestInfoT

//go:linkname LvEventGetCoverArea C.lv_event_get_cover_area
func LvEventGetCoverArea(e *LvEventT) *LvAreaT

//go:linkname LvEventSetCoverRes C.lv_event_set_cover_res
func LvEventSetCoverRes(e *LvEventT, res LvCoverResT)

//go:linkname LvEventGetDrawTask C.lv_event_get_draw_task
func LvEventGetDrawTask(e *LvEventT) *LvDrawTaskT
