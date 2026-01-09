package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Keyboard control button flags
const LV_KEYBOARD_CTRL_BUTTON_FLAGS = LV_BUTTONMATRIX_CTRL_NO_REPEAT | LV_BUTTONMATRIX_CTRL_CLICK_TRIG | LV_BUTTONMATRIX_CTRL_CHECKED

// Keyboard mode
type LvKeyboardModeT uint8

const (
	LV_KEYBOARD_MODE_TEXT_LOWER LvKeyboardModeT = iota
	LV_KEYBOARD_MODE_TEXT_UPPER
	LV_KEYBOARD_MODE_SPECIAL
	LV_KEYBOARD_MODE_NUMBER
	LV_KEYBOARD_MODE_USER_1
	LV_KEYBOARD_MODE_USER_2
	LV_KEYBOARD_MODE_USER_3
	LV_KEYBOARD_MODE_USER_4
	LV_KEYBOARD_MODE_TEXT_ARABIC // Only if LV_USE_ARABIC_PERSIAN_CHARS == 1
)

//go:linkname LvKeyboardCreate C.lv_keyboard_create
func LvKeyboardCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvKeyboardSetTextarea C.lv_keyboard_set_textarea
func LvKeyboardSetTextarea(kb *core.LvObjT, ta *core.LvObjT)

//go:linkname LvKeyboardSetMode C.lv_keyboard_set_mode
func LvKeyboardSetMode(kb *core.LvObjT, mode LvKeyboardModeT)

//go:linkname LvKeyboardSetPopovers C.lv_keyboard_set_popovers
func LvKeyboardSetPopovers(kb *core.LvObjT, en bool)

//go:linkname LvKeyboardSetMap C.lv_keyboard_set_map
func LvKeyboardSetMap(kb *core.LvObjT, mode LvKeyboardModeT, map_ **c.Char, ctrl_map *LvButtonmatrixCtrlT)

//go:linkname LvKeyboardGetTextarea C.lv_keyboard_get_textarea
func LvKeyboardGetTextarea(kb *core.LvObjT) *core.LvObjT

//go:linkname LvKeyboardGetMode C.lv_keyboard_get_mode
func LvKeyboardGetMode(kb *core.LvObjT) LvKeyboardModeT

//go:linkname LvKeyboardGetPopovers C.lv_keyboard_get_popovers
func LvKeyboardGetPopovers(obj *core.LvObjT) bool

//go:linkname LvKeyboardGetMapArray C.lv_keyboard_get_map_array
func LvKeyboardGetMapArray(kb *core.LvObjT) **c.Char

//go:linkname LvKeyboardGetSelectedButton C.lv_keyboard_get_selected_button
func LvKeyboardGetSelectedButton(obj *core.LvObjT) c.Uint32T

//go:linkname LvKeyboardGetButtonText C.lv_keyboard_get_button_text
func LvKeyboardGetButtonText(obj *core.LvObjT, btn_id c.Uint32T) *c.Char

//go:linkname LvKeyboardDefEventCb C.lv_keyboard_def_event_cb
func LvKeyboardDefEventCb(e *core.LvEventT)
