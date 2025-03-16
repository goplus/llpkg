package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvTextareaCreate C.lv_textarea_create
func LvTextareaCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvTextareaAddChar C.lv_textarea_add_char
func LvTextareaAddChar(obj *core.LvObjT, c c.Uint32T)

//go:linkname LvTextareaAddText C.lv_textarea_add_text
func LvTextareaAddText(obj *core.LvObjT, txt *c.Char)

//go:linkname LvTextareaDeleteChar C.lv_textarea_delete_char
func LvTextareaDeleteChar(obj *core.LvObjT)

//go:linkname LvTextareaDeleteCharForward C.lv_textarea_delete_char_forward
func LvTextareaDeleteCharForward(obj *core.LvObjT)

//go:linkname LvTextareaSetText C.lv_textarea_set_text
func LvTextareaSetText(obj *core.LvObjT, txt *c.Char)

//go:linkname LvTextareaSetPlaceholderText C.lv_textarea_set_placeholder_text
func LvTextareaSetPlaceholderText(obj *core.LvObjT, txt *c.Char)

//go:linkname LvTextareaSetCursorPos C.lv_textarea_set_cursor_pos
func LvTextareaSetCursorPos(obj *core.LvObjT, pos c.Int32T)

//go:linkname LvTextareaSetCursorClickPos C.lv_textarea_set_cursor_click_pos
func LvTextareaSetCursorClickPos(obj *core.LvObjT, en c.Char)

//go:linkname LvTextareaSetPasswordMode C.lv_textarea_set_password_mode
func LvTextareaSetPasswordMode(obj *core.LvObjT, en c.Char)

//go:linkname LvTextareaSetPasswordBullet C.lv_textarea_set_password_bullet
func LvTextareaSetPasswordBullet(obj *core.LvObjT, bullet *c.Char)

//go:linkname LvTextareaSetOneLine C.lv_textarea_set_one_line
func LvTextareaSetOneLine(obj *core.LvObjT, en c.Char)

//go:linkname LvTextareaSetAcceptedChars C.lv_textarea_set_accepted_chars
func LvTextareaSetAcceptedChars(obj *core.LvObjT, list *c.Char)

//go:linkname LvTextareaSetMaxLength C.lv_textarea_set_max_length
func LvTextareaSetMaxLength(obj *core.LvObjT, num c.Uint32T)

//go:linkname LvTextareaSetInsertReplace C.lv_textarea_set_insert_replace
func LvTextareaSetInsertReplace(obj *core.LvObjT, txt *c.Char)

//go:linkname LvTextareaSetTextSelection C.lv_textarea_set_text_selection
func LvTextareaSetTextSelection(obj *core.LvObjT, en c.Char)

//go:linkname LvTextareaSetPasswordShowTime C.lv_textarea_set_password_show_time
func LvTextareaSetPasswordShowTime(obj *core.LvObjT, time c.Uint32T)

//go:linkname LvTextareaSetAlign C.lv_textarea_set_align
func LvTextareaSetAlign(obj *core.LvObjT, align c.Int32T)

//go:linkname LvTextareaGetText C.lv_textarea_get_text
func LvTextareaGetText(obj *core.LvObjT) *c.Char

//go:linkname LvTextareaGetPlaceholderText C.lv_textarea_get_placeholder_text
func LvTextareaGetPlaceholderText(obj *core.LvObjT) *c.Char

//go:linkname LvTextareaGetLabel C.lv_textarea_get_label
func LvTextareaGetLabel(obj *core.LvObjT) *core.LvObjT

//go:linkname LvTextareaGetCursorPos C.lv_textarea_get_cursor_pos
func LvTextareaGetCursorPos(obj *core.LvObjT) c.Uint32T

//go:linkname LvTextareaGetCursorClickPos C.lv_textarea_get_cursor_click_pos
func LvTextareaGetCursorClickPos(obj *core.LvObjT) c.Char

//go:linkname LvTextareaGetPasswordMode C.lv_textarea_get_password_mode
func LvTextareaGetPasswordMode(obj *core.LvObjT) c.Char

//go:linkname LvTextareaGetPasswordBullet C.lv_textarea_get_password_bullet
func LvTextareaGetPasswordBullet(obj *core.LvObjT) *c.Char

//go:linkname LvTextareaGetOneLine C.lv_textarea_get_one_line
func LvTextareaGetOneLine(obj *core.LvObjT) c.Char

//go:linkname LvTextareaGetAcceptedChars C.lv_textarea_get_accepted_chars
func LvTextareaGetAcceptedChars(obj *core.LvObjT) *c.Char

//go:linkname LvTextareaGetMaxLength C.lv_textarea_get_max_length
func LvTextareaGetMaxLength(obj *core.LvObjT) c.Uint32T

//go:linkname LvTextareaTextIsSelected C.lv_textarea_text_is_selected
func LvTextareaTextIsSelected(obj *core.LvObjT) c.Char

//go:linkname LvTextareaGetTextSelection C.lv_textarea_get_text_selection
func LvTextareaGetTextSelection(obj *core.LvObjT) c.Char

//go:linkname LvTextareaGetPasswordShowTime C.lv_textarea_get_password_show_time
func LvTextareaGetPasswordShowTime(obj *core.LvObjT) c.Uint32T

//go:linkname LvTextareaGetCurrentChar C.lv_textarea_get_current_char
func LvTextareaGetCurrentChar(obj *core.LvObjT) c.Uint32T

//go:linkname LvTextareaClearSelection C.lv_textarea_clear_selection
func LvTextareaClearSelection(obj *core.LvObjT)

//go:linkname LvTextareaCursorRight C.lv_textarea_cursor_right
func LvTextareaCursorRight(obj *core.LvObjT)

//go:linkname LvTextareaCursorLeft C.lv_textarea_cursor_left
func LvTextareaCursorLeft(obj *core.LvObjT)

//go:linkname LvTextareaCursorDown C.lv_textarea_cursor_down
func LvTextareaCursorDown(obj *core.LvObjT)

//go:linkname LvTextareaCursorUp C.lv_textarea_cursor_up
func LvTextareaCursorUp(obj *core.LvObjT)
