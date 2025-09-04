package widgets

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvListCreate C.lv_list_create
func LvListCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvListAddText C.lv_list_add_text
func LvListAddText(list *core.LvObjT, txt *c.Char) *core.LvObjT

//go:linkname LvListAddButton C.lv_list_add_button
func LvListAddButton(list *core.LvObjT, icon unsafe.Pointer, txt *c.Char) *core.LvObjT

//go:linkname LvListGetButtonText C.lv_list_get_button_text
func LvListGetButtonText(list *core.LvObjT, btn *core.LvObjT) *c.Char

//go:linkname LvListSetButtonText C.lv_list_set_button_text
func LvListSetButtonText(list *core.LvObjT, btn *core.LvObjT, txt *c.Char)
