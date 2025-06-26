package core

import (
	"unsafe"
	_ "unsafe"
)

/**
 * Possible states of a widget.
 * OR-ed values are possible
 */
const (
	LV_STATE_DEFAULT   = 0x0000
	LV_STATE_CHECKED   = 0x0001
	LV_STATE_FOCUSED   = 0x0002
	LV_STATE_FOCUS_KEY = 0x0004
	LV_STATE_EDITED    = 0x0008
	LV_STATE_HOVERED   = 0x0010
	LV_STATE_PRESSED   = 0x0020
	LV_STATE_SCROLLED  = 0x0040
	LV_STATE_DISABLED  = 0x0080
	LV_STATE_USER_1    = 0x1000
	LV_STATE_USER_2    = 0x2000
	LV_STATE_USER_3    = 0x4000
	LV_STATE_USER_4    = 0x8000

	LV_STATE_ANY = 0xFFFF /**< Special value can be used in some functions to target all states*/
)

/**
 * The possible parts of widgets.
 * The parts can be considered as the internal building block of the widgets.
 * E.g. slider = background + indicator + knob
 * Not all parts are used by every widget
 */

const (
	LV_PART_MAIN      = 0x000000 /**< A background like rectangle*/
	LV_PART_SCROLLBAR = 0x010000 /**< The scrollbar(s)*/
	LV_PART_INDICATOR = 0x020000 /**< Indicator, e.g. for slider, bar, switch, or the tick box of the checkbox*/
	LV_PART_KNOB      = 0x030000 /**< Like handle to grab to adjust the value*/
	LV_PART_SELECTED  = 0x040000 /**< Indicate the currently selected option or section*/
	LV_PART_ITEMS     = 0x050000 /**< Used if the widget has multiple similar elements (e.g. table cells)*/
	LV_PART_CURSOR    = 0x060000 /**< Mark a specific place e.g. for text area's cursor or on a chart*/

	LV_PART_CUSTOM_FIRST = 0x080000 /**< Extension point for custom widgets*/

	LV_PART_ANY = 0x0F0000 /**< Special value can be used in some functions to target all parts*/
)

/**
 * On/Off features controlling the object's behavior.
 * OR-ed values are possible
 *
 * Note: update obj flags corresponding properties below
 * whenever add/remove flags or change bit definition of flags.
 */
type LvObjFlagT uint32

const (
	LV_OBJ_FLAG_HIDDEN                LvObjFlagT = (1 << 0) /**< Make the object hidden. (Like it wasn't there at all)*/
	LV_OBJ_FLAG_CLICKABLE             LvObjFlagT = (1 << 1) /**< Make the object clickable by the input devices*/
	LV_OBJ_FLAG_CLICK_FOCUSABLE       LvObjFlagT = (1 << 2) /**< Add focused state to the object when clicked*/
	LV_OBJ_FLAG_CHECKABLE             LvObjFlagT = (1 << 3) /**< Toggle checked state when the object is clicked*/
	LV_OBJ_FLAG_SCROLLABLE            LvObjFlagT = (1 << 4) /**< Make the object scrollable*/
	LV_OBJ_FLAG_SCROLL_ELASTIC        LvObjFlagT = (1 << 5) /**< Allow scrolling inside but with slower speed*/
	LV_OBJ_FLAG_SCROLL_MOMENTUM       LvObjFlagT = (1 << 6) /**< Make the object scroll further when "thrown"*/
	LV_OBJ_FLAG_SCROLL_ONE            LvObjFlagT = (1 << 7) /**< Allow scrolling only one snappable children*/
	LV_OBJ_FLAG_SCROLL_CHAIN_HOR      LvObjFlagT = (1 << 8) /**< Allow propagating the horizontal scroll to a parent*/
	LV_OBJ_FLAG_SCROLL_CHAIN_VER      LvObjFlagT = (1 << 9) /**< Allow propagating the vertical scroll to a parent*/
	LV_OBJ_FLAG_SCROLL_CHAIN          LvObjFlagT = (LV_OBJ_FLAG_SCROLL_CHAIN_HOR | LV_OBJ_FLAG_SCROLL_CHAIN_VER)
	LV_OBJ_FLAG_SCROLL_ON_FOCUS       LvObjFlagT = (1 << 10) /**< Automatically scroll object to make it visible when focused*/
	LV_OBJ_FLAG_SCROLL_WITH_ARROW     LvObjFlagT = (1 << 11) /**< Allow scrolling the focused object with arrow keys*/
	LV_OBJ_FLAG_SNAPPABLE             LvObjFlagT = (1 << 12) /**< If scroll snap is enabled on the parent it can snap to this object*/
	LV_OBJ_FLAG_PRESS_LOCK            LvObjFlagT = (1 << 13) /**< Keep the object pressed even if the press slid from the object*/
	LV_OBJ_FLAG_EVENT_BUBBLE          LvObjFlagT = (1 << 14) /**< Propagate the events to the parent too*/
	LV_OBJ_FLAG_GESTURE_BUBBLE        LvObjFlagT = (1 << 15) /**< Propagate the gestures to the parent*/
	LV_OBJ_FLAG_ADV_HITTEST           LvObjFlagT = (1 << 16) /**< Allow performing more accurate hit (click) test. E.g. consider rounded corners.*/
	LV_OBJ_FLAG_IGNORE_LAYOUT         LvObjFlagT = (1 << 17) /**< Make the object not positioned by the layouts*/
	LV_OBJ_FLAG_FLOATING              LvObjFlagT = (1 << 18) /**< Do not scroll the object when the parent scrolls and ignore layout*/
	LV_OBJ_FLAG_SEND_DRAW_TASK_EVENTS LvObjFlagT = (1 << 19) /**< Send `LV_EVENT_DRAW_TASK_ADDED` events*/
	LV_OBJ_FLAG_OVERFLOW_VISIBLE      LvObjFlagT = (1 << 20) /**< Do not clip the children to the parent's ext draw size*/

	LV_OBJ_FLAG_FLEX_IN_NEW_TRACK LvObjFlagT = (1 << 21) /**< Start a new flex track on this item*/

	LV_OBJ_FLAG_LAYOUT_1 LvObjFlagT = (1 << 23) /**< Custom flag, free to use by layouts*/
	LV_OBJ_FLAG_LAYOUT_2 LvObjFlagT = (1 << 24) /**< Custom flag, free to use by layouts*/

	LV_OBJ_FLAG_WIDGET_1 LvObjFlagT = (1 << 25) /**< Custom flag, free to use by widget*/
	LV_OBJ_FLAG_WIDGET_2 LvObjFlagT = (1 << 26) /**< Custom flag, free to use by widget*/
	LV_OBJ_FLAG_USER_1   LvObjFlagT = (1 << 27) /**< Custom flag, free to use by user*/
	LV_OBJ_FLAG_USER_2   LvObjFlagT = (1 << 28) /**< Custom flag, free to use by user*/
	LV_OBJ_FLAG_USER_3   LvObjFlagT = (1 << 29) /**< Custom flag, free to use by user*/
	LV_OBJ_FLAG_USER_4   LvObjFlagT = (1 << 30) /**< Custom flag, free to use by user*/
)

