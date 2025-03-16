package core

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvStyleSetWidth C.lv_style_set_width
func LvStyleSetWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetBgOpa C.lv_style_set_bg_opa
func LvStyleSetBgOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetBgColor C.lv_style_set_bg_color
func LvStyleSetBgColor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetRadius C.lv_style_set_radius
func LvStyleSetRadius(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetPadRight C.lv_style_set_pad_right
func LvStyleSetPadRight(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetMinWidth C.lv_style_set_min_width
func LvStyleSetMinWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetMaxWidth C.lv_style_set_max_width
func LvStyleSetMaxWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetHeight C.lv_style_set_height
func LvStyleSetHeight(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetMinHeight C.lv_style_set_min_height
func LvStyleSetMinHeight(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetMaxHeight C.lv_style_set_max_height
func LvStyleSetMaxHeight(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetLength C.lv_style_set_length
func LvStyleSetLength(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetX C.lv_style_set_x
func LvStyleSetX(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetY C.lv_style_set_y
func LvStyleSetY(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetAlign C.lv_style_set_align
func LvStyleSetAlign(style *LvStyleT, value LvAlignT)

//go:linkname LvStyleSetTransformWidth C.lv_style_set_transform_width
func LvStyleSetTransformWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTransformHeight C.lv_style_set_transform_height
func LvStyleSetTransformHeight(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTranslateX C.lv_style_set_translate_x
func LvStyleSetTranslateX(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTranslateY C.lv_style_set_translate_y
func LvStyleSetTranslateY(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTranslateRadial C.lv_style_set_translate_radial
func LvStyleSetTranslateRadial(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTransformScaleX C.lv_style_set_transform_scale_x
func LvStyleSetTransformScaleX(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTransformScaleY C.lv_style_set_transform_scale_y
func LvStyleSetTransformScaleY(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTransformRotation C.lv_style_set_transform_rotation
func LvStyleSetTransformRotation(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTransformPivotX C.lv_style_set_transform_pivot_x
func LvStyleSetTransformPivotX(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTransformPivotY C.lv_style_set_transform_pivot_y
func LvStyleSetTransformPivotY(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTransformSkewX C.lv_style_set_transform_skew_x
func LvStyleSetTransformSkewX(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTransformSkewY C.lv_style_set_transform_skew_y
func LvStyleSetTransformSkewY(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetPadTop C.lv_style_set_pad_top
func LvStyleSetPadTop(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetPadBottom C.lv_style_set_pad_bottom
func LvStyleSetPadBottom(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetPadLeft C.lv_style_set_pad_left
func LvStyleSetPadLeft(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetPadRow C.lv_style_set_pad_row
func LvStyleSetPadRow(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetPadColumn C.lv_style_set_pad_column
func LvStyleSetPadColumn(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetPadRadial C.lv_style_set_pad_radial
func LvStyleSetPadRadial(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetMarginTop C.lv_style_set_margin_top
func LvStyleSetMarginTop(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetMarginBottom C.lv_style_set_margin_bottom
func LvStyleSetMarginBottom(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetMarginLeft C.lv_style_set_margin_left
func LvStyleSetMarginLeft(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetMarginRight C.lv_style_set_margin_right
func LvStyleSetMarginRight(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetBgGradColor C.lv_style_set_bg_grad_color
func LvStyleSetBgGradColor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetBgGradDir C.lv_style_set_bg_grad_dir
func LvStyleSetBgGradDir(style *LvStyleT, value LvGradDirT)

//go:linkname LvStyleSetBgMainStop C.lv_style_set_bg_main_stop
func LvStyleSetBgMainStop(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetBgGradStop C.lv_style_set_bg_grad_stop
func LvStyleSetBgGradStop(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetBgMainOpa C.lv_style_set_bg_main_opa
func LvStyleSetBgMainOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetBgGradOpa C.lv_style_set_bg_grad_opa
func LvStyleSetBgGradOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetBgGrad C.lv_style_set_bg_grad
func LvStyleSetBgGrad(style *LvStyleT, value *LvGradDscT)

//go:linkname LvStyleSetBgImageSrc C.lv_style_set_bg_image_src
func LvStyleSetBgImageSrc(style *LvStyleT, value unsafe.Pointer)

//go:linkname LvStyleSetBgImageOpa C.lv_style_set_bg_image_opa
func LvStyleSetBgImageOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetBgImageRecolor C.lv_style_set_bg_image_recolor
func LvStyleSetBgImageRecolor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetBgImageRecolorOpa C.lv_style_set_bg_image_recolor_opa
func LvStyleSetBgImageRecolorOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetBgImageTiled C.lv_style_set_bg_image_tiled
func LvStyleSetBgImageTiled(style *LvStyleT, value bool)

//go:linkname LvStyleSetBorderColor C.lv_style_set_border_color
func LvStyleSetBorderColor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetBorderOpa C.lv_style_set_border_opa
func LvStyleSetBorderOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetBorderWidth C.lv_style_set_border_width
func LvStyleSetBorderWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetBorderSide C.lv_style_set_border_side
func LvStyleSetBorderSide(style *LvStyleT, value LvBorderSideT)

//go:linkname LvStyleSetBorderPost C.lv_style_set_border_post
func LvStyleSetBorderPost(style *LvStyleT, value bool)

//go:linkname LvStyleSetOutlineWidth C.lv_style_set_outline_width
func LvStyleSetOutlineWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetOutlineColor C.lv_style_set_outline_color
func LvStyleSetOutlineColor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetOutlineOpa C.lv_style_set_outline_opa
func LvStyleSetOutlineOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetOutlinePad C.lv_style_set_outline_pad
func LvStyleSetOutlinePad(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetShadowWidth C.lv_style_set_shadow_width
func LvStyleSetShadowWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetShadowOffsetX C.lv_style_set_shadow_offset_x
func LvStyleSetShadowOffsetX(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetShadowOffsetY C.lv_style_set_shadow_offset_y
func LvStyleSetShadowOffsetY(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetShadowSpread C.lv_style_set_shadow_spread
func LvStyleSetShadowSpread(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetShadowColor C.lv_style_set_shadow_color
func LvStyleSetShadowColor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetShadowOpa C.lv_style_set_shadow_opa
func LvStyleSetShadowOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetImageOpa C.lv_style_set_image_opa
func LvStyleSetImageOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetImageRecolor C.lv_style_set_image_recolor
func LvStyleSetImageRecolor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetImageRecolorOpa C.lv_style_set_image_recolor_opa
func LvStyleSetImageRecolorOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetLineWidth C.lv_style_set_line_width
func LvStyleSetLineWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetLineDashWidth C.lv_style_set_line_dash_width
func LvStyleSetLineDashWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetLineDashGap C.lv_style_set_line_dash_gap
func LvStyleSetLineDashGap(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetLineRounded C.lv_style_set_line_rounded
func LvStyleSetLineRounded(style *LvStyleT, value bool)

//go:linkname LvStyleSetLineColor C.lv_style_set_line_color
func LvStyleSetLineColor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetLineOpa C.lv_style_set_line_opa
func LvStyleSetLineOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetArcWidth C.lv_style_set_arc_width
func LvStyleSetArcWidth(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetArcRounded C.lv_style_set_arc_rounded
func LvStyleSetArcRounded(style *LvStyleT, value bool)

//go:linkname LvStyleSetArcColor C.lv_style_set_arc_color
func LvStyleSetArcColor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetArcOpa C.lv_style_set_arc_opa
func LvStyleSetArcOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetArcImageSrc C.lv_style_set_arc_image_src
func LvStyleSetArcImageSrc(style *LvStyleT, value unsafe.Pointer)

//go:linkname LvStyleSetTextColor C.lv_style_set_text_color
func LvStyleSetTextColor(style *LvStyleT, value LvColorT)

//go:linkname LvStyleSetTextOpa C.lv_style_set_text_opa
func LvStyleSetTextOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetTextFont C.lv_style_set_text_font
func LvStyleSetTextFont(style *LvStyleT, value *LvFontT)

//go:linkname LvStyleSetTextLetterSpace C.lv_style_set_text_letter_space
func LvStyleSetTextLetterSpace(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTextLineSpace C.lv_style_set_text_line_space
func LvStyleSetTextLineSpace(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetTextDecor C.lv_style_set_text_decor
func LvStyleSetTextDecor(style *LvStyleT, value LvTextDecorT)

//go:linkname LvStyleSetTextAlign C.lv_style_set_text_align
func LvStyleSetTextAlign(style *LvStyleT, value LvTextAlignT)

//go:linkname LvStyleSetRadialOffset C.lv_style_set_radial_offset
func LvStyleSetRadialOffset(style *LvStyleT, value c.Int)

//go:linkname LvStyleSetClipCorner C.lv_style_set_clip_corner
func LvStyleSetClipCorner(style *LvStyleT, value bool)

//go:linkname LvStyleSetOpa C.lv_style_set_opa
func LvStyleSetOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetOpaLayered C.lv_style_set_opa_layered
func LvStyleSetOpaLayered(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetColorFilterDsc C.lv_style_set_color_filter_dsc
func LvStyleSetColorFilterDsc(style *LvStyleT, value *LvColorFilterDscT)

//go:linkname LvStyleSetColorFilterOpa C.lv_style_set_color_filter_opa
func LvStyleSetColorFilterOpa(style *LvStyleT, value LvOpaT)

//go:linkname LvStyleSetAnim C.lv_style_set_anim
func LvStyleSetAnim(style *LvStyleT, value *LvAnimT)

//go:linkname LvStyleSetAnimDuration C.lv_style_set_anim_duration
func LvStyleSetAnimDuration(style *LvStyleT, value c.Uint)

//go:linkname LvStyleSetTransition C.lv_style_set_transition
func LvStyleSetTransition(style *LvStyleT, value *LvStyleTransitionDscT)

//go:linkname LvStyleSetBlendMode C.lv_style_set_blend_mode
func LvStyleSetBlendMode(style *LvStyleT, value LvBlendModeT)

//go:linkname LvStyleSetLayout C.lv_style_set_layout
func LvStyleSetLayout(style *LvStyleT, value c.Int16T)

//go:linkname LvStyleSetBaseDir C.lv_style_set_base_dir
func LvStyleSetBaseDir(style *LvStyleT, value LvBaseDirT)

//go:linkname LvStyleSetBitmapMaskSrc C.lv_style_set_bitmap_mask_src
func LvStyleSetBitmapMaskSrc(style *LvStyleT, value *c.Void)

//go:linkname LvStyleSetRotarySensitivity C.lv_style_set_rotary_sensitivity
func LvStyleSetRotarySensitivity(style *LvStyleT, value c.Int16T)
