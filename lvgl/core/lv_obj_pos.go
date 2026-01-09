package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvObjSetPos C.lv_obj_set_pos
func LvObjSetPos(obj *LvObjT, x, y c.Int32T)

//go:linkname LvObjSetX C.lv_obj_set_x
func LvObjSetX(obj *LvObjT, x c.Int32T)

//go:linkname LvObjSetY C.lv_obj_set_y
func LvObjSetY(obj *LvObjT, y c.Int32T)

//go:linkname LvObjSetSize C.lv_obj_set_size
func LvObjSetSize(obj *LvObjT, w, h c.Int32T)

//go:linkname LvObjRefrSize C.lv_obj_refr_size
func LvObjRefrSize(obj *LvObjT) bool

//go:linkname LvObjSetWidth C.lv_obj_set_width
func LvObjSetWidth(obj *LvObjT, w c.Int32T)

//go:linkname LvObjSetHeight C.lv_obj_set_height
func LvObjSetHeight(obj *LvObjT, h c.Int32T)

//go:linkname LvObjSetContentWidth C.lv_obj_set_content_width
func LvObjSetContentWidth(obj *LvObjT, w c.Int32T)

//go:linkname LvObjSetContentHeight C.lv_obj_set_content_height
func LvObjSetContentHeight(obj *LvObjT, h c.Int32T)

//go:linkname LvObjSetLayout C.lv_obj_set_layout
func LvObjSetLayout(obj *LvObjT, layout c.Uint32T)

//go:linkname LvObjIsLayoutPositioned C.lv_obj_is_layout_positioned
func LvObjIsLayoutPositioned(obj *LvObjT) bool

//go:linkname LvObjMarkLayoutAsDirty C.lv_obj_mark_layout_as_dirty
func LvObjMarkLayoutAsDirty(obj *LvObjT)

//go:linkname LvObjUpdateLayout C.lv_obj_update_layout
func LvObjUpdateLayout(obj *LvObjT)

//go:linkname LvObjSetAlign C.lv_obj_set_align
func LvObjSetAlign(obj *LvObjT, align LvAlignT)

//go:linkname LvObjAlign C.lv_obj_align
func LvObjAlign(obj *LvObjT, align LvAlignT, x_ofs, y_ofs c.Int32T)

//go:linkname LvObjAlignTo C.lv_obj_align_to
func LvObjAlignTo(obj *LvObjT, base *LvObjT, align LvAlignT, x_ofs, y_ofs c.Int32T)

//go:linkname LvObjCenter C.lv_obj_center
func LvObjCenter(obj *LvObjT)

//go:linkname LvObjSetTransform C.lv_obj_set_transform
func LvObjSetTransform(obj *LvObjT, matrix *LvMatrixT)

//go:linkname LvObjResetTransform C.lv_obj_reset_transform
func LvObjResetTransform(obj *LvObjT)

//go:linkname LvObjGetCoords C.lv_obj_get_coords
func LvObjGetCoords(obj *LvObjT, coords *LvAreaT)

//go:linkname LvObjGetX C.lv_obj_get_x
func LvObjGetX(obj *LvObjT) c.Int32T

//go:linkname LvObjGetX2 C.lv_obj_get_x2
func LvObjGetX2(obj *LvObjT) c.Int32T

//go:linkname LvObjGetY C.lv_obj_get_y
func LvObjGetY(obj *LvObjT) c.Int32T

//go:linkname LvObjGetY2 C.lv_obj_get_y2
func LvObjGetY2(obj *LvObjT) c.Int32T

//go:linkname LvObjGetXAligned C.lv_obj_get_x_aligned
func LvObjGetXAligned(obj *LvObjT) c.Int32T

//go:linkname LvObjGetYAligned C.lv_obj_get_y_aligned
func LvObjGetYAligned(obj *LvObjT) c.Int32T

//go:linkname LvObjGetWidth C.lv_obj_get_width
func LvObjGetWidth(obj *LvObjT) c.Int32T

//go:linkname LvObjGetHeight C.lv_obj_get_height
func LvObjGetHeight(obj *LvObjT) c.Int32T

//go:linkname LvObjGetContentWidth C.lv_obj_get_content_width
func LvObjGetContentWidth(obj *LvObjT) c.Int32T

//go:linkname LvObjGetContentHeight C.lv_obj_get_content_height
func LvObjGetContentHeight(obj *LvObjT) c.Int32T

//go:linkname LvObjGetContentCoords C.lv_obj_get_content_coords
func LvObjGetContentCoords(obj *LvObjT, area *LvAreaT)

//go:linkname LvObjGetSelfWidth C.lv_obj_get_self_width
func LvObjGetSelfWidth(obj *LvObjT) c.Int32T

//go:linkname LvObjGetSelfHeight C.lv_obj_get_self_height
func LvObjGetSelfHeight(obj *LvObjT) c.Int32T

//go:linkname LvObjRefreshSelfSize C.lv_obj_refresh_self_size
func LvObjRefreshSelfSize(obj *LvObjT) bool

//go:linkname LvObjRefrPos C.lv_obj_refr_pos
func LvObjRefrPos(obj *LvObjT)

//go:linkname LvObjMoveTo C.lv_obj_move_to
func LvObjMoveTo(obj *LvObjT, x, y c.Int32T)

//go:linkname LvObjMoveChildrenBy C.lv_obj_move_children_by
func LvObjMoveChildrenBy(obj *LvObjT, x_diff, y_diff c.Int32T, ignore_floating bool)

//go:linkname LvObjGetTransform C.lv_obj_get_transform
func LvObjGetTransform(obj *LvObjT) *LvMatrixT

//go:linkname LvObjTransformPoint C.lv_obj_transform_point
func LvObjTransformPoint(obj *LvObjT, p *LvPointT, flags LvObjPointTransformFlagT)

//go:linkname LvObjTransformPointArray C.lv_obj_transform_point_array
func LvObjTransformPointArray(obj *LvObjT, points *LvPointT, count c.SizeT, flags LvObjPointTransformFlagT)

//go:linkname LvObjGetTransformedArea C.lv_obj_get_transformed_area
func LvObjGetTransformedArea(obj *LvObjT, area *LvAreaT, flags LvObjPointTransformFlagT)

//go:linkname LvObjInvalidateArea C.lv_obj_invalidate_area
func LvObjInvalidateArea(obj *LvObjT, area *LvAreaT)

//go:linkname LvObjInvalidate C.lv_obj_invalidate
func LvObjInvalidate(obj *LvObjT)

//go:linkname LvObjAreaIsVisible C.lv_obj_area_is_visible
func LvObjAreaIsVisible(obj *LvObjT, area *LvAreaT) bool

//go:linkname LvObjIsVisible C.lv_obj_is_visible
func LvObjIsVisible(obj *LvObjT) bool

//go:linkname LvObjSetExtClickArea C.lv_obj_set_ext_click_area
func LvObjSetExtClickArea(obj *LvObjT, size c.Int32T)

//go:linkname LvObjGetClickArea C.lv_obj_get_click_area
func LvObjGetClickArea(obj *LvObjT, area *LvAreaT)

//go:linkname LvObjHitTest C.lv_obj_hit_test
func LvObjHitTest(obj *LvObjT, point *LvPointT) bool

//go:linkname LvClampWidth C.lv_clamp_width
func LvClampWidth(width, min_width, max_width, ref_width c.Int32T) c.Int32T

//go:linkname LvClampHeight C.lv_clamp_height
func LvClampHeight(height, min_height, max_height, ref_height c.Int32T) c.Int32T
