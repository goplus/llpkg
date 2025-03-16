package widgets

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Constants
const LV_DROPDOWN_POS_LAST c.Uint32T = 0xFFFF

//go:linkname LvDropdownCreate C.lv_dropdown_create
func LvDropdownCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvDropdownSetText C.lv_dropdown_set_text
func LvDropdownSetText(obj *core.LvObjT, txt *c.Char)

//go:linkname LvDropdownSetOptions C.lv_dropdown_set_options
func LvDropdownSetOptions(obj *core.LvObjT, options *c.Char)

//go:linkname LvDropdownSetOptionsStatic C.lv_dropdown_set_options_static
func LvDropdownSetOptionsStatic(obj *core.LvObjT, options *c.Char)

//go:linkname LvDropdownAddOption C.lv_dropdown_add_option
func LvDropdownAddOption(obj *core.LvObjT, option *c.Char, pos c.Uint32T)

//go:linkname LvDropdownClearOptions C.lv_dropdown_clear_options
func LvDropdownClearOptions(obj *core.LvObjT)

//go:linkname LvDropdownSetSelected C.lv_dropdown_set_selected
func LvDropdownSetSelected(obj *core.LvObjT, sel_opt c.Uint32T, anim core.LvAnimEnableT)

//go:linkname LvDropdownSetDir C.lv_dropdown_set_dir
func LvDropdownSetDir(obj *core.LvObjT, dir core.LvDirT)

//go:linkname LvDropdownSetSymbol C.lv_dropdown_set_symbol
func LvDropdownSetSymbol(obj *core.LvObjT, symbol unsafe.Pointer)

//go:linkname LvDropdownSetSelectedHighlight C.lv_dropdown_set_selected_highlight
func LvDropdownSetSelectedHighlight(obj *core.LvObjT, en bool)

//go:linkname LvDropdownGetList C.lv_dropdown_get_list
func LvDropdownGetList(obj *core.LvObjT) *core.LvObjT

//go:linkname LvDropdownGetText C.lv_dropdown_get_text
func LvDropdownGetText(obj *core.LvObjT) *c.Char

//go:linkname LvDropdownGetOptions C.lv_dropdown_get_options
func LvDropdownGetOptions(obj *core.LvObjT) *c.Char

//go:linkname LvDropdownGetSelected C.lv_dropdown_get_selected
func LvDropdownGetSelected(obj *core.LvObjT) c.Uint32T

//go:linkname LvDropdownGetOptionCount C.lv_dropdown_get_option_count
func LvDropdownGetOptionCount(obj *core.LvObjT) c.Uint32T

//go:linkname LvDropdownGetSelectedStr C.lv_dropdown_get_selected_str
func LvDropdownGetSelectedStr(obj *core.LvObjT, buf *c.Char, buf_size c.Uint32T)

//go:linkname LvDropdownGetOptionIndex C.lv_dropdown_get_option_index
func LvDropdownGetOptionIndex(obj *core.LvObjT, option *c.Char) c.Int32T

//go:linkname LvDropdownGetSymbol C.lv_dropdown_get_symbol
func LvDropdownGetSymbol(obj *core.LvObjT) *c.Char

//go:linkname LvDropdownGetSelectedHighlight C.lv_dropdown_get_selected_highlight
func LvDropdownGetSelectedHighlight(obj *core.LvObjT) bool

//go:linkname LvDropdownGetDir C.lv_dropdown_get_dir
func LvDropdownGetDir(obj *core.LvObjT) core.LvDirT

//go:linkname LvDropdownOpen C.lv_dropdown_open
func LvDropdownOpen(dropdown_obj *core.LvObjT)

//go:linkname LvDropdownClose C.lv_dropdown_close
func LvDropdownClose(obj *core.LvObjT)

//go:linkname LvDropdownIsOpen C.lv_dropdown_is_open
func LvDropdownIsOpen(obj *core.LvObjT) bool