//go:linkname LvObjCreate C.lv_obj_create
func LvObjCreate(parent *LvObjT) *LvObjT

//go:linkname LvObjAddFlag C.lv_obj_add_flag
func LvObjAddFlag(obj *LvObjT, f LvObjFlagT)

//go:linkname LvObjRemoveFlag C.lv_obj_remove_flag
func LvObjRemoveFlag(obj *LvObjT, f LvObjFlagT)

//go:linkname LvObjUpdateFlag C.lv_obj_update_flag
func LvObjUpdateFlag(obj *LvObjT, f LvObjFlagT, v bool)

//go:linkname LvObjAddState C.lv_obj_add_state
func LvObjAddState(obj *LvObjT, state LvStateT)

//go:linkname LvObjRemoveState C.lv_obj_remove_state
func LvObjRemoveState(obj *LvObjT, state LvStateT)

//go:linkname LvObjSetState C.lv_obj_set_state
func LvObjSetState(obj *LvObjT, state LvStateT, v bool)

//go:linkname LvObjSetUserData C.lv_obj_set_user_data
func LvObjSetUserData(obj *LvObjT, user_data unsafe.Pointer)

//go:linkname LvObjHasFlag C.lv_obj_has_flag
func LvObjHasFlag(obj *LvObjT, f LvObjFlagT) bool

//go:linkname LvObjHasFlagAny C.lv_obj_has_flag_any
func LvObjHasFlagAny(obj *LvObjT, f LvObjFlagT) bool

//go:linkname LvObjGetState C.lv_obj_get_state
func LvObjGetState(obj *LvObjT) LvStateT

//go:linkname LvObjHasState C.lv_obj_has_state
func LvObjHasState(obj *LvObjT, state LvStateT) bool

//go:linkname LvObjGetGroup C.lv_obj_get_group
func LvObjGetGroup(obj *LvObjT) *LvGroupT

//go:linkname LvObjGetUserData C.lv_obj_get_user_data
func LvObjGetUserData(obj *LvObjT) unsafe.Pointer

//go:linkname LvObjAllocateSpecAttr C.lv_obj_allocate_spec_attr
func LvObjAllocateSpecAttr(obj *LvObjT)

//go:linkname LvObjCheckType C.lv_obj_check_type
func LvObjCheckType(obj *LvObjT, class_p *LvObjClassT) bool

//go:linkname LvObjHasClass C.lv_obj_has_class
func LvObjHasClass(obj *LvObjT, class_p *LvObjClassT) bool

//go:linkname LvObjGetClass C.lv_obj_get_class
func LvObjGetClass(obj *LvObjT) *LvObjClassT

//go:linkname LvObjIsValid C.lv_obj_is_valid
func LvObjIsValid(obj *LvObjT) bool

//go:linkname LvObjNullOnDelete C.lv_obj_null_on_delete
func LvObjNullOnDelete(obj_ptr **LvObjT)
