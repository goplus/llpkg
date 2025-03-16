package core

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:type C
type lv_event_t struct {
	Unused [0]byte
}
type LvEventT lv_event_t

// Event codes
type LvEventCodeT uint32

const (
	LV_EVENT_ALL LvEventCodeT = 0

	// Input device events
	LV_EVENT_PRESSED             LvEventCodeT = iota + 1 // Widget has been pressed
	LV_EVENT_PRESSING                                    // Widget is being pressed (sent continuously while pressing)
	LV_EVENT_PRESS_LOST                                  // Widget is still being pressed but slid cursor/finger off Widget
	LV_EVENT_SHORT_CLICKED                               // Widget was pressed for a short period of time, then released. Not sent if scrolled.
	LV_EVENT_SINGLE_CLICKED                              // Sent for first short click within a small distance and short time
	LV_EVENT_DOUBLE_CLICKED                              // Sent for second short click within small distance and short time
	LV_EVENT_TRIPLE_CLICKED                              // Sent for third short click within small distance and short time
	LV_EVENT_LONG_PRESSED                                // Object has been pressed for at least `long_press_time`. Not sent if scrolled.
	LV_EVENT_LONG_PRESSED_REPEAT                         // Sent after `long_press_time` in every `long_press_repeat_time` ms. Not sent if scrolled.
	LV_EVENT_CLICKED                                     // Sent on release if not scrolled (regardless to long press)
	LV_EVENT_RELEASED                                    // Sent in every cases when Widget has been released
	LV_EVENT_SCROLL_BEGIN                                // Scrolling begins. The event parameter is a pointer to the animation of the scroll. Can be modified
	LV_EVENT_SCROLL_THROW_BEGIN
	LV_EVENT_SCROLL_END // Scrolling ends
	LV_EVENT_SCROLL     // Scrolling
	LV_EVENT_GESTURE    // A gesture is detected. Get gesture with `lv_indev_get_gesture_dir(lv_indev_active());`
	LV_EVENT_KEY        // A key is sent to Widget. Get key with `lv_indev_get_key(lv_indev_active());`
	LV_EVENT_ROTARY     // An encoder or wheel was rotated. Get rotation count with `lv_event_get_rotary_diff(e);`
	LV_EVENT_FOCUSED    // Widget received focus
	LV_EVENT_DEFOCUSED  // Widget's focus has been lost
	LV_EVENT_LEAVE      // Widget's focus has been lost but is still selected
	LV_EVENT_HIT_TEST   // Perform advanced hit-testing
	LV_EVENT_INDEV_RESET
	LV_EVENT_HOVER_OVER  // Indev hover over object
	LV_EVENT_HOVER_LEAVE // Indev hover leave object

	// Drawing events
	LV_EVENT_COVER_CHECK        // Check if Widget fully covers an area. The event parameter is `lv_cover_check_info_t *`.
	LV_EVENT_REFR_EXT_DRAW_SIZE // Get required extra draw area around Widget (e.g. for shadow). The event parameter is `int32_t *` to store the size.
	LV_EVENT_DRAW_MAIN_BEGIN    // Starting the main drawing phase
	LV_EVENT_DRAW_MAIN          // Perform the main drawing
	LV_EVENT_DRAW_MAIN_END      // Finishing the main drawing phase
	LV_EVENT_DRAW_POST_BEGIN    // Starting the post draw phase (when all children are drawn)
	LV_EVENT_DRAW_POST          // Perform the post draw phase (when all children are drawn)
	LV_EVENT_DRAW_POST_END      // Finishing the post draw phase (when all children are drawn)
	LV_EVENT_DRAW_TASK_ADDED    // Adding a draw task

	// Special events
	LV_EVENT_VALUE_CHANGED // Widget's value has changed (i.e. slider moved)
	LV_EVENT_INSERT        // Text has been inserted into Widget. The event data is `char *` being inserted.
	LV_EVENT_REFRESH       // Notify Widget to refresh something on it (for user)
	LV_EVENT_READY         // A process has finished
	LV_EVENT_CANCEL        // A process has been cancelled

	// Other events
	LV_EVENT_CREATE              // Object is being created
	LV_EVENT_DELETE              // Object is being deleted
	LV_EVENT_CHILD_CHANGED       // Child was removed, added, or its size, position were changed
	LV_EVENT_CHILD_CREATED       // Child was created, always bubbles up to all parents
	LV_EVENT_CHILD_DELETED       // Child was deleted, always bubbles up to all parents
	LV_EVENT_SCREEN_UNLOAD_START // A screen unload started, fired immediately when scr_load is called
	LV_EVENT_SCREEN_LOAD_START   // A screen load started, fired when the screen change delay is expired
	LV_EVENT_SCREEN_LOADED       // A screen was loaded
	LV_EVENT_SCREEN_UNLOADED     // A screen was unloaded
	LV_EVENT_SIZE_CHANGED        // Object coordinates/size have changed
	LV_EVENT_STYLE_CHANGED       // Object's style has changed
	LV_EVENT_LAYOUT_CHANGED      // A child's position position has changed due to a layout recalculation
	LV_EVENT_GET_SELF_SIZE       // Get internal size of a widget

	// Events of optional LVGL components
	LV_EVENT_INVALIDATE_AREA      // An area is invalidated (marked for redraw)
	LV_EVENT_RESOLUTION_CHANGED   // Sent when the resolution changes due to `lv_display_set_resolution()` or `lv_display_set_rotation()`.
	LV_EVENT_COLOR_FORMAT_CHANGED // Sent as a result of any call to `lv_display_set_color_format()`.
	LV_EVENT_REFR_REQUEST         // Sent when something happened that requires redraw.
	LV_EVENT_REFR_START           // Sent before a refreshing cycle starts. Sent even if there is nothing to redraw.
	LV_EVENT_REFR_READY           // Sent when refreshing has been completed (after rendering and calling flush callback). Sent even if no redraw happened.
	LV_EVENT_RENDER_START         // Sent just before rendering begins.
	LV_EVENT_RENDER_READY         // Sent after rendering has been completed (before calling flush callback)
	LV_EVENT_FLUSH_START          // Sent before flush callback is called.
	LV_EVENT_FLUSH_FINISH         // Sent after flush callback call has returned.
	LV_EVENT_FLUSH_WAIT_START     // Sent before flush wait callback is called.
	LV_EVENT_FLUSH_WAIT_FINISH    // Sent after flush wait callback call has returned.

	LV_EVENT_VSYNC

	LV_EVENT_LAST // Number of default events

	LV_EVENT_PREPROCESS      LvEventCodeT = 0x8000  // This is a flag that can be set with an event so it's processed before the class default event processing
	LV_EVENT_MARKED_DELETING LvEventCodeT = 0x10000 // True: the list has marked deleting objects when some of events are marked as deleting
)

