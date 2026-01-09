package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvSliderCreate C.lv_slider_create
func LvSliderCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvSliderSetValue C.lv_slider_set_value
func LvSliderSetValue(obj *core.LvObjT, value c.Int32T, anim c.Int32T)

//go:linkname LvSliderSetStartValue C.lv_slider_set_start_value
func LvSliderSetStartValue(obj *core.LvObjT, value c.Int32T, anim c.Int32T)

//go:linkname LvSliderSetRange C.lv_slider_set_range
func LvSliderSetRange(obj *core.LvObjT, min c.Int32T, max c.Int32T)

//go:linkname LvSliderSetMode C.lv_slider_set_mode
func LvSliderSetMode(obj *core.LvObjT, mode c.Int32T)

//go:linkname LvSliderSetOrientation C.lv_slider_set_orientation
func LvSliderSetOrientation(obj *core.LvObjT, orientation c.Int32T)

//go:linkname LvSliderGetValue C.lv_slider_get_value
func LvSliderGetValue(obj *core.LvObjT) c.Int32T

//go:linkname LvSliderGetLeftValue C.lv_slider_get_left_value
func LvSliderGetLeftValue(obj *core.LvObjT) c.Int32T

//go:linkname LvSliderGetMinValue C.lv_slider_get_min_value
func LvSliderGetMinValue(obj *core.LvObjT) c.Int32T

//go:linkname LvSliderGetMaxValue C.lv_slider_get_max_value
func LvSliderGetMaxValue(obj *core.LvObjT) c.Int32T

//go:linkname LvSliderIsDragged C.lv_slider_is_dragged
func LvSliderIsDragged(obj *core.LvObjT) c.Char

//go:linkname LvSliderGetMode C.lv_slider_get_mode
func LvSliderGetMode(slider *core.LvObjT) c.Int32T

//go:linkname LvSliderGetOrientation C.lv_slider_get_orientation
func LvSliderGetOrientation(slider *core.LvObjT) c.Int32T

//go:linkname LvSliderIsSymmetrical C.lv_slider_is_symmetrical
func LvSliderIsSymmetrical(obj *core.LvObjT) c.Char
