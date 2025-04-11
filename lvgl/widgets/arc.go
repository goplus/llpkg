package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

type lv_arc_mode_t c.Int

const (
	LV_ARC_MODE_NORMAL lv_arc_mode_t = iota
	LV_ARC_MODE_SYMMETRICAL
	LV_ARC_MODE_REVERSE
)

//go:linkname LvArcCreate C.lv_arc_create
func LvArcCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvArcSetStartAngle C.lv_arc_set_start_angle
func LvArcSetStartAngle(obj *core.LvObjT, start lv_value_precise_t)

//go:linkname LvArcSetEndAngle C.lv_arc_set_end_angle
func LvArcSetEndAngle(obj *core.LvObjT, end lv_value_precise_t)

//go:linkname LvArcSetAngles C.lv_arc_set_angles
func LvArcSetAngles(obj *core.LvObjT, start, end lv_value_precise_t)

//go:linkname LvArcSetBgStartAngle C.lv_arc_set_bg_start_angle
func LvArcSetBgStartAngle(obj *core.LvObjT, start lv_value_precise_t)

//go:linkname LvArcSetBgEndAngle C.lv_arc_set_bg_end_angle
func LvArcSetBgEndAngle(obj *core.LvObjT, end lv_value_precise_t)

//go:linkname LvArcSetBgAngles C.lv_arc_set_bg_angles
func LvArcSetBgAngles(obj *core.LvObjT, start, end lv_value_precise_t)

//go:linkname LvArcSetRotation C.lv_arc_set_rotation
func LvArcSetRotation(obj *core.LvObjT, rotation c.Int32T)

//go:linkname LvArcSetMode C.lv_arc_set_mode
func LvArcSetMode(obj *core.LvObjT, mode lv_arc_mode_t)

//go:linkname LvArcSetValue C.lv_arc_set_value
func LvArcSetValue(obj *core.LvObjT, value c.Int32T)

//go:linkname LvArcSetRange C.lv_arc_set_range
func LvArcSetRange(obj *core.LvObjT, min, max c.Int32T)

//go:linkname LvArcSetChangeRate C.lv_arc_set_change_rate
func LvArcSetChangeRate(obj *core.LvObjT, rate c.Uint32T)

//go:linkname LvArcSetKnobOffset C.lv_arc_set_knob_offset
func LvArcSetKnobOffset(obj *core.LvObjT, offset c.Int32T)

//go:linkname LvArcGetAngleStart C.lv_arc_get_angle_start
func LvArcGetAngleStart(obj *core.LvObjT) lv_value_precise_t

//go:linkname LvArcGetAngleEnd C.lv_arc_get_angle_end
func LvArcGetAngleEnd(obj *core.LvObjT) lv_value_precise_t

//go:linkname LvArcGetBgAngleStart C.lv_arc_get_bg_angle_start
func LvArcGetBgAngleStart(obj *core.LvObjT) lv_value_precise_t

//go:linkname LvArcGetBgAngleEnd C.lv_arc_get_bg_angle_end
func LvArcGetBgAngleEnd(obj *core.LvObjT) lv_value_precise_t

//go:linkname LvArcGetValue C.lv_arc_get_value
func LvArcGetValue(obj *core.LvObjT) c.Int32T

//go:linkname LvArcGetMinValue C.lv_arc_get_min_value
func LvArcGetMinValue(obj *core.LvObjT) c.Int32T

//go:linkname LvArcGetMaxValue C.lv_arc_get_max_value
func LvArcGetMaxValue(obj *core.LvObjT) c.Int32T

//go:linkname LvArcGetMode C.lv_arc_get_mode
func LvArcGetMode(obj *core.LvObjT) lv_arc_mode_t

//go:linkname LvArcGetRotation C.lv_arc_get_rotation
func LvArcGetRotation(obj *core.LvObjT) c.Int32T

//go:linkname LvArcGetKnobOffset C.lv_arc_get_knob_offset
func LvArcGetKnobOffset(obj *core.LvObjT) c.Int32T

//go:linkname LvArcAlignObjToAngle C.lv_arc_align_obj_to_angle
func LvArcAlignObjToAngle(obj *core.LvObjT, obj_to_align *core.LvObjT, r_offset c.Int32T)

//go:linkname LvArcRotateObjToAngle C.lv_arc_rotate_obj_to_angle
func LvArcRotateObjToAngle(obj *core.LvObjT, obj_to_rotate *core.LvObjT, r_offset c.Int32T)

// Note: lv_value_precise_t is typically float or int32_t depending on LV_USE_FLOAT
type lv_value_precise_t c.Float
