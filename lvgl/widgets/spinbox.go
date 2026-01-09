package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvSpinboxCreate C.lv_spinbox_create
func LvSpinboxCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvSpinboxSetValue C.lv_spinbox_set_value
func LvSpinboxSetValue(obj *core.LvObjT, v c.Int32T)

//go:linkname LvSpinboxSetRollover C.lv_spinbox_set_rollover
func LvSpinboxSetRollover(obj *core.LvObjT, rollover c.Char)

//go:linkname LvSpinboxSetDigitFormat C.lv_spinbox_set_digit_format
func LvSpinboxSetDigitFormat(obj *core.LvObjT, digitCount c.Uint32T, sepPos c.Uint32T)

//go:linkname LvSpinboxSetStep C.lv_spinbox_set_step
func LvSpinboxSetStep(obj *core.LvObjT, step c.Uint32T)

//go:linkname LvSpinboxSetRange C.lv_spinbox_set_range
func LvSpinboxSetRange(obj *core.LvObjT, rangeMin c.Int32T, rangeMax c.Int32T)

//go:linkname LvSpinboxSetCursorPos C.lv_spinbox_set_cursor_pos
func LvSpinboxSetCursorPos(obj *core.LvObjT, pos c.Uint32T)

//go:linkname LvSpinboxSetDigitStepDirection C.lv_spinbox_set_digit_step_direction
func LvSpinboxSetDigitStepDirection(obj *core.LvObjT, direction c.Int32T)

//go:linkname LvSpinboxGetRollover C.lv_spinbox_get_rollover
func LvSpinboxGetRollover(obj *core.LvObjT) c.Char

//go:linkname LvSpinboxGetValue C.lv_spinbox_get_value
func LvSpinboxGetValue(obj *core.LvObjT) c.Int32T

//go:linkname LvSpinboxGetStep C.lv_spinbox_get_step
func LvSpinboxGetStep(obj *core.LvObjT) c.Int32T

//go:linkname LvSpinboxStepNext C.lv_spinbox_step_next
func LvSpinboxStepNext(obj *core.LvObjT)

//go:linkname LvSpinboxStepPrev C.lv_spinbox_step_prev
func LvSpinboxStepPrev(obj *core.LvObjT)

//go:linkname LvSpinboxIncrement C.lv_spinbox_increment
func LvSpinboxIncrement(obj *core.LvObjT)

//go:linkname LvSpinboxDecrement C.lv_spinbox_decrement
func LvSpinboxDecrement(obj *core.LvObjT)
