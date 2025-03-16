package widgets

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvWinCreate C.lv_win_create
func LvWinCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvWinAddTitle C.lv_win_add_title
func LvWinAddTitle(win *core.LvObjT, txt *c.Char) *core.LvObjT

//go:linkname LvWinAddButton C.lv_win_add_button
func LvWinAddButton(win *core.LvObjT, icon unsafe.Pointer, btnW c.Int32T) *core.LvObjT

//go:linkname LvWinGetHeader C.lv_win_get_header
func LvWinGetHeader(win *core.LvObjT) *core.LvObjT

//go:linkname LvWinGetContent C.lv_win_get_content
func LvWinGetContent(win *core.LvObjT) *core.LvObjT
