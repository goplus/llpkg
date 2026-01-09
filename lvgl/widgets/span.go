package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Span types
type LvSpanT struct{}
type LvSpanCoordsT struct {
	Heading  core.LvAreaT
	Middle   core.LvAreaT
	Trailing core.LvAreaT
}

//go:linkname LvSpanStackInit C.lv_span_stack_init
func LvSpanStackInit()

//go:linkname LvSpanStackDeinit C.lv_span_stack_deinit
func LvSpanStackDeinit()

//go:linkname LvSpangroupCreate C.lv_spangroup_create
func LvSpangroupCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvSpangroupAddSpan C.lv_spangroup_add_span
func LvSpangroupAddSpan(obj *core.LvObjT) *LvSpanT

//go:linkname LvSpangroupDeleteSpan C.lv_spangroup_delete_span
func LvSpangroupDeleteSpan(obj *core.LvObjT, span *LvSpanT)

//go:linkname LvSpanSetText C.lv_span_set_text
func LvSpanSetText(span *LvSpanT, text *c.Char)

//go:linkname LvSpanSetTextStatic C.lv_span_set_text_static
func LvSpanSetTextStatic(span *LvSpanT, text *c.Char)

//go:linkname LvSpangroupSetSpanText C.lv_spangroup_set_span_text
func LvSpangroupSetSpanText(obj *core.LvObjT, span *LvSpanT, text *c.Char)

//go:linkname LvSpangroupSetSpanTextStatic C.lv_spangroup_set_span_text_static
func LvSpangroupSetSpanTextStatic(obj *core.LvObjT, span *LvSpanT, text *c.Char)

//go:linkname LvSpangroupSetSpanStyle C.lv_spangroup_set_span_style
func LvSpangroupSetSpanStyle(obj *core.LvObjT, span *LvSpanT, style *core.LvStyleT)

//go:linkname LvSpangroupSetAlign C.lv_spangroup_set_align
func LvSpangroupSetAlign(obj *core.LvObjT, align c.Int32T)

//go:linkname LvSpangroupSetOverflow C.lv_spangroup_set_overflow
func LvSpangroupSetOverflow(obj *core.LvObjT, overflow c.Int32T)

//go:linkname LvSpangroupSetIndent C.lv_spangroup_set_indent
func LvSpangroupSetIndent(obj *core.LvObjT, indent c.Int32T)

//go:linkname LvSpangroupSetMode C.lv_spangroup_set_mode
func LvSpangroupSetMode(obj *core.LvObjT, mode c.Int32T)

//go:linkname LvSpangroupSetMaxLines C.lv_spangroup_set_max_lines
func LvSpangroupSetMaxLines(obj *core.LvObjT, lines c.Int32T)

//go:linkname LvSpanGetStyle C.lv_span_get_style
func LvSpanGetStyle(span *LvSpanT) *core.LvStyleT

//go:linkname LvSpanGetText C.lv_span_get_text
func LvSpanGetText(span *LvSpanT) *c.Char

//go:linkname LvSpangroupGetChild C.lv_spangroup_get_child
func LvSpangroupGetChild(obj *core.LvObjT, id c.Int32T) *LvSpanT

//go:linkname LvSpangroupGetSpanCount C.lv_spangroup_get_span_count
func LvSpangroupGetSpanCount(obj *core.LvObjT) c.Uint32T

//go:linkname LvSpangroupGetAlign C.lv_spangroup_get_align
func LvSpangroupGetAlign(obj *core.LvObjT) c.Int32T

//go:linkname LvSpangroupGetOverflow C.lv_spangroup_get_overflow
func LvSpangroupGetOverflow(obj *core.LvObjT) c.Int32T

//go:linkname LvSpangroupGetIndent C.lv_spangroup_get_indent
func LvSpangroupGetIndent(obj *core.LvObjT) c.Int32T

//go:linkname LvSpangroupGetMode C.lv_spangroup_get_mode
func LvSpangroupGetMode(obj *core.LvObjT) c.Int32T

//go:linkname LvSpangroupGetMaxLines C.lv_spangroup_get_max_lines
func LvSpangroupGetMaxLines(obj *core.LvObjT) c.Int32T

//go:linkname LvSpangroupGetMaxLineHeight C.lv_spangroup_get_max_line_height
func LvSpangroupGetMaxLineHeight(obj *core.LvObjT) c.Int32T

//go:linkname LvSpangroupGetExpandWidth C.lv_spangroup_get_expand_width
func LvSpangroupGetExpandWidth(obj *core.LvObjT, maxWidth c.Uint32T) c.Uint32T

//go:linkname LvSpangroupGetExpandHeight C.lv_spangroup_get_expand_height
func LvSpangroupGetExpandHeight(obj *core.LvObjT, width c.Int32T) c.Int32T

//go:linkname LvSpangroupGetSpanCoords C.lv_spangroup_get_span_coords
func LvSpangroupGetSpanCoords(obj *core.LvObjT, span *LvSpanT) LvSpanCoordsT

//go:linkname LvSpangroupGetSpanByPoint C.lv_spangroup_get_span_by_point
func LvSpangroupGetSpanByPoint(obj *core.LvObjT, point *core.LvPointT) *LvSpanT

//go:linkname LvSpangroupRefresh C.lv_spangroup_refresh
func LvSpangroupRefresh(obj *core.LvObjT)
