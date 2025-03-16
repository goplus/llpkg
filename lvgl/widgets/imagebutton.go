package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Image button states
type LvImagebuttonStateT uint8

const (
	LV_IMAGEBUTTON_STATE_RELEASED LvImagebuttonStateT = iota
	LV_IMAGEBUTTON_STATE_PRESSED
	LV_IMAGEBUTTON_STATE_DISABLED
	LV_IMAGEBUTTON_STATE_CHECKED_RELEASED
	LV_IMAGEBUTTON_STATE_CHECKED_PRESSED
	LV_IMAGEBUTTON_STATE_CHECKED_DISABLED
	LV_IMAGEBUTTON_STATE_NUM
)

//go:linkname LvImagebuttonCreate C.lv_imagebutton_create
func LvImagebuttonCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvImagebuttonSetSrc C.lv_imagebutton_set_src
func LvImagebuttonSetSrc(imagebutton *core.LvObjT, state LvImagebuttonStateT, src_left *c.Void, src_mid *c.Void, src_right *c.Void)

//go:linkname LvImagebuttonSetState C.lv_imagebutton_set_state
func LvImagebuttonSetState(imagebutton *core.LvObjT, state LvImagebuttonStateT)

//go:linkname LvImagebuttonGetSrcLeft C.lv_imagebutton_get_src_left
func LvImagebuttonGetSrcLeft(imagebutton *core.LvObjT, state LvImagebuttonStateT) *c.Void

//go:linkname LvImagebuttonGetSrcMiddle C.lv_imagebutton_get_src_middle
func LvImagebuttonGetSrcMiddle(imagebutton *core.LvObjT, state LvImagebuttonStateT) *c.Void

//go:linkname LvImagebuttonGetSrcRight C.lv_imagebutton_get_src_right
func LvImagebuttonGetSrcRight(imagebutton *core.LvObjT, state LvImagebuttonStateT) *c.Void
