package widgets

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvTableCreate C.lv_table_create
func LvTableCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvTableSetCellValue C.lv_table_set_cell_value
func LvTableSetCellValue(obj *core.LvObjT, row c.Uint32T, col c.Uint32T, txt *c.Char)

//go:linkname LvTableSetRowCount C.lv_table_set_row_count
func LvTableSetRowCount(obj *core.LvObjT, rowCnt c.Uint32T)

//go:linkname LvTableSetColumnCount C.lv_table_set_column_count
func LvTableSetColumnCount(obj *core.LvObjT, colCnt c.Uint32T)

//go:linkname LvTableSetColumnWidth C.lv_table_set_column_width
func LvTableSetColumnWidth(obj *core.LvObjT, colId c.Uint32T, w c.Int32T)

//go:linkname LvTableSetCellCtrl C.lv_table_set_cell_ctrl
func LvTableSetCellCtrl(obj *core.LvObjT, row c.Uint32T, col c.Uint32T, ctrl c.Int32T)

//go:linkname LvTableClearCellCtrl C.lv_table_clear_cell_ctrl
func LvTableClearCellCtrl(obj *core.LvObjT, row c.Uint32T, col c.Uint32T, ctrl c.Int32T)

//go:linkname LvTableSetCellUserData C.lv_table_set_cell_user_data
func LvTableSetCellUserData(obj *core.LvObjT, row c.Uint16T, col c.Uint16T, userData unsafe.Pointer)

//go:linkname LvTableSetSelectedCell C.lv_table_set_selected_cell
func LvTableSetSelectedCell(obj *core.LvObjT, row c.Uint16T, col c.Uint16T)

//go:linkname LvTableGetCellValue C.lv_table_get_cell_value
func LvTableGetCellValue(obj *core.LvObjT, row c.Uint32T, col c.Uint32T) *c.Char

//go:linkname LvTableGetRowCount C.lv_table_get_row_count
func LvTableGetRowCount(obj *core.LvObjT) c.Uint32T

//go:linkname LvTableGetColumnCount C.lv_table_get_column_count
func LvTableGetColumnCount(obj *core.LvObjT) c.Uint32T

//go:linkname LvTableGetColumnWidth C.lv_table_get_column_width
func LvTableGetColumnWidth(obj *core.LvObjT, col c.Uint32T) c.Int32T

//go:linkname LvTableHasCellCtrl C.lv_table_has_cell_ctrl
func LvTableHasCellCtrl(obj *core.LvObjT, row c.Uint32T, col c.Uint32T, ctrl c.Int32T) c.Char

//go:linkname LvTableGetSelectedCell C.lv_table_get_selected_cell
func LvTableGetSelectedCell(obj *core.LvObjT, row *c.Uint32T, col *c.Uint32T)

//go:linkname LvTableGetCellUserData C.lv_table_get_cell_user_data
func LvTableGetCellUserData(obj *core.LvObjT, row c.Uint16T, col c.Uint16T) unsafe.Pointer