//go:type C
type lv_event_dsc_t struct {
	Unused [0]byte
}
type LvEventDscT lv_event_dsc_t

// llgo:type C
type lv_event_list_t struct {
	Unused [0]byte
}
type LvEventListT lv_event_list_t

// Event callback function type
//
// llgo:type C
type LvEventCbT func(e *LvEventT)

//go:linkname LvEventSend C.lv_event_send
func LvEventSend(list *LvEventListT, e *LvEventT, preprocess bool) LvResultT

//go:linkname LvEventAdd C.lv_event_add
func LvEventAdd(list *LvEventListT, cb LvEventCbT, filter LvEventCodeT, user_data unsafe.Pointer) *lv_event_dsc_t

//go:linkname LvEventRemoveDsc C.lv_event_remove_dsc
func LvEventRemoveDsc(list *LvEventListT, dsc *lv_event_dsc_t) bool

//go:linkname LvEventGetCount C.lv_event_get_count
func LvEventGetCount(list *LvEventListT) uint32

//go:linkname LvEventGetDsc C.lv_event_get_dsc
func LvEventGetDsc(list *LvEventListT, index uint32) *lv_event_dsc_t

//go:linkname LvEventDscGetCb C.lv_event_dsc_get_cb
func LvEventDscGetCb(dsc *lv_event_dsc_t) LvEventCbT

//go:linkname LvEventDscGetUserData C.lv_event_dsc_get_user_data
func LvEventDscGetUserData(dsc *lv_event_dsc_t) unsafe.Pointer

//go:linkname LvEventRemove C.lv_event_remove
func LvEventRemove(list *LvEventListT, index uint32) bool

//go:linkname LvEventRemoveAll C.lv_event_remove_all
func LvEventRemoveAll(list *LvEventListT)

//go:linkname LvEventGetTarget C.lv_event_get_target
func LvEventGetTarget(e *LvEventT) unsafe.Pointer

//go:linkname LvEventGetCurrentTarget C.lv_event_get_current_target
func LvEventGetCurrentTarget(e *LvEventT) unsafe.Pointer

//go:linkname LvEventGetCode C.lv_event_get_code
func LvEventGetCode(e *LvEventT) LvEventCodeT

//go:linkname LvEventGetParam C.lv_event_get_param
func LvEventGetParam(e *LvEventT) unsafe.Pointer

//go:linkname LvEventGetUserData C.lv_event_get_user_data
func LvEventGetUserData(e *LvEventT) unsafe.Pointer

//go:linkname LvEventStopBubbling C.lv_event_stop_bubbling
func LvEventStopBubbling(e *LvEventT)

//go:linkname LvEventStopProcessing C.lv_event_stop_processing
func LvEventStopProcessing(e *LvEventT)

//go:linkname LvEventRegisterId C.lv_event_register_id
func LvEventRegisterId() uint32

//go:linkname LvEventCodeGetName C.lv_event_code_get_name
func LvEventCodeGetName(code LvEventCodeT) *c.Char
