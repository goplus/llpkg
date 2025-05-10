package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// Special coordinate constants
const (
	LV_COORD_TYPE_SHIFT  = 29
	LV_COORD_TYPE_MASK   = 3 << LV_COORD_TYPE_SHIFT
	LV_COORD_TYPE_PX     = 0 << LV_COORD_TYPE_SHIFT
	LV_COORD_TYPE_SPEC   = 1 << LV_COORD_TYPE_SHIFT
	LV_COORD_TYPE_PX_NEG = 3 << LV_COORD_TYPE_SHIFT
	LV_COORD_MAX         = (1 << LV_COORD_TYPE_SHIFT) - 1
	LV_COORD_MIN         = -LV_COORD_MAX
	LV_SIZE_CONTENT      = LV_COORD_TYPE_SPEC | LV_COORD_MAX
	LV_PCT_STORED_MAX    = LV_COORD_MAX - 1
	LV_PCT_POS_MAX       = LV_PCT_STORED_MAX / 2
)

//go:linkname LvAreaSet C.lv_area_set
func LvAreaSet(areaP *LvAreaT, x1 c.Int32T, y1 c.Int32T, x2 c.Int32T, y2 c.Int32T)

//go:linkname LvAreaGetWidth C.lv_area_get_width
func LvAreaGetWidth(areaP *LvAreaT) c.Int32T

//go:linkname LvAreaGetHeight C.lv_area_get_height
func LvAreaGetHeight(areaP *LvAreaT) c.Int32T

//go:linkname LvAreaSetWidth C.lv_area_set_width
func LvAreaSetWidth(areaP *LvAreaT, w c.Int32T)

//go:linkname LvAreaSetHeight C.lv_area_set_height
func LvAreaSetHeight(areaP *LvAreaT, h c.Int32T)

//go:linkname LvAreaGetSize C.lv_area_get_size
func LvAreaGetSize(areaP *LvAreaT) c.Uint32T

//go:linkname LvAreaIncrease C.lv_area_increase
func LvAreaIncrease(area *LvAreaT, wExtra c.Int32T, hExtra c.Int32T)

//go:linkname LvAreaMove C.lv_area_move
func LvAreaMove(area *LvAreaT, xOfs c.Int32T, yOfs c.Int32T)

//go:linkname LvAreaAlign C.lv_area_align
func LvAreaAlign(base *LvAreaT, toAlign *LvAreaT, align LvAlignT, ofsX c.Int32T, ofsY c.Int32T)

//go:linkname LvPointTransform C.lv_point_transform
func LvPointTransform(point *LvPointT, angle c.Int32T, scaleX c.Int32T, scaleY c.Int32T, pivot *LvPointT, zoomFirst bool)

//go:linkname LvPointArrayTransform C.lv_point_array_transform
func LvPointArrayTransform(points *LvPointT, count c.SizeT, angle c.Int32T, scaleX c.Int32T, scaleY c.Int32T, pivot *LvPointT, zoomFirst bool)

//go:linkname LvPointFromPrecise C.lv_point_from_precise
func LvPointFromPrecise(p *LvPointPreciseT) LvPointT

//go:linkname LvPointToPrecise C.lv_point_to_precise
func LvPointToPrecise(p *LvPointT) LvPointPreciseT

//go:linkname LvPointSet C.lv_point_set
func LvPointSet(p *LvPointT, x c.Int32T, y c.Int32T)

//go:linkname LvPointPreciseSet C.lv_point_precise_set
func LvPointPreciseSet(p *LvPointPreciseT, x LvValuePreciseT, y LvValuePreciseT)

//go:linkname LvPointSwap C.lv_point_swap
func LvPointSwap(p1 *LvPointT, p2 *LvPointT)

//go:linkname LvPointPreciseSwap C.lv_point_precise_swap
func LvPointPreciseSwap(p1 *LvPointPreciseT, p2 *LvPointPreciseT)

//go:linkname LvPct C.lv_pct
func LvPct(x c.Int32T) c.Int32T

//go:linkname LvPctToPx C.lv_pct_to_px
func LvPctToPx(v c.Int32T, base c.Int32T) c.Int32T
