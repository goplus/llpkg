package widgets

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvMsgboxCreate C.lv_msgbox_create
func LvMsgboxCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvMsgboxAddTitle C.lv_msgbox_add_title
func LvMsgboxAddTitle(obj *core.LvObjT, title *c.Char) *core.LvObjT

//go:linkname LvMsgboxAddHeaderButton C.lv_msgbox_add_header_button
func LvMsgboxAddHeaderButton(obj *core.LvObjT, icon unsafe.Pointer) *core.LvObjT

//go:linkname LvMsgboxAddText C.lv_msgbox_add_text
func LvMsgboxAddText(obj *core.LvObjT, text *c.Char) *core.LvObjT

//go:linkname LvMsgboxAddFooterButton C.lv_msgbox_add_footer_button
func LvMsgboxAddFooterButton(obj *core.LvObjT, text *c.Char) *core.LvObjT

//go:linkname LvMsgboxAddCloseButton C.lv_msgbox_add_close_button
func LvMsgboxAddCloseButton(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMsgboxGetHeader C.lv_msgbox_get_header
func LvMsgboxGetHeader(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMsgboxGetFooter C.lv_msgbox_get_footer
func LvMsgboxGetFooter(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMsgboxGetContent C.lv_msgbox_get_content
func LvMsgboxGetContent(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMsgboxGetTitle C.lv_msgbox_get_title
func LvMsgboxGetTitle(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMsgboxClose C.lv_msgbox_close
func LvMsgboxClose(mbox *core.LvObjT)

//go:linkname LvMsgboxCloseAsync C.lv_msgbox_close_async
func LvMsgboxCloseAsync(mbox *core.LvObjT)
