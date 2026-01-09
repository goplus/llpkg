package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvCheckboxCreate C.lv_checkbox_create
func LvCheckboxCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvCheckboxSetText C.lv_checkbox_set_text
func LvCheckboxSetText(obj *core.LvObjT, txt *c.Char)

//go:linkname LvCheckboxSetTextStatic C.lv_checkbox_set_text_static
func LvCheckboxSetTextStatic(obj *core.LvObjT, txt *c.Char)

//go:linkname LvCheckboxGetText C.lv_checkbox_get_text
func LvCheckboxGetText(obj *core.LvObjT) *c.Char
