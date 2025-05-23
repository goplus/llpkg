package core

import (
	"unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvObjSetStyleWidth C.lv_obj_set_style_width
func LvObjSetStyleWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleMinWidth C.lv_obj_set_style_min_width
func LvObjSetStyleMinWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleMaxWidth C.lv_obj_set_style_max_width
func LvObjSetStyleMaxWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleHeight C.lv_obj_set_style_height
func LvObjSetStyleHeight(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleMinHeight C.lv_obj_set_style_min_height
func LvObjSetStyleMinHeight(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleMaxHeight C.lv_obj_set_style_max_height
func LvObjSetStyleMaxHeight(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleLength C.lv_obj_set_style_length
func LvObjSetStyleLength(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleX C.lv_obj_set_style_x
func LvObjSetStyleX(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleY C.lv_obj_set_style_y
func LvObjSetStyleY(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleAlign C.lv_obj_set_style_align
func LvObjSetStyleAlign(obj *LvObjT, value LvAlignT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransformWidth C.lv_obj_set_style_transform_width
func LvObjSetStyleTransformWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransformHeight C.lv_obj_set_style_transform_height
func LvObjSetStyleTransformHeight(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTranslateX C.lv_obj_set_style_translate_x
func LvObjSetStyleTranslateX(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTranslateY C.lv_obj_set_style_translate_y
func LvObjSetStyleTranslateY(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTranslateRadial C.lv_obj_set_style_translate_radial
func LvObjSetStyleTranslateRadial(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransformScaleX C.lv_obj_set_style_transform_scale_x
func LvObjSetStyleTransformScaleX(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransformScaleY C.lv_obj_set_style_transform_scale_y
func LvObjSetStyleTransformScaleY(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransformRotation C.lv_obj_set_style_transform_rotation
func LvObjSetStyleTransformRotation(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransformPivotX C.lv_obj_set_style_transform_pivot_x
func LvObjSetStyleTransformPivotX(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransformPivotY C.lv_obj_set_style_transform_pivot_y
func LvObjSetStyleTransformPivotY(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransformSkewX C.lv_obj_set_style_transform_skew_x
func LvObjSetStyleTransformSkewX(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransformSkewY C.lv_obj_set_style_transform_skew_y
func LvObjSetStyleTransformSkewY(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStylePadTop C.lv_obj_set_style_pad_top
func LvObjSetStylePadTop(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStylePadBottom C.lv_obj_set_style_pad_bottom
func LvObjSetStylePadBottom(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStylePadLeft C.lv_obj_set_style_pad_left
func LvObjSetStylePadLeft(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStylePadRight C.lv_obj_set_style_pad_right
func LvObjSetStylePadRight(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStylePadRow C.lv_obj_set_style_pad_row
func LvObjSetStylePadRow(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStylePadColumn C.lv_obj_set_style_pad_column
func LvObjSetStylePadColumn(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStylePadRadial C.lv_obj_set_style_pad_radial
func LvObjSetStylePadRadial(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleMarginTop C.lv_obj_set_style_margin_top
func LvObjSetStyleMarginTop(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleMarginBottom C.lv_obj_set_style_margin_bottom
func LvObjSetStyleMarginBottom(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleMarginLeft C.lv_obj_set_style_margin_left
func LvObjSetStyleMarginLeft(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleMarginRight C.lv_obj_set_style_margin_right
func LvObjSetStyleMarginRight(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgColor C.lv_obj_set_style_bg_color
func LvObjSetStyleBgColor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgOpa C.lv_obj_set_style_bg_opa
func LvObjSetStyleBgOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgGradColor C.lv_obj_set_style_bg_grad_color
func LvObjSetStyleBgGradColor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgGradDir C.lv_obj_set_style_bg_grad_dir
func LvObjSetStyleBgGradDir(obj *LvObjT, value LvGradDirT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgMainStop C.lv_obj_set_style_bg_main_stop
func LvObjSetStyleBgMainStop(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgGradStop C.lv_obj_set_style_bg_grad_stop
func LvObjSetStyleBgGradStop(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgMainOpa C.lv_obj_set_style_bg_main_opa
func LvObjSetStyleBgMainOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgGradOpa C.lv_obj_set_style_bg_grad_opa
func LvObjSetStyleBgGradOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgGrad C.lv_obj_set_style_bg_grad
func LvObjSetStyleBgGrad(obj *LvObjT, value *LvGradDscT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgImageSrc C.lv_obj_set_style_bg_image_src
func LvObjSetStyleBgImageSrc(obj *LvObjT, value unsafe.Pointer, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgImageOpa C.lv_obj_set_style_bg_image_opa
func LvObjSetStyleBgImageOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgImageRecolor C.lv_obj_set_style_bg_image_recolor
func LvObjSetStyleBgImageRecolor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgImageRecolorOpa C.lv_obj_set_style_bg_image_recolor_opa
func LvObjSetStyleBgImageRecolorOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBgImageTiled C.lv_obj_set_style_bg_image_tiled
func LvObjSetStyleBgImageTiled(obj *LvObjT, value bool, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBorderColor C.lv_obj_set_style_border_color
func LvObjSetStyleBorderColor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBorderOpa C.lv_obj_set_style_border_opa
func LvObjSetStyleBorderOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBorderWidth C.lv_obj_set_style_border_width
func LvObjSetStyleBorderWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBorderSide C.lv_obj_set_style_border_side
func LvObjSetStyleBorderSide(obj *LvObjT, value LvBorderSideT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBorderPost C.lv_obj_set_style_border_post
func LvObjSetStyleBorderPost(obj *LvObjT, value bool, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleOutlineWidth C.lv_obj_set_style_outline_width
func LvObjSetStyleOutlineWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleOutlineColor C.lv_obj_set_style_outline_color
func LvObjSetStyleOutlineColor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleOutlineOpa C.lv_obj_set_style_outline_opa
func LvObjSetStyleOutlineOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleOutlinePad C.lv_obj_set_style_outline_pad
func LvObjSetStyleOutlinePad(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleShadowWidth C.lv_obj_set_style_shadow_width
func LvObjSetStyleShadowWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleShadowOffsetX C.lv_obj_set_style_shadow_offset_x
func LvObjSetStyleShadowOffsetX(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleShadowOffsetY C.lv_obj_set_style_shadow_offset_y
func LvObjSetStyleShadowOffsetY(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleShadowSpread C.lv_obj_set_style_shadow_spread
func LvObjSetStyleShadowSpread(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleShadowColor C.lv_obj_set_style_shadow_color
func LvObjSetStyleShadowColor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleShadowOpa C.lv_obj_set_style_shadow_opa
func LvObjSetStyleShadowOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleImageOpa C.lv_obj_set_style_image_opa
func LvObjSetStyleImageOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleImageRecolor C.lv_obj_set_style_image_recolor
func LvObjSetStyleImageRecolor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleImageRecolorOpa C.lv_obj_set_style_image_recolor_opa
func LvObjSetStyleImageRecolorOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleLineWidth C.lv_obj_set_style_line_width
func LvObjSetStyleLineWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleLineDashWidth C.lv_obj_set_style_line_dash_width
func LvObjSetStyleLineDashWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleLineDashGap C.lv_obj_set_style_line_dash_gap
func LvObjSetStyleLineDashGap(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleLineRounded C.lv_obj_set_style_line_rounded
func LvObjSetStyleLineRounded(obj *LvObjT, value bool, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleLineColor C.lv_obj_set_style_line_color
func LvObjSetStyleLineColor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleLineOpa C.lv_obj_set_style_line_opa
func LvObjSetStyleLineOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleArcWidth C.lv_obj_set_style_arc_width
func LvObjSetStyleArcWidth(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleArcRounded C.lv_obj_set_style_arc_rounded
func LvObjSetStyleArcRounded(obj *LvObjT, value bool, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleArcColor C.lv_obj_set_style_arc_color
func LvObjSetStyleArcColor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleArcOpa C.lv_obj_set_style_arc_opa
func LvObjSetStyleArcOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleArcImageSrc C.lv_obj_set_style_arc_image_src
func LvObjSetStyleArcImageSrc(obj *LvObjT, value unsafe.Pointer, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTextColor C.lv_obj_set_style_text_color
func LvObjSetStyleTextColor(obj *LvObjT, value LvColorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTextOpa C.lv_obj_set_style_text_opa
func LvObjSetStyleTextOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTextFont C.lv_obj_set_style_text_font
func LvObjSetStyleTextFont(obj *LvObjT, value *LvFontT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTextLetterSpace C.lv_obj_set_style_text_letter_space
func LvObjSetStyleTextLetterSpace(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTextLineSpace C.lv_obj_set_style_text_line_space
func LvObjSetStyleTextLineSpace(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTextDecor C.lv_obj_set_style_text_decor
func LvObjSetStyleTextDecor(obj *LvObjT, value LvTextDecorT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTextAlign C.lv_obj_set_style_text_align
func LvObjSetStyleTextAlign(obj *LvObjT, value LvTextAlignT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleRadius C.lv_obj_set_style_radius
func LvObjSetStyleRadius(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleRadialOffset C.lv_obj_set_style_radial_offset
func LvObjSetStyleRadialOffset(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleClipCorner C.lv_obj_set_style_clip_corner
func LvObjSetStyleClipCorner(obj *LvObjT, value bool, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleOpa C.lv_obj_set_style_opa
func LvObjSetStyleOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleOpaLayered C.lv_obj_set_style_opa_layered
func LvObjSetStyleOpaLayered(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleColorFilterDsc C.lv_obj_set_style_color_filter_dsc
func LvObjSetStyleColorFilterDsc(obj *LvObjT, value *LvColorFilterDscT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleColorFilterOpa C.lv_obj_set_style_color_filter_opa
func LvObjSetStyleColorFilterOpa(obj *LvObjT, value LvOpaT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleAnim C.lv_obj_set_style_anim
func LvObjSetStyleAnim(obj *LvObjT, value *LvAnimT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleAnimDuration C.lv_obj_set_style_anim_duration
func LvObjSetStyleAnimDuration(obj *LvObjT, value c.Uint32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleTransition C.lv_obj_set_style_transition
func LvObjSetStyleTransition(obj *LvObjT, value *LvStyleTransitionDscT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBlendMode C.lv_obj_set_style_blend_mode
func LvObjSetStyleBlendMode(obj *LvObjT, value LvBlendModeT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleLayout C.lv_obj_set_style_layout
func LvObjSetStyleLayout(obj *LvObjT, value c.Uint16T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBaseDir C.lv_obj_set_style_base_dir
func LvObjSetStyleBaseDir(obj *LvObjT, value LvBaseDirT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleBitmapMaskSrc C.lv_obj_set_style_bitmap_mask_src
func LvObjSetStyleBitmapMaskSrc(obj *LvObjT, value unsafe.Pointer, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleRotarySensitivity C.lv_obj_set_style_rotary_sensitivity
func LvObjSetStyleRotarySensitivity(obj *LvObjT, value c.Uint32T, selector LvStyleSelectorT)

// Flex layout functions
//
//go:linkname LvObjSetStyleFlexFlow C.lv_obj_set_style_flex_flow
func LvObjSetStyleFlexFlow(obj *LvObjT, value LvFlexFlowT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleFlexMainPlace C.lv_obj_set_style_flex_main_place
func LvObjSetStyleFlexMainPlace(obj *LvObjT, value LvFlexAlignT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleFlexCrossPlace C.lv_obj_set_style_flex_cross_place
func LvObjSetStyleFlexCrossPlace(obj *LvObjT, value LvFlexAlignT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleFlexTrackPlace C.lv_obj_set_style_flex_track_place
func LvObjSetStyleFlexTrackPlace(obj *LvObjT, value LvFlexAlignT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleFlexGrow C.lv_obj_set_style_flex_grow
func LvObjSetStyleFlexGrow(obj *LvObjT, value c.Uint8T, selector LvStyleSelectorT)

// Grid layout functions
//
//go:linkname LvObjSetStyleGridColumnDscArray C.lv_obj_set_style_grid_column_dsc_array
func LvObjSetStyleGridColumnDscArray(obj *LvObjT, value *c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleGridColumnAlign C.lv_obj_set_style_grid_column_align
func LvObjSetStyleGridColumnAlign(obj *LvObjT, value LvGridAlignT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleGridRowDscArray C.lv_obj_set_style_grid_row_dsc_array
func LvObjSetStyleGridRowDscArray(obj *LvObjT, value *c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleGridRowAlign C.lv_obj_set_style_grid_row_align
func LvObjSetStyleGridRowAlign(obj *LvObjT, value LvGridAlignT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleGridCellColumnPos C.lv_obj_set_style_grid_cell_column_pos
func LvObjSetStyleGridCellColumnPos(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleGridCellXAlign C.lv_obj_set_style_grid_cell_x_align
func LvObjSetStyleGridCellXAlign(obj *LvObjT, value LvGridAlignT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleGridCellColumnSpan C.lv_obj_set_style_grid_cell_column_span
func LvObjSetStyleGridCellColumnSpan(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleGridCellRowPos C.lv_obj_set_style_grid_cell_row_pos
func LvObjSetStyleGridCellRowPos(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleGridCellYAlign C.lv_obj_set_style_grid_cell_y_align
func LvObjSetStyleGridCellYAlign(obj *LvObjT, value LvGridAlignT, selector LvStyleSelectorT)

//go:linkname LvObjSetStyleGridCellRowSpan C.lv_obj_set_style_grid_cell_row_span
func LvObjSetStyleGridCellRowSpan(obj *LvObjT, value c.Int32T, selector LvStyleSelectorT)
