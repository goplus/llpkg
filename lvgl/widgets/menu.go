package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvMenuCreate C.lv_menu_create
func LvMenuCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvMenuPageCreate C.lv_menu_page_create
func LvMenuPageCreate(menu *core.LvObjT, title *c.Char) *core.LvObjT

//go:linkname LvMenuContCreate C.lv_menu_cont_create
func LvMenuContCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvMenuSectionCreate C.lv_menu_section_create
func LvMenuSectionCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvMenuSeparatorCreate C.lv_menu_separator_create
func LvMenuSeparatorCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvMenuSetPage C.lv_menu_set_page
func LvMenuSetPage(obj *core.LvObjT, page *core.LvObjT)

//go:linkname LvMenuSetPageTitle C.lv_menu_set_page_title
func LvMenuSetPageTitle(page *core.LvObjT, title *c.Char)

//go:linkname LvMenuSetPageTitleStatic C.lv_menu_set_page_title_static
func LvMenuSetPageTitleStatic(page *core.LvObjT, title *c.Char)

//go:linkname LvMenuSetSidebarPage C.lv_menu_set_sidebar_page
func LvMenuSetSidebarPage(obj *core.LvObjT, page *core.LvObjT)

//go:linkname LvMenuSetModeHeader C.lv_menu_set_mode_header
func LvMenuSetModeHeader(obj *core.LvObjT, mode c.Int32T)

//go:linkname LvMenuSetModeRootBackButton C.lv_menu_set_mode_root_back_button
func LvMenuSetModeRootBackButton(obj *core.LvObjT, mode c.Int32T)

//go:linkname LvMenuSetLoadPageEvent C.lv_menu_set_load_page_event
func LvMenuSetLoadPageEvent(menu *core.LvObjT, obj *core.LvObjT, page *core.LvObjT)

//go:linkname LvMenuGetCurMainPage C.lv_menu_get_cur_main_page
func LvMenuGetCurMainPage(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMenuGetCurSidebarPage C.lv_menu_get_cur_sidebar_page
func LvMenuGetCurSidebarPage(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMenuGetMainHeader C.lv_menu_get_main_header
func LvMenuGetMainHeader(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMenuGetMainHeaderBackButton C.lv_menu_get_main_header_back_button
func LvMenuGetMainHeaderBackButton(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMenuGetSidebarHeader C.lv_menu_get_sidebar_header
func LvMenuGetSidebarHeader(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMenuGetSidebarHeaderBackButton C.lv_menu_get_sidebar_header_back_button
func LvMenuGetSidebarHeaderBackButton(obj *core.LvObjT) *core.LvObjT

//go:linkname LvMenuBackButtonIsRoot C.lv_menu_back_button_is_root
func LvMenuBackButtonIsRoot(menu *core.LvObjT, obj *core.LvObjT) c.Char

//go:linkname LvMenuClearHistory C.lv_menu_clear_history
func LvMenuClearHistory(obj *core.LvObjT)
