package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

type lv_btnmatrix_ctrl_t c.Uint

// Button matrix control flags
const (
	LV_BUTTONMATRIX_CTRL_NONE       lv_btnmatrix_ctrl_t = 0x0000 /**< No extra control, use the default settings*/
	LV_BUTTONMATRIX_CTRL_WIDTH_1    lv_btnmatrix_ctrl_t = 0x0001 /**< Set the width to 1 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_2    lv_btnmatrix_ctrl_t = 0x0002 /**< Set the width to 2 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_3    lv_btnmatrix_ctrl_t = 0x0003 /**< Set the width to 3 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_4    lv_btnmatrix_ctrl_t = 0x0004 /**< Set the width to 4 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_5    lv_btnmatrix_ctrl_t = 0x0005 /**< Set the width to 5 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_6    lv_btnmatrix_ctrl_t = 0x0006 /**< Set the width to 6 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_7    lv_btnmatrix_ctrl_t = 0x0007 /**< Set the width to 7 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_8    lv_btnmatrix_ctrl_t = 0x0008 /**< Set the width to 8 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_9    lv_btnmatrix_ctrl_t = 0x0009 /**< Set the width to 9 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_10   lv_btnmatrix_ctrl_t = 0x000A /**< Set the width to 10 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_11   lv_btnmatrix_ctrl_t = 0x000B /**< Set the width to 11 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_12   lv_btnmatrix_ctrl_t = 0x000C /**< Set the width to 12 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_13   lv_btnmatrix_ctrl_t = 0x000D /**< Set the width to 13 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_14   lv_btnmatrix_ctrl_t = 0x000E /**< Set the width to 14 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_WIDTH_15   lv_btnmatrix_ctrl_t = 0x000F /**< Set the width to 15 relative to the other buttons in the same row */
	LV_BUTTONMATRIX_CTRL_HIDDEN     lv_btnmatrix_ctrl_t = 0x0010 /**< Hides button; it continues to hold its space in layout. */
	LV_BUTTONMATRIX_CTRL_NO_REPEAT  lv_btnmatrix_ctrl_t = 0x0020 /**< Do not emit LV_EVENT_LONG_PRESSED_REPEAT events while button is long-pressed. */
	LV_BUTTONMATRIX_CTRL_DISABLED   lv_btnmatrix_ctrl_t = 0x0040 /**< Disables button like LV_STATE_DISABLED on normal Widgets. */
	LV_BUTTONMATRIX_CTRL_CHECKABLE  lv_btnmatrix_ctrl_t = 0x0080 /**< Enable toggling of LV_STATE_CHECKED when clicked. */
	LV_BUTTONMATRIX_CTRL_CHECKED    lv_btnmatrix_ctrl_t = 0x0100 /**< Make the button checked. It will use the :cpp:enumerator:`LV_STATE_CHECHKED` styles. */
	LV_BUTTONMATRIX_CTRL_CLICK_TRIG lv_btnmatrix_ctrl_t = 0x0200 /**< 1: Enables sending LV_EVENT_VALUE_CHANGE on CLICK, 0: sends LV_EVENT_VALUE_CHANGE on PRESS. */
	LV_BUTTONMATRIX_CTRL_POPOVER    lv_btnmatrix_ctrl_t = 0x0400 /**< Show button text in a pop-over while being pressed. */
	LV_BUTTONMATRIX_CTRL_RECOLOR    lv_btnmatrix_ctrl_t = 0x0800 /**< Enable text recoloring with `#color` */
	LV_BUTTONMATRIX_CTRL_RESERVED_1 lv_btnmatrix_ctrl_t = 0x1000 /**< Reserved for later use */
	LV_BUTTONMATRIX_CTRL_RESERVED_2 lv_btnmatrix_ctrl_t = 0x2000 /**< Reserved for later use */
	LV_BUTTONMATRIX_CTRL_CUSTOM_1   lv_btnmatrix_ctrl_t = 0x4000 /**< Custom free-to-use flag */
	LV_BUTTONMATRIX_CTRL_CUSTOM_2   lv_btnmatrix_ctrl_t = 0x8000 /**< Custom free-to-use flag */
)

type LvButtonmatrixCtrlT lv_btnmatrix_ctrl_t

//go:linkname LvButtonmatrixCreate C.lv_buttonmatrix_create
func LvButtonmatrixCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvButtonmatrixSetMap C.lv_buttonmatrix_set_map
func LvButtonmatrixSetMap(obj *core.LvObjT, map_ **c.Char)

//go:linkname LvButtonmatrixSetCtrlMap C.lv_buttonmatrix_set_ctrl_map
func LvButtonmatrixSetCtrlMap(obj *core.LvObjT, ctrl_map *lv_btnmatrix_ctrl_t)

//go:linkname LvButtonmatrixSetSelectedBtn C.lv_buttonmatrix_set_selected_btn
func LvButtonmatrixSetSelectedBtn(obj *core.LvObjT, btn_id c.Uint16T)

//go:linkname LvButtonmatrixSetBtnCtrl C.lv_buttonmatrix_set_btn_ctrl
func LvButtonmatrixSetBtnCtrl(obj *core.LvObjT, btn_id c.Uint16T, ctrl lv_btnmatrix_ctrl_t)

//go:linkname LvButtonmatrixClearBtnCtrl C.lv_buttonmatrix_clear_btn_ctrl
func LvButtonmatrixClearBtnCtrl(obj *core.LvObjT, btn_id c.Uint16T, ctrl lv_btnmatrix_ctrl_t)

//go:linkname LvButtonmatrixGetSelectedBtn C.lv_buttonmatrix_get_selected_btn
func LvButtonmatrixGetSelectedBtn(obj *core.LvObjT) c.Uint16T

//go:linkname LvButtonmatrixGetBtnText C.lv_buttonmatrix_get_btn_text
func LvButtonmatrixGetBtnText(obj *core.LvObjT, btn_id c.Uint16T) *c.Char

//go:linkname LvButtonmatrixHasBtnCtrl C.lv_buttonmatrix_has_btn_ctrl
func LvButtonmatrixHasBtnCtrl(obj *core.LvObjT, btn_id c.Uint16T, ctrl lv_btnmatrix_ctrl_t) bool
