package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvSpinnerCreate C.lv_spinner_create
func LvSpinnerCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvSpinnerSetAnimParams C.lv_spinner_set_anim_params
func LvSpinnerSetAnimParams(obj *core.LvObjT, t c.Uint32T, angle c.Uint32T)
