package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvRollerCreate C.lv_roller_create
func LvRollerCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvRollerSetOptions C.lv_roller_set_options
func LvRollerSetOptions(obj *core.LvObjT, options *c.Char, mode c.Int32T)

//go:linkname LvRollerSetSelected C.lv_roller_set_selected
func LvRollerSetSelected(obj *core.LvObjT, selOpt c.Uint32T, anim c.Int32T)

//go:linkname LvRollerSetSelectedStr C.lv_roller_set_selected_str
func LvRollerSetSelectedStr(obj *core.LvObjT, selOpt *c.Char, anim c.Int32T) c.Char

//go:linkname LvRollerSetVisibleRowCount C.lv_roller_set_visible_row_count
func LvRollerSetVisibleRowCount(obj *core.LvObjT, rowCnt c.Uint32T)

//go:linkname LvRollerGetSelected C.lv_roller_get_selected
func LvRollerGetSelected(obj *core.LvObjT) c.Uint32T

//go:linkname LvRollerGetSelectedStr C.lv_roller_get_selected_str
func LvRollerGetSelectedStr(obj *core.LvObjT, buf *c.Char, bufSize c.Uint32T)

//go:linkname LvRollerGetOptions C.lv_roller_get_options
func LvRollerGetOptions(obj *core.LvObjT) *c.Char

//go:linkname LvRollerGetOptionCount C.lv_roller_get_option_count
func LvRollerGetOptionCount(obj *core.LvObjT) c.Uint32T
