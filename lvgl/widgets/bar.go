package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

type lv_bar_mode_t c.Int

const (
	LV_BAR_MODE_NORMAL lv_bar_mode_t = iota
	LV_BAR_MODE_SYMMETRICAL
	LV_BAR_MODE_RANGE
)

type lv_bar_orientation_t c.Int

const (
	LV_BAR_ORIENTATION_AUTO lv_bar_orientation_t = iota
	LV_BAR_ORIENTATION_HORIZONTAL
	LV_BAR_ORIENTATION_VERTICAL
)

// lv_anim_enable_t is typically defined in lv_anim.h
type lv_anim_enable_t c.Int

const (
	LV_ANIM_OFF lv_anim_enable_t = 0
	LV_ANIM_ON  lv_anim_enable_t = 1
)

//go:linkname LvBarCreate C.lv_bar_create
func LvBarCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvBarSetValue C.lv_bar_set_value
func LvBarSetValue(obj *core.LvObjT, value c.Int32T, anim lv_anim_enable_t)

//go:linkname LvBarSetStartValue C.lv_bar_set_start_value
func LvBarSetStartValue(obj *core.LvObjT, start_value c.Int32T, anim lv_anim_enable_t)

//go:linkname LvBarSetRange C.lv_bar_set_range
func LvBarSetRange(obj *core.LvObjT, min, max c.Int32T)

//go:linkname LvBarSetMode C.lv_bar_set_mode
func LvBarSetMode(obj *core.LvObjT, mode lv_bar_mode_t)

//go:linkname LvBarSetOrientation C.lv_bar_set_orientation
func LvBarSetOrientation(obj *core.LvObjT, orientation lv_bar_orientation_t)

//go:linkname LvBarGetValue C.lv_bar_get_value
func LvBarGetValue(obj *core.LvObjT) c.Int32T

//go:linkname LvBarGetStartValue C.lv_bar_get_start_value
func LvBarGetStartValue(obj *core.LvObjT) c.Int32T

//go:linkname LvBarGetMinValue C.lv_bar_get_min_value
func LvBarGetMinValue(obj *core.LvObjT) c.Int32T

//go:linkname LvBarGetMaxValue C.lv_bar_get_max_value
func LvBarGetMaxValue(obj *core.LvObjT) c.Int32T

//go:linkname LvBarGetMode C.lv_bar_get_mode
func LvBarGetMode(obj *core.LvObjT) lv_bar_mode_t

//go:linkname LvBarGetOrientation C.lv_bar_get_orientation
func LvBarGetOrientation(obj *core.LvObjT) lv_bar_orientation_t

//go:linkname LvBarIsSymmetrical C.lv_bar_is_symmetrical
func LvBarIsSymmetrical(obj *core.LvObjT) bool
