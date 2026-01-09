package layouts

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

type lv_flex_align_t c.Int

const (
	LV_FLEX_ALIGN_START lv_flex_align_t = iota
	LV_FLEX_ALIGN_END
	LV_FLEX_ALIGN_CENTER
	LV_FLEX_ALIGN_SPACE_EVENLY
	LV_FLEX_ALIGN_SPACE_AROUND
	LV_FLEX_ALIGN_SPACE_BETWEEN
)

// Constants needed for lv_flex_flow_t
const (
	LV_FLEX_COLUMN  = 0x01
	LV_FLEX_WRAP    = 0x02
	LV_FLEX_REVERSE = 0x04
)

type lv_flex_flow_t c.Int

const (
	LV_FLEX_FLOW_ROW                 lv_flex_flow_t = 0x00
	LV_FLEX_FLOW_COLUMN              lv_flex_flow_t = LV_FLEX_COLUMN
	LV_FLEX_FLOW_ROW_WRAP            lv_flex_flow_t = 0x00 | LV_FLEX_WRAP
	LV_FLEX_FLOW_ROW_REVERSE         lv_flex_flow_t = 0x00 | LV_FLEX_REVERSE
	LV_FLEX_FLOW_ROW_WRAP_REVERSE    lv_flex_flow_t = 0x00 | LV_FLEX_WRAP | LV_FLEX_REVERSE
	LV_FLEX_FLOW_COLUMN_WRAP         lv_flex_flow_t = LV_FLEX_COLUMN | LV_FLEX_WRAP
	LV_FLEX_FLOW_COLUMN_REVERSE      lv_flex_flow_t = LV_FLEX_COLUMN | LV_FLEX_REVERSE
	LV_FLEX_FLOW_COLUMN_WRAP_REVERSE lv_flex_flow_t = LV_FLEX_COLUMN | LV_FLEX_WRAP | LV_FLEX_REVERSE
)

//go:linkname LvFlexInit C.lv_flex_init
func LvFlexInit()

//go:linkname LvObjSetFlexFlow C.lv_obj_set_flex_flow
func LvObjSetFlexFlow(obj *core.LvObjT, flow lv_flex_flow_t)

//go:linkname LvObjSetFlexAlign C.lv_obj_set_flex_align
func LvObjSetFlexAlign(obj *core.LvObjT, main_place, cross_place, track_cross_place lv_flex_align_t)

//go:linkname LvObjSetFlexGrow C.lv_obj_set_flex_grow
func LvObjSetFlexGrow(obj *core.LvObjT, grow c.Uint8T)
