package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvTabviewCreate C.lv_tabview_create
func LvTabviewCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvTabviewAddTab C.lv_tabview_add_tab
func LvTabviewAddTab(obj *core.LvObjT, name *c.Char) *core.LvObjT

//go:linkname LvTabviewRenameTab C.lv_tabview_rename_tab
func LvTabviewRenameTab(obj *core.LvObjT, idx c.Uint32T, newName *c.Char)

//go:linkname LvTabviewSetActive C.lv_tabview_set_active
func LvTabviewSetActive(obj *core.LvObjT, idx c.Uint32T, animEn c.Int32T)

//go:linkname LvTabviewSetTabBarPosition C.lv_tabview_set_tab_bar_position
func LvTabviewSetTabBarPosition(obj *core.LvObjT, dir c.Int32T)

//go:linkname LvTabviewSetTabBarSize C.lv_tabview_set_tab_bar_size
func LvTabviewSetTabBarSize(obj *core.LvObjT, size c.Int32T)

//go:linkname LvTabviewGetTabCount C.lv_tabview_get_tab_count
func LvTabviewGetTabCount(obj *core.LvObjT) c.Uint32T

//go:linkname LvTabviewGetTabActive C.lv_tabview_get_tab_active
func LvTabviewGetTabActive(obj *core.LvObjT) c.Uint32T

//go:linkname LvTabviewGetContent C.lv_tabview_get_content
func LvTabviewGetContent(obj *core.LvObjT) *core.LvObjT

//go:linkname LvTabviewGetTabBar C.lv_tabview_get_tab_bar
func LvTabviewGetTabBar(obj *core.LvObjT) *core.LvObjT
