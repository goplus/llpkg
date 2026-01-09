package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvLineCreate C.lv_line_create
func LvLineCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvLineSetPoints C.lv_line_set_points
func LvLineSetPoints(obj *core.LvObjT, points *core.LvPointPreciseT, point_num c.Uint32T)

//go:linkname LvLineSetPointsMutable C.lv_line_set_points_mutable
func LvLineSetPointsMutable(obj *core.LvObjT, points *core.LvPointPreciseT, point_num c.Uint32T)

//go:linkname LvLineSetYInvert C.lv_line_set_y_invert
func LvLineSetYInvert(obj *core.LvObjT, en bool)

//go:linkname LvLineGetPoints C.lv_line_get_points
func LvLineGetPoints(obj *core.LvObjT) *core.LvPointPreciseT

//go:linkname LvLineGetPointCount C.lv_line_get_point_count
func LvLineGetPointCount(obj *core.LvObjT) c.Uint32T

//go:linkname LvLineIsPointArrayMutable C.lv_line_is_point_array_mutable
func LvLineIsPointArrayMutable(obj *core.LvObjT) bool

//go:linkname LvLineGetPointsMutable C.lv_line_get_points_mutable
func LvLineGetPointsMutable(obj *core.LvObjT) *core.LvPointPreciseT

//go:linkname LvLineGetYInvert C.lv_line_get_y_invert
func LvLineGetYInvert(obj *core.LvObjT) bool
