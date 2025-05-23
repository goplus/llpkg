package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Constants
const (
	LV_LABEL_DOT_NUM                   = 3
	LV_LABEL_POS_LAST           uint32 = 0xFFFF
	LV_LABEL_TEXT_SELECTION_OFF        = core.LV_DRAW_LABEL_NO_TXT_SEL
	LV_LABEL_DEFAULT_TEXT              = "Text" // or "" if LV_WIDGETS_HAS_DEFAULT_VALUE is not defined
)

// Label long mode behaviors
type LvLabelLongModeT uint8

const (
	LV_LABEL_LONG_MODE_WRAP            LvLabelLongModeT = iota // Keep the object width, wrap lines longer than object width and expand the object height
	LV_LABEL_LONG_MODE_DOTS                                    // Keep the size and write dots at the end if the text is too long
	LV_LABEL_LONG_MODE_SCROLL                                  // Keep the size and roll the text back and forth
	LV_LABEL_LONG_MODE_SCROLL_CIRCULAR                         // Keep the size and roll the text circularly
	LV_LABEL_LONG_MODE_CLIP                                    // Keep the size and clip the text out of it
)

//go:linkname LvLabelCreate C.lv_label_create
func LvLabelCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvLabelSetText C.lv_label_set_text
func LvLabelSetText(obj *core.LvObjT, text *c.Char)

//go:linkname LvLabelSetTextFmt C.lv_label_set_text_fmt
func LvLabelSetTextFmt(obj *core.LvObjT, fmt *c.Char, args ...interface{})

//go:linkname LvLabelSetTextStatic C.lv_label_set_text_static
func LvLabelSetTextStatic(obj *core.LvObjT, text *c.Char)

//go:linkname LvLabelSetLongMode C.lv_label_set_long_mode
func LvLabelSetLongMode(obj *core.LvObjT, long_mode LvLabelLongModeT)

//go:linkname LvLabelSetTextSelectionStart C.lv_label_set_text_selection_start
func LvLabelSetTextSelectionStart(obj *core.LvObjT, index c.Uint32T)

//go:linkname LvLabelSetTextSelectionEnd C.lv_label_set_text_selection_end
func LvLabelSetTextSelectionEnd(obj *core.LvObjT, index c.Uint32T)

//go:linkname LvLabelSetRecolor C.lv_label_set_recolor
func LvLabelSetRecolor(obj *core.LvObjT, en bool)

//go:linkname LvLabelGetText C.lv_label_get_text
func LvLabelGetText(obj *core.LvObjT) *c.Char

//go:linkname LvLabelGetLongMode C.lv_label_get_long_mode
func LvLabelGetLongMode(obj *core.LvObjT) LvLabelLongModeT

//go:linkname LvLabelGetLetterPos C.lv_label_get_letter_pos
func LvLabelGetLetterPos(obj *core.LvObjT, char_id c.Uint32T, pos *core.LvPointT)

//go:linkname LvLabelGetLetterOn C.lv_label_get_letter_on
func LvLabelGetLetterOn(obj *core.LvObjT, pos_in *core.LvPointT, bidi bool) c.Uint32T

//go:linkname LvLabelIsCharUnderPos C.lv_label_is_char_under_pos
func LvLabelIsCharUnderPos(obj *core.LvObjT, pos *core.LvPointT) bool

//go:linkname LvLabelGetTextSelectionStart C.lv_label_get_text_selection_start
func LvLabelGetTextSelectionStart(obj *core.LvObjT) c.Uint32T

//go:linkname LvLabelGetTextSelectionEnd C.lv_label_get_text_selection_end
func LvLabelGetTextSelectionEnd(obj *core.LvObjT) c.Uint32T

//go:linkname LvLabelGetRecolor C.lv_label_get_recolor
func LvLabelGetRecolor(obj *core.LvObjT) bool

//go:linkname LvLabelInsText C.lv_label_ins_text
func LvLabelInsText(obj *core.LvObjT, pos c.Uint32T, txt *c.Char)

//go:linkname LvLabelCutText C.lv_label_cut_text
func LvLabelCutText(obj *core.LvObjT, pos c.Uint32T, cnt c.Uint32T)
