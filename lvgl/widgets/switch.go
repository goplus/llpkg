package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvSwitchCreate C.lv_switch_create
func LvSwitchCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvSwitchSetOrientation C.lv_switch_set_orientation
func LvSwitchSetOrientation(obj *core.LvObjT, orientation c.Int32T)

//go:linkname LvSwitchGetOrientation C.lv_switch_get_orientation
func LvSwitchGetOrientation(obj *core.LvObjT) c.Int32T
