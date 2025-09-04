package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Scale section type
type LvScaleSectionT struct{}

//go:linkname LvScaleCreate C.lv_scale_create
func LvScaleCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvScaleSetMode C.lv_scale_set_mode
func LvScaleSetMode(obj *core.LvObjT, mode c.Int32T)

//go:linkname LvScaleSetTotalTickCount C.lv_scale_set_total_tick_count
func LvScaleSetTotalTickCount(obj *core.LvObjT, totalTickCount c.Uint32T)

//go:linkname LvScaleSetMajorTickEvery C.lv_scale_set_major_tick_every
func LvScaleSetMajorTickEvery(obj *core.LvObjT, majorTickEvery c.Uint32T)

//go:linkname LvScaleSetLabelShow C.lv_scale_set_label_show
func LvScaleSetLabelShow(obj *core.LvObjT, showLabel c.Char)

//go:linkname LvScaleSetRange C.lv_scale_set_range
func LvScaleSetRange(obj *core.LvObjT, min c.Int32T, max c.Int32T)

//go:linkname LvScaleSetAngleRange C.lv_scale_set_angle_range
func LvScaleSetAngleRange(obj *core.LvObjT, angleRange c.Uint32T)

//go:linkname LvScaleSetRotation C.lv_scale_set_rotation
func LvScaleSetRotation(obj *core.LvObjT, rotation c.Int32T)

//go:linkname LvScaleSetLineNeedleValue C.lv_scale_set_line_needle_value
func LvScaleSetLineNeedleValue(obj *core.LvObjT, needleLine *core.LvObjT, needleLength c.Int32T, value c.Int32T)

//go:linkname LvScaleSetImageNeedleValue C.lv_scale_set_image_needle_value
func LvScaleSetImageNeedleValue(obj *core.LvObjT, needleImg *core.LvObjT, value c.Int32T)

//go:linkname LvScaleSetTextSrc C.lv_scale_set_text_src
func LvScaleSetTextSrc(obj *core.LvObjT, txtSrc **c.Char)

//go:linkname LvScaleSetPostDraw C.lv_scale_set_post_draw
func LvScaleSetPostDraw(obj *core.LvObjT, en c.Char)

//go:linkname LvScaleSetDrawTicksOnTop C.lv_scale_set_draw_ticks_on_top
func LvScaleSetDrawTicksOnTop(obj *core.LvObjT, en c.Char)

//go:linkname LvScaleAddSection C.lv_scale_add_section
func LvScaleAddSection(obj *core.LvObjT) *LvScaleSectionT

//go:linkname LvScaleSetSectionRange C.lv_scale_set_section_range
func LvScaleSetSectionRange(scale *core.LvObjT, section *LvScaleSectionT, min c.Int32T, max c.Int32T)

//go:linkname LvScaleSetSectionStyleMain C.lv_scale_set_section_style_main
func LvScaleSetSectionStyleMain(scale *core.LvObjT, section *LvScaleSectionT, style *core.LvStyleT)

//go:linkname LvScaleSetSectionStyleIndicator C.lv_scale_set_section_style_indicator
func LvScaleSetSectionStyleIndicator(scale *core.LvObjT, section *LvScaleSectionT, style *core.LvStyleT)

//go:linkname LvScaleSetSectionStyleItems C.lv_scale_set_section_style_items
func LvScaleSetSectionStyleItems(scale *core.LvObjT, section *LvScaleSectionT, style *core.LvStyleT)

//go:linkname LvScaleGetMode C.lv_scale_get_mode
func LvScaleGetMode(obj *core.LvObjT) c.Int32T

//go:linkname LvScaleGetTotalTickCount C.lv_scale_get_total_tick_count
func LvScaleGetTotalTickCount(obj *core.LvObjT) c.Int32T

//go:linkname LvScaleGetMajorTickEvery C.lv_scale_get_major_tick_every
func LvScaleGetMajorTickEvery(obj *core.LvObjT) c.Int32T

//go:linkname LvScaleGetRotation C.lv_scale_get_rotation
func LvScaleGetRotation(obj *core.LvObjT) c.Int32T

//go:linkname LvScaleGetLabelShow C.lv_scale_get_label_show
func LvScaleGetLabelShow(obj *core.LvObjT) c.Char

//go:linkname LvScaleGetAngleRange C.lv_scale_get_angle_range
func LvScaleGetAngleRange(obj *core.LvObjT) c.Uint32T

//go:linkname LvScaleGetRangeMinValue C.lv_scale_get_range_min_value
func LvScaleGetRangeMinValue(obj *core.LvObjT) c.Int32T

//go:linkname LvScaleGetRangeMaxValue C.lv_scale_get_range_max_value
func LvScaleGetRangeMaxValue(obj *core.LvObjT) c.Int32T
