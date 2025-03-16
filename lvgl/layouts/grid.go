package layouts

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Constants for special grid values
const (
	LV_COORD_MAX = (1 << 31) - 1000
)

// Grid special values
func LV_GRID_FR(x c.Int) c.Int32T {
	return c.Int32T(LV_COORD_MAX - 100 + int(x))
}

const (
	LV_GRID_CONTENT       = LV_COORD_MAX - 101
	LV_GRID_TEMPLATE_LAST = LV_COORD_MAX
)

type lv_grid_align_t c.Int

const (
	LV_GRID_ALIGN_START lv_grid_align_t = iota
	LV_GRID_ALIGN_CENTER
	LV_GRID_ALIGN_END
	LV_GRID_ALIGN_STRETCH
	LV_GRID_ALIGN_SPACE_EVENLY
	LV_GRID_ALIGN_SPACE_AROUND
	LV_GRID_ALIGN_SPACE_BETWEEN
)

//go:linkname LvGridInit C.lv_grid_init
func LvGridInit()

//go:linkname LvObjSetGridDscArray C.lv_obj_set_grid_dsc_array
func LvObjSetGridDscArray(obj *core.LvObjT, col_dsc, row_dsc *c.Int32T)

//go:linkname LvObjSetGridAlign C.lv_obj_set_grid_align
func LvObjSetGridAlign(obj *core.LvObjT, column_align, row_align lv_grid_align_t)

//go:linkname LvObjSetGridCell C.lv_obj_set_grid_cell
func LvObjSetGridCell(obj *core.LvObjT, column_align lv_grid_align_t, col_pos, col_span c.Int32T,
	row_align lv_grid_align_t, row_pos, row_span c.Int32T)

//go:linkname LvGridFr C.lv_grid_fr
func LvGridFr(x c.Uint8T) c.Int32T
