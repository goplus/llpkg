package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvObjAddStyle C.lv_obj_add_style
func LvObjAddStyle(obj *LvObjT, style *LvStyleT, selector LvStyleSelectorT)

//go:linkname LvObjReplaceStyle C.lv_obj_replace_style
func LvObjReplaceStyle(obj *LvObjT, old_style *LvStyleT, new_style *LvStyleT, selector LvStyleSelectorT) bool

//go:linkname LvObjRemoveStyle C.lv_obj_remove_style
func LvObjRemoveStyle(obj *LvObjT, style *LvStyleT, selector LvStyleSelectorT)

//go:linkname LvObjRemoveStyleAll C.lv_obj_remove_style_all
func LvObjRemoveStyleAll(obj *LvObjT)

//go:linkname LvObjReportStyleChange C.lv_obj_report_style_change
func LvObjReportStyleChange(style *LvStyleT)

//go:linkname LvObjRefreshStyle C.lv_obj_refresh_style
func LvObjRefreshStyle(obj *LvObjT, part LvPartT, prop LvStylePropT)

//go:linkname LvObjEnableStyleRefresh C.lv_obj_enable_style_refresh
func LvObjEnableStyleRefresh(en bool)

//go:linkname LvObjGetStyleProp C.lv_obj_get_style_prop
func LvObjGetStyleProp(obj *LvObjT, part LvPartT, prop LvStylePropT) LvStyleValueT

//go:linkname LvObjHasStyleProp C.lv_obj_has_style_prop
func LvObjHasStyleProp(obj *LvObjT, selector LvStyleSelectorT, prop LvStylePropT) bool

//go:linkname LvObjSetLocalStyleProp C.lv_obj_set_local_style_prop
func LvObjSetLocalStyleProp(obj *LvObjT, prop LvStylePropT, value LvStyleValueT, selector LvStyleSelectorT)

//go:linkname LvObjGetLocalStyleProp C.lv_obj_get_local_style_prop
func LvObjGetLocalStyleProp(obj *LvObjT, prop LvStylePropT, value *LvStyleValueT, selector LvStyleSelectorT) LvStyleResT

//go:linkname LvObjRemoveLocalStyleProp C.lv_obj_remove_local_style_prop
func LvObjRemoveLocalStyleProp(obj *LvObjT, prop LvStylePropT, selector LvStyleSelectorT) bool

//go:linkname LvObjStyleApplyColorFilter C.lv_obj_style_apply_color_filter
func LvObjStyleApplyColorFilter(obj *LvObjT, part LvPartT, v LvStyleValueT) LvStyleValueT

//go:linkname LvObjFadeIn C.lv_obj_fade_in
func LvObjFadeIn(obj *LvObjT, time c.Uint32T, delay c.Uint32T)

//go:linkname LvObjFadeOut C.lv_obj_fade_out
func LvObjFadeOut(obj *LvObjT, time c.Uint32T, delay c.Uint32T)
